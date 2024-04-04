// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.unmango.kubernetesthehardway.remote.enums.Architecture;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeApiServerInstallArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeApiServerInstallArgs Empty = new KubeApiServerInstallArgs();

    /**
     * The kube-apiserver CPU architecture.
     * 
     */
    @Import(name="architecture")
    private @Nullable Output<Architecture> architecture;

    /**
     * @return The kube-apiserver CPU architecture.
     * 
     */
    public Optional<Output<Architecture>> architecture() {
        return Optional.ofNullable(this.architecture);
    }

    /**
     * The connection details.
     * 
     */
    @Import(name="connection")
    private @Nullable Output<ConnectionArgs> connection;

    /**
     * @return The connection details.
     * 
     */
    public Optional<Output<ConnectionArgs>> connection() {
        return Optional.ofNullable(this.connection);
    }

    /**
     * Directory to install the `etcd` and `etcdctl` binaries.
     * 
     */
    @Import(name="installDirectory")
    private @Nullable Output<String> installDirectory;

    /**
     * @return Directory to install the `etcd` and `etcdctl` binaries.
     * 
     */
    public Optional<Output<String>> installDirectory() {
        return Optional.ofNullable(this.installDirectory);
    }

    /**
     * The version of kube-apiserver to install.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The version of kube-apiserver to install.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private KubeApiServerInstallArgs() {}

    private KubeApiServerInstallArgs(KubeApiServerInstallArgs $) {
        this.architecture = $.architecture;
        this.connection = $.connection;
        this.installDirectory = $.installDirectory;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeApiServerInstallArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeApiServerInstallArgs $;

        public Builder() {
            $ = new KubeApiServerInstallArgs();
        }

        public Builder(KubeApiServerInstallArgs defaults) {
            $ = new KubeApiServerInstallArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param architecture The kube-apiserver CPU architecture.
         * 
         * @return builder
         * 
         */
        public Builder architecture(@Nullable Output<Architecture> architecture) {
            $.architecture = architecture;
            return this;
        }

        /**
         * @param architecture The kube-apiserver CPU architecture.
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
        public Builder connection(@Nullable Output<ConnectionArgs> connection) {
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
         * @param installDirectory Directory to install the `etcd` and `etcdctl` binaries.
         * 
         * @return builder
         * 
         */
        public Builder installDirectory(@Nullable Output<String> installDirectory) {
            $.installDirectory = installDirectory;
            return this;
        }

        /**
         * @param installDirectory Directory to install the `etcd` and `etcdctl` binaries.
         * 
         * @return builder
         * 
         */
        public Builder installDirectory(String installDirectory) {
            return installDirectory(Output.of(installDirectory));
        }

        /**
         * @param version The version of kube-apiserver to install.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of kube-apiserver to install.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public KubeApiServerInstallArgs build() {
            $.installDirectory = Codegen.stringProp("installDirectory").output().arg($.installDirectory).def("/usr/local/bin").getNullable();
            return $;
        }
    }

}