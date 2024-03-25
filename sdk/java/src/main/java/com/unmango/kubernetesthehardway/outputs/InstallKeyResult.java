// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.RemoteFile;
import java.util.Objects;

@CustomType
public final class InstallKeyResult {
    private RemoteFile result;

    private InstallKeyResult() {}
    public RemoteFile result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstallKeyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private RemoteFile result;
        public Builder() {}
        public Builder(InstallKeyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(RemoteFile result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("InstallKeyResult", "result");
            }
            this.result = result;
            return this;
        }
        public InstallKeyResult build() {
            final var _resultValue = new InstallKeyResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
