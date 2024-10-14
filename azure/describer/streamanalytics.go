package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func StreamAnalyticsJob(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armstreamanalytics.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	streamingJobsClient := clientFactory.NewStreamingJobsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	var values []Resource
	pager := streamingJobsClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetStreamAnalyticsJob(ctx, diagnosticClient, v)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetStreamAnalyticsJob(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, streamingJob *armstreamanalytics.StreamingJob) (*Resource, error) {
	resourceGroup := strings.Split(*streamingJob.ID, "/")[4]

	pager := diagnosticClient.NewListPager(*streamingJob.ID, nil)
	var streamanalyticsListOp []*armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		streamanalyticsListOp = append(streamanalyticsListOp, page.Value...)
	}

	resource := Resource{
		ID:       *streamingJob.ID,
		Name:     *streamingJob.Name,
		Location: *streamingJob.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.StreamAnalyticsJobDescription{
				StreamingJob:                *streamingJob,
				DiagnosticSettingsResources: streamanalyticsListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}

func StreamAnalyticsCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armstreamanalytics.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	streamingJobsClient := clientFactory.NewClustersClient()

	var values []Resource
	pager := streamingJobsClient.NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetStreamAnalyticsCluster(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetStreamAnalyticsCluster(ctx context.Context, streamingJob *armstreamanalytics.Cluster) *Resource {
	resourceGroup := strings.Split(*streamingJob.ID, "/")[4]

	resource := Resource{
		ID:       *streamingJob.ID,
		Name:     *streamingJob.Name,
		Location: *streamingJob.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.StreamAnalyticsClusterDescription{
				StreamingJob:  *streamingJob,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
