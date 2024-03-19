package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	"github.com/microsoftgraph/msgraph-sdk-go/auditlogs"
	"github.com/microsoftgraph/msgraph-sdk-go/devices"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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
	var itemErr error
	pageIterator, err := msgraphcore.NewPageIterator[models.Userable](result, client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(user models.Userable) bool {
		if user == nil {
			return true
		}
		resource := Resource{
			ID:       *user.GetId(),
			Name:     *user.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdUsersDescription{
					TenantID:              tenantId,
					DisplayName:           user.GetDisplayName(),
					Id:                    user.GetId(),
					UserPrincipalName:     user.GetUserPrincipalName(),
					AccountEnabled:        user.GetAccountEnabled(),
					UserType:              user.GetUserType(),
					GivenName:             user.GetGivenName(),
					Surname:               user.GetSurname(),
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
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
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
	var itemErr error
	pageIterator, err := msgraphcore.NewPageIterator[models.Groupable](result, client.GetAdapter(), models.CreateGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(group models.Groupable) bool {
		if group == nil {
			return true
		}
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
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdServicePrinciple(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []Resource
	var itemErr error
	result, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	pageIterator, err := msgraphcore.NewPageIterator[models.ServicePrincipalable](result, client.GetAdapter(), models.CreateServicePrincipalCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(servicePrincipal models.ServicePrincipalable) bool {
		if servicePrincipal == nil {
			return true
		}
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
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdApplication(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []Resource
	var itemErr error
	result, err := client.Applications().Get(ctx, &applications.ApplicationsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	pageIterator, err := msgraphcore.NewPageIterator[models.Applicationable](result, client.GetAdapter(), models.CreateApplicationCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(app models.Applicationable) bool {
		if app == nil {
			return true
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
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, itemErr
	}
	if err != nil {
		return nil, err
	}

	return values, nil
}

//

func AdSignInReport(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.AuditLogs().SignIns().Get(ctx, &auditlogs.SignInsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get sign in report: %v", err)
	}

	var values []Resource
	var itemErr error

	pageIterator, err := msgraphcore.NewPageIterator[models.SignInable](result, client.GetAdapter(), models.CreateSignInCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(report models.SignInable) bool {
		if report == nil {
			return true
		}
		resource := Resource{
			ID:       *report.GetId(),
			Name:     *report.GetId(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdSignInReportDescription{
					TenantID:                         tenantId,
					Id:                               report.GetId(),
					CreatedDateTime:                  report.GetCreatedDateTime(),
					UserDisplayName:                  report.GetUserDisplayName(),
					UserPrincipalName:                report.GetUserPrincipalName(),
					UserId:                           report.GetUserId(),
					AppId:                            report.GetAppId(),
					AppDisplayName:                   report.GetAppDisplayName(),
					IpAddress:                        report.GetIpAddress(),
					ClientAppUsed:                    report.GetClientAppUsed(),
					CorrelationId:                    report.GetCorrelationId(),
					ConditionalAccessStatus:          report.GetConditionalAccessStatus(),
					IsInteractive:                    report.GetIsInteractive(),
					RiskDetail:                       report.GetRiskDetail(),
					RiskLevelAggregated:              report.GetRiskLevelAggregated(),
					RiskLevelDuringSignIn:            report.GetRiskLevelDuringSignIn(),
					RiskState:                        report.GetRiskState(),
					ResourceDisplayName:              report.GetResourceDisplayName(),
					ResourceId:                       report.GetResourceId(),
					RiskEventTypes:                   report.GetRiskEventTypes(),
					Status:                           report.GetStatus(),
					DeviceDetail:                     report.GetDeviceDetail(),
					Location:                         report.GetLocation(),
					AppliedConditionalAccessPolicies: report.GetAppliedConditionalAccessPolicies(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		// Return true to continue the iteration
		return true
	})
	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}

func AdDevice(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	var values []Resource
	var itemErr error
	result, err := client.Devices().Get(ctx, &devices.DevicesRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get devices: %v", err)
	}
	pageIterator, err := msgraphcore.NewPageIterator[models.Deviceable](result, client.GetAdapter(), models.CreateDeviceCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(device models.Deviceable) bool {
		if device == nil {
			return true
		}
		resource := Resource{
			ID:       *device.GetId(),
			Name:     *device.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdDeviceDescription{
					TenantID:                      tenantId,
					Id:                            device.GetId(),
					DisplayName:                   device.GetDisplayName(),
					AccountEnabled:                device.GetAccountEnabled(),
					DeviceId:                      device.GetDeviceId(),
					ApproximateLastSignInDateTime: device.GetApproximateLastSignInDateTime(),
					IsCompliant:                   device.GetIsCompliant(),
					IsManaged:                     device.GetIsManaged(),
					MdmAppId:                      device.GetMdmAppId(),
					OperatingSystem:               device.GetOperatingSystem(),
					OperatingSystemVersion:        device.GetOperatingSystemVersion(),
					ProfileType:                   device.GetProfileType(),
					TrustType:                     device.GetTrustType(),
					ExtensionAttributes:           device.GetExtensions(),
					MemberOf:                      device.GetMemberOf(),
				},
			},
		}
		if stream != nil {
			if itemErr = (*stream)(resource); itemErr != nil {
				return false
			}
		} else {
			values = append(values, resource)
		}
		return true
	})
	if itemErr != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return values, nil
}
