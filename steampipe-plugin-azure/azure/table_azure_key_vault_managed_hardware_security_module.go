package azure

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureKeyVaultManagedHardwareSecurityModule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_key_vault_managed_hardware_security_module",
		Description: "Azure Key Vault Managed Hardware Security Module",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    kaytu.GetKeyVaultManagedHardwareSecurityModule,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListKeyVaultManagedHardwareSecurityModule,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the managed HSM Pool.",
				Transform:   transform.FromField("Description.ManagedHsm.Name")},
			{
				Name:        "id",
				Description: "The Azure Resource Manager resource ID for the managed HSM Pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.ID")},
			{
				Name:        "type",
				Description: "The resource type of the managed HSM Pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.Type")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state. Possible values include: 'ProvisioningStateSucceeded', 'ProvisioningStateProvisioning', 'ProvisioningStateFailed', 'ProvisioningStateUpdating', 'ProvisioningStateDeleting', 'ProvisioningStateActivated', 'ProvisioningStateSecurityDomainRestore', 'ProvisioningStateRestoring'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.Properties.ProvisioningState")},
			{
				Name:        "hsm_uri",
				Description: "The URI of the managed hsm pool for performing operations on keys.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.Properties.HsmURI")},
			{
				Name:        "enable_soft_delete",
				Description: "Property to specify whether the 'soft delete' functionality is enabled for this managed HSM pool. If it's not set to any value(true or false) when creating new managed HSM pool, it will be set to true by default. Once set to true, it cannot be reverted to false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ManagedHsm.Properties.EnableSoftDelete")},
			{
				Name:        "soft_delete_retention_in_days",
				Description: "Indicates softDelete data retention days. It accepts >=7 and <=90.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ManagedHsm.Properties.SoftDeleteRetentionInDays")},
			{
				Name:        "enable_purge_protection",
				Description: "Property specifying whether protection against purge is enabled for this managed HSM pool. Setting this property to true activates protection against purge for this managed HSM pool and its content - only the Managed HSM service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ManagedHsm.Properties.EnablePurgeProtection")},
			{
				Name:        "status_message",
				Description: "Resource Status Message.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.Properties.StatusMessage")},
			{
				Name:        "create_mode",
				Description: "The create mode to indicate whether the resource is being created or is being recovered from a deleted resource. Possible values include: 'CreateModeRecover', 'CreateModeDefault'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.Properties.CreateMode")},
			{
				Name:        "sku_family",
				Description: "Contains SKU family name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.SKU.Family")},
			{
				Name:        "sku_name",
				Description: "SKU name to specify whether the key vault is a standard vault or a premium vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ManagedHsm.SKU.Name"),
			},
			{
				Name:        "tenant_id",
				Description: "The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ManagedHsm.Properties.TenantID"),
			},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the managed HSM.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.DiagnosticSettingsResources")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedHsm.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedHsm.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ManagedHsm.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ManagedHsm.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit
				Transform: transform.

					//// HYDRATE FUNCTIONS
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Create session

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings