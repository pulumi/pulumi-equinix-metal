// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve a virtual circuit resource from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const exampleConnection = equinix_metal.getConnection({
 *     connectionId: "4347e805-eb46-4699-9eb9-5c116e6a017d",
 * });
 * const exampleVc = exampleConnection.then(exampleConnection => equinix_metal.getVirtualCircuit({
 *     virtualCircuitId: exampleConnection.ports?[1]?.virtualCircuitIds?[0],
 * }));
 * ```
 */
export function getVirtualCircuit(args: GetVirtualCircuitArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualCircuitResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("equinix-metal:index/getVirtualCircuit:getVirtualCircuit", {
        "virtualCircuitId": args.virtualCircuitId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualCircuit.
 */
export interface GetVirtualCircuitArgs {
    /**
     * ID of the virtual circuit resource
     */
    virtualCircuitId: string;
}

/**
 * A collection of values returned by getVirtualCircuit.
 */
export interface GetVirtualCircuitResult {
    /**
     * Description for the Virtual Circuit resource
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the virtual circuit resource
     */
    readonly name: string;
    readonly nniVlan: number;
    readonly nniVnid: number;
    /**
     * ID of project to which the VC belongs
     * * `vnid`, `nniVlan`, `nniNvid` - VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)
     */
    readonly projectId: string;
    /**
     * Status of the virtal circuit
     */
    readonly status: string;
    /**
     * Tags for the Virtual Circuit resource
     */
    readonly tags: string[];
    readonly virtualCircuitId: string;
    readonly vnid: number;
}

export function getVirtualCircuitOutput(args: GetVirtualCircuitOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualCircuitResult> {
    return pulumi.output(args).apply(a => getVirtualCircuit(a, opts))
}

/**
 * A collection of arguments for invoking getVirtualCircuit.
 */
export interface GetVirtualCircuitOutputArgs {
    /**
     * ID of the virtual circuit resource
     */
    virtualCircuitId: pulumi.Input<string>;
}
