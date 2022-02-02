// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Datasource `equinix-metal.Volume` was removed in version 3.0.0, and the API support was deprecated on June 1st 2021. See https://metal.equinix.com/developers/docs/storage/elastic-block-storage/#elastic-block-storage for more details.
 */
export function getVolume(args?: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("equinix-metal:index/getVolume:getVolume", {
        "name": args.name,
        "projectId": args.projectId,
        "volumeId": args.volumeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeArgs {
    name?: string;
    projectId?: string;
    volumeId?: string;
}

/**
 * A collection of values returned by getVolume.
 */
export interface GetVolumeResult {
    readonly billingCycle: string;
    readonly created: string;
    readonly description: string;
    readonly deviceIds: string[];
    readonly facility: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly locked: boolean;
    readonly name: string;
    readonly plan: string;
    readonly projectId: string;
    readonly size: number;
    readonly snapshotPolicies: outputs.GetVolumeSnapshotPolicy[];
    readonly state: string;
    readonly updated: string;
    readonly volumeId: string;
}

export function getVolumeOutput(args?: GetVolumeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeResult> {
    return pulumi.output(args).apply(a => getVolume(a, opts))
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeOutputArgs {
    name?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    volumeId?: pulumi.Input<string>;
}
