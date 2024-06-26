# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import tools as _tools
from ._enums import *
from .file import File
from .kube_api_server_install import KubeApiServerInstall
from .kube_controller_manager_install import KubeControllerManagerInstall
from .kube_scheduler_install import KubeSchedulerInstall
from .kubectl_install import KubectlInstall
from .systemd_service import SystemdService
import pulumi_command

__all__ = ['ControlPlaneNodeArgs', 'ControlPlaneNode']

@pulumi.input_type
class ControlPlaneNodeArgs:
    def __init__(__self__, *,
                 api_server_count: pulumi.Input[int],
                 architecture: pulumi.Input['Architecture'],
                 ca_certificate_path: pulumi.Input[str],
                 ca_private_key_path: pulumi.Input[str],
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 encryption_config_yaml: pulumi.Input[str],
                 kube_api_server_certificate_path: pulumi.Input[str],
                 kube_api_server_private_key_path: pulumi.Input[str],
                 kube_controller_manager_kubeconfig_path: pulumi.Input[str],
                 kube_scheduler_config_yaml: pulumi.Input[str],
                 kube_scheduler_kubeconfig_path: pulumi.Input[str],
                 service_accounts_certificate_path: pulumi.Input[str],
                 service_accounts_private_key_path: pulumi.Input[str],
                 audi_log_path: Optional[pulumi.Input[str]] = None,
                 cluster_cidr: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 kube_api_server_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_controller_manager_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_scheduler_install_directory: Optional[pulumi.Input[str]] = None,
                 kubectl_install_directory: Optional[pulumi.Input[str]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 node_name: Optional[pulumi.Input[str]] = None,
                 service_cluster_ip_range: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ControlPlaneNode resource.
        :param pulumi.Input[int] api_server_count: The number of kube-apiserver instance.
        :param pulumi.Input['Architecture'] architecture: The node's CPU architecture.
        :param pulumi.Input[str] ca_certificate_path: The path to the root certificate authority certificate.
        :param pulumi.Input[str] ca_private_key_path: The path to the root certificate authority private key.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input[str] encryption_config_yaml: The v1/EncryptionConfig yaml.
        :param pulumi.Input[str] kube_api_server_certificate_path: The path to the kube-apiserver certificate.
        :param pulumi.Input[str] kube_api_server_private_key_path: The path to the kube-apiserver private key.
        :param pulumi.Input[str] kube_controller_manager_kubeconfig_path: The path to the kube-controller-manager kubeconfig file.
        :param pulumi.Input[str] kube_scheduler_config_yaml: The kube-scheduler config yaml.
        :param pulumi.Input[str] kube_scheduler_kubeconfig_path: The path to the kube-scheduler kubeconfig file.
        :param pulumi.Input[str] service_accounts_certificate_path: The path to the service accounts certificate.
        :param pulumi.Input[str] service_accounts_private_key_path: The path to the service accounts private key.
        :param pulumi.Input[str] audi_log_path: The path to store the audit log file.
        :param pulumi.Input[str] cluster_cidr: The cluster CIDR.
        :param pulumi.Input[str] cluster_name: The cluster name.
        :param pulumi.Input[str] kube_api_server_install_directory: The directory to store the kube-apiserver binary.
        :param pulumi.Input[str] kube_controller_manager_install_directory: The directory to store the kube-controller-manager binary.
        :param pulumi.Input[str] kube_scheduler_install_directory: The directory to store the kube-scheduler binary.
        :param pulumi.Input[str] kubectl_install_directory: The path to store the kubectl binary.
        :param pulumi.Input[str] kubernetes_version: The version of kubernetes to use.
        :param pulumi.Input[str] node_name: The name of the node.
        :param pulumi.Input[str] service_cluster_ip_range: The IP range to use for cluster services.
        """
        pulumi.set(__self__, "api_server_count", api_server_count)
        pulumi.set(__self__, "architecture", architecture)
        pulumi.set(__self__, "ca_certificate_path", ca_certificate_path)
        pulumi.set(__self__, "ca_private_key_path", ca_private_key_path)
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "encryption_config_yaml", encryption_config_yaml)
        pulumi.set(__self__, "kube_api_server_certificate_path", kube_api_server_certificate_path)
        pulumi.set(__self__, "kube_api_server_private_key_path", kube_api_server_private_key_path)
        pulumi.set(__self__, "kube_controller_manager_kubeconfig_path", kube_controller_manager_kubeconfig_path)
        pulumi.set(__self__, "kube_scheduler_config_yaml", kube_scheduler_config_yaml)
        pulumi.set(__self__, "kube_scheduler_kubeconfig_path", kube_scheduler_kubeconfig_path)
        pulumi.set(__self__, "service_accounts_certificate_path", service_accounts_certificate_path)
        pulumi.set(__self__, "service_accounts_private_key_path", service_accounts_private_key_path)
        if audi_log_path is not None:
            pulumi.set(__self__, "audi_log_path", audi_log_path)
        if cluster_cidr is not None:
            pulumi.set(__self__, "cluster_cidr", cluster_cidr)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if kube_api_server_install_directory is not None:
            pulumi.set(__self__, "kube_api_server_install_directory", kube_api_server_install_directory)
        if kube_controller_manager_install_directory is not None:
            pulumi.set(__self__, "kube_controller_manager_install_directory", kube_controller_manager_install_directory)
        if kube_scheduler_install_directory is not None:
            pulumi.set(__self__, "kube_scheduler_install_directory", kube_scheduler_install_directory)
        if kubectl_install_directory is not None:
            pulumi.set(__self__, "kubectl_install_directory", kubectl_install_directory)
        if kubernetes_version is not None:
            pulumi.set(__self__, "kubernetes_version", kubernetes_version)
        if node_name is not None:
            pulumi.set(__self__, "node_name", node_name)
        if service_cluster_ip_range is not None:
            pulumi.set(__self__, "service_cluster_ip_range", service_cluster_ip_range)

    @property
    @pulumi.getter(name="apiServerCount")
    def api_server_count(self) -> pulumi.Input[int]:
        """
        The number of kube-apiserver instance.
        """
        return pulumi.get(self, "api_server_count")

    @api_server_count.setter
    def api_server_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "api_server_count", value)

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Input['Architecture']:
        """
        The node's CPU architecture.
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: pulumi.Input['Architecture']):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter(name="caCertificatePath")
    def ca_certificate_path(self) -> pulumi.Input[str]:
        """
        The path to the root certificate authority certificate.
        """
        return pulumi.get(self, "ca_certificate_path")

    @ca_certificate_path.setter
    def ca_certificate_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_certificate_path", value)

    @property
    @pulumi.getter(name="caPrivateKeyPath")
    def ca_private_key_path(self) -> pulumi.Input[str]:
        """
        The path to the root certificate authority private key.
        """
        return pulumi.get(self, "ca_private_key_path")

    @ca_private_key_path.setter
    def ca_private_key_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_private_key_path", value)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Input['pulumi_command.remote.ConnectionArgs']:
        """
        The parameters with which to connect to the remote host.
        """
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: pulumi.Input['pulumi_command.remote.ConnectionArgs']):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter(name="encryptionConfigYaml")
    def encryption_config_yaml(self) -> pulumi.Input[str]:
        """
        The v1/EncryptionConfig yaml.
        """
        return pulumi.get(self, "encryption_config_yaml")

    @encryption_config_yaml.setter
    def encryption_config_yaml(self, value: pulumi.Input[str]):
        pulumi.set(self, "encryption_config_yaml", value)

    @property
    @pulumi.getter(name="kubeApiServerCertificatePath")
    def kube_api_server_certificate_path(self) -> pulumi.Input[str]:
        """
        The path to the kube-apiserver certificate.
        """
        return pulumi.get(self, "kube_api_server_certificate_path")

    @kube_api_server_certificate_path.setter
    def kube_api_server_certificate_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_api_server_certificate_path", value)

    @property
    @pulumi.getter(name="kubeApiServerPrivateKeyPath")
    def kube_api_server_private_key_path(self) -> pulumi.Input[str]:
        """
        The path to the kube-apiserver private key.
        """
        return pulumi.get(self, "kube_api_server_private_key_path")

    @kube_api_server_private_key_path.setter
    def kube_api_server_private_key_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_api_server_private_key_path", value)

    @property
    @pulumi.getter(name="kubeControllerManagerKubeconfigPath")
    def kube_controller_manager_kubeconfig_path(self) -> pulumi.Input[str]:
        """
        The path to the kube-controller-manager kubeconfig file.
        """
        return pulumi.get(self, "kube_controller_manager_kubeconfig_path")

    @kube_controller_manager_kubeconfig_path.setter
    def kube_controller_manager_kubeconfig_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_controller_manager_kubeconfig_path", value)

    @property
    @pulumi.getter(name="kubeSchedulerConfigYaml")
    def kube_scheduler_config_yaml(self) -> pulumi.Input[str]:
        """
        The kube-scheduler config yaml.
        """
        return pulumi.get(self, "kube_scheduler_config_yaml")

    @kube_scheduler_config_yaml.setter
    def kube_scheduler_config_yaml(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_scheduler_config_yaml", value)

    @property
    @pulumi.getter(name="kubeSchedulerKubeconfigPath")
    def kube_scheduler_kubeconfig_path(self) -> pulumi.Input[str]:
        """
        The path to the kube-scheduler kubeconfig file.
        """
        return pulumi.get(self, "kube_scheduler_kubeconfig_path")

    @kube_scheduler_kubeconfig_path.setter
    def kube_scheduler_kubeconfig_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_scheduler_kubeconfig_path", value)

    @property
    @pulumi.getter(name="serviceAccountsCertificatePath")
    def service_accounts_certificate_path(self) -> pulumi.Input[str]:
        """
        The path to the service accounts certificate.
        """
        return pulumi.get(self, "service_accounts_certificate_path")

    @service_accounts_certificate_path.setter
    def service_accounts_certificate_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_accounts_certificate_path", value)

    @property
    @pulumi.getter(name="serviceAccountsPrivateKeyPath")
    def service_accounts_private_key_path(self) -> pulumi.Input[str]:
        """
        The path to the service accounts private key.
        """
        return pulumi.get(self, "service_accounts_private_key_path")

    @service_accounts_private_key_path.setter
    def service_accounts_private_key_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_accounts_private_key_path", value)

    @property
    @pulumi.getter(name="audiLogPath")
    def audi_log_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to store the audit log file.
        """
        return pulumi.get(self, "audi_log_path")

    @audi_log_path.setter
    def audi_log_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audi_log_path", value)

    @property
    @pulumi.getter(name="clusterCIDR")
    def cluster_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster CIDR.
        """
        return pulumi.get(self, "cluster_cidr")

    @cluster_cidr.setter
    def cluster_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_cidr", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="kubeApiServerInstallDirectory")
    def kube_api_server_install_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to store the kube-apiserver binary.
        """
        return pulumi.get(self, "kube_api_server_install_directory")

    @kube_api_server_install_directory.setter
    def kube_api_server_install_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kube_api_server_install_directory", value)

    @property
    @pulumi.getter(name="kubeControllerManagerInstallDirectory")
    def kube_controller_manager_install_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to store the kube-controller-manager binary.
        """
        return pulumi.get(self, "kube_controller_manager_install_directory")

    @kube_controller_manager_install_directory.setter
    def kube_controller_manager_install_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kube_controller_manager_install_directory", value)

    @property
    @pulumi.getter(name="kubeSchedulerInstallDirectory")
    def kube_scheduler_install_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to store the kube-scheduler binary.
        """
        return pulumi.get(self, "kube_scheduler_install_directory")

    @kube_scheduler_install_directory.setter
    def kube_scheduler_install_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kube_scheduler_install_directory", value)

    @property
    @pulumi.getter(name="kubectlInstallDirectory")
    def kubectl_install_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The path to store the kubectl binary.
        """
        return pulumi.get(self, "kubectl_install_directory")

    @kubectl_install_directory.setter
    def kubectl_install_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubectl_install_directory", value)

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of kubernetes to use.
        """
        return pulumi.get(self, "kubernetes_version")

    @kubernetes_version.setter
    def kubernetes_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_version", value)

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the node.
        """
        return pulumi.get(self, "node_name")

    @node_name.setter
    def node_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_name", value)

    @property
    @pulumi.getter(name="serviceClusterIpRange")
    def service_cluster_ip_range(self) -> Optional[pulumi.Input[str]]:
        """
        The IP range to use for cluster services.
        """
        return pulumi.get(self, "service_cluster_ip_range")

    @service_cluster_ip_range.setter
    def service_cluster_ip_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_cluster_ip_range", value)


class ControlPlaneNode(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_server_count: Optional[pulumi.Input[int]] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 audi_log_path: Optional[pulumi.Input[str]] = None,
                 ca_certificate_path: Optional[pulumi.Input[str]] = None,
                 ca_private_key_path: Optional[pulumi.Input[str]] = None,
                 cluster_cidr: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 encryption_config_yaml: Optional[pulumi.Input[str]] = None,
                 kube_api_server_certificate_path: Optional[pulumi.Input[str]] = None,
                 kube_api_server_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_api_server_private_key_path: Optional[pulumi.Input[str]] = None,
                 kube_controller_manager_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_controller_manager_kubeconfig_path: Optional[pulumi.Input[str]] = None,
                 kube_scheduler_config_yaml: Optional[pulumi.Input[str]] = None,
                 kube_scheduler_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_scheduler_kubeconfig_path: Optional[pulumi.Input[str]] = None,
                 kubectl_install_directory: Optional[pulumi.Input[str]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 node_name: Optional[pulumi.Input[str]] = None,
                 service_accounts_certificate_path: Optional[pulumi.Input[str]] = None,
                 service_accounts_private_key_path: Optional[pulumi.Input[str]] = None,
                 service_cluster_ip_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A kubernetes control plane node.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] api_server_count: The number of kube-apiserver instance.
        :param pulumi.Input['Architecture'] architecture: The node's CPU architecture.
        :param pulumi.Input[str] audi_log_path: The path to store the audit log file.
        :param pulumi.Input[str] ca_certificate_path: The path to the root certificate authority certificate.
        :param pulumi.Input[str] ca_private_key_path: The path to the root certificate authority private key.
        :param pulumi.Input[str] cluster_cidr: The cluster CIDR.
        :param pulumi.Input[str] cluster_name: The cluster name.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input[str] encryption_config_yaml: The v1/EncryptionConfig yaml.
        :param pulumi.Input[str] kube_api_server_certificate_path: The path to the kube-apiserver certificate.
        :param pulumi.Input[str] kube_api_server_install_directory: The directory to store the kube-apiserver binary.
        :param pulumi.Input[str] kube_api_server_private_key_path: The path to the kube-apiserver private key.
        :param pulumi.Input[str] kube_controller_manager_install_directory: The directory to store the kube-controller-manager binary.
        :param pulumi.Input[str] kube_controller_manager_kubeconfig_path: The path to the kube-controller-manager kubeconfig file.
        :param pulumi.Input[str] kube_scheduler_config_yaml: The kube-scheduler config yaml.
        :param pulumi.Input[str] kube_scheduler_install_directory: The directory to store the kube-scheduler binary.
        :param pulumi.Input[str] kube_scheduler_kubeconfig_path: The path to the kube-scheduler kubeconfig file.
        :param pulumi.Input[str] kubectl_install_directory: The path to store the kubectl binary.
        :param pulumi.Input[str] kubernetes_version: The version of kubernetes to use.
        :param pulumi.Input[str] node_name: The name of the node.
        :param pulumi.Input[str] service_accounts_certificate_path: The path to the service accounts certificate.
        :param pulumi.Input[str] service_accounts_private_key_path: The path to the service accounts private key.
        :param pulumi.Input[str] service_cluster_ip_range: The IP range to use for cluster services.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ControlPlaneNodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A kubernetes control plane node.

        :param str resource_name: The name of the resource.
        :param ControlPlaneNodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ControlPlaneNodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_server_count: Optional[pulumi.Input[int]] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 audi_log_path: Optional[pulumi.Input[str]] = None,
                 ca_certificate_path: Optional[pulumi.Input[str]] = None,
                 ca_private_key_path: Optional[pulumi.Input[str]] = None,
                 cluster_cidr: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 encryption_config_yaml: Optional[pulumi.Input[str]] = None,
                 kube_api_server_certificate_path: Optional[pulumi.Input[str]] = None,
                 kube_api_server_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_api_server_private_key_path: Optional[pulumi.Input[str]] = None,
                 kube_controller_manager_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_controller_manager_kubeconfig_path: Optional[pulumi.Input[str]] = None,
                 kube_scheduler_config_yaml: Optional[pulumi.Input[str]] = None,
                 kube_scheduler_install_directory: Optional[pulumi.Input[str]] = None,
                 kube_scheduler_kubeconfig_path: Optional[pulumi.Input[str]] = None,
                 kubectl_install_directory: Optional[pulumi.Input[str]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 node_name: Optional[pulumi.Input[str]] = None,
                 service_accounts_certificate_path: Optional[pulumi.Input[str]] = None,
                 service_accounts_private_key_path: Optional[pulumi.Input[str]] = None,
                 service_cluster_ip_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ControlPlaneNodeArgs.__new__(ControlPlaneNodeArgs)

            if api_server_count is None and not opts.urn:
                raise TypeError("Missing required property 'api_server_count'")
            __props__.__dict__["api_server_count"] = api_server_count
            if architecture is None and not opts.urn:
                raise TypeError("Missing required property 'architecture'")
            __props__.__dict__["architecture"] = architecture
            __props__.__dict__["audi_log_path"] = audi_log_path
            if ca_certificate_path is None and not opts.urn:
                raise TypeError("Missing required property 'ca_certificate_path'")
            __props__.__dict__["ca_certificate_path"] = ca_certificate_path
            if ca_private_key_path is None and not opts.urn:
                raise TypeError("Missing required property 'ca_private_key_path'")
            __props__.__dict__["ca_private_key_path"] = ca_private_key_path
            __props__.__dict__["cluster_cidr"] = cluster_cidr
            __props__.__dict__["cluster_name"] = cluster_name
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            if encryption_config_yaml is None and not opts.urn:
                raise TypeError("Missing required property 'encryption_config_yaml'")
            __props__.__dict__["encryption_config_yaml"] = encryption_config_yaml
            if kube_api_server_certificate_path is None and not opts.urn:
                raise TypeError("Missing required property 'kube_api_server_certificate_path'")
            __props__.__dict__["kube_api_server_certificate_path"] = kube_api_server_certificate_path
            __props__.__dict__["kube_api_server_install_directory"] = kube_api_server_install_directory
            if kube_api_server_private_key_path is None and not opts.urn:
                raise TypeError("Missing required property 'kube_api_server_private_key_path'")
            __props__.__dict__["kube_api_server_private_key_path"] = kube_api_server_private_key_path
            __props__.__dict__["kube_controller_manager_install_directory"] = kube_controller_manager_install_directory
            if kube_controller_manager_kubeconfig_path is None and not opts.urn:
                raise TypeError("Missing required property 'kube_controller_manager_kubeconfig_path'")
            __props__.__dict__["kube_controller_manager_kubeconfig_path"] = kube_controller_manager_kubeconfig_path
            if kube_scheduler_config_yaml is None and not opts.urn:
                raise TypeError("Missing required property 'kube_scheduler_config_yaml'")
            __props__.__dict__["kube_scheduler_config_yaml"] = kube_scheduler_config_yaml
            __props__.__dict__["kube_scheduler_install_directory"] = kube_scheduler_install_directory
            if kube_scheduler_kubeconfig_path is None and not opts.urn:
                raise TypeError("Missing required property 'kube_scheduler_kubeconfig_path'")
            __props__.__dict__["kube_scheduler_kubeconfig_path"] = kube_scheduler_kubeconfig_path
            __props__.__dict__["kubectl_install_directory"] = kubectl_install_directory
            __props__.__dict__["kubernetes_version"] = kubernetes_version
            __props__.__dict__["node_name"] = node_name
            if service_accounts_certificate_path is None and not opts.urn:
                raise TypeError("Missing required property 'service_accounts_certificate_path'")
            __props__.__dict__["service_accounts_certificate_path"] = service_accounts_certificate_path
            if service_accounts_private_key_path is None and not opts.urn:
                raise TypeError("Missing required property 'service_accounts_private_key_path'")
            __props__.__dict__["service_accounts_private_key_path"] = service_accounts_private_key_path
            __props__.__dict__["service_cluster_ip_range"] = service_cluster_ip_range
            __props__.__dict__["encryption_config_file"] = None
            __props__.__dict__["kube_api_server_install"] = None
            __props__.__dict__["kube_api_server_service"] = None
            __props__.__dict__["kube_controller_manager_install"] = None
            __props__.__dict__["kube_controller_manager_service"] = None
            __props__.__dict__["kube_scheduler_install"] = None
            __props__.__dict__["kube_scheduler_service"] = None
            __props__.__dict__["kubectl_install"] = None
            __props__.__dict__["kubernetes_configuration_mkdir"] = None
            __props__.__dict__["var_lib_kubernetes_mkdir"] = None
        super(ControlPlaneNode, __self__).__init__(
            'kubernetes-the-hard-way:remote:ControlPlaneNode',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter(name="apiServerCount")
    def api_server_count(self) -> pulumi.Output[int]:
        """
        The number of kube-apiserver instance.
        """
        return pulumi.get(self, "api_server_count")

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Output['Architecture']:
        """
        The node's CPU architecture.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="audiLogPath")
    def audi_log_path(self) -> pulumi.Output[Optional[str]]:
        """
        The path to store the audit log file.
        """
        return pulumi.get(self, "audi_log_path")

    @property
    @pulumi.getter(name="caCertificatePath")
    def ca_certificate_path(self) -> pulumi.Output[str]:
        """
        The path to the root certificate authority certificate.
        """
        return pulumi.get(self, "ca_certificate_path")

    @property
    @pulumi.getter(name="caPrivateKeyPath")
    def ca_private_key_path(self) -> pulumi.Output[str]:
        """
        The path to the root certificate authority private key.
        """
        return pulumi.get(self, "ca_private_key_path")

    @property
    @pulumi.getter(name="clusterCIDR")
    def cluster_cidr(self) -> pulumi.Output[Optional[str]]:
        """
        The cluster CIDR.
        """
        return pulumi.get(self, "cluster_cidr")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[Optional[str]]:
        """
        The cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        The parameters with which to connect to the remote host.
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter(name="encryptionConfigFile")
    def encryption_config_file(self) -> pulumi.Output[Optional['File']]:
        """
        The remote encryption config file.
        """
        return pulumi.get(self, "encryption_config_file")

    @property
    @pulumi.getter(name="encryptionConfigYaml")
    def encryption_config_yaml(self) -> pulumi.Output[str]:
        """
        The v1/EncryptionConfig yaml.
        """
        return pulumi.get(self, "encryption_config_yaml")

    @property
    @pulumi.getter(name="kubeApiServerCertificatePath")
    def kube_api_server_certificate_path(self) -> pulumi.Output[str]:
        """
        The path to the kube-apiserver certificate.
        """
        return pulumi.get(self, "kube_api_server_certificate_path")

    @property
    @pulumi.getter(name="kubeApiServerInstall")
    def kube_api_server_install(self) -> pulumi.Output['KubeApiServerInstall']:
        """
        The kube-apiserver install.
        """
        return pulumi.get(self, "kube_api_server_install")

    @property
    @pulumi.getter(name="kubeApiServerInstallDirectory")
    def kube_api_server_install_directory(self) -> pulumi.Output[Optional[str]]:
        """
        The directory to store the kube-apiserver binary.
        """
        return pulumi.get(self, "kube_api_server_install_directory")

    @property
    @pulumi.getter(name="kubeApiServerPrivateKeyPath")
    def kube_api_server_private_key_path(self) -> pulumi.Output[str]:
        """
        The path to the kube-apiserver private key.
        """
        return pulumi.get(self, "kube_api_server_private_key_path")

    @property
    @pulumi.getter(name="kubeApiServerService")
    def kube_api_server_service(self) -> pulumi.Output[Optional['SystemdService']]:
        """
        The kube-apiserver systemd service.
        """
        return pulumi.get(self, "kube_api_server_service")

    @property
    @pulumi.getter(name="kubeControllerManagerInstall")
    def kube_controller_manager_install(self) -> pulumi.Output['KubeControllerManagerInstall']:
        """
        The kube-controller-manager install.
        """
        return pulumi.get(self, "kube_controller_manager_install")

    @property
    @pulumi.getter(name="kubeControllerManagerInstallDirectory")
    def kube_controller_manager_install_directory(self) -> pulumi.Output[Optional[str]]:
        """
        The directory to store the kube-controller-manager binary.
        """
        return pulumi.get(self, "kube_controller_manager_install_directory")

    @property
    @pulumi.getter(name="kubeControllerManagerKubeconfigPath")
    def kube_controller_manager_kubeconfig_path(self) -> pulumi.Output[str]:
        """
        The path to the kube-controller-manager kubeconfig file.
        """
        return pulumi.get(self, "kube_controller_manager_kubeconfig_path")

    @property
    @pulumi.getter(name="kubeControllerManagerService")
    def kube_controller_manager_service(self) -> pulumi.Output[Optional['SystemdService']]:
        """
        The kube-controller-manager systemd service.
        """
        return pulumi.get(self, "kube_controller_manager_service")

    @property
    @pulumi.getter(name="kubeSchedulerConfigYaml")
    def kube_scheduler_config_yaml(self) -> pulumi.Output[str]:
        """
        The kube-scheduler config yaml.
        """
        return pulumi.get(self, "kube_scheduler_config_yaml")

    @property
    @pulumi.getter(name="kubeSchedulerInstall")
    def kube_scheduler_install(self) -> pulumi.Output['KubeSchedulerInstall']:
        """
        The kube-scheduler isntall.
        """
        return pulumi.get(self, "kube_scheduler_install")

    @property
    @pulumi.getter(name="kubeSchedulerInstallDirectory")
    def kube_scheduler_install_directory(self) -> pulumi.Output[Optional[str]]:
        """
        The directory to store the kube-scheduler binary.
        """
        return pulumi.get(self, "kube_scheduler_install_directory")

    @property
    @pulumi.getter(name="kubeSchedulerKubeconfigPath")
    def kube_scheduler_kubeconfig_path(self) -> pulumi.Output[str]:
        """
        The path to the kube-scheduler kubeconfig file.
        """
        return pulumi.get(self, "kube_scheduler_kubeconfig_path")

    @property
    @pulumi.getter(name="kubeSchedulerService")
    def kube_scheduler_service(self) -> pulumi.Output[Optional['SystemdService']]:
        """
        The kube-scheduler systemd service.
        """
        return pulumi.get(self, "kube_scheduler_service")

    @property
    @pulumi.getter(name="kubectlInstall")
    def kubectl_install(self) -> pulumi.Output['KubectlInstall']:
        """
        The kubectl install.
        """
        return pulumi.get(self, "kubectl_install")

    @property
    @pulumi.getter(name="kubectlInstallDirectory")
    def kubectl_install_directory(self) -> pulumi.Output[Optional[str]]:
        """
        The path to store the kubectl binary.
        """
        return pulumi.get(self, "kubectl_install_directory")

    @property
    @pulumi.getter(name="kubernetesConfigurationMkdir")
    def kubernetes_configuration_mkdir(self) -> pulumi.Output['_tools.Mkdir']:
        """
        The kubernetes configuration mkdir operation.
        """
        return pulumi.get(self, "kubernetes_configuration_mkdir")

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of kubernetes to use.
        """
        return pulumi.get(self, "kubernetes_version")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the node.
        """
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter(name="serviceAccountsCertificatePath")
    def service_accounts_certificate_path(self) -> pulumi.Output[str]:
        """
        The path to the service accounts certificate.
        """
        return pulumi.get(self, "service_accounts_certificate_path")

    @property
    @pulumi.getter(name="serviceAccountsPrivateKeyPath")
    def service_accounts_private_key_path(self) -> pulumi.Output[str]:
        """
        The path to the service accounts private key.
        """
        return pulumi.get(self, "service_accounts_private_key_path")

    @property
    @pulumi.getter(name="serviceClusterIpRange")
    def service_cluster_ip_range(self) -> pulumi.Output[Optional[str]]:
        """
        The IP range to use for cluster services.
        """
        return pulumi.get(self, "service_cluster_ip_range")

    @property
    @pulumi.getter(name="varLibKubernetesMkdir")
    def var_lib_kubernetes_mkdir(self) -> pulumi.Output['_tools.Mkdir']:
        """
        The /var/lib/kubernetes mkdir operation.
        """
        return pulumi.get(self, "var_lib_kubernetes_mkdir")

