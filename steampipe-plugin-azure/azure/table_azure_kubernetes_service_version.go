package azure

import (
	"context"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureAKSOrchestractor(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_kubernetes_service_version",
		Description: "Azure Kubernetes Service Version",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListKubernetesServiceVersion,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "location",
					Require: plugin.Required,
				},
				{
					Name:    "resource_type",
					Require: plugin.Optional,
				},
			},
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the orchestrator version profile list result.",
				Transform:   transform.FromField("Name")},
			{
				Name:        "id",
				Description: "ID of the orchestrator version profile list result.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID")},
			{
				Name:        "type",
				Description: "Type of the orchestrator version profile list result.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Type")},
			{
				Name:        "orchestrator_type",
				Description: "The orchestrator type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Orchestrator.OrchestratorType")},
			{
				Name:        "orchestrator_version",
				Description: "Orchestrator version (major, minor, patch).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Orchestrator.OrchestratorVersion")},
			{
				Name:        "default",
				Description: "Installed by default if version is not specified.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Orchestrator.Default")},
			{
				Name:        "is_preview",
				Description: "Whether Kubernetes version is currently in preview.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Orchestrator.IsPreview")},
			{
				Name:        "resource_type",
				Description: "Whether Kubernetes version is currently in preview.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Type"),
			},
			{
				Name:        "upgrades",
				Description: "The list of available upgrade versions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Orchestrator.Upgrades")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "location",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("Location"),
			},
		}),
	}
}
