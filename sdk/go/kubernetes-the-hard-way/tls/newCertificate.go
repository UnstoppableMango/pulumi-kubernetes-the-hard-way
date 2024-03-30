// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Creates a Certificate configured for the current authority.
func CreateCertificate(ctx *pulumi.Context, args *CreateCertificateArgs, opts ...pulumi.InvokeOption) (*CreateCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv CreateCertificateResult
	err := ctx.Invoke("kubernetes-the-hard-way:tls:newCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type CreateCertificateArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm   Algorithm      `pulumi:"algorithm"`
	AllowedUses []AllowedUsage `pulumi:"allowedUses"`
	// The certificate authority to issue the certificate.
	Ca *RootCa `pulumi:"ca"`
	// List of DNS names for which a certificate is being requested.
	DnsNames []string `pulumi:"dnsNames"`
	// TODO
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
	EcdsaCurve *EcdsaCurve `pulumi:"ecdsaCurve"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses     []string                              `pulumi:"ipAddresses"`
	IsCaCertificate *bool                                 `pulumi:"isCaCertificate"`
	Name            string                                `pulumi:"name"`
	Options         *kubernetesthehardway.ResourceOptions `pulumi:"options"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
	RsaBits *int `pulumi:"rsaBits"`
	// Should the generated certificate include an authority key identifier.
	SetAuthorityKeyId *bool `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a subject key identifier.
	SetSubjectKeyId *bool                                    `pulumi:"setSubjectKeyId"`
	Subject         *kubernetesthehardway.CertRequestSubject `pulumi:"subject"`
	// List of URIs for which a certificate is being requested.
	Uris []string `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

type CreateCertificateResult struct {
	Result *Certificate `pulumi:"result"`
}

func CreateCertificateOutput(ctx *pulumi.Context, args CreateCertificateOutputArgs, opts ...pulumi.InvokeOption) CreateCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CreateCertificateResult, error) {
			args := v.(CreateCertificateArgs)
			r, err := CreateCertificate(ctx, &args, opts...)
			var s CreateCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CreateCertificateResultOutput)
}

type CreateCertificateOutputArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm   AlgorithmInput         `pulumi:"algorithm"`
	AllowedUses AllowedUsageArrayInput `pulumi:"allowedUses"`
	// The certificate authority to issue the certificate.
	Ca RootCaInput `pulumi:"ca"`
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayInput `pulumi:"dnsNames"`
	// TODO
	EarlyRenewalHours pulumi.IntPtrInput `pulumi:"earlyRenewalHours"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
	EcdsaCurve EcdsaCurvePtrInput `pulumi:"ecdsaCurve"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses     pulumi.StringArrayInput                   `pulumi:"ipAddresses"`
	IsCaCertificate pulumi.BoolPtrInput                       `pulumi:"isCaCertificate"`
	Name            string                                    `pulumi:"name"`
	Options         *kubernetesthehardway.ResourceOptionsArgs `pulumi:"options"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
	RsaBits pulumi.IntPtrInput `pulumi:"rsaBits"`
	// Should the generated certificate include an authority key identifier.
	SetAuthorityKeyId pulumi.BoolPtrInput `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a subject key identifier.
	SetSubjectKeyId pulumi.BoolPtrInput                             `pulumi:"setSubjectKeyId"`
	Subject         kubernetesthehardway.CertRequestSubjectPtrInput `pulumi:"subject"`
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayInput `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid.
	ValidityPeriodHours pulumi.IntInput `pulumi:"validityPeriodHours"`
}

func (CreateCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateCertificateArgs)(nil)).Elem()
}

type CreateCertificateResultOutput struct{ *pulumi.OutputState }

func (CreateCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateCertificateResult)(nil)).Elem()
}

func (o CreateCertificateResultOutput) ToCreateCertificateResultOutput() CreateCertificateResultOutput {
	return o
}

func (o CreateCertificateResultOutput) ToCreateCertificateResultOutputWithContext(ctx context.Context) CreateCertificateResultOutput {
	return o
}

func (o CreateCertificateResultOutput) Result() CertificateOutput {
	return o.ApplyT(func(v CreateCertificateResult) *Certificate { return v.Result }).(CertificateOutput)
}

func init() {
	pulumi.RegisterOutputType(CreateCertificateResultOutput{})
}
