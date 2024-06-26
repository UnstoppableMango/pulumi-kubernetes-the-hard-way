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
from .. import tls as _tls
from ._enums import *
from ._inputs import *
from .etcd_configuration import EtcdConfiguration
from .etcd_install import EtcdInstall
from .etcd_service import EtcdService
from .start_etcd import StartEtcd
import pulumi_command

__all__ = ['EtcdClusterArgs', 'EtcdCluster']

@pulumi.input_type
class EtcdClusterArgs:
    def __init__(__self__, *,
                 bundle: pulumi.Input['_tls.BundleArgs'],
                 nodes: Mapping[str, pulumi.Input['EtcdNodeArgs']],
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 binary_directory: Optional[pulumi.Input[str]] = None,
                 configuration_directory: Optional[pulumi.Input[str]] = None,
                 data_directory: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EtcdCluster resource.
        :param pulumi.Input['_tls.BundleArgs'] bundle: The TLS bundle.
        :param Mapping[str, pulumi.Input['EtcdNodeArgs']] nodes: Etcd node configuration. The key should be a name used to identify the node.
        :param pulumi.Input['Architecture'] architecture: TODO
        :param pulumi.Input[str] binary_directory: TODO
        :param pulumi.Input[str] configuration_directory: The directory to use for etcd configuration.
        :param pulumi.Input[str] data_directory: The directory to use for etcd data.
        :param pulumi.Input[str] version: The version to install.
        """
        pulumi.set(__self__, "bundle", bundle)
        pulumi.set(__self__, "nodes", nodes)
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if binary_directory is not None:
            pulumi.set(__self__, "binary_directory", binary_directory)
        if configuration_directory is not None:
            pulumi.set(__self__, "configuration_directory", configuration_directory)
        if data_directory is not None:
            pulumi.set(__self__, "data_directory", data_directory)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def bundle(self) -> pulumi.Input['_tls.BundleArgs']:
        """
        The TLS bundle.
        """
        return pulumi.get(self, "bundle")

    @bundle.setter
    def bundle(self, value: pulumi.Input['_tls.BundleArgs']):
        pulumi.set(self, "bundle", value)

    @property
    @pulumi.getter
    def nodes(self) -> Mapping[str, pulumi.Input['EtcdNodeArgs']]:
        """
        Etcd node configuration. The key should be a name used to identify the node.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Mapping[str, pulumi.Input['EtcdNodeArgs']]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input['Architecture']]:
        """
        TODO
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input['Architecture']]):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter(name="binaryDirectory")
    def binary_directory(self) -> Optional[pulumi.Input[str]]:
        """
        TODO
        """
        return pulumi.get(self, "binary_directory")

    @binary_directory.setter
    def binary_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "binary_directory", value)

    @property
    @pulumi.getter(name="configurationDirectory")
    def configuration_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to use for etcd configuration.
        """
        return pulumi.get(self, "configuration_directory")

    @configuration_directory.setter
    def configuration_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration_directory", value)

    @property
    @pulumi.getter(name="dataDirectory")
    def data_directory(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to use for etcd data.
        """
        return pulumi.get(self, "data_directory")

    @data_directory.setter
    def data_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_directory", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version to install.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class EtcdCluster(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 binary_directory: Optional[pulumi.Input[str]] = None,
                 bundle: Optional[pulumi.Input[pulumi.InputType['_tls.BundleArgs']]] = None,
                 configuration_directory: Optional[pulumi.Input[str]] = None,
                 data_directory: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[Mapping[str, pulumi.Input[pulumi.InputType['EtcdNodeArgs']]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an etcd cluster from one or more remote systems.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['Architecture'] architecture: TODO
        :param pulumi.Input[str] binary_directory: TODO
        :param pulumi.Input[pulumi.InputType['_tls.BundleArgs']] bundle: The TLS bundle.
        :param pulumi.Input[str] configuration_directory: The directory to use for etcd configuration.
        :param pulumi.Input[str] data_directory: The directory to use for etcd data.
        :param Mapping[str, pulumi.Input[pulumi.InputType['EtcdNodeArgs']]] nodes: Etcd node configuration. The key should be a name used to identify the node.
        :param pulumi.Input[str] version: The version to install.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EtcdClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an etcd cluster from one or more remote systems.

        :param str resource_name: The name of the resource.
        :param EtcdClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EtcdClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['Architecture']] = None,
                 binary_directory: Optional[pulumi.Input[str]] = None,
                 bundle: Optional[pulumi.Input[pulumi.InputType['_tls.BundleArgs']]] = None,
                 configuration_directory: Optional[pulumi.Input[str]] = None,
                 data_directory: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[Mapping[str, pulumi.Input[pulumi.InputType['EtcdNodeArgs']]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EtcdClusterArgs.__new__(EtcdClusterArgs)

            __props__.__dict__["architecture"] = architecture
            __props__.__dict__["binary_directory"] = binary_directory
            if bundle is None and not opts.urn:
                raise TypeError("Missing required property 'bundle'")
            __props__.__dict__["bundle"] = bundle
            __props__.__dict__["configuration_directory"] = configuration_directory
            __props__.__dict__["data_directory"] = data_directory
            if nodes is None and not opts.urn:
                raise TypeError("Missing required property 'nodes'")
            __props__.__dict__["nodes"] = nodes
            __props__.__dict__["version"] = version
            __props__.__dict__["configuration"] = None
            __props__.__dict__["install"] = None
            __props__.__dict__["service"] = None
            __props__.__dict__["start"] = None
        super(EtcdCluster, __self__).__init__(
            'kubernetes-the-hard-way:remote:EtcdCluster',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Output[Optional['Architecture']]:
        """
        TODO
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="binaryDirectory")
    def binary_directory(self) -> pulumi.Output[Optional[str]]:
        """
        TODO
        """
        return pulumi.get(self, "binary_directory")

    @property
    @pulumi.getter
    def bundle(self) -> pulumi.Output['_tls.outputs.Bundle']:
        """
        The TLS bundle.
        """
        return pulumi.get(self, "bundle")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Mapping[str, 'EtcdConfiguration']]:
        """
        Map of node name to etcd configuration.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="configurationDirectory")
    def configuration_directory(self) -> pulumi.Output[Optional[str]]:
        """
        The directory to use for etcd configuration.
        """
        return pulumi.get(self, "configuration_directory")

    @property
    @pulumi.getter(name="dataDirectory")
    def data_directory(self) -> pulumi.Output[Optional[str]]:
        """
        The directory to use for etcd data.
        """
        return pulumi.get(self, "data_directory")

    @property
    @pulumi.getter
    def install(self) -> pulumi.Output[Mapping[str, 'EtcdInstall']]:
        """
        Map of node name to etcd install.
        """
        return pulumi.get(self, "install")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Mapping[str, 'outputs.EtcdNode']]:
        """
        Etcd node configuration. The key should be a name used to identify the node.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[Mapping[str, 'EtcdService']]:
        """
        Map of node name to etcd systemd service.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def start(self) -> pulumi.Output[Mapping[str, 'StartEtcd']]:
        """
        Map of node name to etcd start commands.
        """
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        The version to install.
        """
        return pulumi.get(self, "version")

