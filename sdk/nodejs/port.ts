// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Port extends pulumi.CustomResource {
    /**
     * Get an existing Port resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortState, opts?: pulumi.CustomResourceOptions): Port {
        return new Port(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix-metal:index/port:Port';

    /**
     * Returns true if the given object is an instance of Port.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Port {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Port.__pulumiType;
    }

    /**
     * UUID of the bond port
     */
    public /*out*/ readonly bondId!: pulumi.Output<string>;
    /**
     * Name of the bond port
     */
    public /*out*/ readonly bondName!: pulumi.Output<string>;
    /**
     * Whether the port should be bonded
     */
    public readonly bonded!: pulumi.Output<boolean>;
    /**
     * Flag indicating whether the port can be removed from a bond
     */
    public /*out*/ readonly disbondSupported!: pulumi.Output<boolean>;
    /**
     * Whether to put the port to Layer 2 mode, valid only for bond ports
     */
    public readonly layer2!: pulumi.Output<boolean | undefined>;
    /**
     * MAC address of the port
     */
    public /*out*/ readonly mac!: pulumi.Output<string>;
    /**
     * Name of the port, e.g. `bond0` or `eth0`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
     */
    public readonly nativeVlanId!: pulumi.Output<string | undefined>;
    /**
     * One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
     */
    public /*out*/ readonly networkType!: pulumi.Output<string>;
    /**
     * ID of the port to read
     */
    public readonly portId!: pulumi.Output<string>;
    /**
     * Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
     */
    public readonly resetOnDelete!: pulumi.Output<boolean | undefined>;
    /**
     * Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
     */
    public readonly vlanIds!: pulumi.Output<string[]>;
    /**
     * List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
     */
    public readonly vxlanIds!: pulumi.Output<number[]>;

    /**
     * Create a Port resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortArgs | PortState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortState | undefined;
            resourceInputs["bondId"] = state ? state.bondId : undefined;
            resourceInputs["bondName"] = state ? state.bondName : undefined;
            resourceInputs["bonded"] = state ? state.bonded : undefined;
            resourceInputs["disbondSupported"] = state ? state.disbondSupported : undefined;
            resourceInputs["layer2"] = state ? state.layer2 : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nativeVlanId"] = state ? state.nativeVlanId : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["portId"] = state ? state.portId : undefined;
            resourceInputs["resetOnDelete"] = state ? state.resetOnDelete : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vlanIds"] = state ? state.vlanIds : undefined;
            resourceInputs["vxlanIds"] = state ? state.vxlanIds : undefined;
        } else {
            const args = argsOrState as PortArgs | undefined;
            if ((!args || args.bonded === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bonded'");
            }
            if ((!args || args.portId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portId'");
            }
            resourceInputs["bonded"] = args ? args.bonded : undefined;
            resourceInputs["layer2"] = args ? args.layer2 : undefined;
            resourceInputs["nativeVlanId"] = args ? args.nativeVlanId : undefined;
            resourceInputs["portId"] = args ? args.portId : undefined;
            resourceInputs["resetOnDelete"] = args ? args.resetOnDelete : undefined;
            resourceInputs["vlanIds"] = args ? args.vlanIds : undefined;
            resourceInputs["vxlanIds"] = args ? args.vxlanIds : undefined;
            resourceInputs["bondId"] = undefined /*out*/;
            resourceInputs["bondName"] = undefined /*out*/;
            resourceInputs["disbondSupported"] = undefined /*out*/;
            resourceInputs["mac"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkType"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Port.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Port resources.
 */
export interface PortState {
    /**
     * UUID of the bond port
     */
    bondId?: pulumi.Input<string>;
    /**
     * Name of the bond port
     */
    bondName?: pulumi.Input<string>;
    /**
     * Whether the port should be bonded
     */
    bonded?: pulumi.Input<boolean>;
    /**
     * Flag indicating whether the port can be removed from a bond
     */
    disbondSupported?: pulumi.Input<boolean>;
    /**
     * Whether to put the port to Layer 2 mode, valid only for bond ports
     */
    layer2?: pulumi.Input<boolean>;
    /**
     * MAC address of the port
     */
    mac?: pulumi.Input<string>;
    /**
     * Name of the port, e.g. `bond0` or `eth0`
     */
    name?: pulumi.Input<string>;
    /**
     * UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
     */
    nativeVlanId?: pulumi.Input<string>;
    /**
     * One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
     */
    networkType?: pulumi.Input<string>;
    /**
     * ID of the port to read
     */
    portId?: pulumi.Input<string>;
    /**
     * Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
     */
    resetOnDelete?: pulumi.Input<boolean>;
    /**
     * Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
     */
    type?: pulumi.Input<string>;
    /**
     * List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
     */
    vlanIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
     */
    vxlanIds?: pulumi.Input<pulumi.Input<number>[]>;
}

/**
 * The set of arguments for constructing a Port resource.
 */
export interface PortArgs {
    /**
     * Whether the port should be bonded
     */
    bonded: pulumi.Input<boolean>;
    /**
     * Whether to put the port to Layer 2 mode, valid only for bond ports
     */
    layer2?: pulumi.Input<boolean>;
    /**
     * UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from `vlanIds` parameter), valid only for physical (non-bond) ports
     */
    nativeVlanId?: pulumi.Input<string>;
    /**
     * ID of the port to read
     */
    portId: pulumi.Input<string>;
    /**
     * Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
     */
    resetOnDelete?: pulumi.Input<boolean>;
    /**
     * List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports
     */
    vlanIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports
     */
    vxlanIds?: pulumi.Input<pulumi.Input<number>[]>;
}
