package terraform

import "github.com/hashicorp/terraform-exec/tfexec"

// terraform object to hold configurations and options
type TerrformObject struct {
	workingDirectory string
	execPath         string
	tf               *tfexec.Terraform
}

func NewTerraformObject(workingDirectory string, execPath string) TerrformObject {

	return TerrformObject{
		workingDirectory: workingDirectory,
		execPath:         execPath,
	}
}

// creates the terraform object using working directory and execPath
func (to *TerrformObject) Initialize() error {

	tf, err := tfexec.NewTerraform(to.workingDirectory, to.execPath)
	if err != nil {
		return err
	}
	to.tf = tf
	return nil
}
