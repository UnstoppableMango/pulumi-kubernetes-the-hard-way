# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *
import pulumi_command

__all__ = [
    'EtcdConfigurationPropsArgs',
    'EtcdNodeArgs',
    'KubeProxyConfigurationPropsArgs',
    'KubeletConfigurationPropsArgs',
    'SystemdInstallSectionArgs',
    'SystemdServiceSectionArgs',
    'SystemdUnitSectionArgs',
]

@pulumi.input_type
class EtcdConfigurationPropsArgs:
    def __init__(__self__, *,
                 ca_file_path: pulumi.Input[str],
                 cert_file_path: pulumi.Input[str],
                 data_directory: pulumi.Input[str],
                 etcd_path: pulumi.Input[str],
                 internal_ip: pulumi.Input[str],
                 key_file_path: pulumi.Input[str],
                 name: pulumi.Input[str]):
        """
        Props for resources that consume etcd configuration.
        :param pulumi.Input[str] ca_file_path: Path to the certificate authority file on the remote system.
        :param pulumi.Input[str] cert_file_path: Path to the certificate file on the remote system.
        :param pulumi.Input[str] data_directory: Etcd's data directory.
        :param pulumi.Input[str] etcd_path: Path to the etcd binary.
        :param pulumi.Input[str] internal_ip: Internal IP of the etcd node.
        :param pulumi.Input[str] key_file_path: Path to the private key file on the remote system.
        :param pulumi.Input[str] name: Name of the etcd node.
        """
        pulumi.set(__self__, "ca_file_path", ca_file_path)
        pulumi.set(__self__, "cert_file_path", cert_file_path)
        pulumi.set(__self__, "data_directory", data_directory)
        pulumi.set(__self__, "etcd_path", etcd_path)
        pulumi.set(__self__, "internal_ip", internal_ip)
        pulumi.set(__self__, "key_file_path", key_file_path)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="caFilePath")
    def ca_file_path(self) -> pulumi.Input[str]:
        """
        Path to the certificate authority file on the remote system.
        """
        return pulumi.get(self, "ca_file_path")

    @ca_file_path.setter
    def ca_file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_file_path", value)

    @property
    @pulumi.getter(name="certFilePath")
    def cert_file_path(self) -> pulumi.Input[str]:
        """
        Path to the certificate file on the remote system.
        """
        return pulumi.get(self, "cert_file_path")

    @cert_file_path.setter
    def cert_file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert_file_path", value)

    @property
    @pulumi.getter(name="dataDirectory")
    def data_directory(self) -> pulumi.Input[str]:
        """
        Etcd's data directory.
        """
        return pulumi.get(self, "data_directory")

    @data_directory.setter
    def data_directory(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_directory", value)

    @property
    @pulumi.getter(name="etcdPath")
    def etcd_path(self) -> pulumi.Input[str]:
        """
        Path to the etcd binary.
        """
        return pulumi.get(self, "etcd_path")

    @etcd_path.setter
    def etcd_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "etcd_path", value)

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Input[str]:
        """
        Internal IP of the etcd node.
        """
        return pulumi.get(self, "internal_ip")

    @internal_ip.setter
    def internal_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_ip", value)

    @property
    @pulumi.getter(name="keyFilePath")
    def key_file_path(self) -> pulumi.Input[str]:
        """
        Path to the private key file on the remote system.
        """
        return pulumi.get(self, "key_file_path")

    @key_file_path.setter
    def key_file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_file_path", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the etcd node.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class EtcdNodeArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 internal_ip: pulumi.Input[str],
                 architecture: Optional[pulumi.Input['Architecture']] = None):
        """
        Etcd node description.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input[str] internal_ip: The internal IP of the node.
        :param pulumi.Input['Architecture'] architecture: The CPU architecture of the node.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "internal_ip", internal_ip)
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)

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
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Input[str]:
        """
        The internal IP of the node.
        """
        return pulumi.get(self, "internal_ip")

    @internal_ip.setter
    def internal_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_ip", value)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input['Architecture']]:
        """
        The CPU architecture of the node.
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input['Architecture']]):
        pulumi.set(self, "architecture", value)


