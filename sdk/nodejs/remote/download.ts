// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

import {Mkdir, Wget} from "../tools";

/**
 * Downloads the file specified by `url` onto a remote system.
 */
export class Download extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:Download';

    /**
     * Returns true if the given object is an instance of Download.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Download {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Download.__pulumiType;
    }

    /**
     * The parameters with which to connect to the remote host.
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * The fully qualified path on the remote system where the file should be downloaded to.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * The mkdir operation.
     */
    public /*out*/ readonly mkdir!: pulumi.Output<Mkdir>;
    /**
     * Remove the downloaded fiel when the resource is deleted.
     */
    public readonly removeOnDelete!: pulumi.Output<boolean>;
    /**
     * The URL of the file to be downloaded.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The wget operation.
     */
    public /*out*/ readonly wget!: pulumi.Output<Wget>;

    /**
     * Create a Download resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DownloadArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["removeOnDelete"] = args ? args.removeOnDelete : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["wget"] = undefined /*out*/;
        } else {
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["destination"] = undefined /*out*/;
            resourceInputs["mkdir"] = undefined /*out*/;
            resourceInputs["removeOnDelete"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["wget"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Download.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Download resource.
 */
export interface DownloadArgs {
    /**
     * The parameters with which to connect to the remote host.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * The fully qualified path on the remote system where the file should be downloaded to.
     */
    destination: pulumi.Input<string>;
    /**
     * Remove the downloaded fiel when the resource is deleted.
     */
    removeOnDelete?: boolean;
    /**
     * The URL of the file to be downloaded.
     */
    url: pulumi.Input<string>;
}
