package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func listResourceGroups(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string) ([]armresources.ResourceGroup, error) {
	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewResourceGroupsClient()
	pager := client.NewListPager(nil)
	var values []armresources.ResourceGroup
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			values = append(values, *v)
		}
	}
	return values, nil
}

func ResourceProvider(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewProvidersClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, provider := range page.Value {
			resource := GetResourceProvider(ctx, provider)
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

func GetResourceProvider(ctx context.Context, provider *armresources.Provider) *Resource {
	resource := Resource{
		ID:       *provider.ID,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			model.ResourceProviderDescription{
				Provider: *provider,
			},
		},
	}

	return &resource
}

func ResourceGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewResourceGroupsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, group := range page.Value {
			resource := GetResourceGroup(ctx, group)
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

func GetResourceGroup(ctx context.Context, group *armresources.ResourceGroup) *Resource {
	resource := Resource{
		ID:       *group.ID,
		Name:     *group.Name,
		Location: *group.Location,
		Description: JSONAllFieldsMarshaller{
			model.ResourceGroupDescription{
				Group: *group,
			},
		},
	}

	return &resource
}
