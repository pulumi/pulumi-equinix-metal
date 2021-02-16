// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// This resource can also be imported using existing device ID
//
// ```sh
//  $ pulumi import equinix-metal:index/deviceNetworkType:DeviceNetworkType metal_device_network_type {existing device_id}
// ```
type DeviceNetworkType struct {
	pulumi.CustomResourceState

	// The ID of the device on which the network type should be set.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeviceNetworkType registers a new resource with the given unique name, arguments, and options.
func NewDeviceNetworkType(ctx *pulumi.Context,
	name string, args *DeviceNetworkTypeArgs, opts ...pulumi.ResourceOption) (*DeviceNetworkType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource DeviceNetworkType
	err := ctx.RegisterResource("equinix-metal:index/deviceNetworkType:DeviceNetworkType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceNetworkType gets an existing DeviceNetworkType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceNetworkType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceNetworkTypeState, opts ...pulumi.ResourceOption) (*DeviceNetworkType, error) {
	var resource DeviceNetworkType
	err := ctx.ReadResource("equinix-metal:index/deviceNetworkType:DeviceNetworkType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceNetworkType resources.
type deviceNetworkTypeState struct {
	// The ID of the device on which the network type should be set.
	DeviceId *string `pulumi:"deviceId"`
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
	Type *string `pulumi:"type"`
}

type DeviceNetworkTypeState struct {
	// The ID of the device on which the network type should be set.
	DeviceId pulumi.StringPtrInput
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
	Type pulumi.StringPtrInput
}

func (DeviceNetworkTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceNetworkTypeState)(nil)).Elem()
}

type deviceNetworkTypeArgs struct {
	// The ID of the device on which the network type should be set.
	DeviceId string `pulumi:"deviceId"`
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DeviceNetworkType resource.
type DeviceNetworkTypeArgs struct {
	// The ID of the device on which the network type should be set.
	DeviceId pulumi.StringInput
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
	Type pulumi.StringInput
}

func (DeviceNetworkTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceNetworkTypeArgs)(nil)).Elem()
}

type DeviceNetworkTypeInput interface {
	pulumi.Input

	ToDeviceNetworkTypeOutput() DeviceNetworkTypeOutput
	ToDeviceNetworkTypeOutputWithContext(ctx context.Context) DeviceNetworkTypeOutput
}

func (*DeviceNetworkType) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceNetworkType)(nil))
}

func (i *DeviceNetworkType) ToDeviceNetworkTypeOutput() DeviceNetworkTypeOutput {
	return i.ToDeviceNetworkTypeOutputWithContext(context.Background())
}

func (i *DeviceNetworkType) ToDeviceNetworkTypeOutputWithContext(ctx context.Context) DeviceNetworkTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceNetworkTypeOutput)
}

func (i *DeviceNetworkType) ToDeviceNetworkTypePtrOutput() DeviceNetworkTypePtrOutput {
	return i.ToDeviceNetworkTypePtrOutputWithContext(context.Background())
}

func (i *DeviceNetworkType) ToDeviceNetworkTypePtrOutputWithContext(ctx context.Context) DeviceNetworkTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceNetworkTypePtrOutput)
}

type DeviceNetworkTypePtrInput interface {
	pulumi.Input

	ToDeviceNetworkTypePtrOutput() DeviceNetworkTypePtrOutput
	ToDeviceNetworkTypePtrOutputWithContext(ctx context.Context) DeviceNetworkTypePtrOutput
}

type deviceNetworkTypePtrType DeviceNetworkTypeArgs

func (*deviceNetworkTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceNetworkType)(nil))
}

func (i *deviceNetworkTypePtrType) ToDeviceNetworkTypePtrOutput() DeviceNetworkTypePtrOutput {
	return i.ToDeviceNetworkTypePtrOutputWithContext(context.Background())
}

func (i *deviceNetworkTypePtrType) ToDeviceNetworkTypePtrOutputWithContext(ctx context.Context) DeviceNetworkTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceNetworkTypePtrOutput)
}

// DeviceNetworkTypeArrayInput is an input type that accepts DeviceNetworkTypeArray and DeviceNetworkTypeArrayOutput values.
// You can construct a concrete instance of `DeviceNetworkTypeArrayInput` via:
//
//          DeviceNetworkTypeArray{ DeviceNetworkTypeArgs{...} }
type DeviceNetworkTypeArrayInput interface {
	pulumi.Input

	ToDeviceNetworkTypeArrayOutput() DeviceNetworkTypeArrayOutput
	ToDeviceNetworkTypeArrayOutputWithContext(context.Context) DeviceNetworkTypeArrayOutput
}

type DeviceNetworkTypeArray []DeviceNetworkTypeInput

func (DeviceNetworkTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DeviceNetworkType)(nil))
}

