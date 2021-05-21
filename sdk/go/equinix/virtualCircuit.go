// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to associate VLAN with a Dedicated Port from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/#associating-a-vlan-with-a-dedicated-port).
//
// ## Example Usage
//
// Pick an existing Project and Connection, create a VLAN and use `VirtualCircuit` to associate it with a Primary Port of the Connection.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v2/go/equinix-metal"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		projectId := "52000fb2-ee46-4673-93a8-de2c2bdba33c"
// 		connId := "73f12f29-3e19-43a0-8e90-ae81580db1e0"
// 		testConnection, err := equinix - metal.LookupConnection(ctx, &equinix-metal.LookupConnectionArgs{
// 			ConnectionId: connId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		testVlan, err := equinix - metal.NewVlan(ctx, "testVlan", &equinix-metal.VlanArgs{
// 			ProjectId: pulumi.String(projectId),
// 			Metro:     pulumi.String(testConnection.Metro),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.NewVirtualCircuit(ctx, "testVirtualCircuit", &equinix-metal.VirtualCircuitArgs{
// 			ConnectionId: pulumi.String(connId),
// 			ProjectId:    pulumi.String(projectId),
// 			PortId:       pulumi.String(testConnection.Ports[0].Id),
// 			VlanId:       testVlan.ID(),
// 			NniVlan:      pulumi.Int(1056),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VirtualCircuit struct {
	pulumi.CustomResourceState

	// UUID of Connection where the VC is scoped to
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// Name of the Virtual Circuit resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Equinix Metal network-to-network VLAN ID
	NniVlan pulumi.IntOutput `pulumi:"nniVlan"`
	// Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	NniVnid pulumi.IntOutput `pulumi:"nniVnid"`
	// UUID of the Connection Port where the VC is scoped to
	PortId pulumi.StringOutput `pulumi:"portId"`
	// UUID of the Project where the VC is scoped to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Status of the virtal circuit
	// * `vnid`
	Status pulumi.StringOutput `pulumi:"status"`
	// UUID of the VLAN to associate
	VlanId pulumi.StringOutput `pulumi:"vlanId"`
	// VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	Vnid pulumi.IntOutput `pulumi:"vnid"`
}

// NewVirtualCircuit registers a new resource with the given unique name, arguments, and options.
func NewVirtualCircuit(ctx *pulumi.Context,
	name string, args *VirtualCircuitArgs, opts ...pulumi.ResourceOption) (*VirtualCircuit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.NniVlan == nil {
		return nil, errors.New("invalid value for required argument 'NniVlan'")
	}
	if args.PortId == nil {
		return nil, errors.New("invalid value for required argument 'PortId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	var resource VirtualCircuit
	err := ctx.RegisterResource("equinix-metal:index/virtualCircuit:VirtualCircuit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualCircuit gets an existing VirtualCircuit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualCircuit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualCircuitState, opts ...pulumi.ResourceOption) (*VirtualCircuit, error) {
	var resource VirtualCircuit
	err := ctx.ReadResource("equinix-metal:index/virtualCircuit:VirtualCircuit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualCircuit resources.
type virtualCircuitState struct {
	// UUID of Connection where the VC is scoped to
	ConnectionId *string `pulumi:"connectionId"`
	// Name of the Virtual Circuit resource
	Name *string `pulumi:"name"`
	// Equinix Metal network-to-network VLAN ID
	NniVlan *int `pulumi:"nniVlan"`
	// Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	NniVnid *int `pulumi:"nniVnid"`
	// UUID of the Connection Port where the VC is scoped to
	PortId *string `pulumi:"portId"`
	// UUID of the Project where the VC is scoped to
	ProjectId *string `pulumi:"projectId"`
	// Status of the virtal circuit
	// * `vnid`
	Status *string `pulumi:"status"`
	// UUID of the VLAN to associate
	VlanId *string `pulumi:"vlanId"`
	// VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	Vnid *int `pulumi:"vnid"`
}

type VirtualCircuitState struct {
	// UUID of Connection where the VC is scoped to
	ConnectionId pulumi.StringPtrInput
	// Name of the Virtual Circuit resource
	Name pulumi.StringPtrInput
	// Equinix Metal network-to-network VLAN ID
	NniVlan pulumi.IntPtrInput
	// Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	NniVnid pulumi.IntPtrInput
	// UUID of the Connection Port where the VC is scoped to
	PortId pulumi.StringPtrInput
	// UUID of the Project where the VC is scoped to
	ProjectId pulumi.StringPtrInput
	// Status of the virtal circuit
	// * `vnid`
	Status pulumi.StringPtrInput
	// UUID of the VLAN to associate
	VlanId pulumi.StringPtrInput
	// VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
	Vnid pulumi.IntPtrInput
}

func (VirtualCircuitState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualCircuitState)(nil)).Elem()
}

type virtualCircuitArgs struct {
	// UUID of Connection where the VC is scoped to
	ConnectionId string `pulumi:"connectionId"`
	// Name of the Virtual Circuit resource
	Name *string `pulumi:"name"`
	// Equinix Metal network-to-network VLAN ID
	NniVlan int `pulumi:"nniVlan"`
	// UUID of the Connection Port where the VC is scoped to
	PortId string `pulumi:"portId"`
	// UUID of the Project where the VC is scoped to
	ProjectId string `pulumi:"projectId"`
	// UUID of the VLAN to associate
	VlanId string `pulumi:"vlanId"`
}

// The set of arguments for constructing a VirtualCircuit resource.
type VirtualCircuitArgs struct {
	// UUID of Connection where the VC is scoped to
	ConnectionId pulumi.StringInput
	// Name of the Virtual Circuit resource
	Name pulumi.StringPtrInput
	// Equinix Metal network-to-network VLAN ID
	NniVlan pulumi.IntInput
	// UUID of the Connection Port where the VC is scoped to
	PortId pulumi.StringInput
	// UUID of the Project where the VC is scoped to
	ProjectId pulumi.StringInput
	// UUID of the VLAN to associate
	VlanId pulumi.StringInput
}

func (VirtualCircuitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualCircuitArgs)(nil)).Elem()
}

type VirtualCircuitInput interface {
	pulumi.Input

	ToVirtualCircuitOutput() VirtualCircuitOutput
	ToVirtualCircuitOutputWithContext(ctx context.Context) VirtualCircuitOutput
}

func (*VirtualCircuit) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualCircuit)(nil))
}

