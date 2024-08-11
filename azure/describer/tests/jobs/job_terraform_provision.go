package jobs

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/pkg/azure"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/pkg/terraform"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/workerpool"
)

type TerraformJob struct {
	resourceType       string
	terraformObject    terraform.TerrformObject
	vars               []string
	workerPool         *workerpool.WorkerPool
	AzureADCredentials azure.AzureADCredentials
	workerpool.TaskProperties
}

func NewTerraformJob(
	resourceType string,
	workingDirectory string,
	execPath string,
	vars []string,
	workerPool *workerpool.WorkerPool,
	azureCredentials azure.AzureADCredentials,
) *TerraformJob {
	return &TerraformJob{
		resourceType: resourceType,
		TaskProperties: workerpool.TaskProperties{
			ID:          uuid.New(),
			Description: fmt.Sprintf("Provisioning resource %s", resourceType),
		},
		terraformObject:    terraform.NewTerraformObject(workingDirectory, execPath),
		vars:               vars,
		workerPool:         workerPool,
		AzureADCredentials: azureCredentials,
	}
}

func (tj *TerraformJob) Properties() workerpool.TaskProperties {
	return tj.TaskProperties
}

func (tj *TerraformJob) Run(ctx context.Context) error {

	planFileName := fmt.Sprintf("%s.tfplan", tj.TaskProperties.ID.String())

	err := tj.terraformObject.Initialize()
	if err != nil {
		return fmt.Errorf("initialize failed: %s %w", tj.TaskProperties.ID.String(), err)
	}
	err = tj.terraformObject.Init(ctx)
	if err != nil {
		return fmt.Errorf("init failed: %w", err)
	}

	err = tj.terraformObject.Plan(ctx, tj.vars, planFileName)
	if err != nil {
		return fmt.Errorf("plan failed: %w", err)
	}

	err = tj.terraformObject.Apply(ctx, tj.vars, planFileName)
	if err != nil {
		return fmt.Errorf("apply failed: %w", err)
	}

	tj.workerPool.AddTask(
		NewDescriberJob(tj.resourceType, tj.AzureADCredentials),
	)

	return nil

}
