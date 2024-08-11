package terraform

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfexec"
)

// basically perform 'terraform init'
func (to *TerrformObject) Init(ctx context.Context) error {

	err := to.tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		return err
	}
	return nil

}
