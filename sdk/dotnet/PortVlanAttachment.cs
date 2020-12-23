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
    /// Provides a resource to attach device ports to VLANs.
    /// 
    /// Device and VLAN must be in the same facility.
    /// 
    /// If you need this resource to add the port back to bond on removal, set `force_bond = true`.
    /// 
    /// To learn more about Layer 2 networking in Equinix Metal, refer to
    /// 
    /// * &lt;https://metal.equinix.com/developers/docs/networking/layer2/&gt;
    /// * &lt;https://metal.equinix.com/developers/docs/networking/layer2-configs/&gt;
    /// 
    /// ## Example Usage
    /// ### Hybrid network type
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using EquinixMetal = Pulumi.EquinixMetal;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testVlan = new EquinixMetal.Vlan("testVlan", new EquinixMetal.VlanArgs
    ///         {
    ///             Description = "VLAN in New Jersey",
    ///             Facility = "ewr1",
    ///             ProjectId = local.Project_id,
    ///         });
    ///         var testDevice = new EquinixMetal.Device("testDevice", new EquinixMetal.DeviceArgs
    ///         {
    ///             Hostname = "test",
    ///             Plan = "m1.xlarge.x86",
    ///             Facilities = 
    ///             {
    ///                 "ewr1",
    ///             },
    ///             OperatingSystem = "ubuntu_16_04",
    ///             BillingCycle = "hourly",
    ///             ProjectId = local.Project_id,
    ///         });
    ///         var testDeviceNetworkType = new EquinixMetal.DeviceNetworkType("testDeviceNetworkType", new EquinixMetal.DeviceNetworkTypeArgs
    ///         {
    ///             DeviceId = testDevice.Id,
    ///             Type = "hybrid",
    ///         });
    ///         var testPortVlanAttachment = new EquinixMetal.PortVlanAttachment("testPortVlanAttachment", new EquinixMetal.PortVlanAttachmentArgs
    ///         {
    ///             DeviceId = testDeviceNetworkType.Id,
    ///             PortName = "eth1",
    ///             VlanVnid = testVlan.Vxlan,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Layer 2 network
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using EquinixMetal = Pulumi.EquinixMetal;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testDevice = new EquinixMetal.Device("testDevice", new EquinixMetal.DeviceArgs
    ///         {
    ///             Hostname = "test",
    ///             Plan = "m1.xlarge.x86",
    ///             Facilities = 
    ///             {
    ///                 "ewr1",
    ///             },
    ///             OperatingSystem = "ubuntu_16_04",
    ///             BillingCycle = "hourly",
    ///             ProjectId = local.Project_id,
    ///         });
    ///         var testDeviceNetworkType = new EquinixMetal.DeviceNetworkType("testDeviceNetworkType", new EquinixMetal.DeviceNetworkTypeArgs
    ///         {
    ///             DeviceId = testDevice.Id,
    ///             Type = "layer2-individual",
    ///         });
    ///         var test1Vlan = new EquinixMetal.Vlan("test1Vlan", new EquinixMetal.VlanArgs
    ///         {
    ///             Description = "VLAN in New Jersey",
    ///             Facility = "ewr1",
    ///             ProjectId = local.Project_id,
    ///         });
    ///         var test2Vlan = new EquinixMetal.Vlan("test2Vlan", new EquinixMetal.VlanArgs
    ///         {
    ///             Description = "VLAN in New Jersey",
    ///             Facility = "ewr1",
    ///             ProjectId = local.Project_id,
    ///         });
    ///         var test1PortVlanAttachment = new EquinixMetal.PortVlanAttachment("test1PortVlanAttachment", new EquinixMetal.PortVlanAttachmentArgs
    ///         {
    ///             DeviceId = testDeviceNetworkType.Id,
    ///             VlanVnid = test1Vlan.Vxlan,
    ///             PortName = "eth1",
    ///         });
    ///         var test2PortVlanAttachment = new EquinixMetal.PortVlanAttachment("test2PortVlanAttachment", new EquinixMetal.PortVlanAttachmentArgs
    ///         {
    ///             DeviceId = testDeviceNetworkType.Id,
    ///             VlanVnid = test2Vlan.Vxlan,
    ///             PortName = "eth1",
    ///             Native = true,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 "metal_port_vlan_attachment.test1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Attribute Referece
    /// 
    /// * `id` - UUID of device port used in the assignment
    /// * `vlan_id` - UUID of VLAN API resource
    /// * `port_id` - UUID of device port
    /// </summary>
    public partial class PortVlanAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of device to be assigned to the VLAN
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;

        /// <summary>
        /// Add port back to the bond when this resource is removed. Default is false.
        /// </summary>
        [Output("forceBond")]
        public Output<bool?> ForceBond { get; private set; } = null!;

        /// <summary>
        /// Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        /// </summary>
        [Output("native")]
        public Output<bool?> Native { get; private set; } = null!;

        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        /// <summary>
        /// Name of network port to be assigned to the VLAN
        /// </summary>
        [Output("portName")]
        public Output<string> PortName { get; private set; } = null!;

        [Output("vlanId")]
        public Output<string> VlanId { get; private set; } = null!;

        /// <summary>
        /// VXLAN Network Identifier, integer
        /// </summary>
        [Output("vlanVnid")]
        public Output<int> VlanVnid { get; private set; } = null!;


        /// <summary>
        /// Create a PortVlanAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PortVlanAttachment(string name, PortVlanAttachmentArgs args, CustomResourceOptions? options = null)
            : base("equinix-metal:index/portVlanAttachment:PortVlanAttachment", name, args ?? new PortVlanAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PortVlanAttachment(string name, Input<string> id, PortVlanAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/portVlanAttachment:PortVlanAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PortVlanAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PortVlanAttachment Get(string name, Input<string> id, PortVlanAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new PortVlanAttachment(name, id, state, options);
        }
    }

    public sealed class PortVlanAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of device to be assigned to the VLAN
        /// </summary>
        [Input("deviceId", required: true)]
        public Input<string> DeviceId { get; set; } = null!;

        /// <summary>
        /// Add port back to the bond when this resource is removed. Default is false.
        /// </summary>
        [Input("forceBond")]
        public Input<bool>? ForceBond { get; set; }

        /// <summary>
        /// Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        /// </summary>
        [Input("native")]
        public Input<bool>? Native { get; set; }

        /// <summary>
        /// Name of network port to be assigned to the VLAN
        /// </summary>
        [Input("portName", required: true)]
        public Input<string> PortName { get; set; } = null!;

        /// <summary>
        /// VXLAN Network Identifier, integer
        /// </summary>
        [Input("vlanVnid", required: true)]
        public Input<int> VlanVnid { get; set; } = null!;

        public PortVlanAttachmentArgs()
        {
        }
    }

    public sealed class PortVlanAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of device to be assigned to the VLAN
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// Add port back to the bond when this resource is removed. Default is false.
        /// </summary>
        [Input("forceBond")]
        public Input<bool>? ForceBond { get; set; }

        /// <summary>
        /// Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use `depends_on` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above.
        /// </summary>
        [Input("native")]
        public Input<bool>? Native { get; set; }

        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// Name of network port to be assigned to the VLAN
        /// </summary>
        [Input("portName")]
        public Input<string>? PortName { get; set; }

        [Input("vlanId")]
        public Input<string>? VlanId { get; set; }

        /// <summary>
        /// VXLAN Network Identifier, integer
        /// </summary>
        [Input("vlanVnid")]
        public Input<int>? VlanVnid { get; set; }

        public PortVlanAttachmentState()
        {
        }
    }
}
