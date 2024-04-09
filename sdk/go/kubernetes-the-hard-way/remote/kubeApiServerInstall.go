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

// Installs kube-apiserver on a remote system.
type KubeApiServerInstall struct {
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

// NewKubeApiServerInstall registers a new resource with the given unique name, arguments, and options.
func NewKubeApiServerInstall(ctx *pulumi.Context,
	name string, args *KubeApiServerInstallArgs, opts ...pulumi.ResourceOption) (*KubeApiServerInstall, error) {
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
	var resource KubeApiServerInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeApiServerInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeApiServerInstallArgs struct {
	// The CPU architecture to install.
	Architecture *Architecture `pulumi:"architecture"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The directory to install the binary to.
	Directory *string `pulumi:"directory"`
	// The version to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a KubeApiServerInstall resource.
type KubeApiServerInstallArgs struct {
	// The CPU architecture to install.
	Architecture ArchitecturePtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The directory to install the binary to.
	Directory pulumi.StringPtrInput
	// The version to install.
	Version pulumi.StringPtrInput
}

func (KubeApiServerInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeApiServerInstallArgs)(nil)).Elem()
}

type KubeApiServerInstallInput interface {
	pulumi.Input

	ToKubeApiServerInstallOutput() KubeApiServerInstallOutput
	ToKubeApiServerInstallOutputWithContext(ctx context.Context) KubeApiServerInstallOutput
}

func (*KubeApiServerInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeApiServerInstall)(nil)).Elem()
}

func (i *KubeApiServerInstall) ToKubeApiServerInstallOutput() KubeApiServerInstallOutput {
	return i.ToKubeApiServerInstallOutputWithContext(context.Background())
}

func (i *KubeApiServerInstall) ToKubeApiServerInstallOutputWithContext(ctx context.Context) KubeApiServerInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerInstallOutput)
}

// KubeApiServerInstallArrayInput is an input type that accepts KubeApiServerInstallArray and KubeApiServerInstallArrayOutput values.
// You can construct a concrete instance of `KubeApiServerInstallArrayInput` via:
//
//	KubeApiServerInstallArray{ KubeApiServerInstallArgs{...} }
type KubeApiServerInstallArrayInput interface {
	pulumi.Input

	ToKubeApiServerInstallArrayOutput() KubeApiServerInstallArrayOutput
	ToKubeApiServerInstallArrayOutputWithContext(context.Context) KubeApiServerInstallArrayOutput
}

type KubeApiServerInstallArray []KubeApiServerInstallInput

func (KubeApiServerInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeApiServerInstall)(nil)).Elem()
}

func (i KubeApiServerInstallArray) ToKubeApiServerInstallArrayOutput() KubeApiServerInstallArrayOutput {
	return i.ToKubeApiServerInstallArrayOutputWithContext(context.Background())
}

func (i KubeApiServerInstallArray) ToKubeApiServerInstallArrayOutputWithContext(ctx context.Context) KubeApiServerInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerInstallArrayOutput)
}

// KubeApiServerInstallMapInput is an input type that accepts KubeApiServerInstallMap and KubeApiServerInstallMapOutput values.
// You can construct a concrete instance of `KubeApiServerInstallMapInput` via:
//
//	KubeApiServerInstallMap{ "key": KubeApiServerInstallArgs{...} }
type KubeApiServerInstallMapInput interface {
	pulumi.Input

	ToKubeApiServerInstallMapOutput() KubeApiServerInstallMapOutput
	ToKubeApiServerInstallMapOutputWithContext(context.Context) KubeApiServerInstallMapOutput
}

type KubeApiServerInstallMap map[string]KubeApiServerInstallInput

func (KubeApiServerInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeApiServerInstall)(nil)).Elem()
}

func (i KubeApiServerInstallMap) ToKubeApiServerInstallMapOutput() KubeApiServerInstallMapOutput {
	return i.ToKubeApiServerInstallMapOutputWithContext(context.Background())
}

func (i KubeApiServerInstallMap) ToKubeApiServerInstallMapOutputWithContext(ctx context.Context) KubeApiServerInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerInstallMapOutput)
}

type KubeApiServerInstallOutput struct{ *pulumi.OutputState }

func (KubeApiServerInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeApiServerInstall)(nil)).Elem()
}

func (o KubeApiServerInstallOutput) ToKubeApiServerInstallOutput() KubeApiServerInstallOutput {
	return o
}

func (o KubeApiServerInstallOutput) ToKubeApiServerInstallOutputWithContext(ctx context.Context) KubeApiServerInstallOutput {
	return o
}

// The CPU architecture to install.
func (o KubeApiServerInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

// The name of the installed binary.
func (o KubeApiServerInstallOutput) BinName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) pulumi.StringPtrOutput { return v.BinName }).(pulumi.StringPtrOutput)
}

// The parameters with which to connect to the remote host.
func (o KubeApiServerInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The directory to install the binary to.
func (o KubeApiServerInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

// The download operation.
func (o KubeApiServerInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

// The mkdir operation.
func (o KubeApiServerInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

// The mktemp operation.
func (o KubeApiServerInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

// The mv operation.
func (o KubeApiServerInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

// The path to the installed binary.
func (o KubeApiServerInstallOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The rm operation.
func (o KubeApiServerInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

// The url used to download the binary.
func (o KubeApiServerInstallOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The version to install.
func (o KubeApiServerInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeApiServerInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubeApiServerInstallArrayOutput struct{ *pulumi.OutputState }

func (KubeApiServerInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeApiServerInstall)(nil)).Elem()
}

func (o KubeApiServerInstallArrayOutput) ToKubeApiServerInstallArrayOutput() KubeApiServerInstallArrayOutput {
	return o
}

func (o KubeApiServerInstallArrayOutput) ToKubeApiServerInstallArrayOutputWithContext(ctx context.Context) KubeApiServerInstallArrayOutput {
	return o
}

func (o KubeApiServerInstallArrayOutput) Index(i pulumi.IntInput) KubeApiServerInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeApiServerInstall {
		return vs[0].([]*KubeApiServerInstall)[vs[1].(int)]
	}).(KubeApiServerInstallOutput)
}

type KubeApiServerInstallMapOutput struct{ *pulumi.OutputState }

func (KubeApiServerInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeApiServerInstall)(nil)).Elem()
}

func (o KubeApiServerInstallMapOutput) ToKubeApiServerInstallMapOutput() KubeApiServerInstallMapOutput {
	return o
}

func (o KubeApiServerInstallMapOutput) ToKubeApiServerInstallMapOutputWithContext(ctx context.Context) KubeApiServerInstallMapOutput {
	return o
}

func (o KubeApiServerInstallMapOutput) MapIndex(k pulumi.StringInput) KubeApiServerInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeApiServerInstall {
		return vs[0].(map[string]*KubeApiServerInstall)[vs[1].(string)]
	}).(KubeApiServerInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerInstallInput)(nil)).Elem(), &KubeApiServerInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerInstallArrayInput)(nil)).Elem(), KubeApiServerInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerInstallMapInput)(nil)).Elem(), KubeApiServerInstallMap{})
	pulumi.RegisterOutputType(KubeApiServerInstallOutput{})
	pulumi.RegisterOutputType(KubeApiServerInstallArrayOutput{})
	pulumi.RegisterOutputType(KubeApiServerInstallMapOutput{})
}
