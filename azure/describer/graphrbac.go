package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	"github.com/microsoftgraph/msgraph-sdk-go/directoryroles"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/groupsettings"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
	users2 "github.com/microsoftgraph/msgraph-sdk-go/users"
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

func AdServicePrinciple(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	for _, servicePrincipal := range result.GetValue() {
		resource := Resource{
			ID:       *servicePrincipal.GetId(),
			Name:     *servicePrincipal.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdServicePrincipalDescription{
					TenantID:       tenantId,
					Id:             servicePrincipal.GetId(),
					DisplayName:    servicePrincipal.GetDisplayName(),
					AppId:          servicePrincipal.GetAppId(),
					AccountEnabled: servicePrincipal.GetAccountEnabled(),
					AppDisplayName: servicePrincipal.GetAppDisplayName(),
					//AppOwnerOrganizationId:    servicePrincipal.GetAppOwnerOrganizationId(),
					AppRoleAssignmentRequired: servicePrincipal.GetAppRoleAssignmentRequired(),
					ServicePrincipalType:      servicePrincipal.GetServicePrincipalType(),
					SignInAudience:            servicePrincipal.GetSignInAudience(),
					AppDescription:            servicePrincipal.GetAppDescription(),
					Description:               servicePrincipal.GetDescription(),
					LoginUrl:                  servicePrincipal.GetLoginUrl(),
					LogoutUrl:                 servicePrincipal.GetLogoutUrl(),
					AddIns:                    servicePrincipal.GetAddIns(),
					AlternativeNames:          servicePrincipal.GetAlternativeNames(),
					AppRoles:                  servicePrincipal.GetAppRoles(),
					//Info: servicePrincipal.GetInfo(),
					KeyCredentials:             servicePrincipal.GetKeyCredentials(),
					NotificationEmailAddresses: servicePrincipal.GetNotificationEmailAddresses(),
					OwnerIds:                   servicePrincipal.GetOwners(),
					PasswordCredentials:        servicePrincipal.GetPasswordCredentials(),
					Oauth2PermissionScopes:     servicePrincipal.GetOauth2PermissionScopes(),
					ReplyUrls:                  servicePrincipal.GetReplyUrls(),
					ServicePrincipalNames:      servicePrincipal.GetServicePrincipalNames(),
					TagsSrc:                    servicePrincipal.GetTags(),
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

func AdApplication(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Applications().Get(ctx, &applications.ApplicationsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	for _, app := range result.GetValue() {
		if app == nil {
			continue
		}

		resource := Resource{
			ID:       *app.GetId(),
			Name:     *app.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdApplicationDescription{
					TenantID:                  tenantId,
					DisplayName:               app.GetDisplayName(),
					Id:                        app.GetId(),
					AppId:                     app.GetAppId(),
					CreatedDateTime:           app.GetCreatedDateTime(),
					Description:               app.GetDescription(),
					Oauth2RequirePostResponse: app.GetOauth2RequirePostResponse(),
					PublisherDomain:           app.GetPublisherDomain(),
					SignInAudience:            app.GetSignInAudience(),
					Api:                       app.GetApi(),
					IdentifierUris:            app.GetIdentifierUris(),
					Info:                      app.GetInfo(),
					KeyCredentials:            app.GetKeyCredentials(),
					OwnerIds:                  app.GetOwners(),
					ParentalControlSettings:   app.GetParentalControlSettings(),
					PasswordCredentials:       app.GetPasswordCredentials(),
					Spa:                       app.GetSpa(),
					TagsSrc:                   app.GetTags(),
					Web:                       app.GetWeb(),
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

func AdDirectoryRole(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.DirectoryRoles().Get(ctx, &directoryroles.DirectoryRolesRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	for _, role := range result.GetValue() {
		if role == nil {
			continue
		}

		var memberIds []*string
		for _, member := range role.GetMembers() {
			memberIds = append(memberIds, member.GetId())
		}

		resource := Resource{
			ID:       *role.GetId(),
			Name:     *role.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdDirectoryRoleDescription{
					TenantID:       tenantId,
					DisplayName:    role.GetDisplayName(),
					Id:             role.GetId(),
					Description:    role.GetDescription(),
					RoleTemplateId: role.GetRoleTemplateId(),
					MemberIds:      memberIds,
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

func AdDirectorySetting(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.GroupSettings().Get(ctx, &groupsettings.GroupSettingsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	for _, setting := range result.GetValue() {
		if setting == nil {
			continue
		}

		for _, v := range setting.GetValues() {
			resource := Resource{
				ID:       *setting.GetId(),
				Name:     *setting.GetDisplayName(),
				Location: "global",
				Description: JSONAllFieldsMarshaller{
					Value: model.AdDirectorySettingDescription{
						TenantID:    tenantId,
						DisplayName: setting.GetDisplayName(),
						Id:          setting.GetId(),
						TemplateId:  setting.GetTemplateId(),
						Name:        v.GetName(),
						Value:       v.GetValue(),
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
	}

	return values, nil
}
