// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.remote.enums.Architecture;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeSchedulerInstallArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeSchedulerInstallArgs Empty = new KubeSchedulerInstallArgs();

    /**
     * The CPU architecture to install.
     * 
     */
    @Import(name="architecture")
    private @Nullable Output<Architecture> architecture;

    /**
     * @return The CPU architecture to install.
     * 
     */
    public Optional<Output<Architecture>> architecture() {
        return Optional.ofNullable(this.architecture);
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
     * The directory to install the binary to.
     * 
     */
    @Import(name="directory")
    private @Nullable Output<String> directory;

    /**
     * @return The directory to install the binary to.
     * 
     */
    public Optional<Output<String>> directory() {
        return Optional.ofNullable(this.directory);
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

    private KubeSchedulerInstallArgs() {}

    private KubeSchedulerInstallArgs(KubeSchedulerInstallArgs $) {
        this.architecture = $.architecture;
        this.connection = $.connection;
        this.directory = $.directory;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeSchedulerInstallArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeSchedulerInstallArgs $;

        public Builder() {
            $ = new KubeSchedulerInstallArgs();
        }

        public Builder(KubeSchedulerInstallArgs defaults) {
            $ = new KubeSchedulerInstallArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param architecture The CPU architecture to install.
         * 
         * @return builder
         * 
         */
        public Builder architecture(@Nullable Output<Architecture> architecture) {
            $.architecture = architecture;
            return this;
        }

        /**
         * @param architecture The CPU architecture to install.
         * 
         * @return builder
         * 
         */
        public Builder architecture(Architecture architecture) {
            return architecture(Output.of(architecture));
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
         * @param directory The directory to install the binary to.
         * 
         * @return builder
         * 
         */
        public Builder directory(@Nullable Output<String> directory) {
            $.directory = directory;
            return this;
        }

        /**
         * @param directory The directory to install the binary to.
         * 
         * @return builder
         * 
         */
        public Builder directory(String directory) {
            return directory(Output.of(directory));
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

        public KubeSchedulerInstallArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("KubeSchedulerInstallArgs", "connection");
            }
            $.directory = Codegen.stringProp("directory").output().arg($.directory).def("/usr/local/bin").getNullable();
            return $;
        }
    }

}
