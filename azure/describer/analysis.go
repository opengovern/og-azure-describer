package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func AnalysisService(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armanalysisservices.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServersClient()

	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resourceGroupName := strings.Split(*server.ID, "/")[4]

			resource := Resource{
				ID:       *server.ID,
				Name:     *server.Name,
				Location: *server.Location,
				Description: JSONAllFieldsMarshaller{
					model.AnalysisServiceServerDescription{
						Server:        *server,
						ResourceGroup: resourceGroupName,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				return nil, nil
			}
		}
	}
	return nil, nil
}
