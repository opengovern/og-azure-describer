package terraform

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func (to *TerrformObject) Destroy(ctx context.Context, tfvars []string) error {

	destroyOptions := []tfexec.DestroyOption{}

	for _, tfvar := range tfvars {
		destroyOptions = append(destroyOptions, tfexec.Var(tfvar))
	}

	err := to.tf.Destroy(ctx, destroyOptions...)
	if err != nil {
		return err
	}
	return nil
}
