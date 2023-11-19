# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the cluster.</td></tr>
	<tr><td>id</td><td>The ID of the cluster.</td></tr>
	<tr><td>type</td><td>The type of the cluster.</td></tr>
	<tr><td>location</td><td>The location where the cluster is created.</td></tr>
	<tr><td>azure_portal_fqdn</td><td>FQDN for the master pool which used by proxy config.</td></tr>
	<tr><td>disk_encryption_set_id</td><td>ResourceId of the disk encryption set to use for enabling encryption at rest.</td></tr>
	<tr><td>dns_prefix</td><td>DNS prefix specified when creating the managed cluster.</td></tr>
	<tr><td>enable_pod_security_policy</td><td>Whether to enable Kubernetes pod security policy (preview).</td></tr>
	<tr><td>enable_rbac</td><td>Whether to enable Kubernetes Role-Based Access Control.</td></tr>
	<tr><td>fqdn</td><td>FQDN for the master pool.</td></tr>
	<tr><td>fqdn_subdomain</td><td>FQDN subdomain specified when creating private cluster with custom private dns zone.</td></tr>
	<tr><td>kubernetes_version</td><td>Version of Kubernetes specified when creating the managed cluster.</td></tr>
	<tr><td>max_agent_pools</td><td>The max number of agent pools for the managed cluster.</td></tr>
	<tr><td>node_resource_group</td><td>Name of the resource group containing agent pool nodes.</td></tr>
	<tr><td>private_fqdn</td><td>FQDN of private cluster.</td></tr>
	<tr><td>provisioning_state</td><td>The current deployment or provisioning state.</td></tr>
	<tr><td>aad_profile</td><td>Profile of Azure Active Directory configuration.</td></tr>
	<tr><td>addon_profiles</td><td>Profile of managed cluster add-on.</td></tr>
	<tr><td>agent_pool_profiles</td><td>Properties of the agent pool.</td></tr>
	<tr><td>api_server_access_profile</td><td>Access profile for managed cluster API server.</td></tr>
	<tr><td>auto_scaler_profile</td><td>Parameters to be applied to the cluster-autoscaler when enabled.</td></tr>
	<tr><td>auto_upgrade_profile</td><td>Profile of auto upgrade configuration.</td></tr>
	<tr><td>identity</td><td>The identity of the managed cluster, if configured.</td></tr>
	<tr><td>identity_profile</td><td>Identities associated with the cluster.</td></tr>
	<tr><td>linux_profile</td><td>Profile for Linux VMs in the container service cluster.</td></tr>
	<tr><td>network_profile</td><td>Profile of network configuration.</td></tr>
	<tr><td>pod_identity_profile</td><td>Profile of managed cluster pod identity.</td></tr>
	<tr><td>power_state</td><td>Represents the Power State of the cluster.</td></tr>
	<tr><td>service_principal_profile</td><td>Information about a service principal identity for the cluster to use for manipulating Azure APIs.</td></tr>
	<tr><td>sku</td><td>The managed cluster SKU.</td></tr>
	<tr><td>windows_profile</td><td>Profile for Windows VMs in the container service cluster.</td></tr>
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