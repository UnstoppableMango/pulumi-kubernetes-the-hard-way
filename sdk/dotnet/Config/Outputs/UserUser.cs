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
    public sealed class UserUser
    {
        public readonly string ClientCertificateData;
        public readonly string ClientKeyData;

        [OutputConstructor]
        private UserUser(
            string clientCertificateData,

            string clientKeyData)
        {
            ClientCertificateData = clientCertificateData;
            ClientKeyData = clientKeyData;
        }
    }
}
