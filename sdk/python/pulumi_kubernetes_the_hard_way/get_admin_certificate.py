# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from . import _utilities
import pulumi_tls

__all__ = [
    'get_admin_certificate',
]

def get_admin_certificate(algorithm: Optional[str] = None,
                          allowed_uses: Optional[Sequence[str]] = None,
                          dns_names: Optional[Sequence[str]] = None,
                          early_renewal_hours: Optional[int] = None,
                          ecdsa_curve: Optional[str] = None,
                          ip_addresses: Optional[Sequence[str]] = None,
                          is_ca_certificate: Optional[bool] = None,
                          rsa_bits: Optional[int] = None,
                          set_authority_key_id: Optional[bool] = None,
                          set_subject_key_id: Optional[bool] = None,
                          subject: Optional[pulumi.InputType['pulumi_tls.CertRequestSubjectArgs']] = None,
                          uris: Optional[Sequence[str]] = None,
                          validity_period_hours: Optional[int] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> Awaitable[None]:
    """
    Creates a Certificate configured for the cluster admin.


    :param str algorithm: Name of the algorithm to use when generating the private key.
    :param Sequence[str] dns_names: List of DNS names for which a certificate is being requested.
    :param int early_renewal_hours: TODO
    :param str ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
    :param Sequence[str] ip_addresses: List of IP addresses for which a certificate is being requested.
    :param int rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
    :param bool set_authority_key_id: Should the generated certificate include an authority key identifier.
    :param bool set_subject_key_id: Should the generated certificate include a subject key identifier.
    :param Sequence[str] uris: List of URIs for which a certificate is being requested.
    :param int validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
    """
    __args__ = dict()
    __args__['algorithm'] = algorithm
    __args__['allowedUses'] = allowed_uses
    __args__['dnsNames'] = dns_names
    __args__['earlyRenewalHours'] = early_renewal_hours
    __args__['ecdsaCurve'] = ecdsa_curve
    __args__['ipAddresses'] = ip_addresses
    __args__['isCaCertificate'] = is_ca_certificate
    __args__['rsaBits'] = rsa_bits
    __args__['setAuthorityKeyId'] = set_authority_key_id
    __args__['setSubjectKeyId'] = set_subject_key_id
    __args__['subject'] = subject
    __args__['uris'] = uris
    __args__['validityPeriodHours'] = validity_period_hours
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kubernetes-the-hard-way:index:getAdminCertificate', __args__, opts=opts).value

