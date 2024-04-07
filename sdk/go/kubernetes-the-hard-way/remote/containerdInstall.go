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

// Installs containerd on a remote system.
type ContainerdInstall struct {
	pulumi.ResourceState

	// The CPU architecture.
	Architecture ArchitectureOutput     `pulumi:"architecture"`
	ArchiveName  pulumi.StringPtrOutput `pulumi:"archiveName"`
	// The connection details.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Directory to install the binary.
	Directory pulumi.StringOutput    `pulumi:"directory"`
	Download  DownloadOutput         `pulumi:"download"`
	Mkdir     tools.MkdirOutput      `pulumi:"mkdir"`
	Mktemp    tools.MktempOutput     `pulumi:"mktemp"`
	Mv        tools.MvOutput         `pulumi:"mv"`
	Path      pulumi.StringOutput    `pulumi:"path"`
	Rm        tools.RmOutput         `pulumi:"rm"`
	Tar       tools.TarOutput        `pulumi:"tar"`
	Url       pulumi.StringPtrOutput `pulumi:"url"`
	// The version to install.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewContainerdInstall registers a new resource with the given unique name, arguments, and options.
func NewContainerdInstall(ctx *pulumi.Context,
	name string, args *ContainerdInstallArgs, opts ...pulumi.ResourceOption) (*ContainerdInstall, error) {
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
	var resource ContainerdInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:ContainerdInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type containerdInstallArgs struct {
	// The CPU architecture.
	Architecture *Architecture `pulumi:"architecture"`
	// The connection details.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Directory to install the binary.
	Directory *string `pulumi:"directory"`
	// The version of to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a ContainerdInstall resource.
type ContainerdInstallArgs struct {
	// The CPU architecture.
	Architecture ArchitecturePtrInput
	// The connection details.
	Connection pulumiCommand.ConnectionInput
	// Directory to install the binary.
	Directory pulumi.StringPtrInput
	// The version of to install.
	Version pulumi.StringPtrInput
}

func (ContainerdInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerdInstallArgs)(nil)).Elem()
}

type ContainerdInstallInput interface {
	pulumi.Input

	ToContainerdInstallOutput() ContainerdInstallOutput
	ToContainerdInstallOutputWithContext(ctx context.Context) ContainerdInstallOutput
}

func (*ContainerdInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerdInstall)(nil)).Elem()
}

func (i *ContainerdInstall) ToContainerdInstallOutput() ContainerdInstallOutput {
	return i.ToContainerdInstallOutputWithContext(context.Background())
}

func (i *ContainerdInstall) ToContainerdInstallOutputWithContext(ctx context.Context) ContainerdInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerdInstallOutput)
}

// ContainerdInstallArrayInput is an input type that accepts ContainerdInstallArray and ContainerdInstallArrayOutput values.
// You can construct a concrete instance of `ContainerdInstallArrayInput` via:
//
//	ContainerdInstallArray{ ContainerdInstallArgs{...} }
type ContainerdInstallArrayInput interface {
	pulumi.Input

	ToContainerdInstallArrayOutput() ContainerdInstallArrayOutput
	ToContainerdInstallArrayOutputWithContext(context.Context) ContainerdInstallArrayOutput
}

type ContainerdInstallArray []ContainerdInstallInput

func (ContainerdInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerdInstall)(nil)).Elem()
}

func (i ContainerdInstallArray) ToContainerdInstallArrayOutput() ContainerdInstallArrayOutput {
	return i.ToContainerdInstallArrayOutputWithContext(context.Background())
}

func (i ContainerdInstallArray) ToContainerdInstallArrayOutputWithContext(ctx context.Context) ContainerdInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerdInstallArrayOutput)
}

// ContainerdInstallMapInput is an input type that accepts ContainerdInstallMap and ContainerdInstallMapOutput values.
// You can construct a concrete instance of `ContainerdInstallMapInput` via:
//
//	ContainerdInstallMap{ "key": ContainerdInstallArgs{...} }
type ContainerdInstallMapInput interface {
	pulumi.Input

	ToContainerdInstallMapOutput() ContainerdInstallMapOutput
	ToContainerdInstallMapOutputWithContext(context.Context) ContainerdInstallMapOutput
}

