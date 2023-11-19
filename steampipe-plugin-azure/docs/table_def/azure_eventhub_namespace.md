# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>The ID of the resource.</td></tr>
	<tr><td>type</td><td>The resource type.</td></tr>
	<tr><td>provisioning_state</td><td>Provisioning state of the namespace.</td></tr>
	<tr><td>created_at</td><td>The time the namespace was created.</td></tr>
	<tr><td>cluster_arm_id</td><td>Cluster ARM ID of the namespace.</td></tr>
	<tr><td>is_auto_inflate_enabled</td><td>Indicates whether auto-inflate is enabled for eventhub namespace.</td></tr>
	<tr><td>kafka_enabled</td><td>Indicates whether kafka is enabled for eventhub namespace, or not.</td></tr>
	<tr><td>maximum_throughput_units</td><td>Upper limit of throughput units when auto-inflate is enabled, value should be within 0 to 20 throughput units.</td></tr>
	<tr><td>metric_id</td><td>Identifier for azure insights metrics.</td></tr>
	<tr><td>service_bus_endpoint</td><td>Endpoint you can use to perform service bus operations.</td></tr>
	<tr><td>sku_capacity</td><td>The Event Hubs throughput units, value should be 0 to 20 throughput units.</td></tr>
	<tr><td>sku_name</td><td>Name of this SKU. Possible values include: 'Basic', 'Standard'.</td></tr>
	<tr><td>sku_tier</td><td>The billing tier of this particular SKU. Valid values are: 'Basic', 'Standard', 'Premium'.</td></tr>
	<tr><td>updated_at</td><td>The time the namespace was updated.</td></tr>
	<tr><td>zone_redundant</td><td>Enabling this property creates a standard event hubs namespace in regions supported availability zones.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the eventhub namespace.</td></tr>
	<tr><td>encryption</td><td>Properties of BYOK encryption description.</td></tr>
	<tr><td>identity</td><td>Describes the properties of BYOK encryption description.</td></tr>
	<tr><td>network_rule_set</td><td>Describes the network rule set for specified namespace. The EventHub Namespace must be Premium in order to attach a EventHub Namespace Network Rule Set.</td></tr>
	<tr><td>private_endpoint_connections</td><td>The private endpoint connections of the namespace.</td></tr>
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