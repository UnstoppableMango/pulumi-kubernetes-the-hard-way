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
	// TODO
	Triggers pulumi.ArrayOutput `pulumi:"triggers"`
	// The command to run on update, if empty, create will
	// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
	// are set to the stdout and stderr properties of the Command resource from previous
	// create or update steps.
	Update pulumi.AnyOutput `pulumi:"update"`
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
	// TODO
	Triggers []interface{} `pulumi:"triggers"`
	// The command to run on update, if empty, create will
	// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
	// are set to the stdout and stderr properties of the Command resource from previous
	// create or update steps.
	Update interface{} `pulumi:"update"`
}

// The set of arguments for constructing a Chmod resource.
type ChmodArgs struct {
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
	// TODO
	Triggers pulumi.ArrayInput
	// The command to run on update, if empty, create will
	// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
	// are set to the stdout and stderr properties of the Command resource from previous
	// create or update steps.
	Update interface{}
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

// The underlying command
func (o ChmodOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Chmod) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o ChmodOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Chmod) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The command to run on create.
func (o ChmodOutput) Create() pulumi.AnyOutput {
	return o.ApplyT(func(v *Chmod) pulumi.AnyOutput { return v.Create }).(pulumi.AnyOutput)
}

// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
// Command resource from previous create or update steps.
func (o ChmodOutput) Delete() pulumi.AnyOutput {
	return o.ApplyT(func(v *Chmod) pulumi.AnyOutput { return v.Delete }).(pulumi.AnyOutput)
}

// Environment variables
func (o ChmodOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Chmod) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
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

// The command to run on update, if empty, create will
// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
// are set to the stdout and stderr properties of the Command resource from previous
// create or update steps.
func (o ChmodOutput) Update() pulumi.AnyOutput {
	return o.ApplyT(func(v *Chmod) pulumi.AnyOutput { return v.Update }).(pulumi.AnyOutput)
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
