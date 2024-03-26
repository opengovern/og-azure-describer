package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"github.com/aws/aws-sdk-go-v2/aws"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func RoleAssignment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armauthorization.NewRoleAssignmentsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewListForSubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, roleAssignment := range page.Value {
			resource := getRoleAssignment(ctx, roleAssignment)
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

func getRoleAssignment(ctx context.Context, v *armauthorization.RoleAssignment) *Resource {
	return &Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.RoleAssignmentDescription{
				RoleAssignment: *v,
			},
		},
	}
}

func RoleDefinition(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armauthorization.NewRoleDefinitionsClient(cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager("/subscriptions/"+subscription, nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, roleDefinition := range page.Value {
			resource := getRoleDefinition(ctx, roleDefinition)
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

func getRoleDefinition(ctx context.Context, v *armauthorization.RoleDefinition) *Resource {
	return &Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.RoleDefinitionDescription{
				RoleDefinition: *v,
			},
		},
	}
}

func PolicyDefinition(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armpolicy.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDefinitionsClient()
	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, definition := range page.Value {
			resource := getPolicyDefinition(ctx, subscription, definition)
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

func getPolicyDefinition(ctx context.Context, subscription string, definition *armpolicy.Definition) *Resource {
	akas := []string{"azure:///subscriptions/" + subscription + *definition.ID, "azure:///subscriptions/" + subscription + strings.ToLower(*definition.ID)}
	turbotData := map[string]interface{}{
		"SubscriptionId": subscription,
		"Akas":           akas,
	}

	return &Resource{
		ID:       *definition.ID,
		Name:     *definition.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.PolicyDefinitionDescription{
				Definition: *definition,
				TurboData:  turbotData,
			},
		},
	}
}

func UserEffectiveAccess(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	client, err := armauthorization.NewRoleAssignmentsClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	pager := client.NewListForSubscriptionPager(nil)
	scopes := []string{"https://graph.microsoft.com/.default"}
	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, roleAssignment := range page.Value {
			if *roleAssignment.Properties.PrincipalType == armauthorization.PrincipalTypeGroup {
				members, err := graphClient.Groups().ByGroupId(*roleAssignment.Properties.PrincipalID).TransitiveMembers().GraphUser().Get(ctx, &groups.ItemTransitiveMembersGraphUserRequestBuilderGetRequestConfiguration{
					QueryParameters: &groups.ItemTransitiveMembersGraphUserRequestBuilderGetQueryParameters{
						Top: aws.Int32(999),
					},
				})
				if err != nil {
					return nil, err
				}
				for _, m := range members.GetValue() {
					id := fmt.Sprintf("%s_%s", *m.GetId(), *roleAssignment.ID)
					resource := Resource{
						ID:       id,
						Name:     *roleAssignment.Name,
						Location: "global",
						Description: JSONAllFieldsMarshaller{
							Value: model.UserEffectiveAccessDescription{
								RoleAssignment: *roleAssignment,
								UserId:         *m.GetId(),
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
			} else if *roleAssignment.Properties.PrincipalType == armauthorization.PrincipalTypeUser {
				id := fmt.Sprintf("%s_%s", *roleAssignment.Properties.PrincipalID, *roleAssignment.ID)
				resource := Resource{
					ID:       id,
					Name:     *roleAssignment.Name,
					Location: "global",
					Description: JSONAllFieldsMarshaller{
						Value: model.UserEffectiveAccessDescription{
							RoleAssignment: *roleAssignment,
							UserId:         *roleAssignment.Properties.PrincipalID,
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
		}
	}
	return values, nil
}
