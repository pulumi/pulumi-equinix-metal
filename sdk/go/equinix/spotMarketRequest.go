// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Equinix Metal Spot Market Request resource to allow you to
// manage spot market requests on your account. For more detail on Spot Market, see [this article in Equinix Metal documentation](https://metal.equinix.com/developers/docs/deploy/spot-market/).
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
// 		_, err := equinix - metal.NewSpotMarketRequest(ctx, "req", &equinix-metal.SpotMarketRequestArgs{
// 			ProjectId:   pulumi.Any(local.Project_id),
// 			MaxBidPrice: pulumi.Float64(0.03),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("ewr1"),
// 			},
// 			DevicesMin: pulumi.Int(1),
// 			DevicesMax: pulumi.Int(1),
// 			InstanceParameters: &equinix - metal.SpotMarketRequestInstanceParametersArgs{
// 				Hostname:        pulumi.String("testspot"),
// 				BillingCycle:    pulumi.String("hourly"),
// 				OperatingSystem: pulumi.String("coreos_stable"),
// 				Plan:            pulumi.String("t1.small.x86"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SpotMarketRequest struct {
	pulumi.CustomResourceState

	// Maximum number devices to be created
	DevicesMax pulumi.IntOutput `pulumi:"devicesMax"`
	// Miniumum number devices to be created
	DevicesMin pulumi.IntOutput `pulumi:"devicesMin"`
	// Facility IDs where devices should be created
	Facilities pulumi.StringArrayOutput `pulumi:"facilities"`
	// Device parameters. See device resource for details
	InstanceParameters SpotMarketRequestInstanceParametersOutput `pulumi:"instanceParameters"`
	// Maximum price user is willing to pay per hour per device
	MaxBidPrice pulumi.Float64Output `pulumi:"maxBidPrice"`
	// Project ID
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
	WaitForDevices pulumi.BoolPtrOutput `pulumi:"waitForDevices"`
}

// NewSpotMarketRequest registers a new resource with the given unique name, arguments, and options.
func NewSpotMarketRequest(ctx *pulumi.Context,
	name string, args *SpotMarketRequestArgs, opts ...pulumi.ResourceOption) (*SpotMarketRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevicesMax == nil {
		return nil, errors.New("invalid value for required argument 'DevicesMax'")
	}
	if args.DevicesMin == nil {
		return nil, errors.New("invalid value for required argument 'DevicesMin'")
	}
	if args.Facilities == nil {
		return nil, errors.New("invalid value for required argument 'Facilities'")
	}
	if args.InstanceParameters == nil {
		return nil, errors.New("invalid value for required argument 'InstanceParameters'")
	}
	if args.MaxBidPrice == nil {
		return nil, errors.New("invalid value for required argument 'MaxBidPrice'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource SpotMarketRequest
	err := ctx.RegisterResource("equinix-metal:index/spotMarketRequest:SpotMarketRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpotMarketRequest gets an existing SpotMarketRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpotMarketRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpotMarketRequestState, opts ...pulumi.ResourceOption) (*SpotMarketRequest, error) {
	var resource SpotMarketRequest
	err := ctx.ReadResource("equinix-metal:index/spotMarketRequest:SpotMarketRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpotMarketRequest resources.
type spotMarketRequestState struct {
	// Maximum number devices to be created
	DevicesMax *int `pulumi:"devicesMax"`
	// Miniumum number devices to be created
	DevicesMin *int `pulumi:"devicesMin"`
	// Facility IDs where devices should be created
	Facilities []string `pulumi:"facilities"`
	// Device parameters. See device resource for details
	InstanceParameters *SpotMarketRequestInstanceParameters `pulumi:"instanceParameters"`
	// Maximum price user is willing to pay per hour per device
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// Project ID
	ProjectId *string `pulumi:"projectId"`
	// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
	WaitForDevices *bool `pulumi:"waitForDevices"`
}

type SpotMarketRequestState struct {
	// Maximum number devices to be created
	DevicesMax pulumi.IntPtrInput
	// Miniumum number devices to be created
	DevicesMin pulumi.IntPtrInput
	// Facility IDs where devices should be created
	Facilities pulumi.StringArrayInput
	// Device parameters. See device resource for details
	InstanceParameters SpotMarketRequestInstanceParametersPtrInput
	// Maximum price user is willing to pay per hour per device
	MaxBidPrice pulumi.Float64PtrInput
	// Project ID
	ProjectId pulumi.StringPtrInput
	// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
	WaitForDevices pulumi.BoolPtrInput
}

func (SpotMarketRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*spotMarketRequestState)(nil)).Elem()
}

type spotMarketRequestArgs struct {
	// Maximum number devices to be created
	DevicesMax int `pulumi:"devicesMax"`
	// Miniumum number devices to be created
	DevicesMin int `pulumi:"devicesMin"`
	// Facility IDs where devices should be created
	Facilities []string `pulumi:"facilities"`
	// Device parameters. See device resource for details
	InstanceParameters SpotMarketRequestInstanceParameters `pulumi:"instanceParameters"`
	// Maximum price user is willing to pay per hour per device
	MaxBidPrice float64 `pulumi:"maxBidPrice"`
	// Project ID
	ProjectId string `pulumi:"projectId"`
	// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
	WaitForDevices *bool `pulumi:"waitForDevices"`
}

// The set of arguments for constructing a SpotMarketRequest resource.
type SpotMarketRequestArgs struct {
	// Maximum number devices to be created
	DevicesMax pulumi.IntInput
	// Miniumum number devices to be created
	DevicesMin pulumi.IntInput
	// Facility IDs where devices should be created
	Facilities pulumi.StringArrayInput
	// Device parameters. See device resource for details
	InstanceParameters SpotMarketRequestInstanceParametersInput
	// Maximum price user is willing to pay per hour per device
	MaxBidPrice pulumi.Float64Input
	// Project ID
	ProjectId pulumi.StringInput
	// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
	WaitForDevices pulumi.BoolPtrInput
}

func (SpotMarketRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spotMarketRequestArgs)(nil)).Elem()
}

type SpotMarketRequestInput interface {
	pulumi.Input

	ToSpotMarketRequestOutput() SpotMarketRequestOutput
	ToSpotMarketRequestOutputWithContext(ctx context.Context) SpotMarketRequestOutput
}

func (*SpotMarketRequest) ElementType() reflect.Type {
	return reflect.TypeOf((*SpotMarketRequest)(nil))
}

func (i *SpotMarketRequest) ToSpotMarketRequestOutput() SpotMarketRequestOutput {
	return i.ToSpotMarketRequestOutputWithContext(context.Background())
}

func (i *SpotMarketRequest) ToSpotMarketRequestOutputWithContext(ctx context.Context) SpotMarketRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotMarketRequestOutput)
}

