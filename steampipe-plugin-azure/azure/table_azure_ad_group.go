package azure

import (
	"context"
	"errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAdGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_ad_group",
		Description: "[DEPRECATED] This table has been deprecated and will be removed in a future release. Please use the azuread_group table in the azuread plugin instead.",
		List: &plugin.ListConfig{
			Hydrate: listAdGroups,
		},

		Columns: []*plugin.Column{
			{
				Name:        "object_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID that identifies a group.",
				Transform:   transform.FromField("Description.AdGroup.DirectoryObject.ODataId"),
			},
			{
				Name:        "object_type",
				Description: "A string that identifies the object type.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AdGroup.DirectoryObject.ODataType"),
			},
			{
				Name:        "display_name",
				Description: "A friendly name that identifies a group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AdGroup.DisplayName")},
			{
				Name:        "mail",
				Description: "The primary email address of the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AdGroup.Mail")},
			{
				Name:        "mail_enabled",
				Description: "Indicates whether the group is mail-enabled. Must be false. This is because only pure security groups can be created using the Graph API.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AdGroup.MailEnabled")},
			{
				Name:        "mail_nickname",
				Description: "The mail alias for the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AdGroup.MailNickname")},
			{
				Name:        "deletion_timestamp",
				Description: "The time at which the directory object was deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AdGroup.DeletedDateTime").Transform(convertDateToTime),
			},
			{
				Name:        "security_enabled",
				Description: "Specifies whether the group is a security group.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AdGroup.SecurityEnabled")},
			{
				Name:        "additional_properties",
				Description: "A list of unmatched properties from the message are deserialized this collection.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AdGroup.SecurityEnabled")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform: transform.

					//// LIST FUNCTION
					FromField("Description.AdGroup.MailNickname")},
		},
	}
}

func listAdGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	err := errors.New("The azure_ad_group table has been deprecated and removed, please use azuread_group table instead.")
	return nil, err
}
