package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func DesktopVirtualizationHostPool(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armdesktopvirtualization.NewHostPoolsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getDesktopVirtualizationHostPool(ctx, v)
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

func getDesktopVirtualizationHostPool(ctx context.Context, v *armdesktopvirtualization.HostPool) *Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.DesktopVirtualizationHostPoolDescription{
				HostPool:      *v,
				ResourceGroup: resourceGroupName,
			},
		},
	}
	return &resource
}
