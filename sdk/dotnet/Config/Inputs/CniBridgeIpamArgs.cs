// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config.Inputs
{

    /// <summary>
    /// The CNI plugins IPAM
    /// </summary>
    public sealed class CniBridgeIpamArgs : global::Pulumi.ResourceArgs
    {
        [Input("ranges")]
        private InputList<ImmutableDictionary<string, string>>? _ranges;

        /// <summary>
        /// IPAM ranges.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Ranges
        {
            get => _ranges ?? (_ranges = new InputList<ImmutableDictionary<string, string>>());
            set => _ranges = value;
        }

        [Input("routes")]
        private InputList<ImmutableDictionary<string, string>>? _routes;

        /// <summary>
        /// IPAM routes.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Routes
        {
            get => _routes ?? (_routes = new InputList<ImmutableDictionary<string, string>>());
            set => _routes = value;
        }

        /// <summary>
        /// CNI bridge IPAM type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CniBridgeIpamArgs()
        {
        }
        public static new CniBridgeIpamArgs Empty => new CniBridgeIpamArgs();
    }
}
