// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
// 		_, err := equinix - metal.NewProjectApiKey(ctx, "test", &equinix-metal.ProjectApiKeyArgs{
// 			ProjectId:   pulumi.Any(local.Existing_project_id),
// 			Description: pulumi.String("Read-only key scoped to a projct"),
// 			ReadOnly:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProjectApiKey struct {
	pulumi.CustomResourceState

	// Description string for the Project API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description pulumi.StringOutput `pulumi:"description"`
	// UUID of the project where the API key is scoped to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewProjectApiKey registers a new resource with the given unique name, arguments, and options.
func NewProjectApiKey(ctx *pulumi.Context,
	name string, args *ProjectApiKeyArgs, opts ...pulumi.ResourceOption) (*ProjectApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ReadOnly == nil {
		return nil, errors.New("invalid value for required argument 'ReadOnly'")
	}
	var resource ProjectApiKey
	err := ctx.RegisterResource("equinix-metal:index/projectApiKey:ProjectApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectApiKey gets an existing ProjectApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectApiKeyState, opts ...pulumi.ResourceOption) (*ProjectApiKey, error) {
	var resource ProjectApiKey
	err := ctx.ReadResource("equinix-metal:index/projectApiKey:ProjectApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectApiKey resources.
type projectApiKeyState struct {
	// Description string for the Project API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description *string `pulumi:"description"`
	// UUID of the project where the API key is scoped to
	ProjectId *string `pulumi:"projectId"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly *bool `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients
	Token *string `pulumi:"token"`
}

type ProjectApiKeyState struct {
	// Description string for the Project API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description pulumi.StringPtrInput
	// UUID of the project where the API key is scoped to
	ProjectId pulumi.StringPtrInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolPtrInput
	// API token which can be used in Equinix Metal API clients
	Token pulumi.StringPtrInput
}

func (ProjectApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectApiKeyState)(nil)).Elem()
}

type projectApiKeyArgs struct {
	// Description string for the Project API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description string `pulumi:"description"`
	// UUID of the project where the API key is scoped to
	ProjectId string `pulumi:"projectId"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly bool `pulumi:"readOnly"`
}

// The set of arguments for constructing a ProjectApiKey resource.
type ProjectApiKeyArgs struct {
	// Description string for the Project API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description pulumi.StringInput
	// UUID of the project where the API key is scoped to
	ProjectId pulumi.StringInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolInput
}

func (ProjectApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectApiKeyArgs)(nil)).Elem()
}

type ProjectApiKeyInput interface {
	pulumi.Input

	ToProjectApiKeyOutput() ProjectApiKeyOutput
	ToProjectApiKeyOutputWithContext(ctx context.Context) ProjectApiKeyOutput
}

func (*ProjectApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectApiKey)(nil))
}

func (i *ProjectApiKey) ToProjectApiKeyOutput() ProjectApiKeyOutput {
	return i.ToProjectApiKeyOutputWithContext(context.Background())
}

func (i *ProjectApiKey) ToProjectApiKeyOutputWithContext(ctx context.Context) ProjectApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApiKeyOutput)
}

func (i *ProjectApiKey) ToProjectApiKeyPtrOutput() ProjectApiKeyPtrOutput {
	return i.ToProjectApiKeyPtrOutputWithContext(context.Background())
}

func (i *ProjectApiKey) ToProjectApiKeyPtrOutputWithContext(ctx context.Context) ProjectApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApiKeyPtrOutput)
}

type ProjectApiKeyPtrInput interface {
	pulumi.Input

	ToProjectApiKeyPtrOutput() ProjectApiKeyPtrOutput
	ToProjectApiKeyPtrOutputWithContext(ctx context.Context) ProjectApiKeyPtrOutput
}

type projectApiKeyPtrType ProjectApiKeyArgs

func (*projectApiKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectApiKey)(nil))
}

func (i *projectApiKeyPtrType) ToProjectApiKeyPtrOutput() ProjectApiKeyPtrOutput {
	return i.ToProjectApiKeyPtrOutputWithContext(context.Background())
}

func (i *projectApiKeyPtrType) ToProjectApiKeyPtrOutputWithContext(ctx context.Context) ProjectApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApiKeyPtrOutput)
}

// ProjectApiKeyArrayInput is an input type that accepts ProjectApiKeyArray and ProjectApiKeyArrayOutput values.
// You can construct a concrete instance of `ProjectApiKeyArrayInput` via:
//
//          ProjectApiKeyArray{ ProjectApiKeyArgs{...} }
type ProjectApiKeyArrayInput interface {
	pulumi.Input

	ToProjectApiKeyArrayOutput() ProjectApiKeyArrayOutput
	ToProjectApiKeyArrayOutputWithContext(context.Context) ProjectApiKeyArrayOutput
}

type ProjectApiKeyArray []ProjectApiKeyInput

func (ProjectApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectApiKey)(nil)).Elem()
}

func (i ProjectApiKeyArray) ToProjectApiKeyArrayOutput() ProjectApiKeyArrayOutput {
	return i.ToProjectApiKeyArrayOutputWithContext(context.Background())
}

func (i ProjectApiKeyArray) ToProjectApiKeyArrayOutputWithContext(ctx context.Context) ProjectApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApiKeyArrayOutput)
}

// ProjectApiKeyMapInput is an input type that accepts ProjectApiKeyMap and ProjectApiKeyMapOutput values.
// You can construct a concrete instance of `ProjectApiKeyMapInput` via:
//
//          ProjectApiKeyMap{ "key": ProjectApiKeyArgs{...} }
type ProjectApiKeyMapInput interface {
	pulumi.Input

	ToProjectApiKeyMapOutput() ProjectApiKeyMapOutput
	ToProjectApiKeyMapOutputWithContext(context.Context) ProjectApiKeyMapOutput
}

type ProjectApiKeyMap map[string]ProjectApiKeyInput

func (ProjectApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectApiKey)(nil)).Elem()
}

func (i ProjectApiKeyMap) ToProjectApiKeyMapOutput() ProjectApiKeyMapOutput {
	return i.ToProjectApiKeyMapOutputWithContext(context.Background())
}

func (i ProjectApiKeyMap) ToProjectApiKeyMapOutputWithContext(ctx context.Context) ProjectApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectApiKeyMapOutput)
}

type ProjectApiKeyOutput struct{ *pulumi.OutputState }

func (ProjectApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectApiKey)(nil))
}

func (o ProjectApiKeyOutput) ToProjectApiKeyOutput() ProjectApiKeyOutput {
	return o
}

func (o ProjectApiKeyOutput) ToProjectApiKeyOutputWithContext(ctx context.Context) ProjectApiKeyOutput {
	return o
}

func (o ProjectApiKeyOutput) ToProjectApiKeyPtrOutput() ProjectApiKeyPtrOutput {
	return o.ToProjectApiKeyPtrOutputWithContext(context.Background())
}

func (o ProjectApiKeyOutput) ToProjectApiKeyPtrOutputWithContext(ctx context.Context) ProjectApiKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectApiKey) *ProjectApiKey {
		return &v
	}).(ProjectApiKeyPtrOutput)
}

type ProjectApiKeyPtrOutput struct{ *pulumi.OutputState }

func (ProjectApiKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectApiKey)(nil))
}

func (o ProjectApiKeyPtrOutput) ToProjectApiKeyPtrOutput() ProjectApiKeyPtrOutput {
	return o
}

func (o ProjectApiKeyPtrOutput) ToProjectApiKeyPtrOutputWithContext(ctx context.Context) ProjectApiKeyPtrOutput {
	return o
}

func (o ProjectApiKeyPtrOutput) Elem() ProjectApiKeyOutput {
	return o.ApplyT(func(v *ProjectApiKey) ProjectApiKey {
		if v != nil {
			return *v
		}
		var ret ProjectApiKey
		return ret
	}).(ProjectApiKeyOutput)
}

type ProjectApiKeyArrayOutput struct{ *pulumi.OutputState }

func (ProjectApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectApiKey)(nil))
}

func (o ProjectApiKeyArrayOutput) ToProjectApiKeyArrayOutput() ProjectApiKeyArrayOutput {
	return o
}

func (o ProjectApiKeyArrayOutput) ToProjectApiKeyArrayOutputWithContext(ctx context.Context) ProjectApiKeyArrayOutput {
	return o
}

func (o ProjectApiKeyArrayOutput) Index(i pulumi.IntInput) ProjectApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProjectApiKey {
		return vs[0].([]ProjectApiKey)[vs[1].(int)]
	}).(ProjectApiKeyOutput)
}

type ProjectApiKeyMapOutput struct{ *pulumi.OutputState }

func (ProjectApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProjectApiKey)(nil))
}

func (o ProjectApiKeyMapOutput) ToProjectApiKeyMapOutput() ProjectApiKeyMapOutput {
	return o
}

func (o ProjectApiKeyMapOutput) ToProjectApiKeyMapOutputWithContext(ctx context.Context) ProjectApiKeyMapOutput {
	return o
}

func (o ProjectApiKeyMapOutput) MapIndex(k pulumi.StringInput) ProjectApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProjectApiKey {
		return vs[0].(map[string]ProjectApiKey)[vs[1].(string)]
	}).(ProjectApiKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApiKeyInput)(nil)).Elem(), &ProjectApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApiKeyPtrInput)(nil)).Elem(), &ProjectApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApiKeyArrayInput)(nil)).Elem(), ProjectApiKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectApiKeyMapInput)(nil)).Elem(), ProjectApiKeyMap{})
	pulumi.RegisterOutputType(ProjectApiKeyOutput{})
	pulumi.RegisterOutputType(ProjectApiKeyPtrOutput{})
	pulumi.RegisterOutputType(ProjectApiKeyArrayOutput{})
	pulumi.RegisterOutputType(ProjectApiKeyMapOutput{})
}
