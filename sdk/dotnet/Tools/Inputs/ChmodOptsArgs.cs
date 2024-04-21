// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Tools.Inputs
{

    /// <summary>
    /// Abstraction over the `chmod` utility on a remote system.
    /// </summary>
    public sealed class ChmodOptsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Like verbose but report only when a change is made.
        /// </summary>
        [Input("changes")]
        public Input<bool>? Changes { get; set; }

        [Input("files", required: true)]
        private InputList<string>? _files;

        /// <summary>
        /// Corresponds to the [FILE] argument.
        /// </summary>
        public InputList<string> Files
        {
            get => _files ?? (_files = new InputList<string>());
            set => _files = value;
        }

        /// <summary>
        /// Display help and exit.
        /// </summary>
        [Input("help")]
        public Input<bool>? Help { get; set; }

        /// <summary>
        /// Modes may be absolute or symbolic. An absolute mode is an octal number...
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        /// <summary>
        /// Do not treat '/' specially (the default).
        /// </summary>
        [Input("noPreserveRoot")]
        public Input<bool>? NoPreserveRoot { get; set; }

        /// <summary>
        /// Fail to operate recursively on '/'.
        /// </summary>
        [Input("preserveRoot")]
        public Input<bool>? PreserveRoot { get; set; }

        /// <summary>
        /// Suppress most error messages. Same as `silent`.
        /// </summary>
        [Input("quiet")]
        public Input<bool>? Quiet { get; set; }

        /// <summary>
        /// Change files and directories recursively.
        /// </summary>
        [Input("recursive")]
        public Input<bool>? Recursive { get; set; }

        /// <summary>
        /// Use RFILE's mode instead of specifying MODE values. RFILE is always dereferenced if a symbolic link.
        /// </summary>
        [Input("reference")]
        public Input<string>? Reference { get; set; }

        /// <summary>
        /// Suppress most error messages. Same as `quiet`.
        /// </summary>
        [Input("silent")]
        public Input<bool>? Silent { get; set; }

        /// <summary>
        /// Output version information and exit.
        /// </summary>
        [Input("version")]
        public Input<bool>? Version { get; set; }

        public ChmodOptsArgs()
        {
        }
        public static new ChmodOptsArgs Empty => new ChmodOptsArgs();
    }
}
