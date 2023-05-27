package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2019-03-01/servicefabric"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func ServiceFabricCluster(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	clusterClient := servicefabric.NewClustersClient(subscription)
	clusterClient.Authorizer = authorizer
	result, err := clusterClient.List(ctx)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for _, cluster := range *result.Value {
		resourceGroup := strings.Split(*cluster.ID, "/")[4]

		resource := Resource{
			ID:       *cluster.ID,
			Name:     *cluster.Name,
			Location: *cluster.Location,
			Description: JSONAllFieldsMarshaller{
				model.ServiceFabricClusterDescription{Cluster: cluster, ResourceGroup: resourceGroup},
			}}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}

	}
	return values, nil
}
