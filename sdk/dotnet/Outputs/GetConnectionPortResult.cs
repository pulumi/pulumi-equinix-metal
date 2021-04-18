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
    public sealed class GetConnectionPortResult
    {
        /// <summary>
        /// Port link status
        /// </summary>
        public readonly string LinkStatus;
        /// <summary>
        /// Port name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Port role - primary or secondary
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Port speed in bits per second
        /// </summary>
        public readonly int Speed;
        /// <summary>
        /// Port status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// List of IDs of virtual cicruits attached to this port
        /// </summary>
        public readonly ImmutableArray<object> VirtualCircuitIds;

        [OutputConstructor]
        private GetConnectionPortResult(
            string linkStatus,

            string name,

            string role,

            int speed,

            string status,

            ImmutableArray<object> virtualCircuitIds)
        {
            LinkStatus = linkStatus;
            Name = name;
            Role = role;
            Speed = speed;
            Status = status;
            VirtualCircuitIds = virtualCircuitIds;
        }
    }
}
