// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetVolume
    {
        /// <summary>
        /// Datasource `equinix-metal.Volume` was removed in version 3.0.0, and the API support was deprecated on June 1st 2021. See https://metal.equinix.com/developers/docs/storage/elastic-block-storage/#elastic-block-storage for more details.
        /// </summary>
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("equinix-metal:index/getVolume:getVolume", args ?? new GetVolumeArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource `equinix-metal.Volume` was removed in version 3.0.0, and the API support was deprecated on June 1st 2021. See https://metal.equinix.com/developers/docs/storage/elastic-block-storage/#elastic-block-storage for more details.
        /// </summary>
        public static Output<GetVolumeResult> Invoke(GetVolumeInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeResult>("equinix-metal:index/getVolume:getVolume", args ?? new GetVolumeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("projectId")]
        public string? ProjectId { get; set; }

        [Input("volumeId")]
        public string? VolumeId { get; set; }

        public GetVolumeArgs()
        {
        }
        public static new GetVolumeArgs Empty => new GetVolumeArgs();
    }

    public sealed class GetVolumeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public GetVolumeInvokeArgs()
        {
        }
        public static new GetVolumeInvokeArgs Empty => new GetVolumeInvokeArgs();
    }


    [OutputType]
    public sealed class GetVolumeResult
    {
        public readonly string BillingCycle;
        public readonly string Created;
        public readonly string Description;
        public readonly ImmutableArray<string> DeviceIds;
        public readonly string Facility;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool Locked;
        public readonly string Name;
        public readonly string Plan;
        public readonly string ProjectId;
        public readonly int Size;
        public readonly ImmutableArray<Outputs.GetVolumeSnapshotPolicyResult> SnapshotPolicies;
        public readonly string State;
        public readonly string Updated;
        public readonly string VolumeId;

        [OutputConstructor]
        private GetVolumeResult(
            string billingCycle,

            string created,

            string description,

            ImmutableArray<string> deviceIds,

            string facility,

            string id,

            bool locked,

            string name,

            string plan,

            string projectId,

            int size,

            ImmutableArray<Outputs.GetVolumeSnapshotPolicyResult> snapshotPolicies,

            string state,

            string updated,

            string volumeId)
        {
            BillingCycle = billingCycle;
            Created = created;
            Description = description;
            DeviceIds = deviceIds;
            Facility = facility;
            Id = id;
            Locked = locked;
            Name = name;
            Plan = plan;
            ProjectId = projectId;
            Size = size;
            SnapshotPolicies = snapshotPolicies;
            State = state;
            Updated = updated;
            VolumeId = volumeId;
        }
    }
}
