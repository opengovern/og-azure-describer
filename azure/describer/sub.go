package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func Location(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSubscriptionsClient()

	var values []Resource
	pager := client.NewListLocationsPager(subscription, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetLocation(ctx, v)
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

func GetLocation(ctx context.Context, location *armsubscription.Location) *Resource {
	resourceGroup := strings.Split(*location.ID, "/")[4]

	resource := Resource{
		ID:       *location.ID,
		Name:     *location.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.LocationDescription{
				Location:      *location,
				ResourceGroup: resourceGroup,
			},
		},
	}

	return &resource
}
