// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to allow users to manage Virtual Networks in their projects.
//
// To learn more about Layer 2 networking in Equinix Metal, refer to
//
// * <https://metal.equinix.com/developers/docs/networking/layer2/>
// * <https://metal.equinix.com/developers/docs/networking/layer2-configs/>
//
// ## Example Usage
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
// 		_, err := equinix - metal.NewVlan(ctx, "vlan1Vlan", &equinix-metal.VlanArgs{
// 			Description: pulumi.String("VLAN in New Jersey"),
// 			Facility:    pulumi.String("sv15"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.NewVlan(ctx, "vlan1Index_vlanVlan", &equinix-metal.VlanArgs{
// 			Description: pulumi.String("VLAN in New Jersey"),
// 			Metro:       pulumi.String("sv"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 			Vxlan:       pulumi.Int(1040),
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
// This resource can be imported using an existing VLAN ID (UUID)
//
// ```sh
//  $ pulumi import equinix-metal:index/vlan:Vlan metal_vlan {existing_vlan_id}
// ```
type Vlan struct {
	pulumi.CustomResourceState

	// Description string
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Facility where to create the VLAN
	Facility pulumi.StringPtrOutput `pulumi:"facility"`
	Metro    pulumi.StringPtrOutput `pulumi:"metro"`
	// ID of parent project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// VLAN ID, must be unique in metro
	Vxlan pulumi.IntOutput `pulumi:"vxlan"`
}

// NewVlan registers a new resource with the given unique name, arguments, and options.
func NewVlan(ctx *pulumi.Context,
	name string, args *VlanArgs, opts ...pulumi.ResourceOption) (*Vlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Vlan
	err := ctx.RegisterResource("equinix-metal:index/vlan:Vlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVlan gets an existing Vlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VlanState, opts ...pulumi.ResourceOption) (*Vlan, error) {
	var resource Vlan
	err := ctx.ReadResource("equinix-metal:index/vlan:Vlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vlan resources.
type vlanState struct {
	// Description string
	Description *string `pulumi:"description"`
	// Facility where to create the VLAN
	Facility *string `pulumi:"facility"`
	Metro    *string `pulumi:"metro"`
	// ID of parent project
	ProjectId *string `pulumi:"projectId"`
	// VLAN ID, must be unique in metro
	Vxlan *int `pulumi:"vxlan"`
}

type VlanState struct {
	// Description string
	Description pulumi.StringPtrInput
	// Facility where to create the VLAN
	Facility pulumi.StringPtrInput
	Metro    pulumi.StringPtrInput
	// ID of parent project
	ProjectId pulumi.StringPtrInput
	// VLAN ID, must be unique in metro
	Vxlan pulumi.IntPtrInput
}

func (VlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanState)(nil)).Elem()
}

type vlanArgs struct {
	// Description string
	Description *string `pulumi:"description"`
	// Facility where to create the VLAN
	Facility *string `pulumi:"facility"`
	Metro    *string `pulumi:"metro"`
	// ID of parent project
	ProjectId string `pulumi:"projectId"`
	// VLAN ID, must be unique in metro
	Vxlan *int `pulumi:"vxlan"`
}

// The set of arguments for constructing a Vlan resource.
type VlanArgs struct {
	// Description string
	Description pulumi.StringPtrInput
	// Facility where to create the VLAN
	Facility pulumi.StringPtrInput
	Metro    pulumi.StringPtrInput
	// ID of parent project
	ProjectId pulumi.StringInput
	// VLAN ID, must be unique in metro
	Vxlan pulumi.IntPtrInput
}

func (VlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vlanArgs)(nil)).Elem()
}

type VlanInput interface {
	pulumi.Input

	ToVlanOutput() VlanOutput
	ToVlanOutputWithContext(ctx context.Context) VlanOutput
}

func (*Vlan) ElementType() reflect.Type {
	return reflect.TypeOf((*Vlan)(nil))
}

func (i *Vlan) ToVlanOutput() VlanOutput {
	return i.ToVlanOutputWithContext(context.Background())
}

func (i *Vlan) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanOutput)
}

func (i *Vlan) ToVlanPtrOutput() VlanPtrOutput {
	return i.ToVlanPtrOutputWithContext(context.Background())
}

func (i *Vlan) ToVlanPtrOutputWithContext(ctx context.Context) VlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanPtrOutput)
}

