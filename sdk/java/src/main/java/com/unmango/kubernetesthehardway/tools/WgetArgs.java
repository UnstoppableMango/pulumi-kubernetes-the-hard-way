// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Either;
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


public final class WgetArgs extends com.pulumi.resources.ResourceArgs {

    public static final WgetArgs Empty = new WgetArgs();

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
     * The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
     * 
     */
    @Import(name="directoryPrefix")
    private @Nullable Output<String> directoryPrefix;

    /**
     * @return The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
     * 
     */
    public Optional<Output<String>> directoryPrefix() {
        return Optional.ofNullable(this.directoryPrefix);
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
     * When in recursive mode, only HTTPS links are followed.
     * 
     */
    @Import(name="httpsOnly")
    private @Nullable Output<Boolean> httpsOnly;

    /**
     * @return When in recursive mode, only HTTPS links are followed.
     * 
     */
    public Optional<Output<Boolean>> httpsOnly() {
        return Optional.ofNullable(this.httpsOnly);
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
     * Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
     * 
     */
    @Import(name="noVerbose")
    private @Nullable Output<Boolean> noVerbose;

    /**
     * @return Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
     * 
     */
    public Optional<Output<Boolean>> noVerbose() {
        return Optional.ofNullable(this.noVerbose);
    }

    /**
     * The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
     * 
     */
    @Import(name="outputDocument")
    private @Nullable Output<String> outputDocument;

    /**
     * @return The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
     * 
     */
    public Optional<Output<String>> outputDocument() {
        return Optional.ofNullable(this.outputDocument);
    }

    /**
     * Turn off Wget&#39;s output.
     * 
     */
    @Import(name="quiet")
    private @Nullable Output<Boolean> quiet;

    /**
     * @return Turn off Wget&#39;s output.
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
     * Turn on time-stamping.
     * 
     */
    @Import(name="timestamping")
    private @Nullable Output<Boolean> timestamping;

    /**
     * @return Turn on time-stamping.
     * 
     */
    public Optional<Output<Boolean>> timestamping() {
        return Optional.ofNullable(this.timestamping);
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
     * Corresponds to the [URL...] argument.
     * 
     */
    @Import(name="url", required=true)
    private Output<Either<String,List<String>>> url;

    /**
     * @return Corresponds to the [URL...] argument.
     * 
     */
    public Output<Either<String,List<String>>> url() {
        return this.url;
    }

    private WgetArgs() {}

    private WgetArgs(WgetArgs $) {
        this.binaryPath = $.binaryPath;
        this.connection = $.connection;
        this.directoryPrefix = $.directoryPrefix;
        this.environment = $.environment;
        this.httpsOnly = $.httpsOnly;
        this.lifecycle = $.lifecycle;
        this.noVerbose = $.noVerbose;
        this.outputDocument = $.outputDocument;
        this.quiet = $.quiet;
        this.stdin = $.stdin;
        this.timestamping = $.timestamping;
        this.triggers = $.triggers;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WgetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WgetArgs $;

        public Builder() {
            $ = new WgetArgs();
        }

        public Builder(WgetArgs defaults) {
            $ = new WgetArgs(Objects.requireNonNull(defaults));
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
         * @param directoryPrefix The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
         * 
         * @return builder
         * 
         */
        public Builder directoryPrefix(@Nullable Output<String> directoryPrefix) {
            $.directoryPrefix = directoryPrefix;
            return this;
        }

        /**
         * @param directoryPrefix The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
         * 
         * @return builder
         * 
         */
        public Builder directoryPrefix(String directoryPrefix) {
            return directoryPrefix(Output.of(directoryPrefix));
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
         * @param httpsOnly When in recursive mode, only HTTPS links are followed.
         * 
         * @return builder
         * 
         */
        public Builder httpsOnly(@Nullable Output<Boolean> httpsOnly) {
            $.httpsOnly = httpsOnly;
            return this;
        }

        /**
         * @param httpsOnly When in recursive mode, only HTTPS links are followed.
         * 
         * @return builder
         * 
         */
        public Builder httpsOnly(Boolean httpsOnly) {
            return httpsOnly(Output.of(httpsOnly));
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
         * @param noVerbose Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
         * 
         * @return builder
         * 
         */
        public Builder noVerbose(@Nullable Output<Boolean> noVerbose) {
            $.noVerbose = noVerbose;
            return this;
        }

        /**
         * @param noVerbose Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
         * 
         * @return builder
         * 
         */
        public Builder noVerbose(Boolean noVerbose) {
            return noVerbose(Output.of(noVerbose));
        }

        /**
         * @param outputDocument The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
         * 
         * @return builder
         * 
         */
        public Builder outputDocument(@Nullable Output<String> outputDocument) {
            $.outputDocument = outputDocument;
            return this;
        }

        /**
         * @param outputDocument The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
         * 
         * @return builder
         * 
         */
        public Builder outputDocument(String outputDocument) {
            return outputDocument(Output.of(outputDocument));
        }

        /**
         * @param quiet Turn off Wget&#39;s output.
         * 
         * @return builder
         * 
         */
        public Builder quiet(@Nullable Output<Boolean> quiet) {
            $.quiet = quiet;
            return this;
        }

        /**
         * @param quiet Turn off Wget&#39;s output.
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
         * @param timestamping Turn on time-stamping.
         * 
         * @return builder
         * 
         */
        public Builder timestamping(@Nullable Output<Boolean> timestamping) {
            $.timestamping = timestamping;
            return this;
        }

        /**
         * @param timestamping Turn on time-stamping.
         * 
         * @return builder
         * 
         */
        public Builder timestamping(Boolean timestamping) {
            return timestamping(Output.of(timestamping));
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
         * @param url Corresponds to the [URL...] argument.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<Either<String,List<String>>> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Corresponds to the [URL...] argument.
         * 
         * @return builder
         * 
         */
        public Builder url(Either<String,List<String>> url) {
            return url(Output.of(url));
        }

        /**
         * @param url Corresponds to the [URL...] argument.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Either.ofLeft(url));
        }

        /**
         * @param url Corresponds to the [URL...] argument.
         * 
         * @return builder
         * 
         */
        public Builder url(List<String> url) {
            return url(Either.ofRight(url));
        }

        public WgetArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("WgetArgs", "connection");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("WgetArgs", "url");
            }
            return $;
        }
    }

}
