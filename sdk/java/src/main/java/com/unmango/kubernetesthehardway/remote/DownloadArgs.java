// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote;

import com.pulumi.command.remote.inputs.ConnectionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DownloadArgs extends com.pulumi.resources.ResourceArgs {

    public static final DownloadArgs Empty = new DownloadArgs();

    /**
     * Connection details for the remote system
     * 
     */
    @Import(name="connection", required=true)
    private Output<ConnectionArgs> connection;

    /**
     * @return Connection details for the remote system
     * 
     */
    public Output<ConnectionArgs> connection() {
        return this.connection;
    }

    /**
     * The fully qualified path on the remote system where the file should be downloaded to.
     * 
     */
    @Import(name="destination", required=true)
    private Output<String> destination;

    /**
     * @return The fully qualified path on the remote system where the file should be downloaded to.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }

    /**
     * Remove the downloaded file when the resource is deleted.
     * 
     */
    @Import(name="removeOnDelete")
    private @Nullable Output<Boolean> removeOnDelete;

    /**
     * @return Remove the downloaded file when the resource is deleted.
     * 
     */
    public Optional<Output<Boolean>> removeOnDelete() {
        return Optional.ofNullable(this.removeOnDelete);
    }

    /**
     * The URL for the file to be downloaded.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The URL for the file to be downloaded.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private DownloadArgs() {}

    private DownloadArgs(DownloadArgs $) {
        this.connection = $.connection;
        this.destination = $.destination;
        this.removeOnDelete = $.removeOnDelete;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DownloadArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DownloadArgs $;

        public Builder() {
            $ = new DownloadArgs();
        }

        public Builder(DownloadArgs defaults) {
            $ = new DownloadArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connection Connection details for the remote system
         * 
         * @return builder
         * 
         */
        public Builder connection(Output<ConnectionArgs> connection) {
            $.connection = connection;
            return this;
        }

        /**
         * @param connection Connection details for the remote system
         * 
         * @return builder
         * 
         */
        public Builder connection(ConnectionArgs connection) {
            return connection(Output.of(connection));
        }

        /**
         * @param destination The fully qualified path on the remote system where the file should be downloaded to.
         * 
         * @return builder
         * 
         */
        public Builder destination(Output<String> destination) {
            $.destination = destination;
            return this;
        }

        /**
         * @param destination The fully qualified path on the remote system where the file should be downloaded to.
         * 
         * @return builder
         * 
         */
        public Builder destination(String destination) {
            return destination(Output.of(destination));
        }

        /**
         * @param removeOnDelete Remove the downloaded file when the resource is deleted.
         * 
         * @return builder
         * 
         */
        public Builder removeOnDelete(@Nullable Output<Boolean> removeOnDelete) {
            $.removeOnDelete = removeOnDelete;
            return this;
        }

        /**
         * @param removeOnDelete Remove the downloaded file when the resource is deleted.
         * 
         * @return builder
         * 
         */
        public Builder removeOnDelete(Boolean removeOnDelete) {
            return removeOnDelete(Output.of(removeOnDelete));
        }

        /**
         * @param url The URL for the file to be downloaded.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL for the file to be downloaded.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public DownloadArgs build() {
            if ($.connection == null) {
                throw new MissingRequiredPropertyException("DownloadArgs", "connection");
            }
            if ($.destination == null) {
                throw new MissingRequiredPropertyException("DownloadArgs", "destination");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("DownloadArgs", "url");
            }
            return $;
        }
    }

}