// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get the `loopback` configuration.
 */
export class CniLoopbackPluginConfiguration extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration';

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

    public /*out*/ readonly result!: pulumi.Output<outputs.config.CniLoopbackPluginConfiguration>;
    /**
     * The yaml representation of the manifest.
     */
    public /*out*/ readonly yaml!: pulumi.Output<string>;

    /**
     * Create a CniLoopbackPluginConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CniLoopbackPluginConfigurationArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["cniVersion"] = args ? args.cniVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["yaml"] = undefined /*out*/;
        } else {
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["yaml"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CniLoopbackPluginConfiguration.__pulumiType, name, resourceInputs, opts, true /*remote*/);
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
