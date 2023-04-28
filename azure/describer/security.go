package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v1.0/security"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func SecurityCenterAutoProvisioning(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := security.NewAutoProvisioningSettingsClient(subscription, "")
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: "global",
				Description: model.SecurityCenterAutoProvisioningDescription{
					AutoProvisioningSetting: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		if !result.NotDone() {
			break
		}

		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func SecurityCenterContact(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := security.NewContactsClient(subscription, "")
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: "global",
				Description: model.SecurityCenterContactDescription{
					Contact: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		if !result.NotDone() {
			break
		}

		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func SecurityCenterJitNetworkAccessPolicy(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := security.NewJitNetworkAccessPoliciesClient(subscription, "")
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: model.SecurityCenterJitNetworkAccessPolicyDescription{
					JitNetworkAccessPolicy: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		if !result.NotDone() {
			break
		}

		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func SecurityCenterSetting(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := security.NewSettingsClient(subscription, "")
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: "global",
				Description: model.SecurityCenterSettingDescription{
					Setting: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		if !result.NotDone() {
			break
		}

		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func SecurityCenterSubscriptionPricing(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := security.NewPricingsClient(subscription, "")
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: "global",
				Description: model.SecurityCenterSubscriptionPricingDescription{
					Pricing: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		if !result.NotDone() {
			break
		}

		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func SecurityCenterAutomation(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := security.NewAutomationsClient(subscription, "")
	client.Authorizer = authorizer

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroup := strings.Split(*v.ID, "/")[4]
			resource := Resource{
				ID:       *v.ID,
				Name:     *v.Name,
				Location: *v.Location,
				Description: model.SecurityCenterAutomationDescription{
					Automation:    v,
					ResourceGroup: resourceGroup,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		if !result.NotDone() {
			break
		}

		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func SecurityCenterSubAssessment(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := security.NewSubAssessmentsClient(subscription, "")
	client.Authorizer = authorizer

	result, err := client.ListAll(ctx, "subscriptions/"+subscription)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, v := range result.Values() {
			resourceGroup := strings.Split(*v.ID, "/")[4]
			resource := Resource{
				ID:   *v.ID,
				Name: *v.Name,
				Description: model.SecurityCenterSubAssessmentDescription{
					SubAssessment: v,
					ResourceGroup: resourceGroup,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}

		if !result.NotDone() {
			break
		}

		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}
