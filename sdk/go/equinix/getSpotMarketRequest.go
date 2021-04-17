// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSpotMarketRequest(ctx *pulumi.Context, args *LookupSpotMarketRequestArgs, opts ...pulumi.InvokeOption) (*LookupSpotMarketRequestResult, error) {
	var rv LookupSpotMarketRequestResult
	err := ctx.Invoke("equinix-metal:index/getSpotMarketRequest:getSpotMarketRequest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpotMarketRequest.
type LookupSpotMarketRequestArgs struct {
	// The id of the Spot Market Request
	RequestId string `pulumi:"requestId"`
}

// A collection of values returned by getSpotMarketRequest.
type LookupSpotMarketRequestResult struct {
	// List of IDs of devices spawned by the referenced Spot Market Request
	DeviceIds []string `pulumi:"deviceIds"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	RequestId string `pulumi:"requestId"`
}