@pulumi.input_type
class KubeProxyConfigurationPropsArgs:
    def __init__(__self__, *,
                 configuration_file_path: pulumi.Input[str],
                 kube_proxy_path: pulumi.Input[str]):
        """
        Props for resources that consume kube-proxy configuration.
        :param pulumi.Input[str] configuration_file_path: Path to the kube proxy configuration file
        :param pulumi.Input[str] kube_proxy_path: Path to the kube-proxy binary.
        """
        pulumi.set(__self__, "configuration_file_path", configuration_file_path)
        pulumi.set(__self__, "kube_proxy_path", kube_proxy_path)

    @property
    @pulumi.getter(name="configurationFilePath")
    def configuration_file_path(self) -> pulumi.Input[str]:
        """
        Path to the kube proxy configuration file
        """
        return pulumi.get(self, "configuration_file_path")

    @configuration_file_path.setter
    def configuration_file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "configuration_file_path", value)

    @property
    @pulumi.getter(name="kubeProxyPath")
    def kube_proxy_path(self) -> pulumi.Input[str]:
        """
        Path to the kube-proxy binary.
        """
        return pulumi.get(self, "kube_proxy_path")

    @kube_proxy_path.setter
    def kube_proxy_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_proxy_path", value)


@pulumi.input_type
class KubeletConfigurationPropsArgs:
    def __init__(__self__, *,
                 configuration_file_path: pulumi.Input[str],
                 kubeconfig_path: pulumi.Input[str],
                 kubelet_path: pulumi.Input[str],
                 register_node: pulumi.Input[bool],
                 v: pulumi.Input[int]):
        """
        Props for resources that consume kubelet configuration.
        :param pulumi.Input[str] configuration_file_path: Path to the kubelet configuration.
        :param pulumi.Input[str] kubeconfig_path: Path to the kubeconfig the kubelet will use
        :param pulumi.Input[str] kubelet_path: Path to the kubelet binary.
        :param pulumi.Input[bool] register_node: Whether to register the node. Defaults to `true`.
        :param pulumi.Input[int] v: Verbosity. Defaults to `2`.
        """
        pulumi.set(__self__, "configuration_file_path", configuration_file_path)
        pulumi.set(__self__, "kubeconfig_path", kubeconfig_path)
        pulumi.set(__self__, "kubelet_path", kubelet_path)
        pulumi.set(__self__, "register_node", register_node)
        pulumi.set(__self__, "v", v)

    @property
    @pulumi.getter(name="configurationFilePath")
    def configuration_file_path(self) -> pulumi.Input[str]:
        """
        Path to the kubelet configuration.
        """
        return pulumi.get(self, "configuration_file_path")

    @configuration_file_path.setter
    def configuration_file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "configuration_file_path", value)

    @property
    @pulumi.getter(name="kubeconfigPath")
    def kubeconfig_path(self) -> pulumi.Input[str]:
        """
        Path to the kubeconfig the kubelet will use
        """
        return pulumi.get(self, "kubeconfig_path")

    @kubeconfig_path.setter
    def kubeconfig_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kubeconfig_path", value)

    @property
    @pulumi.getter(name="kubeletPath")
    def kubelet_path(self) -> pulumi.Input[str]:
        """
        Path to the kubelet binary.
        """
        return pulumi.get(self, "kubelet_path")

    @kubelet_path.setter
    def kubelet_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kubelet_path", value)

    @property
    @pulumi.getter(name="registerNode")
    def register_node(self) -> pulumi.Input[bool]:
        """
        Whether to register the node. Defaults to `true`.
        """
        return pulumi.get(self, "register_node")

    @register_node.setter
    def register_node(self, value: pulumi.Input[bool]):
        pulumi.set(self, "register_node", value)

    @property
    @pulumi.getter
    def v(self) -> pulumi.Input[int]:
        """
        Verbosity. Defaults to `2`.
        """
        return pulumi.get(self, "v")

    @v.setter
    def v(self, value: pulumi.Input[int]):
        pulumi.set(self, "v", value)


