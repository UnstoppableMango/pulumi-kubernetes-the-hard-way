// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetKubeconfigArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKubeconfigArgs Empty = new GetKubeconfigArgs();

    /**
     * Certificate authority data.
     * 
     */
    @Import(name="caPem", required=true)
    private Output<String> caPem;

    /**
     * @return Certificate authority data.
     * 
     */
    public Output<String> caPem() {
        return this.caPem;
    }

    /**
     * The PEM encoded certificate data of the client.
     * 
     */
    @Import(name="clientCert", required=true)
    private Output<String> clientCert;

    /**
     * @return The PEM encoded certificate data of the client.
     * 
     */
    public Output<String> clientCert() {
        return this.clientCert;
    }

    /**
     * The PEM encoded private key data of the client.
     * 
     */
    @Import(name="clientKey", required=true)
    private Output<String> clientKey;

    /**
     * @return The PEM encoded private key data of the client.
     * 
     */
    public Output<String> clientKey() {
        return this.clientKey;
    }

    /**
     * A name to identify the cluster.
     * 
     */
    @Import(name="clusterName", required=true)
    private Output<String> clusterName;

    /**
     * @return A name to identify the cluster.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }

    /**
     * A name to use for the kubeconfig context
     * 
     */
    @Import(name="contextName")
    private @Nullable Output<String> contextName;

    /**
     * @return A name to use for the kubeconfig context
     * 
     */
    public Optional<Output<String>> contextName() {
        return Optional.ofNullable(this.contextName);
    }

    /**
     * The address and port of the Kubernetes API server.
     * 
     */
    @Import(name="server", required=true)
    private Output<String> server;

    /**
     * @return The address and port of the Kubernetes API server.
     * 
     */
    public Output<String> server() {
        return this.server;
    }

    /**
     * The username of the user
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the user
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private GetKubeconfigArgs() {}

    private GetKubeconfigArgs(GetKubeconfigArgs $) {
        this.caPem = $.caPem;
        this.clientCert = $.clientCert;
        this.clientKey = $.clientKey;
        this.clusterName = $.clusterName;
        this.contextName = $.contextName;
        this.server = $.server;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeconfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeconfigArgs $;

        public Builder() {
            $ = new GetKubeconfigArgs();
        }

        public Builder(GetKubeconfigArgs defaults) {
            $ = new GetKubeconfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param caPem Certificate authority data.
         * 
         * @return builder
         * 
         */
        public Builder caPem(Output<String> caPem) {
            $.caPem = caPem;
            return this;
        }

        /**
         * @param caPem Certificate authority data.
         * 
         * @return builder
         * 
         */
        public Builder caPem(String caPem) {
            return caPem(Output.of(caPem));
        }

        /**
         * @param clientCert The PEM encoded certificate data of the client.
         * 
         * @return builder
         * 
         */
        public Builder clientCert(Output<String> clientCert) {
            $.clientCert = clientCert;
            return this;
        }

        /**
         * @param clientCert The PEM encoded certificate data of the client.
         * 
         * @return builder
         * 
         */
        public Builder clientCert(String clientCert) {
            return clientCert(Output.of(clientCert));
        }

        /**
         * @param clientKey The PEM encoded private key data of the client.
         * 
         * @return builder
         * 
         */
        public Builder clientKey(Output<String> clientKey) {
            $.clientKey = clientKey;
            return this;
        }

        /**
         * @param clientKey The PEM encoded private key data of the client.
         * 
         * @return builder
         * 
         */
        public Builder clientKey(String clientKey) {
            return clientKey(Output.of(clientKey));
        }

        /**
         * @param clusterName A name to identify the cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName A name to identify the cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        /**
         * @param contextName A name to use for the kubeconfig context
         * 
         * @return builder
         * 
         */
        public Builder contextName(@Nullable Output<String> contextName) {
            $.contextName = contextName;
            return this;
        }

        /**
         * @param contextName A name to use for the kubeconfig context
         * 
         * @return builder
         * 
         */
        public Builder contextName(String contextName) {
            return contextName(Output.of(contextName));
        }

        /**
         * @param server The address and port of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder server(Output<String> server) {
            $.server = server;
            return this;
        }

        /**
         * @param server The address and port of the Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder server(String server) {
            return server(Output.of(server));
        }

        /**
         * @param username The username of the user
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the user
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public GetKubeconfigArgs build() {
            if ($.caPem == null) {
                throw new MissingRequiredPropertyException("GetKubeconfigArgs", "caPem");
            }
            if ($.clientCert == null) {
                throw new MissingRequiredPropertyException("GetKubeconfigArgs", "clientCert");
            }
            if ($.clientKey == null) {
                throw new MissingRequiredPropertyException("GetKubeconfigArgs", "clientKey");
            }
            if ($.clusterName == null) {
                throw new MissingRequiredPropertyException("GetKubeconfigArgs", "clusterName");
            }
            if ($.server == null) {
                throw new MissingRequiredPropertyException("GetKubeconfigArgs", "server");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("GetKubeconfigArgs", "username");
            }
            return $;
        }
    }

}
