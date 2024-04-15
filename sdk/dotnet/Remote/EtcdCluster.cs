// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Remote
{
    /// <summary>
    /// Creates an etcd cluster from one or more remote systems.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:EtcdCluster")]
    public partial class EtcdCluster : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// TODO
        /// </summary>
        [Output("architecture")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Architecture?> Architecture { get; private set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Output("binaryDirectory")]
        public Output<string?> BinaryDirectory { get; private set; } = null!;

        /// <summary>
        /// The TLS bundle.
        /// </summary>
        [Output("bundle")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Outputs.Bundle> Bundle { get; private set; } = null!;

        /// <summary>
        /// Map of node name to etcd configuration.
        /// </summary>
        [Output("configuration")]
        public Output<ImmutableDictionary<string, UnMango.KubernetesTheHardWay.Remote.EtcdConfiguration>> Configuration { get; private set; } = null!;

        /// <summary>
        /// The directory to use for etcd configuration.
        /// </summary>
        [Output("configurationDirectory")]
        public Output<string?> ConfigurationDirectory { get; private set; } = null!;

        /// <summary>
        /// The directory to use for etcd data.
        /// </summary>
        [Output("dataDirectory")]
        public Output<string?> DataDirectory { get; private set; } = null!;

        /// <summary>
        /// Map of node name to etcd install.
        /// </summary>
        [Output("install")]
        public Output<ImmutableDictionary<string, UnMango.KubernetesTheHardWay.Remote.EtcdInstall>> Install { get; private set; } = null!;

        /// <summary>
        /// Etcd node configuration. The key should be a name used to identify the node.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableDictionary<string, Outputs.EtcdNode>> Nodes { get; private set; } = null!;

        /// <summary>
        /// Map of node name to etcd systemd service.
        /// </summary>
        [Output("service")]
        public Output<ImmutableDictionary<string, UnMango.KubernetesTheHardWay.Remote.EtcdService>> Service { get; private set; } = null!;

        /// <summary>
        /// Map of node name to etcd start commands.
        /// </summary>
        [Output("start")]
        public Output<ImmutableDictionary<string, UnMango.KubernetesTheHardWay.Remote.StartEtcd>> Start { get; private set; } = null!;

        /// <summary>
        /// The version to install.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a EtcdCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EtcdCluster(string name, EtcdClusterArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:EtcdCluster", name, args ?? new EtcdClusterArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class EtcdClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// TODO
        /// </summary>
        [Input("architecture")]
        public Input<UnMango.KubernetesTheHardWay.Remote.Architecture>? Architecture { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("binaryDirectory")]
        public Input<string>? BinaryDirectory { get; set; }

        /// <summary>
        /// The TLS bundle.
        /// </summary>
        [Input("bundle", required: true)]
        public Input<UnMango.KubernetesTheHardWay.Tls.Inputs.BundleArgs> Bundle { get; set; } = null!;

        /// <summary>
        /// The directory to use for etcd configuration.
        /// </summary>
        [Input("configurationDirectory")]
        public Input<string>? ConfigurationDirectory { get; set; }

        /// <summary>
        /// The directory to use for etcd data.
        /// </summary>
        [Input("dataDirectory")]
        public Input<string>? DataDirectory { get; set; }

        [Input("nodes", required: true)]
        private Dictionary<string, Inputs.EtcdNodeArgs>? _nodes;

        /// <summary>
        /// Etcd node configuration. The key should be a name used to identify the node.
        /// </summary>
        public Dictionary<string, Inputs.EtcdNodeArgs> Nodes
        {
            get => _nodes ?? (_nodes = new Dictionary<string, Inputs.EtcdNodeArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// The version to install.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public EtcdClusterArgs()
        {
        }
        public static new EtcdClusterArgs Empty => new EtcdClusterArgs();
    }
}
