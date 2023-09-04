package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func DataProtectionBackupVaults(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armdataprotection.NewBackupVaultsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewGetInSubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getDataProtectionBackupVaults(ctx, v)
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

func getDataProtectionBackupVaults(ctx context.Context, v *armdataprotection.BackupVaultResource) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.DataProtectionBackupVaultsDescription{
				BackupVaults:  *v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}

func DataProtectionBackupVaultsBackupPolicies(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armdataprotection.NewBackupVaultsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	policiesClient, err := armdataprotection.NewBackupPoliciesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewGetInSubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := getDataProtectionBackupVaultsBackupPolicies(ctx, policiesClient, v)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func getDataProtectionBackupVaultsBackupPolicies(ctx context.Context, client *armdataprotection.BackupPoliciesClient, v *armdataprotection.BackupVaultResource) ([]Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	pager := client.NewListPager(resourceGroup, *v.Name, nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, p := range page.Value {
			resourceGroup := strings.Split(*v.ID, "/")[4]

			resource := Resource{
				ID:       *p.ID,
				Name:     *p.Name,
				Location: *v.Location,
				Description: JSONAllFieldsMarshaller{
					Value: model.DataProtectionBackupVaultsBackupPoliciesDescription{
						BackupPolicies: *p,
						ResourceGroup:  resourceGroup,
					},
				},
			}
			values = append(values, resource)
		}
	}
	return values, nil
}
