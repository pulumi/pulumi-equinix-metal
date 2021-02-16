// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to attach elastic IP subnets to devices.
//
// To attach an IP subnet from a reserved block to a provisioned device, you must derive a subnet CIDR belonging to
// one of your reserved blocks in the same project and facility as the target device.
//
// For example, you have reserved IPv4 address block 147.229.10.152/30, you can choose to assign either the whole
// block as one subnet to a device; or 2 subnets with CIDRs 147.229.10.152/31' and 147.229.10.154/31; or 4 subnets
// with mask prefix length 32. More about the elastic IP subnets is [here](https://metal.equinix.com/developers/docs/networking/elastic-ips/).
//
// Device and reserved block must be in the same facility.
type IpAttachment struct {
	pulumi.CustomResourceState

	Address pulumi.StringOutput `pulumi:"address"`
	// Address family as integer (4 or 6)
	AddressFamily pulumi.IntOutput `pulumi:"addressFamily"`
	// length of CIDR prefix of the subnet as integer
	Cidr pulumi.IntOutput `pulumi:"cidr"`
	// CIDR notation of subnet from block reserved in the same
	// project and facility as the device
	CidrNotation pulumi.StringOutput `pulumi:"cidrNotation"`
	// ID of device to which to assign the subnet
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// IP address of gateway for the subnet
	Gateway    pulumi.StringOutput `pulumi:"gateway"`
	Global     pulumi.BoolOutput   `pulumi:"global"`
	Manageable pulumi.BoolOutput   `pulumi:"manageable"`
	Management pulumi.BoolOutput   `pulumi:"management"`
	// Subnet mask in decimal notation, e.g. "255.255.255.0"
	Netmask pulumi.StringOutput `pulumi:"netmask"`
	// Subnet network address
	Network pulumi.StringOutput `pulumi:"network"`
	// boolean flag whether subnet is reachable from the Internet
	Public pulumi.BoolOutput `pulumi:"public"`
}

// NewIpAttachment registers a new resource with the given unique name, arguments, and options.
func NewIpAttachment(ctx *pulumi.Context,
	name string, args *IpAttachmentArgs, opts ...pulumi.ResourceOption) (*IpAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrNotation == nil {
		return nil, errors.New("invalid value for required argument 'CidrNotation'")
	}
	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	var resource IpAttachment
	err := ctx.RegisterResource("equinix-metal:index/ipAttachment:IpAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpAttachment gets an existing IpAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpAttachmentState, opts ...pulumi.ResourceOption) (*IpAttachment, error) {
	var resource IpAttachment
	err := ctx.ReadResource("equinix-metal:index/ipAttachment:IpAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpAttachment resources.
type ipAttachmentState struct {
	Address *string `pulumi:"address"`
	// Address family as integer (4 or 6)
	AddressFamily *int `pulumi:"addressFamily"`
	// length of CIDR prefix of the subnet as integer
	Cidr *int `pulumi:"cidr"`
	// CIDR notation of subnet from block reserved in the same
	// project and facility as the device
	CidrNotation *string `pulumi:"cidrNotation"`
	// ID of device to which to assign the subnet
	DeviceId *string `pulumi:"deviceId"`
	// IP address of gateway for the subnet
	Gateway    *string `pulumi:"gateway"`
	Global     *bool   `pulumi:"global"`
	Manageable *bool   `pulumi:"manageable"`
	Management *bool   `pulumi:"management"`
	// Subnet mask in decimal notation, e.g. "255.255.255.0"
	Netmask *string `pulumi:"netmask"`
	// Subnet network address
	Network *string `pulumi:"network"`
	// boolean flag whether subnet is reachable from the Internet
	Public *bool `pulumi:"public"`
}

type IpAttachmentState struct {
	Address pulumi.StringPtrInput
	// Address family as integer (4 or 6)
	AddressFamily pulumi.IntPtrInput
	// length of CIDR prefix of the subnet as integer
	Cidr pulumi.IntPtrInput
	// CIDR notation of subnet from block reserved in the same
	// project and facility as the device
	CidrNotation pulumi.StringPtrInput
	// ID of device to which to assign the subnet
	DeviceId pulumi.StringPtrInput
	// IP address of gateway for the subnet
	Gateway    pulumi.StringPtrInput
	Global     pulumi.BoolPtrInput
	Manageable pulumi.BoolPtrInput
	Management pulumi.BoolPtrInput
	// Subnet mask in decimal notation, e.g. "255.255.255.0"
	Netmask pulumi.StringPtrInput
	// Subnet network address
	Network pulumi.StringPtrInput
	// boolean flag whether subnet is reachable from the Internet
	Public pulumi.BoolPtrInput
}

func (IpAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAttachmentState)(nil)).Elem()
}

type ipAttachmentArgs struct {
	// CIDR notation of subnet from block reserved in the same
	// project and facility as the device
	CidrNotation string `pulumi:"cidrNotation"`
	// ID of device to which to assign the subnet
	DeviceId string `pulumi:"deviceId"`
}

// The set of arguments for constructing a IpAttachment resource.
type IpAttachmentArgs struct {
	// CIDR notation of subnet from block reserved in the same
	// project and facility as the device
	CidrNotation pulumi.StringInput
	// ID of device to which to assign the subnet
	DeviceId pulumi.StringInput
}

func (IpAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAttachmentArgs)(nil)).Elem()
}

