// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Tls
{
    /// <summary>
    /// A certificate key pair.
    /// </summary>
    [KubernetesTheHardWayResourceType("kubernetes-the-hard-way:tls:Certificate")]
    public partial class Certificate : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Output("algorithm")]
        public Output<UnMango.KubernetesTheHardWay.Tls.Algorithm> Algorithm { get; private set; } = null!;

        /// <summary>
        /// List of key usages allowed for the issued certificate.
        /// </summary>
        [Output("allowedUses")]
        public Output<ImmutableArray<UnMango.KubernetesTheHardWay.Tls.AllowedUsage>> AllowedUses { get; private set; } = null!;

        /// <summary>
        /// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        [Output("caCertPem")]
        public Output<string> CaCertPem { get; private set; } = null!;

        /// <summary>
        /// Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
        /// </summary>
        [Output("caKeyAlgorithm")]
        public Output<string> CaKeyAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        [Output("caPrivateKeyPem")]
        public Output<string> CaPrivateKeyPem { get; private set; } = null!;

        /// <summary>
        /// The certificate.
        /// </summary>
        [Output("cert")]
        public Output<Pulumi.Tls.LocallySignedCert> Cert { get; private set; } = null!;

        /// <summary>
        /// Certificate data in PEM (RFC 1421).
        /// </summary>
        [Output("certPem")]
        public Output<string> CertPem { get; private set; } = null!;

        /// <summary>
        /// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        [Output("certRequestPem")]
        public Output<string> CertRequestPem { get; private set; } = null!;

        /// <summary>
        /// The certificate signing request.
        /// </summary>
        [Output("csr")]
        public Output<Pulumi.Tls.CertRequest> Csr { get; private set; } = null!;

        /// <summary>
        /// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("dnsNames")]
        public Output<ImmutableArray<string>> DnsNames { get; private set; } = null!;

        [Output("earlyRenewalHours")]
        public Output<int> EarlyRenewalHours { get; private set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Output("ecdsaCurve")]
        public Output<UnMango.KubernetesTheHardWay.Tls.EcdsaCurve> EcdsaCurve { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<string>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        /// </summary>
        [Output("isCaCertificate")]
        public Output<bool> IsCaCertificate { get; private set; } = null!;

        /// <summary>
        /// The private key
        /// </summary>
        [Output("key")]
        public Output<Pulumi.Tls.PrivateKey> Key { get; private set; } = null!;

        /// <summary>
        /// Name of the algorithm used when generating the private key provided in `private_key_pem`.
        /// </summary>
        [Output("keyAlgorithm")]
        public Output<string> KeyAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
        /// </summary>
        [Output("privateKeyOpenssh")]
        public Output<string> PrivateKeyOpenssh { get; private set; } = null!;

        /// <summary>
        /// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        [Output("privateKeyPem")]
        public Output<string> PrivateKeyPem { get; private set; } = null!;

        /// <summary>
        /// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
        /// </summary>
        [Output("privateKeyPemPkcs8")]
        public Output<string> PrivateKeyPemPkcs8 { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        /// </summary>
        [Output("publicKeyFingerprintMd5")]
        public Output<string> PublicKeyFingerprintMd5 { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        /// </summary>
        [Output("publicKeyFingerprintSha256")]
        public Output<string> PublicKeyFingerprintSha256 { get; private set; } = null!;

        /// <summary>
        /// The public key data in "Authorized Keys".
        /// </summary>
        [Output("publicKeyOpenssh")]
        public Output<string> PublicKeyOpenssh { get; private set; } = null!;

        /// <summary>
        /// Public key data in PEM (RFC 1421).
        /// </summary>
        [Output("publicKeyPem")]
        public Output<string> PublicKeyPem { get; private set; } = null!;

        /// <summary>
        /// Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
        /// </summary>
        [Output("readyForRenewal")]
        public Output<bool> ReadyForRenewal { get; private set; } = null!;

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        /// </summary>
        [Output("rsaBits")]
        public Output<int> RsaBits { get; private set; } = null!;

        /// <summary>
        /// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Output("setSubjectKeyId")]
        public Output<bool> SetSubjectKeyId { get; private set; } = null!;

        /// <summary>
        /// TODO
        /// </summary>
        [Output("subject")]
        public Output<Pulumi.Tls.Outputs.CertRequestSubject?> Subject { get; private set; } = null!;

        /// <summary>
        /// List of URIs for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("uris")]
        public Output<ImmutableArray<string>> Uris { get; private set; } = null!;

        /// <summary>
        /// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        [Output("validityEndTime")]
        public Output<string> ValidityEndTime { get; private set; } = null!;

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid for.
        /// </summary>
        [Output("validityPeriodHours")]
        public Output<int> ValidityPeriodHours { get; private set; } = null!;

        /// <summary>
        /// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        [Output("validityStartTime")]
        public Output<string> ValidityStartTime { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, ComponentResourceOptions? options = null)
            : base("kubernetes-the-hard-way:tls:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<UnMango.KubernetesTheHardWay.Tls.Algorithm> Algorithm { get; set; } = null!;

        [Input("allowedUses", required: true)]
        private InputList<UnMango.KubernetesTheHardWay.Tls.AllowedUsage>? _allowedUses;

        /// <summary>
        /// List of key usages allowed for the issued certificate.
        /// </summary>
        public InputList<UnMango.KubernetesTheHardWay.Tls.AllowedUsage> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<UnMango.KubernetesTheHardWay.Tls.AllowedUsage>());
            set => _allowedUses = value;
        }

        /// <summary>
        /// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        [Input("caCertPem", required: true)]
        public Input<string> CaCertPem { get; set; } = null!;

        [Input("caPrivateKeyPem", required: true)]
        private Input<string>? _caPrivateKeyPem;

        /// <summary>
        /// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        public Input<string>? CaPrivateKeyPem
        {
            get => _caPrivateKeyPem;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _caPrivateKeyPem = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        [Input("earlyRenewalHours")]
        public Input<int>? EarlyRenewalHours { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<UnMango.KubernetesTheHardWay.Tls.EcdsaCurve>? EcdsaCurve { get; set; }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        /// </summary>
        [Input("isCaCertificate")]
        public Input<bool>? IsCaCertificate { get; set; }

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        /// <summary>
        /// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Input("setSubjectKeyId")]
        public Input<bool>? SetSubjectKeyId { get; set; }

        /// <summary>
        /// TODO
        /// </summary>
        [Input("subject")]
        public Input<Pulumi.Tls.Inputs.CertRequestSubjectArgs>? Subject { get; set; }

        [Input("uris")]
        private InputList<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid for.
        /// </summary>
        [Input("validityPeriodHours", required: true)]
        public Input<int> ValidityPeriodHours { get; set; } = null!;

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}
