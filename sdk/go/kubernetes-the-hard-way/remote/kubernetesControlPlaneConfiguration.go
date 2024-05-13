// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/config"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Configures Kubernetes Control Plane on a remote system.
type KubernetesControlPlaneConfiguration struct {
	pulumi.ResourceState

	// The PEM encoded certificate authority key.
	CaKey pulumi.StringOutput `pulumi:"caKey"`
	// The PEM encoded certificate authority data.
	CaPem pulumi.StringOutput `pulumi:"caPem"`
	// The directory to store Kubernetes Control Plane configuration.
	ConfigurationDirectory pulumi.StringPtrOutput `pulumi:"configurationDirectory"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The YAML encryption configuration manifest.
	EncryptionConfig pulumi.StringOutput `pulumi:"encryptionConfig"`
	// The PEM encoded Kube API Server certificate key.
	KubeApiServerKey pulumi.StringOutput `pulumi:"kubeApiServerKey"`
	// The path to the 'kube-apiserver' binary.
	KubeApiServerPath pulumi.StringPtrOutput `pulumi:"kubeApiServerPath"`
	// The PEM encoded Kube API Server certificate data.
	KubeApiServerPem pulumi.StringOutput `pulumi:"kubeApiServerPem"`
	// The kube-controller-manager kubeconfig configuration
	KubeControllerManagerKubeconfig config.KubeconfigOutput `pulumi:"kubeControllerManagerKubeconfig"`
	// The path to the 'kube-controller-manager' binary.
	KubeControllerManagerPath pulumi.StringPtrOutput `pulumi:"kubeControllerManagerPath"`
	// The kube-scheduler configuration manifest.
	KubeSchedulerConfig pulumi.StringOutput `pulumi:"kubeSchedulerConfig"`
	// The kube-scheduler kubeconfig configuration
	KubeSchedulerKubeconfig config.KubeconfigOutput `pulumi:"kubeSchedulerKubeconfig"`
	// The path to the 'kube-scheduler' binary.
	KubeSchedulerPath pulumi.StringPtrOutput `pulumi:"kubeSchedulerPath"`
	// The path to the 'kubectl' binary.
	KubectlPath pulumi.StringPtrOutput `pulumi:"kubectlPath"`
	// The PEM encoded Service Accounts certificate key.
	ServiceAccountsKey pulumi.StringOutput `pulumi:"serviceAccountsKey"`
	// The PEM encoded Service Accounts certificate data.
	ServiceAccountsPem pulumi.StringOutput `pulumi:"serviceAccountsPem"`
}

// NewKubernetesControlPlaneConfiguration registers a new resource with the given unique name, arguments, and options.
func NewKubernetesControlPlaneConfiguration(ctx *pulumi.Context,
	name string, args *KubernetesControlPlaneConfigurationArgs, opts ...pulumi.ResourceOption) (*KubernetesControlPlaneConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaKey == nil {
		return nil, errors.New("invalid value for required argument 'CaKey'")
	}
	if args.CaPem == nil {
		return nil, errors.New("invalid value for required argument 'CaPem'")
	}
	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.EncryptionConfig == nil {
		return nil, errors.New("invalid value for required argument 'EncryptionConfig'")
	}
	if args.KubeApiServerKey == nil {
		return nil, errors.New("invalid value for required argument 'KubeApiServerKey'")
	}
	if args.KubeApiServerPem == nil {
		return nil, errors.New("invalid value for required argument 'KubeApiServerPem'")
	}
	if args.KubeControllerManagerKubeconfig == nil {
		return nil, errors.New("invalid value for required argument 'KubeControllerManagerKubeconfig'")
	}
	if args.KubeSchedulerConfig == nil {
		return nil, errors.New("invalid value for required argument 'KubeSchedulerConfig'")
	}
	if args.KubeSchedulerKubeconfig == nil {
		return nil, errors.New("invalid value for required argument 'KubeSchedulerKubeconfig'")
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
	var resource KubernetesControlPlaneConfiguration
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubernetesControlPlaneConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubernetesControlPlaneConfigurationArgs struct {
	// The PEM encoded certificate authority key.
	CaKey string `pulumi:"caKey"`
	// The PEM encoded certificate authority data.
	CaPem string `pulumi:"caPem"`
	// The directory to store Kubernetes Control Plane configuration.
	ConfigurationDirectory *string `pulumi:"configurationDirectory"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The YAML encryption configuration manifest.
	EncryptionConfig string `pulumi:"encryptionConfig"`
	// The PEM encoded Kube API Server certificate key.
	KubeApiServerKey string `pulumi:"kubeApiServerKey"`
	// The path to the 'kube-apiserver' binary.
	KubeApiServerPath *string `pulumi:"kubeApiServerPath"`
	// The PEM encoded Kube API Server certificate data.
	KubeApiServerPem string `pulumi:"kubeApiServerPem"`
	// The kube-controller-manager kubeconfig configuration
	KubeControllerManagerKubeconfig config.Kubeconfig `pulumi:"kubeControllerManagerKubeconfig"`
	// The path to the 'kube-controller-manager' binary.
	KubeControllerManagerPath *string `pulumi:"kubeControllerManagerPath"`
	// The kube-scheduler configuration manifest.
	KubeSchedulerConfig string `pulumi:"kubeSchedulerConfig"`
	// The kube-scheduler kubeconfig configuration
	KubeSchedulerKubeconfig config.Kubeconfig `pulumi:"kubeSchedulerKubeconfig"`
	// The path to the 'kube-scheduler' binary.
	KubeSchedulerPath *string `pulumi:"kubeSchedulerPath"`
	// The path to the 'kubectl' binary.
	KubectlPath *string `pulumi:"kubectlPath"`
	// The PEM encoded Service Accounts certificate key.
	ServiceAccountsKey string `pulumi:"serviceAccountsKey"`
	// The PEM encoded Service Accounts certificate data.
	ServiceAccountsPem string `pulumi:"serviceAccountsPem"`
}

