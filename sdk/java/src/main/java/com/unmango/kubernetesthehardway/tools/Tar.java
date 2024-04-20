// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.Command;
import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.tools.TarArgs;
import com.unmango.kubernetesthehardway.tools.outputs.TarOpts;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Abstraction over the `tar` utility on a remote system.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:tools:Tar")
public class Tar extends com.pulumi.resources.CustomResource {
    /**
     * Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
     * 
     */
    @Export(name="binaryPath", refs={String.class}, tree="[0]")
    private Output<String> binaryPath;

    /**
     * @return Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
     * 
     */
    public Output<String> binaryPath() {
        return this.binaryPath;
    }
    /**
     * The underlying command
     * 
     */
    @Export(name="command", refs={Command.class}, tree="[0]")
    private Output<Command> command;

    /**
     * @return The underlying command
     * 
     */
    public Output<Command> command() {
        return this.command;
    }
    /**
     * Connection details for the remote system
     * 
     */
    @Export(name="connection", refs={Connection.class}, tree="[0]")
    private Output<Connection> connection;

    /**
     * @return Connection details for the remote system
     * 
     */
    public Output<Connection> connection() {
        return this.connection;
    }
    /**
     * The command to run on create.
     * 
     */
    @Export(name="create", refs={TarOpts.class}, tree="[0]")
    private Output</* @Nullable */ TarOpts> create;

    /**
     * @return The command to run on create.
     * 
     */
    public Output<Optional<TarOpts>> create() {
        return Codegen.optional(this.create);
    }
    /**
     * The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
     * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
     * Command resource from previous create or update steps.
     * 
     */
    @Export(name="delete", refs={TarOpts.class}, tree="[0]")
    private Output</* @Nullable */ TarOpts> delete;

    /**
     * @return The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
     * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
     * Command resource from previous create or update steps.
     * 
     */
    public Output<Optional<TarOpts>> delete() {
        return Codegen.optional(this.delete);
    }
    /**
     * Environment variables
     * 
     */
    @Export(name="environment", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> environment;

    /**
     * @return Environment variables
     * 
     */
    public Output<Map<String,String>> environment() {
        return this.environment;
    }
    /**
     * TODO
     * 
     */
    @Export(name="stderr", refs={String.class}, tree="[0]")
    private Output<String> stderr;

    /**
     * @return TODO
     * 
     */
    public Output<String> stderr() {
        return this.stderr;
    }
    /**
     * TODO
     * 
     */
    @Export(name="stdin", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stdin;

    /**
     * @return TODO
     * 
     */
    public Output<Optional<String>> stdin() {
        return Codegen.optional(this.stdin);
    }
    /**
     * TODO
     * 
     */
    @Export(name="stdout", refs={String.class}, tree="[0]")
    private Output<String> stdout;

    /**
     * @return TODO
     * 
     */
    public Output<String> stdout() {
        return this.stdout;
    }
    /**
     * TODO
     * 
     */
    @Export(name="triggers", refs={List.class,Object.class}, tree="[0,1]")
    private Output<List<Object>> triggers;

    /**
     * @return TODO
     * 
     */
    public Output<List<Object>> triggers() {
        return this.triggers;
    }
    /**
     * The command to run on update, if empty, create will
     * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
     * are set to the stdout and stderr properties of the Command resource from previous
     * create or update steps.
     * 
     */
    @Export(name="update", refs={TarOpts.class}, tree="[0]")
    private Output</* @Nullable */ TarOpts> update;

    /**
     * @return The command to run on update, if empty, create will
     * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
     * are set to the stdout and stderr properties of the Command resource from previous
     * create or update steps.
     * 
     */
    public Output<Optional<TarOpts>> update() {
        return Codegen.optional(this.update);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Tar(String name) {
        this(name, TarArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Tar(String name, TarArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Tar(String name, TarArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes-the-hard-way:tools:Tar", name, args == null ? TarArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Tar(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("kubernetes-the-hard-way:tools:Tar", name, null, makeResourceOptions(options, id));
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
    public static Tar get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Tar(name, id, options);
    }
}
