// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.config.outputs.KubeProxyConfiguration;
import java.util.Objects;

@CustomType
public final class GetKubeProxyConfigurationResult {
    private KubeProxyConfiguration result;

    private GetKubeProxyConfigurationResult() {}
    public KubeProxyConfiguration result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeProxyConfigurationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private KubeProxyConfiguration result;
        public Builder() {}
        public Builder(GetKubeProxyConfigurationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(KubeProxyConfiguration result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("GetKubeProxyConfigurationResult", "result");
            }
            this.result = result;
            return this;
        }
        public GetKubeProxyConfigurationResult build() {
            final var _resultValue = new GetKubeProxyConfigurationResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}