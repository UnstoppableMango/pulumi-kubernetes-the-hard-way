// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as pulumiKubernetes from "@pulumi/kubernetes";
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

    /**
     * Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.
     *
     * This resource waits until its status is ready before registering success
     * for create/update, and populating output properties from the current state of the resource.
     * The following conditions are used to determine whether the resource creation has
     * succeeded or failed:
     *
     * 1. The Pod is scheduled ("PodScheduled"" '.status.condition' is true).
     * 2. The Pod is initialized ("Initialized" '.status.condition' is true).
     * 3. The Pod is ready ("Ready" '.status.condition' is true) and the '.status.phase' is
     *    set to "Running".
     * Or (for Jobs): The Pod succeeded ('.status.phase' set to "Succeeded").
     *
     * If the Pod has not reached a Ready state after 10 minutes, it will
     * time out and mark the resource update as Failed. You can override the default timeout value
     * by setting the 'customTimeouts' option on the resource.
     *
     * {{% examples %}}
     * ## Example Usage
     * {{% example %}}
     * ### Create a Pod with auto-naming
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as kubernetes from "@pulumi/kubernetes";
     *
     * const pod = new kubernetes.core.v1.Pod("pod", {spec: {
     *     containers: [{
     *         image: "nginx:1.14.2",
     *         name: "nginx",
     *         ports: [{
     *             containerPort: 80,
     *         }],
     *     }],
     * }});
     * ```
     * ```python
     * import pulumi
     * import pulumi_kubernetes as kubernetes
     *
     * pod = kubernetes.core.v1.Pod("pod", spec=kubernetes.core.v1.PodSpecArgs(
     *     containers=[kubernetes.core.v1.ContainerArgs(
     *         image="nginx:1.14.2",
     *         name="nginx",
     *         ports=[kubernetes.core.v1.ContainerPortArgs(
     *             container_port=80,
     *         )],
     *     )],
     * ))
     * ```
     * ```csharp
     * using System.Collections.Generic;
     * using System.Linq;
     * using Pulumi;
     * using Kubernetes = Pulumi.Kubernetes;
     *
     * return await Deployment.RunAsync(() => 
     * {
     *     var pod = new Kubernetes.Core.V1.Pod("pod", new()
     *     {
     *         Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
     *         {
     *             Containers = new[]
     *             {
     *                 new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
     *                 {
     *                     Image = "nginx:1.14.2",
     *                     Name = "nginx",
     *                     Ports = new[]
     *                     {
     *                         new Kubernetes.Types.Inputs.Core.V1.ContainerPortArgs
     *                         {
     *                             ContainerPortValue = 80,
     *                         },
     *                     },
     *                 },
     *             },
     *         },
     *     });
     *
     * });
     *
     * ```
     * ```go
     * package main
     *
     * import (
     * 	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
     * 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
     * )
     *
     * func main() {
     * 	pulumi.Run(func(ctx *pulumi.Context) error {
     * 		_, err := corev1.NewPod(ctx, "pod", &corev1.PodArgs{
     * 			Spec: &corev1.PodSpecArgs{
     * 				Containers: corev1.ContainerArray{
     * 					&corev1.ContainerArgs{
     * 						Image: pulumi.String("nginx:1.14.2"),
     * 						Name:  pulumi.String("nginx"),
     * 						Ports: corev1.ContainerPortArray{
     * 							&corev1.ContainerPortArgs{
     * 								ContainerPort: pulumi.Int(80),
     * 							},
     * 						},
     * 					},
     * 				},
     * 			},
     * 		})
     * 		if err != nil {
     * 			return err
     * 		}
     * 		return nil
     * 	})
     * }
     * ```
     * ```java
     * package generated_program;
     *
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.kubernetes.core_v1.Pod;
     * import com.pulumi.kubernetes.core_v1.PodArgs;
     * import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     *
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     *
     *     public static void stack(Context ctx) {
     *         var pod = new Pod("pod", PodArgs.builder()        
     *             .spec(PodSpecArgs.builder()
     *                 .containers(ContainerArgs.builder()
     *                     .image("nginx:1.14.2")
     *                     .name("nginx")
     *                     .ports(ContainerPortArgs.builder()
     *                         .containerPort(80)
     *                         .build())
     *                     .build())
     *                 .build())
     *             .build());
     *
     *     }
     * }
     * ```
     * ```yaml
     * description: Create a Pod with auto-naming
     * name: yaml-example
     * resources:
     *     pod:
     *         properties:
     *             spec:
     *                 containers:
     *                     - image: nginx:1.14.2
     *                       name: nginx
     *                       ports:
     *                         - containerPort: 80
     *         type: kubernetes:core/v1:Pod
     * runtime: yaml
     * ```
     * {{% /example %}}
     * {{% example %}}
     * ### Create a Pod with a user-specified name
     *
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * import * as kubernetes from "@pulumi/kubernetes";
     *
     * const pod = new kubernetes.core.v1.Pod("pod", {
     *     metadata: {
     *         name: "nginx",
     *     },
     *     spec: {
     *         containers: [{
     *             image: "nginx:1.14.2",
     *             name: "nginx",
     *             ports: [{
     *                 containerPort: 80,
     *             }],
     *         }],
     *     },
     * });
     * ```
     * ```python
     * import pulumi
     * import pulumi_kubernetes as kubernetes
     *
     * pod = kubernetes.core.v1.Pod("pod",
     *     metadata=kubernetes.meta.v1.ObjectMetaArgs(
     *         name="nginx",
     *     ),
     *     spec=kubernetes.core.v1.PodSpecArgs(
     *         containers=[kubernetes.core.v1.ContainerArgs(
     *             image="nginx:1.14.2",
     *             name="nginx",
     *             ports=[kubernetes.core.v1.ContainerPortArgs(
     *                 container_port=80,
     *             )],
     *         )],
     *     ))
     * ```
     * ```csharp
     * using System.Collections.Generic;
     * using System.Linq;
     * using Pulumi;
     * using Kubernetes = Pulumi.Kubernetes;
     *
     * return await Deployment.RunAsync(() => 
     * {
     *     var pod = new Kubernetes.Core.V1.Pod("pod", new()
     *     {
     *         Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
     *         {
     *             Name = "nginx",
     *         },
     *         Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
     *         {
     *             Containers = new[]
     *             {
     *                 new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
     *                 {
     *                     Image = "nginx:1.14.2",
     *                     Name = "nginx",
     *                     Ports = new[]
     *                     {
     *                         new Kubernetes.Types.Inputs.Core.V1.ContainerPortArgs
     *                         {
     *                             ContainerPortValue = 80,
     *                         },
     *                     },
     *                 },
     *             },
     *         },
     *     });
     *
     * });
     *
     * ```
     * ```go
     * package main
     *
     * import (
     * 	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
     * 	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
     * 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
     * )
     *
     * func main() {
     * 	pulumi.Run(func(ctx *pulumi.Context) error {
     * 		_, err := corev1.NewPod(ctx, "pod", &corev1.PodArgs{
     * 			Metadata: &metav1.ObjectMetaArgs{
     * 				Name: pulumi.String("nginx"),
     * 			},
     * 			Spec: &corev1.PodSpecArgs{
     * 				Containers: corev1.ContainerArray{
     * 					&corev1.ContainerArgs{
     * 						Image: pulumi.String("nginx:1.14.2"),
     * 						Name:  pulumi.String("nginx"),
     * 						Ports: corev1.ContainerPortArray{
     * 							&corev1.ContainerPortArgs{
     * 								ContainerPort: pulumi.Int(80),
     * 							},
     * 						},
     * 					},
     * 				},
     * 			},
     * 		})
     * 		if err != nil {
     * 			return err
     * 		}
     * 		return nil
     * 	})
     * }
     * ```
     * ```java
     * package generated_program;
     *
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.kubernetes.core_v1.Pod;
     * import com.pulumi.kubernetes.core_v1.PodArgs;
     * import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
     * import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     *
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     *
     *     public static void stack(Context ctx) {
     *         var pod = new Pod("pod", PodArgs.builder()        
     *             .metadata(ObjectMetaArgs.builder()
     *                 .name("nginx")
     *                 .build())
     *             .spec(PodSpecArgs.builder()
     *                 .containers(ContainerArgs.builder()
     *                     .image("nginx:1.14.2")
     *                     .name("nginx")
     *                     .ports(ContainerPortArgs.builder()
     *                         .containerPort(80)
     *                         .build())
     *                     .build())
     *                 .build())
     *             .build());
     *
     *     }
     * }
     * ```
     * ```yaml
     * description: Create a Pod with a user-specified name
     * name: yaml-example
     * resources:
     *     pod:
     *         properties:
     *             metadata:
     *                 name: nginx
     *             spec:
     *                 containers:
     *                     - image: nginx:1.14.2
     *                       name: nginx
     *                       ports:
     *                         - containerPort: 80
     *         type: kubernetes:core/v1:Pod
     * runtime: yaml
     * ```
     * {{% /example %}}
     * {{% /examples %}}
     */
    export interface PodManifestArgs {
        /**
         * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
         */
        apiVersion?: pulumi.Input<"v1">;
        /**
         * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
         */
        kind?: pulumi.Input<"Pod">;
        /**
         * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
         */
        metadata?: pulumi.Input<pulumiKubernetes.types.input.meta.v1.ObjectMeta>;
        /**
         * Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
         */
        spec?: pulumi.Input<pulumiKubernetes.types.input.core.v1.PodSpec>;
        /**
         * Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
         */
        status?: pulumi.Input<pulumiKubernetes.types.input.core.v1.PodStatus>;
    }

}

export namespace remote {
    /**
     * Props for resources that consume etcd configuration.
     */
    export interface EtcdConfigurationPropsArgs {
        /**
         * Path to the certificate authority file on the remote system.
         */
        caFilePath: pulumi.Input<string>;
        /**
         * Path to the certificate file on the remote system.
         */
        certFilePath: pulumi.Input<string>;
        /**
         * Etcd's data directory.
         */
        dataDirectory: pulumi.Input<string>;
        /**
         * Path to the etcd binary.
         */
        etcdPath: pulumi.Input<string>;
        /**
         * Internal IP of the etcd node.
         */
        internalIp: pulumi.Input<string>;
        /**
         * Path to the private key file on the remote system.
         */
        keyFilePath: pulumi.Input<string>;
        /**
         * Name of the etcd node.
         */
        name: pulumi.Input<string>;
    }

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
