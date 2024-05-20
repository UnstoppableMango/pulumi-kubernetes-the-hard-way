// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config
{
    /// <summary>
    /// Get the `loopback` configuration.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration")]
    public partial class CniLoopbackPluginConfiguration : global::Pulumi.ComponentResource
    {
        [Output("result")]
        public Output<UnMango.KubernetesTheHardWay.Config.Outputs.CniLoopbackPluginConfiguration> Result { get; private set; } = null!;

        /// <summary>
        /// The yaml representation of the manifest.
        /// </summary>
        [Output("yaml")]
        public Output<string> Yaml { get; private set; } = null!;


        /// <summary>
        /// Create a CniLoopbackPluginConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CniLoopbackPluginConfiguration(string name, CniLoopbackPluginConfigurationArgs? args = null, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration", name, args ?? new CniLoopbackPluginConfigurationArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class CniLoopbackPluginConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CNI version.
        /// </summary>
        [Input("cniVersion")]
        public Input<string>? CniVersion { get; set; }

        /// <summary>
        /// CNI plugin name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path to put the configuration file on the remote system
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// CNI plugin type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CniLoopbackPluginConfigurationArgs()
        {
        }
        public static new CniLoopbackPluginConfigurationArgs Empty => new CniLoopbackPluginConfigurationArgs();
    }
}
