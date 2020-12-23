// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal
{
    public static class GetOperatingSystem
    {
        /// <summary>
        /// Use this data source to get Equinix Metal Operating System image.
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
        ///         var example = Output.Create(EquinixMetal.GetOperatingSystem.InvokeAsync(new EquinixMetal.GetOperatingSystemArgs
        ///         {
        ///             Name = "Container Linux",
        ///             Distro = "coreos",
        ///             Version = "alpha",
        ///             ProvisionableOn = "c1.small.x86",
        ///         }));
        ///         var server = new EquinixMetal.Device("server", new EquinixMetal.DeviceArgs
        ///         {
        ///             Hostname = "tf.coreos2",
        ///             Plan = "c1.small.x86",
        ///             Facilities = 
        ///             {
        ///                 "ewr1",
        ///             },
        ///             OperatingSystem = example.Apply(example =&gt; example.Id),
        ///             BillingCycle = "hourly",
        ///             ProjectId = local.Project_id,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOperatingSystemResult> InvokeAsync(GetOperatingSystemArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOperatingSystemResult>("equinix-metal:index/getOperatingSystem:getOperatingSystem", args ?? new GetOperatingSystemArgs(), options.WithVersion());
    }


    public sealed class GetOperatingSystemArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the OS distribution.
        /// </summary>
        [Input("distro")]
        public string? Distro { get; set; }

        /// <summary>
        /// Name or part of the name of the distribution. Case insensitive.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Plan name.
        /// </summary>
        [Input("provisionableOn")]
        public string? ProvisionableOn { get; set; }

        /// <summary>
        /// Version of the distribution
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        public GetOperatingSystemArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOperatingSystemResult
    {
        public readonly string? Distro;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? ProvisionableOn;
        /// <summary>
        /// Operating system slug (same as `id`)
        /// </summary>
        public readonly string Slug;
        public readonly string? Version;

        [OutputConstructor]
        private GetOperatingSystemResult(
            string? distro,

            string id,

            string? name,

            string? provisionableOn,

            string slug,

            string? version)
        {
            Distro = distro;
            Id = id;
            Name = name;
            ProvisionableOn = provisionableOn;
            Slug = slug;
            Version = version;
        }
    }
}
