// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetSpotMarketPrice
    {
        /// <summary>
        /// Use this data source to get Equinix Metal Spot Market Price.
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
        ///         var example = Output.Create(EquinixMetal.GetSpotMarketPrice.InvokeAsync(new EquinixMetal.GetSpotMarketPriceArgs
        ///         {
        ///             Facility = "ny5",
        ///             Plan = "c3.small.x86",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSpotMarketPriceResult> InvokeAsync(GetSpotMarketPriceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSpotMarketPriceResult>("equinix-metal:index/getSpotMarketPrice:getSpotMarketPrice", args ?? new GetSpotMarketPriceArgs(), options.WithVersion());
    }


    public sealed class GetSpotMarketPriceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the facility.
        /// </summary>
        [Input("facility")]
        public string? Facility { get; set; }

        [Input("metro")]
        public string? Metro { get; set; }

        /// <summary>
        /// Name of the plan.
        /// </summary>
        [Input("plan", required: true)]
        public string Plan { get; set; } = null!;

        public GetSpotMarketPriceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSpotMarketPriceResult
    {
        public readonly string? Facility;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Metro;
        public readonly string Plan;
        /// <summary>
        /// Current spot market price for given plan in given facility.
        /// </summary>
        public readonly double Price;

        [OutputConstructor]
        private GetSpotMarketPriceResult(
            string? facility,

            string id,

            string? metro,

            string plan,

            double price)
        {
            Facility = facility;
            Id = id;
            Metro = metro;
            Plan = plan;
            Price = price;
        }
    }
}
