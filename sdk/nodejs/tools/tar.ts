// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

/**
 * Abstracion over the `tar` utility on a remote system.
 */
export class Tar extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:tools:Tar';

    /**
     * Returns true if the given object is an instance of Tar.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tar {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tar.__pulumiType;
    }

    /**
     * Corresponds to the [ARCHIVE] argument.
     */
    public readonly archive!: pulumi.Output<string>;
    /**
     * Represents the remote `tar` operation.
     */
    public /*out*/ readonly command!: pulumi.Output<pulumiCommand.remote.Command>;
    /**
     * Corresponds to the --directory option.
     */
    public readonly directory!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the --extract option.
     */
    public readonly extract!: pulumi.Output<boolean>;
    /**
     * Corresponds to the [FILE] argument.
     */
    public readonly files!: pulumi.Output<string[]>;
    /**
     * Corresponds to the --gzip option.
     */
    public readonly gzip!: pulumi.Output<boolean | undefined>;
    /**
     * The process' stderr.
     */
    public /*out*/ readonly stderr!: pulumi.Output<string>;
    /**
     * The process' stdin.
     */
    public /*out*/ readonly stdin!: pulumi.Output<string | undefined>;
    /**
     * The process' stdout.
     */
    public /*out*/ readonly stdout!: pulumi.Output<string>;
    /**
     * Corresponds to the --strip-components option.
     */
    public readonly stripComponents!: pulumi.Output<number | undefined>;

    /**
     * Create a Tar resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TarArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.archive === undefined) && !opts.urn) {
                throw new Error("Missing required property 'archive'");
            }
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["archive"] = args ? args.archive : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["directory"] = args ? args.directory : undefined;
            resourceInputs["extract"] = args ? args.extract : undefined;
            resourceInputs["files"] = args ? args.files : undefined;
            resourceInputs["gzip"] = args ? args.gzip : undefined;
            resourceInputs["stripComponents"] = args ? args.stripComponents : undefined;
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdin"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
        } else {
            resourceInputs["archive"] = undefined /*out*/;
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["extract"] = undefined /*out*/;
            resourceInputs["files"] = undefined /*out*/;
            resourceInputs["gzip"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdin"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
            resourceInputs["stripComponents"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tar.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Tar resource.
 */
export interface TarArgs {
    /**
     * Corresponds to the [ARCHIVE] argument.
     */
    archive: pulumi.Input<string>;
    /**
     * Connection details for the remote system.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Corresponds to the --directory option.
     */
    directory?: pulumi.Input<string>;
    /**
     * Corresponds to the --extract option.
     */
    extract?: pulumi.Input<boolean>;
    /**
     * Corresponds to the [FILE] argument.
     */
    files?: pulumi.Input<pulumi.Input<string>[] | string>;
    /**
     * Corresponds to the --gzip option.
     */
    gzip?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --strip-components option.
     */
    stripComponents?: pulumi.Input<number>;
}
