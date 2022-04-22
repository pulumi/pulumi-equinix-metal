// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BillingCycle string

const (
	BillingCycleHourly  = BillingCycle("hourly")
	BillingCycleMonthly = BillingCycle("monthly")
)

func (BillingCycle) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingCycle)(nil)).Elem()
}

func (e BillingCycle) ToBillingCycleOutput() BillingCycleOutput {
	return pulumi.ToOutput(e).(BillingCycleOutput)
}

func (e BillingCycle) ToBillingCycleOutputWithContext(ctx context.Context) BillingCycleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BillingCycleOutput)
}

func (e BillingCycle) ToBillingCyclePtrOutput() BillingCyclePtrOutput {
	return e.ToBillingCyclePtrOutputWithContext(context.Background())
}

func (e BillingCycle) ToBillingCyclePtrOutputWithContext(ctx context.Context) BillingCyclePtrOutput {
	return BillingCycle(e).ToBillingCycleOutputWithContext(ctx).ToBillingCyclePtrOutputWithContext(ctx)
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

type BillingCycleOutput struct{ *pulumi.OutputState }

func (BillingCycleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingCycle)(nil)).Elem()
}

func (o BillingCycleOutput) ToBillingCycleOutput() BillingCycleOutput {
	return o
}

func (o BillingCycleOutput) ToBillingCycleOutputWithContext(ctx context.Context) BillingCycleOutput {
	return o
}

func (o BillingCycleOutput) ToBillingCyclePtrOutput() BillingCyclePtrOutput {
	return o.ToBillingCyclePtrOutputWithContext(context.Background())
}

func (o BillingCycleOutput) ToBillingCyclePtrOutputWithContext(ctx context.Context) BillingCyclePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BillingCycle) *BillingCycle {
		return &v
	}).(BillingCyclePtrOutput)
}

func (o BillingCycleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BillingCycleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BillingCycle) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BillingCycleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BillingCycleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BillingCycle) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BillingCyclePtrOutput struct{ *pulumi.OutputState }

func (BillingCyclePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingCycle)(nil)).Elem()
}

func (o BillingCyclePtrOutput) ToBillingCyclePtrOutput() BillingCyclePtrOutput {
	return o
}

func (o BillingCyclePtrOutput) ToBillingCyclePtrOutputWithContext(ctx context.Context) BillingCyclePtrOutput {
	return o
}

func (o BillingCyclePtrOutput) Elem() BillingCycleOutput {
	return o.ApplyT(func(v *BillingCycle) BillingCycle {
		if v != nil {
			return *v
		}
		var ret BillingCycle
		return ret
	}).(BillingCycleOutput)
}

func (o BillingCyclePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BillingCyclePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BillingCycle) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BillingCycleInput is an input type that accepts BillingCycleArgs and BillingCycleOutput values.
// You can construct a concrete instance of `BillingCycleInput` via:
//
//          BillingCycleArgs{...}
type BillingCycleInput interface {
	pulumi.Input

	ToBillingCycleOutput() BillingCycleOutput
	ToBillingCycleOutputWithContext(context.Context) BillingCycleOutput
}

var billingCyclePtrType = reflect.TypeOf((**BillingCycle)(nil)).Elem()

type BillingCyclePtrInput interface {
	pulumi.Input

	ToBillingCyclePtrOutput() BillingCyclePtrOutput
	ToBillingCyclePtrOutputWithContext(context.Context) BillingCyclePtrOutput
}

type billingCyclePtr string

func BillingCyclePtr(v string) BillingCyclePtrInput {
	return (*billingCyclePtr)(&v)
}

func (*billingCyclePtr) ElementType() reflect.Type {
	return billingCyclePtrType
}

func (in *billingCyclePtr) ToBillingCyclePtrOutput() BillingCyclePtrOutput {
	return pulumi.ToOutput(in).(BillingCyclePtrOutput)
}

func (in *billingCyclePtr) ToBillingCyclePtrOutputWithContext(ctx context.Context) BillingCyclePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BillingCyclePtrOutput)
}

type Facility string

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
	FacilityAM6  = Facility("am6")
	FacilityDC13 = Facility("dc13")
	FacilityCH3  = Facility("ch3")
	FacilityDA3  = Facility("da3")
	FacilityDA11 = Facility("da11")
	FacilityLA4  = Facility("la4")
	FacilityNY5  = Facility("ny5")
	FacilitySG1  = Facility("sg1")
	FacilitySV15 = Facility("sv15")
)

