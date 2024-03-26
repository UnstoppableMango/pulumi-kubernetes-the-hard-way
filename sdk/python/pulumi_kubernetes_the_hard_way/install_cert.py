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
from .certificate import Certificate
from .remote_file import RemoteFile

__all__ = [
    'InstallCertResult',
    'AwaitableInstallCertResult',
    'install_cert',
    'install_cert_output',
]

@pulumi.output_type
class InstallCertResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, RemoteFile):
            raise TypeError("Expected argument 'result' to be a RemoteFile")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> 'RemoteFile':
        """
        A resource representing the the file on the remote machine.
        """
        return pulumi.get(self, "result")


class AwaitableInstallCertResult(InstallCertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstallCertResult(
            result=self.result)


def install_cert(cert: Optional['Certificate'] = None,
                 connection: Optional[pulumi.InputType['Connection']] = None,
                 name: Optional[str] = None,
                 options: Optional[pulumi.InputType['ResourceOptions']] = None,
                 path: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstallCertResult:
    """
    Creates a RemoteFile resource representing the copy operation.


    :param 'Certificate' cert: The certificate to install at the remote location.
    :param pulumi.InputType['Connection'] connection: The connection details.
    :param str path: The path to install to.
    """
    __args__ = dict()
    __args__['cert'] = cert
    __args__['connection'] = connection
    __args__['name'] = name
    __args__['options'] = options
    __args__['path'] = path
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kubernetes-the-hard-way:index:installCert', __args__, opts=opts, typ=InstallCertResult).value

    return AwaitableInstallCertResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(install_cert)
def install_cert_output(cert: Optional[pulumi.Input['Certificate']] = None,
                        connection: Optional[pulumi.Input[pulumi.InputType['Connection']]] = None,
                        name: Optional[pulumi.Input[str]] = None,
                        options: Optional[pulumi.Input[Optional[pulumi.InputType['ResourceOptions']]]] = None,
                        path: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstallCertResult]:
    """
    Creates a RemoteFile resource representing the copy operation.


    :param 'Certificate' cert: The certificate to install at the remote location.
    :param pulumi.InputType['Connection'] connection: The connection details.
    :param str path: The path to install to.
    """
    ...