func (i DeviceNetworkTypeArray) ToDeviceNetworkTypeArrayOutput() DeviceNetworkTypeArrayOutput {
	return i.ToDeviceNetworkTypeArrayOutputWithContext(context.Background())
}

func (i DeviceNetworkTypeArray) ToDeviceNetworkTypeArrayOutputWithContext(ctx context.Context) DeviceNetworkTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceNetworkTypeArrayOutput)
}

// DeviceNetworkTypeMapInput is an input type that accepts DeviceNetworkTypeMap and DeviceNetworkTypeMapOutput values.
// You can construct a concrete instance of `DeviceNetworkTypeMapInput` via:
//
//          DeviceNetworkTypeMap{ "key": DeviceNetworkTypeArgs{...} }
type DeviceNetworkTypeMapInput interface {
	pulumi.Input

	ToDeviceNetworkTypeMapOutput() DeviceNetworkTypeMapOutput
	ToDeviceNetworkTypeMapOutputWithContext(context.Context) DeviceNetworkTypeMapOutput
}

type DeviceNetworkTypeMap map[string]DeviceNetworkTypeInput

func (DeviceNetworkTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DeviceNetworkType)(nil))
}

func (i DeviceNetworkTypeMap) ToDeviceNetworkTypeMapOutput() DeviceNetworkTypeMapOutput {
	return i.ToDeviceNetworkTypeMapOutputWithContext(context.Background())
}

func (i DeviceNetworkTypeMap) ToDeviceNetworkTypeMapOutputWithContext(ctx context.Context) DeviceNetworkTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceNetworkTypeMapOutput)
}

type DeviceNetworkTypeOutput struct {
	*pulumi.OutputState
}

func (DeviceNetworkTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceNetworkType)(nil))
}

func (o DeviceNetworkTypeOutput) ToDeviceNetworkTypeOutput() DeviceNetworkTypeOutput {
	return o
}

func (o DeviceNetworkTypeOutput) ToDeviceNetworkTypeOutputWithContext(ctx context.Context) DeviceNetworkTypeOutput {
	return o
}

func (o DeviceNetworkTypeOutput) ToDeviceNetworkTypePtrOutput() DeviceNetworkTypePtrOutput {
	return o.ToDeviceNetworkTypePtrOutputWithContext(context.Background())
}

func (o DeviceNetworkTypeOutput) ToDeviceNetworkTypePtrOutputWithContext(ctx context.Context) DeviceNetworkTypePtrOutput {
	return o.ApplyT(func(v DeviceNetworkType) *DeviceNetworkType {
		return &v
	}).(DeviceNetworkTypePtrOutput)
}

type DeviceNetworkTypePtrOutput struct {
	*pulumi.OutputState
}

func (DeviceNetworkTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceNetworkType)(nil))
}

func (o DeviceNetworkTypePtrOutput) ToDeviceNetworkTypePtrOutput() DeviceNetworkTypePtrOutput {
	return o
}

func (o DeviceNetworkTypePtrOutput) ToDeviceNetworkTypePtrOutputWithContext(ctx context.Context) DeviceNetworkTypePtrOutput {
	return o
}

type DeviceNetworkTypeArrayOutput struct{ *pulumi.OutputState }

func (DeviceNetworkTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeviceNetworkType)(nil))
}

func (o DeviceNetworkTypeArrayOutput) ToDeviceNetworkTypeArrayOutput() DeviceNetworkTypeArrayOutput {
	return o
}

func (o DeviceNetworkTypeArrayOutput) ToDeviceNetworkTypeArrayOutputWithContext(ctx context.Context) DeviceNetworkTypeArrayOutput {
	return o
}

func (o DeviceNetworkTypeArrayOutput) Index(i pulumi.IntInput) DeviceNetworkTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeviceNetworkType {
		return vs[0].([]DeviceNetworkType)[vs[1].(int)]
	}).(DeviceNetworkTypeOutput)
}

type DeviceNetworkTypeMapOutput struct{ *pulumi.OutputState }

func (DeviceNetworkTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeviceNetworkType)(nil))
}

func (o DeviceNetworkTypeMapOutput) ToDeviceNetworkTypeMapOutput() DeviceNetworkTypeMapOutput {
	return o
}

func (o DeviceNetworkTypeMapOutput) ToDeviceNetworkTypeMapOutputWithContext(ctx context.Context) DeviceNetworkTypeMapOutput {
	return o
}

func (o DeviceNetworkTypeMapOutput) MapIndex(k pulumi.StringInput) DeviceNetworkTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeviceNetworkType {
		return vs[0].(map[string]DeviceNetworkType)[vs[1].(string)]
	}).(DeviceNetworkTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceNetworkTypeOutput{})
	pulumi.RegisterOutputType(DeviceNetworkTypePtrOutput{})
	pulumi.RegisterOutputType(DeviceNetworkTypeArrayOutput{})
	pulumi.RegisterOutputType(DeviceNetworkTypeMapOutput{})
}
