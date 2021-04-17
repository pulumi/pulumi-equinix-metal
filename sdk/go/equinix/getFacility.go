// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Equinix Metal facility datasource.
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
// 		ewr1, err := equinix - metal.GetFacility(ctx, &equinix-metal.GetFacilityArgs{
// 			Code: "ewr1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("id", ewr1.Id)
// 		return nil
// 	})
// }
// ```
func GetFacility(ctx *pulumi.Context, args *GetFacilityArgs, opts ...pulumi.InvokeOption) (*GetFacilityResult, error) {
	var rv GetFacilityResult
	err := ctx.Invoke("equinix-metal:index/getFacility:getFacility", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFacility.
type GetFacilityArgs struct {
	// The facility code
	Code string `pulumi:"code"`
	// The features of the facility
	Features []string `pulumi:"features"`
	// The name of the facilityg system running on the device
	Name *string `pulumi:"name"`
}

// A collection of values returned by getFacility.
type GetFacilityResult struct {
	Code string `pulumi:"code"`
	// The features of the facility
	Features []string `pulumi:"features"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the facilityg system running on the device
	Name string `pulumi:"name"`
}
