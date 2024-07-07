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
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/tools"
)

// Installs kubeadm on a remote system.
type KubeadmInstall struct {
	pulumi.ResourceState

	// The CPU architecture to install.
	Architecture ArchitectureOutput `pulumi:"architecture"`
	// The name of the installed binary.
	BinName pulumi.StringPtrOutput `pulumi:"binName"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The directory to install the binary to.
	Directory pulumi.StringOutput `pulumi:"directory"`
	// The download operation.
	Download DownloadOutput `pulumi:"download"`
	// The mkdir operation.
	Mkdir tools.MkdirOutput `pulumi:"mkdir"`
	// The mktemp operation.
	Mktemp tools.MktempOutput `pulumi:"mktemp"`
	// The mv operation.
	Mv tools.MvOutput `pulumi:"mv"`
	// The path to the installed binary.
	Path pulumi.StringOutput `pulumi:"path"`
	// The rm operation.
	Rm tools.RmOutput `pulumi:"rm"`
	// The url used to download the binary.
	Url pulumi.StringOutput `pulumi:"url"`
	// The version to install.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewKubeadmInstall registers a new resource with the given unique name, arguments, and options.
func NewKubeadmInstall(ctx *pulumi.Context,
	name string, args *KubeadmInstallArgs, opts ...pulumi.ResourceOption) (*KubeadmInstall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	if args.Directory == nil {
		args.Directory = pulumi.StringPtr("/usr/local/bin")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubeadmInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeadmInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeadmInstallArgs struct {
	// The CPU architecture to install.
	Architecture *Architecture `pulumi:"architecture"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The directory to install the binary to.
	Directory *string `pulumi:"directory"`
	// The version to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a KubeadmInstall resource.
type KubeadmInstallArgs struct {
	// The CPU architecture to install.
	Architecture ArchitecturePtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The directory to install the binary to.
	Directory pulumi.StringPtrInput
	// The version to install.
	Version pulumi.StringPtrInput
}

func (KubeadmInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeadmInstallArgs)(nil)).Elem()
}

type KubeadmInstallInput interface {
	pulumi.Input

	ToKubeadmInstallOutput() KubeadmInstallOutput
	ToKubeadmInstallOutputWithContext(ctx context.Context) KubeadmInstallOutput
}

func (*KubeadmInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeadmInstall)(nil)).Elem()
}

func (i *KubeadmInstall) ToKubeadmInstallOutput() KubeadmInstallOutput {
	return i.ToKubeadmInstallOutputWithContext(context.Background())
}

func (i *KubeadmInstall) ToKubeadmInstallOutputWithContext(ctx context.Context) KubeadmInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeadmInstallOutput)
}

// KubeadmInstallArrayInput is an input type that accepts KubeadmInstallArray and KubeadmInstallArrayOutput values.
// You can construct a concrete instance of `KubeadmInstallArrayInput` via:
//
//	KubeadmInstallArray{ KubeadmInstallArgs{...} }
type KubeadmInstallArrayInput interface {
	pulumi.Input

	ToKubeadmInstallArrayOutput() KubeadmInstallArrayOutput
	ToKubeadmInstallArrayOutputWithContext(context.Context) KubeadmInstallArrayOutput
}

type KubeadmInstallArray []KubeadmInstallInput

func (KubeadmInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeadmInstall)(nil)).Elem()
}

func (i KubeadmInstallArray) ToKubeadmInstallArrayOutput() KubeadmInstallArrayOutput {
	return i.ToKubeadmInstallArrayOutputWithContext(context.Background())
}

func (i KubeadmInstallArray) ToKubeadmInstallArrayOutputWithContext(ctx context.Context) KubeadmInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeadmInstallArrayOutput)
}

// KubeadmInstallMapInput is an input type that accepts KubeadmInstallMap and KubeadmInstallMapOutput values.
// You can construct a concrete instance of `KubeadmInstallMapInput` via:
//
//	KubeadmInstallMap{ "key": KubeadmInstallArgs{...} }
type KubeadmInstallMapInput interface {
	pulumi.Input

	ToKubeadmInstallMapOutput() KubeadmInstallMapOutput
	ToKubeadmInstallMapOutputWithContext(context.Context) KubeadmInstallMapOutput
}

