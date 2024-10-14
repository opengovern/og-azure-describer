package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/opengovern/og-azure-describer/azure/model"
)

func Tenant(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewTenantsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetTenand(ctx, v)
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

func GetTenand(ctx context.Context, v *armsubscription.TenantIDDescription) *Resource {
	name := ""

	resource := Resource{
		ID:       *v.ID,
		Name:     name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.TenantDescription{
				TenantIDDescription: *v, // TODO has much less values
			},
		},
	}

	return &resource
}

func Subscription(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsubscription.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
	resourceClientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewSubscriptionsClient()
	tagsClient := resourceClientFactory.NewTagsClient()

	op, err := client.Get(ctx, subscription, nil)
	if err != nil {
		return nil, err
	}

	tags := map[string][]string{}
	pager := tagsClient.NewListPager(&armresources.TagsClientListOptions{})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, tag := range page.Value {
			if tag.TagName == nil {
				continue
			}

			var values []string
			for _, v := range tag.Values {
				if v.TagValue != nil {
					values = append(values, *v.TagValue)
				}
			}

			tags[*tag.TagName] = values
		}
	}

	var values []Resource
	resource := Resource{
		ID:       *op.ID,
		Name:     *op.DisplayName,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.SubscriptionDescription{
				Subscription: op.Subscription,
				Tags:         tags,
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

	return values, nil
}