// The set of arguments for constructing a KubernetesControlPlaneConfiguration resource.
type KubernetesControlPlaneConfigurationArgs struct {
	// The PEM encoded certificate authority key.
	CaKey pulumi.StringInput
	// The PEM encoded certificate authority data.
	CaPem pulumi.StringInput
	// The directory to store Kubernetes Control Plane configuration.
	ConfigurationDirectory pulumi.StringPtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The YAML encryption configuration manifest.
	EncryptionConfig pulumi.StringInput
	// The PEM encoded Kube API Server certificate key.
	KubeApiServerKey pulumi.StringInput
	// The path to the 'kube-apiserver' binary.
	KubeApiServerPath pulumi.StringPtrInput
	// The PEM encoded Kube API Server certificate data.
	KubeApiServerPem pulumi.StringInput
	// The kube-controller-manager kubeconfig configuration
	KubeControllerManagerKubeconfig config.KubeconfigInput
	// The path to the 'kube-controller-manager' binary.
	KubeControllerManagerPath pulumi.StringPtrInput
	// The kube-scheduler configuration manifest.
	KubeSchedulerConfig pulumi.StringInput
	// The kube-scheduler kubeconfig configuration
	KubeSchedulerKubeconfig config.KubeconfigInput
	// The path to the 'kube-scheduler' binary.
	KubeSchedulerPath pulumi.StringPtrInput
	// The path to the 'kubectl' binary.
	KubectlPath pulumi.StringPtrInput
	// The PEM encoded Service Accounts certificate key.
	ServiceAccountsKey pulumi.StringInput
	// The PEM encoded Service Accounts certificate data.
	ServiceAccountsPem pulumi.StringInput
}

func (KubernetesControlPlaneConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesControlPlaneConfigurationArgs)(nil)).Elem()
}

type KubernetesControlPlaneConfigurationInput interface {
	pulumi.Input

	ToKubernetesControlPlaneConfigurationOutput() KubernetesControlPlaneConfigurationOutput
	ToKubernetesControlPlaneConfigurationOutputWithContext(ctx context.Context) KubernetesControlPlaneConfigurationOutput
}

func (*KubernetesControlPlaneConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesControlPlaneConfiguration)(nil)).Elem()
}

func (i *KubernetesControlPlaneConfiguration) ToKubernetesControlPlaneConfigurationOutput() KubernetesControlPlaneConfigurationOutput {
	return i.ToKubernetesControlPlaneConfigurationOutputWithContext(context.Background())
}

func (i *KubernetesControlPlaneConfiguration) ToKubernetesControlPlaneConfigurationOutputWithContext(ctx context.Context) KubernetesControlPlaneConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesControlPlaneConfigurationOutput)
}

// KubernetesControlPlaneConfigurationArrayInput is an input type that accepts KubernetesControlPlaneConfigurationArray and KubernetesControlPlaneConfigurationArrayOutput values.
// You can construct a concrete instance of `KubernetesControlPlaneConfigurationArrayInput` via:
//
//	KubernetesControlPlaneConfigurationArray{ KubernetesControlPlaneConfigurationArgs{...} }
type KubernetesControlPlaneConfigurationArrayInput interface {
	pulumi.Input

	ToKubernetesControlPlaneConfigurationArrayOutput() KubernetesControlPlaneConfigurationArrayOutput
	ToKubernetesControlPlaneConfigurationArrayOutputWithContext(context.Context) KubernetesControlPlaneConfigurationArrayOutput
}

