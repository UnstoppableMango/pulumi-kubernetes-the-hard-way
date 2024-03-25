// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

/**
 * Polyfill for `tls.CertRequestSubject`.
 */
export interface CertRequestSubject {
    /**
     * Distinguished name: CN
     */
    commonName?: string;
    /**
     * Distinguished name: C
     */
    country?: string;
    /**
     * Distinguished name: L
     */
    locality?: string;
    /**
     * Distinguished name: O
     */
    organization?: string;
    /**
     * Distinguished name: OU
     */
    organizationalUnit?: string;
    /**
     * Distinguished name: PC
     */
    postalCode?: string;
    /**
     * Distinguished name: ST
     */
    province?: string;
    /**
     * Distinguished name: SERIALNUMBER
     */
    serialNumber?: string;
    /**
     * Distinguished name: STREET
     */
    streetAddresses?: string[];
}

/**
 * Polyfill for `tls.CertRequestSubject`.
 */
export interface CertRequestSubjectArgs {
    /**
     * Distinguished name: CN
     */
    commonName?: pulumi.Input<string>;
    /**
     * Distinguished name: C
     */
    country?: pulumi.Input<string>;
    /**
     * Distinguished name: L
     */
    locality?: pulumi.Input<string>;
    /**
     * Distinguished name: O
     */
    organization?: pulumi.Input<string>;
    /**
     * Distinguished name: OU
     */
    organizationalUnit?: pulumi.Input<string>;
    /**
     * Distinguished name: PC
     */
    postalCode?: pulumi.Input<string>;
    /**
     * Distinguished name: ST
     */
    province?: pulumi.Input<string>;
    /**
     * Distinguished name: SERIALNUMBER
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * Distinguished name: STREET
     */
    streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * Node inputs for the PKI.
 */
export interface ClusterPkiNodeArgs {
    /**
     * The IP address of the node.
     */
    ip?: pulumi.Input<string>;
    role?: pulumi.Input<enums.NodeRole>;
}

/**
 * Instructions for how to connect to a remote endpoint. Polyfill for `command.ConnectionArgs`.
 */
export interface Connection {
    /**
     * SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
     */
    agentSocketPath?: string;
    /**
     * Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
     */
    dialErrorLimit?: number;
    /**
     * The address of the resource to connect to.
     */
    host: string;
    /**
     * The password we should use for the connection.
     */
    password?: string;
    /**
     * Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
     */
    perDialTimeout?: number;
    /**
     * The port to connect to.
     */
    port?: number;
    /**
     * The contents of an SSH key to use for the connection. This takes preference over the password if provided.
     */
    privateKey?: string;
    /**
     * The password to use in case the private key is encrypted.
     */
    privateKeyPassword?: string;
    /**
     * The user that we should use for the connection.
     */
    user?: string;
}
/**
 * connectionProvideDefaults sets the appropriate defaults for Connection
 */
export function connectionProvideDefaults(val: Connection): Connection {
    return {
        ...val,
        dialErrorLimit: (val.dialErrorLimit) ?? 10,
        perDialTimeout: (val.perDialTimeout) ?? 15,
        port: (val.port) ?? 22,
        user: (val.user) ?? "root",
    };
}

/**
 * Instructions for how to connect to a remote endpoint. Polyfill for `command.ConnectionArgs`.
 */
export interface ConnectionArgs {
    /**
     * SSH Agent socket path. Default to environment variable SSH_AUTH_SOCK if present.
     */
    agentSocketPath?: pulumi.Input<string>;
    /**
     * Max allowed errors on trying to dial the remote host. -1 set count to unlimited. Default value is 10.
     */
    dialErrorLimit?: pulumi.Input<number>;
    /**
     * The address of the resource to connect to.
     */
    host: pulumi.Input<string>;
    /**
     * The password we should use for the connection.
     */
    password?: pulumi.Input<string>;
    /**
     * Max number of seconds for each dial attempt. 0 implies no maximum. Default value is 15 seconds.
     */
    perDialTimeout?: pulumi.Input<number>;
    /**
     * The port to connect to.
     */
    port?: pulumi.Input<number>;
    /**
     * The contents of an SSH key to use for the connection. This takes preference over the password if provided.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The password to use in case the private key is encrypted.
     */
    privateKeyPassword?: pulumi.Input<string>;
    /**
     * The user that we should use for the connection.
     */
    user?: pulumi.Input<string>;
}
/**
 * connectionArgsProvideDefaults sets the appropriate defaults for ConnectionArgs
 */
export function connectionArgsProvideDefaults(val: ConnectionArgs): ConnectionArgs {
    return {
        ...val,
        dialErrorLimit: (val.dialErrorLimit) ?? 10,
        perDialTimeout: (val.perDialTimeout) ?? 15,
        port: (val.port) ?? 22,
        user: (val.user) ?? "root",
    };
}

/**
 * Polyfill for `pulumi.ComponentResourceOptions`.
 */
export interface ResourceOptions {
    parent?: any;
}

/**
 * Polyfill for `pulumi.ComponentResourceOptions`.
 */
export interface ResourceOptionsArgs {
    parent?: any;
}
