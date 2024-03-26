// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.RootCa;
import com.unmango.kubernetesthehardway.enums.Algorithm;
import com.unmango.kubernetesthehardway.enums.AllowedUsage;
import com.unmango.kubernetesthehardway.enums.EcdsaCurve;
import com.unmango.kubernetesthehardway.inputs.CertRequestSubjectArgs;
import com.unmango.kubernetesthehardway.inputs.ResourceOptionsArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NewCertificateArgs extends com.pulumi.resources.InvokeArgs {

    public static final NewCertificateArgs Empty = new NewCertificateArgs();

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

    @Import(name="allowedUses", required=true)
    private Output<List<AllowedUsage>> allowedUses;

    public Output<List<AllowedUsage>> allowedUses() {
        return this.allowedUses;
    }

    /**
     * The certificate authority to issue the certificate.
     * 
     */
    @Import(name="ca", required=true)
    private Output<RootCa> ca;

    /**
     * @return The certificate authority to issue the certificate.
     * 
     */
    public Output<RootCa> ca() {
        return this.ca;
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
    private @Nullable Output<EcdsaCurve> ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     * 
     */
    public Optional<Output<EcdsaCurve>> ecdsaCurve() {
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

    @Import(name="isCaCertificate")
    private @Nullable Output<Boolean> isCaCertificate;

    public Optional<Output<Boolean>> isCaCertificate() {
        return Optional.ofNullable(this.isCaCertificate);
    }

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="options")
    private @Nullable ResourceOptionsArgs options;

    public Optional<ResourceOptionsArgs> options() {
        return Optional.ofNullable(this.options);
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
    private @Nullable Output<CertRequestSubjectArgs> subject;

    public Optional<Output<CertRequestSubjectArgs>> subject() {
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

    private NewCertificateArgs() {}

    private NewCertificateArgs(NewCertificateArgs $) {
        this.algorithm = $.algorithm;
        this.allowedUses = $.allowedUses;
        this.ca = $.ca;
        this.dnsNames = $.dnsNames;
        this.earlyRenewalHours = $.earlyRenewalHours;
        this.ecdsaCurve = $.ecdsaCurve;
        this.ipAddresses = $.ipAddresses;
        this.isCaCertificate = $.isCaCertificate;
        this.name = $.name;
        this.options = $.options;
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
    public static Builder builder(NewCertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NewCertificateArgs $;

        public Builder() {
            $ = new NewCertificateArgs();
        }

        public Builder(NewCertificateArgs defaults) {
            $ = new NewCertificateArgs(Objects.requireNonNull(defaults));
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

        public Builder allowedUses(Output<List<AllowedUsage>> allowedUses) {
            $.allowedUses = allowedUses;
            return this;
        }

        public Builder allowedUses(List<AllowedUsage> allowedUses) {
            return allowedUses(Output.of(allowedUses));
        }

        public Builder allowedUses(AllowedUsage... allowedUses) {
            return allowedUses(List.of(allowedUses));
        }

        /**
         * @param ca The certificate authority to issue the certificate.
         * 
         * @return builder
         * 
         */
        public Builder ca(Output<RootCa> ca) {
            $.ca = ca;
            return this;
        }

        /**
         * @param ca The certificate authority to issue the certificate.
         * 
         * @return builder
         * 
         */
        public Builder ca(RootCa ca) {
            return ca(Output.of(ca));
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
        public Builder ecdsaCurve(@Nullable Output<EcdsaCurve> ecdsaCurve) {
            $.ecdsaCurve = ecdsaCurve;
            return this;
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(EcdsaCurve ecdsaCurve) {
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

        public Builder isCaCertificate(@Nullable Output<Boolean> isCaCertificate) {
            $.isCaCertificate = isCaCertificate;
            return this;
        }

        public Builder isCaCertificate(Boolean isCaCertificate) {
            return isCaCertificate(Output.of(isCaCertificate));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder options(@Nullable ResourceOptionsArgs options) {
            $.options = options;
            return this;
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

        public Builder subject(@Nullable Output<CertRequestSubjectArgs> subject) {
            $.subject = subject;
            return this;
        }

        public Builder subject(CertRequestSubjectArgs subject) {
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

        public NewCertificateArgs build() {
            if ($.algorithm == null) {
                throw new MissingRequiredPropertyException("NewCertificateArgs", "algorithm");
            }
            if ($.allowedUses == null) {
                throw new MissingRequiredPropertyException("NewCertificateArgs", "allowedUses");
            }
            if ($.ca == null) {
                throw new MissingRequiredPropertyException("NewCertificateArgs", "ca");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("NewCertificateArgs", "name");
            }
            if ($.validityPeriodHours == null) {
                throw new MissingRequiredPropertyException("NewCertificateArgs", "validityPeriodHours");
            }
            return $;
        }
    }

}