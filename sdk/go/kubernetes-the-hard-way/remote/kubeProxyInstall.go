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

// Installs kube-proxy on a remote system.
type KubeProxyInstall struct {
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

// NewKubeProxyInstall registers a new resource with the given unique name, arguments, and options.
func NewKubeProxyInstall(ctx *pulumi.Context,
	name string, args *KubeProxyInstallArgs, opts ...pulumi.ResourceOption) (*KubeProxyInstall, error) {
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
	var resource KubeProxyInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeProxyInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeProxyInstallArgs struct {
	// The CPU architecture to install.
	Architecture *Architecture `pulumi:"architecture"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The directory to install the binary to.
	Directory *string `pulumi:"directory"`
	// The version to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a KubeProxyInstall resource.
type KubeProxyInstallArgs struct {
	// The CPU architecture to install.
	Architecture ArchitecturePtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The directory to install the binary to.
	Directory pulumi.StringPtrInput
	// The version to install.
	Version pulumi.StringPtrInput
}

func (KubeProxyInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeProxyInstallArgs)(nil)).Elem()
}

type KubeProxyInstallInput interface {
	pulumi.Input

	ToKubeProxyInstallOutput() KubeProxyInstallOutput
	ToKubeProxyInstallOutputWithContext(ctx context.Context) KubeProxyInstallOutput
}

func (*KubeProxyInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeProxyInstall)(nil)).Elem()
}

func (i *KubeProxyInstall) ToKubeProxyInstallOutput() KubeProxyInstallOutput {
	return i.ToKubeProxyInstallOutputWithContext(context.Background())
}

func (i *KubeProxyInstall) ToKubeProxyInstallOutputWithContext(ctx context.Context) KubeProxyInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeProxyInstallOutput)
}

// KubeProxyInstallArrayInput is an input type that accepts KubeProxyInstallArray and KubeProxyInstallArrayOutput values.
// You can construct a concrete instance of `KubeProxyInstallArrayInput` via:
//
//	KubeProxyInstallArray{ KubeProxyInstallArgs{...} }
type KubeProxyInstallArrayInput interface {
	pulumi.Input

	ToKubeProxyInstallArrayOutput() KubeProxyInstallArrayOutput
	ToKubeProxyInstallArrayOutputWithContext(context.Context) KubeProxyInstallArrayOutput
}

type KubeProxyInstallArray []KubeProxyInstallInput

func (KubeProxyInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeProxyInstall)(nil)).Elem()
}

func (i KubeProxyInstallArray) ToKubeProxyInstallArrayOutput() KubeProxyInstallArrayOutput {
	return i.ToKubeProxyInstallArrayOutputWithContext(context.Background())
}

func (i KubeProxyInstallArray) ToKubeProxyInstallArrayOutputWithContext(ctx context.Context) KubeProxyInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeProxyInstallArrayOutput)
}

// KubeProxyInstallMapInput is an input type that accepts KubeProxyInstallMap and KubeProxyInstallMapOutput values.
// You can construct a concrete instance of `KubeProxyInstallMapInput` via:
//
//	KubeProxyInstallMap{ "key": KubeProxyInstallArgs{...} }
type KubeProxyInstallMapInput interface {
	pulumi.Input

	ToKubeProxyInstallMapOutput() KubeProxyInstallMapOutput
	ToKubeProxyInstallMapOutputWithContext(context.Context) KubeProxyInstallMapOutput
}

type KubeProxyInstallMap map[string]KubeProxyInstallInput

func (KubeProxyInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeProxyInstall)(nil)).Elem()
}

func (i KubeProxyInstallMap) ToKubeProxyInstallMapOutput() KubeProxyInstallMapOutput {
	return i.ToKubeProxyInstallMapOutputWithContext(context.Background())
}

func (i KubeProxyInstallMap) ToKubeProxyInstallMapOutputWithContext(ctx context.Context) KubeProxyInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeProxyInstallMapOutput)
}

type KubeProxyInstallOutput struct{ *pulumi.OutputState }

func (KubeProxyInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeProxyInstall)(nil)).Elem()
}

func (o KubeProxyInstallOutput) ToKubeProxyInstallOutput() KubeProxyInstallOutput {
	return o
}

func (o KubeProxyInstallOutput) ToKubeProxyInstallOutputWithContext(ctx context.Context) KubeProxyInstallOutput {
	return o
}

// The CPU architecture to install.
func (o KubeProxyInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *KubeProxyInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

// The name of the installed binary.
func (o KubeProxyInstallOutput) BinName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeProxyInstall) pulumi.StringPtrOutput { return v.BinName }).(pulumi.StringPtrOutput)
}

// The parameters with which to connect to the remote host.
func (o KubeProxyInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeProxyInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The directory to install the binary to.
func (o KubeProxyInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeProxyInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

// The download operation.
func (o KubeProxyInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *KubeProxyInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

// The mkdir operation.
func (o KubeProxyInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *KubeProxyInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

// The mktemp operation.
func (o KubeProxyInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *KubeProxyInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

// The mv operation.
func (o KubeProxyInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *KubeProxyInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

// The path to the installed binary.
func (o KubeProxyInstallOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeProxyInstall) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The rm operation.
func (o KubeProxyInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *KubeProxyInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

// The url used to download the binary.
func (o KubeProxyInstallOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeProxyInstall) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The version to install.
func (o KubeProxyInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeProxyInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubeProxyInstallArrayOutput struct{ *pulumi.OutputState }

func (KubeProxyInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeProxyInstall)(nil)).Elem()
}

func (o KubeProxyInstallArrayOutput) ToKubeProxyInstallArrayOutput() KubeProxyInstallArrayOutput {
	return o
}

func (o KubeProxyInstallArrayOutput) ToKubeProxyInstallArrayOutputWithContext(ctx context.Context) KubeProxyInstallArrayOutput {
	return o
}

func (o KubeProxyInstallArrayOutput) Index(i pulumi.IntInput) KubeProxyInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeProxyInstall {
		return vs[0].([]*KubeProxyInstall)[vs[1].(int)]
	}).(KubeProxyInstallOutput)
}

type KubeProxyInstallMapOutput struct{ *pulumi.OutputState }

func (KubeProxyInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeProxyInstall)(nil)).Elem()
}

func (o KubeProxyInstallMapOutput) ToKubeProxyInstallMapOutput() KubeProxyInstallMapOutput {
	return o
}

func (o KubeProxyInstallMapOutput) ToKubeProxyInstallMapOutputWithContext(ctx context.Context) KubeProxyInstallMapOutput {
	return o
}

func (o KubeProxyInstallMapOutput) MapIndex(k pulumi.StringInput) KubeProxyInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeProxyInstall {
		return vs[0].(map[string]*KubeProxyInstall)[vs[1].(string)]
	}).(KubeProxyInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeProxyInstallInput)(nil)).Elem(), &KubeProxyInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeProxyInstallArrayInput)(nil)).Elem(), KubeProxyInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeProxyInstallMapInput)(nil)).Elem(), KubeProxyInstallMap{})
	pulumi.RegisterOutputType(KubeProxyInstallOutput{})
	pulumi.RegisterOutputType(KubeProxyInstallArrayOutput{})
	pulumi.RegisterOutputType(KubeProxyInstallMapOutput{})
}
