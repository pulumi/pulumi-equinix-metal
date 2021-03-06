// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetDeviceBgpNeighbors
    {
        /// <summary>
        /// Use this datasource to retrieve list of BGP neighbors of a device in the Equinix Metal host.
        /// 
        /// To have any BGP neighbors listed, the device must be in BGP-enabled project and have a BGP session assigned.
        /// 
        /// To learn more about using BGP in Equinix Metal, see the equinix-metal.BgpSession resource documentation.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using EquinixMetal = Pulumi.EquinixMetal;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(EquinixMetal.GetDeviceBgpNeighbors.InvokeAsync(new EquinixMetal.GetDeviceBgpNeighborsArgs
        ///         {
        ///             DeviceId = "4c641195-25e5-4c3c-b2b7-4cd7a42c7b40",
        ///         }));
        ///         this.BgpNeighborsListing = test.Apply(test =&gt; test.BgpNeighbors);
        ///     }
        /// 
        ///     [Output("bgpNeighborsListing")]
        ///     public Output&lt;string&gt; BgpNeighborsListing { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDeviceBgpNeighborsResult> InvokeAsync(GetDeviceBgpNeighborsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeviceBgpNeighborsResult>("equinix-metal:index/getDeviceBgpNeighbors:getDeviceBgpNeighbors", args ?? new GetDeviceBgpNeighborsArgs(), options.WithVersion());
    }


    public sealed class GetDeviceBgpNeighborsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of BGP-enabled device whose neighbors to list
        /// </summary>
        [Input("deviceId", required: true)]
        public string DeviceId { get; set; } = null!;

        public GetDeviceBgpNeighborsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeviceBgpNeighborsResult
    {
        /// <summary>
        /// array of BGP neighbor records with attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDeviceBgpNeighborsBgpNeighborResult> BgpNeighbors;
        public readonly string DeviceId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDeviceBgpNeighborsResult(
            ImmutableArray<Outputs.GetDeviceBgpNeighborsBgpNeighborResult> bgpNeighbors,

            string deviceId,

            string id)
        {
            BgpNeighbors = bgpNeighbors;
            DeviceId = deviceId;
            Id = id;
        }
    }
}
