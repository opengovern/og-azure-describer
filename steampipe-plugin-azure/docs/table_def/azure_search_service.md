# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>Fully qualified resource ID for the resource.</td></tr>
	<tr><td>provisioning_state</td><td>The state of the last provisioning operation performed on the search service.</td></tr>
	<tr><td>status</td><td>The status of the search service. Possible values include: &#39;running&#39;, deleting&#39;, &#39;provisioning&#39;, &#39;degraded&#39;, &#39;disabled&#39;, &#39;error&#39; etc.</td></tr>
	<tr><td>status_details</td><td>The details of the search service status.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>hosting_mode</td><td>Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either &#39;default&#39; or &#39;highDensity&#39;. For all other SKUs, this value must be &#39;default&#39;. Possible values include: &#39;Default&#39;, &#39;HighDensity&#39;.</td></tr>
	<tr><td>partition_count</td><td>The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For &#39;standard3&#39; services with hostingMode set to &#39;highDensity&#39;, the allowed values are between 1 and 3.</td></tr>
	<tr><td>public_network_access</td><td>This value can be set to &#39;enabled&#39; to avoid breaking changes on existing customer resources and templates. If set to &#39;disabled&#39;, traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method. Possible values include: &#39;Enabled&#39;, &#39;Disabled&#39;.</td></tr>
	<tr><td>replica_count</td><td>The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.</td></tr>
	<tr><td>sku_name</td><td>The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new search service.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the search service.</td></tr>
	<tr><td>identity</td><td>The identity of the resource.</td></tr>
	<tr><td>network_rule_set</td><td>Network specific rules that determine how the azure cognitive search service may be reached.</td></tr>
	<tr><td>private_endpoint_connections</td><td>The list of private endpoint connections to the azure cognitive search service.</td></tr>
	<tr><td>shared_private_link_resources</td><td>The list of shared private link resources managed by the azure cognitive search service.</td></tr>
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