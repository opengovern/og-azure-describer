package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/hdinsight/mgmt/2018-06-01/hdinsight"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-04-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func HdInsightCluster(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
	insightsClient.Authorizer = authorizer

	client := hdinsight.NewClustersClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, cluster := range result.Values() {
			resourceGroup := strings.Split(*cluster.ID, "/")[4]

			hdinsightListOp, err := insightsClient.List(ctx, *cluster.ID)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:       *cluster.ID,
				Name:     *cluster.Name,
				Location: *cluster.Location,
				Description: model.HdinsightClusterDescription{
					Cluster:                     cluster,
					DiagnosticSettingsResources: hdinsightListOp.Value,
					ResourceGroup:               resourceGroup,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
