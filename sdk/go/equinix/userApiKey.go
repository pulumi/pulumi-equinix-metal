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
// 		_, err := equinix - metal.NewUserApiKey(ctx, "test", &equinix-metal.UserApiKeyArgs{
// 			Description: pulumi.String("Read-only user key"),
// 			ReadOnly:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type UserApiKey struct {
	pulumi.CustomResourceState

	// Description string for the User API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description pulumi.StringOutput `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients
	Token pulumi.StringOutput `pulumi:"token"`
	// UUID of the owner of the API key
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserApiKey registers a new resource with the given unique name, arguments, and options.
func NewUserApiKey(ctx *pulumi.Context,
	name string, args *UserApiKeyArgs, opts ...pulumi.ResourceOption) (*UserApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ReadOnly == nil {
		return nil, errors.New("invalid value for required argument 'ReadOnly'")
	}
	var resource UserApiKey
	err := ctx.RegisterResource("equinix-metal:index/userApiKey:UserApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserApiKey gets an existing UserApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserApiKeyState, opts ...pulumi.ResourceOption) (*UserApiKey, error) {
	var resource UserApiKey
	err := ctx.ReadResource("equinix-metal:index/userApiKey:UserApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserApiKey resources.
type userApiKeyState struct {
	// Description string for the User API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description *string `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly *bool `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients
	Token *string `pulumi:"token"`
	// UUID of the owner of the API key
	UserId *string `pulumi:"userId"`
}

type UserApiKeyState struct {
	// Description string for the User API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description pulumi.StringPtrInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolPtrInput
	// API token which can be used in Equinix Metal API clients
	Token pulumi.StringPtrInput
	// UUID of the owner of the API key
	UserId pulumi.StringPtrInput
}

func (UserApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userApiKeyState)(nil)).Elem()
}

type userApiKeyArgs struct {
	// Description string for the User API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description string `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly bool `pulumi:"readOnly"`
}

// The set of arguments for constructing a UserApiKey resource.
type UserApiKeyArgs struct {
	// Description string for the User API Key resource
	// * `read-only` - Flag indicating whether the API key shoud be read-only
	Description pulumi.StringInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolInput
}

func (UserApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userApiKeyArgs)(nil)).Elem()
}

type UserApiKeyInput interface {
	pulumi.Input

	ToUserApiKeyOutput() UserApiKeyOutput
	ToUserApiKeyOutputWithContext(ctx context.Context) UserApiKeyOutput
}

func (*UserApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((*UserApiKey)(nil))
}

func (i *UserApiKey) ToUserApiKeyOutput() UserApiKeyOutput {
	return i.ToUserApiKeyOutputWithContext(context.Background())
}

func (i *UserApiKey) ToUserApiKeyOutputWithContext(ctx context.Context) UserApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyOutput)
}

func (i *UserApiKey) ToUserApiKeyPtrOutput() UserApiKeyPtrOutput {
	return i.ToUserApiKeyPtrOutputWithContext(context.Background())
}

func (i *UserApiKey) ToUserApiKeyPtrOutputWithContext(ctx context.Context) UserApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyPtrOutput)
}

type UserApiKeyPtrInput interface {
	pulumi.Input

	ToUserApiKeyPtrOutput() UserApiKeyPtrOutput
	ToUserApiKeyPtrOutputWithContext(ctx context.Context) UserApiKeyPtrOutput
}

type userApiKeyPtrType UserApiKeyArgs

func (*userApiKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserApiKey)(nil))
}

func (i *userApiKeyPtrType) ToUserApiKeyPtrOutput() UserApiKeyPtrOutput {
	return i.ToUserApiKeyPtrOutputWithContext(context.Background())
}

func (i *userApiKeyPtrType) ToUserApiKeyPtrOutputWithContext(ctx context.Context) UserApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyPtrOutput)
}

// UserApiKeyArrayInput is an input type that accepts UserApiKeyArray and UserApiKeyArrayOutput values.
// You can construct a concrete instance of `UserApiKeyArrayInput` via:
//
//          UserApiKeyArray{ UserApiKeyArgs{...} }
type UserApiKeyArrayInput interface {
	pulumi.Input

	ToUserApiKeyArrayOutput() UserApiKeyArrayOutput
	ToUserApiKeyArrayOutputWithContext(context.Context) UserApiKeyArrayOutput
}

