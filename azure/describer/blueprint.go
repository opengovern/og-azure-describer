package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	"strings"

	"github.com/Azure/go-autorest/autorest"
)

//func BlueprintArtifact(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
//	bps, err := blueprintBlueprint(ctx, authorizer, subscription)
//	if err != nil {
//		return nil, err
//	}
//
//	client := blueprint.NewArtifactsClient()
//	client.Authorizer = authorizer
//
//	var values []Resource
//	for _, bp := range bps {
//		it, err := client.ListComplete(ctx, fmt.Sprintf("/subscriptions/%s", subscription), *bp.Name)
//		if err != nil {
//			return nil, err
//		}
//
//		for v := it.Value(); it.NotDone(); v = it.Value() {
//			var (
//				id    string
//				value interface{}
//			)
//			if artifact, ok := v.AsArtifact(); ok {
//				id, value = *artifact.ID, artifact
//			} else if artifact, ok := v.AsTemplateArtifact(); ok {
//				id, value = *artifact.ID, artifact
//			} else if artifact, ok := v.AsPolicyAssignmentArtifact(); ok {
//				id, value = *artifact.ID, artifact
//			} else if artifact, ok := v.AsRoleAssignmentArtifact(); ok {
//				id, value = *artifact.ID, artifact
//			} else {
//				panic("unknown artifact type")
//			}
//
//			resource := Resource{
//				ID:          id,
//				Description: JSONAllFieldsMarshaller{Value: value},
//			}
//			if stream != nil {
//				if err := (*stream)(resource); err != nil {
//					return nil, err
//				}
//			} else {
//				values = append(values, resource)
//			}
//			err := it.NextWithContext(ctx)
//			if err != nil {
//				return nil, err
//			}
//		}
//	}
//
//	return values, nil
//}
//
//func blueprintBlueprint(ctx context.Context, authorizer autorest.Authorizer, subscription string) ([]blueprint.Model, error) {
//	client := blueprint.NewBlueprintsClient()
//	client.Authorizer = authorizer
//
//	it, err := client.ListComplete(ctx, fmt.Sprintf("/subscriptions/%s", subscription))
//	if err != nil {
//		return nil, err
//	}
//
//	var values []blueprint.Model
//	for v := it.Value(); it.NotDone(); v = it.Value() {
//		values = append(values, v)
//
//		err := it.NextWithContext(ctx)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return values, nil
//}

func BlueprintBlueprint(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armblueprint.NewClientFactory(cred, nil)

	client := clientFactory.NewBlueprintsClient()
	pager := client.NewListPager(fmt.Sprintf("/subscriptions/%s", subscription), nil)

	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, b := range page.Value {
			resource := getBlueprintBlueprint(ctx, b)
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

func getBlueprintBlueprint(ctx context.Context, blueprint *armblueprint.Blueprint) *Resource {
	resourceGroupName := strings.Split(string(*blueprint.ID), "/")[4]
	return &Resource{
		ID: *blueprint.ID,
		Description: JSONAllFieldsMarshaller{Value: model.BlueprintDescription{
			Blueprint:     *blueprint,
			ResourceGroup: resourceGroupName,
		}},
	}
}
