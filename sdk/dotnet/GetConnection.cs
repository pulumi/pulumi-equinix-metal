// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetConnection
    {
        /// <summary>
        /// Use this data source to retrieve a connection resource from [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)
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
        ///         var example = Output.Create(EquinixMetal.GetConnection.InvokeAsync(new EquinixMetal.GetConnectionArgs
        ///         {
        ///             ConnectionId = "4347e805-eb46-4699-9eb9-5c116e6a017d",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("equinix-metal:index/getConnection:getConnection", args ?? new GetConnectionArgs(), options.WithVersion());
    }


    public sealed class GetConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the connection resource
        /// </summary>
        [Input("connectionId", required: true)]
        public string ConnectionId { get; set; } = null!;

        public GetConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        public readonly string ConnectionId;
        /// <summary>
        /// Description of the connection resource
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Slug of a facility to which the connection belongs
        /// </summary>
        public readonly string Facility;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Port name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of organization to which the connection belongs
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConnectionPortResult> Ports;
        /// <summary>
        /// Connection redundancy, reduntant or primary
        /// </summary>
        public readonly string Redundancy;
        /// <summary>
        /// Port speed in bits per second
        /// </summary>
        public readonly int Speed;
        /// <summary>
        /// Port status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Fabric Token for the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
        /// </summary>
        public readonly string Token;
        /// <summary>
        /// Connection type, dedicated or shared
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectionResult(
            string connectionId,

            string description,

            string facility,

            string id,

            string name,

            string organizationId,

            ImmutableArray<Outputs.GetConnectionPortResult> ports,

            string redundancy,

            int speed,

            string status,

            string token,

            string type)
        {
            ConnectionId = connectionId;
            Description = description;
            Facility = facility;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            Ports = ports;
            Redundancy = redundancy;
            Speed = speed;
            Status = status;
            Token = token;
            Type = type;
        }
    }
}