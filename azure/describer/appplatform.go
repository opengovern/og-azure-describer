package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

//func SpringCloudService(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
//	resourcesClient := resources.NewGroupsClient(subscription)
//	resourcesClient.Authorizer = authorizer
//
//	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
//	insightsClient.Authorizer = authorizer
//
//	client := appplatform.NewServicesClient(subscription)
//	client.Authorizer = authorizer
//
//	result, err := resourcesClient.List(ctx, "", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var values []Resource
//	for {
//		for _, resourceGroup := range result.Values() {
//			if resourceGroup.Name == nil {
//				continue
//			}
//
//			res, err := client.List(ctx, *resourceGroup.Name)
//			if err != nil {
//				return nil, err
//			}
//			for {
//				for _, service := range res.Values() {
//					id := *service.ID
//					splitID := strings.Split(*service.ID, "/")
//
//					resourceGroup := splitID[4]
//					appplatformListOp, err := insightsClient.List(ctx, id)
//					if err != nil {
//						return nil, err
//					}
//					resource := Resource{
//						ID:       *service.ID,
//						Name:     *service.Name,
//						Location: *service.Location,
//						Description: JSONAllFieldsMarshaller{
//							model.SpringCloudServiceDescription{
//								ServiceResource:            service,
//								DiagnosticSettingsResource: appplatformListOp.Value,
//								ResourceGroup:              resourceGroup,
//							},
//						},
//					}
//					if stream != nil {
//						if err := (*stream)(resource); err != nil {
//							return nil, err
//						}
//					} else {
//						values = append(values, resource)
//					}
//				}
//				if !res.NotDone() {
//					break
//				}
//				err = res.NextWithContext(ctx)
//				if err != nil {
//					return nil, err
//				}
//			}
//		}
//		if !result.NotDone() {
//			break
//		}
//		err = result.NextWithContext(ctx)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return values, nil
//}

func SpringCloudService2(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armresources.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		result, err := pager.NextPage(context.Background())
		if err != nil {
			return nil, err
		}
		for _, resource := range result.Value {
			resource, err := getSpringCloudService(ctx, diagnosticClient, resource)
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
	return values, nil
}

func getSpringCloudService(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, service *armresources.GenericResourceExpanded) (*Resource, error) {
	if service.Name == nil {
		return nil, nil
	}
	splitID := strings.Split(*service.ID, "/")

	resourceGroup := splitID[4]

	var diagnosticList []armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*service.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, diagnostic := range page.Value {
			diagnosticList = append(diagnosticList, *diagnostic)
		}
	}
	resource := Resource{
		ID:       *service.ID,
		Name:     *service.Name,
		Location: *service.Location,
		Description: JSONAllFieldsMarshaller{
			model.SpringCloudServiceDescription{
				ServiceResource:            *service,
				DiagnosticSettingsResource: &diagnosticList,
				ResourceGroup:              resourceGroup,
			},
		},
	}
	return &resource, nil
}
