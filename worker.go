package kaytu_azure_describer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-errors/errors"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/kaytu-io/kaytu-azure-describer/azure"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	"gitlab.com/keibiengine/keibi-engine/pkg/cloudservice"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe/api"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe/es"
	"gitlab.com/keibiengine/keibi-engine/pkg/describe/proto/src/golang"
	"gitlab.com/keibiengine/keibi-engine/pkg/kafka"
	"gitlab.com/keibiengine/keibi-engine/pkg/source"
	"gitlab.com/keibiengine/keibi-engine/pkg/steampipe"
	"gitlab.com/keibiengine/keibi-engine/pkg/vault"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func fixAzureLocation(l string) string {
	return strings.ToLower(strings.ReplaceAll(l, " ", ""))
}

func doDescribeAzure(ctx context.Context, job describe.DescribeJob, config map[string]interface{},
	logger *zap.Logger, client *golang.DescribeServiceClient) ([]kafka.Doc, []string, error) {
	var clientStream *describer.StreamSender
	if client != nil {
		stream, err := (*client).DeliverAzureResources(context.Background())
		if err != nil {
			return nil, nil, err
		}

		f := func(resource describer.Resource) error {
			descriptionJSON, err := json.Marshal(resource.Description)
			if err != nil {
				return err
			}

			return stream.Send(&golang.AzureResource{
				Id:              resource.ID,
				Name:            resource.Name,
				Type:            resource.Type,
				ResourceGroup:   resource.ResourceGroup,
				Location:        resource.Location,
				SubscriptionId:  resource.SubscriptionID,
				DescriptionJson: string(descriptionJSON),
				Job: &golang.DescribeJob{
					JobId:         uint32(job.JobID),
					ScheduleJobId: uint32(job.ScheduleJobID),
					ParentJobId:   uint32(job.ParentJobID),
					ResourceType:  job.ResourceType,
					SourceId:      job.SourceID,
					AccountId:     job.AccountID,
					DescribedAt:   job.DescribedAt,
					SourceType:    string(job.SourceType),
					ConfigReg:     job.CipherText,
					TriggerType:   string(job.TriggerType),
					RetryCounter:  uint32(job.RetryCounter),
				},
			})
		}
		clientStream = (*describer.StreamSender)(&f)
	}

	var resourceIDs []string

	logger.Warn("starting to describe azure subscription", zap.String("resourceType", job.ResourceType), zap.Uint("jobID", job.JobID))
	creds, err := azure.SubscriptionConfigFromMap(config)
	if err != nil {
		return nil, nil, fmt.Errorf("azure subscription credentials: %w", err)
	}

	subscriptionId := job.AccountID
	if len(subscriptionId) == 0 {
		subscriptionId = creds.SubscriptionID
	}

	logger.Warn("getting resources", zap.String("resourceType", job.ResourceType), zap.Uint("jobID", job.JobID))
	output, err := azure.GetResources(
		ctx,
		job.ResourceType,
		job.TriggerType,
		[]string{subscriptionId},
		azure.AuthConfig{
			TenantID:            creds.TenantID,
			ClientID:            creds.ClientID,
			ClientSecret:        creds.ClientSecret,
			CertificatePath:     creds.CertificatePath,
			CertificatePassword: creds.CertificatePass,
			Username:            creds.Username,
			Password:            creds.Password,
		},
		string(azure.AuthEnv),
		"",
		clientStream,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("azure: %w", err)
	}
	logger.Warn("got the resources, finding summaries", zap.String("resourceType", job.ResourceType), zap.Uint("jobID", job.JobID))

	var msgs []kafka.Doc
	var errs []string

	for idx, resource := range output.Resources {
		if resource.Description == nil {
			continue
		}

		output.Resources[idx].Location = fixAzureLocation(resource.Location)

		azureMetadata := model.Metadata{
			ID:               resource.ID,
			Name:             resource.Name,
			SubscriptionID:   strings.Join(output.Metadata.SubscriptionIds, ","),
			Location:         resource.Location,
			CloudEnvironment: output.Metadata.CloudEnvironment,
			ResourceType:     strings.ToLower(resource.Type),
			SourceID:         job.SourceID,
		}
		azureMetadataBytes, err := json.Marshal(azureMetadata)
		if err != nil {
			errs = append(errs, fmt.Sprintf("marshal metadata: %v", err.Error()))
			continue
		}
		metadata := make(map[string]string)
		err = json.Unmarshal(azureMetadataBytes, &metadata)
		if err != nil {
			errs = append(errs, fmt.Sprintf("unmarshal metadata: %v", err.Error()))
			continue
		}

		kafkaResource := es.Resource{
			ID:            resource.UniqueID(),
			Name:          resource.Name,
			ResourceGroup: resource.ResourceGroup,
			Location:      resource.Location,
			SourceType:    source.CloudAzure,
			ResourceType:  strings.ToLower(output.Metadata.ResourceType),
			ResourceJobID: job.JobID,
			SourceJobID:   job.ParentJobID,
			SourceID:      job.SourceID,
			ScheduleJobID: job.ScheduleJobID,
			CreatedAt:     job.DescribedAt,
			Description:   resource.Description,
			Metadata:      metadata,
		}
		lookupResource := es.LookupResource{
			ResourceID:    resource.UniqueID(),
			Name:          resource.Name,
			SourceType:    source.CloudAzure,
			ResourceType:  strings.ToLower(job.ResourceType),
			ResourceGroup: resource.ResourceGroup,
			ServiceName:   cloudservice.ServiceNameByResourceType(job.ResourceType),
			Category:      cloudservice.CategoryByResourceType(job.ResourceType),
			Location:      resource.Location,
			SourceID:      job.SourceID,
			ScheduleJobID: job.ScheduleJobID,
			ResourceJobID: job.JobID,
			SourceJobID:   job.ParentJobID,
			CreatedAt:     job.DescribedAt,
			IsCommon:      cloudservice.IsCommonByResourceType(job.ResourceType),
		}
		resourceIDs = append(resourceIDs, resource.UniqueID())
		pluginTableName := steampipe.ExtractTableName(job.ResourceType)
		desc, err := steampipe.ConvertToDescription(job.ResourceType, kafkaResource)
		if err != nil {
			errs = append(errs, fmt.Sprintf("convertToDescription: %v", err.Error()))
			continue
		}
		pluginProvider := steampipe.ExtractPlugin(job.ResourceType)
		var cells map[string]*proto.Column
		if pluginProvider == steampipe.SteampipePluginAzure {
			cells, err = steampipe.AzureDescriptionToRecord(desc, pluginTableName)
			if err != nil {
				errs = append(errs, fmt.Sprintf("azureDescriptionToRecord: %v", err.Error()))
				continue
			}
		} else {
			cells, err = steampipe.AzureADDescriptionToRecord(desc, pluginTableName)
			if err != nil {
				errs = append(errs, fmt.Sprintf("azureADDescriptionToRecord: %v", err.Error()))
				continue
			}
		}
		for name, v := range cells {
			if name == "title" || name == "name" {
				kafkaResource.Metadata["name"] = v.GetStringValue()
			}
		}

		tags, err := steampipe.ExtractTags(job.ResourceType, kafkaResource)
		if err != nil {
			tags = map[string]string{}
			errs = append(errs, fmt.Sprintf("failed to build tags for service: %v", err.Error()))
		}
		lookupResource.Tags = tags

		msgs = append(msgs, kafkaResource)
		msgs = append(msgs, lookupResource)
	}
	logger.Warn("finished describing azure", zap.String("resourceType", job.ResourceType), zap.Uint("jobID", job.JobID))

	if len(errs) > 0 {
		err = fmt.Errorf("azure: [%s]", strings.Join(errs, ","))
	} else {
		err = nil
	}
	return msgs, resourceIDs, err
}

