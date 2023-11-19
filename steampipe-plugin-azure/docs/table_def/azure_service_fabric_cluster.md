# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>Azure resource name.</td></tr>
	<tr><td>id</td><td>Azure resource identifier.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the cluster resource. Possible values include: &#39;Updating&#39;, &#39;Succeeded&#39;, &#39;Failed&#39;, &#39;Canceled&#39;.</td></tr>
	<tr><td>type</td><td>Azure resource type.</td></tr>
	<tr><td>cluster_code_version</td><td>The service fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to &#39;Manual&#39;.</td></tr>
	<tr><td>cluster_endpoint</td><td>The azure resource provider endpoint. A system service in the cluster connects to this  endpoint.</td></tr>
	<tr><td>cluster_id</td><td>A service generated unique identifier for the cluster resource.</td></tr>
	<tr><td>cluster_state</td><td>The current state of the cluster. Possible values include: &#39;WaitingForNodes&#39;, &#39;Deploying&#39;, &#39;BaselineUpgrade&#39;, &#39;UpdatingUserConfiguration&#39;, &#39;UpdatingUserCertificate&#39;, &#39;UpdatingInfrastructure&#39;, &#39;EnforcingClusterVersion&#39;, &#39;UpgradeServiceUnreachable&#39;, &#39;AutoScale&#39;, &#39;Ready&#39;.</td></tr>
	<tr><td>event_store_service_enabled</td><td>Indicates if the event store service is enabled.</td></tr>
	<tr><td>etag</td><td>Azure resource etag.</td></tr>
	<tr><td>management_endpoint</td><td>The http management endpoint of the cluster.</td></tr>
	<tr><td>reliability_level</td><td>The reliability level sets the replica set size of system services. Possible values include: &#39;None&#39;, &#39;Bronze&#39;, &#39;Silver&#39;, &#39;Gold&#39;, &#39;Platinum&#39;.</td></tr>
	<tr><td>upgrade_mode</td><td>The upgrade mode of the cluster when new service fabric runtime version is available. Possible values include: &#39;Automatic&#39;, &#39;Manual&#39;.</td></tr>
	<tr><td>vm_image</td><td>The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.</td></tr>
	<tr><td>add_on_features</td><td>The list of add-on features to enable in the cluster.</td></tr>
	<tr><td>available_cluster_versions</td><td>The service fabric runtime versions available for this cluster.</td></tr>
	<tr><td>azure_active_directory</td><td>The azure active directory authentication settings of the cluster.</td></tr>
	<tr><td>certificate</td><td>The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.</td></tr>
	<tr><td>certificate_common_names</td><td>Describes a list of server certificates referenced by common name that are used to secure the cluster.</td></tr>
	<tr><td>client_certificate_common_names</td><td>The list of client certificates referenced by common name that are allowed to manage the cluster.</td></tr>
	<tr><td>client_certificate_thumbprints</td><td>The list of client certificates referenced by thumbprint that are allowed to manage the cluster.</td></tr>
	<tr><td>diagnostics_storage_account_config</td><td>The storage account information for storing service fabric diagnostic logs.</td></tr>
	<tr><td>fabric_settings</td><td>The list of custom fabric settings to configure the cluster.</td></tr>
	<tr><td>node_types</td><td>The list of node types in the cluster.</td></tr>
	<tr><td>reverse_proxy_certificate</td><td>The server certificate used by reverse proxy.</td></tr>
	<tr><td>reverse_proxy_certificate_common_names</td><td>Describes a list of server certificates referenced by common name that are used to secure the cluster.</td></tr>
	<tr><td>upgrade_description</td><td>The policy to use when upgrading the cluster.</td></tr>
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