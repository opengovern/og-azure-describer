package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func AutomationAccounts(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armautomation.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewAccountClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range result.Value {
			resource := getAutomationAccount(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
}

func getAutomationAccount(ctx context.Context, account *armautomation.Account) *Resource {
	resourceGroup := strings.Split(*account.ID, "/")[4]

	resource := Resource{
		ID:       *account.ID,
		Name:     *account.Name,
		Location: *account.Location,
		Description: JSONAllFieldsMarshaller{
			model.AutomationAccountsDescription{
				Automation:    *account,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