type UserApiKeyArray []UserApiKeyInput

func (UserApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*UserApiKey)(nil))
}

func (i UserApiKeyArray) ToUserApiKeyArrayOutput() UserApiKeyArrayOutput {
	return i.ToUserApiKeyArrayOutputWithContext(context.Background())
}

func (i UserApiKeyArray) ToUserApiKeyArrayOutputWithContext(ctx context.Context) UserApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyArrayOutput)
}

// UserApiKeyMapInput is an input type that accepts UserApiKeyMap and UserApiKeyMapOutput values.
// You can construct a concrete instance of `UserApiKeyMapInput` via:
//
//          UserApiKeyMap{ "key": UserApiKeyArgs{...} }
type UserApiKeyMapInput interface {
	pulumi.Input

	ToUserApiKeyMapOutput() UserApiKeyMapOutput
	ToUserApiKeyMapOutputWithContext(context.Context) UserApiKeyMapOutput
}

type UserApiKeyMap map[string]UserApiKeyInput

func (UserApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*UserApiKey)(nil))
}

func (i UserApiKeyMap) ToUserApiKeyMapOutput() UserApiKeyMapOutput {
	return i.ToUserApiKeyMapOutputWithContext(context.Background())
}

func (i UserApiKeyMap) ToUserApiKeyMapOutputWithContext(ctx context.Context) UserApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserApiKeyMapOutput)
}

type UserApiKeyOutput struct {
	*pulumi.OutputState
}

func (UserApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserApiKey)(nil))
}

func (o UserApiKeyOutput) ToUserApiKeyOutput() UserApiKeyOutput {
	return o
}

func (o UserApiKeyOutput) ToUserApiKeyOutputWithContext(ctx context.Context) UserApiKeyOutput {
	return o
}

func (o UserApiKeyOutput) ToUserApiKeyPtrOutput() UserApiKeyPtrOutput {
	return o.ToUserApiKeyPtrOutputWithContext(context.Background())
}

func (o UserApiKeyOutput) ToUserApiKeyPtrOutputWithContext(ctx context.Context) UserApiKeyPtrOutput {
	return o.ApplyT(func(v UserApiKey) *UserApiKey {
		return &v
	}).(UserApiKeyPtrOutput)
}

type UserApiKeyPtrOutput struct {
	*pulumi.OutputState
}

func (UserApiKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserApiKey)(nil))
}

func (o UserApiKeyPtrOutput) ToUserApiKeyPtrOutput() UserApiKeyPtrOutput {
	return o
}

func (o UserApiKeyPtrOutput) ToUserApiKeyPtrOutputWithContext(ctx context.Context) UserApiKeyPtrOutput {
	return o
}

type UserApiKeyArrayOutput struct{ *pulumi.OutputState }

func (UserApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserApiKey)(nil))
}

func (o UserApiKeyArrayOutput) ToUserApiKeyArrayOutput() UserApiKeyArrayOutput {
	return o
}

func (o UserApiKeyArrayOutput) ToUserApiKeyArrayOutputWithContext(ctx context.Context) UserApiKeyArrayOutput {
	return o
}

func (o UserApiKeyArrayOutput) Index(i pulumi.IntInput) UserApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserApiKey {
		return vs[0].([]UserApiKey)[vs[1].(int)]
	}).(UserApiKeyOutput)
}

type UserApiKeyMapOutput struct{ *pulumi.OutputState }

func (UserApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserApiKey)(nil))
}

func (o UserApiKeyMapOutput) ToUserApiKeyMapOutput() UserApiKeyMapOutput {
	return o
}

func (o UserApiKeyMapOutput) ToUserApiKeyMapOutputWithContext(ctx context.Context) UserApiKeyMapOutput {
	return o
}

func (o UserApiKeyMapOutput) MapIndex(k pulumi.StringInput) UserApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserApiKey {
		return vs[0].(map[string]UserApiKey)[vs[1].(string)]
	}).(UserApiKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(UserApiKeyOutput{})
	pulumi.RegisterOutputType(UserApiKeyPtrOutput{})
	pulumi.RegisterOutputType(UserApiKeyArrayOutput{})
	pulumi.RegisterOutputType(UserApiKeyMapOutput{})
}
