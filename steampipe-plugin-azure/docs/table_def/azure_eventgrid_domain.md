# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>Fully qualified identifier of the resource.</td></tr>
	<tr><td>type</td><td>The resource type.</td></tr>
	<tr><td>provisioning_state</td><td>Provisioning state of the event grid domain resource. Possible values include: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Canceled', 'Failed'.</td></tr>
	<tr><td>auto_create_topic_with_first_subscription</td><td>This Boolean is used to specify the creation mechanism for 'all' the event grid domain topics associated with this event grid domain resource.</td></tr>
	<tr><td>auto_delete_topic_with_last_subscription</td><td>This Boolean is used to specify the deletion mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.</td></tr>
	<tr><td>created_at</td><td>The timestamp of resource creation (UTC).</td></tr>
	<tr><td>created_by</td><td>The identity that created the resource.</td></tr>
	<tr><td>created_by_type</td><td>The type of identity that created the resource.</td></tr>
	<tr><td>disable_local_auth</td><td>This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the domain.</td></tr>
	<tr><td>endpoint</td><td>Endpoint for the Event Grid Domain Resource which is used for publishing the events.</td></tr>
	<tr><td>identity_type</td><td>The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identity. Possible values include: 'None', 'SystemAssigned', 'UserAssigned', 'SystemAssignedUserAssigned'.</td></tr>
	<tr><td>input_schema</td><td>This determines the format that Event Grid should expect for incoming events published to the Event Grid Domain Resource. Possible values include: 'EventGridSchema', 'CustomEventSchema', 'CloudEventSchemaV10'.</td></tr>
	<tr><td>last_modified_at</td><td>The timestamp of resource last modification (UTC).</td></tr>
	<tr><td>last_modified_by</td><td>The identity that last modified the resource.</td></tr>
	<tr><td>last_modified_by_type</td><td>The type of identity that last modified the resource.</td></tr>
	<tr><td>location</td><td>Location of the resource.</td></tr>
	<tr><td>principal_id</td><td>The principal ID of resource identity.</td></tr>
	<tr><td>public_network_access</td><td>This determines if traffic is allowed over public network. By default it is enabled.</td></tr>
	<tr><td>sku_name</td><td>Name of this SKU. Possible values include: 'Basic', 'Standard'.</td></tr>
	<tr><td>user_assigned_identities</td><td>The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the eventgrid domain.</td></tr>
	<tr><td>inbound_ip_rules</td><td>This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.</td></tr>
	<tr><td>input_schema_mapping</td><td>Information about the InputSchemaMapping which specified the info about mapping event payload.</td></tr>
	<tr><td>private_endpoint_connections</td><td>List of private endpoint connections.</td></tr>
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