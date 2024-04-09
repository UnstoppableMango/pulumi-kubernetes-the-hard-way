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
 * Installs kubelet on a remote system.
 */
export class KubeletInstall extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:KubeletInstall';

    /**
     * Returns true if the given object is an instance of KubeletInstall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubeletInstall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubeletInstall.__pulumiType;
    }

    /**
     * The CPU architecture to install.
     */
    public readonly architecture!: pulumi.Output<enums.remote.Architecture>;
    /**
     * The name of the installed binary.
     */
    public /*out*/ readonly binName!: pulumi.Output<string | undefined>;
    /**
     * The parameters with which to connect to the remote host.
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * The directory to install the binary to.
     */
    public readonly directory!: pulumi.Output<string>;
    /**
     * The download operation.
     */
    public /*out*/ readonly download!: pulumi.Output<Download>;
    /**
     * The mkdir operation.
     */
    public /*out*/ readonly mkdir!: pulumi.Output<Mkdir>;
    /**
     * The mktemp operation.
     */
    public /*out*/ readonly mktemp!: pulumi.Output<Mktemp>;
    /**
     * The mv operation.
     */
    public /*out*/ readonly mv!: pulumi.Output<Mv>;
    /**
     * The path to the installed binary.
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * The rm operation.
     */
    public /*out*/ readonly rm!: pulumi.Output<Rm>;
    /**
     * The url used to download the binary.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * The version to install.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a KubeletInstall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubeletInstallArgs, opts?: pulumi.ComponentResourceOptions) {
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
            resourceInputs["binName"] = undefined /*out*/;
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["mktemp"] = undefined /*out*/;
            resourceInputs["mv"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["rm"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["binName"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["download"] = undefined /*out*/;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["mktemp"] = undefined /*out*/;
            resourceInputs["mv"] = undefined /*out*/;
            resourceInputs["path"] = undefined /*out*/;
            resourceInputs["rm"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KubeletInstall.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a KubeletInstall resource.
 */
export interface KubeletInstallArgs {
    /**
     * The CPU architecture to install.
     */
    architecture?: pulumi.Input<enums.remote.Architecture>;
    /**
     * The parameters with which to connect to the remote host.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * The directory to install the binary to.
     */
    directory?: pulumi.Input<string>;
    /**
     * The version to install.
     */
    version?: pulumi.Input<string>;
}
