// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

import {File} from "./index";

/**
 * The CNI loopback plugin configuration.
 */
export class CniLoopbackPluginConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing CniLoopbackPluginConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CniLoopbackPluginConfiguration {
        return new CniLoopbackPluginConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:CniLoopbackPluginConfiguration';

    /**
     * Returns true if the given object is an instance of CniLoopbackPluginConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CniLoopbackPluginConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CniLoopbackPluginConfiguration.__pulumiType;
    }

    /**
     * CNI version.
     */
    public readonly cniVersion!: pulumi.Output<string>;
    /**
     * The parameters with which to connect to the remote host.
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * The file on the remote system.
     */
    public /*out*/ readonly file!: pulumi.Output<File | undefined>;
    /**
     * CNI plugin name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Path to put the configuration file on the remote system
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * CNI plugin type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a CniLoopbackPluginConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CniLoopbackPluginConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["cniVersion"] = args ? args.cniVersion : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["file"] = undefined /*out*/;
        } else {
            resourceInputs["cniVersion"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["file"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CniLoopbackPluginConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CniLoopbackPluginConfiguration resource.
 */
export interface CniLoopbackPluginConfigurationArgs {
    /**
     * CNI version.
     */
    cniVersion?: pulumi.Input<string>;
    /**
     * The parameters with which to connect to the remote host.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
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