func (Facility) ElementType() reflect.Type {
	return reflect.TypeOf((*Facility)(nil)).Elem()
}

func (e Facility) ToFacilityOutput() FacilityOutput {
	return pulumi.ToOutput(e).(FacilityOutput)
}

func (e Facility) ToFacilityOutputWithContext(ctx context.Context) FacilityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FacilityOutput)
}

func (e Facility) ToFacilityPtrOutput() FacilityPtrOutput {
	return e.ToFacilityPtrOutputWithContext(context.Background())
}

func (e Facility) ToFacilityPtrOutputWithContext(ctx context.Context) FacilityPtrOutput {
	return Facility(e).ToFacilityOutputWithContext(ctx).ToFacilityPtrOutputWithContext(ctx)
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

type FacilityOutput struct{ *pulumi.OutputState }

func (FacilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Facility)(nil)).Elem()
}

func (o FacilityOutput) ToFacilityOutput() FacilityOutput {
	return o
}

func (o FacilityOutput) ToFacilityOutputWithContext(ctx context.Context) FacilityOutput {
	return o
}

func (o FacilityOutput) ToFacilityPtrOutput() FacilityPtrOutput {
	return o.ToFacilityPtrOutputWithContext(context.Background())
}

func (o FacilityOutput) ToFacilityPtrOutputWithContext(ctx context.Context) FacilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Facility) *Facility {
		return &v
	}).(FacilityPtrOutput)
}

func (o FacilityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FacilityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Facility) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FacilityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FacilityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Facility) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FacilityPtrOutput struct{ *pulumi.OutputState }

func (FacilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Facility)(nil)).Elem()
}

func (o FacilityPtrOutput) ToFacilityPtrOutput() FacilityPtrOutput {
	return o
}

func (o FacilityPtrOutput) ToFacilityPtrOutputWithContext(ctx context.Context) FacilityPtrOutput {
	return o
}

func (o FacilityPtrOutput) Elem() FacilityOutput {
	return o.ApplyT(func(v *Facility) Facility {
		if v != nil {
			return *v
		}
		var ret Facility
		return ret
	}).(FacilityOutput)
}

func (o FacilityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FacilityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Facility) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FacilityInput is an input type that accepts FacilityArgs and FacilityOutput values.
// You can construct a concrete instance of `FacilityInput` via:
//
//          FacilityArgs{...}
type FacilityInput interface {
	pulumi.Input

	ToFacilityOutput() FacilityOutput
	ToFacilityOutputWithContext(context.Context) FacilityOutput
}

var facilityPtrType = reflect.TypeOf((**Facility)(nil)).Elem()

type FacilityPtrInput interface {
	pulumi.Input

	ToFacilityPtrOutput() FacilityPtrOutput
	ToFacilityPtrOutputWithContext(context.Context) FacilityPtrOutput
}

type facilityPtr string

func FacilityPtr(v string) FacilityPtrInput {
	return (*facilityPtr)(&v)
}

func (*facilityPtr) ElementType() reflect.Type {
	return facilityPtrType
}

func (in *facilityPtr) ToFacilityPtrOutput() FacilityPtrOutput {
	return pulumi.ToOutput(in).(FacilityPtrOutput)
}

func (in *facilityPtr) ToFacilityPtrOutputWithContext(ctx context.Context) FacilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FacilityPtrOutput)
}

type IpBlockType string

const (
	IpBlockTypeGlobalIPv4 = IpBlockType("global_ipv4")
	IpBlockTypePublicIPv4 = IpBlockType("public_ipv4")
)

func (IpBlockType) ElementType() reflect.Type {
	return reflect.TypeOf((*IpBlockType)(nil)).Elem()
}

func (e IpBlockType) ToIpBlockTypeOutput() IpBlockTypeOutput {
	return pulumi.ToOutput(e).(IpBlockTypeOutput)
}

func (e IpBlockType) ToIpBlockTypeOutputWithContext(ctx context.Context) IpBlockTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpBlockTypeOutput)
}

func (e IpBlockType) ToIpBlockTypePtrOutput() IpBlockTypePtrOutput {
	return e.ToIpBlockTypePtrOutputWithContext(context.Background())
}

func (e IpBlockType) ToIpBlockTypePtrOutputWithContext(ctx context.Context) IpBlockTypePtrOutput {
	return IpBlockType(e).ToIpBlockTypeOutputWithContext(ctx).ToIpBlockTypePtrOutputWithContext(ctx)
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

type IpBlockTypeOutput struct{ *pulumi.OutputState }

func (IpBlockTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpBlockType)(nil)).Elem()
}

