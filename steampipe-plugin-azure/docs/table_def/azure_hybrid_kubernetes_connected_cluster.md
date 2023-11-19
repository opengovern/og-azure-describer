# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>id</td><td>The resource ID.</td></tr>
	<tr><td>connectivity_status</td><td>Represents the connectivity status of the connected cluster.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the connected cluster resource.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>agent_public_key_certificate</td><td>Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.</td></tr>
	<tr><td>agent_version</td><td>Version of the agent running on the connected cluster resource.</td></tr>
	<tr><td>created_at</td><td>The timestamp of resource creation (UTC).</td></tr>
	<tr><td>created_by</td><td>The identity that created the resource.</td></tr>
	<tr><td>created_by_type</td><td>The type of identity that created the resource.</td></tr>
	<tr><td>distribution</td><td>The Kubernetes distribution running on this connected cluster.</td></tr>
	<tr><td>infrastructure</td><td>The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.</td></tr>
	<tr><td>kubernetes_version</td><td>The Kubernetes version of the connected cluster resource.</td></tr>
	<tr><td>last_connectivity_time</td><td>Time representing the last instance when heart beat was received from the cluster.</td></tr>
	<tr><td>last_modified_at</td><td>The timestamp of resource last modification (UTC).</td></tr>
	<tr><td>last_modified_by</td><td>The identity that last modified the resource.</td></tr>
	<tr><td>last_modified_by_type</td><td>The type of identity that last modified the resource.</td></tr>
	<tr><td>location</td><td>Location of the resource.</td></tr>
	<tr><td>managed_identity_certificate_expiration_time</td><td>Expiration time of the managed identity certificate.</td></tr>
	<tr><td>offering</td><td>Connected cluster offering.</td></tr>
	<tr><td>total_core_count</td><td>Number of CPU cores present in the connected cluster resource.</td></tr>
	<tr><td>total_node_count</td><td>Number of nodes present in the connected cluster resource.</td></tr>
	<tr><td>extensions</td><td>The extensions of the connected cluster.</td></tr>
	<tr><td>identity</td><td>The identity of the connected cluster.</td></tr>
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