// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal.Inputs
{

    public sealed class GetMetroCapacityInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// device plan to check
        /// </summary>
        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        /// <summary>
        /// number of device to check
        /// </summary>
        [Input("quantity")]
        public Input<int>? Quantity { get; set; }

        public GetMetroCapacityInputArgs()
        {
        }
    }
}
