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

    public sealed class KubeconfigKubeProxyOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public KubeconfigKubeProxyOptionsArgs()
        {
        }
        public static new KubeconfigKubeProxyOptionsArgs Empty => new KubeconfigKubeProxyOptionsArgs();
    }
}
