// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Provides an Equinix Metal facility datasource.
 */
export function getFacility(args: GetFacilityArgs, opts?: pulumi.InvokeOptions): Promise<GetFacilityResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("equinix-metal:index/getFacility:getFacility", {
        "capacities": args.capacities,
        "code": args.code,
        "featuresRequireds": args.featuresRequireds,
    }, opts);
}

/**
 * A collection of arguments for invoking getFacility.
 */
export interface GetFacilityArgs {
    /**
     * (Optional) Ensure that queried facility has capacity for specified number of given plans
     */
    readonly capacities?: inputs.GetFacilityCapacity[];
    /**
     * The facility code
     */
    readonly code: string;
    /**
     * Set of feature strings that the facility must have
     */
    readonly featuresRequireds?: string[];
}

/**
 * A collection of values returned by getFacility.
 */
export interface GetFacilityResult {
    /**
     * (Optional) Ensure that queried facility has capacity for specified number of given plans
     */
    readonly capacities?: outputs.GetFacilityCapacity[];
    readonly code: string;
    /**
     * The features of the facility
     */
    readonly features: string[];
    readonly featuresRequireds?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The metro code the facility is part of
     */
    readonly metro: string;
    /**
     * The name of the facility
     */
    readonly name: string;
}
