package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func SpringCloudService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	var values []Resource

	clientFactory, err := armspringappdiscovery.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	siteClient := clientFactory.NewSpringbootsitesClient()
	client := clientFactory.NewSpringbootappsClient()
	pager := siteClient.NewListBySubscriptionPager(&armspringappdiscovery.SpringbootsitesClientListBySubscriptionOptions{})
	for pager.More() {
		sites, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, site := range sites.Value {
			appPager := client.NewListBySubscriptionPager(*site.Name, &armspringappdiscovery.SpringbootappsClientListBySubscriptionOptions{})
			for appPager.More() {
				apps, err := appPager.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, app := range apps.Value {
					resource, err := getSpringCloudService(ctx, app, site)
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
		}
	}

	return values, nil
}

func getSpringCloudService(ctx context.Context, service *armspringappdiscovery.SpringbootappsModel, site *armspringappdiscovery.SpringbootsitesModel) (*Resource, error) {
	if service.Name == nil {
		return nil, nil
	}
	splitID := strings.Split(*service.ID, "/")

	resourceGroup := splitID[4]

	resource := Resource{
		ID:       *service.ID,
		Name:     *service.Name,
		Location: *site.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.SpringCloudServiceDescription{
				App:                        *service,
				Site:                       site,
				DiagnosticSettingsResource: nil, // TODO: Arta fix this =)))
				ResourceGroup:              resourceGroup,
			},
		},
	}
	return &resource, nil
}
