# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *
from ._inputs import *
from .certificate import Certificate
from .root_ca import RootCa

__all__ = [
    'NewCertificateResult',
    'AwaitableNewCertificateResult',
    'new_certificate',
    'new_certificate_output',
]

@pulumi.output_type
class NewCertificateResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, Certificate):
            raise TypeError("Expected argument 'result' to be a Certificate")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> 'Certificate':
        return pulumi.get(self, "result")


class AwaitableNewCertificateResult(NewCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return NewCertificateResult(
            result=self.result)


def new_certificate(algorithm: Optional['Algorithm'] = None,
                    allowed_uses: Optional[Sequence['AllowedUsage']] = None,
                    ca: Optional['RootCa'] = None,
                    dns_names: Optional[Sequence[str]] = None,
                    early_renewal_hours: Optional[int] = None,
                    ecdsa_curve: Optional['EcdsaCurve'] = None,
                    ip_addresses: Optional[Sequence[str]] = None,
                    is_ca_certificate: Optional[bool] = None,
                    name: Optional[str] = None,
                    opts: Optional[pulumi.InputType['ResourceOptions']] = None,
                    rsa_bits: Optional[int] = None,
                    set_authority_key_id: Optional[bool] = None,
                    set_subject_key_id: Optional[bool] = None,
                    subject: Optional[pulumi.InputType['CertRequestSubject']] = None,
                    uris: Optional[Sequence[str]] = None,
                    validity_period_hours: Optional[int] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableNewCertificateResult:
    """
    Creates a Certificate configured for the current authority.


    :param 'Algorithm' algorithm: Name of the algorithm to use when generating the private key.
    :param 'RootCa' ca: The certificate authority to issue the certificate.
    :param Sequence[str] dns_names: List of DNS names for which a certificate is being requested.
    :param int early_renewal_hours: TODO
    :param 'EcdsaCurve' ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
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
    __args__['ca'] = ca
    __args__['dnsNames'] = dns_names
    __args__['earlyRenewalHours'] = early_renewal_hours
    __args__['ecdsaCurve'] = ecdsa_curve
    __args__['ipAddresses'] = ip_addresses
    __args__['isCaCertificate'] = is_ca_certificate
    __args__['name'] = name
    __args__['opts'] = opts
    __args__['rsaBits'] = rsa_bits
    __args__['setAuthorityKeyId'] = set_authority_key_id
    __args__['setSubjectKeyId'] = set_subject_key_id
    __args__['subject'] = subject
    __args__['uris'] = uris
    __args__['validityPeriodHours'] = validity_period_hours
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('kubernetes-the-hard-way:index:newCertificate', __args__, opts=opts, typ=NewCertificateResult).value

    return AwaitableNewCertificateResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(new_certificate)
def new_certificate_output(algorithm: Optional[pulumi.Input['Algorithm']] = None,
                           allowed_uses: Optional[pulumi.Input[Sequence['AllowedUsage']]] = None,
                           ca: Optional[pulumi.Input['RootCa']] = None,
                           dns_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           early_renewal_hours: Optional[pulumi.Input[Optional[int]]] = None,
                           ecdsa_curve: Optional[pulumi.Input[Optional['EcdsaCurve']]] = None,
                           ip_addresses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           is_ca_certificate: Optional[pulumi.Input[Optional[bool]]] = None,
                           name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.Input[Optional[pulumi.InputType['ResourceOptions']]]] = None,
                           rsa_bits: Optional[pulumi.Input[Optional[int]]] = None,
                           set_authority_key_id: Optional[pulumi.Input[Optional[bool]]] = None,
                           set_subject_key_id: Optional[pulumi.Input[Optional[bool]]] = None,
                           subject: Optional[pulumi.Input[Optional[pulumi.InputType['CertRequestSubject']]]] = None,
                           uris: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           validity_period_hours: Optional[pulumi.Input[int]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[NewCertificateResult]:
    """
    Creates a Certificate configured for the current authority.


    :param 'Algorithm' algorithm: Name of the algorithm to use when generating the private key.
    :param 'RootCa' ca: The certificate authority to issue the certificate.
    :param Sequence[str] dns_names: List of DNS names for which a certificate is being requested.
    :param int early_renewal_hours: TODO
    :param 'EcdsaCurve' ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
    :param Sequence[str] ip_addresses: List of IP addresses for which a certificate is being requested.
    :param int rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
    :param bool set_authority_key_id: Should the generated certificate include an authority key identifier.
    :param bool set_subject_key_id: Should the generated certificate include a subject key identifier.
    :param Sequence[str] uris: List of URIs for which a certificate is being requested.
    :param int validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
    """
    ...
