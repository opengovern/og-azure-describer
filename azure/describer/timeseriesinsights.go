package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/timeseriesinsights/mgmt/2018-08-15-preview/timeseriesinsights"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func TimeSeriesInsightsEnvironments(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := timeseriesinsights.NewEnvironmentsClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListBySubscription(context.Background())
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, record := range *result.Value {
		v, _ := record.AsEnvironmentResource()
		resourceGroup := strings.Split(*v.ID, "/")[4]

		resource := Resource{
			ID:       *v.ID,
			Name:     *v.Name,
			Location: *v.Location,
			Description: JSONAllFieldsMarshaller{
				model.TimeSeriesInsightsEnvironmentsDescription{
					Environment:   v,
					ResourceGroup: resourceGroup,
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
