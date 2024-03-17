package provider

import (
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Algorithm string
type AllowedUse string

const (
	Ecdsa             Algorithm  = "ECDSA"
	Ed25519           Algorithm  = "ED25519"
	Rsa               Algorithm  = "RSA"
	AnyExtended       AllowedUse = "any_extended"
	CertSigning       AllowedUse = "cert_signing"
	ClientAuth        AllowedUse = "client_auth"
	CodeSigning       AllowedUse = "code_signing"
	ContentCommitment AllowedUse = "content_commitment"
	CrlSigning        AllowedUse = "crl_signing"
	DataEncipherment  AllowedUse = "data_encipherment"
	DecipherOnly      AllowedUse = "decipher_only"
	DigitalSignature  AllowedUse = "digital_signature"
	EmailProtection   AllowedUse = "email_protection"
	EncipherOnly      AllowedUse = "encipher_only"
	IpsecEndSystem    AllowedUse = "ipsec_end_system" // TODO: Review
	IpsecTunnel       AllowedUse = "ipsec_tunnel"
	IpsecUser         AllowedUse = "ipsec_user"
	KeyAgreement      AllowedUse = "key_agreement"
	KeyEncipherment   AllowedUse = "key_encipherment"
	OcspSigning       AllowedUse = "ocsp_signing"
	ServerAuth        AllowedUse = "server_auth"
)

type KeyPairArgs struct {
	tls.PrivateKeyArgs

	Algorithm Algorithm `pulumi:"algorithm"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames pulumi.StringArrayInput `pulumi:"dnsNames"`
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
	// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
	// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
	// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
	// early renewal period. (default: `0`)
	EarlyRenewalHours pulumi.IntPtrInput `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayInput `pulumi:"ipAddresses"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId pulumi.BoolPtrInput `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrInput `pulumi:"setSubjectKeyId"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayInput `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntInput `pulumi:"validityPeriodHours"`
}

type KeyPair struct {
	pulumi.ResourceState

	CertPem pulumi.StringOutput `pulumi:"certPem"`
	Key     *tls.PrivateKey     `pulumi:"key"`
	KeyPem  pulumi.StringOutput `pulumi:"keyPem"`
}
