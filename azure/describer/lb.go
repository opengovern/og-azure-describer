package describer

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-04-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
)

func LoadBalancer(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := network.NewLoadBalancersClient(subscription)
	client.Authorizer = authorizer

	insightsClient := insights.NewDiagnosticSettingsClient(subscription)
	insightsClient.Authorizer = authorizer

	result, err := client.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, loadBalancer := range result.Values() {
			resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

			// Get diagnostic settings
			diagnosticSettings, err := insightsClient.List(ctx, *loadBalancer.ID)
			if err != nil {
				return nil, err
			}

			resource := Resource{
				ID:       *loadBalancer.ID,
				Name:     *loadBalancer.Name,
				Location: *loadBalancer.Location,
				Description: JSONAllFieldsMarshaller{
					model.LoadBalancerDescription{
						ResourceGroup:     resourceGroup,
						DiagnosticSetting: diagnosticSettings.Value,
						LoadBalancer:      loadBalancer,
					},
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

func LoadBalancerBackendAddressPool(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := network.NewLoadBalancersClient(subscription)
	client.Authorizer = authorizer

	poolClient := network.NewLoadBalancerBackendAddressPoolsClient(subscription)
	poolClient.Authorizer = authorizer

	result, err := client.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, loadBalancer := range result.Values() {
			resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

			backendAddressPools, err := poolClient.List(ctx, resourceGroup, *loadBalancer.Name)
			if err != nil {
				return nil, err
			}
			for {
				for _, pool := range backendAddressPools.Values() {
					resourceGroup := strings.Split(*pool.ID, "/")[4]
					resource := Resource{
						ID: *pool.ID,
						Description: JSONAllFieldsMarshaller{
							model.LoadBalancerBackendAddressPoolDescription{
								ResourceGroup: resourceGroup,
								LoadBalancer:  loadBalancer,
								Pool:          pool,
							},
						},
					}
					if pool.Name != nil {
						resource.Name = *pool.Name
					}
					if pool.Location != nil {
						resource.Location = *pool.Location
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, resource)
					}
				}
				if !backendAddressPools.NotDone() {
					break
				}
				err = backendAddressPools.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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

func LoadBalancerNatRule(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := network.NewLoadBalancersClient(subscription)
	client.Authorizer = authorizer

	natRulesClient := network.NewInboundNatRulesClient(subscription)
	natRulesClient.Authorizer = authorizer

	result, err := client.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, loadBalancer := range result.Values() {
			resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

			natRules, err := natRulesClient.List(ctx, resourceGroup, *loadBalancer.Name)
			if err != nil {
				return nil, err
			}
			for {
				for _, natRule := range natRules.Values() {
					resourceGroup := strings.Split(*natRule.ID, "/")[4]
					resource := Resource{
						ID:       *natRule.ID,
						Name:     *natRule.Name,
						Location: *loadBalancer.Location,
						Description: JSONAllFieldsMarshaller{
							model.LoadBalancerNatRuleDescription{
								ResourceGroup:    resourceGroup,
								LoadBalancerName: *loadBalancer.Name,
								Rule:             natRule,
							},
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
				if !natRules.NotDone() {
					break
				}
				err = natRules.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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

func LoadBalancerOutboundRule(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := network.NewLoadBalancersClient(subscription)
	client.Authorizer = authorizer

	outboundRulesClient := network.NewLoadBalancerOutboundRulesClient(subscription)
	outboundRulesClient.Authorizer = authorizer

	result, err := client.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, loadBalancer := range result.Values() {
			resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

			outboundRuleListResultPage, err := outboundRulesClient.List(ctx, resourceGroup, *loadBalancer.Name)
			if err != nil {
				return nil, err
			}
			for {
				for _, outboundRule := range outboundRuleListResultPage.Values() {
					resourceGroup := strings.Split(*outboundRule.ID, "/")[4]
					resource := Resource{
						ID:       *outboundRule.ID,
						Name:     *outboundRule.Name,
						Location: *loadBalancer.Location,
						Description: JSONAllFieldsMarshaller{
							model.LoadBalancerOutboundRuleDescription{
								ResourceGroup:    resourceGroup,
								LoadBalancerName: *loadBalancer.Name,
								Rule:             outboundRule,
							},
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
				if !outboundRuleListResultPage.NotDone() {
					break
				}
				err = outboundRuleListResultPage.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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

func LoadBalancerProbe(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := network.NewLoadBalancersClient(subscription)
	client.Authorizer = authorizer

	probesClient := network.NewLoadBalancerProbesClient(subscription)
	probesClient.Authorizer = authorizer

	result, err := client.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, loadBalancer := range result.Values() {
			resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

			probeListResultPage, err := probesClient.List(ctx, resourceGroup, *loadBalancer.Name)
			if err != nil {
				return nil, err
			}
			for {
				for _, probe := range probeListResultPage.Values() {
					resourceGroup := strings.Split(*probe.ID, "/")[4]
					resource := Resource{
						ID:       *probe.ID,
						Name:     *probe.Name,
						Location: *loadBalancer.Location,
						Description: JSONAllFieldsMarshaller{
							model.LoadBalancerProbeDescription{
								ResourceGroup:    resourceGroup,
								LoadBalancerName: *loadBalancer.Name,
								Probe:            probe,
							},
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
				if !probeListResultPage.NotDone() {
					break
				}
				err = probeListResultPage.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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

func LoadBalancerRule(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := network.NewLoadBalancersClient(subscription)
	client.Authorizer = authorizer

	rulesClient := network.NewLoadBalancerLoadBalancingRulesClient(subscription)
	rulesClient.Authorizer = authorizer

	result, err := client.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, loadBalancer := range result.Values() {
			resourceGroup := strings.Split(*loadBalancer.ID, "/")[4]

			ruleListResultPage, err := rulesClient.List(ctx, resourceGroup, *loadBalancer.Name)
			if err != nil {
				return nil, err
			}
			for {
				for _, rule := range ruleListResultPage.Values() {
					resourceGroup := strings.Split(*rule.ID, "/")[4]
					resource := Resource{
						ID:       *rule.ID,
						Name:     *rule.Name,
						Location: *loadBalancer.Location,
						Description: JSONAllFieldsMarshaller{
							model.LoadBalancerRuleDescription{
								ResourceGroup:    resourceGroup,
								LoadBalancerName: *loadBalancer.Name,
								Rule:             rule,
							},
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
				if !ruleListResultPage.NotDone() {
					break
				}
				err = ruleListResultPage.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}
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
