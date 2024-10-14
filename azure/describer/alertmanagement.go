package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
	"github.com/opengovern/og-azure-describer/azure/model"
)

func AlertManagement(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {

	clientFactory, err := armalertsmanagement.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewAlertsClient()
	pager := client.NewGetAllPager(nil)

	var resources []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getAlertManagement(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				resources = append(resources, *resource)
			}
		}
	}
	return resources, nil
}

func getAlertManagement(_ context.Context, alert *armalertsmanagement.Alert) *Resource {

	resourceGroup := strings.Split(*alert.ID, "/")[4]
	return &Resource{
		ID:   *alert.ID,
		Name: *alert.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.AlertManagementDescription{
				Alert:         *alert,
				ResourceGroup: resourceGroup,
			},
		},
	}
}
