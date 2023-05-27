package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/hybridcompute/mgmt/hybridcompute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func HybridComputeMachine(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	hybridComputeClient := hybridcompute.NewMachineExtensionsClient(subscription)
	hybridComputeClient.Authorizer = authorizer

	client := hybridcompute.NewMachinesClient(subscription)
	client.Authorizer = authorizer

	result, err := client.ListBySubscription(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, machine := range result.Values() {
			resourceGroup := strings.Split(*machine.ID, "/")[4]

			hybridComputeListResult, err := hybridComputeClient.List(ctx, resourceGroup, *machine.Name, "")
			if err != nil {
				return nil, err
			}
			v := hybridComputeListResult.Values()
			for hybridComputeListResult.NotDone() {
				err := hybridComputeListResult.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}

				v = append(v, hybridComputeListResult.Values()...)
			}

			resource := Resource{
				ID:       *machine.ID,
				Name:     *machine.Name,
				Location: *machine.Location,
				Description: JSONAllFieldsMarshaller{
					model.HybridComputeMachineDescription{
						Machine:           machine,
						MachineExtensions: v,
						ResourceGroup:     resourceGroup,
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
