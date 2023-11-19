# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the circuit.</td></tr>
	<tr><td>id</td><td>Resource ID.</td></tr>
	<tr><td>etag</td><td>An unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>sku_name</td><td>The name of the SKU.</td></tr>
	<tr><td>sku_tier</td><td>The tier of the SKU. Possible values include: 'Standard', 'Premium', 'Basic', 'Local'.</td></tr>
	<tr><td>sku_family</td><td>The family of the SKU. Possible values include: 'UnlimitedData', 'MeteredData'.</td></tr>
	<tr><td>allow_classic_operations</td><td>Allow classic operations.</td></tr>
	<tr><td>circuit_provisioning_state</td><td>The CircuitProvisioningState state of the resource.</td></tr>
	<tr><td>service_provider_provisioning_state</td><td>The ServiceProviderProvisioningState state of the resource. Possible values include: 'NotProvisioned', 'Provisioning', 'Provisioned', 'Deprovisioning'.</td></tr>
	<tr><td>authorizations</td><td>The list of authorizations.</td></tr>
	<tr><td>peerings</td><td>The list of peerings.</td></tr>
	<tr><td>service_key</td><td>The ServiceKey.</td></tr>
	<tr><td>service_provider_notes</td><td>The ServiceProviderNotes.</td></tr>
	<tr><td>service_provider_properties</td><td>The ServiceProviderProperties.</td></tr>
	<tr><td>express_route_port</td><td>The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.</td></tr>
	<tr><td>bandwidth_in_gbps</td><td>The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the express route circuit resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.</td></tr>
	<tr><td>global_reach_enabled</td><td>Flag denoting global reach status.</td></tr>
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