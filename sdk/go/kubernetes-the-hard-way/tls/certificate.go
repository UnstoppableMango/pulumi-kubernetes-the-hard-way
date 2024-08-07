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

// A certificate key pair.
type Certificate struct {
	pulumi.ResourceState

	// Name of the algorithm to use when generating the private key.
	Algorithm AlgorithmOutput `pulumi:"algorithm"`
	// List of key usages allowed for the issued certificate.
	AllowedUses AllowedUsageArrayOutput `pulumi:"allowedUses"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem pulumi.StringOutput `pulumi:"caCertPem"`
	// Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
	CaKeyAlgorithm pulumi.StringOutput `pulumi:"caKeyAlgorithm"`
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem pulumi.StringOutput `pulumi:"caPrivateKeyPem"`
	// The certificate.
	Cert tls.LocallySignedCertOutput `pulumi:"cert"`
	// Certificate data in PEM (RFC 1421).
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem pulumi.StringOutput `pulumi:"certRequestPem"`
	// The certificate signing request.
	Csr tls.CertRequestOutput `pulumi:"csr"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          pulumi.StringArrayOutput `pulumi:"dnsNames"`
	EarlyRenewalHours pulumi.IntOutput         `pulumi:"earlyRenewalHours"`
	// TODO
	EcdsaCurve EcdsaCurveOutput `pulumi:"ecdsaCurve"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolOutput `pulumi:"isCaCertificate"`
	// The private key
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
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolOutput `pulumi:"setSubjectKeyId"`
	// TODO
	Subject tls.CertRequestSubjectPtrOutput `pulumi:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayOutput `pulumi:"uris"`
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringOutput `pulumi:"validityEndTime"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntOutput `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringOutput `pulumi:"validityStartTime"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	if args.AllowedUses == nil {
		return nil, errors.New("invalid value for required argument 'AllowedUses'")
	}
	if args.CaCertPem == nil {
		return nil, errors.New("invalid value for required argument 'CaCertPem'")
	}
	if args.CaPrivateKeyPem == nil {
		return nil, errors.New("invalid value for required argument 'CaPrivateKeyPem'")
	}
	if args.ValidityPeriodHours == nil {
		return nil, errors.New("invalid value for required argument 'ValidityPeriodHours'")
	}
	if args.CaPrivateKeyPem != nil {
		args.CaPrivateKeyPem = pulumi.ToSecret(args.CaPrivateKeyPem).(pulumi.StringInput)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tls:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type certificateArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm Algorithm `pulumi:"algorithm"`
	// List of key usages allowed for the issued certificate.
	AllowedUses []AllowedUsage `pulumi:"allowedUses"`
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem string `pulumi:"caCertPem"`
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem string `pulumi:"caPrivateKeyPem"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          []string `pulumi:"dnsNames"`
	EarlyRenewalHours *int     `pulumi:"earlyRenewalHours"`
	// TODO
	EcdsaCurve *EcdsaCurve `pulumi:"ecdsaCurve"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses []string `pulumi:"ipAddresses"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits *int `pulumi:"rsaBits"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// TODO
	Subject *tls.CertRequestSubject `pulumi:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris []string `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm AlgorithmInput
	// List of key usages allowed for the issued certificate.
	AllowedUses AllowedUsageArrayInput
	// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaCertPem pulumi.StringInput
	// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CaPrivateKeyPem pulumi.StringInput
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          pulumi.StringArrayInput
	EarlyRenewalHours pulumi.IntPtrInput
	// TODO
	EcdsaCurve EcdsaCurvePtrInput
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayInput
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolPtrInput
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits pulumi.IntPtrInput
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrInput
	// TODO
	Subject tls.CertRequestSubjectPtrInput
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayInput
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

// CertificateArrayInput is an input type that accepts CertificateArray and CertificateArrayOutput values.
// You can construct a concrete instance of `CertificateArrayInput` via:
//
//	CertificateArray{ CertificateArgs{...} }
type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

// CertificateMapInput is an input type that accepts CertificateMap and CertificateMapOutput values.
// You can construct a concrete instance of `CertificateMapInput` via:
//
//	CertificateMap{ "key": CertificateArgs{...} }
type CertificateMapInput interface {
	pulumi.Input

	ToCertificateMapOutput() CertificateMapOutput
	ToCertificateMapOutputWithContext(context.Context) CertificateMapOutput
}

type CertificateMap map[string]CertificateInput

func (CertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (i CertificateMap) ToCertificateMapOutput() CertificateMapOutput {
	return i.ToCertificateMapOutputWithContext(context.Background())
}

func (i CertificateMap) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateMapOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// Name of the algorithm to use when generating the private key.
func (o CertificateOutput) Algorithm() AlgorithmOutput {
	return o.ApplyT(func(v *Certificate) AlgorithmOutput { return v.Algorithm }).(AlgorithmOutput)
}

// List of key usages allowed for the issued certificate.
func (o CertificateOutput) AllowedUses() AllowedUsageArrayOutput {
	return o.ApplyT(func(v *Certificate) AllowedUsageArrayOutput { return v.AllowedUses }).(AllowedUsageArrayOutput)
}

// Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o CertificateOutput) CaCertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CaCertPem }).(pulumi.StringOutput)
}

// Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
func (o CertificateOutput) CaKeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CaKeyAlgorithm }).(pulumi.StringOutput)
}

// Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o CertificateOutput) CaPrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CaPrivateKeyPem }).(pulumi.StringOutput)
}

// The certificate.
func (o CertificateOutput) Cert() tls.LocallySignedCertOutput {
	return o.ApplyT(func(v *Certificate) tls.LocallySignedCertOutput { return v.Cert }).(tls.LocallySignedCertOutput)
}

// Certificate data in PEM (RFC 1421).
func (o CertificateOutput) CertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertPem }).(pulumi.StringOutput)
}

// Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o CertificateOutput) CertRequestPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertRequestPem }).(pulumi.StringOutput)
}

// The certificate signing request.
func (o CertificateOutput) Csr() tls.CertRequestOutput {
	return o.ApplyT(func(v *Certificate) tls.CertRequestOutput { return v.Csr }).(tls.CertRequestOutput)
}

// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
func (o CertificateOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.DnsNames }).(pulumi.StringArrayOutput)
}

func (o CertificateOutput) EarlyRenewalHours() pulumi.IntOutput {
	return o.ApplyT(func(v *Certificate) pulumi.IntOutput { return v.EarlyRenewalHours }).(pulumi.IntOutput)
}

// TODO
func (o CertificateOutput) EcdsaCurve() EcdsaCurveOutput {
	return o.ApplyT(func(v *Certificate) EcdsaCurveOutput { return v.EcdsaCurve }).(EcdsaCurveOutput)
}

// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
func (o CertificateOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
func (o CertificateOutput) IsCaCertificate() pulumi.BoolOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolOutput { return v.IsCaCertificate }).(pulumi.BoolOutput)
}

// The private key
func (o CertificateOutput) Key() tls.PrivateKeyOutput {
	return o.ApplyT(func(v *Certificate) tls.PrivateKeyOutput { return v.Key }).(tls.PrivateKeyOutput)
}

// Name of the algorithm used when generating the private key provided in `private_key_pem`.
func (o CertificateOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
func (o CertificateOutput) PrivateKeyOpenssh() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PrivateKeyOpenssh }).(pulumi.StringOutput)
}

// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o CertificateOutput) PrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PrivateKeyPem }).(pulumi.StringOutput)
}

// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
func (o CertificateOutput) PrivateKeyPemPkcs8() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PrivateKeyPemPkcs8 }).(pulumi.StringOutput)
}

// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
func (o CertificateOutput) PublicKeyFingerprintMd5() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicKeyFingerprintMd5 }).(pulumi.StringOutput)
}

// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
func (o CertificateOutput) PublicKeyFingerprintSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicKeyFingerprintSha256 }).(pulumi.StringOutput)
}

// The public key data in "Authorized Keys".
func (o CertificateOutput) PublicKeyOpenssh() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicKeyOpenssh }).(pulumi.StringOutput)
}

// Public key data in PEM (RFC 1421).
func (o CertificateOutput) PublicKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicKeyPem }).(pulumi.StringOutput)
}

// Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
func (o CertificateOutput) ReadyForRenewal() pulumi.BoolOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolOutput { return v.ReadyForRenewal }).(pulumi.BoolOutput)
}

// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
func (o CertificateOutput) RsaBits() pulumi.IntOutput {
	return o.ApplyT(func(v *Certificate) pulumi.IntOutput { return v.RsaBits }).(pulumi.IntOutput)
}

// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
func (o CertificateOutput) SetSubjectKeyId() pulumi.BoolOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolOutput { return v.SetSubjectKeyId }).(pulumi.BoolOutput)
}

// TODO
func (o CertificateOutput) Subject() tls.CertRequestSubjectPtrOutput {
	return o.ApplyT(func(v *Certificate) tls.CertRequestSubjectPtrOutput { return v.Subject }).(tls.CertRequestSubjectPtrOutput)
}

// List of URIs for which a certificate is being requested (i.e. certificate subjects).
func (o CertificateOutput) Uris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.Uris }).(pulumi.StringArrayOutput)
}

// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o CertificateOutput) ValidityEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ValidityEndTime }).(pulumi.StringOutput)
}

// Number of hours, after initial issuing, that the certificate will remain valid for.
func (o CertificateOutput) ValidityPeriodHours() pulumi.IntOutput {
	return o.ApplyT(func(v *Certificate) pulumi.IntOutput { return v.ValidityPeriodHours }).(pulumi.IntOutput)
}

// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o CertificateOutput) ValidityStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ValidityStartTime }).(pulumi.StringOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].([]*Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateMapOutput struct{ *pulumi.OutputState }

func (CertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Certificate)(nil)).Elem()
}

func (o CertificateMapOutput) ToCertificateMapOutput() CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) ToCertificateMapOutputWithContext(ctx context.Context) CertificateMapOutput {
	return o
}

func (o CertificateMapOutput) MapIndex(k pulumi.StringInput) CertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Certificate {
		return vs[0].(map[string]*Certificate)[vs[1].(string)]
	}).(CertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateArrayInput)(nil)).Elem(), CertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateMapInput)(nil)).Elem(), CertificateMap{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
