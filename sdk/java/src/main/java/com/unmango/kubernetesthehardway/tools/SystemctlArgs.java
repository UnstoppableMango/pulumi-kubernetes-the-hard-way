// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SystemctlArgs extends com.pulumi.resources.ResourceArgs {

    public static final SystemctlArgs Empty = new SystemctlArgs();

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

    @Import(name="daemonReload")
    private @Nullable Output<Boolean> daemonReload;

    public Optional<Output<Boolean>> daemonReload() {
        return Optional.ofNullable(this.daemonReload);
    }

    @Import(name="enable")
    private @Nullable Output<String> enable;

    public Optional<Output<String>> enable() {
        return Optional.ofNullable(this.enable);
    }

    @Import(name="start")
    private @Nullable Output<String> start;

    public Optional<Output<String>> start() {
        return Optional.ofNullable(this.start);
    }

    private SystemctlArgs() {}

    private SystemctlArgs(SystemctlArgs $) {
        this.connection = $.connection;
        this.daemonReload = $.daemonReload;
        this.enable = $.enable;
        this.start = $.start;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SystemctlArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SystemctlArgs $;

        public Builder() {
            $ = new SystemctlArgs();
        }

        public Builder(SystemctlArgs defaults) {
            $ = new SystemctlArgs(Objects.requireNonNull(defaults));
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

        public Builder daemonReload(@Nullable Output<Boolean> daemonReload) {
            $.daemonReload = daemonReload;
            return this;
        }

        public Builder daemonReload(Boolean daemonReload) {
            return daemonReload(Output.of(daemonReload));
        }

        public Builder enable(@Nullable Output<String> enable) {
            $.enable = enable;
            return this;
        }

        public Builder enable(String enable) {
            return enable(Output.of(enable));
        }

        public Builder start(@Nullable Output<String> start) {
            $.start = start;
            return this;
        }

        public Builder start(String start) {
            return start(Output.of(start));
        }

        public SystemctlArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("SystemctlArgs", "connection");
            }
            return $;
        }
    }

}
