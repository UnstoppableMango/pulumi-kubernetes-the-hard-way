// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

/**
 * Abstraction over the `mkdir` utility on a remote system.
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
     * Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
     */
    public readonly binaryPath!: pulumi.Output<string>;
    /**
     * The underlying command
     */
    public /*out*/ readonly command!: pulumi.Output<pulumiCommand.remote.Command>;
    /**
     * Connection details for the remote system
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * Corresponds to the `--directory` option.
     */
    public readonly directory!: pulumi.Output<boolean | undefined>;
    /**
     * Corresponds to the `--dry-run` option.
     */
    public readonly dryRun!: pulumi.Output<boolean>;
    /**
     * Environment variables
     */
    public readonly environment!: pulumi.Output<{[key: string]: string}>;
    /**
     * At what stage(s) in the resource lifecycle should the command be run
     */
    public readonly lifecycle!: pulumi.Output<enums.tools.CommandLifecycle | undefined>;
    /**
     * Corresponds to the `--quiet` option.
     */
    public readonly quiet!: pulumi.Output<boolean>;
    /**
     * TODO
     */
    public /*out*/ readonly stderr!: pulumi.Output<string>;
    /**
     * TODO
     */
    public readonly stdin!: pulumi.Output<string | undefined>;
    /**
     * TODO
     */
    public /*out*/ readonly stdout!: pulumi.Output<string>;
    /**
     * Corresponds to the `--suffix` option.
     */
    public readonly suffix!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the [TEMPLATE] argument.
     */
    public readonly template!: pulumi.Output<string | undefined>;
    /**
     * Corresponds to the `--tmpdir` option.
     */
    public readonly tmpdir!: pulumi.Output<string | undefined>;
    /**
     * TODO
     */
    public readonly triggers!: pulumi.Output<any[]>;

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
            resourceInputs["binaryPath"] = args ? args.binaryPath : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["directory"] = args ? args.directory : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["lifecycle"] = args ? args.lifecycle : undefined;
            resourceInputs["quiet"] = args ? args.quiet : undefined;
            resourceInputs["stdin"] = args ? args.stdin : undefined;
            resourceInputs["suffix"] = args ? args.suffix : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["tmpdir"] = args ? args.tmpdir : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
        } else {
            resourceInputs["binaryPath"] = undefined /*out*/;
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["dryRun"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["lifecycle"] = undefined /*out*/;
            resourceInputs["quiet"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdin"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
            resourceInputs["suffix"] = undefined /*out*/;
            resourceInputs["template"] = undefined /*out*/;
            resourceInputs["tmpdir"] = undefined /*out*/;
            resourceInputs["triggers"] = undefined /*out*/;
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
     * Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
     */
    binaryPath?: pulumi.Input<string>;
    /**
     * Connection details for the remote system
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Corresponds to the `--directory` option.
     */
    directory?: pulumi.Input<boolean>;
    /**
     * Corresponds to the `--dry-run` option.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Environment variables
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * At what stage(s) in the resource lifecycle should the command be run
     */
    lifecycle?: enums.tools.CommandLifecycle;
    /**
     * Corresponds to the `--quiet` option.
     */
    quiet?: pulumi.Input<boolean>;
    /**
     * TODO
     */
    stdin?: pulumi.Input<string>;
    /**
     * Corresponds to the `--suffix` option.
     */
    suffix?: pulumi.Input<string>;
    /**
     * Corresponds to the [TEMPLATE] argument.
     */
    template?: pulumi.Input<string>;
    /**
     * Corresponds to the `--tmpdir` option.
     */
    tmpdir?: pulumi.Input<string>;
    /**
     * TODO
     */
    triggers?: pulumi.Input<any[]>;
}
