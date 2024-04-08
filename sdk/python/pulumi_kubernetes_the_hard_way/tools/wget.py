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

__all__ = ['WgetArgs', 'Wget']

@pulumi.input_type
class WgetArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 url: pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]],
                 binary_path: Optional[pulumi.Input[str]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 no_verbose: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None):
        """
        The set of arguments for constructing a Wget resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system
        :param pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]] url: Corresponds to the [URL...] argument.
        :param pulumi.Input[str] binary_path: Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        :param pulumi.Input[str] directory_prefix: The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment: Environment variables
        :param pulumi.Input[bool] https_only: When in recursive mode, only HTTPS links are followed.
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run
        :param pulumi.Input[bool] no_verbose: Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
        :param pulumi.Input[str] output_document: The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
        :param pulumi.Input[bool] quiet: Turn off Wget's output.
        :param pulumi.Input[str] stdin: TODO
        :param pulumi.Input[bool] timestamping: Turn on time-stamping.
        :param pulumi.Input[Sequence[Any]] triggers: TODO
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "url", url)
        if binary_path is not None:
            pulumi.set(__self__, "binary_path", binary_path)
        if directory_prefix is not None:
            pulumi.set(__self__, "directory_prefix", directory_prefix)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if https_only is not None:
            pulumi.set(__self__, "https_only", https_only)
        if lifecycle is not None:
            pulumi.set(__self__, "lifecycle", lifecycle)
        if no_verbose is not None:
            pulumi.set(__self__, "no_verbose", no_verbose)
        if output_document is not None:
            pulumi.set(__self__, "output_document", output_document)
        if quiet is not None:
            pulumi.set(__self__, "quiet", quiet)
        if stdin is not None:
            pulumi.set(__self__, "stdin", stdin)
        if timestamping is not None:
            pulumi.set(__self__, "timestamping", timestamping)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Input['pulumi_command.remote.ConnectionArgs']:
        """
        Connection details for the remote system
        """
        return pulumi.get(self, "connection")

    @connection.setter
    def connection(self, value: pulumi.Input['pulumi_command.remote.ConnectionArgs']):
        pulumi.set(self, "connection", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]:
        """
        Corresponds to the [URL...] argument.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="binaryPath")
    def binary_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        """
        return pulumi.get(self, "binary_path")

    @binary_path.setter
    def binary_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "binary_path", value)

    @property
    @pulumi.getter(name="directoryPrefix")
    def directory_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
        """
        return pulumi.get(self, "directory_prefix")

    @directory_prefix.setter
    def directory_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_prefix", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Environment variables
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[pulumi.Input[bool]]:
        """
        When in recursive mode, only HTTPS links are followed.
        """
        return pulumi.get(self, "https_only")

    @https_only.setter
    def https_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_only", value)

    @property
    @pulumi.getter
    def lifecycle(self) -> Optional['CommandLifecycle']:
        """
        At what stage(s) in the resource lifecycle should the command be run
        """
        return pulumi.get(self, "lifecycle")

    @lifecycle.setter
    def lifecycle(self, value: Optional['CommandLifecycle']):
        pulumi.set(self, "lifecycle", value)

    @property
    @pulumi.getter(name="noVerbose")
    def no_verbose(self) -> Optional[pulumi.Input[bool]]:
        """
        Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
        """
        return pulumi.get(self, "no_verbose")

    @no_verbose.setter
    def no_verbose(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_verbose", value)

    @property
    @pulumi.getter(name="outputDocument")
    def output_document(self) -> Optional[pulumi.Input[str]]:
        """
        The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
        """
        return pulumi.get(self, "output_document")

    @output_document.setter
    def output_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_document", value)

    @property
    @pulumi.getter
    def quiet(self) -> Optional[pulumi.Input[bool]]:
        """
        Turn off Wget's output.
        """
        return pulumi.get(self, "quiet")

    @quiet.setter
    def quiet(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "quiet", value)

    @property
    @pulumi.getter
    def stdin(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "stdin")

    @stdin.setter
    def stdin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stdin", value)

    @property
    @pulumi.getter
    def timestamping(self) -> Optional[pulumi.Input[bool]]:
        """
        Turn on time-stamping.
        """
        return pulumi.get(self, "timestamping")

    @timestamping.setter
    def timestamping(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "timestamping", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Sequence[Any]]]:
        """
        TODO
        """
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Sequence[Any]]]):
        pulumi.set(self, "triggers", value)


