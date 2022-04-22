// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get Equinix Metal Spot Market Price for a plan.
//
// ## Example Usage
//
// Lookup by facility:
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
// 		_, err := equinix - metal.GetSpotMarketPrice(ctx, &GetSpotMarketPriceArgs{
// 			Facility: pulumi.StringRef("ny5"),
// 			Plan:     "c3.small.x86",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Lookup by metro:
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
// 		_, err := equinix - metal.GetSpotMarketPrice(ctx, &GetSpotMarketPriceArgs{
// 			Metro: pulumi.StringRef("sv"),
// 			Plan:  "c3.small.x86",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSpotMarketPrice(ctx *pulumi.Context, args *GetSpotMarketPriceArgs, opts ...pulumi.InvokeOption) (*GetSpotMarketPriceResult, error) {
	var rv GetSpotMarketPriceResult
	err := ctx.Invoke("equinix-metal:index/getSpotMarketPrice:getSpotMarketPrice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpotMarketPrice.
type GetSpotMarketPriceArgs struct {
	// Name of the facility.
	Facility *string `pulumi:"facility"`
	// Name of the metro.
	Metro *string `pulumi:"metro"`
	// Name of the plan.
	Plan string `pulumi:"plan"`
}

// A collection of values returned by getSpotMarketPrice.
type GetSpotMarketPriceResult struct {
	Facility *string `pulumi:"facility"`
	// The provider-assigned unique ID for this managed resource.
	Id    string  `pulumi:"id"`
	Metro *string `pulumi:"metro"`
	Plan  string  `pulumi:"plan"`
	// Current spot market price for given plan in given facility.
	Price float64 `pulumi:"price"`
}

func GetSpotMarketPriceOutput(ctx *pulumi.Context, args GetSpotMarketPriceOutputArgs, opts ...pulumi.InvokeOption) GetSpotMarketPriceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSpotMarketPriceResult, error) {
			args := v.(GetSpotMarketPriceArgs)
			r, err := GetSpotMarketPrice(ctx, &args, opts...)
			var s GetSpotMarketPriceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSpotMarketPriceResultOutput)
}

// A collection of arguments for invoking getSpotMarketPrice.
type GetSpotMarketPriceOutputArgs struct {
	// Name of the facility.
	Facility pulumi.StringPtrInput `pulumi:"facility"`
	// Name of the metro.
	Metro pulumi.StringPtrInput `pulumi:"metro"`
	// Name of the plan.
	Plan pulumi.StringInput `pulumi:"plan"`
}

func (GetSpotMarketPriceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpotMarketPriceArgs)(nil)).Elem()
}

// A collection of values returned by getSpotMarketPrice.
type GetSpotMarketPriceResultOutput struct{ *pulumi.OutputState }

func (GetSpotMarketPriceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpotMarketPriceResult)(nil)).Elem()
}

func (o GetSpotMarketPriceResultOutput) ToGetSpotMarketPriceResultOutput() GetSpotMarketPriceResultOutput {
	return o
}

func (o GetSpotMarketPriceResultOutput) ToGetSpotMarketPriceResultOutputWithContext(ctx context.Context) GetSpotMarketPriceResultOutput {
	return o
}

func (o GetSpotMarketPriceResultOutput) Facility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSpotMarketPriceResult) *string { return v.Facility }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSpotMarketPriceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpotMarketPriceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSpotMarketPriceResultOutput) Metro() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSpotMarketPriceResult) *string { return v.Metro }).(pulumi.StringPtrOutput)
}

func (o GetSpotMarketPriceResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpotMarketPriceResult) string { return v.Plan }).(pulumi.StringOutput)
}

// Current spot market price for given plan in given facility.
func (o GetSpotMarketPriceResultOutput) Price() pulumi.Float64Output {
	return o.ApplyT(func(v GetSpotMarketPriceResult) float64 { return v.Price }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(GetSpotMarketPriceResultOutput{})
}
