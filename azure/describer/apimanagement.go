package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2020-12-01/apimanagement"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2022-10-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func APIManagement(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	apiManagementClient := apimanagement.NewServiceClient(subscription)
	apiManagementClient.Authorizer = authorizer

	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
	insightsClient.Authorizer = authorizer

	result, err := apiManagementClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, apiManagement := range result.Values() {
			resourceGroup := strings.Split(*apiManagement.ID, "/")[4]

			op, err := insightsClient.List(ctx, *apiManagement.ID)
			if err != nil {
				return nil, err
			}
			resource := Resource{
				ID:       *apiManagement.ID,
				Name:     *apiManagement.Name,
				Location: *apiManagement.Location,
				Description: JSONAllFieldsMarshaller{
					model.APIManagementDescription{
						APIManagement:               apiManagement,
						DiagnosticSettingsResources: *op.Value,
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
