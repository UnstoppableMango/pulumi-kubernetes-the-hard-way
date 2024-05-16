// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.remote.enums.SystemdServiceRestart;
import com.unmango.kubernetesthehardway.remote.inputs.KubeApiServerProps;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeApiServerServiceArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeApiServerServiceArgs Empty = new KubeApiServerServiceArgs();

    /**
     * If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
     * 
     */
    @Import(name="clientCaFile")
    private @Nullable Output<String> clientCaFile;

    /**
     * @return If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
     * 
     */
    public Optional<Output<String>> clientCaFile() {
        return Optional.ofNullable(this.clientCaFile);
    }

    /**
     * KubeApiServer configuration.
     * 
     */
    @Import(name="configuration", required=true)
    private Output<KubeApiServerProps> configuration;

    /**
     * @return KubeApiServer configuration.
     * 
     */
    public Output<KubeApiServerProps> configuration() {
        return this.configuration;
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
     * Optional systemd unit description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Optional systemd unit description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The location to create the service file.
     * 
     */
    @Import(name="directory")
    private @Nullable Output<String> directory;

    /**
     * @return The location to create the service file.
     * 
     */
    public Optional<Output<String>> directory() {
        return Optional.ofNullable(this.directory);
    }

    /**
     * Optional systemd unit documentation
     * 
     */
    @Import(name="documentation")
    private @Nullable Output<String> documentation;

    /**
     * @return Optional systemd unit documentation
     * 
     */
    public Optional<Output<String>> documentation() {
        return Optional.ofNullable(this.documentation);
    }

    /**
     * List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
     * 
     */
    @Import(name="etcdServers")
    private @Nullable Output<String> etcdServers;

    /**
     * @return List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
     * 
     */
    public Optional<Output<String>> etcdServers() {
        return Optional.ofNullable(this.etcdServers);
    }

    /**
     * Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
     * 
     */
    @Import(name="restart")
    private @Nullable Output<SystemdServiceRestart> restart;

    /**
     * @return Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
     * 
     */
    public Optional<Output<SystemdServiceRestart>> restart() {
        return Optional.ofNullable(this.restart);
    }

    /**
     * Optionally override the systemd service RestartSec. Defaults to `5`.
     * 
     */
    @Import(name="restartSec")
    private @Nullable Output<String> restartSec;

    /**
     * @return Optionally override the systemd service RestartSec. Defaults to `5`.
     * 
     */
    public Optional<Output<String>> restartSec() {
        return Optional.ofNullable(this.restartSec);
    }

    /**
     * Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
     * 
     */
    @Import(name="wantedBy")
    private @Nullable Output<String> wantedBy;

    /**
     * @return Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
     * 
     */
    public Optional<Output<String>> wantedBy() {
        return Optional.ofNullable(this.wantedBy);
    }

    private KubeApiServerServiceArgs() {}

    private KubeApiServerServiceArgs(KubeApiServerServiceArgs $) {
        this.clientCaFile = $.clientCaFile;
        this.configuration = $.configuration;
        this.connection = $.connection;
        this.description = $.description;
        this.directory = $.directory;
        this.documentation = $.documentation;
        this.etcdServers = $.etcdServers;
        this.restart = $.restart;
        this.restartSec = $.restartSec;
        this.wantedBy = $.wantedBy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeApiServerServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeApiServerServiceArgs $;

        public Builder() {
            $ = new KubeApiServerServiceArgs();
        }

        public Builder(KubeApiServerServiceArgs defaults) {
            $ = new KubeApiServerServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientCaFile If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCaFile(@Nullable Output<String> clientCaFile) {
            $.clientCaFile = clientCaFile;
            return this;
        }

        /**
         * @param clientCaFile If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
         * 
         * @return builder
         * 
         */
        public Builder clientCaFile(String clientCaFile) {
            return clientCaFile(Output.of(clientCaFile));
        }

        /**
         * @param configuration KubeApiServer configuration.
         * 
         * @return builder
         * 
         */
        public Builder configuration(Output<KubeApiServerProps> configuration) {
            $.configuration = configuration;
            return this;
        }

        /**
         * @param configuration KubeApiServer configuration.
         * 
         * @return builder
         * 
         */
        public Builder configuration(KubeApiServerProps configuration) {
            return configuration(Output.of(configuration));
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
         * @param description Optional systemd unit description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Optional systemd unit description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param directory The location to create the service file.
         * 
         * @return builder
         * 
         */
        public Builder directory(@Nullable Output<String> directory) {
            $.directory = directory;
            return this;
        }

        /**
         * @param directory The location to create the service file.
         * 
         * @return builder
         * 
         */
        public Builder directory(String directory) {
            return directory(Output.of(directory));
        }

        /**
         * @param documentation Optional systemd unit documentation
         * 
         * @return builder
         * 
         */
        public Builder documentation(@Nullable Output<String> documentation) {
            $.documentation = documentation;
            return this;
        }

        /**
         * @param documentation Optional systemd unit documentation
         * 
         * @return builder
         * 
         */
        public Builder documentation(String documentation) {
            return documentation(Output.of(documentation));
        }

        /**
         * @param etcdServers List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
         * 
         * @return builder
         * 
         */
        public Builder etcdServers(@Nullable Output<String> etcdServers) {
            $.etcdServers = etcdServers;
            return this;
        }

        /**
         * @param etcdServers List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
         * 
         * @return builder
         * 
         */
        public Builder etcdServers(String etcdServers) {
            return etcdServers(Output.of(etcdServers));
        }

        /**
         * @param restart Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
         * 
         * @return builder
         * 
         */
        public Builder restart(@Nullable Output<SystemdServiceRestart> restart) {
            $.restart = restart;
            return this;
        }

        /**
         * @param restart Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
         * 
         * @return builder
         * 
         */
        public Builder restart(SystemdServiceRestart restart) {
            return restart(Output.of(restart));
        }

        /**
         * @param restartSec Optionally override the systemd service RestartSec. Defaults to `5`.
         * 
         * @return builder
         * 
         */
        public Builder restartSec(@Nullable Output<String> restartSec) {
            $.restartSec = restartSec;
            return this;
        }

        /**
         * @param restartSec Optionally override the systemd service RestartSec. Defaults to `5`.
         * 
         * @return builder
         * 
         */
        public Builder restartSec(String restartSec) {
            return restartSec(Output.of(restartSec));
        }

        /**
         * @param wantedBy Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
         * 
         * @return builder
         * 
         */
        public Builder wantedBy(@Nullable Output<String> wantedBy) {
            $.wantedBy = wantedBy;
            return this;
        }

        /**
         * @param wantedBy Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
         * 
         * @return builder
         * 
         */
        public Builder wantedBy(String wantedBy) {
            return wantedBy(Output.of(wantedBy));
        }

        public KubeApiServerServiceArgs build() {
            if ($.configuration == null) {
                throw new MissingRequiredPropertyException("KubeApiServerServiceArgs", "configuration");
            }
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("KubeApiServerServiceArgs", "connection");
            }
            return $;
        }
    }

}
