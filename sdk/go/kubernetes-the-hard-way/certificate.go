// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesthehardway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

type Certificate struct {
	pulumi.ResourceState

	Cert          tls.LocallySignedCertOutput `pulumi:"cert"`
	CertPem       pulumi.StringOutput         `pulumi:"certPem"`
	Csr           tls.CertRequestOutput       `pulumi:"csr"`
	Key           tls.PrivateKeyOutput        `pulumi:"key"`
	PrivateKeyPem pulumi.StringOutput         `pulumi:"privateKeyPem"`
	PublicKeyPem  pulumi.StringOutput         `pulumi:"publicKeyPem"`
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:index:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type certificateArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm       Algorithm      `pulumi:"algorithm"`
	AllowedUses     []AllowedUsage `pulumi:"allowedUses"`
	CaCertPem       string         `pulumi:"caCertPem"`
	CaPrivateKeyPem string         `pulumi:"caPrivateKeyPem"`
	// List of DNS names for which a certificate is being requested.
	DnsNames []string `pulumi:"dnsNames"`
	// TODO
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
	EcdsaCurve *EcdsaCurve `pulumi:"ecdsaCurve"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses     []string `pulumi:"ipAddresses"`
	IsCaCertificate *bool    `pulumi:"isCaCertificate"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
	RsaBits *int `pulumi:"rsaBits"`
	// Should the generated certificate include an authority key identifier.
	SetAuthorityKeyId *bool `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a subject key identifier.
	SetSubjectKeyId *bool                   `pulumi:"setSubjectKeyId"`
	Subject         *tls.CertRequestSubject `pulumi:"subject"`
	// List of URIs for which a certificate is being requested.
	Uris []string `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm       AlgorithmInput
	AllowedUses     AllowedUsageArrayInput
	CaCertPem       pulumi.StringInput
	CaPrivateKeyPem pulumi.StringInput
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayInput
	// TODO
	EarlyRenewalHours pulumi.IntPtrInput
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
	EcdsaCurve EcdsaCurvePtrInput
	// List of IP addresses for which a certificate is being requested.
	IpAddresses     pulumi.StringArrayInput
	IsCaCertificate pulumi.BoolPtrInput
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
	RsaBits pulumi.IntPtrInput
	// Should the generated certificate include an authority key identifier.
	SetAuthorityKeyId pulumi.BoolPtrInput
	// Should the generated certificate include a subject key identifier.
	SetSubjectKeyId pulumi.BoolPtrInput
	Subject         tls.CertRequestSubjectPtrInput
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayInput
	// Number of hours, after initial issuing, that the certificate will remain valid.
	ValidityPeriodHours pulumi.IntInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

// Creates a RemoteFile resource representing the copy operation.
func (r *Certificate) InstallCert(ctx *pulumi.Context, args *CertificateInstallCertArgs) (RemoteFileOutput, error) {
	out, err := ctx.Call("kubernetes-the-hard-way:index:Certificate/installCert", args, certificateInstallCertResultOutput{}, r)
	if err != nil {
		return RemoteFileOutput{}, err
	}
	return out.(certificateInstallCertResultOutput).Result(), nil
}

type certificateInstallCertArgs struct {
	// The connection details.
	Connection Connection       `pulumi:"connection"`
	Name       string           `pulumi:"name"`
	Opts       *ResourceOptions `pulumi:"opts"`
	// The path to install to.
	Path *string `pulumi:"path"`
}

// The set of arguments for the InstallCert method of the Certificate resource.
type CertificateInstallCertArgs struct {
	// The connection details.
	Connection ConnectionInput
	Name       string
	Opts       *ResourceOptionsArgs
	// The path to install to.
	Path pulumi.StringPtrInput
}

func (CertificateInstallCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateInstallCertArgs)(nil)).Elem()
}

type certificateInstallCertResult struct {
	Result *RemoteFile `pulumi:"result"`
}

type certificateInstallCertResultOutput struct{ *pulumi.OutputState }

func (certificateInstallCertResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateInstallCertResult)(nil)).Elem()
}

func (o certificateInstallCertResultOutput) Result() RemoteFileOutput {
	return o.ApplyT(func(v certificateInstallCertResult) *RemoteFile { return v.Result }).(RemoteFileOutput)
}

// Creates a RemoteFile resource representing the copy operation.
func (r *Certificate) InstallKey(ctx *pulumi.Context, args *CertificateInstallKeyArgs) (RemoteFileOutput, error) {
	out, err := ctx.Call("kubernetes-the-hard-way:index:Certificate/installKey", args, certificateInstallKeyResultOutput{}, r)
	if err != nil {
		return RemoteFileOutput{}, err
	}
	return out.(certificateInstallKeyResultOutput).Result(), nil
}

type certificateInstallKeyArgs struct {
	// The connection details.
	Connection Connection       `pulumi:"connection"`
	Name       string           `pulumi:"name"`
	Opts       *ResourceOptions `pulumi:"opts"`
	// The path to install to.
	Path *string `pulumi:"path"`
}

// The set of arguments for the InstallKey method of the Certificate resource.
type CertificateInstallKeyArgs struct {
	// The connection details.
	Connection ConnectionInput
	Name       string
	Opts       *ResourceOptionsArgs
	// The path to install to.
	Path pulumi.StringPtrInput
}

func (CertificateInstallKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateInstallKeyArgs)(nil)).Elem()
}

type certificateInstallKeyResult struct {
	Result *RemoteFile `pulumi:"result"`
}

type certificateInstallKeyResultOutput struct{ *pulumi.OutputState }

func (certificateInstallKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateInstallKeyResult)(nil)).Elem()
}

func (o certificateInstallKeyResultOutput) Result() RemoteFileOutput {
	return o.ApplyT(func(v certificateInstallKeyResult) *RemoteFile { return v.Result }).(RemoteFileOutput)
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

func (o CertificateOutput) Cert() tls.LocallySignedCertOutput {
	return o.ApplyT(func(v *Certificate) tls.LocallySignedCertOutput { return v.Cert }).(tls.LocallySignedCertOutput)
}

func (o CertificateOutput) CertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertPem }).(pulumi.StringOutput)
}

func (o CertificateOutput) Csr() tls.CertRequestOutput {
	return o.ApplyT(func(v *Certificate) tls.CertRequestOutput { return v.Csr }).(tls.CertRequestOutput)
}

func (o CertificateOutput) Key() tls.PrivateKeyOutput {
	return o.ApplyT(func(v *Certificate) tls.PrivateKeyOutput { return v.Key }).(tls.PrivateKeyOutput)
}

func (o CertificateOutput) PrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PrivateKeyPem }).(pulumi.StringOutput)
}

func (o CertificateOutput) PublicKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicKeyPem }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(certificateInstallCertResultOutput{})
	pulumi.RegisterOutputType(certificateInstallKeyResultOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateMapOutput{})
}
