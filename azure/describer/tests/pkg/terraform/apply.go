package terraform

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func (to *TerrformObject) Apply(ctx context.Context, tfvars []string, testPlanName string) error {

	applyOptions := []tfexec.ApplyOption{
		tfexec.DirOrPlan(testPlanName),
	}

	err := to.tf.Apply(ctx, applyOptions...)
	if err != nil {
		return err
	}
	return nil
}
