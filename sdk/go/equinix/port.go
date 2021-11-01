// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Port struct {
	pulumi.CustomResourceState

	// UUID of the bond port
	BondId pulumi.StringOutput `pulumi:"bondId"`
	// Name of the bond port
	BondName pulumi.StringOutput `pulumi:"bondName"`
	// Whether the port should be bonded
	Bonded pulumi.BoolOutput `pulumi:"bonded"`
	// Flag indicating whether the port can be removed from a bond
	DisbondSupported pulumi.BoolOutput `pulumi:"disbondSupported"`
	// Whether to put the port to Layer 2 mode, valid only for bond ports
	Layer2 pulumi.BoolPtrOutput `pulumi:"layer2"`
	// MAC address of the port
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Name of the port, e.g. `bond0` or `eth0`
	Name pulumi.StringOutput `pulumi:"name"`
	// UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
	NativeVlanId pulumi.StringPtrOutput `pulumi:"nativeVlanId"`
	// One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// ID of the port to read
	PortId pulumi.StringOutput `pulumi:"portId"`
	// Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
	ResetOnDelete pulumi.BoolPtrOutput `pulumi:"resetOnDelete"`
	// Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
	Type pulumi.StringOutput `pulumi:"type"`
	// List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
	VlanIds pulumi.StringArrayOutput `pulumi:"vlanIds"`
	// List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
	VxlanIds pulumi.IntArrayOutput `pulumi:"vxlanIds"`
}

// NewPort registers a new resource with the given unique name, arguments, and options.
func NewPort(ctx *pulumi.Context,
	name string, args *PortArgs, opts ...pulumi.ResourceOption) (*Port, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bonded == nil {
		return nil, errors.New("invalid value for required argument 'Bonded'")
	}
	if args.PortId == nil {
		return nil, errors.New("invalid value for required argument 'PortId'")
	}
	var resource Port
	err := ctx.RegisterResource("equinix-metal:index/port:Port", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPort gets an existing Port resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortState, opts ...pulumi.ResourceOption) (*Port, error) {
	var resource Port
	err := ctx.ReadResource("equinix-metal:index/port:Port", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Port resources.
type portState struct {
	// UUID of the bond port
	BondId *string `pulumi:"bondId"`
	// Name of the bond port
	BondName *string `pulumi:"bondName"`
	// Whether the port should be bonded
	Bonded *bool `pulumi:"bonded"`
	// Flag indicating whether the port can be removed from a bond
	DisbondSupported *bool `pulumi:"disbondSupported"`
	// Whether to put the port to Layer 2 mode, valid only for bond ports
	Layer2 *bool `pulumi:"layer2"`
	// MAC address of the port
	Mac *string `pulumi:"mac"`
	// Name of the port, e.g. `bond0` or `eth0`
	Name *string `pulumi:"name"`
	// UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
	NativeVlanId *string `pulumi:"nativeVlanId"`
	// One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
	NetworkType *string `pulumi:"networkType"`
	// ID of the port to read
	PortId *string `pulumi:"portId"`
	// Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
	ResetOnDelete *bool `pulumi:"resetOnDelete"`
	// Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
	Type *string `pulumi:"type"`
	// List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
	VlanIds []string `pulumi:"vlanIds"`
	// List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
	VxlanIds []int `pulumi:"vxlanIds"`
}

type PortState struct {
	// UUID of the bond port
	BondId pulumi.StringPtrInput
	// Name of the bond port
	BondName pulumi.StringPtrInput
	// Whether the port should be bonded
	Bonded pulumi.BoolPtrInput
	// Flag indicating whether the port can be removed from a bond
	DisbondSupported pulumi.BoolPtrInput
	// Whether to put the port to Layer 2 mode, valid only for bond ports
	Layer2 pulumi.BoolPtrInput
	// MAC address of the port
	Mac pulumi.StringPtrInput
	// Name of the port, e.g. `bond0` or `eth0`
	Name pulumi.StringPtrInput
	// UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
	NativeVlanId pulumi.StringPtrInput
	// One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
	NetworkType pulumi.StringPtrInput
	// ID of the port to read
	PortId pulumi.StringPtrInput
	// Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
	ResetOnDelete pulumi.BoolPtrInput
	// Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
	Type pulumi.StringPtrInput
	// List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
	VlanIds pulumi.StringArrayInput
	// List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
	VxlanIds pulumi.IntArrayInput
}

func (PortState) ElementType() reflect.Type {
	return reflect.TypeOf((*portState)(nil)).Elem()
}

type portArgs struct {
	// Whether the port should be bonded
	Bonded bool `pulumi:"bonded"`
	// Whether to put the port to Layer 2 mode, valid only for bond ports
	Layer2 *bool `pulumi:"layer2"`
	// UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
	NativeVlanId *string `pulumi:"nativeVlanId"`
	// ID of the port to read
	PortId string `pulumi:"portId"`
	// Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
	ResetOnDelete *bool `pulumi:"resetOnDelete"`
	// List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
	VlanIds []string `pulumi:"vlanIds"`
	// List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
	VxlanIds []int `pulumi:"vxlanIds"`
}

// The set of arguments for constructing a Port resource.
type PortArgs struct {
	// Whether the port should be bonded
	Bonded pulumi.BoolInput
	// Whether to put the port to Layer 2 mode, valid only for bond ports
	Layer2 pulumi.BoolPtrInput
	// UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
	NativeVlanId pulumi.StringPtrInput
	// ID of the port to read
	PortId pulumi.StringInput
	// Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
	ResetOnDelete pulumi.BoolPtrInput
	// List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
	VlanIds pulumi.StringArrayInput
	// List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
	VxlanIds pulumi.IntArrayInput
}

func (PortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portArgs)(nil)).Elem()
}

type PortInput interface {
	pulumi.Input

	ToPortOutput() PortOutput
	ToPortOutputWithContext(ctx context.Context) PortOutput
}

func (*Port) ElementType() reflect.Type {
	return reflect.TypeOf((*Port)(nil))
}

func (i *Port) ToPortOutput() PortOutput {
	return i.ToPortOutputWithContext(context.Background())
}

func (i *Port) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortOutput)
}

