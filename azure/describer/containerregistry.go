package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func ContainerRegistry(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armcontainerregistry.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewRegistriesClient()
	webhookClient := clientFactory.NewWebhooksClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getContainerRegistry(ctx, client, webhookClient, v)
			if err != nil {
				return nil, err
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
	return values, nil
}

func getContainerRegistry(ctx context.Context, client *armcontainerregistry.RegistriesClient, webhookClient *armcontainerregistry.WebhooksClient, registry *armcontainerregistry.Registry) (*Resource, error) {
	resourceGroup := strings.Split(*registry.ID, "/")[4]
	var containerRegistryListCredentialsOp *armcontainerregistry.RegistryListCredentialsResult
	containerRegistryListCredentialsOpTemp, err := client.ListCredentials(ctx, resourceGroup, *registry.Name, nil)
	if err != nil {
		return nil, err
	} else {
		containerRegistryListCredentialsOp = &containerRegistryListCredentialsOpTemp.RegistryListCredentialsResult
	}
	containerRegistryListUsagesOp, err := client.ListUsages(ctx, resourceGroup, *registry.Name, nil)
	if err != nil {
		return nil, err
	}

	webhooksPager := webhookClient.NewListPager(resourceGroup, *registry.Name, nil)

	var webhooks []*armcontainerregistry.Webhook
	for webhooksPager.More() {
		page, err := webhooksPager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		webhooks = append(webhooks, page.Value...)
	}

	resource := Resource{
		ID:       *registry.ID,
		Name:     *registry.Name,
		Location: *registry.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.ContainerRegistryDescription{
				Registry:                      *registry,
				RegistryListCredentialsResult: containerRegistryListCredentialsOp,
				RegistryUsages:                containerRegistryListUsagesOp.Value,
				Webhooks:                      webhooks,
				ResourceGroup:                 resourceGroup,
			},
		},
	}
	return &resource, nil
}
