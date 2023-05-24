package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/kaytu-io/kaytu-azure-describer/azure"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer"
	azuremodel "github.com/kaytu-io/kaytu-azure-describer/azure/model"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/steampipe"
	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"github.com/kaytu-io/kaytu-util/pkg/source"
	"github.com/kaytu-io/kaytu-util/pkg/vault"
	"github.com/kaytu-io/kaytu-util/proto/src/golang"
	"go.uber.org/zap"
	"strings"
)

func fixAzureLocation(l string) string {
	return strings.ToLower(strings.ReplaceAll(l, " ", ""))
}

func trimEmptyMaps(input map[string]any) {
	for key, value := range input {
		switch value.(type) {
		case map[string]any:
			if len(value.(map[string]any)) != 0 {
				trimEmptyMaps(value.(map[string]any))
			}
			if len(value.(map[string]any)) == 0 {
				delete(input, key)
			}
		}
	}
}

func trimJsonFromEmptyObjects(input []byte) ([]byte, error) {
	unknownData := map[string]any{}
	err := json.Unmarshal(input, &unknownData)
	if err != nil {
		return nil, err
	}
	trimEmptyMaps(unknownData)
	return json.Marshal(unknownData)
}

func doDescribeAzure(
	ctx context.Context,
	logger *zap.Logger,
	job describe.DescribeJob,
	config map[string]any,
	workspaceName string,
	describeEndpoint string,
	describeToken string) ([]string, error) {
	rs, err := NewResourceSender(workspaceName, describeEndpoint, describeToken, job.JobID, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to resource sender: %w", err)
	}

	plg := steampipe.Plugin()
	plgAD := steampipe.ADPlugin()
	creds, err := azure.SubscriptionConfigFromMap(config)
	if err != nil {
		return nil, fmt.Errorf("azure subscription credentials: %w", err)
	}
	subscriptionId := job.AccountID
	if len(subscriptionId) == 0 {
		subscriptionId = creds.SubscriptionID
	}

	f := func(resource describer.Resource) error {
		if resource.Description == nil {
			return nil
		}
		descriptionJSON, err := json.Marshal(resource.Description)
		if err != nil {
			return fmt.Errorf("failed to marshal description: %w", err)
		}
		descriptionJSON, err = trimJsonFromEmptyObjects(descriptionJSON)
		if err != nil {
			return fmt.Errorf("failed to trim json: %w", err)
		}
		resource.Location = fixAzureLocation(resource.Location)

		azureMetadata := azuremodel.Metadata{
			ID:               resource.ID,
			Name:             resource.Name,
			SubscriptionID:   job.AccountID,
			Location:         resource.Location,
			CloudEnvironment: "AzurePublicCloud",
			ResourceType:     strings.ToLower(job.ResourceType),
			SourceID:         job.SourceID,
		}
		azureMetadataBytes, err := json.Marshal(azureMetadata)
		if err != nil {
			return fmt.Errorf("marshal metadata: %v", err.Error())
		}

		metadata := make(map[string]string)
		err = json.Unmarshal(azureMetadataBytes, &metadata)
		if err != nil {
			return fmt.Errorf("unmarshal metadata: %v", err.Error())
		}

		kafkaResource := Resource{
			ID:            resource.UniqueID(),
			Name:          resource.Name,
			ResourceGroup: resource.ResourceGroup,
			Location:      resource.Location,
			SourceType:    source.CloudAzure,
			ResourceType:  strings.ToLower(job.ResourceType),
			ResourceJobID: job.JobID,
			SourceJobID:   job.ParentJobID,
			SourceID:      job.SourceID,
			ScheduleJobID: job.ScheduleJobID,
			CreatedAt:     job.DescribedAt,
			Description:   resource.Description,
			Metadata:      metadata,
		}

		tags, name, err := steampipe.ExtractTagsAndNames(plg, plgAD, job.ResourceType, kafkaResource)
		if err != nil {
			return fmt.Errorf("failed to build tags for service: %v", err.Error())
		}
		if len(name) > 0 {
			kafkaResource.Metadata["name"] = name
		}

		rs.Send(&golang.AzureResource{
			UniqueId:        resource.UniqueID(),
			Id:              resource.ID,
			Name:            resource.Name,
			Type:            resource.Type,
			ResourceGroup:   resource.ResourceGroup,
			Location:        resource.Location,
			SubscriptionId:  resource.SubscriptionID,
			DescriptionJson: string(descriptionJSON),
			Metadata:        metadata,
			Tags:            tags,
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
		return nil
	}
	clientStream := (*describer.StreamSender)(&f)

	_, err = azure.GetResources(
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
		return nil, err
	}

	rs.Finish()

	return rs.GetResourceIDs(), nil
}

func Do(ctx context.Context,
	vlt *vault.KMSVaultSourceConfig,
	logger *zap.Logger,
	job describe.DescribeJob,
	keyARN string,
	describeDeliverEndpoint string,
	describeDeliverToken string,
	workspaceName string) (resourceIDs []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("paniced with error: %v", r)
			logger.Error("paniced with error", zap.Error(err), zap.String("stackTrace", errors.Wrap(r, 2).ErrorStack()))
		}
	}()

	if job.SourceType != source.CloudAzure {
		return nil, fmt.Errorf("unsupported source type %s", job.SourceType)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	config, err := vlt.Decrypt(job.CipherText, keyARN)
	if err != nil {
		return nil, fmt.Errorf("decrypt error: %w", err)
	}

	return doDescribeAzure(ctx, logger, job, config, workspaceName, describeDeliverEndpoint, describeDeliverToken)
}
