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
    /// Installs etcd on a remote system
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:EtcdInstall")]
    public partial class EtcdInstall : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The CPU architecture to install.
        /// </summary>
        [Output("architecture")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Architecture> Architecture { get; private set; } = null!;

        /// <summary>
        /// The name of the downloaded archive.
        /// </summary>
        [Output("archiveName")]
        public Output<string> ArchiveName { get; private set; } = null!;

        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// The directory to install the binary to.
        /// </summary>
        [Output("directory")]
        public Output<string> Directory { get; private set; } = null!;

        /// <summary>
        /// The download operation.
        /// </summary>
        [Output("download")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Outputs.Download> Download { get; private set; } = null!;

        /// <summary>
        /// The etcd mv operation.
        /// </summary>
        [Output("etcdMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> EtcdMv { get; private set; } = null!;

        /// <summary>
        /// The etcd path on the remote system
        /// </summary>
        [Output("etcdPath")]
        public Output<string> EtcdPath { get; private set; } = null!;

        /// <summary>
        /// The etcdctl mv operation.
        /// </summary>
        [Output("etcdctlMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> EtcdctlMv { get; private set; } = null!;

        /// <summary>
        /// The etcdctl path on the remote system
        /// </summary>
        [Output("etcdctlPath")]
        public Output<string> EtcdctlPath { get; private set; } = null!;

        /// <summary>
        /// The mkdir operation.
        /// </summary>
        [Output("mkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Outputs.Mkdir> Mkdir { get; private set; } = null!;

        /// <summary>
        /// The mktemp operation.
        /// </summary>
        [Output("mktemp")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Outputs.Mktemp> Mktemp { get; private set; } = null!;

        /// <summary>
        /// The path to the installed binary.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The rm operation.
        /// </summary>
        [Output("rm")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Rm> Rm { get; private set; } = null!;

        /// <summary>
        /// The tar operation.
        /// </summary>
        [Output("tar")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Tar> Tar { get; private set; } = null!;

        /// <summary>
        /// The url used to download the binary.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The version to install.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a EtcdInstall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EtcdInstall(string name, EtcdInstallArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:EtcdInstall", name, args ?? new EtcdInstallArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class EtcdInstallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU architecture to install.
        /// </summary>
        [Input("architecture")]
        public Input<UnMango.KubernetesTheHardWay.Remote.Architecture>? Architecture { get; set; }

        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The directory to install the binary to.
        /// </summary>
        [Input("directory")]
        public Input<string>? Directory { get; set; }

        /// <summary>
        /// The version to install.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public EtcdInstallArgs()
        {
            Directory = "/usr/local/bin";
        }
        public static new EtcdInstallArgs Empty => new EtcdInstallArgs();
    }
}
