package terraform

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func (to *TerrformObject) Plan(ctx context.Context, tfvars []string, testPlanName string) error {

	// outPath := path.Join(to.workingDirectory, "test.tfplan")

	planOptions := []tfexec.PlanOption{
		tfexec.Out(testPlanName),
	}

	for _, tfvar := range tfvars {
		planOptions = append(planOptions, tfexec.Var(tfvar))
	}

	_, err := to.tf.Plan(ctx, planOptions...)
	if err != nil {
		return err
	}

	return nil

}
