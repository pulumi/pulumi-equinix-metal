// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal.Outputs
{

    [OutputType]
    public sealed class SpotMarketRequestInstanceParameters
    {
        public readonly bool? AlwaysPxe;
        public readonly string BillingCycle;
        public readonly string? Customdata;
        public readonly string? Description;
        public readonly ImmutableArray<string> Features;
        public readonly string Hostname;
        public readonly string? IpxeScriptUrl;
        /// <summary>
        /// Blocks deletion of the SpotMarketRequest device until the lock is disabled
        /// </summary>
        public readonly bool? Locked;
        public readonly string OperatingSystem;
        public readonly string Plan;
        public readonly ImmutableArray<string> ProjectSshKeys;
        public readonly ImmutableArray<string> Tags;
        public readonly string? TermintationTime;
        public readonly ImmutableArray<string> UserSshKeys;
        public readonly string? Userdata;

        [OutputConstructor]
        private SpotMarketRequestInstanceParameters(
            bool? alwaysPxe,

            string billingCycle,

            string? customdata,

            string? description,

            ImmutableArray<string> features,

            string hostname,

            string? ipxeScriptUrl,

            bool? locked,

            string operatingSystem,

            string plan,

            ImmutableArray<string> projectSshKeys,

            ImmutableArray<string> tags,

            string? termintationTime,

            ImmutableArray<string> userSshKeys,

            string? userdata)
        {
            AlwaysPxe = alwaysPxe;
            BillingCycle = billingCycle;
            Customdata = customdata;
            Description = description;
            Features = features;
            Hostname = hostname;
            IpxeScriptUrl = ipxeScriptUrl;
            Locked = locked;
            OperatingSystem = operatingSystem;
            Plan = plan;
            ProjectSshKeys = projectSshKeys;
            Tags = tags;
            TermintationTime = termintationTime;
            UserSshKeys = userSshKeys;
            Userdata = userdata;
        }
    }
}
