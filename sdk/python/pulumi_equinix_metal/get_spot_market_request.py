# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSpotMarketRequestResult',
    'AwaitableGetSpotMarketRequestResult',
    'get_spot_market_request',
    'get_spot_market_request_output',
]

@pulumi.output_type
class GetSpotMarketRequestResult:
    """
    A collection of values returned by getSpotMarketRequest.
    """
    def __init__(__self__, device_ids=None, devices_max=None, devices_min=None, end_at=None, facilities=None, id=None, max_bid_price=None, metro=None, plan=None, project_id=None, request_id=None):
        if device_ids and not isinstance(device_ids, list):
            raise TypeError("Expected argument 'device_ids' to be a list")
        pulumi.set(__self__, "device_ids", device_ids)
        if devices_max and not isinstance(devices_max, int):
            raise TypeError("Expected argument 'devices_max' to be a int")
        pulumi.set(__self__, "devices_max", devices_max)
        if devices_min and not isinstance(devices_min, int):
            raise TypeError("Expected argument 'devices_min' to be a int")
        pulumi.set(__self__, "devices_min", devices_min)
        if end_at and not isinstance(end_at, str):
            raise TypeError("Expected argument 'end_at' to be a str")
        pulumi.set(__self__, "end_at", end_at)
        if facilities and not isinstance(facilities, list):
            raise TypeError("Expected argument 'facilities' to be a list")
        pulumi.set(__self__, "facilities", facilities)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if max_bid_price and not isinstance(max_bid_price, float):
            raise TypeError("Expected argument 'max_bid_price' to be a float")
        pulumi.set(__self__, "max_bid_price", max_bid_price)
        if metro and not isinstance(metro, str):
            raise TypeError("Expected argument 'metro' to be a str")
        pulumi.set(__self__, "metro", metro)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if request_id and not isinstance(request_id, str):
            raise TypeError("Expected argument 'request_id' to be a str")
        pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="deviceIds")
    def device_ids(self) -> Sequence[str]:
        """
        List of IDs of devices spawned by the referenced Spot Market Request
        """
        return pulumi.get(self, "device_ids")

    @property
    @pulumi.getter(name="devicesMax")
    def devices_max(self) -> int:
        """
        Maximum number devices to be created
        """
        return pulumi.get(self, "devices_max")

    @property
    @pulumi.getter(name="devicesMin")
    def devices_min(self) -> int:
        """
        Miniumum number devices to be created
        """
        return pulumi.get(self, "devices_min")

    @property
    @pulumi.getter(name="endAt")
    def end_at(self) -> str:
        """
        Date and time When the spot market request will be ended.
        """
        return pulumi.get(self, "end_at")

    @property
    @pulumi.getter
    def facilities(self) -> Sequence[str]:
        """
        Facility IDs where devices should be created
        """
        return pulumi.get(self, "facilities")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maxBidPrice")
    def max_bid_price(self) -> float:
        """
        Maximum price user is willing to pay per hour per device
        """
        return pulumi.get(self, "max_bid_price")

    @property
    @pulumi.getter
    def metro(self) -> str:
        """
        Metro where devices should be created.
        """
        return pulumi.get(self, "metro")

    @property
    @pulumi.getter
    def plan(self) -> str:
        """
        The device plan slug.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> str:
        return pulumi.get(self, "request_id")


class AwaitableGetSpotMarketRequestResult(GetSpotMarketRequestResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSpotMarketRequestResult(
            device_ids=self.device_ids,
            devices_max=self.devices_max,
            devices_min=self.devices_min,
            end_at=self.end_at,
            facilities=self.facilities,
            id=self.id,
            max_bid_price=self.max_bid_price,
            metro=self.metro,
            plan=self.plan,
            project_id=self.project_id,
            request_id=self.request_id)


def get_spot_market_request(request_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSpotMarketRequestResult:
    """
    Use this data source to access information about an existing resource.

    :param str request_id: The id of the Spot Market Request
    """
    __args__ = dict()
    __args__['requestId'] = request_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('equinix-metal:index/getSpotMarketRequest:getSpotMarketRequest', __args__, opts=opts, typ=GetSpotMarketRequestResult).value

    return AwaitableGetSpotMarketRequestResult(
        device_ids=pulumi.get(__ret__, 'device_ids'),
        devices_max=pulumi.get(__ret__, 'devices_max'),
        devices_min=pulumi.get(__ret__, 'devices_min'),
        end_at=pulumi.get(__ret__, 'end_at'),
        facilities=pulumi.get(__ret__, 'facilities'),
        id=pulumi.get(__ret__, 'id'),
        max_bid_price=pulumi.get(__ret__, 'max_bid_price'),
        metro=pulumi.get(__ret__, 'metro'),
        plan=pulumi.get(__ret__, 'plan'),
        project_id=pulumi.get(__ret__, 'project_id'),
        request_id=pulumi.get(__ret__, 'request_id'))


@_utilities.lift_output_func(get_spot_market_request)
def get_spot_market_request_output(request_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSpotMarketRequestResult]:
    """
    Use this data source to access information about an existing resource.

    :param str request_id: The id of the Spot Market Request
    """
    ...
