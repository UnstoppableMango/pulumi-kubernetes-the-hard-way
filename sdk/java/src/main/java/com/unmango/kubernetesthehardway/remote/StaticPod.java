// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.config.outputs.PodManifest;
import com.unmango.kubernetesthehardway.remote.File;
import com.unmango.kubernetesthehardway.remote.StaticPodArgs;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Create a static pod manifest on a remote system.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:StaticPod")
public class StaticPod extends com.pulumi.resources.ComponentResource {
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
     * The remote manifest file.
     * 
     */
    @Export(name="file", refs={File.class}, tree="[0]")
    private Output<File> file;

    /**
     * @return The remote manifest file.
     * 
     */
    public Output<File> file() {
        return this.file;
    }
    /**
     * The name of the file on the remote system.
     * 
     */
    @Export(name="fileName", refs={String.class}, tree="[0]")
    private Output<String> fileName;

    /**
     * @return The name of the file on the remote system.
     * 
     */
    public Output<String> fileName() {
        return this.fileName;
    }
    /**
     * The path to the manifest on the remote system.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The path to the manifest on the remote system.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * The pod manifest.
     * 
     */
    @Export(name="pod", refs={PodManifest.class}, tree="[0]")
    private Output<PodManifest> pod;

    /**
     * @return The pod manifest.
     * 
     */
    public Output<PodManifest> pod() {
        return this.pod;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StaticPod(String name) {
        this(name, StaticPodArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StaticPod(String name, StaticPodArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StaticPod(String name, StaticPodArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:StaticPod", name, args == null ? StaticPodArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}