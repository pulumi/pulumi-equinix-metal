// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.EquinixMetal.Inputs
{

    public sealed class VolumeAttachmentGetArgs : Pulumi.ResourceArgs
    {
        [Input("href")]
        public Input<string>? Href { get; set; }

        public VolumeAttachmentGetArgs()
        {
        }
    }
}