func (o IpBlockTypeOutput) ToIpBlockTypeOutput() IpBlockTypeOutput {
	return o
}

func (o IpBlockTypeOutput) ToIpBlockTypeOutputWithContext(ctx context.Context) IpBlockTypeOutput {
	return o
}

func (o IpBlockTypeOutput) ToIpBlockTypePtrOutput() IpBlockTypePtrOutput {
	return o.ToIpBlockTypePtrOutputWithContext(context.Background())
}

func (o IpBlockTypeOutput) ToIpBlockTypePtrOutputWithContext(ctx context.Context) IpBlockTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpBlockType) *IpBlockType {
		return &v
	}).(IpBlockTypePtrOutput)
}

func (o IpBlockTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpBlockTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpBlockType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpBlockTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpBlockTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpBlockType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpBlockTypePtrOutput struct{ *pulumi.OutputState }

func (IpBlockTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpBlockType)(nil)).Elem()
}

func (o IpBlockTypePtrOutput) ToIpBlockTypePtrOutput() IpBlockTypePtrOutput {
	return o
}

func (o IpBlockTypePtrOutput) ToIpBlockTypePtrOutputWithContext(ctx context.Context) IpBlockTypePtrOutput {
	return o
}

func (o IpBlockTypePtrOutput) Elem() IpBlockTypeOutput {
	return o.ApplyT(func(v *IpBlockType) IpBlockType {
		if v != nil {
			return *v
		}
		var ret IpBlockType
		return ret
	}).(IpBlockTypeOutput)
}

func (o IpBlockTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpBlockTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpBlockType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IpBlockTypeInput is an input type that accepts IpBlockTypeArgs and IpBlockTypeOutput values.
// You can construct a concrete instance of `IpBlockTypeInput` via:
//
//          IpBlockTypeArgs{...}
type IpBlockTypeInput interface {
	pulumi.Input

	ToIpBlockTypeOutput() IpBlockTypeOutput
	ToIpBlockTypeOutputWithContext(context.Context) IpBlockTypeOutput
}

var ipBlockTypePtrType = reflect.TypeOf((**IpBlockType)(nil)).Elem()

type IpBlockTypePtrInput interface {
	pulumi.Input

	ToIpBlockTypePtrOutput() IpBlockTypePtrOutput
	ToIpBlockTypePtrOutputWithContext(context.Context) IpBlockTypePtrOutput
}

type ipBlockTypePtr string

func IpBlockTypePtr(v string) IpBlockTypePtrInput {
	return (*ipBlockTypePtr)(&v)
}

func (*ipBlockTypePtr) ElementType() reflect.Type {
	return ipBlockTypePtrType
}

func (in *ipBlockTypePtr) ToIpBlockTypePtrOutput() IpBlockTypePtrOutput {
	return pulumi.ToOutput(in).(IpBlockTypePtrOutput)
}

func (in *ipBlockTypePtr) ToIpBlockTypePtrOutputWithContext(ctx context.Context) IpBlockTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpBlockTypePtrOutput)
}

type NetworkType string

const (
	NetworkTypeLayer3           = NetworkType("layer3")
	NetworkTypeLayer2Individual = NetworkType("layer2-individual")
	NetworkTypeLayer2Bonded     = NetworkType("layer2-bonded")
	NetworkTypeHybrid           = NetworkType("hybrid")
)

func (NetworkType) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkType)(nil)).Elem()
}

func (e NetworkType) ToNetworkTypeOutput() NetworkTypeOutput {
	return pulumi.ToOutput(e).(NetworkTypeOutput)
}

func (e NetworkType) ToNetworkTypeOutputWithContext(ctx context.Context) NetworkTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkTypeOutput)
}

func (e NetworkType) ToNetworkTypePtrOutput() NetworkTypePtrOutput {
	return e.ToNetworkTypePtrOutputWithContext(context.Background())
}

func (e NetworkType) ToNetworkTypePtrOutputWithContext(ctx context.Context) NetworkTypePtrOutput {
	return NetworkType(e).ToNetworkTypeOutputWithContext(ctx).ToNetworkTypePtrOutputWithContext(ctx)
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

type NetworkTypeOutput struct{ *pulumi.OutputState }

func (NetworkTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkType)(nil)).Elem()
}

