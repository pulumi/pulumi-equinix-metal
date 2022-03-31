// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to request of create an Interconnection from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := equinix - metal.NewConnection(ctx, "test", &equinix-metal.ConnectionArgs{
// 			OrganizationId: pulumi.Any(local.My_organization_id),
// 			ProjectId:      pulumi.Any(local.My_project_id),
// 			Metro:          pulumi.String("sv"),
// 			Redundancy:     pulumi.String("redundant"),
// 			Type:           pulumi.String("shared"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Connection struct {
	pulumi.CustomResourceState

	// Description for the connection resource
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Facility where the connection will be created
	Facility pulumi.StringOutput `pulumi:"facility"`
	// Metro where the connection will be created
	Metro pulumi.StringOutput `pulumi:"metro"`
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Name of the connection resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization responsible for the connection
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the Connection datasource.
	Ports ConnectionPortArrayOutput `pulumi:"ports"`
	// ID of the project where the connection is scoped to, must be set for shared connection
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Connection redundancy - redundant or primary
	Redundancy pulumi.StringOutput `pulumi:"redundancy"`
	// Port speed in bits per second
	Speed pulumi.IntOutput `pulumi:"speed"`
	// Status of the connection resource
	Status pulumi.StringOutput `pulumi:"status"`
	// String list of tags
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	Token pulumi.StringOutput `pulumi:"token"`
	// Connection type - dedicated or shared
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.Redundancy == nil {
		return nil, errors.New("invalid value for required argument 'Redundancy'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Connection
	err := ctx.RegisterResource("equinix-metal:index/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("equinix-metal:index/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Description for the connection resource
	Description *string `pulumi:"description"`
	// Facility where the connection will be created
	Facility *string `pulumi:"facility"`
	// Metro where the connection will be created
	Metro *string `pulumi:"metro"`
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel
	Mode *string `pulumi:"mode"`
	// Name of the connection resource
	Name *string `pulumi:"name"`
	// ID of the organization responsible for the connection
	OrganizationId *string `pulumi:"organizationId"`
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the Connection datasource.
	Ports []ConnectionPort `pulumi:"ports"`
	// ID of the project where the connection is scoped to, must be set for shared connection
	ProjectId *string `pulumi:"projectId"`
	// Connection redundancy - redundant or primary
	Redundancy *string `pulumi:"redundancy"`
	// Port speed in bits per second
	Speed *int `pulumi:"speed"`
	// Status of the connection resource
	Status *string `pulumi:"status"`
	// String list of tags
	Tags []string `pulumi:"tags"`
	// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	Token *string `pulumi:"token"`
	// Connection type - dedicated or shared
	Type *string `pulumi:"type"`
}

type ConnectionState struct {
	// Description for the connection resource
	Description pulumi.StringPtrInput
	// Facility where the connection will be created
	Facility pulumi.StringPtrInput
	// Metro where the connection will be created
	Metro pulumi.StringPtrInput
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel
	Mode pulumi.StringPtrInput
	// Name of the connection resource
	Name pulumi.StringPtrInput
	// ID of the organization responsible for the connection
	OrganizationId pulumi.StringPtrInput
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the Connection datasource.
	Ports ConnectionPortArrayInput
	// ID of the project where the connection is scoped to, must be set for shared connection
	ProjectId pulumi.StringPtrInput
	// Connection redundancy - redundant or primary
	Redundancy pulumi.StringPtrInput
	// Port speed in bits per second
	Speed pulumi.IntPtrInput
	// Status of the connection resource
	Status pulumi.StringPtrInput
	// String list of tags
	Tags pulumi.StringArrayInput
	// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	Token pulumi.StringPtrInput
	// Connection type - dedicated or shared
	Type pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Description for the connection resource
	Description *string `pulumi:"description"`
	// Facility where the connection will be created
	Facility *string `pulumi:"facility"`
	// Metro where the connection will be created
	Metro *string `pulumi:"metro"`
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel
	Mode *string `pulumi:"mode"`
	// Name of the connection resource
	Name *string `pulumi:"name"`
	// ID of the organization responsible for the connection
	OrganizationId string `pulumi:"organizationId"`
	// ID of the project where the connection is scoped to, must be set for shared connection
	ProjectId *string `pulumi:"projectId"`
	// Connection redundancy - redundant or primary
	Redundancy string `pulumi:"redundancy"`
	// String list of tags
	Tags []string `pulumi:"tags"`
	// Connection type - dedicated or shared
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Description for the connection resource
	Description pulumi.StringPtrInput
	// Facility where the connection will be created
	Facility pulumi.StringPtrInput
	// Metro where the connection will be created
	Metro pulumi.StringPtrInput
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel
	Mode pulumi.StringPtrInput
	// Name of the connection resource
	Name pulumi.StringPtrInput
	// ID of the organization responsible for the connection
	OrganizationId pulumi.StringInput
	// ID of the project where the connection is scoped to, must be set for shared connection
	ProjectId pulumi.StringPtrInput
	// Connection redundancy - redundant or primary
	Redundancy pulumi.StringInput
	// String list of tags
	Tags pulumi.StringArrayInput
	// Connection type - dedicated or shared
	Type pulumi.StringInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//          ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//          ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].([]*Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].(map[string]*Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
