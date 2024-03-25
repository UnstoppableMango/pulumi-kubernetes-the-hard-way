# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from . import _utilities
from ._inputs import *
from .cluster_pki import ClusterPki

__all__ = [
    'install_control_plane',
]

def install_control_plane(connection: Optional[pulumi.InputType['Connection']] = None,
                          opts: Optional[pulumi.InputType['ResourceOptions']] = None,
                          pki: Optional['ClusterPki'] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> Awaitable[None]:
    """
    Install PKI onto a controlplane node.


    :param pulumi.InputType['Connection'] connection: The connection details.
    :param 'ClusterPki' pki: The PKI to install
    """
    __args__ = dict()
    __args__['connection'] = connection
    __args__['opts'] = opts
    __args__['pki'] = pki
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kubernetes-the-hard-way:index:installControlPlane', __args__, opts=opts).value

