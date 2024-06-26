// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Remote.Inputs
{

    /// <summary>
    /// Props for resources that consume kubelet configuration.
    /// </summary>
    public sealed class KubeletConfigurationPropsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to the kubelet configuration.
        /// </summary>
        [Input("configurationFilePath", required: true)]
        public Input<string> ConfigurationFilePath { get; set; } = null!;

        /// <summary>
        /// Path to the kubeconfig the kubelet will use
        /// </summary>
        [Input("kubeconfigPath", required: true)]
        public Input<string> KubeconfigPath { get; set; } = null!;

        /// <summary>
        /// Path to the kubelet binary.
        /// </summary>
        [Input("kubeletPath", required: true)]
        public Input<string> KubeletPath { get; set; } = null!;

        /// <summary>
        /// Whether to register the node. Defaults to `true`.
        /// </summary>
        [Input("registerNode", required: true)]
        public Input<bool> RegisterNode { get; set; } = null!;

        /// <summary>
        /// Verbosity. Defaults to `2`.
        /// </summary>
        [Input("v", required: true)]
        public Input<int> V { get; set; } = null!;

        public KubeletConfigurationPropsArgs()
        {
        }
        public static new KubeletConfigurationPropsArgs Empty => new KubeletConfigurationPropsArgs();
    }
}
