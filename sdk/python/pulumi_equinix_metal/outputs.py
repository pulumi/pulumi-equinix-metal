# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'ConnectionPort',
    'DeviceIpAddress',
    'DeviceNetwork',
    'DevicePort',
    'DeviceReinstall',
    'ProjectBgpConfig',
    'SpotMarketRequestInstanceParameters',
    'VolumeAttachment',
    'VolumeSnapshotPolicy',
    'GetConnectionPortResult',
    'GetDeviceBgpNeighborsBgpNeighborResult',
    'GetDeviceBgpNeighborsBgpNeighborRoutesInResult',
    'GetDeviceBgpNeighborsBgpNeighborRoutesOutResult',
    'GetDeviceNetworkResult',
    'GetDevicePortResult',
    'GetProjectBgpConfigResult',
    'GetVolumeSnapshotPolicyResult',
]

@pulumi.output_type
class ConnectionPort(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "linkStatus":
            suggest = "link_status"
        elif key == "virtualCircuitIds":
            suggest = "virtual_circuit_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConnectionPort. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConnectionPort.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConnectionPort.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: Optional[str] = None,
                 link_status: Optional[str] = None,
                 name: Optional[str] = None,
                 role: Optional[str] = None,
                 speed: Optional[int] = None,
                 status: Optional[str] = None,
                 virtual_circuit_ids: Optional[Sequence[Any]] = None):
        """
        :param str name: Name of the connection resource
        :param int speed: Port speed in bits per second
        :param str status: Status of the connection resource
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if link_status is not None:
            pulumi.set(__self__, "link_status", link_status)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if speed is not None:
            pulumi.set(__self__, "speed", speed)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if virtual_circuit_ids is not None:
            pulumi.set(__self__, "virtual_circuit_ids", virtual_circuit_ids)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="linkStatus")
    def link_status(self) -> Optional[str]:
        return pulumi.get(self, "link_status")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the connection resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def speed(self) -> Optional[int]:
        """
        Port speed in bits per second
        """
        return pulumi.get(self, "speed")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the connection resource
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="virtualCircuitIds")
    def virtual_circuit_ids(self) -> Optional[Sequence[Any]]:
        return pulumi.get(self, "virtual_circuit_ids")


@pulumi.output_type
class DeviceIpAddress(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "reservationIds":
            suggest = "reservation_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DeviceIpAddress. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DeviceIpAddress.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DeviceIpAddress.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 cidr: Optional[int] = None,
                 reservation_ids: Optional[Sequence[str]] = None):
        """
        :param str type: One of [`private_ipv4`, `public_ipv4`, `public_ipv6`]
        :param int cidr: CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
        :param Sequence[str] reservation_ids: List of UUIDs of IP block reservations from which the public IPv4 address should be taken.
        """
        pulumi.set(__self__, "type", type)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if reservation_ids is not None:
            pulumi.set(__self__, "reservation_ids", reservation_ids)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        One of [`private_ipv4`, `public_ipv4`, `public_ipv6`]
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def cidr(self) -> Optional[int]:
        """
        CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="reservationIds")
    def reservation_ids(self) -> Optional[Sequence[str]]:
        """
        List of UUIDs of IP block reservations from which the public IPv4 address should be taken.
        """
        return pulumi.get(self, "reservation_ids")


