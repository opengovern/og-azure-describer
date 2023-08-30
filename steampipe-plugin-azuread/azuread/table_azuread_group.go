package azuread

import (
	"context"

	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAdGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuread_group",
		Description: "Represents an Azure AD group.",
		Get: &plugin.GetConfig{
			Hydrate: kaytu.GetAdGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound", "Invalid object identifier"}),
			},
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAdGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Invalid filter clause"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{Name: "display_name", Require: plugin.Optional},
				{Name: "mail", Require: plugin.Optional},
				{Name: "mail_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "on_premises_sync_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "security_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
			},
		},
		Columns: []*plugin.Column{
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name.", Transform: transform.FromField("Description.AdGroup.DisplayName")},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the group.", Transform: transform.FromField("Description.AdGroup.DirectoryObject.ID")},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional description for the group.", Transform:

			// Other fields
			transform.FromField("Description.AdGroup.Description")},

			{Name: "classification", Type: proto.ColumnType_STRING, Description: "Describes a classification for the group (such as low, medium or high business impact).", Transform: transform.FromField("Description.AdGroup.Classification")},
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time at which the group was created.", Transform: transform.FromField("Description.AdGroup.CreatedDateTime")},
			{Name: "expiration_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the group is set to expire.", Transform: transform.FromField("Description.AdGroup.ExpirationDateTime")},
			{Name: "is_assignable_to_role", Type: proto.ColumnType_BOOL, Description: "Indicates whether this group can be assigned to an Azure Active Directory role or not.", Transform: transform.FromField("Description.AdGroup.Classification")},
			{Name: "is_subscribed_by_mail", Type: proto.ColumnType_BOOL, Description: "Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true.", Transform: transform.FromField("Description.AdGroup.IsAssignableToRole")},
			{Name: "mail", Type: proto.ColumnType_STRING, Description: "The SMTP address for the group, for example, \"serviceadmins@contoso.onmicrosoft.com\".", Transform: transform.FromField("Description.AdGroup.Mail")},
			{Name: "mail_enabled", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group is mail-enabled.", Transform: transform.FromField("Description.AdGroup.MailEnabled")},
			{Name: "mail_nickname", Type: proto.ColumnType_STRING, Description: "The mail alias for the user.", Transform: transform.FromField("Description.AdGroup.MailNickname")},
			{Name: "membership_rule", Type: proto.ColumnType_STRING, Description: "The mail alias for the group, unique in the organization.", Transform: transform.FromField("Description.AdGroup.MembershipRule")},
			{Name: "membership_rule_processing_state", Type: proto.ColumnType_STRING, Description: "Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused.", Transform: transform.FromField("Description.AdGroup.MembershipRuleProcessingState")},
			{Name: "on_premises_domain_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises Domain name synchronized from the on-premises directory.", Transform: transform.FromField("Description.AdGroup.OnPremisesDomainName")},
			{Name: "on_premises_last_sync_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Indicates the last time at which the group was synced with the on-premises directory.", Transform: transform.FromField("Description.AdGroup.OnPremisesLastSyncDateTime")},
			{Name: "on_premises_net_bios_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises NetBiosName synchronized from the on-premises directory.", Transform: transform.FromField("Description.AdGroup.OnPremisesNetBiosName")},
			{Name: "on_premises_sam_account_name", Type: proto.ColumnType_STRING, Description: "Contains the on-premises SAM account name synchronized from the on-premises directory.", Transform: transform.FromField("Description.AdGroup.OnPremisesSamAccountName")},
			{Name: "on_premises_security_identifier", Type: proto.ColumnType_STRING, Description: "Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud.", Transform: transform.FromField("Description.AdGroup.OnPremisesSecurityIdentifier")},
			{Name: "on_premises_sync_enabled", Type: proto.ColumnType_BOOL, Description: "True if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default).", Transform: transform.FromField("Description.AdGroup.OnPremisesSyncEnabled")},
			{Name: "renewed_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action.", Transform: transform.FromField("Description.AdGroup.RenewedDateTime")},
			{Name: "security_enabled", Type: proto.ColumnType_BOOL, Description: "Specifies whether the group is a security group.", Transform: transform.FromField("Description.AdGroup.SecurityEnabled")},
			{Name: "security_identifier", Type: proto.ColumnType_STRING, Description: "Security identifier of the group, used in Windows scenarios.", Transform: transform.FromField("Description.AdGroup.SecurityIdentifier")},
			{Name: "visibility", Type: proto.ColumnType_STRING, Description: "Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership.", Transform:

			// JSON fields
			transform.FromField("Description.AdGroup.Visibility")},

			{Name: "assigned_labels", Type: proto.ColumnType_JSON, Description: "The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group.", Transform: transform.FromField("Description.AdGroup.AssignedLabels")},
			{Name: "group_types", Type: proto.ColumnType_JSON, Description: "Specifies the group type and its membership. If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see [groups overview](https://docs.microsoft.com/en-us/graph/api/resources/groups-overview?view=graph-rest-1.0).", Transform: transform.FromField("Description.AdGroup.GroupTypes")},
			{Name: "member_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.AdGroup.Members"), Description: "Id of Users and groups that are members of this group."},
			{Name: "owner_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.AdGroup.Owners"), Description: "Id od the owners of the group. The owners are a set of non-admin users who are allowed to modify this object."},
			{Name: "proxy_addresses", Type: proto.ColumnType_JSON, Description: "Email addresses for the group that direct to the same group mailbox. For example: [\"SMTP: bob@contoso.com\", \"smtp: bob@sales.contoso.com\"]. The any operator is required to filter expressions on multi-valued properties.", Transform: transform.FromField("Description.AdGroup.ProxyAddresses")},
			{Name: "resource_behavior_options", Type: proto.ColumnType_JSON, Description: "Specifies the group behaviors that can be set for a Microsoft 365 group during creation. Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeNewGroupMembers, WelcomeEmailDisabled.", Transform: transform.FromField("Description.AdGroup.ResourceBehaviorOptions")},
			{Name: "resource_provisioning_options", Type: proto.ColumnType_JSON, Description: "Specifies the group resources that are provisioned as part of Microsoft 365 group creation, that are not normally part of default group creation. Possible value is Team.", Transform: transform.FromField("Description.AdGroup.ResourceProvisioningOptions")},

			{Name: "tags", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTags, Transform: transform.From(adGroupTags)},
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.From(adGroupTitle)},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.

				//// TRANSFORM FUNCTIONS
				FromField("Description.AdGroup.ResourceProvisioningOptions")},
		},
	}
}

func adGroupTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	group := d.HydrateItem.(kaytu.AdGroup).Description.AdGroup

	if group.AssignedLabels == nil {
		return nil, nil
	}

	assignedLabels := *group.AssignedLabels
	if len(assignedLabels) == 0 {
		return nil, nil
	}

	var tags = map[*string]*string{}
	for _, i := range assignedLabels {
		tags[i.LabelId] = i.DisplayName
	}

	return tags, nil
}

func adGroupTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.AdGroup).Description.AdGroup

	title := data.DisplayName
	if title == nil {
		title = data.DirectoryObject.ID
	}

	return title, nil
}
