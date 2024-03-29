// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve a [hardware reservation resource from Equinix Metal](https://metal.equinix.com/developers/docs/deploy/reserved/).
 *
 * You can look up hardware reservation by its ID or by ID of device which occupies it.
 */
export function getHardwareReservation(args?: GetHardwareReservationArgs, opts?: pulumi.InvokeOptions): Promise<GetHardwareReservationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("equinix-metal:index/getHardwareReservation:getHardwareReservation", {
        "deviceId": args.deviceId,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getHardwareReservation.
 */
export interface GetHardwareReservationArgs {
    /**
     * UUID of device occupying the reservation
     */
    deviceId?: string;
    /**
     * ID of the hardware reservation
     */
    id?: string;
}

/**
 * A collection of values returned by getHardwareReservation.
 */
export interface GetHardwareReservationResult {
    /**
     * UUID of device occupying the reservation
     */
    readonly deviceId: string;
    /**
     * Plan type for the reservation
     */
    readonly facility: string;
    /**
     * ID of the hardware reservation to look up
     */
    readonly id: string;
    /**
     * Plan type for the reservation
     */
    readonly plan: string;
    /**
     * UUID of project this reservation is scoped to
     */
    readonly projectId: string;
    /**
     * Flag indicating whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first
     */
    readonly provisionable: boolean;
    /**
     * Reservation short ID
     */
    readonly shortId: string;
    /**
     * Flag indicating whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix
     */
    readonly spare: boolean;
    /**
     * Switch short ID, can be used to determine if two devices are connected to the same switch
     */
    readonly switchUuid: string;
}
/**
 * Use this data source to retrieve a [hardware reservation resource from Equinix Metal](https://metal.equinix.com/developers/docs/deploy/reserved/).
 *
 * You can look up hardware reservation by its ID or by ID of device which occupies it.
 */
export function getHardwareReservationOutput(args?: GetHardwareReservationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHardwareReservationResult> {
    return pulumi.output(args).apply((a: any) => getHardwareReservation(a, opts))
}

/**
 * A collection of arguments for invoking getHardwareReservation.
 */
export interface GetHardwareReservationOutputArgs {
    /**
     * UUID of device occupying the reservation
     */
    deviceId?: pulumi.Input<string>;
    /**
     * ID of the hardware reservation
     */
    id?: pulumi.Input<string>;
}
