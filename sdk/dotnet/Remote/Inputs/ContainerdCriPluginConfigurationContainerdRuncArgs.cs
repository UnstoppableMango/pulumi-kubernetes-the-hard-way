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
    /// containerd cri runc plugin configuration.
    /// </summary>
    public sealed class ContainerdCriPluginConfigurationContainerdRuncArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// runc options.
        /// </summary>
        [Input("options", required: true)]
        public Inputs.ContainerdCriPluginConfigurationContainerdRuncOptionsArgs Options { get; set; } = null!;

        /// <summary>
        /// runtime_type
        /// </summary>
        [Input("runtimeType")]
        public Input<string>? RuntimeType { get; set; }

        public ContainerdCriPluginConfigurationContainerdRuncArgs()
        {
        }
        public static new ContainerdCriPluginConfigurationContainerdRuncArgs Empty => new ContainerdCriPluginConfigurationContainerdRuncArgs();
    }
}
