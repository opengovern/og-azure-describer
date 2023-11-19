# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>The unique id identifying the resource in subscription.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the namespace.</td></tr>
	<tr><td>zone_redundant</td><td>Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</td></tr>
	<tr><td>created_at</td><td>The time the namespace was created.</td></tr>
	<tr><td>metric_id</td><td>The identifier for Azure insights metrics.</td></tr>
	<tr><td>servicebus_endpoint</td><td>Specifies the endpoint used to perform Service Bus operations.</td></tr>
	<tr><td>sku_capacity</td><td>The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</td></tr>
	<tr><td>sku_name</td><td>Name of this SKU. Valid valuer are: 'Basic', 'Standard', 'Premium'.</td></tr>
	<tr><td>sku_tier</td><td>The billing tier of this particular SKU. Valid values are: 'Basic', 'Standard', 'Premium'.</td></tr>
	<tr><td>updated_at</td><td>The time the namespace was updated.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the servicebus namespace.</td></tr>
	<tr><td>encryption</td><td>Specifies the properties of BYOK encryption configuration. Customer-managed key encryption at rest (Bring Your Own Key) is only available on Premium namespaces.</td></tr>
	<tr><td>network_rule_set</td><td>Describes the network rule set for specified namespace. The ServiceBus Namespace must be Premium in order to attach a ServiceBus Namespace Network Rule Set.</td></tr>
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