type ContainerdInstallMap map[string]ContainerdInstallInput

func (ContainerdInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerdInstall)(nil)).Elem()
}

func (i ContainerdInstallMap) ToContainerdInstallMapOutput() ContainerdInstallMapOutput {
	return i.ToContainerdInstallMapOutputWithContext(context.Background())
}

func (i ContainerdInstallMap) ToContainerdInstallMapOutputWithContext(ctx context.Context) ContainerdInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerdInstallMapOutput)
}

type ContainerdInstallOutput struct{ *pulumi.OutputState }

func (ContainerdInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerdInstall)(nil)).Elem()
}

func (o ContainerdInstallOutput) ToContainerdInstallOutput() ContainerdInstallOutput {
	return o
}

func (o ContainerdInstallOutput) ToContainerdInstallOutputWithContext(ctx context.Context) ContainerdInstallOutput {
	return o
}

// The CPU architecture.
func (o ContainerdInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *ContainerdInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

func (o ContainerdInstallOutput) ArchiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerdInstall) pulumi.StringPtrOutput { return v.ArchiveName }).(pulumi.StringPtrOutput)
}

// The connection details.
func (o ContainerdInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *ContainerdInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Directory to install the binary.
func (o ContainerdInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerdInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

func (o ContainerdInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *ContainerdInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

func (o ContainerdInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *ContainerdInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

func (o ContainerdInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *ContainerdInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

func (o ContainerdInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *ContainerdInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

func (o ContainerdInstallOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerdInstall) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o ContainerdInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *ContainerdInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

func (o ContainerdInstallOutput) Tar() tools.TarOutput {
	return o.ApplyT(func(v *ContainerdInstall) tools.TarOutput { return v.Tar }).(tools.TarOutput)
}

func (o ContainerdInstallOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerdInstall) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// The version to install.
func (o ContainerdInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerdInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type ContainerdInstallArrayOutput struct{ *pulumi.OutputState }

func (ContainerdInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerdInstall)(nil)).Elem()
}

func (o ContainerdInstallArrayOutput) ToContainerdInstallArrayOutput() ContainerdInstallArrayOutput {
	return o
}

func (o ContainerdInstallArrayOutput) ToContainerdInstallArrayOutputWithContext(ctx context.Context) ContainerdInstallArrayOutput {
	return o
}

func (o ContainerdInstallArrayOutput) Index(i pulumi.IntInput) ContainerdInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerdInstall {
		return vs[0].([]*ContainerdInstall)[vs[1].(int)]
	}).(ContainerdInstallOutput)
}

type ContainerdInstallMapOutput struct{ *pulumi.OutputState }

func (ContainerdInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerdInstall)(nil)).Elem()
}

func (o ContainerdInstallMapOutput) ToContainerdInstallMapOutput() ContainerdInstallMapOutput {
	return o
}

func (o ContainerdInstallMapOutput) ToContainerdInstallMapOutputWithContext(ctx context.Context) ContainerdInstallMapOutput {
	return o
}

func (o ContainerdInstallMapOutput) MapIndex(k pulumi.StringInput) ContainerdInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerdInstall {
		return vs[0].(map[string]*ContainerdInstall)[vs[1].(string)]
	}).(ContainerdInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerdInstallInput)(nil)).Elem(), &ContainerdInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerdInstallArrayInput)(nil)).Elem(), ContainerdInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerdInstallMapInput)(nil)).Elem(), ContainerdInstallMap{})
	pulumi.RegisterOutputType(ContainerdInstallOutput{})
	pulumi.RegisterOutputType(ContainerdInstallArrayOutput{})
	pulumi.RegisterOutputType(ContainerdInstallMapOutput{})
}