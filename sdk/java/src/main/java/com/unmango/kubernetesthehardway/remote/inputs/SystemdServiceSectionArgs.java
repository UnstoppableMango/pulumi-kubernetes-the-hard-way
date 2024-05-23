// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.unmango.kubernetesthehardway.remote.enums.SystemdKillMode;
import com.unmango.kubernetesthehardway.remote.enums.SystemdServiceExitType;
import com.unmango.kubernetesthehardway.remote.enums.SystemdServiceRestart;
import com.unmango.kubernetesthehardway.remote.enums.SystemdServiceType;
import com.unmango.kubernetesthehardway.remote.inputs.SystemDelegate;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
 * 
 */
public final class SystemdServiceSectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final SystemdServiceSectionArgs Empty = new SystemdServiceSectionArgs();

    /**
     * Turns on delegation of further resource control partitioning to processes of the unit.
     * 
     */
    @Import(name="delegate")
    private @Nullable Output<SystemDelegate> delegate;

    /**
     * @return Turns on delegation of further resource control partitioning to processes of the unit.
     * 
     */
    public Optional<Output<SystemDelegate>> delegate() {
        return Optional.ofNullable(this.delegate);
    }

    /**
     * Commands that are executed when this service is started.
     * 
     */
    @Import(name="execStart")
    private @Nullable Output<String> execStart;

    /**
     * @return Commands that are executed when this service is started.
     * 
     */
    public Optional<Output<String>> execStart() {
        return Optional.ofNullable(this.execStart);
    }

    /**
     * Additional commands that are executed before the command in ExecStart=.
     * 
     */
    @Import(name="execStartPre")
    private @Nullable Output<String> execStartPre;

    /**
     * @return Additional commands that are executed before the command in ExecStart=.
     * 
     */
    public Optional<Output<String>> execStartPre() {
        return Optional.ofNullable(this.execStartPre);
    }

    /**
     * Specifies when the manager should consider the service to be finished.
     * 
     */
    @Import(name="exitType")
    private @Nullable Output<SystemdServiceExitType> exitType;

    /**
     * @return Specifies when the manager should consider the service to be finished.
     * 
     */
    public Optional<Output<SystemdServiceExitType>> exitType() {
        return Optional.ofNullable(this.exitType);
    }

    /**
     * Specifies how processes of this unit shall be killed.
     * 
     */
    @Import(name="killMode")
    private @Nullable Output<SystemdKillMode> killMode;

    /**
     * @return Specifies how processes of this unit shall be killed.
     * 
     */
    public Optional<Output<SystemdKillMode>> killMode() {
        return Optional.ofNullable(this.killMode);
    }

    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    @Import(name="limitCore")
    private @Nullable Output<String> limitCore;

    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    public Optional<Output<String>> limitCore() {
        return Optional.ofNullable(this.limitCore);
    }

    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    @Import(name="limitNProc")
    private @Nullable Output<String> limitNProc;

    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    public Optional<Output<String>> limitNProc() {
        return Optional.ofNullable(this.limitNProc);
    }

    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    @Import(name="limitNoFile")
    private @Nullable Output<Integer> limitNoFile;

    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    public Optional<Output<Integer>> limitNoFile() {
        return Optional.ofNullable(this.limitNoFile);
    }

    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
     * 
     */
    @Import(name="oomScoreAdjust")
    private @Nullable Output<Integer> oomScoreAdjust;

    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
     * 
     */
    public Optional<Output<Integer>> oomScoreAdjust() {
        return Optional.ofNullable(this.oomScoreAdjust);
    }

    /**
     * Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
     * 
     */
    @Import(name="restart")
    private @Nullable Output<SystemdServiceRestart> restart;

    /**
     * @return Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
     * 
     */
    public Optional<Output<SystemdServiceRestart>> restart() {
        return Optional.ofNullable(this.restart);
    }

    /**
     * Configures the time to sleep before restarting a service (as configured with Restart=).
     * 
     */
    @Import(name="restartSec")
    private @Nullable Output<String> restartSec;

    /**
     * @return Configures the time to sleep before restarting a service (as configured with Restart=).
     * 
     */
    public Optional<Output<String>> restartSec() {
        return Optional.ofNullable(this.restartSec);
    }

    /**
     * Configures the mechanism via which the service notifies the manager that the service start-up has finished.
     * 
     */
    @Import(name="type")
    private @Nullable Output<SystemdServiceType> type;

    /**
     * @return Configures the mechanism via which the service notifies the manager that the service start-up has finished.
     * 
     */
    public Optional<Output<SystemdServiceType>> type() {
        return Optional.ofNullable(this.type);
    }

    private SystemdServiceSectionArgs() {}

    private SystemdServiceSectionArgs(SystemdServiceSectionArgs $) {
        this.delegate = $.delegate;
        this.execStart = $.execStart;
        this.execStartPre = $.execStartPre;
        this.exitType = $.exitType;
        this.killMode = $.killMode;
        this.limitCore = $.limitCore;
        this.limitNProc = $.limitNProc;
        this.limitNoFile = $.limitNoFile;
        this.oomScoreAdjust = $.oomScoreAdjust;
        this.restart = $.restart;
        this.restartSec = $.restartSec;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SystemdServiceSectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SystemdServiceSectionArgs $;

        public Builder() {
            $ = new SystemdServiceSectionArgs();
        }

        public Builder(SystemdServiceSectionArgs defaults) {
            $ = new SystemdServiceSectionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param delegate Turns on delegation of further resource control partitioning to processes of the unit.
         * 
         * @return builder
         * 
         */
        public Builder delegate(@Nullable Output<SystemDelegate> delegate) {
            $.delegate = delegate;
            return this;
        }

        /**
         * @param delegate Turns on delegation of further resource control partitioning to processes of the unit.
         * 
         * @return builder
         * 
         */
        public Builder delegate(SystemDelegate delegate) {
            return delegate(Output.of(delegate));
        }

        /**
         * @param execStart Commands that are executed when this service is started.
         * 
         * @return builder
         * 
         */
        public Builder execStart(@Nullable Output<String> execStart) {
            $.execStart = execStart;
            return this;
        }

        /**
         * @param execStart Commands that are executed when this service is started.
         * 
         * @return builder
         * 
         */
        public Builder execStart(String execStart) {
            return execStart(Output.of(execStart));
        }

        /**
         * @param execStartPre Additional commands that are executed before the command in ExecStart=.
         * 
         * @return builder
         * 
         */
        public Builder execStartPre(@Nullable Output<String> execStartPre) {
            $.execStartPre = execStartPre;
            return this;
        }

        /**
         * @param execStartPre Additional commands that are executed before the command in ExecStart=.
         * 
         * @return builder
         * 
         */
        public Builder execStartPre(String execStartPre) {
            return execStartPre(Output.of(execStartPre));
        }

        /**
         * @param exitType Specifies when the manager should consider the service to be finished.
         * 
         * @return builder
         * 
         */
        public Builder exitType(@Nullable Output<SystemdServiceExitType> exitType) {
            $.exitType = exitType;
            return this;
        }

        /**
         * @param exitType Specifies when the manager should consider the service to be finished.
         * 
         * @return builder
         * 
         */
        public Builder exitType(SystemdServiceExitType exitType) {
            return exitType(Output.of(exitType));
        }

        /**
         * @param killMode Specifies how processes of this unit shall be killed.
         * 
         * @return builder
         * 
         */
        public Builder killMode(@Nullable Output<SystemdKillMode> killMode) {
            $.killMode = killMode;
            return this;
        }

        /**
         * @param killMode Specifies how processes of this unit shall be killed.
         * 
         * @return builder
         * 
         */
        public Builder killMode(SystemdKillMode killMode) {
            return killMode(Output.of(killMode));
        }

        /**
         * @param limitCore https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
         * 
         * @return builder
         * 
         */
        public Builder limitCore(@Nullable Output<String> limitCore) {
            $.limitCore = limitCore;
            return this;
        }

        /**
         * @param limitCore https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
         * 
         * @return builder
         * 
         */
        public Builder limitCore(String limitCore) {
            return limitCore(Output.of(limitCore));
        }

        /**
         * @param limitNProc https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
         * 
         * @return builder
         * 
         */
        public Builder limitNProc(@Nullable Output<String> limitNProc) {
            $.limitNProc = limitNProc;
            return this;
        }

        /**
         * @param limitNProc https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
         * 
         * @return builder
         * 
         */
        public Builder limitNProc(String limitNProc) {
            return limitNProc(Output.of(limitNProc));
        }

        /**
         * @param limitNoFile https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
         * 
         * @return builder
         * 
         */
        public Builder limitNoFile(@Nullable Output<Integer> limitNoFile) {
            $.limitNoFile = limitNoFile;
            return this;
        }

        /**
         * @param limitNoFile https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
         * 
         * @return builder
         * 
         */
        public Builder limitNoFile(Integer limitNoFile) {
            return limitNoFile(Output.of(limitNoFile));
        }

        /**
         * @param oomScoreAdjust https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
         * 
         * @return builder
         * 
         */
        public Builder oomScoreAdjust(@Nullable Output<Integer> oomScoreAdjust) {
            $.oomScoreAdjust = oomScoreAdjust;
            return this;
        }

        /**
         * @param oomScoreAdjust https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
         * 
         * @return builder
         * 
         */
        public Builder oomScoreAdjust(Integer oomScoreAdjust) {
            return oomScoreAdjust(Output.of(oomScoreAdjust));
        }

        /**
         * @param restart Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
         * 
         * @return builder
         * 
         */
        public Builder restart(@Nullable Output<SystemdServiceRestart> restart) {
            $.restart = restart;
            return this;
        }

        /**
         * @param restart Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
         * 
         * @return builder
         * 
         */
        public Builder restart(SystemdServiceRestart restart) {
            return restart(Output.of(restart));
        }

        /**
         * @param restartSec Configures the time to sleep before restarting a service (as configured with Restart=).
         * 
         * @return builder
         * 
         */
        public Builder restartSec(@Nullable Output<String> restartSec) {
            $.restartSec = restartSec;
            return this;
        }

        /**
         * @param restartSec Configures the time to sleep before restarting a service (as configured with Restart=).
         * 
         * @return builder
         * 
         */
        public Builder restartSec(String restartSec) {
            return restartSec(Output.of(restartSec));
        }

        /**
         * @param type Configures the mechanism via which the service notifies the manager that the service start-up has finished.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<SystemdServiceType> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Configures the mechanism via which the service notifies the manager that the service start-up has finished.
         * 
         * @return builder
         * 
         */
        public Builder type(SystemdServiceType type) {
            return type(Output.of(type));
        }

        public SystemdServiceSectionArgs build() {
            return $;
        }
    }

}
