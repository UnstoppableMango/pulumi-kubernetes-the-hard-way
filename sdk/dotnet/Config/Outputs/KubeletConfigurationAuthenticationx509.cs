// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config.Outputs
{

    [OutputType]
    public sealed class KubeletConfigurationAuthenticationx509
    {
        /// <summary>
        /// TODO
        /// </summary>
        public readonly string ClientCAFile;

        [OutputConstructor]
        private KubeletConfigurationAuthenticationx509(string clientCAFile)
        {
            ClientCAFile = clientCAFile;
        }
    }
}