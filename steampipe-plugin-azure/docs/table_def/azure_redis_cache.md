# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>The unique id identifying the resource in subscription.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the redis instance at the time the operation was called. Valid values are: 'Creating', 'Deleting', 'Disabled', 'Failed', 'Linking', 'Provisioning', 'RecoveringScaleFailure', 'Scaling', 'Succeeded', 'Unlinking', 'Unprovisioning', and 'Updating'.</td></tr>
	<tr><td>redis_version</td><td>Specifies the version.</td></tr>
	<tr><td>enable_non_ssl_port</td><td>Specifies whether the non-ssl Redis server port (6379) is enabled.</td></tr>
	<tr><td>host_name</td><td>Specifies the name of the redis host.</td></tr>
	<tr><td>minimum_tls_version</td><td>Specifies the TLS version requires to connect.</td></tr>
	<tr><td>port</td><td>Specifies the redis non-SSL port.</td></tr>
	<tr><td>public_network_access</td><td>Indicates whether or not public endpoint access is allowed for this cache. Valid values are: 'Enabled', 'Disabled'.</td></tr>
	<tr><td>sku_capacity</td><td>The size of the Redis cache to deploy.</td></tr>
	<tr><td>sku_family</td><td>The SKU family to use.</td></tr>
	<tr><td>sku_name</td><td>The type of Redis cache to deploy.</td></tr>
	<tr><td>ssl_port</td><td>Specifies the redis SSL port.</td></tr>
	<tr><td>subnet_id</td><td>The full resource ID of a subnet in a virtual network to deploy the Redis cache in.</td></tr>
	<tr><td>static_ip</td><td>Specifies the statis IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.</td></tr>
	<tr><td>replicas_per_master</td><td>The number of replicas to be created per master.</td></tr>
	<tr><td>shard_count</td><td>The number of shards to be created on a premium cluster cache.</td></tr>
	<tr><td>access_keys</td><td>The keys of the Redis cache.</td></tr>
	<tr><td>linked_servers</td><td>A list of the linked servers associated with the cache.</td></tr>
	<tr><td>instances</td><td>A list of the Redis instances associated with the cache.</td></tr>
	<tr><td>private_endpoint_connections</td><td>A list of private endpoint connection associated with the specified redis cache.</td></tr>
	<tr><td>redis_configuration</td><td>Describes the redis cache configuration.</td></tr>
	<tr><td>tenant_settings</td><td>A dictionary of tenant settings.</td></tr>
	<tr><td>zones</td><td>A list of availability zones denoting where the resource needs to come from.</td></tr>
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