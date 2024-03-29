// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	equinix-metal "github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := equinix-metal.NewVlan(ctx, "testVlan", &equinix-metal.VlanArgs{
// Description: pulumi.String("test VLAN in SV"),
// Metro: pulumi.String("sv"),
// ProjectId: pulumi.Any(local.Project_id),
// })
// if err != nil {
// return err
// }
// _, err = equinix-metal.LookupGateway(ctx, &equinix.LookupGatewayArgs{
// GatewayId: local.Gateway_id,
// }, nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayResult
	err := ctx.Invoke("equinix-metal:index/getGateway:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGateway.
type LookupGatewayArgs struct {
	// UUID of the metal gateway resource to retrieve
	GatewayId string `pulumi:"gatewayId"`
}

// A collection of values returned by getGateway.
type LookupGatewayResult struct {
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// UUID of IP reservation block bound to the gateway
	IpReservationId string `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet bound to this metal gateway, one of (8, 16, 32, 64, 128)`
	PrivateIpv4SubnetSize int `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to
	ProjectId string `pulumi:"projectId"`
	// Status of the gateway resource
	State string `pulumi:"state"`
	// UUID of the VLAN where the gateway is scoped to
	VlanId string `pulumi:"vlanId"`
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

// A collection of arguments for invoking getGateway.
type LookupGatewayOutputArgs struct {
	// UUID of the metal gateway resource to retrieve
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getGateway.
type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGatewayResult] {
	return pulumix.Output[LookupGatewayResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupGatewayResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// UUID of IP reservation block bound to the gateway
func (o LookupGatewayResultOutput) IpReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.IpReservationId }).(pulumi.StringOutput)
}

// Size of the private IPv4 subnet bound to this metal gateway, one of (8, 16, 32, 64, 128)`
func (o LookupGatewayResultOutput) PrivateIpv4SubnetSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGatewayResult) int { return v.PrivateIpv4SubnetSize }).(pulumi.IntOutput)
}

// UUID of the project where the gateway is scoped to
func (o LookupGatewayResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Status of the gateway resource
func (o LookupGatewayResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.State }).(pulumi.StringOutput)
}

// UUID of the VLAN where the gateway is scoped to
func (o LookupGatewayResultOutput) VlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.VlanId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}
