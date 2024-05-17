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
    /// containerd cri plugin configuration.
    /// </summary>
    public sealed class ContainerdCriPluginConfigurationCniArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// bin_dir
        /// </summary>
        [Input("binDir")]
        public Input<string>? BinDir { get; set; }

        /// <summary>
        /// conf_dir
        /// </summary>
        [Input("confDir")]
        public Input<string>? ConfDir { get; set; }

        public ContainerdCriPluginConfigurationCniArgs()
        {
        }
        public static new ContainerdCriPluginConfigurationCniArgs Empty => new ContainerdCriPluginConfigurationCniArgs();
    }
}