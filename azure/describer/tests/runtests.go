package tests

import (
	"context"
	"fmt"
	"os/exec"
	"path"

	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/jobs"

	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/workerpool"
)

// Create resource using terraform
// collect verification data from tf files
// run describer
// compare describer output with verification data
// publish result

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

	workerPool.Wait()

	return nil

}
