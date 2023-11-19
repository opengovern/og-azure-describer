# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the virtual network gateway.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a virtual network gateway uniquely.</td></tr>
	<tr><td>type</td><td>Type of the resource.</td></tr>
	<tr><td>gateway_type</td><td>The type of this virtual network gateway. Possible values include: 'Vpn', 'ExpressRoute'.</td></tr>
	<tr><td>vpn_type</td><td>The type of this virtual network gateway. Valid values are: 'PolicyBased', 'RouteBased'.</td></tr>
	<tr><td>vpn_gateway_generation</td><td>The generation for this virtual network gateway. Must be None if gatewayType is not VPN. Valid values are: 'None', 'Generation1', 'Generation2'.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the virtual network gateway resource.</td></tr>
	<tr><td>active_active</td><td>Indicates whether virtual network gateway configured with active-active mode, or not. If true, each Azure gateway instance will have a unique public IP address, and each will establish an IPsec/IKE S2S VPN tunnel to your on-premises VPN device specified in your local network gateway and connection.</td></tr>
	<tr><td>enable_bgp</td><td>Indicates whether BGP is enabled for this virtual network gateway, or not.</td></tr>
	<tr><td>enable_dns_forwarding</td><td>Indicates whether DNS forwarding is enabled, or not.</td></tr>
	<tr><td>enable_private_ip_address</td><td>Indicates whether private IP needs to be enabled on this gateway for connections or not.</td></tr>
	<tr><td>etag</td><td>An unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>gateway_default_site</td><td>The reference to the LocalNetworkGateway resource, which represents local network site having default routes. Assign Null value in case of removing existing default site setting.</td></tr>
	<tr><td>inbound_dns_forwarding_endpoint</td><td>The IP address allocated by the gateway to which dns requests can be sent.</td></tr>
	<tr><td>resource_guid</td><td>The resource GUID property of the virtual network gateway resource.</td></tr>
	<tr><td>sku_name</td><td>Gateway SKU name.</td></tr>
	<tr><td>sku_tier</td><td>Gateway SKU tier.</td></tr>
	<tr><td>sku_capacity</td><td>Gateway SKU capacity.</td></tr>
	<tr><td>bgp_settings</td><td>Virtual network gateway's BGP speaker settings.</td></tr>
	<tr><td>custom_routes_address_prefixes</td><td>A list of address blocks reserved for this virtual network in CIDR notation.</td></tr>
	<tr><td>gateway_connections</td><td>A list of virtual network gateway connection resources that exists in a resource group.</td></tr>
	<tr><td>ip_configurations</td><td>IP configurations for virtual network gateway.</td></tr>
	<tr><td>vpn_client_configuration</td><td>The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.</td></tr>
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