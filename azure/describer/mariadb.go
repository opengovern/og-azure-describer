package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func MariadbServer(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armmariadb.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServersClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resource := getMariadbServer(ctx, server)
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

func getMariadbServer(ctx context.Context, server *armmariadb.Server) *Resource {
	resourceGroup := strings.Split(*server.ID, "/")[4]

	resource := Resource{
		ID:       *server.ID,
		Name:     *server.Name,
		Location: *server.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.MariadbServerDescription{
				Server:        *server,
				ResourceGroup: resourceGroup,
			},
		},
	}

	return &resource
}

func MariadbDatabases(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armmariadb.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServersClient()
	databaseClient := clientFactory.NewDatabasesClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, server := range page.Value {
			resource, err := listMariadbServerDatabases(ctx, databaseClient, server)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				for _, r := range resource {
					if err := (*stream)(r); err != nil {
						return nil, err
					}
				}
			} else {
				values = append(values, resource...)
			}
		}
	}
	return values, nil
}

func listMariadbServerDatabases(ctx context.Context, databaseClient *armmariadb.DatabasesClient, server *armmariadb.Server) ([]Resource, error) {
	resourceGroup := strings.Split(*server.ID, "/")[4]

	pager := databaseClient.NewListByServerPager(resourceGroup, *server.Name, nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, database := range page.Value {
			resource := getMariadbDatabase(ctx, server, database)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getMariadbDatabase(ctx context.Context, server *armmariadb.Server, r *armmariadb.Database) *Resource {
	resourceGroup := strings.Split(*server.ID, "/")[4]

	resource := Resource{
		ID:       *r.ID,
		Name:     *r.Name,
		Location: *server.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.MariadbDatabaseDescription{
				Database:      *r,
				Server:        *server,
				ResourceGroup: resourceGroup,
			},
		},
	}

	return &resource
}
