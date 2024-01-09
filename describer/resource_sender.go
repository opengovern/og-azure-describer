package describer

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/kaytu-io/kaytu-util/pkg/es"
	es2 "github.com/kaytu-io/kaytu-util/pkg/es"
	"github.com/kaytu-io/kaytu-util/pkg/source"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/kaytu-io/kaytu-util/proto/src/golang"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
)

const (
	MinBufferSize   int           = 10
	MaxBufferSize   int           = 100
	ChannelSize     int           = 1000
	BufferEmptyRate time.Duration = 5 * time.Second
)

type ResourceSender struct {
	authToken                 string
	workspaceId               string
	workspaceName             string
	logger                    *zap.Logger
	resourceChannel           chan *golang.AzureResource
	resourceIDs               []string
	doneChannel               chan interface{}
	conn                      *grpc.ClientConn
	describeEndpoint          string
	ingestionPipelineEndpoint string
	jobID                     uint
	kafkaTopic                string

	client     golang.DescribeServiceClient
	httpClient *http.Client

	sendBuffer    []*golang.AzureResource
	useOpenSearch bool
}

func NewResourceSender(workspaceId string, workspaceName string, describeEndpoint, ingestionPipelineEndpoint string, describeToken string, jobID uint, kafkaTopic string, useOpenSearch bool, logger *zap.Logger) (*ResourceSender, error) {
	rs := ResourceSender{
		authToken:                 describeToken,
		workspaceId:               workspaceId,
		workspaceName:             workspaceName,
		logger:                    logger,
		resourceChannel:           make(chan *golang.AzureResource, ChannelSize),
		resourceIDs:               nil,
		doneChannel:               make(chan interface{}),
		conn:                      nil,
		describeEndpoint:          describeEndpoint,
		ingestionPipelineEndpoint: ingestionPipelineEndpoint,
		kafkaTopic:                kafkaTopic,
		jobID:                     jobID,
		useOpenSearch:             useOpenSearch,

		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
	if err := rs.Connect(); err != nil {
		return nil, err
	}

	go rs.ResourceHandler()
	return &rs, nil
}

func (s *ResourceSender) Connect() error {
	conn, err := grpc.Dial(
		s.describeEndpoint,
		grpc.WithTransportCredentials(credentials.NewTLS(nil)),
		grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: s.authToken,
			}),
		}),
	)
	if err != nil {
		return err
	}
	s.conn = conn

	client := golang.NewDescribeServiceClient(conn)
	s.client = client
	return nil
}

func (s *ResourceSender) ResourceHandler() {
	t := time.NewTicker(BufferEmptyRate)
	defer t.Stop()

	for {
		select {
		case resource := <-s.resourceChannel:
			if resource == nil {
				s.flushBuffer(true)
				s.doneChannel <- struct{}{}
				return
			}

			s.resourceIDs = append(s.resourceIDs, resource.UniqueId)
			s.sendBuffer = append(s.sendBuffer, resource)

			if len(s.sendBuffer) > MaxBufferSize {
				s.flushBuffer(true)
			}
		case <-t.C:
			s.flushBuffer(false)
		}
	}
}

func (s *ResourceSender) sendToBackend() {
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"workspace-name":  s.workspaceName,
		"resource-job-id": fmt.Sprintf("%d", s.jobID),
	}))

	_, err := s.client.DeliverAzureResources(grpcCtx, &golang.AzureResources{Resources: s.sendBuffer, KafkaTopic: s.kafkaTopic})
	if err != nil {
		s.logger.Error("failed to send resource", zap.Error(err))
		if errors.Is(err, io.EOF) {
			err = s.Connect()
			if err != nil {
				s.logger.Error("failed to reconnect", zap.Error(err))
			}
		}
		return
	}
}