func (i *SpotMarketRequest) ToSpotMarketRequestPtrOutput() SpotMarketRequestPtrOutput {
	return i.ToSpotMarketRequestPtrOutputWithContext(context.Background())
}

func (i *SpotMarketRequest) ToSpotMarketRequestPtrOutputWithContext(ctx context.Context) SpotMarketRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotMarketRequestPtrOutput)
}

type SpotMarketRequestPtrInput interface {
	pulumi.Input

	ToSpotMarketRequestPtrOutput() SpotMarketRequestPtrOutput
	ToSpotMarketRequestPtrOutputWithContext(ctx context.Context) SpotMarketRequestPtrOutput
}

type spotMarketRequestPtrType SpotMarketRequestArgs

func (*spotMarketRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SpotMarketRequest)(nil))
}

func (i *spotMarketRequestPtrType) ToSpotMarketRequestPtrOutput() SpotMarketRequestPtrOutput {
	return i.ToSpotMarketRequestPtrOutputWithContext(context.Background())
}

func (i *spotMarketRequestPtrType) ToSpotMarketRequestPtrOutputWithContext(ctx context.Context) SpotMarketRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotMarketRequestPtrOutput)
}

// SpotMarketRequestArrayInput is an input type that accepts SpotMarketRequestArray and SpotMarketRequestArrayOutput values.
// You can construct a concrete instance of `SpotMarketRequestArrayInput` via:
//
//          SpotMarketRequestArray{ SpotMarketRequestArgs{...} }
type SpotMarketRequestArrayInput interface {
	pulumi.Input

	ToSpotMarketRequestArrayOutput() SpotMarketRequestArrayOutput
	ToSpotMarketRequestArrayOutputWithContext(context.Context) SpotMarketRequestArrayOutput
}

type SpotMarketRequestArray []SpotMarketRequestInput

func (SpotMarketRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SpotMarketRequest)(nil))
}

func (i SpotMarketRequestArray) ToSpotMarketRequestArrayOutput() SpotMarketRequestArrayOutput {
	return i.ToSpotMarketRequestArrayOutputWithContext(context.Background())
}

