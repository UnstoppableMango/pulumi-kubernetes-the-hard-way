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

__all__ = [
    'ClusterPkiNodeArgs',
]

@pulumi.input_type
class ClusterPkiNodeArgs:
    def __init__(__self__, *,
                 ip: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input['NodeRole']] = None):
        """
        TODO
        :param pulumi.Input[str] ip: The IP address of the node
        :param pulumi.Input['NodeRole'] role: The role a node should be configured for
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the node
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input['NodeRole']]:
        """
        The role a node should be configured for
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input['NodeRole']]):
        pulumi.set(self, "role", value)


