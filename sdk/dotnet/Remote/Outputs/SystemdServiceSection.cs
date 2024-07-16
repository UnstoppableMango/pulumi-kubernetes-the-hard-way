// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Remote.Outputs
{

    /// <summary>
    /// https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
    /// </summary>
    [OutputType]
    public sealed class SystemdServiceSection
    {
        /// <summary>
        /// Turns on delegation of further resource control partitioning to processes of the unit.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Remote.SystemdDelegate? Delegate;
        public readonly ImmutableArray<string> Environment;
        /// <summary>
        /// Commands that are executed when this service is started.
        /// </summary>
        public readonly string? ExecStart;
        /// <summary>
        /// Additional commands that are executed before the command in ExecStart=.
        /// </summary>
        public readonly string? ExecStartPre;
        /// <summary>
        /// Specifies when the manager should consider the service to be finished.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Remote.SystemdServiceExitType? ExitType;
        /// <summary>
        /// Specifies how processes of this unit shall be killed.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Remote.SystemdKillMode? KillMode;
        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        /// </summary>
        public readonly string? LimitCore;
        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        /// </summary>
        public readonly string? LimitNProc;
        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        /// </summary>
        public readonly int? LimitNoFile;
        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
        /// </summary>
        public readonly int? OomScoreAdjust;
        /// <summary>
        /// Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Remote.SystemdServiceRestart? Restart;
        /// <summary>
        /// Configures the time to sleep before restarting a service (as configured with Restart=).
        /// </summary>
        public readonly int? RestartSec;
        /// <summary>
        /// Configure unit start rate limiting. Units which are started more than burst times within an interval time span are not permitted to start any more. Use StartLimitIntervalSec= to configure the checking interval and StartLimitBurst= to configure how many starts per interval are allowed.
        /// </summary>
        public readonly int? StartLimitInterval;
        /// <summary>
        /// Configures the mechanism via which the service notifies the manager that the service start-up has finished.
        /// </summary>
        public readonly UnMango.KubernetesTheHardWay.Remote.SystemdServiceType? Type;

        [OutputConstructor]
        private SystemdServiceSection(
            UnMango.KubernetesTheHardWay.Remote.SystemdDelegate? @delegate,

            ImmutableArray<string> environment,

            string? execStart,

            string? execStartPre,

            UnMango.KubernetesTheHardWay.Remote.SystemdServiceExitType? exitType,

            UnMango.KubernetesTheHardWay.Remote.SystemdKillMode? killMode,

            string? limitCore,

            string? limitNProc,

            int? limitNoFile,

            int? oomScoreAdjust,

            UnMango.KubernetesTheHardWay.Remote.SystemdServiceRestart? restart,

            int? restartSec,

            int? startLimitInterval,

            UnMango.KubernetesTheHardWay.Remote.SystemdServiceType? type)
        {
            Delegate = @delegate;
            Environment = environment;
            ExecStart = execStart;
            ExecStartPre = execStartPre;
            ExitType = exitType;
            KillMode = killMode;
            LimitCore = limitCore;
            LimitNProc = limitNProc;
            LimitNoFile = limitNoFile;
            OomScoreAdjust = oomScoreAdjust;
            Restart = restart;
            RestartSec = restartSec;
            StartLimitInterval = startLimitInterval;
            Type = type;
        }
    }
}
