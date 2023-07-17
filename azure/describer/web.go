package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func AppServiceEnvironment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armappservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewEnvironmentsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetAppServiceEnvironment(ctx, v)
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

func GetAppServiceEnvironment(ctx context.Context, v *armappservice.EnvironmentResource) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			model.AppServiceEnvironmentDescription{
				AppServiceEnvironmentResource: *v,
				ResourceGroup:                 resourceGroup,
			},
		},
	}

	return &resource
}

func AppServiceFunctionApp(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armappservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewWebAppsClient()
	webClient := clientFactory.NewWebAppsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetAppServiceFunctionApp(ctx, webClient, v)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, err
}

func GetAppServiceFunctionApp(ctx context.Context, webClient *armappservice.WebAppsClient, v *armappservice.Site) (*Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	configuration, err := webClient.GetConfiguration(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
	if err != nil {
		return nil, err
	}
	authSettings, err := webClient.GetAuthSettings(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
	if err != nil {
		return nil, err
	}
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			model.AppServiceFunctionAppDescription{
				Site:               *v,
				SiteAuthSettings:   authSettings.SiteAuthSettings,
				SiteConfigResource: configuration.SiteConfigResource,
				ResourceGroup:      resourceGroup,
			},
		},
	}
	return &resource, nil
}

func AppServiceWebApp(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armappservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewWebAppsClient()
	webClient := clientFactory.NewWebAppsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetAppServiceWebApp(ctx, webClient, v)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, err
}

func GetAppServiceWebApp(ctx context.Context, webClient *armappservice.WebAppsClient, v *armappservice.Site) (*Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	configuration, err := webClient.GetConfiguration(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
	if err != nil {
		return nil, err
	}
	authSettings, err := webClient.GetAuthSettings(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
	if err != nil {
		return nil, err
	}

	vnet, err := webClient.GetVnetConnection(ctx, *v.Properties.ResourceGroup, *v.Name, *v.Properties.VirtualNetworkSubnetID, nil)
	if err != nil {
		return nil, err
	}

	location := ""
	if v.Location != nil {
		location = *v.Location
	}

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			model.AppServiceWebAppDescription{
				Site:               *v,
				SiteConfigResource: configuration.SiteConfigResource,
				SiteAuthSettings:   authSettings.SiteAuthSettings,
				VnetInfo:           vnet.VnetInfoResource,
				ResourceGroup:      resourceGroup,
			},
		},
	}

	return &resource, nil
}

func AppServicePlan(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armappservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewPlansClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := GetAppServicePlan(ctx, client, v)
			if err != nil {
				return nil, err
			}
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

func GetAppServicePlan(ctx context.Context, client *armappservice.PlansClient, v *armappservice.Plan) (*Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	location := ""
	if v.Location != nil {
		location = *v.Location
	}

	var webApps []*armappservice.Site

	pager := client.NewListWebAppsPager(resourceGroup, *v.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		webApps = append(webApps, page.Value...)
	}

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			model.AppServicePlanDescription{
				Plan:          *v,
				Apps:          webApps,
				ResourceGroup: resourceGroup,
			},
		},
	}

	return &resource, nil
}

func AppContainerApps(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armappservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewContainerAppsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetAppContainerApps(ctx, v)
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

func GetAppContainerApps(ctx context.Context, server *armappservice.ContainerApp) *Resource {
	resourceGroupName := strings.Split(string(*server.ID), "/")[4]

	resource := Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: JSONAllFieldsMarshaller{
			model.ContainerAppDescription{
				Server:        *server,
				ResourceGroup: resourceGroupName,
			},
		},
	}

	return &resource
}

func WebServerFarms(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armappservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewPlansClient()

	pager := client.NewListByResourceGroupPager(fmt.Sprintf("/subscriptions/%s", subscription), nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetWebServerFarm(ctx, v)
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

func GetWebServerFarm(ctx context.Context, v *armappservice.Plan) *Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			model.WebServerFarmsDescription{
				ServerFarm:    *v,
				ResourceGroup: resourceGroupName,
			},
		},
	}
	return &resource
}
