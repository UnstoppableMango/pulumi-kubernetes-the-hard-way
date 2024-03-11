package provider

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type AdminCertificateArgs struct {
	CertificateArgs
}

type AdminCertificate struct {
	pulumi.ResourceState
	Certificate
}

func NewAdminCertificate(ctx *pulumi.Context,
	name string, args *AdminCertificateArgs, opts ...pulumi.ResourceOption) (*AdminCertificate, error) {
	if args == nil {
		args = &AdminCertificateArgs{}
	}

	component := &AdminCertificate{}
	err := ctx.RegisterComponentResource("kubernetes-the-hard-way:index:AdminCertificate", name, component, opts...)
	if err != nil {
		return nil, err
	}

	cert, err := NewCertificate(ctx, name, &CertificateArgs{
		KeyPairArgs:     args.KeyPairArgs,
		AllowedUses:     args.AllowedUses,
		CaCertPem:       args.CaCertPem,
		CaPrivateKeyPem: args.CaPrivateKeyPem,
		IsCaCertificate: args.IsCaCertificate,
		Subject:         args.Subject,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.Cert = cert.Cert
	component.CertPem = cert.CertPem
	component.Csr = cert.Csr
	component.Key = cert.Key
	component.KeyPem = cert.KeyPem

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"cert":    cert.Cert,
		"certPem": cert.CertPem,
		"csr":     cert.Csr,
		"key":     cert.Key,
		"keyPem":  cert.KeyPem,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
