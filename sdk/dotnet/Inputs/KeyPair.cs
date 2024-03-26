// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.KubernetesTheHardWay.Inputs
{

    /// <summary>
    /// A certificate and key pair.
    /// </summary>
    public sealed class KeyPair : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PEM encoded certificate data
        /// </summary>
        [Input("certPem")]
        public string? CertPem { get; set; }

        /// <summary>
        /// The private key.
        /// </summary>
        [Input("key")]
        public Pulumi.Tls.PrivateKey? Key { get; set; }

        /// <summary>
        /// The PEM encoded private key data.
        /// </summary>
        [Input("privateKeyPem")]
        public string? PrivateKeyPem { get; set; }

        /// <summary>
        /// The PEM encoded public key data.
        /// </summary>
        [Input("publicKeyPem")]
        public string? PublicKeyPem { get; set; }

        public KeyPair()
        {
        }
        public static new KeyPair Empty => new KeyPair();
    }
}
