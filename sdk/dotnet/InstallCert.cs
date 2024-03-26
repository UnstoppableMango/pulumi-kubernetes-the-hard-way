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
    public static class InstallCert
    {
        /// <summary>
        /// Creates a RemoteFile resource representing the copy operation.
        /// </summary>
        public static Task<InstallCertResult> InvokeAsync(InstallCertArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<InstallCertResult>("kubernetes-the-hard-way:index:installCert", args ?? new InstallCertArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a RemoteFile resource representing the copy operation.
        /// </summary>
        public static Output<InstallCertResult> Invoke(InstallCertInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<InstallCertResult>("kubernetes-the-hard-way:index:installCert", args ?? new InstallCertInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstallCertArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The certificate to install at the remote location.
        /// </summary>
        [Input("cert", required: true)]
        public UnMango.KubernetesTheHardWay.Certificate Cert { get; set; } = null!;

        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Inputs.Connection Connection { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("options")]
        public Inputs.ResourceOptions? Options { get; set; }

        /// <summary>
        /// The path to install to.
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        public InstallCertArgs()
        {
        }
        public static new InstallCertArgs Empty => new InstallCertArgs();
    }

    public sealed class InstallCertInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The certificate to install at the remote location.
        /// </summary>
        [Input("cert", required: true)]
        public Input<UnMango.KubernetesTheHardWay.Certificate> Cert { get; set; } = null!;

        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("options")]
        public Inputs.ResourceOptionsArgs? Options { get; set; }

        /// <summary>
        /// The path to install to.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public InstallCertInvokeArgs()
        {
        }
        public static new InstallCertInvokeArgs Empty => new InstallCertInvokeArgs();
    }


    [OutputType]
    public sealed class InstallCertResult
    {
        /// <summary>
        /// A resource representing the the file on the remote machine.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.RemoteFile Result;

        [OutputConstructor]
        private InstallCertResult(UnMango.KubernetesTheHardWay.RemoteFile result)
        {
            Result = result;
        }
    }
}