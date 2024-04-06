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
    /// Installs kube-scheduler on a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:KubeSchedulerInstall")]
    public partial class KubeSchedulerInstall : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The kube-scheduler CPU architecture.
        /// </summary>
        [Output("architecture")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Architecture> Architecture { get; private set; } = null!;

        [Output("archiveName")]
        public Output<string?> ArchiveName { get; private set; } = null!;

        /// <summary>
        /// The connection details.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// Directory to install the `etcd` and `etcdctl` binaries.
        /// </summary>
        [Output("directory")]
        public Output<string> Directory { get; private set; } = null!;

        [Output("mkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir?> Mkdir { get; private set; } = null!;

        /// <summary>
        /// The version of kube-scheduler to install.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a KubeSchedulerInstall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubeSchedulerInstall(string name, KubeSchedulerInstallArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:KubeSchedulerInstall", name, args ?? new KubeSchedulerInstallArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class KubeSchedulerInstallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The kube-scheduler CPU architecture.
        /// </summary>
        [Input("architecture")]
        public Input<UnMango.KubernetesTheHardWay.Remote.Architecture>? Architecture { get; set; }

        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// Directory to install the `kube-scheduler` binary.
        /// </summary>
        [Input("directory")]
        public Input<string>? Directory { get; set; }

        /// <summary>
        /// The version of kube-scheduler to install.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubeSchedulerInstallArgs()
        {
            Directory = "/usr/local/bin";
        }
        public static new KubeSchedulerInstallArgs Empty => new KubeSchedulerInstallArgs();
    }
}