@pulumi.input_type
class SystemdInstallSectionArgs:
    def __init__(__self__, *,
                 wanted_by: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options
        :param pulumi.Input[Sequence[pulumi.Input[str]]] wanted_by: A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
        """
        if wanted_by is not None:
            pulumi.set(__self__, "wanted_by", wanted_by)

    @property
    @pulumi.getter(name="wantedBy")
    def wanted_by(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
        """
        return pulumi.get(self, "wanted_by")

    @wanted_by.setter
    def wanted_by(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "wanted_by", value)


@pulumi.input_type
class SystemdServiceSectionArgs:
    def __init__(__self__, *,
                 delegate: Optional[pulumi.Input['SystemdDelegate']] = None,
                 environment: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exec_start: Optional[pulumi.Input[str]] = None,
                 exec_start_pre: Optional[pulumi.Input[str]] = None,
                 exit_type: Optional[pulumi.Input['SystemdServiceExitType']] = None,
                 kill_mode: Optional[pulumi.Input['SystemdKillMode']] = None,
                 limit_core: Optional[pulumi.Input[str]] = None,
                 limit_n_proc: Optional[pulumi.Input[str]] = None,
                 limit_no_file: Optional[pulumi.Input[int]] = None,
                 oom_score_adjust: Optional[pulumi.Input[int]] = None,
                 restart: Optional[pulumi.Input['SystemdServiceRestart']] = None,
                 restart_sec: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['SystemdServiceType']] = None):
        """
        https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
        :param pulumi.Input['SystemdDelegate'] delegate: Turns on delegation of further resource control partitioning to processes of the unit.
        :param pulumi.Input[str] exec_start: Commands that are executed when this service is started.
        :param pulumi.Input[str] exec_start_pre: Additional commands that are executed before the command in ExecStart=.
        :param pulumi.Input['SystemdServiceExitType'] exit_type: Specifies when the manager should consider the service to be finished.
        :param pulumi.Input['SystemdKillMode'] kill_mode: Specifies how processes of this unit shall be killed.
        :param pulumi.Input[str] limit_core: https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        :param pulumi.Input[str] limit_n_proc: https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        :param pulumi.Input[int] limit_no_file: https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        :param pulumi.Input[int] oom_score_adjust: https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
        :param pulumi.Input['SystemdServiceRestart'] restart: Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
        :param pulumi.Input[str] restart_sec: Configures the time to sleep before restarting a service (as configured with Restart=).
        :param pulumi.Input['SystemdServiceType'] type: Configures the mechanism via which the service notifies the manager that the service start-up has finished.
        """
        if delegate is not None:
            pulumi.set(__self__, "delegate", delegate)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if exec_start is not None:
            pulumi.set(__self__, "exec_start", exec_start)
        if exec_start_pre is not None:
            pulumi.set(__self__, "exec_start_pre", exec_start_pre)
        if exit_type is not None:
            pulumi.set(__self__, "exit_type", exit_type)
        if kill_mode is not None:
            pulumi.set(__self__, "kill_mode", kill_mode)
        if limit_core is not None:
            pulumi.set(__self__, "limit_core", limit_core)
        if limit_n_proc is not None:
            pulumi.set(__self__, "limit_n_proc", limit_n_proc)
        if limit_no_file is not None:
            pulumi.set(__self__, "limit_no_file", limit_no_file)
        if oom_score_adjust is not None:
            pulumi.set(__self__, "oom_score_adjust", oom_score_adjust)
        if restart is not None:
            pulumi.set(__self__, "restart", restart)
        if restart_sec is not None:
            pulumi.set(__self__, "restart_sec", restart_sec)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def delegate(self) -> Optional[pulumi.Input['SystemdDelegate']]:
        """
        Turns on delegation of further resource control partitioning to processes of the unit.
        """
        return pulumi.get(self, "delegate")

    @delegate.setter
    def delegate(self, value: Optional[pulumi.Input['SystemdDelegate']]):
        pulumi.set(self, "delegate", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="execStart")
    def exec_start(self) -> Optional[pulumi.Input[str]]:
        """
        Commands that are executed when this service is started.
        """
        return pulumi.get(self, "exec_start")

    @exec_start.setter
    def exec_start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exec_start", value)

    @property
    @pulumi.getter(name="execStartPre")
    def exec_start_pre(self) -> Optional[pulumi.Input[str]]:
        """
        Additional commands that are executed before the command in ExecStart=.
        """
        return pulumi.get(self, "exec_start_pre")

    @exec_start_pre.setter
    def exec_start_pre(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exec_start_pre", value)

    @property
    @pulumi.getter(name="exitType")
    def exit_type(self) -> Optional[pulumi.Input['SystemdServiceExitType']]:
        """
        Specifies when the manager should consider the service to be finished.
        """
        return pulumi.get(self, "exit_type")

    @exit_type.setter
    def exit_type(self, value: Optional[pulumi.Input['SystemdServiceExitType']]):
        pulumi.set(self, "exit_type", value)

    @property
    @pulumi.getter(name="killMode")
    def kill_mode(self) -> Optional[pulumi.Input['SystemdKillMode']]:
        """
        Specifies how processes of this unit shall be killed.
        """
        return pulumi.get(self, "kill_mode")

    @kill_mode.setter
    def kill_mode(self, value: Optional[pulumi.Input['SystemdKillMode']]):
        pulumi.set(self, "kill_mode", value)

    @property
    @pulumi.getter(name="limitCore")
    def limit_core(self) -> Optional[pulumi.Input[str]]:
        """
        https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        """
        return pulumi.get(self, "limit_core")

    @limit_core.setter
    def limit_core(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_core", value)

    @property
    @pulumi.getter(name="limitNProc")
    def limit_n_proc(self) -> Optional[pulumi.Input[str]]:
        """
        https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        """
        return pulumi.get(self, "limit_n_proc")

    @limit_n_proc.setter
    def limit_n_proc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_n_proc", value)

    @property
    @pulumi.getter(name="limitNoFile")
    def limit_no_file(self) -> Optional[pulumi.Input[int]]:
        """
        https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#Process%20Properties
        """
        return pulumi.get(self, "limit_no_file")

    @limit_no_file.setter
    def limit_no_file(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "limit_no_file", value)

    @property
    @pulumi.getter(name="oomScoreAdjust")
    def oom_score_adjust(self) -> Optional[pulumi.Input[int]]:
        """
        https://www.freedesktop.org/software/systemd/man/latest/systemd.exec.html#OOMScoreAdjust=
        """
        return pulumi.get(self, "oom_score_adjust")

    @oom_score_adjust.setter
    def oom_score_adjust(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "oom_score_adjust", value)

    @property
    @pulumi.getter
    def restart(self) -> Optional[pulumi.Input['SystemdServiceRestart']]:
        """
        Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
        """
        return pulumi.get(self, "restart")

    @restart.setter
    def restart(self, value: Optional[pulumi.Input['SystemdServiceRestart']]):
        pulumi.set(self, "restart", value)

    @property
    @pulumi.getter(name="restartSec")
    def restart_sec(self) -> Optional[pulumi.Input[str]]:
        """
        Configures the time to sleep before restarting a service (as configured with Restart=).
        """
        return pulumi.get(self, "restart_sec")

    @restart_sec.setter
    def restart_sec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restart_sec", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['SystemdServiceType']]:
        """
        Configures the mechanism via which the service notifies the manager that the service start-up has finished.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['SystemdServiceType']]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class SystemdUnitSectionArgs:
    def __init__(__self__, *,
                 after: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 before: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 binds_to: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 documentation: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 requires: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 requisite: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 wants: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#
        :param pulumi.Input[Sequence[pulumi.Input[str]]] after: Those two settings configure ordering dependencies between units.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] before: Those two settings configure ordering dependencies between units.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] binds_to: Configures requirement dependencies, very similar in style to Requires=.
        :param pulumi.Input[str] description: A short human readable title of the unit.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] documentation: A space-separated list of URIs referencing documentation for this unit or its configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requires: Similar to Wants=, but declares a stronger requirement dependency.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requisite: Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] wants: Configures (weak) requirement dependencies on other units.
        """
        if after is not None:
            pulumi.set(__self__, "after", after)
        if before is not None:
            pulumi.set(__self__, "before", before)
        if binds_to is not None:
            pulumi.set(__self__, "binds_to", binds_to)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if documentation is not None:
            pulumi.set(__self__, "documentation", documentation)
        if requires is not None:
            pulumi.set(__self__, "requires", requires)
        if requisite is not None:
            pulumi.set(__self__, "requisite", requisite)
        if wants is not None:
            pulumi.set(__self__, "wants", wants)

    @property
    @pulumi.getter
    def after(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Those two settings configure ordering dependencies between units.
        """
        return pulumi.get(self, "after")

    @after.setter
    def after(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "after", value)

    @property
    @pulumi.getter
    def before(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Those two settings configure ordering dependencies between units.
        """
        return pulumi.get(self, "before")

    @before.setter
    def before(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "before", value)

    @property
    @pulumi.getter(name="bindsTo")
    def binds_to(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Configures requirement dependencies, very similar in style to Requires=.
        """
        return pulumi.get(self, "binds_to")

    @binds_to.setter
    def binds_to(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "binds_to", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A short human readable title of the unit.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def documentation(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A space-separated list of URIs referencing documentation for this unit or its configuration.
        """
        return pulumi.get(self, "documentation")

    @documentation.setter
    def documentation(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "documentation", value)

    @property
    @pulumi.getter
    def requires(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Similar to Wants=, but declares a stronger requirement dependency.
        """
        return pulumi.get(self, "requires")

    @requires.setter
    def requires(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "requires", value)

    @property
    @pulumi.getter
    def requisite(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
        """
        return pulumi.get(self, "requisite")

    @requisite.setter
    def requisite(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "requisite", value)

    @property
    @pulumi.getter
    def wants(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Configures (weak) requirement dependencies on other units.
        """
        return pulumi.get(self, "wants")

    @wants.setter
    def wants(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "wants", value)


