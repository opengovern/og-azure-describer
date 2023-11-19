# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The resource id.</td></tr>
	<tr><td>name</td><td>The resource name.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>description</td><td>The security automation description.</td></tr>
	<tr><td>is_enabled</td><td>Indicates whether the security automation is enabled.</td></tr>
	<tr><td>kind</td><td>Kind of the resource.</td></tr>
	<tr><td>etag</td><td>Entity tag is used for comparing two or more entities from the same requested resource.</td></tr>
	<tr><td>actions</td><td>A collection of the actions which are triggered if all the configured rules evaluations, within at least one rule set, are true.</td></tr>
	<tr><td>scopes</td><td>A collection of scopes on which the security automations logic is applied. Supported scopes are the subscription itself or a resource group under that subscription. The automation will only apply on defined scopes.</td></tr>
	<tr><td>sources</td><td>A collection of the source event types which evaluate the security automation set of rules.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tags</td><td>A list of key value pairs that describe the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>region</td><td>The Azure region/location in which the resource is located.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>