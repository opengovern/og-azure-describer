package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func ContainerInstanceContainerGroups(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armcontainerinstance.NewContainerGroupsClient(subscription, cred, nil)
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
			resource := getContainerInstanceContainerGrou(ctx, v)
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

func getContainerInstanceContainerGrou(ctx context.Context, v *armcontainerinstance.ContainerGroup) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.ContainerInstanceContainerGroupDescription{
				ContainerGroup: *v,
				ResourceGroup:  resourceGroup,
			},
		},
	}

	return &resource
}
