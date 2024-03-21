package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphcore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	"github.com/microsoftgraph/msgraph-sdk-go/auditlogs"
	"github.com/microsoftgraph/msgraph-sdk-go/devices"
	"github.com/microsoftgraph/msgraph-sdk-go/directoryroles"
	"github.com/microsoftgraph/msgraph-sdk-go/domains"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/groupsettings"
	"github.com/microsoftgraph/msgraph-sdk-go/identity"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/policies"
	"github.com/microsoftgraph/msgraph-sdk-go/reports"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
	users2 "github.com/microsoftgraph/msgraph-sdk-go/users"
	"time"
)

func AdUsers(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Users().Get(ctx, &users2.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: &users2.UsersRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
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
			TenantID: tenantId,
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
					ImAddresses: user.GetImAddresses(),
					OtherMails:  user.GetOtherMails(),
					PasswordProfile: struct {
						ForceChangePasswordNextSignIn        *bool
						ForceChangePasswordNextSignInWithMfa *bool
						Password                             *string
					}{
						ForceChangePasswordNextSignIn:        user.GetPasswordProfile().GetForceChangePasswordNextSignIn(),
						ForceChangePasswordNextSignInWithMfa: user.GetPasswordProfile().GetForceChangePasswordNextSignInWithMfa(),
						Password:                             user.GetPasswordProfile().GetPassword(),
					},
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

	result, err := client.Groups().Get(ctx, &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.GroupsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
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

		var assignedLabels []struct {
			DisplayName *string
			LabelId     *string
		}
		for _, l := range group.GetAssignedLabels() {
			assignedLabels = append(assignedLabels, struct {
				DisplayName *string
				LabelId     *string
			}{
				DisplayName: l.GetDisplayName(),
				LabelId:     l.GetLabelId(),
			})
		}

		resource := Resource{
			ID:       *group.GetId(),
			Name:     *group.GetDisplayName(),
			Location: "global",
			TenantID: tenantId,
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
					AssignedLabels:                assignedLabels,
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
	result, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
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
		var orgID *string
		v := servicePrincipal.GetAppOwnerOrganizationId()
		if v != nil {
			tmp := v.String()
			orgID = &tmp
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}

		for _, kc := range servicePrincipal.GetKeyCredentials() {
			keyCredentials = append(keyCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Key                 []byte
				KeyId               string
				StartDateTime       *time.Time
				TypeEscaped         *string
				Usage               *string
			}{
				Key:                 kc.GetKey(),
				TypeEscaped:         kc.GetTypeEscaped(),
				Usage:               kc.GetUsage(),
				DisplayName:         kc.GetDisplayName(),
				CustomKeyIdentifier: kc.GetCustomKeyIdentifier(),
				KeyId:               kc.GetKeyId().String(),
				EndDateTime:         kc.GetEndDateTime(),
				StartDateTime:       kc.GetStartDateTime(),
			})
		}

		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		for _, pc := range servicePrincipal.GetPasswordCredentials() {
			passwordCredentials = append(passwordCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Hint                *string
				KeyId               string
				SecretText          *string
				StartDateTime       *time.Time
			}{
				CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
				DisplayName:         pc.GetDisplayName(),
				EndDateTime:         pc.GetEndDateTime(),
				Hint:                pc.GetHint(),
				KeyId:               pc.GetKeyId().String(),
				SecretText:          pc.GetSecretText(),
				StartDateTime:       pc.GetStartDateTime(),
			})
		}

		var ownerIds []*string
		for _, owner := range servicePrincipal.GetOwners() {
			ownerIds = append(ownerIds, owner.GetId())
		}

		var addIns []struct {
			Id          string
			TypeEscaped *string
			Properties  []struct {
				Key   *string
				Value *string
			}
		}
		for _, addIn := range servicePrincipal.GetAddIns() {
			var properties []struct {
				Key   *string
				Value *string
			}
			for _, p := range addIn.GetProperties() {
				properties = append(properties, struct {
					Key   *string
					Value *string
				}{
					Key:   p.GetKey(),
					Value: p.GetValue(),
				})
			}
			addIns = append(addIns, struct {
				Id          string
				TypeEscaped *string
				Properties  []struct {
					Key   *string
					Value *string
				}
			}{
				Id:          addIn.GetId().String(),
				TypeEscaped: addIn.GetTypeEscaped(),
				Properties:  properties,
			})
		}

		var appRoles []struct {
			AllowedMemberTypes []string
			Description        *string
			DisplayName        *string
			Id                 string
			IsEnabled          *bool
			Origin             *string
			Value              *string
		}
		for _, appRole := range servicePrincipal.GetAppRoles() {
			appRoles = append(appRoles, struct {
				AllowedMemberTypes []string
				Description        *string
				DisplayName        *string
				Id                 string
				IsEnabled          *bool
				Origin             *string
				Value              *string
			}{
				Id:                 appRole.GetId().String(),
				Description:        appRole.GetDescription(),
				DisplayName:        appRole.GetDisplayName(),
				AllowedMemberTypes: appRole.GetAllowedMemberTypes(),
				IsEnabled:          appRole.GetIsEnabled(),
				Origin:             appRole.GetOrigin(),
				Value:              appRole.GetValue(),
			})
		}

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range servicePrincipal.GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}

		resource := Resource{
			ID:       *servicePrincipal.GetId(),
			Name:     *servicePrincipal.GetDisplayName(),
			Location: "global",
			TenantID: tenantId,
			Description: JSONAllFieldsMarshaller{
				Value: model.AdServicePrincipalDescription{
					TenantID:                  tenantId,
					Id:                        servicePrincipal.GetId(),
					DisplayName:               servicePrincipal.GetDisplayName(),
					AppId:                     servicePrincipal.GetAppId(),
					AccountEnabled:            servicePrincipal.GetAccountEnabled(),
					AppDisplayName:            servicePrincipal.GetAppDisplayName(),
					AppOwnerOrganizationId:    orgID,
					AppRoleAssignmentRequired: servicePrincipal.GetAppRoleAssignmentRequired(),
					ServicePrincipalType:      servicePrincipal.GetServicePrincipalType(),
					SignInAudience:            servicePrincipal.GetSignInAudience(),
					AppDescription:            servicePrincipal.GetAppDescription(),
					Description:               servicePrincipal.GetDescription(),
					LoginUrl:                  servicePrincipal.GetLoginUrl(),
					LogoutUrl:                 servicePrincipal.GetLogoutUrl(),
					AddIns:                    addIns,
					AlternativeNames:          servicePrincipal.GetAlternativeNames(),
					AppRoles:                  appRoles,
					//Info: servicePrincipal.GetInfo(),
					KeyCredentials:             keyCredentials,
					NotificationEmailAddresses: servicePrincipal.GetNotificationEmailAddresses(),
					OwnerIds:                   ownerIds,
					PasswordCredentials:        passwordCredentials,
					Oauth2PermissionScopes:     oauth2PermissionScopes,
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
	result, err := client.Applications().Get(ctx, &applications.ApplicationsRequestBuilderGetRequestConfiguration{
		QueryParameters: &applications.ApplicationsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
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

		var oauth2PermissionScopes []struct {
			AdminConsentDescription *string
			AdminConsentDisplayName *string
			Id                      string
			IsEnabled               *bool
			Origin                  *string
			TypeEscaped             *string
			UserConsentDescription  *string
			UserConsentDisplayName  *string
		}
		for _, ps := range app.GetApi().GetOauth2PermissionScopes() {
			oauth2PermissionScopes = append(oauth2PermissionScopes, struct {
				AdminConsentDescription *string
				AdminConsentDisplayName *string
				Id                      string
				IsEnabled               *bool
				Origin                  *string
				TypeEscaped             *string
				UserConsentDescription  *string
				UserConsentDisplayName  *string
			}{
				Id:                      ps.GetId().String(),
				Origin:                  ps.GetOrigin(),
				IsEnabled:               ps.GetIsEnabled(),
				TypeEscaped:             ps.GetTypeEscaped(),
				AdminConsentDescription: ps.GetAdminConsentDescription(),
				UserConsentDescription:  ps.GetUserConsentDescription(),
				UserConsentDisplayName:  ps.GetUserConsentDisplayName(),
				AdminConsentDisplayName: ps.GetAdminConsentDisplayName(),
			})
		}

		var preAuthorizedApplications []struct {
			AppId                  *string
			DelegatedPermissionIds []string
		}
		for _, paa := range app.GetApi().GetPreAuthorizedApplications() {
			preAuthorizedApplications = append(preAuthorizedApplications, struct {
				AppId                  *string
				DelegatedPermissionIds []string
			}{
				AppId:                  paa.GetAppId(),
				DelegatedPermissionIds: paa.GetDelegatedPermissionIds(),
			})
		}

		var knownClientApplications []string
		for _, a := range app.GetApi().GetKnownClientApplications() {
			knownClientApplications = append(knownClientApplications, a.String())
		}

		var keyCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Key                 []byte
			KeyId               string
			StartDateTime       *time.Time
			TypeEscaped         *string
			Usage               *string
		}
		for _, c := range app.GetKeyCredentials() {
			keyCredentials = append(keyCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Key                 []byte
				KeyId               string
				StartDateTime       *time.Time
				TypeEscaped         *string
				Usage               *string
			}{
				TypeEscaped:         c.GetTypeEscaped(),
				Key:                 c.GetKey(),
				DisplayName:         c.GetDisplayName(),
				StartDateTime:       c.GetStartDateTime(),
				KeyId:               c.GetKeyId().String(),
				EndDateTime:         c.GetEndDateTime(),
				Usage:               c.GetUsage(),
				CustomKeyIdentifier: c.GetCustomKeyIdentifier(),
			})
		}

		var ownerIds []*string
		for _, o := range app.GetOwners() {
			ownerIds = append(ownerIds, o.GetId())
		}

		var redirectUriSettings []struct {
			Index *int32
			Uri   *string
		}
		for _, r := range app.GetWeb().GetRedirectUriSettings() {
			redirectUriSettings = append(redirectUriSettings, struct {
				Index *int32
				Uri   *string
			}{
				Uri:   r.GetUri(),
				Index: r.GetIndex(),
			})
		}

		var passwordCredentials []struct {
			CustomKeyIdentifier []byte
			DisplayName         *string
			EndDateTime         *time.Time
			Hint                *string
			KeyId               string
			SecretText          *string
			StartDateTime       *time.Time
		}
		for _, pc := range app.GetPasswordCredentials() {
			passwordCredentials = append(passwordCredentials, struct {
				CustomKeyIdentifier []byte
				DisplayName         *string
				EndDateTime         *time.Time
				Hint                *string
				KeyId               string
				SecretText          *string
				StartDateTime       *time.Time
			}{
				CustomKeyIdentifier: pc.GetCustomKeyIdentifier(),
				DisplayName:         pc.GetDisplayName(),
				EndDateTime:         pc.GetEndDateTime(),
				Hint:                pc.GetHint(),
				KeyId:               pc.GetKeyId().String(),
				SecretText:          pc.GetSecretText(),
				StartDateTime:       pc.GetStartDateTime(),
			})
		}

		resource := Resource{
			ID:       *app.GetId(),
			Name:     *app.GetDisplayName(),
			Location: "global",
			TenantID: tenantId,
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
					Api: struct {
						AcceptMappedClaims      *bool
						KnownClientApplications []string
						Oauth2PermissionScopes  []struct {
							AdminConsentDescription *string
							AdminConsentDisplayName *string
							Id                      string
							IsEnabled               *bool
							Origin                  *string
							TypeEscaped             *string
							UserConsentDescription  *string
							UserConsentDisplayName  *string
						}
						PreAuthorizedApplications []struct {
							AppId                  *string
							DelegatedPermissionIds []string
						}
						RequestedAccessTokenVersion *int32
					}{
						AcceptMappedClaims:          app.GetApi().GetAcceptMappedClaims(),
						KnownClientApplications:     knownClientApplications,
						PreAuthorizedApplications:   preAuthorizedApplications,
						Oauth2PermissionScopes:      oauth2PermissionScopes,
						RequestedAccessTokenVersion: app.GetApi().GetRequestedAccessTokenVersion(),
					},
					IdentifierUris: app.GetIdentifierUris(),
					Info: struct {
						LogoUrl             *string
						MarketingUrl        *string
						PrivacyStatementUrl *string
						SupportUrl          *string
						TermsOfServiceUrl   *string
					}{
						LogoUrl:             app.GetInfo().GetLogoUrl(),
						MarketingUrl:        app.GetInfo().GetMarketingUrl(),
						SupportUrl:          app.GetInfo().GetSupportUrl(),
						PrivacyStatementUrl: app.GetInfo().GetPrivacyStatementUrl(),
						TermsOfServiceUrl:   app.GetInfo().GetTermsOfServiceUrl(),
					},
					KeyCredentials: keyCredentials,
					OwnerIds:       ownerIds,
					ParentalControlSettings: struct {
						CountriesBlockedForMinors []string
						LegalAgeGroupRule         *string
					}{
						CountriesBlockedForMinors: app.GetParentalControlSettings().GetCountriesBlockedForMinors(),
						LegalAgeGroupRule:         app.GetParentalControlSettings().GetLegalAgeGroupRule(),
					},
					PasswordCredentials: passwordCredentials,
					Spa: struct {
						RedirectUris []string
					}{
						RedirectUris: app.GetSpa().GetRedirectUris(),
					},
					TagsSrc: app.GetTags(),
					Web: struct {
						HomePageUrl           *string
						ImplicitGrantSettings struct {
							EnableAccessTokenIssuance *bool
							EnableIdTokenIssuance     *bool
						}
						LogoutUrl           *string
						RedirectUris        []string
						RedirectUriSettings []struct {
							Index *int32
							Uri   *string
						}
					}{
						HomePageUrl:  app.GetWeb().GetHomePageUrl(),
						RedirectUris: app.GetWeb().GetRedirectUris(),
						LogoutUrl:    app.GetWeb().GetLogoutUrl(),
						ImplicitGrantSettings: struct {
							EnableAccessTokenIssuance *bool
							EnableIdTokenIssuance     *bool
						}{
							EnableAccessTokenIssuance: app.GetWeb().GetImplicitGrantSettings().GetEnableAccessTokenIssuance(),
							EnableIdTokenIssuance:     app.GetWeb().GetImplicitGrantSettings().GetEnableIdTokenIssuance(),
						},
						RedirectUriSettings: redirectUriSettings,
					},
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

	result, err := client.AuditLogs().SignIns().Get(ctx, &auditlogs.SignInsRequestBuilderGetRequestConfiguration{
		QueryParameters: &auditlogs.SignInsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
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
			TenantID: tenantId,
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
	result, err := client.Devices().Get(ctx, &devices.DevicesRequestBuilderGetRequestConfiguration{
		QueryParameters: &devices.DevicesRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
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
			TenantID: tenantId,
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

func AdDirectoryRole(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error
	result, err := client.DirectoryRoles().Get(ctx, &directoryroles.DirectoryRolesRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.DirectoryRole](result, client.GetAdapter(), models.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(role models.DirectoryRole) bool {
		var memberIds []*string
		for _, member := range role.GetMembers() {
			memberIds = append(memberIds, member.GetId())
		}

		resource := Resource{
			ID:       *role.GetId(),
			Name:     *role.GetDisplayName(),
			Location: "global",
			TenantID: tenantId,
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

func AdDirectorySetting(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.GroupSettings().Get(ctx, &groupsettings.GroupSettingsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groupsettings.GroupSettingsRequestBuilderGetQueryParameters{
			Top: aws.Int32(9999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.GroupSettingable](result, client.GetAdapter(), models.CreateGroupSettingCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(setting models.GroupSettingable) bool {
		if setting == nil {
			return true
		}

		for _, v := range setting.GetValues() {
			resource := Resource{
				ID:       *setting.GetId(),
				Name:     *setting.GetDisplayName(),
				Location: "global",
				TenantID: tenantId,
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
				if itemErr = (*stream)(resource); itemErr != nil {
					return false
				}
			} else {
				values = append(values, resource)
			}
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

func AdDirectoryAuditReport(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.AuditLogs().DirectoryAudits().Get(ctx, &auditlogs.DirectoryAuditsRequestBuilderGetRequestConfiguration{
		QueryParameters: &auditlogs.DirectoryAuditsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.DirectoryAuditable](result, client.GetAdapter(), models.CreateSignInCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(audit models.DirectoryAuditable) bool {
		if audit == nil {
			return true
		}

		var auditResult *string
		if audit.GetResult() != nil {
			tmpResult := audit.GetResult().String()
			auditResult = &tmpResult
		}

		var additionalDetails []struct {
			Key       *string
			OdataType *string
			Value     *string
		}

		for _, ad := range audit.GetAdditionalDetails() {
			additionalDetails = append(additionalDetails, struct {
				Key       *string
				OdataType *string
				Value     *string
			}{Key: ad.GetKey(), OdataType: ad.GetOdataType(), Value: ad.GetValue()})
		}

		initiatedBy := struct {
			OdataType *string
			App       struct {
				AppId                *string
				DisplayName          *string
				OdataType            *string
				ServicePrincipalId   *string
				ServicePrincipalName *string
			}
			User struct {
				Id                *string
				DisplayName       *string
				OdataType         *string
				IpAddress         *string
				UserPrincipalName *string
			}
		}{
			OdataType: audit.GetInitiatedBy().GetOdataType(),
			App: struct {
				AppId                *string
				DisplayName          *string
				OdataType            *string
				ServicePrincipalId   *string
				ServicePrincipalName *string
			}{
				AppId:                audit.GetInitiatedBy().GetApp().GetAppId(),
				DisplayName:          audit.GetInitiatedBy().GetApp().GetDisplayName(),
				OdataType:            audit.GetInitiatedBy().GetApp().GetOdataType(),
				ServicePrincipalId:   audit.GetInitiatedBy().GetApp().GetServicePrincipalId(),
				ServicePrincipalName: audit.GetInitiatedBy().GetApp().GetServicePrincipalName(),
			},
			User: struct {
				Id                *string
				DisplayName       *string
				OdataType         *string
				IpAddress         *string
				UserPrincipalName *string
			}{
				Id:                audit.GetInitiatedBy().GetUser().GetId(),
				DisplayName:       audit.GetInitiatedBy().GetUser().GetDisplayName(),
				OdataType:         audit.GetInitiatedBy().GetUser().GetOdataType(),
				IpAddress:         audit.GetInitiatedBy().GetUser().GetIpAddress(),
				UserPrincipalName: audit.GetInitiatedBy().GetUser().GetUserPrincipalName(),
			},
		}

		var targetResources []struct {
			DisplayName        *string
			GroupType          string
			Id                 *string
			ModifiedProperties []struct {
				DisplayName *string
				NewValue    *string
				OdataType   *string
				OldValue    *string
			}
			OdataType         *string
			TypeEscaped       *string
			UserPrincipalName *string
		}

		for _, tr := range audit.GetTargetResources() {
			targetResource := struct {
				DisplayName        *string
				GroupType          string
				Id                 *string
				ModifiedProperties []struct {
					DisplayName *string
					NewValue    *string
					OdataType   *string
					OldValue    *string
				}
				OdataType         *string
				TypeEscaped       *string
				UserPrincipalName *string
			}{
				DisplayName:       tr.GetDisplayName(),
				GroupType:         tr.GetGroupType().String(),
				Id:                tr.GetId(),
				OdataType:         tr.GetOdataType(),
				TypeEscaped:       tr.GetTypeEscaped(),
				UserPrincipalName: tr.GetUserPrincipalName(),
			}

			var modifiedProperties []struct {
				DisplayName *string
				NewValue    *string
				OdataType   *string
				OldValue    *string
			}

			for _, mp := range tr.GetModifiedProperties() {
				modifiedProperties = append(modifiedProperties, struct {
					DisplayName *string
					NewValue    *string
					OdataType   *string
					OldValue    *string
				}{
					DisplayName: mp.GetDisplayName(),
					NewValue:    mp.GetNewValue(),
					OdataType:   mp.GetOdataType(),
					OldValue:    mp.GetOldValue(),
				})
			}

			targetResource.ModifiedProperties = modifiedProperties

			targetResources = append(targetResources, targetResource)
		}

		resource := Resource{
			ID:       *audit.GetId(),
			Location: "global",
			TenantID: tenantId,
			Description: JSONAllFieldsMarshaller{
				Value: model.AdDirectoryAuditReportDescription{
					TenantID:            tenantId,
					Id:                  audit.GetId(),
					ActivityDateTime:    audit.GetActivityDateTime(),
					ActivityDisplayName: audit.GetActivityDisplayName(),
					Category:            audit.GetCategory(),
					CorrelationId:       audit.GetCorrelationId(),
					LoggedByService:     audit.GetLoggedByService(),
					OperationType:       audit.GetOperationType(),
					Result:              auditResult,
					ResultReason:        audit.GetResultReason(),
					AdditionalDetails:   additionalDetails,
					InitiatedBy:         initiatedBy,
					TargetResources:     targetResources,
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

func AdDomain(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.Domains().Get(ctx, &domains.DomainsRequestBuilderGetRequestConfiguration{
		QueryParameters: &domains.DomainsRequestBuilderGetQueryParameters{
			Top: aws.Int32(999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.Domainable](result, client.GetAdapter(), models.CreateDomainCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(domain models.Domainable) bool {
		if domain == nil {
			return true
		}

		resource := Resource{
			ID:       *domain.GetId(),
			Location: "global",
			TenantID: tenantId,
			Description: JSONAllFieldsMarshaller{
				Value: model.AdDomainDescription{
					TenantID:           tenantId,
					Id:                 domain.GetId(),
					AuthenticationType: domain.GetAuthenticationType(),
					IsDefault:          domain.GetIsDefault(),
					IsAdminManaged:     domain.GetIsAdminManaged(),
					IsInitial:          domain.GetIsInitial(),
					IsRoot:             domain.GetIsRoot(),
					IsVerified:         domain.GetIsVerified(),
					SupportedServices:  domain.GetSupportedServices(),
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

func AdIdentityProvider(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.Identity().IdentityProviders().Get(ctx, &identity.IdentityProvidersRequestBuilderGetRequestConfiguration{
		QueryParameters: &identity.IdentityProvidersRequestBuilderGetQueryParameters{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	pageIterator, err := msgraphcore.NewPageIterator[models.BuiltInIdentityProvider](result, client.GetAdapter(), models.CreateBuiltInIdentityProviderFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(ip models.BuiltInIdentityProvider) bool {
		clientID := ip.GetAdditionalData()["clientId"]
		clientSecret := ip.GetAdditionalData()["clientSecret"]

		resource := Resource{
			ID:       *ip.GetId(),
			Name:     *ip.GetDisplayName(),
			Location: "global",
			TenantID: tenantId,
			Description: JSONAllFieldsMarshaller{
				Value: model.AdIdentityProviderDescription{
					TenantID:     tenantId,
					Id:           ip.GetId(),
					DisplayName:  ip.GetDisplayName(),
					Type:         ip.GetOdataType(),
					ClientId:     clientID,
					ClientSecret: clientSecret,
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

func AdSecurityDefaultsPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Policies().IdentitySecurityDefaultsEnforcementPolicy().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	if result == nil {
		return values, nil
	}

	resource := Resource{
		ID:       *result.GetId(),
		Name:     *result.GetDisplayName(),
		Location: "global",
		TenantID: tenantId,
		Description: JSONAllFieldsMarshaller{
			Value: model.AdSecurityDefaultsPolicyDescription{
				TenantID:    tenantId,
				Id:          result.GetId(),
				DisplayName: result.GetDisplayName(),
				IsEnabled:   result.GetIsEnabled(),
				Description: result.GetDescription(),
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

	return values, nil
}

func AdAuthorizationPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Policies().AuthorizationPolicy().Get(ctx, &policies.AuthorizationPolicyRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	if result == nil {
		return values, nil
	}

	defaultUserRolePermissions := struct {
		AllowedToCreateApps                      *bool
		AllowedToCreateSecurityGroups            *bool
		AllowedToCreateTenants                   *bool
		AllowedToReadBitlockerKeysForOwnedDevice *bool
		AllowedToReadOtherUsers                  *bool
		OdataType                                *string
		PermissionGrantPoliciesAssigned          []string
	}{
		AllowedToCreateApps:                      result.GetDefaultUserRolePermissions().GetAllowedToCreateApps(),
		AllowedToCreateSecurityGroups:            result.GetDefaultUserRolePermissions().GetAllowedToCreateSecurityGroups(),
		AllowedToCreateTenants:                   result.GetDefaultUserRolePermissions().GetAllowedToCreateTenants(),
		AllowedToReadBitlockerKeysForOwnedDevice: result.GetDefaultUserRolePermissions().GetAllowedToReadBitlockerKeysForOwnedDevice(),
		AllowedToReadOtherUsers:                  result.GetDefaultUserRolePermissions().GetAllowedToReadOtherUsers(),
		OdataType:                                result.GetDefaultUserRolePermissions().GetOdataType(),
		PermissionGrantPoliciesAssigned:          result.GetDefaultUserRolePermissions().GetPermissionGrantPoliciesAssigned(),
	}

	resource := Resource{
		ID:       *result.GetId(),
		Name:     *result.GetDisplayName(),
		Location: "global",
		TenantID: tenantId,
		Description: JSONAllFieldsMarshaller{
			Value: model.AdAuthorizationPolicyDescription{
				TenantID:                               tenantId,
				Id:                                     result.GetId(),
				DisplayName:                            result.GetDisplayName(),
				Description:                            result.GetDescription(),
				AllowedToSignIpEmailBasedSubscriptions: result.GetAllowedToSignUpEmailBasedSubscriptions(),
				AllowedToUseSspr:                       result.GetAllowedToUseSSPR(),
				AllowedEmailVerifiedUsersToJoinOrganization: result.GetAllowEmailVerifiedUsersToJoinOrganization(),
				AllowInvitesFrom:           result.GetAllowInvitesFrom().String(),
				BlockMsolPowershell:        result.GetBlockMsolPowerShell(),
				GuestUserRoleId:            result.GetGuestUserRoleId().String(),
				DefaultUserRolePermissions: defaultUserRolePermissions,
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

	return values, nil
}

func AdConditionalAccessPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var itemErr error

	result, err := client.Identity().ConditionalAccess().Policies().Get(ctx, &identity.ConditionalAccessPoliciesRequestBuilderGetRequestConfiguration{
		QueryParameters: &identity.ConditionalAccessPoliciesRequestBuilderGetQueryParameters{
			Top: aws.Int32(9999),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	if result == nil {
		return values, nil
	}

	pageIterator, err := msgraphcore.NewPageIterator[models.ConditionalAccessPolicyable](result, client.GetAdapter(), models.CreateConditionalAccessPolicyCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(p models.ConditionalAccessPolicyable) bool {
		if p == nil {
			return true
		}

		applications := struct {
			ApplicationFilter struct {
				Mode      *string
				OdataType *string
				Rule      *string
			}
			ExcludeApplications                         []string
			IncludeApplications                         []string
			IncludeAuthenticationContextClassReferences []string
			IncludeUserActions                          []string
			OdataType                                   *string
		}{
			ApplicationFilter: struct {
				Mode      *string
				OdataType *string
				Rule      *string
			}{
				Mode:      p.GetConditions().GetApplications().GetApplicationFilter().GetRule(),
				OdataType: p.GetConditions().GetApplications().GetApplicationFilter().GetOdataType(),
				Rule:      p.GetConditions().GetApplications().GetApplicationFilter().GetRule(),
			},
			ExcludeApplications:                         p.GetConditions().GetApplications().GetExcludeApplications(),
			IncludeApplications:                         p.GetConditions().GetApplications().GetIncludeApplications(),
			IncludeAuthenticationContextClassReferences: p.GetConditions().GetApplications().GetIncludeAuthenticationContextClassReferences(),
			IncludeUserActions:                          p.GetConditions().GetApplications().GetIncludeUserActions(),
			OdataType:                                   p.GetConditions().GetApplications().GetOdataType(),
		}

		var builtInControls []string
		for _, c := range p.GetGrantControls().GetBuiltInControls() {
			builtInControls = append(builtInControls, c.String())
		}

		var clientAppTypes []string
		for _, c := range p.GetConditions().GetClientAppTypes() {
			clientAppTypes = append(builtInControls, c.String())
		}

		var excludePlatforms []string
		for _, ep := range p.GetConditions().GetPlatforms().GetExcludePlatforms() {
			excludePlatforms = append(excludePlatforms, ep.String())
		}

		var includePlatforms []string
		for _, ep := range p.GetConditions().GetPlatforms().GetIncludePlatforms() {
			includePlatforms = append(includePlatforms, ep.String())
		}

		var signInRiskLevel []string
		for _, c := range p.GetConditions().GetSignInRiskLevels() {
			signInRiskLevel = append(signInRiskLevel, c.String())
		}

		var userRiskLevel []string
		for _, c := range p.GetConditions().GetUserRiskLevels() {
			userRiskLevel = append(userRiskLevel, c.String())
		}

		resource := Resource{
			ID:       *p.GetId(),
			Name:     *p.GetDisplayName(),
			Location: "global",
			TenantID: tenantId,
			Description: JSONAllFieldsMarshaller{
				Value: model.AdConditionalAccessPolicyDescription{
					TenantID:         tenantId,
					Id:               p.GetId(),
					DisplayName:      p.GetDisplayName(),
					State:            p.GetState().String(),
					CreatedDateTime:  p.GetCreatedDateTime(),
					ModifiedDateTime: p.GetModifiedDateTime(),
					Operator:         p.GetGrantControls().GetOperator(),
					Applications:     applications,
					ApplicationEnforcedRestrictions: struct {
						IsEnabled *bool
						OdataType *string
					}{IsEnabled: p.GetSessionControls().GetApplicationEnforcedRestrictions().GetIsEnabled(), OdataType: p.GetSessionControls().GetApplicationEnforcedRestrictions().GetOdataType()},
					BuiltInControls:             builtInControls,
					ClientAppTypes:              clientAppTypes,
					CustomAuthenticationFactors: p.GetGrantControls().GetCustomAuthenticationFactors(),
					CloudAppSecurity: struct {
						CloudAppSecurityType string
						OdataType            *string
						IsEnabled            *bool
						AdditionalData       map[string]interface{}
					}{
						CloudAppSecurityType: p.GetSessionControls().GetCloudAppSecurity().GetCloudAppSecurityType().String(),
						OdataType:            p.GetSessionControls().GetCloudAppSecurity().GetOdataType(),
						IsEnabled:            p.GetSessionControls().GetCloudAppSecurity().GetIsEnabled(),
						AdditionalData:       p.GetSessionControls().GetCloudAppSecurity().GetAdditionalData(),
					},
					Locations: struct {
						ExcludeLocations []string
						IncludeLocations []string
					}{
						ExcludeLocations: p.GetConditions().GetLocations().GetExcludeLocations(),
						IncludeLocations: p.GetConditions().GetLocations().GetIncludeLocations()},
					PersistentBrowser: struct {
						OdataType      *string
						IsEnabled      *bool
						Mode           string
						AdditionalData map[string]interface{}
					}{
						OdataType:      p.GetSessionControls().GetPersistentBrowser().GetOdataType(),
						IsEnabled:      p.GetSessionControls().GetPersistentBrowser().GetIsEnabled(),
						Mode:           p.GetSessionControls().GetPersistentBrowser().GetMode().String(),
						AdditionalData: p.GetSessionControls().GetPersistentBrowser().GetAdditionalData(),
					},
					Platforms: struct {
						ExcludePlatforms []string
						IncludePlatforms []string
					}{
						ExcludePlatforms: excludePlatforms,
						IncludePlatforms: includePlatforms,
					},
					SignInFrequency: struct {
						AuthenticationType string
						FrequencyInterval  string
						TypeEscaped        string
						Value              *int32
						IsEnabled          *bool
					}{
						AuthenticationType: p.GetSessionControls().GetSignInFrequency().GetAuthenticationType().String(),
						FrequencyInterval:  p.GetSessionControls().GetSignInFrequency().GetFrequencyInterval().String(),
						TypeEscaped:        p.GetSessionControls().GetSignInFrequency().GetTypeEscaped().String(),
						Value:              p.GetSessionControls().GetSignInFrequency().GetValue(),
						IsEnabled:          p.GetSessionControls().GetSignInFrequency().GetIsEnabled(),
					},
					SignInRiskLevels: signInRiskLevel,
					TermsOfUse:       p.GetGrantControls().GetTermsOfUse(),
					Users: struct {
						ExcludeGroups []string
						IncludeGroups []string
						ExcludeUsers  []string
						IncludeUsers  []string
						ExcludeRoles  []string
						IncludeRoles  []string
					}{
						ExcludeGroups: p.GetConditions().GetUsers().GetExcludeGroups(),
						IncludeGroups: p.GetConditions().GetUsers().GetIncludeGroups(),
						ExcludeUsers:  p.GetConditions().GetUsers().GetExcludeUsers(),
						IncludeUsers:  p.GetConditions().GetUsers().GetIncludeUsers(),
						ExcludeRoles:  p.GetConditions().GetUsers().GetExcludeRoles(),
						IncludeRoles:  p.GetConditions().GetUsers().GetIncludeRoles(),
					},
					UserRiskLevel: userRiskLevel,
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

func AdAdminConsentRequestPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Policies().AdminConsentRequestPolicy().Get(ctx, &policies.AdminConsentRequestPolicyRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}
	var values []Resource
	if result == nil {
		return values, nil
	}

	var reviewers []struct {
		OdataType *string
		Query     *string
		QueryRoot *string
		QueryType *string
	}
	for _, r := range result.GetReviewers() {
		reviewers = append(reviewers, struct {
			OdataType *string
			Query     *string
			QueryRoot *string
			QueryType *string
		}{
			OdataType: r.GetOdataType(),
			Query:     r.GetQuery(),
			QueryRoot: r.GetQueryRoot(),
			QueryType: r.GetQueryType(),
		})
	}

	resource := Resource{
		ID:       *result.GetId(),
		Location: "global",
		TenantID: tenantId,
		Description: JSONAllFieldsMarshaller{
			Value: model.AdAdminConsentRequestPolicyDescription{
				TenantID:              tenantId,
				Id:                    result.GetId(),
				IsEnabled:             result.GetIsEnabled(),
				NotifyReviewers:       result.GetNotifyReviewers(),
				RemindersEnabled:      result.GetRemindersEnabled(),
				RequestDurationInDays: result.GetRequestDurationInDays(),
				Version:               result.GetVersion(),
				Reviewers:             reviewers,
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

	return values, nil
}

func AdUserRegistrationDetails(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	result, err := client.Reports().AuthenticationMethods().UserRegistrationDetails().Get(ctx, &reports.AuthenticationMethodsUserRegistrationDetailsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	var values []Resource
	var itemErr error
	pageIterator, err := msgraphcore.NewPageIterator[*models.UserRegistrationDetails](result, client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	err = pageIterator.Iterate(context.Background(), func(user *models.UserRegistrationDetails) bool {
		if user == nil {
			return true
		}
		resource := Resource{
			ID:       *user.GetId(),
			Location: "global",
			TenantID: tenantId,
			Description: JSONAllFieldsMarshaller{
				Value: model.AdUserRegistrationDetailsDescription{
					TenantID:                             tenantId,
					Id:                                   user.GetId(),
					UserPrincipalName:                    user.GetUserPrincipalName(),
					UserDisplayName:                      user.GetUserDisplayName(),
					UserType:                             user.GetUserType().String(),
					IsAdmin:                              user.GetIsAdmin(),
					IsMfaCapable:                         user.GetIsMfaCapable(),
					IsMfaRegistered:                      user.GetIsMfaRegistered(),
					IsSsprCapable:                        user.GetIsSsprCapable(),
					IsSsprRegistered:                     user.GetIsSsprRegistered(),
					IsSsprEnabled:                        user.GetIsSsprEnabled(),
					IsPasswordlessCapable:                user.GetIsPasswordlessCapable(),
					SystemPreferredAuthenticationMethods: user.GetSystemPreferredAuthenticationMethods(),
					IsSystemPreferredAuthenticationMethodEnabled:  user.GetIsSystemPreferredAuthenticationMethodEnabled(),
					LastUpdatedDateTime:                           user.GetLastUpdatedDateTime(),
					MethodsRegistered:                             user.GetMethodsRegistered(),
					UserPreferredMethodForSecondaryAuthentication: user.GetUserPreferredMethodForSecondaryAuthentication().String(),
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
