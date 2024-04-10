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

__all__ = [
    'GetKubeVipManifestResult',
    'AwaitableGetKubeVipManifestResult',
    'get_kube_vip_manifest',
    'get_kube_vip_manifest_output',
]

@pulumi.output_type
class GetKubeVipManifestResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, dict):
            raise TypeError("Expected argument 'result' to be a dict")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> 'outputs.PodManifest':
        return pulumi.get(self, "result")


class AwaitableGetKubeVipManifestResult(GetKubeVipManifestResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubeVipManifestResult(
            result=self.result)


def get_kube_vip_manifest(address: Optional[str] = None,
                          cp_enable: Optional[bool] = None,
                          cp_namespace: Optional[str] = None,
                          image: Optional[str] = None,
                          kubeconfig_path: Optional[str] = None,
                          port: Optional[int] = None,
                          svc_enable: Optional[bool] = None,
                          version: Optional[str] = None,
                          vip_arp: Optional[bool] = None,
                          vip_cidr: Optional[str] = None,
                          vip_ddns: Optional[bool] = None,
                          vip_interface: Optional[str] = None,
                          vip_leader_election: Optional[bool] = None,
                          vip_lease_duration: Optional[int] = None,
                          vip_renew_deadline: Optional[int] = None,
                          vip_retry_period: Optional[int] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubeVipManifestResult:
    """
    Gets the static pod manifests for KubeVip.


    :param str address: TODO
    :param bool cp_enable: TODO
    :param str cp_namespace: TODO
    :param str image: Override the kube-vip image.
    :param str kubeconfig_path: Path to the kubeconfig on the remote host.
    :param int port: TODO
    :param bool svc_enable: TODO
    :param str version: Version of kube-vip to use.
    :param bool vip_arp: TODO
    :param str vip_cidr: TODO
    :param bool vip_ddns: TODO
    :param str vip_interface: TODO
    :param bool vip_leader_election: TODO
    :param int vip_lease_duration: TODO
    :param int vip_renew_deadline: TODO
    :param int vip_retry_period: TODO
    """
    __args__ = dict()
    __args__['address'] = address
    __args__['cpEnable'] = cp_enable
    __args__['cpNamespace'] = cp_namespace
    __args__['image'] = image
    __args__['kubeconfigPath'] = kubeconfig_path
    __args__['port'] = port
    __args__['svcEnable'] = svc_enable
    __args__['version'] = version
    __args__['vipArp'] = vip_arp
    __args__['vipCidr'] = vip_cidr
    __args__['vipDdns'] = vip_ddns
    __args__['vipInterface'] = vip_interface
    __args__['vipLeaderElection'] = vip_leader_election
    __args__['vipLeaseDuration'] = vip_lease_duration
    __args__['vipRenewDeadline'] = vip_renew_deadline
    __args__['vipRetryPeriod'] = vip_retry_period
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kubernetes-the-hard-way:config:getKubeVipManifest', __args__, opts=opts, typ=GetKubeVipManifestResult).value

    return AwaitableGetKubeVipManifestResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(get_kube_vip_manifest)
def get_kube_vip_manifest_output(address: Optional[pulumi.Input[str]] = None,
                                 cp_enable: Optional[pulumi.Input[Optional[bool]]] = None,
                                 cp_namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                 image: Optional[pulumi.Input[Optional[str]]] = None,
                                 kubeconfig_path: Optional[pulumi.Input[str]] = None,
                                 port: Optional[pulumi.Input[Optional[int]]] = None,
                                 svc_enable: Optional[pulumi.Input[Optional[bool]]] = None,
                                 version: Optional[pulumi.Input[Optional[str]]] = None,
                                 vip_arp: Optional[pulumi.Input[Optional[bool]]] = None,
                                 vip_cidr: Optional[pulumi.Input[str]] = None,
                                 vip_ddns: Optional[pulumi.Input[Optional[bool]]] = None,
                                 vip_interface: Optional[pulumi.Input[Optional[str]]] = None,
                                 vip_leader_election: Optional[pulumi.Input[Optional[bool]]] = None,
                                 vip_lease_duration: Optional[pulumi.Input[Optional[int]]] = None,
                                 vip_renew_deadline: Optional[pulumi.Input[Optional[int]]] = None,
                                 vip_retry_period: Optional[pulumi.Input[Optional[int]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubeVipManifestResult]:
    """
    Gets the static pod manifests for KubeVip.


    :param str address: TODO
    :param bool cp_enable: TODO
    :param str cp_namespace: TODO
    :param str image: Override the kube-vip image.
    :param str kubeconfig_path: Path to the kubeconfig on the remote host.
    :param int port: TODO
    :param bool svc_enable: TODO
    :param str version: Version of kube-vip to use.
    :param bool vip_arp: TODO
    :param str vip_cidr: TODO
    :param bool vip_ddns: TODO
    :param str vip_interface: TODO
    :param bool vip_leader_election: TODO
    :param int vip_lease_duration: TODO
    :param int vip_renew_deadline: TODO
    :param int vip_retry_period: TODO
    """
    ...
