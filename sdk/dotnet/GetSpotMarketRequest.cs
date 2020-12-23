// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetSpotMarketRequest
    {
        public static Task<GetSpotMarketRequestResult> InvokeAsync(GetSpotMarketRequestArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSpotMarketRequestResult>("equinix-metal:index/getSpotMarketRequest:getSpotMarketRequest", args ?? new GetSpotMarketRequestArgs(), options.WithVersion());
    }


    public sealed class GetSpotMarketRequestArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the Spot Market Request
        /// </summary>
        [Input("requestId", required: true)]
        public string RequestId { get; set; } = null!;

        public GetSpotMarketRequestArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSpotMarketRequestResult
    {
        /// <summary>
        /// List of IDs of devices spawned by the referenced Spot Market Request
        /// </summary>
        public readonly ImmutableArray<string> DeviceIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RequestId;

        [OutputConstructor]
        private GetSpotMarketRequestResult(
            ImmutableArray<string> deviceIds,

            string id,

            string requestId)
        {
            DeviceIds = deviceIds;
            Id = id;
            RequestId = requestId;
        }
    }
}
