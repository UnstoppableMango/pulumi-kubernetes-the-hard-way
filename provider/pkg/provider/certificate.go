package provider

import (
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CertificateArgs struct {
	KeyPairArgs

	AllowedUses     pulumi.StringArrayInput        `pulumi:"allowedUses"`
	CaCertPem       pulumi.StringInput             `pulumi:"caCertPem"`
	CaPrivateKeyPem pulumi.StringInput             `pulumi:"caPrivateKeyPem"`
	IsCaCertificate pulumi.BoolInput               `pulumi:"isCaCertificate"`
	Subject         tls.CertRequestSubjectPtrInput `pulumi:"subject"`
}

type Certificate struct {
	KeyPair

	Cert *tls.LocallySignedCert `pulumi:"cert"`
	Csr  *tls.CertRequest       `pulumi:"csr"`
}

func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		args = &CertificateArgs{}
	}

	component := &Certificate{}
	err := ctx.RegisterComponentResource("kubernetes-the-hard-way:index:Certificate", name, component, opts...)
	if err != nil {
		return nil, err
	}

	key, err := tls.NewPrivateKey(ctx, name, &tls.PrivateKeyArgs{
		Algorithm:  args.Algorithm,
		EcdsaCurve: args.EcdsaCurve,
		RsaBits:    args.RsaBits,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	csr, err := tls.NewCertRequest(ctx, name, &tls.CertRequestArgs{
		DnsNames:      args.DnsNames,
		IpAddresses:   args.IpAddresses,
		PrivateKeyPem: key.PrivateKeyPem,
		Subject:       args.Subject,
		Uris:          args.Uris,
	}, pulumi.Parent(component))

	cert, err := tls.NewLocallySignedCert(ctx, name, &tls.LocallySignedCertArgs{
		AllowedUses:         args.AllowedUses,
		CaCertPem:           args.CaCertPem,
		CaPrivateKeyPem:     args.CaPrivateKeyPem,
		CertRequestPem:      csr.PrivateKeyPem,
		EarlyRenewalHours:   args.EarlyRenewalHours,
		IsCaCertificate:     args.IsCaCertificate,
		SetSubjectKeyId:     args.SetSubjectKeyId,
		ValidityPeriodHours: args.ValidityPeriodHours,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.Cert = cert
	component.CertPem = csr.CertRequestPem
	component.Csr = csr
	component.Key = key
	component.KeyPem = key.PrivateKeyPem

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"cert":    cert,
		"certPem": csr.CertRequestPem,
		"csr":     csr,
		"key":     key,
		"keyPem":  key.PrivateKeyPem,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