type IpAttachmentInput interface {
	pulumi.Input

	ToIpAttachmentOutput() IpAttachmentOutput
	ToIpAttachmentOutputWithContext(ctx context.Context) IpAttachmentOutput
}

func (*IpAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAttachment)(nil))
}

func (i *IpAttachment) ToIpAttachmentOutput() IpAttachmentOutput {
	return i.ToIpAttachmentOutputWithContext(context.Background())
}

func (i *IpAttachment) ToIpAttachmentOutputWithContext(ctx context.Context) IpAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentOutput)
}

func (i *IpAttachment) ToIpAttachmentPtrOutput() IpAttachmentPtrOutput {
	return i.ToIpAttachmentPtrOutputWithContext(context.Background())
}

func (i *IpAttachment) ToIpAttachmentPtrOutputWithContext(ctx context.Context) IpAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentPtrOutput)
}

type IpAttachmentPtrInput interface {
	pulumi.Input

	ToIpAttachmentPtrOutput() IpAttachmentPtrOutput
	ToIpAttachmentPtrOutputWithContext(ctx context.Context) IpAttachmentPtrOutput
}

type ipAttachmentPtrType IpAttachmentArgs

func (*ipAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAttachment)(nil))
}

func (i *ipAttachmentPtrType) ToIpAttachmentPtrOutput() IpAttachmentPtrOutput {
	return i.ToIpAttachmentPtrOutputWithContext(context.Background())
}

func (i *ipAttachmentPtrType) ToIpAttachmentPtrOutputWithContext(ctx context.Context) IpAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentPtrOutput)
}

// IpAttachmentArrayInput is an input type that accepts IpAttachmentArray and IpAttachmentArrayOutput values.
// You can construct a concrete instance of `IpAttachmentArrayInput` via:
//
//          IpAttachmentArray{ IpAttachmentArgs{...} }
type IpAttachmentArrayInput interface {
	pulumi.Input

	ToIpAttachmentArrayOutput() IpAttachmentArrayOutput
	ToIpAttachmentArrayOutputWithContext(context.Context) IpAttachmentArrayOutput
}

type IpAttachmentArray []IpAttachmentInput

func (IpAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IpAttachment)(nil))
}

func (i IpAttachmentArray) ToIpAttachmentArrayOutput() IpAttachmentArrayOutput {
	return i.ToIpAttachmentArrayOutputWithContext(context.Background())
}

