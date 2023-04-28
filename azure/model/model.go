//go:generate go run ../../keibi-es-sdk/gen/main.go --file $GOFILE --output ../../keibi-es-sdk/azure_resources_clients.go --type azure

package model

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/healthcareapis/mgmt/healthcareapis"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/hybridcompute/mgmt/hybridcompute"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/hybridkubernetes/mgmt/hybridkubernetes"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/provisioningservices/mgmt/iothub"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/recoveryservices/mgmt/recoveryservices"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/links"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/locks"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managementgroups"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/policy"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	sub "github.com/Azure/azure-sdk-for-go/profiles/latest/subscription/mgmt/subscription"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2020-12-01/apimanagement"
	"github.com/Azure/azure-sdk-for-go/services/appconfiguration/mgmt/2020-06-01/appconfiguration"
	"github.com/Azure/azure-sdk-for-go/services/appplatform/mgmt/2020-07-01/appplatform"
	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2020-09-01/batch"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/mgmt/2021-04-30/cognitiveservices"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-09-01/skus"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-02-01/containerservice"
	"github.com/Azure/azure-sdk-for-go/services/databoxedge/mgmt/2019-07-01/databoxedge"
	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	analytics "github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account"
	store "github.com/Azure/azure-sdk-for-go/services/datalake/store/mgmt/2016-11-01/account"
	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-05-01/frontdoor"
	"github.com/Azure/azure-sdk-for-go/services/guestconfiguration/mgmt/2020-06-25/guestconfiguration"
	"github.com/Azure/azure-sdk-for-go/services/hdinsight/mgmt/2018-06-01/hdinsight"
	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2020-03-01/devices"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	secret "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2021-01-01/kusto"
	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2021-05-01/mysqlflexibleservers"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-05-01/network"
	newnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2020-11-01-preview/containerregistry"
	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	"github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2021-06-01-preview/eventgrid"
	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	previewKeyvault "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/preview/machinelearningservices/mgmt/2020-02-18-preview/machinelearningservices"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-04-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v1.0/security"
	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql"
	sqlv3 "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	sqlv5 "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v5.0/sql"
	"github.com/Azure/azure-sdk-for-go/services/preview/sqlvirtualmachine/mgmt/2017-03-01-preview/sqlvirtualmachine"
	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-06-01/redis"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-06-01/subscriptions"
	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	"github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2019-03-01/servicefabric"
	"github.com/Azure/azure-sdk-for-go/services/signalr/mgmt/2020-05-01/signalr"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	"github.com/Azure/azure-sdk-for-go/services/storagecache/mgmt/2021-05-01/storagecache"
	"github.com/Azure/azure-sdk-for-go/services/storagesync/mgmt/2020-03-01/storagesync"
	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2016-03-01/streamanalytics"
	"github.com/Azure/azure-sdk-for-go/services/synapse/mgmt/2021-03-01/synapse"
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-06-01/web"
	azblobOld "github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/queue/queues"
	"github.com/tombuildsstuff/giovanni/storage/2019-12-12/blob/accounts"
)

type Metadata struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	SubscriptionID   string `json:"subscription_id"`
	Location         string `json:"location"`
	CloudEnvironment string `json:"cloud_environment"`
	ResourceType     string `json:"resource_type"`
	SourceID         string `json:"source_id"`
}

//  ===================  APIManagement ==================

//index:microsoft_apimanagement_service
//getfilter:name=description.APIManagement.name
//getfilter:resource_group=description.ResourceGroup
type APIManagementDescription struct {
	APIManagement               apimanagement.ServiceResource
	DiagnosticSettingsResources []insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  ===================  App Configuration ==================

//index:microsoft_appconfiguration_configurationstores
//getfilter:name=description.ConfigurationStore.name
//getfilter:resource_group=description.ResourceGroup
type AppConfigurationDescription struct {
	ConfigurationStore          appconfiguration.ConfigurationStore
	DiagnosticSettingsResources []insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== web ==================

//index:microsoft_web_hostingenvironments
//getfilter:name=description.AppServiceEnvironmentResource.name
//getfilter:resource_group=description.ResourceGroup
type AppServiceEnvironmentDescription struct {
	AppServiceEnvironmentResource web.AppServiceEnvironmentResource
	ResourceGroup                 string
}

//index:microsoft_web_sites
//getfilter:name=description.Site.name
//getfilter:resource_group=description.ResourceGroup
type AppServiceFunctionAppDescription struct {
	Site               web.Site
	SiteAuthSettings   web.SiteAuthSettings
	SiteConfigResource web.SiteConfigResource
	ResourceGroup      string
}

//index:microsoft_web_staticsites
//getfilter:name=description.Site.name
//getfilter:resource_group=description.ResourceGroup
type AppServiceWebAppDescription struct {
	Site               web.Site
	SiteAuthSettings   web.SiteAuthSettings
	SiteConfigResource web.SiteConfigResource
	VnetInfo           web.VnetInfo
	ResourceGroup      string
}

//index:microsoft_web_plan
//getfilter:name=description.Site.name
//getfilter:resource_group=description.ResourceGroup
type AppServicePlanDescription struct {
	Plan          web.AppServicePlan
	Apps          []web.Site
	ResourceGroup string
}

//  =================== compute ==================

//index:microsoft_compute_disks
//getfilter:name=description.Disk.name
//getfilter:resource_group=description.ResourceGroup
type ComputeDiskDescription struct {
	Disk          compute.Disk
	ResourceGroup string
}

//index:microsoft_compute_disksreadops
type ComputeDiskReadOpsDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_disksreadopsdaily
type ComputeDiskReadOpsDailyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_disksreadopshourly
type ComputeDiskReadOpsHourlyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskswriteops
type ComputeDiskWriteOpsDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskswriteopsdaily
type ComputeDiskWriteOpsDailyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskswriteopshourly
type ComputeDiskWriteOpsHourlyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_diskaccesses
//getfilter:name=description.DiskAccess.name
//getfilter:resource_group=description.ResourceGroup
type ComputeDiskAccessDescription struct {
	DiskAccess    compute.DiskAccess
	ResourceGroup string
}

//index:microsoft_compute_virtualmachinescalesets
//getfilter:name=description.VirtualMachineScaleSet.name
//getfilter:resource_group=description.ResourceGroup
type ComputeVirtualMachineScaleSetDescription struct {
	VirtualMachineScaleSet           compute.VirtualMachineScaleSet
	VirtualMachineScaleSetExtensions []compute.VirtualMachineScaleSetExtension
	ResourceGroup                    string
}

//index:microsoft_compute_virtualmachinescalesetnetworkinterface
type ComputeVirtualMachineScaleSetNetworkInterfaceDescription struct {
	VirtualMachineScaleSet compute.VirtualMachineScaleSet
	NetworkInterface       network.Interface
	ResourceGroup          string
}

//index:microsoft_compute_virtualmachinescalesetvm
//getfilter:scale_set_name=description.VirtualMachineScaleSet.name
//getfilter:instance_id=description.ScaleSetVM.InstanceID
//getfilter:resource_group=description.ResourceGroup
type ComputeVirtualMachineScaleSetVmDescription struct {
	VirtualMachineScaleSet compute.VirtualMachineScaleSet
	ScaleSetVM             compute.VirtualMachineScaleSetVM
	ResourceGroup          string
}

//index:microsoft_compute_snapshots
//getfilter:name=description.Snapshot.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeSnapshotsDescription struct {
	Snapshot      compute.Snapshot
	ResourceGroup string
}

//index:microsoft_compute_availabilityset
//getfilter:name=description.AvailabilitySet.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeAvailabilitySetDescription struct {
	AvailabilitySet compute.AvailabilitySet
	ResourceGroup   string
}

