// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.EquinixMetal
{
    public static class GetPort
    {
        /// <summary>
        /// Use this data source to read ports of existing devices. You can read port by either its UUID, or by a device UUID and port name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Create a device and read it's eth0 port to the datasource.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using EquinixMetal = Pulumi.EquinixMetal;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var projectId = "&lt;UUID_of_your_project&gt;";
        ///         var testDevice = new EquinixMetal.Device("testDevice", new EquinixMetal.DeviceArgs
        ///         {
        ///             Hostname = "tfacc-test-device-port",
        ///             Plan = "c3.medium.x86",
        ///             Facilities = 
        ///             {
        ///                 "sv15",
        ///             },
        ///             OperatingSystem = "ubuntu_20_04",
        ///             BillingCycle = "hourly",
        ///             ProjectId = projectId,
        ///         });
        ///         var testPort = testDevice.Id.Apply(id =&gt; EquinixMetal.GetPort.InvokeAsync(new EquinixMetal.GetPortArgs
        ///         {
        ///             DeviceId = id,
        ///             Name = "eth0",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPortResult> InvokeAsync(GetPortArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPortResult>("equinix-metal:index/getPort:getPort", args ?? new GetPortArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to read ports of existing devices. You can read port by either its UUID, or by a device UUID and port name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Create a device and read it's eth0 port to the datasource.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using EquinixMetal = Pulumi.EquinixMetal;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var projectId = "&lt;UUID_of_your_project&gt;";
        ///         var testDevice = new EquinixMetal.Device("testDevice", new EquinixMetal.DeviceArgs
        ///         {
        ///             Hostname = "tfacc-test-device-port",
        ///             Plan = "c3.medium.x86",
        ///             Facilities = 
        ///             {
        ///                 "sv15",
        ///             },
        ///             OperatingSystem = "ubuntu_20_04",
        ///             BillingCycle = "hourly",
        ///             ProjectId = projectId,
        ///         });
        ///         var testPort = testDevice.Id.Apply(id =&gt; EquinixMetal.GetPort.InvokeAsync(new EquinixMetal.GetPortArgs
        ///         {
        ///             DeviceId = id,
        ///             Name = "eth0",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPortResult> Invoke(GetPortInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPortResult>("equinix-metal:index/getPort:getPort", args ?? new GetPortInvokeArgs(), options.WithVersion());
    }


    public sealed class GetPortArgs : Pulumi.InvokeArgs
    {
        [Input("deviceId")]
        public string? DeviceId { get; set; }

        /// <summary>
        /// Whether to look for public or private block.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("portId")]
        public string? PortId { get; set; }

        public GetPortArgs()
        {
        }
    }

    public sealed class GetPortInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// Whether to look for public or private block.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("portId")]
        public Input<string>? PortId { get; set; }

        public GetPortInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPortResult
    {
        /// <summary>
        /// UUID of the bond port"
        /// </summary>
        public readonly string BondId;
        /// <summary>
        /// Name of the bond port
        /// </summary>
        public readonly string BondName;
        /// <summary>
        /// Flag indicating whether the port is bonded
        /// </summary>
        public readonly bool Bonded;
        public readonly string? DeviceId;
        /// <summary>
        /// Flag indicating whether the port can be removed from a bond
        /// </summary>
        public readonly bool DisbondSupported;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool Layer2;
        /// <summary>
        /// MAC address of the port
        /// </summary>
        public readonly string Mac;
        public readonly string Name;
        /// <summary>
        /// UUID of native VLAN of the port
        /// </summary>
        public readonly string NativeVlanId;
        /// <summary>
        /// One of layer2-bonded, layer2-individual, layer3, hybrid, hybrid-bonded
        /// </summary>
        public readonly string NetworkType;
        public readonly string? PortId;
        /// <summary>
        /// Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// UUIDs of attached VLANs
        /// </summary>
        public readonly ImmutableArray<string> VlanIds;
        /// <summary>
        /// VXLAN ids of attached VLANs
        /// </summary>
        public readonly ImmutableArray<int> VxlanIds;

        [OutputConstructor]
        private GetPortResult(
            string bondId,

            string bondName,

            bool bonded,

            string? deviceId,

            bool disbondSupported,

            string id,

            bool layer2,

            string mac,

            string name,

            string nativeVlanId,

            string networkType,

            string? portId,

            string type,

            ImmutableArray<string> vlanIds,

            ImmutableArray<int> vxlanIds)
        {
            BondId = bondId;
            BondName = bondName;
            Bonded = bonded;
            DeviceId = deviceId;
            DisbondSupported = disbondSupported;
            Id = id;
            Layer2 = layer2;
            Mac = mac;
            Name = name;
            NativeVlanId = nativeVlanId;
            NetworkType = networkType;
            PortId = portId;
            Type = type;
            VlanIds = vlanIds;
            VxlanIds = vxlanIds;
        }
    }
}
