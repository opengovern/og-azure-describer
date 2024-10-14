package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func RedisCache(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armredis.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()

	var values []Resource
	pager := client.NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetRedisCache(ctx, v)
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

func GetRedisCache(ctx context.Context, v *armredis.ResourceInfo) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.RedisCacheDescription{
				ResourceInfo:  *v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}

func CacheRedisEnterprise(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armredisenterprise.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetCacheRedisEnterprise(ctx, v)
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

func GetCacheRedisEnterprise(ctx context.Context, v *armredisenterprise.Cluster) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.RedisEnterpriseCacheDescription{
				RedisEnterprise: *v,
				ResourceGroup:   resourceGroup,
			},
		},
	}
	return &resource
}
