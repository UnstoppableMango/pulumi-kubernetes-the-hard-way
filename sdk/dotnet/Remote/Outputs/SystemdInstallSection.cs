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
    /// https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options
    /// </summary>
    [OutputType]
    public sealed class SystemdInstallSection
    {
        /// <summary>
        /// A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
        /// </summary>
        public readonly ImmutableArray<string> WantedBy;

        [OutputConstructor]
        private SystemdInstallSection(ImmutableArray<string> wantedBy)
        {
            WantedBy = wantedBy;
        }
    }
}
