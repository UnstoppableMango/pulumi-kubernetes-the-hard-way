# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import config as _config
from .. import tools as _tools
from .file import File
import pulumi_command
import pulumi_kubernetes

__all__ = ['StaticPodArgs', 'StaticPod']

@pulumi.input_type
class StaticPodArgs:
    def __init__(__self__, *,
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 pod: pulumi.Input['_config.PodManifestArgs'],
                 file_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StaticPod resource.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input['_config.PodManifestArgs'] pod: The pod manifest.
        :param pulumi.Input[str] file_name: The name of the file on the remote system.
        """
        pulumi.set(__self__, "connection", connection)
        pulumi.set(__self__, "pod", pod)
        if file_name is not None:
            pulumi.set(__self__, "file_name", file_name)

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
    def pod(self) -> pulumi.Input['_config.PodManifestArgs']:
        """
        The pod manifest.
        """
        return pulumi.get(self, "pod")

    @pod.setter
    def pod(self, value: pulumi.Input['_config.PodManifestArgs']):
        pulumi.set(self, "pod", value)

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the file on the remote system.
        """
        return pulumi.get(self, "file_name")

    @file_name.setter
    def file_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_name", value)


class StaticPod(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 pod: Optional[pulumi.Input[pulumi.InputType['_config.PodManifestArgs']]] = None,
                 __props__=None):
        """
        Create a static pod manifest on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: The parameters with which to connect to the remote host.
        :param pulumi.Input[str] file_name: The name of the file on the remote system.
        :param pulumi.Input[pulumi.InputType['_config.PodManifestArgs']] pod: The pod manifest.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StaticPodArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a static pod manifest on a remote system.

        :param str resource_name: The name of the resource.
        :param StaticPodArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StaticPodArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 pod: Optional[pulumi.Input[pulumi.InputType['_config.PodManifestArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StaticPodArgs.__new__(StaticPodArgs)

            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["file_name"] = file_name
            if pod is None and not opts.urn:
                raise TypeError("Missing required property 'pod'")
            __props__.__dict__["pod"] = pod
            __props__.__dict__["file"] = None
            __props__.__dict__["mkdir"] = None
            __props__.__dict__["path"] = None
        super(StaticPod, __self__).__init__(
            'kubernetes-the-hard-way:remote:StaticPod',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        The parameters with which to connect to the remote host.
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output['File']:
        """
        The remote manifest file.
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Output[str]:
        """
        The name of the file on the remote system.
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter
    def mkdir(self) -> pulumi.Output['_tools.Mkdir']:
        """
        The mkdir operation to ensure /etc/kubernetes/manifests exists.
        """
        return pulumi.get(self, "mkdir")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path to the manifest on the remote system.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def pod(self) -> pulumi.Output['_config.outputs.PodManifest']:
        """
        The pod manifest.
        """
        return pulumi.get(self, "pod")

