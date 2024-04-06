// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

import {Mkdir, Mktemp, Mv, Rm} from "../tools";
import {Download} from "./index";

/**
 * Installs runc on a remote system.
 */
export class RuncInstall extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:RuncInstall';

    /**
     * Returns true if the given object is an instance of RuncInstall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuncInstall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuncInstall.__pulumiType;
    }

    /**
     * The CPU architecture.
     */
    public readonly architecture!: pulumi.Output<enums.remote.Architecture>;
    /**
     * The connection details.
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * Directory to install the binary.
     */
    public readonly directory!: pulumi.Output<string>;
    public /*out*/ readonly download!: pulumi.Output<Download>;
    public /*out*/ readonly mkdir!: pulumi.Output<Mkdir | undefined>;
    public /*out*/ readonly mktemp!: pulumi.Output<Mktemp>;
    public /*out*/ readonly mv!: pulumi.Output<Mv>;
    public /*out*/ readonly path!: pulumi.Output<string>;
    public /*out*/ readonly rm!: pulumi.Output<Rm>;
    /**
     * The version to install.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a RuncInstall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuncInstallArgs, opts?: pulumi.ComponentResourceOptions) {
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
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["mktemp"] = undefined /*out*/;
            resourceInputs["mv"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["rm"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["mktemp"] = undefined /*out*/;
            resourceInputs["mv"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["rm"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuncInstall.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a RuncInstall resource.
 */
export interface RuncInstallArgs {
    /**
     * The CPU architecture.
     */
    architecture?: pulumi.Input<enums.remote.Architecture>;
    /**
     * The connection details.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Directory to install the binary.
     */
    directory?: pulumi.Input<string>;
    /**
     * The version of to install.
     */
    version?: pulumi.Input<string>;
}
