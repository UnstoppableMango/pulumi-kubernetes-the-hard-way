// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiKubernetes from "@pulumi/kubernetes";

/**
 * Gets the static pod manifests for KubeVip.
 */
export function getKubeVipManifest(args: GetKubeVipManifestArgs, opts?: pulumi.InvokeOptions): Promise<GetKubeVipManifestResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("kubernetes-the-hard-way:config:getKubeVipManifest", {
        "address": args.address,
        "bgpAs": args.bgpAs,
        "bgpEnable": args.bgpEnable,
        "bgpPeerAddress": args.bgpPeerAddress,
        "bgpPeerAs": args.bgpPeerAs,
        "bgpPeerPass": args.bgpPeerPass,
        "bgpPeers": args.bgpPeers,
        "bgpRouterId": args.bgpRouterId,
        "cpEnable": args.cpEnable,
        "cpNamespace": args.cpNamespace,
        "image": args.image,
        "kubeconfigPath": args.kubeconfigPath,
        "name": args.name,
        "namespace": args.namespace,
        "port": args.port,
        "svcEnable": args.svcEnable,
        "version": args.version,
        "vipArp": args.vipArp,
        "vipCidr": args.vipCidr,
        "vipDdns": args.vipDdns,
        "vipInterface": args.vipInterface,
        "vipLeaderElection": args.vipLeaderElection,
        "vipLeaseDuration": args.vipLeaseDuration,
        "vipRenewDeadline": args.vipRenewDeadline,
        "vipRetryPeriod": args.vipRetryPeriod,
    }, opts);
}

export interface GetKubeVipManifestArgs {
    /**
     * TODO
     */
    address: string;
    /**
     * TODO
     */
    bgpAs?: number;
    /**
     * TODO
     */
    bgpEnable?: boolean;
    /**
     * TODO
     */
    bgpPeerAddress?: string;
    /**
     * TODO
     */
    bgpPeerAs?: number;
    /**
     * TODO
     */
    bgpPeerPass?: string;
    /**
     * TODO
     */
    bgpPeers?: string;
    /**
     * TODO
     */
    bgpRouterId?: string;
    /**
     * TODO
     */
    cpEnable?: boolean;
    /**
     * TODO
     */
    cpNamespace?: string;
    /**
     * Override the kube-vip image.
     */
    image?: string;
    /**
     * Path to the kubeconfig on the remote host.
     */
    kubeconfigPath: string;
    /**
     * Name of the static pod. Defaults to kube-vip.
     */
    name?: string;
    /**
     * Namespace for the static pod. Defaults to kube-system.
     */
    namespace?: string;
    /**
     * TODO
     */
    port?: number;
    /**
     * TODO
     */
    svcEnable?: boolean;
    /**
     * Version of kube-vip to use.
     */
    version?: string;
    /**
     * TODO
     */
    vipArp?: boolean;
    /**
     * TODO
     */
    vipCidr: number;
    /**
     * TODO
     */
    vipDdns?: boolean;
    /**
     * TODO
     */
    vipInterface?: string;
    /**
     * TODO
     */
    vipLeaderElection?: boolean;
    /**
     * TODO
     */
    vipLeaseDuration?: number;
    /**
     * TODO
     */
    vipRenewDeadline?: number;
    /**
     * TODO
     */
    vipRetryPeriod?: number;
}

/**
 * Gets the static pod manifests for KubeVip.
 */
export interface GetKubeVipManifestResult {
    readonly result: outputs.config.PodManifest;
}
/**
 * Gets the static pod manifests for KubeVip.
 */
export function getKubeVipManifestOutput(args: GetKubeVipManifestOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubeVipManifestResult> {
    return pulumi.output(args).apply((a: any) => getKubeVipManifest(a, opts))
}

export interface GetKubeVipManifestOutputArgs {
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
     * Name of the static pod. Defaults to kube-vip.
     */
    name?: pulumi.Input<string>;
    /**
     * Namespace for the static pod. Defaults to kube-system.
     */
    namespace?: pulumi.Input<string>;
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
