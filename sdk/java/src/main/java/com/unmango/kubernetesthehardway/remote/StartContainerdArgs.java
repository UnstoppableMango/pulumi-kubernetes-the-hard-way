// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.Objects;


public final class StartContainerdArgs extends com.pulumi.resources.ResourceArgs {

    public static final StartContainerdArgs Empty = new StartContainerdArgs();

    /**
     * The parameters with which to connect to the remote host.
     * 
     */
    @Import(name="connection", required=true)
    private Output<ConnectionArgs> connection;

    /**
     * @return The parameters with which to connect to the remote host.
     * 
     */
    public Output<ConnectionArgs> connection() {
        return this.connection;
    }

    private StartContainerdArgs() {}

    private StartContainerdArgs(StartContainerdArgs $) {
        this.connection = $.connection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StartContainerdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StartContainerdArgs $;

        public Builder() {
            $ = new StartContainerdArgs();
        }

        public Builder(StartContainerdArgs defaults) {
            $ = new StartContainerdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connection The parameters with which to connect to the remote host.
         * 
         * @return builder
         * 
         */
        public Builder connection(Output<ConnectionArgs> connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param connection The parameters with which to connect to the remote host.
         * 
         * @return builder
         * 
         */
        public Builder connection(ConnectionArgs connection) {
            return connection(Output.of(connection));
        }

        public StartContainerdArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("StartContainerdArgs", "connection");
            }
            return $;
        }
    }

}