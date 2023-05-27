package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/storagesync/mgmt/2020-03-01/storagesync"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func StorageSync(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := storagesync.NewServicesClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListBySubscription(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, storage := range *result.Value {
		resourceGroup := strings.Split(*storage.ID, "/")[4]

		resource := Resource{
			ID:       *storage.ID,
			Name:     *storage.Name,
			Location: *storage.Location,
			Description: JSONAllFieldsMarshaller{
				model.StorageSyncDescription{
					Service:       storage,
					ResourceGroup: resourceGroup,
				},
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
