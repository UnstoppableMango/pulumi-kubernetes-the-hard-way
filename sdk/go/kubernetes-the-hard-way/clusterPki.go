// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesthehardway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

type ClusterPki struct {
	pulumi.ResourceState

	// The admin certificate.
	Admin CertificateOutput `pulumi:"admin"`
	// Name of the algorithm to use when generating the private key.
	Algorithm   AlgorithmOutput     `pulumi:"algorithm"`
	Ca          RootCaOutput        `pulumi:"ca"`
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The controller manager certificate.
	ControllerManager CertificateOutput `pulumi:"controllerManager"`
	// The kube proxy certificate.
	KubeProxy CertificateOutput `pulumi:"kubeProxy"`
	// The kube scheduler certificate.
	KubeScheduler CertificateOutput `pulumi:"kubeScheduler"`
	// Map of node name to kubelet certificate.
	Kubelet CertificateMapOutput `pulumi:"kubelet"`
	// The kubernetes certificate.
	Kubernetes CertificateOutput `pulumi:"kubernetes"`
	// The publicly accessible IP for the cluster.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
	RsaBits pulumi.IntOutput `pulumi:"rsaBits"`
	// The service accounts certificate.
	ServiceAccounts CertificateOutput `pulumi:"serviceAccounts"`
	// Number of hours, after initial issuing, that the certificate will remain valid.
	ValidityPeriodHours pulumi.IntOutput `pulumi:"validityPeriodHours"`
}

// NewClusterPki registers a new resource with the given unique name, arguments, and options.
func NewClusterPki(ctx *pulumi.Context,
	name string, args *ClusterPkiArgs, opts ...pulumi.ResourceOption) (*ClusterPki, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Nodes == nil {
		return nil, errors.New("invalid value for required argument 'Nodes'")
	}
	if args.PublicIp == nil {
		return nil, errors.New("invalid value for required argument 'PublicIp'")
	}
	if args.Algorithm == nil {
		args.Algorithm = Algorithm("RSA")
	}
	if args.RsaBits == nil {
		args.RsaBits = pulumi.IntPtr(2048)
	}
	if args.ValidityPeriodHours == nil {
		args.ValidityPeriodHours = pulumi.IntPtr(8076)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterPki
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:index:ClusterPki", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type clusterPkiArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm *Algorithm `pulumi:"algorithm"`
	// A name to use for the cluster
	ClusterName string `pulumi:"clusterName"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
	EcdsaCurve *EcdsaCurve `pulumi:"ecdsaCurve"`
	// Map of node names to node configuration.
	Nodes map[string]ClusterPkiNode `pulumi:"nodes"`
	// Publicly accessible IP address.
	PublicIp string `pulumi:"publicIp"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
	RsaBits *int `pulumi:"rsaBits"`
	// Number of hours, after initial issuing, that the certificate will remain valid.
	ValidityPeriodHours *int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a ClusterPki resource.
type ClusterPkiArgs struct {
	// Name of the algorithm to use when generating the private key.
	Algorithm AlgorithmPtrInput
	// A name to use for the cluster
	ClusterName pulumi.StringInput
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use.
	EcdsaCurve EcdsaCurvePtrInput
	// Map of node names to node configuration.
	Nodes ClusterPkiNodeMapInput
	// Publicly accessible IP address.
	PublicIp pulumi.StringInput
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
	RsaBits pulumi.IntPtrInput
	// Number of hours, after initial issuing, that the certificate will remain valid.
	ValidityPeriodHours pulumi.IntPtrInput
}

func (ClusterPkiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPkiArgs)(nil)).Elem()
}

type ClusterPkiInput interface {
	pulumi.Input

	ToClusterPkiOutput() ClusterPkiOutput
	ToClusterPkiOutputWithContext(ctx context.Context) ClusterPkiOutput
}

func (*ClusterPki) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPki)(nil)).Elem()
}

func (i *ClusterPki) ToClusterPkiOutput() ClusterPkiOutput {
	return i.ToClusterPkiOutputWithContext(context.Background())
}

func (i *ClusterPki) ToClusterPkiOutputWithContext(ctx context.Context) ClusterPkiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPkiOutput)
}

// ClusterPkiArrayInput is an input type that accepts ClusterPkiArray and ClusterPkiArrayOutput values.
// You can construct a concrete instance of `ClusterPkiArrayInput` via:
//
//	ClusterPkiArray{ ClusterPkiArgs{...} }
type ClusterPkiArrayInput interface {
	pulumi.Input

	ToClusterPkiArrayOutput() ClusterPkiArrayOutput
	ToClusterPkiArrayOutputWithContext(context.Context) ClusterPkiArrayOutput
}

type ClusterPkiArray []ClusterPkiInput

func (ClusterPkiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterPki)(nil)).Elem()
}

func (i ClusterPkiArray) ToClusterPkiArrayOutput() ClusterPkiArrayOutput {
	return i.ToClusterPkiArrayOutputWithContext(context.Background())
}

func (i ClusterPkiArray) ToClusterPkiArrayOutputWithContext(ctx context.Context) ClusterPkiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPkiArrayOutput)
}

// ClusterPkiMapInput is an input type that accepts ClusterPkiMap and ClusterPkiMapOutput values.
// You can construct a concrete instance of `ClusterPkiMapInput` via:
//
//	ClusterPkiMap{ "key": ClusterPkiArgs{...} }
type ClusterPkiMapInput interface {
	pulumi.Input

	ToClusterPkiMapOutput() ClusterPkiMapOutput
	ToClusterPkiMapOutputWithContext(context.Context) ClusterPkiMapOutput
}

