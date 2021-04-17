# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IpAttachmentArgs', 'IpAttachment']

@pulumi.input_type
class IpAttachmentArgs:
    def __init__(__self__, *,
                 cidr_notation: pulumi.Input[str],
                 device_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a IpAttachment resource.
        :param pulumi.Input[str] cidr_notation: CIDR notation of subnet from block reserved in the same
               project and facility as the device
        :param pulumi.Input[str] device_id: ID of device to which to assign the subnet
        """
        pulumi.set(__self__, "cidr_notation", cidr_notation)
        pulumi.set(__self__, "device_id", device_id)

    @property
    @pulumi.getter(name="cidrNotation")
    def cidr_notation(self) -> pulumi.Input[str]:
        """
        CIDR notation of subnet from block reserved in the same
        project and facility as the device
        """
        return pulumi.get(self, "cidr_notation")

    @cidr_notation.setter
    def cidr_notation(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr_notation", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Input[str]:
        """
        ID of device to which to assign the subnet
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_id", value)


@pulumi.input_type
class _IpAttachmentState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 address_family: Optional[pulumi.Input[int]] = None,
                 cidr: Optional[pulumi.Input[int]] = None,
                 cidr_notation: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 global_: Optional[pulumi.Input[bool]] = None,
                 manageable: Optional[pulumi.Input[bool]] = None,
                 management: Optional[pulumi.Input[bool]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering IpAttachment resources.
        :param pulumi.Input[int] address_family: Address family as integer (4 or 6)
        :param pulumi.Input[int] cidr: length of CIDR prefix of the subnet as integer
        :param pulumi.Input[str] cidr_notation: CIDR notation of subnet from block reserved in the same
               project and facility as the device
        :param pulumi.Input[str] device_id: ID of device to which to assign the subnet
        :param pulumi.Input[str] gateway: IP address of gateway for the subnet
        :param pulumi.Input[str] netmask: Subnet mask in decimal notation, e.g. "255.255.255.0"
        :param pulumi.Input[str] network: Subnet network address
        :param pulumi.Input[bool] public: boolean flag whether subnet is reachable from the Internet
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if cidr_notation is not None:
            pulumi.set(__self__, "cidr_notation", cidr_notation)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if global_ is not None:
            pulumi.set(__self__, "global_", global_)
        if manageable is not None:
            pulumi.set(__self__, "manageable", manageable)
        if management is not None:
            pulumi.set(__self__, "management", management)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if public is not None:
            pulumi.set(__self__, "public", public)

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
        length of CIDR prefix of the subnet as integer
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="cidrNotation")
    def cidr_notation(self) -> Optional[pulumi.Input[str]]:
        """
        CIDR notation of subnet from block reserved in the same
        project and facility as the device
        """
        return pulumi.get(self, "cidr_notation")

    @cidr_notation.setter
    def cidr_notation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_notation", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of device to which to assign the subnet
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of gateway for the subnet
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="global")
    def global_(self) -> Optional[pulumi.Input[bool]]:
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
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet mask in decimal notation, e.g. "255.255.255.0"
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet network address
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        boolean flag whether subnet is reachable from the Internet
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)


class IpAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_notation: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to attach elastic IP subnets to devices.

        To attach an IP subnet from a reserved block to a provisioned device, you must derive a subnet CIDR belonging to
        one of your reserved blocks in the same project and facility as the target device.

        For example, you have reserved IPv4 address block 147.229.10.152/30, you can choose to assign either the whole
        block as one subnet to a device; or 2 subnets with CIDRs 147.229.10.152/31' and 147.229.10.154/31; or 4 subnets
        with mask prefix length 32. More about the elastic IP subnets is [here](https://metal.equinix.com/developers/docs/networking/elastic-ips/).

        Device and reserved block must be in the same facility.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_notation: CIDR notation of subnet from block reserved in the same
               project and facility as the device
        :param pulumi.Input[str] device_id: ID of device to which to assign the subnet
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to attach elastic IP subnets to devices.

        To attach an IP subnet from a reserved block to a provisioned device, you must derive a subnet CIDR belonging to
        one of your reserved blocks in the same project and facility as the target device.

        For example, you have reserved IPv4 address block 147.229.10.152/30, you can choose to assign either the whole
        block as one subnet to a device; or 2 subnets with CIDRs 147.229.10.152/31' and 147.229.10.154/31; or 4 subnets
        with mask prefix length 32. More about the elastic IP subnets is [here](https://metal.equinix.com/developers/docs/networking/elastic-ips/).

        Device and reserved block must be in the same facility.

        :param str resource_name: The name of the resource.
        :param IpAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_notation: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpAttachmentArgs.__new__(IpAttachmentArgs)

            if cidr_notation is None and not opts.urn:
                raise TypeError("Missing required property 'cidr_notation'")
            __props__.__dict__["cidr_notation"] = cidr_notation
            if device_id is None and not opts.urn:
                raise TypeError("Missing required property 'device_id'")
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["address"] = None
            __props__.__dict__["address_family"] = None
            __props__.__dict__["cidr"] = None
            __props__.__dict__["gateway"] = None
            __props__.__dict__["global_"] = None
            __props__.__dict__["manageable"] = None
            __props__.__dict__["management"] = None
            __props__.__dict__["netmask"] = None
            __props__.__dict__["network"] = None
            __props__.__dict__["public"] = None
        super(IpAttachment, __self__).__init__(
            'equinix-metal:index/ipAttachment:IpAttachment',
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
            device_id: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            global_: Optional[pulumi.Input[bool]] = None,
            manageable: Optional[pulumi.Input[bool]] = None,
            management: Optional[pulumi.Input[bool]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            public: Optional[pulumi.Input[bool]] = None) -> 'IpAttachment':
        """
        Get an existing IpAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] address_family: Address family as integer (4 or 6)
        :param pulumi.Input[int] cidr: length of CIDR prefix of the subnet as integer
        :param pulumi.Input[str] cidr_notation: CIDR notation of subnet from block reserved in the same
               project and facility as the device
        :param pulumi.Input[str] device_id: ID of device to which to assign the subnet
        :param pulumi.Input[str] gateway: IP address of gateway for the subnet
        :param pulumi.Input[str] netmask: Subnet mask in decimal notation, e.g. "255.255.255.0"
        :param pulumi.Input[str] network: Subnet network address
        :param pulumi.Input[bool] public: boolean flag whether subnet is reachable from the Internet
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpAttachmentState.__new__(_IpAttachmentState)

        __props__.__dict__["address"] = address
        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["cidr_notation"] = cidr_notation
        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["global_"] = global_
        __props__.__dict__["manageable"] = manageable
        __props__.__dict__["management"] = management
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["network"] = network
        __props__.__dict__["public"] = public
        return IpAttachment(resource_name, opts=opts, __props__=__props__)

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
        length of CIDR prefix of the subnet as integer
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="cidrNotation")
    def cidr_notation(self) -> pulumi.Output[str]:
        """
        CIDR notation of subnet from block reserved in the same
        project and facility as the device
        """
        return pulumi.get(self, "cidr_notation")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        ID of device to which to assign the subnet
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        """
        IP address of gateway for the subnet
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="global")
    def global_(self) -> pulumi.Output[bool]:
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
    def netmask(self) -> pulumi.Output[str]:
        """
        Subnet mask in decimal notation, e.g. "255.255.255.0"
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        Subnet network address
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def public(self) -> pulumi.Output[bool]:
        """
        boolean flag whether subnet is reachable from the Internet
        """
        return pulumi.get(self, "public")

