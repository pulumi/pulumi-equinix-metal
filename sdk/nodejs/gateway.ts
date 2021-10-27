// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this resource to create Metal Gateway resources in Equinix Metal.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * // Create Metal Gateway for a VLAN with a private IPv4 block with 8 IP addresses
 * const testVlan = new equinix_metal.Vlan("testVlan", {
 *     description: "test VLAN in SV",
 *     metro: "sv",
 *     projectId: local.project_id,
 * });
 * const testGateway = new equinix_metal.Gateway("testGateway", {
 *     projectId: local.project_id,
 *     vlanId: testVlan.id,
 *     privateIpv4SubnetSize: 8,
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * // Create Metal Gateway for a VLAN and reserved IP address block
 * const testVlan = new equinix_metal.Vlan("testVlan", {
 *     description: "test VLAN in SV",
 *     metro: "sv",
 *     projectId: local.project_id,
 * });
 * const testReservedIpBlock = new equinix_metal.ReservedIpBlock("testReservedIpBlock", {
 *     projectId: local.project_id,
 *     metro: "sv",
 *     quantity: 2,
 * });
 * const testGateway = new equinix_metal.Gateway("testGateway", {
 *     projectId: local.project_id,
 *     vlanId: testVlan.id,
 *     ipReservationId: testReservedIpBlock.id,
 * });
 * ```
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayState, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix-metal:index/gateway:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
     */
    public readonly ipReservationId!: pulumi.Output<string | undefined>;
    /**
     * Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
     */
    public readonly privateIpv4SubnetSize!: pulumi.Output<number | undefined>;
    /**
     * UUID of the project where the gateway is scoped to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Status of the gateway resource
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * UUID of the VLAN where the gateway is scoped to
     */
    public readonly vlanId!: pulumi.Output<string>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayArgs | GatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayState | undefined;
            inputs["ipReservationId"] = state ? state.ipReservationId : undefined;
            inputs["privateIpv4SubnetSize"] = state ? state.privateIpv4SubnetSize : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["vlanId"] = state ? state.vlanId : undefined;
        } else {
            const args = argsOrState as GatewayArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.vlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vlanId'");
            }
            inputs["ipReservationId"] = args ? args.ipReservationId : undefined;
            inputs["privateIpv4SubnetSize"] = args ? args.privateIpv4SubnetSize : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["vlanId"] = args ? args.vlanId : undefined;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Gateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Gateway resources.
 */
export interface GatewayState {
    /**
     * UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
     */
    readonly ipReservationId?: pulumi.Input<string>;
    /**
     * Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
     */
    readonly privateIpv4SubnetSize?: pulumi.Input<number>;
    /**
     * UUID of the project where the gateway is scoped to
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Status of the gateway resource
     */
    readonly state?: pulumi.Input<string>;
    /**
     * UUID of the VLAN where the gateway is scoped to
     */
    readonly vlanId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with `privateIpv4SubnetSize`
     */
    readonly ipReservationId?: pulumi.Input<string>;
    /**
     * Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with `ipReservationId`
     */
    readonly privateIpv4SubnetSize?: pulumi.Input<number>;
    /**
     * UUID of the project where the gateway is scoped to
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * UUID of the VLAN where the gateway is scoped to
     */
    readonly vlanId: pulumi.Input<string>;
}