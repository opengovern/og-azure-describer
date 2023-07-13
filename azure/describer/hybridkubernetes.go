package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func HybridKubernetesConnectedCluster(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}
	clientFactory, err := armhybridkubernetes.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewConnectedClusterClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(nil)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getHybridKubernetesConnectedCluster(ctx, v)
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

func getHybridKubernetesConnectedCluster(ctx context.Context, connectedCluster *armhybridkubernetes.ConnectedCluster) *Resource {
	resourceGroup := strings.Split(*connectedCluster.ID, "/")[4]

	resource := Resource{
		ID:       *connectedCluster.ID,
		Name:     *connectedCluster.Name,
		Location: *connectedCluster.Location,
		Description: JSONAllFieldsMarshaller{
			model.HybridKubernetesConnectedClusterDescription{
				ConnectedCluster: *connectedCluster,
				ResourceGroup:    resourceGroup,
			},
		},
	}
	return &resource
}
