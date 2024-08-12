package jobs

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/google/uuid"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/pkg/azure"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/workerpool"
)

type DescriberJob struct {
	resourceType string
	azureCred    azure.AzureADCredentials
	workerpool.TaskProperties
}

func NewDescriberJob(
	resourceType string,
	azureCred azure.AzureADCredentials,
) *DescriberJob {
	return &DescriberJob{
		resourceType: resourceType,
		TaskProperties: workerpool.TaskProperties{
			ID:          uuid.New(),
			Description: fmt.Sprintf("Describing resource %s", resourceType),
		},
		azureCred: azureCred,
	}
}

func (dj *DescriberJob) Properties() workerpool.TaskProperties {
	return dj.TaskProperties
}

func (dj *DescriberJob) Run(ctx context.Context) error {

	clientCredential, err := azidentity.NewClientSecretCredential(dj.azureCred.TenantID, dj.azureCred.ClientID, dj.azureCred.ClientSecret, nil)
	if err != nil {
		return err
	}

	describerFunc := azure.DescriberMap[dj.resourceType]

	resources, err := describerFunc(ctx, clientCredential, dj.azureCred.SubscriptionID, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources {
		log.Println(resource.ID)
		log.Println(resource.Name)
		log.Println(resource.Type)

	}

	// log.Println("Describer running")

	return nil

}
