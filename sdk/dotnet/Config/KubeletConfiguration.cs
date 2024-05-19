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
    /// Get the kubelet configuration.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:config:KubeletConfiguration")]
    public partial class KubeletConfiguration : global::Pulumi.ComponentResource
    {
        [Output("result")]
        public Output<Outputs.KubeletConfiguration> Result { get; private set; } = null!;

        /// <summary>
        /// The yaml representation of the manifest.
        /// </summary>
        [Output("yaml")]
        public Output<string> Yaml { get; private set; } = null!;


        /// <summary>
        /// Create a KubeletConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubeletConfiguration(string name, KubeletConfigurationArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:config:KubeletConfiguration", name, args ?? new KubeletConfigurationArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class KubeletConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("anonymous")]
        public Input<bool>? Anonymous { get; set; }

        [Input("authorizationMode")]
        public Input<string>? AuthorizationMode { get; set; }

        [Input("cgroupDriver")]
        public Input<string>? CgroupDriver { get; set; }

        [Input("clientCAFile")]
        public Input<string>? ClientCAFile { get; set; }

        [Input("clusterDNS")]
        private InputList<string>? _clusterDNS;
        public InputList<string> ClusterDNS
        {
            get => _clusterDNS ?? (_clusterDNS = new InputList<string>());
            set => _clusterDNS = value;
        }

        [Input("clusterDomain")]
        public Input<string>? ClusterDomain { get; set; }

        [Input("containerRuntimeEndpoint")]
        public Input<string>? ContainerRuntimeEndpoint { get; set; }

        [Input("podCIDR", required: true)]
        public Input<string> PodCIDR { get; set; } = null!;

        [Input("resolvConf")]
        public Input<string>? ResolvConf { get; set; }

        [Input("runtimeRequestTimeout")]
        public Input<string>? RuntimeRequestTimeout { get; set; }

        [Input("tlsCertFile")]
        public Input<string>? TlsCertFile { get; set; }

        [Input("tlsPrivateKeyFile")]
        public Input<string>? TlsPrivateKeyFile { get; set; }

        [Input("webhook")]
        public Input<bool>? Webhook { get; set; }

        public KubeletConfigurationArgs()
        {
        }
        public static new KubeletConfigurationArgs Empty => new KubeletConfigurationArgs();
    }
}
