// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.enums.SystemctlCommand;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SystemctlOpts {
    /**
     * @return Corresponds to the COMMAND argument.
     * 
     */
    private SystemctlCommand command;
    /**
     * @return Corresponds to the [PATTERN] argument
     * 
     */
    private @Nullable String pattern;
    /**
     * @return Corresponds to the [UNIT...] argument.
     * 
     */
    private String unit;

    private SystemctlOpts() {}
    /**
     * @return Corresponds to the COMMAND argument.
     * 
     */
    public SystemctlCommand command() {
        return this.command;
    }
    /**
     * @return Corresponds to the [PATTERN] argument
     * 
     */
    public Optional<String> pattern() {
        return Optional.ofNullable(this.pattern);
    }
    /**
     * @return Corresponds to the [UNIT...] argument.
     * 
     */
    public String unit() {
        return this.unit;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SystemctlOpts defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private SystemctlCommand command;
        private @Nullable String pattern;
        private String unit;
        public Builder() {}
        public Builder(SystemctlOpts defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.command = defaults.command;
    	      this.pattern = defaults.pattern;
    	      this.unit = defaults.unit;
        }

        @CustomType.Setter
        public Builder command(SystemctlCommand command) {
            if (command == null) {
              throw new MissingRequiredPropertyException("SystemctlOpts", "command");
            }
            this.command = command;
            return this;
        }
        @CustomType.Setter
        public Builder pattern(@Nullable String pattern) {

            this.pattern = pattern;
            return this;
        }
        @CustomType.Setter
        public Builder unit(String unit) {
            if (unit == null) {
              throw new MissingRequiredPropertyException("SystemctlOpts", "unit");
            }
            this.unit = unit;
            return this;
        }
        public SystemctlOpts build() {
            final var _resultValue = new SystemctlOpts();
            _resultValue.command = command;
            _resultValue.pattern = pattern;
            _resultValue.unit = unit;
            return _resultValue;
        }
    }
}