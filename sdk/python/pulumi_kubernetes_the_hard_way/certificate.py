# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *
from .remote_file import RemoteFile
import pulumi_tls

__all__ = ['CertificateArgs', 'Certificate']

@pulumi.input_type
class CertificateArgs:
    def __init__(__self__, *,
                 allowed_uses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 ca_cert_pem: pulumi.Input[str],
                 ca_private_key_pem: pulumi.Input[str],
                 validity_period_hours: pulumi.Input[int],
                 algorithm: Optional[pulumi.Input[str]] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ecdsa_curve: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 set_authority_key_id: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subject: Optional[pulumi.Input['pulumi_tls.CertRequestSubjectArgs']] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Certificate resource.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid.
        :param pulumi.Input[str] algorithm: Name of the algorithm to use when generating the private key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: TODO
        :param pulumi.Input[str] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
        :param pulumi.Input[bool] set_authority_key_id: Should the generated certificate include an authority key identifier.
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a subject key identifier.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        """
        pulumi.set(__self__, "allowed_uses", allowed_uses)
        pulumi.set(__self__, "ca_cert_pem", ca_cert_pem)
        pulumi.set(__self__, "ca_private_key_pem", ca_private_key_pem)
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
        if is_ca_certificate is not None:
            pulumi.set(__self__, "is_ca_certificate", is_ca_certificate)
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
    @pulumi.getter(name="allowedUses")
    def allowed_uses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "allowed_uses")

    @allowed_uses.setter
    def allowed_uses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_uses", value)

    @property
    @pulumi.getter(name="caCertPem")
    def ca_cert_pem(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ca_cert_pem")

    @ca_cert_pem.setter
    def ca_cert_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_cert_pem", value)

    @property
    @pulumi.getter(name="caPrivateKeyPem")
    def ca_private_key_pem(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ca_private_key_pem")

    @ca_private_key_pem.setter
    def ca_private_key_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_private_key_pem", value)

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
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the algorithm to use when generating the private key.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
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
    def ecdsa_curve(self) -> Optional[pulumi.Input[str]]:
        """
        When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
        """
        return pulumi.get(self, "ecdsa_curve")

    @ecdsa_curve.setter
    def ecdsa_curve(self, value: Optional[pulumi.Input[str]]):
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
    @pulumi.getter(name="isCaCertificate")
    def is_ca_certificate(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_ca_certificate")

    @is_ca_certificate.setter
    def is_ca_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ca_certificate", value)

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
    def subject(self) -> Optional[pulumi.Input['pulumi_tls.CertRequestSubjectArgs']]:
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input['pulumi_tls.CertRequestSubjectArgs']]):
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


class Certificate(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_cert_pem: Optional[pulumi.Input[str]] = None,
                 ca_private_key_pem: Optional[pulumi.Input[str]] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ecdsa_curve: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 set_authority_key_id: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subject: Optional[pulumi.Input[pulumi.InputType['pulumi_tls.CertRequestSubjectArgs']]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Certificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: Name of the algorithm to use when generating the private key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: TODO
        :param pulumi.Input[str] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
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
                 args: CertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Certificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_cert_pem: Optional[pulumi.Input[str]] = None,
                 ca_private_key_pem: Optional[pulumi.Input[str]] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ecdsa_curve: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 rsa_bits: Optional[pulumi.Input[int]] = None,
                 set_authority_key_id: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subject: Optional[pulumi.Input[pulumi.InputType['pulumi_tls.CertRequestSubjectArgs']]] = None,
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
            __props__ = CertificateArgs.__new__(CertificateArgs)

            __props__.__dict__["algorithm"] = algorithm
            if allowed_uses is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_uses'")
            __props__.__dict__["allowed_uses"] = allowed_uses
            if ca_cert_pem is None and not opts.urn:
                raise TypeError("Missing required property 'ca_cert_pem'")
            __props__.__dict__["ca_cert_pem"] = ca_cert_pem
            if ca_private_key_pem is None and not opts.urn:
                raise TypeError("Missing required property 'ca_private_key_pem'")
            __props__.__dict__["ca_private_key_pem"] = ca_private_key_pem
            __props__.__dict__["dns_names"] = dns_names
            __props__.__dict__["early_renewal_hours"] = early_renewal_hours
            __props__.__dict__["ecdsa_curve"] = ecdsa_curve
            __props__.__dict__["ip_addresses"] = ip_addresses
            __props__.__dict__["is_ca_certificate"] = is_ca_certificate
            __props__.__dict__["rsa_bits"] = rsa_bits
            __props__.__dict__["set_authority_key_id"] = set_authority_key_id
            __props__.__dict__["set_subject_key_id"] = set_subject_key_id
            __props__.__dict__["subject"] = subject
            __props__.__dict__["uris"] = uris
            if validity_period_hours is None and not opts.urn:
                raise TypeError("Missing required property 'validity_period_hours'")
            __props__.__dict__["validity_period_hours"] = validity_period_hours
            __props__.__dict__["cert"] = None
            __props__.__dict__["cert_pem"] = None
            __props__.__dict__["csr"] = None
            __props__.__dict__["key"] = None
            __props__.__dict__["key_pem"] = None
        super(Certificate, __self__).__init__(
            'kubernetes-the-hard-way:index:Certificate',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Output['pulumi_tls.LocallySignedCert']:
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cert_pem")

    @property
    @pulumi.getter
    def csr(self) -> pulumi.Output['pulumi_tls.CertRequest']:
        return pulumi.get(self, "csr")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output['pulumi_tls.PrivateKey']:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyPem")
    def key_pem(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key_pem")

    @pulumi.output_type
    class InstallOnResult:
        def __init__(__self__, file=None):
            if file and not isinstance(file, RemoteFile):
                raise TypeError("Expected argument 'file' to be a RemoteFile")
            pulumi.set(__self__, "file", file)

        @property
        @pulumi.getter
        def file(self) -> 'RemoteFile':
            return pulumi.get(self, "file")

    def install_on(__self__, *,
                   connection: pulumi.Input['ConnectionArgs'],
                   path: Optional[pulumi.Input[str]] = None) -> pulumi.Output['Certificate.InstallOnResult']:
        """
        Creates a RemoteFile resource representing the copy operation.


        :param pulumi.Input['ConnectionArgs'] connection: The connection details.
        :param pulumi.Input[str] path: The path to install to.
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['connection'] = connection
        __args__['path'] = path
        return pulumi.runtime.call('kubernetes-the-hard-way:index:Certificate/installOn', __args__, res=__self__, typ=Certificate.InstallOnResult)

