# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'Algorithm',
    'AllowedUsage',
    'EcdsaCurve',
    'NodeRole',
]


class Algorithm(str, Enum):
    """
    Private key algorithm.
    """
    RSA = "RSA"
    ECDSA = "ECDSA"
    ED25519 = "ED25519"


class AllowedUsage(str, Enum):
    """
    Certificate allowed usage
    """
    CERT_SIGNING = "cert_signing"
    CLIENT_AUTH = "client_auth"
    CRL_SIGNING = "crl_signing"
    DIGITAL_SIGNATURE = "digital_signature"
    KEY_ENCIPHERMENT = "key_encipherment"
    SERVER_AUTH = "server_auth"


class EcdsaCurve(str, Enum):
    """
    ECDSA algorithm curve
    """
    P224 = "P224"
    P256 = "P256"
    P384 = "P384"
    P521 = "P521"


class NodeRole(str, Enum):
    """
    The role a node will play in the final cluster
    """
    CONTROLPLANE = "controlplane"
    WORKER = "worker"
