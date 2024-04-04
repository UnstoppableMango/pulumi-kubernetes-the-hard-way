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
    /// Represents an etcd binary on a remote system.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:remote:EtcdInstall")]
    public partial class EtcdInstall : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The etcd CPU architecture.
        /// </summary>
        [Output("architecture")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Architecture> Architecture { get; private set; } = null!;

        /// <summary>
        /// The name of the etcd release archive.
        /// </summary>
        [Output("archiveName")]
        public Output<string> ArchiveName { get; private set; } = null!;

        /// <summary>
        /// The remote certificate authority file.
        /// </summary>
        [Output("caFile")]
        public Output<UnMango.KubernetesTheHardWay.Remote.File?> CaFile { get; private set; } = null!;

        /// <summary>
        /// The remote certificate file.
        /// </summary>
        [Output("certFile")]
        public Output<UnMango.KubernetesTheHardWay.Remote.File?> CertFile { get; private set; } = null!;

        /// <summary>
        /// The directory to store etcd configuration.
        /// </summary>
        [Output("configurationDirectory")]
        public Output<string> ConfigurationDirectory { get; private set; } = null!;

        /// <summary>
        /// The command used to create the configuration directory.
        /// </summary>
        [Output("configurationMkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir> ConfigurationMkdir { get; private set; } = null!;

        /// <summary>
        /// The directory etcd will use.
        /// </summary>
        [Output("dataDirectory")]
        public Output<string> DataDirectory { get; private set; } = null!;

        /// <summary>
        /// The command used to create the data directory.
        /// </summary>
        [Output("dataMkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir> DataMkdir { get; private set; } = null!;

        /// <summary>
        /// The etcd download operation.
        /// </summary>
        [Output("download")]
        public Output<UnMango.KubernetesTheHardWay.Remote.Download> Download { get; private set; } = null!;

        /// <summary>
        /// The directory where the etcd binary was downloaded to.
        /// </summary>
        [Output("downloadDirectory")]
        public Output<string> DownloadDirectory { get; private set; } = null!;

        /// <summary>
        /// The operation to create the download directory.
        /// </summary>
        [Output("downloadMkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir> DownloadMkdir { get; private set; } = null!;

        /// <summary>
        /// The path to the etcd binary on the remote system.
        /// </summary>
        [Output("etcdPath")]
        public Output<string> EtcdPath { get; private set; } = null!;

        /// <summary>
        /// The path to the etcdctl binary on the remote system.
        /// </summary>
        [Output("etcdctlPath")]
        public Output<string> EtcdctlPath { get; private set; } = null!;

        /// <summary>
        /// Directory to install the `etcd` and `etcdctl` binaries.
        /// </summary>
        [Output("installDirectory")]
        public Output<string> InstallDirectory { get; private set; } = null!;

        /// <summary>
        /// The operation to create the install directory.
        /// </summary>
        [Output("installMkdir")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mkdir> InstallMkdir { get; private set; } = null!;

        /// <summary>
        /// IP used to serve client requests and communicate with etcd peers.
        /// </summary>
        [Output("internalIp")]
        public Output<string> InternalIp { get; private set; } = null!;

        /// <summary>
        /// The remote key file.
        /// </summary>
        [Output("keyFile")]
        public Output<UnMango.KubernetesTheHardWay.Remote.File?> KeyFile { get; private set; } = null!;

        /// <summary>
        /// The operation to move the etcd binary to the install directory.
        /// </summary>
        [Output("mvEtcd")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> MvEtcd { get; private set; } = null!;

        /// <summary>
        /// The operation to move the etcdctl binary to the install directory.
        /// </summary>
        [Output("mvEtcdctl")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Mv> MvEtcdctl { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The remote systemd service.
        /// </summary>
        [Output("systemdService")]
        public Output<UnMango.KubernetesTheHardWay.Remote.SystemdService> SystemdService { get; private set; } = null!;

        /// <summary>
        /// The tar operation.
        /// </summary>
        [Output("tar")]
        public Output<UnMango.KubernetesTheHardWay.Tools.Tar> Tar { get; private set; } = null!;

        /// <summary>
        /// The url used to download etcd.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The version of etcd downloaded.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a EtcdInstall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EtcdInstall(string name, EtcdInstallArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:remote:EtcdInstall", name, args ?? new EtcdInstallArgs(), MakeResourceOptions(options, ""), remote: true)
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

        public global::Pulumi.Output<UnMango.KubernetesTheHardWay.Tools.Etcdctl> Etcdctl()
            => global::Pulumi.Deployment.Instance.Call<EtcdInstallEtcdctlResult>("kubernetes-the-hard-way:remote:EtcdInstall/etcdctl", CallArgs.Empty, this).Apply(v => v.Result);
    }

    public sealed class EtcdInstallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The etcd CPU architecture.
        /// </summary>
        [Input("architecture")]
        public Input<UnMango.KubernetesTheHardWay.Remote.Architecture>? Architecture { get; set; }

        /// <summary>
        /// The PEM encoded CA data.
        /// </summary>
        [Input("caPem", required: true)]
        public Input<string> CaPem { get; set; } = null!;

        /// <summary>
        /// The PEM encoded certificate data.
        /// </summary>
        [Input("certPem", required: true)]
        public Input<string> CertPem { get; set; } = null!;

        /// <summary>
        /// The directory to store etcd configuration.
        /// </summary>
        [Input("configurationDirectory")]
        public Input<string>? ConfigurationDirectory { get; set; }

        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The directory etcd will use.
        /// </summary>
        [Input("dataDirectory")]
        public Input<string>? DataDirectory { get; set; }

        /// <summary>
        /// Temporary directory to download files to. Defaults to `/tmp/&lt;random string&gt;`.
        /// </summary>
        [Input("downloadDirectory")]
        public Input<string>? DownloadDirectory { get; set; }

        /// <summary>
        /// Directory to install the `etcd` and `etcdctl` binaries.
        /// </summary>
        [Input("installDirectory")]
        public Input<string>? InstallDirectory { get; set; }

        /// <summary>
        /// IP used to serve client requests and communicate with etcd peers.
        /// </summary>
        [Input("internalIp", required: true)]
        public Input<string> InternalIp { get; set; } = null!;

        /// <summary>
        /// The PEM encoded key data.
        /// </summary>
        [Input("keyPem", required: true)]
        public Input<string> KeyPem { get; set; } = null!;

        /// <summary>
        /// The systemd service file dirctory.
        /// </summary>
        [Input("systemdDirectory")]
        public Input<string>? SystemdDirectory { get; set; }

        /// <summary>
        /// The version of etcd to install.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public EtcdInstallArgs()
        {
            ConfigurationDirectory = "/etc/etcd";
            DataDirectory = "/var/lib/etcd";
            InstallDirectory = "/usr/local/bin";
            SystemdDirectory = "/etc/system/systemd";
        }
        public static new EtcdInstallArgs Empty => new EtcdInstallArgs();
    }

    /// <summary>
    /// The results of the <see cref="EtcdInstall.Etcdctl"/> method.
    /// </summary>
    [OutputType]
    internal sealed class EtcdInstallEtcdctlResult
    {
        public readonly UnMango.KubernetesTheHardWay.Tools.Etcdctl Result;

        [OutputConstructor]
        private EtcdInstallEtcdctlResult(UnMango.KubernetesTheHardWay.Tools.Etcdctl result)
        {
            Result = result;
        }
    }
}
