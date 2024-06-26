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
    public static class GetKubeProxyConfiguration
    {
        /// <summary>
        /// kube-proxy configuration.
        /// </summary>
        public static Task<GetKubeProxyConfigurationResult> InvokeAsync(GetKubeProxyConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubeProxyConfigurationResult>("kubernetes-the-hard-way:config:getKubeProxyConfiguration", args ?? new GetKubeProxyConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// kube-proxy configuration.
        /// </summary>
        public static Output<GetKubeProxyConfigurationResult> Invoke(GetKubeProxyConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubeProxyConfigurationResult>("kubernetes-the-hard-way:config:getKubeProxyConfiguration", args ?? new GetKubeProxyConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubeProxyConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster CIDR.
        /// </summary>
        [Input("clusterCIDR", required: true)]
        public string ClusterCIDR { get; set; } = null!;

        /// <summary>
        /// Path to the kubeconfig.
        /// </summary>
        [Input("kubeconfig", required: true)]
        public string Kubeconfig { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("mode")]
        public string? Mode { get; set; }

        public GetKubeProxyConfigurationArgs()
        {
        }
        public static new GetKubeProxyConfigurationArgs Empty => new GetKubeProxyConfigurationArgs();
    }

    public sealed class GetKubeProxyConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster CIDR.
        /// </summary>
        [Input("clusterCIDR", required: true)]
        public Input<string> ClusterCIDR { get; set; } = null!;

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

        public GetKubeProxyConfigurationInvokeArgs()
        {
        }
        public static new GetKubeProxyConfigurationInvokeArgs Empty => new GetKubeProxyConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubeProxyConfigurationResult
    {
        public readonly Outputs.KubeProxyConfiguration Result;

        [OutputConstructor]
        private GetKubeProxyConfigurationResult(Outputs.KubeProxyConfiguration result)
        {
            Result = result;
        }
    }
}
