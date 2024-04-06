// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

import {Etcdctl, Mkdir, Mv, Tar} from "../tools";
import {Download} from "./index";

/**
 * Represents an etcd binary on a remote system.
 */
export class EtcdInstall extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:EtcdInstall';

    /**
     * Returns true if the given object is an instance of EtcdInstall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EtcdInstall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EtcdInstall.__pulumiType;
    }

    /**
     * The etcd CPU architecture.
     */
    public readonly architecture!: pulumi.Output<enums.remote.Architecture>;
    /**
     * The name of the etcd release archive.
     */
    public /*out*/ readonly archiveName!: pulumi.Output<string>;
    /**
     * Directory to install the `etcd` and `etcdctl` binaries.
     */
    public readonly directory!: pulumi.Output<string>;
    /**
     * The etcd download operation.
     */
    public /*out*/ readonly download!: pulumi.Output<Download>;
    /**
     * The path to the etcd binary on the remote system.
     */
    public /*out*/ readonly etcdPath!: pulumi.Output<string>;
    /**
     * The path to the etcdctl binary on the remote system.
     */
    public /*out*/ readonly etcdctlPath!: pulumi.Output<string>;
    /**
     * The operation to create the install directory.
     */
    public /*out*/ readonly mkdir!: pulumi.Output<Mkdir>;
    /**
     * The operation to move the etcd binary to the install directory.
     */
    public /*out*/ readonly mvEtcd!: pulumi.Output<Mv>;
    /**
     * The operation to move the etcdctl binary to the install directory.
     */
    public /*out*/ readonly mvEtcdctl!: pulumi.Output<Mv>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The tar operation.
     */
    public /*out*/ readonly tar!: pulumi.Output<Tar>;
    /**
     * The url used to download etcd.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * The version of etcd downloaded.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a EtcdInstall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EtcdInstallArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["directory"] = (args ? args.directory : undefined) ?? "/usr/local/bin";
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["archiveName"] = undefined /*out*/;
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["etcdPath"] = undefined /*out*/;
            resourceInputs["etcdctlPath"] = undefined /*out*/;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["mvEtcd"] = undefined /*out*/;
            resourceInputs["mvEtcdctl"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tar"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["archiveName"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["etcdPath"] = undefined /*out*/;
            resourceInputs["etcdctlPath"] = undefined /*out*/;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["mvEtcd"] = undefined /*out*/;
            resourceInputs["mvEtcdctl"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tar"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EtcdInstall.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }

    etcdctl(): pulumi.Output<Etcdctl> {
        const result: pulumi.Output<EtcdInstall.EtcdctlResult> = pulumi.runtime.call("kubernetes-the-hard-way:remote:EtcdInstall/etcdctl", {
            "__self__": this,
        }, this);
        return result.result;
    }
}

/**
 * The set of arguments for constructing a EtcdInstall resource.
 */
export interface EtcdInstallArgs {
    /**
     * The etcd CPU architecture.
     */
    architecture?: pulumi.Input<enums.remote.Architecture>;
    /**
     * The connection details.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Directory to install the `etcd` and `etcdctl` binaries.
     */
    directory?: pulumi.Input<string>;
    /**
     * The version of etcd to install.
     */
    version?: pulumi.Input<string>;
}

export namespace EtcdInstall {
    /**
     * The results of the EtcdInstall.etcdctl method.
     */
    export interface EtcdctlResult {
        readonly result: Etcdctl;
    }

}
