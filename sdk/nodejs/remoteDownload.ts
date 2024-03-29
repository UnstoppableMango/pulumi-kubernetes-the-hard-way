// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import * as pulumiCommand from "@pulumi/command";

import {Mkdir, Wget} from "tools";

/**
 * Represents a file to be downloaded on a remote system.
 */
export class RemoteDownload extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:index:RemoteDownload';

    /**
     * Returns true if the given object is an instance of RemoteDownload.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemoteDownload {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemoteDownload.__pulumiType;
    }

    /**
     * Connection details for the remote system
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * The fully qualified path on the remote system where the file should be downloaded to.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * Represents the command used to create the remote directory.
     */
    public /*out*/ readonly mkdir!: pulumi.Output<Mkdir>;
    /**
     * The URL for the file to be downloaded.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Represents the wget command used to download the file.
     */
    public /*out*/ readonly wget!: pulumi.Output<Wget>;

    /**
     * Create a RemoteDownload resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteDownloadArgs, opts?: pulumi.ComponentResourceOptions) {
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
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["wget"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RemoteDownload.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a RemoteDownload resource.
 */
export interface RemoteDownloadArgs {
    /**
     * Connection details for the remote system
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * The fully qualified path on the remote system where the file should be downloaded to.
     */
    destination: pulumi.Input<string>;
    /**
     * Remove the downloaded file when the resource is deleted.
     */
    removeOnDelete?: pulumi.Input<boolean>;
    /**
     * The URL for the file to be downloaded.
     */
    url: pulumi.Input<string>;
}
