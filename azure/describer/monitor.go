package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/opengovern/og-azure-describer/azure/model"
)

func MonitorLogProfiles(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	logProfileClient := clientFactory.NewLogProfilesClient()

	pager := logProfileClient.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, logProfile := range page.Value {
			resource, err := getMonitorLogProfile(ctx, logProfile)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getMonitorLogProfile(ctx context.Context, logProfile *armmonitor.LogProfileResource) (*Resource, error) {

	resourceGroup := strings.Split(*logProfile.ID, "/")[4]

	resource := Resource{
		ID:       *logProfile.ID,
		Name:     *logProfile.Name,
		Location: *logProfile.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.MonitorLogProfileDescription{
				LogProfile:    *logProfile,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource, nil

}
