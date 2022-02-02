// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetGateway
    {
        /// <summary>
        /// Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
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
        ///         // Create Metal Gateway for a VLAN with a private IPv4 block with 8 IP addresses
        ///         var testVlan = new EquinixMetal.Vlan("testVlan", new EquinixMetal.VlanArgs
        ///         {
        ///             Description = "test VLAN in SV",
        ///             Metro = "sv",
        ///             ProjectId = local.Project_id,
        ///         });
        ///         var testGateway = Output.Create(EquinixMetal.GetGateway.InvokeAsync(new EquinixMetal.GetGatewayArgs
        ///         {
        ///             GatewayId = local.Gateway_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGatewayResult> InvokeAsync(GetGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewayResult>("equinix-metal:index/getGateway:getGateway", args ?? new GetGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
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
        ///         // Create Metal Gateway for a VLAN with a private IPv4 block with 8 IP addresses
        ///         var testVlan = new EquinixMetal.Vlan("testVlan", new EquinixMetal.VlanArgs
        ///         {
        ///             Description = "test VLAN in SV",
        ///             Metro = "sv",
        ///             ProjectId = local.Project_id,
        ///         });
        ///         var testGateway = Output.Create(EquinixMetal.GetGateway.InvokeAsync(new EquinixMetal.GetGatewayArgs
        ///         {
        ///             GatewayId = local.Gateway_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGatewayResult> Invoke(GetGatewayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGatewayResult>("equinix-metal:index/getGateway:getGateway", args ?? new GetGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the metal gateway resource to retrieve
        /// </summary>
        [Input("gatewayId", required: true)]
        public string GatewayId { get; set; } = null!;

        public GetGatewayArgs()
        {
        }
    }

    public sealed class GetGatewayInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the metal gateway resource to retrieve
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        public GetGatewayInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewayResult
    {
        public readonly string GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// UUID of IP reservation block bound to the gateway
        /// </summary>
        public readonly string IpReservationId;
        /// <summary>
        /// Size of the private IPv4 subnet bound to this metal gateway, one of (8, 16, 32, 64, 128)`
        /// </summary>
        public readonly int PrivateIpv4SubnetSize;
        /// <summary>
        /// UUID of the project where the gateway is scoped to
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Status of the gateway resource
        /// </summary>
        public readonly string State;
        /// <summary>
        /// UUID of the VLAN where the gateway is scoped to
        /// </summary>
        public readonly string VlanId;

        [OutputConstructor]
        private GetGatewayResult(
            string gatewayId,

            string id,

            string ipReservationId,

            int privateIpv4SubnetSize,

            string projectId,

            string state,

            string vlanId)
        {
            GatewayId = gatewayId;
            Id = id;
            IpReservationId = ipReservationId;
            PrivateIpv4SubnetSize = privateIpv4SubnetSize;
            ProjectId = projectId;
            State = state;
            VlanId = vlanId;
        }
    }
}
