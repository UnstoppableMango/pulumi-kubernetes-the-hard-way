// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tools

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/internal"
)

// Abstraction over the `mkdir` utility on a remote system.
type Mktemp struct {
	pulumi.ResourceState

	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringOutput `pulumi:"binaryPath"`
	// The underlying command
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Corresponds to the `--directory` option.
	Directory pulumi.BoolPtrOutput `pulumi:"directory"`
	// Corresponds to the `--dry-run` option.
	DryRun pulumi.BoolOutput `pulumi:"dryRun"`
	// Environment variables
	Environment pulumi.StringMapOutput `pulumi:"environment"`
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle CommandLifecyclePtrOutput `pulumi:"lifecycle"`
	// Corresponds to the `--quiet` option.
	Quiet pulumi.BoolOutput `pulumi:"quiet"`
	// TODO
	Stderr pulumi.StringOutput `pulumi:"stderr"`
	// TODO
	Stdin pulumi.StringPtrOutput `pulumi:"stdin"`
	// TODO
	Stdout pulumi.StringOutput `pulumi:"stdout"`
	// Corresponds to the `--suffix` option.
	Suffix pulumi.StringPtrOutput `pulumi:"suffix"`
	// Corresponds to the [TEMPLATE] argument.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// Corresponds to the `--tmpdir` option.
	Tmpdir pulumi.StringPtrOutput `pulumi:"tmpdir"`
	// TODO
	Triggers pulumi.ArrayOutput `pulumi:"triggers"`
}

// NewMktemp registers a new resource with the given unique name, arguments, and options.
func NewMktemp(ctx *pulumi.Context,
	name string, args *MktempArgs, opts ...pulumi.ResourceOption) (*Mktemp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mktemp
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Mktemp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type mktempArgs struct {
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath *string `pulumi:"binaryPath"`
	// Connection details for the remote system
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Corresponds to the `--directory` option.
	Directory *bool `pulumi:"directory"`
	// Corresponds to the `--dry-run` option.
	DryRun *bool `pulumi:"dryRun"`
	// Environment variables
	Environment map[string]string `pulumi:"environment"`
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle *CommandLifecycle `pulumi:"lifecycle"`
	// Corresponds to the `--quiet` option.
	Quiet *bool `pulumi:"quiet"`
	// TODO
	Stdin *string `pulumi:"stdin"`
	// Corresponds to the `--suffix` option.
	Suffix *string `pulumi:"suffix"`
	// Corresponds to the [TEMPLATE] argument.
	Template *string `pulumi:"template"`
	// Corresponds to the `--tmpdir` option.
	Tmpdir *string `pulumi:"tmpdir"`
	// TODO
	Triggers []interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a Mktemp resource.
type MktempArgs struct {
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringPtrInput
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionInput
	// Corresponds to the `--directory` option.
	Directory pulumi.BoolPtrInput
	// Corresponds to the `--dry-run` option.
	DryRun pulumi.BoolPtrInput
	// Environment variables
	Environment pulumi.StringMapInput
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle *CommandLifecycle
	// Corresponds to the `--quiet` option.
	Quiet pulumi.BoolPtrInput
	// TODO
	Stdin pulumi.StringPtrInput
	// Corresponds to the `--suffix` option.
	Suffix pulumi.StringPtrInput
	// Corresponds to the [TEMPLATE] argument.
	Template pulumi.StringPtrInput
	// Corresponds to the `--tmpdir` option.
	Tmpdir pulumi.StringPtrInput
	// TODO
	Triggers pulumi.ArrayInput
}

func (MktempArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mktempArgs)(nil)).Elem()
}

type MktempInput interface {
	pulumi.Input

	ToMktempOutput() MktempOutput
	ToMktempOutputWithContext(ctx context.Context) MktempOutput
}

func (*Mktemp) ElementType() reflect.Type {
	return reflect.TypeOf((**Mktemp)(nil)).Elem()
}

func (i *Mktemp) ToMktempOutput() MktempOutput {
	return i.ToMktempOutputWithContext(context.Background())
}

func (i *Mktemp) ToMktempOutputWithContext(ctx context.Context) MktempOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MktempOutput)
}

// MktempArrayInput is an input type that accepts MktempArray and MktempArrayOutput values.
// You can construct a concrete instance of `MktempArrayInput` via:
//
//	MktempArray{ MktempArgs{...} }
type MktempArrayInput interface {
	pulumi.Input

	ToMktempArrayOutput() MktempArrayOutput
	ToMktempArrayOutputWithContext(context.Context) MktempArrayOutput
}

type MktempArray []MktempInput

func (MktempArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mktemp)(nil)).Elem()
}

