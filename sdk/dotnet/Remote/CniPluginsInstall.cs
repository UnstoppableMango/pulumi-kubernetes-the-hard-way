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
    /// Installs cni-plugins on a remote system
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:CniPluginsInstall")]
    public partial class CniPluginsInstall : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The CPU architecture to install.
        /// </summary>
        [Output("architecture")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Architecture> Architecture { get; private set; } = null!;

        /// <summary>
        /// The name of the downloaded archive.
        /// </summary>
        [Output("archiveName")]
        public Output<string> ArchiveName { get; private set; } = null!;

        /// <summary>
        /// The bandwidth mv operation.
        /// </summary>
        [Output("bandwidthMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> BandwidthMv { get; private set; } = null!;

        /// <summary>
        /// The bandwidth path on the remote system
        /// </summary>
        [Output("bandwidthPath")]
        public Output<string> BandwidthPath { get; private set; } = null!;

        /// <summary>
        /// The bridge mv operation.
        /// </summary>
        [Output("bridgeMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> BridgeMv { get; private set; } = null!;

        /// <summary>
        /// The bridge path on the remote system
        /// </summary>
        [Output("bridgePath")]
        public Output<string> BridgePath { get; private set; } = null!;

        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Output("connection")]
        public Output<Pulumi.Command.Remote.Outputs.Connection> Connection { get; private set; } = null!;

        /// <summary>
        /// The dhcp mv operation.
        /// </summary>
        [Output("dhcpMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> DhcpMv { get; private set; } = null!;

        /// <summary>
        /// The dhcp path on the remote system
        /// </summary>
        [Output("dhcpPath")]
        public Output<string> DhcpPath { get; private set; } = null!;

        /// <summary>
        /// The directory to install the binary to.
        /// </summary>
        [Output("directory")]
        public Output<string> Directory { get; private set; } = null!;

        /// <summary>
        /// The download operation.
        /// </summary>
        [Output("download")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Download> Download { get; private set; } = null!;

        /// <summary>
        /// The dummy mv operation.
        /// </summary>
        [Output("dummyMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> DummyMv { get; private set; } = null!;

        /// <summary>
        /// The dummy path on the remote system
        /// </summary>
        [Output("dummyPath")]
        public Output<string> DummyPath { get; private set; } = null!;

        /// <summary>
        /// The firewall mv operation.
        /// </summary>
        [Output("firewallMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> FirewallMv { get; private set; } = null!;

        /// <summary>
        /// The firewall path on the remote system
        /// </summary>
        [Output("firewallPath")]
        public Output<string> FirewallPath { get; private set; } = null!;

        /// <summary>
        /// The hostDevice mv operation.
        /// </summary>
        [Output("hostDeviceMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> HostDeviceMv { get; private set; } = null!;

        /// <summary>
        /// The hostDevice path on the remote system
        /// </summary>
        [Output("hostDevicePath")]
        public Output<string> HostDevicePath { get; private set; } = null!;

        /// <summary>
        /// The hostLocal mv operation.
        /// </summary>
        [Output("hostLocalMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> HostLocalMv { get; private set; } = null!;

        /// <summary>
        /// The hostLocal path on the remote system
        /// </summary>
        [Output("hostLocalPath")]
        public Output<string> HostLocalPath { get; private set; } = null!;

        /// <summary>
        /// The ipvlan mv operation.
        /// </summary>
        [Output("ipvlanMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> IpvlanMv { get; private set; } = null!;

        /// <summary>
        /// The ipvlan path on the remote system
        /// </summary>
        [Output("ipvlanPath")]
        public Output<string> IpvlanPath { get; private set; } = null!;

        /// <summary>
        /// The loopback mv operation.
        /// </summary>
        [Output("loopbackMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> LoopbackMv { get; private set; } = null!;

        /// <summary>
        /// The loopback path on the remote system
        /// </summary>
        [Output("loopbackPath")]
        public Output<string> LoopbackPath { get; private set; } = null!;

        /// <summary>
        /// The macvlan mv operation.
        /// </summary>
        [Output("macvlanMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> MacvlanMv { get; private set; } = null!;

        /// <summary>
        /// The macvlan path on the remote system
        /// </summary>
        [Output("macvlanPath")]
        public Output<string> MacvlanPath { get; private set; } = null!;

        /// <summary>
        /// The mkdir operation.
        /// </summary>
        [Output("mkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir> Mkdir { get; private set; } = null!;

        /// <summary>
        /// The mktemp operation.
        /// </summary>
        [Output("mktemp")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mktemp> Mktemp { get; private set; } = null!;

        /// <summary>
        /// The path to the installed binary.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The portmap mv operation.
        /// </summary>
        [Output("portmapMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> PortmapMv { get; private set; } = null!;

        /// <summary>
        /// The portmap path on the remote system
        /// </summary>
        [Output("portmapPath")]
        public Output<string> PortmapPath { get; private set; } = null!;

        /// <summary>
        /// The ptp mv operation.
        /// </summary>
        [Output("ptpMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> PtpMv { get; private set; } = null!;

        /// <summary>
        /// The ptp path on the remote system
        /// </summary>
        [Output("ptpPath")]
        public Output<string> PtpPath { get; private set; } = null!;

        /// <summary>
        /// The rm operation.
        /// </summary>
        [Output("rm")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Rm> Rm { get; private set; } = null!;

        /// <summary>
        /// The sbr mv operation.
        /// </summary>
        [Output("sbrMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> SbrMv { get; private set; } = null!;

        /// <summary>
        /// The sbr path on the remote system
        /// </summary>
        [Output("sbrPath")]
        public Output<string> SbrPath { get; private set; } = null!;

        /// <summary>
        /// The static mv operation.
        /// </summary>
        [Output("staticMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> StaticMv { get; private set; } = null!;

        /// <summary>
        /// The static path on the remote system
        /// </summary>
        [Output("staticPath")]
        public Output<string> StaticPath { get; private set; } = null!;

        /// <summary>
        /// The tap mv operation.
        /// </summary>
        [Output("tapMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> TapMv { get; private set; } = null!;

        /// <summary>
        /// The tap path on the remote system
        /// </summary>
        [Output("tapPath")]
        public Output<string> TapPath { get; private set; } = null!;

        /// <summary>
        /// The tar operation.
        /// </summary>
        [Output("tar")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Tar> Tar { get; private set; } = null!;

        /// <summary>
        /// The tuning mv operation.
        /// </summary>
        [Output("tuningMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> TuningMv { get; private set; } = null!;

        /// <summary>
        /// The tuning path on the remote system
        /// </summary>
        [Output("tuningPath")]
        public Output<string> TuningPath { get; private set; } = null!;

        /// <summary>
        /// The url used to download the binary.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The version to install.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The vlan mv operation.
        /// </summary>
        [Output("vlanMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> VlanMv { get; private set; } = null!;

        /// <summary>
        /// The vlan path on the remote system
        /// </summary>
        [Output("vlanPath")]
        public Output<string> VlanPath { get; private set; } = null!;

        /// <summary>
        /// The vrf mv operation.
        /// </summary>
        [Output("vrfMv")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> VrfMv { get; private set; } = null!;

        /// <summary>
        /// The vrf path on the remote system
        /// </summary>
        [Output("vrfPath")]
        public Output<string> VrfPath { get; private set; } = null!;


        /// <summary>
        /// Create a CniPluginsInstall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CniPluginsInstall(string name, CniPluginsInstallArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:CniPluginsInstall", name, args ?? new CniPluginsInstallArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class CniPluginsInstallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU architecture to install.
        /// </summary>
        [Input("architecture")]
        public Input<UnMango.KubernetesTheHardWay.Remote.Architecture>? Architecture { get; set; }

        /// <summary>
        /// The parameters with which to connect to the remote host.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The directory to install the binary to.
        /// </summary>
        [Input("directory")]
        public Input<string>? Directory { get; set; }

        /// <summary>
        /// The version to install.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public CniPluginsInstallArgs()
        {
            Directory = "/opt/cni/bin";
        }
        public static new CniPluginsInstallArgs Empty => new CniPluginsInstallArgs();
    }
}