func (o NetworkTypeOutput) ToNetworkTypeOutput() NetworkTypeOutput {
	return o
}

func (o NetworkTypeOutput) ToNetworkTypeOutputWithContext(ctx context.Context) NetworkTypeOutput {
	return o
}

func (o NetworkTypeOutput) ToNetworkTypePtrOutput() NetworkTypePtrOutput {
	return o.ToNetworkTypePtrOutputWithContext(context.Background())
}

func (o NetworkTypeOutput) ToNetworkTypePtrOutputWithContext(ctx context.Context) NetworkTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkType) *NetworkType {
		return &v
	}).(NetworkTypePtrOutput)
}

func (o NetworkTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkTypePtrOutput struct{ *pulumi.OutputState }

func (NetworkTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkType)(nil)).Elem()
}

func (o NetworkTypePtrOutput) ToNetworkTypePtrOutput() NetworkTypePtrOutput {
	return o
}

func (o NetworkTypePtrOutput) ToNetworkTypePtrOutputWithContext(ctx context.Context) NetworkTypePtrOutput {
	return o
}

func (o NetworkTypePtrOutput) Elem() NetworkTypeOutput {
	return o.ApplyT(func(v *NetworkType) NetworkType {
		if v != nil {
			return *v
		}
		var ret NetworkType
		return ret
	}).(NetworkTypeOutput)
}

func (o NetworkTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NetworkTypeInput is an input type that accepts NetworkTypeArgs and NetworkTypeOutput values.
// You can construct a concrete instance of `NetworkTypeInput` via:
//
//          NetworkTypeArgs{...}
type NetworkTypeInput interface {
	pulumi.Input

	ToNetworkTypeOutput() NetworkTypeOutput
	ToNetworkTypeOutputWithContext(context.Context) NetworkTypeOutput
}

var networkTypePtrType = reflect.TypeOf((**NetworkType)(nil)).Elem()

type NetworkTypePtrInput interface {
	pulumi.Input

	ToNetworkTypePtrOutput() NetworkTypePtrOutput
	ToNetworkTypePtrOutputWithContext(context.Context) NetworkTypePtrOutput
}

type networkTypePtr string

func NetworkTypePtr(v string) NetworkTypePtrInput {
	return (*networkTypePtr)(&v)
}

func (*networkTypePtr) ElementType() reflect.Type {
	return networkTypePtrType
}

func (in *networkTypePtr) ToNetworkTypePtrOutput() NetworkTypePtrOutput {
	return pulumi.ToOutput(in).(NetworkTypePtrOutput)
}

func (in *networkTypePtr) ToNetworkTypePtrOutputWithContext(ctx context.Context) NetworkTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkTypePtrOutput)
}

type OperatingSystem string

const (
	OperatingSystemAlpine3          = OperatingSystem("alpine_3")
	OperatingSystemCentOS6          = OperatingSystem("centos_6")
	OperatingSystemCentOS7          = OperatingSystem("centos_7")
	OperatingSystemCentOS8          = OperatingSystem("centos_8")
	OperatingSystemCoreOSAlpha      = OperatingSystem("coreos_alpha")
	OperatingSystemCoreOSBeta       = OperatingSystem("coreos_beta")
	OperatingSystemCoreOSStable     = OperatingSystem("coreos_stable")
	OperatingSystemCustomIPXE       = OperatingSystem("custom_ipxe")
	OperatingSystemDebian8          = OperatingSystem("debian_8")
	OperatingSystemDebian9          = OperatingSystem("debian_9")
	OperatingSystemDebian10         = OperatingSystem("debian_10")
	OperatingSystemFlatcarAlpha     = OperatingSystem("flatcar_alpha")
	OperatingSystemFlatcarBeta      = OperatingSystem("flatcar_beta")
	OperatingSystemFlatcarEdge      = OperatingSystem("flatcar_edge")
	OperatingSystemFlatcarStable    = OperatingSystem("flatcar_stable")
	OperatingSystem_FreeBSD10_4     = OperatingSystem("freebsd_10_4")
	OperatingSystem_FreeBSD11_1     = OperatingSystem("freebsd_11_1")
	OperatingSystem_FreeBSD11_2     = OperatingSystem("freebsd_11_2")
	OperatingSystemFreeBSD12Testing = OperatingSystem("freebsd_12_testing")
	OperatingSystem_NixOS18_03      = OperatingSystem("nixos_18_03")
	OperatingSystem_NixOS19_03      = OperatingSystem("nixos_19_03")
	OperatingSystem_OpenSUSE42_3    = OperatingSystem("opensuse_42_3")
	OperatingSystemRancherOS        = OperatingSystem("rancher")
	OperatingSystemRHEL7            = OperatingSystem("rhel_7")
	OperatingSystemRHEL8            = OperatingSystem("rhel_8")
	OperatingSystemScientificLinux6 = OperatingSystem("scientific_6")
	OperatingSystemSLES12SP3        = OperatingSystem("suse_sles12_sp3")
	OperatingSystemUbuntu1404       = OperatingSystem("ubuntu_14_04")
	OperatingSystemUbuntu1604       = OperatingSystem("ubuntu_16_04")
	OperatingSystemUbuntu1710       = OperatingSystem("ubuntu_17_10")
	OperatingSystemUbuntu1804       = OperatingSystem("ubuntu_18_04")
	OperatingSystemUbuntu2004       = OperatingSystem("ubuntu_20_04")
	OperatingSystemUbuntu2010       = OperatingSystem("ubuntu_20_10")
	OperatingSystem_VMWareEsxi6_5   = OperatingSystem("vmware_esxi_6_5")
	OperatingSystem_VMWareEsxi6_7   = OperatingSystem("vmware_esxi_6_7")
	OperatingSystem_VMWareEsxi7_0   = OperatingSystem("vmware_esxi_7_0")
	OperatingSystemWindows2012R2    = OperatingSystem("windows_2012_r2")
	OperatingSystemWindows2016      = OperatingSystem("windows_2016")
	OperatingSystemWindows2019      = OperatingSystem("windows_2019")
)

