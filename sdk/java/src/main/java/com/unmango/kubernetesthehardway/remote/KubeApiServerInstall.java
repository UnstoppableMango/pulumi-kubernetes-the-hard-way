// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.KubeApiServerInstallArgs;
import com.unmango.kubernetesthehardway.remote.enums.Architecture;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Installs kube-apiserver on a remote system.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:KubeApiServerInstall")
public class KubeApiServerInstall extends com.pulumi.resources.ComponentResource {
    /**
     * The kube-apiserver CPU architecture.
     * 
     */
    @Export(name="architecture", refs={Architecture.class}, tree="[0]")
    private Output<Architecture> architecture;

    /**
     * @return The kube-apiserver CPU architecture.
     * 
     */
    public Output<Architecture> architecture() {
        return this.architecture;
    }
    /**
     * The connection details.
     * 
     */
    @Export(name="connection", refs={Connection.class}, tree="[0]")
    private Output<Connection> connection;

    /**
     * @return The connection details.
     * 
     */
    public Output<Connection> connection() {
        return this.connection;
    }
    /**
     * Directory to install the `kube-apiserver` binary.
     * 
     */
    @Export(name="installDirectory", refs={String.class}, tree="[0]")
    private Output<String> installDirectory;

    /**
     * @return Directory to install the `kube-apiserver` binary.
     * 
     */
    public Output<String> installDirectory() {
        return this.installDirectory;
    }
    /**
     * The version of kube-apiserver to install.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return The version of kube-apiserver to install.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KubeApiServerInstall(String name) {
        this(name, KubeApiServerInstallArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KubeApiServerInstall(String name, KubeApiServerInstallArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KubeApiServerInstall(String name, KubeApiServerInstallArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:KubeApiServerInstall", name, args == null ? KubeApiServerInstallArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
