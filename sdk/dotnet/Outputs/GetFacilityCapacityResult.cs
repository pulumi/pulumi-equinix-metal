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
    public sealed class GetFacilityCapacityResult
    {
        /// <summary>
        /// device plan to check
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// number of device to check
        /// </summary>
        public readonly int? Quantity;

        [OutputConstructor]
        private GetFacilityCapacityResult(
            string plan,

            int? quantity)
        {
            Plan = plan;
            Quantity = quantity;
        }
    }
}