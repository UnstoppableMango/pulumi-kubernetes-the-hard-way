// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config.Outputs
{

    /// <summary>
    /// containerd cri plugin configuration.
    /// </summary>
    [OutputType]
    public sealed class ContainerdCriPluginConfigurationContainerd
    {
        /// <summary>
        /// default_runtime_name
        /// </summary>
        public readonly string? DefaultRuntimeName;
        /// <summary>
        /// The containerd runtime configuration.
        /// </summary>
        public readonly Outputs.ContainerdCriPluginConfigurationContainerdRunc? Runtimes;
        /// <summary>
        /// snapshotter
        /// </summary>
        public readonly string? Snapshotter;

        [OutputConstructor]
        private ContainerdCriPluginConfigurationContainerd(
            string? defaultRuntimeName,

            Outputs.ContainerdCriPluginConfigurationContainerdRunc? runtimes,

            string? snapshotter)
        {
            DefaultRuntimeName = defaultRuntimeName;
            Runtimes = runtimes;
            Snapshotter = snapshotter;
        }
    }
}