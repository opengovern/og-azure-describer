# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>Name of the scale set.</td></tr>
	<tr><td>id</td><td>The unique id identifying the resource in subscription.</td></tr>
	<tr><td>provisioning_state</td><td>The provisioning state.</td></tr>
	<tr><td>type</td><td>The type of the resource in Azure.</td></tr>
	<tr><td>location</td><td>The location of the resource.</td></tr>
	<tr><td>do_not_run_extensions_on_overprovisioned_vms</td><td>When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept.</td></tr>
	<tr><td>overprovision</td><td>Specifies whether the Virtual Machine Scale Set should be overprovisioned.</td></tr>
	<tr><td>platform_fault_domain_count</td><td>Fault Domain count for each placement group.</td></tr>
	<tr><td>single_placement_group</td><td>When true this limits the scale set to a single placement group, of max size 100 virtual machines.</td></tr>
	<tr><td>sku_name</td><td>The sku name.</td></tr>
	<tr><td>sku_capacity</td><td>Specifies the tier of virtual machines in a scale set.</td></tr>
	<tr><td>sku_tier</td><td>Specifies the tier of virtual machines in a scale set.</td></tr>
	<tr><td>unique_id</td><td>Specifies the ID which uniquely identifies a Virtual Machine Scale Set.</td></tr>
	<tr><td>extensions</td><td>Specifies the details of VM Scale Set Extensions.</td></tr>
	<tr><td>identity</td><td>The identity of the virtual machine scale set, if configured.</td></tr>
	<tr><td>plan</td><td>Specifies information about the marketplace image used to create the virtual machine.</td></tr>
	<tr><td>scale_in_policy</td><td>Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.</td></tr>
	<tr><td>tags_src</td><td>Resource tags.</td></tr>
	<tr><td>upgrade_policy</td><td>The upgrade policy for the scale set.</td></tr>
	<tr><td>virtual_machine_diagnostics_profile</td><td>Specifies the boot diagnostic settings state.</td></tr>
	<tr><td>virtual_machine_extension_profile</td><td>Specifies a collection of settings for extensions installed on virtual machines in the scale set.</td></tr>
	<tr><td>virtual_machine_network_profile</td><td>Specifies properties of the network interfaces of the virtual machines in the scale set.</td></tr>
	<tr><td>virtual_machine_os_profile</td><td>Specifies the operating system settings for the virtual machines in the scale set.</td></tr>
	<tr><td>virtual_machine_storage_profile</td><td>Specifies the storage settings for the virtual machine disks.</td></tr>
	<tr><td>virtual_machine_security_profile</td><td>Specifies the Security related profile settings for the virtual machines in the scale set.</td></tr>
	<tr><td>zones</td><td>The Logical zone list for scale set.</td></tr>
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