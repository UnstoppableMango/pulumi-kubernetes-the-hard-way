// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Config
{
    public static class GetCniBridgePluginConfiguration
    {
        /// <summary>
        /// Get the `bridge` configuration.
        /// </summary>
        public static Task<GetCniBridgePluginConfigurationResult> InvokeAsync(GetCniBridgePluginConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCniBridgePluginConfigurationResult>("kubernetes-the-hard-way:config:getCniBridgePluginConfiguration", args ?? new GetCniBridgePluginConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Get the `bridge` configuration.
        /// </summary>
        public static Output<GetCniBridgePluginConfigurationResult> Invoke(GetCniBridgePluginConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCniBridgePluginConfigurationResult>("kubernetes-the-hard-way:config:getCniBridgePluginConfiguration", args ?? new GetCniBridgePluginConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCniBridgePluginConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bridge name.
        /// </summary>
        [Input("bridge")]
        public string? Bridge { get; set; }

        /// <summary>
        /// CNI version.
        /// </summary>
        [Input("cniVersion")]
        public string? CniVersion { get; set; }

        /// <summary>
        /// IP masq.
        /// </summary>
        [Input("ipMasq")]
        public bool? IpMasq { get; set; }

        /// <summary>
        /// IPAM
        /// </summary>
        [Input("ipam")]
        public Inputs.CniBridgeIpam? Ipam { get; set; }

        /// <summary>
        /// Is gateway.
        /// </summary>
        [Input("isGateway")]
        public bool? IsGateway { get; set; }

        /// <summary>
        /// CNI plugin name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The subnet to use.
        /// </summary>
        [Input("subnet", required: true)]
        public string Subnet { get; set; } = null!;

        /// <summary>
        /// CNI plugin type.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetCniBridgePluginConfigurationArgs()
        {
        }
        public static new GetCniBridgePluginConfigurationArgs Empty => new GetCniBridgePluginConfigurationArgs();
    }

    public sealed class GetCniBridgePluginConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bridge name.
        /// </summary>
        [Input("bridge")]
        public Input<string>? Bridge { get; set; }

        /// <summary>
        /// CNI version.
        /// </summary>
        [Input("cniVersion")]
        public Input<string>? CniVersion { get; set; }

        /// <summary>
        /// IP masq.
        /// </summary>
        [Input("ipMasq")]
        public Input<bool>? IpMasq { get; set; }

        /// <summary>
        /// IPAM
        /// </summary>
        [Input("ipam")]
        public Input<Inputs.CniBridgeIpamArgs>? Ipam { get; set; }

        /// <summary>
        /// Is gateway.
        /// </summary>
        [Input("isGateway")]
        public Input<bool>? IsGateway { get; set; }

        /// <summary>
        /// CNI plugin name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The subnet to use.
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        /// <summary>
        /// CNI plugin type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetCniBridgePluginConfigurationInvokeArgs()
        {
        }
        public static new GetCniBridgePluginConfigurationInvokeArgs Empty => new GetCniBridgePluginConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetCniBridgePluginConfigurationResult
    {
        public readonly Outputs.CniBridgePluginConfiguration Result;

        [OutputConstructor]
        private GetCniBridgePluginConfigurationResult(Outputs.CniBridgePluginConfiguration result)
        {
            Result = result;
        }
    }
}
