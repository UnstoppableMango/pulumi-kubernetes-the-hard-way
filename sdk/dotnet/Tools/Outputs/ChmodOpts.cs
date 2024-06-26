// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Tools.Outputs
{

    /// <summary>
    /// Abstraction over the `chmod` utility on a remote system.
    /// </summary>
    [OutputType]
    public sealed class ChmodOpts
    {
        /// <summary>
        /// Like verbose but report only when a change is made.
        /// </summary>
        public readonly bool? Changes;
        /// <summary>
        /// Corresponds to the [FILE] argument.
        /// </summary>
        public readonly ImmutableArray<string> Files;
        /// <summary>
        /// Display help and exit.
        /// </summary>
        public readonly bool? Help;
        /// <summary>
        /// Modes may be absolute or symbolic. An absolute mode is an octal number...
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// Do not treat '/' specially (the default).
        /// </summary>
        public readonly bool? NoPreserveRoot;
        /// <summary>
        /// Fail to operate recursively on '/'.
        /// </summary>
        public readonly bool? PreserveRoot;
        /// <summary>
        /// Suppress most error messages. Same as `silent`.
        /// </summary>
        public readonly bool? Quiet;
        /// <summary>
        /// Change files and directories recursively.
        /// </summary>
        public readonly bool? Recursive;
        /// <summary>
        /// Use RFILE's mode instead of specifying MODE values. RFILE is always dereferenced if a symbolic link.
        /// </summary>
        public readonly string? Reference;
        /// <summary>
        /// Suppress most error messages. Same as `quiet`.
        /// </summary>
        public readonly bool? Silent;
        /// <summary>
        /// Output version information and exit.
        /// </summary>
        public readonly bool? Version;

        [OutputConstructor]
        private ChmodOpts(
            bool? changes,

            ImmutableArray<string> files,

            bool? help,

            string mode,

            bool? noPreserveRoot,

            bool? preserveRoot,

            bool? quiet,

            bool? recursive,

            string? reference,

            bool? silent,

            bool? version)
        {
            Changes = changes;
            Files = files;
            Help = help;
            Mode = mode;
            NoPreserveRoot = noPreserveRoot;
            PreserveRoot = preserveRoot;
            Quiet = quiet;
            Recursive = recursive;
            Reference = reference;
            Silent = silent;
            Version = version;
        }
    }
}
