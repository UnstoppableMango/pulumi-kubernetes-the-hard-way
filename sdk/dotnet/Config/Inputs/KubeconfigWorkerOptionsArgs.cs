// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config.Inputs
{

    public sealed class KubeconfigWorkerOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("publicIp", required: true)]
        public Input<string> PublicIp { get; set; } = null!;

        [Input("type")]
        public string? Type { get; set; }

        public KubeconfigWorkerOptionsArgs()
        {
        }
        public static new KubeconfigWorkerOptionsArgs Empty => new KubeconfigWorkerOptionsArgs();
    }
}
