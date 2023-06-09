package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/operationalinsights/mgmt/operationalinsights"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func OperationalInsightsWorkspaces(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := operationalinsights.NewWorkspacesClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(context.Background())
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range *result.Value {
		resourceGroup := strings.Split(*v.ID, "/")[4]

		resource := Resource{
			ID:       *v.ID,
			Name:     *v.Name,
			Location: *v.Location,
			Description: JSONAllFieldsMarshaller{
				model.OperationalInsightsWorkspacesDescription{
					Workspace:     v,
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
