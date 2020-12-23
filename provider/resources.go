// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package equinix

import (
	"strings"
	"unicode"

	"github.com/equinix/terraform-provider-metal/metal"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	pulumiSchema "github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "equinix-metal"
	// modules:
	mainMod = "index" // the y module
)

var namespaceMap = map[string]string{
	"equinix-metal": "EquinixMetal",
}

// makeMember manufactures a type token for the package and the given module, file name, and type.
func makeMember(moduleTitle string, fn string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	if fn != "" {
		moduleName += "/" + fn
	}
	return tokens.ModuleMember(mainPkg + ":" + moduleName + ":" + mem)
}

// makeType manufactures a type token for the package and the given module, file name, and type.
func makeType(mod string, fn string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, fn, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It automatically uses the
// package and names the file by simply lower casing the data source's first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod, fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod, fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(metal.Provider().(*schema.Provider))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "metal",
		Description: "A Pulumi package for creating and managing equinix-metal cloud resources.",
		Keywords:    []string{"pulumi", "equinix-metal"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		GitHubOrg:   "equinix",
		Repository:  "https://github.com/pulumi/pulumi-equinix-metal",
		Config: map[string]*tfbridge.SchemaInfo{
			"auth_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PACKET_AUTH_TOKEN"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"metal_bgp_session": {Tok: makeResource(mainMod, "BgpSession")},
			"metal_device": {
				Tok: makeResource(mainMod, "Device"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"billing_cycle": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "BillingCycle", "BillingCycle")},
					},
					"facilities": {
						Elem: &tfbridge.SchemaInfo{
							Type:     "string",
							AltTypes: []tokens.Type{makeType(mainMod, "Facility", "Facility")},
						},
					},
					"network_type": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "NetworkType", "NetworkType")},
					},
					"operating_system": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "OperatingSystem", "OperatingSystem")},
					},
					"plan": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "Plan", "Plan")},
					},
				},
			},
			"metal_device_network_type":  {Tok: makeResource(mainMod, "DeviceNetworkType")},
			"metal_ip_attachment":        {Tok: makeResource(mainMod, "IpAttachment")},
			"metal_organization":         {Tok: makeResource(mainMod, "Organization")},
			"metal_port_vlan_attachment": {Tok: makeResource(mainMod, "PortVlanAttachment")},
			"metal_project":              {Tok: makeResource(mainMod, "Project")},
			"metal_project_ssh_key":      {Tok: makeResource(mainMod, "ProjectSshKey")},
			"metal_reserved_ip_block": {
				Tok: makeResource(mainMod, "ReservedIpBlock"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"facility": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "Facility", "Facility")},
					},
					"type": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "IpBlockType", "IpBlockType")},
					},
				},
			},
			"metal_spot_market_request": {Tok: makeResource(mainMod, "SpotMarketRequest")},
			"metal_ssh_key":             {Tok: makeResource(mainMod, "SshKey")},
			"metal_vlan": {
				Tok: makeResource(mainMod, "Vlan"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"facility": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "Facility", "Facility")},
					},
				},
			},
			"metal_volume": {
				Tok: makeResource(mainMod, "Volume"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"facility": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "Facility", "Facility")},
					},
					"billing_cycle": {
						Type:     "string",
						AltTypes: []tokens.Type{makeType(mainMod, "BillingCycle", "BillingCycle")},
					},
				},
			},
			"metal_volume_attachment": {Tok: makeResource(mainMod, "VolumeAttachment")},
		},
		ExtraTypes: map[string]pulumiSchema.ComplexTypeSpec{
			"equinix-metal:index/BillingCycle:BillingCycle": {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []*pulumiSchema.EnumValueSpec{
					{Value: "hourly", Name: "Hourly"},
					{Value: "monthly", Name: "Monthly"},
				},
			},
			"equinix-metal:index/IpBlockType:IpBlockType": {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []*pulumiSchema.EnumValueSpec{
					{Value: "global_ipv4", Name: "GlobalIPv4"},
					{Value: "public_ipv4", Name: "PublicIPv4"},
				},
			},
			"equinix-metal:index/NetworkType:NetworkType": {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []*pulumiSchema.EnumValueSpec{
					{Value: "layer3", Name: "Layer3"},
					{Value: "layer2-individual", Name: "Layer2Individual"},
					{Value: "layer2-bonded", Name: "Layer2Bonded"},
					{Value: "hybrid", Name: "Hybrid"},
				},
			},
			"equinix-metal:index/Plan:Plan": {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []*pulumiSchema.EnumValueSpec{
					{Name: "C2LargeARM", Value: "c2.large.arm"},
					{Name: "C2MediumX86", Value: "c2.medium.x86"},
					{Name: "C1SmallX86", Value: "baremetal_1"},
					{Name: "C1LargeARM", Value: "baremetal_2a"},
					{Name: "C1XLargeX86", Value: "baremetal_3"},
					{Name: "X2XLargeX86", Value: "x2.xlarge.x86"},
					{Name: "X1SmallX86", Value: "baremetal_1e"},
					{Name: "G2LargeX86", Value: "g2.large.x86"},
					{Name: "M2XLargeX86", Value: "m2.xlarge.x86"},
					{Name: "M1XLargeX86", Value: "baremetal_2"},
					{Name: "T1SmallX86", Value: "baremetal_0"},
					{Name: "S1LargeX86", Value: "baremetal_s"},
				},
			},
			"equinix-metal:index/OperatingSystem:OperatingSystem": {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []*pulumiSchema.EnumValueSpec{
					{Name: "Alpine3", Value: "alpine_3"},
					{Name: "CentOS6", Value: "centos_6"},
					{Name: "CentOS7", Value: "centos_7"},
					{Name: "CoreOSAlpha", Value: "coreos_alpha"},
					{Name: "CoreOSBeta", Value: "coreos_beta"},
					{Name: "CoreOSStable", Value: "coreos_stable"},
					{Name: "CustomIPXE", Value: "custom_ipxe"},
					{Name: "Debian8", Value: "debian_8"},
					{Name: "Debian9", Value: "debian_9"},
					{Name: "FreeBSD10_4", Value: "freebsd_10_4"},
					{Name: "FreeBSD11_1", Value: "freebsd_11_1"},
					{Name: "FreeBSD11_2", Value: "freebsd_11_2"},
					{Name: "FreeBSD12Testing", Value: "freebsd_12_testing"},
					{Name: "NixOS18_03", Value: "nixos_18_03"},
					{Name: "NixOS19_03", Value: "nixos_19_03"},
					{Name: "OpenSUSE42_3", Value: "opensuse_42_3"},
					{Name: "RancherOS", Value: "rancher"},
					{Name: "RHEL7", Value: "rhel_7"},
					{Name: "ScientificLinux6", Value: "scientific_6"},
					{Name: "SLES12SP3", Value: "suse_sles12_sp3"},
					{Name: "Ubuntu1404", Value: "ubuntu_14_04"},
					{Name: "Ubuntu1604", Value: "ubuntu_16_04"},
					{Name: "Ubuntu1710", Value: "ubuntu_17_10"},
					{Name: "Ubuntu1804", Value: "ubuntu_18_04"},
					{Name: "VMWareEsxi6_5", Value: "vmware_esxi_6_5"},
					{Name: "Windows2012R2", Value: "windows_2012_r2"},
					{Name: "Windows2016", Value: "windows_2016"},
				},
			},
			"equinix-metal:index/Facility:Facility": {
				ObjectTypeSpec: pulumiSchema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []*pulumiSchema.EnumValueSpec{
					{Value: "ewr1", Name: "EWR1"},
					{Value: "sjc1", Name: "SJC1"},
					{Value: "dfw1", Name: "DFW1"},
					{Value: "dfw2", Name: "DFW2"},
					{Value: "ams1", Name: "AMS1"},
					{Value: "nrt1", Name: "NRT1"},
					{Value: "sea1", Name: "SEA1"},
					{Value: "lax1", Name: "LAX1"},
					{Value: "ord1", Name: "ORD1"},
					{Value: "atl1", Name: "ATL1"},
					{Value: "iad1", Name: "IAD1"},
					{Value: "sin1", Name: "SIN1"},
					{Value: "hkg1", Name: "HKG1"},
					{Value: "syd1", Name: "SYD1"},
					{Value: "mrs1", Name: "MRS1"},
					{Value: "yyz1", Name: "YYZ1"},
					{Value: "fra2", Name: "FRA2"},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"metal_device":               {Tok: makeDataSource(mainMod, "getDevice")},
			"metal_device_bgp_neighbors": {Tok: makeDataSource(mainMod, "getDeviceBgpNeighbors")},
			"metal_ip_block_ranges":      {Tok: makeDataSource(mainMod, "getIpBlockRanges")},
			"metal_operating_system":     {Tok: makeDataSource(mainMod, "getOperatingSystem")},
			"metal_organization":         {Tok: makeDataSource(mainMod, "getOrganization")},
			"metal_precreated_ip_block":  {Tok: makeDataSource(mainMod, "getPrecreatedIpBlock")},
			"metal_project":              {Tok: makeDataSource(mainMod, "getProject")},
			"metal_project_ssh_key":      {Tok: makeDataSource(mainMod, "getProjectSshKey")},
			"metal_spot_market_price":    {Tok: makeDataSource(mainMod, "getSpotMarketPrice")},
			"metal_spot_market_request":  {Tok: makeDataSource(mainMod, "getSpotMarketRequest")},
			"metal_volume":               {Tok: makeDataSource(mainMod, "getVolume")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.15.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.15.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: namespaceMap,
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
