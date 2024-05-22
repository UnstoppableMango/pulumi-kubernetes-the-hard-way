// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * containerd cri plugin configuration.
 * 
 */
public final class ContainerdCriPluginConfigurationCni extends com.pulumi.resources.InvokeArgs {

    public static final ContainerdCriPluginConfigurationCni Empty = new ContainerdCriPluginConfigurationCni();

    /**
     * bin_dir
     * 
     */
    @Import(name="binDir")
    private @Nullable String binDir;

    /**
     * @return bin_dir
     * 
     */
    public Optional<String> binDir() {
        return Optional.ofNullable(this.binDir);
    }

    /**
     * conf_dir
     * 
     */
    @Import(name="confDir")
    private @Nullable String confDir;

    /**
     * @return conf_dir
     * 
     */
    public Optional<String> confDir() {
        return Optional.ofNullable(this.confDir);
    }

    private ContainerdCriPluginConfigurationCni() {}

    private ContainerdCriPluginConfigurationCni(ContainerdCriPluginConfigurationCni $) {
        this.binDir = $.binDir;
        this.confDir = $.confDir;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerdCriPluginConfigurationCni defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerdCriPluginConfigurationCni $;

        public Builder() {
            $ = new ContainerdCriPluginConfigurationCni();
        }

        public Builder(ContainerdCriPluginConfigurationCni defaults) {
            $ = new ContainerdCriPluginConfigurationCni(Objects.requireNonNull(defaults));
        }

        /**
         * @param binDir bin_dir
         * 
         * @return builder
         * 
         */
        public Builder binDir(@Nullable String binDir) {
            $.binDir = binDir;
            return this;
        }

        /**
         * @param confDir conf_dir
         * 
         * @return builder
         * 
         */
        public Builder confDir(@Nullable String confDir) {
            $.confDir = confDir;
            return this;
        }

        public ContainerdCriPluginConfigurationCni build() {
            return $;
        }
    }

}
