package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	users2 "github.com/microsoftgraph/msgraph-sdk-go/users"
	"strings"
)

func AdUsers(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Users().Get(ctx, &users2.UsersRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	var values []Resource
	for _, user := range result.GetValue() {
		resource := Resource{
			ID:       *user.GetId(),
			Name:     *user.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdUsersDescription{
					TenantID:          tenantId,
					DisplayName:       user.GetDisplayName(),
					Id:                user.GetId(),
					UserPrincipalName: user.GetUserPrincipalName(),
					AccountEnabled:    user.GetAccountEnabled(),
					UserType:          user.GetUserType(),
					GivenName:         user.GetGivenName(),
					Surname:           user.GetSurname(),
					//Filter:                          user.GetFilter(),
					OnPremisesImmutableId: user.GetOnPremisesImmutableId(),
					CreatedDateTime:       user.GetCreatedDateTime(),
					Mail:                  user.GetMail(),
					MailNickname:          user.GetMailNickname(),
					PasswordPolicies:      user.GetPasswordPolicies(),
					//RefreshTokensValidFromDateTime:  user.GetRefreshTokensValidFromDateTime(),
					SignInSessionsValidFromDateTime: user.GetSignInSessionsValidFromDateTime(),
					UsageLocation:                   user.GetUsageLocation(),
					MemberOf:                        user.GetMemberOf(),
					//AdditionalProperties:            user.GetAdditionalProperties(),
					ImAddresses:     user.GetImAddresses(),
					OtherMails:      user.GetOtherMails(),
					PasswordProfile: user.GetPasswordProfile(),
				},
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, fmt.Errorf("failed to stream due to: %v", err)
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func AdGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Groups().Get(ctx, &groups.GroupsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}

	var values []Resource
	for _, group := range result.GetValue() {
		var memberIds []*string
		for _, m := range group.GetMembers() {
			memberIds = append(memberIds, m.GetId())
		}
		var ownerIds []*string
		for _, m := range group.GetOwners() {
			ownerIds = append(ownerIds, m.GetId())
		}
		resource := Resource{
			ID:       *group.GetId(),
			Name:     *group.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdGroupDescription{
					TenantID:                      tenantId,
					DisplayName:                   group.GetDisplayName(),
					ID:                            group.GetId(),
					Description:                   group.GetDescription(),
					Classification:                group.GetClassification(),
					CreatedDateTime:               group.GetCreatedDateTime(),
					ExpirationDateTime:            group.GetExpirationDateTime(),
					IsAssignableToRole:            group.GetIsAssignableToRole(),
					IsSubscribedByMail:            group.GetIsSubscribedByMail(),
					Mail:                          group.GetMail(),
					MailEnabled:                   group.GetMailEnabled(),
					MailNickname:                  group.GetMailNickname(),
					MembershipRule:                group.GetMembershipRule(),
					MembershipRuleProcessingState: group.GetMembershipRuleProcessingState(),
					OnPremisesDomainName:          group.GetOnPremisesDomainName(),
					OnPremisesLastSyncDateTime:    group.GetOnPremisesLastSyncDateTime(),
					OnPremisesNetBiosName:         group.GetOnPremisesNetBiosName(),
					OnPremisesSamAccountName:      group.GetOnPremisesSamAccountName(),
					OnPremisesSecurityIdentifier:  group.GetOnPremisesSecurityIdentifier(),
					OnPremisesSyncEnabled:         group.GetOnPremisesSyncEnabled(),
					RenewedDateTime:               group.GetRenewedDateTime(),
					SecurityEnabled:               group.GetSecurityEnabled(),
					SecurityIdentifier:            group.GetSecurityIdentifier(),
					Visibility:                    group.GetVisibility(),
					AssignedLabels:                group.GetAssignedLabels(),
					GroupTypes:                    group.GetGroupTypes(),
					MemberIds:                     memberIds,
					OwnerIds:                      ownerIds,
					ProxyAddresses:                group.GetProxyAddresses(),
					//ResourceBehaviorOptions:       group.GetResourceBehaviorOptions(),
					//ResourceProvisioningOptions:   group.GetResourceProvisioningOptions(),
				},
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, fmt.Errorf("failed to stream due to: %v", err)
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func AdServicePrinciple(ctx context.Context, authorizer auth.Authorizer, tenantId string, stream *StreamSender) ([]Resource, error) {
	client := msgraph.NewServicePrincipalsClient(tenantId)
	client.BaseClient.Authorizer = authorizer

	input := odata.Query{}

	servicePrincipals, _, err := client.List(ctx, input)
	if err != nil {
		if strings.Contains(err.Error(), "Request_ResourceNotFound") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, servicePrincipal := range *servicePrincipals {
		resource := Resource{
			ID:       *servicePrincipal.ID,
			Name:     *servicePrincipal.DisplayName,
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdServicePrincipalDescription{
					TenantID:           tenantId,
					AdServicePrincipal: servicePrincipal,
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

	return values, nil
}
