package describer

import (
	"context"
	"strings"

	analytics "github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account"
	"github.com/Azure/azure-sdk-for-go/services/datalake/store/mgmt/2016-11-01/account"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2022-10-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func DataLakeAnalyticsAccount(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := insights.NewDiagnosticSettingsClient(subscription)
	client.Authorizer = authorizer

	accountClient := analytics.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	result, err := accountClient.List(context.Background(), "", nil, nil, "", "", nil)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range result.Values() {
			splitID := strings.Split(*account.ID, "/")
			name := *account.Name
			resourceGroup := splitID[4]

			if name == "" || resourceGroup == "" {
				continue
			}

			accountGetOp, err := accountClient.Get(ctx, resourceGroup, name)
			if err != nil {
				return nil, err
			}

			id := *account.ID
			accountListOp, err := client.List(ctx, id)
			if err != nil {
				return nil, err
			}
			resource := Resource{
				ID:       *account.ID,
				Name:     *account.Name,
				Location: *account.Location,
				Description: JSONAllFieldsMarshaller{
					model.DataLakeAnalyticsAccountDescription{
						DataLakeAnalyticsAccount:   accountGetOp,
						DiagnosticSettingsResource: accountListOp.Value,
						ResourceGroup:              resourceGroup,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func DataLakeStore(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := insights.NewDiagnosticSettingsClient(subscription)
	client.Authorizer = authorizer

	accountClient := account.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	result, err := accountClient.List(ctx, "", nil, nil, "", "", nil)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range result.Values() {
			splitID := strings.Split(*account.ID, "/")
			name := *account.Name
			resourceGroup := splitID[4]
			if name == "" || resourceGroup == "" {
				continue
			}

			accountGetOp, err := accountClient.Get(ctx, resourceGroup, name)
			if err != nil {
				return nil, err
			}
			id := *account.ID
			accountListOp, err := client.List(ctx, id)
			if err != nil {
				return nil, err
			}
			resource := Resource{
				ID:       *account.ID,
				Name:     *account.Name,
				Location: *account.Location,
				Description: JSONAllFieldsMarshaller{
					model.DataLakeStoreDescription{
						DataLakeStoreAccount:       accountGetOp,
						DiagnosticSettingsResource: accountListOp.Value,
						ResourceGroup:              resourceGroup,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
