package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2022-10-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2016-03-01/streamanalytics"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func StreamAnalyticsJob(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := insights.NewDiagnosticSettingsClient(subscription)
	client.Authorizer = authorizer

	streamingJobsClient := streamanalytics.NewStreamingJobsClient(subscription)
	streamingJobsClient.Authorizer = authorizer

	result, err := streamingJobsClient.List(context.Background(), "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, streamingJob := range result.Values() {
			resourceGroup := strings.Split(*streamingJob.ID, "/")[4]

			streamanalyticsListOp, err := client.List(ctx, *streamingJob.ID)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:       *streamingJob.ID,
				Name:     *streamingJob.Name,
				Location: *streamingJob.Location,
				Description: JSONAllFieldsMarshaller{
					model.StreamAnalyticsJobDescription{
						StreamingJob:                streamingJob,
						DiagnosticSettingsResources: streamanalyticsListOp.Value,
						ResourceGroup:               resourceGroup,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
