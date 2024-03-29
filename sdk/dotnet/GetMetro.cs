// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetMetro
    {
        /// <summary>
        /// Provides an Equinix Metal metro datasource.
        /// </summary>
        public static Task<GetMetroResult> InvokeAsync(GetMetroArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMetroResult>("equinix-metal:index/getMetro:getMetro", args ?? new GetMetroArgs(), options.WithDefaults());

        /// <summary>
        /// Provides an Equinix Metal metro datasource.
        /// </summary>
        public static Output<GetMetroResult> Invoke(GetMetroInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMetroResult>("equinix-metal:index/getMetro:getMetro", args ?? new GetMetroInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetroArgs : global::Pulumi.InvokeArgs
    {
        [Input("capacities")]
        private List<Inputs.GetMetroCapacityArgs>? _capacities;

        /// <summary>
        /// (Optional) Ensure that queried metro has capacity for specified number of given plans
        /// </summary>
        public List<Inputs.GetMetroCapacityArgs> Capacities
        {
            get => _capacities ?? (_capacities = new List<Inputs.GetMetroCapacityArgs>());
            set => _capacities = value;
        }

        /// <summary>
        /// The metro code
        /// 
        /// Metros can be looked up by `code`.
        /// </summary>
        [Input("code", required: true)]
        public string Code { get; set; } = null!;

        public GetMetroArgs()
        {
        }
        public static new GetMetroArgs Empty => new GetMetroArgs();
    }

    public sealed class GetMetroInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("capacities")]
        private InputList<Inputs.GetMetroCapacityInputArgs>? _capacities;

        /// <summary>
        /// (Optional) Ensure that queried metro has capacity for specified number of given plans
        /// </summary>
        public InputList<Inputs.GetMetroCapacityInputArgs> Capacities
        {
            get => _capacities ?? (_capacities = new InputList<Inputs.GetMetroCapacityInputArgs>());
            set => _capacities = value;
        }

        /// <summary>
        /// The metro code
        /// 
        /// Metros can be looked up by `code`.
        /// </summary>
        [Input("code", required: true)]
        public Input<string> Code { get; set; } = null!;

        public GetMetroInvokeArgs()
        {
        }
        public static new GetMetroInvokeArgs Empty => new GetMetroInvokeArgs();
    }


    [OutputType]
    public sealed class GetMetroResult
    {
        /// <summary>
        /// (Optional) Ensure that queried metro has capacity for specified number of given plans
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMetroCapacityResult> Capacities;
        /// <summary>
        /// The code of the metro
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// The country of the metro
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// The ID of the metro
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the metro
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetMetroResult(
            ImmutableArray<Outputs.GetMetroCapacityResult> capacities,

            string code,

            string country,

            string id,

            string name)
        {
            Capacities = capacities;
            Code = code;
            Country = country;
            Id = id;
            Name = name;
        }
    }
}
