// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.enums.Architecture;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EtcdArgs extends com.pulumi.resources.ResourceArgs {

    public static final EtcdArgs Empty = new EtcdArgs();

    /**
     * The etcd CPU architecture.
     * 
     */
    @Import(name="architecture")
    private @Nullable Output<Architecture> architecture;

    /**
     * @return The etcd CPU architecture.
     * 
     */
    public Optional<Output<Architecture>> architecture() {
        return Optional.ofNullable(this.architecture);
    }

    /**
     * The connection details.
     * 
     */
    @Import(name="connection", required=true)
    private Output<ConnectionArgs> connection;

    /**
     * @return The connection details.
     * 
     */
    public Output<ConnectionArgs> connection() {
        return this.connection;
    }

    /**
     * The version of etcd to install.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The version of etcd to install.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private EtcdArgs() {}

    private EtcdArgs(EtcdArgs $) {
        this.architecture = $.architecture;
        this.connection = $.connection;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EtcdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EtcdArgs $;

        public Builder() {
            $ = new EtcdArgs();
        }

        public Builder(EtcdArgs defaults) {
            $ = new EtcdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param architecture The etcd CPU architecture.
         * 
         * @return builder
         * 
         */
        public Builder architecture(@Nullable Output<Architecture> architecture) {
            $.architecture = architecture;
            return this;
        }

        /**
         * @param architecture The etcd CPU architecture.
         * 
         * @return builder
         * 
         */
        public Builder architecture(Architecture architecture) {
            return architecture(Output.of(architecture));
        }

        /**
         * @param connection The connection details.
         * 
         * @return builder
         * 
         */
        public Builder connection(Output<ConnectionArgs> connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param connection The connection details.
         * 
         * @return builder
         * 
         */
        public Builder connection(ConnectionArgs connection) {
            return connection(Output.of(connection));
        }

        /**
         * @param version The version of etcd to install.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of etcd to install.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public EtcdArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("EtcdArgs", "connection");
            }
            return $;
        }
    }

}