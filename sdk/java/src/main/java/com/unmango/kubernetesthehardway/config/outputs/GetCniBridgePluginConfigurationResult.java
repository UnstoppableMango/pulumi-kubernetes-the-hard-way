// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.config.outputs.CniBridgePluginConfiguration;
import java.util.Objects;

@CustomType
public final class GetCniBridgePluginConfigurationResult {
    private CniBridgePluginConfiguration result;

    private GetCniBridgePluginConfigurationResult() {}
    public CniBridgePluginConfiguration result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCniBridgePluginConfigurationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private CniBridgePluginConfiguration result;
        public Builder() {}
        public Builder(GetCniBridgePluginConfigurationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(CniBridgePluginConfiguration result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("GetCniBridgePluginConfigurationResult", "result");
            }
            this.result = result;
            return this;
        }
        public GetCniBridgePluginConfigurationResult build() {
            final var _resultValue = new GetCniBridgePluginConfigurationResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
