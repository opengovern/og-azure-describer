package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func MachineLearningWorkspace(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armmachinelearning.NewWorkspacesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, workspace := range page.Value {
			resource, err := getMachineLearningWorkspace(ctx, diagnosticClient, workspace)
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

func getMachineLearningWorkspace(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, workspace *armmachinelearning.Workspace) (*Resource, error) {
	resourceGroup := strings.Split(*workspace.ID, "/")[4]

	var machineLearningServicesListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*workspace.ID, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		machineLearningServicesListOp = append(machineLearningServicesListOp, page.Value...)
	}

	resource := Resource{
		ID:       *workspace.ID,
		Name:     *workspace.Name,
		Location: *workspace.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.MachineLearningWorkspaceDescription{
				Workspace:                   *workspace,
				DiagnosticSettingsResources: machineLearningServicesListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}

	return &resource, nil
}
