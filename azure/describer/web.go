package describer

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	appservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func AppServiceEnvironment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := appservice.NewEnvironmentsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

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

func GetAppServiceEnvironment(ctx context.Context, v *appservice.EnvironmentResource) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AppServiceEnvironmentDescription{
				AppServiceEnvironmentResource: *v,
				ResourceGroup:                 resourceGroup,
			},
		},
	}

	return &resource
}

func AppServiceFunctionApp(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := appservice.NewWebAppsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	webClient, err := appservice.NewWebAppsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

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

func GetAppServiceFunctionApp(ctx context.Context, webClient *appservice.WebAppsClient, v *appservice.Site) (*Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	configuration, err := webClient.GetConfiguration(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
	if err != nil {
		return nil, err
	}
	authSettings, err := webClient.GetAuthSettings(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
	if err != nil {
		//return nil, err
	}
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AppServiceFunctionAppDescription{
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
	client, err := appservice.NewWebAppsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	webClient, err := appservice.NewWebAppsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

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
			if resource == nil {
				continue
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

func GetAppServiceWebApp(ctx context.Context, webClient *appservice.WebAppsClient, v *appservice.Site) (*Resource, error) {
	var err error

	resourceGroup := strings.Split(*v.ID, "/")[4]
	configuration := appservice.WebAppsClientGetConfigurationResponse{}
	if v.Properties != nil && v.Properties.ResourceGroup != nil && v.Name != nil {
		configuration, err = webClient.GetConfiguration(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
		if err != nil {
			return nil, err
		}
	}
	authSettings := appservice.WebAppsClientGetAuthSettingsResponse{}
	if v.Properties != nil && v.Properties.ResourceGroup != nil && v.Name != nil {
		authSettings, err = webClient.GetAuthSettings(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
		if err != nil {
			return nil, err
		}
	}

	vnet := appservice.WebAppsClientGetVnetConnectionResponse{}
	if v.Properties != nil && v.Properties.ResourceGroup != nil && v.Name != nil && v.Properties.VirtualNetworkSubnetID != nil {
		vnet, err = webClient.GetVnetConnection(ctx, *v.Properties.ResourceGroup, *v.Name, *v.Properties.VirtualNetworkSubnetID, nil)
		if err != nil {
			return nil, err
		}
	}

	location := ""
	if v.Location != nil {
		location = *v.Location
	}

	diagnosticLogConfiguration := appservice.WebAppsClientGetDiagnosticLogsConfigurationResponse{}
	if v.Properties != nil && v.Properties.ResourceGroup != nil && v.Name != nil {
		diagnosticLogConfiguration, err = webClient.GetDiagnosticLogsConfiguration(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
		if err != nil {
			return nil, err
		}
	}

	storageAccounts := appservice.WebAppsClientListAzureStorageAccountsResponse{}
	if v.Properties != nil && v.Properties.ResourceGroup != nil && v.Name != nil {
		storageAccounts, err = webClient.ListAzureStorageAccounts(ctx, *v.Properties.ResourceGroup, *v.Name, nil)
		if err != nil {
			return nil, err
		}
	}

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AppServiceWebAppDescription{
				Site:               *v,
				SiteConfigResource: configuration.SiteConfigResource,
				SiteAuthSettings:   authSettings.SiteAuthSettings,
				SiteLogConfig:      diagnosticLogConfiguration.SiteLogsConfig,
				VnetInfo:           vnet.VnetInfoResource,
				StorageAccounts:    storageAccounts.Properties,
				ResourceGroup:      resourceGroup,
			},
		},
	}

	return &resource, nil
}

func AppServiceWebAppSlot(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := appservice.NewWebAppsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := ListAppServiceWebAppSlots(ctx, client, v)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				for _, resource := range resources {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				}
			} else {
				values = append(values, resources...)
			}
		}
	}
	return values, err
}

func ListAppServiceWebAppSlots(ctx context.Context, webClient *appservice.WebAppsClient, v *appservice.Site) ([]Resource, error) {
	pager := webClient.NewListSlotsPager(*v.Properties.ResourceGroup, *v.Name, nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, slot := range page.Value {
			resource := GetAppServiceWebAppSlot(ctx, v, slot)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetAppServiceWebAppSlot(ctx context.Context, app *appservice.Site, v *appservice.Site) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AppServiceWebAppSlotDescription{
				Site:          *v,
				AppName:       *app.Name,
				ResourceGroup: resourceGroup,
			},
		},
	}

	return &resource
}

func AppServicePlan(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := appservice.NewPlansClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

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

func GetAppServicePlan(ctx context.Context, client *appservice.PlansClient, v *appservice.Plan) (*Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	location := ""
	if v.Location != nil {
		location = *v.Location
	}

	var webApps []*appservice.Site

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
			Value: model.AppServicePlanDescription{
				Plan:          *v,
				Apps:          webApps,
				ResourceGroup: resourceGroup,
			},
		},
	}

	return &resource, nil
}

func AppContainerApps(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := appservice.NewContainerAppsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

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

func GetAppContainerApps(ctx context.Context, server *appservice.ContainerApp) *Resource {
	resourceGroupName := strings.Split(string(*server.ID), "/")[4]

	resource := Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.ContainerAppDescription{
				Server:        *server,
				ResourceGroup: resourceGroupName,
			},
		},
	}

	return &resource
}

func WebServerFarms(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := appservice.NewPlansClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

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

func GetWebServerFarm(ctx context.Context, v *appservice.Plan) *Resource {
	resourceGroupName := strings.Split(string(*v.ID), "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.WebServerFarmsDescription{
				ServerFarm:    *v,
				ResourceGroup: resourceGroupName,
			},
		},
	}
	return &resource
}
