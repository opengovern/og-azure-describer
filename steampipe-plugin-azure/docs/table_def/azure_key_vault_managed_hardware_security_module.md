# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the managed HSM Pool.</td></tr>
	<tr><td>id</td><td>The Azure Resource Manager resource ID for the managed HSM Pool.</td></tr>
	<tr><td>type</td><td>The resource type of the managed HSM Pool.</td></tr>
	<tr><td>provisioning_state</td><td>Provisioning state. Possible values include: 'ProvisioningStateSucceeded', 'ProvisioningStateProvisioning', 'ProvisioningStateFailed', 'ProvisioningStateUpdating', 'ProvisioningStateDeleting', 'ProvisioningStateActivated', 'ProvisioningStateSecurityDomainRestore', 'ProvisioningStateRestoring'.</td></tr>
	<tr><td>hsm_uri</td><td>The URI of the managed hsm pool for performing operations on keys.</td></tr>
	<tr><td>enable_soft_delete</td><td>Property to specify whether the 'soft delete' functionality is enabled for this managed HSM pool. If it's not set to any value(true or false) when creating new managed HSM pool, it will be set to true by default. Once set to true, it cannot be reverted to false.</td></tr>
	<tr><td>soft_delete_retention_in_days</td><td>Indicates softDelete data retention days. It accepts >=7 and <=90.</td></tr>
	<tr><td>enable_purge_protection</td><td>Property specifying whether protection against purge is enabled for this managed HSM pool. Setting this property to true activates protection against purge for this managed HSM pool and its content - only the Managed HSM service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible.</td></tr>
	<tr><td>status_message</td><td>Resource Status Message.</td></tr>
	<tr><td>create_mode</td><td>The create mode to indicate whether the resource is being created or is being recovered from a deleted resource. Possible values include: 'CreateModeRecover', 'CreateModeDefault'.</td></tr>
	<tr><td>sku_family</td><td>Contains SKU family name.</td></tr>
	<tr><td>sku_name</td><td>SKU name to specify whether the key vault is a standard vault or a premium vault.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the managed HSM.</td></tr>
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