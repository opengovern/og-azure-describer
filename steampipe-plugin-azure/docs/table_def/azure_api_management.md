# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>A friendly name that identifies an API management service.</td></tr>
	<tr><td>id</td><td>Contains ID to identify an API management service uniquely.</td></tr>
	<tr><td>provisioning_state</td><td>The current provisioning state of the API management service. Possible values include: 'Created', 'Activating', 'Succeeded', 'Updating', 'Failed', 'Stopped', 'Terminating', 'TerminationFailed', 'Deleted'.</td></tr>
	<tr><td>type</td><td>Type of the resource.</td></tr>
	<tr><td>created_at_utc</td><td>Creation UTC date of the API management service.</td></tr>
	<tr><td>developer_portal_url</td><td>Developer Portal endpoint URL of the API management service.</td></tr>
	<tr><td>disable_gateway</td><td>Property only valid for an API management service deployed in multiple locations. This can be used to disable the gateway in master region.</td></tr>
	<tr><td>enable_client_certificate</td><td>Property only meant to be used for Consumption SKU Service. This enforces a client certificate to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the policy on the gateway.</td></tr>
	<tr><td>etag</td><td>An unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>gateway_regional_url</td><td>Gateway URL of the API management service in the default region.</td></tr>
	<tr><td>gateway_url</td><td>Gateway URL of the API management service.</td></tr>
	<tr><td>identity_principal_id</td><td>The principal id of the identity.</td></tr>
	<tr><td>identity_tenant_id</td><td>The client tenant id of the identity.</td></tr>
	<tr><td>identity_type</td><td>The type of identity used for the resource.</td></tr>
	<tr><td>management_api_url</td><td>Management API endpoint URL of the API management service.</td></tr>
	<tr><td>notification_sender_email</td><td>Email address from which the notification will be sent.</td></tr>
	<tr><td>portal_url</td><td>Publisher portal endpoint URL of the API management service.</td></tr>
	<tr><td>publisher_email</td><td>Publisher email of the API management service.</td></tr>
	<tr><td>publisher_name</td><td>Publisher name of the API management service.</td></tr>
	<tr><td>restore</td><td>Undelete API management service if it was previously soft-deleted.</td></tr>
	<tr><td>scm_url</td><td>SCM endpoint URL of the API management service.</td></tr>
	<tr><td>sku_capacity</td><td>Capacity of the SKU (number of deployed units of the SKU)</td></tr>
	<tr><td>sku_name</td><td>Name of the Sku</td></tr>
	<tr><td>target_provisioning_state</td><td>The provisioning state of the API management service, which is targeted by the long running operation started on the service.</td></tr>
	<tr><td>virtual_network_configuration_subnet_name</td><td>The name of the subnet.</td></tr>
	<tr><td>virtual_network_configuration_subnet_resource_id</td><td>The full resource ID of a subnet in a virtual network to deploy the API Management service in.</td></tr>
	<tr><td>virtual_network_configuration_id</td><td>The virtual network ID.</td></tr>
	<tr><td>virtual_network_type</td><td>The type of VPN in which API management service needs to be configured in. None (Default Value) means the API management service is not part of any Virtual Network, External means the API management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only. Possible values include: 'None', 'External', 'Internal'</td></tr>
	<tr><td>additional_locations</td><td>Additional datacenter locations of the API management service.</td></tr>
	<tr><td>api_version_constraint</td><td>Control plane APIs version constraint for the API management service.</td></tr>
	<tr><td>certificates</td><td>List of certificates that need to be installed in the API management service.</td></tr>
	<tr><td>custom_properties</td><td>Custom properties of the API management service.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the API management service.</td></tr>
	<tr><td>host_name_configurations</td><td>Custom hostname configuration of the API management service.</td></tr>
	<tr><td>identity_user_assigned_identities</td><td>The list of user identities associated with the resource.</td></tr>
	<tr><td>private_ip_addresses</td><td>Private static load balanced IP addresses of the API management service in primary region which is deployed in an internal virtual network. Available only for 'Basic', 'Standard', 'Premium' and 'Isolated' SKU.</td></tr>
	<tr><td>public_ip_addresses</td><td>Public static load balanced IP addresses of the API management service in primary region. Available only for 'Basic', 'Standard', 'Premium' and 'Isolated' SKU.</td></tr>
	<tr><td>zones</td><td>A list of availability zones denoting where the resource needs to come from.</td></tr>
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