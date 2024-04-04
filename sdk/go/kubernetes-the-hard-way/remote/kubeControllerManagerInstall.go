// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Installs kube-controller-manager on a remote system.
type KubeControllerManagerInstall struct {
	pulumi.ResourceState

	// The kube-controller-manager CPU architecture.
	Architecture ArchitectureOutput `pulumi:"architecture"`
	// The command resource.
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// The connection details.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Directory to install the `kube-controller-manager` binary.
	InstallDirectory pulumi.StringOutput `pulumi:"installDirectory"`
	// The version of kube-controller-manager to install.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewKubeControllerManagerInstall registers a new resource with the given unique name, arguments, and options.
func NewKubeControllerManagerInstall(ctx *pulumi.Context,
	name string, args *KubeControllerManagerInstallArgs, opts ...pulumi.ResourceOption) (*KubeControllerManagerInstall, error) {
	if args == nil {
		args = &KubeControllerManagerInstallArgs{}
	}

	if args.Connection != nil {
		args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v *pulumiCommand.Connection) *pulumiCommand.Connection { return v.Defaults() }).(*pulumiCommand.ConnectionOutput)
	}
	if args.InstallDirectory == nil {
		args.InstallDirectory = pulumi.StringPtr("/usr/local/bin")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubeControllerManagerInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeControllerManagerInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeControllerManagerInstallArgs struct {
	// The kube-controller-manager CPU architecture.
	Architecture *Architecture `pulumi:"architecture"`
	// The connection details.
	Connection *pulumiCommand.Connection `pulumi:"connection"`
	// Directory to install the `kube-controller-manager` binary.
	InstallDirectory *string `pulumi:"installDirectory"`
	// The version of kube-controller-manager to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a KubeControllerManagerInstall resource.
type KubeControllerManagerInstallArgs struct {
	// The kube-controller-manager CPU architecture.
	Architecture ArchitecturePtrInput
	// The connection details.
	Connection *pulumiCommand.ConnectionInput
	// Directory to install the `kube-controller-manager` binary.
	InstallDirectory pulumi.StringPtrInput
	// The version of kube-controller-manager to install.
	Version pulumi.StringPtrInput
}

func (KubeControllerManagerInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeControllerManagerInstallArgs)(nil)).Elem()
}

type KubeControllerManagerInstallInput interface {
	pulumi.Input

	ToKubeControllerManagerInstallOutput() KubeControllerManagerInstallOutput
	ToKubeControllerManagerInstallOutputWithContext(ctx context.Context) KubeControllerManagerInstallOutput
}

func (*KubeControllerManagerInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeControllerManagerInstall)(nil)).Elem()
}

func (i *KubeControllerManagerInstall) ToKubeControllerManagerInstallOutput() KubeControllerManagerInstallOutput {
	return i.ToKubeControllerManagerInstallOutputWithContext(context.Background())
}

func (i *KubeControllerManagerInstall) ToKubeControllerManagerInstallOutputWithContext(ctx context.Context) KubeControllerManagerInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeControllerManagerInstallOutput)
}

// KubeControllerManagerInstallArrayInput is an input type that accepts KubeControllerManagerInstallArray and KubeControllerManagerInstallArrayOutput values.
// You can construct a concrete instance of `KubeControllerManagerInstallArrayInput` via:
//
//	KubeControllerManagerInstallArray{ KubeControllerManagerInstallArgs{...} }
type KubeControllerManagerInstallArrayInput interface {
	pulumi.Input

	ToKubeControllerManagerInstallArrayOutput() KubeControllerManagerInstallArrayOutput
	ToKubeControllerManagerInstallArrayOutputWithContext(context.Context) KubeControllerManagerInstallArrayOutput
}

type KubeControllerManagerInstallArray []KubeControllerManagerInstallInput

func (KubeControllerManagerInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeControllerManagerInstall)(nil)).Elem()
}

func (i KubeControllerManagerInstallArray) ToKubeControllerManagerInstallArrayOutput() KubeControllerManagerInstallArrayOutput {
	return i.ToKubeControllerManagerInstallArrayOutputWithContext(context.Background())
}

func (i KubeControllerManagerInstallArray) ToKubeControllerManagerInstallArrayOutputWithContext(ctx context.Context) KubeControllerManagerInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeControllerManagerInstallArrayOutput)
}

