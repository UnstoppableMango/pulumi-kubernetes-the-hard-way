// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Tls
{
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:tls:ClusterPki")]
    public partial class ClusterPki : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The admin certificate.
        /// </summary>
        [Output("admin")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Certificate> Admin { get; private set; } = null!;

        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Output("algorithm")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Algorithm> Algorithm { get; private set; } = null!;

        [Output("ca")]
        public Output<UnMango.KubernetesTheHardWay.Tls.RootCa> Ca { get; private set; } = null!;

        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The controller manager certificate.
        /// </summary>
        [Output("controllerManager")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Certificate> ControllerManager { get; private set; } = null!;

        /// <summary>
        /// The kube proxy certificate.
        /// </summary>
        [Output("kubeProxy")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Certificate> KubeProxy { get; private set; } = null!;

        /// <summary>
        /// The kube scheduler certificate.
        /// </summary>
        [Output("kubeScheduler")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Certificate> KubeScheduler { get; private set; } = null!;

        /// <summary>
        /// Map of node name to kubelet certificate.
        /// </summary>
        [Output("kubelet")]
        public Output<ImmutableDictionary<string, UnMango.KubernetesTheHardWay.Tls.Certificate>> Kubelet { get; private set; } = null!;

        /// <summary>
        /// The kubernetes certificate.
        /// </summary>
        [Output("kubernetes")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Certificate> Kubernetes { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible IP for the cluster.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        /// </summary>
        [Output("rsaBits")]
        public Output<int> RsaBits { get; private set; } = null!;

        /// <summary>
        /// The service accounts certificate.
        /// </summary>
        [Output("serviceAccounts")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Certificate> ServiceAccounts { get; private set; } = null!;

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid.
        /// </summary>
        [Output("validityPeriodHours")]
        public Output<int> ValidityPeriodHours { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterPki resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterPki(string name, ClusterPkiArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tls:ClusterPki", name, args ?? new ClusterPkiArgs(), MakeResourceOptions(options, ""), remote: true)
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

        public global::Pulumi.Output<UnMango.KubernetesTheHardWay.Config.Outputs.Kubeconfig> GetKubeconfig(ClusterPkiGetKubeconfigArgs args)
            => global::Pulumi.Deployment.Instance.Call<ClusterPkiGetKubeconfigResult>("kubernetes-the-hard-way:tls:ClusterPki/getKubeconfig", args ?? new ClusterPkiGetKubeconfigArgs(), this).Apply(v => v.Result);
    }

    public sealed class ClusterPkiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<UnMango.KubernetesTheHardWay.Tls.Algorithm>? Algorithm { get; set; }

        /// <summary>
        /// A name to use for the cluster
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<UnMango.KubernetesTheHardWay.Tls.EcdsaCurve>? EcdsaCurve { get; set; }

        [Input("nodes", required: true)]
        private InputMap<Inputs.ClusterPkiNodeArgs>? _nodes;

        /// <summary>
        /// Map of node names to node configuration.
        /// </summary>
        public InputMap<Inputs.ClusterPkiNodeArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputMap<Inputs.ClusterPkiNodeArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// Publicly accessible IP address.
        /// </summary>
        [Input("publicIp", required: true)]
        public Input<string> PublicIp { get; set; } = null!;

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid.
        /// </summary>
        [Input("validityPeriodHours")]
        public Input<int>? ValidityPeriodHours { get; set; }

        public ClusterPkiArgs()
        {
            Algorithm = UnMango.KubernetesTheHardWay.Tls.Algorithm.RSA;
            RsaBits = 2048;
            ValidityPeriodHours = 8076;
        }
        public static new ClusterPkiArgs Empty => new ClusterPkiArgs();
    }

    /// <summary>
    /// The set of arguments for the <see cref="ClusterPki.GetKubeconfig"/> method.
    /// </summary>
    public sealed class ClusterPkiGetKubeconfigArgs : global::Pulumi.CallArgs
    {
        [Input("options", required: true)]
        public object Options { get; set; } = null!;

        public ClusterPkiGetKubeconfigArgs()
        {
        }
        public static new ClusterPkiGetKubeconfigArgs Empty => new ClusterPkiGetKubeconfigArgs();
    }

    /// <summary>
    /// The results of the <see cref="ClusterPki.GetKubeconfig"/> method.
    /// </summary>
    [OutputType]
    internal sealed class ClusterPkiGetKubeconfigResult
    {
        public readonly UnMango.KubernetesTheHardWay.Config.Outputs.Kubeconfig Result;

        [OutputConstructor]
        private ClusterPkiGetKubeconfigResult(UnMango.KubernetesTheHardWay.Config.Outputs.Kubeconfig result)
        {
            Result = result;
        }
    }
}
