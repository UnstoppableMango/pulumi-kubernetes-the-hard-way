// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options
 * 
 */
public final class SystemdInstallSectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final SystemdInstallSectionArgs Empty = new SystemdInstallSectionArgs();

    /**
     * A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
     * 
     */
    @Import(name="wantedBy")
    private @Nullable Output<List<String>> wantedBy;

    /**
     * @return A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
     * 
     */
    public Optional<Output<List<String>>> wantedBy() {
        return Optional.ofNullable(this.wantedBy);
    }

    private SystemdInstallSectionArgs() {}

    private SystemdInstallSectionArgs(SystemdInstallSectionArgs $) {
        this.wantedBy = $.wantedBy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SystemdInstallSectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SystemdInstallSectionArgs $;

        public Builder() {
            $ = new SystemdInstallSectionArgs();
        }

        public Builder(SystemdInstallSectionArgs defaults) {
            $ = new SystemdInstallSectionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param wantedBy A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
         * 
         * @return builder
         * 
         */
        public Builder wantedBy(@Nullable Output<List<String>> wantedBy) {
            $.wantedBy = wantedBy;
            return this;
        }

        /**
         * @param wantedBy A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
         * 
         * @return builder
         * 
         */
        public Builder wantedBy(List<String> wantedBy) {
            return wantedBy(Output.of(wantedBy));
        }

        /**
         * @param wantedBy A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
         * 
         * @return builder
         * 
         */
        public Builder wantedBy(String... wantedBy) {
            return wantedBy(List.of(wantedBy));
        }

        public SystemdInstallSectionArgs build() {
            return $;
        }
    }

}