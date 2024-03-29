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
    /// Use this resource to associate VLAN with a Dedicated Port from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/#associating-a-vlan-with-a-dedicated-port).
    /// 
    /// ## Example Usage
    /// 
    /// Pick an existing Project and Connection, create a VLAN and use `equinix-metal.VirtualCircuit` to associate it with a Primary Port of the Connection.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using EquinixMetal = Pulumi.EquinixMetal;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var projectId = "52000fb2-ee46-4673-93a8-de2c2bdba33c";
    /// 
    ///     var connId = "73f12f29-3e19-43a0-8e90-ae81580db1e0";
    /// 
    ///     var testConnection = EquinixMetal.GetConnection.Invoke(new()
    ///     {
    ///         ConnectionId = connId,
    ///     });
    /// 
    ///     var testVlan = new EquinixMetal.Vlan("testVlan", new()
    ///     {
    ///         ProjectId = projectId,
    ///         Metro = testConnection.Apply(getConnectionResult =&gt; getConnectionResult.Metro),
    ///     });
    /// 
    ///     var testVirtualCircuit = new EquinixMetal.VirtualCircuit("testVirtualCircuit", new()
    ///     {
    ///         ConnectionId = connId,
    ///         ProjectId = projectId,
    ///         PortId = testConnection.Apply(getConnectionResult =&gt; getConnectionResult.Ports[0]?.Id),
    ///         VlanId = testVlan.Id,
    ///         NniVlan = 1056,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [EquinixMetalResourceType("equinix-metal:index/virtualCircuit:VirtualCircuit")]
    public partial class VirtualCircuit : global::Pulumi.CustomResource
    {
        /// <summary>
        /// UUID of Connection where the VC is scoped to
        /// </summary>
        [Output("connectionId")]
        public Output<string> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// Description for the Virtual Circuit resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the Virtual Circuit resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Equinix Metal network-to-network VLAN ID
        /// </summary>
        [Output("nniVlan")]
        public Output<int?> NniVlan { get; private set; } = null!;

        /// <summary>
        /// Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        /// </summary>
        [Output("nniVnid")]
        public Output<int> NniVnid { get; private set; } = null!;

        /// <summary>
        /// UUID of the Connection Port where the VC is scoped to
        /// </summary>
        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        /// <summary>
        /// UUID of the Project where the VC is scoped to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Speed of the Virtual Circuit resource
        /// </summary>
        [Output("speed")]
        public Output<string> Speed { get; private set; } = null!;

        /// <summary>
        /// Status of the virtal circuit
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags for the Virtual Circuit resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// UUID of the VLAN to associate
        /// </summary>
        [Output("vlanId")]
        public Output<string> VlanId { get; private set; } = null!;

        /// <summary>
        /// VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        /// </summary>
        [Output("vnid")]
        public Output<int> Vnid { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualCircuit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualCircuit(string name, VirtualCircuitArgs args, CustomResourceOptions? options = null)
            : base("equinix-metal:index/virtualCircuit:VirtualCircuit", name, args ?? new VirtualCircuitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualCircuit(string name, Input<string> id, VirtualCircuitState? state = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/virtualCircuit:VirtualCircuit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualCircuit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualCircuit Get(string name, Input<string> id, VirtualCircuitState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualCircuit(name, id, state, options);
        }
    }

    public sealed class VirtualCircuitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of Connection where the VC is scoped to
        /// </summary>
        [Input("connectionId", required: true)]
        public Input<string> ConnectionId { get; set; } = null!;

        /// <summary>
        /// Description for the Virtual Circuit resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the Virtual Circuit resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Equinix Metal network-to-network VLAN ID
        /// </summary>
        [Input("nniVlan")]
        public Input<int>? NniVlan { get; set; }

        /// <summary>
        /// UUID of the Connection Port where the VC is scoped to
        /// </summary>
        [Input("portId", required: true)]
        public Input<string> PortId { get; set; } = null!;

        /// <summary>
        /// UUID of the Project where the VC is scoped to
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Speed of the Virtual Circuit resource
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags for the Virtual Circuit resource
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// UUID of the VLAN to associate
        /// </summary>
        [Input("vlanId", required: true)]
        public Input<string> VlanId { get; set; } = null!;

        public VirtualCircuitArgs()
        {
        }
        public static new VirtualCircuitArgs Empty => new VirtualCircuitArgs();
    }

    public sealed class VirtualCircuitState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of Connection where the VC is scoped to
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// Description for the Virtual Circuit resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the Virtual Circuit resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Equinix Metal network-to-network VLAN ID
        /// </summary>
        [Input("nniVlan")]
        public Input<int>? NniVlan { get; set; }

        /// <summary>
        /// Nni VLAN ID parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        /// </summary>
        [Input("nniVnid")]
        public Input<int>? NniVnid { get; set; }

        /// <summary>
        /// UUID of the Connection Port where the VC is scoped to
        /// </summary>
        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// UUID of the Project where the VC is scoped to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Speed of the Virtual Circuit resource
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// Status of the virtal circuit
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags for the Virtual Circuit resource
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// UUID of the VLAN to associate
        /// </summary>
        [Input("vlanId")]
        public Input<string>? VlanId { get; set; }

        /// <summary>
        /// VNID VLAN parameter, see https://metal.equinix.com/developers/docs/networking/fabric/
        /// </summary>
        [Input("vnid")]
        public Input<int>? Vnid { get; set; }

        public VirtualCircuitState()
        {
        }
        public static new VirtualCircuitState Empty => new VirtualCircuitState();
    }
}
