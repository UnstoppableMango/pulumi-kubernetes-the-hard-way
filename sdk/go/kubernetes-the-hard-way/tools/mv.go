// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tools

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Abstraction over the `mv` utility on a remote system.
type Mv struct {
	pulumi.ResourceState

	// Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
	Backup pulumi.BoolOutput `pulumi:"backup"`
	// Represents the command run on the remote system.
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// Corresponds to the --context option.
	Context pulumi.BoolOutput `pulumi:"context"`
	// Corresponds to the [CONTROL] argument for the --backup option.
	Control pulumi.StringPtrOutput `pulumi:"control"`
	// Corresponds to the [DEST] argument.
	Dest pulumi.StringPtrOutput `pulumi:"dest"`
	// Corresponds to the [DIRECTORY] argument.
	Directory pulumi.StringPtrOutput `pulumi:"directory"`
	// Corresponds to the --force option.
	Force pulumi.BoolOutput `pulumi:"force"`
	// At what stage(s) in the resource lifecycle should the command be run.
	Lifecycle CommandLifecyclePtrOutput `pulumi:"lifecycle"`
	// Corresponds to the --no-clobber option.
	NoClobber pulumi.BoolOutput `pulumi:"noClobber"`
	// Corresponds to the --no-target-directory option.
	NoTargetDirectory pulumi.BoolOutput `pulumi:"noTargetDirectory"`
	// Corresponds to the [SOURCE] argument.
	Source pulumi.StringArrayOutput `pulumi:"source"`
	// Corresponds to the --strip-trailing-suffix option.
	StripTrailingSlashes pulumi.BoolOutput `pulumi:"stripTrailingSlashes"`
	// Corresponds to the --suffix option.
	Suffix pulumi.StringPtrOutput `pulumi:"suffix"`
	// Corresponds to the --target-directory option.
	TargetDirectory pulumi.StringPtrOutput `pulumi:"targetDirectory"`
	// Corresponds to the --update option.
	Update pulumi.BoolOutput `pulumi:"update"`
	// Corresponds to the --verbose option.
	Verbose pulumi.BoolOutput `pulumi:"verbose"`
}

// NewMv registers a new resource with the given unique name, arguments, and options.
func NewMv(ctx *pulumi.Context,
	name string, args *MvArgs, opts ...pulumi.ResourceOption) (*Mv, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mv
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Mv", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type mvArgs struct {
	// Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
	Backup *bool `pulumi:"backup"`
	// Connection details for the remote system.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Corresponds to the --context option.
	Context *bool `pulumi:"context"`
	// Corresponds to the [CONTROL] argument for the --backup option.
	Control *string `pulumi:"control"`
	// Corresponds to the [DEST] argument.
	Dest *string `pulumi:"dest"`
	// Corresponds to the [DIRECTORY] argument.
	Directory   *string           `pulumi:"directory"`
	Environment map[string]string `pulumi:"environment"`
	// Corresponds to the --force option.
	Force *bool `pulumi:"force"`
	// Corresponds to the --no-clobber option.
	NoClobber *bool `pulumi:"noClobber"`
	// Corresponds to the --no-target-directory option.
	NoTargetDirectory *bool `pulumi:"noTargetDirectory"`
	// Corresponds to the [SOURCE] argument.
	Source interface{} `pulumi:"source"`
	// Corresponds to the --strip-trailing-suffix option.
	StripTrailingSlashes *bool `pulumi:"stripTrailingSlashes"`
	// Corresponds to the --suffix option.
	Suffix *string `pulumi:"suffix"`
	// Corresponds to the --target-directory option.
	TargetDirectory *string `pulumi:"targetDirectory"`
	// Corresponds to the --update option.
	Update *bool `pulumi:"update"`
	// Corresponds to the --verbose option.
	Verbose *bool `pulumi:"verbose"`
}

// The set of arguments for constructing a Mv resource.
type MvArgs struct {
	// Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
	Backup *bool
	// Connection details for the remote system.
	Connection pulumiCommand.ConnectionInput
	// Corresponds to the --context option.
	Context pulumi.BoolPtrInput
	// Corresponds to the [CONTROL] argument for the --backup option.
	Control pulumi.StringPtrInput
	// Corresponds to the [DEST] argument.
	Dest pulumi.StringPtrInput
	// Corresponds to the [DIRECTORY] argument.
	Directory   pulumi.StringPtrInput
	Environment pulumi.StringMapInput
	// Corresponds to the --force option.
	Force pulumi.BoolPtrInput
	// Corresponds to the --no-clobber option.
	NoClobber pulumi.BoolPtrInput
	// Corresponds to the --no-target-directory option.
	NoTargetDirectory pulumi.BoolPtrInput
	// Corresponds to the [SOURCE] argument.
	Source pulumi.Input
	// Corresponds to the --strip-trailing-suffix option.
	StripTrailingSlashes pulumi.BoolPtrInput
	// Corresponds to the --suffix option.
	Suffix pulumi.StringPtrInput
	// Corresponds to the --target-directory option.
	TargetDirectory pulumi.StringPtrInput
	// Corresponds to the --update option.
	Update pulumi.BoolPtrInput
	// Corresponds to the --verbose option.
	Verbose pulumi.BoolPtrInput
}

func (MvArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mvArgs)(nil)).Elem()
}

type MvInput interface {
	pulumi.Input

	ToMvOutput() MvOutput
	ToMvOutputWithContext(ctx context.Context) MvOutput
}

func (*Mv) ElementType() reflect.Type {
	return reflect.TypeOf((**Mv)(nil)).Elem()
}

func (i *Mv) ToMvOutput() MvOutput {
	return i.ToMvOutputWithContext(context.Background())
}

func (i *Mv) ToMvOutputWithContext(ctx context.Context) MvOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MvOutput)
}