//index:microsoft_compute_diskencryptionset
//getfilter:name=description.DiskEncryptionSet.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeDiskEncryptionSetDescription struct {
	DiskEncryptionSet compute.DiskEncryptionSet
	ResourceGroup     string
}

//index:microsoft_compute_gallery
//getfilter:name=description.Gallery.Name
//getfilter:resource_group=description.ResourceGroup
type ComputeGalleryDescription struct {
	Gallery       compute.Gallery
	ResourceGroup string
}

//index:microsoft_compute_image
//getfilter:name=Description.Image.Name
//getfilter:resource_group=Description.Image.ResourceGroup
type ComputeImageDescription struct {
	Image         compute.Image
	ResourceGroup string
}

//  =================== databoxedge ==================

//index:microsoft_databoxedge_databoxedgedevices
//getfilter:name=description.Device.name
//getfilter:resource_group=description.ResourceGroup
type DataboxEdgeDeviceDescription struct {
	Device        databoxedge.Device
	ResourceGroup string
}

//  =================== healthcareapis ==================

//index:microsoft_healthcareapis_services
//getfilter:name=description.ServicesDescription.name
//getfilter:resource_group=description.ResourceGroup
type HealthcareServiceDescription struct {
	ServicesDescription         healthcareapis.ServicesDescription
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	PrivateEndpointConnections  *[]healthcareapis.PrivateEndpointConnection
	ResourceGroup               string
}

//  =================== storagecache ==================

//index:microsoft_storagecache_caches
//getfilter:name=description.Cache.name
//getfilter:resource_group=description.ResourceGroup
type HpcCacheDescription struct {
	Cache         storagecache.Cache
	ResourceGroup string
}

//  =================== keyvault ==================

//index:microsoft_keyvault_vaults_keys
//getfilter:vault_name=description.Vault.name
//getfilter:name=description.Key.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultKeyDescription struct {
	Vault         keyvault.Resource
	Key           keyvault.Key
	ResourceGroup string
}

//  =================== containerservice ==================

//index:microsoft_containerservice_managedclusters
//getfilter:name=description.ManagedCluster.name
//getfilter:resource_group=description.ResourceGroup
type KubernetesClusterDescription struct {
	ManagedCluster containerservice.ManagedCluster
	ResourceGroup  string
}

//  =================== network ==================

//index:microsoft_network_networkinterfaces
//getfilter:name=description.Interface.name
//getfilter:resource_group=description.ResourceGroup
type NetworkInterfaceDescription struct {
	Interface     network.Interface
	ResourceGroup string
}

//index:microsoft_network_networkwatchers
//getfilter:network_watcher_name=description.NetworkWatcherName
//getfilter:name=description.ManagedCluster.name
//getfilter:resource_group=description.ResourceGroup
type NetworkWatcherFlowLogDescription struct {
	NetworkWatcherName string
	FlowLog            network.FlowLog
	ResourceGroup      string
}

//index:microsoft_network_routetables
//getfilter:name=description.RouteTable.Name
//getfilter:resource_group=description.ResourceGroup
type RouteTablesDescription struct {
	RouteTable    newnetwork.RouteTable
	ResourceGroup string
}

//index:microsoft_network_applicationsecuritygroups
//getfilter:name=description.ApplicationSecurityGroup.Name
//getfilter:resource_group=description.ResourceGroup
type NetworkApplicationSecurityGroupsDescription struct {
	ApplicationSecurityGroup newnetwork.ApplicationSecurityGroup
	ResourceGroup            string
}

