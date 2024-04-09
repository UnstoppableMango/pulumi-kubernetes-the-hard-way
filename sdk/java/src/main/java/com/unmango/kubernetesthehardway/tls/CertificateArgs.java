// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.tls.inputs.CertRequestSubjectArgs;
import com.unmango.kubernetesthehardway.tls.enums.Algorithm;
import com.unmango.kubernetesthehardway.tls.enums.AllowedUsage;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertificateArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertificateArgs Empty = new CertificateArgs();

    /**
     * Name of the algorithm to use when generating the private key.
     * 
     */
    @Import(name="algorithm", required=true)
    private Output<Algorithm> algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key.
     * 
     */
    public Output<Algorithm> algorithm() {
        return this.algorithm;
    }

    /**
     * List of key usages allowed for the issued certificate.
     * 
     */
    @Import(name="allowedUses", required=true)
    private Output<List<AllowedUsage>> allowedUses;

    /**
     * @return List of key usages allowed for the issued certificate.
     * 
     */
    public Output<List<AllowedUsage>> allowedUses() {
        return this.allowedUses;
    }

    /**
     * Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Import(name="caCertPem", required=true)
    private Output<String> caCertPem;

    /**
     * @return Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    public Output<String> caCertPem() {
        return this.caCertPem;
    }

    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Import(name="caPrivateKeyPem", required=true)
    private Output<String> caPrivateKeyPem;

    /**
     * @return Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    public Output<String> caPrivateKeyPem() {
        return this.caPrivateKeyPem;
    }

    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Import(name="dnsNames")
    private @Nullable Output<List<String>> dnsNames;

    /**
     * @return List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Optional<Output<List<String>>> dnsNames() {
        return Optional.ofNullable(this.dnsNames);
    }

    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     * 
     */
    @Import(name="earlyRenewalHours")
    private @Nullable Output<Integer> earlyRenewalHours;

    /**
     * @return The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     * 
     */
    public Optional<Output<Integer>> earlyRenewalHours() {
        return Optional.ofNullable(this.earlyRenewalHours);
    }

    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     * 
     */
    @Import(name="ecdsaCurve")
    private @Nullable Output<String> ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     * 
     */
    public Optional<Output<String>> ecdsaCurve() {
        return Optional.ofNullable(this.ecdsaCurve);
    }

    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Import(name="ipAddresses")
    private @Nullable Output<List<String>> ipAddresses;

    /**
     * @return List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Optional<Output<List<String>>> ipAddresses() {
        return Optional.ofNullable(this.ipAddresses);
    }

    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    @Import(name="isCaCertificate")
    private @Nullable Output<Boolean> isCaCertificate;

    /**
     * @return Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    public Optional<Output<Boolean>> isCaCertificate() {
        return Optional.ofNullable(this.isCaCertificate);
    }

    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    @Import(name="rsaBits")
    private @Nullable Output<Integer> rsaBits;

    /**
     * @return When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    public Optional<Output<Integer>> rsaBits() {
        return Optional.ofNullable(this.rsaBits);
    }

    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    @Import(name="setSubjectKeyId")
    private @Nullable Output<Boolean> setSubjectKeyId;

    /**
     * @return Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    public Optional<Output<Boolean>> setSubjectKeyId() {
        return Optional.ofNullable(this.setSubjectKeyId);
    }

    /**
     * TODO
     * 
     */
    @Import(name="subject")
    private @Nullable Output<CertRequestSubjectArgs> subject;

    /**
     * @return TODO
     * 
     */
    public Optional<Output<CertRequestSubjectArgs>> subject() {
        return Optional.ofNullable(this.subject);
    }

    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Import(name="uris")
    private @Nullable Output<List<String>> uris;

    /**
     * @return List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Optional<Output<List<String>>> uris() {
        return Optional.ofNullable(this.uris);
    }

    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    @Import(name="validityPeriodHours", required=true)
    private Output<Integer> validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    public Output<Integer> validityPeriodHours() {
        return this.validityPeriodHours;
    }

    private CertificateArgs() {}

    private CertificateArgs(CertificateArgs $) {
        this.algorithm = $.algorithm;
        this.allowedUses = $.allowedUses;
        this.caCertPem = $.caCertPem;
        this.caPrivateKeyPem = $.caPrivateKeyPem;
        this.dnsNames = $.dnsNames;
        this.earlyRenewalHours = $.earlyRenewalHours;
        this.ecdsaCurve = $.ecdsaCurve;
        this.ipAddresses = $.ipAddresses;
        this.isCaCertificate = $.isCaCertificate;
        this.rsaBits = $.rsaBits;
        this.setSubjectKeyId = $.setSubjectKeyId;
        this.subject = $.subject;
        this.uris = $.uris;
        this.validityPeriodHours = $.validityPeriodHours;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateArgs $;

        public Builder() {
            $ = new CertificateArgs();
        }

        public Builder(CertificateArgs defaults) {
            $ = new CertificateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(Output<Algorithm> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(Algorithm algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(Output<List<AllowedUsage>> allowedUses) {
            $.allowedUses = allowedUses;
            return this;
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(List<AllowedUsage> allowedUses) {
            return allowedUses(Output.of(allowedUses));
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(AllowedUsage... allowedUses) {
            return allowedUses(List.of(allowedUses));
        }

        /**
         * @param caCertPem Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder caCertPem(Output<String> caCertPem) {
            $.caCertPem = caCertPem;
            return this;
        }

        /**
         * @param caCertPem Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder caCertPem(String caCertPem) {
            return caCertPem(Output.of(caCertPem));
        }

        /**
         * @param caPrivateKeyPem Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder caPrivateKeyPem(Output<String> caPrivateKeyPem) {
            $.caPrivateKeyPem = caPrivateKeyPem;
            return this;
        }

        /**
         * @param caPrivateKeyPem Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder caPrivateKeyPem(String caPrivateKeyPem) {
            return caPrivateKeyPem(Output.of(caPrivateKeyPem));
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(@Nullable Output<List<String>> dnsNames) {
            $.dnsNames = dnsNames;
            return this;
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(List<String> dnsNames) {
            return dnsNames(Output.of(dnsNames));
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(String... dnsNames) {
            return dnsNames(List.of(dnsNames));
        }

        /**
         * @param earlyRenewalHours The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
         * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
         * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
         * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
         * early renewal period. (default: `0`)
         * 
         * @return builder
         * 
         */
        public Builder earlyRenewalHours(@Nullable Output<Integer> earlyRenewalHours) {
            $.earlyRenewalHours = earlyRenewalHours;
            return this;
        }

        /**
         * @param earlyRenewalHours The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
         * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
         * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
         * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
         * early renewal period. (default: `0`)
         * 
         * @return builder
         * 
         */
        public Builder earlyRenewalHours(Integer earlyRenewalHours) {
            return earlyRenewalHours(Output.of(earlyRenewalHours));
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(@Nullable Output<String> ecdsaCurve) {
            $.ecdsaCurve = ecdsaCurve;
            return this;
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(String ecdsaCurve) {
            return ecdsaCurve(Output.of(ecdsaCurve));
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(@Nullable Output<List<String>> ipAddresses) {
            $.ipAddresses = ipAddresses;
            return this;
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(List<String> ipAddresses) {
            return ipAddresses(Output.of(ipAddresses));
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(String... ipAddresses) {
            return ipAddresses(List.of(ipAddresses));
        }

        /**
         * @param isCaCertificate Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder isCaCertificate(@Nullable Output<Boolean> isCaCertificate) {
            $.isCaCertificate = isCaCertificate;
            return this;
        }

        /**
         * @param isCaCertificate Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder isCaCertificate(Boolean isCaCertificate) {
            return isCaCertificate(Output.of(isCaCertificate));
        }

        /**
         * @param rsaBits When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
         * 
         * @return builder
         * 
         */
        public Builder rsaBits(@Nullable Output<Integer> rsaBits) {
            $.rsaBits = rsaBits;
            return this;
        }

        /**
         * @param rsaBits When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
         * 
         * @return builder
         * 
         */
        public Builder rsaBits(Integer rsaBits) {
            return rsaBits(Output.of(rsaBits));
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(@Nullable Output<Boolean> setSubjectKeyId) {
            $.setSubjectKeyId = setSubjectKeyId;
            return this;
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(Boolean setSubjectKeyId) {
            return setSubjectKeyId(Output.of(setSubjectKeyId));
        }

        /**
         * @param subject TODO
         * 
         * @return builder
         * 
         */
        public Builder subject(@Nullable Output<CertRequestSubjectArgs> subject) {
            $.subject = subject;
            return this;
        }

        /**
         * @param subject TODO
         * 
         * @return builder
         * 
         */
        public Builder subject(CertRequestSubjectArgs subject) {
            return subject(Output.of(subject));
        }

        /**
         * @param uris List of URIs for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder uris(@Nullable Output<List<String>> uris) {
            $.uris = uris;
            return this;
        }

        /**
         * @param uris List of URIs for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder uris(List<String> uris) {
            return uris(Output.of(uris));
        }

        /**
         * @param uris List of URIs for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder uris(String... uris) {
            return uris(List.of(uris));
        }

        /**
         * @param validityPeriodHours Number of hours, after initial issuing, that the certificate will remain valid for.
         * 
         * @return builder
         * 
         */
        public Builder validityPeriodHours(Output<Integer> validityPeriodHours) {
            $.validityPeriodHours = validityPeriodHours;
            return this;
        }

        /**
         * @param validityPeriodHours Number of hours, after initial issuing, that the certificate will remain valid for.
         * 
         * @return builder
         * 
         */
        public Builder validityPeriodHours(Integer validityPeriodHours) {
            return validityPeriodHours(Output.of(validityPeriodHours));
        }

        public CertificateArgs build() {
            if ($.algorithm == null) {
                throw new MissingRequiredPropertyException("CertificateArgs", "algorithm");
            }
            if ($.allowedUses == null) {
                throw new MissingRequiredPropertyException("CertificateArgs", "allowedUses");
            }
            if ($.caCertPem == null) {
                throw new MissingRequiredPropertyException("CertificateArgs", "caCertPem");
            }
            if ($.caPrivateKeyPem == null) {
                throw new MissingRequiredPropertyException("CertificateArgs", "caPrivateKeyPem");
            }
            if ($.validityPeriodHours == null) {
                throw new MissingRequiredPropertyException("CertificateArgs", "validityPeriodHours");
            }
            return $;
        }
    }

}