func (OperatingSystem) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystem)(nil)).Elem()
}

func (e OperatingSystem) ToOperatingSystemOutput() OperatingSystemOutput {
	return pulumi.ToOutput(e).(OperatingSystemOutput)
}

func (e OperatingSystem) ToOperatingSystemOutputWithContext(ctx context.Context) OperatingSystemOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemOutput)
}

func (e OperatingSystem) ToOperatingSystemPtrOutput() OperatingSystemPtrOutput {
	return e.ToOperatingSystemPtrOutputWithContext(context.Background())
}

func (e OperatingSystem) ToOperatingSystemPtrOutputWithContext(ctx context.Context) OperatingSystemPtrOutput {
	return OperatingSystem(e).ToOperatingSystemOutputWithContext(ctx).ToOperatingSystemPtrOutputWithContext(ctx)
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

type OperatingSystemOutput struct{ *pulumi.OutputState }

func (OperatingSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystem)(nil)).Elem()
}

func (o OperatingSystemOutput) ToOperatingSystemOutput() OperatingSystemOutput {
	return o
}

func (o OperatingSystemOutput) ToOperatingSystemOutputWithContext(ctx context.Context) OperatingSystemOutput {
	return o
}

func (o OperatingSystemOutput) ToOperatingSystemPtrOutput() OperatingSystemPtrOutput {
	return o.ToOperatingSystemPtrOutputWithContext(context.Background())
}

func (o OperatingSystemOutput) ToOperatingSystemPtrOutputWithContext(ctx context.Context) OperatingSystemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystem) *OperatingSystem {
		return &v
	}).(OperatingSystemPtrOutput)
}

func (o OperatingSystemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemPtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystem)(nil)).Elem()
}

func (o OperatingSystemPtrOutput) ToOperatingSystemPtrOutput() OperatingSystemPtrOutput {
	return o
}

func (o OperatingSystemPtrOutput) ToOperatingSystemPtrOutputWithContext(ctx context.Context) OperatingSystemPtrOutput {
	return o
}

func (o OperatingSystemPtrOutput) Elem() OperatingSystemOutput {
	return o.ApplyT(func(v *OperatingSystem) OperatingSystem {
		if v != nil {
			return *v
		}
		var ret OperatingSystem
		return ret
	}).(OperatingSystemOutput)
}

