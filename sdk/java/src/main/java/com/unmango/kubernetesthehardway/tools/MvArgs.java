// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MvArgs extends com.pulumi.resources.ResourceArgs {

    public static final MvArgs Empty = new MvArgs();

    /**
     * Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
     * 
     */
    @Import(name="backup")
    private @Nullable Boolean backup;

    /**
     * @return Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
     * 
     */
    public Optional<Boolean> backup() {
        return Optional.ofNullable(this.backup);
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
     * Corresponds to the --context option.
     * 
     */
    @Import(name="context")
    private @Nullable Output<Boolean> context;

    /**
     * @return Corresponds to the --context option.
     * 
     */
    public Optional<Output<Boolean>> context() {
        return Optional.ofNullable(this.context);
    }

    /**
     * Corresponds to the [CONTROL] argument for the --backup option.
     * 
     */
    @Import(name="control")
    private @Nullable Output<String> control;

    /**
     * @return Corresponds to the [CONTROL] argument for the --backup option.
     * 
     */
    public Optional<Output<String>> control() {
        return Optional.ofNullable(this.control);
    }

    /**
     * Corresponds to the [DEST] argument.
     * 
     */
    @Import(name="dest")
    private @Nullable Output<String> dest;

    /**
     * @return Corresponds to the [DEST] argument.
     * 
     */
    public Optional<Output<String>> dest() {
        return Optional.ofNullable(this.dest);
    }

    /**
     * Corresponds to the [DIRECTORY] argument.
     * 
     */
    @Import(name="directory")
    private @Nullable Output<String> directory;

    /**
     * @return Corresponds to the [DIRECTORY] argument.
     * 
     */
    public Optional<Output<String>> directory() {
        return Optional.ofNullable(this.directory);
    }

    @Import(name="environment")
    private @Nullable Output<Map<String,String>> environment;

    public Optional<Output<Map<String,String>>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * Corresponds to the --force option.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return Corresponds to the --force option.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    /**
     * Corresponds to the --no-clobber option.
     * 
     */
    @Import(name="noClobber")
    private @Nullable Output<Boolean> noClobber;

    /**
     * @return Corresponds to the --no-clobber option.
     * 
     */
    public Optional<Output<Boolean>> noClobber() {
        return Optional.ofNullable(this.noClobber);
    }

    /**
     * Corresponds to the --no-target-directory option.
     * 
     */
    @Import(name="noTargetDirectory")
    private @Nullable Output<Boolean> noTargetDirectory;

    /**
     * @return Corresponds to the --no-target-directory option.
     * 
     */
    public Optional<Output<Boolean>> noTargetDirectory() {
        return Optional.ofNullable(this.noTargetDirectory);
    }

    /**
     * Corresponds to the [SOURCE] argument.
     * 
     */
    @Import(name="source", required=true)
    private Output<Either<String,List<String>>> source;

    /**
     * @return Corresponds to the [SOURCE] argument.
     * 
     */
    public Output<Either<String,List<String>>> source() {
        return this.source;
    }

    /**
     * Corresponds to the --strip-trailing-suffix option.
     * 
     */
    @Import(name="stripTrailingSlashes")
    private @Nullable Output<Boolean> stripTrailingSlashes;

    /**
     * @return Corresponds to the --strip-trailing-suffix option.
     * 
     */
    public Optional<Output<Boolean>> stripTrailingSlashes() {
        return Optional.ofNullable(this.stripTrailingSlashes);
    }

    /**
     * Corresponds to the --suffix option.
     * 
     */
    @Import(name="suffix")
    private @Nullable Output<String> suffix;

    /**
     * @return Corresponds to the --suffix option.
     * 
     */
    public Optional<Output<String>> suffix() {
        return Optional.ofNullable(this.suffix);
    }

    /**
     * Corresponds to the --target-directory option.
     * 
     */
    @Import(name="targetDirectory")
    private @Nullable Output<String> targetDirectory;

    /**
     * @return Corresponds to the --target-directory option.
     * 
     */
    public Optional<Output<String>> targetDirectory() {
        return Optional.ofNullable(this.targetDirectory);
    }

    /**
     * Corresponds to the --update option.
     * 
     */
    @Import(name="update")
    private @Nullable Output<Boolean> update;

    /**
     * @return Corresponds to the --update option.
     * 
     */
    public Optional<Output<Boolean>> update() {
        return Optional.ofNullable(this.update);
    }

    /**
     * Corresponds to the --verbose option.
     * 
     */
    @Import(name="verbose")
    private @Nullable Output<Boolean> verbose;

    /**
     * @return Corresponds to the --verbose option.
     * 
     */
    public Optional<Output<Boolean>> verbose() {
        return Optional.ofNullable(this.verbose);
    }

    private MvArgs() {}

    private MvArgs(MvArgs $) {
        this.backup = $.backup;
        this.connection = $.connection;
        this.context = $.context;
        this.control = $.control;
        this.dest = $.dest;
        this.directory = $.directory;
        this.environment = $.environment;
        this.force = $.force;
        this.noClobber = $.noClobber;
        this.noTargetDirectory = $.noTargetDirectory;
        this.source = $.source;
        this.stripTrailingSlashes = $.stripTrailingSlashes;
        this.suffix = $.suffix;
        this.targetDirectory = $.targetDirectory;
        this.update = $.update;
        this.verbose = $.verbose;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MvArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MvArgs $;

        public Builder() {
            $ = new MvArgs();
        }

        public Builder(MvArgs defaults) {
            $ = new MvArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backup Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
         * 
         * @return builder
         * 
         */
        public Builder backup(@Nullable Boolean backup) {
            $.backup = backup;
            return this;
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
         * @param context Corresponds to the --context option.
         * 
         * @return builder
         * 
         */
        public Builder context(@Nullable Output<Boolean> context) {
            $.context = context;
            return this;
        }

        /**
         * @param context Corresponds to the --context option.
         * 
         * @return builder
         * 
         */
        public Builder context(Boolean context) {
            return context(Output.of(context));
        }

        /**
         * @param control Corresponds to the [CONTROL] argument for the --backup option.
         * 
         * @return builder
         * 
         */
        public Builder control(@Nullable Output<String> control) {
            $.control = control;
            return this;
        }

        /**
         * @param control Corresponds to the [CONTROL] argument for the --backup option.
         * 
         * @return builder
         * 
         */
        public Builder control(String control) {
            return control(Output.of(control));
        }

        /**
         * @param dest Corresponds to the [DEST] argument.
         * 
         * @return builder
         * 
         */
        public Builder dest(@Nullable Output<String> dest) {
            $.dest = dest;
            return this;
        }

        /**
         * @param dest Corresponds to the [DEST] argument.
         * 
         * @return builder
         * 
         */
        public Builder dest(String dest) {
            return dest(Output.of(dest));
        }

        /**
         * @param directory Corresponds to the [DIRECTORY] argument.
         * 
         * @return builder
         * 
         */
        public Builder directory(@Nullable Output<String> directory) {
            $.directory = directory;
            return this;
        }

        /**
         * @param directory Corresponds to the [DIRECTORY] argument.
         * 
         * @return builder
         * 
         */
        public Builder directory(String directory) {
            return directory(Output.of(directory));
        }

        public Builder environment(@Nullable Output<Map<String,String>> environment) {
            $.environment = environment;
            return this;
        }

        public Builder environment(Map<String,String> environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param force Corresponds to the --force option.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force Corresponds to the --force option.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        /**
         * @param noClobber Corresponds to the --no-clobber option.
         * 
         * @return builder
         * 
         */
        public Builder noClobber(@Nullable Output<Boolean> noClobber) {
            $.noClobber = noClobber;
            return this;
        }

        /**
         * @param noClobber Corresponds to the --no-clobber option.
         * 
         * @return builder
         * 
         */
        public Builder noClobber(Boolean noClobber) {
            return noClobber(Output.of(noClobber));
        }

        /**
         * @param noTargetDirectory Corresponds to the --no-target-directory option.
         * 
         * @return builder
         * 
         */
        public Builder noTargetDirectory(@Nullable Output<Boolean> noTargetDirectory) {
            $.noTargetDirectory = noTargetDirectory;
            return this;
        }

        /**
         * @param noTargetDirectory Corresponds to the --no-target-directory option.
         * 
         * @return builder
         * 
         */
        public Builder noTargetDirectory(Boolean noTargetDirectory) {
            return noTargetDirectory(Output.of(noTargetDirectory));
        }

        /**
         * @param source Corresponds to the [SOURCE] argument.
         * 
         * @return builder
         * 
         */
        public Builder source(Output<Either<String,List<String>>> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source Corresponds to the [SOURCE] argument.
         * 
         * @return builder
         * 
         */
        public Builder source(Either<String,List<String>> source) {
            return source(Output.of(source));
        }

        /**
         * @param source Corresponds to the [SOURCE] argument.
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Either.ofLeft(source));
        }

        /**
         * @param source Corresponds to the [SOURCE] argument.
         * 
         * @return builder
         * 
         */
        public Builder source(List<String> source) {
            return source(Either.ofRight(source));
        }

        /**
         * @param stripTrailingSlashes Corresponds to the --strip-trailing-suffix option.
         * 
         * @return builder
         * 
         */
        public Builder stripTrailingSlashes(@Nullable Output<Boolean> stripTrailingSlashes) {
            $.stripTrailingSlashes = stripTrailingSlashes;
            return this;
        }

        /**
         * @param stripTrailingSlashes Corresponds to the --strip-trailing-suffix option.
         * 
         * @return builder
         * 
         */
        public Builder stripTrailingSlashes(Boolean stripTrailingSlashes) {
            return stripTrailingSlashes(Output.of(stripTrailingSlashes));
        }

        /**
         * @param suffix Corresponds to the --suffix option.
         * 
         * @return builder
         * 
         */
        public Builder suffix(@Nullable Output<String> suffix) {
            $.suffix = suffix;
            return this;
        }

        /**
         * @param suffix Corresponds to the --suffix option.
         * 
         * @return builder
         * 
         */
        public Builder suffix(String suffix) {
            return suffix(Output.of(suffix));
        }

        /**
         * @param targetDirectory Corresponds to the --target-directory option.
         * 
         * @return builder
         * 
         */
        public Builder targetDirectory(@Nullable Output<String> targetDirectory) {
            $.targetDirectory = targetDirectory;
            return this;
        }

        /**
         * @param targetDirectory Corresponds to the --target-directory option.
         * 
         * @return builder
         * 
         */
        public Builder targetDirectory(String targetDirectory) {
            return targetDirectory(Output.of(targetDirectory));
        }

        /**
         * @param update Corresponds to the --update option.
         * 
         * @return builder
         * 
         */
        public Builder update(@Nullable Output<Boolean> update) {
            $.update = update;
            return this;
        }

        /**
         * @param update Corresponds to the --update option.
         * 
         * @return builder
         * 
         */
        public Builder update(Boolean update) {
            return update(Output.of(update));
        }

        /**
         * @param verbose Corresponds to the --verbose option.
         * 
         * @return builder
         * 
         */
        public Builder verbose(@Nullable Output<Boolean> verbose) {
            $.verbose = verbose;
            return this;
        }

        /**
         * @param verbose Corresponds to the --verbose option.
         * 
         * @return builder
         * 
         */
        public Builder verbose(Boolean verbose) {
            return verbose(Output.of(verbose));
        }

        public MvArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("MvArgs", "connection");
            }
            if ($.source == null) {
                throw new MissingRequiredPropertyException("MvArgs", "source");
            }
            return $;
        }
    }

}
