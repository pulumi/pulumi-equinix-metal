// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage organization resource in Equinix Metal.
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
// 		_, err := equinix - metal.NewOrganization(ctx, "tfOrganization1", &equinix-metal.OrganizationArgs{
// 			Description: pulumi.String("quux"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// This resource can be imported using an existing organization ID
//
// ```sh
//  $ pulumi import equinix-metal:index/organization:Organization metal_organization {existing_organization_id}
// ```
type Organization struct {
	pulumi.CustomResourceState

	Created pulumi.StringOutput `pulumi:"created"`
	// Description string
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Logo URL
	Logo pulumi.StringPtrOutput `pulumi:"logo"`
	// The name of the Organization
	Name pulumi.StringOutput `pulumi:"name"`
	// Twitter handle
	Twitter pulumi.StringPtrOutput `pulumi:"twitter"`
	Updated pulumi.StringOutput    `pulumi:"updated"`
	// Website link
	Website pulumi.StringPtrOutput `pulumi:"website"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		args = &OrganizationArgs{}
	}

	var resource Organization
	err := ctx.RegisterResource("equinix-metal:index/organization:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("equinix-metal:index/organization:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
	Created *string `pulumi:"created"`
	// Description string
	Description *string `pulumi:"description"`
	// Logo URL
	Logo *string `pulumi:"logo"`
	// The name of the Organization
	Name *string `pulumi:"name"`
	// Twitter handle
	Twitter *string `pulumi:"twitter"`
	Updated *string `pulumi:"updated"`
	// Website link
	Website *string `pulumi:"website"`
}

type OrganizationState struct {
	Created pulumi.StringPtrInput
	// Description string
	Description pulumi.StringPtrInput
	// Logo URL
	Logo pulumi.StringPtrInput
	// The name of the Organization
	Name pulumi.StringPtrInput
	// Twitter handle
	Twitter pulumi.StringPtrInput
	Updated pulumi.StringPtrInput
	// Website link
	Website pulumi.StringPtrInput
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// Description string
	Description *string `pulumi:"description"`
	// Logo URL
	Logo *string `pulumi:"logo"`
	// The name of the Organization
	Name *string `pulumi:"name"`
	// Twitter handle
	Twitter *string `pulumi:"twitter"`
	// Website link
	Website *string `pulumi:"website"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// Description string
	Description pulumi.StringPtrInput
	// Logo URL
	Logo pulumi.StringPtrInput
	// The name of the Organization
	Name pulumi.StringPtrInput
	// Twitter handle
	Twitter pulumi.StringPtrInput
	// Website link
	Website pulumi.StringPtrInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

func (*Organization) ElementType() reflect.Type {
	return reflect.TypeOf((*Organization)(nil))
}

func (i *Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i *Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

func (i *Organization) ToOrganizationPtrOutput() OrganizationPtrOutput {
	return i.ToOrganizationPtrOutputWithContext(context.Background())
}

func (i *Organization) ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationPtrOutput)
}

type OrganizationPtrInput interface {
	pulumi.Input

	ToOrganizationPtrOutput() OrganizationPtrOutput
	ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput
}

type organizationPtrType OrganizationArgs

func (*organizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil))
}

func (i *organizationPtrType) ToOrganizationPtrOutput() OrganizationPtrOutput {
	return i.ToOrganizationPtrOutputWithContext(context.Background())
}

func (i *organizationPtrType) ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationPtrOutput)
}

// OrganizationArrayInput is an input type that accepts OrganizationArray and OrganizationArrayOutput values.
// You can construct a concrete instance of `OrganizationArrayInput` via:
//
//          OrganizationArray{ OrganizationArgs{...} }
type OrganizationArrayInput interface {
	pulumi.Input

	ToOrganizationArrayOutput() OrganizationArrayOutput
	ToOrganizationArrayOutputWithContext(context.Context) OrganizationArrayOutput
}

type OrganizationArray []OrganizationInput

func (OrganizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Organization)(nil)).Elem()
}

func (i OrganizationArray) ToOrganizationArrayOutput() OrganizationArrayOutput {
	return i.ToOrganizationArrayOutputWithContext(context.Background())
}

func (i OrganizationArray) ToOrganizationArrayOutputWithContext(ctx context.Context) OrganizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationArrayOutput)
}

// OrganizationMapInput is an input type that accepts OrganizationMap and OrganizationMapOutput values.
// You can construct a concrete instance of `OrganizationMapInput` via:
//
//          OrganizationMap{ "key": OrganizationArgs{...} }
type OrganizationMapInput interface {
	pulumi.Input

	ToOrganizationMapOutput() OrganizationMapOutput
	ToOrganizationMapOutputWithContext(context.Context) OrganizationMapOutput
}

type OrganizationMap map[string]OrganizationInput

func (OrganizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Organization)(nil)).Elem()
}

func (i OrganizationMap) ToOrganizationMapOutput() OrganizationMapOutput {
	return i.ToOrganizationMapOutputWithContext(context.Background())
}

func (i OrganizationMap) ToOrganizationMapOutputWithContext(ctx context.Context) OrganizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationMapOutput)
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Organization)(nil))
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationPtrOutput() OrganizationPtrOutput {
	return o.ToOrganizationPtrOutputWithContext(context.Background())
}

func (o OrganizationOutput) ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Organization) *Organization {
		return &v
	}).(OrganizationPtrOutput)
}

type OrganizationPtrOutput struct{ *pulumi.OutputState }

func (OrganizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil))
}

func (o OrganizationPtrOutput) ToOrganizationPtrOutput() OrganizationPtrOutput {
	return o
}

func (o OrganizationPtrOutput) ToOrganizationPtrOutputWithContext(ctx context.Context) OrganizationPtrOutput {
	return o
}

func (o OrganizationPtrOutput) Elem() OrganizationOutput {
	return o.ApplyT(func(v *Organization) Organization {
		if v != nil {
			return *v
		}
		var ret Organization
		return ret
	}).(OrganizationOutput)
}

type OrganizationArrayOutput struct{ *pulumi.OutputState }

func (OrganizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Organization)(nil))
}

func (o OrganizationArrayOutput) ToOrganizationArrayOutput() OrganizationArrayOutput {
	return o
}

func (o OrganizationArrayOutput) ToOrganizationArrayOutputWithContext(ctx context.Context) OrganizationArrayOutput {
	return o
}

func (o OrganizationArrayOutput) Index(i pulumi.IntInput) OrganizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Organization {
		return vs[0].([]Organization)[vs[1].(int)]
	}).(OrganizationOutput)
}

type OrganizationMapOutput struct{ *pulumi.OutputState }

func (OrganizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Organization)(nil))
}

func (o OrganizationMapOutput) ToOrganizationMapOutput() OrganizationMapOutput {
	return o
}

func (o OrganizationMapOutput) ToOrganizationMapOutputWithContext(ctx context.Context) OrganizationMapOutput {
	return o
}

func (o OrganizationMapOutput) MapIndex(k pulumi.StringInput) OrganizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Organization {
		return vs[0].(map[string]Organization)[vs[1].(string)]
	}).(OrganizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationInput)(nil)).Elem(), &Organization{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationPtrInput)(nil)).Elem(), &Organization{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationArrayInput)(nil)).Elem(), OrganizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationMapInput)(nil)).Elem(), OrganizationMap{})
	pulumi.RegisterOutputType(OrganizationOutput{})
	pulumi.RegisterOutputType(OrganizationPtrOutput{})
	pulumi.RegisterOutputType(OrganizationArrayOutput{})
	pulumi.RegisterOutputType(OrganizationMapOutput{})
}
