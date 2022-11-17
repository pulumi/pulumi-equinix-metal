// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.equinixmetal.DeviceArgs;
import com.pulumi.equinixmetal.Utilities;
import com.pulumi.equinixmetal.inputs.DeviceState;
import com.pulumi.equinixmetal.outputs.DeviceIpAddress;
import com.pulumi.equinixmetal.outputs.DeviceNetwork;
import com.pulumi.equinixmetal.outputs.DevicePort;
import com.pulumi.equinixmetal.outputs.DeviceReinstall;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Equinix Metal device resource. This can be used to create,
 * modify, and delete devices.
 * 
 * &gt; **Note:** All arguments including the `root_password` and `user_data` will be stored in
 *  the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * This resource can be imported using an existing device ID
 * 
 * ```sh
 *  $ pulumi import equinix-metal:index/device:Device metal_device {existing_device_id}
 * ```
 * 
 */
@ResourceType(type="equinix-metal:index/device:Device")
public class Device extends com.pulumi.resources.CustomResource {
    /**
     * The ipv4 private IP assigned to the device
     * 
     */
    @Export(name="accessPrivateIpv4", type=String.class, parameters={})
    private Output<String> accessPrivateIpv4;

    /**
     * @return The ipv4 private IP assigned to the device
     * 
     */
    public Output<String> accessPrivateIpv4() {
        return this.accessPrivateIpv4;
    }
    /**
     * The ipv4 maintenance IP assigned to the device
     * 
     */
    @Export(name="accessPublicIpv4", type=String.class, parameters={})
    private Output<String> accessPublicIpv4;

    /**
     * @return The ipv4 maintenance IP assigned to the device
     * 
     */
    public Output<String> accessPublicIpv4() {
        return this.accessPublicIpv4;
    }
    /**
     * The ipv6 maintenance IP assigned to the device
     * 
     */
    @Export(name="accessPublicIpv6", type=String.class, parameters={})
    private Output<String> accessPublicIpv6;

