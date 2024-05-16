// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiCommand from "@pulumi/command";

import {KubeApiServerProps, SystemdService} from "./index";

/**
 * Kube API Server systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.
 */
export class KubeApiServerService extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes-the-hard-way:remote:KubeApiServerService';

    /**
     * Returns true if the given object is an instance of KubeApiServerService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubeApiServerService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubeApiServerService.__pulumiType;
    }

    /**
     * If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
     */
    public readonly clientCaFile!: pulumi.Output<string | undefined>;
    /**
     * KubeApiServer configuration.
     */
    public readonly configuration!: pulumi.Output<KubeApiServerProps>;
    /**
     * The parameters with which to connect to the remote host.
     */
    public readonly connection!: pulumi.Output<pulumiCommand.types.output.remote.Connection>;
    /**
     * Optional systemd unit description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The location to create the service file.
     */
    public readonly directory!: pulumi.Output<string | undefined>;
    /**
     * Optional systemd unit documentation
     */
    public readonly documentation!: pulumi.Output<string | undefined>;
    /**
     * List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
     */
    public readonly etcdServers!: pulumi.Output<string | undefined>;
    /**
     * Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
     */
    public readonly restart!: pulumi.Output<enums.remote.SystemdServiceRestart | undefined>;
    /**
     * Optionally override the systemd service RestartSec. Defaults to `5`.
     */
    public readonly restartSec!: pulumi.Output<string | undefined>;
    /**
     * The remote systemd service.
     */
    public /*out*/ readonly service!: pulumi.Output<SystemdService>;
    /**
     * Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
     */
    public readonly wantedBy!: pulumi.Output<string | undefined>;

    /**
     * Create a KubeApiServerService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubeApiServerServiceArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            resourceInputs["clientCaFile"] = args ? args.clientCaFile : undefined;
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(pulumiCommand.types.input.remote.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["directory"] = args ? args.directory : undefined;
            resourceInputs["documentation"] = args ? args.documentation : undefined;
            resourceInputs["etcdServers"] = args ? args.etcdServers : undefined;
            resourceInputs["restart"] = args ? args.restart : undefined;
            resourceInputs["restartSec"] = args ? args.restartSec : undefined;
            resourceInputs["wantedBy"] = args ? args.wantedBy : undefined;
            resourceInputs["service"] = undefined /*out*/;
        } else {
            resourceInputs["clientCaFile"] = undefined /*out*/;
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["directory"] = undefined /*out*/;
            resourceInputs["documentation"] = undefined /*out*/;
            resourceInputs["etcdServers"] = undefined /*out*/;
            resourceInputs["restart"] = undefined /*out*/;
            resourceInputs["restartSec"] = undefined /*out*/;
            resourceInputs["service"] = undefined /*out*/;
            resourceInputs["wantedBy"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KubeApiServerService.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a KubeApiServerService resource.
 */
export interface KubeApiServerServiceArgs {
    /**
     * If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
     */
    clientCaFile?: pulumi.Input<string>;
    /**
     * KubeApiServer configuration.
     */
    configuration: pulumi.Input<KubeApiServerProps>;
    /**
     * The parameters with which to connect to the remote host.
     */
    connection: pulumi.Input<pulumiCommand.types.input.remote.ConnectionArgs>;
    /**
     * Optional systemd unit description.
     */
    description?: pulumi.Input<string>;
    /**
     * The location to create the service file.
     */
    directory?: pulumi.Input<string>;
    /**
     * Optional systemd unit documentation
     */
    documentation?: pulumi.Input<string>;
    /**
     * List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
     */
    etcdServers?: pulumi.Input<string>;
    /**
     * Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
     */
    restart?: pulumi.Input<enums.remote.SystemdServiceRestart>;
    /**
     * Optionally override the systemd service RestartSec. Defaults to `5`.
     */
    restartSec?: pulumi.Input<string>;
    /**
     * Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
     */
    wantedBy?: pulumi.Input<string>;
}
