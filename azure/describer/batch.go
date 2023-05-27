package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2020-09-01/batch"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-04-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func BatchAccount(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := insights.NewDiagnosticSettingsClient(subscription)
	client.Authorizer = authorizer

	batchAccountClient := batch.NewAccountClient(subscription)
	batchAccountClient.Authorizer = authorizer

	result, err := batchAccountClient.List(context.Background())
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range result.Values() {
			id := *account.ID
			batchListOp, err := client.List(ctx, id)
			if err != nil {
				return nil, err
			}
			splitID := strings.Split(*account.ID, "/")

			resourceGroup := splitID[4]
			resource := Resource{
				ID:       *account.ID,
				Name:     *account.Name,
				Location: *account.Location,
				Description: JSONAllFieldsMarshaller{
					model.BatchAccountDescription{
						Account:                     account,
						DiagnosticSettingsResources: batchListOp.Value,
						ResourceGroup:               resourceGroup,
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