func (i IpAttachmentArray) ToIpAttachmentArrayOutputWithContext(ctx context.Context) IpAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentArrayOutput)
}

// IpAttachmentMapInput is an input type that accepts IpAttachmentMap and IpAttachmentMapOutput values.
// You can construct a concrete instance of `IpAttachmentMapInput` via:
//
//          IpAttachmentMap{ "key": IpAttachmentArgs{...} }
type IpAttachmentMapInput interface {
	pulumi.Input

	ToIpAttachmentMapOutput() IpAttachmentMapOutput
	ToIpAttachmentMapOutputWithContext(context.Context) IpAttachmentMapOutput
}

type IpAttachmentMap map[string]IpAttachmentInput

func (IpAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IpAttachment)(nil))
}

func (i IpAttachmentMap) ToIpAttachmentMapOutput() IpAttachmentMapOutput {
	return i.ToIpAttachmentMapOutputWithContext(context.Background())
}

func (i IpAttachmentMap) ToIpAttachmentMapOutputWithContext(ctx context.Context) IpAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAttachmentMapOutput)
}

type IpAttachmentOutput struct {
	*pulumi.OutputState
}

func (IpAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAttachment)(nil))
}

func (o IpAttachmentOutput) ToIpAttachmentOutput() IpAttachmentOutput {
	return o
}

func (o IpAttachmentOutput) ToIpAttachmentOutputWithContext(ctx context.Context) IpAttachmentOutput {
	return o
}

func (o IpAttachmentOutput) ToIpAttachmentPtrOutput() IpAttachmentPtrOutput {
	return o.ToIpAttachmentPtrOutputWithContext(context.Background())
}

func (o IpAttachmentOutput) ToIpAttachmentPtrOutputWithContext(ctx context.Context) IpAttachmentPtrOutput {
	return o.ApplyT(func(v IpAttachment) *IpAttachment {
		return &v
	}).(IpAttachmentPtrOutput)
}

type IpAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (IpAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAttachment)(nil))
}

func (o IpAttachmentPtrOutput) ToIpAttachmentPtrOutput() IpAttachmentPtrOutput {
	return o
}

func (o IpAttachmentPtrOutput) ToIpAttachmentPtrOutputWithContext(ctx context.Context) IpAttachmentPtrOutput {
	return o
}

type IpAttachmentArrayOutput struct{ *pulumi.OutputState }

func (IpAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAttachment)(nil))
}

func (o IpAttachmentArrayOutput) ToIpAttachmentArrayOutput() IpAttachmentArrayOutput {
	return o
}

func (o IpAttachmentArrayOutput) ToIpAttachmentArrayOutputWithContext(ctx context.Context) IpAttachmentArrayOutput {
	return o
}

func (o IpAttachmentArrayOutput) Index(i pulumi.IntInput) IpAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAttachment {
		return vs[0].([]IpAttachment)[vs[1].(int)]
	}).(IpAttachmentOutput)
}

type IpAttachmentMapOutput struct{ *pulumi.OutputState }

func (IpAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IpAttachment)(nil))
}

func (o IpAttachmentMapOutput) ToIpAttachmentMapOutput() IpAttachmentMapOutput {
	return o
}

func (o IpAttachmentMapOutput) ToIpAttachmentMapOutputWithContext(ctx context.Context) IpAttachmentMapOutput {
	return o
}

func (o IpAttachmentMapOutput) MapIndex(k pulumi.StringInput) IpAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IpAttachment {
		return vs[0].(map[string]IpAttachment)[vs[1].(string)]
	}).(IpAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(IpAttachmentOutput{})
	pulumi.RegisterOutputType(IpAttachmentPtrOutput{})
	pulumi.RegisterOutputType(IpAttachmentArrayOutput{})
	pulumi.RegisterOutputType(IpAttachmentMapOutput{})
}
