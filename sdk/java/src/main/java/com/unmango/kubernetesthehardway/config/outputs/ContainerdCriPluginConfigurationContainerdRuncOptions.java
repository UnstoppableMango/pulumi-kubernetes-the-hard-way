// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerdCriPluginConfigurationContainerdRuncOptions {
    /**
     * @return SystemdCgroup
     * 
     */
    private @Nullable Boolean systemdCgroup;

    private ContainerdCriPluginConfigurationContainerdRuncOptions() {}
    /**
     * @return SystemdCgroup
     * 
     */
    public Optional<Boolean> systemdCgroup() {
        return Optional.ofNullable(this.systemdCgroup);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerdCriPluginConfigurationContainerdRuncOptions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean systemdCgroup;
        public Builder() {}
        public Builder(ContainerdCriPluginConfigurationContainerdRuncOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.systemdCgroup = defaults.systemdCgroup;
        }

        @CustomType.Setter
        public Builder systemdCgroup(@Nullable Boolean systemdCgroup) {

            this.systemdCgroup = systemdCgroup;
            return this;
        }
        public ContainerdCriPluginConfigurationContainerdRuncOptions build() {
            final var _resultValue = new ContainerdCriPluginConfigurationContainerdRuncOptions();
            _resultValue.systemdCgroup = systemdCgroup;
            return _resultValue;
        }
    }
}
