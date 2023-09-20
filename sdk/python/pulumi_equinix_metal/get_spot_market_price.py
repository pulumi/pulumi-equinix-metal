# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSpotMarketPriceResult',
    'AwaitableGetSpotMarketPriceResult',
    'get_spot_market_price',
    'get_spot_market_price_output',
]

@pulumi.output_type
class GetSpotMarketPriceResult:
    """
    A collection of values returned by getSpotMarketPrice.
    """
    def __init__(__self__, facility=None, id=None, metro=None, plan=None, price=None):
        if facility and not isinstance(facility, str):
            raise TypeError("Expected argument 'facility' to be a str")
        pulumi.set(__self__, "facility", facility)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metro and not isinstance(metro, str):
            raise TypeError("Expected argument 'metro' to be a str")
        pulumi.set(__self__, "metro", metro)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if price and not isinstance(price, float):
            raise TypeError("Expected argument 'price' to be a float")
        pulumi.set(__self__, "price", price)

    @property
    @pulumi.getter
    def facility(self) -> Optional[str]:
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
    def metro(self) -> Optional[str]:
        return pulumi.get(self, "metro")

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
            metro=self.metro,
            plan=self.plan,
            price=self.price)


def get_spot_market_price(facility: Optional[str] = None,
                          metro: Optional[str] = None,
                          plan: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSpotMarketPriceResult:
    """
    Use this data source to get Equinix Metal Spot Market Price for a plan.

    ## Example Usage

    Lookup by facility:

    ```python
    import pulumi
    import pulumi_equinix_metal as equinix_metal

    example = equinix_metal.get_spot_market_price(facility="ny5",
        plan="c3.small.x86")
    ```

    Lookup by metro:

    ```python
    import pulumi
    import pulumi_equinix_metal as equinix_metal

    example = equinix_metal.get_spot_market_price(metro="sv",
        plan="c3.small.x86")
    ```


    :param str facility: Name of the facility.
    :param str metro: Name of the metro.
    :param str plan: Name of the plan.
    """
    __args__ = dict()
    __args__['facility'] = facility
    __args__['metro'] = metro
    __args__['plan'] = plan
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('equinix-metal:index/getSpotMarketPrice:getSpotMarketPrice', __args__, opts=opts, typ=GetSpotMarketPriceResult).value

    return AwaitableGetSpotMarketPriceResult(
        facility=pulumi.get(__ret__, 'facility'),
        id=pulumi.get(__ret__, 'id'),
        metro=pulumi.get(__ret__, 'metro'),
        plan=pulumi.get(__ret__, 'plan'),
        price=pulumi.get(__ret__, 'price'))


@_utilities.lift_output_func(get_spot_market_price)
def get_spot_market_price_output(facility: Optional[pulumi.Input[Optional[str]]] = None,
                                 metro: Optional[pulumi.Input[Optional[str]]] = None,
                                 plan: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSpotMarketPriceResult]:
    """
    Use this data source to get Equinix Metal Spot Market Price for a plan.

    ## Example Usage

    Lookup by facility:

    ```python
    import pulumi
    import pulumi_equinix_metal as equinix_metal

    example = equinix_metal.get_spot_market_price(facility="ny5",
        plan="c3.small.x86")
    ```

    Lookup by metro:

    ```python
    import pulumi
    import pulumi_equinix_metal as equinix_metal

    example = equinix_metal.get_spot_market_price(metro="sv",
        plan="c3.small.x86")
    ```


    :param str facility: Name of the facility.
    :param str metro: Name of the metro.
    :param str plan: Name of the plan.
    """
    ...
