// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to read ports of existing devices. You can read port by either its UUID, or by a device UUID and port name.
//
// ## Example Usage
//
// Create a device and read it's eth0 port to the datasource.
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
// 		projectId := "<UUID_of_your_project>"
// 		testDevice, err := equinix - metal.NewDevice(ctx, "testDevice", &equinix-metal.DeviceArgs{
// 			Hostname: pulumi.String("tfacc-test-device-port"),
// 			Plan:     pulumi.String("c3.medium.x86"),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("sv15"),
// 			},
// 			OperatingSystem: pulumi.String("ubuntu_20_04"),
// 			BillingCycle:    pulumi.String("hourly"),
// 			ProjectId:       pulumi.String(projectId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetPort(ctx *pulumi.Context, args *GetPortArgs, opts ...pulumi.InvokeOption) (*GetPortResult, error) {
	var rv GetPortResult
	err := ctx.Invoke("equinix-metal:index/getPort:getPort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPort.
type GetPortArgs struct {
	DeviceId *string `pulumi:"deviceId"`
	// ID of the port to read, conflicts with device_id.
	Id *string `pulumi:"id"`
	// Whether to look for public or private block.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getPort.
type GetPortResult struct {
	// UUID of the bond port"
	BondId string `pulumi:"bondId"`
	// Name of the bond port
	BondName string `pulumi:"bondName"`
	// Flag indicating whether the port is bonded
	Bonded   bool    `pulumi:"bonded"`
	DeviceId *string `pulumi:"deviceId"`
	// Flag indicating whether the port can be removed from a bond
	DisbondSupported bool    `pulumi:"disbondSupported"`
	Id               *string `pulumi:"id"`
	// MAC address of the port
	Mac  string `pulumi:"mac"`
	Name string `pulumi:"name"`
	// UUID of native VLAN of the port
	NativeVlanId string `pulumi:"nativeVlanId"`
	// One of layer2-bonded, layer2-individual, layer3, hybrid, hybrid-bonded
	NetworkType string `pulumi:"networkType"`
	// Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
	Type string `pulumi:"type"`
	// UUIDs of attached VLANs
	VlanIds []string `pulumi:"vlanIds"`
}
