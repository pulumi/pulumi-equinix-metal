// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    [EquinixMetalResourceType("equinix-metal:index/volume:Volume")]
    public partial class Volume : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of attachments, each with it's own `href` attribute
        /// </summary>
        [Output("attachments")]
        public Output<ImmutableArray<Outputs.VolumeAttachment>> Attachments { get; private set; } = null!;

        /// <summary>
        /// The billing cycle, defaults to "hourly"
        /// </summary>
        [Output("billingCycle")]
        public Output<string> BillingCycle { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the volume was created
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Optional description for the volume
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The facility to create the volume in
        /// </summary>
        [Output("facility")]
        public Output<string> Facility { get; private set; } = null!;

        /// <summary>
        /// Lock or unlock the volume
        /// </summary>
        [Output("locked")]
        public Output<bool?> Locked { get; private set; } = null!;

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The service plan slug of the volume
        /// </summary>
        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        /// <summary>
        /// The metal project ID to deploy the volume in
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The size in GB to make the volume
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Optional list of snapshot policies
        /// </summary>
        [Output("snapshotPolicies")]
        public Output<ImmutableArray<Outputs.VolumeSnapshotPolicy>> SnapshotPolicies { get; private set; } = null!;

        /// <summary>
        /// The state of the volume
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The timestamp for the last time the volume was updated
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a Volume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Volume(string name, VolumeArgs args, CustomResourceOptions? options = null)
            : base("equinix-metal:index/volume:Volume", name, args ?? new VolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Volume(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/volume:Volume", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Volume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Volume Get(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new Volume(name, id, state, options);
        }
    }

    public sealed class VolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing cycle, defaults to "hourly"
        /// </summary>
        [Input("billingCycle")]
        public InputUnion<string, Pulumi.EquinixMetal.BillingCycle>? BillingCycle { get; set; }

        /// <summary>
        /// Optional description for the volume
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The facility to create the volume in
        /// </summary>
        [Input("facility", required: true)]
        public InputUnion<string, Pulumi.EquinixMetal.Facility> Facility { get; set; } = null!;

        /// <summary>
        /// Lock or unlock the volume
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// The service plan slug of the volume
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        /// <summary>
        /// The metal project ID to deploy the volume in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The size in GB to make the volume
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("snapshotPolicies")]
        private InputList<Inputs.VolumeSnapshotPolicyArgs>? _snapshotPolicies;

        /// <summary>
        /// Optional list of snapshot policies
        /// </summary>
        public InputList<Inputs.VolumeSnapshotPolicyArgs> SnapshotPolicies
        {
            get => _snapshotPolicies ?? (_snapshotPolicies = new InputList<Inputs.VolumeSnapshotPolicyArgs>());
            set => _snapshotPolicies = value;
        }

        public VolumeArgs()
        {
        }
    }

    public sealed class VolumeState : Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputList<Inputs.VolumeAttachmentGetArgs>? _attachments;

        /// <summary>
        /// A list of attachments, each with it's own `href` attribute
        /// </summary>
        public InputList<Inputs.VolumeAttachmentGetArgs> Attachments
        {
            get => _attachments ?? (_attachments = new InputList<Inputs.VolumeAttachmentGetArgs>());
            set => _attachments = value;
        }

        /// <summary>
        /// The billing cycle, defaults to "hourly"
        /// </summary>
        [Input("billingCycle")]
        public InputUnion<string, Pulumi.EquinixMetal.BillingCycle>? BillingCycle { get; set; }

        /// <summary>
        /// The timestamp for when the volume was created
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Optional description for the volume
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The facility to create the volume in
        /// </summary>
        [Input("facility")]
        public InputUnion<string, Pulumi.EquinixMetal.Facility>? Facility { get; set; }

        /// <summary>
        /// Lock or unlock the volume
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The service plan slug of the volume
        /// </summary>
        [Input("plan")]
        public Input<string>? Plan { get; set; }

        /// <summary>
        /// The metal project ID to deploy the volume in
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The size in GB to make the volume
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("snapshotPolicies")]
        private InputList<Inputs.VolumeSnapshotPolicyGetArgs>? _snapshotPolicies;

        /// <summary>
        /// Optional list of snapshot policies
        /// </summary>
        public InputList<Inputs.VolumeSnapshotPolicyGetArgs> SnapshotPolicies
        {
            get => _snapshotPolicies ?? (_snapshotPolicies = new InputList<Inputs.VolumeSnapshotPolicyGetArgs>());
            set => _snapshotPolicies = value;
        }

        /// <summary>
        /// The state of the volume
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The timestamp for the last time the volume was updated
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        public VolumeState()
        {
        }
    }
}
