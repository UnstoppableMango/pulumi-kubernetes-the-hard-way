// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay
{
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:index:RootCa")]
    public partial class RootCa : global::Pulumi.ComponentResource
    {
        [Output("allowedUses")]
        public Output<ImmutableArray<UnMango.KubernetesTheHardWay.AllowedUsage>> AllowedUses { get; private set; } = null!;

        [Output("cert")]
        public Output<Pulumi.Tls.SelfSignedCert> Cert { get; private set; } = null!;

        [Output("certPem")]
        public Output<string> CertPem { get; private set; } = null!;

        [Output("key")]
        public Output<Pulumi.Tls.PrivateKey> Key { get; private set; } = null!;

        [Output("privateKeyPem")]
        public Output<string> PrivateKeyPem { get; private set; } = null!;

        [Output("publicKeyPem")]
        public Output<string> PublicKeyPem { get; private set; } = null!;


        /// <summary>
        /// Create a RootCa resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RootCa(string name, RootCaArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:index:RootCa", name, args ?? new RootCaArgs(), MakeResourceOptions(options, ""), remote: true)
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

        /// <summary>
        /// Creates a RemoteFile resource representing the copy operation.
        /// </summary>
        public global::Pulumi.Output<UnMango.KubernetesTheHardWay.RemoteFile> InstallCert(RootCaInstallCertArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaInstallCertResult>("kubernetes-the-hard-way:index:RootCa/installCert", args ?? new RootCaInstallCertArgs(), this).Apply(v => v.Result);

        /// <summary>
        /// Creates a RemoteFile resource representing the copy operation.
        /// </summary>
        public global::Pulumi.Output<UnMango.KubernetesTheHardWay.RemoteFile> InstallKey(RootCaInstallKeyArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaInstallKeyResult>("kubernetes-the-hard-way:index:RootCa/installKey", args ?? new RootCaInstallKeyArgs(), this).Apply(v => v.Result);

        /// <summary>
        /// Creates a Certificate configured for the current authority.
        /// </summary>
        public global::Pulumi.Output<UnMango.KubernetesTheHardWay.Certificate> NewCertificate(RootCaNewCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaNewCertificateResult>("kubernetes-the-hard-way:index:RootCa/newCertificate", args ?? new RootCaNewCertificateArgs(), this).Apply(v => v.Result);
    }

    public sealed class RootCaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<UnMango.KubernetesTheHardWay.Algorithm>? Algorithm { get; set; }

        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested.
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("earlyRenewalHours")]
        public Input<int>? EarlyRenewalHours { get; set; }

        /// <summary>
        /// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<UnMango.KubernetesTheHardWay.EcdsaCurve>? EcdsaCurve { get; set; }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        /// <summary>
        /// Should the generated certificate include an authority key identifier.
        /// </summary>
        [Input("setAuthorityKeyId")]
        public Input<bool>? SetAuthorityKeyId { get; set; }

        /// <summary>
        /// Should the generated certificate include a subject key identifier.
        /// </summary>
        [Input("setSubjectKeyId")]
        public Input<bool>? SetSubjectKeyId { get; set; }

        [Input("subject")]
        public Input<Pulumi.Tls.Inputs.SelfSignedCertSubjectArgs>? Subject { get; set; }

        [Input("uris")]
        private InputList<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested.
        /// </summary>
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid.
        /// </summary>
        [Input("validityPeriodHours", required: true)]
        public Input<int> ValidityPeriodHours { get; set; } = null!;

        public RootCaArgs()
        {
        }
        public static new RootCaArgs Empty => new RootCaArgs();
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.InstallCert"/> method.
    /// </summary>
    public sealed class RootCaInstallCertArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("options")]
        public Inputs.ResourceOptionsArgs? Options { get; set; }

        /// <summary>
        /// The path to install to.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public RootCaInstallCertArgs()
        {
        }
        public static new RootCaInstallCertArgs Empty => new RootCaInstallCertArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.InstallCert"/> method.
    /// </summary>
    [OutputType]
    internal sealed class RootCaInstallCertResult
    {
        public readonly UnMango.KubernetesTheHardWay.RemoteFile Result;

        [OutputConstructor]
        private RootCaInstallCertResult(UnMango.KubernetesTheHardWay.RemoteFile result)
        {
            Result = result;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.InstallKey"/> method.
    /// </summary>
    public sealed class RootCaInstallKeyArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("options")]
        public Inputs.ResourceOptionsArgs? Options { get; set; }

        /// <summary>
        /// The path to install to.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public RootCaInstallKeyArgs()
        {
        }
        public static new RootCaInstallKeyArgs Empty => new RootCaInstallKeyArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.InstallKey"/> method.
    /// </summary>
    [OutputType]
    internal sealed class RootCaInstallKeyResult
    {
        public readonly UnMango.KubernetesTheHardWay.RemoteFile Result;

        [OutputConstructor]
        private RootCaInstallKeyResult(UnMango.KubernetesTheHardWay.RemoteFile result)
        {
            Result = result;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.NewCertificate"/> method.
    /// </summary>
    public sealed class RootCaNewCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<UnMango.KubernetesTheHardWay.Algorithm> Algorithm { get; set; } = null!;

        [Input("allowedUses", required: true)]
        private InputList<UnMango.KubernetesTheHardWay.AllowedUsage>? _allowedUses;
        public InputList<UnMango.KubernetesTheHardWay.AllowedUsage> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<UnMango.KubernetesTheHardWay.AllowedUsage>());
            set => _allowedUses = value;
        }

        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested.
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("earlyRenewalHours")]
        public Input<int>? EarlyRenewalHours { get; set; }

        /// <summary>
        /// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<UnMango.KubernetesTheHardWay.EcdsaCurve>? EcdsaCurve { get; set; }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        [Input("isCaCertificate")]
        public Input<bool>? IsCaCertificate { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("options")]
        public Inputs.ResourceOptionsArgs? Options { get; set; }

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        /// <summary>
        /// Should the generated certificate include an authority key identifier.
        /// </summary>
        [Input("setAuthorityKeyId")]
        public Input<bool>? SetAuthorityKeyId { get; set; }

        /// <summary>
        /// Should the generated certificate include a subject key identifier.
        /// </summary>
        [Input("setSubjectKeyId")]
        public Input<bool>? SetSubjectKeyId { get; set; }

        [Input("subject")]
        public Input<Inputs.CertRequestSubjectArgs>? Subject { get; set; }

        [Input("uris")]
        private InputList<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested.
        /// </summary>
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid.
        /// </summary>
        [Input("validityPeriodHours", required: true)]
        public Input<int> ValidityPeriodHours { get; set; } = null!;

        public RootCaNewCertificateArgs()
        {
        }
        public static new RootCaNewCertificateArgs Empty => new RootCaNewCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.NewCertificate"/> method.
    /// </summary>
    [OutputType]
    internal sealed class RootCaNewCertificateResult
    {
        /// <summary>
        /// The issued certificate.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Certificate Result;

        [OutputConstructor]
        private RootCaNewCertificateResult(UnMango.KubernetesTheHardWay.Certificate result)
        {
            Result = result;
        }
    }
}
