// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.CniBridgePluginConfigurationArgs;
import com.unmango.kubernetesthehardway.remote.outputs.CniBridgeIpam;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The CNI bridge plugin configuration.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:CniBridgePluginConfiguration")
public class CniBridgePluginConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * Bridge name.
     * 
     */
    @Export(name="bridge", refs={String.class}, tree="[0]")
    private Output<String> bridge;

    /**
     * @return Bridge name.
     * 
     */
    public Output<String> bridge() {
        return this.bridge;
    }
    /**
     * CNI version.
     * 
     */
    @Export(name="cniVersion", refs={String.class}, tree="[0]")
    private Output<String> cniVersion;

    /**
     * @return CNI version.
     * 
     */
    public Output<String> cniVersion() {
        return this.cniVersion;
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
     * IP masq.
     * 
     */
    @Export(name="ipMasq", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ipMasq;

    /**
     * @return IP masq.
     * 
     */
    public Output<Boolean> ipMasq() {
        return this.ipMasq;
    }
    /**
     * IPAM
     * 
     */
    @Export(name="ipam", refs={CniBridgeIpam.class}, tree="[0]")
    private Output<CniBridgeIpam> ipam;

    /**
     * @return IPAM
     * 
     */
    public Output<CniBridgeIpam> ipam() {
        return this.ipam;
    }
    /**
     * Is gateway.
     * 
     */
    @Export(name="isGateway", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isGateway;

    /**
     * @return Is gateway.
     * 
     */
    public Output<Boolean> isGateway() {
        return this.isGateway;
    }
    /**
     * CNI plugin name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return CNI plugin name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * CNI plugin type.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return CNI plugin type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CniBridgePluginConfiguration(String name) {
        this(name, CniBridgePluginConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CniBridgePluginConfiguration(String name, CniBridgePluginConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CniBridgePluginConfiguration(String name, CniBridgePluginConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes-the-hard-way:remote:CniBridgePluginConfiguration", name, args == null ? CniBridgePluginConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CniBridgePluginConfiguration(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes-the-hard-way:remote:CniBridgePluginConfiguration", name, null, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CniBridgePluginConfiguration get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CniBridgePluginConfiguration(name, id, options);
    }
}
