// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Configures Kubernetes API Server on a remote system.
type KubeApiServerConfiguration struct {
	pulumi.ResourceState

	// The PEM encoded certificate authority key.
	CaKey pulumi.StringOutput `pulumi:"caKey"`
	// The PEM encoded certificate authority data.
	CaPem pulumi.StringOutput `pulumi:"caPem"`
	// The PEM encoded Kube API Server certificate data.
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	// The directory to store Kubernetes Control Plane configuration.
	ConfigurationDirectory pulumi.StringPtrOutput `pulumi:"configurationDirectory"`
	// Configuration mkdir operation
	ConfigurationMkdir tools.MkdirPtrOutput `pulumi:"configurationMkdir"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The YAML encryption configuration manifest.
	EncryptionConfig pulumi.StringOutput `pulumi:"encryptionConfig"`
	// The PEM encoded Kube API Server certificate key.
	KeyPem pulumi.StringOutput `pulumi:"keyPem"`
	// The path to the 'kubectl' binary.
	KubectlPath pulumi.StringPtrOutput `pulumi:"kubectlPath"`
	// The path to the 'kube-apiserver' binary.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The PEM encoded Service Accounts certificate key.
	ServiceAccountsKey pulumi.StringOutput `pulumi:"serviceAccountsKey"`
	// The PEM encoded Service Accounts certificate data.
	ServiceAccountsPem pulumi.StringOutput `pulumi:"serviceAccountsPem"`
}

// NewKubeApiServerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewKubeApiServerConfiguration(ctx *pulumi.Context,
	name string, args *KubeApiServerConfigurationArgs, opts ...pulumi.ResourceOption) (*KubeApiServerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaKey == nil {
		return nil, errors.New("invalid value for required argument 'CaKey'")
	}
	if args.CaPem == nil {
		return nil, errors.New("invalid value for required argument 'CaPem'")
	}
	if args.CertPem == nil {
		return nil, errors.New("invalid value for required argument 'CertPem'")
	}
	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.EncryptionConfig == nil {
		return nil, errors.New("invalid value for required argument 'EncryptionConfig'")
	}
	if args.KeyPem == nil {
		return nil, errors.New("invalid value for required argument 'KeyPem'")
	}
	if args.ServiceAccountsKey == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountsKey'")
	}
	if args.ServiceAccountsPem == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountsPem'")
	}
	if args.ConfigurationDirectory == nil {
		args.ConfigurationDirectory = pulumi.StringPtr("/etc/kubernetes/config")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubeApiServerConfiguration
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeApiServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeApiServerConfigurationArgs struct {
	// The PEM encoded certificate authority key.
	CaKey string `pulumi:"caKey"`
	// The PEM encoded certificate authority data.
	CaPem string `pulumi:"caPem"`
	// The PEM encoded Kube API Server certificate data.
	CertPem string `pulumi:"certPem"`
	// The directory to store Kubernetes Control Plane configuration.
	ConfigurationDirectory *string `pulumi:"configurationDirectory"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The YAML encryption configuration manifest.
	EncryptionConfig string `pulumi:"encryptionConfig"`
	// The PEM encoded Kube API Server certificate key.
	KeyPem string `pulumi:"keyPem"`
	// The path to the 'kubectl' binary.
	KubectlPath *string `pulumi:"kubectlPath"`
	// The path to the 'kube-apiserver' binary.
	Path *string `pulumi:"path"`
	// The PEM encoded Service Accounts certificate key.
	ServiceAccountsKey string `pulumi:"serviceAccountsKey"`
	// The PEM encoded Service Accounts certificate data.
	ServiceAccountsPem string `pulumi:"serviceAccountsPem"`
}

// The set of arguments for constructing a KubeApiServerConfiguration resource.
type KubeApiServerConfigurationArgs struct {
	// The PEM encoded certificate authority key.
	CaKey pulumi.StringInput
	// The PEM encoded certificate authority data.
	CaPem pulumi.StringInput
	// The PEM encoded Kube API Server certificate data.
	CertPem pulumi.StringInput
	// The directory to store Kubernetes Control Plane configuration.
	ConfigurationDirectory pulumi.StringPtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The YAML encryption configuration manifest.
	EncryptionConfig pulumi.StringInput
	// The PEM encoded Kube API Server certificate key.
	KeyPem pulumi.StringInput
	// The path to the 'kubectl' binary.
	KubectlPath pulumi.StringPtrInput
	// The path to the 'kube-apiserver' binary.
	Path pulumi.StringPtrInput
	// The PEM encoded Service Accounts certificate key.
	ServiceAccountsKey pulumi.StringInput
	// The PEM encoded Service Accounts certificate data.
	ServiceAccountsPem pulumi.StringInput
}