type KubernetesControlPlaneConfigurationArray []KubernetesControlPlaneConfigurationInput

func (KubernetesControlPlaneConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesControlPlaneConfiguration)(nil)).Elem()
}

func (i KubernetesControlPlaneConfigurationArray) ToKubernetesControlPlaneConfigurationArrayOutput() KubernetesControlPlaneConfigurationArrayOutput {
	return i.ToKubernetesControlPlaneConfigurationArrayOutputWithContext(context.Background())
}

func (i KubernetesControlPlaneConfigurationArray) ToKubernetesControlPlaneConfigurationArrayOutputWithContext(ctx context.Context) KubernetesControlPlaneConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesControlPlaneConfigurationArrayOutput)
}

// KubernetesControlPlaneConfigurationMapInput is an input type that accepts KubernetesControlPlaneConfigurationMap and KubernetesControlPlaneConfigurationMapOutput values.
// You can construct a concrete instance of `KubernetesControlPlaneConfigurationMapInput` via:
//
//	KubernetesControlPlaneConfigurationMap{ "key": KubernetesControlPlaneConfigurationArgs{...} }
type KubernetesControlPlaneConfigurationMapInput interface {
	pulumi.Input

	ToKubernetesControlPlaneConfigurationMapOutput() KubernetesControlPlaneConfigurationMapOutput
	ToKubernetesControlPlaneConfigurationMapOutputWithContext(context.Context) KubernetesControlPlaneConfigurationMapOutput
}

type KubernetesControlPlaneConfigurationMap map[string]KubernetesControlPlaneConfigurationInput

func (KubernetesControlPlaneConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesControlPlaneConfiguration)(nil)).Elem()
}

func (i KubernetesControlPlaneConfigurationMap) ToKubernetesControlPlaneConfigurationMapOutput() KubernetesControlPlaneConfigurationMapOutput {
	return i.ToKubernetesControlPlaneConfigurationMapOutputWithContext(context.Background())
}

func (i KubernetesControlPlaneConfigurationMap) ToKubernetesControlPlaneConfigurationMapOutputWithContext(ctx context.Context) KubernetesControlPlaneConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesControlPlaneConfigurationMapOutput)
}

type KubernetesControlPlaneConfigurationOutput struct{ *pulumi.OutputState }

func (KubernetesControlPlaneConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesControlPlaneConfiguration)(nil)).Elem()
}

func (o KubernetesControlPlaneConfigurationOutput) ToKubernetesControlPlaneConfigurationOutput() KubernetesControlPlaneConfigurationOutput {
	return o
}

func (o KubernetesControlPlaneConfigurationOutput) ToKubernetesControlPlaneConfigurationOutputWithContext(ctx context.Context) KubernetesControlPlaneConfigurationOutput {
	return o
}

// The PEM encoded certificate authority key.
func (o KubernetesControlPlaneConfigurationOutput) CaKey() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.CaKey }).(pulumi.StringOutput)
}

// The PEM encoded certificate authority data.
func (o KubernetesControlPlaneConfigurationOutput) CaPem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.CaPem }).(pulumi.StringOutput)
}

// The directory to store Kubernetes Control Plane configuration.
func (o KubernetesControlPlaneConfigurationOutput) ConfigurationDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringPtrOutput { return v.ConfigurationDirectory }).(pulumi.StringPtrOutput)
}

// The parameters with which to connect to the remote host.
func (o KubernetesControlPlaneConfigurationOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The YAML encryption configuration manifest.
func (o KubernetesControlPlaneConfigurationOutput) EncryptionConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.EncryptionConfig }).(pulumi.StringOutput)
}

// The PEM encoded Kube API Server certificate key.
func (o KubernetesControlPlaneConfigurationOutput) KubeApiServerKey() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.KubeApiServerKey }).(pulumi.StringOutput)
}

// The path to the 'kube-apiserver' binary.
func (o KubernetesControlPlaneConfigurationOutput) KubeApiServerPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringPtrOutput { return v.KubeApiServerPath }).(pulumi.StringPtrOutput)
}

// The PEM encoded Kube API Server certificate data.
func (o KubernetesControlPlaneConfigurationOutput) KubeApiServerPem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.KubeApiServerPem }).(pulumi.StringOutput)
}

