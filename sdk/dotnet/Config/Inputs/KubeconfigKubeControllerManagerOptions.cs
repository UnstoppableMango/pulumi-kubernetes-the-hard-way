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

    public sealed class KubeconfigKubeControllerManagerOptions : global::Pulumi.InvokeArgs
    {
        [Input("publicIp")]
        public string? PublicIp { get; set; }

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public KubeconfigKubeControllerManagerOptions()
        {
        }
        public static new KubeconfigKubeControllerManagerOptions Empty => new KubeconfigKubeControllerManagerOptions();
    }
}
