// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tls.outputs;

import com.pulumi.core.annotations.CustomType;
import com.unmango.kubernetesthehardway.tls.enums.NodeRole;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterPkiNode {
    /**
     * @return The IP address of the node
     * 
     */
    private @Nullable String ip;
    /**
     * @return The role a node should be configured for
     * 
     */
    private @Nullable NodeRole role;

    private ClusterPkiNode() {}
    /**
     * @return The IP address of the node
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return The role a node should be configured for
     * 
     */
    public Optional<NodeRole> role() {
        return Optional.ofNullable(this.role);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterPkiNode defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ip;
        private @Nullable NodeRole role;
        public Builder() {}
        public Builder(ClusterPkiNode defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.role = defaults.role;
        }

        @CustomType.Setter
        public Builder ip(@Nullable String ip) {

            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder role(@Nullable NodeRole role) {

            this.role = role;
            return this;
        }
        public ClusterPkiNode build() {
            final var _resultValue = new ClusterPkiNode();
            _resultValue.ip = ip;
            _resultValue.role = role;
            return _resultValue;
        }
    }
}