func (o OperatingSystemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// OperatingSystemInput is an input type that accepts OperatingSystemArgs and OperatingSystemOutput values.
// You can construct a concrete instance of `OperatingSystemInput` via:
//
//          OperatingSystemArgs{...}
type OperatingSystemInput interface {
	pulumi.Input

	ToOperatingSystemOutput() OperatingSystemOutput
	ToOperatingSystemOutputWithContext(context.Context) OperatingSystemOutput
}

var operatingSystemPtrType = reflect.TypeOf((**OperatingSystem)(nil)).Elem()

type OperatingSystemPtrInput interface {
	pulumi.Input

	ToOperatingSystemPtrOutput() OperatingSystemPtrOutput
	ToOperatingSystemPtrOutputWithContext(context.Context) OperatingSystemPtrOutput
}

type operatingSystemPtr string

func OperatingSystemPtr(v string) OperatingSystemPtrInput {
	return (*operatingSystemPtr)(&v)
}

func (*operatingSystemPtr) ElementType() reflect.Type {
	return operatingSystemPtrType
}

func (in *operatingSystemPtr) ToOperatingSystemPtrOutput() OperatingSystemPtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemPtrOutput)
}

func (in *operatingSystemPtr) ToOperatingSystemPtrOutputWithContext(ctx context.Context) OperatingSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemPtrOutput)
}

type Plan string

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
	PlanC3SmallX86  = Plan("c3.small.x86")
	PlanC3MediumX86 = Plan("c3.medium.x86")
	PlanF3MediumX86 = Plan("f3.medium.x86")
	PlanF3LargeX86  = Plan("f3.large.x86")
	PlanM3LargeX86  = Plan("m3.large.x86")
	PlanS3XLargeX86 = Plan("s3.xlarge.x86")
	PlanN2XLargeX86 = Plan("n2.xlarge.x86")
)

func (Plan) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (e Plan) ToPlanOutput() PlanOutput {
	return pulumi.ToOutput(e).(PlanOutput)
}

func (e Plan) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PlanOutput)
}

func (e Plan) ToPlanPtrOutput() PlanPtrOutput {
	return e.ToPlanPtrOutputWithContext(context.Background())
}

func (e Plan) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return Plan(e).ToPlanOutputWithContext(ctx).ToPlanPtrOutputWithContext(ctx)
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

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PlanOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Plan) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PlanOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Plan) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PlanPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Plan) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PlanInput is an input type that accepts PlanArgs and PlanOutput values.
// You can construct a concrete instance of `PlanInput` via:
//
//          PlanArgs{...}
type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

var planPtrType = reflect.TypeOf((**Plan)(nil)).Elem()

type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtr string

func PlanPtr(v string) PlanPtrInput {
	return (*planPtr)(&v)
}

func (*planPtr) ElementType() reflect.Type {
	return planPtrType
}

func (in *planPtr) ToPlanPtrOutput() PlanPtrOutput {
	return pulumi.ToOutput(in).(PlanPtrOutput)
}

func (in *planPtr) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PlanPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingCycleInput)(nil)).Elem(), BillingCycle("hourly"))
	pulumi.RegisterInputType(reflect.TypeOf((*BillingCyclePtrInput)(nil)).Elem(), BillingCycle("hourly"))
	pulumi.RegisterInputType(reflect.TypeOf((*FacilityInput)(nil)).Elem(), Facility("ewr1"))
	pulumi.RegisterInputType(reflect.TypeOf((*FacilityPtrInput)(nil)).Elem(), Facility("ewr1"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpBlockTypeInput)(nil)).Elem(), IpBlockType("global_ipv4"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpBlockTypePtrInput)(nil)).Elem(), IpBlockType("global_ipv4"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkTypeInput)(nil)).Elem(), NetworkType("layer3"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkTypePtrInput)(nil)).Elem(), NetworkType("layer3"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemInput)(nil)).Elem(), OperatingSystem("alpine_3"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemPtrInput)(nil)).Elem(), OperatingSystem("alpine_3"))
	pulumi.RegisterInputType(reflect.TypeOf((*PlanInput)(nil)).Elem(), Plan("c2.large.arm"))
	pulumi.RegisterInputType(reflect.TypeOf((*PlanPtrInput)(nil)).Elem(), Plan("c2.large.arm"))
	pulumi.RegisterOutputType(BillingCycleOutput{})
	pulumi.RegisterOutputType(BillingCyclePtrOutput{})
	pulumi.RegisterOutputType(FacilityOutput{})
	pulumi.RegisterOutputType(FacilityPtrOutput{})
	pulumi.RegisterOutputType(IpBlockTypeOutput{})
	pulumi.RegisterOutputType(IpBlockTypePtrOutput{})
	pulumi.RegisterOutputType(NetworkTypeOutput{})
	pulumi.RegisterOutputType(NetworkTypePtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemOutput{})
	pulumi.RegisterOutputType(OperatingSystemPtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
}
