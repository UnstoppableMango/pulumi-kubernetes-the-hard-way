// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.inputs;

import com.pulumi.core.annotations.Import;
import com.unmango.kubernetesthehardway.config.inputs.ContainerdCriPluginConfiguration;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetContainerdConfigurationPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerdConfigurationPlainArgs Empty = new GetContainerdConfigurationPlainArgs();

    /**
     * The cri configuration.
     * 
     */
    @Import(name="cri")
    private @Nullable ContainerdCriPluginConfiguration cri;

    /**
     * @return The cri configuration.
     * 
     */
    public Optional<ContainerdCriPluginConfiguration> cri() {
        return Optional.ofNullable(this.cri);
    }

    private GetContainerdConfigurationPlainArgs() {}

    private GetContainerdConfigurationPlainArgs(GetContainerdConfigurationPlainArgs $) {
        this.cri = $.cri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerdConfigurationPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerdConfigurationPlainArgs $;

        public Builder() {
            $ = new GetContainerdConfigurationPlainArgs();
        }

        public Builder(GetContainerdConfigurationPlainArgs defaults) {
            $ = new GetContainerdConfigurationPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cri The cri configuration.
         * 
         * @return builder
         * 
         */
        public Builder cri(@Nullable ContainerdCriPluginConfiguration cri) {
            $.cri = cri;
            return this;
        }

        public GetContainerdConfigurationPlainArgs build() {
            return $;
        }
    }

}
