package provider

import (
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemoteCertificateArgs struct {
	CertificateArgs

	Connection remote.ConnectionArgs `pulumi:"connection"`
	Path       pulumi.StringInput    `pulumi:"path"`
}

type RemoteCertificate struct {
	pulumi.ResourceState
	Certificate
	RemoteFile
}

func NewRemoteCertificate(ctx *pulumi.Context,
	name string, args *RemoteCertificateArgs, opts ...pulumi.ResourceOption) (*RemoteCertificate, error) {
	if args == nil {
		args = &RemoteCertificateArgs{}
	}

	component := &RemoteCertificate{}
	err := ctx.RegisterComponentResource("kubernetes-the-hard-way:index:RemoteCertificate", name, component, opts...)
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

	file, err := NewRemoteFile(ctx, name, &RemoteFileArgs{
		Connection: args.Connection,
		Content:    cert.CertPem,
		Path:       args.Path,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.Csr = cert.Csr
	component.Cert = cert.Cert
	component.CertPem = cert.CertPem
	component.Key = cert.Key
	component.KeyPem = cert.KeyPem
	component.Command = file.Command

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"cert":    cert.Cert,
		"certPem": cert.CertPem,
		"csr":     cert.Csr,
		"key":     cert.Key,
		"keyPem":  cert.KeyPem,
		"command": file.Command,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
