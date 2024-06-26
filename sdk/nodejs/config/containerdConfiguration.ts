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
export class ContainerdConfiguration extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:config:ContainerdConfiguration';

    /**
     * Returns true if the given object is an instance of ContainerdConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerdConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerdConfiguration.__pulumiType;
    }

    public /*out*/ readonly result!: pulumi.Output<outputs.config.ContainerdConfiguration>;
    /**
     * The toml representation of the configuration.
     */
    public /*out*/ readonly toml!: pulumi.Output<string>;

    /**
     * Create a ContainerdConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerdConfigurationArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["cri"] = args ? args.cri : undefined;
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["toml"] = undefined /*out*/;
        } else {
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["toml"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerdConfiguration.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a ContainerdConfiguration resource.
 */
export interface ContainerdConfigurationArgs {
    /**
     * The cri configuration.
     */
    cri?: inputs.config.ContainerdCriPluginConfigurationArgs;
}
