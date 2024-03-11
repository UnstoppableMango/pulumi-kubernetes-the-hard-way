// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.tls.inputs.SelfSignedCertSubjectArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RootCaArgs extends com.pulumi.resources.ResourceArgs {

    public static final RootCaArgs Empty = new RootCaArgs();

    /**
     * Name of the algorithm to use when generating the private key.
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key.
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * List of DNS names for which a certificate is being requested.
     * 
     */
    @Import(name="dnsNames")
    private @Nullable Output<List<String>> dnsNames;

    /**
     * @return List of DNS names for which a certificate is being requested.
     * 
     */
    public Optional<Output<List<String>>> dnsNames() {
        return Optional.ofNullable(this.dnsNames);
    }

    /**
     * TODO
     * 
     */
    @Import(name="earlyRenewalHours")
    private @Nullable Output<Integer> earlyRenewalHours;

    /**
     * @return TODO
     * 
     */
    public Optional<Output<Integer>> earlyRenewalHours() {
        return Optional.ofNullable(this.earlyRenewalHours);
    }

    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     * 
     */
    @Import(name="ecdsaCurve")
    private @Nullable Output<String> ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     * 
     */
    public Optional<Output<String>> ecdsaCurve() {
        return Optional.ofNullable(this.ecdsaCurve);
    }

    /**
     * List of IP addresses for which a certificate is being requested.
     * 
     */
    @Import(name="ipAddresses")
    private @Nullable Output<List<String>> ipAddresses;

    /**
     * @return List of IP addresses for which a certificate is being requested.
     * 
     */
    public Optional<Output<List<String>>> ipAddresses() {
        return Optional.ofNullable(this.ipAddresses);
    }

    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     * 
     */
    @Import(name="rsaBits")
    private @Nullable Output<Integer> rsaBits;

    /**
     * @return When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     * 
     */
    public Optional<Output<Integer>> rsaBits() {
        return Optional.ofNullable(this.rsaBits);
    }

    /**
     * Should the generated certificate include an authority key identifier.
     * 
     */
    @Import(name="setAuthorityKeyId")
    private @Nullable Output<Boolean> setAuthorityKeyId;

    /**
     * @return Should the generated certificate include an authority key identifier.
     * 
     */
    public Optional<Output<Boolean>> setAuthorityKeyId() {
        return Optional.ofNullable(this.setAuthorityKeyId);
    }

    /**
     * Should the generated certificate include a subject key identifier.
     * 
     */
    @Import(name="setSubjectKeyId")
    private @Nullable Output<Boolean> setSubjectKeyId;

    /**
     * @return Should the generated certificate include a subject key identifier.
     * 
     */
    public Optional<Output<Boolean>> setSubjectKeyId() {
        return Optional.ofNullable(this.setSubjectKeyId);
    }

    @Import(name="subject")
    private @Nullable Output<SelfSignedCertSubjectArgs> subject;

    public Optional<Output<SelfSignedCertSubjectArgs>> subject() {
        return Optional.ofNullable(this.subject);
    }

    /**
     * List of URIs for which a certificate is being requested.
     * 
     */
    @Import(name="uris")
    private @Nullable Output<List<String>> uris;

    /**
     * @return List of URIs for which a certificate is being requested.
     * 
     */
    public Optional<Output<List<String>>> uris() {
        return Optional.ofNullable(this.uris);
    }

    /**
     * Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    @Import(name="validityPeriodHours", required=true)
    private Output<Integer> validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    public Output<Integer> validityPeriodHours() {
        return this.validityPeriodHours;
    }

    private RootCaArgs() {}

    private RootCaArgs(RootCaArgs $) {
        this.algorithm = $.algorithm;
        this.dnsNames = $.dnsNames;
        this.earlyRenewalHours = $.earlyRenewalHours;
        this.ecdsaCurve = $.ecdsaCurve;
        this.ipAddresses = $.ipAddresses;
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
    public static Builder builder(RootCaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RootCaArgs $;

        public Builder() {
            $ = new RootCaArgs();
        }

        public Builder(RootCaArgs defaults) {
            $ = new RootCaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(@Nullable Output<List<String>> dnsNames) {
            $.dnsNames = dnsNames;
            return this;
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(List<String> dnsNames) {
            return dnsNames(Output.of(dnsNames));
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
        public Builder earlyRenewalHours(@Nullable Output<Integer> earlyRenewalHours) {
            $.earlyRenewalHours = earlyRenewalHours;
            return this;
        }

        /**
         * @param earlyRenewalHours TODO
         * 
         * @return builder
         * 
         */
        public Builder earlyRenewalHours(Integer earlyRenewalHours) {
            return earlyRenewalHours(Output.of(earlyRenewalHours));
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(@Nullable Output<String> ecdsaCurve) {
            $.ecdsaCurve = ecdsaCurve;
            return this;
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(String ecdsaCurve) {
            return ecdsaCurve(Output.of(ecdsaCurve));
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(@Nullable Output<List<String>> ipAddresses) {
            $.ipAddresses = ipAddresses;
            return this;
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(List<String> ipAddresses) {
            return ipAddresses(Output.of(ipAddresses));
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

        /**
         * @param rsaBits When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
         * 
         * @return builder
         * 
         */
        public Builder rsaBits(@Nullable Output<Integer> rsaBits) {
            $.rsaBits = rsaBits;
            return this;
        }

        /**
         * @param rsaBits When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
         * 
         * @return builder
         * 
         */
        public Builder rsaBits(Integer rsaBits) {
            return rsaBits(Output.of(rsaBits));
        }

        /**
         * @param setAuthorityKeyId Should the generated certificate include an authority key identifier.
         * 
         * @return builder
         * 
         */
        public Builder setAuthorityKeyId(@Nullable Output<Boolean> setAuthorityKeyId) {
            $.setAuthorityKeyId = setAuthorityKeyId;
            return this;
        }

        /**
         * @param setAuthorityKeyId Should the generated certificate include an authority key identifier.
         * 
         * @return builder
         * 
         */
        public Builder setAuthorityKeyId(Boolean setAuthorityKeyId) {
            return setAuthorityKeyId(Output.of(setAuthorityKeyId));
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a subject key identifier.
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(@Nullable Output<Boolean> setSubjectKeyId) {
            $.setSubjectKeyId = setSubjectKeyId;
            return this;
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a subject key identifier.
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(Boolean setSubjectKeyId) {
            return setSubjectKeyId(Output.of(setSubjectKeyId));
        }

        public Builder subject(@Nullable Output<SelfSignedCertSubjectArgs> subject) {
            $.subject = subject;
            return this;
        }

        public Builder subject(SelfSignedCertSubjectArgs subject) {
            return subject(Output.of(subject));
        }

        /**
         * @param uris List of URIs for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder uris(@Nullable Output<List<String>> uris) {
            $.uris = uris;
            return this;
        }

        /**
         * @param uris List of URIs for which a certificate is being requested.
         * 
         * @return builder
         * 
         */
        public Builder uris(List<String> uris) {
            return uris(Output.of(uris));
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
        public Builder validityPeriodHours(Output<Integer> validityPeriodHours) {
            $.validityPeriodHours = validityPeriodHours;
            return this;
        }

        /**
         * @param validityPeriodHours Number of hours, after initial issuing, that the certificate will remain valid.
         * 
         * @return builder
         * 
         */
        public Builder validityPeriodHours(Integer validityPeriodHours) {
            return validityPeriodHours(Output.of(validityPeriodHours));
        }

        public RootCaArgs build() {
            if ($.validityPeriodHours == null) {
                throw new MissingRequiredPropertyException("RootCaArgs", "validityPeriodHours");
            }
            return $;
        }
    }

}
