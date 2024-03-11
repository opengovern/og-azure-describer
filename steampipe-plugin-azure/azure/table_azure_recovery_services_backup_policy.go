package azure

import (
	"context"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureRecoveryServicesBackupPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_recovery_services_backup_policy",
		Description: "Azure Recovery Services Backup Policy",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRecoveryServicesBackupPolicy,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "404"}),
			},
		},
		Columns: azureColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the table service",
				Transform:   transform.FromField("Description.Policy.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a table service uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.ID"),
			},
			{
				Name:        "vault_name",
				Description: "Backup item vault name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VaultName"),
			},
			{
				Name:        "backup_management_type",
				Description: "Backup management type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Properties.BackupManagementType"),
			},
			{
				Name:        "instant_rp_retention_range_in_days",
				Description: "Backup management type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Properties.InstantRpRetentionRangeInDays"),
			},
			{
				Name:        "policy_type",
				Description: "Backup management type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Properties.PolicyType"),
			},
			{
				Name:        "protected_items_count",
				Description: "ProtectedItemsCount",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Properties.ProtectedItemsCount"),
			},
			{
				Name:        "retention_policy",
				Description: "RetentionPolicy",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Properties.RetentionPolicy"),
			},
			{
				Name:        "schedule_policy",
				Description: "SchedulePolicy",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy.Properties.SchedulePolicy"),
			},
			// Azure standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Policy.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
