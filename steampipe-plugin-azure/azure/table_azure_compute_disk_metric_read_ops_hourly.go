package azure

import (
	"context"

	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeDiskMetricReadOpsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_disk_metric_read_ops_hourly",
		Description: "Azure Compute Disk Metrics - Read Ops (Hourly)",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListComputeDiskReadOpsHourly,
		},
		Columns: kaytuMonitoringMetricColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the disk.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.MonitoringMetric.DimensionValue").Transform(lastPathElement),
			},
		}),
	}
}