    /**
     * @return The ipv6 maintenance IP assigned to the device
     * 
     */
    public Output<String> accessPublicIpv6() {
        return this.accessPublicIpv6;
    }
    /**
     * If true, a device with OS `custom_ipxe` will continue to boot via iPXE on reboots.
     * 
     */
    @Export(name="alwaysPxe", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> alwaysPxe;

    /**
     * @return If true, a device with OS `custom_ipxe` will continue to boot via iPXE on reboots.
     * 
     */
    public Output<Optional<Boolean>> alwaysPxe() {
        return Codegen.optional(this.alwaysPxe);
    }
    /**
     * monthly or hourly
     * 
     */
    @Export(name="billingCycle", type=String.class, parameters={})
    private Output<String> billingCycle;

    /**
     * @return monthly or hourly
     * 
     */
    public Output<String> billingCycle() {
        return this.billingCycle;
    }
    /**
     * The timestamp for when the device was created
     * 
     */
    @Export(name="created", type=String.class, parameters={})
    private Output<String> created;

    /**
     * @return The timestamp for when the device was created
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * A string of the desired Custom Data for the device.
     * 
     */
    @Export(name="customData", type=String.class, parameters={})
    private Output</* @Nullable */ String> customData;

    /**
     * @return A string of the desired Custom Data for the device.
     * 
     */
    public Output<Optional<String>> customData() {
        return Codegen.optional(this.customData);
    }
    /**
     * The facility where the device is deployed.
     * 
     */
    @Export(name="deployedFacility", type=String.class, parameters={})
    private Output<String> deployedFacility;

    /**
     * @return The facility where the device is deployed.
     * 
     */
    public Output<String> deployedFacility() {
        return this.deployedFacility;
    }
    /**
     * ID of hardware reservation where this device was deployed. It is useful when using the `next-available` hardware reservation.
     * 
     */
    @Export(name="deployedHardwareReservationId", type=String.class, parameters={})
    private Output<String> deployedHardwareReservationId;

    /**
     * @return ID of hardware reservation where this device was deployed. It is useful when using the `next-available` hardware reservation.
     * 
     */
    public Output<String> deployedHardwareReservationId() {
        return this.deployedHardwareReservationId;
    }
    /**
     * The device description.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The device description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or `any` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with `metro`.
     * 
     */
    @Export(name="facilities", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> facilities;

    /**
     * @return List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or `any` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with `metro`.
     * 
     */
    public Output<Optional<List<String>>> facilities() {
        return Codegen.optional(this.facilities);
    }
    /**
     * Delete device even if it has volumes attached. Only applies for destroy action.
     * 
     */
    @Export(name="forceDetachVolumes", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceDetachVolumes;

    /**
     * @return Delete device even if it has volumes attached. Only applies for destroy action.
     * 
     */
    public Output<Optional<Boolean>> forceDetachVolumes() {
        return Codegen.optional(this.forceDetachVolumes);
    }
    /**
     * The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your
     * next available reservation automatically
     * 
     */
    @Export(name="hardwareReservationId", type=String.class, parameters={})
    private Output</* @Nullable */ String> hardwareReservationId;

    /**
     * @return The UUID of the hardware reservation where you want this device deployed, or next-available if you want to pick your
     * next available reservation automatically
     * 
     */
    public Output<Optional<String>> hardwareReservationId() {
        return Codegen.optional(this.hardwareReservationId);
    }
    /**
     * The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
     * 
     */
    @Export(name="hostname", type=String.class, parameters={})
    private Output<String> hostname;

    /**
     * @return The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * A list of IP address types for the device (structure is documented below).
     * 
     */
    @Export(name="ipAddresses", type=List.class, parameters={DeviceIpAddress.class})
    private Output</* @Nullable */ List<DeviceIpAddress>> ipAddresses;

    /**
     * @return A list of IP address types for the device (structure is documented below).
     * 
     */
    public Output<Optional<List<DeviceIpAddress>>> ipAddresses() {
        return Codegen.optional(this.ipAddresses);
    }
    /**
     * URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.
     * 
     */
    @Export(name="ipxeScriptUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> ipxeScriptUrl;

    /**
     * @return URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.
     * 
     */
    public Output<Optional<String>> ipxeScriptUrl() {
        return Codegen.optional(this.ipxeScriptUrl);
    }
    /**
     * Whether the device is locked
     * 
     */
    @Export(name="locked", type=Boolean.class, parameters={})
    private Output<Boolean> locked;

    /**
     * @return Whether the device is locked
     * 
     */
    public Output<Boolean> locked() {
        return this.locked;
    }
    /**
     * Metro area for the new device. Conflicts with `facilities`.
     * 
     */
    @Export(name="metro", type=String.class, parameters={})
    private Output</* @Nullable */ String> metro;

    /**
     * @return Metro area for the new device. Conflicts with `facilities`.
     * 
     */
    public Output<Optional<String>> metro() {
        return Codegen.optional(this.metro);
    }
    /**
     * Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/).
     * Will be one of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
     * 
     * @deprecated
     * You should handle Network Type with the new metal_device_network_type resource.
     * 
     */
    @Deprecated /* You should handle Network Type with the new metal_device_network_type resource. */
    @Export(name="networkType", type=String.class, parameters={})
    private Output<String> networkType;

    /**
     * @return Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/).
     * Will be one of layer3, hybrid, hybrid-bonded, layer2-individual, layer2-bonded
     * 
     */
    public Output<String> networkType() {
        return this.networkType;
    }
    /**
     * The device&#39;s private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:
     * * Public IPv4 at `metal_device.name.network.0`
     * * IPv6 at `metal_device.name.network.1`
     * * Private IPv4 at `metal_device.name.network.2`
     *   Elastic addresses then stack by type - an assigned public IPv4 will go after the management public IPv4 (to index 1), and will then shift the indices of the IPv6 and private IPv4. Assigned private IPv4 will go after the management private IPv4 (to the end of the network list).
     *   The fields of the network attributes are:
     * 
     */
    @Export(name="networks", type=List.class, parameters={DeviceNetwork.class})
    private Output<List<DeviceNetwork>> networks;

    /**
     * @return The device&#39;s private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:
     * * Public IPv4 at `metal_device.name.network.0`
     * * IPv6 at `metal_device.name.network.1`
     * * Private IPv4 at `metal_device.name.network.2`
     *   Elastic addresses then stack by type - an assigned public IPv4 will go after the management public IPv4 (to index 1), and will then shift the indices of the IPv6 and private IPv4. Assigned private IPv4 will go after the management private IPv4 (to the end of the network list).
     *   The fields of the network attributes are:
     * 
     */
    public Output<List<DeviceNetwork>> networks() {
        return this.networks;
    }
    /**
     * The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.
     * 
     */
    @Export(name="operatingSystem", type=String.class, parameters={})
    private Output<String> operatingSystem;

    /**
     * @return The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.
     * 
     */
    public Output<String> operatingSystem() {
        return this.operatingSystem;
    }
    /**
     * The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.
     * 
     */
    @Export(name="plan", type=String.class, parameters={})
    private Output<String> plan;

    /**
     * @return The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.
     * 
     */
    public Output<String> plan() {
        return this.plan;
    }
    /**
     * Ports assigned to the device
     * 
     */
    @Export(name="ports", type=List.class, parameters={DevicePort.class})
    private Output<List<DevicePort>> ports;

    /**
     * @return Ports assigned to the device
     * 
     */
    public Output<List<DevicePort>> ports() {
        return this.ports;
    }
    /**
     * The ID of the project in which to create the device
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The ID of the project in which to create the device
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the equinix-metal.ProjectSshKey resource.
     * 
     */
    @Export(name="projectSshKeyIds", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> projectSshKeyIds;

    /**
     * @return Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the equinix-metal.ProjectSshKey resource.
     * 
     */
    public Output<Optional<List<String>>> projectSshKeyIds() {
        return Codegen.optional(this.projectSshKeyIds);
    }
    /**
     * Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.
     * 
     */
    @Export(name="reinstall", type=DeviceReinstall.class, parameters={})
    private Output</* @Nullable */ DeviceReinstall> reinstall;

    /**
     * @return Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.
     * 
     */
    public Output<Optional<DeviceReinstall>> reinstall() {
        return Codegen.optional(this.reinstall);
    }
    /**
     * Root password to the server (disabled after 24 hours)
     * 
     */
    @Export(name="rootPassword", type=String.class, parameters={})
    private Output<String> rootPassword;

    /**
     * @return Root password to the server (disabled after 24 hours)
     * 
     */
    public Output<String> rootPassword() {
        return this.rootPassword;
    }
    /**
     * List of IDs of SSH keys deployed in the device, can be both user and project SSH keys
     * 
     */
    @Export(name="sshKeyIds", type=List.class, parameters={String.class})
    private Output<List<String>> sshKeyIds;

    /**
     * @return List of IDs of SSH keys deployed in the device, can be both user and project SSH keys
     * 
     */
    public Output<List<String>> sshKeyIds() {
        return this.sshKeyIds;
    }
    /**
     * The status of the device
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return The status of the device
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. &#34;4G&#34; or &#34;8M&#34; (for gigabytes and megabytes).
     * 
     */
    @Export(name="storage", type=String.class, parameters={})
    private Output</* @Nullable */ String> storage;

    /**
     * @return JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. &#34;4G&#34; or &#34;8M&#34; (for gigabytes and megabytes).
     * 
     */
    public Output<Optional<String>> storage() {
        return Codegen.optional(this.storage);
    }
    /**
     * Tags attached to the device
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags attached to the device
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Timestamp for device termination. For example `2021-09-03T16:32:00+03:00`. If you don&#39;t supply timezone info, timestamp is assumed to be in UTC.
     * 
     */
    @Export(name="terminationTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> terminationTime;

    /**
     * @return Timestamp for device termination. For example `2021-09-03T16:32:00+03:00`. If you don&#39;t supply timezone info, timestamp is assumed to be in UTC.
     * 
     */
    public Output<Optional<String>> terminationTime() {
        return Codegen.optional(this.terminationTime);
    }
    /**
     * The timestamp for the last time the device was updated
     * 
     */
    @Export(name="updated", type=String.class, parameters={})
    private Output<String> updated;

    /**
     * @return The timestamp for the last time the device was updated
     * 
     */
    public Output<String> updated() {
        return this.updated;
    }
    /**
     * A string of the desired User Data for the device.
     * 
     */
    @Export(name="userData", type=String.class, parameters={})
    private Output</* @Nullable */ String> userData;

    /**
     * @return A string of the desired User Data for the device.
     * 
     */
    public Output<Optional<String>> userData() {
        return Codegen.optional(this.userData);
    }
    /**
     * Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019).
     * 
     */
    @Export(name="waitForReservationDeprovision", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> waitForReservationDeprovision;

    /**
     * @return Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019).
     * 
     */
    public Output<Optional<Boolean>> waitForReservationDeprovision() {
        return Codegen.optional(this.waitForReservationDeprovision);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Device(String name) {
        this(name, DeviceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Device(String name, DeviceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Device(String name, DeviceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/device:Device", name, args == null ? DeviceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Device(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/device:Device", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Device get(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Device(name, id, state, options);
    }
}
