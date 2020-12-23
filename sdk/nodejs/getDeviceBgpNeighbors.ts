// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Use this datasource to retrieve list of BGP neighbors of a device in the Equinix Metal host.
 *
 * To have any BGP neighbors listed, the device must be in BGP-enabled project and have a BGP session assigned.
 *
 * To learn more about using BGP in Equinix Metal, see the equinix-metal.BgpSession resource documentation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const test = equinix_metal.getDeviceBgpNeighbors({
 *     deviceId: "4c641195-25e5-4c3c-b2b7-4cd7a42c7b40",
 * });
 * export const bgpNeighborsListing = test.then(test => test.bgpNeighbors);
 * ```
 */
export function getDeviceBgpNeighbors(args: GetDeviceBgpNeighborsArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceBgpNeighborsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("equinix-metal:index/getDeviceBgpNeighbors:getDeviceBgpNeighbors", {
        "deviceId": args.deviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDeviceBgpNeighbors.
 */
export interface GetDeviceBgpNeighborsArgs {
    /**
     * UUID of BGP-enabled device whose neighbors to list
     */
    readonly deviceId: string;
}

/**
 * A collection of values returned by getDeviceBgpNeighbors.
 */
export interface GetDeviceBgpNeighborsResult {
    /**
     * array of BGP neighbor records with attributes:
     */
    readonly bgpNeighbors: outputs.GetDeviceBgpNeighborsBgpNeighbor[];
    readonly deviceId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
