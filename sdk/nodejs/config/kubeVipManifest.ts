// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiKubernetes from "@pulumi/kubernetes";

/**
 * Pseudo resource for generating the kube-vip manifest.
 */
export class KubeVipManifest extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:config:KubeVipManifest';

    /**
     * Returns true if the given object is an instance of KubeVipManifest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubeVipManifest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubeVipManifest.__pulumiType;
    }

    public /*out*/ readonly result!: pulumi.Output<outputs.config.PodManifest>;

    /**
     * Create a KubeVipManifest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubeVipManifestArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.kubeconfigPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubeconfigPath'");
            }
            if ((!args || args.vipCidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vipCidr'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["bgpAs"] = args ? args.bgpAs : undefined;
            resourceInputs["bgpEnable"] = args ? args.bgpEnable : undefined;
            resourceInputs["bgpPeerAddress"] = args ? args.bgpPeerAddress : undefined;
            resourceInputs["bgpPeerAs"] = args ? args.bgpPeerAs : undefined;
            resourceInputs["bgpPeerPass"] = args ? args.bgpPeerPass : undefined;
            resourceInputs["bgpPeers"] = args ? args.bgpPeers : undefined;
            resourceInputs["bgpRouterId"] = args ? args.bgpRouterId : undefined;
            resourceInputs["cpEnable"] = args ? args.cpEnable : undefined;
            resourceInputs["cpNamespace"] = args ? args.cpNamespace : undefined;
            resourceInputs["image"] = args ? args.image : undefined;
            resourceInputs["kubeconfigPath"] = args ? args.kubeconfigPath : undefined;
            resourceInputs["port"] = (args ? args.port : undefined) ?? 6443;
            resourceInputs["svcEnable"] = args ? args.svcEnable : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["vipArp"] = args ? args.vipArp : undefined;
            resourceInputs["vipCidr"] = args ? args.vipCidr : undefined;
            resourceInputs["vipDdns"] = args ? args.vipDdns : undefined;
            resourceInputs["vipInterface"] = args ? args.vipInterface : undefined;
            resourceInputs["vipLeaderElection"] = args ? args.vipLeaderElection : undefined;
            resourceInputs["vipLeaseDuration"] = args ? args.vipLeaseDuration : undefined;
            resourceInputs["vipRenewDeadline"] = args ? args.vipRenewDeadline : undefined;
            resourceInputs["vipRetryPeriod"] = args ? args.vipRetryPeriod : undefined;
            resourceInputs["result"] = undefined /*out*/;
        } else {
            resourceInputs["result"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KubeVipManifest.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a KubeVipManifest resource.
 */
export interface KubeVipManifestArgs {
    /**
     * TODO
     */
    address: pulumi.Input<string>;
    /**
     * TODO
     */
    bgpAs?: pulumi.Input<number>;
    /**
     * TODO
     */
    bgpEnable?: pulumi.Input<boolean>;
    /**
     * TODO
     */
    bgpPeerAddress?: pulumi.Input<string>;
    /**
     * TODO
     */
    bgpPeerAs?: pulumi.Input<number>;
    /**
     * TODO
     */
    bgpPeerPass?: pulumi.Input<string>;
    /**
     * TODO
     */
    bgpPeers?: pulumi.Input<string>;
    /**
     * TODO
     */
    bgpRouterId?: pulumi.Input<string>;
    /**
     * TODO
     */
    cpEnable?: pulumi.Input<boolean>;
    /**
     * TODO
     */
    cpNamespace?: pulumi.Input<string>;
    /**
     * Override the kube-vip image.
     */
    image?: pulumi.Input<string>;
    /**
     * Path to the kubeconfig on the remote host.
     */
    kubeconfigPath: pulumi.Input<string>;
    /**
     * TODO
     */
    port?: pulumi.Input<number>;
    /**
     * TODO
     */
    svcEnable?: pulumi.Input<boolean>;
    /**
     * Version of kube-vip to use.
     */
    version?: pulumi.Input<string>;
    /**
     * TODO
     */
    vipArp?: pulumi.Input<boolean>;
    /**
     * TODO
     */
    vipCidr: pulumi.Input<number>;
    /**
     * TODO
     */
    vipDdns?: pulumi.Input<boolean>;
    /**
     * TODO
     */
    vipInterface?: pulumi.Input<string>;
    /**
     * TODO
     */
    vipLeaderElection?: pulumi.Input<boolean>;
    /**
     * TODO
     */
    vipLeaseDuration?: pulumi.Input<number>;
    /**
     * TODO
     */
    vipRenewDeadline?: pulumi.Input<number>;
    /**
     * TODO
     */
    vipRetryPeriod?: pulumi.Input<number>;
}
