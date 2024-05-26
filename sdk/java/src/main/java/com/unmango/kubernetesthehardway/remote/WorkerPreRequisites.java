// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.Command;
import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.remote.WorkerPreRequisitesArgs;
import java.lang.Object;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Verifies that all worker node pre-requisites have been met.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:remote:WorkerPreRequisites")
public class WorkerPreRequisites extends com.pulumi.resources.ComponentResource {
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
     * Verifies that the conntrack binary exists.
     * 
     */
    @Export(name="conntrack", refs={Command.class}, tree="[0]")
    private Output<Command> conntrack;

    /**
     * @return Verifies that the conntrack binary exists.
     * 
     */
    public Output<Command> conntrack() {
        return this.conntrack;
    }
    /**
     * Verifies that the ipset binary exists.
     * 
     */
    @Export(name="ipset", refs={Command.class}, tree="[0]")
    private Output<Command> ipset;

    /**
     * @return Verifies that the ipset binary exists.
     * 
     */
    public Output<Command> ipset() {
        return this.ipset;
    }
    /**
     * Verifies that the socat binary exists.
     * 
     */
    @Export(name="socat", refs={Command.class}, tree="[0]")
    private Output<Command> socat;

    /**
     * @return Verifies that the socat binary exists.
     * 
     */
    public Output<Command> socat() {
        return this.socat;
    }
    /**
     * Verifies that swap is disabled.
     * 
     */
    @Export(name="swap", refs={Command.class}, tree="[0]")
    private Output<Command> swap;

    /**
     * @return Verifies that swap is disabled.
     * 
     */
    public Output<Command> swap() {
        return this.swap;
    }
    /**
     * Trigger recheck on changes to this input.
     * 
     */
    @Export(name="triggers", refs={List.class,Object.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Object>> triggers;

    /**
     * @return Trigger recheck on changes to this input.
     * 
     */
    public Output<Optional<List<Object>>> triggers() {
        return Codegen.optional(this.triggers);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WorkerPreRequisites(String name) {
        this(name, WorkerPreRequisitesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WorkerPreRequisites(String name, WorkerPreRequisitesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WorkerPreRequisites(String name, WorkerPreRequisitesArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:remote:WorkerPreRequisites", name, args == null ? WorkerPreRequisitesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
