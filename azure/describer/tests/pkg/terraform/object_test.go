package terraform

import (
	"os/exec"
	"testing"
)

func TestNewTerraformObject(t *testing.T) {

	execPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Errorf("Cannot find 'terraform' path")
	}

	_ = NewTerraformObject("templates/resource_group", execPath)

}

func TestInitialize(t *testing.T) {

	execPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Errorf("Cannot find 'terraform' path")
	}

	to := NewTerraformObject("templates/resource_group", execPath)

	err = to.Initialize()
	if err != nil {
		t.Log(err)
		t.Errorf("Initilize failed")
	}

}
