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
from .remote_file import RemoteFile
import pulumi_tls

__all__ = ['RootCaArgs', 'RootCa']

@pulumi.input_type
class RootCaArgs:
    def __init__(__self__, *,
                 validity_period_hours: pulumi.Input[int],
                 algorithm: Optional[pulumi.Input['Algorithm']] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ecdsa_curve: Optional[pulumi.Input['EcdsaCurve']] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 set_authority_key_id: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subject: Optional[pulumi.Input['pulumi_tls.SelfSignedCertSubjectArgs']] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RootCa resource.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
        :param pulumi.Input['Algorithm'] algorithm: Name of the algorithm to use when generating the private key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: TODO
        :param pulumi.Input['EcdsaCurve'] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        :param pulumi.Input[bool] set_authority_key_id: Should the generated certificate include an authority key identifier.
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a subject key identifier.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        """
        pulumi.set(__self__, "validity_period_hours", validity_period_hours)
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if dns_names is not None:
            pulumi.set(__self__, "dns_names", dns_names)
        if early_renewal_hours is not None:
            pulumi.set(__self__, "early_renewal_hours", early_renewal_hours)
        if ecdsa_curve is not None:
            pulumi.set(__self__, "ecdsa_curve", ecdsa_curve)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if rsa_bits is not None:
            pulumi.set(__self__, "rsa_bits", rsa_bits)
        if set_authority_key_id is not None:
            pulumi.set(__self__, "set_authority_key_id", set_authority_key_id)
        if set_subject_key_id is not None:
            pulumi.set(__self__, "set_subject_key_id", set_subject_key_id)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if uris is not None:
            pulumi.set(__self__, "uris", uris)

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> pulumi.Input[int]:
        """
        Number of hours, after initial issuing, that the certificate will remain valid.
        """
        return pulumi.get(self, "validity_period_hours")

    @validity_period_hours.setter
    def validity_period_hours(self, value: pulumi.Input[int]):
        pulumi.set(self, "validity_period_hours", value)

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
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DNS names for which a certificate is being requested.
        """
        return pulumi.get(self, "dns_names")

    @dns_names.setter
    def dns_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_names", value)

    @property
    @pulumi.getter(name="earlyRenewalHours")
    def early_renewal_hours(self) -> Optional[pulumi.Input[int]]:
        """
        TODO
        """
        return pulumi.get(self, "early_renewal_hours")

    @early_renewal_hours.setter
    def early_renewal_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "early_renewal_hours", value)

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
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses for which a certificate is being requested.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

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
    @pulumi.getter(name="setAuthorityKeyId")
    def set_authority_key_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the generated certificate include an authority key identifier.
        """
        return pulumi.get(self, "set_authority_key_id")

    @set_authority_key_id.setter
    def set_authority_key_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_authority_key_id", value)

    @property
    @pulumi.getter(name="setSubjectKeyId")
    def set_subject_key_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the generated certificate include a subject key identifier.
        """
        return pulumi.get(self, "set_subject_key_id")

    @set_subject_key_id.setter
    def set_subject_key_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_subject_key_id", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input['pulumi_tls.SelfSignedCertSubjectArgs']]:
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input['pulumi_tls.SelfSignedCertSubjectArgs']]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter
    def uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URIs for which a certificate is being requested.
        """
        return pulumi.get(self, "uris")

    @uris.setter
    def uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "uris", value)


class RootCa(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input['Algorithm']] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ecdsa_curve: Optional[pulumi.Input['EcdsaCurve']] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 set_authority_key_id: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subject: Optional[pulumi.Input[pulumi.InputType['pulumi_tls.SelfSignedCertSubjectArgs']]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a RootCa resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['Algorithm'] algorithm: Name of the algorithm to use when generating the private key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: TODO
        :param pulumi.Input['EcdsaCurve'] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        :param pulumi.Input[bool] set_authority_key_id: Should the generated certificate include an authority key identifier.
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a subject key identifier.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RootCaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RootCa resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RootCaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RootCaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input['Algorithm']] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ecdsa_curve: Optional[pulumi.Input['EcdsaCurve']] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 set_authority_key_id: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subject: Optional[pulumi.Input[pulumi.InputType['pulumi_tls.SelfSignedCertSubjectArgs']]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = RootCaArgs.__new__(RootCaArgs)

            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["dns_names"] = dns_names
            __props__.__dict__["early_renewal_hours"] = early_renewal_hours
            __props__.__dict__["ecdsa_curve"] = ecdsa_curve
            __props__.__dict__["ip_addresses"] = ip_addresses
            __props__.__dict__["rsa_bits"] = rsa_bits
            __props__.__dict__["set_authority_key_id"] = set_authority_key_id
            __props__.__dict__["set_subject_key_id"] = set_subject_key_id
            __props__.__dict__["subject"] = subject
            __props__.__dict__["uris"] = uris
            if validity_period_hours is None and not opts.urn:
                raise TypeError("Missing required property 'validity_period_hours'")
            __props__.__dict__["validity_period_hours"] = validity_period_hours
            __props__.__dict__["allowed_uses"] = None
            __props__.__dict__["cert"] = None
            __props__.__dict__["cert_pem"] = None
            __props__.__dict__["key"] = None
            __props__.__dict__["private_key_pem"] = None
            __props__.__dict__["public_key_pem"] = None
        super(RootCa, __self__).__init__(
            'kubernetes-the-hard-way:index:RootCa',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter(name="allowedUses")
    def allowed_uses(self) -> pulumi.Output[Sequence['AllowedUsage']]:
        return pulumi.get(self, "allowed_uses")

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Output['pulumi_tls.SelfSignedCert']:
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cert_pem")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output['pulumi_tls.PrivateKey']:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_key_pem")

    @property
    @pulumi.getter(name="publicKeyPem")
    def public_key_pem(self) -> pulumi.Output[str]:
        return pulumi.get(self, "public_key_pem")

    @pulumi.output_type
    class InstallCertResult:
        def __init__(__self__, result=None):
            if result and not isinstance(result, RemoteFile):
                raise TypeError("Expected argument 'result' to be a RemoteFile")
            pulumi.set(__self__, "result", result)

        @property
        @pulumi.getter
        def result(self) -> 'RemoteFile':
            return pulumi.get(self, "result")

    def install_cert(__self__, *,
                     connection: pulumi.Input['ConnectionArgs'],
                     name: str,
                     options: Optional['ResourceOptionsArgs'] = None,
                     path: Optional[pulumi.Input[str]] = None) -> pulumi.Output['RemoteFile']:
        """
        Creates a RemoteFile resource representing the copy operation.


        :param pulumi.Input['ConnectionArgs'] connection: The connection details.
        :param pulumi.Input[str] path: The path to install to.
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['connection'] = connection
        __args__['name'] = name
        __args__['options'] = options
        __args__['path'] = path
        __result__ = pulumi.runtime.call('kubernetes-the-hard-way:index:RootCa/installCert', __args__, res=__self__, typ=RootCa.InstallCertResult)
        return __result__.result

    @pulumi.output_type
    class InstallKeyResult:
        def __init__(__self__, result=None):
            if result and not isinstance(result, RemoteFile):
                raise TypeError("Expected argument 'result' to be a RemoteFile")
            pulumi.set(__self__, "result", result)

        @property
        @pulumi.getter
        def result(self) -> 'RemoteFile':
            return pulumi.get(self, "result")

    def install_key(__self__, *,
                    connection: pulumi.Input['ConnectionArgs'],
                    name: str,
                    options: Optional['ResourceOptionsArgs'] = None,
                    path: Optional[pulumi.Input[str]] = None) -> pulumi.Output['RemoteFile']:
        """
        Creates a RemoteFile resource representing the copy operation.


        :param pulumi.Input['ConnectionArgs'] connection: The connection details.
        :param pulumi.Input[str] path: The path to install to.
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['connection'] = connection
        __args__['name'] = name
        __args__['options'] = options
        __args__['path'] = path
        __result__ = pulumi.runtime.call('kubernetes-the-hard-way:index:RootCa/installKey', __args__, res=__self__, typ=RootCa.InstallKeyResult)
        return __result__.result

    @pulumi.output_type
    class NewCertificateResult:
        def __init__(__self__, result=None):
            if result and not isinstance(result, Certificate):
                raise TypeError("Expected argument 'result' to be a Certificate")
            pulumi.set(__self__, "result", result)

        @property
        @pulumi.getter
        def result(self) -> 'Certificate':
            """
            The issued certificate.
            """
            return pulumi.get(self, "result")

    def new_certificate(__self__, *,
                        algorithm: pulumi.Input['Algorithm'],
                        allowed_uses: pulumi.Input[Sequence[pulumi.Input['AllowedUsage']]],
                        name: str,
                        validity_period_hours: pulumi.Input[int],
                        dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                        early_renewal_hours: Optional[pulumi.Input[int]] = None,
                        ecdsa_curve: Optional[pulumi.Input['EcdsaCurve']] = None,
                        ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                        is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                        options: Optional['ResourceOptionsArgs'] = None,
                        rsa_bits: Optional[pulumi.Input[int]] = None,
                        set_authority_key_id: Optional[pulumi.Input[bool]] = None,
                        set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                        subject: Optional[pulumi.Input['CertRequestSubjectArgs']] = None,
                        uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> pulumi.Output['Certificate']:
        """
        Creates a Certificate configured for the current authority.


        :param pulumi.Input['Algorithm'] algorithm: Name of the algorithm to use when generating the private key.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: TODO
        :param pulumi.Input['EcdsaCurve'] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        :param pulumi.Input[bool] set_authority_key_id: Should the generated certificate include an authority key identifier.
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a subject key identifier.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['algorithm'] = algorithm
        __args__['allowedUses'] = allowed_uses
        __args__['name'] = name
        __args__['validityPeriodHours'] = validity_period_hours
        __args__['dnsNames'] = dns_names
        __args__['earlyRenewalHours'] = early_renewal_hours
        __args__['ecdsaCurve'] = ecdsa_curve
        __args__['ipAddresses'] = ip_addresses
        __args__['isCaCertificate'] = is_ca_certificate
        __args__['options'] = options
        __args__['rsaBits'] = rsa_bits
        __args__['setAuthorityKeyId'] = set_authority_key_id
        __args__['setSubjectKeyId'] = set_subject_key_id
        __args__['subject'] = subject
        __args__['uris'] = uris
        __result__ = pulumi.runtime.call('kubernetes-the-hard-way:index:RootCa/newCertificate', __args__, res=__self__, typ=RootCa.NewCertificateResult)
        return __result__.result

