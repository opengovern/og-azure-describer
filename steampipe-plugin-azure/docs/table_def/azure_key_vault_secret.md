# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the secret.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a secret uniquely.</td></tr>
	<tr><td>vault_name</td><td>The friendly name that identifies the vault.</td></tr>
	<tr><td>enabled</td><td>Indicates whether the secret is enabled, or not.</td></tr>
	<tr><td>managed</td><td>Indicates whether the secret's lifetime is managed by key vault, or not.</td></tr>
	<tr><td>content_type</td><td>Specifies the type of the secret value such as a password.</td></tr>
	<tr><td>created_at</td><td>Specifies the time when the secret is created.</td></tr>
	<tr><td>expires_at</td><td>Specifies the time when the secret will expire.</td></tr>
	<tr><td>kid</td><td>If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV certificate.</td></tr>
	<tr><td>not_before</td><td>Specifies the time before which the secret is not usable.</td></tr>
	<tr><td>recoverable_days</td><td>Specifies the soft delete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.</td></tr>
	<tr><td>recovery_level</td><td>The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then the object can be permanently deleted by a privileged user; otherwise, only the system can purge the object at the end of the retention interval.</td></tr>
	<tr><td>updated_at</td><td>Specifies the time when the secret was last updated.</td></tr>
	<tr><td>value</td><td>Specifies the secret value.</td></tr>
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