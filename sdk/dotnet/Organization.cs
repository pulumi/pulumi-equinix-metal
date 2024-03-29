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
    /// Provides a resource to manage organization resource in Equinix Metal.
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
    ///     // Create a new Project
    ///     var tfOrganization1 = new EquinixMetal.Organization("tfOrganization1", new()
    ///     {
    ///         Description = "quux",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an existing organization ID
    /// 
    /// ```sh
    ///  $ pulumi import equinix-metal:index/organization:Organization metal_organization {existing_organization_id}
    /// ```
    /// </summary>
    [EquinixMetalResourceType("equinix-metal:index/organization:Organization")]
    public partial class Organization : global::Pulumi.CustomResource
    {
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Description string
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Logo URL
        /// </summary>
        [Output("logo")]
        public Output<string?> Logo { get; private set; } = null!;

        /// <summary>
        /// The name of the Organization
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Twitter handle
        /// </summary>
        [Output("twitter")]
        public Output<string?> Twitter { get; private set; } = null!;

        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;

        /// <summary>
        /// Website link
        /// </summary>
        [Output("website")]
        public Output<string?> Website { get; private set; } = null!;


        /// <summary>
        /// Create a Organization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Organization(string name, OrganizationArgs? args = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/organization:Organization", name, args ?? new OrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Organization(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/organization:Organization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Organization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Organization Get(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
        {
            return new Organization(name, id, state, options);
        }
    }

    public sealed class OrganizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description string
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Logo URL
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// The name of the Organization
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Twitter handle
        /// </summary>
        [Input("twitter")]
        public Input<string>? Twitter { get; set; }

        /// <summary>
        /// Website link
        /// </summary>
        [Input("website")]
        public Input<string>? Website { get; set; }

        public OrganizationArgs()
        {
        }
        public static new OrganizationArgs Empty => new OrganizationArgs();
    }

    public sealed class OrganizationState : global::Pulumi.ResourceArgs
    {
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Description string
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Logo URL
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// The name of the Organization
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Twitter handle
        /// </summary>
        [Input("twitter")]
        public Input<string>? Twitter { get; set; }

        [Input("updated")]
        public Input<string>? Updated { get; set; }

        /// <summary>
        /// Website link
        /// </summary>
        [Input("website")]
        public Input<string>? Website { get; set; }

        public OrganizationState()
        {
        }
        public static new OrganizationState Empty => new OrganizationState();
    }
}
