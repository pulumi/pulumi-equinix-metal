// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides an Equinix Metal facility datasource.
 */
export function getFacility(args: GetFacilityArgs, opts?: pulumi.InvokeOptions): Promise<GetFacilityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
    capacities?: inputs.GetFacilityCapacity[];
    /**
     * The facility code
     */
    code: string;
    /**
     * Set of feature strings that the facility must have
     *
     * Facilities can be looked up by `code`.
     */
    featuresRequireds?: string[];
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
/**
 * Provides an Equinix Metal facility datasource.
 */
export function getFacilityOutput(args: GetFacilityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFacilityResult> {
    return pulumi.output(args).apply((a: any) => getFacility(a, opts))
}

/**
 * A collection of arguments for invoking getFacility.
 */
export interface GetFacilityOutputArgs {
    /**
     * (Optional) Ensure that queried facility has capacity for specified number of given plans
     */
    capacities?: pulumi.Input<pulumi.Input<inputs.GetFacilityCapacityArgs>[]>;
    /**
     * The facility code
     */
    code: pulumi.Input<string>;
    /**
     * Set of feature strings that the facility must have
     *
     * Facilities can be looked up by `code`.
     */
    featuresRequireds?: pulumi.Input<pulumi.Input<string>[]>;
}
