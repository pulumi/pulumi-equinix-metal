// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    [EquinixMetalResourceType("equinix-metal:index/volumeAttachment:VolumeAttachment")]
    public partial class VolumeAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the device to which the volume should be attached
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the volume to attach
        /// </summary>
        [Output("volumeId")]
        public Output<string> VolumeId { get; private set; } = null!;


        /// <summary>
        /// Create a VolumeAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VolumeAttachment(string name, VolumeAttachmentArgs args, CustomResourceOptions? options = null)
            : base("equinix-metal:index/volumeAttachment:VolumeAttachment", name, args ?? new VolumeAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VolumeAttachment(string name, Input<string> id, VolumeAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/volumeAttachment:VolumeAttachment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VolumeAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VolumeAttachment Get(string name, Input<string> id, VolumeAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new VolumeAttachment(name, id, state, options);
        }
    }

    public sealed class VolumeAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the device to which the volume should be attached
        /// </summary>
        [Input("deviceId", required: true)]
        public Input<string> DeviceId { get; set; } = null!;

        /// <summary>
        /// The ID of the volume to attach
        /// </summary>
        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public VolumeAttachmentArgs()
        {
        }
    }

    public sealed class VolumeAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the device to which the volume should be attached
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// The ID of the volume to attach
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public VolumeAttachmentState()
        {
        }
    }
}
