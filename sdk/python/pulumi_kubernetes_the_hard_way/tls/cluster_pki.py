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
from .. import config as _config
from ._enums import *
from ._inputs import *
from .certificate import Certificate
from .root_ca import RootCa

__all__ = ['ClusterPkiArgs', 'ClusterPki']

@pulumi.input_type
class ClusterPkiArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 nodes: pulumi.Input[Mapping[str, pulumi.Input['ClusterPkiNodeArgs']]],
                 public_ip: pulumi.Input[str],
                 algorithm: Optional[pulumi.Input['Algorithm']] = None,
                 ecdsa_curve: Optional[pulumi.Input['EcdsaCurve']] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ClusterPki resource.
        :param pulumi.Input[str] cluster_name: A name to use for the cluster
        :param pulumi.Input[Mapping[str, pulumi.Input['ClusterPkiNodeArgs']]] nodes: Map of node name to node configuration
        :param pulumi.Input[str] public_ip: Publicly accessible IP address.
        :param pulumi.Input['Algorithm'] algorithm: Name of the algorithm to use when generating the private key.
        :param pulumi.Input['EcdsaCurve'] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        :param pulumi.Input[int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "nodes", nodes)
        pulumi.set(__self__, "public_ip", public_ip)
        if algorithm is None:
            algorithm = 'RSA'
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if ecdsa_curve is not None:
            pulumi.set(__self__, "ecdsa_curve", ecdsa_curve)
        if rsa_bits is None:
            rsa_bits = 2048
        if rsa_bits is not None:
            pulumi.set(__self__, "rsa_bits", rsa_bits)
        if validity_period_hours is None:
            validity_period_hours = 8076
        if validity_period_hours is not None:
            pulumi.set(__self__, "validity_period_hours", validity_period_hours)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        A name to use for the cluster
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Input[Mapping[str, pulumi.Input['ClusterPkiNodeArgs']]]:
        """
        Map of node name to node configuration
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: pulumi.Input[Mapping[str, pulumi.Input['ClusterPkiNodeArgs']]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> pulumi.Input[str]:
        """
        Publicly accessible IP address.
        """
        return pulumi.get(self, "public_ip")

    @public_ip.setter
    def public_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_ip", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input['Algorithm']]:
        """
        Name of the algorithm to use when generating the private key.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input['Algorithm']]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="ecdsaCurve")
    def ecdsa_curve(self) -> Optional[pulumi.Input['EcdsaCurve']]:
        """
        When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        """
        return pulumi.get(self, "ecdsa_curve")

    @ecdsa_curve.setter
    def ecdsa_curve(self, value: Optional[pulumi.Input['EcdsaCurve']]):
        pulumi.set(self, "ecdsa_curve", value)

    @property
    @pulumi.getter(name="rsaBits")
    def rsa_bits(self) -> Optional[pulumi.Input[int]]:
        """
        When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        """
        return pulumi.get(self, "rsa_bits")

    @rsa_bits.setter
    def rsa_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rsa_bits", value)

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> Optional[pulumi.Input[int]]:
        """
        Number of hours, after initial issuing, that the certificate will remain valid.
        """
        return pulumi.get(self, "validity_period_hours")

    @validity_period_hours.setter
    def validity_period_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "validity_period_hours", value)


