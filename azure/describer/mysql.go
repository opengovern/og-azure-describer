package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func MysqlServer(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	keysClient := mysql.NewServerKeysClient(subscription)
	keysClient.Authorizer = authorizer

	mysqlClient := mysql.NewConfigurationsClient(subscription)
	mysqlClient.Authorizer = authorizer

	client := mysql.NewServersClient(subscription)
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, server := range *result.Value {
		resourceGroup := strings.Split(string(*server.ID), "/")[4]
		serverName := *server.Name

		mysqlListByServerOp, err := mysqlClient.ListByServer(ctx, resourceGroup, serverName)
		if err != nil {
			return nil, err
		}

		keysListOp, err := keysClient.List(ctx, resourceGroup, serverName)
		if err != nil {
			return nil, err
		}

		var keys []mysql.ServerKey
		keys = append(keys, keysListOp.Values()...)
		for keysListOp.NotDone() {
			err = keysListOp.NextWithContext(ctx)
			if err != nil {
				return nil, err
			}
			keys = append(keys, keysListOp.Values()...)
		}

		resource := Resource{
			ID:       *server.ID,
			Name:     *server.Name,
			Location: *server.Location,
			Description: JSONAllFieldsMarshaller{
				model.MysqlServerDescription{
					Server:         server,
					Configurations: mysqlListByServerOp.Value,
					ServerKeys:     keys,
					ResourceGroup:  resourceGroup,
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
