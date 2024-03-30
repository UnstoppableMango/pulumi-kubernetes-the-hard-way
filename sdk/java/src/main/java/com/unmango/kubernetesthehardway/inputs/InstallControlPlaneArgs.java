// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.inputs.ConnectionArgs;
import com.unmango.kubernetesthehardway.inputs.ResourceOptionsArgs;
import com.unmango.kubernetesthehardway.tls.ClusterPki;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstallControlPlaneArgs extends com.pulumi.resources.InvokeArgs {

    public static final InstallControlPlaneArgs Empty = new InstallControlPlaneArgs();

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

    @Import(name="options")
    private @Nullable ResourceOptionsArgs options;

    public Optional<ResourceOptionsArgs> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * The PKI to install
     * 
     */
    @Import(name="pki", required=true)
    private Output<ClusterPki> pki;

    /**
     * @return The PKI to install
     * 
     */
    public Output<ClusterPki> pki() {
        return this.pki;
    }

    private InstallControlPlaneArgs() {}

    private InstallControlPlaneArgs(InstallControlPlaneArgs $) {
        this.connection = $.connection;
        this.options = $.options;
        this.pki = $.pki;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstallControlPlaneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstallControlPlaneArgs $;

        public Builder() {
            $ = new InstallControlPlaneArgs();
        }

        public Builder(InstallControlPlaneArgs defaults) {
            $ = new InstallControlPlaneArgs(Objects.requireNonNull(defaults));
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

        public Builder options(@Nullable ResourceOptionsArgs options) {
            $.options = options;
            return this;
        }

        /**
         * @param pki The PKI to install
         * 
         * @return builder
         * 
         */
        public Builder pki(Output<ClusterPki> pki) {
            $.pki = pki;
            return this;
        }

        /**
         * @param pki The PKI to install
         * 
         * @return builder
         * 
         */
        public Builder pki(ClusterPki pki) {
            return pki(Output.of(pki));
        }

        public InstallControlPlaneArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("InstallControlPlaneArgs", "connection");
            }
            if ($.pki == null) {
                throw new MissingRequiredPropertyException("InstallControlPlaneArgs", "pki");
            }
            return $;
        }
    }

}
