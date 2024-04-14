// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Remote.Inputs
{

    /// <summary>
    /// Etcd node description.
    /// </summary>
    public sealed class EtcdNodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The internal IP of the node.
        /// </summary>
        [Input("internalIp", required: true)]
        public Input<string> InternalIp { get; set; } = null!;

        public EtcdNodeArgs()
        {
        }
        public static new EtcdNodeArgs Empty => new EtcdNodeArgs();
    }
}
