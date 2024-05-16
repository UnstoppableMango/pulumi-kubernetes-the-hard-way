// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Remote
{
    /// <summary>
    /// The CNI bridge plugin configuration.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:CniBridgePluginConfiguration")]
    public partial class CniBridgePluginConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bridge name.
        /// </summary>
        [Output("bridge")]
        public Output<string> Bridge { get; private set; } = null!;

        /// <summary>
        /// CNI version.
        /// </summary>
        [Output("cniVersion")]
        public Output<string> CniVersion { get; private set; } = null!;

        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// IP masq.
        /// </summary>
        [Output("ipMasq")]
        public Output<bool> IpMasq { get; private set; } = null!;

        /// <summary>
        /// IPAM
        /// </summary>
        [Output("ipam")]
        public Output<Outputs.CniBridgeIpam> Ipam { get; private set; } = null!;

        /// <summary>
        /// Is gateway.
        /// </summary>
        [Output("isGateway")]
        public Output<bool> IsGateway { get; private set; } = null!;

        /// <summary>
        /// CNI plugin name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// CNI plugin type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CniBridgePluginConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CniBridgePluginConfiguration(string name, CniBridgePluginConfigurationArgs args, CustomResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:CniBridgePluginConfiguration", name, args ?? new CniBridgePluginConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }
        internal CniBridgePluginConfiguration(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:CniBridgePluginConfiguration", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private CniBridgePluginConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:CniBridgePluginConfiguration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CniBridgePluginConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CniBridgePluginConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CniBridgePluginConfiguration(name, id, options);
        }
    }

    public sealed class CniBridgePluginConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bridge name.
        /// </summary>
        [Input("bridge")]
        public Input<string>? Bridge { get; set; }

        /// <summary>
        /// CNI version.
        /// </summary>
        [Input("cniVersion")]
        public Input<string>? CniVersion { get; set; }

        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// IP masq.
        /// </summary>
        [Input("ipMasq")]
        public Input<bool>? IpMasq { get; set; }

        /// <summary>
        /// IPAM
        /// </summary>
        [Input("ipam")]
        public Input<Inputs.CniBridgeIpamArgs>? Ipam { get; set; }

        /// <summary>
        /// Is gateway.
        /// </summary>
        [Input("isGateway")]
        public Input<bool>? IsGateway { get; set; }

        /// <summary>
        /// CNI plugin name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// CNI plugin type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CniBridgePluginConfigurationArgs()
        {
        }
        public static new CniBridgePluginConfigurationArgs Empty => new CniBridgePluginConfigurationArgs();
    }
}