func (i MktempArray) ToMktempArrayOutput() MktempArrayOutput {
	return i.ToMktempArrayOutputWithContext(context.Background())
}

func (i MktempArray) ToMktempArrayOutputWithContext(ctx context.Context) MktempArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MktempArrayOutput)
}

// MktempMapInput is an input type that accepts MktempMap and MktempMapOutput values.
// You can construct a concrete instance of `MktempMapInput` via:
//
//	MktempMap{ "key": MktempArgs{...} }
type MktempMapInput interface {
	pulumi.Input

	ToMktempMapOutput() MktempMapOutput
	ToMktempMapOutputWithContext(context.Context) MktempMapOutput
}

type MktempMap map[string]MktempInput

func (MktempMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mktemp)(nil)).Elem()
}

func (i MktempMap) ToMktempMapOutput() MktempMapOutput {
	return i.ToMktempMapOutputWithContext(context.Background())
}

func (i MktempMap) ToMktempMapOutputWithContext(ctx context.Context) MktempMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MktempMapOutput)
}

type MktempOutput struct{ *pulumi.OutputState }

func (MktempOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mktemp)(nil)).Elem()
}

func (o MktempOutput) ToMktempOutput() MktempOutput {
	return o
}

func (o MktempOutput) ToMktempOutputWithContext(ctx context.Context) MktempOutput {
	return o
}

// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
func (o MktempOutput) BinaryPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringOutput { return v.BinaryPath }).(pulumi.StringOutput)
}

// The underlying command
func (o MktempOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Mktemp) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o MktempOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Mktemp) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Corresponds to the `--directory` option.
func (o MktempOutput) Directory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.BoolPtrOutput { return v.Directory }).(pulumi.BoolPtrOutput)
}

// Corresponds to the `--dry-run` option.
func (o MktempOutput) DryRun() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.BoolOutput { return v.DryRun }).(pulumi.BoolOutput)
}

// Environment variables
func (o MktempOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
}

// At what stage(s) in the resource lifecycle should the command be run
func (o MktempOutput) Lifecycle() CommandLifecyclePtrOutput {
	return o.ApplyT(func(v *Mktemp) CommandLifecyclePtrOutput { return v.Lifecycle }).(CommandLifecyclePtrOutput)
}

// Corresponds to the `--quiet` option.
func (o MktempOutput) Quiet() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.BoolOutput { return v.Quiet }).(pulumi.BoolOutput)
}

// TODO
func (o MktempOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// TODO
func (o MktempOutput) Stdin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringPtrOutput { return v.Stdin }).(pulumi.StringPtrOutput)
}

// TODO
func (o MktempOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// Corresponds to the `--suffix` option.
func (o MktempOutput) Suffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringPtrOutput { return v.Suffix }).(pulumi.StringPtrOutput)
}

// Corresponds to the [TEMPLATE] argument.
func (o MktempOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

// Corresponds to the `--tmpdir` option.
func (o MktempOutput) Tmpdir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringPtrOutput { return v.Tmpdir }).(pulumi.StringPtrOutput)
}

// TODO
func (o MktempOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

type MktempArrayOutput struct{ *pulumi.OutputState }

func (MktempArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mktemp)(nil)).Elem()
}

func (o MktempArrayOutput) ToMktempArrayOutput() MktempArrayOutput {
	return o
}

func (o MktempArrayOutput) ToMktempArrayOutputWithContext(ctx context.Context) MktempArrayOutput {
	return o
}

func (o MktempArrayOutput) Index(i pulumi.IntInput) MktempOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mktemp {
		return vs[0].([]*Mktemp)[vs[1].(int)]
	}).(MktempOutput)
}

type MktempMapOutput struct{ *pulumi.OutputState }

func (MktempMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mktemp)(nil)).Elem()
}

func (o MktempMapOutput) ToMktempMapOutput() MktempMapOutput {
	return o
}

func (o MktempMapOutput) ToMktempMapOutputWithContext(ctx context.Context) MktempMapOutput {
	return o
}

func (o MktempMapOutput) MapIndex(k pulumi.StringInput) MktempOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mktemp {
		return vs[0].(map[string]*Mktemp)[vs[1].(string)]
	}).(MktempOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MktempInput)(nil)).Elem(), &Mktemp{})
	pulumi.RegisterInputType(reflect.TypeOf((*MktempArrayInput)(nil)).Elem(), MktempArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MktempMapInput)(nil)).Elem(), MktempMap{})
	pulumi.RegisterOutputType(MktempOutput{})
	pulumi.RegisterOutputType(MktempArrayOutput{})
	pulumi.RegisterOutputType(MktempMapOutput{})
}
