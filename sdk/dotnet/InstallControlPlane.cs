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
    public static class InstallControlPlane
    {
        /// <summary>
        /// Install PKI onto a controlplane node.
        /// </summary>
        public static Task InvokeAsync(InstallControlPlaneArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync("kubernetes-the-hard-way:index:installControlPlane", args ?? new InstallControlPlaneArgs(), options.WithDefaults());
    }


    public sealed class InstallControlPlaneArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Inputs.Connection Connection { get; set; } = null!;

        [Input("opts")]
        public Inputs.ResourceOptions? Opts { get; set; }

        /// <summary>
        /// The PKI to install
        /// </summary>
        [Input("pki", required: true)]
        public UnMango.KubernetesTheHardWay.ClusterPki Pki { get; set; } = null!;

        public InstallControlPlaneArgs()
        {
        }
        public static new InstallControlPlaneArgs Empty => new InstallControlPlaneArgs();
    }
}
