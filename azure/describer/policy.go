package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func PolicyAssignment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armpolicy.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAssignmentsClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetPolicyAssignment(ctx, v)
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

func GetPolicyAssignment(ctx context.Context, v *armpolicy.Assignment) *Resource {
	location := "global"
	if v.Location != nil {
		location = *v.Location
	}

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			model.PolicyAssignmentDescription{
				Assignment: *v,
			},
		},
	}

	return &resource
}
