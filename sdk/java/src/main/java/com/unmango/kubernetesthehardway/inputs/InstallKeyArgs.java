// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.Certificate;
import com.unmango.kubernetesthehardway.inputs.ConnectionArgs;
import com.unmango.kubernetesthehardway.inputs.ResourceOptionsArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstallKeyArgs extends com.pulumi.resources.InvokeArgs {

    public static final InstallKeyArgs Empty = new InstallKeyArgs();

    /**
     * The certificate to install.
     * 
     */
    @Import(name="cert", required=true)
    private Output<Certificate> cert;

    /**
     * @return The certificate to install.
     * 
     */
    public Output<Certificate> cert() {
        return this.cert;
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

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="options")
    private @Nullable ResourceOptionsArgs options;

    public Optional<ResourceOptionsArgs> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * The path to install to.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to install to.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    private InstallKeyArgs() {}

    private InstallKeyArgs(InstallKeyArgs $) {
        this.cert = $.cert;
        this.connection = $.connection;
        this.name = $.name;
        this.options = $.options;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstallKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstallKeyArgs $;

        public Builder() {
            $ = new InstallKeyArgs();
        }

        public Builder(InstallKeyArgs defaults) {
            $ = new InstallKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cert The certificate to install.
         * 
         * @return builder
         * 
         */
        public Builder cert(Output<Certificate> cert) {
            $.cert = cert;
            return this;
        }

        /**
         * @param cert The certificate to install.
         * 
         * @return builder
         * 
         */
        public Builder cert(Certificate cert) {
            return cert(Output.of(cert));
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

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder options(@Nullable ResourceOptionsArgs options) {
            $.options = options;
            return this;
        }

        /**
         * @param path The path to install to.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to install to.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        public InstallKeyArgs build() {
            if ($.cert == null) {
                throw new MissingRequiredPropertyException("InstallKeyArgs", "cert");
            }
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("InstallKeyArgs", "connection");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("InstallKeyArgs", "name");
            }
            return $;
        }
    }

}
