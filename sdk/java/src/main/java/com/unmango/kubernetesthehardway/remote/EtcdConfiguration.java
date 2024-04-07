// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.EtcdConfigurationArgs;
import com.unmango.kubernetesthehardway.remote.File;
import com.unmango.kubernetesthehardway.remote.SystemdService;
import com.unmango.kubernetesthehardway.tools.Mkdir;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="kubernetes-the-hard-way:remote:EtcdConfiguration")
public class EtcdConfiguration extends com.pulumi.resources.ComponentResource {
    /**
     * The remote certificate authority file.
     * 
     */
    @Export(name="caFile", refs={File.class}, tree="[0]")
    private Output</* @Nullable */ File> caFile;

    /**
     * @return The remote certificate authority file.
     * 
     */
    public Output<Optional<File>> caFile() {
        return Codegen.optional(this.caFile);
    }
    /**
     * The remote certificate file.
     * 
     */
    @Export(name="certFile", refs={File.class}, tree="[0]")
    private Output</* @Nullable */ File> certFile;

    /**
     * @return The remote certificate file.
     * 
     */
    public Output<Optional<File>> certFile() {
        return Codegen.optional(this.certFile);
    }
    /**
     * The directory to store etcd configuration.
     * 
     */
    @Export(name="configurationDirectory", refs={String.class}, tree="[0]")
    private Output<String> configurationDirectory;

    /**
     * @return The directory to store etcd configuration.
     * 
     */
    public Output<String> configurationDirectory() {
        return this.configurationDirectory;
    }
    /**
     * The command used to create the configuration directory.
     * 
     */
    @Export(name="configurationMkdir", refs={Mkdir.class}, tree="[0]")
    private Output<Mkdir> configurationMkdir;

    /**
     * @return The command used to create the configuration directory.
     * 
     */
    public Output<Mkdir> configurationMkdir() {
        return this.configurationMkdir;
    }
    /**
     * The directory etcd will use.
     * 
     */
    @Export(name="dataDirectory", refs={String.class}, tree="[0]")
    private Output<String> dataDirectory;

    /**
     * @return The directory etcd will use.
     * 
     */
    public Output<String> dataDirectory() {
        return this.dataDirectory;
    }
    /**
     * The command used to create the data directory.
     * 
     */
    @Export(name="dataMkdir", refs={Mkdir.class}, tree="[0]")
    private Output<Mkdir> dataMkdir;

    /**
     * @return The command used to create the data directory.
     * 
     */
    public Output<Mkdir> dataMkdir() {
        return this.dataMkdir;
    }
    /**
     * IP used to serve client requests and communicate with etcd peers.
     * 
     */
    @Export(name="internalIp", refs={String.class}, tree="[0]")
    private Output<String> internalIp;

    /**
     * @return IP used to serve client requests and communicate with etcd peers.
     * 
     */
    public Output<String> internalIp() {
        return this.internalIp;
    }
    /**
     * The remote key file.
     * 
     */
    @Export(name="keyFile", refs={File.class}, tree="[0]")
    private Output</* @Nullable */ File> keyFile;

    /**
     * @return The remote key file.
     * 
     */
    public Output<Optional<File>> keyFile() {
        return Codegen.optional(this.keyFile);
    }
    /**
     * The remote systemd service.
     * 
     */
    @Export(name="systemdService", refs={SystemdService.class}, tree="[0]")
    private Output<SystemdService> systemdService;

    /**
     * @return The remote systemd service.
     * 
     */
    public Output<SystemdService> systemdService() {
        return this.systemdService;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EtcdConfiguration(String name) {
        this(name, EtcdConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EtcdConfiguration(String name, EtcdConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EtcdConfiguration(String name, EtcdConfigurationArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:EtcdConfiguration", name, args == null ? EtcdConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}