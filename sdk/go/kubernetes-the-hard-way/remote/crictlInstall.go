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

// Installs crictl on a remote system.
type CrictlInstall struct {
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

// NewCrictlInstall registers a new resource with the given unique name, arguments, and options.
func NewCrictlInstall(ctx *pulumi.Context,
	name string, args *CrictlInstallArgs, opts ...pulumi.ResourceOption) (*CrictlInstall, error) {
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
	var resource CrictlInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:CrictlInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type crictlInstallArgs struct {
	// The CPU architecture.
	Architecture *Architecture `pulumi:"architecture"`
	// The connection details.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Directory to install the binary.
	Directory *string `pulumi:"directory"`
	// The version of to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a CrictlInstall resource.
type CrictlInstallArgs struct {
	// The CPU architecture.
	Architecture ArchitecturePtrInput
	// The connection details.
	Connection pulumiCommand.ConnectionInput
	// Directory to install the binary.
	Directory pulumi.StringPtrInput
	// The version of to install.
	Version pulumi.StringPtrInput
}

func (CrictlInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*crictlInstallArgs)(nil)).Elem()
}

type CrictlInstallInput interface {
	pulumi.Input

	ToCrictlInstallOutput() CrictlInstallOutput
	ToCrictlInstallOutputWithContext(ctx context.Context) CrictlInstallOutput
}

func (*CrictlInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**CrictlInstall)(nil)).Elem()
}

func (i *CrictlInstall) ToCrictlInstallOutput() CrictlInstallOutput {
	return i.ToCrictlInstallOutputWithContext(context.Background())
}

func (i *CrictlInstall) ToCrictlInstallOutputWithContext(ctx context.Context) CrictlInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrictlInstallOutput)
}

// CrictlInstallArrayInput is an input type that accepts CrictlInstallArray and CrictlInstallArrayOutput values.
// You can construct a concrete instance of `CrictlInstallArrayInput` via:
//
//	CrictlInstallArray{ CrictlInstallArgs{...} }
type CrictlInstallArrayInput interface {
	pulumi.Input

	ToCrictlInstallArrayOutput() CrictlInstallArrayOutput
	ToCrictlInstallArrayOutputWithContext(context.Context) CrictlInstallArrayOutput
}

type CrictlInstallArray []CrictlInstallInput

func (CrictlInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CrictlInstall)(nil)).Elem()
}

func (i CrictlInstallArray) ToCrictlInstallArrayOutput() CrictlInstallArrayOutput {
	return i.ToCrictlInstallArrayOutputWithContext(context.Background())
}

func (i CrictlInstallArray) ToCrictlInstallArrayOutputWithContext(ctx context.Context) CrictlInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrictlInstallArrayOutput)
}

// CrictlInstallMapInput is an input type that accepts CrictlInstallMap and CrictlInstallMapOutput values.
// You can construct a concrete instance of `CrictlInstallMapInput` via:
//
//	CrictlInstallMap{ "key": CrictlInstallArgs{...} }
type CrictlInstallMapInput interface {
	pulumi.Input

	ToCrictlInstallMapOutput() CrictlInstallMapOutput
	ToCrictlInstallMapOutputWithContext(context.Context) CrictlInstallMapOutput
}

type CrictlInstallMap map[string]CrictlInstallInput

func (CrictlInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CrictlInstall)(nil)).Elem()
}

func (i CrictlInstallMap) ToCrictlInstallMapOutput() CrictlInstallMapOutput {
	return i.ToCrictlInstallMapOutputWithContext(context.Background())
}

func (i CrictlInstallMap) ToCrictlInstallMapOutputWithContext(ctx context.Context) CrictlInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrictlInstallMapOutput)
}

type CrictlInstallOutput struct{ *pulumi.OutputState }

func (CrictlInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrictlInstall)(nil)).Elem()
}

func (o CrictlInstallOutput) ToCrictlInstallOutput() CrictlInstallOutput {
	return o
}

func (o CrictlInstallOutput) ToCrictlInstallOutputWithContext(ctx context.Context) CrictlInstallOutput {
	return o
}

// The CPU architecture.
func (o CrictlInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *CrictlInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

func (o CrictlInstallOutput) ArchiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrictlInstall) pulumi.StringPtrOutput { return v.ArchiveName }).(pulumi.StringPtrOutput)
}

// The connection details.
func (o CrictlInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *CrictlInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Directory to install the binary.
func (o CrictlInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *CrictlInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

func (o CrictlInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *CrictlInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

func (o CrictlInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *CrictlInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

func (o CrictlInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *CrictlInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

func (o CrictlInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *CrictlInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

func (o CrictlInstallOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *CrictlInstall) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o CrictlInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *CrictlInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

func (o CrictlInstallOutput) Tar() tools.TarOutput {
	return o.ApplyT(func(v *CrictlInstall) tools.TarOutput { return v.Tar }).(tools.TarOutput)
}

func (o CrictlInstallOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrictlInstall) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// The version to install.
func (o CrictlInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *CrictlInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type CrictlInstallArrayOutput struct{ *pulumi.OutputState }

func (CrictlInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CrictlInstall)(nil)).Elem()
}

func (o CrictlInstallArrayOutput) ToCrictlInstallArrayOutput() CrictlInstallArrayOutput {
	return o
}

func (o CrictlInstallArrayOutput) ToCrictlInstallArrayOutputWithContext(ctx context.Context) CrictlInstallArrayOutput {
	return o
}

func (o CrictlInstallArrayOutput) Index(i pulumi.IntInput) CrictlInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CrictlInstall {
		return vs[0].([]*CrictlInstall)[vs[1].(int)]
	}).(CrictlInstallOutput)
}

type CrictlInstallMapOutput struct{ *pulumi.OutputState }

func (CrictlInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CrictlInstall)(nil)).Elem()
}

func (o CrictlInstallMapOutput) ToCrictlInstallMapOutput() CrictlInstallMapOutput {
	return o
}

func (o CrictlInstallMapOutput) ToCrictlInstallMapOutputWithContext(ctx context.Context) CrictlInstallMapOutput {
	return o
}

func (o CrictlInstallMapOutput) MapIndex(k pulumi.StringInput) CrictlInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CrictlInstall {
		return vs[0].(map[string]*CrictlInstall)[vs[1].(string)]
	}).(CrictlInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CrictlInstallInput)(nil)).Elem(), &CrictlInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrictlInstallArrayInput)(nil)).Elem(), CrictlInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CrictlInstallMapInput)(nil)).Elem(), CrictlInstallMap{})
	pulumi.RegisterOutputType(CrictlInstallOutput{})
	pulumi.RegisterOutputType(CrictlInstallArrayOutput{})
	pulumi.RegisterOutputType(CrictlInstallMapOutput{})
}
