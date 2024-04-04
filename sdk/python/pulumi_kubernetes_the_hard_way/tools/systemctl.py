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

__all__ = ['SystemctlArgs', 'Systemctl']

@pulumi.input_type
class SystemctlArgs:
    def __init__(__self__, *,
                 commands: pulumi.Input[Sequence[pulumi.Input['SystemctlCommand']]],
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Systemctl resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system.
        """
        pulumi.set(__self__, "commands", commands)
        pulumi.set(__self__, "connection", connection)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def commands(self) -> pulumi.Input[Sequence[pulumi.Input['SystemctlCommand']]]:
        return pulumi.get(self, "commands")

    @commands.setter
    def commands(self, value: pulumi.Input[Sequence[pulumi.Input['SystemctlCommand']]]):
        pulumi.set(self, "commands", value)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Input['pulumi_command.remote.ConnectionArgs']:
        """
        Connection details for the remote system.
        """
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: pulumi.Input['pulumi_command.remote.ConnectionArgs']):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class Systemctl(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 commands: Optional[pulumi.Input[Sequence[pulumi.Input['SystemctlCommand']]]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Abstraction over the `systemctl` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemctlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Abstraction over the `systemctl` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param SystemctlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemctlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 commands: Optional[pulumi.Input[Sequence[pulumi.Input['SystemctlCommand']]]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemctlArgs.__new__(SystemctlArgs)

            if commands is None and not opts.urn:
                raise TypeError("Missing required property 'commands'")
            __props__.__dict__["commands"] = commands
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["command"] = None
        super(Systemctl, __self__).__init__(
            'kubernetes-the-hard-way:tools:Systemctl',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output['pulumi_command.remote.Command']:
        """
        Represents the command run on the remote system.
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def commands(self) -> pulumi.Output[Sequence['SystemctlCommand']]:
        return pulumi.get(self, "commands")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        Connection details for the remote system.
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "service_name")
