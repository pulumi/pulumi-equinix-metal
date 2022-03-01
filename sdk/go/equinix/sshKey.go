// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage User SSH keys on your Equinix Metal user account. If you create a new device in a project, all the keys of the project's collaborators will be injected to the device.
//
// The link between User SSH key and device is implicit. If you want to make sure that a key will be copied to a device, you must ensure that the device resource `dependsOn` the key resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix-metal"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := equinix - metal.NewSshKey(ctx, "key1", &equinix-metal.SshKeyArgs{
// 			PublicKey: readFileOrPanic("/home/terraform/.ssh/id_rsa.pub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.NewDevice(ctx, "test", &equinix-metal.DeviceArgs{
// 			Hostname: pulumi.String("test-device"),
// 			Plan:     pulumi.String("c3.small.x86"),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("sjc1"),
// 			},
// 			OperatingSystem: pulumi.String("ubuntu_20_04"),
// 			BillingCycle:    pulumi.String("hourly"),
// 			ProjectId:       pulumi.Any(local.Project_id),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			pulumi.Resource("metal_ssh_key.key1"),
// 		}))
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
// This resource can be imported using an existing SSH Key ID
//
// ```sh
//  $ pulumi import equinix-metal:index/sshKey:SshKey metal_ssh_key {existing_sshkey_id}
// ```
type SshKey struct {
	pulumi.CustomResourceState

	// The timestamp for when the SSH key was created
	Created pulumi.StringOutput `pulumi:"created"`
	// The fingerprint of the SSH key
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUID of the Equinix Metal API User who owns this key
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The timestamp for the last time the SSH key was updated
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewSshKey registers a new resource with the given unique name, arguments, and options.
func NewSshKey(ctx *pulumi.Context,
	name string, args *SshKeyArgs, opts ...pulumi.ResourceOption) (*SshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	var resource SshKey
	err := ctx.RegisterResource("equinix-metal:index/sshKey:SshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshKey gets an existing SshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshKeyState, opts ...pulumi.ResourceOption) (*SshKey, error) {
	var resource SshKey
	err := ctx.ReadResource("equinix-metal:index/sshKey:SshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshKey resources.
type sshKeyState struct {
	// The timestamp for when the SSH key was created
	Created *string `pulumi:"created"`
	// The fingerprint of the SSH key
	Fingerprint *string `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The UUID of the Equinix Metal API User who owns this key
	OwnerId *string `pulumi:"ownerId"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey *string `pulumi:"publicKey"`
	// The timestamp for the last time the SSH key was updated
	Updated *string `pulumi:"updated"`
}

type SshKeyState struct {
	// The timestamp for when the SSH key was created
	Created pulumi.StringPtrInput
	// The fingerprint of the SSH key
	Fingerprint pulumi.StringPtrInput
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The UUID of the Equinix Metal API User who owns this key
	OwnerId pulumi.StringPtrInput
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringPtrInput
	// The timestamp for the last time the SSH key was updated
	Updated pulumi.StringPtrInput
}

func (SshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyState)(nil)).Elem()
}

type sshKeyArgs struct {
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a SshKey resource.
type SshKeyArgs struct {
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringInput
}

func (SshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyArgs)(nil)).Elem()
}

type SshKeyInput interface {
	pulumi.Input

	ToSshKeyOutput() SshKeyOutput
	ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput
}

func (*SshKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SshKey)(nil)).Elem()
}

func (i *SshKey) ToSshKeyOutput() SshKeyOutput {
	return i.ToSshKeyOutputWithContext(context.Background())
}

func (i *SshKey) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshKeyOutput)
}

// SshKeyArrayInput is an input type that accepts SshKeyArray and SshKeyArrayOutput values.
// You can construct a concrete instance of `SshKeyArrayInput` via:
//
//          SshKeyArray{ SshKeyArgs{...} }
type SshKeyArrayInput interface {
	pulumi.Input

	ToSshKeyArrayOutput() SshKeyArrayOutput
	ToSshKeyArrayOutputWithContext(context.Context) SshKeyArrayOutput
}

type SshKeyArray []SshKeyInput

func (SshKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SshKey)(nil)).Elem()
}

func (i SshKeyArray) ToSshKeyArrayOutput() SshKeyArrayOutput {
	return i.ToSshKeyArrayOutputWithContext(context.Background())
}

func (i SshKeyArray) ToSshKeyArrayOutputWithContext(ctx context.Context) SshKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshKeyArrayOutput)
}

// SshKeyMapInput is an input type that accepts SshKeyMap and SshKeyMapOutput values.
// You can construct a concrete instance of `SshKeyMapInput` via:
//
//          SshKeyMap{ "key": SshKeyArgs{...} }
type SshKeyMapInput interface {
	pulumi.Input

	ToSshKeyMapOutput() SshKeyMapOutput
	ToSshKeyMapOutputWithContext(context.Context) SshKeyMapOutput
}

type SshKeyMap map[string]SshKeyInput

func (SshKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SshKey)(nil)).Elem()
}

func (i SshKeyMap) ToSshKeyMapOutput() SshKeyMapOutput {
	return i.ToSshKeyMapOutputWithContext(context.Background())
}

func (i SshKeyMap) ToSshKeyMapOutputWithContext(ctx context.Context) SshKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshKeyMapOutput)
}

type SshKeyOutput struct{ *pulumi.OutputState }

func (SshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshKey)(nil)).Elem()
}

func (o SshKeyOutput) ToSshKeyOutput() SshKeyOutput {
	return o
}

func (o SshKeyOutput) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return o
}

type SshKeyArrayOutput struct{ *pulumi.OutputState }

func (SshKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SshKey)(nil)).Elem()
}

func (o SshKeyArrayOutput) ToSshKeyArrayOutput() SshKeyArrayOutput {
	return o
}

func (o SshKeyArrayOutput) ToSshKeyArrayOutputWithContext(ctx context.Context) SshKeyArrayOutput {
	return o
}

func (o SshKeyArrayOutput) Index(i pulumi.IntInput) SshKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SshKey {
		return vs[0].([]*SshKey)[vs[1].(int)]
	}).(SshKeyOutput)
}

type SshKeyMapOutput struct{ *pulumi.OutputState }

func (SshKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SshKey)(nil)).Elem()
}

func (o SshKeyMapOutput) ToSshKeyMapOutput() SshKeyMapOutput {
	return o
}

func (o SshKeyMapOutput) ToSshKeyMapOutputWithContext(ctx context.Context) SshKeyMapOutput {
	return o
}

func (o SshKeyMapOutput) MapIndex(k pulumi.StringInput) SshKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SshKey {
		return vs[0].(map[string]*SshKey)[vs[1].(string)]
	}).(SshKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SshKeyInput)(nil)).Elem(), &SshKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshKeyArrayInput)(nil)).Elem(), SshKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshKeyMapInput)(nil)).Elem(), SshKeyMap{})
	pulumi.RegisterOutputType(SshKeyOutput{})
	pulumi.RegisterOutputType(SshKeyArrayOutput{})
	pulumi.RegisterOutputType(SshKeyMapOutput{})
}
