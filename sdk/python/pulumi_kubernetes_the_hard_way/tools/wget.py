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

__all__ = ['WgetArgs', 'Wget']

@pulumi.input_type
class WgetArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 url: pulumi.Input[str],
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Wget resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system.
        :param pulumi.Input[str] url: Corresponse to the [URL] argument.
        :param pulumi.Input[str] directory_prefix: Corresponds to the --directory-prefix option.
        :param pulumi.Input[bool] https_only: Corresponds to the --https-only option.
        :param pulumi.Input[str] output_document: Corresponds to the --output-document option.
        :param pulumi.Input[bool] quiet: Corresponds to the --quiet option.
        :param pulumi.Input[bool] timestamping: Corresponds to the --timestamping option.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "url", url)
        if directory_prefix is not None:
            pulumi.set(__self__, "directory_prefix", directory_prefix)
        if https_only is not None:
            pulumi.set(__self__, "https_only", https_only)
        if output_document is not None:
            pulumi.set(__self__, "output_document", output_document)
        if quiet is not None:
            pulumi.set(__self__, "quiet", quiet)
        if timestamping is not None:
            pulumi.set(__self__, "timestamping", timestamping)

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
    def url(self) -> pulumi.Input[str]:
        """
        Corresponse to the [URL] argument.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="directoryPrefix")
    def directory_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Corresponds to the --directory-prefix option.
        """
        return pulumi.get(self, "directory_prefix")

    @directory_prefix.setter
    def directory_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_prefix", value)

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the --https-only option.
        """
        return pulumi.get(self, "https_only")

    @https_only.setter
    def https_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_only", value)

    @property
    @pulumi.getter(name="outputDocument")
    def output_document(self) -> Optional[pulumi.Input[str]]:
        """
        Corresponds to the --output-document option.
        """
        return pulumi.get(self, "output_document")

    @output_document.setter
    def output_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_document", value)

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
    def timestamping(self) -> Optional[pulumi.Input[bool]]:
        """
        Corresponds to the --timestamping option.
        """
        return pulumi.get(self, "timestamping")

    @timestamping.setter
    def timestamping(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "timestamping", value)


class Wget(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Abstraction over the `wget` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system.
        :param pulumi.Input[str] directory_prefix: Corresponds to the --directory-prefix option.
        :param pulumi.Input[bool] https_only: Corresponds to the --https-only option.
        :param pulumi.Input[str] output_document: Corresponds to the --output-document option.
        :param pulumi.Input[bool] quiet: Corresponds to the --quiet option.
        :param pulumi.Input[bool] timestamping: Corresponds to the --timestamping option.
        :param pulumi.Input[str] url: Corresponse to the [URL] argument.
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
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
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

            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["directory_prefix"] = directory_prefix
            __props__.__dict__["https_only"] = https_only
            __props__.__dict__["output_document"] = output_document
            __props__.__dict__["quiet"] = quiet
            __props__.__dict__["timestamping"] = timestamping
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["command"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdin"] = None
            __props__.__dict__["stdout"] = None
        super(Wget, __self__).__init__(
            'kubernetes-the-hard-way:tools:Wget',
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
    @pulumi.getter(name="directoryPrefix")
    def directory_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Corresponds to the --directory-prefix option.
        """
        return pulumi.get(self, "directory_prefix")

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> pulumi.Output[bool]:
        """
        Corresponds to the --https-only option.
        """
        return pulumi.get(self, "https_only")

    @property
    @pulumi.getter(name="outputDocument")
    def output_document(self) -> pulumi.Output[Optional[str]]:
        """
        Corresponds to the --output-document option.
        """
        return pulumi.get(self, "output_document")

    @property
    @pulumi.getter
    def quiet(self) -> pulumi.Output[bool]:
        """
        Corresponds to the --quiet option.
        """
        return pulumi.get(self, "quiet")

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
    @pulumi.getter
    def timestamping(self) -> pulumi.Output[bool]:
        """
        Corresponds to the --timestamping option.
        """
        return pulumi.get(self, "timestamping")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        Corresponse to the [URL] argument.
        """
        return pulumi.get(self, "url")