type VlanPtrInput interface {
	pulumi.Input

	ToVlanPtrOutput() VlanPtrOutput
	ToVlanPtrOutputWithContext(ctx context.Context) VlanPtrOutput
}

type vlanPtrType VlanArgs

func (*vlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlan)(nil))
}

func (i *vlanPtrType) ToVlanPtrOutput() VlanPtrOutput {
	return i.ToVlanPtrOutputWithContext(context.Background())
}

func (i *vlanPtrType) ToVlanPtrOutputWithContext(ctx context.Context) VlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanPtrOutput)
}

// VlanArrayInput is an input type that accepts VlanArray and VlanArrayOutput values.
// You can construct a concrete instance of `VlanArrayInput` via:
//
//          VlanArray{ VlanArgs{...} }
type VlanArrayInput interface {
	pulumi.Input

	ToVlanArrayOutput() VlanArrayOutput
	ToVlanArrayOutputWithContext(context.Context) VlanArrayOutput
}

type VlanArray []VlanInput

func (VlanArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Vlan)(nil))
}

func (i VlanArray) ToVlanArrayOutput() VlanArrayOutput {
	return i.ToVlanArrayOutputWithContext(context.Background())
}

func (i VlanArray) ToVlanArrayOutputWithContext(ctx context.Context) VlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanArrayOutput)
}

// VlanMapInput is an input type that accepts VlanMap and VlanMapOutput values.
// You can construct a concrete instance of `VlanMapInput` via:
//
//          VlanMap{ "key": VlanArgs{...} }
type VlanMapInput interface {
	pulumi.Input

	ToVlanMapOutput() VlanMapOutput
	ToVlanMapOutputWithContext(context.Context) VlanMapOutput
}

type VlanMap map[string]VlanInput

func (VlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Vlan)(nil))
}

func (i VlanMap) ToVlanMapOutput() VlanMapOutput {
	return i.ToVlanMapOutputWithContext(context.Background())
}

func (i VlanMap) ToVlanMapOutputWithContext(ctx context.Context) VlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanMapOutput)
}

type VlanOutput struct {
	*pulumi.OutputState
}

func (VlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vlan)(nil))
}

func (o VlanOutput) ToVlanOutput() VlanOutput {
	return o
}

func (o VlanOutput) ToVlanOutputWithContext(ctx context.Context) VlanOutput {
	return o
}

func (o VlanOutput) ToVlanPtrOutput() VlanPtrOutput {
	return o.ToVlanPtrOutputWithContext(context.Background())
}

func (o VlanOutput) ToVlanPtrOutputWithContext(ctx context.Context) VlanPtrOutput {
	return o.ApplyT(func(v Vlan) *Vlan {
		return &v
	}).(VlanPtrOutput)
}

type VlanPtrOutput struct {
	*pulumi.OutputState
}

func (VlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vlan)(nil))
}

func (o VlanPtrOutput) ToVlanPtrOutput() VlanPtrOutput {
	return o
}

func (o VlanPtrOutput) ToVlanPtrOutputWithContext(ctx context.Context) VlanPtrOutput {
	return o
}

type VlanArrayOutput struct{ *pulumi.OutputState }

func (VlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Vlan)(nil))
}

func (o VlanArrayOutput) ToVlanArrayOutput() VlanArrayOutput {
	return o
}

func (o VlanArrayOutput) ToVlanArrayOutputWithContext(ctx context.Context) VlanArrayOutput {
	return o
}

func (o VlanArrayOutput) Index(i pulumi.IntInput) VlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Vlan {
		return vs[0].([]Vlan)[vs[1].(int)]
	}).(VlanOutput)
}

type VlanMapOutput struct{ *pulumi.OutputState }

func (VlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Vlan)(nil))
}

func (o VlanMapOutput) ToVlanMapOutput() VlanMapOutput {
	return o
}

func (o VlanMapOutput) ToVlanMapOutputWithContext(ctx context.Context) VlanMapOutput {
	return o
}

func (o VlanMapOutput) MapIndex(k pulumi.StringInput) VlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Vlan {
		return vs[0].(map[string]Vlan)[vs[1].(string)]
	}).(VlanOutput)
}

func init() {
	pulumi.RegisterOutputType(VlanOutput{})
	pulumi.RegisterOutputType(VlanPtrOutput{})
	pulumi.RegisterOutputType(VlanArrayOutput{})
	pulumi.RegisterOutputType(VlanMapOutput{})
}
