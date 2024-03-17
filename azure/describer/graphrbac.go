package describer

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	users2 "github.com/microsoftgraph/msgraph-sdk-go/users"
	"reflect"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

func AdUsers(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, err
	}

	result, err := client.Users().Get(ctx, &users2.UsersRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, user := range result.GetValue() {
		userModel, ok := user.(*models.User)
		if !ok {
			return nil, fmt.Errorf("failed to convert user to userModel, type: %s, value: %v", reflect.TypeOf(user).String(), user)
		}
		resource := Resource{
			ID:       *user.GetId(),
			Name:     *user.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdUsersDescription{
					TenantID: tenantId,
					AdUsers:  userModel,
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

	return values, nil
}

func AdGroup(ctx context.Context, cred *azidentity.ClientSecretCredential, tenantId string, stream *StreamSender) ([]Resource, error) {
	scopes := []string{"https://graph.microsoft.com/.default"}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, err
	}

	result, err := client.Groups().Get(ctx, &groups.GroupsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, group := range result.GetValue() {
		groupModel, ok := group.(*models.Group)
		if !ok {
			return nil, fmt.Errorf("failed to convert group to groupModel, type: %s, value: %v", reflect.TypeOf(group).String(), group)
		}
		resource := Resource{
			ID:       *group.GetId(),
			Name:     *group.GetDisplayName(),
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdGroupDescription{
					TenantID: tenantId,
					AdGroup:  groupModel,
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

	return values, nil
}

func AdServicePrinciple(ctx context.Context, authorizer auth.Authorizer, tenantId string, stream *StreamSender) ([]Resource, error) {
	client := msgraph.NewServicePrincipalsClient(tenantId)
	client.BaseClient.Authorizer = authorizer

	input := odata.Query{}

	servicePrincipals, _, err := client.List(ctx, input)
	if err != nil {
		if strings.Contains(err.Error(), "Request_ResourceNotFound") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, servicePrincipal := range *servicePrincipals {
		resource := Resource{
			ID:       *servicePrincipal.ID,
			Name:     *servicePrincipal.DisplayName,
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdServicePrincipalDescription{
					TenantID:           tenantId,
					AdServicePrincipal: servicePrincipal,
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

	return values, nil
}
