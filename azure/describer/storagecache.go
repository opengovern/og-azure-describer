package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v2"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func HpcCache(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armstoragecache.NewCachesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	var values []Resource
	pager := client.NewListPager(nil)
	if pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetHpcCache(ctx, v)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, err
}

func GetHpcCache(ctx context.Context, v *armstoragecache.Cache) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.HpcCacheDescription{
				Cache:         *v,
				ResourceGroup: resourceGroup,
			},
		},
	}

	return &resource
}
