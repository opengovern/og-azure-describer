package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func KubernetesCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armcontainerservice.NewManagedClustersClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getKubernatesCluster(ctx, v)
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

func getKubernatesCluster(ctx context.Context, v *armcontainerservice.ManagedCluster) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.KubernetesClusterDescription{
				ManagedCluster: *v,
				ResourceGroup:  resourceGroup,
			},
		},
	}
	return &resource
}

func KubernetesService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	subClient, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		return nil, err
	}
	clientFactory, err := armhybridcontainerservice.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClient()

	var values []Resource
	pager := subClient.NewListLocationsPager(subscription, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, location := range page.Value {
			services, err := listLocationKubernatesServices(ctx, client, location)
			if err != nil {
				return nil, err
			}
			values = append(values, services...)
		}
	}
	return values, nil
}

func listLocationKubernatesServices(ctx context.Context, client *armhybridcontainerservice.Client, location *armsubscriptions.Location) ([]Resource, error) {
	orchestrators, err := client.ListOrchestrators(ctx, *location.ID, nil)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for _, v := range orchestrators.Orchestrators {
		resource := getKubernatesService(ctx, location, orchestrators, v)
		values = append(values, *resource)
	}
	return values, nil
}

func getKubernatesService(ctx context.Context, location *armsubscriptions.Location, orchestrators armhybridcontainerservice.ClientListOrchestratorsResponse, v *armhybridcontainerservice.OrchestratorVersionProfile) *Resource {
	resourceGroup := strings.Split(*orchestrators.ID, "/")[4]

	resource := Resource{
		ID:       *orchestrators.ID,
		Name:     *orchestrators.Name,
		Type:     *orchestrators.Type,
		Location: *location.ID,
		Description: JSONAllFieldsMarshaller{
			Value: model.KubernetesServiceVersionDescription{
				Orchestrator:  *v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
