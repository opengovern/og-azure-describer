package azure

import (
	"context"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureManagementGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_management_group",
		Description: "Azure Management Group.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    kaytu.GetManagementGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListManagementGroup,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The fully qualified ID for the management group.",
				Transform:   transform.FromField("Description.Group.ID"),
			},
			{
				Name:        "name",
				Description: "The name of the management group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Name")},
			{
				Name:        "type",
				Description: "The type of the management group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Type")},
			{
				Name:        "display_name",
				Description: "The friendly name of the management group.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Group.Properties.DisplayName")},
			{
				Name:        "tenant_id",
				Description: "The AAD Tenant ID associated with the management group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Properties.TenantID")},
			{
				Name:        "parent",
				Description: "The associated parent management group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Group.Properties.Details.Parent"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Group.ID").Transform(idToAkas),
			},
		},
	}
}