//index:microsoft_network_azurefirewall
//getfilter:name=description.AzureFirewall.Name
//getfilter:resource_group=description.ResourceGroup
type NetworkAzureFirewallDescription struct {
	AzureFirewall newnetwork.AzureFirewall
	ResourceGroup string
}

//index:microsoft_network_expressroutecircuit
//getfilter:name=description.ExpressRouteCircuit.name
//getfilter:resource_group=description.ResourceGroup
type ExpressRouteCircuitDescription struct {
	ExpressRouteCircuit newnetwork.ExpressRouteCircuit
	ResourceGroup       string
}

//index:microsoft_network_virtualnetworkgateway
//getfilter:name=description.VirtualNetworkGateway.Name
//getfilter:resource_group=description.ResourceGroup
type VirtualNetworkGatewayDescription struct {
	VirtualNetworkGateway           newnetwork.VirtualNetworkGateway
	VirtualNetworkGatewayConnection newnetwork.VirtualNetworkGatewayConnection
	ResourceGroup                   string
}

//index:microsoft_network_dnszone
//getfilter:name=description.Zone.Name
//getfilter:resource_group=description.ResourceGroup
type DNSZoneDescription struct { // TODO: Implement describer func
	Zone          armdns.Zone
	ResourceGroup string
}

//index:microsoft_network_firewallpolicy
//getfilter:name=description.FirewallPolicy.Name
//getfilter:resource_group=description.ResourceGroup
type FirewallPolicyDescription struct {
	FirewallPolicy newnetwork.FirewallPolicy
	ResourceGroup  string
}

//index:microsoft_network_frontdoorwebapplicationfirewallpolicy
//getfilter:name=description.WebApplicationFirewallPolicy.Name
//getfilter:resource_group=description.ResourceGroup
type FrontdoorWebApplicationFirewallPolicyDescription struct { // TODO: Implement describer func
	WebApplicationFirewallPolicy frontdoor.WebApplicationFirewallPolicy
	ResourceGroup                string
}

//index:microsoft_network_localnetworkgateway
//getfilter:name=description.LocalNetworkGateway.Name
//getfilter:resource_group=description.ResourceGroup
type LocalNetworkGatewayDescription struct {
	LocalNetworkGateway newnetwork.LocalNetworkGateway
	ResourceGroup       string
}

//index:microsoft_network_natgateways
//getfilter:name=description.NatGateway.Name
//getfilter:resource_group=description.ResourceGroup
type NatGatewayDescription struct {
	NatGateway    newnetwork.NatGateway
	ResourceGroup string
}

//index:microsoft_network_privatelinkservice
//getfilter:name=description.PrivateLinkService.Name
//getfilter:resource_group=description.ResourceGroup
type PrivateLinkServiceDescription struct {
	PrivateLinkService newnetwork.PrivateLinkService
	ResourceGroup      string
}

//index:microsoft_network_routefilter
//getfilter:name=description.RouteFilter.Name
//getfilter:resource_group=description.ResourceGroup
type RouteFilterDescription struct {
	RouteFilter   newnetwork.RouteFilter
	ResourceGroup string
}

//index:microsoft_network_vpngateway
//getfilter:name=description.VpnGateway.Name
//getfilter:resource_group=description.ResourceGroup
type VpnGatewayDescription struct {
	VpnGateway    newnetwork.VpnGateway
	ResourceGroup string
}

//index:microsoft_network_publicipaddresses
//getfilter:name=description.PublicIPAddress.Name
//getfilter:resource_group=description.ResourceGroup
type PublicIPAddressDescription struct {
	PublicIPAddress newnetwork.PublicIPAddress
	ResourceGroup   string
}

//  =================== policy ==================

//index:microsoft_authorization_policyassignments
//getfilter:name=description.Assignment.name
type PolicyAssignmentDescription struct {
	Assignment policy.Assignment
}

//  =================== redis ==================

//index:microsoft_cache_redis
//getfilter:name=description.ResourceType.name
//getfilter:resource_group=description.ResourceGroup
type RedisCacheDescription struct {
	ResourceType  redis.ResourceType
	ResourceGroup string
}

//  =================== links ==================

//index:microsoft_resources_links
//getfilter:id=description.ResourceLink.id
type ResourceLinkDescription struct {
	ResourceLink links.ResourceLink
}

//  =================== authorization ==================

//index:microsoft_authorization_elevateaccessroleassignment
//getfilter:id=description.RoleAssignment.id
type RoleAssignmentDescription struct {
	RoleAssignment authorization.RoleAssignment
}

//index:microsoft_authorization_roledefinitions
//getfilter:name=description.RoleDefinition.name
type RoleDefinitionDescription struct {
	RoleDefinition authorization.RoleDefinition
}

//index:microsoft_authorization_policydefinition
//getfilter:name=description.Definition.Name
type PolicyDefinitionDescription struct {
	Definition policy.Definition
	TurboData  map[string]interface{}
}

//  =================== security ==================

//index:microsoft_security_autoprovisioningsettings
//getfilter:name=description.AutoProvisioningSetting.name
type SecurityCenterAutoProvisioningDescription struct {
	AutoProvisioningSetting security.AutoProvisioningSetting
}

//index:microsoft_security_securitycontacts
//getfilter:name=description.Contact.name
type SecurityCenterContactDescription struct {
	Contact security.Contact
}

