package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/locks"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managementgroups"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func ManagementGroup(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := managementgroups.NewClient()
	client.Authorizer = authorizer

	result, err := client.List(ctx, "", "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, info := range result.Values() {
			group, err := client.Get(ctx, *info.Name, "children", nil, "", "")
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:   *info.ID,
				Name: *info.Name,
				Description: JSONAllFieldsMarshaller{
					model.ManagementGroupDescription{
						Group: group,
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

func ManagementLock(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := locks.NewManagementLocksClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListAtSubscriptionLevel(ctx, "")
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, lockObject := range result.Values() {
			resourceGroup := strings.Split(*lockObject.ID, "/")[4]
			resource := Resource{
				ID:       *lockObject.ID,
				Name:     *lockObject.Name,
				Location: "global",
				Description: JSONAllFieldsMarshaller{
					model.ManagementLockDescription{
						Lock:          lockObject,
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
