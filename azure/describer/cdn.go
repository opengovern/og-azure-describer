package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cdn/mgmt/cdn"

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

func CdnEndpoint(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
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
			endpointClient := cdn.NewEndpointsClient(subscription)
			endpointClient.Authorizer = authorizer
			endpointResult, err := endpointClient.ListByProfile(ctx, resourceGroup, *v.Name)
			if err != nil {
				return nil, err
			}
			for _, endpoint := range endpointResult.Values() {
				resource := Resource{
					ID:       *v.ID,
					Name:     *v.Name,
					Location: *v.Location,
					Description: JSONAllFieldsMarshaller{
						model.CDNEndpointDescription{
							Endpoint:      endpoint,
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
			if !endpointResult.NotDone() {
				break
			}
			err = endpointResult.NextWithContext(ctx)
			if err != nil {
				return nil, err
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
