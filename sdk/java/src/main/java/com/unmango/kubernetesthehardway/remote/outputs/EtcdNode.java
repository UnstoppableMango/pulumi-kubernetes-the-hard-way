// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.outputs;

import com.pulumi.command.remote.outputs.Connection;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class EtcdNode {
    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    private Connection connection;
    /**
     * @return The internal IP of the node.
     * 
     */
    private String internalIp;

    private EtcdNode() {}
    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    public Connection connection() {
        return this.connection;
    }
    /**
     * @return The internal IP of the node.
     * 
     */
    public String internalIp() {
        return this.internalIp;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EtcdNode defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Connection connection;
        private String internalIp;
        public Builder() {}
        public Builder(EtcdNode defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connection = defaults.connection;
    	      this.internalIp = defaults.internalIp;
        }

        @CustomType.Setter
        public Builder connection(Connection connection) {
            if (connection == null) {
              throw new MissingRequiredPropertyException("EtcdNode", "connection");
            }
            this.connection = connection;
            return this;
        }
        @CustomType.Setter
        public Builder internalIp(String internalIp) {
            if (internalIp == null) {
              throw new MissingRequiredPropertyException("EtcdNode", "internalIp");
            }
            this.internalIp = internalIp;
            return this;
        }
        public EtcdNode build() {
            final var _resultValue = new EtcdNode();
            _resultValue.connection = connection;
            _resultValue.internalIp = internalIp;
            return _resultValue;
        }
    }
}
