// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.CrictlInstallArgs;
import com.unmango.kubernetesthehardway.remote.Download;
import com.unmango.kubernetesthehardway.remote.enums.Architecture;
import com.unmango.kubernetesthehardway.tools.Mkdir;
import com.unmango.kubernetesthehardway.tools.Mktemp;
import com.unmango.kubernetesthehardway.tools.Mv;
import com.unmango.kubernetesthehardway.tools.Rm;
import com.unmango.kubernetesthehardway.tools.Tar;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Installs crictl on a remote system
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:CrictlInstall")
public class CrictlInstall extends com.pulumi.resources.ComponentResource {
    /**
     * The CPU architecture to install.
     * 
     */
    @Export(name="architecture", refs={Architecture.class}, tree="[0]")
    private Output<Architecture> architecture;

    /**
     * @return The CPU architecture to install.
     * 
     */
    public Output<Architecture> architecture() {
        return this.architecture;
    }
    /**
     * The name of the downloaded archive.
     * 
     */
    @Export(name="archiveName", refs={String.class}, tree="[0]")
    private Output<String> archiveName;

    /**
     * @return The name of the downloaded archive.
     * 
     */
    public Output<String> archiveName() {
        return this.archiveName;
    }
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
     * The crictl mv operation.
     * 
     */
    @Export(name="crictlMv", refs={Mv.class}, tree="[0]")
    private Output<Mv> crictlMv;

    /**
     * @return The crictl mv operation.
     * 
     */
    public Output<Mv> crictlMv() {
        return this.crictlMv;
    }
    /**
     * The crictl path on the remote system
     * 
     */
    @Export(name="crictlPath", refs={String.class}, tree="[0]")
    private Output<String> crictlPath;

    /**
     * @return The crictl path on the remote system
     * 
     */
    public Output<String> crictlPath() {
        return this.crictlPath;
    }
    /**
     * The directory to install the binary to.
     * 
     */
    @Export(name="directory", refs={String.class}, tree="[0]")
    private Output<String> directory;

    /**
     * @return The directory to install the binary to.
     * 
     */
    public Output<String> directory() {
        return this.directory;
    }
    /**
     * The download operation.
     * 
     */
    @Export(name="download", refs={Download.class}, tree="[0]")
    private Output<Download> download;

    /**
     * @return The download operation.
     * 
     */
    public Output<Download> download() {
        return this.download;
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
     * The mktemp operation.
     * 
     */
    @Export(name="mktemp", refs={Mktemp.class}, tree="[0]")
    private Output<Mktemp> mktemp;

    /**
     * @return The mktemp operation.
     * 
     */
    public Output<Mktemp> mktemp() {
        return this.mktemp;
    }
    /**
     * The path to the installed binary.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return The path to the installed binary.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The rm operation.
     * 
     */
    @Export(name="rm", refs={Rm.class}, tree="[0]")
    private Output<Rm> rm;

    /**
     * @return The rm operation.
     * 
     */
    public Output<Rm> rm() {
        return this.rm;
    }
    /**
     * The tar operation.
     * 
     */
    @Export(name="tar", refs={Tar.class}, tree="[0]")
    private Output<Tar> tar;

    /**
     * @return The tar operation.
     * 
     */
    public Output<Tar> tar() {
        return this.tar;
    }
    /**
     * The url used to download the binary.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The url used to download the binary.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * The version to install.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return The version to install.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CrictlInstall(String name) {
        this(name, CrictlInstallArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CrictlInstall(String name, CrictlInstallArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CrictlInstall(String name, CrictlInstallArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:CrictlInstall", name, args == null ? CrictlInstallArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
