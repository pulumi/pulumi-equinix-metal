// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create Metal Gateway resources in Equinix Metal.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix-metal"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testVlan, err := equinix - metal.NewVlan(ctx, "testVlan", &equinix-metal.VlanArgs{
// 			Description: pulumi.String("test VLAN in SV"),
// 			Metro:       pulumi.String("sv"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.NewGateway(ctx, "testGateway", &equinix-metal.GatewayArgs{
// 			ProjectId:             pulumi.Any(local.Project_id),
// 			VlanId:                testVlan.ID(),
// 			PrivateIpv4SubnetSize: pulumi.Int(8),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix-metal"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testVlan, err := equinix - metal.NewVlan(ctx, "testVlan", &equinix-metal.VlanArgs{
// 			Description: pulumi.String("test VLAN in SV"),
// 			Metro:       pulumi.String("sv"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testReservedIpBlock, err := equinix - metal.NewReservedIpBlock(ctx, "testReservedIpBlock", &equinix-metal.ReservedIpBlockArgs{
// 			ProjectId: pulumi.Any(local.Project_id),
// 			Metro:     pulumi.String("sv"),
// 			Quantity:  pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.NewGateway(ctx, "testGateway", &equinix-metal.GatewayArgs{
// 			ProjectId:       pulumi.Any(local.Project_id),
// 			VlanId:          testVlan.ID(),
// 			IpReservationId: testReservedIpBlock.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
	IpReservationId pulumi.StringPtrOutput `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
	PrivateIpv4SubnetSize pulumi.IntPtrOutput `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Status of the gateway resource
	State pulumi.StringOutput `pulumi:"state"`
	// UUID of the VLAN where the gateway is scoped to
	VlanId pulumi.StringOutput `pulumi:"vlanId"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	var resource Gateway
	err := ctx.RegisterResource("equinix-metal:index/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("equinix-metal:index/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
	IpReservationId *string `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
	PrivateIpv4SubnetSize *int `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to
	ProjectId *string `pulumi:"projectId"`
	// Status of the gateway resource
	State *string `pulumi:"state"`
	// UUID of the VLAN where the gateway is scoped to
	VlanId *string `pulumi:"vlanId"`
}

type GatewayState struct {
	// UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
	IpReservationId pulumi.StringPtrInput
	// Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
	PrivateIpv4SubnetSize pulumi.IntPtrInput
	// UUID of the project where the gateway is scoped to
	ProjectId pulumi.StringPtrInput
	// Status of the gateway resource
	State pulumi.StringPtrInput
	// UUID of the VLAN where the gateway is scoped to
	VlanId pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
	IpReservationId *string `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
	PrivateIpv4SubnetSize *int `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to
	ProjectId string `pulumi:"projectId"`
	// UUID of the VLAN where the gateway is scoped to
	VlanId string `pulumi:"vlanId"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
	IpReservationId pulumi.StringPtrInput
	// Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
	PrivateIpv4SubnetSize pulumi.IntPtrInput
	// UUID of the project where the gateway is scoped to
	ProjectId pulumi.StringInput
	// UUID of the VLAN where the gateway is scoped to
	VlanId pulumi.StringInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil))
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

func (i *Gateway) ToGatewayPtrOutput() GatewayPtrOutput {
	return i.ToGatewayPtrOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPtrOutput)
}

type GatewayPtrInput interface {
	pulumi.Input

	ToGatewayPtrOutput() GatewayPtrOutput
	ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput
}

type gatewayPtrType GatewayArgs

func (*gatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil))
}

func (i *gatewayPtrType) ToGatewayPtrOutput() GatewayPtrOutput {
	return i.ToGatewayPtrOutputWithContext(context.Background())
}

func (i *gatewayPtrType) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPtrOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//          GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//          GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil))
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayPtrOutput() GatewayPtrOutput {
	return o.ToGatewayPtrOutputWithContext(context.Background())
}

func (o GatewayOutput) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Gateway) *Gateway {
		return &v
	}).(GatewayPtrOutput)
}

type GatewayPtrOutput struct{ *pulumi.OutputState }

func (GatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil))
}

func (o GatewayPtrOutput) ToGatewayPtrOutput() GatewayPtrOutput {
	return o
}

func (o GatewayPtrOutput) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return o
}

func (o GatewayPtrOutput) Elem() GatewayOutput {
	return o.ApplyT(func(v *Gateway) Gateway {
		if v != nil {
			return *v
		}
		var ret Gateway
		return ret
	}).(GatewayOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Gateway)(nil))
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Gateway {
		return vs[0].([]Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Gateway)(nil))
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Gateway {
		return vs[0].(map[string]Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayPtrInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayArrayInput)(nil)).Elem(), GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayMapInput)(nil)).Elem(), GatewayMap{})
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayPtrOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
