# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the key version.</td></tr>
	<tr><td>key_name</td><td>The friendly name that identifies the key.</td></tr>
	<tr><td>key_id</td><td>Contains ID to identify a key uniquely.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a key version uniquely.</td></tr>
	<tr><td>vault_name</td><td>The friendly name that identifies the vault.</td></tr>
	<tr><td>enabled</td><td>Indicates whether the key version is enabled, or not.</td></tr>
	<tr><td>key_type</td><td>The type of the key. Possible values are: 'EC', 'ECHSM', 'RSA', 'RSAHSM'.</td></tr>
	<tr><td>created_at</td><td>Specifies the time when the key version is created.</td></tr>
	<tr><td>curve_name</td><td>The elliptic curve name. Possible values are: 'P256', 'P384', 'P521', 'P256K'.</td></tr>
	<tr><td>expires_at</td><td>Specifies the time when the key version wil expire.</td></tr>
	<tr><td>key_size</td><td>The key size in bits.</td></tr>
	<tr><td>key_uri</td><td>The URI to retrieve the current version of the key.</td></tr>
	<tr><td>key_uri_with_version</td><td>The URI to retrieve the specific version of the key.</td></tr>
	<tr><td>location</td><td>Azure location of the key vault resource.</td></tr>
	<tr><td>not_before</td><td>Specifies the time before which the key version is not usable.</td></tr>
	<tr><td>recovery_level</td><td>The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then the object can be permanently deleted by a privileged user; otherwise, only the system can purge the object at the end of the retention interval.</td></tr>
	<tr><td>type</td><td>Type of the resource</td></tr>
	<tr><td>updated_at</td><td>Specifies the time when the key was last updated.</td></tr>
	<tr><td>key_ops</td><td>A list of key operations.</td></tr>
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