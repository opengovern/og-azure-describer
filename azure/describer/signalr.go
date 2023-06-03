package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2022-10-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/signalr/mgmt/2020-05-01/signalr"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func SignalrService(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
	insightsClient.Authorizer = authorizer

	client := signalr.NewClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListBySubscription(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, service := range result.Values() {
			resourceGroup := strings.Split(*service.ID, "/")[4]

			signalrListOp, err := insightsClient.List(ctx, *service.ID)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:       *service.ID,
				Name:     *service.Name,
				Location: *service.Location,
				Description: JSONAllFieldsMarshaller{
					model.SignalrServiceDescription{
						ResourceType:                service,
						DiagnosticSettingsResources: signalrListOp.Value,
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
