# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = ['ReservedIpBlockArgs', 'ReservedIpBlock']

@pulumi.input_type
class ReservedIpBlockArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 quantity: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[Union[str, 'Facility']]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[Union[str, 'IpBlockType']]] = None):
        """
        The set of arguments for constructing a ReservedIpBlock resource.
        :param pulumi.Input[str] project_id: The metal project ID where to allocate the address block
        :param pulumi.Input[int] quantity: The number of allocated /32 addresses, a power of 2
        :param pulumi.Input[str] description: Arbitrary description
        :param pulumi.Input[Union[str, 'Facility']] facility: Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
        :param pulumi.Input[str] metro: Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: String list of tags
        :param pulumi.Input[Union[str, 'IpBlockType']] type: Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility
        """
        ReservedIpBlockArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project_id=project_id,
            quantity=quantity,
            description=description,
            facility=facility,
            metro=metro,
            tags=tags,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project_id: pulumi.Input[str],
             quantity: pulumi.Input[int],
             description: Optional[pulumi.Input[str]] = None,
             facility: Optional[pulumi.Input[Union[str, 'Facility']]] = None,
             metro: Optional[pulumi.Input[str]] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             type: Optional[pulumi.Input[Union[str, 'IpBlockType']]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("project_id", project_id)
        _setter("quantity", quantity)
        if description is not None:
            _setter("description", description)
        if facility is not None:
            _setter("facility", facility)
        if metro is not None:
            _setter("metro", metro)
        if tags is not None:
            _setter("tags", tags)
        if type is not None:
            _setter("type", type)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The metal project ID where to allocate the address block
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def quantity(self) -> pulumi.Input[int]:
        """
        The number of allocated /32 addresses, a power of 2
        """
        return pulumi.get(self, "quantity")

    @quantity.setter
    def quantity(self, value: pulumi.Input[int]):
        pulumi.set(self, "quantity", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Arbitrary description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[Union[str, 'Facility']]]:
        """
        Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
        """
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[Union[str, 'Facility']]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def metro(self) -> Optional[pulumi.Input[str]]:
        """
        Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
        """
        return pulumi.get(self, "metro")

    @metro.setter
    def metro(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metro", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        String list of tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[str, 'IpBlockType']]]:
        """
        Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[str, 'IpBlockType']]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _ReservedIpBlockState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 address_family: Optional[pulumi.Input[int]] = None,
                 cidr: Optional[pulumi.Input[int]] = None,
                 cidr_notation: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[Union[str, 'Facility']]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 global_: Optional[pulumi.Input[bool]] = None,
                 manageable: Optional[pulumi.Input[bool]] = None,
                 management: Optional[pulumi.Input[bool]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 quantity: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[Union[str, 'IpBlockType']]] = None):
        """
        Input properties used for looking up and filtering ReservedIpBlock resources.
        :param pulumi.Input[int] address_family: Address family as integer (4 or 6)
        :param pulumi.Input[int] cidr: length of CIDR prefix of the block as integer
        :param pulumi.Input[str] cidr_notation: Address and mask in CIDR notation, e.g. "147.229.15.30/31"
        :param pulumi.Input[str] description: Arbitrary description
        :param pulumi.Input[Union[str, 'Facility']] facility: Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
        :param pulumi.Input[bool] global_: boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
        :param pulumi.Input[str] metro: Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
        :param pulumi.Input[str] netmask: Mask in decimal notation, e.g. "255.255.255.0"
        :param pulumi.Input[str] network: Network IP address portion of the block specification
        :param pulumi.Input[str] project_id: The metal project ID where to allocate the address block
        :param pulumi.Input[bool] public: boolean flag whether addresses from a block are public
        :param pulumi.Input[int] quantity: The number of allocated /32 addresses, a power of 2
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: String list of tags
        :param pulumi.Input[Union[str, 'IpBlockType']] type: Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility
        """
        _ReservedIpBlockState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address=address,
            address_family=address_family,
            cidr=cidr,
            cidr_notation=cidr_notation,
            description=description,
            facility=facility,
            gateway=gateway,
            global_=global_,
            manageable=manageable,
            management=management,
            metro=metro,
            netmask=netmask,
            network=network,
            project_id=project_id,
            public=public,
            quantity=quantity,
            tags=tags,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address: Optional[pulumi.Input[str]] = None,
             address_family: Optional[pulumi.Input[int]] = None,
             cidr: Optional[pulumi.Input[int]] = None,
             cidr_notation: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             facility: Optional[pulumi.Input[Union[str, 'Facility']]] = None,
             gateway: Optional[pulumi.Input[str]] = None,
             global_: Optional[pulumi.Input[bool]] = None,
             manageable: Optional[pulumi.Input[bool]] = None,
             management: Optional[pulumi.Input[bool]] = None,
             metro: Optional[pulumi.Input[str]] = None,
             netmask: Optional[pulumi.Input[str]] = None,
             network: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             public: Optional[pulumi.Input[bool]] = None,
             quantity: Optional[pulumi.Input[int]] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             type: Optional[pulumi.Input[Union[str, 'IpBlockType']]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if address is not None:
            _setter("address", address)
        if address_family is not None:
            _setter("address_family", address_family)
        if cidr is not None:
            _setter("cidr", cidr)
        if cidr_notation is not None:
            _setter("cidr_notation", cidr_notation)
        if description is not None:
            _setter("description", description)
        if facility is not None:
            _setter("facility", facility)
        if gateway is not None:
            _setter("gateway", gateway)
        if global_ is not None:
            _setter("global_", global_)
        if manageable is not None:
            _setter("manageable", manageable)
        if management is not None:
            _setter("management", management)
        if metro is not None:
            _setter("metro", metro)
        if netmask is not None:
            _setter("netmask", netmask)
        if network is not None:
            _setter("network", network)
        if project_id is not None:
            _setter("project_id", project_id)
        if public is not None:
            _setter("public", public)
        if quantity is not None:
            _setter("quantity", quantity)
        if tags is not None:
            _setter("tags", tags)
        if type is not None:
            _setter("type", type)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[int]]:
        """
        Address family as integer (4 or 6)
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[int]]:
        """
        length of CIDR prefix of the block as integer
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="cidrNotation")
    def cidr_notation(self) -> Optional[pulumi.Input[str]]:
        """
        Address and mask in CIDR notation, e.g. "147.229.15.30/31"
        """
        return pulumi.get(self, "cidr_notation")

    @cidr_notation.setter
    def cidr_notation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_notation", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Arbitrary description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[Union[str, 'Facility']]]:
        """
        Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
        """
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[Union[str, 'Facility']]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="global")
    def global_(self) -> Optional[pulumi.Input[bool]]:
        """
        boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
        """
        return pulumi.get(self, "global_")

    @global_.setter
    def global_(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "global_", value)

    @property
    @pulumi.getter
    def manageable(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "manageable")

    @manageable.setter
    def manageable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manageable", value)

    @property
    @pulumi.getter
    def management(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "management")

    @management.setter
    def management(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "management", value)

    @property
    @pulumi.getter
    def metro(self) -> Optional[pulumi.Input[str]]:
        """
        Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
        """
        return pulumi.get(self, "metro")

    @metro.setter
    def metro(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metro", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        Mask in decimal notation, e.g. "255.255.255.0"
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        Network IP address portion of the block specification
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The metal project ID where to allocate the address block
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        boolean flag whether addresses from a block are public
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)

    @property
    @pulumi.getter
    def quantity(self) -> Optional[pulumi.Input[int]]:
        """
        The number of allocated /32 addresses, a power of 2
        """
        return pulumi.get(self, "quantity")

    @quantity.setter
    def quantity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "quantity", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        String list of tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[str, 'IpBlockType']]]:
        """
        Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[str, 'IpBlockType']]]):
        pulumi.set(self, "type", value)


