// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tls.PrivateKey;
import com.pulumi.tls.SelfSignedCert;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.tls.RootCaArgs;
import com.unmango.kubernetesthehardway.tls.enums.Algorithm;
import com.unmango.tls.outputs.SelfSignedCertSubject;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Root certificate authority for a cluster.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:tls:RootCa")
public class RootCa extends com.pulumi.resources.ComponentResource {
    /**
     * Name of the algorithm to use when generating the private key.
     * 
     */
    @Export(name="algorithm", refs={Algorithm.class}, tree="[0]")
    private Output<Algorithm> algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key.
     * 
     */
    public Output<Algorithm> algorithm() {
        return this.algorithm;
    }
    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
     * 
     */
    @Export(name="allowedUses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedUses;

    /**
     * @return List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
     * 
     */
    public Output<List<String>> allowedUses() {
        return this.allowedUses;
    }
    /**
     * The certificate authority certificate.
     * 
     */
    @Export(name="cert", refs={SelfSignedCert.class}, tree="[0]")
    private Output<SelfSignedCert> cert;

    /**
     * @return The certificate authority certificate.
     * 
     */
    public Output<SelfSignedCert> cert() {
        return this.cert;
    }
    /**
     * Certificate data in PEM (RFC 1421).
     * 
     */
    @Export(name="certPem", refs={String.class}, tree="[0]")
    private Output<String> certPem;

    /**
     * @return Certificate data in PEM (RFC 1421).
     * 
     */
    public Output<String> certPem() {
        return this.certPem;
    }
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="dnsNames", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dnsNames;

    /**
     * @return List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> dnsNames() {
        return Codegen.optional(this.dnsNames);
    }
    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     * 
     */
    @Export(name="earlyRenewalHours", refs={Integer.class}, tree="[0]")
    private Output<Integer> earlyRenewalHours;

    /**
     * @return The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     * 
     */
    public Output<Integer> earlyRenewalHours() {
        return this.earlyRenewalHours;
    }
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     * 
     */
    @Export(name="ecdsaCurve", refs={String.class}, tree="[0]")
    private Output<String> ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     * 
     */
    public Output<String> ecdsaCurve() {
        return this.ecdsaCurve;
    }
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="ipAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ipAddresses;

    /**
     * @return List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> ipAddresses() {
        return Codegen.optional(this.ipAddresses);
    }
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    @Export(name="isCaCertificate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isCaCertificate;

    /**
     * @return Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    public Output<Boolean> isCaCertificate() {
        return this.isCaCertificate;
    }
    /**
     * The certificate authority key.
     * 
     */
    @Export(name="key", refs={PrivateKey.class}, tree="[0]")
    private Output<PrivateKey> key;

    /**
     * @return The certificate authority key.
     * 
     */
    public Output<PrivateKey> key() {
        return this.key;
    }
    /**
     * Name of the algorithm used when generating the private key provided in `private_key_pem`.
     * 
     */
    @Export(name="keyAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> keyAlgorithm;

    /**
     * @return Name of the algorithm used when generating the private key provided in `private_key_pem`.
     * 
     */
    public Output<String> keyAlgorithm() {
        return this.keyAlgorithm;
    }
    /**
     * Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
     * 
     */
    @Export(name="privateKeyOpenssh", refs={String.class}, tree="[0]")
    private Output<String> privateKeyOpenssh;

    /**
     * @return Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
     * 
     */
    public Output<String> privateKeyOpenssh() {
        return this.privateKeyOpenssh;
    }
    /**
     * Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Export(name="privateKeyPem", refs={String.class}, tree="[0]")
    private Output<String> privateKeyPem;

    /**
     * @return Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    public Output<String> privateKeyPem() {
        return this.privateKeyPem;
    }
    /**
     * Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
     * 
     */
    @Export(name="privateKeyPemPkcs8", refs={String.class}, tree="[0]")
    private Output<String> privateKeyPemPkcs8;

    /**
     * @return Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
     * 
     */
    public Output<String> privateKeyPemPkcs8() {
        return this.privateKeyPemPkcs8;
    }
    /**
     * The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    @Export(name="publicKeyFingerprintMd5", refs={String.class}, tree="[0]")
    private Output<String> publicKeyFingerprintMd5;

    /**
     * @return The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    public Output<String> publicKeyFingerprintMd5() {
        return this.publicKeyFingerprintMd5;
    }
    /**
     * The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    @Export(name="publicKeyFingerprintSha256", refs={String.class}, tree="[0]")
    private Output<String> publicKeyFingerprintSha256;

    /**
     * @return The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    public Output<String> publicKeyFingerprintSha256() {
        return this.publicKeyFingerprintSha256;
    }
    /**
     * The public key data in &#34;Authorized Keys&#34;.
     * 
     */
    @Export(name="publicKeyOpenssh", refs={String.class}, tree="[0]")
    private Output<String> publicKeyOpenssh;

    /**
     * @return The public key data in &#34;Authorized Keys&#34;.
     * 
     */
    public Output<String> publicKeyOpenssh() {
        return this.publicKeyOpenssh;
    }
    /**
     * Public key data in PEM (RFC 1421).
     * 
     */
    @Export(name="publicKeyPem", refs={String.class}, tree="[0]")
    private Output<String> publicKeyPem;

    /**
     * @return Public key data in PEM (RFC 1421).
     * 
     */
    public Output<String> publicKeyPem() {
        return this.publicKeyPem;
    }
    /**
     * Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
     * 
     */
    @Export(name="readyForRenewal", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> readyForRenewal;

    /**
     * @return Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
     * 
     */
    public Output<Boolean> readyForRenewal() {
        return this.readyForRenewal;
    }
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    @Export(name="rsaBits", refs={Integer.class}, tree="[0]")
    private Output<Integer> rsaBits;

    /**
     * @return When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    public Output<Integer> rsaBits() {
        return this.rsaBits;
    }
    /**
     * Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    @Export(name="setAuthorityKeyId", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> setAuthorityKeyId;

    /**
     * @return Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    public Output<Boolean> setAuthorityKeyId() {
        return this.setAuthorityKeyId;
    }
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    @Export(name="setSubjectKeyId", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> setSubjectKeyId;

    /**
     * @return Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    public Output<Boolean> setSubjectKeyId() {
        return this.setSubjectKeyId;
    }
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    @Export(name="subject", refs={SelfSignedCertSubject.class}, tree="[0]")
    private Output</* @Nullable */ SelfSignedCertSubject> subject;

    /**
     * @return The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    public Output<Optional<SelfSignedCertSubject>> subject() {
        return Codegen.optional(this.subject);
    }
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="uris", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> uris;

    /**
     * @return List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> uris() {
        return Codegen.optional(this.uris);
    }
    /**
     * The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    @Export(name="validityEndTime", refs={String.class}, tree="[0]")
    private Output<String> validityEndTime;

    /**
     * @return The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    public Output<String> validityEndTime() {
        return this.validityEndTime;
    }
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    @Export(name="validityPeriodHours", refs={Integer.class}, tree="[0]")
    private Output<Integer> validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    public Output<Integer> validityPeriodHours() {
        return this.validityPeriodHours;
    }
    /**
     * The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    @Export(name="validityStartTime", refs={String.class}, tree="[0]")
    private Output<String> validityStartTime;

    /**
     * @return The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    public Output<String> validityStartTime() {
        return this.validityStartTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RootCa(String name) {
        this(name, RootCaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RootCa(String name, RootCaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RootCa(String name, RootCaArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:tls:RootCa", name, args == null ? RootCaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "privateKeyOpenssh",
                "privateKeyPem",
                "privateKeyPemPkcs8"
            ))
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
