package describer

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/synapse/mgmt/2021-03-01/synapse"
	"github.com/gdexlab/go-render/render"
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
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
				t.Errorf("JSONAllFieldsMarshaller.MarshalJSON() = %v, want %v", string(got), tt.want)
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
