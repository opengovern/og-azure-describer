package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func APIManagement(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armapimanagement.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServiceClient()

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
		for _, apiManagement := range page.Value {
			resource, err := getAPIMangement(ctx, diagnosticClient, apiManagement)
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

func getAPIMangement(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, apiManagement *armapimanagement.ServiceResource) (*Resource, error) {
	resourceGroup := strings.Split(*apiManagement.ID, "/")[4]
	accountListOpTemp := diagnosticClient.NewListPager(*apiManagement.ID, nil)
	var op []armmonitor.DiagnosticSettingsResource
	for accountListOpTemp.More() {
		accountOpPage, err := accountListOpTemp.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, accountOp := range accountOpPage.Value {
			op = append(op, *accountOp)
		}
	}
	resource := Resource{
		ID:       *apiManagement.ID,
		Name:     *apiManagement.Name,
		Location: *apiManagement.Location,
		Description: JSONAllFieldsMarshaller{
			model.APIManagementDescription{
				APIManagement:               *apiManagement,
				DiagnosticSettingsResources: &op,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}
