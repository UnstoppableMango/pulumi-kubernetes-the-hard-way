// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.unmango.kubernetesthehardway.tools.enums.HostnamectlCommand;
import com.unmango.kubernetesthehardway.tools.enums.HostnamectlJsonMode;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * Abstraction over the `hostnamectl` utility on a remote system.
 * 
 */
public final class HostnamectlOptsArgs extends com.pulumi.resources.ResourceArgs {

    public static final HostnamectlOptsArgs Empty = new HostnamectlOptsArgs();

    /**
     * The argument for the specified `command`.
     * 
     */
    @Import(name="arg")
    private @Nullable Output<String> arg;

    /**
     * @return The argument for the specified `command`.
     * 
     */
    public Optional<Output<String>> arg() {
        return Optional.ofNullable(this.arg);
    }

    /**
     * Corresponds to the {COMMAND} argument.
     * 
     */
    @Import(name="command", required=true)
    private Output<HostnamectlCommand> command;

    /**
     * @return Corresponds to the {COMMAND} argument.
     * 
     */
    public Output<HostnamectlCommand> command() {
        return this.command;
    }

    /**
     * Print a short help text and exit.
     * 
     */
    @Import(name="help")
    private @Nullable Output<Boolean> help;

    /**
     * @return Print a short help text and exit.
     * 
     */
    public Optional<Output<Boolean>> help() {
        return Optional.ofNullable(this.help);
    }

    /**
     * Execute the operation remotely. Specify a hostname, or a username and hostname separated by &#39;{@literal @}&#39;, to connect to.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return Execute the operation remotely. Specify a hostname, or a username and hostname separated by &#39;{@literal @}&#39;, to connect to.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * Shows output formatted as JSON.
     * 
     */
    @Import(name="json")
    private @Nullable Output<HostnamectlJsonMode> json;

    /**
     * @return Shows output formatted as JSON.
     * 
     */
    public Optional<Output<HostnamectlJsonMode>> json() {
        return Optional.ofNullable(this.json);
    }

    /**
     * Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating &#39;{@literal @}&#39; character.
     * 
     */
    @Import(name="machine")
    private @Nullable Output<String> machine;

    /**
     * @return Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating &#39;{@literal @}&#39; character.
     * 
     */
    public Optional<Output<String>> machine() {
        return Optional.ofNullable(this.machine);
    }

    /**
     * Do not query the user for authentication for privileged operations.
     * 
     */
    @Import(name="noAskPassword")
    private @Nullable Output<Boolean> noAskPassword;

    /**
     * @return Do not query the user for authentication for privileged operations.
     * 
     */
    public Optional<Output<Boolean>> noAskPassword() {
        return Optional.ofNullable(this.noAskPassword);
    }

    /**
     * If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
     * 
     */
    @Import(name="pretty")
    private @Nullable Output<Boolean> pretty;

    /**
     * @return If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
     * 
     */
    public Optional<Output<Boolean>> pretty() {
        return Optional.ofNullable(this.pretty);
    }

    /**
     * If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
     * 
     */
    @Import(name="static")
    private @Nullable Output<Boolean> static_;

    /**
     * @return If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
     * 
     */
    public Optional<Output<Boolean>> static_() {
        return Optional.ofNullable(this.static_);
    }

    /**
     * If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
     * 
     */
    @Import(name="transient")
    private @Nullable Output<Boolean> transient_;

    /**
     * @return If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
     * 
     */
    public Optional<Output<Boolean>> transient_() {
        return Optional.ofNullable(this.transient_);
    }

    /**
     * Print a short version string and exit.
     * 
     */
    @Import(name="version")
    private @Nullable Output<Boolean> version;

    /**
     * @return Print a short version string and exit.
     * 
     */
    public Optional<Output<Boolean>> version() {
        return Optional.ofNullable(this.version);
    }

    private HostnamectlOptsArgs() {}

