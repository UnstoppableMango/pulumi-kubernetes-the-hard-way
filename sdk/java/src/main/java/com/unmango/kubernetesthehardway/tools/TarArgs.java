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


public final class TarArgs extends com.pulumi.resources.ResourceArgs {

    public static final TarArgs Empty = new TarArgs();

    /**
     * Corresponds to the [ARCHIVE] argument.
     * 
     */
    @Import(name="archive", required=true)
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
    private @Nullable Output<String> directory;

    /**
     * @return Corresponds to the `--directory` option.
     * 
     */
    public Optional<Output<String>> directory() {
        return Optional.ofNullable(this.directory);
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
     * Corresponds to the `--extract` option.
     * 
     */
    @Import(name="extract")
    private @Nullable Output<Boolean> extract;

    /**
     * @return Corresponds to the `--extract` option.
     * 
     */
    public Optional<Output<Boolean>> extract() {
        return Optional.ofNullable(this.extract);
    }

    /**
     * Corresponds to the [FILE] argument.
     * 
     */
    @Import(name="files")
    private @Nullable Output<Either<String,List<String>>> files;

    /**
     * @return Corresponds to the [FILE] argument.
     * 
     */
    public Optional<Output<Either<String,List<String>>>> files() {
        return Optional.ofNullable(this.files);
    }

    /**
     * Corresponds to the `--gzip` option.
     * 
     */
    @Import(name="gzip")
    private @Nullable Output<Boolean> gzip;

    /**
     * @return Corresponds to the `--gzip` option.
     * 
     */
    public Optional<Output<Boolean>> gzip() {
        return Optional.ofNullable(this.gzip);
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
     * Corresponds to the `--strip-components` option.
     * 
     */
    @Import(name="stripComponents")
    private @Nullable Output<Boolean> stripComponents;

    /**
     * @return Corresponds to the `--strip-components` option.
     * 
     */
    public Optional<Output<Boolean>> stripComponents() {
        return Optional.ofNullable(this.stripComponents);
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

    private TarArgs() {}

    private TarArgs(TarArgs $) {
        this.archive = $.archive;
        this.binaryPath = $.binaryPath;
        this.connection = $.connection;
        this.directory = $.directory;
        this.environment = $.environment;
        this.extract = $.extract;
        this.files = $.files;
        this.gzip = $.gzip;
        this.lifecycle = $.lifecycle;
        this.onDelete = $.onDelete;
        this.recursive = $.recursive;
        this.stdin = $.stdin;
        this.stripComponents = $.stripComponents;
        this.triggers = $.triggers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TarArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TarArgs $;

        public Builder() {
            $ = new TarArgs();
        }

        public Builder(TarArgs defaults) {
            $ = new TarArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param archive Corresponds to the [ARCHIVE] argument.
         * 
         * @return builder
         * 
         */
        public Builder archive(Output<String> archive) {
            $.archive = archive;
            return this;
        }

        /**
         * @param archive Corresponds to the [ARCHIVE] argument.
         * 
         * @return builder
         * 
         */
        public Builder archive(String archive) {
            return archive(Output.of(archive));
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
        public Builder directory(@Nullable Output<String> directory) {
            $.directory = directory;
            return this;
        }

        /**
         * @param directory Corresponds to the `--directory` option.
         * 
         * @return builder
         * 
         */
        public Builder directory(String directory) {
            return directory(Output.of(directory));
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
         * @param extract Corresponds to the `--extract` option.
         * 
         * @return builder
         * 
         */
        public Builder extract(@Nullable Output<Boolean> extract) {
            $.extract = extract;
            return this;
        }

        /**
         * @param extract Corresponds to the `--extract` option.
         * 
         * @return builder
         * 
         */
        public Builder extract(Boolean extract) {
            return extract(Output.of(extract));
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(@Nullable Output<Either<String,List<String>>> files) {
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
         * @param gzip Corresponds to the `--gzip` option.
         * 
         * @return builder
         * 
         */
        public Builder gzip(@Nullable Output<Boolean> gzip) {
            $.gzip = gzip;
            return this;
        }

        /**
         * @param gzip Corresponds to the `--gzip` option.
         * 
         * @return builder
         * 
         */
        public Builder gzip(Boolean gzip) {
            return gzip(Output.of(gzip));
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
         * @param stripComponents Corresponds to the `--strip-components` option.
         * 
         * @return builder
         * 
         */
        public Builder stripComponents(@Nullable Output<Boolean> stripComponents) {
            $.stripComponents = stripComponents;
            return this;
        }

        /**
         * @param stripComponents Corresponds to the `--strip-components` option.
         * 
         * @return builder
         * 
         */
        public Builder stripComponents(Boolean stripComponents) {
            return stripComponents(Output.of(stripComponents));
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

        public TarArgs build() {
            if ($.archive == null) {
                throw new MissingRequiredPropertyException("TarArgs", "archive");
            }
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("TarArgs", "connection");
            }
            return $;
        }
    }

}
