# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the snapshot</td></tr>
	<tr><td>id</td><td>The unique id identifying the resource in subscription</td></tr>
	<tr><td>type</td><td>The type of the resource in Azure</td></tr>
	<tr><td>provisioning_state</td><td>The disk provisioning state</td></tr>
	<tr><td>create_option</td><td>Specifies the possible sources of a disk&#39;s creation</td></tr>
	<tr><td>disk_access_id</td><td>ARM id of the DiskAccess resource for using private endpoints on disks</td></tr>
	<tr><td>disk_encryption_set_id</td><td>ResourceId of the disk encryption set to use for enabling encryption at rest</td></tr>
	<tr><td>disk_size_bytes</td><td>The size of the disk in bytes</td></tr>
	<tr><td>disk_size_gb</td><td>The size of the disk to create</td></tr>
	<tr><td>encryption_setting_collection_enabled</td><td>Specifies whether the encryption is enables, or not</td></tr>
	<tr><td>encryption_setting_version</td><td>Describes what type of encryption is used for the disks</td></tr>
	<tr><td>encryption_type</td><td>The type of the encryption</td></tr>
	<tr><td>gallery_image_reference_id</td><td>A relative uri containing either a Platform Image Repository or user image reference</td></tr>
	<tr><td>gallery_reference_lun</td><td>Specifies the index that indicates which of the data disks in the image to use</td></tr>
	<tr><td>hyperv_generation</td><td>Specifies the hypervisor generation of the Virtual Machine</td></tr>
	<tr><td>image_reference_id</td><td>A relative uri containing either a Platform Image Repository or user image reference</td></tr>
	<tr><td>image_reference_lun</td><td>Specifies the index that indicates which of the data disks in the image to use</td></tr>
	<tr><td>incremental</td><td>Specifies whether a snapshot is incremental, or not</td></tr>
	<tr><td>network_access_policy</td><td>Contains the type of access</td></tr>
	<tr><td>os_type</td><td>Contains the type of operating system</td></tr>
	<tr><td>sku_name</td><td>The snapshot sku name</td></tr>
	<tr><td>sku_tier</td><td>The sku tier</td></tr>
	<tr><td>source_resource_id</td><td>ARM id of the source snapshot or disk</td></tr>
	<tr><td>source_unique_id</td><td>An unique id identifying the source of this resource</td></tr>
	<tr><td>source_uri</td><td>An URI of a blob to be imported into a managed disk</td></tr>
	<tr><td>storage_account_id</td><td>The Azure Resource Manager identifier of the storage account containing the blob to import as a disk</td></tr>
	<tr><td>time_created</td><td>The time when the snapshot was created</td></tr>
	<tr><td>unique_id</td><td>An unique Guid identifying the resource</td></tr>
	<tr><td>upload_size_bytes</td><td>The size of the contents of the upload including the VHD footer</td></tr>
	<tr><td>encryption_settings</td><td>A list of encryption settings, one for each disk volume</td></tr>
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