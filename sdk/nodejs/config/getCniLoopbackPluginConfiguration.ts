// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {CniLoopbackPluginConfiguration} from "./index";

/**
 * Get the `loopback` configuration.
 */
export function getCniLoopbackPluginConfiguration(args?: GetCniLoopbackPluginConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetCniLoopbackPluginConfigurationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("kubernetes-the-hard-way:config:getCniLoopbackPluginConfiguration", {
        "cniVersion": args.cniVersion,
        "name": args.name,
        "path": args.path,
        "type": args.type,
    }, opts);
}

export interface GetCniLoopbackPluginConfigurationArgs {
    /**
     * CNI version.
     */
    cniVersion?: string;
    /**
     * CNI plugin name.
     */
    name?: string;
    /**
     * Path to put the configuration file on the remote system
     */
    path?: string;
    /**
     * CNI plugin type.
     */
    type?: string;
}

/**
 * Get the `loopback` configuration.
 */
export interface GetCniLoopbackPluginConfigurationResult {
    readonly result: CniLoopbackPluginConfiguration;
}
/**
 * Get the `loopback` configuration.
 */
export function getCniLoopbackPluginConfigurationOutput(args?: GetCniLoopbackPluginConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCniLoopbackPluginConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getCniLoopbackPluginConfiguration(a, opts))
}

export interface GetCniLoopbackPluginConfigurationOutputArgs {
    /**
     * CNI version.
     */
    cniVersion?: pulumi.Input<string>;
    /**
     * CNI plugin name.
     */
    name?: pulumi.Input<string>;
    /**
     * Path to put the configuration file on the remote system
     */
    path?: pulumi.Input<string>;
    /**
     * CNI plugin type.
     */
    type?: pulumi.Input<string>;
}
