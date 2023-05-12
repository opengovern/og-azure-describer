package steampipe

import (
	"encoding/json"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"gitlab.com/keibiengine/steampipe-plugin-azure/azure"
	"gitlab.com/keibiengine/steampipe-plugin-azuread/azuread"
	"reflect"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
)

func Plugin() *plugin.Plugin {
	return azure.Plugin(buildContext())
}
func ADPlugin() *plugin.Plugin {
	return azuread.Plugin(buildContext())
}
func ExtractTagsAndNames(plg, adPlg *plugin.Plugin, resourceType string, source interface{}) (map[string]string, string, error) {
	var cells map[string]*proto.Column
	pluginProvider := ExtractPlugin(resourceType)
	pluginTableName := ExtractTableName(resourceType)
	if pluginTableName == "" {
		return nil, "", fmt.Errorf("cannot find table name for resourceType: %s", resourceType)
	}
	if pluginProvider == SteampipePluginAzure || pluginProvider == SteampipePluginAzureAD {
		if pluginProvider == SteampipePluginAzure {
			desc, err := ConvertToDescription(resourceType, source)
			if err != nil {
				return nil, "", err
			}

			cells, err = DescriptionToRecord(plg, desc, pluginTableName)
			if err != nil {
				return nil, "", err
			}
		} else {
			desc, err := ConvertToDescription(resourceType, source)
			if err != nil {
				return nil, "", err
			}

			cells, err = DescriptionToRecord(adPlg, desc, pluginTableName)
			if err != nil {
				return nil, "", err
			}
		}
	} else {
		return nil, "", fmt.Errorf("invalid provider for resource type: %s", resourceType)
	}

	tags := map[string]string{}
	var name string
	for k, v := range cells {
		if k == "title" || k == "name" {
			name = v.GetStringValue()
		}
		if k == "tags" {
			if jsonBytes := v.GetJsonValue(); jsonBytes != nil && len(jsonBytes) > 0 && string(jsonBytes) != "null" {
				var t interface{}
				err := json.Unmarshal(jsonBytes, &t)
				if err != nil {
					return nil, "", err
				}

				if tmap, ok := t.(map[string]string); ok {
					tags = tmap
				} else if t == nil {
					return tags, "", nil
				} else if tmap, ok := t.(map[string]interface{}); ok {
					for tk, tv := range tmap {
						if ts, ok := tv.(string); ok {
							tags[tk] = ts
						} else {
							return nil, "", fmt.Errorf("invalid tags value type: %s", reflect.TypeOf(tv))
						}
					}
				} else if tarr, ok := t.([]interface{}); ok {
					for _, tr := range tarr {
						if tmap, ok := tr.(map[string]string); ok {
							var key string
							for tk, tv := range tmap {
								if tk == "TagKey" {
									key = tv
								} else if tk == "TagValue" {
									tags[key] = tv
								}
							}
						} else if tmap, ok := tr.(map[string]interface{}); ok {
							var key string
							for tk, tv := range tmap {
								if ts, ok := tv.(string); ok {
									if tk == "TagKey" {
										key = ts
									} else if tk == "TagValue" {
										tags[key] = ts
									}
								} else {
									return nil, "", fmt.Errorf("invalid tags js value type: %s", reflect.TypeOf(tv))
								}
							}
						}
					}
				} else {
					fmt.Printf("invalid tag type for: %s\n", string(jsonBytes))
					return nil, "", fmt.Errorf("invalid tags type: %s", reflect.TypeOf(t))
				}
			}
		}
	}
	return tags, name, nil
}
