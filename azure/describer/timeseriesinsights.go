package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func TimeSeriesInsightsEnvironments(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armtimeseriesinsights.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewEnvironmentsClient()

	list, err := client.ListBySubscription(ctx, nil)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, record := range list.Value {
		resource := GetTimeSeriesInsightsEnvironments(ctx, record)
		if stream != nil {
			if err := (*stream)(*resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetTimeSeriesInsightsEnvironments(ctx context.Context, record armtimeseriesinsights.EnvironmentResourceClassification) *Resource {
	v := record.GetEnvironmentResource()
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.TimeSeriesInsightsEnvironmentsDescription{
				Environment:   v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
