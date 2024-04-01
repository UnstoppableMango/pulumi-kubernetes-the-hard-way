// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import {Certificate, RootCa} from "./index";

export class ClusterPki extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:tls:ClusterPki';

    /**
     * Returns true if the given object is an instance of ClusterPki.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterPki {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterPki.__pulumiType;
    }

    /**
     * The admin certificate.
     */
    public /*out*/ readonly admin!: pulumi.Output<Certificate>;
    /**
     * Name of the algorithm to use when generating the private key.
     */
    public readonly algorithm!: pulumi.Output<enums.tls.Algorithm>;
    public /*out*/ readonly ca!: pulumi.Output<RootCa>;
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * The controller manager certificate.
     */
    public /*out*/ readonly controllerManager!: pulumi.Output<Certificate>;
    /**
     * The kube proxy certificate.
     */
    public /*out*/ readonly kubeProxy!: pulumi.Output<Certificate>;
    /**
     * The kube scheduler certificate.
     */
    public /*out*/ readonly kubeScheduler!: pulumi.Output<Certificate>;
    /**
     * Map of node name to kubelet certificate.
     */
    public /*out*/ readonly kubelet!: pulumi.Output<{[key: string]: Certificate}>;
    /**
     * The kubernetes certificate.
     */
    public /*out*/ readonly kubernetes!: pulumi.Output<Certificate>;
    /**
     * The publicly accessible IP for the cluster.
     */
    public readonly publicIp!: pulumi.Output<string>;
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     */
    public readonly rsaBits!: pulumi.Output<number>;
    /**
     * The service accounts certificate.
     */
    public /*out*/ readonly serviceAccounts!: pulumi.Output<Certificate>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid.
     */
    public readonly validityPeriodHours!: pulumi.Output<number>;

    /**
     * Create a ClusterPki resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterPkiArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.nodes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodes'");
            }
            if ((!args || args.publicIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicIp'");
            }
            resourceInputs["algorithm"] = (args ? args.algorithm : undefined) ?? "RSA";
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["ecdsaCurve"] = args ? args.ecdsaCurve : undefined;
            resourceInputs["nodes"] = args ? args.nodes : undefined;
            resourceInputs["publicIp"] = args ? args.publicIp : undefined;
            resourceInputs["rsaBits"] = (args ? args.rsaBits : undefined) ?? 2048;
            resourceInputs["validityPeriodHours"] = (args ? args.validityPeriodHours : undefined) ?? 8076;
            resourceInputs["admin"] = undefined /*out*/;
            resourceInputs["ca"] = undefined /*out*/;
            resourceInputs["controllerManager"] = undefined /*out*/;
            resourceInputs["kubeProxy"] = undefined /*out*/;
            resourceInputs["kubeScheduler"] = undefined /*out*/;
            resourceInputs["kubelet"] = undefined /*out*/;
            resourceInputs["kubernetes"] = undefined /*out*/;
            resourceInputs["serviceAccounts"] = undefined /*out*/;
        } else {
            resourceInputs["admin"] = undefined /*out*/;
            resourceInputs["algorithm"] = undefined /*out*/;
            resourceInputs["ca"] = undefined /*out*/;
            resourceInputs["clusterName"] = undefined /*out*/;
            resourceInputs["controllerManager"] = undefined /*out*/;
            resourceInputs["kubeProxy"] = undefined /*out*/;
            resourceInputs["kubeScheduler"] = undefined /*out*/;
            resourceInputs["kubelet"] = undefined /*out*/;
            resourceInputs["kubernetes"] = undefined /*out*/;
            resourceInputs["publicIp"] = undefined /*out*/;
            resourceInputs["rsaBits"] = undefined /*out*/;
            resourceInputs["serviceAccounts"] = undefined /*out*/;
            resourceInputs["validityPeriodHours"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterPki.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }

    getKubeconfig(args: ClusterPki.GetKubeconfigArgs): pulumi.Output<outputs.config.Kubeconfig> {
        const result: pulumi.Output<ClusterPki.GetKubeconfigResult> = pulumi.runtime.call("kubernetes-the-hard-way:tls:ClusterPki/getKubeconfig", {
            "__self__": this,
            "options": args.options,
        }, this);
        return result.result;
    }
}

/**
 * The set of arguments for constructing a ClusterPki resource.
 */
export interface ClusterPkiArgs {
    /**
     * Name of the algorithm to use when generating the private key.
     */
    algorithm?: pulumi.Input<enums.tls.Algorithm>;
    /**
     * A name to use for the cluster
     */
    clusterName: pulumi.Input<string>;
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
     */
    ecdsaCurve?: pulumi.Input<enums.tls.EcdsaCurve>;
    /**
     * Map of node names to node configuration.
     */
    nodes: pulumi.Input<{[key: string]: pulumi.Input<inputs.tls.ClusterPkiNodeArgs>}>;
    /**
     * Publicly accessible IP address.
     */
    publicIp: pulumi.Input<string>;
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
     */
    rsaBits?: pulumi.Input<number>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid.
     */
    validityPeriodHours?: pulumi.Input<number>;
}

export namespace ClusterPki {
    /**
     * The set of arguments for the ClusterPki.getKubeconfig method.
     */
    export interface GetKubeconfigArgs {
        options: pulumi.Input<inputs.config.KubeconfigAdminOptionsArgs> | pulumi.Input<inputs.config.KubeconfigKubeControllerManagerOptionsArgs> | pulumi.Input<inputs.config.KubeconfigKubeProxyOptionsArgs> | pulumi.Input<inputs.config.KubeconfigKubeSchedulerOptionsArgs> | pulumi.Input<inputs.config.KubeconfigWorkerOptionsArgs>;
    }

    /**
     * The results of the ClusterPki.getKubeconfig method.
     */
    export interface GetKubeconfigResult {
        readonly result: outputs.config.Kubeconfig;
    }

}
