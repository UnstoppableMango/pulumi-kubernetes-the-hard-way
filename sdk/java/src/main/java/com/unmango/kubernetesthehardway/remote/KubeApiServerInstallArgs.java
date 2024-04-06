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
     * Directory to install the `etcd` and `etcdctl` binaries.
     * 
     */
    @Import(name="directory")
    private @Nullable Output<String> directory;

    /**
     * @return Directory to install the `etcd` and `etcdctl` binaries.
     * 
     */
    public Optional<Output<String>> directory() {
        return Optional.ofNullable(this.directory);
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
        this.directory = $.directory;
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
         * @param directory Directory to install the `etcd` and `etcdctl` binaries.
         * 
         * @return builder
         * 
         */
        public Builder directory(@Nullable Output<String> directory) {
            $.directory = directory;
            return this;
        }

        /**
         * @param directory Directory to install the `etcd` and `etcdctl` binaries.
         * 
         * @return builder
         * 
         */
        public Builder directory(String directory) {
            return directory(Output.of(directory));
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
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("KubeApiServerInstallArgs", "connection");
            }
            $.directory = Codegen.stringProp("directory").output().arg($.directory).def("/usr/local/bin").getNullable();
            return $;
        }
    }

}