type ClusterPkiMap map[string]ClusterPkiInput

func (ClusterPkiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterPki)(nil)).Elem()
}

func (i ClusterPkiMap) ToClusterPkiMapOutput() ClusterPkiMapOutput {
	return i.ToClusterPkiMapOutputWithContext(context.Background())
}

func (i ClusterPkiMap) ToClusterPkiMapOutputWithContext(ctx context.Context) ClusterPkiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPkiMapOutput)
}

type ClusterPkiOutput struct{ *pulumi.OutputState }

func (ClusterPkiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPki)(nil)).Elem()
}

func (o ClusterPkiOutput) ToClusterPkiOutput() ClusterPkiOutput {
	return o
}

func (o ClusterPkiOutput) ToClusterPkiOutputWithContext(ctx context.Context) ClusterPkiOutput {
	return o
}

// The admin certificate.
func (o ClusterPkiOutput) Admin() CertificateOutput {
	return o.ApplyT(func(v *ClusterPki) CertificateOutput { return v.Admin }).(CertificateOutput)
}

// Name of the algorithm to use when generating the private key.
func (o ClusterPkiOutput) Algorithm() AlgorithmOutput {
	return o.ApplyT(func(v *ClusterPki) AlgorithmOutput { return v.Algorithm }).(AlgorithmOutput)
}

func (o ClusterPkiOutput) Ca() RootCaOutput {
	return o.ApplyT(func(v *ClusterPki) RootCaOutput { return v.Ca }).(RootCaOutput)
}

func (o ClusterPkiOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPki) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// The controller manager certificate.
func (o ClusterPkiOutput) ControllerManager() CertificateOutput {
	return o.ApplyT(func(v *ClusterPki) CertificateOutput { return v.ControllerManager }).(CertificateOutput)
}

// The kube proxy certificate.
func (o ClusterPkiOutput) KubeProxy() CertificateOutput {
	return o.ApplyT(func(v *ClusterPki) CertificateOutput { return v.KubeProxy }).(CertificateOutput)
}

// The kube scheduler certificate.
func (o ClusterPkiOutput) KubeScheduler() CertificateOutput {
	return o.ApplyT(func(v *ClusterPki) CertificateOutput { return v.KubeScheduler }).(CertificateOutput)
}

// Map of node name to kubelet certificate.
func (o ClusterPkiOutput) Kubelet() CertificateMapOutput {
	return o.ApplyT(func(v *ClusterPki) CertificateMapOutput { return v.Kubelet }).(CertificateMapOutput)
}

// The kubernetes certificate.
func (o ClusterPkiOutput) Kubernetes() CertificateOutput {
	return o.ApplyT(func(v *ClusterPki) CertificateOutput { return v.Kubernetes }).(CertificateOutput)
}

// The publicly accessible IP for the cluster.
func (o ClusterPkiOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPki) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

// When `algorithm` is `RSA`, the size of the generated RSA key, in bits.
func (o ClusterPkiOutput) RsaBits() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterPki) pulumi.IntOutput { return v.RsaBits }).(pulumi.IntOutput)
}

// The service accounts certificate.
func (o ClusterPkiOutput) ServiceAccounts() CertificateOutput {
	return o.ApplyT(func(v *ClusterPki) CertificateOutput { return v.ServiceAccounts }).(CertificateOutput)
}

// Number of hours, after initial issuing, that the certificate will remain valid.
func (o ClusterPkiOutput) ValidityPeriodHours() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterPki) pulumi.IntOutput { return v.ValidityPeriodHours }).(pulumi.IntOutput)
}

type ClusterPkiArrayOutput struct{ *pulumi.OutputState }

func (ClusterPkiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterPki)(nil)).Elem()
}

func (o ClusterPkiArrayOutput) ToClusterPkiArrayOutput() ClusterPkiArrayOutput {
	return o
}

func (o ClusterPkiArrayOutput) ToClusterPkiArrayOutputWithContext(ctx context.Context) ClusterPkiArrayOutput {
	return o
}

func (o ClusterPkiArrayOutput) Index(i pulumi.IntInput) ClusterPkiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterPki {
		return vs[0].([]*ClusterPki)[vs[1].(int)]
	}).(ClusterPkiOutput)
}

type ClusterPkiMapOutput struct{ *pulumi.OutputState }

func (ClusterPkiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterPki)(nil)).Elem()
}

func (o ClusterPkiMapOutput) ToClusterPkiMapOutput() ClusterPkiMapOutput {
	return o
}

func (o ClusterPkiMapOutput) ToClusterPkiMapOutputWithContext(ctx context.Context) ClusterPkiMapOutput {
	return o
}

func (o ClusterPkiMapOutput) MapIndex(k pulumi.StringInput) ClusterPkiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterPki {
		return vs[0].(map[string]*ClusterPki)[vs[1].(string)]
	}).(ClusterPkiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPkiInput)(nil)).Elem(), &ClusterPki{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPkiArrayInput)(nil)).Elem(), ClusterPkiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPkiMapInput)(nil)).Elem(), ClusterPkiMap{})
	pulumi.RegisterOutputType(ClusterPkiOutput{})
	pulumi.RegisterOutputType(ClusterPkiArrayOutput{})
	pulumi.RegisterOutputType(ClusterPkiMapOutput{})
}
