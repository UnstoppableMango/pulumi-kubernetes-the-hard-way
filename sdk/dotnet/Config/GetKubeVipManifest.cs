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
    public static class GetKubeVipManifest
    {
        /// <summary>
        /// Gets the static pod manifests for KubeVip.
        /// </summary>
        public static Task<GetKubeVipManifestResult> InvokeAsync(GetKubeVipManifestArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubeVipManifestResult>("kubernetes-the-hard-way:config:getKubeVipManifest", args ?? new GetKubeVipManifestArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the static pod manifests for KubeVip.
        /// </summary>
        public static Output<GetKubeVipManifestResult> Invoke(GetKubeVipManifestInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubeVipManifestResult>("kubernetes-the-hard-way:config:getKubeVipManifest", args ?? new GetKubeVipManifestInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubeVipManifestArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// TODO
        /// </summary>
        [Input("address", required: true)]
        public string Address { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("cpEnable")]
        public bool? CpEnable { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("cpNamespace")]
        public string? CpNamespace { get; set; }

        /// <summary>
        /// Override the kube-vip image.
        /// </summary>
        [Input("image")]
        public string? Image { get; set; }

        /// <summary>
        /// Path to the kubeconfig on the remote host.
        /// </summary>
        [Input("kubeconfigPath", required: true)]
        public string KubeconfigPath { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("port")]
        public int? Port { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("svcEnable")]
        public bool? SvcEnable { get; set; }

        /// <summary>
        /// Version of kube-vip to use.
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipArp")]
        public bool? VipArp { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipCidr", required: true)]
        public string VipCidr { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipDdns")]
        public bool? VipDdns { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipInterface")]
        public string? VipInterface { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipLeaderElection")]
        public bool? VipLeaderElection { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipLeaseDuration")]
        public int? VipLeaseDuration { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipRenewDeadline")]
        public int? VipRenewDeadline { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipRetryPeriod")]
        public int? VipRetryPeriod { get; set; }

        public GetKubeVipManifestArgs()
        {
            Port = 6443;
        }
        public static new GetKubeVipManifestArgs Empty => new GetKubeVipManifestArgs();
    }

    public sealed class GetKubeVipManifestInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// TODO
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("cpEnable")]
        public Input<bool>? CpEnable { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("cpNamespace")]
        public Input<string>? CpNamespace { get; set; }

        /// <summary>
        /// Override the kube-vip image.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// Path to the kubeconfig on the remote host.
        /// </summary>
        [Input("kubeconfigPath", required: true)]
        public Input<string> KubeconfigPath { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("svcEnable")]
        public Input<bool>? SvcEnable { get; set; }

        /// <summary>
        /// Version of kube-vip to use.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipArp")]
        public Input<bool>? VipArp { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipCidr", required: true)]
        public Input<string> VipCidr { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipDdns")]
        public Input<bool>? VipDdns { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipInterface")]
        public Input<string>? VipInterface { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipLeaderElection")]
        public Input<bool>? VipLeaderElection { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipLeaseDuration")]
        public Input<int>? VipLeaseDuration { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipRenewDeadline")]
        public Input<int>? VipRenewDeadline { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("vipRetryPeriod")]
        public Input<int>? VipRetryPeriod { get; set; }

        public GetKubeVipManifestInvokeArgs()
        {
            Port = 6443;
        }
        public static new GetKubeVipManifestInvokeArgs Empty => new GetKubeVipManifestInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubeVipManifestResult
    {
        public readonly Outputs.PodManifest Result;

        [OutputConstructor]
        private GetKubeVipManifestResult(Outputs.PodManifest result)
        {
            Result = result;
        }
    }
}
