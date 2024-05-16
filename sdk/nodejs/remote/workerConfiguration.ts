// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

/**
 * Kubernetes worker node configuration.
 */
export class WorkerConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing WorkerConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkerConfiguration {
        return new WorkerConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:WorkerConfiguration';

    /**
     * Returns true if the given object is an instance of WorkerConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkerConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkerConfiguration.__pulumiType;
    }

    /**
     * The parameters with which to connect to the remote host.
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;

    /**
     * Create a WorkerConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkerConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
        } else {
            resourceInputs["connection"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkerConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkerConfiguration resource.
 */
export interface WorkerConfigurationArgs {
    /**
     * The parameters with which to connect to the remote host.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
}