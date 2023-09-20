// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.equinixmetal.Utilities;
import com.pulumi.equinixmetal.VolumeAttachmentArgs;
import com.pulumi.equinixmetal.inputs.VolumeAttachmentState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Resource `equinix-metal.VolumeAttachment` was removed in version 3.0.0, and the API support was deprecated on June 1st 2021. See https://metal.equinix.com/developers/docs/storage/elastic-block-storage/#elastic-block-storage for more details.
 * 
 */
@ResourceType(type="equinix-metal:index/volumeAttachment:VolumeAttachment")
public class VolumeAttachment extends com.pulumi.resources.CustomResource {
    @Export(name="deviceId", refs={String.class}, tree="[0]")
    private Output<String> deviceId;

    public Output<String> deviceId() {
        return this.deviceId;
    }
    @Export(name="volumeId", refs={String.class}, tree="[0]")
    private Output<String> volumeId;

    public Output<String> volumeId() {
        return this.volumeId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VolumeAttachment(String name) {
        this(name, VolumeAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VolumeAttachment(String name, VolumeAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VolumeAttachment(String name, VolumeAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/volumeAttachment:VolumeAttachment", name, args == null ? VolumeAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VolumeAttachment(String name, Output<String> id, @Nullable VolumeAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/volumeAttachment:VolumeAttachment", name, state, makeResourceOptions(options, id));
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
    public static VolumeAttachment get(String name, Output<String> id, @Nullable VolumeAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VolumeAttachment(name, id, state, options);
    }
}
