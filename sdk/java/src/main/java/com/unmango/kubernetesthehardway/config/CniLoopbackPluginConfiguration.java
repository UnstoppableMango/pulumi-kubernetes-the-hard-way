// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.config.CniLoopbackPluginConfigurationArgs;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Get the `loopback` configuration.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration")
public class CniLoopbackPluginConfiguration extends com.pulumi.resources.ComponentResource {
    @Export(name="result", refs={com.unmango.kubernetesthehardway.config.outputs.CniLoopbackPluginConfiguration.class}, tree="[0]")
    private Output<com.unmango.kubernetesthehardway.config.outputs.CniLoopbackPluginConfiguration> result;

    public Output<com.unmango.kubernetesthehardway.config.outputs.CniLoopbackPluginConfiguration> result() {
        return this.result;
    }
    /**
     * The yaml representation of the manifest.
     * 
     */
    @Export(name="yaml", refs={String.class}, tree="[0]")
    private Output<String> yaml;

    /**
     * @return The yaml representation of the manifest.
     * 
     */
    public Output<String> yaml() {
        return this.yaml;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CniLoopbackPluginConfiguration(String name) {
        this(name, CniLoopbackPluginConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CniLoopbackPluginConfiguration(String name, @Nullable CniLoopbackPluginConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CniLoopbackPluginConfiguration(String name, @Nullable CniLoopbackPluginConfigurationArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration", name, args == null ? CniLoopbackPluginConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
