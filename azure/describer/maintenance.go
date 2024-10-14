package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"github.com/opengovern/og-azure-describer/azure/model"
)

func MaintenanceConfiguration(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {

	clientFactory, err := armmaintenance.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	configurationsClient := clientFactory.NewConfigurationsClient()

	pager := configurationsClient.NewListPager(nil)
	var resources []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, configuration := range page.Value {
			resource, err := getMaintenanceConfiguration(ctx, configuration)
			if err != nil {
				return nil, err
			}
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

func getMaintenanceConfiguration(ctx context.Context, configuration *armmaintenance.Configuration) (*Resource, error) {
	resourceGroup := strings.Split(*configuration.ID, "/")[4]

	resource := Resource{
		ID:   *configuration.ID,
		Name: *configuration.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.MaintenanceConfigurationDescription{
				MaintenanceConfiguration: *configuration,
				ResourceGroup:            resourceGroup,
			},
		},
	}
	return &resource, nil
}
