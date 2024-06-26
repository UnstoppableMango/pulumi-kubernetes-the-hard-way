// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.inputs;

import com.pulumi.core.annotations.Import;
import com.unmango.kubernetesthehardway.config.inputs.ContainerdCriPluginConfigurationArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetContainerdConfigurationArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerdConfigurationArgs Empty = new GetContainerdConfigurationArgs();

    /**
     * The cri configuration.
     * 
     */
    @Import(name="cri")
    private @Nullable ContainerdCriPluginConfigurationArgs cri;

    /**
     * @return The cri configuration.
     * 
     */
    public Optional<ContainerdCriPluginConfigurationArgs> cri() {
        return Optional.ofNullable(this.cri);
    }

    private GetContainerdConfigurationArgs() {}

    private GetContainerdConfigurationArgs(GetContainerdConfigurationArgs $) {
        this.cri = $.cri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerdConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerdConfigurationArgs $;

        public Builder() {
            $ = new GetContainerdConfigurationArgs();
        }

        public Builder(GetContainerdConfigurationArgs defaults) {
            $ = new GetContainerdConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cri The cri configuration.
         * 
         * @return builder
         * 
         */
        public Builder cri(@Nullable ContainerdCriPluginConfigurationArgs cri) {
            $.cri = cri;
            return this;
        }

        public GetContainerdConfigurationArgs build() {
            return $;
        }
    }

}
