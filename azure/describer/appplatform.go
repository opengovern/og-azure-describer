package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func SpringCloudService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()
	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, resource := range result.Value {
			resource, err := getSpringCloudService(ctx, resource)
			if err != nil {
				return nil, err
			}
			if resource == nil {
				continue
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

func getSpringCloudService(ctx context.Context, service *armresources.GenericResourceExpanded) (*Resource, error) {
	if service.Name == nil {
		return nil, nil
	}
	splitID := strings.Split(*service.ID, "/")

	resourceGroup := splitID[4]

	resource := Resource{
		ID:       *service.ID,
		Name:     *service.Name,
		Location: *service.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.SpringCloudServiceDescription{
				ServiceResource:            *service,
				DiagnosticSettingsResource: nil, // TODO: Arta fix this =)))
				ResourceGroup:              resourceGroup,
			},
		},
	}
	return &resource, nil
}
