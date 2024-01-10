//go:generate go run ../../kaytu-deploy/kaytu/inventory-data/resource_types_generator.go --provider azure --output resource_types.go --index-map ../pkg/steampipe/table_index_map.go && gofmt -w -s resource_types.go  && goimports -w resource_types.go

package azure

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"sort"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"github.com/kaytu-io/kaytu-util/pkg/describe/enums"
	"github.com/kaytu-io/kaytu-util/pkg/source"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resourcegraph/mgmt/resourcegraph"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer"
	hamiltonAuthAutoRest "github.com/manicminer/hamilton-autorest/auth"
	hamiltonAuth "github.com/manicminer/hamilton/auth"
)

const AzureAuthLocation = "AZURE_AUTH_LOCATION"

type AuthType string

const (
	AuthEnv  AuthType = "ENV"
	AuthFile AuthType = "FILE"
	AuthCLI  AuthType = "CLI"
)

type ResourceDescriber interface {
	DescribeResources(context.Context, *azidentity.ClientSecretCredential, hamiltonAuth.Authorizer, []string, string, enums.DescribeTriggerType, *describer.StreamSender) ([]describer.Resource, error)
}

type ResourceDescribeFunc func(context.Context, *azidentity.ClientSecretCredential, hamiltonAuth.Authorizer, []string, string, enums.DescribeTriggerType, *describer.StreamSender) ([]describer.Resource, error)

func (fn ResourceDescribeFunc) DescribeResources(c context.Context, a *azidentity.ClientSecretCredential, ah hamiltonAuth.Authorizer, s []string, t string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
	return fn(c, a, ah, s, t, triggerType, stream)
}

type ResourceType struct {
	Connector source.Type

	ResourceName  string
	ResourceLabel string
	ServiceName   string

	Tags map[string][]string

	ListDescriber ResourceDescriber
	GetDescriber  ResourceDescriber // TODO: Change the type?

	TerraformName        []string
	TerraformServiceName string

	FastDiscovery bool
	CostDiscovery bool
	Summarize     bool
}

func (r ResourceType) GetConnector() source.Type {
	return r.Connector
}

func (r ResourceType) GetResourceName() string {
	return r.ResourceName
}

func (r ResourceType) GetResourceLabel() string {
	return r.ResourceLabel
}

func (r ResourceType) GetServiceName() string {
	return r.ServiceName
}

func (r ResourceType) GetTags() map[string][]string {
	return r.Tags
}

func (r ResourceType) GetTerraformName() []string {
	return r.TerraformName
}

func (r ResourceType) GetTerraformServiceName() string {
	return r.TerraformServiceName
}

func (r ResourceType) IsFastDiscovery() bool {
	return r.FastDiscovery
}

func (r ResourceType) IsCostDiscovery() bool {
	return r.CostDiscovery
}

func (r ResourceType) IsSummarized() bool {
	return r.Summarize
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

func ListSummarizeResourceTypes() []string {
	var list []string
	for k, v := range resourceTypes {
		if v.Summarize {
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
	resourceType = strings.ToLower(resourceType)
	for k, v := range resourceTypes {
		k := strings.ToLower(k)
		v := v
		if k == resourceType {
			return &v, nil
		}
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
	logger *zap.Logger,
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

	cred, err := azidentity.NewClientSecretCredential(cfg.TenantID, cfg.ClientID, cfg.ClientSecret, nil)
	if err != nil {
		return nil, err
	}

	resources, err := describe(ctx, logger, cred, hamiltonAuthorizer, resourceType, subscriptions, cfg.TenantID, triggerType, stream)
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

func describe(ctx context.Context, logger *zap.Logger, cred *azidentity.ClientSecretCredential, hamiltonAuth hamiltonAuth.Authorizer, resourceType string, subscriptions []string, tenantId string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
	resourceTypeObject, ok := resourceTypes[resourceType]
	if !ok {
		return nil, fmt.Errorf("unsupported resource type: %s", resourceType)
	}

	listDescriber := resourceTypeObject.ListDescriber
	if listDescriber == nil {
		listDescriber = describer.GenericResourceGraph{Table: "Resources", Type: resourceType}
	}
	ctx = describer.WithLogger(ctx, logger)

	return listDescriber.DescribeResources(ctx, cred, hamiltonAuth, subscriptions, tenantId, triggerType, stream)
}

func DescribeBySubscription(describe func(context.Context, *azidentity.ClientSecretCredential, string, *describer.StreamSender) ([]describer.Resource, error)) ResourceDescriber {
	return ResourceDescribeFunc(func(ctx context.Context, client *azidentity.ClientSecretCredential, hamiltonAuth hamiltonAuth.Authorizer, subscriptions []string, tenantId string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
		values := []describer.Resource{}
		for _, subscription := range subscriptions {
			result, err := describe(ctx, client, subscription, stream)
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
	return ResourceDescribeFunc(func(ctx context.Context, cred *azidentity.ClientSecretCredential, hamiltonAuth hamiltonAuth.Authorizer, subscription []string, tenantId string, triggerType enums.DescribeTriggerType, stream *describer.StreamSender) ([]describer.Resource, error) {
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

func GetResourceTypeByTerraform(terraformType string) string {
	for t, v := range resourceTypes {
		for _, name := range v.TerraformName {
			if name == terraformType {
				return t
			}
		}
	}
	return ""
}

func GetUnsupportedCostQuotaIds() []string {
	return []string{
		"PayAsYouGo_2014-09-01",
		"CSP_2015-05-01",
		"DreamSpark_2015-02-01",
		"AzureForStudents_2018-01-01",
		"Sponsored_2016-01-01",
		"Default_2014-09-01",
	}
}
