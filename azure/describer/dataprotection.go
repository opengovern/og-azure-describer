package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/dataprotection/mgmt/2021-07-01/dataprotection"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func DataProtectionBackupVaults(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := dataprotection.NewBackupVaultsClient(subscription)
	client.Authorizer = authorizer

	result, err := client.GetInSubscription(ctx)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroup := strings.Split(*v.ID, "/")[4]

			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: JSONAllFieldsMarshaller{
					model.DataProtectionBackupVaultsDescription{
						BackupVaults:  v,
						ResourceGroup: resourceGroup,
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

func DataProtectionBackupVaultsBackupPolicies(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	clientVault := dataprotection.NewBackupVaultsClient(subscription)
	clientVault.Authorizer = authorizer

	clientPolicy := dataprotection.NewBackupPoliciesClient(subscription)
	clientPolicy.Authorizer = authorizer

	result, err := clientVault.GetInSubscription(ctx)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroup := strings.Split(*v.ID, "/")[4]

			resultPolicy, err := clientPolicy.List(ctx, *v.Name, resourceGroup)
			if err != nil {
				return nil, err
			}

			for {
				for _, p := range resultPolicy.Values() {
					resourceGroup := strings.Split(*v.ID, "/")[4]

					resource := Resource{
						ID:       *p.ID,
						Name:     *p.Name,
						Location: *v.Location,
						Description: JSONAllFieldsMarshaller{
							model.DataProtectionBackupVaultsBackupPoliciesDescription{
								BackupPolicies: p,
								ResourceGroup:  resourceGroup,
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

				if !resultPolicy.NotDone() {
					break
				}
				err = resultPolicy.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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
