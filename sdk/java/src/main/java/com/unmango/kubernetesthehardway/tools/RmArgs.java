// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.enums.CommandLifecycle;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RmArgs extends com.pulumi.resources.ResourceArgs {

    public static final RmArgs Empty = new RmArgs();

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
     * Corresponds to the `--dir` option.
     * 
     */
    @Import(name="dir")
    private @Nullable Output<Boolean> dir;

    /**
     * @return Corresponds to the `--dir` option.
     * 
     */
    public Optional<Output<Boolean>> dir() {
        return Optional.ofNullable(this.dir);
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
     * Corresponds to the [FILE] argument.
     * 
     */
    @Import(name="files", required=true)
    private Output<Either<String,List<String>>> files;

    /**
     * @return Corresponds to the [FILE] argument.
     * 
     */
    public Output<Either<String,List<String>>> files() {
        return this.files;
    }

    /**
     * Corresponds to the `--force` option.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return Corresponds to the `--force` option.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
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
     * Whether rm should be run when the resource is created or deleted.
     * 
     */
    @Import(name="onDelete")
    private @Nullable Output<Boolean> onDelete;

    /**
     * @return Whether rm should be run when the resource is created or deleted.
     * 
     */
    public Optional<Output<Boolean>> onDelete() {
        return Optional.ofNullable(this.onDelete);
    }

    /**
     * Corresponds to the `--recursive` option.
     * 
     */
    @Import(name="recursive")
    private @Nullable Output<Boolean> recursive;

    /**
     * @return Corresponds to the `--recursive` option.
     * 
     */
    public Optional<Output<Boolean>> recursive() {
        return Optional.ofNullable(this.recursive);
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

    /**
     * Corresponds to the `--verbose` option.
     * 
     */
    @Import(name="verbose")
    private @Nullable Output<Boolean> verbose;

    /**
     * @return Corresponds to the `--verbose` option.
     * 
     */
    public Optional<Output<Boolean>> verbose() {
        return Optional.ofNullable(this.verbose);
    }

    private RmArgs() {}

    private RmArgs(RmArgs $) {
        this.binaryPath = $.binaryPath;
        this.connection = $.connection;
        this.dir = $.dir;
        this.environment = $.environment;
        this.files = $.files;
        this.force = $.force;
        this.lifecycle = $.lifecycle;
        this.onDelete = $.onDelete;
        this.recursive = $.recursive;
        this.stdin = $.stdin;
        this.triggers = $.triggers;
        this.verbose = $.verbose;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RmArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RmArgs $;

        public Builder() {
            $ = new RmArgs();
        }

        public Builder(RmArgs defaults) {
            $ = new RmArgs(Objects.requireNonNull(defaults));
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
         * @param dir Corresponds to the `--dir` option.
         * 
         * @return builder
         * 
         */
        public Builder dir(@Nullable Output<Boolean> dir) {
            $.dir = dir;
            return this;
        }

        /**
         * @param dir Corresponds to the `--dir` option.
         * 
         * @return builder
         * 
         */
        public Builder dir(Boolean dir) {
            return dir(Output.of(dir));
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
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(Output<Either<String,List<String>>> files) {
            $.files = files;
            return this;
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(Either<String,List<String>> files) {
            return files(Output.of(files));
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(String files) {
            return files(Either.ofLeft(files));
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(List<String> files) {
            return files(Either.ofRight(files));
        }

        /**
         * @param force Corresponds to the `--force` option.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force Corresponds to the `--force` option.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
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
         * @param onDelete Whether rm should be run when the resource is created or deleted.
         * 
         * @return builder
         * 
         */
        public Builder onDelete(@Nullable Output<Boolean> onDelete) {
            $.onDelete = onDelete;
            return this;
        }

        /**
         * @param onDelete Whether rm should be run when the resource is created or deleted.
         * 
         * @return builder
         * 
         */
        public Builder onDelete(Boolean onDelete) {
            return onDelete(Output.of(onDelete));
        }

        /**
         * @param recursive Corresponds to the `--recursive` option.
         * 
         * @return builder
         * 
         */
        public Builder recursive(@Nullable Output<Boolean> recursive) {
            $.recursive = recursive;
            return this;
        }

        /**
         * @param recursive Corresponds to the `--recursive` option.
         * 
         * @return builder
         * 
         */
        public Builder recursive(Boolean recursive) {
            return recursive(Output.of(recursive));
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

        /**
         * @param verbose Corresponds to the `--verbose` option.
         * 
         * @return builder
         * 
         */
        public Builder verbose(@Nullable Output<Boolean> verbose) {
            $.verbose = verbose;
            return this;
        }

        /**
         * @param verbose Corresponds to the `--verbose` option.
         * 
         * @return builder
         * 
         */
        public Builder verbose(Boolean verbose) {
            return verbose(Output.of(verbose));
        }

        public RmArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("RmArgs", "connection");
            }
            if ($.files == null) {
                throw new MissingRequiredPropertyException("RmArgs", "files");
            }
            return $;
        }
    }

}
