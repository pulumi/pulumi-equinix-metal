# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ConnectionArgs', 'Connection']

@pulumi.input_type
class ConnectionArgs:
    def __init__(__self__, *,
                 organization_id: pulumi.Input[str],
                 redundancy: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Connection resource.
        :param pulumi.Input[str] organization_id: ID of the organization responsible for the connection
        :param pulumi.Input[str] redundancy: Connection redundancy - redundant or primary
        :param pulumi.Input[str] type: Connection type - dedicated or shared
        :param pulumi.Input[str] description: Description for the connection resource
        :param pulumi.Input[str] facility: Facility where the connection will be created
        :param pulumi.Input[str] metro: Metro where the connection will be created
        :param pulumi.Input[str] name: Name of the connection resource
        :param pulumi.Input[str] project_id: ID of the project where the connection is scoped to, must be set for shared connection
        """
        pulumi.set(__self__, "organization_id", organization_id)
        pulumi.set(__self__, "redundancy", redundancy)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if facility is not None:
            pulumi.set(__self__, "facility", facility)
        if metro is not None:
            pulumi.set(__self__, "metro", metro)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        """
        ID of the organization responsible for the connection
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def redundancy(self) -> pulumi.Input[str]:
        """
        Connection redundancy - redundant or primary
        """
        return pulumi.get(self, "redundancy")

    @redundancy.setter
    def redundancy(self, value: pulumi.Input[str]):
        pulumi.set(self, "redundancy", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Connection type - dedicated or shared
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the connection resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[str]]:
        """
        Facility where the connection will be created
        """
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def metro(self) -> Optional[pulumi.Input[str]]:
        """
        Metro where the connection will be created
        """
        return pulumi.get(self, "metro")

    @metro.setter
    def metro(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metro", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the connection resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the project where the connection is scoped to, must be set for shared connection
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _ConnectionState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input['ConnectionPortArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 redundancy: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Connection resources.
        :param pulumi.Input[str] description: Description for the connection resource
        :param pulumi.Input[str] facility: Facility where the connection will be created
        :param pulumi.Input[str] metro: Metro where the connection will be created
        :param pulumi.Input[str] name: Name of the connection resource
        :param pulumi.Input[str] organization_id: ID of the organization responsible for the connection
        :param pulumi.Input[Sequence[pulumi.Input['ConnectionPortArgs']]] ports: List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the Connection datasource.
        :param pulumi.Input[str] project_id: ID of the project where the connection is scoped to, must be set for shared connection
        :param pulumi.Input[str] redundancy: Connection redundancy - redundant or primary
        :param pulumi.Input[int] speed: Port speed in bits per second
        :param pulumi.Input[str] status: Status of the connection resource
        :param pulumi.Input[str] token: Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
        :param pulumi.Input[str] type: Connection type - dedicated or shared
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if facility is not None:
            pulumi.set(__self__, "facility", facility)
        if metro is not None:
            pulumi.set(__self__, "metro", metro)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if redundancy is not None:
            pulumi.set(__self__, "redundancy", redundancy)
        if speed is not None:
            pulumi.set(__self__, "speed", speed)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the connection resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[str]]:
        """
        Facility where the connection will be created
        """
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def metro(self) -> Optional[pulumi.Input[str]]:
        """
        Metro where the connection will be created
        """
        return pulumi.get(self, "metro")

    @metro.setter
    def metro(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metro", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the connection resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization responsible for the connection
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConnectionPortArgs']]]]:
        """
        List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the Connection datasource.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConnectionPortArgs']]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the project where the connection is scoped to, must be set for shared connection
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def redundancy(self) -> Optional[pulumi.Input[str]]:
        """
        Connection redundancy - redundant or primary
        """
        return pulumi.get(self, "redundancy")

    @redundancy.setter
    def redundancy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redundancy", value)

    @property
    @pulumi.getter
    def speed(self) -> Optional[pulumi.Input[int]]:
        """
        Port speed in bits per second
        """
        return pulumi.get(self, "speed")

    @speed.setter
    def speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "speed", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the connection resource
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Connection type - dedicated or shared
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Connection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 redundancy: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to request of create an Interconnection from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test = equinix_metal.Connection("test",
            organization_id=local["my_organization_id"],
            project_id=local["my_project_id"],
            metro="sv",
            redundancy="redundant",
            type="shared")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for the connection resource
        :param pulumi.Input[str] facility: Facility where the connection will be created
        :param pulumi.Input[str] metro: Metro where the connection will be created
        :param pulumi.Input[str] name: Name of the connection resource
        :param pulumi.Input[str] organization_id: ID of the organization responsible for the connection
        :param pulumi.Input[str] project_id: ID of the project where the connection is scoped to, must be set for shared connection
        :param pulumi.Input[str] redundancy: Connection redundancy - redundant or primary
        :param pulumi.Input[str] type: Connection type - dedicated or shared
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to request of create an Interconnection from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test = equinix_metal.Connection("test",
            organization_id=local["my_organization_id"],
            project_id=local["my_project_id"],
            metro="sv",
            redundancy="redundant",
            type="shared")
        ```

        :param str resource_name: The name of the resource.
        :param ConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 metro: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 redundancy: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConnectionArgs.__new__(ConnectionArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["facility"] = facility
            __props__.__dict__["metro"] = metro
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["project_id"] = project_id
            if redundancy is None and not opts.urn:
                raise TypeError("Missing required property 'redundancy'")
            __props__.__dict__["redundancy"] = redundancy
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["ports"] = None
            __props__.__dict__["speed"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["token"] = None
        super(Connection, __self__).__init__(
            'equinix-metal:index/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            facility: Optional[pulumi.Input[str]] = None,
            metro: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConnectionPortArgs']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            redundancy: Optional[pulumi.Input[str]] = None,
            speed: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for the connection resource
        :param pulumi.Input[str] facility: Facility where the connection will be created
        :param pulumi.Input[str] metro: Metro where the connection will be created
        :param pulumi.Input[str] name: Name of the connection resource
        :param pulumi.Input[str] organization_id: ID of the organization responsible for the connection
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConnectionPortArgs']]]] ports: List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the Connection datasource.
        :param pulumi.Input[str] project_id: ID of the project where the connection is scoped to, must be set for shared connection
        :param pulumi.Input[str] redundancy: Connection redundancy - redundant or primary
        :param pulumi.Input[int] speed: Port speed in bits per second
        :param pulumi.Input[str] status: Status of the connection resource
        :param pulumi.Input[str] token: Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
        :param pulumi.Input[str] type: Connection type - dedicated or shared
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConnectionState.__new__(_ConnectionState)

        __props__.__dict__["description"] = description
        __props__.__dict__["facility"] = facility
        __props__.__dict__["metro"] = metro
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["ports"] = ports
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["redundancy"] = redundancy
        __props__.__dict__["speed"] = speed
        __props__.__dict__["status"] = status
        __props__.__dict__["token"] = token
        __props__.__dict__["type"] = type
        return Connection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description for the connection resource
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def facility(self) -> pulumi.Output[str]:
        """
        Facility where the connection will be created
        """
        return pulumi.get(self, "facility")

    @property
    @pulumi.getter
    def metro(self) -> pulumi.Output[str]:
        """
        Metro where the connection will be created
        """
        return pulumi.get(self, "metro")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the connection resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        ID of the organization responsible for the connection
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Output[Sequence['outputs.ConnectionPort']]:
        """
        List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of port is described in documentation of the Connection datasource.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the project where the connection is scoped to, must be set for shared connection
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def redundancy(self) -> pulumi.Output[str]:
        """
        Connection redundancy - redundant or primary
        """
        return pulumi.get(self, "redundancy")

    @property
    @pulumi.getter
    def speed(self) -> pulumi.Output[int]:
        """
        Port speed in bits per second
        """
        return pulumi.get(self, "speed")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the connection resource
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Connection type - dedicated or shared
        """
        return pulumi.get(self, "type")

