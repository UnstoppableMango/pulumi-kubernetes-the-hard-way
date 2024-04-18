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
    /// Abstraction over the `sed` utility on a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:tools:Sed")]
    public partial class Sed : global::Pulumi.ComponentResource
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
        /// annotate program execution.
        /// </summary>
        [Output("debug")]
        public Output<bool> Debug { get; private set; } = null!;

        /// <summary>
        /// Environment variables
        /// </summary>
        [Output("environment")]
        public Output<ImmutableDictionary<string, string>> Environment { get; private set; } = null!;

        /// <summary>
        /// add the script to the commands to be executed.
        /// </summary>
        [Output("expressions")]
        public Output<Union<string, ImmutableArray<string>>> Expressions { get; private set; } = null!;

        /// <summary>
        /// add the contents of script-file to the commands to be executed.
        /// </summary>
        [Output("files")]
        public Output<Union<string, ImmutableArray<string>>> Files { get; private set; } = null!;

        /// <summary>
        /// follow symlinks when processing in place
        /// </summary>
        [Output("followSymlinks")]
        public Output<bool> FollowSymlinks { get; private set; } = null!;

        /// <summary>
        /// display this help and exit.
        /// </summary>
        [Output("help")]
        public Output<bool> Help { get; private set; } = null!;

        /// <summary>
        /// edit files in place (makes backup if SUFFIX supplied)
        /// </summary>
        [Output("inPlace")]
        public Output<string?> InPlace { get; private set; } = null!;

        /// <summary>
        /// corresponds to the [input-file]... argument(s).
        /// </summary>
        [Output("inputFiles")]
        public Output<Union<string, ImmutableArray<string>>> InputFiles { get; private set; } = null!;

        /// <summary>
        /// At what stage(s) in the resource lifecycle should the command be run
        /// </summary>
        [Output("lifecycle")]
        public Output<UnMango.KubernetesTheHardWay.Tools.CommandLifecycle?> Lifecycle { get; private set; } = null!;

        /// <summary>
        /// specify the desired line-wrap length for the `l' command
        /// </summary>
        [Output("lineLength")]
        public Output<int?> LineLength { get; private set; } = null!;

        /// <summary>
        /// separate lines by NUL characters
        /// </summary>
        [Output("nullData")]
        public Output<bool> NullData { get; private set; } = null!;

        /// <summary>
        /// disable all GNU extensions.
        /// </summary>
        [Output("posix")]
        public Output<bool> Posix { get; private set; } = null!;

        /// <summary>
        /// suppress automatic printing of pattern space. Same as `silent`.
        /// </summary>
        [Output("quiet")]
        public Output<bool> Quiet { get; private set; } = null!;

        /// <summary>
        /// use extended regular expressions in the script (for portability use POSIX -E).
        /// </summary>
        [Output("regexpExtended")]
        public Output<bool> RegexpExtended { get; private set; } = null!;

        /// <summary>
        /// operate in sandbox mode (disable e/r/w commands).
        /// </summary>
        [Output("sandbox")]
        public Output<bool> Sandbox { get; private set; } = null!;

        /// <summary>
        /// script only if no other script.
        /// </summary>
        [Output("script")]
        public Output<string?> Script { get; private set; } = null!;

        /// <summary>
        /// consider files as separate rather than as a single, continuous long stream.
        /// </summary>
        [Output("separate")]
        public Output<bool> Separate { get; private set; } = null!;

        /// <summary>
        /// suppress automatic printing of pattern space. Same as `quiet`.
        /// </summary>
        [Output("silent")]
        public Output<bool> Silent { get; private set; } = null!;

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
        /// load minimal amounts of data from the input files and flush the output buffers more often.
        /// </summary>
        [Output("unbuffered")]
        public Output<bool> Unbuffered { get; private set; } = null!;

        /// <summary>
        /// output version information and exit.
        /// </summary>
        [Output("version")]
        public Output<bool> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Sed resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sed(string name, SedArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Sed", name, args ?? new SedArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class SedArgs : global::Pulumi.ResourceArgs
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
        /// annotate program execution.
        /// </summary>
        [Input("debug")]
        public Input<bool>? Debug { get; set; }

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
        /// add the script to the commands to be executed.
        /// </summary>
        [Input("expressions")]
        public InputUnion<string, ImmutableArray<string>>? Expressions { get; set; }

        /// <summary>
        /// add the contents of script-file to the commands to be executed.
        /// </summary>
        [Input("files")]
        public InputUnion<string, ImmutableArray<string>>? Files { get; set; }

        /// <summary>
        /// follow symlinks when processing in place
        /// </summary>
        [Input("followSymlinks")]
        public Input<bool>? FollowSymlinks { get; set; }

        /// <summary>
        /// display this help and exit.
        /// </summary>
        [Input("help")]
        public Input<bool>? Help { get; set; }

        /// <summary>
        /// edit files in place (makes backup if SUFFIX supplied)
        /// </summary>
        [Input("inPlace")]
        public Input<string>? InPlace { get; set; }

        /// <summary>
        /// corresponds to the [input-file]... argument(s).
        /// </summary>
        [Input("inputFiles")]
        public InputUnion<string, ImmutableArray<string>>? InputFiles { get; set; }

        /// <summary>
        /// At what stage(s) in the resource lifecycle should the command be run
        /// </summary>
        [Input("lifecycle")]
        public UnMango.KubernetesTheHardWay.Tools.CommandLifecycle? Lifecycle { get; set; }

        /// <summary>
        /// specify the desired line-wrap length for the `l' command
        /// </summary>
        [Input("lineLength")]
        public Input<int>? LineLength { get; set; }

        /// <summary>
        /// separate lines by NUL characters
        /// </summary>
        [Input("nullData")]
        public Input<bool>? NullData { get; set; }

        /// <summary>
        /// disable all GNU extensions.
        /// </summary>
        [Input("posix")]
        public Input<bool>? Posix { get; set; }

        /// <summary>
        /// suppress automatic printing of pattern space. Same as `silent`.
        /// </summary>
        [Input("quiet")]
        public Input<bool>? Quiet { get; set; }

        /// <summary>
        /// use extended regular expressions in the script (for portability use POSIX -E).
        /// </summary>
        [Input("regexpExtended")]
        public Input<bool>? RegexpExtended { get; set; }

        /// <summary>
        /// operate in sandbox mode (disable e/r/w commands).
        /// </summary>
        [Input("sandbox")]
        public Input<bool>? Sandbox { get; set; }

        /// <summary>
        /// script only if no other script.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// consider files as separate rather than as a single, continuous long stream.
        /// </summary>
        [Input("separate")]
        public Input<bool>? Separate { get; set; }

        /// <summary>
        /// suppress automatic printing of pattern space. Same as `quiet`.
        /// </summary>
        [Input("silent")]
        public Input<bool>? Silent { get; set; }

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
        /// load minimal amounts of data from the input files and flush the output buffers more often.
        /// </summary>
        [Input("unbuffered")]
        public Input<bool>? Unbuffered { get; set; }

        /// <summary>
        /// output version information and exit.
        /// </summary>
        [Input("version")]
        public Input<bool>? Version { get; set; }

        public SedArgs()
        {
        }
        public static new SedArgs Empty => new SedArgs();
    }
}