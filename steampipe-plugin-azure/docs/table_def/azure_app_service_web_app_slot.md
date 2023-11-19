# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>Resource Name.</td></tr>
	<tr><td>app_name</td><td>The name of the application.</td></tr>
	<tr><td>id</td><td>Resource ID of the app slot.</td></tr>
	<tr><td>kind</td><td>Contains the kind of the resource.</td></tr>
	<tr><td>state</td><td>Current state of the app.</td></tr>
	<tr><td>type</td><td>Resource type.</td></tr>
	<tr><td>last_modified_time_utc</td><td>Last time the app was modified, in UTC.</td></tr>
	<tr><td>repository_site_name</td><td>Name of the repository site.</td></tr>
	<tr><td>usage_state</td><td>State indicating whether the app has exceeded its quota usage. Read-only. Possible values include: 'UsageStateNormal', 'UsageStateExceeded'.</td></tr>
	<tr><td>enabled</td><td>Indicates wheather the app is enabled.</td></tr>
	<tr><td>availability_state</td><td>Management information availability state for the app. Possible values include: 'Normal', 'Limited', 'DisasterRecoveryMode'.</td></tr>
	<tr><td>server_farm_id</td><td>Resource ID of the associated App Service plan, formatted as: '/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}'.</td></tr>
	<tr><td>reserved</td><td>True if reserved; otherwise, false.</td></tr>
	<tr><td>is_xenon</td><td>Obsolete: Hyper-V sandbox.</td></tr>
	<tr><td>hyper_v</td><td>Hyper-V sandbox.</td></tr>
	<tr><td>scm_site_also_stopped</td><td>True to stop SCM (KUDU) site when the app is stopped; otherwise, false. The default is false.</td></tr>
	<tr><td>target_swap_slot</td><td>Specifies which deployment slot this app will swap into.</td></tr>
	<tr><td>client_affinity_enabled</td><td>True to enable client affinity; false to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is true.</td></tr>
	<tr><td>client_cert_mode</td><td>This composes with ClientCertEnabled setting. ClientCertEnabled: false means ClientCert is ignored. ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted. Possible values include: 'Required', 'Optional'.</td></tr>
	<tr><td>client_cert_exclusion_paths</td><td>Client certificate authentication comma-separated exclusion paths.</td></tr>
	<tr><td>host_names_disabled</td><td>True to disable the public hostnames of the app; otherwise, false. If true, the app is only accessible via API management process.</td></tr>
	<tr><td>custom_domain_verification_id</td><td>Unique identifier that verifies the custom domains assigned to the app. The customer will add this ID to a text record for verification.</td></tr>
	<tr><td>outbound_ip_addresses</td><td>List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings.</td></tr>
	<tr><td>possible_outbound_ip_addresses</td><td>List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants except dataComponent.</td></tr>
	<tr><td>container_size</td><td>Size of the function container.</td></tr>
	<tr><td>suspended_till</td><td>App suspended till in case memory-time quota is exceeded.</td></tr>
	<tr><td>is_default_container</td><td>True if the app is a default container; otherwise, false.</td></tr>
	<tr><td>default_host_name</td><td>Default hostname of the app.</td></tr>
	<tr><td>https_only</td><td>Configures a web site to accept only https requests.</td></tr>
	<tr><td>redundancy_mode</td><td>Site redundancy mode. Possible values include: 'RedundancyModeNone', 'RedundancyModeManual', 'RedundancyModeFailover', 'RedundancyModeActiveActive', 'RedundancyModeGeoRedundant'.</td></tr>
	<tr><td>identity</td><td>Managed service identity.</td></tr>
	<tr><td>host_names</td><td>Hostnames associated with the app.</td></tr>
	<tr><td>enabled_host_names</td><td>Enabled hostnames for the app. Hostnames need to be assigned (see HostNames) AND enabled. Otherwise, the app is not served on those hostnames.</td></tr>
	<tr><td>host_name_ssl_states</td><td>Hostname SSL states are used to manage the SSL bindings for app's hostnames.</td></tr>
	<tr><td>site_config</td><td>Configuration of the app.</td></tr>
	<tr><td>traffic_manager_host_names</td><td>Azure Traffic Manager hostnames associated with the app.</td></tr>
	<tr><td>hosting_environment_profile</td><td>App Service Environment to use for the app.</td></tr>
	<tr><td>slot_swap_status</td><td>Status of the last deployment slot swap operation.</td></tr>
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