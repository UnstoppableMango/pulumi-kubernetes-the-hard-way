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
    /// Abstraction over the `chmod` utility on a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:tools:Chmod")]
    public partial class Chmod : global::Pulumi.ComponentResource
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
        /// The command to run on create.
        /// </summary>
        [Output("create")]
        public Output<Union<string, Outputs.ChmodOpts>?> Create { get; private set; } = null!;

        /// <summary>
        /// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
        /// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
        /// Command resource from previous create or update steps.
        /// </summary>
        [Output("delete")]
        public Output<Union<string, Outputs.ChmodOpts>?> Delete { get; private set; } = null!;

        /// <summary>
        /// Environment variables
        /// </summary>
        [Output("environment")]
        public Output<ImmutableDictionary<string, string>> Environment { get; private set; } = null!;

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
        /// Trigger replacements on changes to this input.
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableArray<object>> Triggers { get; private set; } = null!;

        /// <summary>
        /// The command to run on update, if empty, create will 
        /// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR 
        /// are set to the stdout and stderr properties of the Command resource from previous 
        /// create or update steps.
        /// </summary>
        [Output("update")]
        public Output<Union<string, Outputs.ChmodOpts>?> Update { get; private set; } = null!;


        /// <summary>
        /// Create a Chmod resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Chmod(string name, ChmodArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Chmod", name, args ?? new ChmodArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class ChmodArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        /// </summary>
        [Input("binaryPath")]
        public Input<string>? BinaryPath { get; set; }

        /// <summary>
        /// Connection details for the remote system
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The command to run on create.
        /// </summary>
        [Input("create")]
        public Union<Input<string>, Inputs.ChmodOptsArgs>? Create { get; set; }

        /// <summary>
        /// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
        /// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
        /// Command resource from previous create or update steps.
        /// </summary>
        [Input("delete")]
        public Union<Input<string>, Inputs.ChmodOptsArgs>? Delete { get; set; }

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
        /// TODO
        /// </summary>
        [Input("stdin")]
        public Input<string>? Stdin { get; set; }

        [Input("triggers")]
        private InputList<object>? _triggers;

        /// <summary>
        /// Trigger replacements on changes to this input.
        /// </summary>
        public InputList<object> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<object>());
            set => _triggers = value;
        }

        /// <summary>
        /// The command to run on update, if empty, create will 
        /// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR 
        /// are set to the stdout and stderr properties of the Command resource from previous 
        /// create or update steps.
        /// </summary>
        [Input("update")]
        public Union<Input<string>, Inputs.ChmodOptsArgs>? Update { get; set; }

        public ChmodArgs()
        {
        }
        public static new ChmodArgs Empty => new ChmodArgs();
    }
}
