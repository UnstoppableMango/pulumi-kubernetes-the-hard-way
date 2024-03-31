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

__all__ = ['TarArgs', 'Tar']

@pulumi.input_type
class TarArgs:
    def __init__(__self__, *,
                 archive: pulumi.Input[str],
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 directory: Optional[pulumi.Input[str]] = None,
                 extract: Optional[pulumi.Input[bool]] = None,
                 files: Optional[pulumi.Input[Union[Sequence[pulumi.Input[str]], str]]] = None,
                 gzip: Optional[pulumi.Input[bool]] = None,
                 strip_components: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Tar resource.
        :param pulumi.Input[str] archive: Corresponds to the [ARCHIVE] argument.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system.
        :param pulumi.Input[str] directory: Corresponds to the --directory option.
        :param pulumi.Input[bool] extract: Corresponds to the --extract option.
        :param pulumi.Input[Union[Sequence[pulumi.Input[str]], str]] files: Corresponds to the [FILE] argument.
        :param pulumi.Input[bool] gzip: Corresponds to the --gzip option.
        :param pulumi.Input[int] strip_components: Corresponds to the --strip-components option.
        """
        pulumi.set(__self__, "archive", archive)
        pulumi.set(__self__, "connection", connection)
        if directory is not None:
            pulumi.set(__self__, "directory", directory)
        if extract is not None:
            pulumi.set(__self__, "extract", extract)
        if files is not None:
            pulumi.set(__self__, "files", files)
        if gzip is not None:
            pulumi.set(__self__, "gzip", gzip)
        if strip_components is not None:
            pulumi.set(__self__, "strip_components", strip_components)

    @property
    @pulumi.getter
    def archive(self) -> pulumi.Input[str]:
        """
        Corresponds to the [ARCHIVE] argument.
        """
        return pulumi.get(self, "archive")

    @archive.setter
    def archive(self, value: pulumi.Input[str]):
        pulumi.set(self, "archive", value)

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
    @pulumi.getter
    def directory(self) -> Optional[pulumi.Input[str]]:
        """
        Corresponds to the --directory option.
        """
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory", value)

    @property
    @pulumi.getter
    def extract(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the --extract option.
        """
        return pulumi.get(self, "extract")

    @extract.setter
    def extract(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "extract", value)

    @property
    @pulumi.getter
    def files(self) -> Optional[pulumi.Input[Union[Sequence[pulumi.Input[str]], str]]]:
        """
        Corresponds to the [FILE] argument.
        """
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: Optional[pulumi.Input[Union[Sequence[pulumi.Input[str]], str]]]):
        pulumi.set(self, "files", value)

    @property
    @pulumi.getter
    def gzip(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the --gzip option.
        """
        return pulumi.get(self, "gzip")

    @gzip.setter
    def gzip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "gzip", value)

    @property
    @pulumi.getter(name="stripComponents")
    def strip_components(self) -> Optional[pulumi.Input[int]]:
        """
        Corresponds to the --strip-components option.
        """
        return pulumi.get(self, "strip_components")

    @strip_components.setter
    def strip_components(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "strip_components", value)


class Tar(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 extract: Optional[pulumi.Input[bool]] = None,
                 files: Optional[pulumi.Input[Union[Sequence[pulumi.Input[str]], str]]] = None,
                 gzip: Optional[pulumi.Input[bool]] = None,
                 strip_components: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Abstracion over the `tar` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] archive: Corresponds to the [ARCHIVE] argument.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system.
        :param pulumi.Input[str] directory: Corresponds to the --directory option.
        :param pulumi.Input[bool] extract: Corresponds to the --extract option.
        :param pulumi.Input[Union[Sequence[pulumi.Input[str]], str]] files: Corresponds to the [FILE] argument.
        :param pulumi.Input[bool] gzip: Corresponds to the --gzip option.
        :param pulumi.Input[int] strip_components: Corresponds to the --strip-components option.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TarArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Abstracion over the `tar` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param TarArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TarArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 extract: Optional[pulumi.Input[bool]] = None,
                 files: Optional[pulumi.Input[Union[Sequence[pulumi.Input[str]], str]]] = None,
                 gzip: Optional[pulumi.Input[bool]] = None,
                 strip_components: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TarArgs.__new__(TarArgs)

            if archive is None and not opts.urn:
                raise TypeError("Missing required property 'archive'")
            __props__.__dict__["archive"] = archive
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["directory"] = directory
            __props__.__dict__["extract"] = extract
            __props__.__dict__["files"] = files
            __props__.__dict__["gzip"] = gzip
            __props__.__dict__["strip_components"] = strip_components
            __props__.__dict__["command"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdin"] = None
            __props__.__dict__["stdout"] = None
        super(Tar, __self__).__init__(
            'kubernetes-the-hard-way:tools:Tar',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def archive(self) -> pulumi.Output[str]:
        """
        Corresponds to the [ARCHIVE] argument.
        """
        return pulumi.get(self, "archive")

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output['pulumi_command.remote.Command']:
        """
        Represents the remote `tar` operation.
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def directory(self) -> pulumi.Output[Optional[str]]:
        """
        Corresponds to the --directory option.
        """
        return pulumi.get(self, "directory")

    @property
    @pulumi.getter
    def extract(self) -> pulumi.Output[bool]:
        """
        Corresponds to the --extract option.
        """
        return pulumi.get(self, "extract")

    @property
    @pulumi.getter
    def files(self) -> pulumi.Output[Sequence[str]]:
        """
        Corresponds to the [FILE] argument.
        """
        return pulumi.get(self, "files")

    @property
    @pulumi.getter
    def gzip(self) -> pulumi.Output[Optional[bool]]:
        """
        Corresponds to the --gzip option.
        """
        return pulumi.get(self, "gzip")

    @property
    @pulumi.getter
    def stderr(self) -> pulumi.Output[str]:
        """
        The process' stderr.
        """
        return pulumi.get(self, "stderr")

    @property
    @pulumi.getter
    def stdin(self) -> pulumi.Output[Optional[str]]:
        """
        The process' stdin.
        """
        return pulumi.get(self, "stdin")

    @property
    @pulumi.getter
    def stdout(self) -> pulumi.Output[str]:
        """
        The process' stdout.
        """
        return pulumi.get(self, "stdout")

    @property
    @pulumi.getter(name="stripComponents")
    def strip_components(self) -> pulumi.Output[Optional[int]]:
        """
        Corresponds to the --strip-components option.
        """
        return pulumi.get(self, "strip_components")

