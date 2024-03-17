package provider

import (
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RootCaArgs struct {
	KeyPairArgs

	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subject tls.SelfSignedCertSubjectPtrInput `pulumi:"subject"`
}

// The RootCa for a cluster
type RootCa struct {
	pulumi.ResourceState
	KeyPair

	Cert *tls.SelfSignedCert `pulumi:"cert"`
}

// NewRootCa creates a new RootCa component resource.
func NewRootCa(ctx *pulumi.Context,
	name string, args *RootCaArgs, opts ...pulumi.ResourceOption) (*RootCa, error) {
	if args == nil {
		args = &RootCaArgs{}
	}

	component := &RootCa{}
	err := ctx.RegisterComponentResource("kubernetes-the-hard-way:index:RootCa", name, component, opts...)
	if err != nil {
		return nil, err
	}

	key, err := tls.NewPrivateKey(ctx, name, &tls.PrivateKeyArgs{
		Algorithm:  pulumi.String(args.Algorithm),
		EcdsaCurve: args.EcdsaCurve,
		RsaBits:    args.RsaBits,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	cert, err := tls.NewSelfSignedCert(ctx, name, &tls.SelfSignedCertArgs{
		AllowedUses: pulumi.StringArray{
			pulumi.String(CertSigning),
			pulumi.String(KeyEncipherment),
			pulumi.String(ServerAuth),
			pulumi.String(ClientAuth),
		},
		DnsNames:            args.DnsNames,
		EarlyRenewalHours:   args.EarlyRenewalHours,
		IpAddresses:         args.IpAddresses,
		IsCaCertificate:     pulumi.Bool(true),
		PrivateKeyPem:       key.PrivateKeyPem,
		SetAuthorityKeyId:   args.SetAuthorityKeyId,
		SetSubjectKeyId:     args.SetSubjectKeyId,
		Subject:             args.Subject,
		Uris:                args.Uris,
		ValidityPeriodHours: args.ValidityPeriodHours,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.Cert = cert
	component.CertPem = cert.CertPem
	component.Key = key
	component.KeyPem = key.PrivateKeyPem // TODO: Have I had this backwards the whole time??

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"cert":    cert,
		"certPem": cert.CertPem,
		"key":     key,
		"keyPem":  key.PrivateKeyPem,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

// CreateCertificateArgs is the set of arguments for creating a Certificate resource
type CreateCertificateArgs struct {
	// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
	Algorithm string `pulumi:"algorithm"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
	EcdsaCurve *string `pulumi:"ecdsaCurve"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits *int `pulumi:"rsaBits"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames []string `pulumi:"dnsNames"`
	// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
	// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
	// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
	// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
	// early renewal period. (default: `0`)
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses []string `pulumi:"ipAddresses"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId *bool `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris []string `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`

	AllowedUses     []string                   `pulumi:"allowedUses"`
	IsCaCertificate bool                       `pulumi:"isCaCertificate"`
	Name            string                     `pulumi:"name"`
	Subject         tls.CertRequestSubjectArgs `pulumi:"subject"`
}

type CreateCertificateResult struct {
	Cert *Certificate `pulumi:"cert"`
}

func (c *RootCa) CreateCertificate(ctx *pulumi.Context, args CreateCertificateArgs) (*CreateCertificateResult, error) {
	cert, err := NewCertificate(ctx, args.Name, &CertificateArgs{
		KeyPairArgs: KeyPairArgs{
			DnsNames:            pulumi.ToStringArray(args.DnsNames),
			EarlyRenewalHours:   pulumi.IntPtrFromPtr(args.EarlyRenewalHours),
			IpAddresses:         pulumi.ToStringArray(args.IpAddresses),
			SetAuthorityKeyId:   pulumi.BoolPtrFromPtr(args.SetAuthorityKeyId),
			SetSubjectKeyId:     pulumi.BoolPtrFromPtr(args.SetSubjectKeyId),
			Uris:                pulumi.ToStringArray(args.Uris),
			ValidityPeriodHours: pulumi.Int(args.ValidityPeriodHours),
		},
		AllowedUses:     pulumi.ToStringArray(args.AllowedUses),
		IsCaCertificate: pulumi.Bool(args.IsCaCertificate),
		Subject:         tls.CertRequestSubjectPtr(&args.Subject),
	})
	if err != nil {
		return nil, err
	}

	return &CreateCertificateResult{Cert: cert}, nil
}

func (c *RootCa) InstallOn(ctx *pulumi.Context, args InstallOnArgs) (*InstallOnResult, error) {
	file, err := NewRemoteFile(ctx, args.Name, &RemoteFileArgs{
		Connection: args.Connection,
		Content:    c.CertPem,
		Path:       args.Path,
	})
	if err != nil {
		return nil, err
	}

	return &InstallOnResult{File: file}, nil
}
