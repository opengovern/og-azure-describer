package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/cdn/mgmt/cdn"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func CdnProfiles(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := cdn.NewProfilesClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroup := strings.Split(*v.ID, "/")[4]

			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: JSONAllFieldsMarshaller{
					model.CDNProfileDescription{
						Profile:       v,
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
