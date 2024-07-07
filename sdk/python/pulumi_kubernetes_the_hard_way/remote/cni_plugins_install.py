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
import pulumi_command

__all__ = ['CniPluginsInstallArgs', 'CniPluginsInstall']

@pulumi.input_type
class CniPluginsInstallArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CniPluginsInstall resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input['Architecture'] architecture: The CPU architecture to install.
        :param pulumi.Input[str] directory: The directory to install the binary to.
        :param pulumi.Input[str] version: The version to install.
        """
        pulumi.set(__self__, "connection", connection)
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if directory is None:
            directory = '/opt/cni/bin'
        if directory is not None:
            pulumi.set(__self__, "directory", directory)
        if version is not None:
            pulumi.set(__self__, "version", version)

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
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input['Architecture']]:
        """
        The CPU architecture to install.
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input['Architecture']]):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter
    def directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to install the binary to.
        """
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version to install.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class CniPluginsInstall(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Installs cni-plugins on a remote system

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['Architecture'] architecture: The CPU architecture to install.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input[str] directory: The directory to install the binary to.
        :param pulumi.Input[str] version: The version to install.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CniPluginsInstallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Installs cni-plugins on a remote system

        :param str resource_name: The name of the resource.
        :param CniPluginsInstallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CniPluginsInstallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
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
            __props__ = CniPluginsInstallArgs.__new__(CniPluginsInstallArgs)

            __props__.__dict__["architecture"] = architecture
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            if directory is None:
                directory = '/opt/cni/bin'
            __props__.__dict__["directory"] = directory
            __props__.__dict__["version"] = version
            __props__.__dict__["archive_name"] = None
            __props__.__dict__["bandwidth_mv"] = None
            __props__.__dict__["bandwidth_path"] = None
            __props__.__dict__["bridge_mv"] = None
            __props__.__dict__["bridge_path"] = None
            __props__.__dict__["dhcp_mv"] = None
            __props__.__dict__["dhcp_path"] = None
            __props__.__dict__["download"] = None
            __props__.__dict__["dummy_mv"] = None
            __props__.__dict__["dummy_path"] = None
            __props__.__dict__["firewall_mv"] = None
            __props__.__dict__["firewall_path"] = None
            __props__.__dict__["host_device_mv"] = None
            __props__.__dict__["host_device_path"] = None
            __props__.__dict__["host_local_mv"] = None
            __props__.__dict__["host_local_path"] = None
            __props__.__dict__["ipvlan_mv"] = None
            __props__.__dict__["ipvlan_path"] = None
            __props__.__dict__["loopback_mv"] = None
            __props__.__dict__["loopback_path"] = None
            __props__.__dict__["macvlan_mv"] = None
            __props__.__dict__["macvlan_path"] = None
            __props__.__dict__["mkdir"] = None
            __props__.__dict__["mktemp"] = None
            __props__.__dict__["path"] = None
            __props__.__dict__["portmap_mv"] = None
            __props__.__dict__["portmap_path"] = None
            __props__.__dict__["ptp_mv"] = None
            __props__.__dict__["ptp_path"] = None
            __props__.__dict__["rm"] = None
            __props__.__dict__["sbr_mv"] = None
            __props__.__dict__["sbr_path"] = None
            __props__.__dict__["static_mv"] = None
            __props__.__dict__["static_path"] = None
            __props__.__dict__["tap_mv"] = None
            __props__.__dict__["tap_path"] = None
            __props__.__dict__["tar"] = None
            __props__.__dict__["tuning_mv"] = None
            __props__.__dict__["tuning_path"] = None
            __props__.__dict__["url"] = None
            __props__.__dict__["vlan_mv"] = None
            __props__.__dict__["vlan_path"] = None
            __props__.__dict__["vrf_mv"] = None
            __props__.__dict__["vrf_path"] = None
        super(CniPluginsInstall, __self__).__init__(
            'kubernetes-the-hard-way:remote:CniPluginsInstall',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Output['Architecture']:
        """
        The CPU architecture to install.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="archiveName")
    def archive_name(self) -> pulumi.Output[str]:
        """
        The name of the downloaded archive.
        """
        return pulumi.get(self, "archive_name")

    @property
    @pulumi.getter(name="bandwidthMv")
    def bandwidth_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The bandwidth mv operation.
        """
        return pulumi.get(self, "bandwidth_mv")

    @property
    @pulumi.getter(name="bandwidthPath")
    def bandwidth_path(self) -> pulumi.Output[str]:
        """
        The bandwidth path on the remote system
        """
        return pulumi.get(self, "bandwidth_path")

    @property
    @pulumi.getter(name="bridgeMv")
    def bridge_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The bridge mv operation.
        """
        return pulumi.get(self, "bridge_mv")

    @property
    @pulumi.getter(name="bridgePath")
    def bridge_path(self) -> pulumi.Output[str]:
        """
        The bridge path on the remote system
        """
        return pulumi.get(self, "bridge_path")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        The parameters with which to connect to the remote host.
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter(name="dhcpMv")
    def dhcp_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The dhcp mv operation.
        """
        return pulumi.get(self, "dhcp_mv")

    @property
    @pulumi.getter(name="dhcpPath")
    def dhcp_path(self) -> pulumi.Output[str]:
        """
        The dhcp path on the remote system
        """
        return pulumi.get(self, "dhcp_path")

    @property
    @pulumi.getter
    def directory(self) -> pulumi.Output[str]:
        """
        The directory to install the binary to.
        """
        return pulumi.get(self, "directory")

    @property
    @pulumi.getter
    def download(self) -> pulumi.Output['Download']:
        """
        The download operation.
        """
        return pulumi.get(self, "download")

    @property
    @pulumi.getter(name="dummyMv")
    def dummy_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The dummy mv operation.
        """
        return pulumi.get(self, "dummy_mv")

    @property
    @pulumi.getter(name="dummyPath")
    def dummy_path(self) -> pulumi.Output[str]:
        """
        The dummy path on the remote system
        """
        return pulumi.get(self, "dummy_path")

    @property
    @pulumi.getter(name="firewallMv")
    def firewall_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The firewall mv operation.
        """
        return pulumi.get(self, "firewall_mv")

    @property
    @pulumi.getter(name="firewallPath")
    def firewall_path(self) -> pulumi.Output[str]:
        """
        The firewall path on the remote system
        """
        return pulumi.get(self, "firewall_path")

    @property
    @pulumi.getter(name="hostDeviceMv")
    def host_device_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The hostDevice mv operation.
        """
        return pulumi.get(self, "host_device_mv")

    @property
    @pulumi.getter(name="hostDevicePath")
    def host_device_path(self) -> pulumi.Output[str]:
        """
        The hostDevice path on the remote system
        """
        return pulumi.get(self, "host_device_path")

    @property
    @pulumi.getter(name="hostLocalMv")
    def host_local_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The hostLocal mv operation.
        """
        return pulumi.get(self, "host_local_mv")

    @property
    @pulumi.getter(name="hostLocalPath")
    def host_local_path(self) -> pulumi.Output[str]:
        """
        The hostLocal path on the remote system
        """
        return pulumi.get(self, "host_local_path")

    @property
    @pulumi.getter(name="ipvlanMv")
    def ipvlan_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The ipvlan mv operation.
        """
        return pulumi.get(self, "ipvlan_mv")

    @property
    @pulumi.getter(name="ipvlanPath")
    def ipvlan_path(self) -> pulumi.Output[str]:
        """
        The ipvlan path on the remote system
        """
        return pulumi.get(self, "ipvlan_path")

    @property
    @pulumi.getter(name="loopbackMv")
    def loopback_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The loopback mv operation.
        """
        return pulumi.get(self, "loopback_mv")

    @property
    @pulumi.getter(name="loopbackPath")
    def loopback_path(self) -> pulumi.Output[str]:
        """
        The loopback path on the remote system
        """
        return pulumi.get(self, "loopback_path")

    @property
    @pulumi.getter(name="macvlanMv")
    def macvlan_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The macvlan mv operation.
        """
        return pulumi.get(self, "macvlan_mv")

    @property
    @pulumi.getter(name="macvlanPath")
    def macvlan_path(self) -> pulumi.Output[str]:
        """
        The macvlan path on the remote system
        """
        return pulumi.get(self, "macvlan_path")

    @property
    @pulumi.getter
    def mkdir(self) -> pulumi.Output['_tools.Mkdir']:
        """
        The mkdir operation.
        """
        return pulumi.get(self, "mkdir")

    @property
    @pulumi.getter
    def mktemp(self) -> pulumi.Output['_tools.Mktemp']:
        """
        The mktemp operation.
        """
        return pulumi.get(self, "mktemp")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        The path to the installed binary.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="portmapMv")
    def portmap_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The portmap mv operation.
        """
        return pulumi.get(self, "portmap_mv")

    @property
    @pulumi.getter(name="portmapPath")
    def portmap_path(self) -> pulumi.Output[str]:
        """
        The portmap path on the remote system
        """
        return pulumi.get(self, "portmap_path")

    @property
    @pulumi.getter(name="ptpMv")
    def ptp_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The ptp mv operation.
        """
        return pulumi.get(self, "ptp_mv")

    @property
    @pulumi.getter(name="ptpPath")
    def ptp_path(self) -> pulumi.Output[str]:
        """
        The ptp path on the remote system
        """
        return pulumi.get(self, "ptp_path")

    @property
    @pulumi.getter
    def rm(self) -> pulumi.Output['_tools.Rm']:
        """
        The rm operation.
        """
        return pulumi.get(self, "rm")

    @property
    @pulumi.getter(name="sbrMv")
    def sbr_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The sbr mv operation.
        """
        return pulumi.get(self, "sbr_mv")

    @property
    @pulumi.getter(name="sbrPath")
    def sbr_path(self) -> pulumi.Output[str]:
        """
        The sbr path on the remote system
        """
        return pulumi.get(self, "sbr_path")

    @property
    @pulumi.getter(name="staticMv")
    def static_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The static mv operation.
        """
        return pulumi.get(self, "static_mv")

    @property
    @pulumi.getter(name="staticPath")
    def static_path(self) -> pulumi.Output[str]:
        """
        The static path on the remote system
        """
        return pulumi.get(self, "static_path")

    @property
    @pulumi.getter(name="tapMv")
    def tap_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The tap mv operation.
        """
        return pulumi.get(self, "tap_mv")

    @property
    @pulumi.getter(name="tapPath")
    def tap_path(self) -> pulumi.Output[str]:
        """
        The tap path on the remote system
        """
        return pulumi.get(self, "tap_path")

    @property
    @pulumi.getter
    def tar(self) -> pulumi.Output['_tools.Tar']:
        """
        The tar operation.
        """
        return pulumi.get(self, "tar")

    @property
    @pulumi.getter(name="tuningMv")
    def tuning_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The tuning mv operation.
        """
        return pulumi.get(self, "tuning_mv")

    @property
    @pulumi.getter(name="tuningPath")
    def tuning_path(self) -> pulumi.Output[str]:
        """
        The tuning path on the remote system
        """
        return pulumi.get(self, "tuning_path")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The url used to download the binary.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version to install.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="vlanMv")
    def vlan_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The vlan mv operation.
        """
        return pulumi.get(self, "vlan_mv")

    @property
    @pulumi.getter(name="vlanPath")
    def vlan_path(self) -> pulumi.Output[str]:
        """
        The vlan path on the remote system
        """
        return pulumi.get(self, "vlan_path")

    @property
    @pulumi.getter(name="vrfMv")
    def vrf_mv(self) -> pulumi.Output['_tools.Mv']:
        """
        The vrf mv operation.
        """
        return pulumi.get(self, "vrf_mv")

    @property
    @pulumi.getter(name="vrfPath")
    def vrf_path(self) -> pulumi.Output[str]:
        """
        The vrf path on the remote system
        """
        return pulumi.get(self, "vrf_path")

