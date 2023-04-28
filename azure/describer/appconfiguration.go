package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/appconfiguration/mgmt/2020-06-01/appconfiguration"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-04-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func AppConfiguration(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	configurationStoresClient := appconfiguration.NewConfigurationStoresClient(subscription)
	configurationStoresClient.Authorizer = authorizer

	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
	insightsClient.Authorizer = authorizer

	result, err := configurationStoresClient.List(ctx, "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, config := range result.Values() {
			resourceGroup := strings.Split(*config.ID, "/")[4]

			op, err := insightsClient.List(ctx, *config.ID)
			if err != nil {
				return nil, err
			}
			resource := Resource{
				ID:       *config.ID,
				Name:     *config.Name,
				Location: *config.Location,
				Description: model.AppConfigurationDescription{
					ConfigurationStore:          config,
					DiagnosticSettingsResources: *op.Value,
					ResourceGroup:               resourceGroup,
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
