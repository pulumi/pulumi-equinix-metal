// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get Equinix Metal Spot Market Price for a plan.
 *
 * ## Example Usage
 *
 * Lookup by facility:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const example = equinix_metal.getSpotMarketPrice({
 *     facility: "ny5",
 *     plan: "c3.small.x86",
 * });
 * ```
 *
 * Lookup by metro:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const example = equinix_metal.getSpotMarketPrice({
 *     metro: "sv",
 *     plan: "c3.small.x86",
 * });
 * ```
 */
export function getSpotMarketPrice(args: GetSpotMarketPriceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpotMarketPriceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("equinix-metal:index/getSpotMarketPrice:getSpotMarketPrice", {
        "facility": args.facility,
        "metro": args.metro,
        "plan": args.plan,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpotMarketPrice.
 */
export interface GetSpotMarketPriceArgs {
    /**
     * Name of the facility.
     */
    facility?: string;
    /**
     * Name of the metro.
     */
    metro?: string;
    /**
     * Name of the plan.
     */
    plan: string;
}

/**
 * A collection of values returned by getSpotMarketPrice.
 */
export interface GetSpotMarketPriceResult {
    readonly facility?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly metro?: string;
    readonly plan: string;
    /**
     * Current spot market price for given plan in given facility.
     */
    readonly price: number;
}
/**
 * Use this data source to get Equinix Metal Spot Market Price for a plan.
 *
 * ## Example Usage
 *
 * Lookup by facility:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const example = equinix_metal.getSpotMarketPrice({
 *     facility: "ny5",
 *     plan: "c3.small.x86",
 * });
 * ```
 *
 * Lookup by metro:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const example = equinix_metal.getSpotMarketPrice({
 *     metro: "sv",
 *     plan: "c3.small.x86",
 * });
 * ```
 */
export function getSpotMarketPriceOutput(args: GetSpotMarketPriceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpotMarketPriceResult> {
    return pulumi.output(args).apply((a: any) => getSpotMarketPrice(a, opts))
}

/**
 * A collection of arguments for invoking getSpotMarketPrice.
 */
export interface GetSpotMarketPriceOutputArgs {
    /**
     * Name of the facility.
     */
    facility?: pulumi.Input<string>;
    /**
     * Name of the metro.
     */
    metro?: pulumi.Input<string>;
    /**
     * Name of the plan.
     */
    plan: pulumi.Input<string>;
}
