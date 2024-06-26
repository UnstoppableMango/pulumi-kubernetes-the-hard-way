// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config.Outputs
{

    /// <summary>
    /// The CNI plugins IPAM
    /// </summary>
    [OutputType]
    public sealed class CniBridgeIpam
    {
        /// <summary>
        /// IPAM ranges.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, string>> Ranges;
        /// <summary>
        /// IPAM routes.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, string>> Routes;
        /// <summary>
        /// CNI bridge IPAM type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private CniBridgeIpam(
            ImmutableArray<ImmutableDictionary<string, string>> ranges,

            ImmutableArray<ImmutableDictionary<string, string>> routes,

            string? type)
        {
            Ranges = ranges;
            Routes = routes;
            Type = type;
        }
    }
}