//index:microsoft_security_locations_jitnetworkaccesspolicies
type SecurityCenterJitNetworkAccessPolicyDescription struct {
	JitNetworkAccessPolicy security.JitNetworkAccessPolicy
}

//index:microsoft_security_settings
//getfilter:name=description.Setting.name
type SecurityCenterSettingDescription struct {
	Setting security.Setting
}

//index:microsoft_security_pricings
//getfilter:name=description.Pricing.name
type SecurityCenterSubscriptionPricingDescription struct {
	Pricing security.Pricing
}

//index:microsoft_security_automations
//getfilter:name=description.Automation.name
//getfilter:resource_group=description.ResourceGroup
type SecurityCenterAutomationDescription struct {
	Automation    security.Automation
	ResourceGroup string
}

//index:microsoft_security_subassessments
type SecurityCenterSubAssessmentDescription struct {
	SubAssessment security.SubAssessment
	ResourceGroup string
}

//  =================== storage ==================

//index:microsoft_storage_storageaccounts_containers
//getfilter:name=description.ListContainerItem.name
//getfilter:resource_group=description.ResourceGroup
//getfilter:account_name=description.AccountName
type StorageContainerDescription struct {
	AccountName        string
	ListContainerItem  storage.ListContainerItem
	ImmutabilityPolicy storage.ImmutabilityPolicy
	ResourceGroup      string
}

//index:microsoft_storage_blobs
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageBlobDescription struct {
	Blob          azblobOld.BlobItemInternal
	AccountName   string
	IsSnapshot    bool
	ContainerName string
	ResourceGroup string
}

//index:microsoft_storage_blobservices
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageBlobServiceDescription struct {
	BlobService   storage.BlobServiceProperties
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_queues
//listfilter:name=description.Queue.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageQueueDescription struct {
	Queue         storage.ListQueue
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_fileshares
//listfilter:name=description.FileShare.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageFileShareDescription struct {
	FileShare     storage.FileShareItem
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_tables
//listfilter:name=description.Table.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageTableDescription struct {
	Table         storage.Table
	AccountName   string
	Location      string
	ResourceGroup string
}

//index:microsoft_storage_tableservices
//listfilter:name=description.TableService.Name
//listfilter:storage_account_name=description.AccountName
//listfilter:resource_group=description.ResourceGroup
type StorageTableServiceDescription struct {
	TableService  storage.TableServiceProperties
	AccountName   string
	Location      string
	ResourceGroup string
}

//  =================== network ==================

//index:microsoft_network_virtualnetworks_subnets
//getfilter:name=description.Subnet.name
//getfilter:resource_group=description.ResourceGroup
//getfilter:virtual_network_name=description.VirtualNetworkName
type SubnetDescription struct {
	VirtualNetworkName string
	Subnet             network.Subnet
	ResourceGroup      string
}

//index:microsoft_network_virtualnetworks
//getfilter:name=description.VirtualNetwork.name
//getfilter:resource_group=description.ResourceGroup
type VirtualNetworkDescription struct {
	VirtualNetwork network.VirtualNetwork
	ResourceGroup  string
}

//  =================== subscriptions ==================

//index:microsoft_resources_tenants
type TenantDescription struct {
	TenantIDDescription subscriptions.TenantIDDescription
}

//index:microsoft_resources_subscriptions
type SubscriptionDescription struct {
	Subscription subscriptions.Subscription
}

//  =================== network ==================

