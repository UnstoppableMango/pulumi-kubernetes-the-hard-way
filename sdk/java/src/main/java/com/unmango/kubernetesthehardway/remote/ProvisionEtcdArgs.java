// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.remote.enums.Architecture;
import com.unmango.kubernetesthehardway.tls.inputs.BundleArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProvisionEtcdArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProvisionEtcdArgs Empty = new ProvisionEtcdArgs();

    /**
     * TODO
     * 
     */
    @Import(name="architecture")
    private @Nullable Output<Architecture> architecture;

    /**
     * @return TODO
     * 
     */
    public Optional<Output<Architecture>> architecture() {
        return Optional.ofNullable(this.architecture);
    }

    /**
     * TODO
     * 
     */
    @Import(name="binaryDirectory")
    private @Nullable Output<String> binaryDirectory;

    /**
     * @return TODO
     * 
     */
    public Optional<Output<String>> binaryDirectory() {
        return Optional.ofNullable(this.binaryDirectory);
    }

    /**
     * The TLS bundle.
     * 
     */
    @Import(name="bundle", required=true)
    private Output<BundleArgs> bundle;

    /**
     * @return The TLS bundle.
     * 
     */
    public Output<BundleArgs> bundle() {
        return this.bundle;
    }

    /**
     * The directory to use for etcd configuration.
     * 
     */
    @Import(name="configurationDirectory")
    private @Nullable Output<String> configurationDirectory;

    /**
     * @return The directory to use for etcd configuration.
     * 
     */
    public Optional<Output<String>> configurationDirectory() {
        return Optional.ofNullable(this.configurationDirectory);
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
     * The directory to use for etcd data.
     * 
     */
    @Import(name="dataDirectory")
    private @Nullable Output<String> dataDirectory;

    /**
     * @return The directory to use for etcd data.
     * 
     */
    public Optional<Output<String>> dataDirectory() {
        return Optional.ofNullable(this.dataDirectory);
    }

    /**
     * The internal IP of the etcd node
     * 
     */
    @Import(name="internalIp", required=true)
    private Output<String> internalIp;

    /**
     * @return The internal IP of the etcd node
     * 
     */
    public Output<String> internalIp() {
        return this.internalIp;
    }

    /**
     * The version to install.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The version to install.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private ProvisionEtcdArgs() {}

    private ProvisionEtcdArgs(ProvisionEtcdArgs $) {
        this.architecture = $.architecture;
        this.binaryDirectory = $.binaryDirectory;
        this.bundle = $.bundle;
        this.configurationDirectory = $.configurationDirectory;
        this.connection = $.connection;
        this.dataDirectory = $.dataDirectory;
        this.internalIp = $.internalIp;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProvisionEtcdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProvisionEtcdArgs $;

        public Builder() {
            $ = new ProvisionEtcdArgs();
        }

        public Builder(ProvisionEtcdArgs defaults) {
            $ = new ProvisionEtcdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param architecture TODO
         * 
         * @return builder
         * 
         */
        public Builder architecture(@Nullable Output<Architecture> architecture) {
            $.architecture = architecture;
            return this;
        }

        /**
         * @param architecture TODO
         * 
         * @return builder
         * 
         */
        public Builder architecture(Architecture architecture) {
            return architecture(Output.of(architecture));
        }

        /**
         * @param binaryDirectory TODO
         * 
         * @return builder
         * 
         */
        public Builder binaryDirectory(@Nullable Output<String> binaryDirectory) {
            $.binaryDirectory = binaryDirectory;
            return this;
        }

        /**
         * @param binaryDirectory TODO
         * 
         * @return builder
         * 
         */
        public Builder binaryDirectory(String binaryDirectory) {
            return binaryDirectory(Output.of(binaryDirectory));
        }

        /**
         * @param bundle The TLS bundle.
         * 
         * @return builder
         * 
         */
        public Builder bundle(Output<BundleArgs> bundle) {
            $.bundle = bundle;
            return this;
        }

        /**
         * @param bundle The TLS bundle.
         * 
         * @return builder
         * 
         */
        public Builder bundle(BundleArgs bundle) {
            return bundle(Output.of(bundle));
        }

        /**
         * @param configurationDirectory The directory to use for etcd configuration.
         * 
         * @return builder
         * 
         */
        public Builder configurationDirectory(@Nullable Output<String> configurationDirectory) {
            $.configurationDirectory = configurationDirectory;
            return this;
        }

        /**
         * @param configurationDirectory The directory to use for etcd configuration.
         * 
         * @return builder
         * 
         */
        public Builder configurationDirectory(String configurationDirectory) {
            return configurationDirectory(Output.of(configurationDirectory));
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
         * @param dataDirectory The directory to use for etcd data.
         * 
         * @return builder
         * 
         */
        public Builder dataDirectory(@Nullable Output<String> dataDirectory) {
            $.dataDirectory = dataDirectory;
            return this;
        }

        /**
         * @param dataDirectory The directory to use for etcd data.
         * 
         * @return builder
         * 
         */
        public Builder dataDirectory(String dataDirectory) {
            return dataDirectory(Output.of(dataDirectory));
        }

        /**
         * @param internalIp The internal IP of the etcd node
         * 
         * @return builder
         * 
         */
        public Builder internalIp(Output<String> internalIp) {
            $.internalIp = internalIp;
            return this;
        }

        /**
         * @param internalIp The internal IP of the etcd node
         * 
         * @return builder
         * 
         */
        public Builder internalIp(String internalIp) {
            return internalIp(Output.of(internalIp));
        }

        /**
         * @param version The version to install.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version to install.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public ProvisionEtcdArgs build() {
            if ($.bundle == null) {
                throw new MissingRequiredPropertyException("ProvisionEtcdArgs", "bundle");
            }
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("ProvisionEtcdArgs", "connection");
            }
            if ($.internalIp == null) {
                throw new MissingRequiredPropertyException("ProvisionEtcdArgs", "internalIp");
            }
            return $;
        }
    }

}
