// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.enums.CommandLifecycle;
import com.unmango.kubernetesthehardway.tools.enums.SystemctlCommand;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SystemctlArgs extends com.pulumi.resources.ResourceArgs {

    public static final SystemctlArgs Empty = new SystemctlArgs();

    @Import(name="commands", required=true)
    private Output<List<SystemctlCommand>> commands;

    public Output<List<SystemctlCommand>> commands() {
        return this.commands;
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

    @Import(name="environment")
    private @Nullable Output<Map<String,String>> environment;

    public Optional<Output<Map<String,String>>> environment() {
        return Optional.ofNullable(this.environment);
    }

    @Import(name="lifecycle")
    private @Nullable CommandLifecycle lifecycle;

    public Optional<CommandLifecycle> lifecycle() {
        return Optional.ofNullable(this.lifecycle);
    }

    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private SystemctlArgs() {}

    private SystemctlArgs(SystemctlArgs $) {
        this.commands = $.commands;
        this.connection = $.connection;
        this.environment = $.environment;
        this.lifecycle = $.lifecycle;
        this.serviceName = $.serviceName;
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

        public Builder commands(Output<List<SystemctlCommand>> commands) {
            $.commands = commands;
            return this;
        }

        public Builder commands(List<SystemctlCommand> commands) {
            return commands(Output.of(commands));
        }

        public Builder commands(SystemctlCommand... commands) {
            return commands(List.of(commands));
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

        public Builder environment(@Nullable Output<Map<String,String>> environment) {
            $.environment = environment;
            return this;
        }

        public Builder environment(Map<String,String> environment) {
            return environment(Output.of(environment));
        }

        public Builder lifecycle(@Nullable CommandLifecycle lifecycle) {
            $.lifecycle = lifecycle;
            return this;
        }

        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public SystemctlArgs build() {
            if ($.commands == null) {
                throw new MissingRequiredPropertyException("SystemctlArgs", "commands");
            }
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("SystemctlArgs", "connection");
            }
            return $;
        }
    }

}
