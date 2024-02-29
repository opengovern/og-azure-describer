package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
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

func RecoveryServicesBackupJobs(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	vaultClientFactory, err := armrecoveryservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	vaultClient := vaultClientFactory.NewVaultsClient()

	clientFactory, err := armrecoveryservicesbackup.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewBackupJobsClient()

	var values []Resource
	pager := vaultClient.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, vault := range page.Value {
			resourceGroup := strings.Split(*vault.ID, "/")[4]
			vaultBackupJobs, err := ListRecoveryServicesVaultBackupJobs(ctx, client, *vault.Name, resourceGroup)
			if err != nil {
				return nil, err
			}
			for _, resource := range vaultBackupJobs {
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

func ListRecoveryServicesVaultBackupJobs(ctx context.Context, client *armrecoveryservicesbackup.BackupJobsClient, vaultName, resourceGroup string) ([]Resource, error) {
	pager := client.NewListPager(vaultName, resourceGroup, &armrecoveryservicesbackup.BackupJobsClientListOptions{Filter: nil,
		SkipToken: nil,
	})
	var resources []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, job := range page.Value {
			resource, err := GetRecoveryServicesBackupJob(resourceGroup, vaultName, job)
			if err != nil {
				return nil, err
			}
			resources = append(resources, *resource)
		}
	}
	return resources, nil
}

func GetRecoveryServicesBackupJob(resourceGroup, vaultName string, job *armrecoveryservicesbackup.JobResource) (*Resource, error) {
	properties, err := backupJobProperties(job)
	if err != nil {
		return nil, err
	}
	resource := Resource{
		ID:   *job.ID,
		Name: *job.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.RecoveryServicesBackupJobDescription{
				Job: struct {
					Name     *string
					ID       *string
					Type     *string
					ETag     *string
					Tags     map[string]*string
					Location *string
				}{
					Name:     job.Name,
					ID:       job.ID,
					Location: job.Location,
					Type:     job.Type,
					Tags:     job.Tags,
					ETag:     job.ETag,
				},
				VaultName:     vaultName,
				Properties:    properties,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource, nil
}

func backupJobProperties(data *armrecoveryservicesbackup.JobResource) (map[string]interface{}, error) {
	output := make(map[string]interface{})

	if data.Properties != nil {
		if data.Properties.GetJob() != nil {
			if data.Properties.GetJob().ActivityID != nil {
				output["ActivityID"] = data.Properties.GetJob().ActivityID
			}
			if data.Properties.GetJob().BackupManagementType != nil {
				output["BackupManagementType"] = data.Properties.GetJob().BackupManagementType
			}
			if data.Properties.GetJob().JobType != nil {
				output["JobType"] = data.Properties.GetJob().JobType
			}
			if data.Properties.GetJob().EndTime != nil {
				output["EndTime"] = data.Properties.GetJob().EndTime
			}
			if data.Properties.GetJob().EntityFriendlyName != nil {
				output["EntityFriendlyName"] = data.Properties.GetJob().EntityFriendlyName
			}
			if data.Properties.GetJob().Operation != nil {
				output["Operation"] = data.Properties.GetJob().Operation
			}
			if data.Properties.GetJob().StartTime != nil {
				output["StartTime"] = data.Properties.GetJob().StartTime
			}
			if data.Properties.GetJob().Status != nil {
				output["Status"] = data.Properties.GetJob().Status
			}
		}
	}
	return output, nil
}
