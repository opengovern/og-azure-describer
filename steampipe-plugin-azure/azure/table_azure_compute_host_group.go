package azure

import (
	"context"

	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureComputeHostGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_host_group",
		Description: "Azure Compute HostGroup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetComputeHostGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListComputeHostGroup,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the hostgroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostGroup.ID")},
			{
				Name:        "name",
				Description: "The name of the hostgroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostGroup.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HostGroup.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.HostGroup.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.HostGroup.ID").Transform(idToAkas),
			},
		}),
	}
}