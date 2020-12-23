// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";

export interface DeviceIpAddress {
    /**
     * CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
     */
    cidr?: pulumi.Input<number>;
    /**
     * String of UUID of IP block reservations from which the public IPv4 address should be taken.
     */
    reservationIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One of [`privateIpv4`, `publicIpv4`, `publicIpv6`]
     */
    type: pulumi.Input<string>;
}

export interface DeviceNetwork {
    /**
     * IPv4 or IPv6 address string
     */
    address?: pulumi.Input<string>;
    /**
     * CIDR suffix for IP address block to be assigned, i.e. amount of addresses.
     */
    cidr?: pulumi.Input<number>;
    /**
     * IP version - "4" or "6"
     * * `networkType` Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Will be one of `layer3`, `hybrid`, `layer2-individual` and `layer2-bonded`.
     */
    family?: pulumi.Input<number>;
    /**
     * address of router
     */
    gateway?: pulumi.Input<string>;
    /**
     * whether the address is routable from the Internet
     */
    public?: pulumi.Input<boolean>;
}

export interface DevicePort {
    /**
     * Whether this port is part of a bond in bonded network setup
     * * `projectId`- The ID of the project the device belongs to
     */
    bonded?: pulumi.Input<boolean>;
    /**
     * ID of the port
     */
    id?: pulumi.Input<string>;
    /**
     * MAC address assigned to the port
     */
    mac?: pulumi.Input<string>;
    /**
     * Name of the port (e.g. `eth0`, or `bond0`)
     */
    name?: pulumi.Input<string>;
    /**
     * One of [`privateIpv4`, `publicIpv4`, `publicIpv6`]
     */
    type?: pulumi.Input<string>;
}

export interface ProjectBgpConfig {
    /**
     * Autonomous System Number for local BGP deployment
     */
    asn: pulumi.Input<number>;
    /**
     * `private` or `public`, the `private` is likely to be usable immediately, the `public` will need to be review by Equinix Metal engineers
     */
    deploymentType: pulumi.Input<string>;
    /**
     * The maximum number of route filters allowed per server
     */
    maxPrefix?: pulumi.Input<number>;
    /**
     * Password for BGP session in plaintext (not a checksum)
     */
    md5?: pulumi.Input<string>;
    /**
     * status of BGP configuration in the project
     */
    status?: pulumi.Input<string>;
}

export interface SpotMarketRequestInstanceParameters {
    alwaysPxe?: pulumi.Input<boolean>;
    billingCycle: pulumi.Input<string>;
    customdata?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    features?: pulumi.Input<pulumi.Input<string>[]>;
    hostname: pulumi.Input<string>;
    ipxeScriptUrl?: pulumi.Input<string>;
    locked?: pulumi.Input<string>;
    operatingSystem: pulumi.Input<string>;
    plan: pulumi.Input<string>;
    projectSshKeys?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    termintationTime?: pulumi.Input<string>;
    userSshKeys?: pulumi.Input<pulumi.Input<string>[]>;
    userdata?: pulumi.Input<string>;
}

export interface VolumeAttachment {
    href?: pulumi.Input<string>;
}

export interface VolumeSnapshotPolicy {
    snapshotCount: pulumi.Input<number>;
    snapshotFrequency: pulumi.Input<string>;
}
