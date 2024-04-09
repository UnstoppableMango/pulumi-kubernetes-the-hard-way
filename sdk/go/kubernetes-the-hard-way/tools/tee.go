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

// Abstraction over the `rm` utility on a remote system.
type Tee struct {
	pulumi.ResourceState

	// Append to the given FILEs, do not overwrite
	Append pulumi.BoolOutput `pulumi:"append"`
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringOutput `pulumi:"binaryPath"`
	// The underlying command
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Environment variables
	Environment pulumi.StringMapOutput `pulumi:"environment"`
	// Corresponds to the [FILE] argument.
	Files pulumi.AnyOutput `pulumi:"files"`
	// Ignore interrupt signals.
	IgnoreInterrupts pulumi.BoolOutput `pulumi:"ignoreInterrupts"`
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle CommandLifecyclePtrOutput `pulumi:"lifecycle"`
	// Set behavior on write error.
	OutputError TeeModePtrOutput `pulumi:"outputError"`
	// Operate in a more appropriate MODE with pipes.
	Pipe pulumi.BoolOutput `pulumi:"pipe"`
	// TODO
	Stderr pulumi.StringOutput `pulumi:"stderr"`
	// TODO
	Stdin pulumi.StringOutput `pulumi:"stdin"`
	// TODO
	Stdout pulumi.StringOutput `pulumi:"stdout"`
	// TODO
	Triggers pulumi.ArrayOutput `pulumi:"triggers"`
	// Output version information and exit.
	Version pulumi.BoolOutput `pulumi:"version"`
}

// NewTee registers a new resource with the given unique name, arguments, and options.
func NewTee(ctx *pulumi.Context,
	name string, args *TeeArgs, opts ...pulumi.ResourceOption) (*Tee, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Files == nil {
		return nil, errors.New("invalid value for required argument 'Files'")
	}
	if args.Stdin == nil {
		return nil, errors.New("invalid value for required argument 'Stdin'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Tee
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Tee", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type teeArgs struct {
	// Append to the given FILEs, do not overwrite
	Append *bool `pulumi:"append"`
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath *string `pulumi:"binaryPath"`
	// Connection details for the remote system
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Environment variables
	Environment map[string]string `pulumi:"environment"`
	// Corresponds to the [FILE] argument.
	Files interface{} `pulumi:"files"`
	// Ignore interrupt signals.
	IgnoreInterrupts *bool `pulumi:"ignoreInterrupts"`
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle *CommandLifecycle `pulumi:"lifecycle"`
	// Set behavior on write error.
	OutputError *TeeMode `pulumi:"outputError"`
	// Operate in a more appropriate MODE with pipes.
	Pipe *bool `pulumi:"pipe"`
	// TODO
	Stdin string `pulumi:"stdin"`
	// TODO
	Triggers []interface{} `pulumi:"triggers"`
	// Output version information and exit.
	Version *bool `pulumi:"version"`
}

// The set of arguments for constructing a Tee resource.
type TeeArgs struct {
	// Append to the given FILEs, do not overwrite
	Append pulumi.BoolPtrInput
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringPtrInput
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionInput
	// Environment variables
	Environment pulumi.StringMapInput
	// Corresponds to the [FILE] argument.
	Files pulumi.Input
	// Ignore interrupt signals.
	IgnoreInterrupts pulumi.BoolPtrInput
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle *CommandLifecycle
	// Set behavior on write error.
	OutputError TeeModePtrInput
	// Operate in a more appropriate MODE with pipes.
	Pipe pulumi.BoolPtrInput
	// TODO
	Stdin pulumi.StringInput
	// TODO
	Triggers pulumi.ArrayInput
	// Output version information and exit.
	Version pulumi.BoolPtrInput
}

func (TeeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teeArgs)(nil)).Elem()
}

type TeeInput interface {
	pulumi.Input

	ToTeeOutput() TeeOutput
	ToTeeOutputWithContext(ctx context.Context) TeeOutput
}

func (*Tee) ElementType() reflect.Type {
	return reflect.TypeOf((**Tee)(nil)).Elem()
}

func (i *Tee) ToTeeOutput() TeeOutput {
	return i.ToTeeOutputWithContext(context.Background())
}

func (i *Tee) ToTeeOutputWithContext(ctx context.Context) TeeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeeOutput)
}

// TeeArrayInput is an input type that accepts TeeArray and TeeArrayOutput values.
// You can construct a concrete instance of `TeeArrayInput` via:
//
//	TeeArray{ TeeArgs{...} }
type TeeArrayInput interface {
	pulumi.Input

	ToTeeArrayOutput() TeeArrayOutput
	ToTeeArrayOutputWithContext(context.Context) TeeArrayOutput
}

type TeeArray []TeeInput

func (TeeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tee)(nil)).Elem()
}

func (i TeeArray) ToTeeArrayOutput() TeeArrayOutput {
	return i.ToTeeArrayOutputWithContext(context.Background())
}

