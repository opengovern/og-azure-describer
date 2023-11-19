# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the container.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a container uniquely.</td></tr>
	<tr><td>account_name</td><td>The friendly name that identifies the storage account.</td></tr>
	<tr><td>deleted</td><td>Indicates whether the blob container was deleted.</td></tr>
	<tr><td>public_access</td><td>Specifies whether data in the container may be accessed publicly and the level of access.</td></tr>
	<tr><td>type</td><td>Specifies the type of the container.</td></tr>
	<tr><td>default_encryption_scope</td><td>Default the container to use specified encryption scope for all writes.</td></tr>
	<tr><td>deleted_time</td><td>Specifies the time when the container was deleted.</td></tr>
	<tr><td>deny_encryption_scope_override</td><td>Indicates whether block override of encryption scope from the container default, or not.</td></tr>
	<tr><td>has_immutability_policy</td><td>The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.</td></tr>
	<tr><td>has_legal_hold</td><td>The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.</td></tr>
	<tr><td>last_modified_time</td><td>Specifies the date and time the container was last modified.</td></tr>
	<tr><td>lease_status</td><td>Specifies the lease status of the container.</td></tr>
	<tr><td>lease_state</td><td>Specifies the lease state of the container.</td></tr>
	<tr><td>lease_duration</td><td>Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased. Possible values are: &#39;Infinite&#39;, &#39;Fixed&#39;.</td></tr>
	<tr><td>remaining_retention_days</td><td>Remaining retention days for soft deleted blob container.</td></tr>
	<tr><td>version</td><td>The version of the deleted blob container.</td></tr>
	<tr><td>immutability_policy</td><td>The ImmutabilityPolicy property of the container.</td></tr>
	<tr><td>legal_hold</td><td>The LegalHold property of the container.</td></tr>
	<tr><td>metadata</td><td>A name-value pair to associate with the container as metadata.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>