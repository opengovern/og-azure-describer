# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>Fully qualified resource ID for the resource.</td></tr>
	<tr><td>kind</td><td>The kind of the resource.</td></tr>
	<tr><td>provisioning_state</td><td>The status of the cognitive services account at the time the operation was called. Possible values include: 'Accepted', 'Creating', 'Deleting', 'Moving', 'Failed', 'Succeeded', 'ResolvingDNS'.</td></tr>
	<tr><td>type</td><td>The type of the resource. E.g. 'Microsoft.Compute/virtualMachines' or 'Microsoft.Storage/storageAccounts'.</td></tr>
	<tr><td>custom_sub_domain_name</td><td>The subdomain name used for token-based authentication.</td></tr>
	<tr><td>date_created</td><td>The date of cognitive services account creation.</td></tr>
	<tr><td>disable_local_auth</td><td>Checks if local auth is disabled for the resource.</td></tr>
	<tr><td>endpoint</td><td>The endpoint of the created account.</td></tr>
	<tr><td>etag</td><td>The resource etag.</td></tr>
	<tr><td>is_migrated</td><td>Checks if the resource is migrated from an existing key.</td></tr>
	<tr><td>migration_token</td><td>The resource migration token.</td></tr>
	<tr><td>public_network_access</td><td>Whether or not public endpoint access is allowed for this account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'. Possible values include: 'Enabled', 'Disabled'.</td></tr>
	<tr><td>restore</td><td>Checks if restore is enabled for the resource.</td></tr>
	<tr><td>restrict_outbound_network_access</td><td>Checks if outbound network access is restricted for the resource.</td></tr>
	<tr><td>allowed_fqdn_list</td><td>The allowed FQDN list for the resource.</td></tr>
	<tr><td>api_properties</td><td>The api properties for special APIs.</td></tr>
	<tr><td>call_rate_limit</td><td>The call rate limit of the resource.</td></tr>
	<tr><td>capabilities</td><td>The capabilities of the cognitive services account. Each item indicates the capability of a specific feature. The values are read-only and for reference only.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the cognitive service account.</td></tr>
	<tr><td>encryption</td><td>The encryption properties for the resource.</td></tr>
	<tr><td>endpoints</td><td>All endpoints of the cognitive services account.</td></tr>
	<tr><td>identity</td><td>The identity for the resource.</td></tr>
	<tr><td>network_acls</td><td>A collection of rules governing the accessibility from specific network locations.</td></tr>
	<tr><td>private_endpoint_connections</td><td>The private endpoint connection associated with the cognitive services account.</td></tr>
	<tr><td>quota_limit</td><td>The quota limit of the resource.</td></tr>
	<tr><td>sku</td><td>The resource model definition representing SKU.</td></tr>
	<tr><td>sku_change_info</td><td>Sku change info of the resource.</td></tr>
	<tr><td>system_data</td><td>The metadata pertaining to creation and last modification of the resource.</td></tr>
	<tr><td>user_owned_storage</td><td>The storage accounts for the resource.</td></tr>
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