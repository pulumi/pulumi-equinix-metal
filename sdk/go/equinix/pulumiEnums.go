// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type BillingCycle pulumi.String

const (
	BillingCycleHourly  = BillingCycle("hourly")
	BillingCycleMonthly = BillingCycle("monthly")
)

func (BillingCycle) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e BillingCycle) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BillingCycle) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BillingCycle) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BillingCycle) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type Facility pulumi.String

const (
	FacilityEWR1 = Facility("ewr1")
	FacilitySJC1 = Facility("sjc1")
	FacilityDFW1 = Facility("dfw1")
	FacilityDFW2 = Facility("dfw2")
	FacilityAMS1 = Facility("ams1")
	FacilityNRT1 = Facility("nrt1")
	FacilitySEA1 = Facility("sea1")
	FacilityLAX1 = Facility("lax1")
	FacilityORD1 = Facility("ord1")
	FacilityATL1 = Facility("atl1")
	FacilityIAD1 = Facility("iad1")
	FacilitySIN1 = Facility("sin1")
	FacilityHKG1 = Facility("hkg1")
	FacilitySYD1 = Facility("syd1")
	FacilityMRS1 = Facility("mrs1")
	FacilityYYZ1 = Facility("yyz1")
	FacilityFRA2 = Facility("fra2")
)

func (Facility) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Facility) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Facility) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Facility) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Facility) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpBlockType pulumi.String

const (
	IpBlockTypeGlobalIPv4 = IpBlockType("global_ipv4")
	IpBlockTypePublicIPv4 = IpBlockType("public_ipv4")
)

func (IpBlockType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IpBlockType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpBlockType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpBlockType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpBlockType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkType pulumi.String

const (
	NetworkTypeLayer3           = NetworkType("layer3")
	NetworkTypeLayer2Individual = NetworkType("layer2-individual")
	NetworkTypeLayer2Bonded     = NetworkType("layer2-bonded")
	NetworkTypeHybrid           = NetworkType("hybrid")
)

func (NetworkType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e NetworkType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystem pulumi.String

const (
	OperatingSystemAlpine3          = OperatingSystem("alpine_3")
	OperatingSystemCentOS6          = OperatingSystem("centos_6")
	OperatingSystemCentOS7          = OperatingSystem("centos_7")
	OperatingSystemCoreOSAlpha      = OperatingSystem("coreos_alpha")
	OperatingSystemCoreOSBeta       = OperatingSystem("coreos_beta")
	OperatingSystemCoreOSStable     = OperatingSystem("coreos_stable")
	OperatingSystemCustomIPXE       = OperatingSystem("custom_ipxe")
	OperatingSystemDebian8          = OperatingSystem("debian_8")
	OperatingSystemDebian9          = OperatingSystem("debian_9")
	OperatingSystem_FreeBSD10_4     = OperatingSystem("freebsd_10_4")
	OperatingSystem_FreeBSD11_1     = OperatingSystem("freebsd_11_1")
	OperatingSystem_FreeBSD11_2     = OperatingSystem("freebsd_11_2")
	OperatingSystemFreeBSD12Testing = OperatingSystem("freebsd_12_testing")
	OperatingSystem_NixOS18_03      = OperatingSystem("nixos_18_03")
	OperatingSystem_NixOS19_03      = OperatingSystem("nixos_19_03")
	OperatingSystem_OpenSUSE42_3    = OperatingSystem("opensuse_42_3")
	OperatingSystemRancherOS        = OperatingSystem("rancher")
	OperatingSystemRHEL7            = OperatingSystem("rhel_7")
	OperatingSystemScientificLinux6 = OperatingSystem("scientific_6")
	OperatingSystemSLES12SP3        = OperatingSystem("suse_sles12_sp3")
	OperatingSystemUbuntu1404       = OperatingSystem("ubuntu_14_04")
	OperatingSystemUbuntu1604       = OperatingSystem("ubuntu_16_04")
	OperatingSystemUbuntu1710       = OperatingSystem("ubuntu_17_10")
	OperatingSystemUbuntu1804       = OperatingSystem("ubuntu_18_04")
	OperatingSystem_VMWareEsxi6_5   = OperatingSystem("vmware_esxi_6_5")
	OperatingSystemWindows2012R2    = OperatingSystem("windows_2012_r2")
	OperatingSystemWindows2016      = OperatingSystem("windows_2016")
)

func (OperatingSystem) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e OperatingSystem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type Plan pulumi.String

const (
	PlanC2LargeARM  = Plan("c2.large.arm")
	PlanC2MediumX86 = Plan("c2.medium.x86")
	PlanC1SmallX86  = Plan("baremetal_1")
	PlanC1LargeARM  = Plan("baremetal_2a")
	PlanC1XLargeX86 = Plan("baremetal_3")
	PlanX2XLargeX86 = Plan("x2.xlarge.x86")
	PlanX1SmallX86  = Plan("baremetal_1e")
	PlanG2LargeX86  = Plan("g2.large.x86")
	PlanM2XLargeX86 = Plan("m2.xlarge.x86")
	PlanM1XLargeX86 = Plan("baremetal_2")
	PlanT1SmallX86  = Plan("baremetal_0")
	PlanS1LargeX86  = Plan("baremetal_s")
)

func (Plan) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Plan) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Plan) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Plan) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Plan) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
