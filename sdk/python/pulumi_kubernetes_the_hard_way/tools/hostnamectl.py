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

__all__ = ['HostnamectlArgs', 'Hostnamectl']

@pulumi.input_type
class HostnamectlArgs:
    def __init__(__self__, *,
                 command: pulumi.Input['HostnamectlCommand'],
                 connection: pulumi.Input['pulumi_command.remote.ConnectionArgs'],
                 arg: Optional[pulumi.Input[str]] = None,
                 binary_path: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 json: Optional[pulumi.Input['HostnamectlJsonMode']] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 machine: Optional[pulumi.Input[str]] = None,
                 no_ask_password: Optional[pulumi.Input[bool]] = None,
                 pretty: Optional[pulumi.Input[bool]] = None,
                 static: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 transient: Optional[pulumi.Input[bool]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 version: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Hostnamectl resource.
        :param pulumi.Input['HostnamectlCommand'] command: Corresponds to the {COMMAND} argument.
        :param pulumi.Input['pulumi_command.remote.ConnectionArgs'] connection: Connection details for the remote system
        :param pulumi.Input[str] arg: The argument for the specified `command`.
        :param pulumi.Input[str] binary_path: Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment: Environment variables
        :param pulumi.Input[bool] help: Print a short help text and exit.
        :param pulumi.Input[str] host: Execute the operation remotely. Specify a hostname, or a username and hostname separated by '@', to connect to.
        :param pulumi.Input['HostnamectlJsonMode'] json: Shows output formatted as JSON.
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run
        :param pulumi.Input[str] machine: Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating '@' character.
        :param pulumi.Input[bool] no_ask_password: Do not query the user for authentication for privileged operations.
        :param pulumi.Input[bool] pretty: If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
        :param pulumi.Input[bool] static: If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
        :param pulumi.Input[str] stdin: TODO
        :param pulumi.Input[bool] transient: If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
        :param pulumi.Input[Sequence[Any]] triggers: TODO
        :param pulumi.Input[bool] version: Print a short version string and exit.
        """
        pulumi.set(__self__, "command", command)
        pulumi.set(__self__, "connection", connection)
        if arg is not None:
            pulumi.set(__self__, "arg", arg)
        if binary_path is not None:
            pulumi.set(__self__, "binary_path", binary_path)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if help is not None:
            pulumi.set(__self__, "help", help)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if json is not None:
            pulumi.set(__self__, "json", json)
        if lifecycle is not None:
            pulumi.set(__self__, "lifecycle", lifecycle)
        if machine is not None:
            pulumi.set(__self__, "machine", machine)
        if no_ask_password is not None:
            pulumi.set(__self__, "no_ask_password", no_ask_password)
        if pretty is not None:
            pulumi.set(__self__, "pretty", pretty)
        if static is not None:
            pulumi.set(__self__, "static", static)
        if stdin is not None:
            pulumi.set(__self__, "stdin", stdin)
        if transient is not None:
            pulumi.set(__self__, "transient", transient)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Input['HostnamectlCommand']:
        """
        Corresponds to the {COMMAND} argument.
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: pulumi.Input['HostnamectlCommand']):
        pulumi.set(self, "command", value)

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
    def arg(self) -> Optional[pulumi.Input[str]]:
        """
        The argument for the specified `command`.
        """
        return pulumi.get(self, "arg")

    @arg.setter
    def arg(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arg", value)

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
    def help(self) -> Optional[pulumi.Input[bool]]:
        """
        Print a short help text and exit.
        """
        return pulumi.get(self, "help")

    @help.setter
    def help(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "help", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Execute the operation remotely. Specify a hostname, or a username and hostname separated by '@', to connect to.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def json(self) -> Optional[pulumi.Input['HostnamectlJsonMode']]:
        """
        Shows output formatted as JSON.
        """
        return pulumi.get(self, "json")

    @json.setter
    def json(self, value: Optional[pulumi.Input['HostnamectlJsonMode']]):
        pulumi.set(self, "json", value)

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
    def machine(self) -> Optional[pulumi.Input[str]]:
        """
        Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating '@' character.
        """
        return pulumi.get(self, "machine")

    @machine.setter
    def machine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine", value)

    @property
    @pulumi.getter(name="noAskPassword")
    def no_ask_password(self) -> Optional[pulumi.Input[bool]]:
        """
        Do not query the user for authentication for privileged operations.
        """
        return pulumi.get(self, "no_ask_password")

    @no_ask_password.setter
    def no_ask_password(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_ask_password", value)

    @property
    @pulumi.getter
    def pretty(self) -> Optional[pulumi.Input[bool]]:
        """
        If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
        """
        return pulumi.get(self, "pretty")

    @pretty.setter
    def pretty(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "pretty", value)

    @property
    @pulumi.getter
    def static(self) -> Optional[pulumi.Input[bool]]:
        """
        If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
        """
        return pulumi.get(self, "static")

    @static.setter
    def static(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "static", value)

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
    def transient(self) -> Optional[pulumi.Input[bool]]:
        """
        If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
        """
        return pulumi.get(self, "transient")

    @transient.setter
    def transient(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transient", value)

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

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[bool]]:
        """
        Print a short version string and exit.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "version", value)


class Hostnamectl(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arg: Optional[pulumi.Input[str]] = None,
                 binary_path: Optional[pulumi.Input[str]] = None,
                 command: Optional[pulumi.Input['HostnamectlCommand']] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 json: Optional[pulumi.Input['HostnamectlJsonMode']] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 machine: Optional[pulumi.Input[str]] = None,
                 no_ask_password: Optional[pulumi.Input[bool]] = None,
                 pretty: Optional[pulumi.Input[bool]] = None,
                 static: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 transient: Optional[pulumi.Input[bool]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 version: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Abstraction over the `hostnamectl` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arg: The argument for the specified `command`.
        :param pulumi.Input[str] binary_path: Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
        :param pulumi.Input['HostnamectlCommand'] command: Corresponds to the {COMMAND} argument.
        :param pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']] connection: Connection details for the remote system
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] environment: Environment variables
        :param pulumi.Input[bool] help: Print a short help text and exit.
        :param pulumi.Input[str] host: Execute the operation remotely. Specify a hostname, or a username and hostname separated by '@', to connect to.
        :param pulumi.Input['HostnamectlJsonMode'] json: Shows output formatted as JSON.
        :param 'CommandLifecycle' lifecycle: At what stage(s) in the resource lifecycle should the command be run
        :param pulumi.Input[str] machine: Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating '@' character.
        :param pulumi.Input[bool] no_ask_password: Do not query the user for authentication for privileged operations.
        :param pulumi.Input[bool] pretty: If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
        :param pulumi.Input[bool] static: If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
        :param pulumi.Input[str] stdin: TODO
        :param pulumi.Input[bool] transient: If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
        :param pulumi.Input[Sequence[Any]] triggers: TODO
        :param pulumi.Input[bool] version: Print a short version string and exit.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostnamectlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Abstraction over the `hostnamectl` utility on a remote system.

        :param str resource_name: The name of the resource.
        :param HostnamectlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostnamectlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arg: Optional[pulumi.Input[str]] = None,
                 binary_path: Optional[pulumi.Input[str]] = None,
                 command: Optional[pulumi.Input['HostnamectlCommand']] = None,
                 connection: Optional[pulumi.Input[pulumi.InputType['pulumi_command.remote.ConnectionArgs']]] = None,
                 environment: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 json: Optional[pulumi.Input['HostnamectlJsonMode']] = None,
                 lifecycle: Optional['CommandLifecycle'] = None,
                 machine: Optional[pulumi.Input[str]] = None,
                 no_ask_password: Optional[pulumi.Input[bool]] = None,
                 pretty: Optional[pulumi.Input[bool]] = None,
                 static: Optional[pulumi.Input[bool]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 transient: Optional[pulumi.Input[bool]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
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
            __props__ = HostnamectlArgs.__new__(HostnamectlArgs)

            __props__.__dict__["arg"] = arg
            __props__.__dict__["binary_path"] = binary_path
            if command is None and not opts.urn:
                raise TypeError("Missing required property 'command'")
            __props__.__dict__["command"] = command
            if connection is None and not opts.urn:
                raise TypeError("Missing required property 'connection'")
            __props__.__dict__["connection"] = connection
            __props__.__dict__["environment"] = environment
            __props__.__dict__["help"] = help
            __props__.__dict__["host"] = host
            __props__.__dict__["json"] = json
            __props__.__dict__["lifecycle"] = lifecycle
            __props__.__dict__["machine"] = machine
            __props__.__dict__["no_ask_password"] = no_ask_password
            __props__.__dict__["pretty"] = pretty
            __props__.__dict__["static"] = static
            __props__.__dict__["stdin"] = stdin
            __props__.__dict__["transient"] = transient
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["version"] = version
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Hostnamectl, __self__).__init__(
            'kubernetes-the-hard-way:tools:Hostnamectl',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def arg(self) -> pulumi.Output[Optional[str]]:
        """
        The argument for the specified `command`.
        """
        return pulumi.get(self, "arg")

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
    @pulumi.getter
    def environment(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Environment variables
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def help(self) -> pulumi.Output[bool]:
        """
        Print a short help text and exit.
        """
        return pulumi.get(self, "help")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        Execute the operation remotely. Specify a hostname, or a username and hostname separated by '@', to connect to.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def json(self) -> pulumi.Output['HostnamectlJsonMode']:
        """
        Shows output formatted as JSON.
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def lifecycle(self) -> pulumi.Output[Optional['CommandLifecycle']]:
        """
        At what stage(s) in the resource lifecycle should the command be run
        """
        return pulumi.get(self, "lifecycle")

    @property
    @pulumi.getter
    def machine(self) -> pulumi.Output[Optional[str]]:
        """
        Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating '@' character.
        """
        return pulumi.get(self, "machine")

    @property
    @pulumi.getter(name="noAskPassword")
    def no_ask_password(self) -> pulumi.Output[bool]:
        """
        Do not query the user for authentication for privileged operations.
        """
        return pulumi.get(self, "no_ask_password")

    @property
    @pulumi.getter
    def pretty(self) -> pulumi.Output[bool]:
        """
        If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`.
        """
        return pulumi.get(self, "pretty")

    @property
    @pulumi.getter
    def static(self) -> pulumi.Output[bool]:
        """
        If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`.
        """
        return pulumi.get(self, "static")

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
    def transient(self) -> pulumi.Output[bool]:
        """
        If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`.
        """
        return pulumi.get(self, "transient")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Sequence[Any]]:
        """
        TODO
        """
        return pulumi.get(self, "triggers")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[bool]:
        """
        Print a short version string and exit.
        """
        return pulumi.get(self, "version")

