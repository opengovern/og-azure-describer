package azure

import (
	"context"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureFirewallPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_firewall_policy",
		Description: "Azure Firewall Policy",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetFirewallPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListFirewallPolicy,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the firewall policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicy.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a firewall policy uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicy.ID")},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicies.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the firewall policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicies.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the firewall policy resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.ProvisioningState")},
			{
				Name:        "intrusion_detection_mode",
				Description: "Intrusion detection general state. Possible values include: 'FirewallPolicyIntrusionDetectionStateTypeOff', 'FirewallPolicyIntrusionDetectionStateTypeAlert', 'FirewallPolicyIntrusionDetectionStateTypeDeny'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.IntrusionDetection.Mode")},
			{
				Name:        "sku_tier",
				Description: "Tier of Firewall Policy. Possible values include: 'FirewallPolicySkuTierStandard', 'FirewallPolicySkuTierPremium'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.SKU.Tier")},
			{
				Name:        "threat_intel_mode",
				Description: "The operation mode for Threat Intelligence. Possible values include: 'AzureFirewallThreatIntelModeAlert', 'AzureFirewallThreatIntelModeDeny', 'AzureFirewallThreatIntelModeOff'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.ThreatIntelMode")},
			{
				Name:        "base_policy",
				Description: "The parent firewall policy from which rules are inherited.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.BasePolicy")},
			{
				Name:        "child_policies",
				Description: "List of references to Child Firewall Policies.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.ChildPolicies")},
			{
				Name:        "dns_settings",
				Description: "DNS Proxy Settings definition.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.DNSSettings")},
			{
				Name:        "firewalls",
				Description: "List of references to Azure Firewalls that this Firewall Policy is associated with.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.Firewalls")},
			{
				Name:        "identity",
				Description: "The identity of the firewall policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Identity")},
			{
				Name:        "intrusion_detection_configuration",
				Description: "Intrusion detection configuration properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.IntrusionDetection.Configuration")},
			{
				Name:        "rule_collection_groups",
				Description: "List of references to FirewallPolicyRuleCollectionGroups.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.RuleCollectionGroups")},
			{
				Name:        "threat_intel_whitelist_ip_addresses",
				Description: "List of IP addresses for the ThreatIntel Whitelist.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.ThreatIntelWhitelist.IPAddresses")},
			{
				Name:        "threat_intel_whitelist_fqdns",
				Description: "List of FQDNs for the ThreatIntel Whitelist.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.ThreatIntelWhitelist.Fqdns")},
			{
				Name:        "transport_security_certificate_authority",
				Description: "The CA used for intermediate CA generation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicies.Properties.TransportSecurity.CertificateAuthority")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicy.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicy.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallPolicy.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicy.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}