func (i SpotMarketRequestArray) ToSpotMarketRequestArrayOutputWithContext(ctx context.Context) SpotMarketRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotMarketRequestArrayOutput)
}

// SpotMarketRequestMapInput is an input type that accepts SpotMarketRequestMap and SpotMarketRequestMapOutput values.
// You can construct a concrete instance of `SpotMarketRequestMapInput` via:
//
//          SpotMarketRequestMap{ "key": SpotMarketRequestArgs{...} }
type SpotMarketRequestMapInput interface {
	pulumi.Input

	ToSpotMarketRequestMapOutput() SpotMarketRequestMapOutput
	ToSpotMarketRequestMapOutputWithContext(context.Context) SpotMarketRequestMapOutput
}

type SpotMarketRequestMap map[string]SpotMarketRequestInput

func (SpotMarketRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SpotMarketRequest)(nil))
}

func (i SpotMarketRequestMap) ToSpotMarketRequestMapOutput() SpotMarketRequestMapOutput {
	return i.ToSpotMarketRequestMapOutputWithContext(context.Background())
}

func (i SpotMarketRequestMap) ToSpotMarketRequestMapOutputWithContext(ctx context.Context) SpotMarketRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotMarketRequestMapOutput)
}

type SpotMarketRequestOutput struct {
	*pulumi.OutputState
}

func (SpotMarketRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpotMarketRequest)(nil))
}

func (o SpotMarketRequestOutput) ToSpotMarketRequestOutput() SpotMarketRequestOutput {
	return o
}

func (o SpotMarketRequestOutput) ToSpotMarketRequestOutputWithContext(ctx context.Context) SpotMarketRequestOutput {
	return o
}

func (o SpotMarketRequestOutput) ToSpotMarketRequestPtrOutput() SpotMarketRequestPtrOutput {
	return o.ToSpotMarketRequestPtrOutputWithContext(context.Background())
}

func (o SpotMarketRequestOutput) ToSpotMarketRequestPtrOutputWithContext(ctx context.Context) SpotMarketRequestPtrOutput {
	return o.ApplyT(func(v SpotMarketRequest) *SpotMarketRequest {
		return &v
	}).(SpotMarketRequestPtrOutput)
}

type SpotMarketRequestPtrOutput struct {
	*pulumi.OutputState
}

func (SpotMarketRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpotMarketRequest)(nil))
}

func (o SpotMarketRequestPtrOutput) ToSpotMarketRequestPtrOutput() SpotMarketRequestPtrOutput {
	return o
}

func (o SpotMarketRequestPtrOutput) ToSpotMarketRequestPtrOutputWithContext(ctx context.Context) SpotMarketRequestPtrOutput {
	return o
}

type SpotMarketRequestArrayOutput struct{ *pulumi.OutputState }

func (SpotMarketRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpotMarketRequest)(nil))
}

func (o SpotMarketRequestArrayOutput) ToSpotMarketRequestArrayOutput() SpotMarketRequestArrayOutput {
	return o
}

func (o SpotMarketRequestArrayOutput) ToSpotMarketRequestArrayOutputWithContext(ctx context.Context) SpotMarketRequestArrayOutput {
	return o
}

func (o SpotMarketRequestArrayOutput) Index(i pulumi.IntInput) SpotMarketRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpotMarketRequest {
		return vs[0].([]SpotMarketRequest)[vs[1].(int)]
	}).(SpotMarketRequestOutput)
}

type SpotMarketRequestMapOutput struct{ *pulumi.OutputState }

func (SpotMarketRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SpotMarketRequest)(nil))
}

func (o SpotMarketRequestMapOutput) ToSpotMarketRequestMapOutput() SpotMarketRequestMapOutput {
	return o
}

func (o SpotMarketRequestMapOutput) ToSpotMarketRequestMapOutputWithContext(ctx context.Context) SpotMarketRequestMapOutput {
	return o
}

func (o SpotMarketRequestMapOutput) MapIndex(k pulumi.StringInput) SpotMarketRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SpotMarketRequest {
		return vs[0].(map[string]SpotMarketRequest)[vs[1].(string)]
	}).(SpotMarketRequestOutput)
}

func init() {
	pulumi.RegisterOutputType(SpotMarketRequestOutput{})
	pulumi.RegisterOutputType(SpotMarketRequestPtrOutput{})
	pulumi.RegisterOutputType(SpotMarketRequestArrayOutput{})
	pulumi.RegisterOutputType(SpotMarketRequestMapOutput{})
}
