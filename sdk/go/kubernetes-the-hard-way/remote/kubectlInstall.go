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

// Installs kubectl on a remote system.
type KubectlInstall struct {
	pulumi.ResourceState

	// The CPU architecture.
	Architecture ArchitectureOutput     `pulumi:"architecture"`
	ArchiveName  pulumi.StringPtrOutput `pulumi:"archiveName"`
	// The connection details.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Directory to install the binary.
	Directory pulumi.StringOutput `pulumi:"directory"`
	Download  DownloadOutput      `pulumi:"download"`
	Mkdir     tools.MkdirOutput   `pulumi:"mkdir"`
	Mktemp    tools.MktempOutput  `pulumi:"mktemp"`
	Mv        tools.MvOutput      `pulumi:"mv"`
	Path      pulumi.StringOutput `pulumi:"path"`
	Rm        tools.RmOutput      `pulumi:"rm"`
	// The version to install.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewKubectlInstall registers a new resource with the given unique name, arguments, and options.
func NewKubectlInstall(ctx *pulumi.Context,
	name string, args *KubectlInstallArgs, opts ...pulumi.ResourceOption) (*KubectlInstall, error) {
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
	var resource KubectlInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubectlInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubectlInstallArgs struct {
	// The CPU architecture.
	Architecture *Architecture `pulumi:"architecture"`
	// The connection details.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Directory to install the binary.
	Directory *string `pulumi:"directory"`
	// The version of to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a KubectlInstall resource.
type KubectlInstallArgs struct {
	// The CPU architecture.
	Architecture ArchitecturePtrInput
	// The connection details.
	Connection pulumiCommand.ConnectionInput
	// Directory to install the binary.
	Directory pulumi.StringPtrInput
	// The version of to install.
	Version pulumi.StringPtrInput
}

func (KubectlInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubectlInstallArgs)(nil)).Elem()
}

type KubectlInstallInput interface {
	pulumi.Input

	ToKubectlInstallOutput() KubectlInstallOutput
	ToKubectlInstallOutputWithContext(ctx context.Context) KubectlInstallOutput
}

func (*KubectlInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**KubectlInstall)(nil)).Elem()
}

func (i *KubectlInstall) ToKubectlInstallOutput() KubectlInstallOutput {
	return i.ToKubectlInstallOutputWithContext(context.Background())
}

func (i *KubectlInstall) ToKubectlInstallOutputWithContext(ctx context.Context) KubectlInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubectlInstallOutput)
}

// KubectlInstallArrayInput is an input type that accepts KubectlInstallArray and KubectlInstallArrayOutput values.
// You can construct a concrete instance of `KubectlInstallArrayInput` via:
//
//	KubectlInstallArray{ KubectlInstallArgs{...} }
type KubectlInstallArrayInput interface {
	pulumi.Input

	ToKubectlInstallArrayOutput() KubectlInstallArrayOutput
	ToKubectlInstallArrayOutputWithContext(context.Context) KubectlInstallArrayOutput
}

type KubectlInstallArray []KubectlInstallInput

func (KubectlInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubectlInstall)(nil)).Elem()
}

func (i KubectlInstallArray) ToKubectlInstallArrayOutput() KubectlInstallArrayOutput {
	return i.ToKubectlInstallArrayOutputWithContext(context.Background())
}

func (i KubectlInstallArray) ToKubectlInstallArrayOutputWithContext(ctx context.Context) KubectlInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubectlInstallArrayOutput)
}

// KubectlInstallMapInput is an input type that accepts KubectlInstallMap and KubectlInstallMapOutput values.
// You can construct a concrete instance of `KubectlInstallMapInput` via:
//
//	KubectlInstallMap{ "key": KubectlInstallArgs{...} }
type KubectlInstallMapInput interface {
	pulumi.Input

	ToKubectlInstallMapOutput() KubectlInstallMapOutput
	ToKubectlInstallMapOutputWithContext(context.Context) KubectlInstallMapOutput
}

type KubectlInstallMap map[string]KubectlInstallInput

func (KubectlInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubectlInstall)(nil)).Elem()
}

func (i KubectlInstallMap) ToKubectlInstallMapOutput() KubectlInstallMapOutput {
	return i.ToKubectlInstallMapOutputWithContext(context.Background())
}

func (i KubectlInstallMap) ToKubectlInstallMapOutputWithContext(ctx context.Context) KubectlInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubectlInstallMapOutput)
}

type KubectlInstallOutput struct{ *pulumi.OutputState }

func (KubectlInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubectlInstall)(nil)).Elem()
}

func (o KubectlInstallOutput) ToKubectlInstallOutput() KubectlInstallOutput {
	return o
}

func (o KubectlInstallOutput) ToKubectlInstallOutputWithContext(ctx context.Context) KubectlInstallOutput {
	return o
}

// The CPU architecture.
func (o KubectlInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *KubectlInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

func (o KubectlInstallOutput) ArchiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubectlInstall) pulumi.StringPtrOutput { return v.ArchiveName }).(pulumi.StringPtrOutput)
}

// The connection details.
func (o KubectlInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubectlInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Directory to install the binary.
func (o KubectlInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *KubectlInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

func (o KubectlInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *KubectlInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

func (o KubectlInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *KubectlInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

func (o KubectlInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *KubectlInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

func (o KubectlInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *KubectlInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

func (o KubectlInstallOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *KubectlInstall) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o KubectlInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *KubectlInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

// The version to install.
func (o KubectlInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *KubectlInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubectlInstallArrayOutput struct{ *pulumi.OutputState }

func (KubectlInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubectlInstall)(nil)).Elem()
}

func (o KubectlInstallArrayOutput) ToKubectlInstallArrayOutput() KubectlInstallArrayOutput {
	return o
}

func (o KubectlInstallArrayOutput) ToKubectlInstallArrayOutputWithContext(ctx context.Context) KubectlInstallArrayOutput {
	return o
}

func (o KubectlInstallArrayOutput) Index(i pulumi.IntInput) KubectlInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubectlInstall {
		return vs[0].([]*KubectlInstall)[vs[1].(int)]
	}).(KubectlInstallOutput)
}

type KubectlInstallMapOutput struct{ *pulumi.OutputState }

func (KubectlInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubectlInstall)(nil)).Elem()
}

func (o KubectlInstallMapOutput) ToKubectlInstallMapOutput() KubectlInstallMapOutput {
	return o
}

func (o KubectlInstallMapOutput) ToKubectlInstallMapOutputWithContext(ctx context.Context) KubectlInstallMapOutput {
	return o
}

func (o KubectlInstallMapOutput) MapIndex(k pulumi.StringInput) KubectlInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubectlInstall {
		return vs[0].(map[string]*KubectlInstall)[vs[1].(string)]
	}).(KubectlInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubectlInstallInput)(nil)).Elem(), &KubectlInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubectlInstallArrayInput)(nil)).Elem(), KubectlInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubectlInstallMapInput)(nil)).Elem(), KubectlInstallMap{})
	pulumi.RegisterOutputType(KubectlInstallOutput{})
	pulumi.RegisterOutputType(KubectlInstallArrayOutput{})
	pulumi.RegisterOutputType(KubectlInstallMapOutput{})
}
