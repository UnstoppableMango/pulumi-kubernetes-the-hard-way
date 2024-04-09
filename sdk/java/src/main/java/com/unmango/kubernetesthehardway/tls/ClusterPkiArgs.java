// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tls.enums.Algorithm;
import com.unmango.kubernetesthehardway.tls.enums.EcdsaCurve;
import com.unmango.kubernetesthehardway.tls.inputs.ClusterPkiNodeArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterPkiArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterPkiArgs Empty = new ClusterPkiArgs();

    /**
     * Name of the algorithm to use when generating the private key.
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<Algorithm> algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key.
     * 
     */
    public Optional<Output<Algorithm>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * A name to use for the cluster
     * 
     */
    @Import(name="clusterName", required=true)
    private Output<String> clusterName;

    /**
     * @return A name to use for the cluster
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
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
     * Map of node name to node configuration
     * 
     */
    @Import(name="nodes", required=true)
    private Output<Map<String,ClusterPkiNodeArgs>> nodes;

    /**
     * @return Map of node name to node configuration
     * 
     */
    public Output<Map<String,ClusterPkiNodeArgs>> nodes() {
        return this.nodes;
    }

    /**
     * Publicly accessible IP address.
     * 
     */
    @Import(name="publicIp", required=true)
    private Output<String> publicIp;

    /**
     * @return Publicly accessible IP address.
     * 
     */
    public Output<String> publicIp() {
        return this.publicIp;
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
     * Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    @Import(name="validityPeriodHours")
    private @Nullable Output<Integer> validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    public Optional<Output<Integer>> validityPeriodHours() {
        return Optional.ofNullable(this.validityPeriodHours);
    }

    private ClusterPkiArgs() {}

    private ClusterPkiArgs(ClusterPkiArgs $) {
        this.algorithm = $.algorithm;
        this.clusterName = $.clusterName;
        this.ecdsaCurve = $.ecdsaCurve;
        this.nodes = $.nodes;
        this.publicIp = $.publicIp;
        this.rsaBits = $.rsaBits;
        this.validityPeriodHours = $.validityPeriodHours;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterPkiArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterPkiArgs $;

        public Builder() {
            $ = new ClusterPkiArgs();
        }

        public Builder(ClusterPkiArgs defaults) {
            $ = new ClusterPkiArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<Algorithm> algorithm) {
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
         * @param clusterName A name to use for the cluster
         * 
         * @return builder
         * 
         */
        public Builder clusterName(Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName A name to use for the cluster
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
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
         * @param nodes Map of node name to node configuration
         * 
         * @return builder
         * 
         */
        public Builder nodes(Output<Map<String,ClusterPkiNodeArgs>> nodes) {
            $.nodes = nodes;
            return this;
        }

        /**
         * @param nodes Map of node name to node configuration
         * 
         * @return builder
         * 
         */
        public Builder nodes(Map<String,ClusterPkiNodeArgs> nodes) {
            return nodes(Output.of(nodes));
        }

        /**
         * @param publicIp Publicly accessible IP address.
         * 
         * @return builder
         * 
         */
        public Builder publicIp(Output<String> publicIp) {
            $.publicIp = publicIp;
            return this;
        }

        /**
         * @param publicIp Publicly accessible IP address.
         * 
         * @return builder
         * 
         */
        public Builder publicIp(String publicIp) {
            return publicIp(Output.of(publicIp));
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
         * @param validityPeriodHours Number of hours, after initial issuing, that the certificate will remain valid.
         * 
         * @return builder
         * 
         */
        public Builder validityPeriodHours(@Nullable Output<Integer> validityPeriodHours) {
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

        public ClusterPkiArgs build() {
            $.algorithm = Codegen.objectProp("algorithm", Algorithm.class).output().arg($.algorithm).def(Algorithm.RSA).getNullable();
            if ($.clusterName == null) {
                throw new MissingRequiredPropertyException("ClusterPkiArgs", "clusterName");
            }
            if ($.nodes == null) {
                throw new MissingRequiredPropertyException("ClusterPkiArgs", "nodes");
            }
            if ($.publicIp == null) {
                throw new MissingRequiredPropertyException("ClusterPkiArgs", "publicIp");
            }
            $.rsaBits = Codegen.integerProp("rsaBits").output().arg($.rsaBits).def(2048).getNullable();
            $.validityPeriodHours = Codegen.integerProp("validityPeriodHours").output().arg($.validityPeriodHours).def(8076).getNullable();
            return $;
        }
    }

}
