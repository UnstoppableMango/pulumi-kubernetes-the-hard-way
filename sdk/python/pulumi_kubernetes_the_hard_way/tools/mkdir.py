# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
import pulumi_command

__all__ = ['MkdirArgs', 'Mkdir']

@pulumi.input_type
class MkdirArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 directory: pulumi.Input[str],
                 parents: Optional[pulumi.Input[bool]] = None,
                 remove_on_delete: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Mkdir resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: The connection details for the remote system.
        :param pulumi.Input[str] directory: The fully qualified path of the directory on the remote system.
        :param pulumi.Input[bool] parents: Corresponds to the `--parents` option.
        :param pulumi.Input[bool] remove_on_delete: Remove the created directory when the `Mkdir` resource is deleted or updated.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "directory", directory)
        if parents is not None:
            pulumi.set(__self__, "parents", parents)
        if remove_on_delete is not None:
            pulumi.set(__self__, "remove_on_delete", remove_on_delete)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Input['pulumi_command.remote.ConnectionArgs']:
        """
        The connection details for the remote system.
        """
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: pulumi.Input['pulumi_command.remote.ConnectionArgs']):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter
    def directory(self) -> pulumi.Input[str]:
        """
        The fully qualified path of the directory on the remote system.
        """
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: pulumi.Input[str]):
        pulumi.set(self, "directory", value)

    @property
    @pulumi.getter
    def parents(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the `--parents` option.
        """
        return pulumi.get(self, "parents")

    @parents.setter
    def parents(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "parents", value)

    @property
    @pulumi.getter(name="removeOnDelete")
    def remove_on_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Remove the created directory when the `Mkdir` resource is deleted or updated.
        """
        return pulumi.get(self, "remove_on_delete")

    @remove_on_delete.setter
    def remove_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "remove_on_delete", value)


class Mkdir(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 parents: Optional[pulumi.Input[bool]] = None,
                 remove_on_delete: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Represents the `mkdir` utility.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: The connection details for the remote system.
        :param pulumi.Input[str] directory: The fully qualified path of the directory on the remote system.
        :param pulumi.Input[bool] parents: Corresponds to the `--parents` option.
        :param pulumi.Input[bool] remove_on_delete: Remove the created directory when the `Mkdir` resource is deleted or updated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MkdirArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents the `mkdir` utility.

        :param str resource_name: The name of the resource.
        :param MkdirArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MkdirArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 parents: Optional[pulumi.Input[bool]] = None,
                 remove_on_delete: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MkdirArgs.__new__(MkdirArgs)

            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            if directory is None and not opts.urn:
                raise TypeError("Missing required property 'directory'")
            __props__.__dict__["directory"] = directory
            __props__.__dict__["parents"] = parents
            __props__.__dict__["remove_on_delete"] = remove_on_delete
            __props__.__dict__["command"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Mkdir, __self__).__init__(
            'kubernetes-the-hard-way:tools:Mkdir',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output['pulumi_command.remote.Command']:
        """
        The remote command.
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def directory(self) -> pulumi.Output[str]:
        """
        The fully qualified path of the directory on the remote system.
        """
        return pulumi.get(self, "directory")

    @property
    @pulumi.getter
    def parents(self) -> pulumi.Output[bool]:
        """
        Corresponds to the `--parents` option.
        """
        return pulumi.get(self, "parents")

    @property
    @pulumi.getter(name="removeOnDelete")
    def remove_on_delete(self) -> pulumi.Output[bool]:
        """
        Remove the created directory when the `Mkdir` resource is deleted or updated.
        """
        return pulumi.get(self, "remove_on_delete")

    @property
    @pulumi.getter
    def stderr(self) -> pulumi.Output[str]:
        """
        The command's stderr.
        """
        return pulumi.get(self, "stderr")

    @property
    @pulumi.getter
    def stdout(self) -> pulumi.Output[str]:
        """
        The command's stdout.
        """
        return pulumi.get(self, "stdout")
