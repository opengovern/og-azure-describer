package describer

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/synapse/mgmt/2021-03-01/synapse"
	"github.com/gdexlab/go-render/render"
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
	"time"
)

func TestJSONAllFieldsMarshaller(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{
			name: "Struct/Pointer",
			value: compute.VirtualMachine{
				ID:   String("MyVirtualMachine"),
				Type: String("MyVirtualMachineType"),
			},
			want: `{"ID":"MyVirtualMachine","Tags":null,"Type":"MyVirtualMachineType"}`,
		},
		{
			name: "Struct/Pointer 2",
			value: compute.VirtualMachine{
				ID:   String("MyVirtualMachine"),
				Type: String("MyVirtualMachineType"),
				Plan: &compute.Plan{
					Name:      String("MyPlan"),
					Publisher: String("MyPublisher"),
				},
			},
			want: `{"ID":"MyVirtualMachine","Plan":{"Name":"MyPlan","Publisher":"MyPublisher"},"Tags":null,"Type":"MyVirtualMachineType"}`,
		},
		{
			name: "Struct/Pointer/Slice",
			value: compute.VirtualMachine{
				ID:   String("MyVirtualMachine"),
				Type: String("MyVirtualMachineType"),
				Plan: &compute.Plan{
					Name:      String("MyPlan"),
					Publisher: String("MyPublisher"),
				},
				Resources: &[]compute.VirtualMachineExtension{
					{
						ID: String("MyVirtualMachineExtension"),
					},
				},
			},
			want: `{"ID":"MyVirtualMachine","Plan":{"Name":"MyPlan","Publisher":"MyPublisher"},"Resources":[{"ID":"MyVirtualMachineExtension","Tags":null}],"Tags":null,"Type":"MyVirtualMachineType"}`,
		},
		{
			name: "Array/Slice",
			value: compute.VirtualMachine{
				ID:   String("MyVirtualMachine"),
				Type: String("MyVirtualMachineType"),
				Resources: &[]compute.VirtualMachineExtension{
					{
						ID:   String("MyVirtualMachineExtension"),
						Name: String("MyVirtualMachineExtensionName"),
						VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
							Publisher: String("MyPublisher"),
						},
					},
				},
			},
			want: `{"ID":"MyVirtualMachine","Resources":[{"ID":"MyVirtualMachineExtension","Name":"MyVirtualMachineExtensionName","Tags":null,"VirtualMachineExtensionProperties":{"Publisher":"MyPublisher"}}],"Tags":null,"Type":"MyVirtualMachineType"}`,
		},
		{
			name: "UUID",
			value: synapse.Workspace{
				ID: String("MyWorkspace"),
				WorkspaceProperties: &synapse.WorkspaceProperties{
					WorkspaceUID: UUID(uuid.Must(uuid.FromString("7eae5af9-b353-4d53-89b6-15a1a664b2c2"))),
				},
			},
			want: `{"ID":"MyWorkspace","Tags":null,"WorkspaceProperties":{"ConnectivityEndpoints":null,"ExtraProperties":null,"WorkspaceUID":"7eae5af9-b353-4d53-89b6-15a1a664b2c2"}}`,
		},
		{
			name: "Enum",
			value: armautomation.Account{
				Etag: nil,
				Identity: &armautomation.Identity{
					Type: PTR(armautomation.ResourceIdentityTypeSystemAssigned),
				},
				Location:   nil,
				Properties: nil,
				Tags:       nil,
				ID:         nil,
				Name:       nil,
				SystemData: nil,
				Type:       nil,
			},
			want: `{"Etag":null,"ID":null,"Identity":{"PrincipalID":null,"TenantID":null,"Type":"SystemAssigned","UserAssignedIdentities":null},"Location":null,"Name":null,"Properties":null,"SystemData":null,"Tags":null,"Type":null}`,
		},
		{
			name: "Nested Struct",
			value: &armautomation.Account{
				Etag: nil,
				Identity: &armautomation.Identity{
					Type: PTR((armautomation.ResourceIdentityType)("UserAssigned")),
					UserAssignedIdentities: map[string]*armautomation.ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties{
						"/subscriptions/xyx/resourcegroups/yyy/providers/Microsoft.ManagedIdentity/userAssignedIdentities/yyz": &armautomation.ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties{
							ClientID:    nil,
							PrincipalID: nil,
						},
					},
					PrincipalID: nil,
					TenantID:    nil,
				},
				Location: PTR("westeurope"),
				Properties: &armautomation.AccountProperties{
					AutomationHybridServiceURL: nil,
					Description:                nil,
					DisableLocalAuth:           PTR(false),
					Encryption:                 nil,
					LastModifiedBy:             nil,
					PrivateEndpointConnections: nil,
					PublicNetworkAccess:        PTR(true),
					SKU:                        nil,
					CreationTime:               PTR(time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)),
					LastModifiedTime:           PTR(time.Date(2023, 2, 15, 0, 0, 0, 0, time.UTC)),
					State:                      nil,
				},
				Tags: map[string]*string{
					"app_support_group":   PTR("1"),
					"application_bit_id":  PTR("2"),
					"application_name":    PTR("3"),
					"bu_code":             PTR("4"),
					"business_owner":      PTR("5"),
					"data_classification": PTR("6"),
				},
				ID:         PTR("/subscriptions/xxx/resourceGroups/yyy/providers/Microsoft.Automation/automationAccounts/zzz"),
				Name:       PTR("zzz"),
				SystemData: nil,
				Type:       PTR("Microsoft.Automation/AutomationAccounts")},
			want: "{\"Etag\":null,\"ID\":\"/subscriptions/xxx/resourceGroups/yyy/providers/Microsoft.Automation/automationAccounts/zzz\",\"Identity\":{\"PrincipalID\":null,\"TenantID\":null,\"Type\":\"UserAssigned\",\"UserAssignedIdentities\":{\"/subscriptions/xyx/resourcegroups/yyy/providers/Microsoft.ManagedIdentity/userAssignedIdentities/yyz\":{}}},\"Location\":\"westeurope\",\"Name\":\"zzz\",\"Properties\":{\"AutomationHybridServiceURL\":null,\"CreationTime\":\"2022-12-01T00:00:00Z\",\"Description\":null,\"DisableLocalAuth\":false,\"Encryption\":null,\"LastModifiedBy\":null,\"LastModifiedTime\":\"2023-02-15T00:00:00Z\",\"PrivateEndpointConnections\":null,\"PublicNetworkAccess\":true,\"SKU\":null,\"State\":null},\"SystemData\":null,\"Tags\":{\"app_support_group\":\"1\",\"application_bit_id\":\"2\",\"application_name\":\"3\",\"bu_code\":\"4\",\"business_owner\":\"5\",\"data_classification\":\"6\"},\"Type\":\"Microsoft.Automation/AutomationAccounts\"}",
		},
		//The unmarshaller does not work for interfaces
		{
			name: "Interface struct",
			value: armcosmos.DatabaseAccountGetResults{
				Properties: &armcosmos.DatabaseAccountGetProperties{
					BackupPolicy: &armcosmos.BackupPolicy{
						Type:           PTR(armcosmos.BackupPolicyTypePeriodic),
						MigrationState: nil,
					},
				},
			},
			want: "{\"ID\":null,\"Identity\":null,\"Kind\":null,\"Location\":null,\"Name\":null,\"Properties\":{\"APIProperties\":null,\"AnalyticalStorageConfiguration\":null,\"BackupPolicy\":{\"MigrationState\":null,\"Type\":\"Periodic\"},\"Capabilities\":null,\"Capacity\":null,\"ConnectorOffer\":null,\"ConsistencyPolicy\":null,\"Cors\":null,\"CreateMode\":null,\"DatabaseAccountOfferType\":null,\"DefaultIdentity\":null,\"DisableKeyBasedMetadataWriteAccess\":null,\"DisableLocalAuth\":null,\"DocumentEndpoint\":null,\"EnableAnalyticalStorage\":null,\"EnableAutomaticFailover\":null,\"EnableCassandraConnector\":null,\"EnableFreeTier\":null,\"EnableMultipleWriteLocations\":null,\"EnablePartitionMerge\":null,\"FailoverPolicies\":null,\"IPRules\":null,\"InstanceID\":null,\"IsVirtualNetworkFilterEnabled\":null,\"KeyVaultKeyURI\":null,\"KeysMetadata\":null,\"Locations\":null,\"MinimalTLSVersion\":null,\"NetworkACLBypass\":null,\"NetworkACLBypassResourceIDs\":null,\"PrivateEndpointConnections\":null,\"ProvisioningState\":null,\"PublicNetworkAccess\":null,\"ReadLocations\":null,\"RestoreParameters\":null,\"VirtualNetworkRules\":null,\"WriteLocations\":null},\"SystemData\":null,\"Tags\":null,\"Type\":null}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := JSONAllFieldsMarshaller{
				Value: tt.value,
			}
			got, err := x.MarshalJSON()
			if err != nil {
				t.Errorf("JSONAllFieldsMarshaller.MarshalJSON() error = %v", err)
				return
			}
			if string(got) != tt.want {
				t.Errorf("JSONAllFieldsMarshaller.MarshalJSON() = %v\n want %v", string(got), tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			x := JSONAllFieldsMarshaller{
				Value: reflect.New(reflect.TypeOf(tt.value)).Elem().Interface(),
			}
			err := x.UnmarshalJSON([]byte(tt.want))
			if err != nil {
				t.Errorf("JSONAllFieldsMarshaller.MarshalJSON() error = %v", err)
				return
			}
			if render.AsCode(x.Value) != render.AsCode(tt.value) {
				t.Errorf("JSONAllFieldsMarshaller.UnmarshalJSON() = %v\nwant %v\noriginal: %s", render.AsCode(x.Value), render.AsCode(tt.value), tt.want)
			}
		})
	}
}

func String(s string) *string {
	return &s
}

func UUID(u uuid.UUID) *uuid.UUID {
	return &u
}

func PTR[T any](t T) *T {
	return &t
}
