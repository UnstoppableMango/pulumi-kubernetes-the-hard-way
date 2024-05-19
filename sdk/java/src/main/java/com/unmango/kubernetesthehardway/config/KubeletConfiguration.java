// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.config.KubeletConfigurationArgs;
import java.lang.String;
import javax.annotation.Nullable;

@ResourceType(type="kubernetes-the-hard-way:config:KubeletConfiguration")
public class KubeletConfiguration extends com.pulumi.resources.ComponentResource {
    @Export(name="result", refs={com.unmango.kubernetesthehardway.config.outputs.KubeletConfiguration.class}, tree="[0]")
    private Output<com.unmango.kubernetesthehardway.config.outputs.KubeletConfiguration> result;

    public Output<com.unmango.kubernetesthehardway.config.outputs.KubeletConfiguration> result() {
        return this.result;
    }
    /**
     * The yaml representation of the manifest
     * 
     */
    @Export(name="yaml", refs={String.class}, tree="[0]")
    private Output<String> yaml;

    /**
     * @return The yaml representation of the manifest
     * 
     */
    public Output<String> yaml() {
        return this.yaml;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KubeletConfiguration(String name) {
        this(name, KubeletConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KubeletConfiguration(String name, KubeletConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KubeletConfiguration(String name, KubeletConfigurationArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:config:KubeletConfiguration", name, args == null ? KubeletConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
