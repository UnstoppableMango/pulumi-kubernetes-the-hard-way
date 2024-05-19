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

__all__ = ['KubeletConfigurationArgs', 'KubeletConfiguration']

@pulumi.input_type
class KubeletConfigurationArgs:
    def __init__(__self__, *,
                 pod_cidr: pulumi.Input[str],
                 anonymous: Optional[pulumi.Input[bool]] = None,
                 authorization_mode: Optional[pulumi.Input[str]] = None,
                 cgroup_driver: Optional[pulumi.Input[str]] = None,
                 client_ca_file: Optional[pulumi.Input[str]] = None,
                 cluster_dns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_domain: Optional[pulumi.Input[str]] = None,
                 container_runtime_endpoint: Optional[pulumi.Input[str]] = None,
                 resolv_conf: Optional[pulumi.Input[str]] = None,
                 runtime_request_timeout: Optional[pulumi.Input[str]] = None,
                 tls_cert_file: Optional[pulumi.Input[str]] = None,
                 tls_private_key_file: Optional[pulumi.Input[str]] = None,
                 webhook: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a KubeletConfiguration resource.
        """
        pulumi.set(__self__, "pod_cidr", pod_cidr)
        if anonymous is not None:
            pulumi.set(__self__, "anonymous", anonymous)
        if authorization_mode is not None:
            pulumi.set(__self__, "authorization_mode", authorization_mode)
        if cgroup_driver is not None:
            pulumi.set(__self__, "cgroup_driver", cgroup_driver)
        if client_ca_file is not None:
            pulumi.set(__self__, "client_ca_file", client_ca_file)
        if cluster_dns is not None:
            pulumi.set(__self__, "cluster_dns", cluster_dns)
        if cluster_domain is not None:
            pulumi.set(__self__, "cluster_domain", cluster_domain)
        if container_runtime_endpoint is not None:
            pulumi.set(__self__, "container_runtime_endpoint", container_runtime_endpoint)
        if resolv_conf is not None:
            pulumi.set(__self__, "resolv_conf", resolv_conf)
        if runtime_request_timeout is not None:
            pulumi.set(__self__, "runtime_request_timeout", runtime_request_timeout)
        if tls_cert_file is not None:
            pulumi.set(__self__, "tls_cert_file", tls_cert_file)
        if tls_private_key_file is not None:
            pulumi.set(__self__, "tls_private_key_file", tls_private_key_file)
        if webhook is not None:
            pulumi.set(__self__, "webhook", webhook)

    @property
    @pulumi.getter(name="podCIDR")
    def pod_cidr(self) -> pulumi.Input[str]:
        return pulumi.get(self, "pod_cidr")

    @pod_cidr.setter
    def pod_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "pod_cidr", value)

    @property
    @pulumi.getter
    def anonymous(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "anonymous")

    @anonymous.setter
    def anonymous(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "anonymous", value)

    @property
    @pulumi.getter(name="authorizationMode")
    def authorization_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "authorization_mode")

    @authorization_mode.setter
    def authorization_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_mode", value)

    @property
    @pulumi.getter(name="cgroupDriver")
    def cgroup_driver(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cgroup_driver")

    @cgroup_driver.setter
    def cgroup_driver(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cgroup_driver", value)

    @property
    @pulumi.getter(name="clientCAFile")
    def client_ca_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_ca_file")

    @client_ca_file.setter
    def client_ca_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_ca_file", value)

    @property
    @pulumi.getter(name="clusterDNS")
    def cluster_dns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "cluster_dns")

    @cluster_dns.setter
    def cluster_dns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cluster_dns", value)

    @property
    @pulumi.getter(name="clusterDomain")
    def cluster_domain(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cluster_domain")

    @cluster_domain.setter
    def cluster_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_domain", value)

    @property
    @pulumi.getter(name="containerRuntimeEndpoint")
    def container_runtime_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container_runtime_endpoint")

    @container_runtime_endpoint.setter
    def container_runtime_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_runtime_endpoint", value)

    @property
    @pulumi.getter(name="resolvConf")
    def resolv_conf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resolv_conf")

    @resolv_conf.setter
    def resolv_conf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resolv_conf", value)

    @property
    @pulumi.getter(name="runtimeRequestTimeout")
    def runtime_request_timeout(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "runtime_request_timeout")

    @runtime_request_timeout.setter
    def runtime_request_timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "runtime_request_timeout", value)

    @property
    @pulumi.getter(name="tlsCertFile")
    def tls_cert_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_cert_file")

    @tls_cert_file.setter
    def tls_cert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_cert_file", value)

    @property
    @pulumi.getter(name="tlsPrivateKeyFile")
    def tls_private_key_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_private_key_file")

    @tls_private_key_file.setter
    def tls_private_key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_private_key_file", value)

    @property
    @pulumi.getter
    def webhook(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "webhook")

    @webhook.setter
    def webhook(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "webhook", value)


class KubeletConfiguration(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anonymous: Optional[pulumi.Input[bool]] = None,
                 authorization_mode: Optional[pulumi.Input[str]] = None,
                 cgroup_driver: Optional[pulumi.Input[str]] = None,
                 client_ca_file: Optional[pulumi.Input[str]] = None,
                 cluster_dns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_domain: Optional[pulumi.Input[str]] = None,
                 container_runtime_endpoint: Optional[pulumi.Input[str]] = None,
                 pod_cidr: Optional[pulumi.Input[str]] = None,
                 resolv_conf: Optional[pulumi.Input[str]] = None,
                 runtime_request_timeout: Optional[pulumi.Input[str]] = None,
                 tls_cert_file: Optional[pulumi.Input[str]] = None,
                 tls_private_key_file: Optional[pulumi.Input[str]] = None,
                 webhook: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a KubeletConfiguration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubeletConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a KubeletConfiguration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param KubeletConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubeletConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anonymous: Optional[pulumi.Input[bool]] = None,
                 authorization_mode: Optional[pulumi.Input[str]] = None,
                 cgroup_driver: Optional[pulumi.Input[str]] = None,
                 client_ca_file: Optional[pulumi.Input[str]] = None,
                 cluster_dns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_domain: Optional[pulumi.Input[str]] = None,
                 container_runtime_endpoint: Optional[pulumi.Input[str]] = None,
                 pod_cidr: Optional[pulumi.Input[str]] = None,
                 resolv_conf: Optional[pulumi.Input[str]] = None,
                 runtime_request_timeout: Optional[pulumi.Input[str]] = None,
                 tls_cert_file: Optional[pulumi.Input[str]] = None,
                 tls_private_key_file: Optional[pulumi.Input[str]] = None,
                 webhook: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubeletConfigurationArgs.__new__(KubeletConfigurationArgs)

            __props__.__dict__["anonymous"] = anonymous
            __props__.__dict__["authorization_mode"] = authorization_mode
            __props__.__dict__["cgroup_driver"] = cgroup_driver
            __props__.__dict__["client_ca_file"] = client_ca_file
            __props__.__dict__["cluster_dns"] = cluster_dns
            __props__.__dict__["cluster_domain"] = cluster_domain
            __props__.__dict__["container_runtime_endpoint"] = container_runtime_endpoint
            if pod_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'pod_cidr'")
            __props__.__dict__["pod_cidr"] = pod_cidr
            __props__.__dict__["resolv_conf"] = resolv_conf
            __props__.__dict__["runtime_request_timeout"] = runtime_request_timeout
            __props__.__dict__["tls_cert_file"] = tls_cert_file
            __props__.__dict__["tls_private_key_file"] = tls_private_key_file
            __props__.__dict__["webhook"] = webhook
            __props__.__dict__["result"] = None
            __props__.__dict__["yaml"] = None
        super(KubeletConfiguration, __self__).__init__(
            'kubernetes-the-hard-way:config:KubeletConfiguration',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output['outputs.KubeletConfiguration']:
        return pulumi.get(self, "result")

    @property
    @pulumi.getter
    def yaml(self) -> pulumi.Output[str]:
        """
        The yaml representation of the manifest
        """
        return pulumi.get(self, "yaml")
