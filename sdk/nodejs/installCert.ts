// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

import * as pulumiTls from "@pulumi/tls";

import {File} from "remote";

/**
 * Creates a RemoteFile resource representing the copy operation.
 */
export function installCert(args: InstallCertArgs, opts?: pulumi.InvokeOptions): Promise<InstallCertResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("kubernetes-the-hard-way:index:installCert", {
        "connection": args.connection ? inputs.connectionProvideDefaults(args.connection) : undefined,
        "keypair": args.keypair,
        "name": args.name,
        "options": args.options,
        "path": args.path,
    }, opts);
}

export interface InstallCertArgs {
    /**
     * The connection details.
     */
    connection: inputs.Connection;
    /**
     * The certificate to install at the remote location.
     */
    keypair: inputs.KeyPair;
    name: string;
    options?: inputs.ResourceOptions;
    /**
     * The path to install to.
     */
    path?: string;
}

export interface InstallCertResult {
    /**
     * A resource representing the the file on the remote machine.
     */
    readonly result: File;
}
/**
 * Creates a RemoteFile resource representing the copy operation.
 */
export function installCertOutput(args: InstallCertOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<InstallCertResult> {
    return pulumi.output(args).apply((a: any) => installCert(a, opts))
}

export interface InstallCertOutputArgs {
    /**
     * The connection details.
     */
    connection: pulumi.Input<inputs.ConnectionArgs>;
    /**
     * The certificate to install at the remote location.
     */
    keypair: pulumi.Input<inputs.KeyPairArgs>;
    name: string;
    options?: inputs.ResourceOptionsArgs;
    /**
     * The path to install to.
     */
    path?: pulumi.Input<string>;
}
