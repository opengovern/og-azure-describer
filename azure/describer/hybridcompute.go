package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func HybridComputeMachine(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armhybridcompute.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewMachinesClient()
	extentionClient := clientFactory.NewMachineExtensionsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getHybridComputeMachine(ctx, extentionClient, v)
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

func getHybridComputeMachine(ctx context.Context, extentionClient *armhybridcompute.MachineExtensionsClient, machine *armhybridcompute.Machine) (*Resource, error) {
	resourceGroup := strings.Split(*machine.ID, "/")[4]

	var hybridComputeListResult []*armhybridcompute.MachineExtension
	pager := extentionClient.NewListPager(resourceGroup, *machine.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		hybridComputeListResult = append(hybridComputeListResult, page.Value...)
	}

	resource := Resource{
		ID:       *machine.ID,
		Name:     *machine.Name,
		Location: *machine.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.HybridComputeMachineDescription{
				Machine:           *machine,
				MachineExtensions: hybridComputeListResult,
				ResourceGroup:     resourceGroup,
			},
		},
	}

	return &resource, nil
}
