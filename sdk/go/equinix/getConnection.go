// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a connection resource from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)
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
// 		_, err := equinix - metal.GetConnection(ctx, &equinix-metal.GetConnectionArgs{
// 			ConnectionId: "4347e805-eb46-4699-9eb9-5c116e6a017d",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetConnection(ctx *pulumi.Context, args *GetConnectionArgs, opts ...pulumi.InvokeOption) (*GetConnectionResult, error) {
	var rv GetConnectionResult
	err := ctx.Invoke("equinix-metal:index/getConnection:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConnection.
type GetConnectionArgs struct {
	// ID of the connection resource
	ConnectionId string `pulumi:"connectionId"`
}

// A collection of values returned by getConnection.
type GetConnectionResult struct {
	ConnectionId string `pulumi:"connectionId"`
	// Description of the connection resource
	Description string `pulumi:"description"`
	// Slug of a facility to which the connection belongs
	Facility string `pulumi:"facility"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Port name
	Name string `pulumi:"name"`
	// ID of organization to which the connection belongs
	OrganizationId string `pulumi:"organizationId"`
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
	Ports []GetConnectionPort `pulumi:"ports"`
	// Connection redundancy, reduntant or primary
	Redundancy string `pulumi:"redundancy"`
	// Port speed in bits per second
	Speed int `pulumi:"speed"`
	// Port status
	Status string `pulumi:"status"`
	// Fabric Token for the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	Token string `pulumi:"token"`
	// Connection type, dedicated or shared
	Type string `pulumi:"type"`
}
