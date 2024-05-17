// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.ContainerdConfigurationArgs;
import com.unmango.kubernetesthehardway.remote.File;
import com.unmango.kubernetesthehardway.remote.outputs.ContainerdCriPluginConfiguration;
import javax.annotation.Nullable;

/**
 * The containerd configuration file.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:ContainerdConfiguration")
public class ContainerdConfiguration extends com.pulumi.resources.ComponentResource {
    /**
     * The parameters with which to connect to the remote host.
     * 
     */
    @Export(name="connection", refs={Connection.class}, tree="[0]")
    private Output<Connection> connection;

    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    public Output<Connection> connection() {
        return this.connection;
    }
    /**
     * The cri configuration.
     * 
     */
    @Export(name="cri", refs={ContainerdCriPluginConfiguration.class}, tree="[0]")
    private Output<ContainerdCriPluginConfiguration> cri;

    /**
     * @return The cri configuration.
     * 
     */
    public Output<ContainerdCriPluginConfiguration> cri() {
        return this.cri;
    }
    /**
     * The remote configuration file.
     * 
     */
    @Export(name="file", refs={File.class}, tree="[0]")
    private Output<File> file;

    /**
     * @return The remote configuration file.
     * 
     */
    public Output<File> file() {
        return this.file;
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
    public ContainerdConfiguration(String name, ContainerdConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerdConfiguration(String name, ContainerdConfigurationArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:ContainerdConfiguration", name, args == null ? ContainerdConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}