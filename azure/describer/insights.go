package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/opengovern/og-azure-describer/azure/model"
	"strings"
	"time"
)

func DiagnosticSetting(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := monitorClientFactory.NewDiagnosticSettingsClient()
	resourceURI := "/subscriptions/" + subscription
	pager := client.NewListPager(resourceURI, nil)

	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, diagnosticSetting := range page.Value {
			resource := getDiagnosticSetting(ctx, diagnosticSetting)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getDiagnosticSetting(ctx context.Context, diagnosticSetting *armmonitor.DiagnosticSettingsResource) *Resource {
	var resourceGroup string
	if diagnosticSetting.Properties.StorageAccountID != nil {
		resourceGroup = strings.Split(*diagnosticSetting.Properties.StorageAccountID, "/")[4]
	} else if diagnosticSetting.Properties.EventHubAuthorizationRuleID != nil {
		resourceGroup = strings.Split(*diagnosticSetting.Properties.EventHubAuthorizationRuleID, "/")[4]
	} else {
		resourceGroup = strings.Split(*diagnosticSetting.Properties.WorkspaceID, "/")[4]
	}
	resource := Resource{
		ID:       *diagnosticSetting.ID,
		Name:     *diagnosticSetting.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.DiagnosticSettingDescription{
				DiagnosticSettingsResource: *diagnosticSetting,
				ResourceGroup:              resourceGroup,
			},
		},
	}
	return &resource
}

func LogAlert(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := monitorClientFactory.NewActivityLogAlertsClient()

	pager := client.NewListBySubscriptionIDPager(nil)

	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, logAlert := range page.Value {
			resource := getLogAlert(ctx, logAlert)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getLogAlert(ctx context.Context, logAlert *armmonitor.ActivityLogAlertResource) *Resource {
	resourceGroup := strings.Split(*logAlert.ID, "/")[4]

	resource := Resource{
		ID:       *logAlert.ID,
		Name:     *logAlert.Name,
		Location: *logAlert.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.LogAlertDescription{
				ActivityLogAlertResource: *logAlert,
				ResourceGroup:            resourceGroup,
			},
		},
	}

	return &resource
}

func LogProfile(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := monitorClientFactory.NewLogProfilesClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, r := range page.Value {
			resource := getLogProfile(ctx, r)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getLogProfile(ctx context.Context, logProfile *armmonitor.LogProfileResource) *Resource {
	resourceGroup := strings.Split(*logProfile.ID, "/")[4]
	location := "global"
	if logProfile.Location != nil {
		location = *logProfile.Location
	}
	resource := Resource{
		ID:       *logProfile.ID,
		Name:     *logProfile.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			Value: model.LogProfileDescription{
				LogProfileResource: *logProfile,
				ResourceGroup:      resourceGroup,
			},
		},
	}

	return &resource
}

func getMonitoringIntervalForGranularity(granularity string) string {
	switch strings.ToUpper(granularity) {
	case "DAILY":
		// 24 hours
		return "PT24H"
	case "HOURLY":
		// 1 hour
		return "PT1H"
	}
	// else 5 minutes
	return "PT5M"
}

func getMonitoringStartDateForGranularity(granularity string) string {
	switch strings.ToUpper(granularity) {
	case "DAILY":
		// Last 1 year
		return time.Now().UTC().AddDate(-1, 0, 0).Format(time.RFC3339)
	case "HOURLY":
		// Last 60 days
		return time.Now().UTC().AddDate(0, 0, -60).Format(time.RFC3339)
	}
	// Last 5 days
	return time.Now().UTC().AddDate(0, 0, -5).Format(time.RFC3339)
}

func listAzureMonitorMetricStatistics(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, granularity string, metricNameSpace string, metricNames string, dimensionValue string) ([]model.MonitoringMetric, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	metricsClient := monitorClientFactory.NewMetricsClient()

	interval := getMonitoringIntervalForGranularity(granularity)
	aggregation := "average,count,maximum,minimum,total"
	timeSpan := getMonitoringStartDateForGranularity(granularity) + "/" + time.Now().UTC().AddDate(0, 0, 1).Format(time.RFC3339) // Retrieve data within a year
	orderBy := "timestamp"
	top := int32(1000) // Maximum number of record fetch with given interval
	filter := ""
	resultType := armmonitor.ResultTypeData
	options := armmonitor.MetricsClientListOptions{
		Aggregation:     &aggregation,
		Interval:        &interval,
		Timespan:        &timeSpan,
		Orderby:         &orderBy,
		Top:             &top,
		Filter:          &filter,
		Metricnames:     &metricNames,
		ResultType:      &resultType,
		Metricnamespace: &metricNameSpace,
	}

	result, err := metricsClient.List(ctx, dimensionValue, &options)
	if err != nil {
		return nil, err
	}
	var values []model.MonitoringMetric
	for _, metric := range result.Value {
		for _, timeseries := range metric.Timeseries {
			for _, data := range timeseries.Data {
				if data.Average != nil {
					values = append(values, model.MonitoringMetric{
						DimensionValue: dimensionValue,
						TimeStamp:      data.TimeStamp.Format(time.RFC3339),
						Maximum:        data.Maximum,
						Minimum:        data.Minimum,
						Average:        data.Average,
						Sum:            data.Total,
						SampleCount:    data.Count,
						Unit:           string(*metric.Unit),
					})
				}
			}
		}
	}
	return values, nil
}

func AutoscaleSetting(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := monitorClientFactory.NewAutoscaleSettingsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getAutoscaleSetting(v)
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func getAutoscaleSetting(v *armmonitor.AutoscaleSettingResource) Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	return Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AutoscaleSettingDescription{
				AutoscaleSettingsResource: *v,
				ResourceGroup:             resourceGroup,
			},
		},
	}
}