func (i *Port) ToPortPtrOutput() PortPtrOutput {
	return i.ToPortPtrOutputWithContext(context.Background())
}

func (i *Port) ToPortPtrOutputWithContext(ctx context.Context) PortPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortPtrOutput)
}

type PortPtrInput interface {
	pulumi.Input

	ToPortPtrOutput() PortPtrOutput
	ToPortPtrOutputWithContext(ctx context.Context) PortPtrOutput
}

type portPtrType PortArgs

func (*portPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Port)(nil))
}

func (i *portPtrType) ToPortPtrOutput() PortPtrOutput {
	return i.ToPortPtrOutputWithContext(context.Background())
}

func (i *portPtrType) ToPortPtrOutputWithContext(ctx context.Context) PortPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortPtrOutput)
}

// PortArrayInput is an input type that accepts PortArray and PortArrayOutput values.
// You can construct a concrete instance of `PortArrayInput` via:
//
//          PortArray{ PortArgs{...} }
type PortArrayInput interface {
	pulumi.Input

	ToPortArrayOutput() PortArrayOutput
	ToPortArrayOutputWithContext(context.Context) PortArrayOutput
}

type PortArray []PortInput

func (PortArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Port)(nil))
}

func (i PortArray) ToPortArrayOutput() PortArrayOutput {
	return i.ToPortArrayOutputWithContext(context.Background())
}

func (i PortArray) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortArrayOutput)
}

// PortMapInput is an input type that accepts PortMap and PortMapOutput values.
// You can construct a concrete instance of `PortMapInput` via:
//
//          PortMap{ "key": PortArgs{...} }
type PortMapInput interface {
	pulumi.Input

	ToPortMapOutput() PortMapOutput
	ToPortMapOutputWithContext(context.Context) PortMapOutput
}

type PortMap map[string]PortInput

func (PortMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Port)(nil))
}

func (i PortMap) ToPortMapOutput() PortMapOutput {
	return i.ToPortMapOutputWithContext(context.Background())
}

func (i PortMap) ToPortMapOutputWithContext(ctx context.Context) PortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortMapOutput)
}

type PortOutput struct {
	*pulumi.OutputState
}

func (PortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Port)(nil))
}

func (o PortOutput) ToPortOutput() PortOutput {
	return o
}

func (o PortOutput) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return o
}

func (o PortOutput) ToPortPtrOutput() PortPtrOutput {
	return o.ToPortPtrOutputWithContext(context.Background())
}

func (o PortOutput) ToPortPtrOutputWithContext(ctx context.Context) PortPtrOutput {
	return o.ApplyT(func(v Port) *Port {
		return &v
	}).(PortPtrOutput)
}

type PortPtrOutput struct {
	*pulumi.OutputState
}

func (PortPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Port)(nil))
}

func (o PortPtrOutput) ToPortPtrOutput() PortPtrOutput {
	return o
}

func (o PortPtrOutput) ToPortPtrOutputWithContext(ctx context.Context) PortPtrOutput {
	return o
}

type PortArrayOutput struct{ *pulumi.OutputState }

func (PortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Port)(nil))
}

func (o PortArrayOutput) ToPortArrayOutput() PortArrayOutput {
	return o
}

func (o PortArrayOutput) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return o
}

func (o PortArrayOutput) Index(i pulumi.IntInput) PortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Port {
		return vs[0].([]Port)[vs[1].(int)]
	}).(PortOutput)
}

type PortMapOutput struct{ *pulumi.OutputState }

func (PortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Port)(nil))
}

func (o PortMapOutput) ToPortMapOutput() PortMapOutput {
	return o
}

func (o PortMapOutput) ToPortMapOutputWithContext(ctx context.Context) PortMapOutput {
	return o
}

func (o PortMapOutput) MapIndex(k pulumi.StringInput) PortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Port {
		return vs[0].(map[string]Port)[vs[1].(string)]
	}).(PortOutput)
}

func init() {
	pulumi.RegisterOutputType(PortOutput{})
	pulumi.RegisterOutputType(PortPtrOutput{})
	pulumi.RegisterOutputType(PortArrayOutput{})
	pulumi.RegisterOutputType(PortMapOutput{})
}
