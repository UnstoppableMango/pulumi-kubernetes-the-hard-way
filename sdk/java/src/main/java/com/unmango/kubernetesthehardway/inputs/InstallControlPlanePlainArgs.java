// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.ClusterPki;
import com.unmango.kubernetesthehardway.inputs.Connection;
import com.unmango.kubernetesthehardway.inputs.ResourceOptions;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstallControlPlanePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final InstallControlPlanePlainArgs Empty = new InstallControlPlanePlainArgs();

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

    @Import(name="options")
    private @Nullable ResourceOptions options;

    public Optional<ResourceOptions> options() {
        return Optional.ofNullable(this.options);
    }

    /**
     * The PKI to install
     * 
     */
    @Import(name="pki", required=true)
    private ClusterPki pki;

    /**
     * @return The PKI to install
     * 
     */
    public ClusterPki pki() {
        return this.pki;
    }

    private InstallControlPlanePlainArgs() {}

    private InstallControlPlanePlainArgs(InstallControlPlanePlainArgs $) {
        this.connection = $.connection;
        this.options = $.options;
        this.pki = $.pki;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstallControlPlanePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstallControlPlanePlainArgs $;

        public Builder() {
            $ = new InstallControlPlanePlainArgs();
        }

        public Builder(InstallControlPlanePlainArgs defaults) {
            $ = new InstallControlPlanePlainArgs(Objects.requireNonNull(defaults));
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

        public Builder options(@Nullable ResourceOptions options) {
            $.options = options;
            return this;
        }

        /**
         * @param pki The PKI to install
         * 
         * @return builder
         * 
         */
        public Builder pki(ClusterPki pki) {
            $.pki = pki;
            return this;
        }

        public InstallControlPlanePlainArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("InstallControlPlanePlainArgs", "connection");
            }
            if ($.pki == null) {
                throw new MissingRequiredPropertyException("InstallControlPlanePlainArgs", "pki");
            }
            return $;
        }
    }

}