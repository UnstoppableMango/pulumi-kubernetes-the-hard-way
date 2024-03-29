// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.inputs.Connection;
import com.unmango.kubernetesthehardway.inputs.KeyPair;
import com.unmango.kubernetesthehardway.inputs.ResourceOptions;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstallCertPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final InstallCertPlainArgs Empty = new InstallCertPlainArgs();

    /**
     * The connection details.
     * 
     */
    @Import(name="connection", required=true)
    private Connection connection;

    /**
     * @return The connection details.
     * 
     */
    public Connection connection() {
        return this.connection;
    }

    /**
     * The certificate to install at the remote location.
     * 
     */
    @Import(name="keypair", required=true)
    private KeyPair keypair;

    /**
     * @return The certificate to install at the remote location.
     * 
     */
    public KeyPair keypair() {
        return this.keypair;
    }

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="options")
    private @Nullable ResourceOptions options;

    public Optional<ResourceOptions> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * The path to install to.
     * 
     */
    @Import(name="path")
    private @Nullable String path;

    /**
     * @return The path to install to.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }

    private InstallCertPlainArgs() {}

    private InstallCertPlainArgs(InstallCertPlainArgs $) {
        this.connection = $.connection;
        this.keypair = $.keypair;
        this.name = $.name;
        this.options = $.options;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstallCertPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstallCertPlainArgs $;

        public Builder() {
            $ = new InstallCertPlainArgs();
        }

        public Builder(InstallCertPlainArgs defaults) {
            $ = new InstallCertPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connection The connection details.
         * 
         * @return builder
         * 
         */
        public Builder connection(Connection connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param keypair The certificate to install at the remote location.
         * 
         * @return builder
         * 
         */
        public Builder keypair(KeyPair keypair) {
            $.keypair = keypair;
            return this;
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder options(@Nullable ResourceOptions options) {
            $.options = options;
            return this;
        }

        /**
         * @param path The path to install to.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable String path) {
            $.path = path;
            return this;
        }

        public InstallCertPlainArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("InstallCertPlainArgs", "connection");
            }
            if ($.keypair == null) {
                throw new MissingRequiredPropertyException("InstallCertPlainArgs", "keypair");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("InstallCertPlainArgs", "name");
            }
            return $;
        }
    }

}
