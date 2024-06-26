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

__all__ = ['CniBridgePluginConfigurationArgs', 'CniBridgePluginConfiguration']

@pulumi.input_type
class CniBridgePluginConfigurationArgs:
    def __init__(__self__, *,
                 subnet: pulumi.Input[str],
                 bridge: Optional[pulumi.Input[str]] = None,
                 cni_version: Optional[pulumi.Input[str]] = None,
                 ip_masq: Optional[pulumi.Input[bool]] = None,
                 ipam: Optional[pulumi.Input['CniBridgeIpamArgs']] = None,
                 is_gateway: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CniBridgePluginConfiguration resource.
        :param pulumi.Input[str] subnet: The subnet to use.
        :param pulumi.Input[str] bridge: Bridge name.
        :param pulumi.Input[str] cni_version: CNI version.
        :param pulumi.Input[bool] ip_masq: IP masq.
        :param pulumi.Input['CniBridgeIpamArgs'] ipam: IPAM
        :param pulumi.Input[bool] is_gateway: Is gateway.
        :param pulumi.Input[str] name: CNI plugin name.
        :param pulumi.Input[str] type: CNI plugin type.
        """
        pulumi.set(__self__, "subnet", subnet)
        if bridge is not None:
            pulumi.set(__self__, "bridge", bridge)
        if cni_version is not None:
            pulumi.set(__self__, "cni_version", cni_version)
        if ip_masq is not None:
            pulumi.set(__self__, "ip_masq", ip_masq)
        if ipam is not None:
            pulumi.set(__self__, "ipam", ipam)
        if is_gateway is not None:
            pulumi.set(__self__, "is_gateway", is_gateway)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Input[str]:
        """
        The subnet to use.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def bridge(self) -> Optional[pulumi.Input[str]]:
        """
        Bridge name.
        """
        return pulumi.get(self, "bridge")

    @bridge.setter
    def bridge(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bridge", value)

    @property
    @pulumi.getter(name="cniVersion")
    def cni_version(self) -> Optional[pulumi.Input[str]]:
        """
        CNI version.
        """
        return pulumi.get(self, "cni_version")

    @cni_version.setter
    def cni_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cni_version", value)

    @property
    @pulumi.getter(name="ipMasq")
    def ip_masq(self) -> Optional[pulumi.Input[bool]]:
        """
        IP masq.
        """
        return pulumi.get(self, "ip_masq")

    @ip_masq.setter
    def ip_masq(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ip_masq", value)

    @property
    @pulumi.getter
    def ipam(self) -> Optional[pulumi.Input['CniBridgeIpamArgs']]:
        """
        IPAM
        """
        return pulumi.get(self, "ipam")

    @ipam.setter
    def ipam(self, value: Optional[pulumi.Input['CniBridgeIpamArgs']]):
        pulumi.set(self, "ipam", value)

    @property
    @pulumi.getter(name="isGateway")
    def is_gateway(self) -> Optional[pulumi.Input[bool]]:
        """
        Is gateway.
        """
        return pulumi.get(self, "is_gateway")

    @is_gateway.setter
    def is_gateway(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_gateway", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        CNI plugin name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        CNI plugin type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class CniBridgePluginConfiguration(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bridge: Optional[pulumi.Input[str]] = None,
                 cni_version: Optional[pulumi.Input[str]] = None,
                 ip_masq: Optional[pulumi.Input[bool]] = None,
                 ipam: Optional[pulumi.Input[pulumi.InputType['CniBridgeIpamArgs']]] = None,
                 is_gateway: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Get the `bridge` configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bridge: Bridge name.
        :param pulumi.Input[str] cni_version: CNI version.
        :param pulumi.Input[bool] ip_masq: IP masq.
        :param pulumi.Input[pulumi.InputType['CniBridgeIpamArgs']] ipam: IPAM
        :param pulumi.Input[bool] is_gateway: Is gateway.
        :param pulumi.Input[str] name: CNI plugin name.
        :param pulumi.Input[str] subnet: The subnet to use.
        :param pulumi.Input[str] type: CNI plugin type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CniBridgePluginConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Get the `bridge` configuration.

        :param str resource_name: The name of the resource.
        :param CniBridgePluginConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CniBridgePluginConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bridge: Optional[pulumi.Input[str]] = None,
                 cni_version: Optional[pulumi.Input[str]] = None,
                 ip_masq: Optional[pulumi.Input[bool]] = None,
                 ipam: Optional[pulumi.Input[pulumi.InputType['CniBridgeIpamArgs']]] = None,
                 is_gateway: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CniBridgePluginConfigurationArgs.__new__(CniBridgePluginConfigurationArgs)

            __props__.__dict__["bridge"] = bridge
            __props__.__dict__["cni_version"] = cni_version
            __props__.__dict__["ip_masq"] = ip_masq
            __props__.__dict__["ipam"] = ipam
            __props__.__dict__["is_gateway"] = is_gateway
            __props__.__dict__["name"] = name
            if subnet is None and not opts.urn:
                raise TypeError("Missing required property 'subnet'")
            __props__.__dict__["subnet"] = subnet
            __props__.__dict__["type"] = type
            __props__.__dict__["result"] = None
            __props__.__dict__["yaml"] = None
        super(CniBridgePluginConfiguration, __self__).__init__(
            'kubernetes-the-hard-way:config:CniBridgePluginConfiguration',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output['outputs.CniBridgePluginConfiguration']:
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def yaml(self) -> pulumi.Output[str]:
        """
        The yaml representation of the manifest.
        """
        return pulumi.get(self, "yaml")

