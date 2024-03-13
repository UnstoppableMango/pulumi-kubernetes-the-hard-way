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
    public static class GetAdminCertificate
    {
        /// <summary>
        /// Creates a Certificate configured for the cluster admin.
        /// </summary>
        public static Task<GetAdminCertificateResult> InvokeAsync(GetAdminCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAdminCertificateResult>("kubernetes-the-hard-way:index:getAdminCertificate", args ?? new GetAdminCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a Certificate configured for the cluster admin.
        /// </summary>
        public static Output<GetAdminCertificateResult> Invoke(GetAdminCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAdminCertificateResult>("kubernetes-the-hard-way:index:getAdminCertificate", args ?? new GetAdminCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAdminCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm")]
        public string? Algorithm { get; set; }

        [Input("allowedUses", required: true)]
        private List<string>? _allowedUses;
        public List<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new List<string>());
            set => _allowedUses = value;
        }

        [Input("dnsNames")]
        private List<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested.
        /// </summary>
        public List<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new List<string>());
            set => _dnsNames = value;
        }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("earlyRenewalHours")]
        public int? EarlyRenewalHours { get; set; }

        /// <summary>
        /// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        /// </summary>
        [Input("ecdsaCurve")]
        public string? EcdsaCurve { get; set; }

        [Input("ipAddresses")]
        private List<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested.
        /// </summary>
        public List<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new List<string>());
            set => _ipAddresses = value;
        }

        [Input("isCaCertificate")]
        public bool? IsCaCertificate { get; set; }

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        /// </summary>
        [Input("rsaBits")]
        public int? RsaBits { get; set; }

        /// <summary>
        /// Should the generated certificate include an authority key identifier.
        /// </summary>
        [Input("setAuthorityKeyId")]
        public bool? SetAuthorityKeyId { get; set; }

        /// <summary>
        /// Should the generated certificate include a subject key identifier.
        /// </summary>
        [Input("setSubjectKeyId")]
        public bool? SetSubjectKeyId { get; set; }

        [Input("subject")]
        public Pulumi.Tls.Inputs.CertRequestSubject? Subject { get; set; }

        [Input("uris")]
        private List<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested.
        /// </summary>
        public List<string> Uris
        {
            get => _uris ?? (_uris = new List<string>());
            set => _uris = value;
        }

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid.
        /// </summary>
        [Input("validityPeriodHours", required: true)]
        public int ValidityPeriodHours { get; set; }

        public GetAdminCertificateArgs()
        {
        }
        public static new GetAdminCertificateArgs Empty => new GetAdminCertificateArgs();
    }

    public sealed class GetAdminCertificateInvokeArgs : global::Pulumi.InvokeArgs
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
        public Input<Pulumi.Tls.Inputs.CertRequestSubjectArgs>? Subject { get; set; }

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

        public GetAdminCertificateInvokeArgs()
        {
        }
        public static new GetAdminCertificateInvokeArgs Empty => new GetAdminCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetAdminCertificateResult
    {
        public readonly UnMango.KubernetesTheHardWay.Certificate Cert;

        [OutputConstructor]
        private GetAdminCertificateResult(UnMango.KubernetesTheHardWay.Certificate cert)
        {
            Cert = cert;
        }
    }
}
