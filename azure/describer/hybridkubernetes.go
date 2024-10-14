package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func HybridKubernetesConnectedCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armhybridkubernetes.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewConnectedClusterClient()

	confClientFactory, err := armkubernetesconfiguration.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	extClient := confClientFactory.NewExtensionsClient()

	pager := client.NewListBySubscriptionPager(nil)

	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getHybridKubernetesConnectedCluster(ctx, extClient, v)
			if err != nil {
				return nil, err
			}
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

func getHybridKubernetesConnectedCluster(ctx context.Context, extClient *armkubernetesconfiguration.ExtensionsClient, connectedCluster *armhybridkubernetes.ConnectedCluster) (*Resource, error) {
	resourceGroup := strings.Split(*connectedCluster.ID, "/")[4]

	pager := extClient.NewListPager(resourceGroup, "Microsoft.Kubernetes", "connectedClusters", *connectedCluster.Name, nil)
	var extensions []*armkubernetesconfiguration.Extension
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		extensions = append(extensions, page.Value...)
	}
	resource := Resource{
		ID:       *connectedCluster.ID,
		Name:     *connectedCluster.Name,
		Location: *connectedCluster.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.HybridKubernetesConnectedClusterDescription{
				ConnectedCluster:           *connectedCluster,
				ConnectedClusterExtensions: extensions,
				ResourceGroup:              resourceGroup,
			},
		},
	}
	return &resource, nil
}
