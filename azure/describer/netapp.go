package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/netapp/mgmt/netapp"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func NetAppAccount(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := netapp.NewAccountsClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListBySubscription(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroupName := strings.Split(string(*v.ID), "/")[4]
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: JSONAllFieldsMarshaller{
					model.NetAppAccountDescription{
						Account:       v,
						ResourceGroup: resourceGroupName,
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

func NetAppCapacityPool(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := netapp.NewAccountsClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListBySubscription(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroupName := strings.Split(string(*v.ID), "/")[4]

			poolsClient := netapp.NewPoolsClient(subscription)
			poolsClient.Authorizer = authorizer
			poolsResult, err := poolsClient.List(ctx, resourceGroupName, *v.Name)
			if err != nil {
				return nil, err
			}

			for _, pool := range poolsResult.Values() {
				resource := Resource{
					ID:       *v.ID,
					Name:     *v.Name,
					Location: *v.Location,
					Description: JSONAllFieldsMarshaller{
						model.NetAppCapacityPoolDescription{
							CapacityPool:  pool,
							ResourceGroup: resourceGroupName,
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
			if !poolsResult.NotDone() {
				break
			}
			err = poolsResult.NextWithContext(ctx)
			if err != nil {
				return nil, err
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
