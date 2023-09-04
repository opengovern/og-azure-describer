package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
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

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	var values []Resource
	pager := client.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resource, err := GetRecoveryServicesVault(ctx, diagnosticClient, vault)
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

func GetRecoveryServicesVault(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, vault *armrecoveryservices.Vault) (*Resource, error) {
	resourceGroup := strings.Split(*vault.ID, "/")[4]

	var diagnostic []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*vault.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		diagnostic = append(diagnostic, page.Value...)
	}

	resource := Resource{
		ID:       *vault.ID,
		Name:     *vault.Name,
		Location: *vault.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.RecoveryServicesVaultDescription{
				Vault:                      *vault,
				DiagnosticSettingsResource: diagnostic,
				ResourceGroup:              resourceGroup,
			},
		},
	}
	return &resource, nil
}
