// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config
{
    /// <summary>
    /// kube-proxy configuration.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:config:KubeProxyConfiguration")]
    public partial class KubeProxyConfiguration : global::Pulumi.ComponentResource
    {
        [Output("result")]
        public Output<Outputs.KubeProxyConfiguration> Result { get; private set; } = null!;

        /// <summary>
        /// The yaml representation of the manifest.
        /// </summary>
        [Output("yaml")]
        public Output<string> Yaml { get; private set; } = null!;


        /// <summary>
        /// Create a KubeProxyConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubeProxyConfiguration(string name, KubeProxyConfigurationArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:config:KubeProxyConfiguration", name, args ?? new KubeProxyConfigurationArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/UnstoppableMango",
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class KubeProxyConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster CIDR.
        /// </summary>
        [Input("clusterCIDR")]
        public Input<string>? ClusterCIDR { get; set; }

        /// <summary>
        /// Path to the kubeconfig.
        /// </summary>
        [Input("kubeconfig", required: true)]
        public Input<string> Kubeconfig { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public KubeProxyConfigurationArgs()
        {
        }
        public static new KubeProxyConfigurationArgs Empty => new KubeProxyConfigurationArgs();
    }
}
