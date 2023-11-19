# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>The unique id identifying the resource in subscription.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the container registry at the time the operation was called. Valid values are: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'.</td></tr>
	<tr><td>creation_date</td><td>The creation date of the container registry.</td></tr>
	<tr><td>admin_user_enabled</td><td>Indicates whether the admin user is enabled, or not.</td></tr>
	<tr><td>data_endpoint_enabled</td><td>Enable a single data endpoint per region for serving data.</td></tr>
	<tr><td>login_server</td><td>The URL that can be used to log into the container registry.</td></tr>
	<tr><td>network_rule_bypass_options</td><td>Indicates whether to allow trusted Azure services to access a network restricted registry. Valid values are: 'AzureServices', 'None'.</td></tr>
	<tr><td>public_network_access</td><td>Indicates whether or not public network access is allowed for the container registry. Valid values are: 'Enabled', 'Disabled'.</td></tr>
	<tr><td>sku_name</td><td>The SKU name of the container registry. Required for registry creation. Valid values are: 'Classic', 'Basic', 'Standard', 'Premium'.</td></tr>
	<tr><td>sku_tier</td><td>The SKU tier based on the SKU name. Valid values are: 'Classic', 'Basic', 'Standard', 'Premium'.</td></tr>
	<tr><td>status</td><td>The current status of the resource.</td></tr>
	<tr><td>status_message</td><td>The detailed message for the status, including alerts and error messages.</td></tr>
	<tr><td>status_timestamp</td><td>The timestamp when the status was changed to the current value.</td></tr>
	<tr><td>storage_account_id</td><td>The resource ID of the storage account. Only applicable to Classic SKU.</td></tr>
	<tr><td>zone_redundancy</td><td>Indicates whether or not zone redundancy is enabled for this container registry. Valid values are: 'Enabled', 'Disabled'.</td></tr>
	<tr><td>data_endpoint_host_names</td><td>A list of host names that will serve data when dataEndpointEnabled is true.</td></tr>
	<tr><td>encryption</td><td>The encryption settings of container registry.</td></tr>
	<tr><td>identity</td><td>The identity of the container registry.</td></tr>
	<tr><td>login_credentials</td><td>The login credentials for the specified container registry.</td></tr>
	<tr><td>network_rule_set</td><td>The network rule set for a container registry.</td></tr>
	<tr><td>policies</td><td>The policies for a container registry.</td></tr>
	<tr><td>private_endpoint_connections</td><td>A list of private endpoint connections for a container registry.</td></tr>
	<tr><td>system_data</td><td>Metadata pertaining to creation and last modification of the resource.</td></tr>
	<tr><td>usages</td><td>Specifies the quota usages for the specified container registry.</td></tr>
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