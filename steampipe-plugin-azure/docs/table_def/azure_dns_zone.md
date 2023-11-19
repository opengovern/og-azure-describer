# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the DNS zone.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a DNS zone uniquely.</td></tr>
	<tr><td>etag</td><td>An unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>type</td><td>The resource type of the DNS zone.</td></tr>
	<tr><td>max_number_of_record_sets</td><td>The maximum number of record sets that can be created in this DNS zone.</td></tr>
	<tr><td>max_number_of_records_per_record_set</td><td>The maximum number of records per record set that can be created in this DNS zone.</td></tr>
	<tr><td>number_of_record_sets</td><td>The current number of record sets in this DNS zone.</td></tr>
	<tr><td>name_servers</td><td>The name servers for this DNS zone.</td></tr>
	<tr><td>zone_type</td><td>The type of this DNS zone (always `Public`, see `azure_private_dns_zone` table for private DNS zones).</td></tr>
	<tr><td>registration_virtual_networks</td><td>A list of references to virtual networks that register hostnames in this DNS zone.</td></tr>
	<tr><td>resolution_virtual_networks</td><td>A list of references to virtual networks that resolve records in this DNS zone.</td></tr>
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