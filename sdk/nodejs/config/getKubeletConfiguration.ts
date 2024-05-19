// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get the kubelet configuration.
 */
export function getKubeletConfiguration(args: GetKubeletConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetKubeletConfigurationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("kubernetes-the-hard-way:config:getKubeletConfiguration", {
        "anonymous": args.anonymous,
        "authorizationMode": args.authorizationMode,
        "cgroupDriver": args.cgroupDriver,
        "clientCAFile": args.clientCAFile,
        "clusterDNS": args.clusterDNS,
        "clusterDomain": args.clusterDomain,
        "containerRuntimeEndpoint": args.containerRuntimeEndpoint,
        "podCIDR": args.podCIDR,
        "resolvConf": args.resolvConf,
        "runtimeRequestTimeout": args.runtimeRequestTimeout,
        "tlsCertFile": args.tlsCertFile,
        "tlsPrivateKeyFile": args.tlsPrivateKeyFile,
        "webhook": args.webhook,
    }, opts);
}

export interface GetKubeletConfigurationArgs {
    anonymous?: boolean;
    authorizationMode?: string;
    cgroupDriver?: string;
    clientCAFile?: string;
    clusterDNS?: string[];
    clusterDomain?: string;
    containerRuntimeEndpoint?: string;
    podCIDR: string;
    resolvConf?: string;
    runtimeRequestTimeout?: string;
    tlsCertFile?: string;
    tlsPrivateKeyFile?: string;
    webhook?: boolean;
}

export interface GetKubeletConfigurationResult {
    readonly result: outputs.config.KubeletConfiguration;
}
/**
 * Get the kubelet configuration.
 */
export function getKubeletConfigurationOutput(args: GetKubeletConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubeletConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getKubeletConfiguration(a, opts))
}

export interface GetKubeletConfigurationOutputArgs {
    anonymous?: pulumi.Input<boolean>;
    authorizationMode?: pulumi.Input<string>;
    cgroupDriver?: pulumi.Input<string>;
    clientCAFile?: pulumi.Input<string>;
    clusterDNS?: pulumi.Input<pulumi.Input<string>[]>;
    clusterDomain?: pulumi.Input<string>;
    containerRuntimeEndpoint?: pulumi.Input<string>;
    podCIDR: pulumi.Input<string>;
    resolvConf?: pulumi.Input<string>;
    runtimeRequestTimeout?: pulumi.Input<string>;
    tlsCertFile?: pulumi.Input<string>;
    tlsPrivateKeyFile?: pulumi.Input<string>;
    webhook?: pulumi.Input<boolean>;
}