// The kube-controller-manager kubeconfig configuration
func (o KubernetesControlPlaneConfigurationOutput) KubeControllerManagerKubeconfig() config.KubeconfigOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) config.KubeconfigOutput {
		return v.KubeControllerManagerKubeconfig
	}).(config.KubeconfigOutput)
}

// The path to the 'kube-controller-manager' binary.
func (o KubernetesControlPlaneConfigurationOutput) KubeControllerManagerPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringPtrOutput {
		return v.KubeControllerManagerPath
	}).(pulumi.StringPtrOutput)
}

// The kube-scheduler configuration manifest.
func (o KubernetesControlPlaneConfigurationOutput) KubeSchedulerConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.KubeSchedulerConfig }).(pulumi.StringOutput)
}

// The kube-scheduler kubeconfig configuration
func (o KubernetesControlPlaneConfigurationOutput) KubeSchedulerKubeconfig() config.KubeconfigOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) config.KubeconfigOutput { return v.KubeSchedulerKubeconfig }).(config.KubeconfigOutput)
}

// The path to the 'kube-scheduler' binary.
func (o KubernetesControlPlaneConfigurationOutput) KubeSchedulerPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringPtrOutput { return v.KubeSchedulerPath }).(pulumi.StringPtrOutput)
}

// The path to the 'kubectl' binary.
func (o KubernetesControlPlaneConfigurationOutput) KubectlPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringPtrOutput { return v.KubectlPath }).(pulumi.StringPtrOutput)
}

// The PEM encoded Service Accounts certificate key.
func (o KubernetesControlPlaneConfigurationOutput) ServiceAccountsKey() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.ServiceAccountsKey }).(pulumi.StringOutput)
}

// The PEM encoded Service Accounts certificate data.
func (o KubernetesControlPlaneConfigurationOutput) ServiceAccountsPem() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesControlPlaneConfiguration) pulumi.StringOutput { return v.ServiceAccountsPem }).(pulumi.StringOutput)
}

type KubernetesControlPlaneConfigurationArrayOutput struct{ *pulumi.OutputState }

func (KubernetesControlPlaneConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesControlPlaneConfiguration)(nil)).Elem()
}

func (o KubernetesControlPlaneConfigurationArrayOutput) ToKubernetesControlPlaneConfigurationArrayOutput() KubernetesControlPlaneConfigurationArrayOutput {
	return o
}

func (o KubernetesControlPlaneConfigurationArrayOutput) ToKubernetesControlPlaneConfigurationArrayOutputWithContext(ctx context.Context) KubernetesControlPlaneConfigurationArrayOutput {
	return o
}

func (o KubernetesControlPlaneConfigurationArrayOutput) Index(i pulumi.IntInput) KubernetesControlPlaneConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubernetesControlPlaneConfiguration {
		return vs[0].([]*KubernetesControlPlaneConfiguration)[vs[1].(int)]
	}).(KubernetesControlPlaneConfigurationOutput)
}

type KubernetesControlPlaneConfigurationMapOutput struct{ *pulumi.OutputState }

func (KubernetesControlPlaneConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesControlPlaneConfiguration)(nil)).Elem()
}

func (o KubernetesControlPlaneConfigurationMapOutput) ToKubernetesControlPlaneConfigurationMapOutput() KubernetesControlPlaneConfigurationMapOutput {
	return o
}

func (o KubernetesControlPlaneConfigurationMapOutput) ToKubernetesControlPlaneConfigurationMapOutputWithContext(ctx context.Context) KubernetesControlPlaneConfigurationMapOutput {
	return o
}

func (o KubernetesControlPlaneConfigurationMapOutput) MapIndex(k pulumi.StringInput) KubernetesControlPlaneConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubernetesControlPlaneConfiguration {
		return vs[0].(map[string]*KubernetesControlPlaneConfiguration)[vs[1].(string)]
	}).(KubernetesControlPlaneConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesControlPlaneConfigurationInput)(nil)).Elem(), &KubernetesControlPlaneConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesControlPlaneConfigurationArrayInput)(nil)).Elem(), KubernetesControlPlaneConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesControlPlaneConfigurationMapInput)(nil)).Elem(), KubernetesControlPlaneConfigurationMap{})
	pulumi.RegisterOutputType(KubernetesControlPlaneConfigurationOutput{})
	pulumi.RegisterOutputType(KubernetesControlPlaneConfigurationArrayOutput{})
	pulumi.RegisterOutputType(KubernetesControlPlaneConfigurationMapOutput{})
}
