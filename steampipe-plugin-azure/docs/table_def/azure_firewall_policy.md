# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the firewall policy.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a firewall policy uniquely.</td></tr>
	<tr><td>etag</td><td>A unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>type</td><td>The resource type of the firewall policy.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the firewall policy resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.</td></tr>
	<tr><td>intrusion_detection_mode</td><td>Intrusion detection general state. Possible values include: 'FirewallPolicyIntrusionDetectionStateTypeOff', 'FirewallPolicyIntrusionDetectionStateTypeAlert', 'FirewallPolicyIntrusionDetectionStateTypeDeny'.</td></tr>
	<tr><td>sku_tier</td><td>Tier of Firewall Policy. Possible values include: 'FirewallPolicySkuTierStandard', 'FirewallPolicySkuTierPremium'.</td></tr>
	<tr><td>threat_intel_mode</td><td>The operation mode for Threat Intelligence. Possible values include: 'AzureFirewallThreatIntelModeAlert', 'AzureFirewallThreatIntelModeDeny', 'AzureFirewallThreatIntelModeOff'.</td></tr>
	<tr><td>base_policy</td><td>The parent firewall policy from which rules are inherited.</td></tr>
	<tr><td>child_policies</td><td>List of references to Child Firewall Policies.</td></tr>
	<tr><td>dns_settings</td><td>DNS Proxy Settings definition.</td></tr>
	<tr><td>firewalls</td><td>List of references to Azure Firewalls that this Firewall Policy is associated with.</td></tr>
	<tr><td>identity</td><td>The identity of the firewall policy.</td></tr>
	<tr><td>intrusion_detection_configuration</td><td>Intrusion detection configuration properties.</td></tr>
	<tr><td>rule_collection_groups</td><td>List of references to FirewallPolicyRuleCollectionGroups.</td></tr>
	<tr><td>threat_intel_whitelist_ip_addresses</td><td>List of IP addresses for the ThreatIntel Whitelist.</td></tr>
	<tr><td>threat_intel_whitelist_fqdns</td><td>List of FQDNs for the ThreatIntel Whitelist.</td></tr>
	<tr><td>transport_security_certificate_authority</td><td>The CA used for intermediate CA generation.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>region</td><td>The Azure region/location in which the resource is located.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>