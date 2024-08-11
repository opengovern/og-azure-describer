package tests

import (
	"context"
	"fmt"
	"os/exec"
	"path"

	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/jobs"

	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/workerpool"
)

func Controller() error {

	execPath, err := exec.LookPath("terraform")
	if err != nil {
		return fmt.Errorf("cannot find 'terraform' path")
	}

	workerPool := workerpool.NewWorkerPool(ConcurrentWorkers)
	workerPool.Start(context.Background())

	parameters, err := ParseParameters()
	if err != nil {
		return err
	}

	for _, parameter := range parameters {

		workingDirectory := path.Join(WorkingDirectory, parameter.ResourceType)
		terraformJob := jobs.NewTerraformJob(
			parameter.ResourceType,
			workingDirectory,
			execPath,
			parameter.Vars,
			workerPool,
			parameter.Credentials,
		)

		workerPool.AddTask(terraformJob)

	}

	// workerPool.Wait()
	workerPool.Wg.Wait()

	return nil

}
