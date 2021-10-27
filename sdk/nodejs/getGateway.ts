// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * // Create Metal Gateway for a VLAN with a private IPv4 block with 8 IP addresses
 * const testVlan = new equinix_metal.Vlan("testVlan", {
 *     description: "test VLAN in SV",
 *     metro: "sv",
 *     projectId: local.project_id,
 * });
 * const testGateway = equinix_metal.getGateway({
 *     gatewayId: local.gateway_id,
 * });
 * ```
 */
export function getGateway(args: GetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("equinix-metal:index/getGateway:getGateway", {
        "gatewayId": args.gatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGateway.
 */
export interface GetGatewayArgs {
    /**
     * UUID of the metal gateway resource to retrieve
     */
    readonly gatewayId: string;
}

/**
 * A collection of values returned by getGateway.
 */
export interface GetGatewayResult {
    readonly gatewayId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * UUID of IP reservation block bound to the gateway
     */
    readonly ipReservationId: string;
    /**
     * Size of the private IPv4 subnet bound to this metal gateway, one of (8, 16, 32, 64, 128)`
     */
    readonly privateIpv4SubnetSize: number;
    /**
     * UUID of the project where the gateway is scoped to
     */
    readonly projectId: string;
    /**
     * Status of the gateway resource
     */
    readonly state: string;
    /**
     * UUID of the VLAN where the gateway is scoped to
     */
    readonly vlanId: string;
}