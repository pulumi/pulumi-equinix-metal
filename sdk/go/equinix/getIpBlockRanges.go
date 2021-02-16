// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this datasource to get CIDR expressions for allocated IP blocks of all the types in a project, optionally filtered by facility.
//
// There are four types of IP blocks in Equinix Metal: global IPv4, public IPv4, private IPv4 and IPv6. Both global and public IPv4 are routable from the Internet. Public IPv4 block is allocated in a facility, and addresses from it can only be assigned to devices in that facility. Addresses from Global IPv4 block can be assigned to a device in any facility.
//
// The datasource has 4 list attributes: `globalIpv4`, `publicIpv4`, `privateIpv4` and `ipv6`, each listing CIDR notation (`<network>/<mask>`) of respective blocks from the project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/go/equinix-metal"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		projectId := "<UUID_of_your_project>"
// 		test, err := equinix - metal.GetIpBlockRanges(ctx, &equinix-metal.GetIpBlockRangesArgs{
// 			ProjectId: projectId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("out", test)
// 		return nil
// 	})
// }
// ```
func GetIpBlockRanges(ctx *pulumi.Context, args *GetIpBlockRangesArgs, opts ...pulumi.InvokeOption) (*GetIpBlockRangesResult, error) {
	var rv GetIpBlockRangesResult
	err := ctx.Invoke("equinix-metal:index/getIpBlockRanges:getIpBlockRanges", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpBlockRanges.
type GetIpBlockRangesArgs struct {
	// Facility code filtering the IP blocks. Global IPv4 blcoks will be listed anyway. If you omit this, all the block from the project will be listed.
	Facility *string `pulumi:"facility"`
	// ID of the project from which to list the blocks.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getIpBlockRanges.
type GetIpBlockRangesResult struct {
	Facility *string `pulumi:"facility"`
	// list of CIDR expressions for Global IPv4 blocks in the project
	GlobalIpv4s []string `pulumi:"globalIpv4s"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// list of CIDR expressions for IPv6 blocks in the project
	Ipv6s []string `pulumi:"ipv6s"`
	// list of CIDR expressions for Private IPv4 blocks in the project
	PrivateIpv4s []string `pulumi:"privateIpv4s"`
	ProjectId    string   `pulumi:"projectId"`
	// list of CIDR expressions for Public IPv4 blocks in the project
	PublicIpv4s []string `pulumi:"publicIpv4s"`
}
