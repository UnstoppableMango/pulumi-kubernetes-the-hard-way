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
    /// Downloads the file specified by `url` onto a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:Download")]
    public partial class Download : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// The fully qualified path on the remote system where the file should be downloaded to.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// The mkdir operation.
        /// </summary>
        [Output("mkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir> Mkdir { get; private set; } = null!;

        /// <summary>
        /// Remove the downloaded fiel when the resource is deleted.
        /// </summary>
        [Output("removeOnDelete")]
        public Output<bool> RemoveOnDelete { get; private set; } = null!;

        /// <summary>
        /// The URL of the file to be downloaded.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The wget operation.
        /// </summary>
        [Output("wget")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Wget> Wget { get; private set; } = null!;


        /// <summary>
        /// Create a Download resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Download(string name, DownloadArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:Download", name, args ?? new DownloadArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class DownloadArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The fully qualified path on the remote system where the file should be downloaded to.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// Remove the downloaded fiel when the resource is deleted.
        /// </summary>
        [Input("removeOnDelete")]
        public Input<bool>? RemoveOnDelete { get; set; }

        /// <summary>
        /// The URL of the file to be downloaded.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public DownloadArgs()
        {
        }
        public static new DownloadArgs Empty => new DownloadArgs();
    }
}
