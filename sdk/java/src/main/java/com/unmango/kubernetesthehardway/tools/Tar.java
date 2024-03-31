// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.Command;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.Utilities;
import com.unmango.kubernetesthehardway.tools.TarArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Abstracion over the `tar` utility on a remote system.
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
     * Represents the remote `tar` operation.
     * 
     */
    @Export(name="command", refs={Command.class}, tree="[0]")
    private Output<Command> command;

    /**
     * @return Represents the remote `tar` operation.
     * 
     */
    public Output<Command> command() {
        return this.command;
    }
    /**
     * Corresponds to the --directory option.
     * 
     */
    @Export(name="directory", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> directory;

    /**
     * @return Corresponds to the --directory option.
     * 
     */
    public Output<Optional<String>> directory() {
        return Codegen.optional(this.directory);
    }
    /**
     * Corresponds to the --extract option.
     * 
     */
    @Export(name="extract", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> extract;

    /**
     * @return Corresponds to the --extract option.
     * 
     */
    public Output<Boolean> extract() {
        return this.extract;
    }
    /**
     * Corresponds to the [FILE] argument.
     * 
     */
    @Export(name="files", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> files;

    /**
     * @return Corresponds to the [FILE] argument.
     * 
     */
    public Output<List<String>> files() {
        return this.files;
    }
    /**
     * Corresponds to the --gzip option.
     * 
     */
    @Export(name="gzip", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> gzip;

    /**
     * @return Corresponds to the --gzip option.
     * 
     */
    public Output<Optional<Boolean>> gzip() {
        return Codegen.optional(this.gzip);
    }
    /**
     * The process&#39; stderr.
     * 
     */
    @Export(name="stderr", refs={String.class}, tree="[0]")
    private Output<String> stderr;

    /**
     * @return The process&#39; stderr.
     * 
     */
    public Output<String> stderr() {
        return this.stderr;
    }
    /**
     * The process&#39; stdin.
     * 
     */
    @Export(name="stdin", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stdin;

    /**
     * @return The process&#39; stdin.
     * 
     */
    public Output<Optional<String>> stdin() {
        return Codegen.optional(this.stdin);
    }
    /**
     * The process&#39; stdout.
     * 
     */
    @Export(name="stdout", refs={String.class}, tree="[0]")
    private Output<String> stdout;

    /**
     * @return The process&#39; stdout.
     * 
     */
    public Output<String> stdout() {
        return this.stdout;
    }
    /**
     * Corresponds to the --strip-components option.
     * 
     */
    @Export(name="stripComponents", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stripComponents;

    /**
     * @return Corresponds to the --strip-components option.
     * 
     */
    public Output<Optional<Integer>> stripComponents() {
        return Codegen.optional(this.stripComponents);
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
