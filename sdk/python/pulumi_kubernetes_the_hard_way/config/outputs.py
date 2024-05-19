# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
import pulumi_kubernetes

__all__ = [
    'Cluster',
    'Context',
    'Kubeconfig',
    'KubeletConfiguration',
    'KubeletConfigurationAuthentication',
    'KubeletConfigurationAuthenticationAnonymous',
    'KubeletConfigurationAuthenticationWebhook',
    'KubeletConfigurationAuthenticationx509',
    'KubeletConfigurationAuthorization',
    'PodManifest',
    'User',
]

@pulumi.output_type
class Cluster(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateAuthorityData":
            suggest = "certificate_authority_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Cluster. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Cluster.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Cluster.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_authority_data: str,
                 server: str):
        """
        :param str certificate_authority_data: TODO
        :param str server: TODO
        """
        pulumi.set(__self__, "certificate_authority_data", certificate_authority_data)
        pulumi.set(__self__, "server", server)

    @property
    @pulumi.getter(name="certificateAuthorityData")
    def certificate_authority_data(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "certificate_authority_data")

    @property
    @pulumi.getter
    def server(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "server")


@pulumi.output_type
class Context(dict):
    def __init__(__self__, *,
                 cluster: str,
                 user: str):
        """
        :param str cluster: TODO
        :param str user: TODO
        """
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def cluster(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter
    def user(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "user")


@pulumi.output_type
class Kubeconfig(dict):
    def __init__(__self__, *,
                 clusters: Sequence['outputs.Cluster'],
                 contexts: Sequence['outputs.Context'],
                 users: Sequence['outputs.User']):
        pulumi.set(__self__, "clusters", clusters)
        pulumi.set(__self__, "contexts", contexts)
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.Cluster']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter
    def contexts(self) -> Sequence['outputs.Context']:
        return pulumi.get(self, "contexts")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.User']:
        return pulumi.get(self, "users")


@pulumi.output_type
class KubeletConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"
        elif key == "cgroupDriver":
            suggest = "cgroup_driver"
        elif key == "clusterDNS":
            suggest = "cluster_dns"
        elif key == "clusterDomain":
            suggest = "cluster_domain"
        elif key == "containerRuntimeEndpoint":
            suggest = "container_runtime_endpoint"
        elif key == "podCIDR":
            suggest = "pod_cidr"
        elif key == "resolvConf":
            suggest = "resolv_conf"
        elif key == "runtimeRequestTimeout":
            suggest = "runtime_request_timeout"
        elif key == "tlsCertFile":
            suggest = "tls_cert_file"
        elif key == "tlsPrivateKeyFile":
            suggest = "tls_private_key_file"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KubeletConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KubeletConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KubeletConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_version: str,
                 authentication: 'outputs.KubeletConfigurationAuthentication',
                 authorization: 'outputs.KubeletConfigurationAuthorization',
                 cgroup_driver: str,
                 cluster_dns: Sequence[str],
                 cluster_domain: str,
                 container_runtime_endpoint: str,
                 kind: str,
                 pod_cidr: str,
                 resolv_conf: str,
                 runtime_request_timeout: str,
                 tls_cert_file: str,
                 tls_private_key_file: str):
        """
        :param str cgroup_driver: TODO
        :param Sequence[str] cluster_dns: TODO
        :param str cluster_domain: TODO
        :param str container_runtime_endpoint: TODO
        :param str pod_cidr: TODO
        :param str resolv_conf: TODO
        :param str runtime_request_timeout: TODO
        :param str tls_cert_file: TODO
        :param str tls_private_key_file: TODO
        """
        pulumi.set(__self__, "api_version", 'kubelet.config.k8s.io/v1beta1')
        pulumi.set(__self__, "authentication", authentication)
        pulumi.set(__self__, "authorization", authorization)
        pulumi.set(__self__, "cgroup_driver", cgroup_driver)
        pulumi.set(__self__, "cluster_dns", cluster_dns)
        pulumi.set(__self__, "cluster_domain", cluster_domain)
        pulumi.set(__self__, "container_runtime_endpoint", container_runtime_endpoint)
        pulumi.set(__self__, "kind", 'KubeletConfiguration')
        pulumi.set(__self__, "pod_cidr", pod_cidr)
        pulumi.set(__self__, "resolv_conf", resolv_conf)
        pulumi.set(__self__, "runtime_request_timeout", runtime_request_timeout)
        pulumi.set(__self__, "tls_cert_file", tls_cert_file)
        pulumi.set(__self__, "tls_private_key_file", tls_private_key_file)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> str:
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def authentication(self) -> 'outputs.KubeletConfigurationAuthentication':
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter
    def authorization(self) -> 'outputs.KubeletConfigurationAuthorization':
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="cgroupDriver")
    def cgroup_driver(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "cgroup_driver")

    @property
    @pulumi.getter(name="clusterDNS")
    def cluster_dns(self) -> Sequence[str]:
        """
        TODO
        """
        return pulumi.get(self, "cluster_dns")

    @property
    @pulumi.getter(name="clusterDomain")
    def cluster_domain(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "cluster_domain")

    @property
    @pulumi.getter(name="containerRuntimeEndpoint")
    def container_runtime_endpoint(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "container_runtime_endpoint")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="podCIDR")
    def pod_cidr(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "pod_cidr")

    @property
    @pulumi.getter(name="resolvConf")
    def resolv_conf(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "resolv_conf")

    @property
    @pulumi.getter(name="runtimeRequestTimeout")
    def runtime_request_timeout(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "runtime_request_timeout")

    @property
    @pulumi.getter(name="tlsCertFile")
    def tls_cert_file(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "tls_cert_file")

    @property
    @pulumi.getter(name="tlsPrivateKeyFile")
    def tls_private_key_file(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "tls_private_key_file")


