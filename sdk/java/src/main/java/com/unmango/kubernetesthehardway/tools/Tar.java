// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.Command;
import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.tools.TarArgs;
import com.unmango.kubernetesthehardway.tools.outputs.CommandLifecycle;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Abstraction over the `rm` utility on a remote system.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:tools:Tar")
public class Tar extends com.pulumi.resources.ComponentResource {
    /**
     * Corresponds to the [ARCHIVE] argument.
     * 
     */
    @Export(name="archive", refs={String.class}, tree="[0]")
    private Output<String> archive;

    /**
     * @return Corresponds to the [ARCHIVE] argument.
     * 
     */
    public Output<String> archive() {
        return this.archive;
    }
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
     * Corresponds to the `--directory` option.
     * 
     */
    @Export(name="directory", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> directory;

    /**
     * @return Corresponds to the `--directory` option.
     * 
     */
    public Output<Optional<String>> directory() {
        return Codegen.optional(this.directory);
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
     * Corresponds to the `--extract` option.
     * 
     */
    @Export(name="extract", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> extract;

    /**
     * @return Corresponds to the `--extract` option.
     * 
     */
    public Output<Boolean> extract() {
        return this.extract;
    }
    /**
     * Corresponds to the [FILE] argument.
     * 
     */
    @Export(name="files", refs={Either.class,String.class,List.class}, tree="[0,1,[2,1]]")
    private Output<Either<String,List<String>>> files;

    /**
     * @return Corresponds to the [FILE] argument.
     * 
     */
    public Output<Either<String,List<String>>> files() {
        return this.files;
    }
    /**
     * Corresponds to the `--gzip` option.
     * 
     */
    @Export(name="gzip", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> gzip;

    /**
     * @return Corresponds to the `--gzip` option.
     * 
     */
    public Output<Optional<Boolean>> gzip() {
        return Codegen.optional(this.gzip);
    }
    /**
     * At what stage(s) in the resource lifecycle should the command be run
     * 
     */
    @Export(name="lifecycle", refs={CommandLifecycle.class}, tree="[0]")
    private Output</* @Nullable */ CommandLifecycle> lifecycle;

    /**
     * @return At what stage(s) in the resource lifecycle should the command be run
     * 
     */
    public Output<Optional<CommandLifecycle>> lifecycle() {
        return Codegen.optional(this.lifecycle);
    }
    /**
     * Whether rm should be run when the resource is created or deleted.
     * 
     */
    @Export(name="onDelete", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> onDelete;

    /**
     * @return Whether rm should be run when the resource is created or deleted.
     * 
     */
    public Output<Optional<Boolean>> onDelete() {
        return Codegen.optional(this.onDelete);
    }
    /**
     * Corresponds to the `--strip-components` option.
     * 
     */
    @Export(name="recursive", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> recursive;

    /**
     * @return Corresponds to the `--strip-components` option.
     * 
     */
    public Output<Optional<Integer>> recursive() {
        return Codegen.optional(this.recursive);
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
    public Tar(String name, TarArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:tools:Tar", name, args == null ? TarArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}
