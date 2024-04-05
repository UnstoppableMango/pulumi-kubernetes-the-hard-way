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

__all__ = ['TeeArgs', 'Tee']

@pulumi.input_type
class TeeArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 files: pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]],
                 stdin: pulumi.Input[str],
                 append: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ignore_interrupts: Optional[pulumi.Input[bool]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 output_error: Optional[pulumi.Input['TeeMode']] = None,
                 pipe: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Tee resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system.
        :param pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]] files: The file(s) to write to.
        :param pulumi.Input[bool] append: Append to the given FILEs, do not overwrite.
        :param pulumi.Input[bool] ignore_interrupts: Ignore interrupt signals.
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run.
        :param pulumi.Input['TeeMode'] output_error: Set behavior on write error.
        :param pulumi.Input[bool] pipe: Operate in a more appropriate MODE with pipes.
        :param pulumi.Input[bool] version: Output version information and exit.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "files", files)
        pulumi.set(__self__, "stdin", stdin)
        if append is not None:
            pulumi.set(__self__, "append", append)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if ignore_interrupts is not None:
            pulumi.set(__self__, "ignore_interrupts", ignore_interrupts)
        if lifecycle is not None:
            pulumi.set(__self__, "lifecycle", lifecycle)
        if output_error is not None:
            pulumi.set(__self__, "output_error", output_error)
        if pipe is not None:
            pulumi.set(__self__, "pipe", pipe)
        if version is not None:
            pulumi.set(__self__, "version", version)

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
    def files(self) -> pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]:
        """
        The file(s) to write to.
        """
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "files", value)

    @property
    @pulumi.getter
    def stdin(self) -> pulumi.Input[str]:
        return pulumi.get(self, "stdin")

    @stdin.setter
    def stdin(self, value: pulumi.Input[str]):
        pulumi.set(self, "stdin", value)

    @property
    @pulumi.getter
    def append(self) -> Optional[pulumi.Input[bool]]:
        """
        Append to the given FILEs, do not overwrite.
        """
        return pulumi.get(self, "append")

    @append.setter
    def append(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "append", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="ignoreInterrupts")
    def ignore_interrupts(self) -> Optional[pulumi.Input[bool]]:
        """
        Ignore interrupt signals.
        """
        return pulumi.get(self, "ignore_interrupts")

    @ignore_interrupts.setter
    def ignore_interrupts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_interrupts", value)

    @property
    @pulumi.getter
    def lifecycle(self) -> Optional['CommandLifecycle']:
        """
        At what stage(s) in the resource lifecycle should the command be run.
        """
        return pulumi.get(self, "lifecycle")

    @lifecycle.setter
    def lifecycle(self, value: Optional['CommandLifecycle']):
        pulumi.set(self, "lifecycle", value)

    @property
    @pulumi.getter(name="outputError")
    def output_error(self) -> Optional[pulumi.Input['TeeMode']]:
        """
        Set behavior on write error.
        """
        return pulumi.get(self, "output_error")

    @output_error.setter
    def output_error(self, value: Optional[pulumi.Input['TeeMode']]):
        pulumi.set(self, "output_error", value)

    @property
    @pulumi.getter
    def pipe(self) -> Optional[pulumi.Input[bool]]:
        """
        Operate in a more appropriate MODE with pipes.
        """
        return pulumi.get(self, "pipe")

    @pipe.setter
    def pipe(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "pipe", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[bool]]:
        """
        Output version information and exit.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "version", value)


class Tee(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 append: Optional[pulumi.Input[bool]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 files: Optional[pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]] = None,
                 ignore_interrupts: Optional[pulumi.Input[bool]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 output_error: Optional[pulumi.Input['TeeMode']] = None,
                 pipe: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Read from standard input and write to standard output and files.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] append: Append to the given FILEs, do not overwrite.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system.
        :param pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]] files: The file(s) to write to.
        :param pulumi.Input[bool] ignore_interrupts: Ignore interrupt signals.
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run.
        :param pulumi.Input['TeeMode'] output_error: Set behavior on write error.
        :param pulumi.Input[bool] pipe: Operate in a more appropriate MODE with pipes.
        :param pulumi.Input[bool] version: Output version information and exit.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Read from standard input and write to standard output and files.

        :param str resource_name: The name of the resource.
        :param TeeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 append: Optional[pulumi.Input[bool]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 files: Optional[pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]] = None,
                 ignore_interrupts: Optional[pulumi.Input[bool]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 output_error: Optional[pulumi.Input['TeeMode']] = None,
                 pipe: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeeArgs.__new__(TeeArgs)

            __props__.__dict__["append"] = append
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["environment"] = environment
            if files is None and not opts.urn:
                raise TypeError("Missing required property 'files'")
            __props__.__dict__["files"] = files
            __props__.__dict__["ignore_interrupts"] = ignore_interrupts
            __props__.__dict__["lifecycle"] = lifecycle
            __props__.__dict__["output_error"] = output_error
            __props__.__dict__["pipe"] = pipe
            if stdin is None and not opts.urn:
                raise TypeError("Missing required property 'stdin'")
            __props__.__dict__["stdin"] = stdin
            __props__.__dict__["version"] = version
            __props__.__dict__["command"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Tee, __self__).__init__(
            'kubernetes-the-hard-way:tools:Tee',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def append(self) -> pulumi.Output[bool]:
        """
        Append to the given FILEs, do not overwrite.
        """
        return pulumi.get(self, "append")

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output['pulumi_command.remote.Command']:
        """
        Represents the command run on the remote system.
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        Connection details for the remote system.
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def files(self) -> pulumi.Output[Any]:
        """
        The file(s) to write to.
        """
        return pulumi.get(self, "files")

    @property
    @pulumi.getter(name="ignoreInterrupts")
    def ignore_interrupts(self) -> pulumi.Output[bool]:
        """
        Ignore interrupt signals.
        """
        return pulumi.get(self, "ignore_interrupts")

    @property
    @pulumi.getter
    def lifecycle(self) -> pulumi.Output[Optional['CommandLifecycle']]:
        """
        At what stage(s) in the resource lifecycle should the command be run.
        """
        return pulumi.get(self, "lifecycle")

    @property
    @pulumi.getter(name="outputError")
    def output_error(self) -> pulumi.Output[Optional['TeeMode']]:
        """
        Set behavior on write error.
        """
        return pulumi.get(self, "output_error")

    @property
    @pulumi.getter
    def pipe(self) -> pulumi.Output[bool]:
        """
        Operate in a more appropriate MODE with pipes.
        """
        return pulumi.get(self, "pipe")

    @property
    @pulumi.getter
    def stderr(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "stderr")

    @property
    @pulumi.getter
    def stdin(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stdin")

    @property
    @pulumi.getter
    def stdout(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "stdout")