@pulumi.output_type
class KubeletConfigurationAuthentication(dict):
    def __init__(__self__, *,
                 anonymous: 'outputs.KubeletConfigurationAuthenticationAnonymous',
                 webhook: 'outputs.KubeletConfigurationAuthenticationWebhook',
                 x509: 'outputs.KubeletConfigurationAuthenticationx509'):
        pulumi.set(__self__, "anonymous", anonymous)
        pulumi.set(__self__, "webhook", webhook)
        pulumi.set(__self__, "x509", x509)

    @property
    @pulumi.getter
    def anonymous(self) -> 'outputs.KubeletConfigurationAuthenticationAnonymous':
        return pulumi.get(self, "anonymous")

    @property
    @pulumi.getter
    def webhook(self) -> 'outputs.KubeletConfigurationAuthenticationWebhook':
        return pulumi.get(self, "webhook")

    @property
    @pulumi.getter
    def x509(self) -> 'outputs.KubeletConfigurationAuthenticationx509':
        return pulumi.get(self, "x509")


@pulumi.output_type
class KubeletConfigurationAuthenticationAnonymous(dict):
    def __init__(__self__, *,
                 enabled: bool):
        """
        :param bool enabled: TODO
        """
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        TODO
        """
        return pulumi.get(self, "enabled")


@pulumi.output_type
class KubeletConfigurationAuthenticationWebhook(dict):
    def __init__(__self__, *,
                 enabled: bool):
        """
        :param bool enabled: TODO
        """
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        TODO
        """
        return pulumi.get(self, "enabled")


@pulumi.output_type
class KubeletConfigurationAuthenticationx509(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientCAFile":
            suggest = "client_ca_file"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KubeletConfigurationAuthenticationx509. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KubeletConfigurationAuthenticationx509.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KubeletConfigurationAuthenticationx509.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_ca_file: str):
        """
        :param str client_ca_file: TODO
        """
        pulumi.set(__self__, "client_ca_file", client_ca_file)

    @property
    @pulumi.getter(name="clientCAFile")
    def client_ca_file(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "client_ca_file")


@pulumi.output_type
class KubeletConfigurationAuthorization(dict):
    def __init__(__self__, *,
                 mode: str):
        pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> str:
        return pulumi.get(self, "mode")


