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
	Cert    *tls.SelfSignedCert `pulumi:"cert"`
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	Key     *tls.PrivateKey     `pulumi:"key"`
	KeyPem  pulumi.StringOutput `pulumi:"keyPem"`
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
		Algorithm:  args.Algorithm,
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
		IsCaCertificate:     pulumi.BoolPtr(true),
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