func (i *VirtualCircuit) ToVirtualCircuitOutput() VirtualCircuitOutput {
	return i.ToVirtualCircuitOutputWithContext(context.Background())
}

func (i *VirtualCircuit) ToVirtualCircuitOutputWithContext(ctx context.Context) VirtualCircuitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCircuitOutput)
}

func (i *VirtualCircuit) ToVirtualCircuitPtrOutput() VirtualCircuitPtrOutput {
	return i.ToVirtualCircuitPtrOutputWithContext(context.Background())
}

func (i *VirtualCircuit) ToVirtualCircuitPtrOutputWithContext(ctx context.Context) VirtualCircuitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCircuitPtrOutput)
}

type VirtualCircuitPtrInput interface {
	pulumi.Input

	ToVirtualCircuitPtrOutput() VirtualCircuitPtrOutput
	ToVirtualCircuitPtrOutputWithContext(ctx context.Context) VirtualCircuitPtrOutput
}

type virtualCircuitPtrType VirtualCircuitArgs

func (*virtualCircuitPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCircuit)(nil))
}

func (i *virtualCircuitPtrType) ToVirtualCircuitPtrOutput() VirtualCircuitPtrOutput {
	return i.ToVirtualCircuitPtrOutputWithContext(context.Background())
}

func (i *virtualCircuitPtrType) ToVirtualCircuitPtrOutputWithContext(ctx context.Context) VirtualCircuitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCircuitPtrOutput)
}

// VirtualCircuitArrayInput is an input type that accepts VirtualCircuitArray and VirtualCircuitArrayOutput values.
// You can construct a concrete instance of `VirtualCircuitArrayInput` via:
//
//          VirtualCircuitArray{ VirtualCircuitArgs{...} }
type VirtualCircuitArrayInput interface {
	pulumi.Input

	ToVirtualCircuitArrayOutput() VirtualCircuitArrayOutput
	ToVirtualCircuitArrayOutputWithContext(context.Context) VirtualCircuitArrayOutput
}

type VirtualCircuitArray []VirtualCircuitInput

func (VirtualCircuitArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VirtualCircuit)(nil))
}

