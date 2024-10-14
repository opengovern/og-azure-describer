package describer

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"github.com/opengovern/og-azure-describer/azure/model"
	"golang.org/x/net/context"
	"strings"
)

func ApplicationInsights(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armapplicationinsights.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewComponentsClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, component := range page.Value {
			resource := GetApplicationInsights(ctx, component)
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

func GetApplicationInsights(ctx context.Context, component *armapplicationinsights.Component) *Resource {
	resourceGroup := strings.Split(*component.ID, "/")[4]

	resource := Resource{
		ID:       *component.ID,
		Name:     *component.Name,
		Location: *component.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.ApplicationInsightsComponentDescription{
				Component:     *component,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
