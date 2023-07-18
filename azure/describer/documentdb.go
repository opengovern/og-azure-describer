package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"strings"

	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func DocumentDBSQLDatabase(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSQLResourcesClient()

	var values []Resource
	for _, rg := range rgs {
		accounts, err := documentDBDatabaseAccounts(ctx, cred, subscription, *rg.Name)
		if err != nil {
			return nil, err
		}

		for _, account := range accounts {

			pager := client.NewListSQLDatabasesPager(*rg.Name, *account.Name, nil)
			var it []*armcosmos.SQLDatabaseGetResults
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, v := range page.Value {
					it = append(it, v)
				}
			}
			for _, v := range it {
				resource := getDocumentDBSQLDatabase(ctx, v, account, rg)
				if stream != nil {
					if err := (*stream)(*resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, *resource)
				}
			}
		}
	}

	return values, nil
}

func getDocumentDBSQLDatabase(ctx context.Context, v *armcosmos.SQLDatabaseGetResults, account *armcosmos.DatabaseAccountGetResults, rg armresources.ResourceGroup) *Resource {
	location := "global"
	if v.Location != nil {
		location = *v.Location
	}

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			model.CosmosdbSqlDatabaseDescription{
				Account:       *account,
				SqlDatabase:   *v,
				ResourceGroup: *rg.Name,
			},
		},
	}
	return &resource
}

func DocumentDBMongoDatabase(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewMongoDBResourcesClient()

	var values []Resource
	for _, rg := range rgs {
		accounts, err := documentDBDatabaseAccounts(ctx, cred, subscription, *rg.Name)
		if err != nil {
			return nil, err
		}

		for _, account := range accounts {
			pager := client.NewListMongoDBDatabasesPager(*rg.Name, *account.Name, nil)
			var it []*armcosmos.MongoDBDatabaseGetResults
			for pager.More() {
				page, err := pager.NextPage(ctx)
				if err != nil {
					return nil, err
				}
				for _, v := range page.Value {
					it = append(it, v)
				}
			}
			for _, v := range it {
				resource := getDocumentDBMongoDatabase(ctx, v, account, rg)
				if stream != nil {
					if err := (*stream)(*resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, *resource)
				}
			}
		}
	}
	return values, nil
}

func getDocumentDBMongoDatabase(ctx context.Context, v *armcosmos.MongoDBDatabaseGetResults, account *armcosmos.DatabaseAccountGetResults, rg armresources.ResourceGroup) *Resource {
	location := ""
	if v.Location != nil {
		location = *v.Location
	}

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			model.CosmosdbMongoDatabaseDescription{
				Account:       *account,
				MongoDatabase: *v,
				ResourceGroup: *rg.Name,
			},
		},
	}
	return &resource
}

func DocumentDBCassandraCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewCassandraClustersClient()

	var values []Resource
	pager := client.NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getDocumentDBCassandraCluster(ctx, v)
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

func getDocumentDBCassandraCluster(ctx context.Context, v *armcosmos.ClusterResource) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	location := "global"
	if v.Location != nil {
		location = *v.Location
	}
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			model.CosmosdbCassandraClusterDescription{
				CassandraCluster: *v,
				ResourceGroup:    resourceGroup,
			},
		},
	}
	return &resource
}

func documentDBDatabaseAccounts(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, resourceGroup string) ([]*armcosmos.DatabaseAccountGetResults, error) {
	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDatabaseAccountsClient()

	var values []*armcosmos.DatabaseAccountGetResults
	pager := client.NewListByResourceGroupPager(resourceGroup, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		values = append(values, page.Value...)
	}

	return values, nil
}

func CosmosdbAccount(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armcosmos.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDatabaseAccountsClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getCosmosdbAccount(ctx, v)
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

func getCosmosdbAccount(ctx context.Context, v *armcosmos.DatabaseAccountGetResults) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	location := ""
	if v.Location != nil {
		location = *v.Location
	}
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: location,
		Description: JSONAllFieldsMarshaller{
			model.CosmosdbAccountDescription{
				DatabaseAccountGetResults: *v,
				ResourceGroup:             resourceGroup,
			},
		},
	}
	return &resource
}
