# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVirtualCircuitResult',
    'AwaitableGetVirtualCircuitResult',
    'get_virtual_circuit',
]

@pulumi.output_type
class GetVirtualCircuitResult:
    """
    A collection of values returned by getVirtualCircuit.
    """
    def __init__(__self__, id=None, name=None, nni_vlan=None, nni_vnid=None, project_id=None, status=None, virtual_circuit_id=None, vnid=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nni_vlan and not isinstance(nni_vlan, int):
            raise TypeError("Expected argument 'nni_vlan' to be a int")
        pulumi.set(__self__, "nni_vlan", nni_vlan)
        if nni_vnid and not isinstance(nni_vnid, int):
            raise TypeError("Expected argument 'nni_vnid' to be a int")
        pulumi.set(__self__, "nni_vnid", nni_vnid)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if virtual_circuit_id and not isinstance(virtual_circuit_id, str):
            raise TypeError("Expected argument 'virtual_circuit_id' to be a str")
        pulumi.set(__self__, "virtual_circuit_id", virtual_circuit_id)
        if vnid and not isinstance(vnid, int):
            raise TypeError("Expected argument 'vnid' to be a int")
        pulumi.set(__self__, "vnid", vnid)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the virtual circuit resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nniVlan")
    def nni_vlan(self) -> int:
        return pulumi.get(self, "nni_vlan")

    @property
    @pulumi.getter(name="nniVnid")
    def nni_vnid(self) -> int:
        return pulumi.get(self, "nni_vnid")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        ID of project to which the VC belongs
        * `vnid`, `nni_vlan`, `nni_nvid` - VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the virtal circuit
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="virtualCircuitId")
    def virtual_circuit_id(self) -> str:
        return pulumi.get(self, "virtual_circuit_id")

    @property
    @pulumi.getter
    def vnid(self) -> int:
        return pulumi.get(self, "vnid")


class AwaitableGetVirtualCircuitResult(GetVirtualCircuitResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualCircuitResult(
            id=self.id,
            name=self.name,
            nni_vlan=self.nni_vlan,
            nni_vnid=self.nni_vnid,
            project_id=self.project_id,
            status=self.status,
            virtual_circuit_id=self.virtual_circuit_id,
            vnid=self.vnid)


def get_virtual_circuit(virtual_circuit_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualCircuitResult:
    """
    Use this data source to retrieve a virtual circuit resource from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix_metal as equinix_metal

    example_connection = equinix_metal.get_connection(connection_id="4347e805-eb46-4699-9eb9-5c116e6a017d")
    example_vc = equinix_metal.get_virtual_circuit(virtual_circuit_id=example_connection.ports[1].virtual_circuit_ids[0])
    ```


    :param str virtual_circuit_id: ID of the virtual circuit resource
    """
    __args__ = dict()
    __args__['virtualCircuitId'] = virtual_circuit_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('equinix-metal:index/getVirtualCircuit:getVirtualCircuit', __args__, opts=opts, typ=GetVirtualCircuitResult).value

    return AwaitableGetVirtualCircuitResult(
        id=__ret__.id,
        name=__ret__.name,
        nni_vlan=__ret__.nni_vlan,
        nni_vnid=__ret__.nni_vnid,
        project_id=__ret__.project_id,
        status=__ret__.status,
        virtual_circuit_id=__ret__.virtual_circuit_id,
        vnid=__ret__.vnid)