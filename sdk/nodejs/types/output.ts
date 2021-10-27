// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";

export interface ConnectionPort {
    id: string;
    linkStatus: string;
    /**
     * Name of the connection resource
     */
    name: string;
    role: string;
    /**
     * Port speed in bits per second
     */
    speed: number;
    /**
     * Status of the connection resource
     */
    status: string;
    virtualCircuitIds: any[];
}

export interface DeviceIpAddress {
    /**
     * CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
     */
    cidr?: number;
    /**
     * List of UUIDs of IP block reservations from which the public IPv4 address should be taken.
     */
    reservationIds?: string[];
    /**
     * One of [`privateIpv4`, `publicIpv4`, `publicIpv6`]
     */
    type: string;
}

export interface DeviceNetwork {
    /**
     * IPv4 or IPv6 address string
     */
    address: string;
    /**
     * CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
     */
    cidr: number;
    /**
     * IP version - "4" or "6"
     * * `networkType` Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Will be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
     */
    family: number;
    /**
     * address of router
     */
    gateway: string;
    /**
     * whether the address is routable from the Internet
     */
    public: boolean;
}

export interface DevicePort {
    /**
     * Whether this port is part of a bond in bonded network setup
     */
    bonded: boolean;
    /**
     * ID of the port
     */
    id: string;
    /**
     * MAC address assigned to the port
     */
    mac: string;
    /**
     * Name of the port (e.g. `eth0`, or `bond0`)
     */
    name: string;
    /**
     * One of [`privateIpv4`, `publicIpv4`, `publicIpv6`]
     */
    type: string;
}

export interface DeviceReinstall {
    /**
     * Whether the OS disk should be filled with `00h` bytes before reinstall. Defaults to `false`.
     */
    deprovisionFast?: boolean;
    /**
     * Whether the provider should favour reinstall over destroy and create. Defaults to `false`.
     */
    enabled?: boolean;
    /**
     * Whether the non-OS disks should be kept or wiped during reinstall. Defaults to `false`.
     */
    preserveData?: boolean;
}

export interface GetConnectionPort {
    /**
     * Port UUID
     */
    id: string;
    /**
     * Port link status
     */
    linkStatus: string;
    /**
     * Port name
     */
    name: string;
    /**
     * Port role - primary or secondary
     */
    role: string;
    /**
     * Port speed in bits per second
     */
    speed: number;
    /**
     * Port status
     */
    status: string;
    /**
     * List of IDs of virtual cicruits attached to this port
     */
    virtualCircuitIds: any[];
}

export interface GetDeviceBgpNeighborsBgpNeighbor {
    /**
     * IP address version, 4 or 6
     */
    addressFamily: number;
    /**
     * Local autonomous system number
     */
    customerAs: number;
    /**
     * Local used peer IP address
     */
    customerIp: string;
    /**
     * Whether BGP session is password enabled
     */
    md5Enabled: boolean;
    /**
     * BGP session password in plaintext (not a checksum)
     */
    md5Password: string;
    /**
     * Whether the neighbor is in EBGP multihop session
     */
    multihop: boolean;
    /**
     * Peer AS number (different than customerAs for EBGP)
     */
    peerAs: number;
    /**
     * Array of IP addresses of this neighbor's peers
     */
    peerIps?: string[];
    /**
     * Array of incoming routes. Each route has attributes:
     */
    routesIns: outputs.GetDeviceBgpNeighborsBgpNeighborRoutesIn[];
    /**
     * Array of outgoing routes in the same format
     */
    routesOuts: outputs.GetDeviceBgpNeighborsBgpNeighborRoutesOut[];
}

export interface GetDeviceBgpNeighborsBgpNeighborRoutesIn {
    /**
     * (bool) Whether the route is exact
     */
    exact: boolean;
    /**
     * CIDR expression of route (IP/mask)
     */
    route: string;
}

export interface GetDeviceBgpNeighborsBgpNeighborRoutesOut {
    /**
     * (bool) Whether the route is exact
     */
    exact: boolean;
    /**
     * CIDR expression of route (IP/mask)
     */
    route: string;
}

export interface GetDeviceNetwork {
    /**
     * IPv4 or IPv6 address string
     */
    address: string;
    /**
     * Bit length of the network mask of the address
     */
    cidr: number;
    /**
     * IP version - "4" or "6"
     */
    family: number;
    /**
     * Address of router
     */
    gateway: string;
    /**
     * Whether the address is routable from the Internet
     */
    public: boolean;
}

export interface GetDevicePort {
    /**
     * Whether this port is part of a bond in bonded network setup
     */
    bonded: boolean;
    /**
     * ID of the port
     */
    id: string;
    /**
     * MAC address assigned to the port
     */
    mac: string;
    /**
     * Name of the port (e.g. `eth0`, or `bond0`)
     */
    name: string;
    /**
     * Type of the port (e.g. `NetworkPort` or `NetworkBondPort`)
     */
    type: string;
}

export interface GetFacilityCapacity {
    /**
     * device plan to check
     */
    plan: string;
    /**
     * number of device to check
     */
    quantity?: number;
}

export interface GetMetroCapacity {
    /**
     * device plan to check
     */
    plan: string;
    /**
     * number of device to check
     */
    quantity?: number;
}

export interface GetProjectBgpConfig {
    /**
     * Autonomous System Number for local BGP deployment
     */
    asn: number;
    /**
     * `private` or `public`, the `private` is likely to be usable immediately, the `public` will need to be review by Equinix Metal engineers
     */
    deploymentType: string;
    /**
     * The maximum number of route filters allowed per server
     */
    maxPrefix: number;
    /**
     * Password for BGP session in plaintext (not a checksum)
     */
    md5?: string;
    /**
     * status of BGP configuration in the project
     */
    status: string;
}

export interface GetVolumeSnapshotPolicy {
    snapshotCount: number;
    snapshotFrequency: string;
}

export interface ProjectBgpConfig {
    /**
     * Autonomous System Number for local BGP deployment
     */
    asn: number;
    /**
     * `private` or `public`, the `private` is likely to be usable immediately, the `public` will need to be review by Equinix Metal engineers
     */
    deploymentType: string;
    /**
     * The maximum number of route filters allowed per server
     */
    maxPrefix: number;
    /**
     * Password for BGP session in plaintext (not a checksum)
     */
    md5?: string;
    /**
     * status of BGP configuration in the project
     */
    status: string;
}

export interface SpotMarketRequestInstanceParameters {
    alwaysPxe?: boolean;
    billingCycle: string;
    customdata?: string;
    description?: string;
    features?: string[];
    hostname: string;
    ipxeScriptUrl?: string;
    /**
     * Blocks deletion of the SpotMarketRequest device until the lock is disabled
     */
    locked?: boolean;
    operatingSystem: string;
    plan: string;
    projectSshKeys?: string[];
    tags?: string[];
    termintationTime: string;
    userSshKeys?: string[];
    userdata?: string;
}

export interface VolumeAttachment {
    href: string;
}

export interface VolumeSnapshotPolicy {
    snapshotCount: number;
    snapshotFrequency: string;
}
