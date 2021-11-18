// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Provides an Equinix Metal metro datasource.
 */
export function getMetro(args: GetMetroArgs, opts?: pulumi.InvokeOptions): Promise<GetMetroResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("equinix-metal:index/getMetro:getMetro", {
        "capacities": args.capacities,
        "code": args.code,
    }, opts);
}

/**
 * A collection of arguments for invoking getMetro.
 */
export interface GetMetroArgs {
    /**
     * (Optional) Ensure that queried metro has capacity for specified number of given plans
     */
    capacities?: inputs.GetMetroCapacity[];
    /**
     * The metro code
     */
    code: string;
}

/**
 * A collection of values returned by getMetro.
 */
export interface GetMetroResult {
    /**
     * (Optional) Ensure that queried metro has capacity for specified number of given plans
     */
    readonly capacities?: outputs.GetMetroCapacity[];
    /**
     * The code of the metro
     */
    readonly code: string;
    /**
     * The country of the metro
     */
    readonly country: string;
    /**
     * The ID of the metro
     */
    readonly id: string;
    /**
     * The name of the metro
     */
    readonly name: string;
}

export function getMetroOutput(args: GetMetroOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMetroResult> {
    return pulumi.output(args).apply(a => getMetro(a, opts))
}

/**
 * A collection of arguments for invoking getMetro.
 */
export interface GetMetroOutputArgs {
    /**
     * (Optional) Ensure that queried metro has capacity for specified number of given plans
     */
    capacities?: pulumi.Input<pulumi.Input<inputs.GetMetroCapacityArgs>[]>;
    /**
     * The metro code
     */
    code: pulumi.Input<string>;
}
