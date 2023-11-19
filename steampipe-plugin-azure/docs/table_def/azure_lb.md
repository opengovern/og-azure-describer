# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The resource name.</td></tr>
	<tr><td>id</td><td>The resource ID.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the load balancer resource. Possible values include: &#39;Succeeded&#39;, &#39;Updating&#39;, &#39;Deleting&#39;, &#39;Failed&#39;.</td></tr>
	<tr><td>type</td><td>The resource type.</td></tr>
	<tr><td>etag</td><td>A unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>extended_location_name</td><td>The name of the extended location.</td></tr>
	<tr><td>extended_location_type</td><td>The type of the extended location. Possible values include: &#39;ExtendedLocationTypesEdgeZone&#39;.</td></tr>
	<tr><td>resource_guid</td><td>The resource GUID property of the load balancer resource.</td></tr>
	<tr><td>sku_name</td><td>Name of the load balancer SKU. Possible values include: &#39;Basic&#39;, &#39;Standard&#39;, &#39;Gateway&#39;.</td></tr>
	<tr><td>sku_tier</td><td>Tier of the load balancer SKU. Possible values include: &#39;Regional&#39;, &#39;Global&#39;.</td></tr>
	<tr><td>backend_address_pools</td><td>Collection of backend address pools used by the load balancer.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the load balancer.</td></tr>
	<tr><td>frontend_ip_configurations</td><td>Object representing the frontend IPs to be used for the load balancer.</td></tr>
	<tr><td>inbound_nat_pools</td><td>Defines an external port range for inbound NAT to a single backend port on NICs associated with the load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on the Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.</td></tr>
	<tr><td>inbound_nat_rules</td><td>Collection of inbound NAT Rules used by the load balancer. Defining inbound NAT rules on the load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.</td></tr>
	<tr><td>load_balancing_rules</td><td>Object collection representing the load balancing rules Gets the provisioning.</td></tr>
	<tr><td>outbound_rules</td><td>The outbound rules.</td></tr>
	<tr><td>probes</td><td>Collection of probe objects used in the load balancer.</td></tr>
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