@pulumi.output_type
class DeviceNetwork(dict):
    def __init__(__self__, *,
                 address: Optional[str] = None,
                 cidr: Optional[int] = None,
                 family: Optional[int] = None,
                 gateway: Optional[str] = None,
                 public: Optional[bool] = None):
        """
        :param str address: IPv4 or IPv6 address string
        :param int cidr: CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
        :param int family: IP version - "4" or "6"
               * `network_type` Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Will be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
        :param str gateway: address of router
        :param bool public: whether the address is routable from the Internet
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if family is not None:
            pulumi.set(__self__, "family", family)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if public is not None:
            pulumi.set(__self__, "public", public)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        """
        IPv4 or IPv6 address string
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def cidr(self) -> Optional[int]:
        """
        CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def family(self) -> Optional[int]:
        """
        IP version - "4" or "6"
        * `network_type` Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Will be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def gateway(self) -> Optional[str]:
        """
        address of router
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def public(self) -> Optional[bool]:
        """
        whether the address is routable from the Internet
        """
        return pulumi.get(self, "public")


@pulumi.output_type
class DevicePort(dict):
    def __init__(__self__, *,
                 bonded: Optional[bool] = None,
                 id: Optional[str] = None,
                 mac: Optional[str] = None,
                 name: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param bool bonded: Whether this port is part of a bond in bonded network setup
        :param str id: ID of the port
        :param str mac: MAC address assigned to the port
        :param str name: Name of the port (e.g. `eth0`, or `bond0`)
        :param str type: One of [`private_ipv4`, `public_ipv4`, `public_ipv6`]
        """
        if bonded is not None:
            pulumi.set(__self__, "bonded", bonded)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def bonded(self) -> Optional[bool]:
        """
        Whether this port is part of a bond in bonded network setup
        """
        return pulumi.get(self, "bonded")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        ID of the port
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mac(self) -> Optional[str]:
        """
        MAC address assigned to the port
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the port (e.g. `eth0`, or `bond0`)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        One of [`private_ipv4`, `public_ipv4`, `public_ipv6`]
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class DeviceReinstall(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deprovisionFast":
            suggest = "deprovision_fast"
        elif key == "preserveData":
            suggest = "preserve_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DeviceReinstall. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DeviceReinstall.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DeviceReinstall.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 deprovision_fast: Optional[bool] = None,
                 enabled: Optional[bool] = None,
                 preserve_data: Optional[bool] = None):
        """
        :param bool deprovision_fast: Whether the OS disk should be filled with `00h` bytes before reinstall. Defaults to `false`.
               *
        :param bool enabled: Whether the provider should favour reinstall over destroy and create. Defaults to `false`.
        :param bool preserve_data: Whether the non-OS disks should be kept or wiped during reinstall. Defaults to `false`.
        """
        if deprovision_fast is not None:
            pulumi.set(__self__, "deprovision_fast", deprovision_fast)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if preserve_data is not None:
            pulumi.set(__self__, "preserve_data", preserve_data)

    @property
    @pulumi.getter(name="deprovisionFast")
    def deprovision_fast(self) -> Optional[bool]:
        """
        Whether the OS disk should be filled with `00h` bytes before reinstall. Defaults to `false`.
        *
        """
        return pulumi.get(self, "deprovision_fast")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether the provider should favour reinstall over destroy and create. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="preserveData")
    def preserve_data(self) -> Optional[bool]:
        """
        Whether the non-OS disks should be kept or wiped during reinstall. Defaults to `false`.
        """
        return pulumi.get(self, "preserve_data")


@pulumi.output_type
class ProjectBgpConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deploymentType":
            suggest = "deployment_type"
        elif key == "maxPrefix":
            suggest = "max_prefix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectBgpConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectBgpConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectBgpConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 asn: int,
                 deployment_type: str,
                 max_prefix: Optional[int] = None,
                 md5: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param int asn: Autonomous System Number for local BGP deployment
        :param str deployment_type: `private` or `public`, the `private` is likely to be usable immediately, the `public` will need to be review by Equinix Metal engineers
        :param int max_prefix: The maximum number of route filters allowed per server
        :param str md5: Password for BGP session in plaintext (not a checksum)
        :param str status: status of BGP configuration in the project
        """
        pulumi.set(__self__, "asn", asn)
        pulumi.set(__self__, "deployment_type", deployment_type)
        if max_prefix is not None:
            pulumi.set(__self__, "max_prefix", max_prefix)
        if md5 is not None:
            pulumi.set(__self__, "md5", md5)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def asn(self) -> int:
        """
        Autonomous System Number for local BGP deployment
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> str:
        """
        `private` or `public`, the `private` is likely to be usable immediately, the `public` will need to be review by Equinix Metal engineers
        """
        return pulumi.get(self, "deployment_type")

    @property
    @pulumi.getter(name="maxPrefix")
    def max_prefix(self) -> Optional[int]:
        """
        The maximum number of route filters allowed per server
        """
        return pulumi.get(self, "max_prefix")

    @property
    @pulumi.getter
    def md5(self) -> Optional[str]:
        """
        Password for BGP session in plaintext (not a checksum)
        """
        return pulumi.get(self, "md5")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        status of BGP configuration in the project
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class SpotMarketRequestInstanceParameters(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "billingCycle":
            suggest = "billing_cycle"
        elif key == "operatingSystem":
            suggest = "operating_system"
        elif key == "alwaysPxe":
            suggest = "always_pxe"
        elif key == "ipxeScriptUrl":
            suggest = "ipxe_script_url"
        elif key == "projectSshKeys":
            suggest = "project_ssh_keys"
        elif key == "termintationTime":
            suggest = "termintation_time"
        elif key == "userSshKeys":
            suggest = "user_ssh_keys"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SpotMarketRequestInstanceParameters. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SpotMarketRequestInstanceParameters.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SpotMarketRequestInstanceParameters.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 billing_cycle: str,
                 hostname: str,
                 operating_system: str,
                 plan: str,
                 always_pxe: Optional[bool] = None,
                 customdata: Optional[str] = None,
                 description: Optional[str] = None,
                 features: Optional[Sequence[str]] = None,
                 ipxe_script_url: Optional[str] = None,
                 locked: Optional[bool] = None,
                 project_ssh_keys: Optional[Sequence[str]] = None,
                 tags: Optional[Sequence[str]] = None,
                 termintation_time: Optional[str] = None,
                 user_ssh_keys: Optional[Sequence[str]] = None,
                 userdata: Optional[str] = None):
        """
        :param bool locked: Blocks deletion of the SpotMarketRequest device until the lock is disabled
        """
        pulumi.set(__self__, "billing_cycle", billing_cycle)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "operating_system", operating_system)
        pulumi.set(__self__, "plan", plan)
        if always_pxe is not None:
            pulumi.set(__self__, "always_pxe", always_pxe)
        if customdata is not None:
            pulumi.set(__self__, "customdata", customdata)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if features is not None:
            pulumi.set(__self__, "features", features)
        if ipxe_script_url is not None:
            pulumi.set(__self__, "ipxe_script_url", ipxe_script_url)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if project_ssh_keys is not None:
            pulumi.set(__self__, "project_ssh_keys", project_ssh_keys)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if termintation_time is not None:
            pulumi.set(__self__, "termintation_time", termintation_time)
        if user_ssh_keys is not None:
            pulumi.set(__self__, "user_ssh_keys", user_ssh_keys)
        if userdata is not None:
            pulumi.set(__self__, "userdata", userdata)

    @property
    @pulumi.getter(name="billingCycle")
    def billing_cycle(self) -> str:
        return pulumi.get(self, "billing_cycle")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="operatingSystem")
    def operating_system(self) -> str:
        return pulumi.get(self, "operating_system")

    @property
    @pulumi.getter
    def plan(self) -> str:
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="alwaysPxe")
    def always_pxe(self) -> Optional[bool]:
        return pulumi.get(self, "always_pxe")

    @property
    @pulumi.getter
    def customdata(self) -> Optional[str]:
        return pulumi.get(self, "customdata")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def features(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="ipxeScriptUrl")
    def ipxe_script_url(self) -> Optional[str]:
        return pulumi.get(self, "ipxe_script_url")

    @property
    @pulumi.getter
    def locked(self) -> Optional[bool]:
        """
        Blocks deletion of the SpotMarketRequest device until the lock is disabled
        """
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter(name="projectSshKeys")
    def project_ssh_keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "project_ssh_keys")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="termintationTime")
    def termintation_time(self) -> Optional[str]:
        return pulumi.get(self, "termintation_time")

    @property
    @pulumi.getter(name="userSshKeys")
    def user_ssh_keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "user_ssh_keys")

    @property
    @pulumi.getter
    def userdata(self) -> Optional[str]:
        return pulumi.get(self, "userdata")


@pulumi.output_type
class VolumeAttachment(dict):
    def __init__(__self__, *,
                 href: Optional[str] = None):
        if href is not None:
            pulumi.set(__self__, "href", href)

    @property
    @pulumi.getter
    def href(self) -> Optional[str]:
        return pulumi.get(self, "href")


@pulumi.output_type
class VolumeSnapshotPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "snapshotCount":
            suggest = "snapshot_count"
        elif key == "snapshotFrequency":
            suggest = "snapshot_frequency"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeSnapshotPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeSnapshotPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeSnapshotPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 snapshot_count: int,
                 snapshot_frequency: str):
        pulumi.set(__self__, "snapshot_count", snapshot_count)
        pulumi.set(__self__, "snapshot_frequency", snapshot_frequency)

    @property
    @pulumi.getter(name="snapshotCount")
    def snapshot_count(self) -> int:
        return pulumi.get(self, "snapshot_count")

    @property
    @pulumi.getter(name="snapshotFrequency")
    def snapshot_frequency(self) -> str:
        return pulumi.get(self, "snapshot_frequency")


@pulumi.output_type
class GetConnectionPortResult(dict):
    def __init__(__self__, *,
                 id: str,
                 link_status: str,
                 name: str,
                 role: str,
                 speed: int,
                 status: str,
                 virtual_circuit_ids: Sequence[Any]):
        """
        :param str id: Port UUID
        :param str link_status: Port link status
        :param str name: Port name
        :param str role: Port role - primary or secondary
        :param int speed: Port speed in bits per second
        :param str status: Port status
        :param Sequence[Any] virtual_circuit_ids: List of IDs of virtual cicruits attached to this port
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "link_status", link_status)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "speed", speed)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "virtual_circuit_ids", virtual_circuit_ids)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Port UUID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="linkStatus")
    def link_status(self) -> str:
        """
        Port link status
        """
        return pulumi.get(self, "link_status")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Port name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Port role - primary or secondary
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def speed(self) -> int:
        """
        Port speed in bits per second
        """
        return pulumi.get(self, "speed")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Port status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="virtualCircuitIds")
    def virtual_circuit_ids(self) -> Sequence[Any]:
        """
        List of IDs of virtual cicruits attached to this port
        """
        return pulumi.get(self, "virtual_circuit_ids")


