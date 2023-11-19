# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The resource name.</td></tr>
	<tr><td>friendly_name</td><td>The friendly name for this workspace. This name in mutable.</td></tr>
	<tr><td>id</td><td>The resource identifier.</td></tr>
	<tr><td>provisioning_state</td><td>The current deployment state of workspace resource, The provisioningState is to indicate states for resource provisioning. Possible values include: &#39;Unknown&#39;, &#39;Updating&#39;, &#39;Creating&#39;, &#39;Deleting&#39;, &#39;Succeeded&#39;, &#39;Failed&#39;, &#39;Canceled&#39;.</td></tr>
	<tr><td>creation_time</td><td>The creation time for this workspace resource.</td></tr>
	<tr><td>workspace_id</td><td>The immutable id associated with this workspace.</td></tr>
	<tr><td>application_insights</td><td>ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created.</td></tr>
	<tr><td>container_registry</td><td>ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created.</td></tr>
	<tr><td>description</td><td>The description of this workspace.</td></tr>
	<tr><td>discovery_url</td><td>ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created.</td></tr>
	<tr><td>hbi_workspace</td><td>The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service.</td></tr>
	<tr><td>key_vault</td><td>ARM id of the key vault associated with this workspace, This cannot be changed once the workspace has been created.</td></tr>
	<tr><td>location</td><td>The location of the resource. This cannot be changed after the resource is created.</td></tr>
	<tr><td>service_provisioned_resource_group</td><td>The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace.</td></tr>
	<tr><td>sku_name</td><td>Name of the sku.</td></tr>
	<tr><td>sku_tier</td><td>Tier of the sku like Basic or Enterprise.</td></tr>
	<tr><td>storage_account</td><td>ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created.</td></tr>
	<tr><td>type</td><td>The resource type.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the azure ML workspace.</td></tr>
	<tr><td>encryption</td><td>The encryption settings of Azure ML workspace.</td></tr>
	<tr><td>identity</td><td>The identity of the resource.</td></tr>
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