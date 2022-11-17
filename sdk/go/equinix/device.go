// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Equinix Metal device resource. This can be used to create,
// modify, and delete devices.
//
// > **Note:** All arguments including the `rootPassword` and `userData` will be stored in
//
//	the raw state as plain-text.
//
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
//
// # Create a device and add it to coolProject
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := equinix - metal.NewDevice(ctx, "web1", &equinix-metal.DeviceArgs{
//				Hostname:        pulumi.String("tf.coreos2"),
//				Plan:            pulumi.String("c3.small.x86"),
//				Metro:           pulumi.String("sv"),
//				OperatingSystem: pulumi.String("ubuntu_20_04"),
//				BillingCycle:    pulumi.String("hourly"),
//				ProjectId:       pulumi.Any(local.Project_id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Same as above, but boot via iPXE initially, using the Ignition Provider for provisioning
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := equinix - metal.NewDevice(ctx, "pxe1", &equinix-metal.DeviceArgs{
//				Hostname:        pulumi.String("tf.coreos2-pxe"),
//				Plan:            pulumi.String("c3.small.x86"),
//				Metro:           pulumi.String("sv"),
//				OperatingSystem: pulumi.String("custom_ipxe"),
//				BillingCycle:    pulumi.String("hourly"),
//				ProjectId:       pulumi.Any(local.Project_id),
//				IpxeScriptUrl:   pulumi.String("https://rawgit.com/cloudnativelabs/pxe/master/metal/coreos-stable-metal.ipxe"),
//				AlwaysPxe:       pulumi.Bool(false),
//				UserData:        pulumi.Any(data.Ignition_config.Example.Rendered),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Create a device without a public IP address in facility ny5, with only a /30 private IPv4 subnet (4 IP addresses)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
//	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix-metal"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := equinix - metal.NewDevice(ctx, "web1", &equinix-metal.DeviceArgs{
//				Hostname: pulumi.String("tf.coreos2"),
//				Plan:     pulumi.String("c3.small.x86"),
//				Facilities: pulumi.StringArray{
//					pulumi.String("ny5"),
//				},
//				OperatingSystem: pulumi.String("ubuntu_20_04"),
//				BillingCycle:    pulumi.String("hourly"),
//				ProjectId:       pulumi.Any(local.Project_id),
//				IpAddresses: DeviceIpAddressArray{
//					&DeviceIpAddressArgs{
//						Type: pulumi.String("private_ipv4"),
//						Cidr: pulumi.Int(30),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deploy device on next-available reserved hardware and do custom partitioning.
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix"
//	"github.com/pulumi/pulumi-equinix-metal/sdk/v3/go/equinix-metal"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := equinix - metal.NewDevice(ctx, "web1", &equinix-metal.DeviceArgs{
//				Hostname: pulumi.String("tftest"),
//				Plan:     pulumi.String("c3.small.x86"),
//				Facilities: pulumi.StringArray{
//					pulumi.String("ny5"),
//				},
//				OperatingSystem:       pulumi.String("ubuntu_20_04"),
//				BillingCycle:          pulumi.String("hourly"),
//				ProjectId:             pulumi.Any(local.Project_id),
//				HardwareReservationId: pulumi.String("next-available"),
//				Storage:               pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"disks\": [\n", "    {\n", "      \"device\": \"/dev/sda\",\n", "      \"wipeTable\": true,\n", "      \"partitions\": [\n", "        {\n", "          \"label\": \"BIOS\",\n", "          \"number\": 1,\n", "          \"size\": \"4096\"\n", "        },\n", "        {\n", "          \"label\": \"SWAP\",\n", "          \"number\": 2,\n", "          \"size\": \"3993600\"\n", "        },\n", "        {\n", "          \"label\": \"ROOT\",\n", "          \"number\": 3,\n", "          \"size\": \"0\"\n", "        }\n", "      ]\n", "    }\n", "  ],\n", "  \"filesystems\": [\n", "    {\n", "      \"mount\": {\n", "        \"device\": \"/dev/sda3\",\n", "        \"format\": \"ext4\",\n", "        \"point\": \"/\",\n", "        \"create\": {\n", "          \"options\": [\n", "            \"-L\",\n", "            \"ROOT\"\n", "          ]\n", "        }\n", "      }\n", "    },\n", "    {\n", "      \"mount\": {\n", "        \"device\": \"/dev/sda2\",\n", "        \"format\": \"swap\",\n", "        \"point\": \"none\",\n", "        \"create\": {\n", "          \"options\": [\n", "            \"-L\",\n", "            \"SWAP\"\n", "          ]\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n")),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # This resource can be imported using an existing device ID
//
// ```sh
//
//	$ pulumi import equinix-metal:index/device:Device metal_device {existing_device_id}
//
// ```
type Device struct {
	pulumi.CustomResourceState

	// The ipv4 private IP assigned to the device
	AccessPrivateIpv4 pulumi.StringOutput `pulumi:"accessPrivateIpv4"`
	// The ipv4 maintenance IP assigned to the device
	AccessPublicIpv4 pulumi.StringOutput `pulumi:"accessPublicIpv4"`
	// The ipv6 maintenance IP assigned to the device
	AccessPublicIpv6 pulumi.StringOutput `pulumi:"accessPublicIpv6"`
	// If true, a device with OS `customIpxe` will continue to boot via iPXE on reboots.
	AlwaysPxe pulumi.BoolPtrOutput `pulumi:"alwaysPxe"`
	// monthly or hourly
	BillingCycle pulumi.StringOutput `pulumi:"billingCycle"`
	// The timestamp for when the device was created
	Created pulumi.StringOutput `pulumi:"created"`
	// A string of the desired Custom Data for the device.
	CustomData pulumi.StringPtrOutput `pulumi:"customData"`
	// The facility where the device is deployed.
	DeployedFacility pulumi.StringOutput `pulumi:"deployedFacility"`
	// ID of hardware reservation where this device was deployed. It is useful when using the `next-available` hardware reservation.
	DeployedHardwareReservationId pulumi.StringOutput `pulumi:"deployedHardwareReservationId"`
	// The device description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or `any` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with `metro`.
	Facilities pulumi.StringArrayOutput `pulumi:"facilities"`
	// Delete device even if it has volumes attached. Only applies for destroy action.
	ForceDetachVolumes pulumi.BoolPtrOutput `pulumi:"forceDetachVolumes"`
	// The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your
	// next available reservation automatically
	HardwareReservationId pulumi.StringPtrOutput `pulumi:"hardwareReservationId"`
	// The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// A list of IP address types for the device (structure is documented below).
	IpAddresses DeviceIpAddressArrayOutput `pulumi:"ipAddresses"`
	// URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.
	IpxeScriptUrl pulumi.StringPtrOutput `pulumi:"ipxeScriptUrl"`
	// Whether the device is locked
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// Metro area for the new device. Conflicts with `facilities`.
	Metro pulumi.StringPtrOutput `pulumi:"metro"`
	// Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/).
	// Will be one of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
	//
	// Deprecated: You should handle Network Type with the new metal_device_network_type resource.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:
	// * Public IPv4 at `metal_device.name.network.0`
	// * IPv6 at `metal_device.name.network.1`
	// * Private IPv4 at `metal_device.name.network.2`
	//   Elastic addresses then stack by type - an assigned public IPv4 will go after the management public IPv4 (to index 1), and will then shift the indices of the IPv6 and private IPv4. Assigned private IPv4 will go after the management private IPv4 (to the end of the network list).
	//   The fields of the network attributes are:
	Networks DeviceNetworkArrayOutput `pulumi:"networks"`
	// The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.
	OperatingSystem pulumi.StringOutput `pulumi:"operatingSystem"`
	// The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.
	Plan pulumi.StringOutput `pulumi:"plan"`
	// Ports assigned to the device
	Ports DevicePortArrayOutput `pulumi:"ports"`
	// The ID of the project in which to create the device
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the ProjectSshKey resource.
	ProjectSshKeyIds pulumi.StringArrayOutput `pulumi:"projectSshKeyIds"`
	// Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.
	Reinstall DeviceReinstallPtrOutput `pulumi:"reinstall"`
	// Root password to the server (disabled after 24 hours)
	RootPassword pulumi.StringOutput `pulumi:"rootPassword"`
	// List of IDs of SSH keys deployed in the device, can be both user and project SSH keys
	SshKeyIds pulumi.StringArrayOutput `pulumi:"sshKeyIds"`
	// The status of the device
	State pulumi.StringOutput `pulumi:"state"`
	// JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).
	Storage pulumi.StringPtrOutput `pulumi:"storage"`
	// Tags attached to the device
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Timestamp for device termination. For example `2021-09-03T16:32:00+03:00`. If you don't supply timezone info, timestamp is assumed to be in UTC.
	TerminationTime pulumi.StringPtrOutput `pulumi:"terminationTime"`
	// The timestamp for the last time the device was updated
	Updated pulumi.StringOutput `pulumi:"updated"`
	// A string of the desired User Data for the device.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019).
	WaitForReservationDeprovision pulumi.BoolPtrOutput `pulumi:"waitForReservationDeprovision"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OperatingSystem == nil {
		return nil, errors.New("invalid value for required argument 'OperatingSystem'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Device
	err := ctx.RegisterResource("equinix-metal:index/device:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("equinix-metal:index/device:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
	// The ipv4 private IP assigned to the device
	AccessPrivateIpv4 *string `pulumi:"accessPrivateIpv4"`
	// The ipv4 maintenance IP assigned to the device
	AccessPublicIpv4 *string `pulumi:"accessPublicIpv4"`
	// The ipv6 maintenance IP assigned to the device
	AccessPublicIpv6 *string `pulumi:"accessPublicIpv6"`
	// If true, a device with OS `customIpxe` will continue to boot via iPXE on reboots.
	AlwaysPxe *bool `pulumi:"alwaysPxe"`
	// monthly or hourly
	BillingCycle *string `pulumi:"billingCycle"`
	// The timestamp for when the device was created
	Created *string `pulumi:"created"`
	// A string of the desired Custom Data for the device.
	CustomData *string `pulumi:"customData"`
	// The facility where the device is deployed.
	DeployedFacility *string `pulumi:"deployedFacility"`
	// ID of hardware reservation where this device was deployed. It is useful when using the `next-available` hardware reservation.
	DeployedHardwareReservationId *string `pulumi:"deployedHardwareReservationId"`
	// The device description.
	Description *string `pulumi:"description"`
	// List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or `any` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with `metro`.
	Facilities []string `pulumi:"facilities"`
	// Delete device even if it has volumes attached. Only applies for destroy action.
	ForceDetachVolumes *bool `pulumi:"forceDetachVolumes"`
	// The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your
	// next available reservation automatically
	HardwareReservationId *string `pulumi:"hardwareReservationId"`
	// The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
	Hostname *string `pulumi:"hostname"`
	// A list of IP address types for the device (structure is documented below).
	IpAddresses []DeviceIpAddress `pulumi:"ipAddresses"`
	// URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.
	IpxeScriptUrl *string `pulumi:"ipxeScriptUrl"`
	// Whether the device is locked
	Locked *bool `pulumi:"locked"`
	// Metro area for the new device. Conflicts with `facilities`.
	Metro *string `pulumi:"metro"`
	// Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/).
	// Will be one of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
	//
	// Deprecated: You should handle Network Type with the new metal_device_network_type resource.
	NetworkType *string `pulumi:"networkType"`
	// The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:
	// * Public IPv4 at `metal_device.name.network.0`
	// * IPv6 at `metal_device.name.network.1`
	// * Private IPv4 at `metal_device.name.network.2`
	//   Elastic addresses then stack by type - an assigned public IPv4 will go after the management public IPv4 (to index 1), and will then shift the indices of the IPv6 and private IPv4. Assigned private IPv4 will go after the management private IPv4 (to the end of the network list).
	//   The fields of the network attributes are:
	Networks []DeviceNetwork `pulumi:"networks"`
	// The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.
	Plan *string `pulumi:"plan"`
	// Ports assigned to the device
	Ports []DevicePort `pulumi:"ports"`
	// The ID of the project in which to create the device
	ProjectId *string `pulumi:"projectId"`
	// Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the ProjectSshKey resource.
	ProjectSshKeyIds []string `pulumi:"projectSshKeyIds"`
	// Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.
	Reinstall *DeviceReinstall `pulumi:"reinstall"`
	// Root password to the server (disabled after 24 hours)
	RootPassword *string `pulumi:"rootPassword"`
	// List of IDs of SSH keys deployed in the device, can be both user and project SSH keys
	SshKeyIds []string `pulumi:"sshKeyIds"`
	// The status of the device
	State *string `pulumi:"state"`
	// JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).
	Storage *string `pulumi:"storage"`
	// Tags attached to the device
	Tags []string `pulumi:"tags"`
	// Timestamp for device termination. For example `2021-09-03T16:32:00+03:00`. If you don't supply timezone info, timestamp is assumed to be in UTC.
	TerminationTime *string `pulumi:"terminationTime"`
	// The timestamp for the last time the device was updated
	Updated *string `pulumi:"updated"`
	// A string of the desired User Data for the device.
	UserData *string `pulumi:"userData"`
	// Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019).
	WaitForReservationDeprovision *bool `pulumi:"waitForReservationDeprovision"`
}

type DeviceState struct {
	// The ipv4 private IP assigned to the device
	AccessPrivateIpv4 pulumi.StringPtrInput
	// The ipv4 maintenance IP assigned to the device
	AccessPublicIpv4 pulumi.StringPtrInput
	// The ipv6 maintenance IP assigned to the device
	AccessPublicIpv6 pulumi.StringPtrInput
	// If true, a device with OS `customIpxe` will continue to boot via iPXE on reboots.
	AlwaysPxe pulumi.BoolPtrInput
	// monthly or hourly
	BillingCycle pulumi.StringPtrInput
	// The timestamp for when the device was created
	Created pulumi.StringPtrInput
	// A string of the desired Custom Data for the device.
	CustomData pulumi.StringPtrInput
	// The facility where the device is deployed.
	DeployedFacility pulumi.StringPtrInput
	// ID of hardware reservation where this device was deployed. It is useful when using the `next-available` hardware reservation.
	DeployedHardwareReservationId pulumi.StringPtrInput
	// The device description.
	Description pulumi.StringPtrInput
	// List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or `any` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with `metro`.
	Facilities pulumi.StringArrayInput
	// Delete device even if it has volumes attached. Only applies for destroy action.
	ForceDetachVolumes pulumi.BoolPtrInput
	// The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your
	// next available reservation automatically
	HardwareReservationId pulumi.StringPtrInput
	// The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
	Hostname pulumi.StringPtrInput
	// A list of IP address types for the device (structure is documented below).
	IpAddresses DeviceIpAddressArrayInput
	// URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.
	IpxeScriptUrl pulumi.StringPtrInput
	// Whether the device is locked
	Locked pulumi.BoolPtrInput
	// Metro area for the new device. Conflicts with `facilities`.
	Metro pulumi.StringPtrInput
	// Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/).
	// Will be one of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
	//
	// Deprecated: You should handle Network Type with the new metal_device_network_type resource.
	NetworkType pulumi.StringPtrInput
	// The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:
	// * Public IPv4 at `metal_device.name.network.0`
	// * IPv6 at `metal_device.name.network.1`
	// * Private IPv4 at `metal_device.name.network.2`
	//   Elastic addresses then stack by type - an assigned public IPv4 will go after the management public IPv4 (to index 1), and will then shift the indices of the IPv6 and private IPv4. Assigned private IPv4 will go after the management private IPv4 (to the end of the network list).
	//   The fields of the network attributes are:
	Networks DeviceNetworkArrayInput
	// The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.
	OperatingSystem pulumi.StringPtrInput
	// The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.
	Plan pulumi.StringPtrInput
	// Ports assigned to the device
	Ports DevicePortArrayInput
	// The ID of the project in which to create the device
	ProjectId pulumi.StringPtrInput
	// Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the ProjectSshKey resource.
	ProjectSshKeyIds pulumi.StringArrayInput
	// Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.
	Reinstall DeviceReinstallPtrInput
	// Root password to the server (disabled after 24 hours)
	RootPassword pulumi.StringPtrInput
	// List of IDs of SSH keys deployed in the device, can be both user and project SSH keys
	SshKeyIds pulumi.StringArrayInput
	// The status of the device
	State pulumi.StringPtrInput
	// JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).
	Storage pulumi.StringPtrInput
	// Tags attached to the device
	Tags pulumi.StringArrayInput
	// Timestamp for device termination. For example `2021-09-03T16:32:00+03:00`. If you don't supply timezone info, timestamp is assumed to be in UTC.
	TerminationTime pulumi.StringPtrInput
	// The timestamp for the last time the device was updated
	Updated pulumi.StringPtrInput
	// A string of the desired User Data for the device.
	UserData pulumi.StringPtrInput
	// Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019).
	WaitForReservationDeprovision pulumi.BoolPtrInput
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// If true, a device with OS `customIpxe` will continue to boot via iPXE on reboots.
	AlwaysPxe *bool `pulumi:"alwaysPxe"`
	// monthly or hourly
	BillingCycle *string `pulumi:"billingCycle"`
	// A string of the desired Custom Data for the device.
	CustomData *string `pulumi:"customData"`
	// The device description.
	Description *string `pulumi:"description"`
	// List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or `any` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with `metro`.
	Facilities []string `pulumi:"facilities"`
	// Delete device even if it has volumes attached. Only applies for destroy action.
	ForceDetachVolumes *bool `pulumi:"forceDetachVolumes"`
	// The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your
	// next available reservation automatically
	HardwareReservationId *string `pulumi:"hardwareReservationId"`
	// The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
	Hostname *string `pulumi:"hostname"`
	// A list of IP address types for the device (structure is documented below).
	IpAddresses []DeviceIpAddress `pulumi:"ipAddresses"`
	// URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.
	IpxeScriptUrl *string `pulumi:"ipxeScriptUrl"`
	// Metro area for the new device. Conflicts with `facilities`.
	Metro *string `pulumi:"metro"`
	// The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.
	OperatingSystem string `pulumi:"operatingSystem"`
	// The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.
	Plan string `pulumi:"plan"`
	// The ID of the project in which to create the device
	ProjectId string `pulumi:"projectId"`
	// Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the ProjectSshKey resource.
	ProjectSshKeyIds []string `pulumi:"projectSshKeyIds"`
	// Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.
	Reinstall *DeviceReinstall `pulumi:"reinstall"`
	// JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).
	Storage *string `pulumi:"storage"`
	// Tags attached to the device
	Tags []string `pulumi:"tags"`
	// Timestamp for device termination. For example `2021-09-03T16:32:00+03:00`. If you don't supply timezone info, timestamp is assumed to be in UTC.
	TerminationTime *string `pulumi:"terminationTime"`
	// A string of the desired User Data for the device.
	UserData *string `pulumi:"userData"`
	// Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019).
	WaitForReservationDeprovision *bool `pulumi:"waitForReservationDeprovision"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// If true, a device with OS `customIpxe` will continue to boot via iPXE on reboots.
	AlwaysPxe pulumi.BoolPtrInput
	// monthly or hourly
	BillingCycle pulumi.StringPtrInput
	// A string of the desired Custom Data for the device.
	CustomData pulumi.StringPtrInput
	// The device description.
	Description pulumi.StringPtrInput
	// List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or `any` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with `metro`.
	Facilities pulumi.StringArrayInput
	// Delete device even if it has volumes attached. Only applies for destroy action.
	ForceDetachVolumes pulumi.BoolPtrInput
	// The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your
	// next available reservation automatically
	HardwareReservationId pulumi.StringPtrInput
	// The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
	Hostname pulumi.StringPtrInput
	// A list of IP address types for the device (structure is documented below).
	IpAddresses DeviceIpAddressArrayInput
	// URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.
	IpxeScriptUrl pulumi.StringPtrInput
	// Metro area for the new device. Conflicts with `facilities`.
	Metro pulumi.StringPtrInput
	// The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.
	OperatingSystem pulumi.StringInput
	// The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.
	Plan pulumi.StringInput
	// The ID of the project in which to create the device
	ProjectId pulumi.StringInput
	// Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the ProjectSshKey resource.
	ProjectSshKeyIds pulumi.StringArrayInput
	// Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.
	Reinstall DeviceReinstallPtrInput
	// JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).
	Storage pulumi.StringPtrInput
	// Tags attached to the device
	Tags pulumi.StringArrayInput
	// Timestamp for device termination. For example `2021-09-03T16:32:00+03:00`. If you don't supply timezone info, timestamp is assumed to be in UTC.
	TerminationTime pulumi.StringPtrInput
	// A string of the desired User Data for the device.
	UserData pulumi.StringPtrInput
	// Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019).
	WaitForReservationDeprovision pulumi.BoolPtrInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

// DeviceArrayInput is an input type that accepts DeviceArray and DeviceArrayOutput values.
// You can construct a concrete instance of `DeviceArrayInput` via:
//
//	DeviceArray{ DeviceArgs{...} }
type DeviceArrayInput interface {
	pulumi.Input

	ToDeviceArrayOutput() DeviceArrayOutput
	ToDeviceArrayOutputWithContext(context.Context) DeviceArrayOutput
}

type DeviceArray []DeviceInput

func (DeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (i DeviceArray) ToDeviceArrayOutput() DeviceArrayOutput {
	return i.ToDeviceArrayOutputWithContext(context.Background())
}

func (i DeviceArray) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceArrayOutput)
}

// DeviceMapInput is an input type that accepts DeviceMap and DeviceMapOutput values.
// You can construct a concrete instance of `DeviceMapInput` via:
//
//	DeviceMap{ "key": DeviceArgs{...} }
type DeviceMapInput interface {
	pulumi.Input

	ToDeviceMapOutput() DeviceMapOutput
	ToDeviceMapOutputWithContext(context.Context) DeviceMapOutput
}

type DeviceMap map[string]DeviceInput

func (DeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (i DeviceMap) ToDeviceMapOutput() DeviceMapOutput {
	return i.ToDeviceMapOutputWithContext(context.Background())
}

func (i DeviceMap) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceMapOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

type DeviceArrayOutput struct{ *pulumi.OutputState }

func (DeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Device)(nil)).Elem()
}

func (o DeviceArrayOutput) ToDeviceArrayOutput() DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) ToDeviceArrayOutputWithContext(ctx context.Context) DeviceArrayOutput {
	return o
}

func (o DeviceArrayOutput) Index(i pulumi.IntInput) DeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Device {
		return vs[0].([]*Device)[vs[1].(int)]
	}).(DeviceOutput)
}

type DeviceMapOutput struct{ *pulumi.OutputState }

func (DeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Device)(nil)).Elem()
}

func (o DeviceMapOutput) ToDeviceMapOutput() DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) ToDeviceMapOutputWithContext(ctx context.Context) DeviceMapOutput {
	return o
}

func (o DeviceMapOutput) MapIndex(k pulumi.StringInput) DeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Device {
		return vs[0].(map[string]*Device)[vs[1].(string)]
	}).(DeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceInput)(nil)).Elem(), &Device{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceArrayInput)(nil)).Elem(), DeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceMapInput)(nil)).Elem(), DeviceMap{})
	pulumi.RegisterOutputType(DeviceOutput{})
	pulumi.RegisterOutputType(DeviceArrayOutput{})
	pulumi.RegisterOutputType(DeviceMapOutput{})
}
