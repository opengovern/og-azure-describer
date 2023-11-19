# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the database.</td></tr>
	<tr><td>id</td><td>Contains ID to identify a database uniquely.</td></tr>
	<tr><td>server_name</td><td>The name of the parent server of the database.</td></tr>
	<tr><td>status</td><td>The status of the database.</td></tr>
	<tr><td>type</td><td>Type of the database.</td></tr>
	<tr><td>collation</td><td>The collation of the database.</td></tr>
	<tr><td>containment_state</td><td>The containment state of the database.</td></tr>
	<tr><td>creation_date</td><td>The creation date of the database.</td></tr>
	<tr><td>current_service_objective_id</td><td>The current service level objective ID of the database.</td></tr>
	<tr><td>database_id</td><td>The ID of the database.</td></tr>
	<tr><td>default_secondary_location</td><td>The default secondary region for this database.</td></tr>
	<tr><td>earliest_restore_date</td><td>This records the earliest start date and time that restore is available for this database.</td></tr>
	<tr><td>edition</td><td>The edition of the database.</td></tr>
	<tr><td>elastic_pool_name</td><td>The name of the elastic pool the database is in.</td></tr>
	<tr><td>failover_group_id</td><td>The resource identifier of the failover group containing this database.</td></tr>
	<tr><td>kind</td><td>Kind of the database.</td></tr>
	<tr><td>location</td><td>Location of the database.</td></tr>
	<tr><td>max_size_bytes</td><td>The max size of the database expressed in bytes.</td></tr>
	<tr><td>recovery_services_recovery_point_resource_id</td><td>Specifies the resource ID of the recovery point to restore from if createMode is RestoreLongTermRetentionBackup.</td></tr>
	<tr><td>requested_service_objective_id</td><td>The configured service level objective ID of the database.</td></tr>
	<tr><td>restore_point_in_time</td><td>Specifies the point in time of the source database that will be restored to create the new database.</td></tr>
	<tr><td>requested_service_objective_name</td><td>The name of the configured service level objective of the database.</td></tr>
	<tr><td>retention_policy_id</td><td>Retention policy ID.</td></tr>
	<tr><td>retention_policy_name</td><td>Retention policy Name.</td></tr>
	<tr><td>retention_policy_type</td><td>Long term Retention policy Type.</td></tr>
	<tr><td>source_database_deletion_date</td><td>Specifies the time that the database was deleted when createMode is Restore and sourceDatabaseId is the deleted database's original resource id.</td></tr>
	<tr><td>source_database_id</td><td>Specifies the resource ID of the source database if createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore.</td></tr>
	<tr><td>zone_redundant</td><td>Indicates if the database is zone redundant or not.</td></tr>
	<tr><td>create_mode</td><td>Specifies the mode of database creation.</td></tr>
	<tr><td>read_scale</td><td>ReadScale indicates whether read-only connections are allowed to this database or not if the database is a geo-secondary.</td></tr>
	<tr><td>recommended_index</td><td>The recommended indices for this database.</td></tr>
	<tr><td>retention_policy_property</td><td>Long term Retention policy Property.</td></tr>
	<tr><td>sample_name</td><td>Indicates the name of the sample schema to apply when creating this database.</td></tr>
	<tr><td>service_level_objective</td><td>The current service level objective of the database.</td></tr>
	<tr><td>service_tier_advisors</td><td>The list of service tier advisors for this database.</td></tr>
	<tr><td>transparent_data_encryption</td><td>The transparent data encryption info for this database.</td></tr>
	<tr><td>vulnerability_assessments</td><td>The vulnerability assessments for this database.</td></tr>
	<tr><td>vulnerability_assessment_scan_records</td><td>The vulnerability assessment scan records for this database.</td></tr>
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