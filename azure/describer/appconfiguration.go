package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func AppConfiguration(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armappconfiguration.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewConfigurationStoresClient()

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
		for _, config := range page.Value {
			resource, err := getAppConfiguration(ctx, diagnosticClient, config)
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

func getAppConfiguration(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, config *armappconfiguration.ConfigurationStore) (*Resource, error) {
	resourceGroup := strings.Split(*config.ID, "/")[4]

	var op []armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*config.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, config := range page.Value {
			op = append(op, *config)
		}
	}
	resource := Resource{
		ID:       *config.ID,
		Name:     *config.Name,
		Location: *config.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AppConfigurationDescription{
				ConfigurationStore:          *config,
				DiagnosticSettingsResources: &op,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}
