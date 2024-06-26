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

	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringOutput `pulumi:"binaryPath"`
	// The underlying command
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The command to run on create.
	Create pulumi.AnyOutput `pulumi:"create"`
	// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
	// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
	// Command resource from previous create or update steps.
	Delete pulumi.AnyOutput `pulumi:"delete"`
	// Environment variables
	Environment pulumi.StringMapOutput `pulumi:"environment"`
	// TODO
	Stderr pulumi.StringOutput `pulumi:"stderr"`
	// TODO
	Stdin pulumi.StringPtrOutput `pulumi:"stdin"`
	// TODO
	Stdout pulumi.StringOutput `pulumi:"stdout"`
	// Trigger replacements on changes to this input.
	Triggers pulumi.ArrayOutput `pulumi:"triggers"`
	// The command to run on update, if empty, create will
	// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
	// are set to the stdout and stderr properties of the Command resource from previous
	// create or update steps.
	Update pulumi.AnyOutput `pulumi:"update"`
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
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath *string `pulumi:"binaryPath"`
	// Connection details for the remote system
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The command to run on create.
	Create interface{} `pulumi:"create"`
	// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
	// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
	// Command resource from previous create or update steps.
	Delete interface{} `pulumi:"delete"`
	// Environment variables
	Environment map[string]string `pulumi:"environment"`
	// TODO
	Stdin *string `pulumi:"stdin"`
	// Trigger replacements on changes to this input.
	Triggers []interface{} `pulumi:"triggers"`
	// The command to run on update, if empty, create will
	// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
	// are set to the stdout and stderr properties of the Command resource from previous
	// create or update steps.
	Update interface{} `pulumi:"update"`
}

// The set of arguments for constructing a Mv resource.
type MvArgs struct {
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringPtrInput
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionInput
	// The command to run on create.
	Create interface{}
	// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
	// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
	// Command resource from previous create or update steps.
	Delete interface{}
	// Environment variables
	Environment pulumi.StringMapInput
	// TODO
	Stdin pulumi.StringPtrInput
	// Trigger replacements on changes to this input.
	Triggers pulumi.ArrayInput
	// The command to run on update, if empty, create will
	// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
	// are set to the stdout and stderr properties of the Command resource from previous
	// create or update steps.
	Update interface{}
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

// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
func (o MvOutput) BinaryPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringOutput { return v.BinaryPath }).(pulumi.StringOutput)
}

// The underlying command
func (o MvOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Mv) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o MvOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Mv) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The command to run on create.
func (o MvOutput) Create() pulumi.AnyOutput {
	return o.ApplyT(func(v *Mv) pulumi.AnyOutput { return v.Create }).(pulumi.AnyOutput)
}

// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
// Command resource from previous create or update steps.
func (o MvOutput) Delete() pulumi.AnyOutput {
	return o.ApplyT(func(v *Mv) pulumi.AnyOutput { return v.Delete }).(pulumi.AnyOutput)
}

// Environment variables
func (o MvOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
}

// TODO
func (o MvOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// TODO
func (o MvOutput) Stdin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringPtrOutput { return v.Stdin }).(pulumi.StringPtrOutput)
}

// TODO
func (o MvOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Mv) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// Trigger replacements on changes to this input.
func (o MvOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Mv) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

// The command to run on update, if empty, create will
// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
// are set to the stdout and stderr properties of the Command resource from previous
// create or update steps.
func (o MvOutput) Update() pulumi.AnyOutput {
	return o.ApplyT(func(v *Mv) pulumi.AnyOutput { return v.Update }).(pulumi.AnyOutput)
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
