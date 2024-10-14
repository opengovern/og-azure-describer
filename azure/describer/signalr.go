package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func SignalrService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsignalr.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	var values []Resource
	pager := client.NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, service := range page.Value {
			resource, err := GetSignalrService(ctx, diagnosticClient, service)
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

func GetSignalrService(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, service *armsignalr.ResourceInfo) (*Resource, error) {
	resourceGroup := strings.Split(*service.ID, "/")[4]

	var signalrListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*service.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		signalrListOp = append(signalrListOp, page.Value...)
	}

	resource := Resource{
		ID:       *service.ID,
		Name:     *service.Name,
		Location: *service.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.SignalrServiceDescription{
				ResourceInfo:                *service,
				DiagnosticSettingsResources: signalrListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}
