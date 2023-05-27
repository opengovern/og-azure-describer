package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func DataFactory(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	connClient := datafactory.NewPrivateEndPointConnectionsClient(subscription)
	connClient.Authorizer = authorizer
	factoryClient := datafactory.NewFactoriesClient(subscription)
	factoryClient.Authorizer = authorizer
	result, err := factoryClient.List(ctx)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for {
		for _, factory := range result.Values() {
			factoryName := factory.Name
			resourceGroup := strings.Split(*factory.ID, "/")[4]

			datafactoryListByFactoryOp, err := connClient.ListByFactory(ctx, resourceGroup, *factoryName)
			if err != nil {
				return nil, err
			}
			v := datafactoryListByFactoryOp.Values()
			for datafactoryListByFactoryOp.NotDone() {
				err := datafactoryListByFactoryOp.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}

				v = append(v, datafactoryListByFactoryOp.Values()...)
			}

			resource := Resource{
				ID:       *factory.ID,
				Name:     *factory.Name,
				Location: *factory.Location,
				Description: JSONAllFieldsMarshaller{
					model.DataFactoryDescription{
						Factory:                    factory,
						PrivateEndPointConnections: v,
						ResourceGroup:              resourceGroup,
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
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func DataFactoryDataset(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	factoryClient := datafactory.NewFactoriesClient(subscription)
	factoryClient.Authorizer = authorizer

	datasetClient := datafactory.NewDatasetsClient(subscription)
	datasetClient.Authorizer = authorizer

	result, err := factoryClient.List(ctx)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for {
		for _, factory := range result.Values() {
			factoryName := factory.Name
			factoryResourceGroup := strings.Split(*factory.ID, "/")[4]

			datasetListResponsePage, err := datasetClient.ListByFactory(ctx, factoryResourceGroup, *factoryName)
			if err != nil {
				return nil, err
			}

			for datasetListResponsePage.NotDone() {
				err := datasetListResponsePage.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
				for _, dataset := range datasetListResponsePage.Values() {
					resource := Resource{
						ID:       *dataset.ID,
						Name:     *dataset.Name,
						Location: *factory.Location,
						Description: JSONAllFieldsMarshaller{
							model.DataFactoryDatasetDescription{
								Factory:       factory,
								Dataset:       dataset,
								ResourceGroup: factoryResourceGroup,
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
				err = datasetListResponsePage.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
			}

		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func DataFactoryPipeline(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	factoryClient := datafactory.NewFactoriesClient(subscription)
	factoryClient.Authorizer = authorizer

	pipelineClient := datafactory.NewPipelinesClient(subscription)
	pipelineClient.Authorizer = authorizer

	result, err := factoryClient.List(ctx)
	if err != nil {
		return nil, err
	}
	var values []Resource
	for {
		for _, factory := range result.Values() {
			factoryName := factory.Name
			factoryResourceGroup := strings.Split(*factory.ID, "/")[4]

			pipelineListResponsePage, err := pipelineClient.ListByFactory(ctx, factoryResourceGroup, *factoryName)
			if err != nil {
				return nil, err
			}

			for pipelineListResponsePage.NotDone() {
				err := pipelineListResponsePage.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
				for _, pipelineResource := range pipelineListResponsePage.Values() {
					resource := Resource{
						ID:       *pipelineResource.ID,
						Name:     *pipelineResource.Name,
						Location: *factory.Location,
						Description: JSONAllFieldsMarshaller{
							model.DataFactoryPipelineDescription{
								Factory:       factory,
								Pipeline:      pipelineResource,
								ResourceGroup: factoryResourceGroup,
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
				err = pipelineListResponsePage.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
			}

		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
