// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.DownloadArgs;
import com.unmango.kubernetesthehardway.tools.Mkdir;
import com.unmango.kubernetesthehardway.tools.Wget;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Downloads the file specified by `url` onto a remote system.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:Download")
public class Download extends com.pulumi.resources.ComponentResource {
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
     * The fully qualified path on the remote system where the file should be downloaded to.
     * 
     */
    @Export(name="destination", refs={String.class}, tree="[0]")
    private Output<String> destination;

    /**
     * @return The fully qualified path on the remote system where the file should be downloaded to.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }
    /**
     * The mkdir operation.
     * 
     */
    @Export(name="mkdir", refs={Mkdir.class}, tree="[0]")
    private Output<Mkdir> mkdir;

    /**
     * @return The mkdir operation.
     * 
     */
    public Output<Mkdir> mkdir() {
        return this.mkdir;
    }
    /**
     * Remove the downloaded fiel when the resource is deleted.
     * 
     */
    @Export(name="removeOnDelete", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> removeOnDelete;

    /**
     * @return Remove the downloaded fiel when the resource is deleted.
     * 
     */
    public Output<Boolean> removeOnDelete() {
        return this.removeOnDelete;
    }
    /**
     * The URL of the file to be downloaded.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the file to be downloaded.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The wget operation.
     * 
     */
    @Export(name="wget", refs={Wget.class}, tree="[0]")
    private Output<Wget> wget;

    /**
     * @return The wget operation.
     * 
     */
    public Output<Wget> wget() {
        return this.wget;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Download(String name) {
        this(name, DownloadArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Download(String name, DownloadArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Download(String name, DownloadArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:Download", name, args == null ? DownloadArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
