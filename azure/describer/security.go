package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func SecurityCenterAutoProvisioning(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAutoProvisioningSettingsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterAutoProvisioning(ctx, v)
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

func GetSecurityCenterAutoProvisioning(ctx context.Context, v *armsecurity.AutoProvisioningSetting) *Resource {
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.SecurityCenterAutoProvisioningDescription{
				AutoProvisioningSetting: *v,
			},
		},
	}

	return &resource
}

func SecurityCenterContact(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewContactsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.ContactList.Value {
			resource := GetSecurityCenterContact(ctx, v)
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

func GetSecurityCenterContact(ctx context.Context, v *armsecurity.Contact) *Resource {
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.SecurityCenterContactDescription{
				Contact: *v,
			},
		},
	}
	return &resource
}

func SecurityCenterJitNetworkAccessPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewJitNetworkAccessPoliciesClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterJitNetworkAccessPolicy(ctx, v)
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

func GetSecurityCenterJitNetworkAccessPolicy(ctx context.Context, v *armsecurity.JitNetworkAccessPolicy) *Resource {
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.SecurityCenterJitNetworkAccessPolicyDescription{
				JitNetworkAccessPolicy: *v,
			},
		},
	}
	return &resource
}

func SecurityCenterSetting(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSettingsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterSetting(ctx, v)
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

func GetSecurityCenterSetting(ctx context.Context, v armsecurity.SettingClassification) *Resource {
	var settingStatus bool
	if *v.GetSetting().Kind == armsecurity.SettingKindDataExportSettings {
		settingStatus = true
	} else {
		settingStatus = false
	}
	resource := Resource{
		ID:       *v.GetSetting().ID,
		Name:     *v.GetSetting().Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.SecurityCenterSettingDescription{
				Setting:             *v.GetSetting(),
				ExportSettingStatus: settingStatus,
			},
		},
	}
	return &resource
}

func SecurityCenterSubscriptionPricing(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewPricingsClient()

	var values []Resource
	list, err := client.List(ctx, nil)
	if err != nil {
		return nil, err
	}
	for _, v := range list.Value {
		resource := GetSecurityCenterSubscriptionPricing(ctx, v)
		if stream != nil {
			if err := (*stream)(*resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSecurityCenterSubscriptionPricing(ctx context.Context, v *armsecurity.Pricing) *Resource {
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: JSONAllFieldsMarshaller{
			Value: model.SecurityCenterSubscriptionPricingDescription{
				Pricing: *v,
			},
		},
	}
	return &resource
}

func SecurityCenterAutomation(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAutomationsClient()

	var values []Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterAutomation(ctx, v)
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

func GetSecurityCenterAutomation(ctx context.Context, v *armsecurity.Automation) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.SecurityCenterAutomationDescription{
				Automation:    *v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}

func SecurityCenterSubAssessment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSubAssessmentsClient()

	var values []Resource
	pager := client.NewListAllPager("subscriptions/"+subscription, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterSubAssessment(ctx, v)
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

func GetSecurityCenterSubAssessment(ctx context.Context, v *armsecurity.SubAssessment) *Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := Resource{
		ID:       *v.ID,
		Location: "global",
		Name:     *v.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.SecurityCenterSubAssessmentDescription{
				SubAssessment: *v,
				ResourceGroup: resourceGroup,
			},
		},
	}
	return &resource
}
