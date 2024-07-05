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
    /// Abstraction over the `systemctl` utility on a remote system.
    /// </summary>
    [OutputType]
    public sealed class SystemctlOpts
    {
        /// <summary>
        /// Corresponds to the COMMAND argument.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Tools.SystemctlCommand Command;
        /// <summary>
        /// Corresponds to the `--now` option.
        /// </summary>
        public readonly bool? Now;
        /// <summary>
        /// Corresponds to the [PATTERN] argument
        /// </summary>
        public readonly string? Pattern;
        /// <summary>
        /// Corresponds to the [UNIT...] argument.
        /// </summary>
        public readonly string? Unit;

        [OutputConstructor]
        private SystemctlOpts(
            UnMango.KubernetesTheHardWay.Tools.SystemctlCommand command,

            bool? now,

            string? pattern,

            string? unit)
        {
            Command = command;
            Now = now;
            Pattern = pattern;
            Unit = unit;
        }
    }
}