func (KubeApiServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeApiServerConfigurationArgs)(nil)).Elem()
}

type KubeApiServerConfigurationInput interface {
	pulumi.Input

	ToKubeApiServerConfigurationOutput() KubeApiServerConfigurationOutput
	ToKubeApiServerConfigurationOutputWithContext(ctx context.Context) KubeApiServerConfigurationOutput
}

func (*KubeApiServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeApiServerConfiguration)(nil)).Elem()
}

func (i *KubeApiServerConfiguration) ToKubeApiServerConfigurationOutput() KubeApiServerConfigurationOutput {
	return i.ToKubeApiServerConfigurationOutputWithContext(context.Background())
}

func (i *KubeApiServerConfiguration) ToKubeApiServerConfigurationOutputWithContext(ctx context.Context) KubeApiServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerConfigurationOutput)
}

// KubeApiServerConfigurationArrayInput is an input type that accepts KubeApiServerConfigurationArray and KubeApiServerConfigurationArrayOutput values.
// You can construct a concrete instance of `KubeApiServerConfigurationArrayInput` via:
//
//	KubeApiServerConfigurationArray{ KubeApiServerConfigurationArgs{...} }
type KubeApiServerConfigurationArrayInput interface {
	pulumi.Input

	ToKubeApiServerConfigurationArrayOutput() KubeApiServerConfigurationArrayOutput
	ToKubeApiServerConfigurationArrayOutputWithContext(context.Context) KubeApiServerConfigurationArrayOutput
}

type KubeApiServerConfigurationArray []KubeApiServerConfigurationInput

func (KubeApiServerConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeApiServerConfiguration)(nil)).Elem()
}

func (i KubeApiServerConfigurationArray) ToKubeApiServerConfigurationArrayOutput() KubeApiServerConfigurationArrayOutput {
	return i.ToKubeApiServerConfigurationArrayOutputWithContext(context.Background())
}

func (i KubeApiServerConfigurationArray) ToKubeApiServerConfigurationArrayOutputWithContext(ctx context.Context) KubeApiServerConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerConfigurationArrayOutput)
}

// KubeApiServerConfigurationMapInput is an input type that accepts KubeApiServerConfigurationMap and KubeApiServerConfigurationMapOutput values.
// You can construct a concrete instance of `KubeApiServerConfigurationMapInput` via:
//
//	KubeApiServerConfigurationMap{ "key": KubeApiServerConfigurationArgs{...} }
type KubeApiServerConfigurationMapInput interface {
	pulumi.Input

	ToKubeApiServerConfigurationMapOutput() KubeApiServerConfigurationMapOutput
	ToKubeApiServerConfigurationMapOutputWithContext(context.Context) KubeApiServerConfigurationMapOutput
}

type KubeApiServerConfigurationMap map[string]KubeApiServerConfigurationInput

func (KubeApiServerConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeApiServerConfiguration)(nil)).Elem()
}

func (i KubeApiServerConfigurationMap) ToKubeApiServerConfigurationMapOutput() KubeApiServerConfigurationMapOutput {
	return i.ToKubeApiServerConfigurationMapOutputWithContext(context.Background())
}

func (i KubeApiServerConfigurationMap) ToKubeApiServerConfigurationMapOutputWithContext(ctx context.Context) KubeApiServerConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerConfigurationMapOutput)
}

type KubeApiServerConfigurationOutput struct{ *pulumi.OutputState }

func (KubeApiServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeApiServerConfiguration)(nil)).Elem()
}

func (o KubeApiServerConfigurationOutput) ToKubeApiServerConfigurationOutput() KubeApiServerConfigurationOutput {
	return o
}

func (o KubeApiServerConfigurationOutput) ToKubeApiServerConfigurationOutputWithContext(ctx context.Context) KubeApiServerConfigurationOutput {
	return o
}

