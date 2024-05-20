// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get the containerd configuration.
 */
export function getContainerdConfiguration(args?: GetContainerdConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerdConfigurationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("kubernetes-the-hard-way:config:getContainerdConfiguration", {
        "cri": args.cri,
    }, opts);
}

export interface GetContainerdConfigurationArgs {
    /**
     * The cri configuration.
     */
    cri?: inputs.config.ContainerdCriPluginConfiguration;
}

/**
 * Get the containerd configuration.
 */
export interface GetContainerdConfigurationResult {
    readonly result: outputs.config.ContainerdConfiguration;
}
/**
 * Get the containerd configuration.
 */
export function getContainerdConfigurationOutput(args?: GetContainerdConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContainerdConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getContainerdConfiguration(a, opts))
}

export interface GetContainerdConfigurationOutputArgs {
    /**
     * The cri configuration.
     */
    cri?: inputs.config.ContainerdCriPluginConfigurationArgs;
}