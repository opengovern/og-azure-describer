package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func AutomationAccounts(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
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
	return values, nil
}

func getAutomationAccount(ctx context.Context, account *armautomation.Account) *Resource {
	resourceGroup := strings.Split(*account.ID, "/")[4]

	resource := Resource{
		ID:       *account.ID,
		Name:     *account.Name,
		Location: *account.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AutomationAccountsDescription{
				Automation:    *account,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}

func AutomationVariables(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armautomation.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewAccountClient()
	variablesClient := clientFactory.NewVariableClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		result, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range result.Value {
			resources, err := ListAutomationAccountVariables(ctx, variablesClient, v)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func ListAutomationAccountVariables(ctx context.Context, variablesClient *armautomation.VariableClient, account *armautomation.Account) ([]Resource, error) {
	resourceGroup := strings.Split(*account.ID, "/")[4]
	pager := variablesClient.NewListByAutomationAccountPager(resourceGroup, *account.Name, nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetAutomationVariable(ctx, account, v)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetAutomationVariable(ctx context.Context, account *armautomation.Account, v *armautomation.Variable) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:   *v.ID,
		Name: *v.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.AutomationVariablesDescription{
				Automation:    *v,
				AccountName:   *account.Name,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
