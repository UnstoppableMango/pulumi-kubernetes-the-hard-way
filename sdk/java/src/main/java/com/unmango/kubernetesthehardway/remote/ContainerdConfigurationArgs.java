// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.remote.inputs.ContainerdCriPluginConfigurationArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerdConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerdConfigurationArgs Empty = new ContainerdConfigurationArgs();

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
     * The cri configuration.
     * 
     */
    @Import(name="cri")
    private @Nullable ContainerdCriPluginConfigurationArgs cri;

    /**
     * @return The cri configuration.
     * 
     */
    public Optional<ContainerdCriPluginConfigurationArgs> cri() {
        return Optional.ofNullable(this.cri);
    }

    private ContainerdConfigurationArgs() {}

    private ContainerdConfigurationArgs(ContainerdConfigurationArgs $) {
        this.connection = $.connection;
        this.cri = $.cri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerdConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerdConfigurationArgs $;

        public Builder() {
            $ = new ContainerdConfigurationArgs();
        }

        public Builder(ContainerdConfigurationArgs defaults) {
            $ = new ContainerdConfigurationArgs(Objects.requireNonNull(defaults));
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
         * @param cri The cri configuration.
         * 
         * @return builder
         * 
         */
        public Builder cri(@Nullable ContainerdCriPluginConfigurationArgs cri) {
            $.cri = cri;
            return this;
        }

        public ContainerdConfigurationArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("ContainerdConfigurationArgs", "connection");
            }
            return $;
        }
    }

}
