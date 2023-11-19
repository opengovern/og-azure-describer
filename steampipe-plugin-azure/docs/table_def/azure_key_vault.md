# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the vault.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a vault uniquely.</td></tr>
	<tr><td>vault_uri</td><td>Contains URI of the vault for performing operations on keys and secrets.</td></tr>
	<tr><td>type</td><td>Type of the resource.</td></tr>
	<tr><td>create_mode</td><td>The vault&#39;s create mode to indicate whether the vault need to be recovered or not. Possible values include: &#39;default&#39;, &#39;recover&#39;.</td></tr>
	<tr><td>enabled_for_deployment</td><td>Indicates whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.</td></tr>
	<tr><td>enabled_for_disk_encryption</td><td>Indicates whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.</td></tr>
	<tr><td>enabled_for_template_deployment</td><td>Indicates whether Azure Resource Manager is permitted to retrieve secrets from the key vault.</td></tr>
	<tr><td>enable_rbac_authorization</td><td>Property that controls how data actions are authorized.</td></tr>
	<tr><td>purge_protection_enabled</td><td>Indicates whether protection against purge is enabled for this vault.</td></tr>
	<tr><td>soft_delete_enabled</td><td>Indicates whether the &#39;soft delete&#39; functionality is enabled for this key vault.</td></tr>
	<tr><td>soft_delete_retention_in_days</td><td>Contains softDelete data retention days.</td></tr>
	<tr><td>sku_family</td><td>Contains SKU family name.</td></tr>
	<tr><td>sku_name</td><td>SKU name to specify whether the key vault is a standard vault or a premium vault.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</td></tr>
	<tr><td>access_policies</td><td>A list of 0 to 1024 identities that have access to the key vault.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the vault.</td></tr>
	<tr><td>network_acls</td><td>Rules governing the accessibility of the key vault from specific network locations.</td></tr>
	<tr><td>private_endpoint_connections</td><td>List of private endpoint connections associated with the key vault.</td></tr>
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