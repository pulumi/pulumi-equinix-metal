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
    /// Provides a resource to manage User SSH keys on your Equinix Metal user account. If you create a new device in a project, all the keys of the project's collaborators will be injected to the device.
    /// 
    /// The link between User SSH key and device is implicit. If you want to make sure that a key will be copied to a device, you must ensure that the device resource `depends_on` the key resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using EquinixMetal = Pulumi.EquinixMetal;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new SSH key
    ///     var key1 = new EquinixMetal.SshKey("key1", new()
    ///     {
    ///         PublicKey = File.ReadAllText("/home/terraform/.ssh/id_rsa.pub"),
    ///     });
    /// 
    ///     // Create new device with "key1" included. The device resource "depends_on" the
    ///     // key, in order to make sure the key is created before the device.
    ///     var test = new EquinixMetal.Device("test", new()
    ///     {
    ///         Hostname = "test-device",
    ///         Plan = "c3.small.x86",
    ///         Facilities = new[]
    ///         {
    ///             "sjc1",
    ///         },
    ///         OperatingSystem = "ubuntu_20_04",
    ///         BillingCycle = "hourly",
    ///         ProjectId = local.Project_id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             "metal_ssh_key.key1",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an existing SSH Key ID
    /// 
    /// ```sh
    ///  $ pulumi import equinix-metal:index/sshKey:SshKey metal_ssh_key {existing_sshkey_id}
    /// ```
    /// </summary>
    [EquinixMetalResourceType("equinix-metal:index/sshKey:SshKey")]
    public partial class SshKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp for when the SSH key was created
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the SSH key
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The name of the SSH key for identification
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The UUID of the Equinix Metal API User who owns this key
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The public key. If this is a file, it
        /// can be read using the file interpolation function
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        /// <summary>
        /// The timestamp for the last time the SSH key was updated
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a SshKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SshKey(string name, SshKeyArgs args, CustomResourceOptions? options = null)
            : base("equinix-metal:index/sshKey:SshKey", name, args ?? new SshKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SshKey(string name, Input<string> id, SshKeyState? state = null, CustomResourceOptions? options = null)
            : base("equinix-metal:index/sshKey:SshKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SshKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SshKey Get(string name, Input<string> id, SshKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new SshKey(name, id, state, options);
        }
    }

    public sealed class SshKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the SSH key for identification
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public key. If this is a file, it
        /// can be read using the file interpolation function
        /// </summary>
        [Input("publicKey", required: true)]
        public Input<string> PublicKey { get; set; } = null!;

        public SshKeyArgs()
        {
        }
        public static new SshKeyArgs Empty => new SshKeyArgs();
    }

    public sealed class SshKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp for when the SSH key was created
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// The fingerprint of the SSH key
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The name of the SSH key for identification
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The UUID of the Equinix Metal API User who owns this key
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The public key. If this is a file, it
        /// can be read using the file interpolation function
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// The timestamp for the last time the SSH key was updated
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        public SshKeyState()
        {
        }
        public static new SshKeyState Empty => new SshKeyState();
    }
}
