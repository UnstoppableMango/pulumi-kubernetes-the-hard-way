// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.config.ContainerdConfigurationArgs;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Get the containerd configuration.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:config:ContainerdConfiguration")
public class ContainerdConfiguration extends com.pulumi.resources.ComponentResource {
    @Export(name="result", refs={com.unmango.kubernetesthehardway.config.outputs.ContainerdConfiguration.class}, tree="[0]")
    private Output<com.unmango.kubernetesthehardway.config.outputs.ContainerdConfiguration> result;

    public Output<com.unmango.kubernetesthehardway.config.outputs.ContainerdConfiguration> result() {
        return this.result;
    }
    /**
     * The toml representation of the configuration.
     * 
     */
    @Export(name="toml", refs={String.class}, tree="[0]")
    private Output<String> toml;

    /**
     * @return The toml representation of the configuration.
     * 
     */
    public Output<String> toml() {
        return this.toml;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerdConfiguration(String name) {
        this(name, ContainerdConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerdConfiguration(String name, @Nullable ContainerdConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerdConfiguration(String name, @Nullable ContainerdConfigurationArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:config:ContainerdConfiguration", name, args == null ? ContainerdConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
