// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

/**
 * Abstraction over the `mv` utility on a remote system.
 */
export class Mv extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:tools:Mv';

    /**
     * Returns true if the given object is an instance of Mv.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Mv {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Mv.__pulumiType;
    }

    /**
     * Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
     */
    public readonly backup!: pulumi.Output<boolean>;
    /**
     * Represents the command run on the remote system.
     */
    public /*out*/ readonly command!: pulumi.Output<pulumiCommand.remote.Command>;
    /**
     * Corresponds to the --context option.
     */
    public readonly context!: pulumi.Output<boolean>;
    /**
     * Corresponds to the [CONTROL] argument for the --backup option.
     */
    public readonly control!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the [DEST] argument.
     */
    public readonly dest!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the [DIRECTORY] argument.
     */
    public readonly directory!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the --force option.
     */
    public readonly force!: pulumi.Output<boolean>;
    /**
     * Corresponds to the --no-clobber option.
     */
    public readonly noClobber!: pulumi.Output<boolean>;
    /**
     * Corresponds to the --no-target-directory option.
     */
    public readonly noTargetDirectory!: pulumi.Output<boolean>;
    /**
     * Corresponds to the [SOURCE] argument.
     */
    public readonly source!: pulumi.Output<string[]>;
    /**
     * Corresponds to the --strip-trailing-suffix option.
     */
    public readonly stripTrailingSlashes!: pulumi.Output<boolean>;
    /**
     * Corresponds to the --suffix option.
     */
    public readonly suffix!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the --target-directory option.
     */
    public readonly targetDirectory!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the --update option.
     */
    public readonly update!: pulumi.Output<boolean>;
    /**
     * Corresponds to the --verbose option.
     */
    public readonly verbose!: pulumi.Output<boolean>;

    /**
     * Create a Mv resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MvArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["backup"] = args ? args.backup : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["context"] = args ? args.context : undefined;
            resourceInputs["control"] = args ? args.control : undefined;
            resourceInputs["dest"] = args ? args.dest : undefined;
            resourceInputs["directory"] = args ? args.directory : undefined;
            resourceInputs["force"] = args ? args.force : undefined;
            resourceInputs["noClobber"] = args ? args.noClobber : undefined;
            resourceInputs["noTargetDirectory"] = args ? args.noTargetDirectory : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["stripTrailingSlashes"] = args ? args.stripTrailingSlashes : undefined;
            resourceInputs["suffix"] = args ? args.suffix : undefined;
            resourceInputs["targetDirectory"] = args ? args.targetDirectory : undefined;
            resourceInputs["update"] = args ? args.update : undefined;
            resourceInputs["verbose"] = args ? args.verbose : undefined;
            resourceInputs["command"] = undefined /*out*/;
        } else {
            resourceInputs["backup"] = undefined /*out*/;
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["context"] = undefined /*out*/;
            resourceInputs["control"] = undefined /*out*/;
            resourceInputs["dest"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["force"] = undefined /*out*/;
            resourceInputs["noClobber"] = undefined /*out*/;
            resourceInputs["noTargetDirectory"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["stripTrailingSlashes"] = undefined /*out*/;
            resourceInputs["suffix"] = undefined /*out*/;
            resourceInputs["targetDirectory"] = undefined /*out*/;
            resourceInputs["update"] = undefined /*out*/;
            resourceInputs["verbose"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Mv.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Mv resource.
 */
export interface MvArgs {
    /**
     * Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
     */
    backup?: boolean;
    /**
     * Connection details for the remote system.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Corresponds to the --context option.
     */
    context?: pulumi.Input<boolean>;
    /**
     * Corresponds to the [CONTROL] argument for the --backup option.
     */
    control?: pulumi.Input<string>;
    /**
     * Corresponds to the [DEST] argument.
     */
    dest?: pulumi.Input<string>;
    /**
     * Corresponds to the [DIRECTORY] argument.
     */
    directory?: pulumi.Input<string>;
    /**
     * Corresponds to the --force option.
     */
    force?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --no-clobber option.
     */
    noClobber?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --no-target-directory option.
     */
    noTargetDirectory?: pulumi.Input<boolean>;
    /**
     * Corresponds to the [SOURCE] argument.
     */
    source: pulumi.Input<string | pulumi.Input<string>[]>;
    /**
     * Corresponds to the --strip-trailing-suffix option.
     */
    stripTrailingSlashes?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --suffix option.
     */
    suffix?: pulumi.Input<string>;
    /**
     * Corresponds to the --target-directory option.
     */
    targetDirectory?: pulumi.Input<string>;
    /**
     * Corresponds to the --update option.
     */
    update?: pulumi.Input<boolean>;
    /**
     * Corresponds to the --verbose option.
     */
    verbose?: pulumi.Input<boolean>;
}