//index:microsoft_network_applicationgateways
//getfilter:name=description.ApplicationGateway.name
//getfilter:resource_group=description.ResourceGroup
type ApplicationGatewayDescription struct {
	ApplicationGateway          newnetwork.ApplicationGateway
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== batch ==================

//index:microsoft_batch_batchaccounts
//getfilter:name=description.Account.name
//getfilter:resource_group=description.ResourceGroup
type BatchAccountDescription struct {
	Account                     batch.Account
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== cognitiveservices ==================

//index:microsoft_cognitiveservices_accounts
//getfilter:name=description.Account.name
//getfilter:resource_group=description.ResourceGroup
type CognitiveAccountDescription struct {
	Account                     cognitiveservices.Account
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== compute ==================

//index:microsoft_compute_virtualmachines
//getfilter:name=description.VirtualMachine.name
//getfilter:resource_group=description.ResourceGroup
type ComputeVirtualMachineDescription struct {
	VirtualMachine             compute.VirtualMachine
	VirtualMachineInstanceView compute.VirtualMachineInstanceView
	InterfaceIPConfigurations  []network.InterfaceIPConfiguration
	PublicIPs                  []string
	VirtualMachineExtension    *[]compute.VirtualMachineExtension
	Assignments                *[]guestconfiguration.Assignment
	ResourceGroup              string
}

//index:microsoft_compute_resourcesku
type ComputeResourceSKUDescription struct {
	ResourceSKU skus.ResourceSku
}

//index:microsoft_compute_virtualmachinecpuutilization
type ComputeVirtualMachineCpuUtilizationDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_virtualmachinecpuutilizationdaily
type ComputeVirtualMachineCpuUtilizationDailyDescription struct {
	MonitoringMetric
}

//index:microsoft_compute_virtualmachinecpuutilizationhourly
type ComputeVirtualMachineCpuUtilizationHourlyDescription struct {
	MonitoringMetric
}

//  =================== containerregistry ==================

//index:microsoft_containerregistry_registries
//getfilter:name=description.Registry.name
//getfilter:resource_group=description.ResourceGroup
type ContainerRegistryDescription struct {
	Registry                      containerregistry.Registry
	RegistryListCredentialsResult containerregistry.RegistryListCredentialsResult
	RegistryUsages                *[]containerregistry.RegistryUsage
	ResourceGroup                 string
}

//  =================== documentdb ==================

//index:microsoft_documentdb_databaseaccounts
//getfilter:name=description.DatabaseAccountGetResults.name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbAccountDescription struct {
	DatabaseAccountGetResults documentdb.DatabaseAccountGetResults
	ResourceGroup             string
}

//index:microsoft_documentdb_mongodatabases
//getfilter:account_name=description.Account.name
//getfilter:name=description.MongoDatabase.name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbMongoDatabaseDescription struct {
	Account       documentdb.DatabaseAccountGetResults
	MongoDatabase documentdb.MongoDBDatabaseGetResults
	ResourceGroup string
}

//index:microsoft_documentdb_sqldatabases
//getfilter:account_name=description.Account.name
//getfilter:name=description.SqlDatabase.name
//getfilter:resource_group=description.ResourceGroup
type CosmosdbSqlDatabaseDescription struct {
	Account       documentdb.DatabaseAccountGetResults
	SqlDatabase   documentdb.SQLDatabaseGetResults
	ResourceGroup string
}

//  =================== datafactory ==================

//index:microsoft_datafactory_datafactories
//getfilter:name=description.Factory.name
//getfilter:resource_group=description.ResourceGroup
type DataFactoryDescription struct {
	Factory                    datafactory.Factory
	PrivateEndPointConnections []datafactory.PrivateEndpointConnectionResource
	ResourceGroup              string
}

//index:microsoft_datafactory_datafactorydatasets
//getfilter:factory_name=description.Factory.name
//getfilter:name=description.Dataset.name
//getfilter:resource_group=description.ResourceGroup
type DataFactoryDatasetDescription struct {
	Factory       datafactory.Factory
	Dataset       datafactory.DatasetResource
	ResourceGroup string
}

//index:microsoft_datafactory_datafactorypipelines
//getfilter:factory_name=description.Factory.name
//getfilter:name=description.Pipeline.name
//getfilter:resource_group=description.ResourceGroup
type DataFactoryPipelineDescription struct {
	Factory       datafactory.Factory
	Pipeline      datafactory.PipelineResource
	ResourceGroup string
}

//  =================== account ==================

//index:microsoft_datalakeanalytics_accounts
//getfilter:name=description.DataLakeAnalyticsAccount.name
//getfilter:resource_group=description.ResourceGroup
type DataLakeAnalyticsAccountDescription struct {
	DataLakeAnalyticsAccount   analytics.DataLakeAnalyticsAccount
	DiagnosticSettingsResource *[]insights.DiagnosticSettingsResource
	ResourceGroup              string
}

//  =================== account ==================

//index:microsoft_datalakestore_accounts
//getfilter:name=description.DataLakeStoreAccount.name
//getfilter:resource_group=description.ResourceGroup
type DataLakeStoreDescription struct {
	DataLakeStoreAccount       store.DataLakeStoreAccount
	DiagnosticSettingsResource *[]insights.DiagnosticSettingsResource
	ResourceGroup              string
}

//  =================== insights ==================

type MonitoringMetric struct {
	// Resource Name
	DimensionValue string
	// MetadataValue represents a metric metadata value.
	MetaData *insights.MetadataValue
	// Metric the result data of a query.
	Metric *insights.Metric
	// The maximum metric value for the data point.
	Maximum *float64
	// The minimum metric value for the data point.
	Minimum *float64
	// The average of the metric values that correspond to the data point.
	Average *float64
	// The number of metric values that contributed to the aggregate value of this data point.
	SampleCount *float64
	// The sum of the metric values for the data point.
	Sum *float64
	// The time stamp used for the data point.
	TimeStamp string
	// The units in which the metric value is reported.
	Unit string
}

//index:microsoft_insights_guestdiagnosticsettings
//getfilter:name=description.DiagnosticSettingsResource.name
//getfilter:resource_group=description.ResourceGroup
type DiagnosticSettingDescription struct {
	DiagnosticSettingsResource insights.DiagnosticSettingsResource
	ResourceGroup              string
}

//  =================== eventgrid ==================

//index:microsoft_eventgrid_domains
//getfilter:name=description.Domain.name
//getfilter:resource_group=description.ResourceGroup
type EventGridDomainDescription struct {
	Domain                      eventgrid.Domain
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== eventgrid ==================

//index:microsoft_eventgrid_topics
//getfilter:name=description.Topic.name
//getfilter:resource_group=description.ResourceGroup
type EventGridTopicDescription struct {
	Topic                       eventgrid.Topic
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== eventhub ==================

//index:microsoft_eventhub_namespaces
//getfilter:name=description.EHNamespace.name
//getfilter:resource_group=description.ResourceGroup
type EventhubNamespaceDescription struct {
	EHNamespace                 eventhub.EHNamespace
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	NetworkRuleSet              eventhub.NetworkRuleSet
	PrivateEndpointConnection   []eventhub.PrivateEndpointConnection
	ResourceGroup               string
}

//  =================== frontdoor ==================

//index:microsoft_network_frontdoors
//getfilter:name=description.FrontDoor.name
//getfilter:resource_group=description.ResourceGroup
type FrontdoorDescription struct {
	FrontDoor                   frontdoor.FrontDoor
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== hdinsight ==================

//index:microsoft_hdinsight_clusterpools
//getfilter:name=description.Cluster.name
//getfilter:resource_group=description.ResourceGroup
type HdinsightClusterDescription struct {
	Cluster                     hdinsight.Cluster
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== hybridcompute ==================

//index:microsoft_hybridcompute_machines
//getfilter:name=description.Machine.name
//getfilter:resource_group=description.ResourceGroup
type HybridComputeMachineDescription struct {
	Machine           hybridcompute.Machine
	MachineExtensions []hybridcompute.MachineExtension
	ResourceGroup     string
}

//  =================== devices ==================

//index:microsoft_devices_iothubs
//getfilter:name=description.IotHubDescription.name
//getfilter:resource_group=description.ResourceGroup
type IOTHubDescription struct {
	IotHubDescription           devices.IotHubDescription
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_devices_iothubdpses
//getfilter:name=description.IotHubDps.name
//getfilter:resource_group=description.ResourceGroup
type IOTHubDpsDescription struct {
	IotHubDps                   iothub.ProvisioningServiceDescription
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== keyvault ==================

//index:microsoft_keyvault_vaults
//getfilter:name=description.Resource.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultDescription struct {
	Resource                    keyvault.Resource
	Vault                       keyvault.Vault
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_keyvault_deletedvaults
//getfilter:name=description.Vault.name
//getfilter:region=description.Vault.Properties.Location
type KeyVaultDeletedVaultDescription struct {
	Vault         keyvault.DeletedVault
	ResourceGroup string
}

//  =================== keyvault ==================

//index:microsoft_keyvault_managedhsms
//getfilter:name=description.ManagedHsm.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultManagedHardwareSecurityModuleDescription struct {
	ManagedHsm                  previewKeyvault.ManagedHsm
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== secret ==================

//index:microsoft_keyvault_vaults_secrets
//getfilter:name=description.SecretItem.name
//getfilter:resource_group=description.ResourceGroup
type KeyVaultSecretDescription struct {
	SecretItem    secret.SecretItem
	SecretBundle  secret.SecretBundle
	TurboData     map[string]interface{}
	ResourceGroup string
}

//  =================== kusto ==================

//index:microsoft_kusto_clusters
//getfilter:name=description.Cluster.name
//getfilter:resource_group=description.ResourceGroup
type KustoClusterDescription struct {
	Cluster       kusto.Cluster
	ResourceGroup string
}

//  =================== insights ==================

//index:microsoft_insights_activitylogalerts
//getfilter:name=description.ActivityLogAlertResource.name
//getfilter:resource_group=description.ResourceGroup
type LogAlertDescription struct {
	ActivityLogAlertResource insights.ActivityLogAlertResource
	ResourceGroup            string
}

//  =================== insights ==================

//index:microsoft_insights_logprofiles
//getfilter:name=description.LogProfileResource.name
//getfilter:resource_group=description.ResourceGroup
type LogProfileDescription struct {
	LogProfileResource insights.LogProfileResource
	ResourceGroup      string
}

//  =================== logic ==================

//index:microsoft_logic_workflows
//getfilter:name=description.Workflow.name
//getfilter:resource_group=description.ResourceGroup
type LogicAppWorkflowDescription struct {
	Workflow                    logic.Workflow
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== machinelearningservices ==================

//index:microsoft_machinelearning_workspaces
//getfilter:name=description.Workspace.name
//getfilter:resource_group=description.ResourceGroup
type MachineLearningWorkspaceDescription struct {
	Workspace                   machinelearningservices.Workspace
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== mariadb ==================

//index:microsoft_dbformariadb_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type MariadbServerDescription struct {
	Server        mariadb.Server
	ResourceGroup string
}

//  =================== mysql ==================

//index:microsoft_dbformysql_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type MysqlServerDescription struct {
	Server         mysql.Server
	Configurations *[]mysql.Configuration
	ServerKeys     []mysql.ServerKey
	ResourceGroup  string
}

//  =================== network ==================

//index:microsoft_classicnetwork_networksecuritygroups
//getfilter:name=description.SecurityGroup.name
//getfilter:resource_group=description.ResourceGroup
type NetworkSecurityGroupDescription struct {
	SecurityGroup               newnetwork.SecurityGroup
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//index:microsoft_network_networkwatchers
//getfilter:name=description.Watcher.name
//getfilter:resource_group=description.ResourceGroup
type NetworkWatcherDescription struct {
	Watcher       newnetwork.Watcher
	ResourceGroup string
}

//  =================== search ==================

//index:microsoft_search_searchservices
//getfilter:name=description.Service.name
//getfilter:resource_group=description.ResourceGroup
type SearchServiceDescription struct {
	Service                     search.Service
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== servicefabric ==================

//index:microsoft_servicefabric_clusters
//getfilter:name=description.Cluster.name
//getfilter:resource_group=description.ResourceGroup
type ServiceFabricClusterDescription struct {
	Cluster       servicefabric.Cluster
	ResourceGroup string
}

//  =================== servicebus ==================

//index:microsoft_servicebus_namespaces
//getfilter:name=description.SBNamespace.name
//getfilter:resource_group=description.ResourceGroup
type ServicebusNamespaceDescription struct {
	SBNamespace                 servicebus.SBNamespace
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	NetworkRuleSet              servicebus.NetworkRuleSet
	PrivateEndpointConnections  []servicebus.PrivateEndpointConnection
	ResourceGroup               string
}

//  =================== signalr ==================

//index:microsoft_signalrservice_signalr
//getfilter:name=description.ResourceType.name
//getfilter:resource_group=description.ResourceGroup
type SignalrServiceDescription struct {
	ResourceType                signalr.ResourceType
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== appplatform ==================

//index:microsoft_appplatform_spring
//getfilter:name=description.ServiceResource.name
//getfilter:resource_group=description.ResourceGroup
type SpringCloudServiceDescription struct {
	ServiceResource            appplatform.ServiceResource
	DiagnosticSettingsResource *[]insights.DiagnosticSettingsResource
	ResourceGroup              string
}

//  =================== streamanalytics ==================

//index:microsoft_streamanalytics_streamingjobs
//getfilter:name=description.StreamingJob.name
//getfilter:resource_group=description.ResourceGroup
type StreamAnalyticsJobDescription struct {
	StreamingJob                streamanalytics.StreamingJob
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	ResourceGroup               string
}

//  =================== synapse ==================

//index:microsoft_synapse_workspaces
//getfilter:name=description.Workspace.name
//getfilter:resource_group=description.ResourceGroup
type SynapseWorkspaceDescription struct {
	Workspace                      synapse.Workspace
	ServerVulnerabilityAssessments []synapse.ServerVulnerabilityAssessment
	DiagnosticSettingsResources    *[]insights.DiagnosticSettingsResource
	ResourceGroup                  string
}

//  =================== sub ==================

//index:microsoft_resources_subscriptions_locations
//getfilter:name=description.Location.name
//getfilter:resource_group=description.ResourceGroup
type LocationDescription struct {
	Location      sub.Location
	ResourceGroup string
}

//  =================== msgraph ==================

//index:microsoft_resources_users
//getfilter:id=description.AdUsers.DirectoryObject.id
//listfilter:id=description.AdUsers.DirectoryObject.id
//listfilter:user_principal_name=description.AdUsers.userPrincipalName
//listfilter:filter=description.AdUsers.filter
//listfilter:user_type=description.AdUsers.userType
//listfilter:account_enabled=description.AdUsers.accountEnabled
//listfilter:display_name=description.AdUsers.displayName
//listfilter:surname=description.AdUsers.surname
type AdUsersDescription struct {
	TenantID string
	AdUsers  msgraph.User
}

//index:microsoft_resources_groups
//getfilter:id=description.AdGroup.DirectoryObject.ID
//listfilter:display_name=description.AdGroup.DisplayName
//listfilter:mail=description.AdGroup.Mail
//listfilter:mail_enabled=description.AdGroup.MailEnabled
//listfilter:on_premises_sync_enabled=description.AdGroup.OnPremisesSyncEnabled
//listfilter:security_enabled=description.AdGroup.SecurityEnabled
type AdGroupDescription struct {
	TenantID string
	AdGroup  msgraph.Group
}

//index:microsoft_resources_serviceprincipals
//getfilter:id=description.AdServicePrincipal.DirectoryObject.ID
//listfilter:display_name=description.AdServicePrincipal.DisplayName
//listfilter:account_enabled=description.AdServicePrincipal.AccountEnabled
//listfilter:service_principal_type=description.AdServicePrincipal.ServicePrincipalType
type AdServicePrincipalDescription struct {
	TenantID           string
	AdServicePrincipal msgraph.ServicePrincipal
}

//  =================== postgresql ==================

//index:microsoft_dbforpostgresql_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type PostgresqlServerDescription struct {
	Server                       postgresql.Server
	ServerAdministratorResources *[]postgresql.ServerAdministratorResource
	Configurations               *[]postgresql.Configuration
	ServerKeys                   []postgresql.ServerKey
	FirewallRules                *[]postgresql.FirewallRule
	ResourceGroup                string
}

//  =================== storagesync ==================

//index:microsoft_storagesync_storagesyncservices
//getfilter:name=description.Service.name
//getfilter:resource_group=description.ResourceGroup
type StorageSyncDescription struct {
	Service       storagesync.Service
	ResourceGroup string
}

//  =================== sql ==================

//index:microsoft_sql_managedinstances
//getfilter:name=description.ManagedInstance.name
//getfilter:resource_group=description.ResourceGroup
type MssqlManagedInstanceDescription struct {
	ManagedInstance                         sqlv5.ManagedInstance
	ManagedInstanceVulnerabilityAssessments []sqlv5.ManagedInstanceVulnerabilityAssessment
	ManagedDatabaseSecurityAlertPolicies    []sqlv5.ManagedServerSecurityAlertPolicy
	ManagedInstanceEncryptionProtectors     []sqlv5.ManagedInstanceEncryptionProtector
	ResourceGroup                           string
}

//index:microsoft_sql_servers_databases
//getfilter:name=description.Database.name
//getfilter:resource_group=description.ResourceGroup
type SqlDatabaseDescription struct {
	Database                           sql.Database
	LongTermRetentionPolicy            sqlv5.LongTermRetentionPolicy
	TransparentDataEncryption          sql.TransparentDataEncryption
	DatabaseVulnerabilityAssessments   []sqlv5.DatabaseVulnerabilityAssessment
	VulnerabilityAssessmentScanRecords []sqlv5.VulnerabilityAssessmentScanRecord
	ResourceGroup                      string
}

//  =================== sqlv3 ==================

//index:microsoft_sql_servers
//getfilter:name=description.Server.name
//getfilter:resource_group=description.ResourceGroup
type SqlServerDescription struct {
	Server                         sqlv3.Server
	ServerBlobAuditingPolicies     []sql.ServerBlobAuditingPolicy
	ServerSecurityAlertPolicies    []sql.ServerSecurityAlertPolicy
	ServerAzureADAdministrators    *[]sql.ServerAzureADAdministrator
	ServerVulnerabilityAssessments []sqlv3.ServerVulnerabilityAssessment
	FirewallRules                  *[]sql.FirewallRule
	EncryptionProtectors           []sql.EncryptionProtector
	PrivateEndpointConnections     []sqlv3.PrivateEndpointConnection
	VirtualNetworkRules            []sql.VirtualNetworkRule
	ResourceGroup                  string
}

//index:microsoft_sql_elasticpools
//getfilter:name=description.Pool.Name
//getfilter:server_name=description.ServerName
//getfilter:resource_group=description.ResourceGroup
type SqlServerElasticPoolDescription struct {
	Pool          sql.ElasticPool
	ServerName    string
	ResourceGroup string
}

//index:microsoft_sql_virtualmachines
//getfilter:name=description.VirtualMachine.Name
//getfilter:resource_group=description.ResourceGroup
type SqlServerVirtualMachineDescription struct {
	VirtualMachine sqlvirtualmachine.SQLVirtualMachine
	ResourceGroup  string
}

//index:microsoft_sql_flexibleservers
//getfilter:name=description.FlexibleServer.Name
//getfilter:resource_group=description.ResourceGroup
type SqlServerFlexibleServerDescription struct {
	FlexibleServer mysqlflexibleservers.Server
	ResourceGroup  string
}

//  =================== storage ==================

//index:microsoft_classicstorage_storageaccounts
//getfilter:name=description.Account.name
//getfilter:resource_group=description.ResourceGroup
type StorageAccountDescription struct {
	Account                     storage.Account
	ManagementPolicy            *storage.ManagementPolicy
	BlobServiceProperties       *storage.BlobServiceProperties
	Logging                     *accounts.Logging
	StorageServiceProperties    *queues.StorageServiceProperties
	FileServiceProperties       *storage.FileServiceProperties
	DiagnosticSettingsResources *[]insights.DiagnosticSettingsResource
	EncryptionScopes            []storage.EncryptionScope
	ResourceGroup               string
}

//  =================== recoveryservice ==================

//index:microsoft_recoveryservices_vault
//getfilter:name=description.Vault.Name
//getfilter:resource_group=description.ResourceGroup
type RecoveryServicesVaultDescription struct {
	Vault         recoveryservices.Vault
	ResourceGroup string
}

//  =================== kubernetes ==================

//index:microsoft_hybridkubernetes_connectedcluster
//getfilter:name=description.ConnectedCluster.Name
//getfilter:resource_group=description.ResourceGroup
type HybridKubernetesConnectedClusterDescription struct {
	ConnectedCluster hybridkubernetes.ConnectedCluster
	ResourceGroup    string
}

//  =================== Cost ==================

type CostManagementQueryRow struct {
	UsageDate      int     `json:"UsageDate"`
	Cost           float64 `json:"Cost"`
	Currency       string  `json:"Currency"`
	ResourceType   *string `json:"resourceType,omitempty"`
	SubscriptionID *string `json:"SubscriptionId,omitempty"`
}

//index:microsoft_costmanagement_costbyresourcetype
type CostManagementCostByResourceTypeDescription struct {
	CostManagementCostByResourceType CostManagementQueryRow
}

//index:microsoft_costmanagement_costbysubscription
type CostManagementCostBySubscriptionDescription struct {
	CostManagementCostBySubscription CostManagementQueryRow
}

// =================== LB (loadbalancer) ==================

//index:microsoft_network_loadbalancers
//getfilter:name=description.LoadBalancer.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerDescription struct {
	LoadBalancer      newnetwork.LoadBalancer
	DiagnosticSetting *[]insights.DiagnosticSettingsResource
	ResourceGroup     string
}

//index:microsoft_lb_backendaddresspools
//getfilter:load_balancer_name=description.LoadBalancer.Name
//getfilter:name=description.Pool.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerBackendAddressPoolDescription struct {
	LoadBalancer  newnetwork.LoadBalancer
	Pool          newnetwork.BackendAddressPool
	ResourceGroup string
}

//index:microsoft_lb_natrules
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Rule.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerNatRuleDescription struct {
	Rule             newnetwork.InboundNatRule
	LoadBalancerName string
	ResourceGroup    string
}

//index:microsoft_lb_outboundrules
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Rule.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerOutboundRuleDescription struct {
	Rule             newnetwork.OutboundRule
	LoadBalancerName string
	ResourceGroup    string
}

//index:microsoft_lb_probes
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Probe.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerProbeDescription struct {
	Probe            newnetwork.Probe
	LoadBalancerName string
	ResourceGroup    string
}

//index:microsoft_lb_rules
//getfilter:load_balancer_name=description.LoadBalancerName
//getfilter:name=description.Rule.Name
//getfilter:resource_group=description.ResourceGroup
type LoadBalancerRuleDescription struct {
	Rule             newnetwork.LoadBalancingRule
	LoadBalancerName string
	ResourceGroup    string
}

// =================== Management ==================

//index:microsoft_management_groups
//getfilter:name=description.Group.Name
type ManagementGroupDescription struct {
	Group managementgroups.ManagementGroup
}

//index:microsoft_management_locks
//getfilter:name=description.Lock.Name
//getfilter:resource_group=description.ResourceGroup
type ManagementLockDescription struct {
	Lock          locks.ManagementLockObject
	ResourceGroup string
}

// =================== Resources ==================

//index:microsoft_resources_providers
//getfilter:namespace=description.Provider.Namespace
type ResourceProviderDescription struct {
	Provider resources.Provider
}

//index:microsoft_resources_resourcegroups
//getfilter:name=description.Group.Name
type ResourceGroupDescription struct {
	Group resources.Group
}
