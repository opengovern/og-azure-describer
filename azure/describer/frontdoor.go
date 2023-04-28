package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-05-01/frontdoor"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-04-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func FrontDoor(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
	insightsClient.Authorizer = authorizer

	client := frontdoor.NewFrontDoorsClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, door := range result.Values() {
			resourceGroup := strings.Split(*door.ID, "/")[4]

			frontDoorListOp, err := insightsClient.List(ctx, *door.ID)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:       *door.ID,
				Name:     *door.Name,
				Location: *door.Location,
				Description: model.FrontdoorDescription{
					FrontDoor:                   door,
					DiagnosticSettingsResources: frontDoorListOp.Value,
					ResourceGroup:               resourceGroup,
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
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
