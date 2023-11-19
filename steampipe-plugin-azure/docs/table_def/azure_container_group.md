# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>The resource ID.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the container group. This only appears in the response.</td></tr>
	<tr><td>restart_policy</td><td>Restart policy for all containers within the container group. Possible values include: 'ContainerGroupRestartPolicyAlways', 'ContainerGroupRestartPolicyOnFailure', 'ContainerGroupRestartPolicyNever'.</td></tr>
	<tr><td>sku</td><td>The SKU for a container group. Possible values include: 'ContainerGroupSkuStandard', 'ContainerGroupSkuDedicated'.</td></tr>
	<tr><td>os_type</td><td>The operating system type required by the containers in the container group. Possible values include: 'OperatingSystemTypesWindows', 'OperatingSystemTypesLinux'.</td></tr>
	<tr><td>encryption_properties</td><td>The encryption settings of container registry.</td></tr>
	<tr><td>containers</td><td>The containers within the container group.</td></tr>
	<tr><td>ip_address</td><td>The IP address type of the container group.</td></tr>
	<tr><td>volumes</td><td>The instance view of the container group. Only valid in response.</td></tr>
	<tr><td>instance_view</td><td>The instance view of the container group. Only valid in response.</td></tr>
	<tr><td>diagnostics</td><td>The diagnostic information for a container group.</td></tr>
	<tr><td>subnet_ids</td><td>The subnet resource IDs for a container group.</td></tr>
	<tr><td>dns_config</td><td>The DNS config information for a container group.</td></tr>
	<tr><td>init_containers</td><td>The init containers for a container group.</td></tr>
	<tr><td>image_registry_credentials</td><td>The image registry credentials by which the container group is created from.</td></tr>
	<tr><td>identity</td><td>The identity of the container group.</td></tr>
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