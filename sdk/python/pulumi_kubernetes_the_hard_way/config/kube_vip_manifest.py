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
import pulumi_kubernetes

__all__ = ['KubeVipManifestArgs', 'KubeVipManifest']

@pulumi.input_type
class KubeVipManifestArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 kubeconfig_path: pulumi.Input[str],
                 vip_cidr: pulumi.Input[int],
                 bgp_as: Optional[pulumi.Input[int]] = None,
                 bgp_enable: Optional[pulumi.Input[bool]] = None,
                 bgp_peer_address: Optional[pulumi.Input[str]] = None,
                 bgp_peer_as: Optional[pulumi.Input[int]] = None,
                 bgp_peer_pass: Optional[pulumi.Input[str]] = None,
                 bgp_peers: Optional[pulumi.Input[str]] = None,
                 bgp_router_id: Optional[pulumi.Input[str]] = None,
                 cp_enable: Optional[pulumi.Input[bool]] = None,
                 cp_namespace: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 svc_enable: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 vip_arp: Optional[pulumi.Input[bool]] = None,
                 vip_ddns: Optional[pulumi.Input[bool]] = None,
                 vip_interface: Optional[pulumi.Input[str]] = None,
                 vip_leader_election: Optional[pulumi.Input[bool]] = None,
                 vip_lease_duration: Optional[pulumi.Input[int]] = None,
                 vip_renew_deadline: Optional[pulumi.Input[int]] = None,
                 vip_retry_period: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a KubeVipManifest resource.
        :param pulumi.Input[str] address: TODO
        :param pulumi.Input[str] kubeconfig_path: Path to the kubeconfig on the remote host.
        :param pulumi.Input[int] vip_cidr: TODO
        :param pulumi.Input[int] bgp_as: TODO
        :param pulumi.Input[bool] bgp_enable: TODO
        :param pulumi.Input[str] bgp_peer_address: TODO
        :param pulumi.Input[int] bgp_peer_as: TODO
        :param pulumi.Input[str] bgp_peer_pass: TODO
        :param pulumi.Input[str] bgp_peers: TODO
        :param pulumi.Input[str] bgp_router_id: TODO
        :param pulumi.Input[bool] cp_enable: TODO
        :param pulumi.Input[str] cp_namespace: TODO
        :param pulumi.Input[str] image: Override the kube-vip image.
        :param pulumi.Input[str] name: Name of the static pod. Defaults to kube-vip.
        :param pulumi.Input[str] namespace: Namespace for the static pod. Defaults to kube-system.
        :param pulumi.Input[int] port: TODO
        :param pulumi.Input[bool] svc_enable: TODO
        :param pulumi.Input[str] version: Version of kube-vip to use.
        :param pulumi.Input[bool] vip_arp: TODO
        :param pulumi.Input[bool] vip_ddns: TODO
        :param pulumi.Input[str] vip_interface: TODO
        :param pulumi.Input[bool] vip_leader_election: TODO
        :param pulumi.Input[int] vip_lease_duration: TODO
        :param pulumi.Input[int] vip_renew_deadline: TODO
        :param pulumi.Input[int] vip_retry_period: TODO
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "kubeconfig_path", kubeconfig_path)
        pulumi.set(__self__, "vip_cidr", vip_cidr)
        if bgp_as is not None:
            pulumi.set(__self__, "bgp_as", bgp_as)
        if bgp_enable is not None:
            pulumi.set(__self__, "bgp_enable", bgp_enable)
        if bgp_peer_address is not None:
            pulumi.set(__self__, "bgp_peer_address", bgp_peer_address)
        if bgp_peer_as is not None:
            pulumi.set(__self__, "bgp_peer_as", bgp_peer_as)
        if bgp_peer_pass is not None:
            pulumi.set(__self__, "bgp_peer_pass", bgp_peer_pass)
        if bgp_peers is not None:
            pulumi.set(__self__, "bgp_peers", bgp_peers)
        if bgp_router_id is not None:
            pulumi.set(__self__, "bgp_router_id", bgp_router_id)
        if cp_enable is not None:
            pulumi.set(__self__, "cp_enable", cp_enable)
        if cp_namespace is not None:
            pulumi.set(__self__, "cp_namespace", cp_namespace)
        if image is not None:
            pulumi.set(__self__, "image", image)
        if name is None:
            name = 'kube-vip'
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is None:
            namespace = 'kube-system'
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if port is None:
            port = 6443
        if port is not None:
            pulumi.set(__self__, "port", port)
        if svc_enable is not None:
            pulumi.set(__self__, "svc_enable", svc_enable)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if vip_arp is not None:
            pulumi.set(__self__, "vip_arp", vip_arp)
        if vip_ddns is not None:
            pulumi.set(__self__, "vip_ddns", vip_ddns)
        if vip_interface is not None:
            pulumi.set(__self__, "vip_interface", vip_interface)
        if vip_leader_election is not None:
            pulumi.set(__self__, "vip_leader_election", vip_leader_election)
        if vip_lease_duration is not None:
            pulumi.set(__self__, "vip_lease_duration", vip_lease_duration)
        if vip_renew_deadline is not None:
            pulumi.set(__self__, "vip_renew_deadline", vip_renew_deadline)
        if vip_retry_period is not None:
            pulumi.set(__self__, "vip_retry_period", vip_retry_period)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        TODO
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="kubeconfigPath")
    def kubeconfig_path(self) -> pulumi.Input[str]:
        """
        Path to the kubeconfig on the remote host.
        """
        return pulumi.get(self, "kubeconfig_path")

    @kubeconfig_path.setter
    def kubeconfig_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "kubeconfig_path", value)

    @property
    @pulumi.getter(name="vipCidr")
    def vip_cidr(self) -> pulumi.Input[int]:
        """
        TODO
        """
        return pulumi.get(self, "vip_cidr")

    @vip_cidr.setter
    def vip_cidr(self, value: pulumi.Input[int]):
        pulumi.set(self, "vip_cidr", value)

    @property
    @pulumi.getter(name="bgpAs")
    def bgp_as(self) -> Optional[pulumi.Input[int]]:
        """
        TODO
        """
        return pulumi.get(self, "bgp_as")

    @bgp_as.setter
    def bgp_as(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bgp_as", value)

    @property
    @pulumi.getter(name="bgpEnable")
    def bgp_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO
        """
        return pulumi.get(self, "bgp_enable")

    @bgp_enable.setter
    def bgp_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bgp_enable", value)

    @property
    @pulumi.getter(name="bgpPeerAddress")
    def bgp_peer_address(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "bgp_peer_address")

    @bgp_peer_address.setter
    def bgp_peer_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_peer_address", value)

    @property
    @pulumi.getter(name="bgpPeerAs")
    def bgp_peer_as(self) -> Optional[pulumi.Input[int]]:
        """
        TODO
        """
        return pulumi.get(self, "bgp_peer_as")

    @bgp_peer_as.setter
    def bgp_peer_as(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bgp_peer_as", value)

    @property
    @pulumi.getter(name="bgpPeerPass")
    def bgp_peer_pass(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "bgp_peer_pass")

    @bgp_peer_pass.setter
    def bgp_peer_pass(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_peer_pass", value)

    @property
    @pulumi.getter(name="bgpPeers")
    def bgp_peers(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "bgp_peers")

    @bgp_peers.setter
    def bgp_peers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_peers", value)

    @property
    @pulumi.getter(name="bgpRouterId")
    def bgp_router_id(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "bgp_router_id")

    @bgp_router_id.setter
    def bgp_router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_router_id", value)

    @property
    @pulumi.getter(name="cpEnable")
    def cp_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO
        """
        return pulumi.get(self, "cp_enable")

    @cp_enable.setter
    def cp_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cp_enable", value)

    @property
    @pulumi.getter(name="cpNamespace")
    def cp_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "cp_namespace")

    @cp_namespace.setter
    def cp_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cp_namespace", value)

    @property
    @pulumi.getter
    def image(self) -> Optional[pulumi.Input[str]]:
        """
        Override the kube-vip image.
        """
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the static pod. Defaults to kube-vip.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace for the static pod. Defaults to kube-system.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        TODO
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="svcEnable")
    def svc_enable(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO
        """
        return pulumi.get(self, "svc_enable")

    @svc_enable.setter
    def svc_enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "svc_enable", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of kube-vip to use.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="vipArp")
    def vip_arp(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO
        """
        return pulumi.get(self, "vip_arp")

    @vip_arp.setter
    def vip_arp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "vip_arp", value)

    @property
    @pulumi.getter(name="vipDdns")
    def vip_ddns(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO
        """
        return pulumi.get(self, "vip_ddns")

    @vip_ddns.setter
    def vip_ddns(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "vip_ddns", value)

    @property
    @pulumi.getter(name="vipInterface")
    def vip_interface(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "vip_interface")

    @vip_interface.setter
    def vip_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip_interface", value)

    @property
    @pulumi.getter(name="vipLeaderElection")
    def vip_leader_election(self) -> Optional[pulumi.Input[bool]]:
        """
        TODO
        """
        return pulumi.get(self, "vip_leader_election")

    @vip_leader_election.setter
    def vip_leader_election(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "vip_leader_election", value)

    @property
    @pulumi.getter(name="vipLeaseDuration")
    def vip_lease_duration(self) -> Optional[pulumi.Input[int]]:
        """
        TODO
        """
        return pulumi.get(self, "vip_lease_duration")

    @vip_lease_duration.setter
    def vip_lease_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vip_lease_duration", value)

    @property
    @pulumi.getter(name="vipRenewDeadline")
    def vip_renew_deadline(self) -> Optional[pulumi.Input[int]]:
        """
        TODO
        """
        return pulumi.get(self, "vip_renew_deadline")

    @vip_renew_deadline.setter
    def vip_renew_deadline(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vip_renew_deadline", value)

    @property
    @pulumi.getter(name="vipRetryPeriod")
    def vip_retry_period(self) -> Optional[pulumi.Input[int]]:
        """
        TODO
        """
        return pulumi.get(self, "vip_retry_period")

    @vip_retry_period.setter
    def vip_retry_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vip_retry_period", value)


class KubeVipManifest(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 bgp_as: Optional[pulumi.Input[int]] = None,
                 bgp_enable: Optional[pulumi.Input[bool]] = None,
                 bgp_peer_address: Optional[pulumi.Input[str]] = None,
                 bgp_peer_as: Optional[pulumi.Input[int]] = None,
                 bgp_peer_pass: Optional[pulumi.Input[str]] = None,
                 bgp_peers: Optional[pulumi.Input[str]] = None,
                 bgp_router_id: Optional[pulumi.Input[str]] = None,
                 cp_enable: Optional[pulumi.Input[bool]] = None,
                 cp_namespace: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[str]] = None,
                 kubeconfig_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 svc_enable: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 vip_arp: Optional[pulumi.Input[bool]] = None,
                 vip_cidr: Optional[pulumi.Input[int]] = None,
                 vip_ddns: Optional[pulumi.Input[bool]] = None,
                 vip_interface: Optional[pulumi.Input[str]] = None,
                 vip_leader_election: Optional[pulumi.Input[bool]] = None,
                 vip_lease_duration: Optional[pulumi.Input[int]] = None,
                 vip_renew_deadline: Optional[pulumi.Input[int]] = None,
                 vip_retry_period: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Pseudo resource for generating the kube-vip manifest.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: TODO
        :param pulumi.Input[int] bgp_as: TODO
        :param pulumi.Input[bool] bgp_enable: TODO
        :param pulumi.Input[str] bgp_peer_address: TODO
        :param pulumi.Input[int] bgp_peer_as: TODO
        :param pulumi.Input[str] bgp_peer_pass: TODO
        :param pulumi.Input[str] bgp_peers: TODO
        :param pulumi.Input[str] bgp_router_id: TODO
        :param pulumi.Input[bool] cp_enable: TODO
        :param pulumi.Input[str] cp_namespace: TODO
        :param pulumi.Input[str] image: Override the kube-vip image.
        :param pulumi.Input[str] kubeconfig_path: Path to the kubeconfig on the remote host.
        :param pulumi.Input[str] name: Name of the static pod. Defaults to kube-vip.
        :param pulumi.Input[str] namespace: Namespace for the static pod. Defaults to kube-system.
        :param pulumi.Input[int] port: TODO
        :param pulumi.Input[bool] svc_enable: TODO
        :param pulumi.Input[str] version: Version of kube-vip to use.
        :param pulumi.Input[bool] vip_arp: TODO
        :param pulumi.Input[int] vip_cidr: TODO
        :param pulumi.Input[bool] vip_ddns: TODO
        :param pulumi.Input[str] vip_interface: TODO
        :param pulumi.Input[bool] vip_leader_election: TODO
        :param pulumi.Input[int] vip_lease_duration: TODO
        :param pulumi.Input[int] vip_renew_deadline: TODO
        :param pulumi.Input[int] vip_retry_period: TODO
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubeVipManifestArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Pseudo resource for generating the kube-vip manifest.

        :param str resource_name: The name of the resource.
        :param KubeVipManifestArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubeVipManifestArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 bgp_as: Optional[pulumi.Input[int]] = None,
                 bgp_enable: Optional[pulumi.Input[bool]] = None,
                 bgp_peer_address: Optional[pulumi.Input[str]] = None,
                 bgp_peer_as: Optional[pulumi.Input[int]] = None,
                 bgp_peer_pass: Optional[pulumi.Input[str]] = None,
                 bgp_peers: Optional[pulumi.Input[str]] = None,
                 bgp_router_id: Optional[pulumi.Input[str]] = None,
                 cp_enable: Optional[pulumi.Input[bool]] = None,
                 cp_namespace: Optional[pulumi.Input[str]] = None,
                 image: Optional[pulumi.Input[str]] = None,
                 kubeconfig_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 svc_enable: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 vip_arp: Optional[pulumi.Input[bool]] = None,
                 vip_cidr: Optional[pulumi.Input[int]] = None,
                 vip_ddns: Optional[pulumi.Input[bool]] = None,
                 vip_interface: Optional[pulumi.Input[str]] = None,
                 vip_leader_election: Optional[pulumi.Input[bool]] = None,
                 vip_lease_duration: Optional[pulumi.Input[int]] = None,
                 vip_renew_deadline: Optional[pulumi.Input[int]] = None,
                 vip_retry_period: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubeVipManifestArgs.__new__(KubeVipManifestArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["bgp_as"] = bgp_as
            __props__.__dict__["bgp_enable"] = bgp_enable
            __props__.__dict__["bgp_peer_address"] = bgp_peer_address
            __props__.__dict__["bgp_peer_as"] = bgp_peer_as
            __props__.__dict__["bgp_peer_pass"] = bgp_peer_pass
            __props__.__dict__["bgp_peers"] = bgp_peers
            __props__.__dict__["bgp_router_id"] = bgp_router_id
            __props__.__dict__["cp_enable"] = cp_enable
            __props__.__dict__["cp_namespace"] = cp_namespace
            __props__.__dict__["image"] = image
            if kubeconfig_path is None and not opts.urn:
                raise TypeError("Missing required property 'kubeconfig_path'")
            __props__.__dict__["kubeconfig_path"] = kubeconfig_path
            if name is None:
                name = 'kube-vip'
            __props__.__dict__["name"] = name
            if namespace is None:
                namespace = 'kube-system'
            __props__.__dict__["namespace"] = namespace
            if port is None:
                port = 6443
            __props__.__dict__["port"] = port
            __props__.__dict__["svc_enable"] = svc_enable
            __props__.__dict__["version"] = version
            __props__.__dict__["vip_arp"] = vip_arp
            if vip_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'vip_cidr'")
            __props__.__dict__["vip_cidr"] = vip_cidr
            __props__.__dict__["vip_ddns"] = vip_ddns
            __props__.__dict__["vip_interface"] = vip_interface
            __props__.__dict__["vip_leader_election"] = vip_leader_election
            __props__.__dict__["vip_lease_duration"] = vip_lease_duration
            __props__.__dict__["vip_renew_deadline"] = vip_renew_deadline
            __props__.__dict__["vip_retry_period"] = vip_retry_period
            __props__.__dict__["result"] = None
            __props__.__dict__["yaml"] = None
        super(KubeVipManifest, __self__).__init__(
            'kubernetes-the-hard-way:config:KubeVipManifest',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output['outputs.PodManifest']:
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def yaml(self) -> pulumi.Output[str]:
        """
        The yaml representation of the manifest
        """
        return pulumi.get(self, "yaml")

