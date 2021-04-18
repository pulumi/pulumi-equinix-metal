# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFacilityResult',
    'AwaitableGetFacilityResult',
    'get_facility',
]

@pulumi.output_type
class GetFacilityResult:
    """
    A collection of values returned by getFacility.
    """
    def __init__(__self__, code=None, features=None, id=None, metro=None, name=None):
        if code and not isinstance(code, str):
            raise TypeError("Expected argument 'code' to be a str")
        pulumi.set(__self__, "code", code)
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        pulumi.set(__self__, "features", features)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metro and not isinstance(metro, str):
            raise TypeError("Expected argument 'metro' to be a str")
        pulumi.set(__self__, "metro", metro)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def code(self) -> str:
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def features(self) -> Sequence[str]:
        """
        The features of the facility
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metro(self) -> str:
        """
        The metro code the facility is part of
        """
        return pulumi.get(self, "metro")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the facility
        """
        return pulumi.get(self, "name")


class AwaitableGetFacilityResult(GetFacilityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFacilityResult(
            code=self.code,
            features=self.features,
            id=self.id,
            metro=self.metro,
            name=self.name)


def get_facility(code: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFacilityResult:
    """
    Provides an Equinix Metal facility datasource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix_metal as equinix_metal

    ny5 = equinix_metal.get_facility(code="ny5")
    pulumi.export("id", ny5.id)
    ```


    :param str code: The facility code
    """
    __args__ = dict()
    __args__['code'] = code
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('equinix-metal:index/getFacility:getFacility', __args__, opts=opts, typ=GetFacilityResult).value

    return AwaitableGetFacilityResult(
        code=__ret__.code,
        features=__ret__.features,
        id=__ret__.id,
        metro=__ret__.metro,
        name=__ret__.name)
