# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The resource name.</td></tr>
	<tr><td>id</td><td>The resource ID.</td></tr>
	<tr><td>provisioning_state</td><td>Provisioning state to track the async operation status.</td></tr>
	<tr><td>type</td><td>The resource type.</td></tr>
	<tr><td>sql_image_offer</td><td>SQL image offer for the SQL virtual machine.</td></tr>
	<tr><td>sql_image_sku</td><td>SQL Server edition type. Possible values include: 'Developer', 'Express', 'Standard', 'Enterprise', 'Web'.</td></tr>
	<tr><td>sql_management</td><td>SQL Server Management type. Possible values include: 'Full', 'LightWeight', 'NoAgent'.</td></tr>
	<tr><td>sql_server_license_type</td><td>SQL server license type for the SQL virtual machine. Possible values include: 'PAYG', 'AHUB', 'DR'.</td></tr>
	<tr><td>sql_virtual_machine_group_resource_id</td><td>ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.</td></tr>
	<tr><td>virtual_machine_resource_id</td><td>ARM resource id of underlying virtual machine created from SQL marketplace image.</td></tr>
	<tr><td>auto_backup_settings</td><td>Auto backup settings for SQL Server.</td></tr>
	<tr><td>auto_patching_settings</td><td>Auto patching settings for applying critical security updates to SQL virtual machine.</td></tr>
	<tr><td>identity</td><td>Azure Active Directory identity for the SQL virtual machine.</td></tr>
	<tr><td>key_vault_credential_settings</td><td>Key vault credential settings for the SQL virtual machine.</td></tr>
	<tr><td>server_configurations_management_settings</td><td>SQL server configuration management settings for the SQL virtual machine.</td></tr>
	<tr><td>storage_configuration_settings</td><td>Storage configuration settings for the SQL virtual machine.</td></tr>
	<tr><td>wsfc_domain_credentials</td><td>Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>akas</td><td>Array of globally unique identifier strings (also known as) for the resource.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>region</td><td>The Azure region/location in which the resource is located.</td></tr>
	<tr><td>resource_group</td><td>The resource group which holds this resource.</td></tr>
	<tr><td>cloud_environment</td><td>The Azure Cloud Environment.</td></tr>
	<tr><td>subscription_id</td><td>The Azure Subscription ID in which the resource is located.</td></tr>
	<tr><td>kaytu_account_id</td><td>The Kaytu Account ID in which the resource is located.</td></tr>
	<tr><td>kaytu_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
	<tr><td>kaytu_metadata</td><td>Kaytu Metadata of the Azure resource.</td></tr>
</table>