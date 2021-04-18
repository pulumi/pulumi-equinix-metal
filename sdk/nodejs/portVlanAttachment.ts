// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to attach device ports to VLANs.
 *
 * Device and VLAN must be in the same facility.
 *
 * If you need this resource to add the port back to bond on removal, set `forceBond = true`.
 *
 * To learn more about Layer 2 networking in Equinix Metal, refer to
 *
 * * <https://metal.equinix.com/developers/docs/networking/layer2/>
 * * <https://metal.equinix.com/developers/docs/networking/layer2-configs/>
 *
 * ## Example Usage
 * ### Hybrid network type
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const testVlan = new equinix_metal.Vlan("testVlan", {
 *     description: "VLAN in New Jersey",
 *     facility: "ny5",
 *     projectId: local.project_id,
 * });
 * const testDevice = new equinix_metal.Device("testDevice", {
 *     hostname: "test",
 *     plan: "c3.small.x86",
 *     facilities: ["ny5"],
 *     operatingSystem: "ubuntu_20_04",
 *     billingCycle: "hourly",
 *     projectId: local.project_id,
 * });
 * const testDeviceNetworkType = new equinix_metal.DeviceNetworkType("testDeviceNetworkType", {
 *     deviceId: testDevice.id,
 *     type: "hybrid",
 * });
 * const testPortVlanAttachment = new equinix_metal.PortVlanAttachment("testPortVlanAttachment", {
 *     deviceId: testDeviceNetworkType.id,
 *     portName: "eth1",
 *     vlanVnid: testVlan.vxlan,
 * });
 * ```
 * ### Layer 2 network
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const testDevice = new equinix_metal.Device("testDevice", {
 *     hostname: "test",
 *     plan: "c3.small.x86",
 *     facilities: ["ny5"],
 *     operatingSystem: "ubuntu_20_04",
 *     billingCycle: "hourly",
 *     projectId: local.project_id,
 * });
 * const testDeviceNetworkType = new equinix_metal.DeviceNetworkType("testDeviceNetworkType", {
 *     deviceId: testDevice.id,
 *     type: "layer2-individual",
 * });
 * const test1Vlan = new equinix_metal.Vlan("test1Vlan", {
 *     description: "VLAN in New Jersey",
 *     facility: "ny5",
 *     projectId: local.project_id,
 * });
 * const test2Vlan = new equinix_metal.Vlan("test2Vlan", {
 *     description: "VLAN in New Jersey",
 *     facility: "ny5",
 *     projectId: local.project_id,
 * });
 * const test1PortVlanAttachment = new equinix_metal.PortVlanAttachment("test1PortVlanAttachment", {
 *     deviceId: testDeviceNetworkType.id,
 *     vlanVnid: test1Vlan.vxlan,
 *     portName: "eth1",
 * });
 * const test2PortVlanAttachment = new equinix_metal.PortVlanAttachment("test2PortVlanAttachment", {
 *     deviceId: testDeviceNetworkType.id,
 *     vlanVnid: test2Vlan.vxlan,
 *     portName: "eth1",
 *     native: true,
 * }, {
 *     dependsOn: ["metal_port_vlan_attachment.test1"],
 * });
 * ```
 * ## Attribute Referece
 *
 * * `id` - UUID of device port used in the assignment
 * * `vlanId` - UUID of VLAN API resource
 * * `portId` - UUID of device port
 */
export class PortVlanAttachment extends pulumi.CustomResource {
    /**
     * Get an existing PortVlanAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortVlanAttachmentState, opts?: pulumi.CustomResourceOptions): PortVlanAttachment {
        return new PortVlanAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix-metal:index/portVlanAttachment:PortVlanAttachment';

    /**
     * Returns true if the given object is an instance of PortVlanAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PortVlanAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PortVlanAttachment.__pulumiType;
    }

    /**
     * ID of device to be assigned to the VLAN
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * Add port back to the bond when this resource is removed. Default is false.
     */
    public readonly forceBond!: pulumi.Output<boolean | undefined>;
    /**
     * Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `dependsOn` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
     */
    public readonly native!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly portId!: pulumi.Output<string>;
    /**
     * Name of network port to be assigned to the VLAN
     */
    public readonly portName!: pulumi.Output<string>;
    public /*out*/ readonly vlanId!: pulumi.Output<string>;
    /**
     * VXLAN Network Identifier, integer
     */
    public readonly vlanVnid!: pulumi.Output<number>;

    /**
     * Create a PortVlanAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortVlanAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortVlanAttachmentArgs | PortVlanAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortVlanAttachmentState | undefined;
            inputs["deviceId"] = state ? state.deviceId : undefined;
            inputs["forceBond"] = state ? state.forceBond : undefined;
            inputs["native"] = state ? state.native : undefined;
            inputs["portId"] = state ? state.portId : undefined;
            inputs["portName"] = state ? state.portName : undefined;
            inputs["vlanId"] = state ? state.vlanId : undefined;
            inputs["vlanVnid"] = state ? state.vlanVnid : undefined;
        } else {
            const args = argsOrState as PortVlanAttachmentArgs | undefined;
            if ((!args || args.deviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceId'");
            }
            if ((!args || args.portName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portName'");
            }
            if ((!args || args.vlanVnid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vlanVnid'");
            }
            inputs["deviceId"] = args ? args.deviceId : undefined;
            inputs["forceBond"] = args ? args.forceBond : undefined;
            inputs["native"] = args ? args.native : undefined;
            inputs["portName"] = args ? args.portName : undefined;
            inputs["vlanVnid"] = args ? args.vlanVnid : undefined;
            inputs["portId"] = undefined /*out*/;
            inputs["vlanId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PortVlanAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PortVlanAttachment resources.
 */
export interface PortVlanAttachmentState {
    /**
     * ID of device to be assigned to the VLAN
     */
    readonly deviceId?: pulumi.Input<string>;
    /**
     * Add port back to the bond when this resource is removed. Default is false.
     */
    readonly forceBond?: pulumi.Input<boolean>;
    /**
     * Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `dependsOn` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
     */
    readonly native?: pulumi.Input<boolean>;
    readonly portId?: pulumi.Input<string>;
    /**
     * Name of network port to be assigned to the VLAN
     */
    readonly portName?: pulumi.Input<string>;
    readonly vlanId?: pulumi.Input<string>;
    /**
     * VXLAN Network Identifier, integer
     */
    readonly vlanVnid?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PortVlanAttachment resource.
 */
export interface PortVlanAttachmentArgs {
    /**
     * ID of device to be assigned to the VLAN
     */
    readonly deviceId: pulumi.Input<string>;
    /**
     * Add port back to the bond when this resource is removed. Default is false.
     */
    readonly forceBond?: pulumi.Input<boolean>;
    /**
     * Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `dependsOn` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
     */
    readonly native?: pulumi.Input<boolean>;
    /**
     * Name of network port to be assigned to the VLAN
     */
    readonly portName: pulumi.Input<string>;
    /**
     * VXLAN Network Identifier, integer
     */
    readonly vlanVnid: pulumi.Input<number>;
}
