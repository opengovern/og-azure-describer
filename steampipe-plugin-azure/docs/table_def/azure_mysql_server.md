# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the server.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a server uniquely.</td></tr>
	<tr><td>type</td><td>The resource type of the server.</td></tr>
	<tr><td>state</td><td>The state of the server.</td></tr>
	<tr><td>version</td><td>Specifies the version of the server.</td></tr>
	<tr><td>user_visible_state</td><td>A state of a server that is visible to user. Possible values include: 'Ready', 'Dropping', 'Disabled', 'Inaccessible'.</td></tr>
	<tr><td>location</td><td>The resource location.</td></tr>
	<tr><td>administrator_login</td><td>Specifies the username of the administrator for this server.</td></tr>
	<tr><td>backup_retention_days</td><td>Backup retention days for the server.</td></tr>
	<tr><td>byok_enforcement</td><td>Status showing whether the server data encryption is enabled with customer-managed keys.</td></tr>
	<tr><td>earliest_restore_date</td><td>Specifies the earliest restore point creation time.</td></tr>
	<tr><td>fully_qualified_domain_name</td><td>The fully qualified domain name of the server.</td></tr>
	<tr><td>geo_redundant_backup</td><td>Indicates whether Geo-redundant is enabled, or not for server backup.</td></tr>
	<tr><td>infrastructure_encryption</td><td>Status showing whether the server enabled infrastructure encryption. Possible values include: 'Enabled', 'Disabled'.</td></tr>
	<tr><td>master_server_id</td><td>The master server id of a replica server.</td></tr>
	<tr><td>minimal_tls_version</td><td>Enforce a minimal Tls version for the server. Possible values include: 'TLS10', 'TLS11', 'TLS12', 'TLSEnforcementDisabled'.</td></tr>
	<tr><td>public_network_access</td><td>Indicates whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'. Possible values include: 'Enabled', 'Disabled'.</td></tr>
	<tr><td>replica_capacity</td><td>The maximum number of replicas that a master server can have.</td></tr>
	<tr><td>replication_role</td><td>The replication role of the server.</td></tr>
	<tr><td>sku_capacity</td><td>The scale up/out capacity, representing server's compute units.</td></tr>
	<tr><td>sku_family</td><td>The family of hardware.</td></tr>
	<tr><td>sku_name</td><td>The name of the sku. For example: 'B_Gen4_1', 'GP_Gen5_8'.</td></tr>
	<tr><td>sku_size</td><td>The size code, to be interpreted by resource as appropriate.</td></tr>
	<tr><td>sku_tier</td><td>The tier of the particular SKU. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'.</td></tr>
	<tr><td>ssl_enforcement</td><td>Enable ssl enforcement or not when connect to server. Possible values include: 'Enabled', 'Disabled'.</td></tr>
	<tr><td>storage_auto_grow</td><td>Indicates whether storage auto grow is enabled, or not.</td></tr>
	<tr><td>storage_mb</td><td>Indicates max storage allowed for a server.</td></tr>
	<tr><td>private_endpoint_connections</td><td>A list of private endpoint connections on a server.</td></tr>
	<tr><td>server_configurations</td><td>The server configurations(parameters) details of the server.</td></tr>
	<tr><td>server_keys</td><td>The server keys of the server.</td></tr>
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