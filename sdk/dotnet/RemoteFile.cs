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
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:index:RemoteFile")]
    public partial class RemoteFile : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The command resource.
        /// </summary>
        [Output("command")]
        public Output<Pulumi.Command.Remote.Command> Command { get; private set; } = null!;


        /// <summary>
        /// Create a RemoteFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RemoteFile(string name, RemoteFileArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:index:RemoteFile", name, args ?? new RemoteFileArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class RemoteFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The content of the file.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The path to the file on the remote host.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public RemoteFileArgs()
        {
        }
        public static new RemoteFileArgs Empty => new RemoteFileArgs();
    }
}
