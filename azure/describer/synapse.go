package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-04-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/synapse/mgmt/2021-03-01/synapse"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func SynapseWorkspace(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
	insightsClient.Authorizer = authorizer

	synapseClient := synapse.NewWorkspaceManagedSQLServerVulnerabilityAssessmentsClient(subscription)
	synapseClient.Authorizer = authorizer

	client := synapse.NewWorkspacesClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, config := range result.Values() {
			resourceGroup := strings.Split(*config.ID, "/")[4]

			ignoreAssesment := false
			synapseListResult, err := synapseClient.List(ctx, resourceGroup, *config.Name)
			if err != nil {
				if !strings.Contains(err.Error(), "UnsupportedOperation") {
					ignoreAssesment = true
				} else {
					return nil, err
				}
			}

			var serverVulnerabilityAssessments []synapse.ServerVulnerabilityAssessment
			if !ignoreAssesment {
				serverVulnerabilityAssessments = append(serverVulnerabilityAssessments, synapseListResult.Values()...)

				for synapseListResult.NotDone() {
					err = synapseListResult.NextWithContext(ctx)
					if err != nil {
						return nil, err
					}
					serverVulnerabilityAssessments = append(serverVulnerabilityAssessments, synapseListResult.Values()...)
				}
			}

			synapseListOp, err := insightsClient.List(ctx, *config.ID)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:       *config.ID,
				Name:     *config.Name,
				Location: *config.Location,
				Description: JSONAllFieldsMarshaller{
					model.SynapseWorkspaceDescription{
						Workspace:                      config,
						ServerVulnerabilityAssessments: serverVulnerabilityAssessments,
						DiagnosticSettingsResources:    synapseListOp.Value,
						ResourceGroup:                  resourceGroup,
					},
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
