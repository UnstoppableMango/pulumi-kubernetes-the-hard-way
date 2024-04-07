// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.tls.ClusterPkiArgs;
import com.unmango.kubernetesthehardway.tls.enums.Algorithm;
import com.unmango.kubernetesthehardway.tls.enums.EcdsaCurve;
import com.unmango.kubernetesthehardway.tls.outputs.Certificate;
import com.unmango.kubernetesthehardway.tls.outputs.ClusterPkiNode;
import com.unmango.kubernetesthehardway.tls.outputs.RootCa;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The private key infrastructure for a cluster
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:tls:ClusterPki")
public class ClusterPki extends com.pulumi.resources.ComponentResource {
    /**
     * The admin certificate.
     * 
     */
    @Export(name="admin", refs={Certificate.class}, tree="[0]")
    private Output<Certificate> admin;

    /**
     * @return The admin certificate.
     * 
     */
    public Output<Certificate> admin() {
        return this.admin;
    }
    /**
     * Name of the algorithm to use when generating the private key.
     * 
     */
    @Export(name="algorithm", refs={Algorithm.class}, tree="[0]")
    private Output</* @Nullable */ Algorithm> algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key.
     * 
     */
    public Output<Optional<Algorithm>> algorithm() {
        return Codegen.optional(this.algorithm);
    }
    /**
     * The cluster certificate authority.
     * 
     */
    @Export(name="ca", refs={RootCa.class}, tree="[0]")
    private Output<RootCa> ca;

    /**
     * @return The cluster certificate authority.
     * 
     */
    public Output<RootCa> ca() {
        return this.ca;
    }
    /**
     * A name to use for the cluster
     * 
     */
    @Export(name="clusterName", refs={String.class}, tree="[0]")
    private Output<String> clusterName;

    /**
     * @return A name to use for the cluster
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * The controller manager certificate.
     * 
     */
    @Export(name="controllerManager", refs={Certificate.class}, tree="[0]")
    private Output<Certificate> controllerManager;

    /**
     * @return The controller manager certificate.
     * 
     */
    public Output<Certificate> controllerManager() {
        return this.controllerManager;
    }
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     * 
     */
    @Export(name="ecdsaCurve", refs={EcdsaCurve.class}, tree="[0]")
    private Output</* @Nullable */ EcdsaCurve> ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     * 
     */
    public Output<Optional<EcdsaCurve>> ecdsaCurve() {
        return Codegen.optional(this.ecdsaCurve);
    }
    /**
     * The kube proxy certificate.
     * 
     */
    @Export(name="kubeProxy", refs={Certificate.class}, tree="[0]")
    private Output<Certificate> kubeProxy;

    /**
     * @return The kube proxy certificate.
     * 
     */
    public Output<Certificate> kubeProxy() {
        return this.kubeProxy;
    }
    /**
     * The kube scheduler certificate.
     * 
     */
    @Export(name="kubeScheduler", refs={Certificate.class}, tree="[0]")
    private Output<Certificate> kubeScheduler;

    /**
     * @return The kube scheduler certificate.
     * 
     */
    public Output<Certificate> kubeScheduler() {
        return this.kubeScheduler;
    }
    /**
     * Map of node name to kubelet certificate.
     * 
     */
    @Export(name="kubelet", refs={Map.class,String.class,Certificate.class}, tree="[0,1,2]")
    private Output<Map<String,Certificate>> kubelet;

    /**
     * @return Map of node name to kubelet certificate.
     * 
     */
    public Output<Map<String,Certificate>> kubelet() {
        return this.kubelet;
    }
    /**
     * The kubernetes certificate.
     * 
     */
    @Export(name="kubernetes", refs={Certificate.class}, tree="[0]")
    private Output<Certificate> kubernetes;

    /**
     * @return The kubernetes certificate.
     * 
     */
    public Output<Certificate> kubernetes() {
        return this.kubernetes;
    }
    /**
     * Map of node name to node configuration
     * 
     */
    @Export(name="nodes", refs={Map.class,String.class,ClusterPkiNode.class}, tree="[0,1,2]")
    private Output<Map<String,ClusterPkiNode>> nodes;

    /**
     * @return Map of node name to node configuration
     * 
     */
    public Output<Map<String,ClusterPkiNode>> nodes() {
        return this.nodes;
    }
    /**
     * Publicly accessible IP address.
     * 
     */
    @Export(name="publicIp", refs={String.class}, tree="[0]")
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
    @Export(name="rsaBits", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> rsaBits;

    /**
     * @return When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     * 
     */
    public Output<Optional<Integer>> rsaBits() {
        return Codegen.optional(this.rsaBits);
    }
    /**
     * The service accounts certificate
     * 
     */
    @Export(name="serviceAccounts", refs={Certificate.class}, tree="[0]")
    private Output<Certificate> serviceAccounts;

    /**
     * @return The service accounts certificate
     * 
     */
    public Output<Certificate> serviceAccounts() {
        return this.serviceAccounts;
    }
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    @Export(name="validityPeriodHours", refs={Integer.class}, tree="[0]")
    private Output<Integer> validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid.
     * 
     */
    public Output<Integer> validityPeriodHours() {
        return this.validityPeriodHours;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterPki(String name) {
        this(name, ClusterPkiArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterPki(String name, ClusterPkiArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterPki(String name, ClusterPkiArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:tls:ClusterPki", name, args == null ? ClusterPkiArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
