// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.config.outputs.KubeletConfiguration;
import java.util.Objects;

@CustomType
public final class GetKubeletConfigurationResult {
    private KubeletConfiguration result;

    private GetKubeletConfigurationResult() {}
    public KubeletConfiguration result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeletConfigurationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private KubeletConfiguration result;
        public Builder() {}
        public Builder(GetKubeletConfigurationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(KubeletConfiguration result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("GetKubeletConfigurationResult", "result");
            }
            this.result = result;
            return this;
        }
        public GetKubeletConfigurationResult build() {
            final var _resultValue = new GetKubeletConfigurationResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
