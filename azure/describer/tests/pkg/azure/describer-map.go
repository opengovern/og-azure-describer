package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/kaytu-io/kaytu-azure-describer/azure/describer"
)

var DescriberMap = map[string]func(context.Context, *azidentity.ClientSecretCredential, string, *describer.StreamSender) ([]describer.Resource, error){
	"resource_group": describer.ResourceGroup,
}
