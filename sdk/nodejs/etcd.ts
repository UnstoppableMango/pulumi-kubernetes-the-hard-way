// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

import * as pulumiCommand from "@pulumi/command";

import {Download} from "./remote";
import {Mkdir, Mv, Tar} from "./tools";

/**
 * Represents an etcd binary on a remote system.
 */
export class Etcd extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:index:Etcd';

    /**
     * Returns true if the given object is an instance of Etcd.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Etcd {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Etcd.__pulumiType;
    }

    /**
     * The etcd CPU architecture.
     */
    public readonly architecture!: pulumi.Output<enums.Architecture>;
    /**
     * The name of the etcd release archive.
     */
    public /*out*/ readonly archiveName!: pulumi.Output<string>;
    /**
     * The etcd download operation.
     */
    public /*out*/ readonly download!: pulumi.Output<Download>;
    /**
     * The directory where the etcd binary was downloaded to.
     */
    public readonly downloadDirectory!: pulumi.Output<string>;
    /**
     * The operation to create the download directory.
     */
    public /*out*/ readonly downloadMkdir!: pulumi.Output<Mkdir>;
    /**
     * The path to the etcd binary on the remote system.
     */
    public /*out*/ readonly etcdPath!: pulumi.Output<string>;
    /**
     * The path to the etcdctl binary on the remote system.
     */
    public /*out*/ readonly etcdctlPath!: pulumi.Output<string>;
    /**
     * Directory to install the `etcd` and `etcdctl` binaries.
     */
    public readonly installDirectory!: pulumi.Output<string>;
    /**
     * The operation to create the install directory.
     */
    public /*out*/ readonly installMkdir!: pulumi.Output<Mkdir>;
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
     * Create a Etcd resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EtcdArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["downloadDirectory"] = args ? args.downloadDirectory : undefined;
            resourceInputs["installDirectory"] = (args ? args.installDirectory : undefined) ?? "/usr/local/bin";
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["archiveName"] = undefined /*out*/;
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["downloadMkdir"] = undefined /*out*/;
            resourceInputs["etcdPath"] = undefined /*out*/;
            resourceInputs["etcdctlPath"] = undefined /*out*/;
            resourceInputs["installMkdir"] = undefined /*out*/;
            resourceInputs["mvEtcd"] = undefined /*out*/;
            resourceInputs["mvEtcdctl"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tar"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["archiveName"] = undefined /*out*/;
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["downloadDirectory"] = undefined /*out*/;
            resourceInputs["downloadMkdir"] = undefined /*out*/;
            resourceInputs["etcdPath"] = undefined /*out*/;
            resourceInputs["etcdctlPath"] = undefined /*out*/;
            resourceInputs["installDirectory"] = undefined /*out*/;
            resourceInputs["installMkdir"] = undefined /*out*/;
            resourceInputs["mvEtcd"] = undefined /*out*/;
            resourceInputs["mvEtcdctl"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tar"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Etcd.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Etcd resource.
 */
export interface EtcdArgs {
    /**
     * The etcd CPU architecture.
     */
    architecture?: pulumi.Input<enums.Architecture>;
    /**
     * The connection details.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Temporary directory to download files to. Defaults to `/tmp/<random string>`.
     */
    downloadDirectory?: pulumi.Input<string>;
    /**
     * Directory to install the `etcd` and `etcdctl` binaries.
     */
    installDirectory?: pulumi.Input<string>;
    /**
     * The version of etcd to install.
     */
    version?: pulumi.Input<string>;
}
