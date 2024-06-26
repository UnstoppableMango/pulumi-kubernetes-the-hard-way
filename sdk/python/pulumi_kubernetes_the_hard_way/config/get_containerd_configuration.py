# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetContainerdConfigurationResult',
    'AwaitableGetContainerdConfigurationResult',
    'get_containerd_configuration',
    'get_containerd_configuration_output',
]

@pulumi.output_type
class GetContainerdConfigurationResult:
    """
    Get the containerd configuration.
    """
    def __init__(__self__, result=None):
        if result and not isinstance(result, dict):
            raise TypeError("Expected argument 'result' to be a dict")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> 'outputs.ContainerdConfiguration':
        return pulumi.get(self, "result")


class AwaitableGetContainerdConfigurationResult(GetContainerdConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerdConfigurationResult(
            result=self.result)


def get_containerd_configuration(cri: Optional[pulumi.InputType['ContainerdCriPluginConfiguration']] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerdConfigurationResult:
    """
    Get the containerd configuration.


    :param pulumi.InputType['ContainerdCriPluginConfiguration'] cri: The cri configuration.
    """
    __args__ = dict()
    __args__['cri'] = cri
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kubernetes-the-hard-way:config:getContainerdConfiguration', __args__, opts=opts, typ=GetContainerdConfigurationResult).value

    return AwaitableGetContainerdConfigurationResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(get_containerd_configuration)
def get_containerd_configuration_output(cri: Optional[pulumi.Input[Optional[pulumi.InputType['ContainerdCriPluginConfiguration']]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContainerdConfigurationResult]:
    """
    Get the containerd configuration.


    :param pulumi.InputType['ContainerdCriPluginConfiguration'] cri: The cri configuration.
    """
    ...
