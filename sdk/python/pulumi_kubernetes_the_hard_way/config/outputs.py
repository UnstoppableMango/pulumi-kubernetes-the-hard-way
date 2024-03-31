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
from ._enums import *

__all__ = [
    'Cluster',
    'Context',
    'Kubeconfig',
    'User',
]

@pulumi.output_type
class Cluster(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateAuthorityData":
            suggest = "certificate_authority_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Cluster. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Cluster.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Cluster.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_authority_data: str,
                 server: str):
        pulumi.set(__self__, "certificate_authority_data", certificate_authority_data)
        pulumi.set(__self__, "server", server)

    @property
    @pulumi.getter(name="certificateAuthorityData")
    def certificate_authority_data(self) -> str:
        return pulumi.get(self, "certificate_authority_data")

    @property
    @pulumi.getter
    def server(self) -> str:
        return pulumi.get(self, "server")


@pulumi.output_type
class Context(dict):
    def __init__(__self__, *,
                 cluster: str,
                 user: str):
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def cluster(self) -> str:
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter
    def user(self) -> str:
        return pulumi.get(self, "user")


@pulumi.output_type
class Kubeconfig(dict):
    def __init__(__self__, *,
                 clusters: Sequence['outputs.Cluster'],
                 contexts: Sequence['outputs.Context'],
                 users: Sequence['outputs.User']):
        pulumi.set(__self__, "clusters", clusters)
        pulumi.set(__self__, "contexts", contexts)
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.Cluster']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter
    def contexts(self) -> Sequence['outputs.Context']:
        return pulumi.get(self, "contexts")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.User']:
        return pulumi.get(self, "users")


@pulumi.output_type
class User(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientCertificateData":
            suggest = "client_certificate_data"
        elif key == "clientKeyData":
            suggest = "client_key_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in User. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        User.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        User.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_certificate_data: str,
                 client_key_data: str):
        pulumi.set(__self__, "client_certificate_data", client_certificate_data)
        pulumi.set(__self__, "client_key_data", client_key_data)

    @property
    @pulumi.getter(name="clientCertificateData")
    def client_certificate_data(self) -> str:
        return pulumi.get(self, "client_certificate_data")

    @property
    @pulumi.getter(name="clientKeyData")
    def client_key_data(self) -> str:
        return pulumi.get(self, "client_key_data")


