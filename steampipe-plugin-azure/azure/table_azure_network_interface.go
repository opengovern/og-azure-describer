package azure

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/kaytu-io/kaytu-azure-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureNetworkInterface(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_network_interface",
		Description: "Azure Network Interface",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    kaytu.GetNetworkInterface,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListNetworkInterface,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the network interface",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a network interface uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the network interface",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Type")},
			{
				Name:        "provisioning_state",
				Description: "Providsioning state of the network interface resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Interface.Properties.ProvisioningState"),
			},
			{
				Name:        "enable_accelerated_networking",
				Description: "Indicates whether the network interface is accelerated networking enabled",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Interface.Properties.EnableAcceleratedNetworking")},
			{
				Name:        "enable_ip_forwarding",
				Description: "Indicates whether IP forwarding is enabled on this network interface",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Interface.Properties.EnableIPForwarding")},
			{
				Name:        "internal_dns_name_label",
				Description: "Relative DNS name for this NIC used for internal communications between VMs in the same virtual network",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Properties.DNSSettings.InternalDNSNameLabel")},
			{
				Name:        "internal_domain_name_suffix",
				Description: "Contains domain name suffix for the network interface",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Properties.DNSSettings.InternalDomainNameSuffix")},
			{
				Name:        "internal_fqdn",
				Description: "Fully qualified DNS name supporting internal communications between VMs in the same virtual network",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Properties.DNSSettings.InternalFqdn")},
			{
				Name:        "is_primary",
				Description: "Indicates whether this is a primary network interface on a virtual machine",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Interface.Properties.Primary")},
			{
				Name:        "mac_address",
				Description: "The MAC address of the network interface",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Properties.MacAddress")},
			{
				Name:        "network_security_group_id",
				Description: "The reference to the NetworkSecurityGroup resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Properties.NetworkSecurityGroup.ID")},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the network interface resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Properties.ResourceGUID")},
			{
				Name:        "virtual_machine_id",
				Description: "The reference to a virtual machine",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Properties.VirtualMachine.ID")},
			{
				Name:        "applied_dns_servers",
				Description: "A list of applied dns servers",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Interface.Properties.DNSSettings.AppliedDNSServers")},
			{
				Name:        "dns_servers",
				Description: "A collection of DNS servers IP addresses",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Interface.Properties.DNSSettings.DNSServers")},
			{
				Name:        "hosted_workloads",
				Description: "A collection of references to linked BareMetal resources",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Interface.Properties.HostedWorkloads")},
			{
				Name:        "ip_configurations",
				Description: "A list of IPConfigurations of the network interface",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Interface.Properties.IPConfigurations")},
			{
				Name:        "tap_configurations",
				Description: "A collection of TapConfigurations of the network interface",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Interface.Properties.TapConfigurations")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Interface.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Interface.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Interface.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Interface.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// FETCH FUNCTIONS ////

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS ////

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data