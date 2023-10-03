# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PortArgs', 'Port']

@pulumi.input_type
class PortArgs:
    def __init__(__self__, *,
                 bonded: pulumi.Input[bool],
                 port_id: pulumi.Input[str],
                 layer2: Optional[pulumi.Input[bool]] = None,
                 native_vlan_id: Optional[pulumi.Input[str]] = None,
                 reset_on_delete: Optional[pulumi.Input[bool]] = None,
                 vlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vxlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None):
        """
        The set of arguments for constructing a Port resource.
        :param pulumi.Input[bool] bonded: Whether the port should be bonded
        :param pulumi.Input[str] port_id: ID of the port to read
        :param pulumi.Input[bool] layer2: Whether to put the port to Layer 2 mode, valid only for bond ports
        :param pulumi.Input[str] native_vlan_id: UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlan_ids` parameter), valid only for physical (non-bond) ports
        :param pulumi.Input[bool] reset_on_delete: Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vlan_ids: List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
        :param pulumi.Input[Sequence[pulumi.Input[int]]] vxlan_ids: List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
        """
        PortArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            bonded=bonded,
            port_id=port_id,
            layer2=layer2,
            native_vlan_id=native_vlan_id,
            reset_on_delete=reset_on_delete,
            vlan_ids=vlan_ids,
            vxlan_ids=vxlan_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             bonded: pulumi.Input[bool],
             port_id: pulumi.Input[str],
             layer2: Optional[pulumi.Input[bool]] = None,
             native_vlan_id: Optional[pulumi.Input[str]] = None,
             reset_on_delete: Optional[pulumi.Input[bool]] = None,
             vlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             vxlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("bonded", bonded)
        _setter("port_id", port_id)
        if layer2 is not None:
            _setter("layer2", layer2)
        if native_vlan_id is not None:
            _setter("native_vlan_id", native_vlan_id)
        if reset_on_delete is not None:
            _setter("reset_on_delete", reset_on_delete)
        if vlan_ids is not None:
            _setter("vlan_ids", vlan_ids)
        if vxlan_ids is not None:
            _setter("vxlan_ids", vxlan_ids)

    @property
    @pulumi.getter
    def bonded(self) -> pulumi.Input[bool]:
        """
        Whether the port should be bonded
        """
        return pulumi.get(self, "bonded")

    @bonded.setter
    def bonded(self, value: pulumi.Input[bool]):
        pulumi.set(self, "bonded", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Input[str]:
        """
        ID of the port to read
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def layer2(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to put the port to Layer 2 mode, valid only for bond ports
        """
        return pulumi.get(self, "layer2")

    @layer2.setter
    def layer2(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "layer2", value)

    @property
    @pulumi.getter(name="nativeVlanId")
    def native_vlan_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlan_ids` parameter), valid only for physical (non-bond) ports
        """
        return pulumi.get(self, "native_vlan_id")

    @native_vlan_id.setter
    def native_vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "native_vlan_id", value)

    @property
    @pulumi.getter(name="resetOnDelete")
    def reset_on_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
        """
        return pulumi.get(self, "reset_on_delete")

    @reset_on_delete.setter
    def reset_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset_on_delete", value)

    @property
    @pulumi.getter(name="vlanIds")
    def vlan_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
        """
        return pulumi.get(self, "vlan_ids")

    @vlan_ids.setter
    def vlan_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vlan_ids", value)

    @property
    @pulumi.getter(name="vxlanIds")
    def vxlan_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
        """
        return pulumi.get(self, "vxlan_ids")

    @vxlan_ids.setter
    def vxlan_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "vxlan_ids", value)


@pulumi.input_type
class _PortState:
    def __init__(__self__, *,
                 bond_id: Optional[pulumi.Input[str]] = None,
                 bond_name: Optional[pulumi.Input[str]] = None,
                 bonded: Optional[pulumi.Input[bool]] = None,
                 disbond_supported: Optional[pulumi.Input[bool]] = None,
                 layer2: Optional[pulumi.Input[bool]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 native_vlan_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 reset_on_delete: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vxlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None):
        """
        Input properties used for looking up and filtering Port resources.
        :param pulumi.Input[str] bond_id: UUID of the bond port
        :param pulumi.Input[str] bond_name: Name of the bond port
        :param pulumi.Input[bool] bonded: Whether the port should be bonded
        :param pulumi.Input[bool] disbond_supported: Flag indicating whether the port can be removed from a bond
        :param pulumi.Input[bool] layer2: Whether to put the port to Layer 2 mode, valid only for bond ports
        :param pulumi.Input[str] mac: MAC address of the port
        :param pulumi.Input[str] name: Name of the port, e.g. `bond0` or `eth0`
        :param pulumi.Input[str] native_vlan_id: UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlan_ids` parameter), valid only for physical (non-bond) ports
        :param pulumi.Input[str] network_type: One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
        :param pulumi.Input[str] port_id: ID of the port to read
        :param pulumi.Input[bool] reset_on_delete: Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
        :param pulumi.Input[str] type: Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vlan_ids: List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
        :param pulumi.Input[Sequence[pulumi.Input[int]]] vxlan_ids: List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
        """
        _PortState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            bond_id=bond_id,
            bond_name=bond_name,
            bonded=bonded,
            disbond_supported=disbond_supported,
            layer2=layer2,
            mac=mac,
            name=name,
            native_vlan_id=native_vlan_id,
            network_type=network_type,
            port_id=port_id,
            reset_on_delete=reset_on_delete,
            type=type,
            vlan_ids=vlan_ids,
            vxlan_ids=vxlan_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             bond_id: Optional[pulumi.Input[str]] = None,
             bond_name: Optional[pulumi.Input[str]] = None,
             bonded: Optional[pulumi.Input[bool]] = None,
             disbond_supported: Optional[pulumi.Input[bool]] = None,
             layer2: Optional[pulumi.Input[bool]] = None,
             mac: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             native_vlan_id: Optional[pulumi.Input[str]] = None,
             network_type: Optional[pulumi.Input[str]] = None,
             port_id: Optional[pulumi.Input[str]] = None,
             reset_on_delete: Optional[pulumi.Input[bool]] = None,
             type: Optional[pulumi.Input[str]] = None,
             vlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             vxlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if bond_id is not None:
            _setter("bond_id", bond_id)
        if bond_name is not None:
            _setter("bond_name", bond_name)
        if bonded is not None:
            _setter("bonded", bonded)
        if disbond_supported is not None:
            _setter("disbond_supported", disbond_supported)
        if layer2 is not None:
            _setter("layer2", layer2)
        if mac is not None:
            _setter("mac", mac)
        if name is not None:
            _setter("name", name)
        if native_vlan_id is not None:
            _setter("native_vlan_id", native_vlan_id)
        if network_type is not None:
            _setter("network_type", network_type)
        if port_id is not None:
            _setter("port_id", port_id)
        if reset_on_delete is not None:
            _setter("reset_on_delete", reset_on_delete)
        if type is not None:
            _setter("type", type)
        if vlan_ids is not None:
            _setter("vlan_ids", vlan_ids)
        if vxlan_ids is not None:
            _setter("vxlan_ids", vxlan_ids)

    @property
    @pulumi.getter(name="bondId")
    def bond_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the bond port
        """
        return pulumi.get(self, "bond_id")

    @bond_id.setter
    def bond_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bond_id", value)

    @property
    @pulumi.getter(name="bondName")
    def bond_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the bond port
        """
        return pulumi.get(self, "bond_name")

    @bond_name.setter
    def bond_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bond_name", value)

    @property
    @pulumi.getter
    def bonded(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the port should be bonded
        """
        return pulumi.get(self, "bonded")

    @bonded.setter
    def bonded(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bonded", value)

    @property
    @pulumi.getter(name="disbondSupported")
    def disbond_supported(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag indicating whether the port can be removed from a bond
        """
        return pulumi.get(self, "disbond_supported")

    @disbond_supported.setter
    def disbond_supported(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disbond_supported", value)

    @property
    @pulumi.getter
    def layer2(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to put the port to Layer 2 mode, valid only for bond ports
        """
        return pulumi.get(self, "layer2")

    @layer2.setter
    def layer2(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "layer2", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address of the port
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the port, e.g. `bond0` or `eth0`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nativeVlanId")
    def native_vlan_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlan_ids` parameter), valid only for physical (non-bond) ports
        """
        return pulumi.get(self, "native_vlan_id")

    @native_vlan_id.setter
    def native_vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "native_vlan_id", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the port to read
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="resetOnDelete")
    def reset_on_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
        """
        return pulumi.get(self, "reset_on_delete")

    @reset_on_delete.setter
    def reset_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset_on_delete", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vlanIds")
    def vlan_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
        """
        return pulumi.get(self, "vlan_ids")

    @vlan_ids.setter
    def vlan_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vlan_ids", value)

    @property
    @pulumi.getter(name="vxlanIds")
    def vxlan_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
        """
        return pulumi.get(self, "vxlan_ids")

    @vxlan_ids.setter
    def vxlan_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "vxlan_ids", value)


class Port(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bonded: Optional[pulumi.Input[bool]] = None,
                 layer2: Optional[pulumi.Input[bool]] = None,
                 native_vlan_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 reset_on_delete: Optional[pulumi.Input[bool]] = None,
                 vlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vxlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        See the Network Types Guide for examples of this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bonded: Whether the port should be bonded
        :param pulumi.Input[bool] layer2: Whether to put the port to Layer 2 mode, valid only for bond ports
        :param pulumi.Input[str] native_vlan_id: UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlan_ids` parameter), valid only for physical (non-bond) ports
        :param pulumi.Input[str] port_id: ID of the port to read
        :param pulumi.Input[bool] reset_on_delete: Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vlan_ids: List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
        :param pulumi.Input[Sequence[pulumi.Input[int]]] vxlan_ids: List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PortArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        See the Network Types Guide for examples of this resource.

        :param str resource_name: The name of the resource.
        :param PortArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PortArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            PortArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bonded: Optional[pulumi.Input[bool]] = None,
                 layer2: Optional[pulumi.Input[bool]] = None,
                 native_vlan_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 reset_on_delete: Optional[pulumi.Input[bool]] = None,
                 vlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vxlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PortArgs.__new__(PortArgs)

            if bonded is None and not opts.urn:
                raise TypeError("Missing required property 'bonded'")
            __props__.__dict__["bonded"] = bonded
            __props__.__dict__["layer2"] = layer2
            __props__.__dict__["native_vlan_id"] = native_vlan_id
            if port_id is None and not opts.urn:
                raise TypeError("Missing required property 'port_id'")
            __props__.__dict__["port_id"] = port_id
            __props__.__dict__["reset_on_delete"] = reset_on_delete
            __props__.__dict__["vlan_ids"] = vlan_ids
            __props__.__dict__["vxlan_ids"] = vxlan_ids
            __props__.__dict__["bond_id"] = None
            __props__.__dict__["bond_name"] = None
            __props__.__dict__["disbond_supported"] = None
            __props__.__dict__["mac"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["network_type"] = None
            __props__.__dict__["type"] = None
        super(Port, __self__).__init__(
            'equinix-metal:index/port:Port',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bond_id: Optional[pulumi.Input[str]] = None,
            bond_name: Optional[pulumi.Input[str]] = None,
            bonded: Optional[pulumi.Input[bool]] = None,
            disbond_supported: Optional[pulumi.Input[bool]] = None,
            layer2: Optional[pulumi.Input[bool]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            native_vlan_id: Optional[pulumi.Input[str]] = None,
            network_type: Optional[pulumi.Input[str]] = None,
            port_id: Optional[pulumi.Input[str]] = None,
            reset_on_delete: Optional[pulumi.Input[bool]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vxlan_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None) -> 'Port':
        """
        Get an existing Port resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bond_id: UUID of the bond port
        :param pulumi.Input[str] bond_name: Name of the bond port
        :param pulumi.Input[bool] bonded: Whether the port should be bonded
        :param pulumi.Input[bool] disbond_supported: Flag indicating whether the port can be removed from a bond
        :param pulumi.Input[bool] layer2: Whether to put the port to Layer 2 mode, valid only for bond ports
        :param pulumi.Input[str] mac: MAC address of the port
        :param pulumi.Input[str] name: Name of the port, e.g. `bond0` or `eth0`
        :param pulumi.Input[str] native_vlan_id: UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlan_ids` parameter), valid only for physical (non-bond) ports
        :param pulumi.Input[str] network_type: One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
        :param pulumi.Input[str] port_id: ID of the port to read
        :param pulumi.Input[bool] reset_on_delete: Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
        :param pulumi.Input[str] type: Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vlan_ids: List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
        :param pulumi.Input[Sequence[pulumi.Input[int]]] vxlan_ids: List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PortState.__new__(_PortState)

        __props__.__dict__["bond_id"] = bond_id
        __props__.__dict__["bond_name"] = bond_name
        __props__.__dict__["bonded"] = bonded
        __props__.__dict__["disbond_supported"] = disbond_supported
        __props__.__dict__["layer2"] = layer2
        __props__.__dict__["mac"] = mac
        __props__.__dict__["name"] = name
        __props__.__dict__["native_vlan_id"] = native_vlan_id
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["reset_on_delete"] = reset_on_delete
        __props__.__dict__["type"] = type
        __props__.__dict__["vlan_ids"] = vlan_ids
        __props__.__dict__["vxlan_ids"] = vxlan_ids
        return Port(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bondId")
    def bond_id(self) -> pulumi.Output[str]:
        """
        UUID of the bond port
        """
        return pulumi.get(self, "bond_id")

    @property
    @pulumi.getter(name="bondName")
    def bond_name(self) -> pulumi.Output[str]:
        """
        Name of the bond port
        """
        return pulumi.get(self, "bond_name")

    @property
    @pulumi.getter
    def bonded(self) -> pulumi.Output[bool]:
        """
        Whether the port should be bonded
        """
        return pulumi.get(self, "bonded")

    @property
    @pulumi.getter(name="disbondSupported")
    def disbond_supported(self) -> pulumi.Output[bool]:
        """
        Flag indicating whether the port can be removed from a bond
        """
        return pulumi.get(self, "disbond_supported")

    @property
    @pulumi.getter
    def layer2(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to put the port to Layer 2 mode, valid only for bond ports
        """
        return pulumi.get(self, "layer2")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        MAC address of the port
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the port, e.g. `bond0` or `eth0`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nativeVlanId")
    def native_vlan_id(self) -> pulumi.Output[Optional[str]]:
        """
        UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlan_ids` parameter), valid only for physical (non-bond) ports
        """
        return pulumi.get(self, "native_vlan_id")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[str]:
        """
        One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[str]:
        """
        ID of the port to read
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="resetOnDelete")
    def reset_on_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
        """
        return pulumi.get(self, "reset_on_delete")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vlanIds")
    def vlan_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
        """
        return pulumi.get(self, "vlan_ids")

    @property
    @pulumi.getter(name="vxlanIds")
    def vxlan_ids(self) -> pulumi.Output[Sequence[int]]:
        """
        List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
        """
        return pulumi.get(self, "vxlan_ids")

