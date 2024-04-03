// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Tools
{
    /// <summary>
    /// Abstraction over the `etcdctl` utility on a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:tools:Etcdctl")]
    public partial class Etcdctl : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// Represents the command run on the remote system.
        /// </summary>
        [Output("command")]
        public Output<Pulumi.Command.Remote.Command?> Command { get; private set; } = null!;

        /// <summary>
        /// Connection details for the remote system.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection?> Connection { get; private set; } = null!;


        /// <summary>
        /// Create a Etcdctl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Etcdctl(string name, EtcdctlArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tools:Etcdctl", name, args ?? new EtcdctlArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class EtcdctlArgs : global::Pulumi.ResourceArgs
    {
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        [Input("cert")]
        public Input<string>? Cert { get; set; }

        [Input("commands")]
        private InputList<UnMango.KubernetesTheHardWay.Tools.EtcdctlCommand>? _commands;
        public InputList<UnMango.KubernetesTheHardWay.Tools.EtcdctlCommand> Commands
        {
            get => _commands ?? (_commands = new InputList<UnMango.KubernetesTheHardWay.Tools.EtcdctlCommand>());
            set => _commands = value;
        }

        /// <summary>
        /// Connection details for the remote system.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("endpoints")]
        public Input<string>? Endpoints { get; set; }

        [Input("env")]
        private InputMap<string>? _env;
        public InputMap<string> Env
        {
            get => _env ?? (_env = new InputMap<string>());
            set => _env = value;
        }

        [Input("key")]
        public Input<string>? Key { get; set; }

        public EtcdctlArgs()
        {
        }
        public static new EtcdctlArgs Empty => new EtcdctlArgs();
    }
}
