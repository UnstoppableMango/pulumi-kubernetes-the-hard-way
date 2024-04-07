// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.inputs.CommandLifecycle;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MktempArgs extends com.pulumi.resources.ResourceArgs {

    public static final MktempArgs Empty = new MktempArgs();

    /**
     * Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
     * 
     */
    @Import(name="binaryPath")
    private @Nullable Output<String> binaryPath;

    /**
     * @return Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
     * 
     */
    public Optional<Output<String>> binaryPath() {
        return Optional.ofNullable(this.binaryPath);
    }

    /**
     * Connection details for the remote system
     * 
     */
    @Import(name="connection", required=true)
    private Output<ConnectionArgs> connection;

    /**
     * @return Connection details for the remote system
     * 
     */
    public Output<ConnectionArgs> connection() {
        return this.connection;
    }

    /**
     * Corresponds to the `--directory` option.
     * 
     */
    @Import(name="directory")
    private @Nullable Output<Boolean> directory;

    /**
     * @return Corresponds to the `--directory` option.
     * 
     */
    public Optional<Output<Boolean>> directory() {
        return Optional.ofNullable(this.directory);
    }

    /**
     * Corresponds to the `--dry-run` option.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Corresponds to the `--dry-run` option.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * Environment variables
     * 
     */
    @Import(name="environment")
    private @Nullable Output<Map<String,String>> environment;

    /**
     * @return Environment variables
     * 
     */
    public Optional<Output<Map<String,String>>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * At what stage(s) in the resource lifecycle should the command be run
     * 
     */
    @Import(name="lifecycle")
    private @Nullable CommandLifecycle lifecycle;

    /**
     * @return At what stage(s) in the resource lifecycle should the command be run
     * 
     */
    public Optional<CommandLifecycle> lifecycle() {
        return Optional.ofNullable(this.lifecycle);
    }

    /**
     * Corresponds to the `--quiet` option.
     * 
     */
    @Import(name="quiet")
    private @Nullable Output<Boolean> quiet;

    /**
     * @return Corresponds to the `--quiet` option.
     * 
     */
    public Optional<Output<Boolean>> quiet() {
        return Optional.ofNullable(this.quiet);
    }

    /**
     * TODO
     * 
     */
    @Import(name="stdin")
    private @Nullable Output<String> stdin;

    /**
     * @return TODO
     * 
     */
    public Optional<Output<String>> stdin() {
        return Optional.ofNullable(this.stdin);
    }

    /**
     * Corresponds to the `--suffix` option.
     * 
     */
    @Import(name="suffix")
    private @Nullable Output<String> suffix;

    /**
     * @return Corresponds to the `--suffix` option.
     * 
     */
    public Optional<Output<String>> suffix() {
        return Optional.ofNullable(this.suffix);
    }

    /**
     * Corresponds to the [TEMPLATE] argument.
     * 
     */
    @Import(name="template")
    private @Nullable Output<Boolean> template;

    /**
     * @return Corresponds to the [TEMPLATE] argument.
     * 
     */
    public Optional<Output<Boolean>> template() {
        return Optional.ofNullable(this.template);
    }

    /**
     * Corresponds to the `--tmpdir` option.
     * 
     */
    @Import(name="tmpdir")
    private @Nullable Output<String> tmpdir;

    /**
     * @return Corresponds to the `--tmpdir` option.
     * 
     */
    public Optional<Output<String>> tmpdir() {
        return Optional.ofNullable(this.tmpdir);
    }

    /**
     * TODO
     * 
     */
    @Import(name="triggers")
    private @Nullable Output<List<Object>> triggers;

    /**
     * @return TODO
     * 
     */
    public Optional<Output<List<Object>>> triggers() {
        return Optional.ofNullable(this.triggers);
    }

    private MktempArgs() {}

    private MktempArgs(MktempArgs $) {
        this.binaryPath = $.binaryPath;
        this.connection = $.connection;
        this.directory = $.directory;
        this.dryRun = $.dryRun;
        this.environment = $.environment;
        this.lifecycle = $.lifecycle;
        this.quiet = $.quiet;
        this.stdin = $.stdin;
        this.suffix = $.suffix;
        this.template = $.template;
        this.tmpdir = $.tmpdir;
        this.triggers = $.triggers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MktempArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MktempArgs $;

        public Builder() {
            $ = new MktempArgs();
        }

        public Builder(MktempArgs defaults) {
            $ = new MktempArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param binaryPath Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
         * 
         * @return builder
         * 
         */
        public Builder binaryPath(@Nullable Output<String> binaryPath) {
            $.binaryPath = binaryPath;
            return this;
        }

        /**
         * @param binaryPath Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
         * 
         * @return builder
         * 
         */
        public Builder binaryPath(String binaryPath) {
            return binaryPath(Output.of(binaryPath));
        }

        /**
         * @param connection Connection details for the remote system
         * 
         * @return builder
         * 
         */
        public Builder connection(Output<ConnectionArgs> connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param connection Connection details for the remote system
         * 
         * @return builder
         * 
         */
        public Builder connection(ConnectionArgs connection) {
            return connection(Output.of(connection));
        }

        /**
         * @param directory Corresponds to the `--directory` option.
         * 
         * @return builder
         * 
         */
        public Builder directory(@Nullable Output<Boolean> directory) {
            $.directory = directory;
            return this;
        }

        /**
         * @param directory Corresponds to the `--directory` option.
         * 
         * @return builder
         * 
         */
        public Builder directory(Boolean directory) {
            return directory(Output.of(directory));
        }

        /**
         * @param dryRun Corresponds to the `--dry-run` option.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Corresponds to the `--dry-run` option.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param environment Environment variables
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<Map<String,String>> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment Environment variables
         * 
         * @return builder
         * 
         */
        public Builder environment(Map<String,String> environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param lifecycle At what stage(s) in the resource lifecycle should the command be run
         * 
         * @return builder
         * 
         */
        public Builder lifecycle(@Nullable CommandLifecycle lifecycle) {
            $.lifecycle = lifecycle;
            return this;
        }

        /**
         * @param quiet Corresponds to the `--quiet` option.
         * 
         * @return builder
         * 
         */
        public Builder quiet(@Nullable Output<Boolean> quiet) {
            $.quiet = quiet;
            return this;
        }

        /**
         * @param quiet Corresponds to the `--quiet` option.
         * 
         * @return builder
         * 
         */
        public Builder quiet(Boolean quiet) {
            return quiet(Output.of(quiet));
        }

        /**
         * @param stdin TODO
         * 
         * @return builder
         * 
         */
        public Builder stdin(@Nullable Output<String> stdin) {
            $.stdin = stdin;
            return this;
        }

        /**
         * @param stdin TODO
         * 
         * @return builder
         * 
         */
        public Builder stdin(String stdin) {
            return stdin(Output.of(stdin));
        }

        /**
         * @param suffix Corresponds to the `--suffix` option.
         * 
         * @return builder
         * 
         */
        public Builder suffix(@Nullable Output<String> suffix) {
            $.suffix = suffix;
            return this;
        }

        /**
         * @param suffix Corresponds to the `--suffix` option.
         * 
         * @return builder
         * 
         */
        public Builder suffix(String suffix) {
            return suffix(Output.of(suffix));
        }

        /**
         * @param template Corresponds to the [TEMPLATE] argument.
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<Boolean> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template Corresponds to the [TEMPLATE] argument.
         * 
         * @return builder
         * 
         */
        public Builder template(Boolean template) {
            return template(Output.of(template));
        }

        /**
         * @param tmpdir Corresponds to the `--tmpdir` option.
         * 
         * @return builder
         * 
         */
        public Builder tmpdir(@Nullable Output<String> tmpdir) {
            $.tmpdir = tmpdir;
            return this;
        }

        /**
         * @param tmpdir Corresponds to the `--tmpdir` option.
         * 
         * @return builder
         * 
         */
        public Builder tmpdir(String tmpdir) {
            return tmpdir(Output.of(tmpdir));
        }

        /**
         * @param triggers TODO
         * 
         * @return builder
         * 
         */
        public Builder triggers(@Nullable Output<List<Object>> triggers) {
            $.triggers = triggers;
            return this;
        }

        /**
         * @param triggers TODO
         * 
         * @return builder
         * 
         */
        public Builder triggers(List<Object> triggers) {
            return triggers(Output.of(triggers));
        }

        /**
         * @param triggers TODO
         * 
         * @return builder
         * 
         */
        public Builder triggers(Object... triggers) {
            return triggers(List.of(triggers));
        }

        public MktempArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("MktempArgs", "connection");
            }
            return $;
        }
    }

}
