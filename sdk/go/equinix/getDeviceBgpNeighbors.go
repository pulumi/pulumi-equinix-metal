// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this datasource to retrieve list of BGP neighbors of a device in the Equinix Metal host.
//
// To have any BGP neighbors listed, the device must be in BGP-enabled project and have a BGP session assigned.
//
// To learn more about using BGP in Equinix Metal, see the BgpSession resource documentation.
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
// 		test, err := equinix - metal.GetDeviceBgpNeighbors(ctx, &equinix-metal.GetDeviceBgpNeighborsArgs{
// 			DeviceId: "4c641195-25e5-4c3c-b2b7-4cd7a42c7b40",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("bgpNeighborsListing", test.BgpNeighbors)
// 		return nil
// 	})
// }
// ```
func GetDeviceBgpNeighbors(ctx *pulumi.Context, args *GetDeviceBgpNeighborsArgs, opts ...pulumi.InvokeOption) (*GetDeviceBgpNeighborsResult, error) {
	var rv GetDeviceBgpNeighborsResult
	err := ctx.Invoke("equinix-metal:index/getDeviceBgpNeighbors:getDeviceBgpNeighbors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeviceBgpNeighbors.
type GetDeviceBgpNeighborsArgs struct {
	// UUID of BGP-enabled device whose neighbors to list
	DeviceId string `pulumi:"deviceId"`
}

// A collection of values returned by getDeviceBgpNeighbors.
type GetDeviceBgpNeighborsResult struct {
	// array of BGP neighbor records with attributes:
	BgpNeighbors []GetDeviceBgpNeighborsBgpNeighbor `pulumi:"bgpNeighbors"`
	DeviceId     string                             `pulumi:"deviceId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
