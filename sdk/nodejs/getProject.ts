// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Use this datasource to retrieve attributes of the Project API resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const tfProject1 = equinix_metal.getProject({
 *     name: "Terraform Fun",
 * });
 * export const usersOfTerraformFun = tfProject1.then(tfProject1 => tfProject1.userIds);
 * ```
 */
export function getProject(args?: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("equinix-metal:index/getProject:getProject", {
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * The name which is used to look up the project
     */
    name?: string;
    /**
     * The UUID by which to look up the project
     */
    projectId?: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * Whether Backend Transfer is enabled for this project
     */
    readonly backendTransfer: boolean;
    /**
     * Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
     */
    readonly bgpConfigs: outputs.GetProjectBgpConfig[];
    /**
     * The timestamp for when the project was created
     */
    readonly created: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The UUID of this project's parent organization
     */
    readonly organizationId: string;
    /**
     * The UUID of payment method for this project
     */
    readonly paymentMethodId: string;
    readonly projectId: string;
    /**
     * The timestamp for the last time the project was updated
     */
    readonly updated: string;
    /**
     * List of UUIDs of user accounts which belong to this project
     */
    readonly userIds: string[];
}

export function getProjectOutput(args?: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply(a => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    /**
     * The name which is used to look up the project
     */
    name?: pulumi.Input<string>;
    /**
     * The UUID by which to look up the project
     */
    projectId?: pulumi.Input<string>;
}
