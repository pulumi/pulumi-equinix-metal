# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .bgp_session import *
from .connection import *
from .device import *
from .device_network_type import *
from .gateway import *
from .get_connection import *
from .get_device import *
from .get_device_bgp_neighbors import *
from .get_facility import *
from .get_gateway import *
from .get_hardware_reservation import *
from .get_ip_block_ranges import *
from .get_metro import *
from .get_operating_system import *
from .get_organization import *
from .get_port import *
from .get_precreated_ip_block import *
from .get_project import *
from .get_project_ssh_key import *
from .get_reserved_ip_block import *
from .get_spot_market_price import *
from .get_spot_market_request import *
from .get_virtual_circuit import *
from .get_vlan import *
from .get_volume import *
from .ip_attachment import *
from .organization import *
from .port_vlan_attachment import *
from .project import *
from .project_api_key import *
from .project_ssh_key import *
from .provider import *
from .reserved_ip_block import *
from .spot_market_request import *
from .ssh_key import *
from .user_api_key import *
from .virtual_circuit import *
from .vlan import *
from .volume import *
from .volume_attachment import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    config,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "equinix-metal:index/bgpSession:BgpSession":
                return BgpSession(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/connection:Connection":
                return Connection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/device:Device":
                return Device(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/deviceNetworkType:DeviceNetworkType":
                return DeviceNetworkType(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/gateway:Gateway":
                return Gateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/ipAttachment:IpAttachment":
                return IpAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/organization:Organization":
                return Organization(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/portVlanAttachment:PortVlanAttachment":
                return PortVlanAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/project:Project":
                return Project(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/projectApiKey:ProjectApiKey":
                return ProjectApiKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/projectSshKey:ProjectSshKey":
                return ProjectSshKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/reservedIpBlock:ReservedIpBlock":
                return ReservedIpBlock(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/spotMarketRequest:SpotMarketRequest":
                return SpotMarketRequest(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/sshKey:SshKey":
                return SshKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/userApiKey:UserApiKey":
                return UserApiKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/virtualCircuit:VirtualCircuit":
                return VirtualCircuit(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/vlan:Vlan":
                return Vlan(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/volume:Volume":
                return Volume(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "equinix-metal:index/volumeAttachment:VolumeAttachment":
                return VolumeAttachment(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("equinix-metal", "index/bgpSession", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/connection", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/device", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/deviceNetworkType", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/gateway", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/ipAttachment", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/organization", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/portVlanAttachment", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/project", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/projectApiKey", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/projectSshKey", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/reservedIpBlock", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/spotMarketRequest", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/sshKey", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/userApiKey", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/virtualCircuit", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/vlan", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/volume", _module_instance)
    pulumi.runtime.register_resource_module("equinix-metal", "index/volumeAttachment", _module_instance)


    class Package(pulumi.runtime.ResourcePackage):
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:equinix-metal":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("equinix-metal", Package())

_register_module()
