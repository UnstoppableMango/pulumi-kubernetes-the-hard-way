// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

/**
 * Abstracion over the `mktemp` utility on a remote system.
 */
export class Mktemp extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:tools:Mktemp';

    /**
     * Returns true if the given object is an instance of Mktemp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Mktemp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Mktemp.__pulumiType;
    }

    /**
     * Represents the remote `tar` operation.
     */
    public /*out*/ readonly command!: pulumi.Output<pulumiCommand.remote.Command>;
    /**
     * Corresponds to the --directory option.
     */
    public readonly directory!: pulumi.Output<boolean>;
    /**
     * Corresponds to the --dry-run option.
     */
    public readonly dryRun!: pulumi.Output<boolean>;
    /**
     * Corresponds to the --quiet option.
     */
    public readonly quiet!: pulumi.Output<boolean>;
    /**
     * Corresponds to the --suffix option.
     */
    public readonly suffix!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the [TEMPLATE] arg.
     */
    public readonly template!: pulumi.Output<string>;
    /**
     * Corresponds to the --tmpdir option.
     */
    public readonly tmpdir!: pulumi.Output<string | undefined>;

    /**
     * Create a Mktemp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MktempArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.template === undefined) && !opts.urn) {
                throw new Error("Missing required property 'template'");
            }
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["directory"] = args ? args.directory : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["quiet"] = args ? args.quiet : undefined;
            resourceInputs["suffix"] = args ? args.suffix : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["tmpdir"] = args ? args.tmpdir : undefined;
            resourceInputs["command"] = undefined /*out*/;
        } else {
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["dryRun"] = undefined /*out*/;
            resourceInputs["quiet"] = undefined /*out*/;
            resourceInputs["suffix"] = undefined /*out*/;
            resourceInputs["template"] = undefined /*out*/;
            resourceInputs["tmpdir"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Mktemp.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Mktemp resource.
 */
export interface MktempArgs {
    /**
     * Connection details for the remote system.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Corresponds to the --directory option.
     */
    directory?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --dry-run option.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --quiet option.
     */
    quiet?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --suffix option.
     */
    suffix?: pulumi.Input<string>;
    /**
     * Corresponds to the [TEMPLATE] arg.
     */
    template: pulumi.Input<string>;
    /**
     * Corresponds to the --tmpdir option.
     */
    tmpdir?: pulumi.Input<string>;
}
