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

// Abstraction over the `sed` utility on a remote system.
type Sed struct {
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

// NewSed registers a new resource with the given unique name, arguments, and options.
func NewSed(ctx *pulumi.Context,
	name string, args *SedArgs, opts ...pulumi.ResourceOption) (*Sed, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sed
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Sed", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type sedArgs struct {
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

// The set of arguments for constructing a Sed resource.
type SedArgs struct {
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

func (SedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sedArgs)(nil)).Elem()
}

type SedInput interface {
	pulumi.Input

	ToSedOutput() SedOutput
	ToSedOutputWithContext(ctx context.Context) SedOutput
}

func (*Sed) ElementType() reflect.Type {
	return reflect.TypeOf((**Sed)(nil)).Elem()
}

func (i *Sed) ToSedOutput() SedOutput {
	return i.ToSedOutputWithContext(context.Background())
}

func (i *Sed) ToSedOutputWithContext(ctx context.Context) SedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SedOutput)
}

// SedArrayInput is an input type that accepts SedArray and SedArrayOutput values.
// You can construct a concrete instance of `SedArrayInput` via:
//
//	SedArray{ SedArgs{...} }
type SedArrayInput interface {
	pulumi.Input

	ToSedArrayOutput() SedArrayOutput
	ToSedArrayOutputWithContext(context.Context) SedArrayOutput
}

type SedArray []SedInput

func (SedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sed)(nil)).Elem()
}

func (i SedArray) ToSedArrayOutput() SedArrayOutput {
	return i.ToSedArrayOutputWithContext(context.Background())
}

func (i SedArray) ToSedArrayOutputWithContext(ctx context.Context) SedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SedArrayOutput)
}

// SedMapInput is an input type that accepts SedMap and SedMapOutput values.
// You can construct a concrete instance of `SedMapInput` via:
//
//	SedMap{ "key": SedArgs{...} }
type SedMapInput interface {
	pulumi.Input

	ToSedMapOutput() SedMapOutput
	ToSedMapOutputWithContext(context.Context) SedMapOutput
}

type SedMap map[string]SedInput

func (SedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sed)(nil)).Elem()
}

func (i SedMap) ToSedMapOutput() SedMapOutput {
	return i.ToSedMapOutputWithContext(context.Background())
}

func (i SedMap) ToSedMapOutputWithContext(ctx context.Context) SedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SedMapOutput)
}

type SedOutput struct{ *pulumi.OutputState }

func (SedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sed)(nil)).Elem()
}

func (o SedOutput) ToSedOutput() SedOutput {
	return o
}

func (o SedOutput) ToSedOutputWithContext(ctx context.Context) SedOutput {
	return o
}

// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
func (o SedOutput) BinaryPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Sed) pulumi.StringOutput { return v.BinaryPath }).(pulumi.StringOutput)
}

// The underlying command
func (o SedOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Sed) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o SedOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Sed) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The command to run on create.
func (o SedOutput) Create() pulumi.AnyOutput {
	return o.ApplyT(func(v *Sed) pulumi.AnyOutput { return v.Create }).(pulumi.AnyOutput)
}

// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
// Command resource from previous create or update steps.
func (o SedOutput) Delete() pulumi.AnyOutput {
	return o.ApplyT(func(v *Sed) pulumi.AnyOutput { return v.Delete }).(pulumi.AnyOutput)
}

// Environment variables
func (o SedOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Sed) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
}

// TODO
func (o SedOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Sed) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// TODO
func (o SedOutput) Stdin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sed) pulumi.StringPtrOutput { return v.Stdin }).(pulumi.StringPtrOutput)
}

// TODO
func (o SedOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Sed) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// Trigger replacements on changes to this input.
func (o SedOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Sed) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

// The command to run on update, if empty, create will
// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
// are set to the stdout and stderr properties of the Command resource from previous
// create or update steps.
func (o SedOutput) Update() pulumi.AnyOutput {
	return o.ApplyT(func(v *Sed) pulumi.AnyOutput { return v.Update }).(pulumi.AnyOutput)
}

type SedArrayOutput struct{ *pulumi.OutputState }

func (SedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sed)(nil)).Elem()
}

func (o SedArrayOutput) ToSedArrayOutput() SedArrayOutput {
	return o
}

func (o SedArrayOutput) ToSedArrayOutputWithContext(ctx context.Context) SedArrayOutput {
	return o
}

func (o SedArrayOutput) Index(i pulumi.IntInput) SedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sed {
		return vs[0].([]*Sed)[vs[1].(int)]
	}).(SedOutput)
}

type SedMapOutput struct{ *pulumi.OutputState }

func (SedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sed)(nil)).Elem()
}

func (o SedMapOutput) ToSedMapOutput() SedMapOutput {
	return o
}

func (o SedMapOutput) ToSedMapOutputWithContext(ctx context.Context) SedMapOutput {
	return o
}

func (o SedMapOutput) MapIndex(k pulumi.StringInput) SedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sed {
		return vs[0].(map[string]*Sed)[vs[1].(string)]
	}).(SedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SedInput)(nil)).Elem(), &Sed{})
	pulumi.RegisterInputType(reflect.TypeOf((*SedArrayInput)(nil)).Elem(), SedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SedMapInput)(nil)).Elem(), SedMap{})
	pulumi.RegisterOutputType(SedOutput{})
	pulumi.RegisterOutputType(SedArrayOutput{})
	pulumi.RegisterOutputType(SedMapOutput{})
}