class ReservedIpBlock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[Union[str, 'Facility']]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 quantity: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[Union[str, 'IpBlockType']]] = None,
                 __props__=None):
        """
        Provides a resource to create and manage blocks of reserved IP addresses in a project.

        When a user provisions first device in a facility, Equinix Metal API automatically allocates IPv6/56 and private IPv4/25 blocks.
        The new device then gets IPv6 and private IPv4 addresses from those block. It also gets a public IPv4/31 address.
        Every new device in the project and facility will automatically get IPv6 and private IPv4 addresses from these pre-allocated blocks.
        The IPv6 and private IPv4 blocks can't be created, only imported. With this resource, it's possible to create either public IPv4 blocks or global IPv4 blocks.

        Public blocks are allocated in a facility. Addresses from public blocks can only be assigned to devices in the facility. Public blocks can have mask from /24 (256 addresses) to /32 (1 address). If you create public block with this resource, you must fill the facility argmument.

        Addresses from global blocks can be assigned in any facility. Global blocks can have mask from /30 (4 addresses), to /32 (1 address). If you create global block with this resource, you must specify type = "global_ipv4" and you must omit the facility argument.

        Once IP block is allocated or imported, an address from it can be assigned to device with the `IpAttachment` resource.

        ## Example Usage

        Allocate reserved IP blocks:

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        # Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility for myproject
        two_elastic_addresses = equinix_metal.ReservedIpBlock("twoElasticAddresses",
            project_id=local["project_id"],
            facility="sv15",
            quantity=2)
        # Allocate 1 floating IP in Sillicon Valley (sv) metro
        test_reserved_ip_block = equinix_metal.ReservedIpBlock("testReservedIpBlock",
            project_id=local["project_id"],
            type="public_ipv4",
            metro="sv",
            quantity=1)
        # Allocate 1 global floating IP, which can be assigned to device in any facility
        test_index_reserved_ip_block_reserved_ip_block = equinix_metal.ReservedIpBlock("testIndex/reservedIpBlockReservedIpBlock",
            project_id=local["project_id"],
            type="global_ipv4",
            quantity=1)
        ```

        Allocate a block and run a device with public IPv4 from the block

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        # Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility
        example = equinix_metal.ReservedIpBlock("example",
            project_id=local["project_id"],
            facility="sv15",
            quantity=2)
        # Run a device with both public IPv4 from the block assigned
        nodes = equinix_metal.Device("nodes",
            project_id=local["project_id"],
            facilities=["sv15"],
            plan="c3.small.x86",
            operating_system="ubuntu_20_04",
            hostname="test",
            billing_cycle="hourly",
            ip_addresses=[
                equinix_metal.DeviceIpAddressArgs(
                    type="public_ipv4",
                    cidr=31,
                    reservation_ids=[example.id],
                ),
                equinix_metal.DeviceIpAddressArgs(
                    type="private_ipv4",
                ),
            ])
        ```

        ## Import

        This resource can be imported using an existing IP reservation ID

        ```sh
         $ pulumi import equinix-metal:index/reservedIpBlock:ReservedIpBlock metal_reserved_ip_block {existing_ip_reservation_id}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Arbitrary description
        :param pulumi.Input[Union[str, 'Facility']] facility: Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
        :param pulumi.Input[str] metro: Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
        :param pulumi.Input[str] project_id: The metal project ID where to allocate the address block
        :param pulumi.Input[int] quantity: The number of allocated /32 addresses, a power of 2
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: String list of tags
        :param pulumi.Input[Union[str, 'IpBlockType']] type: Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReservedIpBlockArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create and manage blocks of reserved IP addresses in a project.

        When a user provisions first device in a facility, Equinix Metal API automatically allocates IPv6/56 and private IPv4/25 blocks.
        The new device then gets IPv6 and private IPv4 addresses from those block. It also gets a public IPv4/31 address.
        Every new device in the project and facility will automatically get IPv6 and private IPv4 addresses from these pre-allocated blocks.
        The IPv6 and private IPv4 blocks can't be created, only imported. With this resource, it's possible to create either public IPv4 blocks or global IPv4 blocks.

        Public blocks are allocated in a facility. Addresses from public blocks can only be assigned to devices in the facility. Public blocks can have mask from /24 (256 addresses) to /32 (1 address). If you create public block with this resource, you must fill the facility argmument.

        Addresses from global blocks can be assigned in any facility. Global blocks can have mask from /30 (4 addresses), to /32 (1 address). If you create global block with this resource, you must specify type = "global_ipv4" and you must omit the facility argument.

        Once IP block is allocated or imported, an address from it can be assigned to device with the `IpAttachment` resource.

        ## Example Usage

        Allocate reserved IP blocks:

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        # Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility for myproject
        two_elastic_addresses = equinix_metal.ReservedIpBlock("twoElasticAddresses",
            project_id=local["project_id"],
            facility="sv15",
            quantity=2)
        # Allocate 1 floating IP in Sillicon Valley (sv) metro
        test_reserved_ip_block = equinix_metal.ReservedIpBlock("testReservedIpBlock",
            project_id=local["project_id"],
            type="public_ipv4",
            metro="sv",
            quantity=1)
        # Allocate 1 global floating IP, which can be assigned to device in any facility
        test_index_reserved_ip_block_reserved_ip_block = equinix_metal.ReservedIpBlock("testIndex/reservedIpBlockReservedIpBlock",
            project_id=local["project_id"],
            type="global_ipv4",
            quantity=1)
        ```

        Allocate a block and run a device with public IPv4 from the block

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        # Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility
        example = equinix_metal.ReservedIpBlock("example",
            project_id=local["project_id"],
            facility="sv15",
            quantity=2)
        # Run a device with both public IPv4 from the block assigned
        nodes = equinix_metal.Device("nodes",
            project_id=local["project_id"],
            facilities=["sv15"],
            plan="c3.small.x86",
            operating_system="ubuntu_20_04",
            hostname="test",
            billing_cycle="hourly",
            ip_addresses=[
                equinix_metal.DeviceIpAddressArgs(
                    type="public_ipv4",
                    cidr=31,
                    reservation_ids=[example.id],
                ),
                equinix_metal.DeviceIpAddressArgs(
                    type="private_ipv4",
                ),
            ])
        ```

        ## Import

        This resource can be imported using an existing IP reservation ID

        ```sh
         $ pulumi import equinix-metal:index/reservedIpBlock:ReservedIpBlock metal_reserved_ip_block {existing_ip_reservation_id}
        ```

        :param str resource_name: The name of the resource.
        :param ReservedIpBlockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReservedIpBlockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ReservedIpBlockArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[Union[str, 'Facility']]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 quantity: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[Union[str, 'IpBlockType']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReservedIpBlockArgs.__new__(ReservedIpBlockArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["facility"] = facility
            __props__.__dict__["metro"] = metro
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if quantity is None and not opts.urn:
                raise TypeError("Missing required property 'quantity'")
            __props__.__dict__["quantity"] = quantity
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["address"] = None
            __props__.__dict__["address_family"] = None
            __props__.__dict__["cidr"] = None
            __props__.__dict__["cidr_notation"] = None
            __props__.__dict__["gateway"] = None
            __props__.__dict__["global_"] = None
            __props__.__dict__["manageable"] = None
            __props__.__dict__["management"] = None
            __props__.__dict__["netmask"] = None
            __props__.__dict__["network"] = None
            __props__.__dict__["public"] = None
        super(ReservedIpBlock, __self__).__init__(
            'equinix-metal:index/reservedIpBlock:ReservedIpBlock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            address_family: Optional[pulumi.Input[int]] = None,
            cidr: Optional[pulumi.Input[int]] = None,
            cidr_notation: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            facility: Optional[pulumi.Input[Union[str, 'Facility']]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            global_: Optional[pulumi.Input[bool]] = None,
            manageable: Optional[pulumi.Input[bool]] = None,
            management: Optional[pulumi.Input[bool]] = None,
            metro: Optional[pulumi.Input[str]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            public: Optional[pulumi.Input[bool]] = None,
            quantity: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[Union[str, 'IpBlockType']]] = None) -> 'ReservedIpBlock':
        """
        Get an existing ReservedIpBlock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] address_family: Address family as integer (4 or 6)
        :param pulumi.Input[int] cidr: length of CIDR prefix of the block as integer
        :param pulumi.Input[str] cidr_notation: Address and mask in CIDR notation, e.g. "147.229.15.30/31"
        :param pulumi.Input[str] description: Arbitrary description
        :param pulumi.Input[Union[str, 'Facility']] facility: Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
        :param pulumi.Input[bool] global_: boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
        :param pulumi.Input[str] metro: Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
        :param pulumi.Input[str] netmask: Mask in decimal notation, e.g. "255.255.255.0"
        :param pulumi.Input[str] network: Network IP address portion of the block specification
        :param pulumi.Input[str] project_id: The metal project ID where to allocate the address block
        :param pulumi.Input[bool] public: boolean flag whether addresses from a block are public
        :param pulumi.Input[int] quantity: The number of allocated /32 addresses, a power of 2
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: String list of tags
        :param pulumi.Input[Union[str, 'IpBlockType']] type: Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReservedIpBlockState.__new__(_ReservedIpBlockState)

        __props__.__dict__["address"] = address
        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["cidr_notation"] = cidr_notation
        __props__.__dict__["description"] = description
        __props__.__dict__["facility"] = facility
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["global_"] = global_
        __props__.__dict__["manageable"] = manageable
        __props__.__dict__["management"] = management
        __props__.__dict__["metro"] = metro
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["network"] = network
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["public"] = public
        __props__.__dict__["quantity"] = quantity
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        return ReservedIpBlock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[int]:
        """
        Address family as integer (4 or 6)
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[int]:
        """
        length of CIDR prefix of the block as integer
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="cidrNotation")
    def cidr_notation(self) -> pulumi.Output[str]:
        """
        Address and mask in CIDR notation, e.g. "147.229.15.30/31"
        """
        return pulumi.get(self, "cidr_notation")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Arbitrary description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def facility(self) -> pulumi.Output[Optional[str]]:
        """
        Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
        """
        return pulumi.get(self, "facility")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="global")
    def global_(self) -> pulumi.Output[bool]:
        """
        boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
        """
        return pulumi.get(self, "global_")

    @property
    @pulumi.getter
    def manageable(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "manageable")

    @property
    @pulumi.getter
    def management(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "management")

    @property
    @pulumi.getter
    def metro(self) -> pulumi.Output[Optional[str]]:
        """
        Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
        """
        return pulumi.get(self, "metro")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[str]:
        """
        Mask in decimal notation, e.g. "255.255.255.0"
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        Network IP address portion of the block specification
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The metal project ID where to allocate the address block
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def public(self) -> pulumi.Output[bool]:
        """
        boolean flag whether addresses from a block are public
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def quantity(self) -> pulumi.Output[int]:
        """
        The number of allocated /32 addresses, a power of 2
        """
        return pulumi.get(self, "quantity")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        String list of tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility
        """
        return pulumi.get(self, "type")

