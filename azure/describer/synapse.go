package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2022-10-01-preview/insights"
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
				if strings.Contains(err.Error(), "UnsupportedOperation") {
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

func SynapseWorkspaceBigdataPools(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := synapse.NewWorkspacesClient(subscription)
	client.Authorizer = authorizer

	bpClient := synapse.NewBigDataPoolsClient(subscription)
	bpClient.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroup := strings.Split(*v.ID, "/")[4]

			wResult, err := bpClient.ListByWorkspace(ctx, resourceGroup, *v.Name)
			if err != nil {
				return nil, err
			}
			for {
				for _, bp := range wResult.Values() {

					resource := Resource{
						ID:       *v.ID,
						Name:     *v.Name,
						Location: *v.Location,
						Description: JSONAllFieldsMarshaller{
							model.SynapseWorkspaceBigdatapoolsDescription{
								Workspace:     v,
								BigDataPool:   bp,
								ResourceGroup: resourceGroup,
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
				if !wResult.NotDone() {
					break
				}
				err = wResult.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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

func SynapseWorkspaceSqlpools(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := synapse.NewWorkspacesClient(subscription)
	client.Authorizer = authorizer

	bpClient := synapse.NewSQLPoolsClient(subscription)
	bpClient.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroup := strings.Split(*v.ID, "/")[4]

			wResult, err := bpClient.ListByWorkspace(ctx, resourceGroup, *v.Name)
			if err != nil {
				if strings.Contains(err.Error(), "UnsupportedOperation") {
					continue
				}
				return nil, err
			}
			for {
				for _, bp := range wResult.Values() {

					resource := Resource{
						ID:       *v.ID,
						Name:     *v.Name,
						Location: *v.Location,
						Description: JSONAllFieldsMarshaller{
							model.SynapseWorkspaceSqlpoolsDescription{
								Workspace:     v,
								SqlPool:       bp,
								ResourceGroup: resourceGroup,
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
				if !wResult.NotDone() {
					break
				}
				err = wResult.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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
