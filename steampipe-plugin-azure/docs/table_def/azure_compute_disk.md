# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>Name of the disk</td></tr>
	<tr><td>id</td><td>The unique id identifying the resource in subscription</td></tr>
	<tr><td>type</td><td>The type of the resource in Azure</td></tr>
	<tr><td>provisioning_state</td><td>The disk provisioning state</td></tr>
	<tr><td>managed_by</td><td>A relative URI containing the ID of the VM that has the disk attached</td></tr>
	<tr><td>sku_name</td><td>The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS</td></tr>
	<tr><td>sku_tier</td><td>The sku tier</td></tr>
	<tr><td>time_created</td><td>The time when the disk was created</td></tr>
	<tr><td>unique_id</td><td>Unique Guid identifying the resource</td></tr>
	<tr><td>disk_access_id</td><td>ARM id of the DiskAccess resource for using private endpoints on disks</td></tr>
	<tr><td>disk_size_bytes</td><td>The size of the disk in bytes</td></tr>
	<tr><td>disk_size_gb</td><td>If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size</td></tr>
	<tr><td>disk_state</td><td>This enumerates the possible state of the disk</td></tr>
	<tr><td>hyper_v_generation</td><td>The hypervisor generation of the Virtual Machine. Applicable to OS disks only</td></tr>
	<tr><td>disk_iops_read_only</td><td>The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes</td></tr>
	<tr><td>disk_iops_read_write</td><td>The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes</td></tr>
	<tr><td>disk_iops_mbps_read_only</td><td>The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10</td></tr>
	<tr><td>disk_iops_mbps_read_write</td><td>The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10</td></tr>
	<tr><td>max_shares</td><td>The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time</td></tr>
	<tr><td>os_type</td><td>The Operating System type</td></tr>
	<tr><td>encryption_settings_collection_enabled</td><td>Shows the status of the encryption settings for the disk</td></tr>
	<tr><td>encryption_settings_collection_version</td><td>Describes the type of encryption is used for the disks. '1.0' corresponds to Azure Disk Encryption with AAD app. '1.1' corresponds to Azure Disk Encryption</td></tr>
	<tr><td>encryption_disk_encryption_set_id</td><td>ResourceId of the disk encryption set to use for enabling encryption at rest</td></tr>
	<tr><td>encryption_type</td><td>The type of key used to encrypt the data of the disk</td></tr>
	<tr><td>network_access_policy</td><td>Policy for accessing the disk via network</td></tr>
	<tr><td>creation_data_option</td><td>This enumerates the possible sources of a disk's creation</td></tr>
	<tr><td>creation_data_storage_account_id</td><td>The Azure Resource Manager identifier of the storage account containing the blob to import as a disk</td></tr>
	<tr><td>creation_data_source_uri</td><td>The URI of a blob to be imported into a managed disk</td></tr>
	<tr><td>creation_data_source_resource_id</td><td>The ARM id of the source snapshot or disk</td></tr>
	<tr><td>creation_data_source_unique_id</td><td>An unique id identifying the source of this resource</td></tr>
	<tr><td>creation_data_upload_size_bytes</td><td>This is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer)</td></tr>
	<tr><td>creation_data_image_reference_id</td><td>A relative uri containing either a Platform Image Repository or user image reference</td></tr>
	<tr><td>creation_data_image_reference_lun</td><td>If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null</td></tr>
	<tr><td>creation_data_gallery_image_reference_id</td><td>The ARM id of the shared galley image version from which disk was created</td></tr>
	<tr><td>creation_data_gallery_image_reference_lun</td><td>An index that indicates which of the data disks in the image to use, if the disk is created from an image's data disk</td></tr>
	<tr><td>encryption_settings_collection_settings</td><td>A collection of encryption settings, one for each disk volume</td></tr>
	<tr><td>share_info</td><td>Details of the list of all VMs that have the disk attached</td></tr>
	<tr><td>zones</td><td>The Logical zone list for Disk</td></tr>
	<tr><td>managed_by_extended</td><td>List of relative URIs containing the IDs of the VMs that have the disk attached</td></tr>
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