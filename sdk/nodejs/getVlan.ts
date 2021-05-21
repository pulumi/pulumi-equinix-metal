// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Provides an Equinix Metal Virtual Network datasource. VLANs data sources can be
 * searched by VLAN UUID, or project UUID and vxlan number.
 *
 * ## Example Usage
 *
 * Fetch a vlan by ID:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const foovlan = new equinix_metal.Vlan("foovlan", {
 *     projectId: local.project_id,
 *     metro: "sv",
 *     vxlan: 5,
 * });
 * const dsvlan = foovlan.id.apply(id => equinix_metal.getVlan({
 *     vlanId: id,
 * }));
 * ```
 *
 * Fetch a vlan by project ID, vxlan and metro
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const foovlan = new equinix_metal.Vlan("foovlan", {
 *     projectId: local.project_id,
 *     metro: "sv",
 *     vxlan: 5,
 * });
 * const dsvlan = equinix_metal.getVlan({
 *     projectId: local.project_id,
 *     vxlan: 5,
 *     metro: "sv",
 * });
 * ```
 */
export function getVlan(args?: GetVlanArgs, opts?: pulumi.InvokeOptions): Promise<GetVlanResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("equinix-metal:index/getVlan:getVlan", {
        "facility": args.facility,
        "metro": args.metro,
        "projectId": args.projectId,
        "vlanId": args.vlanId,
        "vxlan": args.vxlan,
    }, opts);
}

/**
 * A collection of arguments for invoking getVlan.
 */
export interface GetVlanArgs {
    /**
     * Facility where the VLAN is deployed
     */
    readonly facility?: string;
    /**
     * Metro where the VLAN is deployed
     */
    readonly metro?: string;
    /**
     * UUID of parent project of the VLAN. Use together with the vxlan number and metro or facility
     */
    readonly projectId?: string;
    /**
     * Metal UUID of the VLAN resource to look up
     */
    readonly vlanId?: string;
    /**
     * vxlan number of the VLAN to look up. Use together with the projectId and metro or facility
     */
    readonly vxlan?: number;
}

/**
 * A collection of values returned by getVlan.
 */
export interface GetVlanResult {
    /**
     * List of device ID to which this VLAN is assigned
     */
    readonly assignedDevicesIds: string[];
    /**
     * Description text of the VLAN resource
     */
    readonly description: string;
    readonly facility: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly metro: string;
    readonly projectId: string;
    readonly vlanId: string;
    readonly vxlan: number;
}
