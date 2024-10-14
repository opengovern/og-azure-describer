package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/opengovern/og-azure-describer/azure/model"
)

func EventGridDomainTopic(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	rgs, err := listResourceGroups(ctx, cred, subscription)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armeventgrid.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDomainTopicsClient()

	var values []Resource
	for _, rg := range rgs {
		domains, err := eventGridDomain(ctx, cred, subscription, *rg.Name)
		if err != nil {
			return nil, err
		}

		for _, domain := range domains {
			it := client.NewListByDomainPager(*rg.Name, *domain.Name, nil)
			for it.More() {
				page, err := it.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				for _, v := range page.Value {
					resource := getEventGridDomainTopic(ctx, v)
					if stream != nil {
						if err := (*stream)(*resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, *resource)
					}
				}
			}
		}
	}
	return values, nil
}

func getEventGridDomainTopic(ctx context.Context, v *armeventgrid.DomainTopic) *Resource {
	return &Resource{
		ID:          *v.ID,
		Name:        *v.Name,
		Location:    "global",
		Description: JSONAllFieldsMarshaller{Value: v},
	}
}

func eventGridDomain(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, resourceGroup string) ([]*armeventgrid.Domain, error) {
	clientFactory, err := armeventgrid.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDomainsClient()

	pager := client.NewListByResourceGroupPager(resourceGroup, nil)
	var values []*armeventgrid.Domain
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		values = append(values, page.Value...)
	}
	return values, nil
}

func EventGridDomain(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armeventgrid.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewDomainsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getEventGridDomain(ctx, v, diagnosticClient)
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

func getEventGridDomain(ctx context.Context, domain *armeventgrid.Domain, client *armmonitor.DiagnosticSettingsClient) (*Resource, error) {
	resourceGroup := strings.Split(*domain.ID, "/")[4]

	id := *domain.ID
	pager := client.NewListPager(id, nil)
	var eventgridListOp []*armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		eventgridListOp = append(eventgridListOp, page.Value...)
	}

	resource := Resource{
		ID:       *domain.ID,
		Name:     *domain.Name,
		Location: *domain.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.EventGridDomainDescription{
				Domain:                      *domain,
				DiagnosticSettingsResources: eventgridListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}

func EventGridTopic(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armeventgrid.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewTopicsClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource, err := getEventGridTopic(ctx, v, diagnosticClient)
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

func getEventGridTopic(ctx context.Context, v *armeventgrid.Topic, client *armmonitor.DiagnosticSettingsClient) (*Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	id := *v.ID
	pager := client.NewListPager(id, nil)
	var eventgridListOp []*armmonitor.DiagnosticSettingsResource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		eventgridListOp = append(eventgridListOp, page.Value...)
	}

	resource := Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.EventGridTopicDescription{
				Topic:                       *v,
				DiagnosticSettingsResources: eventgridListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}
