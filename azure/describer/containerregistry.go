package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func ContainerRegistry(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armcontainerregistry.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewRegistriesClient()
	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getContainerRegistry(ctx, client, v)
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

func getContainerRegistry(ctx context.Context, client *armcontainerregistry.RegistriesClient, registry *armcontainerregistry.Registry) (*Resource, error) {
	resourceGroup := strings.Split(*registry.ID, "/")[4]
	containerRegistryListCredentialsOp, err := client.ListCredentials(ctx, resourceGroup, *registry.Name, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "does not have authorization to perform action 'Microsoft.ContainerRegistry/registries/listCredentials/action'") {
			return nil, err
		}
	}

	containerRegistryListUsagesOp, err := client.ListUsages(ctx, resourceGroup, *registry.Name, nil)
	if err != nil {
		return nil, err
	}
	resource := Resource{
		ID:       *registry.ID,
		Name:     *registry.Name,
		Location: *registry.Location,
		Description: JSONAllFieldsMarshaller{
			model.ContainerRegistryDescription{
				Registry:                      *registry,
				RegistryListCredentialsResult: containerRegistryListCredentialsOp.RegistryListCredentialsResult,
				RegistryUsages:                containerRegistryListUsagesOp.Value,
				ResourceGroup:                 resourceGroup,
			},
		},
	}
	return &resource, nil
}
