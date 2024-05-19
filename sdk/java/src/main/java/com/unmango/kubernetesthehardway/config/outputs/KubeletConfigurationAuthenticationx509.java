// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class KubeletConfigurationAuthenticationx509 {
    /**
     * @return TODO
     * 
     */
    private String clientCAFile;

    private KubeletConfigurationAuthenticationx509() {}
    /**
     * @return TODO
     * 
     */
    public String clientCAFile() {
        return this.clientCAFile;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubeletConfigurationAuthenticationx509 defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clientCAFile;
        public Builder() {}
        public Builder(KubeletConfigurationAuthenticationx509 defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientCAFile = defaults.clientCAFile;
        }

        @CustomType.Setter
        public Builder clientCAFile(String clientCAFile) {
            if (clientCAFile == null) {
              throw new MissingRequiredPropertyException("KubeletConfigurationAuthenticationx509", "clientCAFile");
            }
            this.clientCAFile = clientCAFile;
            return this;
        }
        public KubeletConfigurationAuthenticationx509 build() {
            final var _resultValue = new KubeletConfigurationAuthenticationx509();
            _resultValue.clientCAFile = clientCAFile;
            return _resultValue;
        }
    }
}