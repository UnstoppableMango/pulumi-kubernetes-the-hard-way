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

__all__ = ['EtcdctlArgs', 'Etcdctl']

@pulumi.input_type
class EtcdctlArgs:
    def __init__(__self__, *,
                 commands: pulumi.Input['EtcdctlCommand'],
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 binary_path: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None):
        """
        The set of arguments for constructing a Etcdctl resource.
        :param pulumi.Input['EtcdctlCommand'] commands: TODO
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system
        :param pulumi.Input[str] binary_path: Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        :param pulumi.Input[str] ca_cert: TODO
        :param pulumi.Input[str] cert: TODO
        :param pulumi.Input[str] endpoints: TODO
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment: Environment variables
        :param pulumi.Input[str] key: TODO
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run
        :param pulumi.Input[str] stdin: TODO
        :param pulumi.Input[Sequence[Any]] triggers: TODO
        """
        pulumi.set(__self__, "commands", commands)
        pulumi.set(__self__, "connection", connection)
        if binary_path is not None:
            pulumi.set(__self__, "binary_path", binary_path)
        if ca_cert is not None:
            pulumi.set(__self__, "ca_cert", ca_cert)
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if lifecycle is not None:
            pulumi.set(__self__, "lifecycle", lifecycle)
        if stdin is not None:
            pulumi.set(__self__, "stdin", stdin)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def commands(self) -> pulumi.Input['EtcdctlCommand']:
        """
        TODO
        """
        return pulumi.get(self, "commands")

    @commands.setter
    def commands(self, value: pulumi.Input['EtcdctlCommand']):
        pulumi.set(self, "commands", value)

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
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "ca_cert")

    @ca_cert.setter
    def ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert", value)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoints", value)

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
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

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
    def triggers(self) -> Optional[pulumi.Input[Sequence[Any]]]:
        """
        TODO
        """
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Sequence[Any]]]):
        pulumi.set(self, "triggers", value)


class Etcdctl(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 binary_path: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 commands: Optional[pulumi.Input['EtcdctlCommand']] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 endpoints: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 __props__=None):
        """
        Abstraction over the `etcdctl` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] binary_path: Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        :param pulumi.Input[str] ca_cert: TODO
        :param pulumi.Input[str] cert: TODO
        :param pulumi.Input['EtcdctlCommand'] commands: TODO
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system
        :param pulumi.Input[str] endpoints: TODO
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment: Environment variables
        :param pulumi.Input[str] key: TODO
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run
        :param pulumi.Input[str] stdin: TODO
        :param pulumi.Input[Sequence[Any]] triggers: TODO
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EtcdctlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Abstraction over the `etcdctl` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param EtcdctlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EtcdctlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 binary_path: Optional[pulumi.Input[str]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 commands: Optional[pulumi.Input['EtcdctlCommand']] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 endpoints: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EtcdctlArgs.__new__(EtcdctlArgs)

            __props__.__dict__["binary_path"] = binary_path
            __props__.__dict__["ca_cert"] = ca_cert
            __props__.__dict__["cert"] = cert
            if commands is None and not opts.urn:
                raise TypeError("Missing required property 'commands'")
            __props__.__dict__["commands"] = commands
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["endpoints"] = endpoints
            __props__.__dict__["environment"] = environment
            __props__.__dict__["key"] = key
            __props__.__dict__["lifecycle"] = lifecycle
            __props__.__dict__["stdin"] = stdin
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["command"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Etcdctl, __self__).__init__(
            'kubernetes-the-hard-way:tools:Etcdctl',
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
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> pulumi.Output[Optional[str]]:
        """
        TODO
        """
        return pulumi.get(self, "ca_cert")

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Output[Optional[str]]:
        """
        TODO
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output['pulumi_command.remote.Command']:
        """
        The underlying command
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def commands(self) -> pulumi.Output['EtcdctlCommand']:
        """
        TODO
        """
        return pulumi.get(self, "commands")

    @property
    @pulumi.getter
    def connection(self) -> pulumi.Output['pulumi_command.remote.outputs.Connection']:
        """
        Connection details for the remote system
        """
        return pulumi.get(self, "connection")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Optional[str]]:
        """
        TODO
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Environment variables
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[str]]:
        """
        TODO
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def lifecycle(self) -> pulumi.Output[Optional['CommandLifecycle']]:
        """
        At what stage(s) in the resource lifecycle should the command be run
        """
        return pulumi.get(self, "lifecycle")

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
    def triggers(self) -> pulumi.Output[Sequence[Any]]:
        """
        TODO
        """
        return pulumi.get(self, "triggers")

