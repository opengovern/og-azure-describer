# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the blob.</td></tr>
	<tr><td>storage_account_name</td><td>The friendly name that identifies the storage account, in which the blob is located.</td></tr>
	<tr><td>container_name</td><td>The friendly name that identifies the container, in which the blob has been uploaded.</td></tr>
	<tr><td>type</td><td>Specifies the type of the blob.</td></tr>
	<tr><td>is_snapshot</td><td>Specifies whether the resource is snapshot of a blob, or not.</td></tr>
	<tr><td>access_tier</td><td>The tier of the blob.</td></tr>
	<tr><td>creation_time</td><td>Indicates the time, when the blob was uploaded.</td></tr>
	<tr><td>deleted</td><td>Specifies whether the blob was deleted, or not.</td></tr>
	<tr><td>deleted_time</td><td>Specifies the deletion time of blob container.</td></tr>
	<tr><td>etag</td><td>An unique read-only string that changes whenever the resource is updated.</td></tr>
	<tr><td>last_modified</td><td>Specifies the date and time the container was last modified.</td></tr>
	<tr><td>snapshot</td><td>Specifies the time, when the snapshot is taken.</td></tr>
	<tr><td>version_id</td><td>Specifies the version id.</td></tr>
	<tr><td>server_encrypted</td><td>Indicates whether the blob is encrypted on the server, or not.</td></tr>
	<tr><td>encryption_scope</td><td>The name of the encryption scope under which the blob is encrypted.</td></tr>
	<tr><td>encryption_key_sha256</td><td>The SHA-256 hash of the provided encryption key.</td></tr>
	<tr><td>is_current_version</td><td>Specifies whether the blob container was deleted, or not.</td></tr>
	<tr><td>access_tier_change_time</td><td>Species the time, when the access tier has been updated.</td></tr>
	<tr><td>access_tier_inferred</td><td>Indicates whether the access tier was inferred by the service.</td></tr>
	<tr><td>blob_sequence_number</td><td>Specifies the sequence number for page blob used for coordinating concurrent writes.</td></tr>
	<tr><td>content_length</td><td>Specifies the size of the content returned.</td></tr>
	<tr><td>cache_control</td><td>Indicates the cache control specified for the blob.</td></tr>
	<tr><td>content_disposition</td><td>Specifies additional information about how to process the response payload, and also can be used to attach additional metadata.</td></tr>
	<tr><td>content_encoding</td><td>Indicates content encoding specified for the blob.</td></tr>
	<tr><td>content_language</td><td>Indicates content language specified for the blob.</td></tr>
	<tr><td>content_md5</td><td>If the content_md5 has been set for the blob, this response header is stored so that the client can check for message content integrity.</td></tr>
	<tr><td>content_type</td><td>Specifies the content type specified for the blob. If no content type was specified, the default content type is application/octet-stream.</td></tr>
	<tr><td>copy_completion_time</td><td>Conclusion time of the last attempted Copy Blob operation where this blob was the destination blob.</td></tr>
	<tr><td>copy_id</td><td>A String identifier for the last attempted Copy Blob operation where this blob was the destination blob.</td></tr>
	<tr><td>copy_progress</td><td>Contains the number of bytes copied and the total bytes in the source in the last attempted Copy Blob operation where this blob was the destination blob.</td></tr>
	<tr><td>copy_source</td><td>An URL up to 2 KB in length that specifies the source blob used in the last attempted Copy Blob operation where this blob was the destination blob.</td></tr>
	<tr><td>copy_status</td><td>Specifies the state of the copy operation identified by Copy ID.</td></tr>
	<tr><td>copy_status_description</td><td>Describes cause of fatal or non-fatal copy operation failure.</td></tr>
	<tr><td>destination_snapshot</td><td>Included if the blob is incremental copy blob or incremental copy snapshot, if x-ms-copy-status is success. Snapshot time of the last successful incremental copy snapshot for this blob.</td></tr>
	<tr><td>lease_duration</td><td>Specifies whether the lease is of infinite or fixed duration, when a blob is leased.</td></tr>
	<tr><td>lease_state</td><td>Specifies lease state of the blob.</td></tr>
	<tr><td>lease_status</td><td>Specifies the lease status of the blob.</td></tr>
	<tr><td>incremental_copy</td><td>Copies the snapshot of the source page blob to a destination page blob. The snapshot is copied such that only the differential changes between the previously copied snapshot are transferred to the destination.</td></tr>
	<tr><td>is_sealed</td><td>Indicate if the append blob is sealed or not.</td></tr>
	<tr><td>remaining_retention_days</td><td>The number of days that the blob will be retained before being permanently deleted by the service.</td></tr>
	<tr><td>archive_status</td><td>Specifies the archive status of the blob.</td></tr>
	<tr><td>blob_tag_set</td><td>A list of blob tags.</td></tr>
	<tr><td>metadata</td><td>A name-value pair to associate with the container as metadata.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>region</td><td>The Azure region/location in which the resource is located.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>