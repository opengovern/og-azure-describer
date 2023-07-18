package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func RecoveryServicesVault(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armrecoveryservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewVaultsClient()

	var values []Resource
	pager := client.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resource := GetRecoveryServicesVault(ctx, vault)
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

func GetRecoveryServicesVault(ctx context.Context, vault *armrecoveryservices.Vault) *Resource {
	resourceGroup := strings.Split(*vault.ID, "/")[4]

	resource := Resource{
		ID:       *vault.ID,
		Name:     *vault.Name,
		Location: *vault.Location,
		Description: JSONAllFieldsMarshaller{
			model.RecoveryServicesVaultDescription{
				Vault:         *vault,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
