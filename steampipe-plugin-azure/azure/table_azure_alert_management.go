package azure

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureAlertMangement(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_alert_management",
		Description: "Azure Alert Management Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    getAlertManagement, // this will be updated from the auto generated code from `pkg/kaytu-es-sdk/azure_resources_clients.go`
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "InvalidApiVersionParameter", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: listAlertManagements, // this will be updated from the auto generated code from `pkg/kaytu-es-sdk/azure_resources_clients.go`
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "target_resource",
					Require: plugin.Optional,
				},
				{
					Name:    "target_resource_type",
					Require: plugin.Optional,
				},
				{
					Name:    "resource_group",
					Require: plugin.Optional,
				},
				{
					Name:    "alert_rule",
					Require: plugin.Optional,
				},
				{
					Name:    "smart_group_id",
					Require: plugin.Optional,
				},
				{
					Name:    "sort_order",
					Require: plugin.Optional,
				},
				{
					Name:    "custom_time_range",
					Require: plugin.Optional,
				},
				{
					Name:    "sort_by",
					Require: plugin.Optional,
				},
				{
					Name:    "monitor_service",
					Require: plugin.Optional,
				},
				{
					Name:    "monitor_condition",
					Require: plugin.Optional,
				},
				{
					Name:    "severity",
					Require: plugin.Optional,
				},
				{
					Name:    "alert_state",
					Require: plugin.Optional,
				},
				{
					Name:    "time_range",
					Require: plugin.Optional,
				},
			},
		},
		Columns: azureColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "A friendly name that identifies an Alert management service.",
			},
			{
				Name:        "id",
				Description: "Azure resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sort_order",
				Description: "Sort order of the alert management.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("sort_order"),
			},
			{
				Name:        "sort_by",
				Description: "Sort the query results by input field, default value is 'lastModifiedDateTime'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("sort_by"),
			},
			{
				Name:        "custom_time_range",
				Description: "Filter by custom time range in the format <start-time>/<end-time> where time is in (ISO-8601 format).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("custom_time_range"),
			},
			{
				Name:        "time_range",
				Description: "Filter by time range. Possible values are '1h', '1d', '7d' or '30d'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("time_range"),
			},
			{
				Name:        "severity",
				Description: "Severity of alert Sev0 being highest and Sev4 being lowest. Possible values include: 'Sev0', 'Sev1', 'Sev2', 'Sev3', 'Sev4'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.Severity"),
			},
			{
				Name:        "signal_type",
				Description: "The type of signal the alert is based on, which could be metrics, logs or activity logs. Possible values include: 'Metric', 'Log', 'Unknown'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.SignalType"),
			},
			{
				Name:        "alert_state",
				Description: "Alert object state, which can be modified by the user. Possible values include: 'AlertStateNew', 'AlertStateAcknowledged', 'AlertStateClosed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.AlertState"),
			},
			{
				Name:        "monitor_condition",
				Description: "Can be 'Fired' or 'Resolved', which represents whether the underlying conditions have crossed the defined alert rule thresholds. Possible values include: 'Fired', 'Resolved'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.MonitorCondition"),
			},
			{
				Name:        "target_resource",
				Description: "Target ARM resource, on which alert got created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.TargetResource"),
			},
			{
				Name:        "target_resource_name",
				Description: "Name of the target ARM resource, on which alert got created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.TargetResourceName"),
			},
			{
				Name:        "target_resource_type",
				Description: "Resource type of target ARM resource, on which alert got created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.TargetResourceType"),
			},
			{
				Name:        "monitor_service",
				Description: "Monitor service on which the rule(monitor) is set. Possible values include: 'ApplicationInsights', 'ActivityLogAdministrative', 'ActivityLogSecurity', 'ActivityLogRecommendation', 'ActivityLogPolicy', 'ActivityLogAutoscale', 'LogAnalytics', 'Nagios', 'Platform', 'SCOM', 'ServiceHealth', 'SmartDetector', 'VMInsights', 'Zabbix', 'ResourceHealth'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.MonitorService"),
			},
			{
				Name:        "alert_rule",
				Description: "Rule(monitor) which fired alert instance. Depending on the monitor service, this would be ARM ID or name of the rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.AlertRule"),
			},
			{
				Name:        "source_created_id",
				Description: "Unique ID created by monitor service for each alert instance. This could be used to track the issue at the monitor service, in case of Nagios, Zabbix, SCOM, etc.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.SourceCreatedID"),
			},
			{
				Name:        "smart_group_id",
				Description: "Unique ID of the smart group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.SmartGroupID"),
			},
			{
				Name:        "smart_grouping_reason",
				Description: "Verbose reason describing the reason why this alert instance is added to a smart group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.SmartGroupingReason"),
			},
			{
				Name:        "start_date_time",
				Description: "Creation time(ISO-8601 format) of alert instance.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Properties.Essentials.StartDateTime").Transform(convertDateToTime),
			},
			{
				Name:        "last_modified_date_time",
				Description: "Last modification time(ISO-8601 format) of alert instance.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Properties.Essentials.LastModifiedDateTime").Transform(convertDateToTime),
			},
			{
				Name:        "monitor_condition_resolved_date_time",
				Description: "Resolved time(ISO-8601 format) of alert instance. This will be updated when monitor service resolves the alert instance because the rule condition is no longer met.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Properties.Essentials.MonitorConditionResolvedDateTime").Transform(convertDateToTime),
			},
			{
				Name:        "last_modified_user_name",
				Description: "User who last modified the alert, in case of monitor service updates user would be 'system', otherwise name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Essentials.LastModifiedUserName"),
			},
			{
				Name:        "context",
				Description: "The context of the alert management.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.Context"),
			},
			{
				Name:        "egress_config",
				Description: "The egress config for the context management.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Properties.EgressConfig"),
			},

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
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("resource_group"),
			},
		}),
	}
}
