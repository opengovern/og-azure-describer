package azure

import (
	"context"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDataMigrationServices(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_datamigration_services",
		Description: "Azure DataMigration Services",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetDataMigrationService,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDataMigrationService,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the services.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ID")},
			{
				Name:        "name",
				Description: "The name of the services.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Service.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Service.ID").Transform(idToAkas),
			},
		}),
	}
}