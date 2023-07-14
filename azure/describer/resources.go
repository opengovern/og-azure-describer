package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
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
	client := resources.NewProvidersClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx, "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, provider := range result.Values() {
			resource := Resource{
				ID:       *provider.ID,
				Location: "global",
				Description: JSONAllFieldsMarshaller{
					model.ResourceProviderDescription{
						Provider: provider,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func ResourceGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client := resources.NewGroupsClient(subscription)
	client.Authorizer = authorizer

	groupListResultPage, err := client.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, group := range groupListResultPage.Values() {
			resource := Resource{
				ID:       *group.ID,
				Name:     *group.Name,
				Location: *group.Location,
				Description: JSONAllFieldsMarshaller{
					model.ResourceGroupDescription{
						Group: group,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		if !groupListResultPage.NotDone() {
			break
		}
		err = groupListResultPage.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}
