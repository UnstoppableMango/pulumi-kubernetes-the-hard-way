// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Root certificate authority for a cluster.
type RootCa struct {
	pulumi.ResourceState

	// Name of the algorithm to use when generating the private key.
	Algorithm AlgorithmOutput `pulumi:"algorithm"`
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
	AllowedUses pulumi.StringArrayOutput `pulumi:"allowedUses"`
	// The certificate authority certificate.
	Cert tls.SelfSignedCertOutput `pulumi:"cert"`
	// Certificate data in PEM (RFC 1421).
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          pulumi.StringArrayOutput `pulumi:"dnsNames"`
	EarlyRenewalHours pulumi.IntOutput         `pulumi:"earlyRenewalHours"`
	// TODO
	EcdsaCurve EcdsaCurveOutput `pulumi:"ecdsaCurve"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolOutput `pulumi:"isCaCertificate"`
	// The certificate authority key.
	Key tls.PrivateKeyOutput `pulumi:"key"`
	// Name of the algorithm used when generating the private key provided in `private_key_pem`.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
	PrivateKeyOpenssh pulumi.StringOutput `pulumi:"privateKeyOpenssh"`
	// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	PrivateKeyPem pulumi.StringOutput `pulumi:"privateKeyPem"`
	// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
	PrivateKeyPemPkcs8 pulumi.StringOutput `pulumi:"privateKeyPemPkcs8"`
	// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintMd5 pulumi.StringOutput `pulumi:"publicKeyFingerprintMd5"`
	// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintSha256 pulumi.StringOutput `pulumi:"publicKeyFingerprintSha256"`
	// The public key data in "Authorized Keys".
	PublicKeyOpenssh pulumi.StringOutput `pulumi:"publicKeyOpenssh"`
	// Public key data in PEM (RFC 1421).
	PublicKeyPem pulumi.StringOutput `pulumi:"publicKeyPem"`
	// Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
	ReadyForRenewal pulumi.BoolOutput `pulumi:"readyForRenewal"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits pulumi.IntOutput `pulumi:"rsaBits"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId pulumi.BoolOutput `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolOutput `pulumi:"setSubjectKeyId"`
	// TODO
	Subject tls.SelfSignedCertSubjectPtrOutput `pulumi:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayOutput `pulumi:"uris"`
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringOutput `pulumi:"validityEndTime"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntOutput `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringOutput `pulumi:"validityStartTime"`
}

// NewRootCa registers a new resource with the given unique name, arguments, and options.
func NewRootCa(ctx *pulumi.Context,
	name string, args *RootCaArgs, opts ...pulumi.ResourceOption) (*RootCa, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ValidityPeriodHours == nil {
		return nil, errors.New("invalid value for required argument 'ValidityPeriodHours'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RootCa
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tls:RootCa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type rootCaArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm *Algorithm `pulumi:"algorithm"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          []string `pulumi:"dnsNames"`
	EarlyRenewalHours *int     `pulumi:"earlyRenewalHours"`
	// TODO
	EcdsaCurve *EcdsaCurve `pulumi:"ecdsaCurve"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses []string `pulumi:"ipAddresses"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits *int `pulumi:"rsaBits"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId *bool `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// TODO
	Subject *tls.SelfSignedCertSubject `pulumi:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris []string `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a RootCa resource.
type RootCaArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm AlgorithmPtrInput
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          pulumi.StringArrayInput
	EarlyRenewalHours pulumi.IntPtrInput
	// TODO
	EcdsaCurve EcdsaCurvePtrInput
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayInput
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits pulumi.IntPtrInput
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId pulumi.BoolPtrInput
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrInput
	// TODO
	Subject tls.SelfSignedCertSubjectPtrInput
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayInput
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntInput
}

func (RootCaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rootCaArgs)(nil)).Elem()
}

type RootCaInput interface {
	pulumi.Input

	ToRootCaOutput() RootCaOutput
	ToRootCaOutputWithContext(ctx context.Context) RootCaOutput
}

func (*RootCa) ElementType() reflect.Type {
	return reflect.TypeOf((**RootCa)(nil)).Elem()
}

func (i *RootCa) ToRootCaOutput() RootCaOutput {
	return i.ToRootCaOutputWithContext(context.Background())
}

func (i *RootCa) ToRootCaOutputWithContext(ctx context.Context) RootCaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RootCaOutput)
}

