package azure

import (
	"context"
	"encoding/json"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// column definitions for the common columns
func commonColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "cloud_environment",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCloudEnvironment,
			Description: ColumnDescriptionCloudEnvironment,
			Transform:   transform.FromValue(),
		},
		{
			Name:        "subscription_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getSubscriptionID,
			Description: ColumnDescriptionSubscription,
			Transform:   transform.FromValue(),
		},
	}
}

// append the common azure columns onto the column list
func azureColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumns()...)
}

// column definitions for the common columns
func commonKaytuColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "cloud_environment",
			Type:        proto.ColumnType_STRING,
			Description: ColumnDescriptionCloudEnvironment,
			Transform:   transform.FromField("Metadata.CloudEnvironment"),
		},
		{
			Name:        "subscription_id",
			Type:        proto.ColumnType_STRING,
			Description: ColumnDescriptionSubscription,
			Transform:   transform.FromField("Metadata.SubscriptionID"),
		},
		{
			Name:        "kaytu_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Kaytu Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.SourceID"),
		},
		{
			Name:        "kaytu_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in Kaytu.",
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "kaytu_metadata",
			Type:        proto.ColumnType_STRING,
			Description: ColumnDescriptionMetadata,
			Transform:   transform.FromField("Metadata").Transform(marshalJSON),
		},
	}
}

// append the common azure columns onto the column list
func azureKaytuColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonKaytuColumns()...)
}

func getSubscriptionID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getSubscriptionID")
	cacheKey := "getSubscriptionID"

	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(string), nil
	}

	session, err := GetNewSession(ctx, d, "MANAGEMENT")
	if err != nil {
		return nil, err
	}

	// cache subscription id for the session
	d.ConnectionManager.Cache.Set(cacheKey, session.SubscriptionID)

	return session.SubscriptionID, nil
}

func getCloudEnvironment(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCloudEnvironment")
	cacheKey := "getCloudEnvironment"

	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(string), nil
	}

	session, err := GetNewSession(ctx, d, "MANAGEMENT")
	if err != nil {
		return nil, err
	}

	// cache environment name for the session
	d.ConnectionManager.Cache.Set(cacheKey, session.CloudEnvironment)

	return session.CloudEnvironment, nil
}

func marshalJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	b, err := json.Marshal(d.Value)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}