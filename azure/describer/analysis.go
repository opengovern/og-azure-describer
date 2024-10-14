package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func AnalysisService(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armanalysisservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServersClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resource := getAnalysisService(ctx, server)
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

func getAnalysisService(ctx context.Context, server *armanalysisservices.Server) *Resource {
	resourceGroupName := strings.Split(*server.ID, "/")[4]

	resource := Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.AnalysisServiceServerDescription{
				Server:        *server,
				ResourceGroup: resourceGroupName,
			},
		},
	}
	return &resource
}
