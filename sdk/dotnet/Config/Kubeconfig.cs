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
    /// https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:config:Kubeconfig")]
    public partial class Kubeconfig : global::Pulumi.ComponentResource
    {
        [Output("result")]
        public Output<Outputs.Kubeconfig> Result { get; private set; } = null!;

        /// <summary>
        /// The yaml representation of the manifest.
        /// </summary>
        [Output("yaml")]
        public Output<string> Yaml { get; private set; } = null!;


        /// <summary>
        /// Create a Kubeconfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Kubeconfig(string name, KubeconfigArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:config:Kubeconfig", name, args ?? new KubeconfigArgs(), MakeResourceOptions(options, ""), remote: true)
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
        /// <summary>
        /// Certificate authority data.
        /// </summary>
        [Input("caPem", required: true)]
        public Input<string> CaPem { get; set; } = null!;

        /// <summary>
        /// The PEM encoded certificate data of the client.
        /// </summary>
        [Input("clientCert", required: true)]
        public Input<string> ClientCert { get; set; } = null!;

        /// <summary>
        /// The PEM encoded private key data of the client.
        /// </summary>
        [Input("clientKey", required: true)]
        public Input<string> ClientKey { get; set; } = null!;

        /// <summary>
        /// A name to identify the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// A name to use for the kubeconfig context
        /// </summary>
        [Input("contextName")]
        public Input<string>? ContextName { get; set; }

        /// <summary>
        /// The address and port of the Kubernetes API server.
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// The username of the user
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public KubeconfigArgs()
        {
        }
        public static new KubeconfigArgs Empty => new KubeconfigArgs();
    }
}
