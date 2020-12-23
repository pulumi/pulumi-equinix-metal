# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetSpotMarketPriceResult',
    'AwaitableGetSpotMarketPriceResult',
    'get_spot_market_price',
]

@pulumi.output_type
class GetSpotMarketPriceResult:
    """
    A collection of values returned by getSpotMarketPrice.
    """
    def __init__(__self__, facility=None, id=None, plan=None, price=None):
        if facility and not isinstance(facility, str):
            raise TypeError("Expected argument 'facility' to be a str")
        pulumi.set(__self__, "facility", facility)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if price and not isinstance(price, float):
            raise TypeError("Expected argument 'price' to be a float")
        pulumi.set(__self__, "price", price)

    @property
    @pulumi.getter
    def facility(self) -> str:
        return pulumi.get(self, "facility")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def plan(self) -> str:
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def price(self) -> float:
        """
        Current spot market price for given plan in given facility.
        """
        return pulumi.get(self, "price")


class AwaitableGetSpotMarketPriceResult(GetSpotMarketPriceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSpotMarketPriceResult(
            facility=self.facility,
            id=self.id,
            plan=self.plan,
            price=self.price)


def get_spot_market_price(facility: Optional[str] = None,
                          plan: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSpotMarketPriceResult:
    """
    Use this data source to get Equinix Metal Spot Market Price.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix_metal as equinix_metal

    example = equinix_metal.get_spot_market_price(facility="ewr1",
        plan="c1.small.x86")
    ```


    :param str facility: Name of the facility.
    :param str plan: Name of the plan.
    """
    __args__ = dict()
    __args__['facility'] = facility
    __args__['plan'] = plan
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('equinix-metal:index/getSpotMarketPrice:getSpotMarketPrice', __args__, opts=opts, typ=GetSpotMarketPriceResult).value

    return AwaitableGetSpotMarketPriceResult(
        facility=__ret__.facility,
        id=__ret__.id,
        plan=__ret__.plan,
        price=__ret__.price)
