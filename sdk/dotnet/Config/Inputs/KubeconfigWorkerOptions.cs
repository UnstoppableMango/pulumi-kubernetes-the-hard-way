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

    public sealed class KubeconfigWorkerOptions : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// TODO
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Input("publicIp", required: true)]
        public string PublicIp { get; set; } = null!;

        [Input("type")]
        public string? Type { get; set; }

        public KubeconfigWorkerOptions()
        {
        }
        public static new KubeconfigWorkerOptions Empty => new KubeconfigWorkerOptions();
    }
}
