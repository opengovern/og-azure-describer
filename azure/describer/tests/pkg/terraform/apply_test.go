package terraform

import (
	"context"
	"os/exec"
	"testing"
)

func TestApply(t *testing.T) {

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

	tfvars := []string{
		"resource_group_name=test-resource-group",
		"location=Canada Central",
		"resourceCount=2",
	}

	err = to.Plan(context.Background(), tfvars, "test.tfplan")
	if err != nil {
		t.Log(err)
		t.Errorf("Cannot create terraform plan")
	}

	err = to.Apply(context.Background(), tfvars, "test.tfplan")
	if err != nil {
		t.Log(err)
	}

}
