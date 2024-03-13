// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.tls.inputs.CertRequestSubject;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAdminCertificatePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAdminCertificatePlainArgs Empty = new GetAdminCertificatePlainArgs();

    /**
     * Name of the algorithm to use when generating the private key.
     * 
     */
    @Import(name="algorithm")
    private @Nullable String algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key.
     * 
     */
    public Optional<String> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    @Import(name="allowedUses", required=true)
    private List<String> allowedUses;

    public List<String> allowedUses() {
        return this.allowedUses;
    }

    /**
     * List of DNS names for which a certificate is being requested.
     * 
     */
    @Import(name="dnsNames")
    private @Nullable List<String> dnsNames;

    /**
     * @return List of DNS names for which a certificate is being requested.
     * 
     */
    public Optional<List<String>> dnsNames() {
        return Optional.ofNullable(this.dnsNames);
    }

    /**
     * TODO
     * 
     */
    @Import(name="earlyRenewalHours")
    private @Nullable Integer earlyRenewalHours;

    /**
     * @return TODO
     * 
     */
    public Optional<Integer> earlyRenewalHours() {
        return Optional.ofNullable(this.earlyRenewalHours);
    }

    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     * 
     */
    @Import(name="ecdsaCurve")
    private @Nullable String ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     * 
     */
    public Optional<String> ecdsaCurve() {
        return Optional.ofNullable(this.ecdsaCurve);
    }

    /**
     * List of IP addresses for which a certificate is being requested.
     * 
     */
    @Import(name="ipAddresses")
    private @Nullable List<String> ipAddresses;

    /**
     * @return List of IP addresses for which a certificate is being requested.
     * 
     */
    public Optional<List<String>> ipAddresses() {
        return Optional.ofNullable(this.ipAddresses);
    }

    @Import(name="isCaCertificate")
    private @Nullable Boolean isCaCertificate;

    public Optional<Boolean> isCaCertificate() {
        return Optional.ofNullable(this.isCaCertificate);
    }

    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     * 
     */
    @Import(name="rsaBits")
    private @Nullable Integer rsaBits;

    /**
     * @return When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     * 
     */
    public Optional<Integer> rsaBits() {
        return Optional.ofNullable(this.rsaBits);
    }

    /**
     * Should the generated certificate include an authority key identifier.
     * 
     */
    @Import(name="setAuthorityKeyId")
    private @Nullable Boolean setAuthorityKeyId;

    /**
     * @return Should the generated certificate include an authority key identifier.
     * 
     */
    public Optional<Boolean> setAuthorityKeyId() {
        return Optional.ofNullable(this.setAuthorityKeyId);
    }

    /**
     * Should the generated certificate include a subject key identifier.
     * 
     */
    @Import(name="setSubjectKeyId")
    private @Nullable Boolean setSubjectKeyId;

    /**
     * @return Should the generated certificate include a subject key identifier.
     * 
     */
    public Optional<Boolean> setSubjectKeyId() {
        return Optional.ofNullable(this.setSubjectKeyId);
    }

    @Import(name="subject")
    private @Nullable CertRequestSubject subject;

    public Optional<CertRequestSubject> subject() {
        return Optional.ofNullable(this.subject);
    }

    /**
     * List of URIs for which a certificate is being requested.
     * 
     */
    @Import(name="uris")
    private @Nullable List<String> uris;

    /**
     * @return List of URIs for which a certificate is being requested.
     * 
     */
    public Optional<List<String>> uris() {
        return Optional.ofNullable(this.uris);
    }

    /**
     * Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    @Import(name="validityPeriodHours", required=true)
    private Integer validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    public Integer validityPeriodHours() {
        return this.validityPeriodHours;
    }

    private GetAdminCertificatePlainArgs() {}

    private GetAdminCertificatePlainArgs(GetAdminCertificatePlainArgs $) {
        this.algorithm = $.algorithm;
        this.allowedUses = $.allowedUses;
        this.dnsNames = $.dnsNames;
        this.earlyRenewalHours = $.earlyRenewalHours;
        this.ecdsaCurve = $.ecdsaCurve;
        this.ipAddresses = $.ipAddresses;
        this.isCaCertificate = $.isCaCertificate;
        this.rsaBits = $.rsaBits;
        this.setAuthorityKeyId = $.setAuthorityKeyId;
        this.setSubjectKeyId = $.setSubjectKeyId;
        this.subject = $.subject;
        this.uris = $.uris;
        this.validityPeriodHours = $.validityPeriodHours;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAdminCertificatePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAdminCertificatePlainArgs $;

        public Builder() {
            $ = new GetAdminCertificatePlainArgs();
        }

        public Builder(GetAdminCertificatePlainArgs defaults) {
            $ = new GetAdminCertificatePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable String algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        public Builder allowedUses(List<String> allowedUses) {
            $.allowedUses = allowedUses;
            return this;
        }

        public Builder allowedUses(String... allowedUses) {
            return allowedUses(List.of(allowedUses));
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(@Nullable List<String> dnsNames) {
            $.dnsNames = dnsNames;
            return this;
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(String... dnsNames) {
            return dnsNames(List.of(dnsNames));
        }

        /**
         * @param earlyRenewalHours TODO
         * 
         * @return builder
         * 
         */
        public Builder earlyRenewalHours(@Nullable Integer earlyRenewalHours) {
            $.earlyRenewalHours = earlyRenewalHours;
            return this;
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(@Nullable String ecdsaCurve) {
            $.ecdsaCurve = ecdsaCurve;
            return this;
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(@Nullable List<String> ipAddresses) {
            $.ipAddresses = ipAddresses;
            return this;
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(String... ipAddresses) {
            return ipAddresses(List.of(ipAddresses));
        }

        public Builder isCaCertificate(@Nullable Boolean isCaCertificate) {
            $.isCaCertificate = isCaCertificate;
            return this;
        }

        /**
         * @param rsaBits When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
         * 
         * @return builder
         * 
         */
        public Builder rsaBits(@Nullable Integer rsaBits) {
            $.rsaBits = rsaBits;
            return this;
        }

        /**
         * @param setAuthorityKeyId Should the generated certificate include an authority key identifier.
         * 
         * @return builder
         * 
         */
        public Builder setAuthorityKeyId(@Nullable Boolean setAuthorityKeyId) {
            $.setAuthorityKeyId = setAuthorityKeyId;
            return this;
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a subject key identifier.
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(@Nullable Boolean setSubjectKeyId) {
            $.setSubjectKeyId = setSubjectKeyId;
            return this;
        }

        public Builder subject(@Nullable CertRequestSubject subject) {
            $.subject = subject;
            return this;
        }

        /**
         * @param uris List of URIs for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder uris(@Nullable List<String> uris) {
            $.uris = uris;
            return this;
        }

        /**
         * @param uris List of URIs for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder uris(String... uris) {
            return uris(List.of(uris));
        }

        /**
         * @param validityPeriodHours Number of hours, after initial issuing, that the certificate will remain valid.
         * 
         * @return builder
         * 
         */
        public Builder validityPeriodHours(Integer validityPeriodHours) {
            $.validityPeriodHours = validityPeriodHours;
            return this;
        }

        public GetAdminCertificatePlainArgs build() {
            if ($.allowedUses == null) {
                throw new MissingRequiredPropertyException("GetAdminCertificatePlainArgs", "allowedUses");
            }
            if ($.validityPeriodHours == null) {
                throw new MissingRequiredPropertyException("GetAdminCertificatePlainArgs", "validityPeriodHours");
            }
            return $;
        }
    }

}