func (i TeeArray) ToTeeArrayOutputWithContext(ctx context.Context) TeeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeeArrayOutput)
}

// TeeMapInput is an input type that accepts TeeMap and TeeMapOutput values.
// You can construct a concrete instance of `TeeMapInput` via:
//
//	TeeMap{ "key": TeeArgs{...} }
type TeeMapInput interface {
	pulumi.Input

	ToTeeMapOutput() TeeMapOutput
	ToTeeMapOutputWithContext(context.Context) TeeMapOutput
}

type TeeMap map[string]TeeInput

func (TeeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tee)(nil)).Elem()
}

func (i TeeMap) ToTeeMapOutput() TeeMapOutput {
	return i.ToTeeMapOutputWithContext(context.Background())
}

func (i TeeMap) ToTeeMapOutputWithContext(ctx context.Context) TeeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeeMapOutput)
}

type TeeOutput struct{ *pulumi.OutputState }

func (TeeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tee)(nil)).Elem()
}

func (o TeeOutput) ToTeeOutput() TeeOutput {
	return o
}

func (o TeeOutput) ToTeeOutputWithContext(ctx context.Context) TeeOutput {
	return o
}

// Append to the given FILEs, do not overwrite
func (o TeeOutput) Append() pulumi.BoolOutput {
	return o.ApplyT(func(v *Tee) pulumi.BoolOutput { return v.Append }).(pulumi.BoolOutput)
}

// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
func (o TeeOutput) BinaryPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Tee) pulumi.StringOutput { return v.BinaryPath }).(pulumi.StringOutput)
}

// The underlying command
func (o TeeOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Tee) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o TeeOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Tee) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Environment variables
func (o TeeOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Tee) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
}

// Corresponds to the [FILE] argument.
func (o TeeOutput) Files() pulumi.AnyOutput {
	return o.ApplyT(func(v *Tee) pulumi.AnyOutput { return v.Files }).(pulumi.AnyOutput)
}

// Ignore interrupt signals.
func (o TeeOutput) IgnoreInterrupts() pulumi.BoolOutput {
	return o.ApplyT(func(v *Tee) pulumi.BoolOutput { return v.IgnoreInterrupts }).(pulumi.BoolOutput)
}

// At what stage(s) in the resource lifecycle should the command be run
func (o TeeOutput) Lifecycle() CommandLifecyclePtrOutput {
	return o.ApplyT(func(v *Tee) CommandLifecyclePtrOutput { return v.Lifecycle }).(CommandLifecyclePtrOutput)
}

// Set behavior on write error.
func (o TeeOutput) OutputError() TeeModePtrOutput {
	return o.ApplyT(func(v *Tee) TeeModePtrOutput { return v.OutputError }).(TeeModePtrOutput)
}

// Operate in a more appropriate MODE with pipes.
func (o TeeOutput) Pipe() pulumi.BoolOutput {
	return o.ApplyT(func(v *Tee) pulumi.BoolOutput { return v.Pipe }).(pulumi.BoolOutput)
}

// TODO
func (o TeeOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Tee) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// TODO
func (o TeeOutput) Stdin() pulumi.StringOutput {
	return o.ApplyT(func(v *Tee) pulumi.StringOutput { return v.Stdin }).(pulumi.StringOutput)
}

// TODO
func (o TeeOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Tee) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// TODO
func (o TeeOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Tee) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

// Output version information and exit.
func (o TeeOutput) Version() pulumi.BoolOutput {
	return o.ApplyT(func(v *Tee) pulumi.BoolOutput { return v.Version }).(pulumi.BoolOutput)
}

type TeeArrayOutput struct{ *pulumi.OutputState }

func (TeeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tee)(nil)).Elem()
}

func (o TeeArrayOutput) ToTeeArrayOutput() TeeArrayOutput {
	return o
}

func (o TeeArrayOutput) ToTeeArrayOutputWithContext(ctx context.Context) TeeArrayOutput {
	return o
}

func (o TeeArrayOutput) Index(i pulumi.IntInput) TeeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tee {
		return vs[0].([]*Tee)[vs[1].(int)]
	}).(TeeOutput)
}

type TeeMapOutput struct{ *pulumi.OutputState }

func (TeeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tee)(nil)).Elem()
}

func (o TeeMapOutput) ToTeeMapOutput() TeeMapOutput {
	return o
}

func (o TeeMapOutput) ToTeeMapOutputWithContext(ctx context.Context) TeeMapOutput {
	return o
}

func (o TeeMapOutput) MapIndex(k pulumi.StringInput) TeeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tee {
		return vs[0].(map[string]*Tee)[vs[1].(string)]
	}).(TeeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeeInput)(nil)).Elem(), &Tee{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeeArrayInput)(nil)).Elem(), TeeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeeMapInput)(nil)).Elem(), TeeMap{})
	pulumi.RegisterOutputType(TeeOutput{})
	pulumi.RegisterOutputType(TeeArrayOutput{})
	pulumi.RegisterOutputType(TeeMapOutput{})
}
