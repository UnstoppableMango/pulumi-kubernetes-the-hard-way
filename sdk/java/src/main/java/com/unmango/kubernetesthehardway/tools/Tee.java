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
import com.unmango.kubernetesthehardway.tools.TeeArgs;
import com.unmango.kubernetesthehardway.tools.enums.CommandLifecycle;
import com.unmango.kubernetesthehardway.tools.enums.TeeMode;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Read from standard input and write to standard output and files.
 * 
 */
@ResourceType(type="kubernetes-the-hard-way:tools:Tee")
public class Tee extends com.pulumi.resources.ComponentResource {
    /**
     * Append to the given FILEs, do not overwrite.
     * 
     */
    @Export(name="append", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> append;

    /**
     * @return Append to the given FILEs, do not overwrite.
     * 
     */
    public Output<Boolean> append() {
        return this.append;
    }
    /**
     * Represents the command run on the remote system.
     * 
     */
    @Export(name="command", refs={Command.class}, tree="[0]")
    private Output<Command> command;

    /**
     * @return Represents the command run on the remote system.
     * 
     */
    public Output<Command> command() {
        return this.command;
    }
    /**
     * Connection details for the remote system.
     * 
     */
    @Export(name="connection", refs={Connection.class}, tree="[0]")
    private Output<Connection> connection;

    /**
     * @return Connection details for the remote system.
     * 
     */
    public Output<Connection> connection() {
        return this.connection;
    }
    /**
     * The file(s) to write to.
     * 
     */
    @Export(name="files", refs={Either.class,String.class,List.class}, tree="[0,1,[2,1]]")
    private Output<Either<String,List<String>>> files;

    /**
     * @return The file(s) to write to.
     * 
     */
    public Output<Either<String,List<String>>> files() {
        return this.files;
    }
    /**
     * Ignore interrupt signals.
     * 
     */
    @Export(name="ignoreInterrupts", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ignoreInterrupts;

    /**
     * @return Ignore interrupt signals.
     * 
     */
    public Output<Boolean> ignoreInterrupts() {
        return this.ignoreInterrupts;
    }
    /**
     * At what stage(s) in the resource lifecycle should the command be run.
     * 
     */
    @Export(name="lifecycle", refs={List.class,CommandLifecycle.class}, tree="[0,1]")
    private Output<List<CommandLifecycle>> lifecycle;

    /**
     * @return At what stage(s) in the resource lifecycle should the command be run.
     * 
     */
    public Output<List<CommandLifecycle>> lifecycle() {
        return this.lifecycle;
    }
    /**
     * Set behavior on write error.
     * 
     */
    @Export(name="outputError", refs={TeeMode.class}, tree="[0]")
    private Output</* @Nullable */ TeeMode> outputError;

    /**
     * @return Set behavior on write error.
     * 
     */
    public Output<Optional<TeeMode>> outputError() {
        return Codegen.optional(this.outputError);
    }
    /**
     * Operate in a more appropriate MODE with pipes.
     * 
     */
    @Export(name="pipe", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> pipe;

    /**
     * @return Operate in a more appropriate MODE with pipes.
     * 
     */
    public Output<Boolean> pipe() {
        return this.pipe;
    }
    @Export(name="stdin", refs={String.class}, tree="[0]")
    private Output<String> stdin;

    public Output<String> stdin() {
        return this.stdin;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Tee(String name) {
        this(name, TeeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Tee(String name, TeeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Tee(String name, TeeArgs args, @Nullable com.pulumi.resources.ComponentResourceOptions options) {
        super("kubernetes-the-hard-way:tools:Tee", name, args == null ? TeeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()), true);
    }

    private static com.pulumi.resources.ComponentResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.ComponentResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.ComponentResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.ComponentResourceOptions.merge(defaultOptions, options, id);
    }

}