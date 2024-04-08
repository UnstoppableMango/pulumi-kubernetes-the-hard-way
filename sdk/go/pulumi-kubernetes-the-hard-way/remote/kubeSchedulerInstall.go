// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/internal"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/tools"
)

// Installs kube-scheduler on a remote system.
type KubeSchedulerInstall struct {
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

// NewKubeSchedulerInstall registers a new resource with the given unique name, arguments, and options.
func NewKubeSchedulerInstall(ctx *pulumi.Context,
	name string, args *KubeSchedulerInstallArgs, opts ...pulumi.ResourceOption) (*KubeSchedulerInstall, error) {
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
	var resource KubeSchedulerInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeSchedulerInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeSchedulerInstallArgs struct {
	// The CPU architecture to install.
	Architecture *Architecture `pulumi:"architecture"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The directory to install the binary to.
	Directory *string `pulumi:"directory"`
	// The version to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a KubeSchedulerInstall resource.
type KubeSchedulerInstallArgs struct {
	// The CPU architecture to install.
	Architecture ArchitecturePtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The directory to install the binary to.
	Directory pulumi.StringPtrInput
	// The version to install.
	Version pulumi.StringPtrInput
}

func (KubeSchedulerInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeSchedulerInstallArgs)(nil)).Elem()
}

type KubeSchedulerInstallInput interface {
	pulumi.Input

	ToKubeSchedulerInstallOutput() KubeSchedulerInstallOutput
	ToKubeSchedulerInstallOutputWithContext(ctx context.Context) KubeSchedulerInstallOutput
}

func (*KubeSchedulerInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeSchedulerInstall)(nil)).Elem()
}

func (i *KubeSchedulerInstall) ToKubeSchedulerInstallOutput() KubeSchedulerInstallOutput {
	return i.ToKubeSchedulerInstallOutputWithContext(context.Background())
}

func (i *KubeSchedulerInstall) ToKubeSchedulerInstallOutputWithContext(ctx context.Context) KubeSchedulerInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeSchedulerInstallOutput)
}

// KubeSchedulerInstallArrayInput is an input type that accepts KubeSchedulerInstallArray and KubeSchedulerInstallArrayOutput values.
// You can construct a concrete instance of `KubeSchedulerInstallArrayInput` via:
//
//	KubeSchedulerInstallArray{ KubeSchedulerInstallArgs{...} }
type KubeSchedulerInstallArrayInput interface {
	pulumi.Input

	ToKubeSchedulerInstallArrayOutput() KubeSchedulerInstallArrayOutput
	ToKubeSchedulerInstallArrayOutputWithContext(context.Context) KubeSchedulerInstallArrayOutput
}

type KubeSchedulerInstallArray []KubeSchedulerInstallInput

func (KubeSchedulerInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeSchedulerInstall)(nil)).Elem()
}

func (i KubeSchedulerInstallArray) ToKubeSchedulerInstallArrayOutput() KubeSchedulerInstallArrayOutput {
	return i.ToKubeSchedulerInstallArrayOutputWithContext(context.Background())
}

func (i KubeSchedulerInstallArray) ToKubeSchedulerInstallArrayOutputWithContext(ctx context.Context) KubeSchedulerInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeSchedulerInstallArrayOutput)
}

// KubeSchedulerInstallMapInput is an input type that accepts KubeSchedulerInstallMap and KubeSchedulerInstallMapOutput values.
// You can construct a concrete instance of `KubeSchedulerInstallMapInput` via:
//
//	KubeSchedulerInstallMap{ "key": KubeSchedulerInstallArgs{...} }
type KubeSchedulerInstallMapInput interface {
	pulumi.Input

	ToKubeSchedulerInstallMapOutput() KubeSchedulerInstallMapOutput
	ToKubeSchedulerInstallMapOutputWithContext(context.Context) KubeSchedulerInstallMapOutput
}

type KubeSchedulerInstallMap map[string]KubeSchedulerInstallInput

func (KubeSchedulerInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeSchedulerInstall)(nil)).Elem()
}

func (i KubeSchedulerInstallMap) ToKubeSchedulerInstallMapOutput() KubeSchedulerInstallMapOutput {
	return i.ToKubeSchedulerInstallMapOutputWithContext(context.Background())
}

func (i KubeSchedulerInstallMap) ToKubeSchedulerInstallMapOutputWithContext(ctx context.Context) KubeSchedulerInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeSchedulerInstallMapOutput)
}

type KubeSchedulerInstallOutput struct{ *pulumi.OutputState }

func (KubeSchedulerInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeSchedulerInstall)(nil)).Elem()
}

func (o KubeSchedulerInstallOutput) ToKubeSchedulerInstallOutput() KubeSchedulerInstallOutput {
	return o
}

func (o KubeSchedulerInstallOutput) ToKubeSchedulerInstallOutputWithContext(ctx context.Context) KubeSchedulerInstallOutput {
	return o
}

// The CPU architecture to install.
func (o KubeSchedulerInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

// The name of the installed binary.
func (o KubeSchedulerInstallOutput) BinName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) pulumi.StringPtrOutput { return v.BinName }).(pulumi.StringPtrOutput)
}

// The parameters with which to connect to the remote host.
func (o KubeSchedulerInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The directory to install the binary to.
func (o KubeSchedulerInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

// The download operation.
func (o KubeSchedulerInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

// The mkdir operation.
func (o KubeSchedulerInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

// The mktemp operation.
func (o KubeSchedulerInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

// The mv operation.
func (o KubeSchedulerInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

// The path to the installed binary.
func (o KubeSchedulerInstallOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The rm operation.
func (o KubeSchedulerInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

// The url used to download the binary.
func (o KubeSchedulerInstallOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The version to install.
func (o KubeSchedulerInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeSchedulerInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubeSchedulerInstallArrayOutput struct{ *pulumi.OutputState }

func (KubeSchedulerInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeSchedulerInstall)(nil)).Elem()
}

func (o KubeSchedulerInstallArrayOutput) ToKubeSchedulerInstallArrayOutput() KubeSchedulerInstallArrayOutput {
	return o
}

func (o KubeSchedulerInstallArrayOutput) ToKubeSchedulerInstallArrayOutputWithContext(ctx context.Context) KubeSchedulerInstallArrayOutput {
	return o
}

func (o KubeSchedulerInstallArrayOutput) Index(i pulumi.IntInput) KubeSchedulerInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeSchedulerInstall {
		return vs[0].([]*KubeSchedulerInstall)[vs[1].(int)]
	}).(KubeSchedulerInstallOutput)
}

type KubeSchedulerInstallMapOutput struct{ *pulumi.OutputState }

func (KubeSchedulerInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeSchedulerInstall)(nil)).Elem()
}

func (o KubeSchedulerInstallMapOutput) ToKubeSchedulerInstallMapOutput() KubeSchedulerInstallMapOutput {
	return o
}

func (o KubeSchedulerInstallMapOutput) ToKubeSchedulerInstallMapOutputWithContext(ctx context.Context) KubeSchedulerInstallMapOutput {
	return o
}

func (o KubeSchedulerInstallMapOutput) MapIndex(k pulumi.StringInput) KubeSchedulerInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeSchedulerInstall {
		return vs[0].(map[string]*KubeSchedulerInstall)[vs[1].(string)]
	}).(KubeSchedulerInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeSchedulerInstallInput)(nil)).Elem(), &KubeSchedulerInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeSchedulerInstallArrayInput)(nil)).Elem(), KubeSchedulerInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeSchedulerInstallMapInput)(nil)).Elem(), KubeSchedulerInstallMap{})
	pulumi.RegisterOutputType(KubeSchedulerInstallOutput{})
	pulumi.RegisterOutputType(KubeSchedulerInstallArrayOutput{})
	pulumi.RegisterOutputType(KubeSchedulerInstallMapOutput{})
}
