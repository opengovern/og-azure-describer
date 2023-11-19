# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the storage account.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a storage account uniquely.</td></tr>
	<tr><td>type</td><td>Type of the resource.</td></tr>
	<tr><td>access_tier</td><td>The access tier used for billing.</td></tr>
	<tr><td>kind</td><td>The kind of the resource.</td></tr>
	<tr><td>sku_name</td><td>Contains sku name of the storage account.</td></tr>
	<tr><td>sku_tier</td><td>Contains sku tier of the storage account.</td></tr>
	<tr><td>creation_time</td><td>Creation date and time of the storage account.</td></tr>
	<tr><td>allow_blob_public_access</td><td>Specifies whether allow or disallow public access to all blobs or containers in the storage account.</td></tr>
	<tr><td>blob_change_feed_enabled</td><td>Specifies whether change feed event logging is enabled for the Blob service.</td></tr>
	<tr><td>blob_container_soft_delete_enabled</td><td>Specifies whether DeleteRetentionPolicy is enabled.</td></tr>
	<tr><td>blob_container_soft_delete_retention_days</td><td>Indicates the number of days that the deleted item should be retained.</td></tr>
	<tr><td>blob_restore_policy_days</td><td>Specifies how long the blob can be restored.</td></tr>
	<tr><td>blob_restore_policy_enabled</td><td>Specifies whether blob restore is enabled.</td></tr>
	<tr><td>blob_service_logging</td><td>Specifies the blob service properties for logging access.</td></tr>
	<tr><td>blob_soft_delete_enabled</td><td>Specifies whether DeleteRetentionPolicy is enabled.</td></tr>
	<tr><td>blob_soft_delete_retention_days</td><td>Indicates the number of days that the deleted item should be retained.</td></tr>
	<tr><td>blob_versioning_enabled</td><td>Specifies whether versioning is enabled.</td></tr>
	<tr><td>enable_https_traffic_only</td><td>Allows https traffic only to storage service if sets to true.</td></tr>
	<tr><td>encryption_key_source</td><td>Contains the encryption keySource (provider).</td></tr>
	<tr><td>encryption_key_vault_properties_key_current_version_id</td><td>The object identifier of the current versioned Key Vault Key in use.</td></tr>
	<tr><td>encryption_key_vault_properties_key_name</td><td>The name of KeyVault key.</td></tr>
	<tr><td>encryption_key_vault_properties_key_vault_uri</td><td>The Uri of KeyVault.</td></tr>
	<tr><td>encryption_key_vault_properties_key_version</td><td>The version of KeyVault key.</td></tr>
	<tr><td>encryption_key_vault_properties_last_rotation_time</td><td>Timestamp of last rotation of the Key Vault Key.</td></tr>
	<tr><td>failover_in_progress</td><td>Specifies whether the failover is in progress.</td></tr>
	<tr><td>file_soft_delete_enabled</td><td>Specifies whether DeleteRetentionPolicy is enabled.</td></tr>
	<tr><td>file_soft_delete_retention_days</td><td>Indicates the number of days that the deleted item should be retained.</td></tr>
	<tr><td>is_hns_enabled</td><td>Specifies whether account HierarchicalNamespace is enabled.</td></tr>
	<tr><td>queue_logging_delete</td><td>Specifies whether all delete requests should be logged.</td></tr>
	<tr><td>queue_logging_read</td><td>Specifies whether all read requests should be logged.</td></tr>
	<tr><td>queue_logging_retention_days</td><td>Indicates the number of days that metrics or logging data should be retained.</td></tr>
	<tr><td>queue_logging_retention_enabled</td><td>Specifies whether a retention policy is enabled for the storage service.</td></tr>
	<tr><td>queue_logging_version</td><td>The version of Storage Analytics to configure.</td></tr>
	<tr><td>queue_logging_write</td><td>Specifies whether all write requests should be logged.</td></tr>
	<tr><td>table_logging_read</td><td>Indicates whether all read requests should be logged.</td></tr>
	<tr><td>table_logging_write</td><td>Indicates whether all write requests should be logged.</td></tr>
	<tr><td>table_logging_delete</td><td>Indicates whether all delete requests should be logged.</td></tr>
	<tr><td>table_logging_version</td><td>The version of Analytics to configure.</td></tr>
	<tr><td>table_logging_retention_policy</td><td>The retention policy.</td></tr>
	<tr><td>minimum_tls_version</td><td>Contains the minimum TLS version to be permitted on requests to storage.</td></tr>
	<tr><td>network_rule_bypass</td><td>Specifies whether traffic is bypassed for Logging/Metrics/AzureServices.</td></tr>
	<tr><td>network_rule_default_action</td><td>Specifies the default action of allow or deny when no other rules match.</td></tr>
	<tr><td>primary_blob_endpoint</td><td>Contains the blob endpoint.</td></tr>
	<tr><td>primary_dfs_endpoint</td><td>Contains the dfs endpoint.</td></tr>
	<tr><td>primary_file_endpoint</td><td>Contains the file endpoint.</td></tr>
	<tr><td>primary_location</td><td>Contains the location of the primary data center for the storage account.</td></tr>
	<tr><td>primary_queue_endpoint</td><td>Contains the queue endpoint.</td></tr>
	<tr><td>primary_table_endpoint</td><td>Contains the table endpoint.</td></tr>
	<tr><td>primary_web_endpoint</td><td>Contains the web endpoint.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state of the storage account resource.</td></tr>
	<tr><td>require_infrastructure_encryption</td><td>Specifies whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest.</td></tr>
	<tr><td>secondary_location</td><td>Contains the location of the geo-replicated secondary for the storage account.</td></tr>
	<tr><td>diagnostic_settings</td><td>A list of active diagnostic settings for the storage account.</td></tr>
	<tr><td>encryption_scope</td><td>Encryption scope details for the storage account.</td></tr>
	<tr><td>encryption_services</td><td>A list of services which support encryption.</td></tr>
	<tr><td>lifecycle_management_policy</td><td>The managementpolicy associated with the specified storage account.</td></tr>
	<tr><td>network_ip_rules</td><td>A list of IP ACL rules.</td></tr>
	<tr><td>private_endpoint_connections</td><td>A list of private endpoint connection associated with the specified storage account.</td></tr>
	<tr><td>virtual_network_rules</td><td>A list of virtual network rules.</td></tr>
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