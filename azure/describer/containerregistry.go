package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2020-11-01-preview/containerregistry"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func ContainerRegistry(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	containerRegistryClient := containerregistry.NewRegistriesClient(subscription)
	containerRegistryClient.Authorizer = authorizer

	client := containerregistry.NewRegistriesClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, registry := range result.Values() {
			resourceGroup := strings.Split(*registry.ID, "/")[4]

			containerRegistryListCredentialsOp, err := containerRegistryClient.ListCredentials(ctx, resourceGroup, *registry.Name)
			if err != nil {
				if !strings.Contains(err.Error(), "UnAuthorizedForCredentialOperations") {
					return nil, err
				}
			}

			containerRegistryListUsagesOp, err := containerRegistryClient.ListUsages(ctx, resourceGroup, *registry.Name)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:       *registry.ID,
				Name:     *registry.Name,
				Location: *registry.Location,
				Description: model.ContainerRegistryDescription{
					Registry:                      registry,
					RegistryListCredentialsResult: containerRegistryListCredentialsOp,
					RegistryUsages:                containerRegistryListUsagesOp.Value,
					ResourceGroup:                 resourceGroup,
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
