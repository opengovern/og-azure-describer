package azure

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/kaytu-io/kaytu-util/pkg/describe/enums"
	"github.com/kaytu-io/kaytu-util/pkg/source"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resourcegraph/mgmt/resourcegraph"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer"
	hamiltonAuthAutoRest "github.com/manicminer/hamilton-autorest/auth"
	hamiltonAuth "github.com/manicminer/hamilton/auth"
)

type ResourceDescriber interface {
	DescribeResources(context.Context, autorest.Authorizer, hamiltonAuth.Authorizer, []string, string, enums.DescribeTriggerType, *describer.StreamSender) ([]describer.Resource, error)
}

type ResourceDescribeFunc func(context.Context, autorest.Authorizer, hamiltonAuth.Authorizer, []string, string, enums.DescribeTriggerType, *describer.StreamSender) ([]describer.Resource, error)

func (fn ResourceDescribeFunc) DescribeResources(c context.Context, a autorest.Authorizer, ah hamiltonAuth.Authorizer, s []string, t string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
	return fn(c, a, ah, s, t, triggerType, stream)
}

type ResourceType struct {
	Connector source.Type

	ResourceName  string
	ResourceLabel string
	ServiceName   string

	ListDescriber ResourceDescriber
	GetDescriber  ResourceDescriber // TODO: Change the type?

	TerraformName        []string
	TerraformServiceName string

	FastDiscovery bool
}

func ListResourceTypes() []string {
	var list []string
	for k := range resourceTypes {
		list = append(list, k)
	}

	sort.Strings(list)
	return list
}

func ListFastDiscoveryResourceTypes() []string {
	var list []string
	for k, v := range resourceTypes {
		if v.FastDiscovery {
			list = append(list, k)
		}
	}

	sort.Strings(list)
	return list
}

func GetResourceType(resourceType string) (*ResourceType, error) {
	if r, ok := resourceTypes[resourceType]; ok {
		return &r, nil
	}

	return nil, fmt.Errorf("resource type %s not found", resourceType)
}

func GetResourceTypesMap() map[string]ResourceType {
	return resourceTypes
}

type ResourceDescriptionMetadata struct {
	ResourceType     string
	SubscriptionIds  []string
	CloudEnvironment string
}

type Resources struct {
	Resources []describer.Resource
	Metadata  ResourceDescriptionMetadata
}

func GetResources(
	ctx context.Context,
	resourceType string,
	triggerType enums.DescribeTriggerType,
	subscriptions []string,
	cfg AuthConfig,
	azureAuth string,
	azureAuthLoc string,
	stream *describer.StreamSender,
) (*Resources, error) {
	// Create and authorize a ResourceGraph client
	var authorizer autorest.Authorizer
	var err error
	switch v := AuthType(strings.ToUpper(azureAuth)); v {
	case AuthEnv:
		authorizer, err = NewAuthorizerFromConfig(cfg)
	case AuthFile:
		setEnvIfNotEmpty(AzureAuthLocation, azureAuthLoc)
		authorizer, err = auth.NewAuthorizerFromFile(resourcegraph.DefaultBaseURI)
	case AuthCLI:
		authorizer, err = auth.NewAuthorizerFromCLI()
	default:
		err = fmt.Errorf("invalid auth type: %s", v)
	}

	if err != nil {
		return nil, err
	}

	hamiltonAuthorizer, err := hamiltonAuthAutoRest.NewAuthorizerWrapper(authorizer)
	if err != nil {
		return nil, err
	}

	env, err := auth.GetSettingsFromEnvironment()
	if err != nil {
		return nil, err
	}

	resources, err := describe(ctx, authorizer, hamiltonAuthorizer, resourceType, subscriptions, cfg.TenantID, triggerType, stream)
	if err != nil {
		return nil, err
	}

	for i, resource := range resources {
		resources[i].Type = resourceType
		if parts := strings.Split(resources[i].ID, "/"); len(parts) > 4 {
			resources[i].ResourceGroup = strings.Split(resources[i].ID, "/")[4]
		}
		resources[i].Description = describer.JSONAllFieldsMarshaller{
			Value: resource.Description,
		}
	}

	output := &Resources{
		Resources: resources,
		Metadata: ResourceDescriptionMetadata{
			ResourceType:     resourceType,
			SubscriptionIds:  subscriptions,
			CloudEnvironment: env.Environment.Name,
		},
	}

	return output, err
}

func setEnvIfNotEmpty(env, s string) {
	if s != "" {
		err := os.Setenv(env, s)
		if err != nil {
			panic(err)
		}
	}
}

func describe(ctx context.Context, authorizer autorest.Authorizer, hamiltonAuth hamiltonAuth.Authorizer, resourceType string, subscriptions []string, tenantId string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
	resourceTypeObject, ok := resourceTypes[resourceType]
	if !ok {
		return nil, fmt.Errorf("unsupported resource type: %s", resourceType)
	}

	listDescriber := resourceTypeObject.ListDescriber
	if listDescriber == nil {
		listDescriber = describer.GenericResourceGraph{Table: "Resources", Type: resourceType}
	}

	return listDescriber.DescribeResources(ctx, authorizer, hamiltonAuth, subscriptions, tenantId, triggerType, stream)
}

func DescribeBySubscription(describe func(context.Context, autorest.Authorizer, string, *describer.StreamSender) ([]describer.Resource, error)) ResourceDescriber {
	return ResourceDescribeFunc(func(ctx context.Context, authorizer autorest.Authorizer, hamiltonAuth hamiltonAuth.Authorizer, subscriptions []string, tenantId string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
		values := []describer.Resource{}
		for _, subscription := range subscriptions {
			result, err := describe(ctx, authorizer, subscription, stream)
			if err != nil {
				return nil, err
			}

			for _, resource := range result {
				resource.SubscriptionID = subscription
			}
			values = append(values, result...)
		}

		return values, nil
	})
}

func DescribeADByTenantID(describe func(context.Context, hamiltonAuth.Authorizer, string, *describer.StreamSender) ([]describer.Resource, error)) ResourceDescriber {
	return ResourceDescribeFunc(func(ctx context.Context, authorizer autorest.Authorizer, hamiltonAuth hamiltonAuth.Authorizer, subscription []string, tenantId string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
		var values []describer.Resource
		result, err := describe(ctx, hamiltonAuth, tenantId, stream)
		if err != nil {
			return nil, err
		}

		values = append(values, result...)

		return values, nil
	})
}
