// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Remote.Inputs
{

    /// <summary>
    /// https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
    /// </summary>
    public sealed class SystemdServiceSectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Turns on delegation of further resource control partitioning to processes of the unit.
        /// </summary>
        [Input("delegate")]
        public Input<UnMango.KubernetesTheHardWay.Remote.SystemdDelegate>? Delegate { get; set; }

        [Input("environment")]
        private InputList<string>? _environment;
        public InputList<string> Environment
        {
            get => _environment ?? (_environment = new InputList<string>());
            set => _environment = value;
        }

        /// <summary>
        /// Commands that are executed when this service is started.
        /// </summary>
        [Input("execStart")]
        public Input<string>? ExecStart { get; set; }

        /// <summary>
        /// Additional commands that are executed before the command in ExecStart=.
        /// </summary>
        [Input("execStartPre")]
        public Input<string>? ExecStartPre { get; set; }

        /// <summary>
        /// Specifies when the manager should consider the service to be finished.
        /// </summary>
        [Input("exitType")]
        public Input<UnMango.KubernetesTheHardWay.Remote.SystemdServiceExitType>? ExitType { get; set; }

        /// <summary>
        /// Specifies how processes of this unit shall be killed.
        /// </summary>
        [Input("killMode")]
        public Input<UnMango.KubernetesTheHardWay.Remote.SystemdKillMode>? KillMode { get; set; }

        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        /// </summary>
        [Input("limitCore")]
        public Input<string>? LimitCore { get; set; }

        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        /// </summary>
        [Input("limitNProc")]
        public Input<string>? LimitNProc { get; set; }

        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        /// </summary>
        [Input("limitNoFile")]
        public Input<int>? LimitNoFile { get; set; }

        /// <summary>
        /// https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
        /// </summary>
        [Input("oomScoreAdjust")]
        public Input<int>? OomScoreAdjust { get; set; }

        /// <summary>
        /// Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
        /// </summary>
        [Input("restart")]
        public Input<UnMango.KubernetesTheHardWay.Remote.SystemdServiceRestart>? Restart { get; set; }

        /// <summary>
        /// Configures the time to sleep before restarting a service (as configured with Restart=).
        /// </summary>
        [Input("restartSec")]
        public Input<string>? RestartSec { get; set; }

        /// <summary>
        /// Configures the mechanism via which the service notifies the manager that the service start-up has finished.
        /// </summary>
        [Input("type")]
        public Input<UnMango.KubernetesTheHardWay.Remote.SystemdServiceType>? Type { get; set; }

        public SystemdServiceSectionArgs()
        {
        }
        public static new SystemdServiceSectionArgs Empty => new SystemdServiceSectionArgs();
    }
}
