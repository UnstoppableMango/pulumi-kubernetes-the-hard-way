// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as pulumiTls from "@pulumi/tls";

export namespace config {
    export interface KubeconfigAdminOptions {
        publicIp?: string;
        type: "admin";
    }

    export interface KubeconfigAdminOptionsArgs {
        publicIp?: pulumi.Input<string>;
        type: "admin";
    }

    export interface KubeconfigKubeControllerManagerOptions {
        publicIp?: string;
        type: "kube-controller-manager";
    }

    export interface KubeconfigKubeControllerManagerOptionsArgs {
        publicIp?: pulumi.Input<string>;
        type: "kube-controller-manager";
    }

    export interface KubeconfigKubeProxyOptions {
        publicIp?: string;
        type: "kube-proxy";
    }

    export interface KubeconfigKubeProxyOptionsArgs {
        publicIp?: pulumi.Input<string>;
        type: "kube-proxy";
    }

    export interface KubeconfigKubeSchedulerOptions {
        publicIp?: string;
        type: "kube-scheduler";
    }

    export interface KubeconfigKubeSchedulerOptionsArgs {
        publicIp?: pulumi.Input<string>;
        type: "kube-scheduler";
    }

    export interface KubeconfigWorkerOptions {
        name: string;
        publicIp: string;
        type?: "worker";
    }

    export interface KubeconfigWorkerOptionsArgs {
        name: pulumi.Input<string>;
        publicIp: pulumi.Input<string>;
        type?: "worker";
    }

}

export namespace remote {
    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options
     */
    export interface SystemdInstallSectionArgs {
        /**
         * A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
         */
        wantedBy?: pulumi.Input<pulumi.Input<string>[]>;
    }

    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
     */
    export interface SystemdServiceSectionArgs {
        /**
         * Commands that are executed when this service is started.
         */
        execStart?: pulumi.Input<string>;
        /**
         * Specifies when the manager should consider the service to be finished.
         */
        exitType?: pulumi.Input<enums.remote.SystemdServiceExitType>;
        /**
         * Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
         */
        restart?: pulumi.Input<enums.remote.SystemdServiceRestart>;
        /**
         * Configures the time to sleep before restarting a service (as configured with Restart=).
         */
        restartSec?: pulumi.Input<string>;
        /**
         * Configures the mechanism via which the service notifies the manager that the service start-up has finished.
         */
        type?: pulumi.Input<enums.remote.SystemdServiceType>;
    }

    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#
     */
    export interface SystemdUnitSectionArgs {
        /**
         * Configures requirement dependencies, very similar in style to Requires=.
         */
        bindsTo?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * A short human readable title of the unit.
         */
        description?: pulumi.Input<string>;
        /**
         * A space-separated list of URIs referencing documentation for this unit or its configuration.
         */
        documentation?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Similar to Wants=, but declares a stronger requirement dependency.
         */
        requires?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
         */
        requisite?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * Configures (weak) requirement dependencies on other units.
         */
        wants?: pulumi.Input<pulumi.Input<string>[]>;
    }
}

export namespace tls {
    /**
     * TODO
     */
    export interface ClusterPkiNodeArgs {
        /**
         * The IP address of the node
         */
        ip?: pulumi.Input<string>;
        /**
         * The role a node should be configured for
         */
        role?: pulumi.Input<enums.tls.NodeRole>;
    }

}