type KubeadmInstallMap map[string]KubeadmInstallInput

func (KubeadmInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeadmInstall)(nil)).Elem()
}

func (i KubeadmInstallMap) ToKubeadmInstallMapOutput() KubeadmInstallMapOutput {
	return i.ToKubeadmInstallMapOutputWithContext(context.Background())
}

func (i KubeadmInstallMap) ToKubeadmInstallMapOutputWithContext(ctx context.Context) KubeadmInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeadmInstallMapOutput)
}

type KubeadmInstallOutput struct{ *pulumi.OutputState }

func (KubeadmInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeadmInstall)(nil)).Elem()
}

func (o KubeadmInstallOutput) ToKubeadmInstallOutput() KubeadmInstallOutput {
	return o
}

func (o KubeadmInstallOutput) ToKubeadmInstallOutputWithContext(ctx context.Context) KubeadmInstallOutput {
	return o
}

// The CPU architecture to install.
func (o KubeadmInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *KubeadmInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

// The name of the installed binary.
func (o KubeadmInstallOutput) BinName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeadmInstall) pulumi.StringPtrOutput { return v.BinName }).(pulumi.StringPtrOutput)
}

// The parameters with which to connect to the remote host.
func (o KubeadmInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeadmInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The directory to install the binary to.
func (o KubeadmInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeadmInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

// The download operation.
func (o KubeadmInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *KubeadmInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

// The mkdir operation.
func (o KubeadmInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *KubeadmInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

// The mktemp operation.
func (o KubeadmInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *KubeadmInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

// The mv operation.
func (o KubeadmInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *KubeadmInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

// The path to the installed binary.
func (o KubeadmInstallOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeadmInstall) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The rm operation.
func (o KubeadmInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *KubeadmInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

// The url used to download the binary.
func (o KubeadmInstallOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeadmInstall) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The version to install.
func (o KubeadmInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeadmInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubeadmInstallArrayOutput struct{ *pulumi.OutputState }

func (KubeadmInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeadmInstall)(nil)).Elem()
}

func (o KubeadmInstallArrayOutput) ToKubeadmInstallArrayOutput() KubeadmInstallArrayOutput {
	return o
}

func (o KubeadmInstallArrayOutput) ToKubeadmInstallArrayOutputWithContext(ctx context.Context) KubeadmInstallArrayOutput {
	return o
}

func (o KubeadmInstallArrayOutput) Index(i pulumi.IntInput) KubeadmInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeadmInstall {
		return vs[0].([]*KubeadmInstall)[vs[1].(int)]
	}).(KubeadmInstallOutput)
}

type KubeadmInstallMapOutput struct{ *pulumi.OutputState }

func (KubeadmInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeadmInstall)(nil)).Elem()
}

func (o KubeadmInstallMapOutput) ToKubeadmInstallMapOutput() KubeadmInstallMapOutput {
	return o
}

func (o KubeadmInstallMapOutput) ToKubeadmInstallMapOutputWithContext(ctx context.Context) KubeadmInstallMapOutput {
	return o
}

func (o KubeadmInstallMapOutput) MapIndex(k pulumi.StringInput) KubeadmInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeadmInstall {
		return vs[0].(map[string]*KubeadmInstall)[vs[1].(string)]
	}).(KubeadmInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeadmInstallInput)(nil)).Elem(), &KubeadmInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeadmInstallArrayInput)(nil)).Elem(), KubeadmInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeadmInstallMapInput)(nil)).Elem(), KubeadmInstallMap{})
	pulumi.RegisterOutputType(KubeadmInstallOutput{})
	pulumi.RegisterOutputType(KubeadmInstallArrayOutput{})
	pulumi.RegisterOutputType(KubeadmInstallMapOutput{})
}
