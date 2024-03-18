package azuread

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureAdUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuread_user",
		Description: "Represents an Azure AD user account.",
		Get: &plugin.GetConfig{
			Hydrate:    kaytu.GetAdUsers,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAdUsers,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{
					Name: "id", Require: plugin.Optional},
				{
					Name: "user_principal_name", Require: plugin.Optional},
				{
					Name: "filter", Require: plugin.Optional},

				// Other fields for filtering OData
				{
					Name: "user_type", Require: plugin.Optional},
				{
					Name: "account_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{
					Name: "display_name", Require: plugin.Optional},
				{
					Name: "surname", Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name.",

				Transform: transform.FromField("Description.DisplayName")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier for the user. Should be treated as an opaque identifier.",
				Transform:   transform.FromField("Description.ID"),
			},
			{
				Name:        "user_principal_name",
				Type:        proto.ColumnType_STRING,
				Description: "Principal email of the active directory user.",

				Transform: transform.FromField("Description.UserPrincipalName")},
			{
				Name:        "account_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "True if the account is enabled; otherwise, false.",

				Transform: transform.FromField("Description.AccountEnabled")},
			{
				Name:        "user_type",
				Type:        proto.ColumnType_STRING,
				Description: "A string value that can be used to classify user types in your directory.",

				Transform: transform.FromField("Description.UserType")},
			{
				Name:        "given_name",
				Type:        proto.ColumnType_STRING,
				Description: "The given name (first name) of the user.",

				Transform: transform.FromField("Description.GivenName")},
			{
				Name:        "surname",
				Type:        proto.ColumnType_STRING,
				Description: "Family name or last name of the active directory user.",

				Transform: transform.FromField("Description.Surname")},
			{
				Name:        "filter",
				Type:        proto.ColumnType_STRING,
				Description: "Odata query to search for resources.",
				Transform:   transform.FromQual("filter"),
			},

			// Other fields
			{
				Name:        "on_premises_immutable_id",
				Type:        proto.ColumnType_STRING,
				Description: "Used to associate an on-premises Active Directory user account with their Azure AD user object.",

				Transform: transform.FromField("Description.OnPremisesImmutableId")},
			{
				Name:        "created_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the user was created.",

				Transform: transform.FromField("Description.CreatedDateTime")},
			{
				Name:        "mail",
				Type:        proto.ColumnType_STRING,
				Description: "The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.",

				Transform: transform.FromField("Description.Mail")},
			{
				Name:        "mail_nickname",
				Type:        proto.ColumnType_STRING,
				Description: "The mail alias for the user.",

				Transform: transform.FromField("Description.MailNickname")},
			{
				Name:        "password_policies",
				Type:        proto.ColumnType_STRING,
				Description: "Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.",

				Transform: transform.FromField("Description.PasswordPolicies")},
			{
				Name:        "refresh_tokens_valid_from_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).",

				Transform: transform.FromField("Description.RefreshTokensValidFromDateTime")},
			{
				Name:        "sign_in_sessions_valid_from_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).",

				Transform: transform.FromField("Description.SignInSessionsValidFromDateTime")},
			{
				Name:        "usage_location",
				Type:        proto.ColumnType_STRING,
				Description: "A two letter country code (ISO standard 3166), required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.",
				Transform:   transform.FromField("Description.UsageLocation")},

			{
				Name:        "member_of",
				Type:        proto.ColumnType_JSON,
				Description: "A list the groups and directory roles that the user is a direct member of.",

				Transform: transform.FromField("Description.MemberOf")},
			{
				Name:        "additional_properties",
				Type:        proto.ColumnType_JSON,
				Description: "A list of unmatched properties from the message are deserialized this collection.",
				Transform:   transform.FromField("Description.AdditionalProperties")},
			{
				Name:        "im_addresses",
				Type:        proto.ColumnType_JSON,
				Description: "The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.",

				Transform: transform.FromField("Description.ImAddresses")},
			{
				Name:        "other_mails",
				Type:        proto.ColumnType_JSON,
				Description: "A list of additional email addresses for the user.",

				Transform: transform.FromField("Description.OtherMails")},
			{
				Name:        "password_profile",
				Type:        proto.ColumnType_JSON,
				Description: "Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created.",

				Transform:
				// Standard columns
				transform.FromField("Description.PasswordProfile")},

			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,

				Transform: transform.FromField("Description.DisplayName")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
			{
				Name:        "metadata",
				Description: "Metadata of the Azure resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata").Transform(marshalJSON),
			},
			{
				Name:        "kaytu_account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Kaytu Account ID in which the resource is located.",
				Transform:   transform.FromField("Metadata.SourceID")},
			{
				Name:        "kaytu_resource_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the resource in Kaytu.",
				Transform:   transform.FromField("ID")},
		},
	}
}

func adUserTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*ADUserInfo)
	if data == nil {
		return nil, nil
	}

	title := data.GetDisplayName()
	if title == nil {
		title = data.GetUserPrincipalName()
	}

	return title, nil
}

func buildQueryFilter(equalQuals plugin.KeyColumnEqualsQualMap) []string {
	filters := []string{}

	filterQuals := map[string]string{
		"display_name":             "string",
		"id":                       "string",
		"surname":                  "string",
		"user_principal_name":      "string",
		"user_type":                "string",
		"account_enabled":          "bool",
		"mail_enabled":             "bool",
		"security_enabled":         "bool",
		"on_premises_sync_enabled": "bool",
	}

	for qual, qualType := range filterQuals {
		switch qualType {
		case "string":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq '%s'", strcase.ToCamel(qual), equalQuals[qual].GetStringValue()))
			}
		case "bool":
			if equalQuals[qual] != nil {
				filters = append(filters, fmt.Sprintf("%s eq %t", strcase.ToCamel(qual), equalQuals[qual].GetBoolValue()))
			}
		}
	}

	return filters
}

func buildBoolNEFilter(quals plugin.KeyColumnQualMap) []string {
	filters := []string{}

	filterQuals := []string{
		"account_enabled",
		"mail_enabled",
		"on_premises_sync_enabled",
		"security_enabled",
	}

	for _, qual := range filterQuals {
		if quals[qual] != nil {
			for _, q := range quals[qual].Quals {
				value := q.Value.GetBoolValue()
				if q.Operator == "<>" {
					filters = append(filters, fmt.Sprintf("%s eq %t", strcase.ToCamel(qual), !value))
					break
				}
			}
		}
	}

	return filters
}

func marshalJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	b, err := json.Marshal(d.Value)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
