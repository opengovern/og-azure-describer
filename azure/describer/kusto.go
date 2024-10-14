package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func KustoCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armkusto.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClustersClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, kusto := range page.Value {
			resource := getKustoCluster(ctx, kusto)
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

func getKustoCluster(ctx context.Context, kusto *armkusto.Cluster) *Resource {
	resourceGroup := strings.Split(*kusto.ID, "/")[4]

	resource := Resource{
		ID:       *kusto.ID,
		Name:     *kusto.Name,
		Location: *kusto.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.KustoClusterDescription{
				Cluster:       *kusto,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
