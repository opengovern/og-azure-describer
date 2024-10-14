package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func CognitiveAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armcognitiveservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAccountsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticsClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListPager(nil)

	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			resource, err := getCognitiveAccount(ctx, diagnosticsClient, account)
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

func getCognitiveAccount(ctx context.Context, diagnosticsClient *armmonitor.DiagnosticSettingsClient, account *armcognitiveservices.Account) (*Resource, error) {
	resourceGroupName := strings.Split(string(*account.ID), "/")[4]

	var diagnosticSettings []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticsClient.NewListPager(*account.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		diagnosticSettings = append(diagnosticSettings, page.Value...)
	}
	return &Resource{
		ID: *account.ID,
		Description: JSONAllFieldsMarshaller{Value: model.CognitiveAccountDescription{
			Account:                     *account,
			DiagnosticSettingsResources: diagnosticSettings,
			ResourceGroup:               resourceGroupName,
		}},
	}, nil
}