// MvArrayInput is an input type that accepts MvArray and MvArrayOutput values.
// You can construct a concrete instance of `MvArrayInput` via:
//
//	MvArray{ MvArgs{...} }
type MvArrayInput interface {
	pulumi.Input

	ToMvArrayOutput() MvArrayOutput
	ToMvArrayOutputWithContext(context.Context) MvArrayOutput
}

type MvArray []MvInput

func (MvArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mv)(nil)).Elem()
}

func (i MvArray) ToMvArrayOutput() MvArrayOutput {
	return i.ToMvArrayOutputWithContext(context.Background())
}

func (i MvArray) ToMvArrayOutputWithContext(ctx context.Context) MvArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MvArrayOutput)
}

// MvMapInput is an input type that accepts MvMap and MvMapOutput values.
// You can construct a concrete instance of `MvMapInput` via:
//
//	MvMap{ "key": MvArgs{...} }
type MvMapInput interface {
	pulumi.Input

	ToMvMapOutput() MvMapOutput
	ToMvMapOutputWithContext(context.Context) MvMapOutput
}

type MvMap map[string]MvInput

func (MvMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mv)(nil)).Elem()
}

func (i MvMap) ToMvMapOutput() MvMapOutput {
	return i.ToMvMapOutputWithContext(context.Background())
}

func (i MvMap) ToMvMapOutputWithContext(ctx context.Context) MvMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MvMapOutput)
}

type MvOutput struct{ *pulumi.OutputState }

func (MvOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mv)(nil)).Elem()
}

func (o MvOutput) ToMvOutput() MvOutput {
	return o
}

func (o MvOutput) ToMvOutputWithContext(ctx context.Context) MvOutput {
	return o
}

// Corresponds to both the -b and --backup options depending on whether [CONTROL] is supplied.
func (o MvOutput) Backup() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.Backup }).(pulumi.BoolOutput)
}

// Represents the command run on the remote system.
func (o MvOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Mv) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Corresponds to the --context option.
func (o MvOutput) Context() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.Context }).(pulumi.BoolOutput)
}

// Corresponds to the [CONTROL] argument for the --backup option.
func (o MvOutput) Control() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringPtrOutput { return v.Control }).(pulumi.StringPtrOutput)
}

// Corresponds to the [DEST] argument.
func (o MvOutput) Dest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringPtrOutput { return v.Dest }).(pulumi.StringPtrOutput)
}

// Corresponds to the [DIRECTORY] argument.
func (o MvOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringPtrOutput { return v.Directory }).(pulumi.StringPtrOutput)
}

// Corresponds to the --force option.
func (o MvOutput) Force() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.Force }).(pulumi.BoolOutput)
}

// At what stage(s) in the resource lifecycle should the command be run.
func (o MvOutput) Lifecycle() CommandLifecyclePtrOutput {
	return o.ApplyT(func(v *Mv) CommandLifecyclePtrOutput { return v.Lifecycle }).(CommandLifecyclePtrOutput)
}

// Corresponds to the --no-clobber option.
func (o MvOutput) NoClobber() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.NoClobber }).(pulumi.BoolOutput)
}

// Corresponds to the --no-target-directory option.
func (o MvOutput) NoTargetDirectory() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.NoTargetDirectory }).(pulumi.BoolOutput)
}

// Corresponds to the [SOURCE] argument.
func (o MvOutput) Source() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringArrayOutput { return v.Source }).(pulumi.StringArrayOutput)
}

// Corresponds to the --strip-trailing-suffix option.
func (o MvOutput) StripTrailingSlashes() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.StripTrailingSlashes }).(pulumi.BoolOutput)
}

// Corresponds to the --suffix option.
func (o MvOutput) Suffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringPtrOutput { return v.Suffix }).(pulumi.StringPtrOutput)
}

// Corresponds to the --target-directory option.
func (o MvOutput) TargetDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringPtrOutput { return v.TargetDirectory }).(pulumi.StringPtrOutput)
}

// Corresponds to the --update option.
func (o MvOutput) Update() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.Update }).(pulumi.BoolOutput)
}

// Corresponds to the --verbose option.
func (o MvOutput) Verbose() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mv) pulumi.BoolOutput { return v.Verbose }).(pulumi.BoolOutput)
}

type MvArrayOutput struct{ *pulumi.OutputState }

func (MvArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mv)(nil)).Elem()
}

func (o MvArrayOutput) ToMvArrayOutput() MvArrayOutput {
	return o
}

func (o MvArrayOutput) ToMvArrayOutputWithContext(ctx context.Context) MvArrayOutput {
	return o
}

func (o MvArrayOutput) Index(i pulumi.IntInput) MvOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mv {
		return vs[0].([]*Mv)[vs[1].(int)]
	}).(MvOutput)
}

type MvMapOutput struct{ *pulumi.OutputState }

func (MvMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mv)(nil)).Elem()
}

func (o MvMapOutput) ToMvMapOutput() MvMapOutput {
	return o
}

func (o MvMapOutput) ToMvMapOutputWithContext(ctx context.Context) MvMapOutput {
	return o
}

func (o MvMapOutput) MapIndex(k pulumi.StringInput) MvOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mv {
		return vs[0].(map[string]*Mv)[vs[1].(string)]
	}).(MvOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MvInput)(nil)).Elem(), &Mv{})
	pulumi.RegisterInputType(reflect.TypeOf((*MvArrayInput)(nil)).Elem(), MvArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MvMapInput)(nil)).Elem(), MvMap{})
	pulumi.RegisterOutputType(MvOutput{})
	pulumi.RegisterOutputType(MvArrayOutput{})
	pulumi.RegisterOutputType(MvMapOutput{})
}