func (i VirtualCircuitArray) ToVirtualCircuitArrayOutput() VirtualCircuitArrayOutput {
	return i.ToVirtualCircuitArrayOutputWithContext(context.Background())
}

func (i VirtualCircuitArray) ToVirtualCircuitArrayOutputWithContext(ctx context.Context) VirtualCircuitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCircuitArrayOutput)
}

// VirtualCircuitMapInput is an input type that accepts VirtualCircuitMap and VirtualCircuitMapOutput values.
// You can construct a concrete instance of `VirtualCircuitMapInput` via:
//
//          VirtualCircuitMap{ "key": VirtualCircuitArgs{...} }
type VirtualCircuitMapInput interface {
	pulumi.Input

	ToVirtualCircuitMapOutput() VirtualCircuitMapOutput
	ToVirtualCircuitMapOutputWithContext(context.Context) VirtualCircuitMapOutput
}

type VirtualCircuitMap map[string]VirtualCircuitInput

func (VirtualCircuitMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VirtualCircuit)(nil))
}

func (i VirtualCircuitMap) ToVirtualCircuitMapOutput() VirtualCircuitMapOutput {
	return i.ToVirtualCircuitMapOutputWithContext(context.Background())
}

func (i VirtualCircuitMap) ToVirtualCircuitMapOutputWithContext(ctx context.Context) VirtualCircuitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCircuitMapOutput)
}

type VirtualCircuitOutput struct {
	*pulumi.OutputState
}

func (VirtualCircuitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualCircuit)(nil))
}

func (o VirtualCircuitOutput) ToVirtualCircuitOutput() VirtualCircuitOutput {
	return o
}

func (o VirtualCircuitOutput) ToVirtualCircuitOutputWithContext(ctx context.Context) VirtualCircuitOutput {
	return o
}

func (o VirtualCircuitOutput) ToVirtualCircuitPtrOutput() VirtualCircuitPtrOutput {
	return o.ToVirtualCircuitPtrOutputWithContext(context.Background())
}

func (o VirtualCircuitOutput) ToVirtualCircuitPtrOutputWithContext(ctx context.Context) VirtualCircuitPtrOutput {
	return o.ApplyT(func(v VirtualCircuit) *VirtualCircuit {
		return &v
	}).(VirtualCircuitPtrOutput)
}

type VirtualCircuitPtrOutput struct {
	*pulumi.OutputState
}

func (VirtualCircuitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCircuit)(nil))
}

func (o VirtualCircuitPtrOutput) ToVirtualCircuitPtrOutput() VirtualCircuitPtrOutput {
	return o
}

func (o VirtualCircuitPtrOutput) ToVirtualCircuitPtrOutputWithContext(ctx context.Context) VirtualCircuitPtrOutput {
	return o
}

type VirtualCircuitArrayOutput struct{ *pulumi.OutputState }

func (VirtualCircuitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualCircuit)(nil))
}

func (o VirtualCircuitArrayOutput) ToVirtualCircuitArrayOutput() VirtualCircuitArrayOutput {
	return o
}

func (o VirtualCircuitArrayOutput) ToVirtualCircuitArrayOutputWithContext(ctx context.Context) VirtualCircuitArrayOutput {
	return o
}

func (o VirtualCircuitArrayOutput) Index(i pulumi.IntInput) VirtualCircuitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualCircuit {
		return vs[0].([]VirtualCircuit)[vs[1].(int)]
	}).(VirtualCircuitOutput)
}

type VirtualCircuitMapOutput struct{ *pulumi.OutputState }

func (VirtualCircuitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualCircuit)(nil))
}

func (o VirtualCircuitMapOutput) ToVirtualCircuitMapOutput() VirtualCircuitMapOutput {
	return o
}

func (o VirtualCircuitMapOutput) ToVirtualCircuitMapOutputWithContext(ctx context.Context) VirtualCircuitMapOutput {
	return o
}

func (o VirtualCircuitMapOutput) MapIndex(k pulumi.StringInput) VirtualCircuitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualCircuit {
		return vs[0].(map[string]VirtualCircuit)[vs[1].(string)]
	}).(VirtualCircuitOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualCircuitOutput{})
	pulumi.RegisterOutputType(VirtualCircuitPtrOutput{})
	pulumi.RegisterOutputType(VirtualCircuitArrayOutput{})
	pulumi.RegisterOutputType(VirtualCircuitMapOutput{})
}
