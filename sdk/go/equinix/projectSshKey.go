// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Equinix Metal project SSH key resource to manage project-specific SSH keys.
// Project SSH keys will only be populated onto servers that belong to that project, in contrast to User SSH Keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/go/equinix/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		projectId := "<UUID_of_your_project>"
// 		testProjectSshKey, err := equinix - metal.NewProjectSshKey(ctx, "testProjectSshKey", &equinix-metal.ProjectSshKeyArgs{
// 			PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDM/unxJeFqxsTJcu6mhqsMHSaVlpu+Jj/P+44zrm6X/MAoHSX3X9oLgujEjjZ74yLfdfe0bJrbL2YgJzNaEkIQQ1VPMHB5EhTKUBGnzlPP0hHTnxsjAm9qDHgUPgvgFDQSAMzdJRJ0Cexo16Ph9VxCoLh3dxiE7s2gaM2FdVg7P8aSxKypsxAhYV3D0AwqzoOyT6WWhBoQ0xZ85XevOTnJCpImSemEGs6nVGEsWcEc1d1YvdxFjAK4SdsKUMkj4Dsy/leKsdi/DEAf356vbMT1UHsXXvy5TlHu/Pa6qF53v32Enz+nhKy7/8W2Yt2yWx8HnQcT2rug9lvCXagJO6oauqRTO77C4QZn13ZLMZgLT66S/tNh2EX0gi6vmIs5dth8uF+K6nxIyKJXbcA4ASg7F1OJrHKFZdTc5v1cPeq6PcbqGgc+8SrPYQmzvQqLoMBuxyos2hUkYOmw3aeWJj9nFa8Wu5WaN89mUeOqSkU4S5cgUzWUOmKey56B/j/s1sVys9rMhZapVs0wL4L9GBBM48N5jAQZnnpo85A8KsZq5ME22bTLqnxsDXqDYZvS7PSI6Dxi7eleOFE/NYYDkrgDLHTQri8ucDMVeVWHgoMY2bPXdn7KKy5jW5jKsf8EPARXg77A4gRYmgKrcwIKqJEUPqyxJBe0CPoGTqgXPRsUiQ== tomk@hp2"),
// 			ProjectId: pulumi.String(projectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.NewDevice(ctx, "testDevice", &equinix-metal.DeviceArgs{
// 			Hostname: pulumi.String("test"),
// 			Plan:     pulumi.String("baremetal_0"),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("ewr1"),
// 			},
// 			OperatingSystem: pulumi.String("ubuntu_16_04"),
// 			BillingCycle:    pulumi.String("hourly"),
// 			ProjectSshKeyIds: pulumi.StringArray{
// 				testProjectSshKey.ID(),
// 			},
// 			ProjectId: pulumi.String(projectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProjectSshKey struct {
	pulumi.CustomResourceState

	// The timestamp for when the SSH key was created
	Created pulumi.StringOutput `pulumi:"created"`
	// The fingerprint of the SSH key
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of parent project (same as project_id)
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The ID of parent project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The public key. If this is a file, it can be read using the file interpolation function
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The timestamp for the last time the SSH key was updated
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewProjectSshKey registers a new resource with the given unique name, arguments, and options.
func NewProjectSshKey(ctx *pulumi.Context,
	name string, args *ProjectSshKeyArgs, opts ...pulumi.ResourceOption) (*ProjectSshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	var resource ProjectSshKey
	err := ctx.RegisterResource("equinix-metal:index/projectSshKey:ProjectSshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectSshKey gets an existing ProjectSshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectSshKeyState, opts ...pulumi.ResourceOption) (*ProjectSshKey, error) {
	var resource ProjectSshKey
	err := ctx.ReadResource("equinix-metal:index/projectSshKey:ProjectSshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectSshKey resources.
type projectSshKeyState struct {
	// The timestamp for when the SSH key was created
	Created *string `pulumi:"created"`
	// The fingerprint of the SSH key
	Fingerprint *string `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The ID of parent project (same as project_id)
	OwnerId *string `pulumi:"ownerId"`
	// The ID of parent project
	ProjectId *string `pulumi:"projectId"`
	// The public key. If this is a file, it can be read using the file interpolation function
	PublicKey *string `pulumi:"publicKey"`
	// The timestamp for the last time the SSH key was updated
	Updated *string `pulumi:"updated"`
}

type ProjectSshKeyState struct {
	// The timestamp for when the SSH key was created
	Created pulumi.StringPtrInput
	// The fingerprint of the SSH key
	Fingerprint pulumi.StringPtrInput
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The ID of parent project (same as project_id)
	OwnerId pulumi.StringPtrInput
	// The ID of parent project
	ProjectId pulumi.StringPtrInput
	// The public key. If this is a file, it can be read using the file interpolation function
	PublicKey pulumi.StringPtrInput
	// The timestamp for the last time the SSH key was updated
	Updated pulumi.StringPtrInput
}

func (ProjectSshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSshKeyState)(nil)).Elem()
}

type projectSshKeyArgs struct {
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The ID of parent project
	ProjectId string `pulumi:"projectId"`
	// The public key. If this is a file, it can be read using the file interpolation function
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a ProjectSshKey resource.
type ProjectSshKeyArgs struct {
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The ID of parent project
	ProjectId pulumi.StringInput
	// The public key. If this is a file, it can be read using the file interpolation function
	PublicKey pulumi.StringInput
}

func (ProjectSshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSshKeyArgs)(nil)).Elem()
}

type ProjectSshKeyInput interface {
	pulumi.Input

	ToProjectSshKeyOutput() ProjectSshKeyOutput
	ToProjectSshKeyOutputWithContext(ctx context.Context) ProjectSshKeyOutput
}

func (*ProjectSshKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSshKey)(nil))
}

func (i *ProjectSshKey) ToProjectSshKeyOutput() ProjectSshKeyOutput {
	return i.ToProjectSshKeyOutputWithContext(context.Background())
}

func (i *ProjectSshKey) ToProjectSshKeyOutputWithContext(ctx context.Context) ProjectSshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSshKeyOutput)
}

type ProjectSshKeyOutput struct {
	*pulumi.OutputState
}

func (ProjectSshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectSshKey)(nil))
}

func (o ProjectSshKeyOutput) ToProjectSshKeyOutput() ProjectSshKeyOutput {
	return o
}

func (o ProjectSshKeyOutput) ToProjectSshKeyOutputWithContext(ctx context.Context) ProjectSshKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProjectSshKeyOutput{})
}
