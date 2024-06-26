// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.unmango.kubernetesthehardway.config.inputs.ContainerdCriPluginConfigurationArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * The containerd configuration.
 * 
 */
public final class ContainerdConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerdConfigurationArgs Empty = new ContainerdConfigurationArgs();

    /**
     * The cri configuration.
     * 
     */
    @Import(name="cri")
    private @Nullable Output<ContainerdCriPluginConfigurationArgs> cri;

    /**
     * @return The cri configuration.
     * 
     */
    public Optional<Output<ContainerdCriPluginConfigurationArgs>> cri() {
        return Optional.ofNullable(this.cri);
    }

    private ContainerdConfigurationArgs() {}

    private ContainerdConfigurationArgs(ContainerdConfigurationArgs $) {
        this.cri = $.cri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerdConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerdConfigurationArgs $;

        public Builder() {
            $ = new ContainerdConfigurationArgs();
        }

        public Builder(ContainerdConfigurationArgs defaults) {
            $ = new ContainerdConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cri The cri configuration.
         * 
         * @return builder
         * 
         */
        public Builder cri(@Nullable Output<ContainerdCriPluginConfigurationArgs> cri) {
            $.cri = cri;
            return this;
        }

        /**
         * @param cri The cri configuration.
         * 
         * @return builder
         * 
         */
        public Builder cri(ContainerdCriPluginConfigurationArgs cri) {
            return cri(Output.of(cri));
        }

        public ContainerdConfigurationArgs build() {
            return $;
        }
    }

}