// The PEM encoded certificate authority key.
func (o KubeApiServerConfigurationOutput) CaKey() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringOutput { return v.CaKey }).(pulumi.StringOutput)
}

// The PEM encoded certificate authority data.
func (o KubeApiServerConfigurationOutput) CaPem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringOutput { return v.CaPem }).(pulumi.StringOutput)
}

// The PEM encoded Kube API Server certificate data.
func (o KubeApiServerConfigurationOutput) CertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringOutput { return v.CertPem }).(pulumi.StringOutput)
}

// The directory to store Kubernetes Control Plane configuration.
func (o KubeApiServerConfigurationOutput) ConfigurationDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringPtrOutput { return v.ConfigurationDirectory }).(pulumi.StringPtrOutput)
}

// Configuration mkdir operation
func (o KubeApiServerConfigurationOutput) ConfigurationMkdir() tools.MkdirPtrOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) tools.MkdirPtrOutput { return v.ConfigurationMkdir }).(tools.MkdirPtrOutput)
}

// The parameters with which to connect to the remote host.
func (o KubeApiServerConfigurationOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The YAML encryption configuration manifest.
func (o KubeApiServerConfigurationOutput) EncryptionConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringOutput { return v.EncryptionConfig }).(pulumi.StringOutput)
}

// The PEM encoded Kube API Server certificate key.
func (o KubeApiServerConfigurationOutput) KeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringOutput { return v.KeyPem }).(pulumi.StringOutput)
}

// The path to the 'kubectl' binary.
func (o KubeApiServerConfigurationOutput) KubectlPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringPtrOutput { return v.KubectlPath }).(pulumi.StringPtrOutput)
}

// The path to the 'kube-apiserver' binary.
func (o KubeApiServerConfigurationOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The PEM encoded Service Accounts certificate key.
func (o KubeApiServerConfigurationOutput) ServiceAccountsKey() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringOutput { return v.ServiceAccountsKey }).(pulumi.StringOutput)
}

// The PEM encoded Service Accounts certificate data.
func (o KubeApiServerConfigurationOutput) ServiceAccountsPem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerConfiguration) pulumi.StringOutput { return v.ServiceAccountsPem }).(pulumi.StringOutput)
}

type KubeApiServerConfigurationArrayOutput struct{ *pulumi.OutputState }

func (KubeApiServerConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeApiServerConfiguration)(nil)).Elem()
}

func (o KubeApiServerConfigurationArrayOutput) ToKubeApiServerConfigurationArrayOutput() KubeApiServerConfigurationArrayOutput {
	return o
}

func (o KubeApiServerConfigurationArrayOutput) ToKubeApiServerConfigurationArrayOutputWithContext(ctx context.Context) KubeApiServerConfigurationArrayOutput {
	return o
}

func (o KubeApiServerConfigurationArrayOutput) Index(i pulumi.IntInput) KubeApiServerConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeApiServerConfiguration {
		return vs[0].([]*KubeApiServerConfiguration)[vs[1].(int)]
	}).(KubeApiServerConfigurationOutput)
}

type KubeApiServerConfigurationMapOutput struct{ *pulumi.OutputState }

func (KubeApiServerConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeApiServerConfiguration)(nil)).Elem()
}

func (o KubeApiServerConfigurationMapOutput) ToKubeApiServerConfigurationMapOutput() KubeApiServerConfigurationMapOutput {
	return o
}

func (o KubeApiServerConfigurationMapOutput) ToKubeApiServerConfigurationMapOutputWithContext(ctx context.Context) KubeApiServerConfigurationMapOutput {
	return o
}

func (o KubeApiServerConfigurationMapOutput) MapIndex(k pulumi.StringInput) KubeApiServerConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeApiServerConfiguration {
		return vs[0].(map[string]*KubeApiServerConfiguration)[vs[1].(string)]
	}).(KubeApiServerConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerConfigurationInput)(nil)).Elem(), &KubeApiServerConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerConfigurationArrayInput)(nil)).Elem(), KubeApiServerConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerConfigurationMapInput)(nil)).Elem(), KubeApiServerConfigurationMap{})
	pulumi.RegisterOutputType(KubeApiServerConfigurationOutput{})
	pulumi.RegisterOutputType(KubeApiServerConfigurationArrayOutput{})
	pulumi.RegisterOutputType(KubeApiServerConfigurationMapOutput{})
}