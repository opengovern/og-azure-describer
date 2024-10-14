package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func BatchAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armbatch.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAccountClient()

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
		for _, account := range page.Value {
			resource, err := getBatchAccount(ctx, account, diagnosticClient)
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

func getBatchAccount(ctx context.Context, account *armbatch.Account, diagnosticClient *armmonitor.DiagnosticSettingsClient) (*Resource, error) {
	id := *account.ID
	var batchListOp []armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(id, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range page.Value {
			batchListOp = append(batchListOp, *item)
		}
	}
	splitID := strings.Split(*account.ID, "/")

	resourceGroup := splitID[4]
	resource := Resource{
		ID:       *account.ID,
		Name:     *account.Name,
		Location: *account.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.BatchAccountDescription{
				Account:                     *account,
				DiagnosticSettingsResources: &batchListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}
