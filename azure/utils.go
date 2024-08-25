package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/authorization/mgmt/authorization"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/subscription/mgmt/subscription"
)

const (
	DefaultReaderRoleDefinitionIDTemplate        = "/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7"
	DefaultBillingReaderRoleDefinitionIDTemplate = "/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/fa23ad8b-c56e-40d8-ac0c-ce449e1d2c64"
)

func CheckSPNAccessPermission(authConf AuthConfig) error {
	authorizer, err := NewAuthorizerFromConfig(authConf)
	if err != nil {
		return err
	}
	// list subscriptions
	client := subscription.NewSubscriptionsClient()
	client.Authorizer = authorizer
	authorizer.WithAuthorization()

	_, err = client.ListComplete(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

type EntraIdExtraData struct {
	DefaultDomain *string `json:"default_domain"`
}

func CheckEntraIDPermission(authConf AuthConfig) (*EntraIdExtraData, error) {
	creds, err := azidentity.NewClientSecretCredential(authConf.TenantID, authConf.ClientID, authConf.ClientSecret, nil)
	if err != nil {
		return nil, err
	}

	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(creds, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		return nil, err
	}

	orgs, err := graphClient.Organization().Get(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	if len(orgs.GetValue()) == 0 {
		return nil, fmt.Errorf("no organization found")
	}

	org := orgs.GetValue()[0]
	tenantType := org.GetTenantType()
	if tenantType == nil || *tenantType != "AAD" {
		return nil, fmt.Errorf("organization is not AAD")
	}

	hasAADPremiumServicePlan := false
	for _, assignedPlan := range org.GetAssignedPlans() {
		capabilityStatus := assignedPlan.GetCapabilityStatus()
		if capabilityStatus == nil || *capabilityStatus != "Enabled" {
			continue
		}
		service := assignedPlan.GetService()
		if service == nil || *service != "AADPremiumService" {
			continue
		}
		hasAADPremiumServicePlan = true
		break
	}

	if !hasAADPremiumServicePlan {
		return nil, fmt.Errorf("organization does not have AAD Premium service plan")
	}

	for _, domain := range org.GetVerifiedDomains() {
		isDefault := domain.GetIsDefault()
		if isDefault != nil && *isDefault {
			return &EntraIdExtraData{
				DefaultDomain: domain.GetName(),
			}, nil
		}
	}

	return &EntraIdExtraData{}, nil
}

func CheckRole(authConf AuthConfig, subscriptionID string, roleDefinitionIDTemplate string) (bool, error) {
	if roleDefinitionIDTemplate == "" {
		return false, fmt.Errorf("roleDefinitionIDTemplate is empty")
	}
	roleDefinitionID := fmt.Sprintf(roleDefinitionIDTemplate, subscriptionID)

	authorizer, err := NewAuthorizerFromConfig(authConf)
	if err != nil {
		return false, err
	}

	client := authorization.NewRoleAssignmentsClient(subscriptionID)
	client.Authorizer = authorizer
	authorizer.WithAuthorization()

	it, err := client.ListComplete(context.TODO(), "")
	if err != nil {
		return false, err
	}

	for it.NotDone() {
		v := it.Value()

		if v.Properties.RoleDefinitionID != nil && *v.Properties.RoleDefinitionID == roleDefinitionID {
			return true, nil
		}

		if it.NotDone() {
			err := it.NextWithContext(context.TODO())
			if err != nil {
				return false, err
			}
		} else {
			break
		}
	}

	return false, nil
}
