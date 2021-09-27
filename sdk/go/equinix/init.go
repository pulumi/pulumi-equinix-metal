// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "equinix-metal:index/bgpSession:BgpSession":
		r = &BgpSession{}
	case "equinix-metal:index/connection:Connection":
		r = &Connection{}
	case "equinix-metal:index/device:Device":
		r = &Device{}
	case "equinix-metal:index/deviceNetworkType:DeviceNetworkType":
		r = &DeviceNetworkType{}
	case "equinix-metal:index/gateway:Gateway":
		r = &Gateway{}
	case "equinix-metal:index/ipAttachment:IpAttachment":
		r = &IpAttachment{}
	case "equinix-metal:index/organization:Organization":
		r = &Organization{}
	case "equinix-metal:index/portVlanAttachment:PortVlanAttachment":
		r = &PortVlanAttachment{}
	case "equinix-metal:index/project:Project":
		r = &Project{}
	case "equinix-metal:index/projectApiKey:ProjectApiKey":
		r = &ProjectApiKey{}
	case "equinix-metal:index/projectSshKey:ProjectSshKey":
		r = &ProjectSshKey{}
	case "equinix-metal:index/reservedIpBlock:ReservedIpBlock":
		r = &ReservedIpBlock{}
	case "equinix-metal:index/spotMarketRequest:SpotMarketRequest":
		r = &SpotMarketRequest{}
	case "equinix-metal:index/sshKey:SshKey":
		r = &SshKey{}
	case "equinix-metal:index/userApiKey:UserApiKey":
		r = &UserApiKey{}
	case "equinix-metal:index/virtualCircuit:VirtualCircuit":
		r = &VirtualCircuit{}
	case "equinix-metal:index/vlan:Vlan":
		r = &Vlan{}
	case "equinix-metal:index/volume:Volume":
		r = &Volume{}
	case "equinix-metal:index/volumeAttachment:VolumeAttachment":
		r = &VolumeAttachment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:equinix-metal" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/bgpSession",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/connection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/device",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/deviceNetworkType",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/gateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/ipAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/organization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/portVlanAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/projectApiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/projectSshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/reservedIpBlock",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/spotMarketRequest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/sshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/userApiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/virtualCircuit",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/vlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/volume",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix-metal",
		"index/volumeAttachment",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"equinix-metal",
		&pkg{version},
	)
}
