// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.enums.EtcdctlCommand;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EtcdctlOpts {
    /**
     * @return TODO
     * 
     */
    private @Nullable String caCert;
    /**
     * @return TODO
     * 
     */
    private @Nullable String cert;
    /**
     * @return TODO
     * 
     */
    private EtcdctlCommand commands;
    /**
     * @return TODO
     * 
     */
    private @Nullable String endpoints;
    /**
     * @return TODO
     * 
     */
    private @Nullable String key;

    private EtcdctlOpts() {}
    /**
     * @return TODO
     * 
     */
    public Optional<String> caCert() {
        return Optional.ofNullable(this.caCert);
    }
    /**
     * @return TODO
     * 
     */
    public Optional<String> cert() {
        return Optional.ofNullable(this.cert);
    }
    /**
     * @return TODO
     * 
     */
    public EtcdctlCommand commands() {
        return this.commands;
    }
    /**
     * @return TODO
     * 
     */
    public Optional<String> endpoints() {
        return Optional.ofNullable(this.endpoints);
    }
    /**
     * @return TODO
     * 
     */
    public Optional<String> key() {
        return Optional.ofNullable(this.key);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EtcdctlOpts defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String caCert;
        private @Nullable String cert;
        private EtcdctlCommand commands;
        private @Nullable String endpoints;
        private @Nullable String key;
        public Builder() {}
        public Builder(EtcdctlOpts defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.caCert = defaults.caCert;
    	      this.cert = defaults.cert;
    	      this.commands = defaults.commands;
    	      this.endpoints = defaults.endpoints;
    	      this.key = defaults.key;
        }

        @CustomType.Setter
        public Builder caCert(@Nullable String caCert) {

            this.caCert = caCert;
            return this;
        }
        @CustomType.Setter
        public Builder cert(@Nullable String cert) {

            this.cert = cert;
            return this;
        }
        @CustomType.Setter
        public Builder commands(EtcdctlCommand commands) {
            if (commands == null) {
              throw new MissingRequiredPropertyException("EtcdctlOpts", "commands");
            }
            this.commands = commands;
            return this;
        }
        @CustomType.Setter
        public Builder endpoints(@Nullable String endpoints) {

            this.endpoints = endpoints;
            return this;
        }
        @CustomType.Setter
        public Builder key(@Nullable String key) {

            this.key = key;
            return this;
        }
        public EtcdctlOpts build() {
            final var _resultValue = new EtcdctlOpts();
            _resultValue.caCert = caCert;
            _resultValue.cert = cert;
            _resultValue.commands = commands;
            _resultValue.endpoints = endpoints;
            _resultValue.key = key;
            return _resultValue;
        }
    }
}
