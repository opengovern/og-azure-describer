package describer

import (
	"context"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

func AdUsers(ctx context.Context, authorizer auth.Authorizer, tenantId string, stream *StreamSender) ([]Resource, error) {
	client := msgraph.NewUsersClient(tenantId)
	client.BaseClient.Authorizer = authorizer

	input := odata.Query{}
	input.Expand = odata.Expand{
		Relationship: "memberOf",
		Select:       []string{"id", "displayName"},
	}

	users, _, err := client.List(ctx, input)
	if err != nil {
		if strings.Contains(err.Error(), "Request_ResourceNotFound") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, user := range *users {

		resource := Resource{
			ID:       *user.ID,
			Name:     *user.DisplayName,
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdUsersDescription{
					TenantID: tenantId,
					AdUsers:  user,
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

func AdGroup(ctx context.Context, authorizer auth.Authorizer, tenantId string, stream *StreamSender) ([]Resource, error) {
	client := msgraph.NewGroupsClient(tenantId)
	client.BaseClient.Authorizer = authorizer

	input := odata.Query{}

	groups, _, err := client.List(ctx, input)
	if err != nil {
		if strings.Contains(err.Error(), "Request_ResourceNotFound") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, group := range *groups {
		resource := Resource{
			ID:       *group.ID,
			Name:     *group.DisplayName,
			Location: "global",
			Description: JSONAllFieldsMarshaller{
				Value: model.AdGroupDescription{
					TenantID: tenantId,
					AdGroup:  group,
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
