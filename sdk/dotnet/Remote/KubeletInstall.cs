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
    /// Installs kubelet on a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:KubeletInstall")]
    public partial class KubeletInstall : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The CPU architecture.
        /// </summary>
        [Output("architecture")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Architecture> Architecture { get; private set; } = null!;

        [Output("binName")]
        public Output<string?> BinName { get; private set; } = null!;

        /// <summary>
        /// The connection details.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// Directory to install the binary.
        /// </summary>
        [Output("directory")]
        public Output<string> Directory { get; private set; } = null!;

        [Output("download")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Download> Download { get; private set; } = null!;

        [Output("mkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir?> Mkdir { get; private set; } = null!;

        [Output("mktemp")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mktemp> Mktemp { get; private set; } = null!;

        [Output("mv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> Mv { get; private set; } = null!;

        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        [Output("rm")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Rm> Rm { get; private set; } = null!;

        /// <summary>
        /// The version to install.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a KubeletInstall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubeletInstall(string name, KubeletInstallArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:KubeletInstall", name, args ?? new KubeletInstallArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class KubeletInstallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU architecture.
        /// </summary>
        [Input("architecture")]
        public Input<UnMango.KubernetesTheHardWay.Remote.Architecture>? Architecture { get; set; }

        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// Directory to install the binary.
        /// </summary>
        [Input("directory")]
        public Input<string>? Directory { get; set; }

        /// <summary>
        /// The version of to install.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubeletInstallArgs()
        {
            Directory = "/usr/local/bin";
        }
        public static new KubeletInstallArgs Empty => new KubeletInstallArgs();
    }
}
