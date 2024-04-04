// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.enums.CommandLifecycle;
import com.unmango.kubernetesthehardway.tools.enums.TeeMode;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeeArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeeArgs Empty = new TeeArgs();

    /**
     * Append to the given FILEs, do not overwrite.
     * 
     */
    @Import(name="append")
    private @Nullable Output<Boolean> append;

    /**
     * @return Append to the given FILEs, do not overwrite.
     * 
     */
    public Optional<Output<Boolean>> append() {
        return Optional.ofNullable(this.append);
    }

    /**
     * Connection details for the remote system.
     * 
     */
    @Import(name="connection", required=true)
    private Output<ConnectionArgs> connection;

    /**
     * @return Connection details for the remote system.
     * 
     */
    public Output<ConnectionArgs> connection() {
        return this.connection;
    }

    /**
     * The file(s) to write to.
     * 
     */
    @Import(name="files", required=true)
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
    @Import(name="ignoreInterrupts")
    private @Nullable Output<Boolean> ignoreInterrupts;

    /**
     * @return Ignore interrupt signals.
     * 
     */
    public Optional<Output<Boolean>> ignoreInterrupts() {
        return Optional.ofNullable(this.ignoreInterrupts);
    }

    /**
     * At what stage(s) in the resource lifecycle should the command be run.
     * 
     */
    @Import(name="lifecycle")
    private @Nullable Either<CommandLifecycle,List<CommandLifecycle>> lifecycle;

    /**
     * @return At what stage(s) in the resource lifecycle should the command be run.
     * 
     */
    public Optional<Either<CommandLifecycle,List<CommandLifecycle>>> lifecycle() {
        return Optional.ofNullable(this.lifecycle);
    }

    /**
     * Set behavior on write error.
     * 
     */
    @Import(name="outputError")
    private @Nullable Output<TeeMode> outputError;

    /**
     * @return Set behavior on write error.
     * 
     */
    public Optional<Output<TeeMode>> outputError() {
        return Optional.ofNullable(this.outputError);
    }

    /**
     * Operate in a more appropriate MODE with pipes.
     * 
     */
    @Import(name="pipe")
    private @Nullable Output<Boolean> pipe;

    /**
     * @return Operate in a more appropriate MODE with pipes.
     * 
     */
    public Optional<Output<Boolean>> pipe() {
        return Optional.ofNullable(this.pipe);
    }

    @Import(name="stdin", required=true)
    private Output<String> stdin;

    public Output<String> stdin() {
        return this.stdin;
    }

    /**
     * Output version information and exit.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Boolean> version;

    /**
     * @return Output version information and exit.
     * 
     */
    public Optional<Output<Boolean>> version() {
        return Optional.ofNullable(this.version);
    }

    private TeeArgs() {}

    private TeeArgs(TeeArgs $) {
        this.append = $.append;
        this.connection = $.connection;
        this.files = $.files;
        this.ignoreInterrupts = $.ignoreInterrupts;
        this.lifecycle = $.lifecycle;
        this.outputError = $.outputError;
        this.pipe = $.pipe;
        this.stdin = $.stdin;
        this.version = $.version;
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
         * @param append Append to the given FILEs, do not overwrite.
         * 
         * @return builder
         * 
         */
        public Builder append(@Nullable Output<Boolean> append) {
            $.append = append;
            return this;
        }

        /**
         * @param append Append to the given FILEs, do not overwrite.
         * 
         * @return builder
         * 
         */
        public Builder append(Boolean append) {
            return append(Output.of(append));
        }

        /**
         * @param connection Connection details for the remote system.
         * 
         * @return builder
         * 
         */
        public Builder connection(Output<ConnectionArgs> connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param connection Connection details for the remote system.
         * 
         * @return builder
         * 
         */
        public Builder connection(ConnectionArgs connection) {
            return connection(Output.of(connection));
        }

        /**
         * @param files The file(s) to write to.
         * 
         * @return builder
         * 
         */
        public Builder files(Output<Either<String,List<String>>> files) {
            $.files = files;
            return this;
        }

        /**
         * @param files The file(s) to write to.
         * 
         * @return builder
         * 
         */
        public Builder files(Either<String,List<String>> files) {
            return files(Output.of(files));
        }

        /**
         * @param files The file(s) to write to.
         * 
         * @return builder
         * 
         */
        public Builder files(String files) {
            return files(Either.ofLeft(files));
        }

        /**
         * @param files The file(s) to write to.
         * 
         * @return builder
         * 
         */
        public Builder files(List<String> files) {
            return files(Either.ofRight(files));
        }

        /**
         * @param ignoreInterrupts Ignore interrupt signals.
         * 
         * @return builder
         * 
         */
        public Builder ignoreInterrupts(@Nullable Output<Boolean> ignoreInterrupts) {
            $.ignoreInterrupts = ignoreInterrupts;
            return this;
        }

        /**
         * @param ignoreInterrupts Ignore interrupt signals.
         * 
         * @return builder
         * 
         */
        public Builder ignoreInterrupts(Boolean ignoreInterrupts) {
            return ignoreInterrupts(Output.of(ignoreInterrupts));
        }

        /**
         * @param lifecycle At what stage(s) in the resource lifecycle should the command be run.
         * 
         * @return builder
         * 
         */
        public Builder lifecycle(@Nullable Either<CommandLifecycle,List<CommandLifecycle>> lifecycle) {
            $.lifecycle = lifecycle;
            return this;
        }

        /**
         * @param lifecycle At what stage(s) in the resource lifecycle should the command be run.
         * 
         * @return builder
         * 
         */
        public Builder lifecycle(CommandLifecycle lifecycle) {
            return lifecycle(Either.ofLeft(lifecycle));
        }

        /**
         * @param lifecycle At what stage(s) in the resource lifecycle should the command be run.
         * 
         * @return builder
         * 
         */
        public Builder lifecycle(List<CommandLifecycle> lifecycle) {
            return lifecycle(Either.ofRight(lifecycle));
        }

        /**
         * @param outputError Set behavior on write error.
         * 
         * @return builder
         * 
         */
        public Builder outputError(@Nullable Output<TeeMode> outputError) {
            $.outputError = outputError;
            return this;
        }

        /**
         * @param outputError Set behavior on write error.
         * 
         * @return builder
         * 
         */
        public Builder outputError(TeeMode outputError) {
            return outputError(Output.of(outputError));
        }

        /**
         * @param pipe Operate in a more appropriate MODE with pipes.
         * 
         * @return builder
         * 
         */
        public Builder pipe(@Nullable Output<Boolean> pipe) {
            $.pipe = pipe;
            return this;
        }

        /**
         * @param pipe Operate in a more appropriate MODE with pipes.
         * 
         * @return builder
         * 
         */
        public Builder pipe(Boolean pipe) {
            return pipe(Output.of(pipe));
        }

        public Builder stdin(Output<String> stdin) {
            $.stdin = stdin;
            return this;
        }

        public Builder stdin(String stdin) {
            return stdin(Output.of(stdin));
        }

        /**
         * @param version Output version information and exit.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Boolean> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Output version information and exit.
         * 
         * @return builder
         * 
         */
        public Builder version(Boolean version) {
            return version(Output.of(version));
        }

        public TeeArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("TeeArgs", "connection");
            }
            if ($.files == null) {
                throw new MissingRequiredPropertyException("TeeArgs", "files");
            }
            if ($.stdin == null) {
                throw new MissingRequiredPropertyException("TeeArgs", "stdin");
            }
            return $;
        }
    }

}
