package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kaytu-io/kaytu-util/pkg/describe/enums"

	"github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2019-11-01/costmanagement"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

const resourceTypeDimension = "resourceType"
const subscriptionDimension = "SubscriptionId"

func cost(ctx context.Context, authorizer autorest.Authorizer, subscription string, from time.Time, to time.Time, dimension string) ([]model.CostManagementQueryRow, *string, error) {
	client := costmanagement.NewQueryClient(subscription)
	client.Authorizer = authorizer

	scope := fmt.Sprintf("subscriptions/%s", subscription)

	groupings := []costmanagement.QueryGrouping{
		{
			Type: costmanagement.QueryColumnTypeDimension,
			Name: &dimension,
		},
	}

	costAggregationString := "Cost"

	var costs, err = client.Usage(ctx, scope, costmanagement.QueryDefinition{
		Type:      costmanagement.ExportTypeAmortizedCost,
		Timeframe: costmanagement.TimeframeTypeCustom,
		TimePeriod: &costmanagement.QueryTimePeriod{
			From: &date.Time{Time: from},
			To:   &date.Time{Time: to},
		},
		Dataset: &costmanagement.QueryDataset{
			Granularity: costmanagement.GranularityTypeDaily,
			Grouping:    &groupings,
			Aggregation: map[string]*costmanagement.QueryAggregation{
				"Cost": {
					Name:     &costAggregationString,
					Function: costmanagement.FunctionTypeSum,
				},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}
	mapResult := make([]map[string]any, 0)
	for _, row := range *costs.Rows {
		rowMap := make(map[string]any)
		for i, column := range *costs.Columns {
			rowMap[*column.Name] = row[i]
		}
		mapResult = append(mapResult, rowMap)
	}
	jsonMapResult, err := json.Marshal(mapResult)
	if err != nil {
		return nil, nil, err
	}

	result := make([]model.CostManagementQueryRow, 0, len(mapResult))
	err = json.Unmarshal(jsonMapResult, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, costs.Location, nil
}

func DailyCostByResourceType(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := costmanagement.NewQueryClient(subscription)
	client.Authorizer = authorizer

	triggerType := GetTriggerTypeFromContext(ctx)
	from := time.Now().AddDate(0, 0, -7)
	if triggerType == enums.DescribeTriggerTypeInitialDiscovery {
		from = time.Now().AddDate(0, -1, -7)
	}

	costResult, locationPtr, err := cost(ctx, authorizer, subscription, from, time.Now(), resourceTypeDimension)
	if err != nil {
		return nil, err
	}
	location := "global"
	if locationPtr != nil {
		location = *locationPtr
	}
	var values []Resource
	for _, row := range costResult {
		resource := Resource{
			ID:       fmt.Sprintf("resource-cost-%s/%s-%d", subscription, *row.ResourceType, row.UsageDate),
			Location: location,
			Description: JSONAllFieldsMarshaller{
				model.CostManagementCostByResourceTypeDescription{
					CostManagementCostByResourceType: row,
				},
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func DailyCostBySubscription(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := costmanagement.NewQueryClient(subscription)
	client.Authorizer = authorizer

	triggerType := GetTriggerTypeFromContext(ctx)
	from := time.Now().AddDate(0, 0, -7)
	if triggerType == enums.DescribeTriggerTypeInitialDiscovery {
		from = time.Now().AddDate(0, -1, -7)
	}

	costResult, locationPtr, err := cost(ctx, authorizer, subscription, from, time.Now(), subscriptionDimension)
	if err != nil {
		return nil, err
	}
	location := "global"
	if locationPtr != nil {
		location = *locationPtr
	}
	var values []Resource
	for _, row := range costResult {
		resource := Resource{
			ID:       fmt.Sprintf("resource-cost-%s/%d", subscription, row.UsageDate),
			Location: location,
			Description: JSONAllFieldsMarshaller{
				model.CostManagementCostBySubscriptionDescription{
					CostManagementCostBySubscription: row,
				},
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}
