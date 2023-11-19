# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource that is unique within the set of probes used by the load balancer. This name can be used to access the resource.</td></tr>
	<tr><td>id</td><td>The resource ID.</td></tr>
	<tr><td>load_balancer_name</td><td>The friendly name that identifies the load balancer.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the probe resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.</td></tr>
	<tr><td>type</td><td>Type of the resource.</td></tr>
	<tr><td>etag</td><td>A unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>interval_in_seconds</td><td>The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking the instance out of rotation. The default value is 15, the minimum value is 5.</td></tr>
	<tr><td>number_of_probes</td><td>The number of probes where if no response, will result in stopping further traffic from being delivered to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used in Azure.</td></tr>
	<tr><td>port</td><td>The port for communicating the probe. Possible values range from 1 to 65535, inclusive.</td></tr>
	<tr><td>protocol</td><td>The protocol of the end point. If 'Tcp' is specified, a received ACK is required for the probe to be successful. If 'Http' or 'Https' is specified, a 200 OK response from the specifies URI is required for the probe to be successful. Possible values include: 'HTTP', 'TCP', 'HTTPS'.</td></tr>
	<tr><td>request_path</td><td>The URI used for requesting health status from the VM. Path is required if a protocol is set to http. Otherwise, it is not allowed. There is no default value.</td></tr>
	<tr><td>load_balancing_rules</td><td>The load balancer rules that use this probe.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>