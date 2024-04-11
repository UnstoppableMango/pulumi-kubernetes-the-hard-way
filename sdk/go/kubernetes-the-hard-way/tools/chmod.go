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

// Abstraction over the `chmod` utility on a remote system.
type Chmod struct {
	pulumi.ResourceState

	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringOutput `pulumi:"binaryPath"`
	// Like verbose but report only when a change is made.
	Changes pulumi.BoolOutput `pulumi:"changes"`
	// The underlying command
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Environment variables
	Environment pulumi.StringMapOutput `pulumi:"environment"`
	// Corresponds to the [FILE] argument.
	Files pulumi.AnyOutput `pulumi:"files"`
	// Display help and exit.
	Help pulumi.BoolOutput `pulumi:"help"`
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle CommandLifecyclePtrOutput `pulumi:"lifecycle"`
	// Modes may be absolute or symbolic. An absolute mode is an octal number...
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Do not trea '/' spcially (the default).
	NoPreserveRoot pulumi.BoolOutput `pulumi:"noPreserveRoot"`
	// Fail to operate recursively on '/'.
	PreserveRoot pulumi.BoolOutput `pulumi:"preserveRoot"`
	// Suppress most error messages. Same as `silent`.
	Quiet pulumi.BoolOutput `pulumi:"quiet"`
	// Change files and directories recursively.
	Recursive pulumi.BoolOutput `pulumi:"recursive"`
	// Use RFILE's mode instead of specifying MODE values. RFILE is always dereferenced if a symbolic link.
	Reference pulumi.StringPtrOutput `pulumi:"reference"`
	// Suppress most error messages. Same as `quiet`.
	Silent pulumi.BoolOutput `pulumi:"silent"`
	// TODO
	Stderr pulumi.StringOutput `pulumi:"stderr"`
	// TODO
	Stdin pulumi.StringPtrOutput `pulumi:"stdin"`
	// TODO
	Stdout pulumi.StringOutput `pulumi:"stdout"`
	// TODO
	Triggers pulumi.ArrayOutput `pulumi:"triggers"`
	// Output version information and exit.
	Version pulumi.BoolOutput `pulumi:"version"`
}

