// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.config.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * The CNI plugins IPAM
 * 
 */
public final class CniBridgeIpamArgs extends com.pulumi.resources.ResourceArgs {

    public static final CniBridgeIpamArgs Empty = new CniBridgeIpamArgs();

    /**
     * IPAM ranges.
     * 
     */
    @Import(name="ranges")
    private @Nullable Output<List<Map<String,String>>> ranges;

    /**
     * @return IPAM ranges.
     * 
     */
    public Optional<Output<List<Map<String,String>>>> ranges() {
        return Optional.ofNullable(this.ranges);
    }

    /**
     * IPAM routes.
     * 
     */
    @Import(name="routes")
    private @Nullable Output<List<Map<String,String>>> routes;

    /**
     * @return IPAM routes.
     * 
     */
    public Optional<Output<List<Map<String,String>>>> routes() {
        return Optional.ofNullable(this.routes);
    }

    /**
     * CNI bridge IPAM type
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return CNI bridge IPAM type
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private CniBridgeIpamArgs() {}

    private CniBridgeIpamArgs(CniBridgeIpamArgs $) {
        this.ranges = $.ranges;
        this.routes = $.routes;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CniBridgeIpamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CniBridgeIpamArgs $;

        public Builder() {
            $ = new CniBridgeIpamArgs();
        }

        public Builder(CniBridgeIpamArgs defaults) {
            $ = new CniBridgeIpamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ranges IPAM ranges.
         * 
         * @return builder
         * 
         */
        public Builder ranges(@Nullable Output<List<Map<String,String>>> ranges) {
            $.ranges = ranges;
            return this;
        }

        /**
         * @param ranges IPAM ranges.
         * 
         * @return builder
         * 
         */
        public Builder ranges(List<Map<String,String>> ranges) {
            return ranges(Output.of(ranges));
        }

        /**
         * @param ranges IPAM ranges.
         * 
         * @return builder
         * 
         */
        public Builder ranges(Map<String,String>... ranges) {
            return ranges(List.of(ranges));
        }

        /**
         * @param routes IPAM routes.
         * 
         * @return builder
         * 
         */
        public Builder routes(@Nullable Output<List<Map<String,String>>> routes) {
            $.routes = routes;
            return this;
        }

        /**
         * @param routes IPAM routes.
         * 
         * @return builder
         * 
         */
        public Builder routes(List<Map<String,String>> routes) {
            return routes(Output.of(routes));
        }

        /**
         * @param routes IPAM routes.
         * 
         * @return builder
         * 
         */
        public Builder routes(Map<String,String>... routes) {
            return routes(List.of(routes));
        }

        /**
         * @param type CNI bridge IPAM type
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type CNI bridge IPAM type
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public CniBridgeIpamArgs build() {
            return $;
        }
    }

}
