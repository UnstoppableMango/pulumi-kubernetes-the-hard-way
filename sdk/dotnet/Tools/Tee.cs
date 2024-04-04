// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Tools
{
    /// <summary>
    /// Read from standard input and write to standard output and files.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:tools:Tee")]
    public partial class Tee : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// Append to the given FILEs, do not overwrite.
        /// </summary>
        [Output("append")]
        public Output<bool> Append { get; private set; } = null!;

        /// <summary>
        /// Represents the command run on the remote system.
        /// </summary>
        [Output("command")]
        public Output<Pulumi.Command.Remote.Command> Command { get; private set; } = null!;

        /// <summary>
        /// Connection details for the remote system.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// The file(s) to write to.
        /// </summary>
        [Output("files")]
        public Output<Union<string, ImmutableArray<string>>> Files { get; private set; } = null!;

        /// <summary>
        /// Ignore interrupt signals.
        /// </summary>
        [Output("ignoreInterrupts")]
        public Output<bool> IgnoreInterrupts { get; private set; } = null!;

        /// <summary>
        /// At what stage(s) in the resource lifecycle should the command be run.
        /// </summary>
        [Output("lifecycle")]
        public Output<ImmutableArray<UnMango.KubernetesTheHardWay.Tools.CommandLifecycle>> Lifecycle { get; private set; } = null!;

        /// <summary>
        /// Set behavior on write error.
        /// </summary>
        [Output("outputError")]
        public Output<UnMango.KubernetesTheHardWay.Tools.TeeMode?> OutputError { get; private set; } = null!;

        /// <summary>
        /// Operate in a more appropriate MODE with pipes.
        /// </summary>
        [Output("pipe")]
        public Output<bool> Pipe { get; private set; } = null!;

        [Output("stdin")]
        public Output<string> Stdin { get; private set; } = null!;


        /// <summary>
        /// Create a Tee resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tee(string name, TeeArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Tee", name, args ?? new TeeArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class TeeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Append to the given FILEs, do not overwrite.
        /// </summary>
        [Input("append")]
        public Input<bool>? Append { get; set; }

        /// <summary>
        /// Connection details for the remote system.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The file(s) to write to.
        /// </summary>
        [Input("files", required: true)]
        public InputUnion<string, ImmutableArray<string>> Files { get; set; } = null!;

        /// <summary>
        /// Ignore interrupt signals.
        /// </summary>
        [Input("ignoreInterrupts")]
        public Input<bool>? IgnoreInterrupts { get; set; }

        /// <summary>
        /// At what stage(s) in the resource lifecycle should the command be run.
        /// </summary>
        [Input("lifecycle")]
        public Union<UnMango.KubernetesTheHardWay.Tools.CommandLifecycle, ImmutableArray<UnMango.KubernetesTheHardWay.Tools.CommandLifecycle>>? Lifecycle { get; set; }

        /// <summary>
        /// Set behavior on write error.
        /// </summary>
        [Input("outputError")]
        public Input<UnMango.KubernetesTheHardWay.Tools.TeeMode>? OutputError { get; set; }

        /// <summary>
        /// Operate in a more appropriate MODE with pipes.
        /// </summary>
        [Input("pipe")]
        public Input<bool>? Pipe { get; set; }

        [Input("stdin", required: true)]
        public Input<string> Stdin { get; set; } = null!;

        /// <summary>
        /// Output version information and exit.
        /// </summary>
        [Input("version")]
        public Input<bool>? Version { get; set; }

        public TeeArgs()
        {
        }
        public static new TeeArgs Empty => new TeeArgs();
    }
}
