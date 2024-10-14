package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func PowerBIDedicatedCapacity(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armpowerbidedicated.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewCapacitiesClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetPowerBIDedicatedCapacity(ctx, v)
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

func GetPowerBIDedicatedCapacity(ctx context.Context, v *armpowerbidedicated.DedicatedCapacity) *Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.PowerBIDedicatedCapacityDescription{
				Capacity:      *v,
				ResourceGroup: resourceGroupName,
			},
		},
	}
	return &resource
}
