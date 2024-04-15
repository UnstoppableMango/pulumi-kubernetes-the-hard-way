// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Remote.Outputs
{

    /// <summary>
    /// Etcd node description.
    /// </summary>
    [OutputType]
    public sealed class EtcdNode
    {
        /// <summary>
        /// The CPU architecture of the node.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Remote.Architecture? Architecture;
        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        public readonly Pulumi.Command.Remote.Outputs.Connection Connection;
        /// <summary>
        /// The internal IP of the node.
        /// </summary>
        public readonly string InternalIp;

        [OutputConstructor]
        private EtcdNode(
            UnMango.KubernetesTheHardWay.Remote.Architecture? architecture,

            Pulumi.Command.Remote.Outputs.Connection connection,

            string internalIp)
        {
            Architecture = architecture;
            Connection = connection;
            InternalIp = internalIp;
        }
    }
}
