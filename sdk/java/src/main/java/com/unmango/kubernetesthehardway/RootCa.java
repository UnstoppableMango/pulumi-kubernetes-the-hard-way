// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tls.PrivateKey;
import com.pulumi.tls.SelfSignedCert;
import com.unmango.kubernetesthehardway.RootCaArgs;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.enums.AllowedUsage;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

@ResourceType(type="kubernetes-the-hard-way:index:RootCa")
public class RootCa extends com.pulumi.resources.ComponentResource {
    @Export(name="allowedUses", refs={List.class,AllowedUsage.class}, tree="[0,1]")
    private Output<List<AllowedUsage>> allowedUses;

    public Output<List<AllowedUsage>> allowedUses() {
        return this.allowedUses;
    }
    @Export(name="cert", refs={SelfSignedCert.class}, tree="[0]")
    private Output<SelfSignedCert> cert;

    public Output<SelfSignedCert> cert() {
        return this.cert;
    }
    @Export(name="certPem", refs={String.class}, tree="[0]")
    private Output<String> certPem;

    public Output<String> certPem() {
        return this.certPem;
    }
    @Export(name="key", refs={PrivateKey.class}, tree="[0]")
    private Output<PrivateKey> key;

    public Output<PrivateKey> key() {
        return this.key;
    }
    @Export(name="privateKeyPem", refs={String.class}, tree="[0]")
    private Output<String> privateKeyPem;

    public Output<String> privateKeyPem() {
        return this.privateKeyPem;
    }
    @Export(name="publicKeyPem", refs={String.class}, tree="[0]")
    private Output<String> publicKeyPem;

    public Output<String> publicKeyPem() {
        return this.publicKeyPem;
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
        super("kubernetes-the-hard-way:index:RootCa", name, args == null ? RootCaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
