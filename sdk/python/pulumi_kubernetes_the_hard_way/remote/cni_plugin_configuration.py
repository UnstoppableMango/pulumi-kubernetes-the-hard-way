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
from .cni_bridge_plugin_configuration import CniBridgePluginConfiguration
from .cni_loopback_plugin_configuration import CniLoopbackPluginConfiguration
import pulumi_command

__all__ = ['CniPluginConfigurationArgs', 'CniPluginConfiguration']

@pulumi.input_type
class CniPluginConfigurationArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 subnet: pulumi.Input[str],
                 directory: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CniPluginConfiguration resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input[str] subnet: The subnet to use for the CNI bridge plugin configuration.
        :param pulumi.Input[str] directory: The plugin configuration directory.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "subnet", subnet)
        if directory is not None:
            pulumi.set(__self__, "directory", directory)

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
    def subnet(self) -> pulumi.Input[str]:
        """
        The subnet to use for the CNI bridge plugin configuration.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def directory(self) -> Optional[pulumi.Input[str]]:
        """
        The plugin configuration directory.
        """
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory", value)


class CniPluginConfiguration(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The CNI plugin configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input[str] directory: The plugin configuration directory.
        :param pulumi.Input[str] subnet: The subnet to use for the CNI bridge plugin configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CniPluginConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The CNI plugin configuration.

        :param str resource_name: The name of the resource.
        :param CniPluginConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CniPluginConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CniPluginConfigurationArgs.__new__(CniPluginConfigurationArgs)

            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["directory"] = directory
            if subnet is None and not opts.urn:
                raise TypeError("Missing required property 'subnet'")
            __props__.__dict__["subnet"] = subnet
            __props__.__dict__["bridge"] = None
            __props__.__dict__["loopback"] = None
            __props__.__dict__["mkdir"] = None
        super(CniPluginConfiguration, __self__).__init__(
            'kubernetes-the-hard-way:remote:CniPluginConfiguration',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def bridge(self) -> pulumi.Output['CniBridgePluginConfiguration']:
        """
        The bridge plugin configuration.
        """
        return pulumi.get(self, "bridge")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        The parameters with which to connect to the remote host.
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter
    def directory(self) -> pulumi.Output[str]:
        """
        The plugin configuration directory.
        """
        return pulumi.get(self, "directory")

    @property
    @pulumi.getter
    def loopback(self) -> pulumi.Output['CniLoopbackPluginConfiguration']:
        """
        The loopback plugin configuration.
        """
        return pulumi.get(self, "loopback")

    @property
    @pulumi.getter
    def mkdir(self) -> pulumi.Output['_tools.Mkdir']:
        """
        The `directory` mkdir operation.
        """
        return pulumi.get(self, "mkdir")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[str]:
        """
        The subnet to use for the CNI bridge plugin configuration.
        """
        return pulumi.get(self, "subnet")
