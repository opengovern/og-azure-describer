package azure

import (
	"context"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureRecoveryServicesBackupItem(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_recovery_services_backup_item",
		Description: "Azure Recovery Services Backup Item",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRecoveryServicesBackupItem,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "404"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "vault_name",
					Require: plugin.Optional,
				},
				{
					Name:    "resource_group",
					Require: plugin.Optional,
				},
			},
		},
		Columns: azureColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the table service",
				Transform:   transform.FromField("Description.Item.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a table service uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Item.ID"),
			},
			{
				Name:        "vault_name",
				Description: "Backup item vault name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VaultName"),
			},
			{
				Name:        "policy_name",
				Description: "Backup item policy name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Item.Properties.PolicyName"),
			},
			{
				Name:        "policy_id",
				Description: "Backup item policy id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Item.Properties.PolicyID"),
			},
			{
				Name:        "source_resource_id",
				Description: "Backup item source resource identifier",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Item.Properties.SourceResourceID"),
			},
			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Item.Location").Transform(formatRegion).Transform(toLower),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Item.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Item.ID").Transform(idToAkas),
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
