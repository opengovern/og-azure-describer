package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks"
	"github.com/opengovern/og-azure-describer/azure/model"
)

func ResourceLink(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armlinks.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewResourceLinksClient()

	pager := client.NewListAtSubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getResourceLink(ctx, v)
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

func getResourceLink(ctx context.Context, v *armlinks.ResourceLink) *Resource {
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.ResourceLinkDescription{
				ResourceLink: *v,
			},
		},
	}
	return &resource
}