@pulumi.output_type
class PodManifest(dict):
    """
    Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.

    This resource waits until its status is ready before registering success
    for create/update, and populating output properties from the current state of the resource.
    The following conditions are used to determine whether the resource creation has
    succeeded or failed:

    1. The Pod is scheduled ("PodScheduled"" '.status.condition' is true).
    2. The Pod is initialized ("Initialized" '.status.condition' is true).
    3. The Pod is ready ("Ready" '.status.condition' is true) and the '.status.phase' is
       set to "Running".
    Or (for Jobs): The Pod succeeded ('.status.phase' set to "Succeeded").

    If the Pod has not reached a Ready state after 10 minutes, it will
    time out and mark the resource update as Failed. You can override the default timeout value
    by setting the 'customTimeouts' option on the resource.

    {{% examples %}}
    ## Example Usage
    {{% example %}}
    ### Create a Pod with auto-naming

    ```typescript
    import * as pulumi from "@pulumi/pulumi";
    import * as kubernetes from "@pulumi/kubernetes";

    const pod = new kubernetes.core.v1.Pod("pod", {spec: {
        containers: [{
            image: "nginx:1.14.2",
            name: "nginx",
            ports: [{
                containerPort: 80,
            }],
        }],
    }});
    ```
    ```python
    import pulumi
    import pulumi_kubernetes as kubernetes

    pod = kubernetes.core.v1.Pod("pod", spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            image="nginx:1.14.2",
            name="nginx",
            ports=[kubernetes.core.v1.ContainerPortArgs(
                container_port=80,
            )],
        )],
    ))
    ```
    ```csharp
    using System.Collections.Generic;
    using System.Linq;
    using Pulumi;
    using Kubernetes = Pulumi.Kubernetes;

    return await Deployment.RunAsync(() => 
    {
        var pod = new Kubernetes.Core.V1.Pod("pod", new()
        {
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {
                Containers = new[]
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Image = "nginx:1.14.2",
                        Name = "nginx",
                        Ports = new[]
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerPortArgs
                            {
                                ContainerPortValue = 80,
                            },
                        },
                    },
                },
            },
        });

    });

    ```
    ```go
    package main

    import (
    	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
    	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    )

    func main() {
    	pulumi.Run(func(ctx *pulumi.Context) error {
    		_, err := corev1.NewPod(ctx, "pod", &corev1.PodArgs{
    			Spec: &corev1.PodSpecArgs{
    				Containers: corev1.ContainerArray{
    					&corev1.ContainerArgs{
    						Image: pulumi.String("nginx:1.14.2"),
    						Name:  pulumi.String("nginx"),
    						Ports: corev1.ContainerPortArray{
    							&corev1.ContainerPortArgs{
    								ContainerPort: pulumi.Int(80),
    							},
    						},
    					},
    				},
    			},
    		})
    		if err != nil {
    			return err
    		}
    		return nil
    	})
    }
    ```
    ```java
    package generated_program;

    import com.pulumi.Context;
    import com.pulumi.Pulumi;
    import com.pulumi.core.Output;
    import com.pulumi.kubernetes.core_v1.Pod;
    import com.pulumi.kubernetes.core_v1.PodArgs;
    import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
    import java.util.List;
    import java.util.ArrayList;
    import java.util.Map;
    import java.io.File;
    import java.nio.file.Files;
    import java.nio.file.Paths;

    public class App {
        public static void main(String[] args) {
            Pulumi.run(App::stack);
        }

        public static void stack(Context ctx) {
            var pod = new Pod("pod", PodArgs.builder()        
                .spec(PodSpecArgs.builder()
                    .containers(ContainerArgs.builder()
                        .image("nginx:1.14.2")
                        .name("nginx")
                        .ports(ContainerPortArgs.builder()
                            .containerPort(80)
                            .build())
                        .build())
                    .build())
                .build());

        }
    }
    ```
    ```yaml
    description: Create a Pod with auto-naming
    name: yaml-example
    resources:
        pod:
            properties:
                spec:
                    containers:
                        - image: nginx:1.14.2
                          name: nginx
                          ports:
                            - containerPort: 80
            type: kubernetes:core/v1:Pod
    runtime: yaml
    ```
    {{% /example %}}
    {{% example %}}
    ### Create a Pod with a user-specified name

    ```typescript
    import * as pulumi from "@pulumi/pulumi";
    import * as kubernetes from "@pulumi/kubernetes";

    const pod = new kubernetes.core.v1.Pod("pod", {
        metadata: {
            name: "nginx",
        },
        spec: {
            containers: [{
                image: "nginx:1.14.2",
                name: "nginx",
                ports: [{
                    containerPort: 80,
                }],
            }],
        },
    });
    ```
    ```python
    import pulumi
    import pulumi_kubernetes as kubernetes

    pod = kubernetes.core.v1.Pod("pod",
        metadata=kubernetes.meta.v1.ObjectMetaArgs(
            name="nginx",
        ),
        spec=kubernetes.core.v1.PodSpecArgs(
            containers=[kubernetes.core.v1.ContainerArgs(
                image="nginx:1.14.2",
                name="nginx",
                ports=[kubernetes.core.v1.ContainerPortArgs(
                    container_port=80,
                )],
            )],
        ))
    ```
    ```csharp
    using System.Collections.Generic;
    using System.Linq;
    using Pulumi;
    using Kubernetes = Pulumi.Kubernetes;

    return await Deployment.RunAsync(() => 
    {
        var pod = new Kubernetes.Core.V1.Pod("pod", new()
        {
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "nginx",
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {
                Containers = new[]
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Image = "nginx:1.14.2",
                        Name = "nginx",
                        Ports = new[]
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerPortArgs
                            {
                                ContainerPortValue = 80,
                            },
                        },
                    },
                },
            },
        });

    });

    ```
    ```go
    package main

    import (
    	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
    	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
    	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    )

    func main() {
    	pulumi.Run(func(ctx *pulumi.Context) error {
    		_, err := corev1.NewPod(ctx, "pod", &corev1.PodArgs{
    			Metadata: &metav1.ObjectMetaArgs{
    				Name: pulumi.String("nginx"),
    			},
    			Spec: &corev1.PodSpecArgs{
    				Containers: corev1.ContainerArray{
    					&corev1.ContainerArgs{
    						Image: pulumi.String("nginx:1.14.2"),
    						Name:  pulumi.String("nginx"),
    						Ports: corev1.ContainerPortArray{
    							&corev1.ContainerPortArgs{
    								ContainerPort: pulumi.Int(80),
    							},
    						},
    					},
    				},
    			},
    		})
    		if err != nil {
    			return err
    		}
    		return nil
    	})
    }
    ```
    ```java
    package generated_program;

    import com.pulumi.Context;
    import com.pulumi.Pulumi;
    import com.pulumi.core.Output;
    import com.pulumi.kubernetes.core_v1.Pod;
    import com.pulumi.kubernetes.core_v1.PodArgs;
    import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
    import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
    import java.util.List;
    import java.util.ArrayList;
    import java.util.Map;
    import java.io.File;
    import java.nio.file.Files;
    import java.nio.file.Paths;

    public class App {
        public static void main(String[] args) {
            Pulumi.run(App::stack);
        }

        public static void stack(Context ctx) {
            var pod = new Pod("pod", PodArgs.builder()        
                .metadata(ObjectMetaArgs.builder()
                    .name("nginx")
                    .build())
                .spec(PodSpecArgs.builder()
                    .containers(ContainerArgs.builder()
                        .image("nginx:1.14.2")
                        .name("nginx")
                        .ports(ContainerPortArgs.builder()
                            .containerPort(80)
                            .build())
                        .build())
                    .build())
                .build());

        }
    }
    ```
    ```yaml
    description: Create a Pod with a user-specified name
    name: yaml-example
    resources:
        pod:
            properties:
                metadata:
                    name: nginx
                spec:
                    containers:
                        - image: nginx:1.14.2
                          name: nginx
                          ports:
                            - containerPort: 80
            type: kubernetes:core/v1:Pod
    runtime: yaml
    ```
    {{% /example %}}
    {{% /examples %}}
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PodManifest. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PodManifest.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PodManifest.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 kind: Optional[str] = None,
                 metadata: Optional['pulumi_kubernetes.meta.v1.outputs.ObjectMeta'] = None,
                 spec: Optional['pulumi_kubernetes.core.v1.outputs.PodSpec'] = None,
                 status: Optional['pulumi_kubernetes.core.v1.outputs.PodStatus'] = None):
        """
        Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.

        This resource waits until its status is ready before registering success
        for create/update, and populating output properties from the current state of the resource.
        The following conditions are used to determine whether the resource creation has
        succeeded or failed:

        1. The Pod is scheduled ("PodScheduled"" '.status.condition' is true).
        2. The Pod is initialized ("Initialized" '.status.condition' is true).
        3. The Pod is ready ("Ready" '.status.condition' is true) and the '.status.phase' is
           set to "Running".
        Or (for Jobs): The Pod succeeded ('.status.phase' set to "Succeeded").

        If the Pod has not reached a Ready state after 10 minutes, it will
        time out and mark the resource update as Failed. You can override the default timeout value
        by setting the 'customTimeouts' option on the resource.

        {{% examples %}}
        ## Example Usage
        {{% example %}}
        ### Create a Pod with auto-naming

        ```typescript
        import * as pulumi from "@pulumi/pulumi";
        import * as kubernetes from "@pulumi/kubernetes";

        const pod = new kubernetes.core.v1.Pod("pod", {spec: {
            containers: [{
                image: "nginx:1.14.2",
                name: "nginx",
                ports: [{
                    containerPort: 80,
                }],
            }],
        }});
        ```
        ```python
        import pulumi
        import pulumi_kubernetes as kubernetes

        pod = kubernetes.core.v1.Pod("pod", spec=kubernetes.core.v1.PodSpecArgs(
            containers=[kubernetes.core.v1.ContainerArgs(
                image="nginx:1.14.2",
                name="nginx",
                ports=[kubernetes.core.v1.ContainerPortArgs(
                    container_port=80,
                )],
            )],
        ))
        ```
        ```csharp
        using System.Collections.Generic;
        using System.Linq;
        using Pulumi;
        using Kubernetes = Pulumi.Kubernetes;

        return await Deployment.RunAsync(() => 
        {
            var pod = new Kubernetes.Core.V1.Pod("pod", new()
            {
                Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                {
                    Containers = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Image = "nginx:1.14.2",
                            Name = "nginx",
                            Ports = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.ContainerPortArgs
                                {
                                    ContainerPortValue = 80,
                                },
                            },
                        },
                    },
                },
            });

        });

        ```
        ```go
        package main

        import (
        	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
        	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
        )

        func main() {
        	pulumi.Run(func(ctx *pulumi.Context) error {
        		_, err := corev1.NewPod(ctx, "pod", &corev1.PodArgs{
        			Spec: &corev1.PodSpecArgs{
        				Containers: corev1.ContainerArray{
        					&corev1.ContainerArgs{
        						Image: pulumi.String("nginx:1.14.2"),
        						Name:  pulumi.String("nginx"),
        						Ports: corev1.ContainerPortArray{
        							&corev1.ContainerPortArgs{
        								ContainerPort: pulumi.Int(80),
        							},
        						},
        					},
        				},
        			},
        		})
        		if err != nil {
        			return err
        		}
        		return nil
        	})
        }
        ```
        ```java
        package generated_program;

        import com.pulumi.Context;
        import com.pulumi.Pulumi;
        import com.pulumi.core.Output;
        import com.pulumi.kubernetes.core_v1.Pod;
        import com.pulumi.kubernetes.core_v1.PodArgs;
        import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
        import java.util.List;
        import java.util.ArrayList;
        import java.util.Map;
        import java.io.File;
        import java.nio.file.Files;
        import java.nio.file.Paths;

        public class App {
            public static void main(String[] args) {
                Pulumi.run(App::stack);
            }

            public static void stack(Context ctx) {
                var pod = new Pod("pod", PodArgs.builder()        
                    .spec(PodSpecArgs.builder()
                        .containers(ContainerArgs.builder()
                            .image("nginx:1.14.2")
                            .name("nginx")
                            .ports(ContainerPortArgs.builder()
                                .containerPort(80)
                                .build())
                            .build())
                        .build())
                    .build());

            }
        }
        ```
        ```yaml
        description: Create a Pod with auto-naming
        name: yaml-example
        resources:
            pod:
                properties:
                    spec:
                        containers:
                            - image: nginx:1.14.2
                              name: nginx
                              ports:
                                - containerPort: 80
                type: kubernetes:core/v1:Pod
        runtime: yaml
        ```
        {{% /example %}}
        {{% example %}}
        ### Create a Pod with a user-specified name

        ```typescript
        import * as pulumi from "@pulumi/pulumi";
        import * as kubernetes from "@pulumi/kubernetes";

        const pod = new kubernetes.core.v1.Pod("pod", {
            metadata: {
                name: "nginx",
            },
            spec: {
                containers: [{
                    image: "nginx:1.14.2",
                    name: "nginx",
                    ports: [{
                        containerPort: 80,
                    }],
                }],
            },
        });
        ```
        ```python
        import pulumi
        import pulumi_kubernetes as kubernetes

        pod = kubernetes.core.v1.Pod("pod",
            metadata=kubernetes.meta.v1.ObjectMetaArgs(
                name="nginx",
            ),
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    image="nginx:1.14.2",
                    name="nginx",
                    ports=[kubernetes.core.v1.ContainerPortArgs(
                        container_port=80,
                    )],
                )],
            ))
        ```
        ```csharp
        using System.Collections.Generic;
        using System.Linq;
        using Pulumi;
        using Kubernetes = Pulumi.Kubernetes;

        return await Deployment.RunAsync(() => 
        {
            var pod = new Kubernetes.Core.V1.Pod("pod", new()
            {
                Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
                {
                    Name = "nginx",
                },
                Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                {
                    Containers = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Image = "nginx:1.14.2",
                            Name = "nginx",
                            Ports = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.ContainerPortArgs
                                {
                                    ContainerPortValue = 80,
                                },
                            },
                        },
                    },
                },
            });

        });

        ```
        ```go
        package main

        import (
        	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
        	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
        	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
        )

        func main() {
        	pulumi.Run(func(ctx *pulumi.Context) error {
        		_, err := corev1.NewPod(ctx, "pod", &corev1.PodArgs{
        			Metadata: &metav1.ObjectMetaArgs{
        				Name: pulumi.String("nginx"),
        			},
        			Spec: &corev1.PodSpecArgs{
        				Containers: corev1.ContainerArray{
        					&corev1.ContainerArgs{
        						Image: pulumi.String("nginx:1.14.2"),
        						Name:  pulumi.String("nginx"),
        						Ports: corev1.ContainerPortArray{
        							&corev1.ContainerPortArgs{
        								ContainerPort: pulumi.Int(80),
        							},
        						},
        					},
        				},
        			},
        		})
        		if err != nil {
        			return err
        		}
        		return nil
        	})
        }
        ```
        ```java
        package generated_program;

        import com.pulumi.Context;
        import com.pulumi.Pulumi;
        import com.pulumi.core.Output;
        import com.pulumi.kubernetes.core_v1.Pod;
        import com.pulumi.kubernetes.core_v1.PodArgs;
        import com.pulumi.kubernetes.meta_v1.inputs.ObjectMetaArgs;
        import com.pulumi.kubernetes.core_v1.inputs.PodSpecArgs;
        import java.util.List;
        import java.util.ArrayList;
        import java.util.Map;
        import java.io.File;
        import java.nio.file.Files;
        import java.nio.file.Paths;

        public class App {
            public static void main(String[] args) {
                Pulumi.run(App::stack);
            }

            public static void stack(Context ctx) {
                var pod = new Pod("pod", PodArgs.builder()        
                    .metadata(ObjectMetaArgs.builder()
                        .name("nginx")
                        .build())
                    .spec(PodSpecArgs.builder()
                        .containers(ContainerArgs.builder()
                            .image("nginx:1.14.2")
                            .name("nginx")
                            .ports(ContainerPortArgs.builder()
                                .containerPort(80)
                                .build())
                            .build())
                        .build())
                    .build());

            }
        }
        ```
        ```yaml
        description: Create a Pod with a user-specified name
        name: yaml-example
        resources:
            pod:
                properties:
                    metadata:
                        name: nginx
                    spec:
                        containers:
                            - image: nginx:1.14.2
                              name: nginx
                              ports:
                                - containerPort: 80
                type: kubernetes:core/v1:Pod
        runtime: yaml
        ```
        {{% /example %}}
        {{% /examples %}}

        :param str api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param str kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param 'pulumi_kubernetes.meta.v1.ObjectMetaArgs' metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param 'pulumi_kubernetes.core.v1.PodSpecArgs' spec: Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        :param 'pulumi_kubernetes.core.v1.PodStatusArgs' status: Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'v1')
        if kind is not None:
            pulumi.set(__self__, "kind", 'Pod')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['pulumi_kubernetes.meta.v1.outputs.ObjectMeta']:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> Optional['pulumi_kubernetes.core.v1.outputs.PodSpec']:
        """
        Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> Optional['pulumi_kubernetes.core.v1.outputs.PodStatus']:
        """
        Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class User(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientCertificateData":
            suggest = "client_certificate_data"
        elif key == "clientKeyData":
            suggest = "client_key_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in User. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        User.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        User.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_certificate_data: str,
                 client_key_data: str):
        """
        :param str client_certificate_data: TODO
        :param str client_key_data: TODO
        """
        pulumi.set(__self__, "client_certificate_data", client_certificate_data)
        pulumi.set(__self__, "client_key_data", client_key_data)

    @property
    @pulumi.getter(name="clientCertificateData")
    def client_certificate_data(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "client_certificate_data")

    @property
    @pulumi.getter(name="clientKeyData")
    def client_key_data(self) -> str:
        """
        TODO
        """
        return pulumi.get(self, "client_key_data")


