package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func CognitiveAccount(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armcognitiveservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAccountsClient()

	pager := client.NewListPager(nil)

	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, account := range page.Value {
			resource := getCognitiveAccount(ctx, account)
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

func getCognitiveAccount(ctx context.Context, account *armcognitiveservices.Account) *Resource {
	resourceGroupName := strings.Split(string(*account.ID), "/")[4]
	return &Resource{
		ID: *account.ID,
		Description: JSONAllFieldsMarshaller{Value: model.CognitiveAccountDescription{
			Account:       *account,
			ResourceGroup: resourceGroupName,
		}},
	}
}
