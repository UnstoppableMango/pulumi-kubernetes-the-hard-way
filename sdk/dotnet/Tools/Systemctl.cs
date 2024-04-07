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
    /// Abstraction over the `systemctl` utility on a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:tools:Systemctl")]
    public partial class Systemctl : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        /// </summary>
        [Output("binaryPath")]
        public Output<string> BinaryPath { get; private set; } = null!;

        /// <summary>
        /// The underlying command
        /// </summary>
        [Output("command")]
        public Output<Pulumi.Command.Remote.Command> Command { get; private set; } = null!;

        /// <summary>
        /// Connection details for the remote system
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// Environment variables
        /// </summary>
        [Output("environment")]
        public Output<ImmutableDictionary<string, string>> Environment { get; private set; } = null!;

        /// <summary>
        /// At what stage(s) in the resource lifecycle should the command be run
        /// </summary>
        [Output("lifecycle")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Outputs.CommandLifecycle?> Lifecycle { get; private set; } = null!;

        /// <summary>
        /// Corresponds to the [PATTERN] argument
        /// </summary>
        [Output("pattern")]
        public Output<string?> Pattern { get; private set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Output("stderr")]
        public Output<string> Stderr { get; private set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Output("stdin")]
        public Output<string?> Stdin { get; private set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Output("stdout")]
        public Output<string> Stdout { get; private set; } = null!;

        /// <summary>
        /// Corresponds to the COMMAND argument.
        /// </summary>
        [Output("systemctlCommand")]
        public Output<UnMango.KubernetesTheHardWay.Tools.SystemctlCommand> SystemctlCommand { get; private set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableArray<object>> Triggers { get; private set; } = null!;

        /// <summary>
        /// Corresponds to the [UNIT...] argument.
        /// </summary>
        [Output("unit")]
        public Output<Union<string, ImmutableArray<string>>> Unit { get; private set; } = null!;


        /// <summary>
        /// Create a Systemctl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Systemctl(string name, SystemctlArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Systemctl", name, args ?? new SystemctlArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class SystemctlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        /// </summary>
        [Input("binaryPath")]
        public Input<string>? BinaryPath { get; set; }

        /// <summary>
        /// Corresponds to the COMMAND argument.
        /// </summary>
        [Input("command", required: true)]
        public UnMango.KubernetesTheHardWay.Tools.SystemctlCommand Command { get; set; }

        /// <summary>
        /// Connection details for the remote system
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("environment")]
        private InputMap<string>? _environment;

        /// <summary>
        /// Environment variables
        /// </summary>
        public InputMap<string> Environment
        {
            get => _environment ?? (_environment = new InputMap<string>());
            set => _environment = value;
        }

        /// <summary>
        /// At what stage(s) in the resource lifecycle should the command be run
        /// </summary>
        [Input("lifecycle")]
        public UnMango.KubernetesTheHardWay.Tools.Inputs.CommandLifecycle? Lifecycle { get; set; }

        /// <summary>
        /// Corresponds to the [PATTERN] argument
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("stdin")]
        public Input<string>? Stdin { get; set; }

        [Input("triggers")]
        private InputList<object>? _triggers;

        /// <summary>
        /// TODO
        /// </summary>
        public InputList<object> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<object>());
            set => _triggers = value;
        }

        /// <summary>
        /// Corresponds to the [UNIT...] argument.
        /// </summary>
        [Input("unit", required: true)]
        public InputUnion<string, ImmutableArray<string>> Unit { get; set; } = null!;

        public SystemctlArgs()
        {
        }
        public static new SystemctlArgs Empty => new SystemctlArgs();
    }
}
