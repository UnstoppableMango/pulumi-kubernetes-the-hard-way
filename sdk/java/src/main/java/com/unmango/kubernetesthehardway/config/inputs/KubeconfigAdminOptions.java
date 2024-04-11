// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeconfigAdminOptions extends com.pulumi.resources.InvokeArgs {

    public static final KubeconfigAdminOptions Empty = new KubeconfigAdminOptions();

    @Import(name="publicIp")
    private @Nullable String publicIp;

    public Optional<String> publicIp() {
        return Optional.ofNullable(this.publicIp);
    }

    @Import(name="type", required=true)
    private String type;

    public String type() {
        return this.type;
    }

    private KubeconfigAdminOptions() {}

    private KubeconfigAdminOptions(KubeconfigAdminOptions $) {
        this.publicIp = $.publicIp;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeconfigAdminOptions defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeconfigAdminOptions $;

        public Builder() {
            $ = new KubeconfigAdminOptions();
        }

        public Builder(KubeconfigAdminOptions defaults) {
            $ = new KubeconfigAdminOptions(Objects.requireNonNull(defaults));
        }

        public Builder publicIp(@Nullable String publicIp) {
            $.publicIp = publicIp;
            return this;
        }

        public Builder type(String type) {
            $.type = type;
            return this;
        }

        public KubeconfigAdminOptions build() {
            $.type = Codegen.stringProp("type").arg($.type).require();
            return $;
        }
    }

}