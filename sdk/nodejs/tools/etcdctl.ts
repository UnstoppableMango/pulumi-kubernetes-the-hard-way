// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

/**
 * Abstraction over the `etcdctl` utility on a remote system.
 */
export class Etcdctl extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:tools:Etcdctl';

    /**
     * Returns true if the given object is an instance of Etcdctl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Etcdctl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Etcdctl.__pulumiType;
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
     * The command to run on create.
     */
    public readonly create!: pulumi.Output<outputs.tools.EtcdctlOpts | undefined>;
    /**
     * The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
     * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
     * Command resource from previous create or update steps.
     */
    public readonly delete!: pulumi.Output<outputs.tools.EtcdctlOpts | undefined>;
    /**
     * Environment variables
     */
    public readonly environment!: pulumi.Output<{[key: string]: string}>;
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
     * TODO
     */
    public readonly triggers!: pulumi.Output<any[]>;
    /**
     * The command to run on update, if empty, create will 
     * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR 
     * are set to the stdout and stderr properties of the Command resource from previous 
     * create or update steps.
     */
    public readonly update!: pulumi.Output<outputs.tools.EtcdctlOpts | undefined>;

    /**
     * Create a Etcdctl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EtcdctlArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["binaryPath"] = args ? args.binaryPath : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["create"] = args ? args.create : undefined;
            resourceInputs["delete"] = args ? args.delete : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["stdin"] = args ? args.stdin : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["update"] = args ? args.update : undefined;
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
        } else {
            resourceInputs["binaryPath"] = undefined /*out*/;
            resourceInputs["command"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["create"] = undefined /*out*/;
            resourceInputs["delete"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdin"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
            resourceInputs["triggers"] = undefined /*out*/;
            resourceInputs["update"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Etcdctl.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Etcdctl resource.
 */
export interface EtcdctlArgs {
    /**
     * Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
     */
    binaryPath?: pulumi.Input<string>;
    /**
     * Connection details for the remote system
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * The command to run on create.
     */
    create?: inputs.tools.EtcdctlOptsArgs;
    /**
     * The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
     * and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
     * Command resource from previous create or update steps.
     */
    delete?: inputs.tools.EtcdctlOptsArgs;
    /**
     * Environment variables
     */
    environment?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * TODO
     */
    stdin?: pulumi.Input<string>;
    /**
     * TODO
     */
    triggers?: pulumi.Input<any[]>;
    /**
     * The command to run on update, if empty, create will 
     * run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR 
     * are set to the stdout and stderr properties of the Command resource from previous 
     * create or update steps.
     */
    update?: inputs.tools.EtcdctlOptsArgs;
}
