// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay
{
    /// <summary>
    /// Kubeconfig
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:index:Kubeconfig")]
    public partial class Kubeconfig : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// Create a Kubeconfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Kubeconfig(string name, KubeconfigArgs? args = null, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:index:Kubeconfig", name, args ?? new KubeconfigArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class KubeconfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("options")]
        public object? Options { get; set; }

        /// <summary>
        /// The PKI containing certificate data.
        /// </summary>
        [Input("pki")]
        public Input<UnMango.KubernetesTheHardWay.Tls.ClusterPki>? Pki { get; set; }

        public KubeconfigArgs()
        {
        }
        public static new KubeconfigArgs Empty => new KubeconfigArgs();
    }
}