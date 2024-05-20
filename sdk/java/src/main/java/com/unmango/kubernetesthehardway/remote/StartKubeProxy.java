// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.StartKubeProxyArgs;
import com.unmango.kubernetesthehardway.tools.Systemctl;
import javax.annotation.Nullable;

/**
 * Starts `kube-proxy` on a remote system
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:StartKubeProxy")
public class StartKubeProxy extends com.pulumi.resources.ComponentResource {
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
     * The daemon-reload command.
     * 
     */
    @Export(name="daemonReload", refs={Systemctl.class}, tree="[0]")
    private Output<Systemctl> daemonReload;

    /**
     * @return The daemon-reload command.
     * 
     */
    public Output<Systemctl> daemonReload() {
        return this.daemonReload;
    }
    /**
     * The enable command.
     * 
     */
    @Export(name="enable", refs={Systemctl.class}, tree="[0]")
    private Output<Systemctl> enable;

    /**
     * @return The enable command.
     * 
     */
    public Output<Systemctl> enable() {
        return this.enable;
    }
    /**
     * The start command.
     * 
     */
    @Export(name="start", refs={Systemctl.class}, tree="[0]")
    private Output<Systemctl> start;

    /**
     * @return The start command.
     * 
     */
    public Output<Systemctl> start() {
        return this.start;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StartKubeProxy(String name) {
        this(name, StartKubeProxyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StartKubeProxy(String name, StartKubeProxyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StartKubeProxy(String name, StartKubeProxyArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:StartKubeProxy", name, args == null ? StartKubeProxyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