// NewChmod registers a new resource with the given unique name, arguments, and options.
func NewChmod(ctx *pulumi.Context,
	name string, args *ChmodArgs, opts ...pulumi.ResourceOption) (*Chmod, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Files == nil {
		return nil, errors.New("invalid value for required argument 'Files'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Chmod
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Chmod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type chmodArgs struct {
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath *string `pulumi:"binaryPath"`
	// Like verbose but report only when a change is made.
	Changes *bool `pulumi:"changes"`
	// Connection details for the remote system
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Environment variables
	Environment map[string]string `pulumi:"environment"`
	// Corresponds to the [FILE] argument.
	Files interface{} `pulumi:"files"`
	// Display help and exit.
	Help *bool `pulumi:"help"`
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle *CommandLifecycle `pulumi:"lifecycle"`
	// Modes may be absolute or symbolic. An absolute mode is an octal number...
	Mode string `pulumi:"mode"`
	// Do not trea '/' spcially (the default).
	NoPreserveRoot *bool `pulumi:"noPreserveRoot"`
	// Fail to operate recursively on '/'.
	PreserveRoot *bool `pulumi:"preserveRoot"`
	// Suppress most error messages. Same as `silent`.
	Quiet *bool `pulumi:"quiet"`
	// Change files and directories recursively.
	Recursive *bool `pulumi:"recursive"`
	// Use RFILE's mode instead of specifying MODE values. RFILE is always dereferenced if a symbolic link.
	Reference *string `pulumi:"reference"`
	// Suppress most error messages. Same as `quiet`.
	Silent *bool `pulumi:"silent"`
	// TODO
	Stdin *string `pulumi:"stdin"`
	// TODO
	Triggers []interface{} `pulumi:"triggers"`
	// Output version information and exit.
	Version *bool `pulumi:"version"`
}

// The set of arguments for constructing a Chmod resource.
type ChmodArgs struct {
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringPtrInput
	// Like verbose but report only when a change is made.
	Changes pulumi.BoolPtrInput
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionInput
	// Environment variables
	Environment pulumi.StringMapInput
	// Corresponds to the [FILE] argument.
	Files pulumi.Input
	// Display help and exit.
	Help pulumi.BoolPtrInput
	// At what stage(s) in the resource lifecycle should the command be run
	Lifecycle *CommandLifecycle
	// Modes may be absolute or symbolic. An absolute mode is an octal number...
	Mode pulumi.StringInput
	// Do not trea '/' spcially (the default).
	NoPreserveRoot pulumi.BoolPtrInput
	// Fail to operate recursively on '/'.
	PreserveRoot pulumi.BoolPtrInput
	// Suppress most error messages. Same as `silent`.
	Quiet pulumi.BoolPtrInput
	// Change files and directories recursively.
	Recursive pulumi.BoolPtrInput
	// Use RFILE's mode instead of specifying MODE values. RFILE is always dereferenced if a symbolic link.
	Reference pulumi.StringPtrInput
	// Suppress most error messages. Same as `quiet`.
	Silent pulumi.BoolPtrInput
	// TODO
	Stdin pulumi.StringPtrInput
	// TODO
	Triggers pulumi.ArrayInput
	// Output version information and exit.
	Version pulumi.BoolPtrInput
}

func (ChmodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chmodArgs)(nil)).Elem()
}

type ChmodInput interface {
	pulumi.Input

	ToChmodOutput() ChmodOutput
	ToChmodOutputWithContext(ctx context.Context) ChmodOutput
}

func (*Chmod) ElementType() reflect.Type {
	return reflect.TypeOf((**Chmod)(nil)).Elem()
}

func (i *Chmod) ToChmodOutput() ChmodOutput {
	return i.ToChmodOutputWithContext(context.Background())
}

func (i *Chmod) ToChmodOutputWithContext(ctx context.Context) ChmodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChmodOutput)
}

// ChmodArrayInput is an input type that accepts ChmodArray and ChmodArrayOutput values.
// You can construct a concrete instance of `ChmodArrayInput` via:
//
//	ChmodArray{ ChmodArgs{...} }
type ChmodArrayInput interface {
	pulumi.Input

	ToChmodArrayOutput() ChmodArrayOutput
	ToChmodArrayOutputWithContext(context.Context) ChmodArrayOutput
}

type ChmodArray []ChmodInput

func (ChmodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Chmod)(nil)).Elem()
}

func (i ChmodArray) ToChmodArrayOutput() ChmodArrayOutput {
	return i.ToChmodArrayOutputWithContext(context.Background())
}

func (i ChmodArray) ToChmodArrayOutputWithContext(ctx context.Context) ChmodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChmodArrayOutput)
}

// ChmodMapInput is an input type that accepts ChmodMap and ChmodMapOutput values.
// You can construct a concrete instance of `ChmodMapInput` via:
//
//	ChmodMap{ "key": ChmodArgs{...} }
type ChmodMapInput interface {
	pulumi.Input

	ToChmodMapOutput() ChmodMapOutput
	ToChmodMapOutputWithContext(context.Context) ChmodMapOutput
}

type ChmodMap map[string]ChmodInput

func (ChmodMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Chmod)(nil)).Elem()
}

func (i ChmodMap) ToChmodMapOutput() ChmodMapOutput {
	return i.ToChmodMapOutputWithContext(context.Background())
}

func (i ChmodMap) ToChmodMapOutputWithContext(ctx context.Context) ChmodMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChmodMapOutput)
}

type ChmodOutput struct{ *pulumi.OutputState }

func (ChmodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Chmod)(nil)).Elem()
}

func (o ChmodOutput) ToChmodOutput() ChmodOutput {
	return o
}

func (o ChmodOutput) ToChmodOutputWithContext(ctx context.Context) ChmodOutput {
	return o
}

// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
func (o ChmodOutput) BinaryPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringOutput { return v.BinaryPath }).(pulumi.StringOutput)
}