class ClusterPki(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input['Algorithm']] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 ecdsa_curve: Optional[pulumi.Input['EcdsaCurve']] = None,
                 nodes: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ClusterPkiNodeArgs']]]]] = None,
                 public_ip: Optional[pulumi.Input[str]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The private key infrastructure for a cluster

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['Algorithm'] algorithm: Name of the algorithm to use when generating the private key.
        :param pulumi.Input[str] cluster_name: A name to use for the cluster
        :param pulumi.Input['EcdsaCurve'] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        :param pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ClusterPkiNodeArgs']]]] nodes: Map of node name to node configuration
        :param pulumi.Input[str] public_ip: Publicly accessible IP address.
        :param pulumi.Input[int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterPkiArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The private key infrastructure for a cluster

        :param str resource_name: The name of the resource.
        :param ClusterPkiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterPkiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input['Algorithm']] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 ecdsa_curve: Optional[pulumi.Input['EcdsaCurve']] = None,
                 nodes: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ClusterPkiNodeArgs']]]]] = None,
                 public_ip: Optional[pulumi.Input[str]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterPkiArgs.__new__(ClusterPkiArgs)

            if algorithm is None:
                algorithm = 'RSA'
            __props__.__dict__["algorithm"] = algorithm
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["ecdsa_curve"] = ecdsa_curve
            if nodes is None and not opts.urn:
                raise TypeError("Missing required property 'nodes'")
            __props__.__dict__["nodes"] = nodes
            if public_ip is None and not opts.urn:
                raise TypeError("Missing required property 'public_ip'")
            __props__.__dict__["public_ip"] = public_ip
            if rsa_bits is None:
                rsa_bits = 2048
            __props__.__dict__["rsa_bits"] = rsa_bits
            if validity_period_hours is None:
                validity_period_hours = 8076
            __props__.__dict__["validity_period_hours"] = validity_period_hours
            __props__.__dict__["admin"] = None
            __props__.__dict__["ca"] = None
            __props__.__dict__["controller_manager"] = None
            __props__.__dict__["kube_proxy"] = None
            __props__.__dict__["kube_scheduler"] = None
            __props__.__dict__["kubelet"] = None
            __props__.__dict__["kubernetes"] = None
            __props__.__dict__["service_accounts"] = None
        super(ClusterPki, __self__).__init__(
            'kubernetes-the-hard-way:tls:ClusterPki',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def admin(self) -> pulumi.Output['Certificate']:
        """
        The admin certificate.
        """
        return pulumi.get(self, "admin")

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[Optional['Algorithm']]:
        """
        Name of the algorithm to use when generating the private key.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def ca(self) -> pulumi.Output['RootCa']:
        """
        The cluster certificate authority.
        """
        return pulumi.get(self, "ca")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        A name to use for the cluster
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="controllerManager")
    def controller_manager(self) -> pulumi.Output['Certificate']:
        """
        The controller manager certificate.
        """
        return pulumi.get(self, "controller_manager")

    @property
    @pulumi.getter(name="ecdsaCurve")
    def ecdsa_curve(self) -> pulumi.Output[Optional['EcdsaCurve']]:
        """
        When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        """
        return pulumi.get(self, "ecdsa_curve")

    @property
    @pulumi.getter(name="kubeProxy")
    def kube_proxy(self) -> pulumi.Output['Certificate']:
        """
        The kube proxy certificate.
        """
        return pulumi.get(self, "kube_proxy")

    @property
    @pulumi.getter(name="kubeScheduler")
    def kube_scheduler(self) -> pulumi.Output['Certificate']:
        """
        The kube scheduler certificate.
        """
        return pulumi.get(self, "kube_scheduler")

    @property
    @pulumi.getter
    def kubelet(self) -> pulumi.Output[Mapping[str, 'Certificate']]:
        """
        Map of node name to kubelet certificate.
        """
        return pulumi.get(self, "kubelet")

    @property
    @pulumi.getter
    def kubernetes(self) -> pulumi.Output['Certificate']:
        """
        The kubernetes certificate.
        """
        return pulumi.get(self, "kubernetes")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Mapping[str, 'outputs.ClusterPkiNode']]:
        """
        Map of node name to node configuration
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> pulumi.Output[str]:
        """
        Publicly accessible IP address.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="rsaBits")
    def rsa_bits(self) -> pulumi.Output[Optional[int]]:
        """
        When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        """
        return pulumi.get(self, "rsa_bits")

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> pulumi.Output['Certificate']:
        """
        The service accounts certificate
        """
        return pulumi.get(self, "service_accounts")

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> pulumi.Output[int]:
        """
        Number of hours, after initial issuing, that the certificate will remain valid.
        """
        return pulumi.get(self, "validity_period_hours")

    @pulumi.output_type
    class GetKubeconfigResult:
        def __init__(__self__, result=None):
            if result and not isinstance(result, dict):
                raise TypeError("Expected argument 'result' to be a dict")
            pulumi.set(__self__, "result", result)

        @property
        @pulumi.getter
        def result(self) -> '_config.outputs.Kubeconfig':
            return pulumi.get(self, "result")

    def get_kubeconfig(__self__, *,
                       options: Union[pulumi.Input['_config.KubeconfigAdminOptionsArgs'], pulumi.Input['_config.KubeconfigKubeControllerManagerOptionsArgs'], pulumi.Input['_config.KubeconfigKubeProxyOptionsArgs'], pulumi.Input['_config.KubeconfigKubeSchedulerOptionsArgs'], pulumi.Input['_config.KubeconfigWorkerOptionsArgs']]) -> pulumi.Output['dict']:
        """
        Get a kubeconfig configured from this PKI.


        :param Union[pulumi.Input['_config.KubeconfigAdminOptionsArgs'], pulumi.Input['_config.KubeconfigKubeControllerManagerOptionsArgs'], pulumi.Input['_config.KubeconfigKubeProxyOptionsArgs'], pulumi.Input['_config.KubeconfigKubeSchedulerOptionsArgs'], pulumi.Input['_config.KubeconfigWorkerOptionsArgs']] options: Options for creating the kubeconfig.
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['options'] = options
        __result__ = pulumi.runtime.call('kubernetes-the-hard-way:tls:ClusterPki/getKubeconfig', __args__, res=__self__, typ=ClusterPki.GetKubeconfigResult)
        return __result__.result

