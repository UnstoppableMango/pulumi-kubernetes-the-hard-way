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
from .download import Download
from .file import File
import pulumi_command

__all__ = ['EtcdInstallArgs', 'EtcdInstall']

@pulumi.input_type
class EtcdInstallArgs:
    def __init__(__self__, *,
                 ca_pem: pulumi.Input[str],
                 cert_pem: pulumi.Input[str],
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 internal_ip: pulumi.Input[str],
                 key_pem: pulumi.Input[str],
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 configuration_directory: Optional[pulumi.Input[str]] = None,
                 data_directory: Optional[pulumi.Input[str]] = None,
                 download_directory: Optional[pulumi.Input[str]] = None,
                 install_directory: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EtcdInstall resource.
        :param pulumi.Input[str] ca_pem: The PEM encoded CA data.
        :param pulumi.Input[str] cert_pem: The PEM encoded certificate data.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: The connection details.
        :param pulumi.Input[str] internal_ip: IP used to serve client requests and communicate with etcd peers.
        :param pulumi.Input[str] key_pem: The PEM encoded key data.
        :param pulumi.Input['Architecture'] architecture: The etcd CPU architecture.
        :param pulumi.Input[str] configuration_directory: The directory to store etcd configuration.
        :param pulumi.Input[str] data_directory: The directory etcd will use.
        :param pulumi.Input[str] download_directory: Temporary directory to download files to. Defaults to `/tmp/<random string>`.
        :param pulumi.Input[str] install_directory: Directory to install the `etcd` and `etcdctl` binaries.
        :param pulumi.Input[str] version: The version of etcd to install.
        """
        pulumi.set(__self__, "ca_pem", ca_pem)
        pulumi.set(__self__, "cert_pem", cert_pem)
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "internal_ip", internal_ip)
        pulumi.set(__self__, "key_pem", key_pem)
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if configuration_directory is None:
            configuration_directory = '/etc/etcd'
        if configuration_directory is not None:
            pulumi.set(__self__, "configuration_directory", configuration_directory)
        if data_directory is None:
            data_directory = '/var/lib/etcd'
        if data_directory is not None:
            pulumi.set(__self__, "data_directory", data_directory)
        if download_directory is not None:
            pulumi.set(__self__, "download_directory", download_directory)
        if install_directory is None:
            install_directory = '/usr/local/bin'
        if install_directory is not None:
            pulumi.set(__self__, "install_directory", install_directory)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="caPem")
    def ca_pem(self) -> pulumi.Input[str]:
        """
        The PEM encoded CA data.
        """
        return pulumi.get(self, "ca_pem")

    @ca_pem.setter
    def ca_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_pem", value)

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> pulumi.Input[str]:
        """
        The PEM encoded certificate data.
        """
        return pulumi.get(self, "cert_pem")

    @cert_pem.setter
    def cert_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert_pem", value)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Input['pulumi_command.remote.ConnectionArgs']:
        """
        The connection details.
        """
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: pulumi.Input['pulumi_command.remote.ConnectionArgs']):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Input[str]:
        """
        IP used to serve client requests and communicate with etcd peers.
        """
        return pulumi.get(self, "internal_ip")

    @internal_ip.setter
    def internal_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_ip", value)

    @property
    @pulumi.getter(name="keyPem")
    def key_pem(self) -> pulumi.Input[str]:
        """
        The PEM encoded key data.
        """
        return pulumi.get(self, "key_pem")

    @key_pem.setter
    def key_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_pem", value)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input['Architecture']]:
        """
        The etcd CPU architecture.
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input['Architecture']]):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter(name="configurationDirectory")
    def configuration_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to store etcd configuration.
        """
        return pulumi.get(self, "configuration_directory")

    @configuration_directory.setter
    def configuration_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration_directory", value)

    @property
    @pulumi.getter(name="dataDirectory")
    def data_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory etcd will use.
        """
        return pulumi.get(self, "data_directory")

    @data_directory.setter
    def data_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_directory", value)

    @property
    @pulumi.getter(name="downloadDirectory")
    def download_directory(self) -> Optional[pulumi.Input[str]]:
        """
        Temporary directory to download files to. Defaults to `/tmp/<random string>`.
        """
        return pulumi.get(self, "download_directory")

    @download_directory.setter
    def download_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_directory", value)

    @property
    @pulumi.getter(name="installDirectory")
    def install_directory(self) -> Optional[pulumi.Input[str]]:
        """
        Directory to install the `etcd` and `etcdctl` binaries.
        """
        return pulumi.get(self, "install_directory")

    @install_directory.setter
    def install_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "install_directory", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of etcd to install.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class EtcdInstall(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 ca_pem: Optional[pulumi.Input[str]] = None,
                 cert_pem: Optional[pulumi.Input[str]] = None,
                 configuration_directory: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 data_directory: Optional[pulumi.Input[str]] = None,
                 download_directory: Optional[pulumi.Input[str]] = None,
                 install_directory: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 key_pem: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents an etcd binary on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['Architecture'] architecture: The etcd CPU architecture.
        :param pulumi.Input[str] ca_pem: The PEM encoded CA data.
        :param pulumi.Input[str] cert_pem: The PEM encoded certificate data.
        :param pulumi.Input[str] configuration_directory: The directory to store etcd configuration.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: The connection details.
        :param pulumi.Input[str] data_directory: The directory etcd will use.
        :param pulumi.Input[str] download_directory: Temporary directory to download files to. Defaults to `/tmp/<random string>`.
        :param pulumi.Input[str] install_directory: Directory to install the `etcd` and `etcdctl` binaries.
        :param pulumi.Input[str] internal_ip: IP used to serve client requests and communicate with etcd peers.
        :param pulumi.Input[str] key_pem: The PEM encoded key data.
        :param pulumi.Input[str] version: The version of etcd to install.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EtcdInstallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents an etcd binary on a remote system.

        :param str resource_name: The name of the resource.
        :param EtcdInstallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EtcdInstallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 ca_pem: Optional[pulumi.Input[str]] = None,
                 cert_pem: Optional[pulumi.Input[str]] = None,
                 configuration_directory: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 data_directory: Optional[pulumi.Input[str]] = None,
                 download_directory: Optional[pulumi.Input[str]] = None,
                 install_directory: Optional[pulumi.Input[str]] = None,
                 internal_ip: Optional[pulumi.Input[str]] = None,
                 key_pem: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EtcdInstallArgs.__new__(EtcdInstallArgs)

            __props__.__dict__["architecture"] = architecture
            if ca_pem is None and not opts.urn:
                raise TypeError("Missing required property 'ca_pem'")
            __props__.__dict__["ca_pem"] = ca_pem
            if cert_pem is None and not opts.urn:
                raise TypeError("Missing required property 'cert_pem'")
            __props__.__dict__["cert_pem"] = cert_pem
            if configuration_directory is None:
                configuration_directory = '/etc/etcd'
            __props__.__dict__["configuration_directory"] = configuration_directory
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            if data_directory is None:
                data_directory = '/var/lib/etcd'
            __props__.__dict__["data_directory"] = data_directory
            __props__.__dict__["download_directory"] = download_directory
            if install_directory is None:
                install_directory = '/usr/local/bin'
            __props__.__dict__["install_directory"] = install_directory
            if internal_ip is None and not opts.urn:
                raise TypeError("Missing required property 'internal_ip'")
            __props__.__dict__["internal_ip"] = internal_ip
            if key_pem is None and not opts.urn:
                raise TypeError("Missing required property 'key_pem'")
            __props__.__dict__["key_pem"] = key_pem
            __props__.__dict__["version"] = version
            __props__.__dict__["archive_name"] = None
            __props__.__dict__["ca_file"] = None
            __props__.__dict__["cert_file"] = None
            __props__.__dict__["download"] = None
            __props__.__dict__["download_mkdir"] = None
            __props__.__dict__["etcd_path"] = None
            __props__.__dict__["etcdctl_path"] = None
            __props__.__dict__["install_mkdir"] = None
            __props__.__dict__["key_file"] = None
            __props__.__dict__["mv_etcd"] = None
            __props__.__dict__["mv_etcdctl"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["systemd_service_file"] = None
            __props__.__dict__["tar"] = None
            __props__.__dict__["url"] = None
        super(EtcdInstall, __self__).__init__(
            'kubernetes-the-hard-way:remote:EtcdInstall',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Output['Architecture']:
        """
        The etcd CPU architecture.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="archiveName")
    def archive_name(self) -> pulumi.Output[str]:
        """
        The name of the etcd release archive.
        """
        return pulumi.get(self, "archive_name")

    @property
    @pulumi.getter(name="caFile")
    def ca_file(self) -> pulumi.Output[Optional['File']]:
        """
        The remote certificate authority file.
        """
        return pulumi.get(self, "ca_file")

    @property
    @pulumi.getter(name="certFile")
    def cert_file(self) -> pulumi.Output[Optional['File']]:
        """
        The remote certificate file.
        """
        return pulumi.get(self, "cert_file")

    @property
    @pulumi.getter
    def download(self) -> pulumi.Output['Download']:
        """
        The etcd download operation.
        """
        return pulumi.get(self, "download")

    @property
    @pulumi.getter(name="downloadDirectory")
    def download_directory(self) -> pulumi.Output[str]:
        """
        The directory where the etcd binary was downloaded to.
        """
        return pulumi.get(self, "download_directory")

    @property
    @pulumi.getter(name="downloadMkdir")
    def download_mkdir(self) -> pulumi.Output['_tools.Mkdir']:
        """
        The operation to create the download directory.
        """
        return pulumi.get(self, "download_mkdir")

    @property
    @pulumi.getter(name="etcdPath")
    def etcd_path(self) -> pulumi.Output[str]:
        """
        The path to the etcd binary on the remote system.
        """
        return pulumi.get(self, "etcd_path")

    @property
    @pulumi.getter(name="etcdctlPath")
    def etcdctl_path(self) -> pulumi.Output[str]:
        """
        The path to the etcdctl binary on the remote system.
        """
        return pulumi.get(self, "etcdctl_path")

    @property
    @pulumi.getter(name="installDirectory")
    def install_directory(self) -> pulumi.Output[str]:
        """
        Directory to install the `etcd` and `etcdctl` binaries.
        """
        return pulumi.get(self, "install_directory")

    @property
    @pulumi.getter(name="installMkdir")
    def install_mkdir(self) -> pulumi.Output['_tools.Mkdir']:
        """
        The operation to create the install directory.
        """
        return pulumi.get(self, "install_mkdir")

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> pulumi.Output[str]:
        """
        IP used to serve client requests and communicate with etcd peers.
        """
        return pulumi.get(self, "internal_ip")

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> pulumi.Output[Optional['File']]:
        """
        The remote key file.
        """
        return pulumi.get(self, "key_file")

    @property
    @pulumi.getter(name="mvEtcd")
    def mv_etcd(self) -> pulumi.Output['_tools.Mv']:
        """
        The operation to move the etcd binary to the install directory.
        """
        return pulumi.get(self, "mv_etcd")

    @property
    @pulumi.getter(name="mvEtcdctl")
    def mv_etcdctl(self) -> pulumi.Output['_tools.Mv']:
        """
        The operation to move the etcdctl binary to the install directory.
        """
        return pulumi.get(self, "mv_etcdctl")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemdServiceFile")
    def systemd_service_file(self) -> pulumi.Output['File']:
        """
        The remote systemd service file.
        """
        return pulumi.get(self, "systemd_service_file")

    @property
    @pulumi.getter
    def tar(self) -> pulumi.Output['_tools.Tar']:
        """
        The tar operation.
        """
        return pulumi.get(self, "tar")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The url used to download etcd.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version of etcd downloaded.
        """
        return pulumi.get(self, "version")

