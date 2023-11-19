# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The name of the resource.</td></tr>
	<tr><td>storage_account_name</td><td>The name of the storage account.</td></tr>
	<tr><td>id</td><td>Fully qualified resource ID for the resource.</td></tr>
	<tr><td>type</td><td>The type of the resource.</td></tr>
	<tr><td>access_tier</td><td>Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool.</td></tr>
	<tr><td>access_tier_change_time</td><td>Indicates the last modification time for share access tier.</td></tr>
	<tr><td>access_tier_status</td><td>Indicates if there is a pending transition for access tier.</td></tr>
	<tr><td>last_modified_time</td><td>Returns the date and time the share was last modified.</td></tr>
	<tr><td>deleted</td><td>Indicates whether the share was deleted.</td></tr>
	<tr><td>deleted_time</td><td>The deleted time if the share was deleted.</td></tr>
	<tr><td>enabled_protocols</td><td>The authentication protocol that is used for the file share. Can only be specified when creating a share. Possible values include: 'SMB', 'NFS'.</td></tr>
	<tr><td>remaining_retention_days</td><td>Remaining retention days for share that was soft deleted.</td></tr>
	<tr><td>root_squash</td><td>The property is for NFS share only. The default is NoRootSquash. Possible values include: 'NoRootSquash', 'RootSquash', 'AllSquash'.</td></tr>
	<tr><td>share_quota</td><td>The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.</td></tr>
	<tr><td>share_usage_bytes</td><td>The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.</td></tr>
	<tr><td>version</td><td>The version of the share.</td></tr>
	<tr><td>metadata</td><td>A name-value pair to associate with the share as metadata.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>