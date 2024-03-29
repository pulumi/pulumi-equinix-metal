// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a resource to create and manage blocks of reserved IP addresses in a project.
 *
 * When a user provisions first device in a facility, Equinix Metal API automatically allocates IPv6/56 and private IPv4/25 blocks.
 * The new device then gets IPv6 and private IPv4 addresses from those block. It also gets a public IPv4/31 address.
 * Every new device in the project and facility will automatically get IPv6 and private IPv4 addresses from these pre-allocated blocks.
 * The IPv6 and private IPv4 blocks can't be created, only imported. With this resource, it's possible to create either public IPv4 blocks or global IPv4 blocks.
 *
 * Public blocks are allocated in a facility. Addresses from public blocks can only be assigned to devices in the facility. Public blocks can have mask from /24 (256 addresses) to /32 (1 address). If you create public block with this resource, you must fill the facility argmument.
 *
 * Addresses from global blocks can be assigned in any facility. Global blocks can have mask from /30 (4 addresses), to /32 (1 address). If you create global block with this resource, you must specify type = "globalIpv4" and you must omit the facility argument.
 *
 * Once IP block is allocated or imported, an address from it can be assigned to device with the `equinix-metal.IpAttachment` resource.
 *
 * ## Example Usage
 *
 * Allocate reserved IP blocks:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * // Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility for myproject
 * const twoElasticAddresses = new equinix_metal.ReservedIpBlock("twoElasticAddresses", {
 *     projectId: local.project_id,
 *     facility: "sv15",
 *     quantity: 2,
 * });
 * // Allocate 1 floating IP in Sillicon Valley (sv) metro
 * const testReservedIpBlock = new equinix_metal.ReservedIpBlock("testReservedIpBlock", {
 *     projectId: local.project_id,
 *     type: "public_ipv4",
 *     metro: "sv",
 *     quantity: 1,
 * });
 * // Allocate 1 global floating IP, which can be assigned to device in any facility
 * const testIndex_reservedIpBlockReservedIpBlock = new equinix_metal.ReservedIpBlock("testIndex/reservedIpBlockReservedIpBlock", {
 *     projectId: local.project_id,
 *     type: "global_ipv4",
 *     quantity: 1,
 * });
 * ```
 *
 * Allocate a block and run a device with public IPv4 from the block
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * // Allocate /31 block of max 2 public IPv4 addresses in Silicon Valley (sv15) facility
 * const example = new equinix_metal.ReservedIpBlock("example", {
 *     projectId: local.project_id,
 *     facility: "sv15",
 *     quantity: 2,
 * });
 * // Run a device with both public IPv4 from the block assigned
 * const nodes = new equinix_metal.Device("nodes", {
 *     projectId: local.project_id,
 *     facilities: ["sv15"],
 *     plan: "c3.small.x86",
 *     operatingSystem: "ubuntu_20_04",
 *     hostname: "test",
 *     billingCycle: "hourly",
 *     ipAddresses: [
 *         {
 *             type: "public_ipv4",
 *             cidr: 31,
 *             reservationIds: [example.id],
 *         },
 *         {
 *             type: "private_ipv4",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using an existing IP reservation ID
 *
 * ```sh
 *  $ pulumi import equinix-metal:index/reservedIpBlock:ReservedIpBlock metal_reserved_ip_block {existing_ip_reservation_id}
 * ```
 */
export class ReservedIpBlock extends pulumi.CustomResource {
    /**
     * Get an existing ReservedIpBlock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReservedIpBlockState, opts?: pulumi.CustomResourceOptions): ReservedIpBlock {
        return new ReservedIpBlock(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix-metal:index/reservedIpBlock:ReservedIpBlock';

    /**
     * Returns true if the given object is an instance of ReservedIpBlock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReservedIpBlock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReservedIpBlock.__pulumiType;
    }

    public /*out*/ readonly address!: pulumi.Output<string>;
    /**
     * Address family as integer (4 or 6)
     */
    public /*out*/ readonly addressFamily!: pulumi.Output<number>;
    /**
     * length of CIDR prefix of the block as integer
     */
    public /*out*/ readonly cidr!: pulumi.Output<number>;
    /**
     * Address and mask in CIDR notation, e.g. "147.229.15.30/31"
     */
    public /*out*/ readonly cidrNotation!: pulumi.Output<string>;
    /**
     * Arbitrary description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
     */
    public readonly facility!: pulumi.Output<string | undefined>;
    public /*out*/ readonly gateway!: pulumi.Output<string>;
    /**
     * boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
     */
    public /*out*/ readonly global!: pulumi.Output<boolean>;
    public /*out*/ readonly manageable!: pulumi.Output<boolean>;
    public /*out*/ readonly management!: pulumi.Output<boolean>;
    /**
     * Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
     */
    public readonly metro!: pulumi.Output<string | undefined>;
    /**
     * Mask in decimal notation, e.g. "255.255.255.0"
     */
    public /*out*/ readonly netmask!: pulumi.Output<string>;
    /**
     * Network IP address portion of the block specification
     */
    public /*out*/ readonly network!: pulumi.Output<string>;
    /**
     * The metal project ID where to allocate the address block
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * boolean flag whether addresses from a block are public
     */
    public /*out*/ readonly public!: pulumi.Output<boolean>;
    /**
     * The number of allocated /32 addresses, a power of 2
     */
    public readonly quantity!: pulumi.Output<number>;
    /**
     * String list of tags
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Either "globalIpv4" or "publicIpv4", defaults to "publicIpv4" for backward compatibility
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ReservedIpBlock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReservedIpBlockArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReservedIpBlockArgs | ReservedIpBlockState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReservedIpBlockState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["cidr"] = state ? state.cidr : undefined;
            resourceInputs["cidrNotation"] = state ? state.cidrNotation : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["facility"] = state ? state.facility : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["global"] = state ? state.global : undefined;
            resourceInputs["manageable"] = state ? state.manageable : undefined;
            resourceInputs["management"] = state ? state.management : undefined;
            resourceInputs["metro"] = state ? state.metro : undefined;
            resourceInputs["netmask"] = state ? state.netmask : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["public"] = state ? state.public : undefined;
            resourceInputs["quantity"] = state ? state.quantity : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ReservedIpBlockArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.quantity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'quantity'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["facility"] = args ? args.facility : undefined;
            resourceInputs["metro"] = args ? args.metro : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["quantity"] = args ? args.quantity : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["address"] = undefined /*out*/;
            resourceInputs["addressFamily"] = undefined /*out*/;
            resourceInputs["cidr"] = undefined /*out*/;
            resourceInputs["cidrNotation"] = undefined /*out*/;
            resourceInputs["gateway"] = undefined /*out*/;
            resourceInputs["global"] = undefined /*out*/;
            resourceInputs["manageable"] = undefined /*out*/;
            resourceInputs["management"] = undefined /*out*/;
            resourceInputs["netmask"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["public"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReservedIpBlock.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReservedIpBlock resources.
 */
export interface ReservedIpBlockState {
    address?: pulumi.Input<string>;
    /**
     * Address family as integer (4 or 6)
     */
    addressFamily?: pulumi.Input<number>;
    /**
     * length of CIDR prefix of the block as integer
     */
    cidr?: pulumi.Input<number>;
    /**
     * Address and mask in CIDR notation, e.g. "147.229.15.30/31"
     */
    cidrNotation?: pulumi.Input<string>;
    /**
     * Arbitrary description
     */
    description?: pulumi.Input<string>;
    /**
     * Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
     */
    facility?: pulumi.Input<string | enums.Facility>;
    gateway?: pulumi.Input<string>;
    /**
     * boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)
     */
    global?: pulumi.Input<boolean>;
    manageable?: pulumi.Input<boolean>;
    management?: pulumi.Input<boolean>;
    /**
     * Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
     */
    metro?: pulumi.Input<string>;
    /**
     * Mask in decimal notation, e.g. "255.255.255.0"
     */
    netmask?: pulumi.Input<string>;
    /**
     * Network IP address portion of the block specification
     */
    network?: pulumi.Input<string>;
    /**
     * The metal project ID where to allocate the address block
     */
    projectId?: pulumi.Input<string>;
    /**
     * boolean flag whether addresses from a block are public
     */
    public?: pulumi.Input<boolean>;
    /**
     * The number of allocated /32 addresses, a power of 2
     */
    quantity?: pulumi.Input<number>;
    /**
     * String list of tags
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either "globalIpv4" or "publicIpv4", defaults to "publicIpv4" for backward compatibility
     */
    type?: pulumi.Input<string | enums.IpBlockType>;
}

/**
 * The set of arguments for constructing a ReservedIpBlock resource.
 */
export interface ReservedIpBlockArgs {
    /**
     * Arbitrary description
     */
    description?: pulumi.Input<string>;
    /**
     * Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `metro`
     */
    facility?: pulumi.Input<string | enums.Facility>;
    /**
     * Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with `facility`
     */
    metro?: pulumi.Input<string>;
    /**
     * The metal project ID where to allocate the address block
     */
    projectId: pulumi.Input<string>;
    /**
     * The number of allocated /32 addresses, a power of 2
     */
    quantity: pulumi.Input<number>;
    /**
     * String list of tags
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either "globalIpv4" or "publicIpv4", defaults to "publicIpv4" for backward compatibility
     */
    type?: pulumi.Input<string | enums.IpBlockType>;
}