// KubeControllerManagerInstallMapInput is an input type that accepts KubeControllerManagerInstallMap and KubeControllerManagerInstallMapOutput values.
// You can construct a concrete instance of `KubeControllerManagerInstallMapInput` via:
//
//	KubeControllerManagerInstallMap{ "key": KubeControllerManagerInstallArgs{...} }
type KubeControllerManagerInstallMapInput interface {
	pulumi.Input

	ToKubeControllerManagerInstallMapOutput() KubeControllerManagerInstallMapOutput
	ToKubeControllerManagerInstallMapOutputWithContext(context.Context) KubeControllerManagerInstallMapOutput
}

type KubeControllerManagerInstallMap map[string]KubeControllerManagerInstallInput

func (KubeControllerManagerInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeControllerManagerInstall)(nil)).Elem()
}

func (i KubeControllerManagerInstallMap) ToKubeControllerManagerInstallMapOutput() KubeControllerManagerInstallMapOutput {
	return i.ToKubeControllerManagerInstallMapOutputWithContext(context.Background())
}

func (i KubeControllerManagerInstallMap) ToKubeControllerManagerInstallMapOutputWithContext(ctx context.Context) KubeControllerManagerInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeControllerManagerInstallMapOutput)
}

type KubeControllerManagerInstallOutput struct{ *pulumi.OutputState }

func (KubeControllerManagerInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeControllerManagerInstall)(nil)).Elem()
}

func (o KubeControllerManagerInstallOutput) ToKubeControllerManagerInstallOutput() KubeControllerManagerInstallOutput {
	return o
}

func (o KubeControllerManagerInstallOutput) ToKubeControllerManagerInstallOutputWithContext(ctx context.Context) KubeControllerManagerInstallOutput {
	return o
}

// The kube-controller-manager CPU architecture.
func (o KubeControllerManagerInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *KubeControllerManagerInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

// The command resource.
func (o KubeControllerManagerInstallOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *KubeControllerManagerInstall) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// The connection details.
func (o KubeControllerManagerInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeControllerManagerInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Directory to install the `kube-controller-manager` binary.
func (o KubeControllerManagerInstallOutput) InstallDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeControllerManagerInstall) pulumi.StringOutput { return v.InstallDirectory }).(pulumi.StringOutput)
}

// The version of kube-controller-manager to install.
func (o KubeControllerManagerInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeControllerManagerInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubeControllerManagerInstallArrayOutput struct{ *pulumi.OutputState }

func (KubeControllerManagerInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeControllerManagerInstall)(nil)).Elem()
}

func (o KubeControllerManagerInstallArrayOutput) ToKubeControllerManagerInstallArrayOutput() KubeControllerManagerInstallArrayOutput {
	return o
}

func (o KubeControllerManagerInstallArrayOutput) ToKubeControllerManagerInstallArrayOutputWithContext(ctx context.Context) KubeControllerManagerInstallArrayOutput {
	return o
}

func (o KubeControllerManagerInstallArrayOutput) Index(i pulumi.IntInput) KubeControllerManagerInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeControllerManagerInstall {
		return vs[0].([]*KubeControllerManagerInstall)[vs[1].(int)]
	}).(KubeControllerManagerInstallOutput)
}

type KubeControllerManagerInstallMapOutput struct{ *pulumi.OutputState }

func (KubeControllerManagerInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeControllerManagerInstall)(nil)).Elem()
}

func (o KubeControllerManagerInstallMapOutput) ToKubeControllerManagerInstallMapOutput() KubeControllerManagerInstallMapOutput {
	return o
}

func (o KubeControllerManagerInstallMapOutput) ToKubeControllerManagerInstallMapOutputWithContext(ctx context.Context) KubeControllerManagerInstallMapOutput {
	return o
}

func (o KubeControllerManagerInstallMapOutput) MapIndex(k pulumi.StringInput) KubeControllerManagerInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeControllerManagerInstall {
		return vs[0].(map[string]*KubeControllerManagerInstall)[vs[1].(string)]
	}).(KubeControllerManagerInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeControllerManagerInstallInput)(nil)).Elem(), &KubeControllerManagerInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeControllerManagerInstallArrayInput)(nil)).Elem(), KubeControllerManagerInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeControllerManagerInstallMapInput)(nil)).Elem(), KubeControllerManagerInstallMap{})
	pulumi.RegisterOutputType(KubeControllerManagerInstallOutput{})
	pulumi.RegisterOutputType(KubeControllerManagerInstallArrayOutput{})
	pulumi.RegisterOutputType(KubeControllerManagerInstallMapOutput{})
}