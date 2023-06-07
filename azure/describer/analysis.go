package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2017-08-01/analysisservices"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func AnalysisService(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := analysisservices.NewServersClient(subscription)
	client.Authorizer = authorizer
	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for _, server := range *result.Value {
		resourceGroupName := strings.Split(string(*server.ID), "/")[4]

		resource := Resource{
			ID:       *server.ID,
			Name:     *server.Name,
			Location: *server.Location,
			Description: JSONAllFieldsMarshaller{
				model.AnalysisServiceServerDescription{
					Server:        server,
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
