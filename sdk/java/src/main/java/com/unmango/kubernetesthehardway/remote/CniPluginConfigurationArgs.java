// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.remote.inputs.CniBridgePluginConfigurationArgs;
import com.unmango.kubernetesthehardway.remote.inputs.CniLoopbackPluginConfigurationArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CniPluginConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final CniPluginConfigurationArgs Empty = new CniPluginConfigurationArgs();

    /**
     * The CNI bridge configuration.
     * 
     */
    @Import(name="bridge")
    private @Nullable Output<CniBridgePluginConfigurationArgs> bridge;

    /**
     * @return The CNI bridge configuration.
     * 
     */
    public Optional<Output<CniBridgePluginConfigurationArgs>> bridge() {
        return Optional.ofNullable(this.bridge);
    }

    /**
     * The parameters with which to connect to the remote host.
     * 
     */
    @Import(name="connection", required=true)
    private Output<ConnectionArgs> connection;

    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    public Output<ConnectionArgs> connection() {
        return this.connection;
    }

    /**
     * The CNI loopback configuration.
     * 
     */
    @Import(name="loopback")
    private @Nullable Output<CniLoopbackPluginConfigurationArgs> loopback;

    /**
     * @return The CNI loopback configuration.
     * 
     */
    public Optional<Output<CniLoopbackPluginConfigurationArgs>> loopback() {
        return Optional.ofNullable(this.loopback);
    }

    private CniPluginConfigurationArgs() {}

    private CniPluginConfigurationArgs(CniPluginConfigurationArgs $) {
        this.bridge = $.bridge;
        this.connection = $.connection;
        this.loopback = $.loopback;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CniPluginConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CniPluginConfigurationArgs $;

        public Builder() {
            $ = new CniPluginConfigurationArgs();
        }

        public Builder(CniPluginConfigurationArgs defaults) {
            $ = new CniPluginConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bridge The CNI bridge configuration.
         * 
         * @return builder
         * 
         */
        public Builder bridge(@Nullable Output<CniBridgePluginConfigurationArgs> bridge) {
            $.bridge = bridge;
            return this;
        }

        /**
         * @param bridge The CNI bridge configuration.
         * 
         * @return builder
         * 
         */
        public Builder bridge(CniBridgePluginConfigurationArgs bridge) {
            return bridge(Output.of(bridge));
        }

        /**
         * @param connection The parameters with which to connect to the remote host.
         * 
         * @return builder
         * 
         */
        public Builder connection(Output<ConnectionArgs> connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param connection The parameters with which to connect to the remote host.
         * 
         * @return builder
         * 
         */
        public Builder connection(ConnectionArgs connection) {
            return connection(Output.of(connection));
        }

        /**
         * @param loopback The CNI loopback configuration.
         * 
         * @return builder
         * 
         */
        public Builder loopback(@Nullable Output<CniLoopbackPluginConfigurationArgs> loopback) {
            $.loopback = loopback;
            return this;
        }

        /**
         * @param loopback The CNI loopback configuration.
         * 
         * @return builder
         * 
         */
        public Builder loopback(CniLoopbackPluginConfigurationArgs loopback) {
            return loopback(Output.of(loopback));
        }

        public CniPluginConfigurationArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("CniPluginConfigurationArgs", "connection");
            }
            return $;
        }
    }

}
