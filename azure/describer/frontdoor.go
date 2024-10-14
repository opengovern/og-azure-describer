package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func FrontDoor(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armfrontdoor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewFrontDoorsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, door := range page.Value {
			resource, err := getFrontDoor(ctx, diagnosticClient, door)
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

func getFrontDoor(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, door *armfrontdoor.FrontDoor) (*Resource, error) {
	resourceGroup := strings.Split(*door.ID, "/")[4]

	pager := diagnosticClient.NewListPager(*door.ID, nil)
	var frontDoorListOp []*armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		frontDoorListOp = append(frontDoorListOp, page.Value...)
	}

	resource := Resource{
		ID:       *door.ID,
		Name:     *door.Name,
		Location: *door.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.FrontdoorDescription{
				FrontDoor:                   *door,
				DiagnosticSettingsResources: frontDoorListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}
