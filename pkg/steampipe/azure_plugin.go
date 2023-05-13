package steampipe

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/kaytu-io/kaytu-util/pkg/steampipe"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/context_key"
	"gitlab.com/keibiengine/steampipe-plugin-azure/azure"
	"gitlab.com/keibiengine/steampipe-plugin-azuread/azuread"
	"strings"
)

func buildContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, context_key.Logger, hclog.New(nil))
	return ctx
}

func AzureDescriptionToRecord(resource interface{}, indexName string) (map[string]*proto.Column, error) {
	return steampipe.DescriptionToRecord(azure.Plugin(buildContext()), resource, indexName)
}

func AzureADDescriptionToRecord(resource interface{}, indexName string) (map[string]*proto.Column, error) {
	return steampipe.DescriptionToRecord(azuread.Plugin(buildContext()), resource, indexName)
}

func AzureCells(indexName string) ([]string, error) {
	return steampipe.Cells(azure.Plugin(buildContext()), indexName)
}
func AzureADCells(indexName string) ([]string, error) {
	return steampipe.Cells(azuread.Plugin(buildContext()), indexName)
}

func Plugin() *plugin.Plugin {
	return azure.Plugin(buildContext())
}
func ADPlugin() *plugin.Plugin {
	return azuread.Plugin(buildContext())
}
func ExtractTagsAndNames(plg, adPlg *plugin.Plugin, resourceType string, source interface{}) (map[string]string, string, error) {
	pluginTableName := ExtractTableName(resourceType)
	if pluginTableName == "" {
		return nil, "", fmt.Errorf("cannot find table name for resourceType: %s", resourceType)
	}

	switch steampipe.ExtractPlugin(resourceType) {
	case steampipe.SteampipePluginAzure:
		return steampipe.ExtractTagsAndNames(plg, pluginTableName, resourceType, source, AzureDescriptionMap)
	case steampipe.SteampipePluginAzureAD:
		return steampipe.ExtractTagsAndNames(adPlg, pluginTableName, resourceType, source, AzureDescriptionMap)
	default:
		return nil, "", fmt.Errorf("invalid provider for resource type: %s", resourceType)
	}
}

func ExtractTableName(resourceType string) string {
	resourceType = strings.ToLower(resourceType)
	if strings.HasPrefix(resourceType, "microsoft.") {
		for k, v := range azureMap {
			if resourceType == strings.ToLower(k) {
				return v
			}
		}
	}
	return ""
}

func GetResourceTypeByTableName(tableName string) string {
	tableName = strings.ToLower(tableName)
	for k, v := range azureMap {
		if tableName == strings.ToLower(v) {
			return k
		}
	}
	return ""
}