    private HostnamectlOptsArgs(HostnamectlOptsArgs $) {
        this.arg = $.arg;
        this.command = $.command;
        this.help = $.help;
        this.host = $.host;
        this.json = $.json;
        this.machine = $.machine;
        this.noAskPassword = $.noAskPassword;
        this.pretty = $.pretty;
        this.static_ = $.static_;
        this.transient_ = $.transient_;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HostnamectlOptsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HostnamectlOptsArgs $;

        public Builder() {
            $ = new HostnamectlOptsArgs();
        }

        public Builder(HostnamectlOptsArgs defaults) {
            $ = new HostnamectlOptsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param arg The argument for the specified `command`.
         * 
         * @return builder
         * 
         */
        public Builder arg(@Nullable Output<String> arg) {
            $.arg = arg;
            return this;
        }

        /**
         * @param arg The argument for the specified `command`.
         * 
         * @return builder
         * 
         */
        public Builder arg(String arg) {
            return arg(Output.of(arg));
        }

        /**
         * @param command Corresponds to the {COMMAND} argument.
         * 
         * @return builder
         * 
         */
        public Builder command(Output<HostnamectlCommand> command) {
            $.command = command;
            return this;
        }

        /**
         * @param command Corresponds to the {COMMAND} argument.
         * 
         * @return builder
         * 
         */
        public Builder command(HostnamectlCommand command) {
            return command(Output.of(command));
        }

        /**
         * @param help Print a short help text and exit.
         * 
         * @return builder
         * 
         */
        public Builder help(@Nullable Output<Boolean> help) {
            $.help = help;
            return this;
        }

        /**
         * @param help Print a short help text and exit.
         * 
         * @return builder
         * 
         */
        public Builder help(Boolean help) {
            return help(Output.of(help));
        }

        /**
         * @param host Execute the operation remotely. Specify a hostname, or a username and hostname separated by &#39;{@literal @}&#39;, to connect to.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host Execute the operation remotely. Specify a hostname, or a username and hostname separated by &#39;{@literal @}&#39;, to connect to.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param json Shows output formatted as JSON.
         * 
         * @return builder
         * 
         */
        public Builder json(@Nullable Output<HostnamectlJsonMode> json) {
            $.json = json;
            return this;
        }

        /**
         * @param json Shows output formatted as JSON.
         * 
         * @return builder
         * 
         */
        public Builder json(HostnamectlJsonMode json) {
            return json(Output.of(json));
        }

        /**
         * @param machine Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating &#39;{@literal @}&#39; character.
         * 
         * @return builder
         * 
         */
        public Builder machine(@Nullable Output<String> machine) {
            $.machine = machine;
            return this;
        }

        /**
         * @param machine Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating &#39;{@literal @}&#39; character.
         * 
         * @return builder
         * 
         */
        public Builder machine(String machine) {
            return machine(Output.of(machine));
        }

        /**
         * @param noAskPassword Do not query the user for authentication for privileged operations.
         * 
         * @return builder
         * 
         */
        public Builder noAskPassword(@Nullable Output<Boolean> noAskPassword) {
            $.noAskPassword = noAskPassword;
            return this;
        }

        /**
         * @param noAskPassword Do not query the user for authentication for privileged operations.
         * 
         * @return builder
         * 
         */
        public Builder noAskPassword(Boolean noAskPassword) {
            return noAskPassword(Output.of(noAskPassword));
        }

        /**
         * @param pretty If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
         * 
         * @return builder
         * 
         */
        public Builder pretty(@Nullable Output<Boolean> pretty) {
            $.pretty = pretty;
            return this;
        }

        /**
         * @param pretty If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
         * 
         * @return builder
         * 
         */
        public Builder pretty(Boolean pretty) {
            return pretty(Output.of(pretty));
        }

        /**
         * @param static_ If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
         * 
         * @return builder
         * 
         */
        public Builder static_(@Nullable Output<Boolean> static_) {
            $.static_ = static_;
            return this;
        }

        /**
         * @param static_ If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
         * 
         * @return builder
         * 
         */
        public Builder static_(Boolean static_) {
            return static_(Output.of(static_));
        }

        /**
         * @param transient_ If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
         * 
         * @return builder
         * 
         */
        public Builder transient_(@Nullable Output<Boolean> transient_) {
            $.transient_ = transient_;
            return this;
        }

        /**
         * @param transient_ If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
         * 
         * @return builder
         * 
         */
        public Builder transient_(Boolean transient_) {
            return transient_(Output.of(transient_));
        }

        /**
         * @param version Print a short version string and exit.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Boolean> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Print a short version string and exit.
         * 
         * @return builder
         * 
         */
        public Builder version(Boolean version) {
            return version(Output.of(version));
        }

        public HostnamectlOptsArgs build() {
            if ($.command == null) {
                throw new MissingRequiredPropertyException("HostnamectlOptsArgs", "command");
            }
            return $;
        }
    }

}
