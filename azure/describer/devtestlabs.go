package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/devtestlabs/mgmt/dtl"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func DevTestLabLab(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := dtl.NewLabsClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListBySubscription(ctx, "", "", nil, "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroupName := strings.Split(string(*v.ID), "/")[4]
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: JSONAllFieldsMarshaller{
					model.DevTestLabLabDescription{
						Lab:           v,
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
