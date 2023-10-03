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
    'GetHardwareReservationResult',
    'AwaitableGetHardwareReservationResult',
    'get_hardware_reservation',
    'get_hardware_reservation_output',
]

@pulumi.output_type
class GetHardwareReservationResult:
    """
    A collection of values returned by getHardwareReservation.
    """
    def __init__(__self__, device_id=None, facility=None, id=None, plan=None, project_id=None, provisionable=None, short_id=None, spare=None, switch_uuid=None):
        if device_id and not isinstance(device_id, str):
            raise TypeError("Expected argument 'device_id' to be a str")
        pulumi.set(__self__, "device_id", device_id)
        if facility and not isinstance(facility, str):
            raise TypeError("Expected argument 'facility' to be a str")
        pulumi.set(__self__, "facility", facility)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if provisionable and not isinstance(provisionable, bool):
            raise TypeError("Expected argument 'provisionable' to be a bool")
        pulumi.set(__self__, "provisionable", provisionable)
        if short_id and not isinstance(short_id, str):
            raise TypeError("Expected argument 'short_id' to be a str")
        pulumi.set(__self__, "short_id", short_id)
        if spare and not isinstance(spare, bool):
            raise TypeError("Expected argument 'spare' to be a bool")
        pulumi.set(__self__, "spare", spare)
        if switch_uuid and not isinstance(switch_uuid, str):
            raise TypeError("Expected argument 'switch_uuid' to be a str")
        pulumi.set(__self__, "switch_uuid", switch_uuid)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> str:
        """
        UUID of device occupying the reservation
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter
    def facility(self) -> str:
        """
        Plan type for the reservation
        """
        return pulumi.get(self, "facility")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the hardware reservation to look up
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def plan(self) -> str:
        """
        Plan type for the reservation
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        UUID of project this reservation is scoped to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def provisionable(self) -> bool:
        """
        Flag indicating whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first
        """
        return pulumi.get(self, "provisionable")

    @property
    @pulumi.getter(name="shortId")
    def short_id(self) -> str:
        """
        Reservation short ID
        """
        return pulumi.get(self, "short_id")

    @property
    @pulumi.getter
    def spare(self) -> bool:
        """
        Flag indicating whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix
        """
        return pulumi.get(self, "spare")

    @property
    @pulumi.getter(name="switchUuid")
    def switch_uuid(self) -> str:
        """
        Switch short ID, can be used to determine if two devices are connected to the same switch
        """
        return pulumi.get(self, "switch_uuid")


class AwaitableGetHardwareReservationResult(GetHardwareReservationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHardwareReservationResult(
            device_id=self.device_id,
            facility=self.facility,
            id=self.id,
            plan=self.plan,
            project_id=self.project_id,
            provisionable=self.provisionable,
            short_id=self.short_id,
            spare=self.spare,
            switch_uuid=self.switch_uuid)


def get_hardware_reservation(device_id: Optional[str] = None,
                             id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHardwareReservationResult:
    """
    Use this data source to retrieve a [hardware reservation resource from Equinix Metal](https://metal.equinix.com/developers/docs/deploy/reserved/).

    You can look up hardware reservation by its ID or by ID of device which occupies it.


    :param str device_id: UUID of device occupying the reservation
    :param str id: ID of the hardware reservation
    """
    __args__ = dict()
    __args__['deviceId'] = device_id
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('equinix-metal:index/getHardwareReservation:getHardwareReservation', __args__, opts=opts, typ=GetHardwareReservationResult).value

    return AwaitableGetHardwareReservationResult(
        device_id=pulumi.get(__ret__, 'device_id'),
        facility=pulumi.get(__ret__, 'facility'),
        id=pulumi.get(__ret__, 'id'),
        plan=pulumi.get(__ret__, 'plan'),
        project_id=pulumi.get(__ret__, 'project_id'),
        provisionable=pulumi.get(__ret__, 'provisionable'),
        short_id=pulumi.get(__ret__, 'short_id'),
        spare=pulumi.get(__ret__, 'spare'),
        switch_uuid=pulumi.get(__ret__, 'switch_uuid'))


@_utilities.lift_output_func(get_hardware_reservation)
def get_hardware_reservation_output(device_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    id: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHardwareReservationResult]:
    """
    Use this data source to retrieve a [hardware reservation resource from Equinix Metal](https://metal.equinix.com/developers/docs/deploy/reserved/).

    You can look up hardware reservation by its ID or by ID of device which occupies it.


    :param str device_id: UUID of device occupying the reservation
    :param str id: ID of the hardware reservation
    """
    ...
