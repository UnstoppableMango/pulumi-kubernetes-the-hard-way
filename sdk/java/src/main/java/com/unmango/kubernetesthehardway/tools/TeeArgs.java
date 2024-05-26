// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.inputs.TeeOptsArgs;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeeArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeeArgs Empty = new TeeArgs();

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
     * The command to run on create.
     * 
     */
    @Import(name="create")
    private @Nullable Either<String,TeeOptsArgs> create;

    /**
     * @return The command to run on create.
     * 
     */
    public Optional<Either<String,TeeOptsArgs>> create() {
        return Optional.ofNullable(this.create);
    }

    /**
     * The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
     * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
     * Command resource from previous create or update steps.
     * 
     */
    @Import(name="delete")
    private @Nullable Either<String,TeeOptsArgs> delete;

    /**
     * @return The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
     * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
     * Command resource from previous create or update steps.
     * 
     */
    public Optional<Either<String,TeeOptsArgs>> delete() {
        return Optional.ofNullable(this.delete);
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
     * Trigger replacements on changes to this input.
     * 
     */
    @Import(name="triggers")
    private @Nullable Output<List<Object>> triggers;

    /**
     * @return Trigger replacements on changes to this input.
     * 
     */
    public Optional<Output<List<Object>>> triggers() {
        return Optional.ofNullable(this.triggers);
    }

    /**
     * The command to run on update, if empty, create will
     * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
     * are set to the stdout and stderr properties of the Command resource from previous
     * create or update steps.
     * 
     */
    @Import(name="update")
    private @Nullable Either<String,TeeOptsArgs> update;

    /**
     * @return The command to run on update, if empty, create will
     * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
     * are set to the stdout and stderr properties of the Command resource from previous
     * create or update steps.
     * 
     */
    public Optional<Either<String,TeeOptsArgs>> update() {
        return Optional.ofNullable(this.update);
    }

    private TeeArgs() {}

    private TeeArgs(TeeArgs $) {
        this.binaryPath = $.binaryPath;
        this.connection = $.connection;
        this.create = $.create;
        this.delete = $.delete;
        this.environment = $.environment;
        this.stdin = $.stdin;
        this.triggers = $.triggers;
        this.update = $.update;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeeArgs $;

        public Builder() {
            $ = new TeeArgs();
        }

        public Builder(TeeArgs defaults) {
            $ = new TeeArgs(Objects.requireNonNull(defaults));
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
         * @param create The command to run on create.
         * 
         * @return builder
         * 
         */
        public Builder create(@Nullable Either<String,TeeOptsArgs> create) {
            $.create = create;
            return this;
        }

        /**
         * @param create The command to run on create.
         * 
         * @return builder
         * 
         */
        public Builder create(String create) {
            return create(Either.ofLeft(create));
        }

        /**
         * @param create The command to run on create.
         * 
         * @return builder
         * 
         */
        public Builder create(TeeOptsArgs create) {
            return create(Either.ofRight(create));
        }

        /**
         * @param delete The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
         * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
         * Command resource from previous create or update steps.
         * 
         * @return builder
         * 
         */
        public Builder delete(@Nullable Either<String,TeeOptsArgs> delete) {
            $.delete = delete;
            return this;
        }

        /**
         * @param delete The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
         * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
         * Command resource from previous create or update steps.
         * 
         * @return builder
         * 
         */
        public Builder delete(String delete) {
            return delete(Either.ofLeft(delete));
        }

        /**
         * @param delete The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
         * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
         * Command resource from previous create or update steps.
         * 
         * @return builder
         * 
         */
        public Builder delete(TeeOptsArgs delete) {
            return delete(Either.ofRight(delete));
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
         * @param triggers Trigger replacements on changes to this input.
         * 
         * @return builder
         * 
         */
        public Builder triggers(@Nullable Output<List<Object>> triggers) {
            $.triggers = triggers;
            return this;
        }

        /**
         * @param triggers Trigger replacements on changes to this input.
         * 
         * @return builder
         * 
         */
        public Builder triggers(List<Object> triggers) {
            return triggers(Output.of(triggers));
        }

        /**
         * @param triggers Trigger replacements on changes to this input.
         * 
         * @return builder
         * 
         */
        public Builder triggers(Object... triggers) {
            return triggers(List.of(triggers));
        }

        /**
         * @param update The command to run on update, if empty, create will
         * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
         * are set to the stdout and stderr properties of the Command resource from previous
         * create or update steps.
         * 
         * @return builder
         * 
         */
        public Builder update(@Nullable Either<String,TeeOptsArgs> update) {
            $.update = update;
            return this;
        }

        /**
         * @param update The command to run on update, if empty, create will
         * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
         * are set to the stdout and stderr properties of the Command resource from previous
         * create or update steps.
         * 
         * @return builder
         * 
         */
        public Builder update(String update) {
            return update(Either.ofLeft(update));
        }

        /**
         * @param update The command to run on update, if empty, create will
         * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
         * are set to the stdout and stderr properties of the Command resource from previous
         * create or update steps.
         * 
         * @return builder
         * 
         */
        public Builder update(TeeOptsArgs update) {
            return update(Either.ofRight(update));
        }

        public TeeArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("TeeArgs", "connection");
            }
            return $;
        }
    }

}