@pulumi.output_type
class GetDeviceBgpNeighborsBgpNeighborResult(dict):
    def __init__(__self__, *,
                 address_family: int,
                 customer_as: int,
                 customer_ip: str,
                 md5_enabled: bool,
                 md5_password: str,
                 multihop: bool,
                 peer_as: int,
                 routes_ins: Sequence['outputs.GetDeviceBgpNeighborsBgpNeighborRoutesInResult'],
                 routes_outs: Sequence['outputs.GetDeviceBgpNeighborsBgpNeighborRoutesOutResult'],
                 peer_ips: Optional[Sequence[str]] = None):
        """
        :param int address_family: IP address version, 4 or 6
        :param int customer_as: Local autonomous system number
        :param str customer_ip: Local used peer IP address
        :param bool md5_enabled: Whether BGP session is password enabled
        :param str md5_password: BGP session password in plaintext (not a checksum)
        :param bool multihop: Whether the neighbor is in EBGP multihop session
        :param int peer_as: Peer AS number (different than customer_as for EBGP)
        :param Sequence['GetDeviceBgpNeighborsBgpNeighborRoutesInArgs'] routes_ins: Array of incoming routes. Each route has attributes:
        :param Sequence['GetDeviceBgpNeighborsBgpNeighborRoutesOutArgs'] routes_outs: Array of outgoing routes in the same format
        :param Sequence[str] peer_ips: Array of IP addresses of this neighbor's peers
        """
        pulumi.set(__self__, "address_family", address_family)
        pulumi.set(__self__, "customer_as", customer_as)
        pulumi.set(__self__, "customer_ip", customer_ip)
        pulumi.set(__self__, "md5_enabled", md5_enabled)
        pulumi.set(__self__, "md5_password", md5_password)
        pulumi.set(__self__, "multihop", multihop)
        pulumi.set(__self__, "peer_as", peer_as)
        pulumi.set(__self__, "routes_ins", routes_ins)
        pulumi.set(__self__, "routes_outs", routes_outs)
        if peer_ips is not None:
            pulumi.set(__self__, "peer_ips", peer_ips)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> int:
        """
        IP address version, 4 or 6
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter(name="customerAs")
    def customer_as(self) -> int:
        """
        Local autonomous system number
        """
        return pulumi.get(self, "customer_as")

    @property
    @pulumi.getter(name="customerIp")
    def customer_ip(self) -> str:
        """
        Local used peer IP address
        """
        return pulumi.get(self, "customer_ip")

    @property
    @pulumi.getter(name="md5Enabled")
    def md5_enabled(self) -> bool:
        """
        Whether BGP session is password enabled
        """
        return pulumi.get(self, "md5_enabled")

    @property
    @pulumi.getter(name="md5Password")
    def md5_password(self) -> str:
        """
        BGP session password in plaintext (not a checksum)
        """
        return pulumi.get(self, "md5_password")

    @property
    @pulumi.getter
    def multihop(self) -> bool:
        """
        Whether the neighbor is in EBGP multihop session
        """
        return pulumi.get(self, "multihop")

    @property
    @pulumi.getter(name="peerAs")
    def peer_as(self) -> int:
        """
        Peer AS number (different than customer_as for EBGP)
        """
        return pulumi.get(self, "peer_as")

    @property
    @pulumi.getter(name="routesIns")
    def routes_ins(self) -> Sequence['outputs.GetDeviceBgpNeighborsBgpNeighborRoutesInResult']:
        """
        Array of incoming routes. Each route has attributes:
        """
        return pulumi.get(self, "routes_ins")

    @property
    @pulumi.getter(name="routesOuts")
    def routes_outs(self) -> Sequence['outputs.GetDeviceBgpNeighborsBgpNeighborRoutesOutResult']:
        """
        Array of outgoing routes in the same format
        """
        return pulumi.get(self, "routes_outs")

    @property
    @pulumi.getter(name="peerIps")
    def peer_ips(self) -> Optional[Sequence[str]]:
        """
        Array of IP addresses of this neighbor's peers
        """
        return pulumi.get(self, "peer_ips")


@pulumi.output_type
class GetDeviceBgpNeighborsBgpNeighborRoutesInResult(dict):
    def __init__(__self__, *,
                 exact: bool,
                 route: str):
        """
        :param bool exact: (bool) Whether the route is exact
        :param str route: CIDR expression of route (IP/mask)
        """
        pulumi.set(__self__, "exact", exact)
        pulumi.set(__self__, "route", route)

    @property
    @pulumi.getter
    def exact(self) -> bool:
        """
        (bool) Whether the route is exact
        """
        return pulumi.get(self, "exact")

    @property
    @pulumi.getter
    def route(self) -> str:
        """
        CIDR expression of route (IP/mask)
        """
        return pulumi.get(self, "route")


@pulumi.output_type
class GetDeviceBgpNeighborsBgpNeighborRoutesOutResult(dict):
    def __init__(__self__, *,
                 exact: bool,
                 route: str):
        """
        :param bool exact: (bool) Whether the route is exact
        :param str route: CIDR expression of route (IP/mask)
        """
        pulumi.set(__self__, "exact", exact)
        pulumi.set(__self__, "route", route)

    @property
    @pulumi.getter
    def exact(self) -> bool:
        """
        (bool) Whether the route is exact
        """
        return pulumi.get(self, "exact")

    @property
    @pulumi.getter
    def route(self) -> str:
        """
        CIDR expression of route (IP/mask)
        """
        return pulumi.get(self, "route")


@pulumi.output_type
class GetDeviceNetworkResult(dict):
    def __init__(__self__, *,
                 address: str,
                 cidr: int,
                 family: int,
                 gateway: str,
                 public: bool):
        """
        :param str address: IPv4 or IPv6 address string
        :param int cidr: Bit length of the network mask of the address
        :param int family: IP version - "4" or "6"
        :param str gateway: Address of router
        :param bool public: Whether the address is routable from the Internet
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "cidr", cidr)
        pulumi.set(__self__, "family", family)
        pulumi.set(__self__, "gateway", gateway)
        pulumi.set(__self__, "public", public)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        IPv4 or IPv6 address string
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def cidr(self) -> int:
        """
        Bit length of the network mask of the address
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def family(self) -> int:
        """
        IP version - "4" or "6"
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def gateway(self) -> str:
        """
        Address of router
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def public(self) -> bool:
        """
        Whether the address is routable from the Internet
        """
        return pulumi.get(self, "public")


@pulumi.output_type
class GetDevicePortResult(dict):
    def __init__(__self__, *,
                 bonded: bool,
                 id: str,
                 mac: str,
                 name: str,
                 type: str):
        """
        :param bool bonded: Whether this port is part of a bond in bonded network setup
        :param str id: ID of the port
        :param str mac: MAC address assigned to the port
        :param str name: Name of the port (e.g. `eth0`, or `bond0`)
        :param str type: Type of the port (e.g. `NetworkPort` or `NetworkBondPort`)
        """
        pulumi.set(__self__, "bonded", bonded)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "mac", mac)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def bonded(self) -> bool:
        """
        Whether this port is part of a bond in bonded network setup
        """
        return pulumi.get(self, "bonded")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the port
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mac(self) -> str:
        """
        MAC address assigned to the port
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the port (e.g. `eth0`, or `bond0`)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of the port (e.g. `NetworkPort` or `NetworkBondPort`)
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetProjectBgpConfigResult(dict):
    def __init__(__self__, *,
                 asn: int,
                 deployment_type: str,
                 max_prefix: int,
                 status: str,
                 md5: Optional[str] = None):
        """
        :param int asn: Autonomous System Number for local BGP deployment
        :param str deployment_type: `private` or `public`, the `private` is likely to be usable immediately, the `public` will need to be review by Equinix Metal engineers
        :param int max_prefix: The maximum number of route filters allowed per server
        :param str status: status of BGP configuration in the project
        :param str md5: Password for BGP session in plaintext (not a checksum)
        """
        pulumi.set(__self__, "asn", asn)
        pulumi.set(__self__, "deployment_type", deployment_type)
        pulumi.set(__self__, "max_prefix", max_prefix)
        pulumi.set(__self__, "status", status)
        if md5 is not None:
            pulumi.set(__self__, "md5", md5)

    @property
    @pulumi.getter
    def asn(self) -> int:
        """
        Autonomous System Number for local BGP deployment
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> str:
        """
        `private` or `public`, the `private` is likely to be usable immediately, the `public` will need to be review by Equinix Metal engineers
        """
        return pulumi.get(self, "deployment_type")

    @property
    @pulumi.getter(name="maxPrefix")
    def max_prefix(self) -> int:
        """
        The maximum number of route filters allowed per server
        """
        return pulumi.get(self, "max_prefix")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        status of BGP configuration in the project
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def md5(self) -> Optional[str]:
        """
        Password for BGP session in plaintext (not a checksum)
        """
        return pulumi.get(self, "md5")


@pulumi.output_type
class GetVolumeSnapshotPolicyResult(dict):
    def __init__(__self__, *,
                 snapshot_count: int,
                 snapshot_frequency: str):
        pulumi.set(__self__, "snapshot_count", snapshot_count)
        pulumi.set(__self__, "snapshot_frequency", snapshot_frequency)

    @property
    @pulumi.getter(name="snapshotCount")
    def snapshot_count(self) -> int:
        return pulumi.get(self, "snapshot_count")

    @property
    @pulumi.getter(name="snapshotFrequency")
    def snapshot_frequency(self) -> str:
        return pulumi.get(self, "snapshot_frequency")


