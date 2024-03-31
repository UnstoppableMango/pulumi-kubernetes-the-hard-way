// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.config.outputs.ContextContext;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class Context {
    private ContextContext context;
    private String name;

    private Context() {}
    public ContextContext context() {
        return this.context;
    }
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(Context defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private ContextContext context;
        private String name;
        public Builder() {}
        public Builder(Context defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.context = defaults.context;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder context(ContextContext context) {
            if (context == null) {
              throw new MissingRequiredPropertyException("Context", "context");
            }
            this.context = context;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("Context", "name");
            }
            this.name = name;
            return this;
        }
        public Context build() {
            final var _resultValue = new Context();
            _resultValue.context = context;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
