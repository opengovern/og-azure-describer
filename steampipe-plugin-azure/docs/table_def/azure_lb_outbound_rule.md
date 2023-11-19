# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can be used to access the resource.</td></tr>
	<tr><td>id</td><td>The resource ID.</td></tr>
	<tr><td>load_balancer_name</td><td>The friendly name that identifies the load balancer.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the outbound rule resource. Possible values include: &#39;ProvisioningStateSucceeded&#39;, &#39;ProvisioningStateUpdating&#39;, &#39;ProvisioningStateDeleting&#39;, &#39;ProvisioningStateFailed&#39;.</td></tr>
	<tr><td>type</td><td>Type of the resource.</td></tr>
	<tr><td>etag</td><td>A unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>allocated_outbound_ports</td><td>The number of outbound ports to be used for NAT.</td></tr>
	<tr><td>enable_tcp_reset</td><td>Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.</td></tr>
	<tr><td>idle_timeout_in_minutes</td><td>The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.</td></tr>
	<tr><td>protocol</td><td>The protocol for the outbound rule in load balancer. Possible values include: &#39;LoadBalancerOutboundRuleProtocolTCP&#39;, &#39;LoadBalancerOutboundRuleProtocolUDP&#39;, &#39;LoadBalancerOutboundRuleProtocolAll&#39;.</td></tr>
	<tr><td>backend_address_pools</td><td>A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend IPs.</td></tr>
	<tr><td>frontend_ip_configurations</td><td>The Frontend IP addresses of the load balancer.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>