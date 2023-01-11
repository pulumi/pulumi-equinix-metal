// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.equinixmetal.SshKeyArgs;
import com.pulumi.equinixmetal.Utilities;
import com.pulumi.equinixmetal.inputs.SshKeyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage User SSH keys on your Equinix Metal user account. If you create a new device in a project, all the keys of the project&#39;s collaborators will be injected to the device.
 * 
 * The link between User SSH key and device is implicit. If you want to make sure that a key will be copied to a device, you must ensure that the device resource `depends_on` the key resource.
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * This resource can be imported using an existing SSH Key ID
 * 
 * ```sh
 *  $ pulumi import equinix-metal:index/sshKey:SshKey metal_ssh_key {existing_sshkey_id}
 * ```
 * 
 */
@ResourceType(type="equinix-metal:index/sshKey:SshKey")
public class SshKey extends com.pulumi.resources.CustomResource {
    /**
     * The timestamp for when the SSH key was created
     * 
     */
    @Export(name="created", type=String.class, parameters={})
    private Output<String> created;

    /**
     * @return The timestamp for when the SSH key was created
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * The fingerprint of the SSH key
     * 
     */
    @Export(name="fingerprint", type=String.class, parameters={})
    private Output<String> fingerprint;

    /**
     * @return The fingerprint of the SSH key
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * The name of the SSH key for identification
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the SSH key for identification
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The UUID of the Equinix Metal API User who owns this key
     * 
     */
    @Export(name="ownerId", type=String.class, parameters={})
    private Output<String> ownerId;

    /**
     * @return The UUID of the Equinix Metal API User who owns this key
     * 
     */
    public Output<String> ownerId() {
        return this.ownerId;
    }
    /**
     * The public key. If this is a file, it
     * can be read using the file interpolation function
     * 
     */
    @Export(name="publicKey", type=String.class, parameters={})
    private Output<String> publicKey;

    /**
     * @return The public key. If this is a file, it
     * can be read using the file interpolation function
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }
    /**
     * The timestamp for the last time the SSH key was updated
     * 
     */
    @Export(name="updated", type=String.class, parameters={})
    private Output<String> updated;

    /**
     * @return The timestamp for the last time the SSH key was updated
     * 
     */
    public Output<String> updated() {
        return this.updated;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SshKey(String name) {
        this(name, SshKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SshKey(String name, SshKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SshKey(String name, SshKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/sshKey:SshKey", name, args == null ? SshKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SshKey(String name, Output<String> id, @Nullable SshKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/sshKey:SshKey", name, state, makeResourceOptions(options, id));
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
    public static SshKey get(String name, Output<String> id, @Nullable SshKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SshKey(name, id, state, options);
    }
}