class Wget(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 binary_path: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 no_verbose: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 url: Optional[pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]] = None,
                 __props__=None):
        """
        Abstraction over the `wget` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] binary_path: Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system
        :param pulumi.Input[str] directory_prefix: The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment: Environment variables
        :param pulumi.Input[bool] https_only: When in recursive mode, only HTTPS links are followed.
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run
        :param pulumi.Input[bool] no_verbose: Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
        :param pulumi.Input[str] output_document: The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
        :param pulumi.Input[bool] quiet: Turn off Wget's output.
        :param pulumi.Input[str] stdin: TODO
        :param pulumi.Input[bool] timestamping: Turn on time-stamping.
        :param pulumi.Input[Sequence[Any]] triggers: TODO
        :param pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]] url: Corresponds to the [URL...] argument.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WgetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Abstraction over the `wget` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param WgetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WgetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 binary_path: Optional[pulumi.Input[str]] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 no_verbose: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 url: Optional[pulumi.Input[Union[str, Sequence[pulumi.Input[str]]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WgetArgs.__new__(WgetArgs)

            __props__.__dict__["binary_path"] = binary_path
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["directory_prefix"] = directory_prefix
            __props__.__dict__["environment"] = environment
            __props__.__dict__["https_only"] = https_only
            __props__.__dict__["lifecycle"] = lifecycle
            __props__.__dict__["no_verbose"] = no_verbose
            __props__.__dict__["output_document"] = output_document
            __props__.__dict__["quiet"] = quiet
            __props__.__dict__["stdin"] = stdin
            __props__.__dict__["timestamping"] = timestamping
            __props__.__dict__["triggers"] = triggers
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["command"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Wget, __self__).__init__(
            'kubernetes-the-hard-way:tools:Wget',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter(name="binaryPath")
    def binary_path(self) -> pulumi.Output[str]:
        """
        Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        """
        return pulumi.get(self, "binary_path")

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output['pulumi_command.remote.Command']:
        """
        The underlying command
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        Connection details for the remote system
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter(name="directoryPrefix")
    def directory_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).
        """
        return pulumi.get(self, "directory_prefix")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Environment variables
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> pulumi.Output[bool]:
        """
        When in recursive mode, only HTTPS links are followed.
        """
        return pulumi.get(self, "https_only")

    @property
    @pulumi.getter
    def lifecycle(self) -> pulumi.Output[Optional['CommandLifecycle']]:
        """
        At what stage(s) in the resource lifecycle should the command be run
        """
        return pulumi.get(self, "lifecycle")

    @property
    @pulumi.getter(name="noVerbose")
    def no_verbose(self) -> pulumi.Output[bool]:
        """
        Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.
        """
        return pulumi.get(self, "no_verbose")

    @property
    @pulumi.getter(name="outputDocument")
    def output_document(self) -> pulumi.Output[Optional[str]]:
        """
        The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.
        """
        return pulumi.get(self, "output_document")

    @property
    @pulumi.getter
    def quiet(self) -> pulumi.Output[bool]:
        """
        Turn off Wget's output.
        """
        return pulumi.get(self, "quiet")

    @property
    @pulumi.getter
    def stderr(self) -> pulumi.Output[str]:
        """
        TODO
        """
        return pulumi.get(self, "stderr")

    @property
    @pulumi.getter
    def stdin(self) -> pulumi.Output[Optional[str]]:
        """
        TODO
        """
        return pulumi.get(self, "stdin")

    @property
    @pulumi.getter
    def stdout(self) -> pulumi.Output[str]:
        """
        TODO
        """
        return pulumi.get(self, "stdout")

    @property
    @pulumi.getter
    def timestamping(self) -> pulumi.Output[bool]:
        """
        Turn on time-stamping.
        """
        return pulumi.get(self, "timestamping")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Sequence[Any]]:
        """
        TODO
        """
        return pulumi.get(self, "triggers")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Any]:
        """
        Corresponds to the [URL...] argument.
        """
        return pulumi.get(self, "url")

