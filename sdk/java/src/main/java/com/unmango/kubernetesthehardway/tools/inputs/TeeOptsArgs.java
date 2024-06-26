// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.enums.TeeMode;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * Abstraction over the `rm` utility on a remote system.
 * 
 */
public final class TeeOptsArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeeOptsArgs Empty = new TeeOptsArgs();

    /**
     * Append to the given FILEs, do not overwrite
     * 
     */
    @Import(name="append")
    private @Nullable Output<Boolean> append;

    /**
     * @return Append to the given FILEs, do not overwrite
     * 
     */
    public Optional<Output<Boolean>> append() {
        return Optional.ofNullable(this.append);
    }

    /**
     * Corresponds to the [FILE] argument.
     * 
     */
    @Import(name="files", required=true)
    private Output<List<String>> files;

    /**
     * @return Corresponds to the [FILE] argument.
     * 
     */
    public Output<List<String>> files() {
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

    private TeeOptsArgs() {}

    private TeeOptsArgs(TeeOptsArgs $) {
        this.append = $.append;
        this.files = $.files;
        this.ignoreInterrupts = $.ignoreInterrupts;
        this.outputError = $.outputError;
        this.pipe = $.pipe;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeeOptsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeeOptsArgs $;

        public Builder() {
            $ = new TeeOptsArgs();
        }

        public Builder(TeeOptsArgs defaults) {
            $ = new TeeOptsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param append Append to the given FILEs, do not overwrite
         * 
         * @return builder
         * 
         */
        public Builder append(@Nullable Output<Boolean> append) {
            $.append = append;
            return this;
        }

        /**
         * @param append Append to the given FILEs, do not overwrite
         * 
         * @return builder
         * 
         */
        public Builder append(Boolean append) {
            return append(Output.of(append));
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(Output<List<String>> files) {
            $.files = files;
            return this;
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(List<String> files) {
            return files(Output.of(files));
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(String... files) {
            return files(List.of(files));
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

        public TeeOptsArgs build() {
            if ($.files == null) {
                throw new MissingRequiredPropertyException("TeeOptsArgs", "files");
            }
            return $;
        }
    }

}
