# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PortVlanAttachmentArgs', 'PortVlanAttachment']

@pulumi.input_type
class PortVlanAttachmentArgs:
    def __init__(__self__, *,
                 device_id: pulumi.Input[str],
                 port_name: pulumi.Input[str],
                 vlan_vnid: pulumi.Input[int],
                 force_bond: Optional[pulumi.Input[bool]] = None,
                 native: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a PortVlanAttachment resource.
        :param pulumi.Input[str] device_id: ID of device to be assigned to the VLAN
        :param pulumi.Input[str] port_name: Name of network port to be assigned to the VLAN
        :param pulumi.Input[int] vlan_vnid: VXLAN Network Identifier, integer
        :param pulumi.Input[bool] force_bond: Add port back to the bond when this resource is removed. Default is false.
        :param pulumi.Input[bool] native: Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        """
        pulumi.set(__self__, "device_id", device_id)
        pulumi.set(__self__, "port_name", port_name)
        pulumi.set(__self__, "vlan_vnid", vlan_vnid)
        if force_bond is not None:
            pulumi.set(__self__, "force_bond", force_bond)
        if native is not None:
            pulumi.set(__self__, "native", native)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Input[str]:
        """
        ID of device to be assigned to the VLAN
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="portName")
    def port_name(self) -> pulumi.Input[str]:
        """
        Name of network port to be assigned to the VLAN
        """
        return pulumi.get(self, "port_name")

    @port_name.setter
    def port_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "port_name", value)

    @property
    @pulumi.getter(name="vlanVnid")
    def vlan_vnid(self) -> pulumi.Input[int]:
        """
        VXLAN Network Identifier, integer
        """
        return pulumi.get(self, "vlan_vnid")

    @vlan_vnid.setter
    def vlan_vnid(self, value: pulumi.Input[int]):
        pulumi.set(self, "vlan_vnid", value)

    @property
    @pulumi.getter(name="forceBond")
    def force_bond(self) -> Optional[pulumi.Input[bool]]:
        """
        Add port back to the bond when this resource is removed. Default is false.
        """
        return pulumi.get(self, "force_bond")

    @force_bond.setter
    def force_bond(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_bond", value)

    @property
    @pulumi.getter
    def native(self) -> Optional[pulumi.Input[bool]]:
        """
        Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        """
        return pulumi.get(self, "native")

    @native.setter
    def native(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "native", value)


@pulumi.input_type
class _PortVlanAttachmentState:
    def __init__(__self__, *,
                 device_id: Optional[pulumi.Input[str]] = None,
                 force_bond: Optional[pulumi.Input[bool]] = None,
                 native: Optional[pulumi.Input[bool]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 port_name: Optional[pulumi.Input[str]] = None,
                 vlan_id: Optional[pulumi.Input[str]] = None,
                 vlan_vnid: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering PortVlanAttachment resources.
        :param pulumi.Input[str] device_id: ID of device to be assigned to the VLAN
        :param pulumi.Input[bool] force_bond: Add port back to the bond when this resource is removed. Default is false.
        :param pulumi.Input[bool] native: Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        :param pulumi.Input[str] port_name: Name of network port to be assigned to the VLAN
        :param pulumi.Input[int] vlan_vnid: VXLAN Network Identifier, integer
        """
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if force_bond is not None:
            pulumi.set(__self__, "force_bond", force_bond)
        if native is not None:
            pulumi.set(__self__, "native", native)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if port_name is not None:
            pulumi.set(__self__, "port_name", port_name)
        if vlan_id is not None:
            pulumi.set(__self__, "vlan_id", vlan_id)
        if vlan_vnid is not None:
            pulumi.set(__self__, "vlan_vnid", vlan_vnid)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of device to be assigned to the VLAN
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="forceBond")
    def force_bond(self) -> Optional[pulumi.Input[bool]]:
        """
        Add port back to the bond when this resource is removed. Default is false.
        """
        return pulumi.get(self, "force_bond")

    @force_bond.setter
    def force_bond(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_bond", value)

    @property
    @pulumi.getter
    def native(self) -> Optional[pulumi.Input[bool]]:
        """
        Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        """
        return pulumi.get(self, "native")

    @native.setter
    def native(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "native", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="portName")
    def port_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of network port to be assigned to the VLAN
        """
        return pulumi.get(self, "port_name")

    @port_name.setter
    def port_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_name", value)

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_id", value)

    @property
    @pulumi.getter(name="vlanVnid")
    def vlan_vnid(self) -> Optional[pulumi.Input[int]]:
        """
        VXLAN Network Identifier, integer
        """
        return pulumi.get(self, "vlan_vnid")

    @vlan_vnid.setter
    def vlan_vnid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlan_vnid", value)


class PortVlanAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 force_bond: Optional[pulumi.Input[bool]] = None,
                 native: Optional[pulumi.Input[bool]] = None,
                 port_name: Optional[pulumi.Input[str]] = None,
                 vlan_vnid: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to attach device ports to VLANs.

        Device and VLAN must be in the same facility.

        If you need this resource to add the port back to bond on removal, set `force_bond = true`.

        To learn more about Layer 2 networking in Equinix Metal, refer to

        * <https://metal.equinix.com/developers/docs/networking/layer2/>
        * <https://metal.equinix.com/developers/docs/networking/layer2-configs/>

        ## Example Usage
        ### Hybrid network type

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test_vlan = equinix_metal.Vlan("testVlan",
            description="VLAN in New Jersey",
            facility="ny5",
            project_id=local["project_id"])
        test_device = equinix_metal.Device("testDevice",
            hostname="test",
            plan="c3.small.x86",
            facilities=["ny5"],
            operating_system="ubuntu_20_04",
            billing_cycle="hourly",
            project_id=local["project_id"])
        test_device_network_type = equinix_metal.DeviceNetworkType("testDeviceNetworkType",
            device_id=test_device.id,
            type="hybrid")
        test_port_vlan_attachment = equinix_metal.PortVlanAttachment("testPortVlanAttachment",
            device_id=test_device_network_type.id,
            port_name="eth1",
            vlan_vnid=test_vlan.vxlan)
        ```
        ### Layer 2 network

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test_device = equinix_metal.Device("testDevice",
            hostname="test",
            plan="c3.small.x86",
            facilities=["ny5"],
            operating_system="ubuntu_20_04",
            billing_cycle="hourly",
            project_id=local["project_id"])
        test_device_network_type = equinix_metal.DeviceNetworkType("testDeviceNetworkType",
            device_id=test_device.id,
            type="layer2-individual")
        test1_vlan = equinix_metal.Vlan("test1Vlan",
            description="VLAN in New Jersey",
            facility="ny5",
            project_id=local["project_id"])
        test2_vlan = equinix_metal.Vlan("test2Vlan",
            description="VLAN in New Jersey",
            facility="ny5",
            project_id=local["project_id"])
        test1_port_vlan_attachment = equinix_metal.PortVlanAttachment("test1PortVlanAttachment",
            device_id=test_device_network_type.id,
            vlan_vnid=test1_vlan.vxlan,
            port_name="eth1")
        test2_port_vlan_attachment = equinix_metal.PortVlanAttachment("test2PortVlanAttachment",
            device_id=test_device_network_type.id,
            vlan_vnid=test2_vlan.vxlan,
            port_name="eth1",
            native=True,
            opts=pulumi.ResourceOptions(depends_on=["metal_port_vlan_attachment.test1"]))
        ```
        ## Attribute Referece

        * `id` - UUID of device port used in the assignment
        * `vlan_id` - UUID of VLAN API resource
        * `port_id` - UUID of device port

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: ID of device to be assigned to the VLAN
        :param pulumi.Input[bool] force_bond: Add port back to the bond when this resource is removed. Default is false.
        :param pulumi.Input[bool] native: Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        :param pulumi.Input[str] port_name: Name of network port to be assigned to the VLAN
        :param pulumi.Input[int] vlan_vnid: VXLAN Network Identifier, integer
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PortVlanAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to attach device ports to VLANs.

        Device and VLAN must be in the same facility.

        If you need this resource to add the port back to bond on removal, set `force_bond = true`.

        To learn more about Layer 2 networking in Equinix Metal, refer to

        * <https://metal.equinix.com/developers/docs/networking/layer2/>
        * <https://metal.equinix.com/developers/docs/networking/layer2-configs/>

        ## Example Usage
        ### Hybrid network type

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test_vlan = equinix_metal.Vlan("testVlan",
            description="VLAN in New Jersey",
            facility="ny5",
            project_id=local["project_id"])
        test_device = equinix_metal.Device("testDevice",
            hostname="test",
            plan="c3.small.x86",
            facilities=["ny5"],
            operating_system="ubuntu_20_04",
            billing_cycle="hourly",
            project_id=local["project_id"])
        test_device_network_type = equinix_metal.DeviceNetworkType("testDeviceNetworkType",
            device_id=test_device.id,
            type="hybrid")
        test_port_vlan_attachment = equinix_metal.PortVlanAttachment("testPortVlanAttachment",
            device_id=test_device_network_type.id,
            port_name="eth1",
            vlan_vnid=test_vlan.vxlan)
        ```
        ### Layer 2 network

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test_device = equinix_metal.Device("testDevice",
            hostname="test",
            plan="c3.small.x86",
            facilities=["ny5"],
            operating_system="ubuntu_20_04",
            billing_cycle="hourly",
            project_id=local["project_id"])
        test_device_network_type = equinix_metal.DeviceNetworkType("testDeviceNetworkType",
            device_id=test_device.id,
            type="layer2-individual")
        test1_vlan = equinix_metal.Vlan("test1Vlan",
            description="VLAN in New Jersey",
            facility="ny5",
            project_id=local["project_id"])
        test2_vlan = equinix_metal.Vlan("test2Vlan",
            description="VLAN in New Jersey",
            facility="ny5",
            project_id=local["project_id"])
        test1_port_vlan_attachment = equinix_metal.PortVlanAttachment("test1PortVlanAttachment",
            device_id=test_device_network_type.id,
            vlan_vnid=test1_vlan.vxlan,
            port_name="eth1")
        test2_port_vlan_attachment = equinix_metal.PortVlanAttachment("test2PortVlanAttachment",
            device_id=test_device_network_type.id,
            vlan_vnid=test2_vlan.vxlan,
            port_name="eth1",
            native=True,
            opts=pulumi.ResourceOptions(depends_on=["metal_port_vlan_attachment.test1"]))
        ```
        ## Attribute Referece

        * `id` - UUID of device port used in the assignment
        * `vlan_id` - UUID of VLAN API resource
        * `port_id` - UUID of device port

        :param str resource_name: The name of the resource.
        :param PortVlanAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PortVlanAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 force_bond: Optional[pulumi.Input[bool]] = None,
                 native: Optional[pulumi.Input[bool]] = None,
                 port_name: Optional[pulumi.Input[str]] = None,
                 vlan_vnid: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PortVlanAttachmentArgs.__new__(PortVlanAttachmentArgs)

            if device_id is None and not opts.urn:
                raise TypeError("Missing required property 'device_id'")
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["force_bond"] = force_bond
            __props__.__dict__["native"] = native
            if port_name is None and not opts.urn:
                raise TypeError("Missing required property 'port_name'")
            __props__.__dict__["port_name"] = port_name
            if vlan_vnid is None and not opts.urn:
                raise TypeError("Missing required property 'vlan_vnid'")
            __props__.__dict__["vlan_vnid"] = vlan_vnid
            __props__.__dict__["port_id"] = None
            __props__.__dict__["vlan_id"] = None
        super(PortVlanAttachment, __self__).__init__(
            'equinix-metal:index/portVlanAttachment:PortVlanAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_id: Optional[pulumi.Input[str]] = None,
            force_bond: Optional[pulumi.Input[bool]] = None,
            native: Optional[pulumi.Input[bool]] = None,
            port_id: Optional[pulumi.Input[str]] = None,
            port_name: Optional[pulumi.Input[str]] = None,
            vlan_id: Optional[pulumi.Input[str]] = None,
            vlan_vnid: Optional[pulumi.Input[int]] = None) -> 'PortVlanAttachment':
        """
        Get an existing PortVlanAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: ID of device to be assigned to the VLAN
        :param pulumi.Input[bool] force_bond: Add port back to the bond when this resource is removed. Default is false.
        :param pulumi.Input[bool] native: Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        :param pulumi.Input[str] port_name: Name of network port to be assigned to the VLAN
        :param pulumi.Input[int] vlan_vnid: VXLAN Network Identifier, integer
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PortVlanAttachmentState.__new__(_PortVlanAttachmentState)

        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["force_bond"] = force_bond
        __props__.__dict__["native"] = native
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["port_name"] = port_name
        __props__.__dict__["vlan_id"] = vlan_id
        __props__.__dict__["vlan_vnid"] = vlan_vnid
        return PortVlanAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        ID of device to be assigned to the VLAN
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="forceBond")
    def force_bond(self) -> pulumi.Output[Optional[bool]]:
        """
        Add port back to the bond when this resource is removed. Default is false.
        """
        return pulumi.get(self, "force_bond")

    @property
    @pulumi.getter
    def native(self) -> pulumi.Output[Optional[bool]]:
        """
        Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        """
        return pulumi.get(self, "native")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="portName")
    def port_name(self) -> pulumi.Output[str]:
        """
        Name of network port to be assigned to the VLAN
        """
        return pulumi.get(self, "port_name")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vlan_id")

    @property
    @pulumi.getter(name="vlanVnid")
    def vlan_vnid(self) -> pulumi.Output[int]:
        """
        VXLAN Network Identifier, integer
        """
        return pulumi.get(self, "vlan_vnid")

