# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VirtualCircuitArgs', 'VirtualCircuit']

@pulumi.input_type
class VirtualCircuitArgs:
    def __init__(__self__, *,
                 connection_id: pulumi.Input[str],
                 port_id: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 vlan_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nni_vlan: Optional[pulumi.Input[int]] = None,
                 speed: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VirtualCircuit resource.
        :param pulumi.Input[str] connection_id: UUID of Connection where the VC is scoped to
        :param pulumi.Input[str] port_id: UUID of the Connection Port where the VC is scoped to
        :param pulumi.Input[str] project_id: UUID of the Project where the VC is scoped to
        :param pulumi.Input[str] vlan_id: UUID of the VLAN to associate
        :param pulumi.Input[str] description: Description for the Virtual Circuit resource
        :param pulumi.Input[str] name: Name of the Virtual Circuit resource
        :param pulumi.Input[int] nni_vlan: Equinix Metal network-to-network VLAN ID
        :param pulumi.Input[str] speed: Speed of the Virtual Circuit resource
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for the Virtual Circuit resource
        """
        VirtualCircuitArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            connection_id=connection_id,
            port_id=port_id,
            project_id=project_id,
            vlan_id=vlan_id,
            description=description,
            name=name,
            nni_vlan=nni_vlan,
            speed=speed,
            tags=tags,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             connection_id: pulumi.Input[str],
             port_id: pulumi.Input[str],
             project_id: pulumi.Input[str],
             vlan_id: pulumi.Input[str],
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             nni_vlan: Optional[pulumi.Input[int]] = None,
             speed: Optional[pulumi.Input[str]] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("connection_id", connection_id)
        _setter("port_id", port_id)
        _setter("project_id", project_id)
        _setter("vlan_id", vlan_id)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if nni_vlan is not None:
            _setter("nni_vlan", nni_vlan)
        if speed is not None:
            _setter("speed", speed)
        if tags is not None:
            _setter("tags", tags)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Input[str]:
        """
        UUID of Connection where the VC is scoped to
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Input[str]:
        """
        UUID of the Connection Port where the VC is scoped to
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        UUID of the Project where the VC is scoped to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Input[str]:
        """
        UUID of the VLAN to associate
        """
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vlan_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the Virtual Circuit resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Virtual Circuit resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nniVlan")
    def nni_vlan(self) -> Optional[pulumi.Input[int]]:
        """
        Equinix Metal network-to-network VLAN ID
        """
        return pulumi.get(self, "nni_vlan")

    @nni_vlan.setter
    def nni_vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nni_vlan", value)

    @property
    @pulumi.getter
    def speed(self) -> Optional[pulumi.Input[str]]:
        """
        Speed of the Virtual Circuit resource
        """
        return pulumi.get(self, "speed")

    @speed.setter
    def speed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "speed", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags for the Virtual Circuit resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VirtualCircuitState:
    def __init__(__self__, *,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nni_vlan: Optional[pulumi.Input[int]] = None,
                 nni_vnid: Optional[pulumi.Input[int]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vlan_id: Optional[pulumi.Input[str]] = None,
                 vnid: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering VirtualCircuit resources.
        :param pulumi.Input[str] connection_id: UUID of Connection where the VC is scoped to
        :param pulumi.Input[str] description: Description for the Virtual Circuit resource
        :param pulumi.Input[str] name: Name of the Virtual Circuit resource
        :param pulumi.Input[int] nni_vlan: Equinix Metal network-to-network VLAN ID
        :param pulumi.Input[int] nni_vnid: Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        :param pulumi.Input[str] port_id: UUID of the Connection Port where the VC is scoped to
        :param pulumi.Input[str] project_id: UUID of the Project where the VC is scoped to
        :param pulumi.Input[str] speed: Speed of the Virtual Circuit resource
        :param pulumi.Input[str] status: Status of the virtal circuit
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for the Virtual Circuit resource
        :param pulumi.Input[str] vlan_id: UUID of the VLAN to associate
        :param pulumi.Input[int] vnid: VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        """
        _VirtualCircuitState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            connection_id=connection_id,
            description=description,
            name=name,
            nni_vlan=nni_vlan,
            nni_vnid=nni_vnid,
            port_id=port_id,
            project_id=project_id,
            speed=speed,
            status=status,
            tags=tags,
            vlan_id=vlan_id,
            vnid=vnid,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             connection_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             nni_vlan: Optional[pulumi.Input[int]] = None,
             nni_vnid: Optional[pulumi.Input[int]] = None,
             port_id: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             speed: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             vlan_id: Optional[pulumi.Input[str]] = None,
             vnid: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if connection_id is not None:
            _setter("connection_id", connection_id)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if nni_vlan is not None:
            _setter("nni_vlan", nni_vlan)
        if nni_vnid is not None:
            _setter("nni_vnid", nni_vnid)
        if port_id is not None:
            _setter("port_id", port_id)
        if project_id is not None:
            _setter("project_id", project_id)
        if speed is not None:
            _setter("speed", speed)
        if status is not None:
            _setter("status", status)
        if tags is not None:
            _setter("tags", tags)
        if vlan_id is not None:
            _setter("vlan_id", vlan_id)
        if vnid is not None:
            _setter("vnid", vnid)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of Connection where the VC is scoped to
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the Virtual Circuit resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Virtual Circuit resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nniVlan")
    def nni_vlan(self) -> Optional[pulumi.Input[int]]:
        """
        Equinix Metal network-to-network VLAN ID
        """
        return pulumi.get(self, "nni_vlan")

    @nni_vlan.setter
    def nni_vlan(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nni_vlan", value)

    @property
    @pulumi.getter(name="nniVnid")
    def nni_vnid(self) -> Optional[pulumi.Input[int]]:
        """
        Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        """
        return pulumi.get(self, "nni_vnid")

    @nni_vnid.setter
    def nni_vnid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nni_vnid", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the Connection Port where the VC is scoped to
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the Project where the VC is scoped to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def speed(self) -> Optional[pulumi.Input[str]]:
        """
        Speed of the Virtual Circuit resource
        """
        return pulumi.get(self, "speed")

    @speed.setter
    def speed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "speed", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the virtal circuit
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags for the Virtual Circuit resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the VLAN to associate
        """
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_id", value)

    @property
    @pulumi.getter
    def vnid(self) -> Optional[pulumi.Input[int]]:
        """
        VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        """
        return pulumi.get(self, "vnid")

    @vnid.setter
    def vnid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vnid", value)


class VirtualCircuit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nni_vlan: Optional[pulumi.Input[int]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vlan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to associate VLAN with a Dedicated Port from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/#associating-a-vlan-with-a-dedicated-port).

        ## Example Usage

        Pick an existing Project and Connection, create a VLAN and use `VirtualCircuit` to associate it with a Primary Port of the Connection.

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        project_id = "52000fb2-ee46-4673-93a8-de2c2bdba33c"
        conn_id = "73f12f29-3e19-43a0-8e90-ae81580db1e0"
        test_connection = equinix_metal.get_connection(connection_id=conn_id)
        test_vlan = equinix_metal.Vlan("testVlan",
            project_id=project_id,
            metro=test_connection.metro)
        test_virtual_circuit = equinix_metal.VirtualCircuit("testVirtualCircuit",
            connection_id=conn_id,
            project_id=project_id,
            port_id=test_connection.ports[0].id,
            vlan_id=test_vlan.id,
            nni_vlan=1056)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_id: UUID of Connection where the VC is scoped to
        :param pulumi.Input[str] description: Description for the Virtual Circuit resource
        :param pulumi.Input[str] name: Name of the Virtual Circuit resource
        :param pulumi.Input[int] nni_vlan: Equinix Metal network-to-network VLAN ID
        :param pulumi.Input[str] port_id: UUID of the Connection Port where the VC is scoped to
        :param pulumi.Input[str] project_id: UUID of the Project where the VC is scoped to
        :param pulumi.Input[str] speed: Speed of the Virtual Circuit resource
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for the Virtual Circuit resource
        :param pulumi.Input[str] vlan_id: UUID of the VLAN to associate
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualCircuitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to associate VLAN with a Dedicated Port from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/#associating-a-vlan-with-a-dedicated-port).

        ## Example Usage

        Pick an existing Project and Connection, create a VLAN and use `VirtualCircuit` to associate it with a Primary Port of the Connection.

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        project_id = "52000fb2-ee46-4673-93a8-de2c2bdba33c"
        conn_id = "73f12f29-3e19-43a0-8e90-ae81580db1e0"
        test_connection = equinix_metal.get_connection(connection_id=conn_id)
        test_vlan = equinix_metal.Vlan("testVlan",
            project_id=project_id,
            metro=test_connection.metro)
        test_virtual_circuit = equinix_metal.VirtualCircuit("testVirtualCircuit",
            connection_id=conn_id,
            project_id=project_id,
            port_id=test_connection.ports[0].id,
            vlan_id=test_vlan.id,
            nni_vlan=1056)
        ```

        :param str resource_name: The name of the resource.
        :param VirtualCircuitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualCircuitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VirtualCircuitArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nni_vlan: Optional[pulumi.Input[int]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vlan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualCircuitArgs.__new__(VirtualCircuitArgs)

            if connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'connection_id'")
            __props__.__dict__["connection_id"] = connection_id
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["nni_vlan"] = nni_vlan
            if port_id is None and not opts.urn:
                raise TypeError("Missing required property 'port_id'")
            __props__.__dict__["port_id"] = port_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["speed"] = speed
            __props__.__dict__["tags"] = tags
            if vlan_id is None and not opts.urn:
                raise TypeError("Missing required property 'vlan_id'")
            __props__.__dict__["vlan_id"] = vlan_id
            __props__.__dict__["nni_vnid"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["vnid"] = None
        super(VirtualCircuit, __self__).__init__(
            'equinix-metal:index/virtualCircuit:VirtualCircuit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nni_vlan: Optional[pulumi.Input[int]] = None,
            nni_vnid: Optional[pulumi.Input[int]] = None,
            port_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            speed: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            vlan_id: Optional[pulumi.Input[str]] = None,
            vnid: Optional[pulumi.Input[int]] = None) -> 'VirtualCircuit':
        """
        Get an existing VirtualCircuit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_id: UUID of Connection where the VC is scoped to
        :param pulumi.Input[str] description: Description for the Virtual Circuit resource
        :param pulumi.Input[str] name: Name of the Virtual Circuit resource
        :param pulumi.Input[int] nni_vlan: Equinix Metal network-to-network VLAN ID
        :param pulumi.Input[int] nni_vnid: Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        :param pulumi.Input[str] port_id: UUID of the Connection Port where the VC is scoped to
        :param pulumi.Input[str] project_id: UUID of the Project where the VC is scoped to
        :param pulumi.Input[str] speed: Speed of the Virtual Circuit resource
        :param pulumi.Input[str] status: Status of the virtal circuit
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for the Virtual Circuit resource
        :param pulumi.Input[str] vlan_id: UUID of the VLAN to associate
        :param pulumi.Input[int] vnid: VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualCircuitState.__new__(_VirtualCircuitState)

        __props__.__dict__["connection_id"] = connection_id
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["nni_vlan"] = nni_vlan
        __props__.__dict__["nni_vnid"] = nni_vnid
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["speed"] = speed
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vlan_id"] = vlan_id
        __props__.__dict__["vnid"] = vnid
        return VirtualCircuit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Output[str]:
        """
        UUID of Connection where the VC is scoped to
        """
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description for the Virtual Circuit resource
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the Virtual Circuit resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nniVlan")
    def nni_vlan(self) -> pulumi.Output[Optional[int]]:
        """
        Equinix Metal network-to-network VLAN ID
        """
        return pulumi.get(self, "nni_vlan")

    @property
    @pulumi.getter(name="nniVnid")
    def nni_vnid(self) -> pulumi.Output[int]:
        """
        Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        """
        return pulumi.get(self, "nni_vnid")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[str]:
        """
        UUID of the Connection Port where the VC is scoped to
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        UUID of the Project where the VC is scoped to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def speed(self) -> pulumi.Output[str]:
        """
        Speed of the Virtual Circuit resource
        """
        return pulumi.get(self, "speed")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the virtal circuit
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Tags for the Virtual Circuit resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Output[str]:
        """
        UUID of the VLAN to associate
        """
        return pulumi.get(self, "vlan_id")

    @property
    @pulumi.getter
    def vnid(self) -> pulumi.Output[int]:
        """
        VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        """
        return pulumi.get(self, "vnid")

