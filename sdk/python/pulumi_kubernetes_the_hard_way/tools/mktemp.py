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

__all__ = ['MktempArgs', 'Mktemp']

@pulumi.input_type
class MktempArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 template: pulumi.Input[str],
                 directory: Optional[pulumi.Input[bool]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 suffix: Optional[pulumi.Input[str]] = None,
                 tmpdir: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Mktemp resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system.
        :param pulumi.Input[str] template: Corresponds to the [TEMPLATE] arg.
        :param pulumi.Input[bool] directory: Corresponds to the --directory option.
        :param pulumi.Input[bool] dry_run: Corresponds to the --dry-run option.
        :param pulumi.Input[bool] quiet: Corresponds to the --quiet option.
        :param pulumi.Input[str] suffix: Corresponds to the --suffix option.
        :param pulumi.Input[str] tmpdir: Corresponds to the --tmpdir option.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "template", template)
        if directory is not None:
            pulumi.set(__self__, "directory", directory)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if quiet is not None:
            pulumi.set(__self__, "quiet", quiet)
        if suffix is not None:
            pulumi.set(__self__, "suffix", suffix)
        if tmpdir is not None:
            pulumi.set(__self__, "tmpdir", tmpdir)

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
    def template(self) -> pulumi.Input[str]:
        """
        Corresponds to the [TEMPLATE] arg.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: pulumi.Input[str]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter
    def directory(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the --directory option.
        """
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "directory", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the --dry-run option.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter
    def quiet(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the --quiet option.
        """
        return pulumi.get(self, "quiet")

    @quiet.setter
    def quiet(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "quiet", value)

    @property
    @pulumi.getter
    def suffix(self) -> Optional[pulumi.Input[str]]:
        """
        Corresponds to the --suffix option.
        """
        return pulumi.get(self, "suffix")

    @suffix.setter
    def suffix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "suffix", value)

    @property
    @pulumi.getter
    def tmpdir(self) -> Optional[pulumi.Input[str]]:
        """
        Corresponds to the --tmpdir option.
        """
        return pulumi.get(self, "tmpdir")

    @tmpdir.setter
    def tmpdir(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tmpdir", value)


class Mktemp(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[bool]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 suffix: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 tmpdir: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Abstracion over the `mktemp` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system.
        :param pulumi.Input[bool] directory: Corresponds to the --directory option.
        :param pulumi.Input[bool] dry_run: Corresponds to the --dry-run option.
        :param pulumi.Input[bool] quiet: Corresponds to the --quiet option.
        :param pulumi.Input[str] suffix: Corresponds to the --suffix option.
        :param pulumi.Input[str] template: Corresponds to the [TEMPLATE] arg.
        :param pulumi.Input[str] tmpdir: Corresponds to the --tmpdir option.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MktempArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Abstracion over the `mktemp` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param MktempArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MktempArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory: Optional[pulumi.Input[bool]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 suffix: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 tmpdir: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MktempArgs.__new__(MktempArgs)

            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["directory"] = directory
            __props__.__dict__["dry_run"] = dry_run
            __props__.__dict__["quiet"] = quiet
            __props__.__dict__["suffix"] = suffix
            if template is None and not opts.urn:
                raise TypeError("Missing required property 'template'")
            __props__.__dict__["template"] = template
            __props__.__dict__["tmpdir"] = tmpdir
            __props__.__dict__["command"] = None
        super(Mktemp, __self__).__init__(
            'kubernetes-the-hard-way:tools:Mktemp',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output['pulumi_command.remote.Command']:
        """
        Represents the remote `tar` operation.
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def directory(self) -> pulumi.Output[bool]:
        """
        Corresponds to the --directory option.
        """
        return pulumi.get(self, "directory")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[bool]:
        """
        Corresponds to the --dry-run option.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter
    def quiet(self) -> pulumi.Output[bool]:
        """
        Corresponds to the --quiet option.
        """
        return pulumi.get(self, "quiet")

    @property
    @pulumi.getter
    def suffix(self) -> pulumi.Output[Optional[str]]:
        """
        Corresponds to the --suffix option.
        """
        return pulumi.get(self, "suffix")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[str]:
        """
        Corresponds to the [TEMPLATE] arg.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def tmpdir(self) -> pulumi.Output[Optional[str]]:
        """
        Corresponds to the --tmpdir option.
        """
        return pulumi.get(self, "tmpdir")

