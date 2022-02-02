// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Resource `equinix-metal.Volume` was removed in version 3.0.0, and the API support was deprecated on June 1st 2021. See https://metal.equinix.com/developers/docs/storage/elastic-block-storage/#elastic-block-storage for more details.
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix-metal:index/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    public /*out*/ readonly attachments!: pulumi.Output<outputs.VolumeAttachment[]>;
    public readonly billingCycle!: pulumi.Output<string>;
    public /*out*/ readonly created!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly facility!: pulumi.Output<string>;
    public readonly locked!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly plan!: pulumi.Output<string>;
    public readonly projectId!: pulumi.Output<string>;
    public readonly size!: pulumi.Output<number>;
    public readonly snapshotPolicies!: pulumi.Output<outputs.VolumeSnapshotPolicy[] | undefined>;
    public /*out*/ readonly state!: pulumi.Output<string>;
    public /*out*/ readonly updated!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            resourceInputs["attachments"] = state ? state.attachments : undefined;
            resourceInputs["billingCycle"] = state ? state.billingCycle : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["facility"] = state ? state.facility : undefined;
            resourceInputs["locked"] = state ? state.locked : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotPolicies"] = state ? state.snapshotPolicies : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["updated"] = state ? state.updated : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.facility === undefined) && !opts.urn) {
                throw new Error("Missing required property 'facility'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["billingCycle"] = args ? args.billingCycle : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["facility"] = args ? args.facility : undefined;
            resourceInputs["locked"] = args ? args.locked : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotPolicies"] = args ? args.snapshotPolicies : undefined;
            resourceInputs["attachments"] = undefined /*out*/;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    attachments?: pulumi.Input<pulumi.Input<inputs.VolumeAttachment>[]>;
    billingCycle?: pulumi.Input<string | enums.BillingCycle>;
    created?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    facility?: pulumi.Input<string | enums.Facility>;
    locked?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    plan?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    size?: pulumi.Input<number>;
    snapshotPolicies?: pulumi.Input<pulumi.Input<inputs.VolumeSnapshotPolicy>[]>;
    state?: pulumi.Input<string>;
    updated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    billingCycle?: pulumi.Input<string | enums.BillingCycle>;
    description?: pulumi.Input<string>;
    facility: pulumi.Input<string | enums.Facility>;
    locked?: pulumi.Input<boolean>;
    plan: pulumi.Input<string>;
    projectId: pulumi.Input<string>;
    size: pulumi.Input<number>;
    snapshotPolicies?: pulumi.Input<pulumi.Input<inputs.VolumeSnapshotPolicy>[]>;
}
