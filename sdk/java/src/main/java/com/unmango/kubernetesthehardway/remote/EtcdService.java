// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.outputs.SystemdServiceRestart;
import com.unmango.kubernetesthehardway.remote.EtcdServiceArgs;
import com.unmango.kubernetesthehardway.remote.SystemdService;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Etcd systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:EtcdService")
public class EtcdService extends com.pulumi.resources.ComponentResource {
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
     * Optional systemd unit description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional systemd unit description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Optional systemd unit documentation
     * 
     */
    @Export(name="documentation", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> documentation;

    /**
     * @return Optional systemd unit documentation
     * 
     */
    public Output<Optional<String>> documentation() {
        return Codegen.optional(this.documentation);
    }
    /**
     * Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
     * 
     */
    @Export(name="restart", refs={SystemdServiceRestart.class}, tree="[0]")
    private Output</* @Nullable */ SystemdServiceRestart> restart;

    /**
     * @return Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
     * 
     */
    public Output<Optional<SystemdServiceRestart>> restart() {
        return Codegen.optional(this.restart);
    }
    /**
     * Optionally override the systemd service RestartSec. Defaults to `5`.
     * 
     */
    @Export(name="restartSec", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> restartSec;

    /**
     * @return Optionally override the systemd service RestartSec. Defaults to `5`.
     * 
     */
    public Output<Optional<Integer>> restartSec() {
        return Codegen.optional(this.restartSec);
    }
    /**
     * The remote systemd service.
     * 
     */
    @Export(name="service", refs={SystemdService.class}, tree="[0]")
    private Output<SystemdService> service;

    /**
     * @return The remote systemd service.
     * 
     */
    public Output<SystemdService> service() {
        return this.service;
    }
    /**
     * Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
     * 
     */
    @Export(name="wantedBy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> wantedBy;

    /**
     * @return Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
     * 
     */
    public Output<Optional<String>> wantedBy() {
        return Codegen.optional(this.wantedBy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EtcdService(String name) {
        this(name, EtcdServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EtcdService(String name, EtcdServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EtcdService(String name, EtcdServiceArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:EtcdService", name, args == null ? EtcdServiceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
