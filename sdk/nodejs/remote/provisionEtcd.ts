// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

import {EtcdConfiguration, EtcdInstall, StartEtcd, SystemdService} from "./index";

/**
 * Starts etcd on a remote system.
 */
export class ProvisionEtcd extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:ProvisionEtcd';

    /**
     * Returns true if the given object is an instance of ProvisionEtcd.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProvisionEtcd {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProvisionEtcd.__pulumiType;
    }

    /**
     * TODO
     */
    public readonly architecture!: pulumi.Output<enums.remote.Architecture | undefined>;
    /**
     * TODO
     */
    public readonly binaryDirectory!: pulumi.Output<string | undefined>;
    /**
     * The TLS bundle.
     */
    public readonly bundle!: pulumi.Output<outputs.tls.Bundle>;
    /**
     * Etcd configuration.
     */
    public /*out*/ readonly configuration!: pulumi.Output<EtcdConfiguration>;
    /**
     * The directory to use for etcd configuration.
     */
    public readonly configurationDirectory!: pulumi.Output<string | undefined>;
    /**
     * The parameters with which to connect to the remote host.
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * The directory to use for etcd data.
     */
    public readonly dataDirectory!: pulumi.Output<string | undefined>;
    /**
     * Install etcd.
     */
    public /*out*/ readonly install!: pulumi.Output<EtcdInstall>;
    /**
     * The internal IP of the etcd node
     */
    public readonly internalIp!: pulumi.Output<string>;
    /**
     * Systemd service.
     */
    public /*out*/ readonly service!: pulumi.Output<SystemdService>;
    /**
     * Start etcd
     */
    public /*out*/ readonly start!: pulumi.Output<StartEtcd>;
    /**
     * The version to install.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a ProvisionEtcd resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProvisionEtcdArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bundle === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bundle'");
            }
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.internalIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internalIp'");
            }
            resourceInputs["architecture"] = args ? args.architecture : undefined;
            resourceInputs["binaryDirectory"] = args ? args.binaryDirectory : undefined;
            resourceInputs["bundle"] = args ? args.bundle : undefined;
            resourceInputs["configurationDirectory"] = args ? args.configurationDirectory : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["dataDirectory"] = args ? args.dataDirectory : undefined;
            resourceInputs["internalIp"] = args ? args.internalIp : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["install"] = undefined /*out*/;
            resourceInputs["service"] = undefined /*out*/;
            resourceInputs["start"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["binaryDirectory"] = undefined /*out*/;
            resourceInputs["bundle"] = undefined /*out*/;
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["configurationDirectory"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["dataDirectory"] = undefined /*out*/;
            resourceInputs["install"] = undefined /*out*/;
            resourceInputs["internalIp"] = undefined /*out*/;
            resourceInputs["service"] = undefined /*out*/;
            resourceInputs["start"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProvisionEtcd.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a ProvisionEtcd resource.
 */
export interface ProvisionEtcdArgs {
    /**
     * TODO
     */
    architecture?: pulumi.Input<enums.remote.Architecture>;
    /**
     * TODO
     */
    binaryDirectory?: pulumi.Input<string>;
    /**
     * The TLS bundle.
     */
    bundle: pulumi.Input<inputs.tls.BundleArgs>;
    /**
     * The directory to use for etcd configuration.
     */
    configurationDirectory?: pulumi.Input<string>;
    /**
     * The parameters with which to connect to the remote host.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * The directory to use for etcd data.
     */
    dataDirectory?: pulumi.Input<string>;
    /**
     * The internal IP of the etcd node
     */
    internalIp: pulumi.Input<string>;
    /**
     * The version to install.
     */
    version?: pulumi.Input<string>;
}
