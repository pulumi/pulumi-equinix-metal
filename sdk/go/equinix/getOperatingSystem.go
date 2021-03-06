// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get Equinix Metal Operating System image.
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
// 		opt0 := "ubuntu"
// 		opt1 := "20.04"
// 		opt2 := "c3.medium.x86"
// 		example, err := equinix - metal.GetOperatingSystem(ctx, &equinix-metal.GetOperatingSystemArgs{
// 			Distro:          &opt0,
// 			Version:         &opt1,
// 			ProvisionableOn: &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix - metal.NewDevice(ctx, "server", &equinix-metal.DeviceArgs{
// 			Hostname: pulumi.String("tf.ubuntu"),
// 			Plan:     pulumi.String("c3.medium.x86"),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("ny5"),
// 			},
// 			OperatingSystem: pulumi.String(example.Id),
// 			BillingCycle:    pulumi.String("hourly"),
// 			ProjectId:       pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOperatingSystem(ctx *pulumi.Context, args *GetOperatingSystemArgs, opts ...pulumi.InvokeOption) (*GetOperatingSystemResult, error) {
	var rv GetOperatingSystemResult
	err := ctx.Invoke("equinix-metal:index/getOperatingSystem:getOperatingSystem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOperatingSystem.
type GetOperatingSystemArgs struct {
	// Name of the OS distribution.
	Distro *string `pulumi:"distro"`
	// Name or part of the name of the distribution. Case insensitive.
	Name *string `pulumi:"name"`
	// Plan name.
	ProvisionableOn *string `pulumi:"provisionableOn"`
	// Version of the distribution
	Version *string `pulumi:"version"`
}

// A collection of values returned by getOperatingSystem.
type GetOperatingSystemResult struct {
	Distro *string `pulumi:"distro"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	Name            *string `pulumi:"name"`
	ProvisionableOn *string `pulumi:"provisionableOn"`
	// Operating system slug (same as `id`)
	Slug    string  `pulumi:"slug"`
	Version *string `pulumi:"version"`
}
