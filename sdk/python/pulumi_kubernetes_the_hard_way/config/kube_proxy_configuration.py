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

__all__ = ['KubeProxyConfigurationArgs', 'KubeProxyConfiguration']

@pulumi.input_type
class KubeProxyConfigurationArgs:
    def __init__(__self__, *,
                 cluster_cidr: pulumi.Input[str],
                 kubeconfig: pulumi.Input[str],
                 mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a KubeProxyConfiguration resource.
        :param pulumi.Input[str] cluster_cidr: Cluster CIDR.
        :param pulumi.Input[str] kubeconfig: Path to the kubeconfig.
        :param pulumi.Input[str] mode: TODO
        """
        pulumi.set(__self__, "cluster_cidr", cluster_cidr)
        pulumi.set(__self__, "kubeconfig", kubeconfig)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter(name="clusterCIDR")
    def cluster_cidr(self) -> pulumi.Input[str]:
        """
        Cluster CIDR.
        """
        return pulumi.get(self, "cluster_cidr")

    @cluster_cidr.setter
    def cluster_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_cidr", value)

    @property
    @pulumi.getter
    def kubeconfig(self) -> pulumi.Input[str]:
        """
        Path to the kubeconfig.
        """
        return pulumi.get(self, "kubeconfig")

    @kubeconfig.setter
    def kubeconfig(self, value: pulumi.Input[str]):
        pulumi.set(self, "kubeconfig", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)


class KubeProxyConfiguration(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_cidr: Optional[pulumi.Input[str]] = None,
                 kubeconfig: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        kube-proxy configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_cidr: Cluster CIDR.
        :param pulumi.Input[str] kubeconfig: Path to the kubeconfig.
        :param pulumi.Input[str] mode: TODO
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubeProxyConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        kube-proxy configuration.

        :param str resource_name: The name of the resource.
        :param KubeProxyConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubeProxyConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_cidr: Optional[pulumi.Input[str]] = None,
                 kubeconfig: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubeProxyConfigurationArgs.__new__(KubeProxyConfigurationArgs)

            if cluster_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_cidr'")
            __props__.__dict__["cluster_cidr"] = cluster_cidr
            if kubeconfig is None and not opts.urn:
                raise TypeError("Missing required property 'kubeconfig'")
            __props__.__dict__["kubeconfig"] = kubeconfig
            __props__.__dict__["mode"] = mode
            __props__.__dict__["result"] = None
            __props__.__dict__["yaml"] = None
        super(KubeProxyConfiguration, __self__).__init__(
            'kubernetes-the-hard-way:config:KubeProxyConfiguration',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output['outputs.KubeProxyConfiguration']:
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def yaml(self) -> pulumi.Output[str]:
        """
        The yaml representation of the manifest.
        """
        return pulumi.get(self, "yaml")

