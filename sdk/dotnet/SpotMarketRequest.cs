// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    /// <summary>
    /// Provides an Equinix Metal Spot Market Request resource to allow you to
    /// manage spot market requests on your account. For more detail on Spot Market, see [this article in Equinix Metal documentation](https://metal.equinix.com/developers/docs/deploy/spot-market/).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using EquinixMetal = Pulumi.EquinixMetal;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a spot market request
    ///     var req = new EquinixMetal.SpotMarketRequest("req", new()
    ///     {
    ///         ProjectId = local.Project_id,
    ///         MaxBidPrice = 0.03,
    ///         Facilities = new[]
    ///         {
    ///             "ny5",
    ///         },
    ///         DevicesMin = 1,
    ///         DevicesMax = 1,
    ///         InstanceParameters = new EquinixMetal.Inputs.SpotMarketRequestInstanceParametersArgs
    ///         {
    ///             Hostname = "testspot",
    ///             BillingCycle = "hourly",
    ///             OperatingSystem = "ubuntu_20_04",
    ///             Plan = "c3.small.x86",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an existing spot market request ID
    /// 
    /// ```sh
    ///  $ pulumi import equinix-metal:index/spotMarketRequest:SpotMarketRequest metal_spot_market_request {existing_spot_market_request_id}
    /// ```
    /// </summary>
    [EquinixMetalResourceType("equinix-metal:index/spotMarketRequest:SpotMarketRequest")]
    public partial class SpotMarketRequest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum number devices to be created
        /// </summary>
        [Output("devicesMax")]
        public Output<int> DevicesMax { get; private set; } = null!;

        /// <summary>
        /// Miniumum number devices to be created
        /// </summary>
        [Output("devicesMin")]
        public Output<int> DevicesMin { get; private set; } = null!;

        /// <summary>
        /// Facility IDs where devices should be created
        /// </summary>
        [Output("facilities")]
        public Output<ImmutableArray<string>> Facilities { get; private set; } = null!;

        /// <summary>
        /// Parameters for devices provisioned from this request. You can find the parameter description from the equinix-metal.Device doc.
        /// </summary>
        [Output("instanceParameters")]
        public Output<Outputs.SpotMarketRequestInstanceParameters> InstanceParameters { get; private set; } = null!;

        /// <summary>
        /// Maximum price user is willing to pay per hour per device
        /// </summary>
        [Output("maxBidPrice")]
        public Output<double> MaxBidPrice { get; private set; } = null!;

        /// <summary>
        /// Metro where devices should be created
        /// </summary>
        [Output("metro")]
        public Output<string?> Metro { get; private set; } = null!;

        /// <summary>
        /// Project ID
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
        /// </summary>
        [Output("waitForDevices")]
        public Output<bool?> WaitForDevices { get; private set; } = null!;


        /// <summary>
        /// Create a SpotMarketRequest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SpotMarketRequest(string name, SpotMarketRequestArgs args, CustomResourceOptions? options = null)
            : base("equinix-metal:index/spotMarketRequest:SpotMarketRequest", name, args ?? new SpotMarketRequestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SpotMarketRequest(string name, Input<string> id, SpotMarketRequestState? state = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/spotMarketRequest:SpotMarketRequest", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SpotMarketRequest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SpotMarketRequest Get(string name, Input<string> id, SpotMarketRequestState? state = null, CustomResourceOptions? options = null)
        {
            return new SpotMarketRequest(name, id, state, options);
        }
    }

    public sealed class SpotMarketRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number devices to be created
        /// </summary>
        [Input("devicesMax", required: true)]
        public Input<int> DevicesMax { get; set; } = null!;

        /// <summary>
        /// Miniumum number devices to be created
        /// </summary>
        [Input("devicesMin", required: true)]
        public Input<int> DevicesMin { get; set; } = null!;

        [Input("facilities")]
        private InputList<string>? _facilities;

        /// <summary>
        /// Facility IDs where devices should be created
        /// </summary>
        public InputList<string> Facilities
        {
            get => _facilities ?? (_facilities = new InputList<string>());
            set => _facilities = value;
        }

        /// <summary>
        /// Parameters for devices provisioned from this request. You can find the parameter description from the equinix-metal.Device doc.
        /// </summary>
        [Input("instanceParameters", required: true)]
        public Input<Inputs.SpotMarketRequestInstanceParametersArgs> InstanceParameters { get; set; } = null!;

        /// <summary>
        /// Maximum price user is willing to pay per hour per device
        /// </summary>
        [Input("maxBidPrice", required: true)]
        public Input<double> MaxBidPrice { get; set; } = null!;

        /// <summary>
        /// Metro where devices should be created
        /// </summary>
        [Input("metro")]
        public Input<string>? Metro { get; set; }

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
        /// </summary>
        [Input("waitForDevices")]
        public Input<bool>? WaitForDevices { get; set; }

        public SpotMarketRequestArgs()
        {
        }
        public static new SpotMarketRequestArgs Empty => new SpotMarketRequestArgs();
    }

    public sealed class SpotMarketRequestState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number devices to be created
        /// </summary>
        [Input("devicesMax")]
        public Input<int>? DevicesMax { get; set; }

        /// <summary>
        /// Miniumum number devices to be created
        /// </summary>
        [Input("devicesMin")]
        public Input<int>? DevicesMin { get; set; }

        [Input("facilities")]
        private InputList<string>? _facilities;

        /// <summary>
        /// Facility IDs where devices should be created
        /// </summary>
        public InputList<string> Facilities
        {
            get => _facilities ?? (_facilities = new InputList<string>());
            set => _facilities = value;
        }

        /// <summary>
        /// Parameters for devices provisioned from this request. You can find the parameter description from the equinix-metal.Device doc.
        /// </summary>
        [Input("instanceParameters")]
        public Input<Inputs.SpotMarketRequestInstanceParametersGetArgs>? InstanceParameters { get; set; }

        /// <summary>
        /// Maximum price user is willing to pay per hour per device
        /// </summary>
        [Input("maxBidPrice")]
        public Input<double>? MaxBidPrice { get; set; }

        /// <summary>
        /// Metro where devices should be created
        /// </summary>
        [Input("metro")]
        public Input<string>? Metro { get; set; }

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
        /// </summary>
        [Input("waitForDevices")]
        public Input<bool>? WaitForDevices { get; set; }

        public SpotMarketRequestState()
        {
        }
        public static new SpotMarketRequestState Empty => new SpotMarketRequestState();
    }
}
