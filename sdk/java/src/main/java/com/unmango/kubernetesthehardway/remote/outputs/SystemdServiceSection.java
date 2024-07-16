// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.outputs;

import com.pulumi.core.annotations.CustomType;
import com.unmango.kubernetesthehardway.remote.enums.SystemdDelegate;
import com.unmango.kubernetesthehardway.remote.enums.SystemdKillMode;
import com.unmango.kubernetesthehardway.remote.enums.SystemdServiceExitType;
import com.unmango.kubernetesthehardway.remote.enums.SystemdServiceRestart;
import com.unmango.kubernetesthehardway.remote.enums.SystemdServiceType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SystemdServiceSection {
    /**
     * @return Turns on delegation of further resource control partitioning to processes of the unit.
     * 
     */
    private @Nullable SystemdDelegate delegate;
    private @Nullable List<String> environment;
    /**
     * @return Commands that are executed when this service is started.
     * 
     */
    private @Nullable String execStart;
    /**
     * @return Additional commands that are executed before the command in ExecStart=.
     * 
     */
    private @Nullable String execStartPre;
    /**
     * @return Specifies when the manager should consider the service to be finished.
     * 
     */
    private @Nullable SystemdServiceExitType exitType;
    /**
     * @return Specifies how processes of this unit shall be killed.
     * 
     */
    private @Nullable SystemdKillMode killMode;
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    private @Nullable String limitCore;
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    private @Nullable String limitNProc;
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    private @Nullable Integer limitNoFile;
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
     * 
     */
    private @Nullable Integer oomScoreAdjust;
    /**
     * @return Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
     * 
     */
    private @Nullable SystemdServiceRestart restart;
    /**
     * @return Configures the time to sleep before restarting a service (as configured with Restart=).
     * 
     */
    private @Nullable String restartSec;
    /**
     * @return Configure unit start rate limiting. Units which are started more than burst times within an interval time span are not permitted to start any more. Use StartLimitIntervalSec= to configure the checking interval and StartLimitBurst= to configure how many starts per interval are allowed.
     * 
     */
    private @Nullable String startLimitInterval;
    /**
     * @return Configures the mechanism via which the service notifies the manager that the service start-up has finished.
     * 
     */
    private @Nullable SystemdServiceType type;

    private SystemdServiceSection() {}
    /**
     * @return Turns on delegation of further resource control partitioning to processes of the unit.
     * 
     */
    public Optional<SystemdDelegate> delegate() {
        return Optional.ofNullable(this.delegate);
    }
    public List<String> environment() {
        return this.environment == null ? List.of() : this.environment;
    }
    /**
     * @return Commands that are executed when this service is started.
     * 
     */
    public Optional<String> execStart() {
        return Optional.ofNullable(this.execStart);
    }
    /**
     * @return Additional commands that are executed before the command in ExecStart=.
     * 
     */
    public Optional<String> execStartPre() {
        return Optional.ofNullable(this.execStartPre);
    }
    /**
     * @return Specifies when the manager should consider the service to be finished.
     * 
     */
    public Optional<SystemdServiceExitType> exitType() {
        return Optional.ofNullable(this.exitType);
    }
    /**
     * @return Specifies how processes of this unit shall be killed.
     * 
     */
    public Optional<SystemdKillMode> killMode() {
        return Optional.ofNullable(this.killMode);
    }
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    public Optional<String> limitCore() {
        return Optional.ofNullable(this.limitCore);
    }
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    public Optional<String> limitNProc() {
        return Optional.ofNullable(this.limitNProc);
    }
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
     * 
     */
    public Optional<Integer> limitNoFile() {
        return Optional.ofNullable(this.limitNoFile);
    }
    /**
     * @return https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
     * 
     */
    public Optional<Integer> oomScoreAdjust() {
        return Optional.ofNullable(this.oomScoreAdjust);
    }
    /**
     * @return Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
     * 
     */
    public Optional<SystemdServiceRestart> restart() {
        return Optional.ofNullable(this.restart);
    }
    /**
     * @return Configures the time to sleep before restarting a service (as configured with Restart=).
     * 
     */
    public Optional<String> restartSec() {
        return Optional.ofNullable(this.restartSec);
    }
    /**
     * @return Configure unit start rate limiting. Units which are started more than burst times within an interval time span are not permitted to start any more. Use StartLimitIntervalSec= to configure the checking interval and StartLimitBurst= to configure how many starts per interval are allowed.
     * 
     */
    public Optional<String> startLimitInterval() {
        return Optional.ofNullable(this.startLimitInterval);
    }
    /**
     * @return Configures the mechanism via which the service notifies the manager that the service start-up has finished.
     * 
     */
    public Optional<SystemdServiceType> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SystemdServiceSection defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable SystemdDelegate delegate;
        private @Nullable List<String> environment;
        private @Nullable String execStart;
        private @Nullable String execStartPre;
        private @Nullable SystemdServiceExitType exitType;
        private @Nullable SystemdKillMode killMode;
        private @Nullable String limitCore;
        private @Nullable String limitNProc;
        private @Nullable Integer limitNoFile;
        private @Nullable Integer oomScoreAdjust;
        private @Nullable SystemdServiceRestart restart;
        private @Nullable String restartSec;
        private @Nullable String startLimitInterval;
        private @Nullable SystemdServiceType type;
        public Builder() {}
        public Builder(SystemdServiceSection defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.delegate = defaults.delegate;
    	      this.environment = defaults.environment;
    	      this.execStart = defaults.execStart;
    	      this.execStartPre = defaults.execStartPre;
    	      this.exitType = defaults.exitType;
    	      this.killMode = defaults.killMode;
    	      this.limitCore = defaults.limitCore;
    	      this.limitNProc = defaults.limitNProc;
    	      this.limitNoFile = defaults.limitNoFile;
    	      this.oomScoreAdjust = defaults.oomScoreAdjust;
    	      this.restart = defaults.restart;
    	      this.restartSec = defaults.restartSec;
    	      this.startLimitInterval = defaults.startLimitInterval;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder delegate(@Nullable SystemdDelegate delegate) {

            this.delegate = delegate;
            return this;
        }
        @CustomType.Setter
        public Builder environment(@Nullable List<String> environment) {

            this.environment = environment;
            return this;
        }
        public Builder environment(String... environment) {
            return environment(List.of(environment));
        }
        @CustomType.Setter
        public Builder execStart(@Nullable String execStart) {

            this.execStart = execStart;
            return this;
        }
        @CustomType.Setter
        public Builder execStartPre(@Nullable String execStartPre) {

            this.execStartPre = execStartPre;
            return this;
        }
        @CustomType.Setter
        public Builder exitType(@Nullable SystemdServiceExitType exitType) {

            this.exitType = exitType;
            return this;
        }
        @CustomType.Setter
        public Builder killMode(@Nullable SystemdKillMode killMode) {

            this.killMode = killMode;
            return this;
        }
        @CustomType.Setter
        public Builder limitCore(@Nullable String limitCore) {

            this.limitCore = limitCore;
            return this;
        }
        @CustomType.Setter
        public Builder limitNProc(@Nullable String limitNProc) {

            this.limitNProc = limitNProc;
            return this;
        }
        @CustomType.Setter
        public Builder limitNoFile(@Nullable Integer limitNoFile) {

            this.limitNoFile = limitNoFile;
            return this;
        }
        @CustomType.Setter
        public Builder oomScoreAdjust(@Nullable Integer oomScoreAdjust) {

            this.oomScoreAdjust = oomScoreAdjust;
            return this;
        }
        @CustomType.Setter
        public Builder restart(@Nullable SystemdServiceRestart restart) {

            this.restart = restart;
            return this;
        }
        @CustomType.Setter
        public Builder restartSec(@Nullable String restartSec) {

            this.restartSec = restartSec;
            return this;
        }
        @CustomType.Setter
        public Builder startLimitInterval(@Nullable String startLimitInterval) {

            this.startLimitInterval = startLimitInterval;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable SystemdServiceType type) {

            this.type = type;
            return this;
        }
        public SystemdServiceSection build() {
            final var _resultValue = new SystemdServiceSection();
            _resultValue.delegate = delegate;
            _resultValue.environment = environment;
            _resultValue.execStart = execStart;
            _resultValue.execStartPre = execStartPre;
            _resultValue.exitType = exitType;
            _resultValue.killMode = killMode;
            _resultValue.limitCore = limitCore;
            _resultValue.limitNProc = limitNProc;
            _resultValue.limitNoFile = limitNoFile;
            _resultValue.oomScoreAdjust = oomScoreAdjust;
            _resultValue.restart = restart;
            _resultValue.restartSec = restartSec;
            _resultValue.startLimitInterval = startLimitInterval;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
