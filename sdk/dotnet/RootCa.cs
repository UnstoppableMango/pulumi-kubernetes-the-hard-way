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
        [Output("cert")]
        public Output<Pulumi.Tls.SelfSignedCert> Cert { get; private set; } = null!;

        [Output("certPem")]
        public Output<string> CertPem { get; private set; } = null!;

        [Output("key")]
        public Output<Pulumi.Tls.PrivateKey> Key { get; private set; } = null!;

        [Output("keyPem")]
        public Output<string> KeyPem { get; private set; } = null!;


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
        /// Creates a Certificate configured for the cluster admin.
        /// </summary>
        public global::Pulumi.Output<RootCaGetAdminCertificateResult> GetAdminCertificate(RootCaGetAdminCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaGetAdminCertificateResult>("kubernetes-the-hard-way:index:RootCa/getAdminCertificate", args ?? new RootCaGetAdminCertificateArgs(), this);

        /// <summary>
        /// Creates a Certificate configured for the api-server.
        /// </summary>
        public global::Pulumi.Output<RootCaGetApiServerCertificateResult> GetApiServerCertificate(RootCaGetApiServerCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaGetApiServerCertificateResult>("kubernetes-the-hard-way:index:RootCa/getApiServerCertificate", args ?? new RootCaGetApiServerCertificateArgs(), this);

        /// <summary>
        /// Creates a Certificate configured for the controller manager.
        /// </summary>
        public global::Pulumi.Output<RootCaGetControllerManagerCertificateResult> GetControllerManagerCertificate(RootCaGetControllerManagerCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaGetControllerManagerCertificateResult>("kubernetes-the-hard-way:index:RootCa/getControllerManagerCertificate", args ?? new RootCaGetControllerManagerCertificateArgs(), this);

        /// <summary>
        /// Creates a Certificate configured for the kube-proxy.
        /// </summary>
        public global::Pulumi.Output<RootCaGetKubeProxyCertificateResult> GetKubeProxyCertificate(RootCaGetKubeProxyCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaGetKubeProxyCertificateResult>("kubernetes-the-hard-way:index:RootCa/getKubeProxyCertificate", args ?? new RootCaGetKubeProxyCertificateArgs(), this);

        /// <summary>
        /// Creates a Certificate configured for the kube-scheduler.
        /// </summary>
        public global::Pulumi.Output<RootCaGetKubeSchedulerCertificateResult> GetKubeSchedulerCertificate(RootCaGetKubeSchedulerCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaGetKubeSchedulerCertificateResult>("kubernetes-the-hard-way:index:RootCa/getKubeSchedulerCertificate", args ?? new RootCaGetKubeSchedulerCertificateArgs(), this);

        /// <summary>
        /// Creates a Certificate configured for a kubelet.
        /// </summary>
        public global::Pulumi.Output<RootCaGetKubeletCertificateResult> GetKubeletCertificate(RootCaGetKubeletCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaGetKubeletCertificateResult>("kubernetes-the-hard-way:index:RootCa/getKubeletCertificate", args ?? new RootCaGetKubeletCertificateArgs(), this);

        /// <summary>
        /// Creates a Certificate configured for the kube-scheduler.
        /// </summary>
        public global::Pulumi.Output<RootCaGetServiceAccountsCertificateResult> GetServiceAccountsCertificate(RootCaGetServiceAccountsCertificateArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaGetServiceAccountsCertificateResult>("kubernetes-the-hard-way:index:RootCa/getServiceAccountsCertificate", args ?? new RootCaGetServiceAccountsCertificateArgs(), this);

        /// <summary>
        /// Creates a RemoteFile resource representing the copy operation.
        /// </summary>
        public global::Pulumi.Output<RootCaInstallOnResult> InstallOn(RootCaInstallOnArgs args)
            => global::Pulumi.Deployment.Instance.Call<RootCaInstallOnResult>("kubernetes-the-hard-way:index:RootCa/installOn", args ?? new RootCaInstallOnArgs(), this);
    }

    public sealed class RootCaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

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
        public Input<string>? EcdsaCurve { get; set; }

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
    /// The set of arguments for the <see cref="RootCa.GetAdminCertificate"/> method.
    /// </summary>
    public sealed class RootCaGetAdminCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
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
        public Input<string>? EcdsaCurve { get; set; }

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

        public RootCaGetAdminCertificateArgs()
        {
        }
        public static new RootCaGetAdminCertificateArgs Empty => new RootCaGetAdminCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.GetAdminCertificate"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaGetAdminCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private RootCaGetAdminCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.GetApiServerCertificate"/> method.
    /// </summary>
    public sealed class RootCaGetApiServerCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
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
        public Input<string>? EcdsaCurve { get; set; }

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

        public RootCaGetApiServerCertificateArgs()
        {
        }
        public static new RootCaGetApiServerCertificateArgs Empty => new RootCaGetApiServerCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.GetApiServerCertificate"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaGetApiServerCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private RootCaGetApiServerCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.GetControllerManagerCertificate"/> method.
    /// </summary>
    public sealed class RootCaGetControllerManagerCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
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
        public Input<string>? EcdsaCurve { get; set; }

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

        public RootCaGetControllerManagerCertificateArgs()
        {
        }
        public static new RootCaGetControllerManagerCertificateArgs Empty => new RootCaGetControllerManagerCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.GetControllerManagerCertificate"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaGetControllerManagerCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private RootCaGetControllerManagerCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.GetKubeProxyCertificate"/> method.
    /// </summary>
    public sealed class RootCaGetKubeProxyCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
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
        public Input<string>? EcdsaCurve { get; set; }

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

        public RootCaGetKubeProxyCertificateArgs()
        {
        }
        public static new RootCaGetKubeProxyCertificateArgs Empty => new RootCaGetKubeProxyCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.GetKubeProxyCertificate"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaGetKubeProxyCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private RootCaGetKubeProxyCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.GetKubeSchedulerCertificate"/> method.
    /// </summary>
    public sealed class RootCaGetKubeSchedulerCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
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
        public Input<string>? EcdsaCurve { get; set; }

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

        public RootCaGetKubeSchedulerCertificateArgs()
        {
        }
        public static new RootCaGetKubeSchedulerCertificateArgs Empty => new RootCaGetKubeSchedulerCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.GetKubeSchedulerCertificate"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaGetKubeSchedulerCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private RootCaGetKubeSchedulerCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.GetKubeletCertificate"/> method.
    /// </summary>
    public sealed class RootCaGetKubeletCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
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
        public Input<string>? EcdsaCurve { get; set; }

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

        public RootCaGetKubeletCertificateArgs()
        {
        }
        public static new RootCaGetKubeletCertificateArgs Empty => new RootCaGetKubeletCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.GetKubeletCertificate"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaGetKubeletCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private RootCaGetKubeletCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.GetServiceAccountsCertificate"/> method.
    /// </summary>
    public sealed class RootCaGetServiceAccountsCertificateArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
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
        public Input<string>? EcdsaCurve { get; set; }

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

        public RootCaGetServiceAccountsCertificateArgs()
        {
        }
        public static new RootCaGetServiceAccountsCertificateArgs Empty => new RootCaGetServiceAccountsCertificateArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.GetServiceAccountsCertificate"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaGetServiceAccountsCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private RootCaGetServiceAccountsCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="RootCa.InstallOn"/> method.
    /// </summary>
    public sealed class RootCaInstallOnArgs : global::Pulumi.CallArgs
    {
        /// <summary>
        /// The connection details.
        /// </summary>
        [Input("connection", required: true)]
        public Input<Inputs.ConnectionArgs> Connection { get; set; } = null!;

        /// <summary>
        /// The path to install to.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public RootCaInstallOnArgs()
        {
        }
        public static new RootCaInstallOnArgs Empty => new RootCaInstallOnArgs();
    }

    /// <summary>
    /// The results of the <see cref="RootCa.InstallOn"/> method.
    /// </summary>
    [OutputType]
    public sealed class RootCaInstallOnResult
    {
        public readonly UnMango.KubernetesTheHardWay.RemoteFile File;

        [OutputConstructor]
        private RootCaInstallOnResult(UnMango.KubernetesTheHardWay.RemoteFile file)
        {
            File = file;
        }
    }
}