func (s *ResourceSender) sendToOpenSearchIngestPipeline() {
	resourcesToSend := make([]es2.Doc, 0, 2*len(s.sendBuffer))
	for _, resource := range s.sendBuffer {
		var description any
		err := json.Unmarshal([]byte(resource.DescriptionJson), &description)
		if err != nil {
			s.logger.Error("failed to parse resource description json", zap.Error(err), zap.Uint32("jobID", resource.Job.JobId), zap.String("resourceID", resource.Id))
			continue
		}

		tags := make([]es.Tag, 0, len(resource.Tags))
		for k, v := range resource.Tags {
			tags = append(tags, es.Tag{
				// tags should be case-insensitive
				Key:   strings.ToLower(k),
				Value: strings.ToLower(v),
			})
		}

		kafkaResource := es.Resource{
			ID:            resource.UniqueId,
			ARN:           "",
			Description:   description,
			SourceType:    source.CloudAzure,
			ResourceType:  strings.ToLower(resource.Job.ResourceType),
			ResourceJobID: uint(resource.Job.JobId),
			SourceID:      resource.Job.SourceId,
			SourceJobID:   uint(resource.Job.ParentJobId),
			Metadata:      resource.Metadata,
			Name:          resource.Name,
			ResourceGroup: resource.ResourceGroup,
			Location:      resource.Location,
			ScheduleJobID: uint(resource.Job.ScheduleJobId),
			CreatedAt:     resource.Job.DescribedAt,
			CanonicalTags: tags,
		}
		keys, idx := kafkaResource.KeysAndIndex()
		kafkaResource.EsID = es2.HashOf(keys...)
		kafkaResource.EsIndex = idx

		lookupResource := es.LookupResource{
			ResourceID:    resource.UniqueId,
			Name:          resource.Name,
			SourceType:    source.CloudAzure,
			ResourceType:  strings.ToLower(resource.Job.ResourceType),
			ResourceGroup: resource.ResourceGroup,
			Location:      resource.Location,
			SourceID:      resource.Job.SourceId,
			ResourceJobID: uint(resource.Job.JobId),
			SourceJobID:   uint(resource.Job.ParentJobId),
			ScheduleJobID: uint(resource.Job.ScheduleJobId),
			CreatedAt:     resource.Job.DescribedAt,
			Tags:          tags,
		}
		lookupKeys, lookupIdx := lookupResource.KeysAndIndex()
		lookupResource.EsID = es2.HashOf(lookupKeys...)
		lookupResource.EsIndex = lookupIdx

		resourcesToSend = append(resourcesToSend, kafkaResource)
		resourcesToSend = append(resourcesToSend, lookupResource)
	}

	if len(resourcesToSend) == 0 {
		return
	}

	jsonResourcesToSend, err := json.Marshal(resourcesToSend)
	if err != nil {
		s.logger.Error("failed to marshal resources", zap.Error(err))
		return
	}

	req, err := http.NewRequest(
		http.MethodPost,
		s.ingestionPipelineEndpoint,
		strings.NewReader(string(jsonResourcesToSend)),
	)
	req.Header.Add("Content-Type", "application/json")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		s.logger.Error("failed to load configuration", zap.Error(err))
		return
	}
	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		s.logger.Error("failed to retrieve credentials", zap.Error(err))
		return
	}

	signer := v4.NewSigner()
	err = signer.SignHTTP(context.TODO(), creds, req,
		fmt.Sprintf("%x", sha256.Sum256(jsonResourcesToSend)),
		"osis", "us-east-2", time.Now())
	if err != nil {
		s.logger.Error("failed to sign request", zap.Error(err))
		return
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		s.logger.Error("failed to send request", zap.Error(err))
		return
	}
	defer resp.Body.Close()
	// check status
	if resp.StatusCode != http.StatusOK {
		bodyStr := ""
		if resp.Body != nil {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				s.logger.Warn("failed to read response body", zap.Error(err))
			} else {
				bodyStr = string(bodyBytes)
			}
		}
		s.logger.Error("failed to send resources to OpenSearch",
			zap.Int("statusCode", resp.StatusCode),
			zap.String("body", bodyStr),
		)
		return
	}
}

func (s *ResourceSender) flushBuffer(force bool) {
	if len(s.sendBuffer) == 0 {
		return
	}

	if !force && len(s.sendBuffer) < MinBufferSize {
		return
	}

	if s.useOpenSearch {
		s.sendToOpenSearchIngestPipeline()
	} else {
		s.sendToBackend()
	}

	s.sendBuffer = nil
}

func (s *ResourceSender) Finish() {
	s.resourceChannel <- nil
	_ = <-s.doneChannel
	s.conn.Close()
}

func (s *ResourceSender) GetResourceIDs() []string {
	return s.resourceIDs
}

func (s *ResourceSender) Send(resource *golang.AzureResource) {
	s.resourceChannel <- resource
}