// RootCaArrayInput is an input type that accepts RootCaArray and RootCaArrayOutput values.
// You can construct a concrete instance of `RootCaArrayInput` via:
//
//	RootCaArray{ RootCaArgs{...} }
type RootCaArrayInput interface {
	pulumi.Input

	ToRootCaArrayOutput() RootCaArrayOutput
	ToRootCaArrayOutputWithContext(context.Context) RootCaArrayOutput
}

type RootCaArray []RootCaInput

func (RootCaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RootCa)(nil)).Elem()
}

func (i RootCaArray) ToRootCaArrayOutput() RootCaArrayOutput {
	return i.ToRootCaArrayOutputWithContext(context.Background())
}

func (i RootCaArray) ToRootCaArrayOutputWithContext(ctx context.Context) RootCaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RootCaArrayOutput)
}

// RootCaMapInput is an input type that accepts RootCaMap and RootCaMapOutput values.
// You can construct a concrete instance of `RootCaMapInput` via:
//
//	RootCaMap{ "key": RootCaArgs{...} }
type RootCaMapInput interface {
	pulumi.Input

	ToRootCaMapOutput() RootCaMapOutput
	ToRootCaMapOutputWithContext(context.Context) RootCaMapOutput
}

type RootCaMap map[string]RootCaInput

func (RootCaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RootCa)(nil)).Elem()
}

func (i RootCaMap) ToRootCaMapOutput() RootCaMapOutput {
	return i.ToRootCaMapOutputWithContext(context.Background())
}

func (i RootCaMap) ToRootCaMapOutputWithContext(ctx context.Context) RootCaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RootCaMapOutput)
}

type RootCaOutput struct{ *pulumi.OutputState }

func (RootCaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RootCa)(nil)).Elem()
}

func (o RootCaOutput) ToRootCaOutput() RootCaOutput {
	return o
}

func (o RootCaOutput) ToRootCaOutputWithContext(ctx context.Context) RootCaOutput {
	return o
}

// Name of the algorithm to use when generating the private key.
func (o RootCaOutput) Algorithm() AlgorithmOutput {
	return o.ApplyT(func(v *RootCa) AlgorithmOutput { return v.Algorithm }).(AlgorithmOutput)
}

// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
func (o RootCaOutput) AllowedUses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringArrayOutput { return v.AllowedUses }).(pulumi.StringArrayOutput)
}

// The certificate authority certificate.
func (o RootCaOutput) Cert() tls.SelfSignedCertOutput {
	return o.ApplyT(func(v *RootCa) tls.SelfSignedCertOutput { return v.Cert }).(tls.SelfSignedCertOutput)
}

// Certificate data in PEM (RFC 1421).
func (o RootCaOutput) CertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.CertPem }).(pulumi.StringOutput)
}

// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
func (o RootCaOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringArrayOutput { return v.DnsNames }).(pulumi.StringArrayOutput)
}

func (o RootCaOutput) EarlyRenewalHours() pulumi.IntOutput {
	return o.ApplyT(func(v *RootCa) pulumi.IntOutput { return v.EarlyRenewalHours }).(pulumi.IntOutput)
}

// TODO
func (o RootCaOutput) EcdsaCurve() EcdsaCurveOutput {
	return o.ApplyT(func(v *RootCa) EcdsaCurveOutput { return v.EcdsaCurve }).(EcdsaCurveOutput)
}

// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
func (o RootCaOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
func (o RootCaOutput) IsCaCertificate() pulumi.BoolOutput {
	return o.ApplyT(func(v *RootCa) pulumi.BoolOutput { return v.IsCaCertificate }).(pulumi.BoolOutput)
}

// The certificate authority key.
func (o RootCaOutput) Key() tls.PrivateKeyOutput {
	return o.ApplyT(func(v *RootCa) tls.PrivateKeyOutput { return v.Key }).(tls.PrivateKeyOutput)
}

// Name of the algorithm used when generating the private key provided in `private_key_pem`.
func (o RootCaOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
func (o RootCaOutput) PrivateKeyOpenssh() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.PrivateKeyOpenssh }).(pulumi.StringOutput)
}

// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o RootCaOutput) PrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.PrivateKeyPem }).(pulumi.StringOutput)
}

// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
func (o RootCaOutput) PrivateKeyPemPkcs8() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.PrivateKeyPemPkcs8 }).(pulumi.StringOutput)
}

// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
func (o RootCaOutput) PublicKeyFingerprintMd5() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.PublicKeyFingerprintMd5 }).(pulumi.StringOutput)
}

// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
func (o RootCaOutput) PublicKeyFingerprintSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.PublicKeyFingerprintSha256 }).(pulumi.StringOutput)
}

// The public key data in "Authorized Keys".
func (o RootCaOutput) PublicKeyOpenssh() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.PublicKeyOpenssh }).(pulumi.StringOutput)
}

// Public key data in PEM (RFC 1421).
func (o RootCaOutput) PublicKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.PublicKeyPem }).(pulumi.StringOutput)
}

// Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
func (o RootCaOutput) ReadyForRenewal() pulumi.BoolOutput {
	return o.ApplyT(func(v *RootCa) pulumi.BoolOutput { return v.ReadyForRenewal }).(pulumi.BoolOutput)
}

// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
func (o RootCaOutput) RsaBits() pulumi.IntOutput {
	return o.ApplyT(func(v *RootCa) pulumi.IntOutput { return v.RsaBits }).(pulumi.IntOutput)
}

// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
func (o RootCaOutput) SetAuthorityKeyId() pulumi.BoolOutput {
	return o.ApplyT(func(v *RootCa) pulumi.BoolOutput { return v.SetAuthorityKeyId }).(pulumi.BoolOutput)
}

// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
func (o RootCaOutput) SetSubjectKeyId() pulumi.BoolOutput {
	return o.ApplyT(func(v *RootCa) pulumi.BoolOutput { return v.SetSubjectKeyId }).(pulumi.BoolOutput)
}

// TODO
func (o RootCaOutput) Subject() tls.SelfSignedCertSubjectPtrOutput {
	return o.ApplyT(func(v *RootCa) tls.SelfSignedCertSubjectPtrOutput { return v.Subject }).(tls.SelfSignedCertSubjectPtrOutput)
}

// List of URIs for which a certificate is being requested (i.e. certificate subjects).
func (o RootCaOutput) Uris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringArrayOutput { return v.Uris }).(pulumi.StringArrayOutput)
}

// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o RootCaOutput) ValidityEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.ValidityEndTime }).(pulumi.StringOutput)
}

// Number of hours, after initial issuing, that the certificate will remain valid for.
func (o RootCaOutput) ValidityPeriodHours() pulumi.IntOutput {
	return o.ApplyT(func(v *RootCa) pulumi.IntOutput { return v.ValidityPeriodHours }).(pulumi.IntOutput)
}

// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o RootCaOutput) ValidityStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RootCa) pulumi.StringOutput { return v.ValidityStartTime }).(pulumi.StringOutput)
}

type RootCaArrayOutput struct{ *pulumi.OutputState }

func (RootCaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RootCa)(nil)).Elem()
}

func (o RootCaArrayOutput) ToRootCaArrayOutput() RootCaArrayOutput {
	return o
}

func (o RootCaArrayOutput) ToRootCaArrayOutputWithContext(ctx context.Context) RootCaArrayOutput {
	return o
}

func (o RootCaArrayOutput) Index(i pulumi.IntInput) RootCaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RootCa {
		return vs[0].([]*RootCa)[vs[1].(int)]
	}).(RootCaOutput)
}

type RootCaMapOutput struct{ *pulumi.OutputState }

func (RootCaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RootCa)(nil)).Elem()
}

func (o RootCaMapOutput) ToRootCaMapOutput() RootCaMapOutput {
	return o
}

func (o RootCaMapOutput) ToRootCaMapOutputWithContext(ctx context.Context) RootCaMapOutput {
	return o
}

func (o RootCaMapOutput) MapIndex(k pulumi.StringInput) RootCaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RootCa {
		return vs[0].(map[string]*RootCa)[vs[1].(string)]
	}).(RootCaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RootCaInput)(nil)).Elem(), &RootCa{})
	pulumi.RegisterInputType(reflect.TypeOf((*RootCaArrayInput)(nil)).Elem(), RootCaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RootCaMapInput)(nil)).Elem(), RootCaMap{})
	pulumi.RegisterOutputType(RootCaOutput{})
	pulumi.RegisterOutputType(RootCaArrayOutput{})
	pulumi.RegisterOutputType(RootCaMapOutput{})
}
