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
    public sealed class DeviceReinstall
    {
        /// <summary>
        /// Whether the OS disk should be filled with `00h` bytes before reinstall. Defaults to `false`.
        /// *
        /// </summary>
        public readonly bool? DeprovisionFast;
        /// <summary>
        /// Whether the provider should favour reinstall over destroy and create. Defaults to `false`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Whether the non-OS disks should be kept or wiped during reinstall. Defaults to `false`.
        /// </summary>
        public readonly bool? PreserveData;

        [OutputConstructor]
        private DeviceReinstall(
            bool? deprovisionFast,

            bool? enabled,

            bool? preserveData)
        {
            DeprovisionFast = deprovisionFast;
            Enabled = enabled;
            PreserveData = preserveData;
        }
    }
}
