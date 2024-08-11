package terraform

import (
	"context"
	"os/exec"
	"testing"
)

func TestInit(t *testing.T) {

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

	err = to.Init(context.Background())
	if err != nil {
		t.Log(err)
		t.Error("Init failed")
	}

}
