package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
	"github.com/opengovern/og-azure-describer/azure/model"
	"strings"
)

func DashboardGrafana(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armdashboard.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewGrafanaClient()

	pager := client.NewListPager(&armdashboard.GrafanaClientListOptions{})
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getDashboardGrafana(ctx, v)
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

func getDashboardGrafana(ctx context.Context, v *armdashboard.ManagedGrafana) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.DashboardGrafanaDescription{
				Grafana:       *v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
