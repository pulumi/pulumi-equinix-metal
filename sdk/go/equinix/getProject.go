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

// Use this datasource to retrieve attributes of the Project API resource.
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
// tfProject1, err := equinix-metal.LookupProject(ctx, &equinix.LookupProjectArgs{
// Name: pulumi.StringRef("Terraform Fun"),
// }, nil);
// if err != nil {
// return err
// }
// ctx.Export("usersOfTerraformFun", tfProject1.UserIds)
// return nil
// })
// }
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("equinix-metal:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// The name which is used to look up the project
	Name *string `pulumi:"name"`
	// The UUID by which to look up the project
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// Whether Backend Transfer is enabled for this project
	BackendTransfer bool `pulumi:"backendTransfer"`
	// Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
	BgpConfigs []GetProjectBgpConfig `pulumi:"bgpConfigs"`
	// The timestamp for when the project was created
	Created string `pulumi:"created"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The UUID of this project's parent organization
	OrganizationId string `pulumi:"organizationId"`
	// The UUID of payment method for this project
	PaymentMethodId string `pulumi:"paymentMethodId"`
	ProjectId       string `pulumi:"projectId"`
	// The timestamp for the last time the project was updated
	Updated string `pulumi:"updated"`
	// List of UUIDs of user accounts which belong to this project
	UserIds []string `pulumi:"userIds"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// The name which is used to look up the project
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The UUID by which to look up the project
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupProjectResult] {
	return pulumix.Output[LookupProjectResult]{
		OutputState: o.OutputState,
	}
}

// Whether Backend Transfer is enabled for this project
func (o LookupProjectResultOutput) BackendTransfer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.BackendTransfer }).(pulumi.BoolOutput)
}

// Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
func (o LookupProjectResultOutput) BgpConfigs() GetProjectBgpConfigArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectBgpConfig { return v.BgpConfigs }).(GetProjectBgpConfigArrayOutput)
}

// The timestamp for when the project was created
func (o LookupProjectResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Created }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The UUID of this project's parent organization
func (o LookupProjectResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The UUID of payment method for this project
func (o LookupProjectResultOutput) PaymentMethodId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.PaymentMethodId }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The timestamp for the last time the project was updated
func (o LookupProjectResultOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Updated }).(pulumi.StringOutput)
}

// List of UUIDs of user accounts which belong to this project
func (o LookupProjectResultOutput) UserIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []string { return v.UserIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
