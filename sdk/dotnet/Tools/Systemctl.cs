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
    public partial class Systemctl : global::Pulumi.CustomResource
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
        public Output<Outputs.SystemctlOpts?> Create { get; private set; } = null!;

        /// <summary>
        /// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
        /// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
        /// Command resource from previous create or update steps.
        /// </summary>
        [Output("delete")]
        public Output<Outputs.SystemctlOpts?> Delete { get; private set; } = null!;

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
        /// TODO
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
        public Output<Outputs.SystemctlOpts?> Update { get; private set; } = null!;


        /// <summary>
        /// Create a Systemctl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Systemctl(string name, SystemctlArgs args, CustomResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Systemctl", name, args ?? new SystemctlArgs(), MakeResourceOptions(options, ""))
        {
        }
        internal Systemctl(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Systemctl", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private Systemctl(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Systemctl", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/UnstoppableMango",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Systemctl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Systemctl Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Systemctl(name, id, options);
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
        /// Connection details for the remote system
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The command to run on create.
        /// </summary>
        [Input("create")]
        public Inputs.SystemctlOptsArgs? Create { get; set; }

        /// <summary>
        /// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
        /// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
        /// Command resource from previous create or update steps.
        /// </summary>
        [Input("delete")]
        public Inputs.SystemctlOptsArgs? Delete { get; set; }

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
        /// TODO
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
        public Inputs.SystemctlOptsArgs? Update { get; set; }

        public SystemctlArgs()
        {
        }
        public static new SystemctlArgs Empty => new SystemctlArgs();
    }
}
