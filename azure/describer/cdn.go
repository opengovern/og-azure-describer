package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func CdnProfiles(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armcdn.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewProfilesClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getCdnProfiles(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getCdnProfiles(ctx context.Context, v *armcdn.Profile) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.CDNProfileDescription{
				Profile:       *v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}

func CdnEndpoint(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armcdn.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewProfilesClient()
	endpointsClient := clientFactory.NewEndpointsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := getCdnProfilesEndpoints(ctx, endpointsClient, v)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func getCdnProfilesEndpoints(ctx context.Context, endpointsClient *armcdn.EndpointsClient, v *armcdn.Profile) ([]Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	pager := endpointsClient.NewListByProfilePager(resourceGroup, *v.Name, nil)

	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, endpoint := range page.Value {
			resource := getCdnEndpoint(ctx, v, endpoint, resourceGroup)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getCdnEndpoint(ctx context.Context, v *armcdn.Profile, endpoint *armcdn.Endpoint, resourceGroup string) *Resource {
	return &Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.CDNEndpointDescription{
				Endpoint:      *endpoint,
				ResourceGroup: resourceGroup,
			},
		},
	}
}
