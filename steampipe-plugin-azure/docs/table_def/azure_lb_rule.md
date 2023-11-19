# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource that is unique within the set of load balancing rules used by the load balancer. This name can be used to access the resource.</td></tr>
	<tr><td>id</td><td>The resource ID.</td></tr>
	<tr><td>load_balancer_name</td><td>The friendly name that identifies the load balancer.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the load balancing rule resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.</td></tr>
	<tr><td>type</td><td>Type of the resource.</td></tr>
	<tr><td>etag</td><td>A unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>backend_address_pool_id</td><td>A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend IPs.</td></tr>
	<tr><td>backend_port</td><td>The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note that value 0 enables 'Any Port'.</td></tr>
	<tr><td>disable_outbound_snat</td><td>Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the frontend of the load balancing rule.</td></tr>
	<tr><td>enable_floating_ip</td><td>Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.</td></tr>
	<tr><td>enable_tcp_reset</td><td>Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.</td></tr>
	<tr><td>frontend_ip_configuration_id</td><td>A reference to frontend IP addresses.</td></tr>
	<tr><td>frontend_port</td><td>The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 0 and 65534. Note that value 0 enables 'Any Port'.</td></tr>
	<tr><td>idle_timeout_in_minutes</td><td>The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.</td></tr>
	<tr><td>load_distribution</td><td>The load distribution policy for this rule. Possible values include: 'Default', 'SourceIP', 'SourceIPProtocol'.</td></tr>
	<tr><td>probe_id</td><td>The reference to the load balancer probe used by the load balancing rule.</td></tr>
	<tr><td>protocol</td><td>The reference to the transport protocol used by the load balancing rule. Possible values include: 'UDP', 'TCP', 'All'.</td></tr>
	<tr><td>backend_address_pools</td><td>An array of references to pool of DIPs.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>