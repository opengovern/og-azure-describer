package steampipe

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/kaytu-io/kaytu-azure-describer/steampipe-plugin-azure/azure"
	"github.com/kaytu-io/kaytu-azure-describer/steampipe-plugin-azuread/azuread"
	"github.com/kaytu-io/kaytu-util/pkg/steampipe"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/context_key"
)

func buildContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, context_key.Logger, hclog.New(nil))
	return ctx
}

func AzureDescriptionToRecord(logger *zap.Logger, resource interface{}, indexName string) (map[string]*proto.Column, error) {
	return steampipe.DescriptionToRecord(logger, azure.Plugin(buildContext()), resource, indexName)
}

func AzureADDescriptionToRecord(logger *zap.Logger, resource interface{}, indexName string) (map[string]*proto.Column, error) {
	return steampipe.DescriptionToRecord(logger, azuread.Plugin(buildContext()), resource, indexName)
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
func ExtractTagsAndNames(logger *zap.Logger, plg, adPlg *plugin.Plugin, resourceType string, source interface{}) (map[string]string, string, error) {
	pluginTableName := ExtractTableName(resourceType)
	if pluginTableName == "" {
		return nil, "", fmt.Errorf("cannot find table name for resourceType: %s", resourceType)
	}

	switch steampipe.ExtractPlugin(resourceType) {
	case steampipe.SteampipePluginAzure:
		return steampipe.ExtractTagsAndNames(plg, logger, pluginTableName, resourceType, source, AzureDescriptionMap)
	case steampipe.SteampipePluginAzureAD:
		return steampipe.ExtractTagsAndNames(adPlg, logger, pluginTableName, resourceType, source, AzureDescriptionMap)
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

func ExtractResourceType(tableName string) string {
	tableName = strings.ToLower(tableName)
	return strings.ToLower(AzureReverseMap[tableName])
}

func GetResourceTypeByTableName(tableName string) string {
	return ExtractResourceType(tableName)
}
