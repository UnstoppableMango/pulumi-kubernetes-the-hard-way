# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *
from .remote_file import RemoteFile
import pulumi_tls

__all__ = [
    'InstallKeyResult',
    'AwaitableInstallKeyResult',
    'install_key',
    'install_key_output',
]

@pulumi.output_type
class InstallKeyResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, RemoteFile):
            raise TypeError("Expected argument 'result' to be a RemoteFile")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> 'RemoteFile':
        return pulumi.get(self, "result")


class AwaitableInstallKeyResult(InstallKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstallKeyResult(
            result=self.result)


def install_key(connection: Optional[pulumi.InputType['Connection']] = None,
                keypair: Optional[pulumi.InputType['KeyPair']] = None,
                name: Optional[str] = None,
                options: Optional[pulumi.InputType['ResourceOptions']] = None,
                path: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstallKeyResult:
    """
    Creates a RemoteFile resource representing the copy operation.


    :param pulumi.InputType['Connection'] connection: The connection details.
    :param pulumi.InputType['KeyPair'] keypair: The certificate to install.
    :param str path: The path to install to.
    """
    __args__ = dict()
    __args__['connection'] = connection
    __args__['keypair'] = keypair
    __args__['name'] = name
    __args__['options'] = options
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kubernetes-the-hard-way:index:installKey', __args__, opts=opts, typ=InstallKeyResult).value

    return AwaitableInstallKeyResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(install_key)
def install_key_output(connection: Optional[pulumi.Input[pulumi.InputType['Connection']]] = None,
                       keypair: Optional[pulumi.Input[pulumi.InputType['KeyPair']]] = None,
                       name: Optional[pulumi.Input[str]] = None,
                       options: Optional[pulumi.Input[Optional[pulumi.InputType['ResourceOptions']]]] = None,
                       path: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstallKeyResult]:
    """
    Creates a RemoteFile resource representing the copy operation.


    :param pulumi.InputType['Connection'] connection: The connection details.
    :param pulumi.InputType['KeyPair'] keypair: The certificate to install.
    :param str path: The path to install to.
    """
    ...
