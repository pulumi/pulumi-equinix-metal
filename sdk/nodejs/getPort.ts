// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to read ports of existing devices. You can read port by either its UUID, or by a device UUID and port name.
 *
 * ## Example Usage
 *
 * Create a device and read it's eth0 port to the datasource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix_metal from "@pulumi/equinix-metal";
 *
 * const projectId = "<UUID_of_your_project>";
 * const testDevice = new equinix_metal.Device("testDevice", {
 *     hostname: "tfacc-test-device-port",
 *     plan: "c3.medium.x86",
 *     facilities: ["sv15"],
 *     operatingSystem: "ubuntu_20_04",
 *     billingCycle: "hourly",
 *     projectId: projectId,
 * });
 * const testPort = equinix_metal.getPortOutput({
 *     deviceId: testDevice.id,
 *     name: "eth0",
 * });
 * ```
 */
export function getPort(args?: GetPortArgs, opts?: pulumi.InvokeOptions): Promise<GetPortResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("equinix-metal:index/getPort:getPort", {
        "deviceId": args.deviceId,
        "name": args.name,
        "portId": args.portId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPort.
 */
export interface GetPortArgs {
    deviceId?: string;
    /**
     * Whether to look for public or private block.
     */
    name?: string;
    portId?: string;
}

/**
 * A collection of values returned by getPort.
 */
export interface GetPortResult {
    /**
     * UUID of the bond port"
     */
    readonly bondId: string;
    /**
     * Name of the bond port
     */
    readonly bondName: string;
    /**
     * Flag indicating whether the port is bonded
     */
    readonly bonded: boolean;
    readonly deviceId?: string;
    /**
     * Flag indicating whether the port can be removed from a bond
     */
    readonly disbondSupported: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly layer2: boolean;
    /**
     * MAC address of the port
     */
    readonly mac: string;
    readonly name: string;
    /**
     * UUID of native VLAN of the port
     */
    readonly nativeVlanId: string;
    /**
     * One of layer2-bonded, layer2-individual, layer3, hybrid, hybrid-bonded
     */
    readonly networkType: string;
    readonly portId?: string;
    /**
     * Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
     */
    readonly type: string;
    /**
     * UUIDs of attached VLANs
     */
    readonly vlanIds: string[];
    /**
     * VXLAN ids of attached VLANs
     */
    readonly vxlanIds: number[];
}

export function getPortOutput(args?: GetPortOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortResult> {
    return pulumi.output(args).apply(a => getPort(a, opts))
}

/**
 * A collection of arguments for invoking getPort.
 */
export interface GetPortOutputArgs {
    deviceId?: pulumi.Input<string>;
    /**
     * Whether to look for public or private block.
     */
    name?: pulumi.Input<string>;
    portId?: pulumi.Input<string>;
}
