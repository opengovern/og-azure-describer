package azure

import (
	"context"

	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDevTestLabLab(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_devtestlab_lab",
		Description: "Azure DevTestLab Lab",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetDevTestLabLab,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDevTestLabLab,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the lab.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lab.ID")},
			{
				Name:        "name",
				Description: "The name of the lab.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lab.Properties.VaultName")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lab.Properties.VaultName")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Lab.Properties.LabStorageType")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Lab.ID").Transform(idToAkas),
			},
		}),
	}
}