// Like verbose but report only when a change is made.
func (o ChmodOutput) Changes() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.Changes }).(pulumi.BoolOutput)
}

// The underlying command
func (o ChmodOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Chmod) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o ChmodOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Chmod) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Environment variables
func (o ChmodOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
}

// Corresponds to the [FILE] argument.
func (o ChmodOutput) Files() pulumi.AnyOutput {
	return o.ApplyT(func(v *Chmod) pulumi.AnyOutput { return v.Files }).(pulumi.AnyOutput)
}

// Display help and exit.
func (o ChmodOutput) Help() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.Help }).(pulumi.BoolOutput)
}

// At what stage(s) in the resource lifecycle should the command be run
func (o ChmodOutput) Lifecycle() CommandLifecyclePtrOutput {
	return o.ApplyT(func(v *Chmod) CommandLifecyclePtrOutput { return v.Lifecycle }).(CommandLifecyclePtrOutput)
}

// Modes may be absolute or symbolic. An absolute mode is an octal number...
func (o ChmodOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Do not trea '/' spcially (the default).
func (o ChmodOutput) NoPreserveRoot() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.NoPreserveRoot }).(pulumi.BoolOutput)
}

// Fail to operate recursively on '/'.
func (o ChmodOutput) PreserveRoot() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.PreserveRoot }).(pulumi.BoolOutput)
}

// Suppress most error messages. Same as `silent`.
func (o ChmodOutput) Quiet() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.Quiet }).(pulumi.BoolOutput)
}

// Change files and directories recursively.
func (o ChmodOutput) Recursive() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.Recursive }).(pulumi.BoolOutput)
}

// Use RFILE's mode instead of specifying MODE values. RFILE is always dereferenced if a symbolic link.
func (o ChmodOutput) Reference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringPtrOutput { return v.Reference }).(pulumi.StringPtrOutput)
}

// Suppress most error messages. Same as `quiet`.
func (o ChmodOutput) Silent() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.Silent }).(pulumi.BoolOutput)
}

// TODO
func (o ChmodOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// TODO
func (o ChmodOutput) Stdin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringPtrOutput { return v.Stdin }).(pulumi.StringPtrOutput)
}

// TODO
func (o ChmodOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// TODO
func (o ChmodOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Chmod) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

// Output version information and exit.
func (o ChmodOutput) Version() pulumi.BoolOutput {
	return o.ApplyT(func(v *Chmod) pulumi.BoolOutput { return v.Version }).(pulumi.BoolOutput)
}

type ChmodArrayOutput struct{ *pulumi.OutputState }

func (ChmodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Chmod)(nil)).Elem()
}

func (o ChmodArrayOutput) ToChmodArrayOutput() ChmodArrayOutput {
	return o
}

func (o ChmodArrayOutput) ToChmodArrayOutputWithContext(ctx context.Context) ChmodArrayOutput {
	return o
}

func (o ChmodArrayOutput) Index(i pulumi.IntInput) ChmodOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Chmod {
		return vs[0].([]*Chmod)[vs[1].(int)]
	}).(ChmodOutput)
}

type ChmodMapOutput struct{ *pulumi.OutputState }

func (ChmodMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Chmod)(nil)).Elem()
}

func (o ChmodMapOutput) ToChmodMapOutput() ChmodMapOutput {
	return o
}

func (o ChmodMapOutput) ToChmodMapOutputWithContext(ctx context.Context) ChmodMapOutput {
	return o
}

func (o ChmodMapOutput) MapIndex(k pulumi.StringInput) ChmodOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Chmod {
		return vs[0].(map[string]*Chmod)[vs[1].(string)]
	}).(ChmodOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChmodInput)(nil)).Elem(), &Chmod{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChmodArrayInput)(nil)).Elem(), ChmodArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChmodMapInput)(nil)).Elem(), ChmodMap{})
	pulumi.RegisterOutputType(ChmodOutput{})
	pulumi.RegisterOutputType(ChmodArrayOutput{})
	pulumi.RegisterOutputType(ChmodMapOutput{})
}
