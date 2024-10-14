package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
	"github.com/opengovern/og-azure-describer/azure/model"
	"strings"
)

func BlueprintArtifact(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}

	artifactClient := clientFactory.NewArtifactsClient()
	client := clientFactory.NewBlueprintsClient()

	pager := client.NewListPager(fmt.Sprintf("/subscriptions/%s", subscription), nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, b := range page.Value {
			pager2 := artifactClient.NewListPager(fmt.Sprintf("/subscriptions/%s", subscription), *b.Name, nil)
			if err != nil {
				return nil, err
			}
			var it []armblueprint.ArtifactClassification
			for pager2.More() {
				page2, err := pager2.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				it = append(it, page2.Value...)
			}
			for _, v := range it {
				resource := getBluePrintArtifact(ctx, v)
				if stream != nil {
					if err := (*stream)(*resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, *resource)
				}
			}
		}
	}
	return values, nil
}

func getBluePrintArtifact(ctx context.Context, v armblueprint.ArtifactClassification) *Resource {
	return &Resource{
		ID:          *v.GetArtifact().ID,
		Description: JSONAllFieldsMarshaller{Value: v.GetArtifact()},
	}
}

func BlueprintBlueprint(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		return nil, err
	}
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
	resourceGroupName := strings.Split(*blueprint.ID, "/")[4]
	return &Resource{
		ID: *blueprint.ID,
		Description: JSONAllFieldsMarshaller{Value: model.BlueprintDescription{
			Blueprint:     *blueprint,
			ResourceGroup: resourceGroupName,
		}},
	}
}
