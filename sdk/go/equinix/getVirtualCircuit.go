// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a virtual circuit resource from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix-metal"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleConnection, err := equinix - metal.LookupConnection(ctx, &GetConnectionArgs{
// 			ConnectionId: "4347e805-eb46-4699-9eb9-5c116e6a017d",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.LookupVirtualCircuit(ctx, &GetVirtualCircuitArgs{
// 			VirtualCircuitId: exampleConnection.Ports[1].VirtualCircuitIds[0],
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVirtualCircuit(ctx *pulumi.Context, args *LookupVirtualCircuitArgs, opts ...pulumi.InvokeOption) (*LookupVirtualCircuitResult, error) {
	var rv LookupVirtualCircuitResult
	err := ctx.Invoke("equinix-metal:index/getVirtualCircuit:getVirtualCircuit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualCircuit.
type LookupVirtualCircuitArgs struct {
	// ID of the virtual circuit resource
	VirtualCircuitId string `pulumi:"virtualCircuitId"`
}

// A collection of values returned by getVirtualCircuit.
type LookupVirtualCircuitResult struct {
	// Description for the Virtual Circuit resource
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the virtual circuit resource
	Name    string `pulumi:"name"`
	NniVlan int    `pulumi:"nniVlan"`
	NniVnid int    `pulumi:"nniVnid"`
	// ID of project to which the VC belongs
	// * `vnid`, `nniVlan`, `nniNvid` - VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)
	ProjectId string `pulumi:"projectId"`
	// Status of the virtal circuit
	Status string `pulumi:"status"`
	// Tags for the Virtual Circuit resource
	Tags             []string `pulumi:"tags"`
	VirtualCircuitId string   `pulumi:"virtualCircuitId"`
	Vnid             int      `pulumi:"vnid"`
}

func LookupVirtualCircuitOutput(ctx *pulumi.Context, args LookupVirtualCircuitOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualCircuitResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualCircuitResult, error) {
			args := v.(LookupVirtualCircuitArgs)
			r, err := LookupVirtualCircuit(ctx, &args, opts...)
			var s LookupVirtualCircuitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualCircuitResultOutput)
}

// A collection of arguments for invoking getVirtualCircuit.
type LookupVirtualCircuitOutputArgs struct {
	// ID of the virtual circuit resource
	VirtualCircuitId pulumi.StringInput `pulumi:"virtualCircuitId"`
}

func (LookupVirtualCircuitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualCircuitArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualCircuit.
type LookupVirtualCircuitResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualCircuitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualCircuitResult)(nil)).Elem()
}

func (o LookupVirtualCircuitResultOutput) ToLookupVirtualCircuitResultOutput() LookupVirtualCircuitResultOutput {
	return o
}

func (o LookupVirtualCircuitResultOutput) ToLookupVirtualCircuitResultOutputWithContext(ctx context.Context) LookupVirtualCircuitResultOutput {
	return o
}

// Description for the Virtual Circuit resource
func (o LookupVirtualCircuitResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualCircuitResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the virtual circuit resource
func (o LookupVirtualCircuitResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualCircuitResultOutput) NniVlan() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) int { return v.NniVlan }).(pulumi.IntOutput)
}

func (o LookupVirtualCircuitResultOutput) NniVnid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) int { return v.NniVnid }).(pulumi.IntOutput)
}

// ID of project to which the VC belongs
// * `vnid`, `nniVlan`, `nniNvid` - VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)
func (o LookupVirtualCircuitResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Status of the virtal circuit
func (o LookupVirtualCircuitResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) string { return v.Status }).(pulumi.StringOutput)
}

// Tags for the Virtual Circuit resource
func (o LookupVirtualCircuitResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualCircuitResultOutput) VirtualCircuitId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) string { return v.VirtualCircuitId }).(pulumi.StringOutput)
}

func (o LookupVirtualCircuitResultOutput) Vnid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualCircuitResult) int { return v.Vnid }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualCircuitResultOutput{})
}
