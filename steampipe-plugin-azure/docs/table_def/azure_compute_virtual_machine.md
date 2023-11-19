# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>name</td><td>The friendly name that identifies the virtual machine.</td></tr>
	<tr><td>power_state</td><td>Specifies the power state of the vm.</td></tr>
	<tr><td>id</td><td>The unique id identifying the resource in subscription.</td></tr>
	<tr><td>type</td><td>The type of the resource in Azure.</td></tr>
	<tr><td>provisioning_state</td><td>The virtual machine provisioning state.</td></tr>
	<tr><td>vm_id</td><td>Specifies an unique ID for VM, which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.</td></tr>
	<tr><td>size</td><td>Specifies the size of the virtual machine.</td></tr>
	<tr><td>admin_user_name</td><td>Specifies the name of the administrator account.</td></tr>
	<tr><td>allow_extension_operations</td><td>Specifies whether extension operations should be allowed on the virtual machine.</td></tr>
	<tr><td>availability_set_id</td><td>Specifies the ID of the availability set.</td></tr>
	<tr><td>billing_profile_max_price</td><td>Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS.</td></tr>
	<tr><td>boot_diagnostics_enabled</td><td>Specifies whether boot diagnostics should be enabled on the Virtual Machine, or not.</td></tr>
	<tr><td>boot_diagnostics_storage_uri</td><td>Contains the Uri of the storage account to use for placing the console output and screenshot.</td></tr>
	<tr><td>computer_name</td><td>Specifies the host OS name of the virtual machine.</td></tr>
	<tr><td>disable_password_authentication</td><td>Specifies whether password authentication should be disabled.</td></tr>
	<tr><td>enable_automatic_updates</td><td>Indicates whether automatic updates is enabled for the windows virtual machine.</td></tr>
	<tr><td>eviction_policy</td><td>Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set.</td></tr>
	<tr><td>image_exact_version</td><td>Specifies in decimal numbers, the version of platform image or marketplace image used to create the virtual machine.</td></tr>
	<tr><td>image_id</td><td>Specifies the ID of the image to use.</td></tr>
	<tr><td>image_offer</td><td>Specifies the offer of the platform image or marketplace image used to create the virtual machine.</td></tr>
	<tr><td>image_publisher</td><td>Specifies the publisher of the image to use.</td></tr>
	<tr><td>image_sku</td><td>Specifies the sku of the image to use.</td></tr>
	<tr><td>image_version</td><td>Specifies the version of the platform image or marketplace image used to create the virtual machine.</td></tr>
	<tr><td>managed_disk_id</td><td>Specifies the ID of the managed disk used by the virtual machine.</td></tr>
	<tr><td>os_disk_caching</td><td>Specifies the caching requirements of the operating system disk used by the virtual machine.</td></tr>
	<tr><td>os_disk_create_option</td><td>Specifies how the virtual machine should be created.</td></tr>
	<tr><td>os_disk_name</td><td>Specifies the name of the operating system disk used by the virtual machine.</td></tr>
	<tr><td>os_disk_vhd_uri</td><td>Specifies the virtual hard disk's uri.</td></tr>
	<tr><td>os_name</td><td>The Operating System running on the virtual machine.</td></tr>
	<tr><td>os_version</td><td>The version of Operating System running on the virtual machine.</td></tr>
	<tr><td>os_type</td><td>Specifies the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD.</td></tr>
	<tr><td>priority</td><td>Specifies the priority for the virtual machine.</td></tr>
	<tr><td>provision_vm_agent</td><td>Specifies whether virtual machine agent should be provisioned on the virtual machine for linux configuration.</td></tr>
	<tr><td>provision_vm_agent_windows</td><td>Specifies whether virtual machine agent should be provisioned on the virtual machine for windows configuration.</td></tr>
	<tr><td>require_guest_provision_signal</td><td>Specifies whether the guest provision signal is required to infer provision success of the virtual machine.</td></tr>
	<tr><td>ultra_ssd_enabled</td><td>Specifies whether managed disks with storage account type UltraSSD_LRS can be added to a virtual machine or virtual machine scale set, or not.</td></tr>
	<tr><td>time_zone</td><td>Specifies the time zone of the virtual machine.</td></tr>
	<tr><td>additional_unattend_content</td><td>Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by windows setup.</td></tr>
	<tr><td>data_disks</td><td>A list of parameters that are used to add a data disk to a virtual machine.</td></tr>
	<tr><td>linux_configuration_ssh_public_keys</td><td>A list of ssh key configuration for a Linux OS</td></tr>
	<tr><td>network_interfaces</td><td>A list of resource Ids for the network interfaces associated with the virtual machine.</td></tr>
	<tr><td>patch_settings</td><td>Specifies settings related to in-guest patching (KBs).</td></tr>
	<tr><td>private_ips</td><td>An array of private ip addesses associated with the vm.</td></tr>
	<tr><td>public_ips</td><td>An array of public ip addesses associated with the vm.</td></tr>
	<tr><td>secrets</td><td>A list of certificates that should be installed onto the virtual machine.</td></tr>
	<tr><td>statuses</td><td>Specifies the resource status information.</td></tr>
	<tr><td>extensions</td><td>Specifies the details of VM Extensions.</td></tr>
	<tr><td>guest_configuration_assignments</td><td>Guest configuration assignments for a virtual machine.</td></tr>
	<tr><td>identity</td><td>The identity of the virtual machine, if configured.</td></tr>
	<tr><td>security_profile</td><td>Specifies the security related profile settings for the virtual machine.</td></tr>
	<tr><td>win_rm</td><td>Specifies the windows remote management listeners. This enables remote windows powershell.</td></tr>
	<tr><td>zones</td><td>A list of virtual machine zones.</td></tr>
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