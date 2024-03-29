// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to find IP address blocks in Equinix Metal. You can use IP address or a block ID for lookup.
 */
export function getReservedIpBlock(args?: GetReservedIpBlockArgs, opts?: pulumi.InvokeOptions): Promise<GetReservedIpBlockResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("equinix-metal:index/getReservedIpBlock:getReservedIpBlock", {
        "id": args.id,
        "ipAddress": args.ipAddress,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getReservedIpBlock.
 */
export interface GetReservedIpBlockArgs {
    /**
     * UUID of the IP address block to look up
     */
    id?: string;
    /**
     * Block containing this IP address will be returned
     */
    ipAddress?: string;
    /**
     * UUID of the project where the searched block should be
     */
    projectId?: string;
}

/**
 * A collection of values returned by getReservedIpBlock.
 */
export interface GetReservedIpBlockResult {
    readonly address: string;
    readonly addressFamily: number;
    readonly cidr: number;
    readonly cidrNotation: string;
    readonly facility: string;
    readonly gateway: string;
    readonly global: boolean;
    readonly id: string;
    readonly ipAddress?: string;
    readonly manageable: boolean;
    readonly management: boolean;
    readonly metro: string;
    readonly netmask: string;
    readonly network: string;
    readonly projectId: string;
    readonly public: boolean;
    readonly quantity: number;
    readonly type: string;
}
/**
 * Use this data source to find IP address blocks in Equinix Metal. You can use IP address or a block ID for lookup.
 */
export function getReservedIpBlockOutput(args?: GetReservedIpBlockOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReservedIpBlockResult> {
    return pulumi.output(args).apply((a: any) => getReservedIpBlock(a, opts))
}

/**
 * A collection of arguments for invoking getReservedIpBlock.
 */
export interface GetReservedIpBlockOutputArgs {
    /**
     * UUID of the IP address block to look up
     */
    id?: pulumi.Input<string>;
    /**
     * Block containing this IP address will be returned
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * UUID of the project where the searched block should be
     */
    projectId?: pulumi.Input<string>;
}