func Do(ctx context.Context,
	vlt *vault.KMSVaultSourceConfig,
	job describe.DescribeJob,
	keyARN string,
	logger *zap.Logger,
	describeDeliverEndpoint *string) error {
	logger.Info("Starting DescribeJob",
		zap.Uint("jobID", job.JobID),
		zap.Uint("parentJobID", job.ParentJobID),
		zap.String("resourceType", job.ResourceType),
		zap.String("sourceID", job.SourceID),
		zap.String("accountID", job.AccountID),
		zap.Int64("describedAt", job.DescribedAt),
		zap.String("sourceType", string(job.SourceType)),
		zap.String("cipherText", job.CipherText),
		zap.String("triggerType", string(job.TriggerType)),
		zap.Uint("retryCounter", job.RetryCounter))

	if job.SourceType != source.CloudAzure {
		return fmt.Errorf("unsupported source type %s", job.SourceType)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("paniced with error:", err)
			fmt.Println(errors.Wrap(err, 2).ErrorStack())
		}
	}()

	// Assume it succeeded unless it fails somewhere
	var (
		status               = api.DescribeResourceJobSucceeded
		firstErr    error    = nil
		resourceIDs []string = nil
	)

	fail := func(err error) {
		status = api.DescribeResourceJobFailed
		if firstErr == nil {
			firstErr = err
		}
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if conn, err := grpc.Dial(*describeDeliverEndpoint); err == nil {
		defer conn.Close()
		client := golang.NewDescribeServiceClient(conn)

		if config, err := vlt.Decrypt(job.CipherText, keyARN); err == nil {
			_, resourceIDs, err = doDescribeAzure(ctx, job, config, logger, &client)
			if err != nil {
				fail(fmt.Errorf("describe resources: %w", err))
			}
		} else if config == nil {
			fail(fmt.Errorf("config is null! path is: %s", job.CipherText))
		} else {
			fail(fmt.Errorf("resource source config: %w", err))
		}

		errMsg := ""
		if firstErr != nil {
			errMsg = firstErr.Error()
		}

		_, err := client.DeliverResult(ctx, &golang.DeliverResultRequest{
			JobId:       uint32(job.JobID),
			ParentJobId: uint32(job.ParentJobID),
			Status:      string(status),
			Error:       errMsg,
			DescribeJob: &golang.DescribeJob{
				JobId:         uint32(job.JobID),
				ScheduleJobId: uint32(job.ScheduleJobID),
				ParentJobId:   uint32(job.ParentJobID),
				ResourceType:  job.ResourceType,
				SourceId:      job.SourceID,
				AccountId:     job.AccountID,
				DescribedAt:   job.DescribedAt,
				SourceType:    string(job.SourceType),
				ConfigReg:     job.CipherText,
				TriggerType:   string(job.TriggerType),
				RetryCounter:  uint32(job.RetryCounter),
			},
			DescribedResourceIds: resourceIDs,
		})
		if err != nil {
			return fmt.Errorf("DeliverResult: %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("grpc: %v", err)
	}
}
