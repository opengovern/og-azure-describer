package tests

import (
	"encoding/json"
	"io"
	"os"

	"github.com/kaytu-io/kaytu-azure-describer/azure/describer/tests/pkg/azure"
)

type ResourceParameters struct {
	ResourceType string                   `json:"resource_type"`
	Vars         []string                 `json:"vars"`
	Subscription string                   `json:"subscription"`
	Credentials  azure.AzureADCredentials `json:"azure_credentials"`
}

var ConcurrentWorkers = 1
var WorkingDirectory = "pkg/terraform/templates"

func ParseParameters() ([]*ResourceParameters, error) {

	var params []*ResourceParameters

	jsonFile, err := os.Open("parameters.json")
	if err != nil {
		return nil, err
	}

	byteArray, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(byteArray, &params); err != nil {
		return nil, err
	}

	return params, nil

}
