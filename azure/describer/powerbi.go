package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/powerbidedicated/mgmt/powerbidedicated"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func PowerBIDedicatedCapacity(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := powerbidedicated.NewCapacitiesClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}
	if result.Value == nil {
		return nil, nil
	}

	var values []Resource

	for _, v := range *result.Value {
		resourceGroupName := strings.Split(string(*v.ID), "/")[4]
		resource := Resource{
			ID:       *v.ID,
			Name:     *v.Name,
			Location: *v.Location,
			Description: JSONAllFieldsMarshaller{
				model.PowerBIDedicatedCapacityDescription{
					Capacity:      v,
					ResourceGroup: resourceGroupName,
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
