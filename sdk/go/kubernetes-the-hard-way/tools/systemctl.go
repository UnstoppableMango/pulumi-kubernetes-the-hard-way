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

// Abstraction over the `systemctl` utility on a remote system.
type Systemctl struct {
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

// NewSystemctl registers a new resource with the given unique name, arguments, and options.
func NewSystemctl(ctx *pulumi.Context,
	name string, args *SystemctlArgs, opts ...pulumi.ResourceOption) (*Systemctl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Systemctl
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Systemctl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type systemctlArgs struct {
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

// The set of arguments for constructing a Systemctl resource.
type SystemctlArgs struct {
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

func (SystemctlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemctlArgs)(nil)).Elem()
}

type SystemctlInput interface {
	pulumi.Input

	ToSystemctlOutput() SystemctlOutput
	ToSystemctlOutputWithContext(ctx context.Context) SystemctlOutput
}

func (*Systemctl) ElementType() reflect.Type {
	return reflect.TypeOf((**Systemctl)(nil)).Elem()
}

func (i *Systemctl) ToSystemctlOutput() SystemctlOutput {
	return i.ToSystemctlOutputWithContext(context.Background())
}

func (i *Systemctl) ToSystemctlOutputWithContext(ctx context.Context) SystemctlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemctlOutput)
}

// SystemctlArrayInput is an input type that accepts SystemctlArray and SystemctlArrayOutput values.
// You can construct a concrete instance of `SystemctlArrayInput` via:
//
//	SystemctlArray{ SystemctlArgs{...} }
type SystemctlArrayInput interface {
	pulumi.Input

	ToSystemctlArrayOutput() SystemctlArrayOutput
	ToSystemctlArrayOutputWithContext(context.Context) SystemctlArrayOutput
}

type SystemctlArray []SystemctlInput

func (SystemctlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Systemctl)(nil)).Elem()
}

func (i SystemctlArray) ToSystemctlArrayOutput() SystemctlArrayOutput {
	return i.ToSystemctlArrayOutputWithContext(context.Background())
}

func (i SystemctlArray) ToSystemctlArrayOutputWithContext(ctx context.Context) SystemctlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemctlArrayOutput)
}

// SystemctlMapInput is an input type that accepts SystemctlMap and SystemctlMapOutput values.
// You can construct a concrete instance of `SystemctlMapInput` via:
//
//	SystemctlMap{ "key": SystemctlArgs{...} }
type SystemctlMapInput interface {
	pulumi.Input

	ToSystemctlMapOutput() SystemctlMapOutput
	ToSystemctlMapOutputWithContext(context.Context) SystemctlMapOutput
}

type SystemctlMap map[string]SystemctlInput

func (SystemctlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Systemctl)(nil)).Elem()
}

func (i SystemctlMap) ToSystemctlMapOutput() SystemctlMapOutput {
	return i.ToSystemctlMapOutputWithContext(context.Background())
}

func (i SystemctlMap) ToSystemctlMapOutputWithContext(ctx context.Context) SystemctlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemctlMapOutput)
}

type SystemctlOutput struct{ *pulumi.OutputState }

func (SystemctlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Systemctl)(nil)).Elem()
}

func (o SystemctlOutput) ToSystemctlOutput() SystemctlOutput {
	return o
}

func (o SystemctlOutput) ToSystemctlOutputWithContext(ctx context.Context) SystemctlOutput {
	return o
}

// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
func (o SystemctlOutput) BinaryPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.StringOutput { return v.BinaryPath }).(pulumi.StringOutput)
}

// The underlying command
func (o SystemctlOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Systemctl) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o SystemctlOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Systemctl) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The command to run on create.
func (o SystemctlOutput) Create() pulumi.AnyOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.AnyOutput { return v.Create }).(pulumi.AnyOutput)
}

// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
// Command resource from previous create or update steps.
func (o SystemctlOutput) Delete() pulumi.AnyOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.AnyOutput { return v.Delete }).(pulumi.AnyOutput)
}

// Environment variables
func (o SystemctlOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
}

// TODO
func (o SystemctlOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// TODO
func (o SystemctlOutput) Stdin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.StringPtrOutput { return v.Stdin }).(pulumi.StringPtrOutput)
}

// TODO
func (o SystemctlOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// TODO
func (o SystemctlOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

// The command to run on update, if empty, create will
// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
// are set to the stdout and stderr properties of the Command resource from previous
// create or update steps.
func (o SystemctlOutput) Update() pulumi.AnyOutput {
	return o.ApplyT(func(v *Systemctl) pulumi.AnyOutput { return v.Update }).(pulumi.AnyOutput)
}

type SystemctlArrayOutput struct{ *pulumi.OutputState }

func (SystemctlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Systemctl)(nil)).Elem()
}

func (o SystemctlArrayOutput) ToSystemctlArrayOutput() SystemctlArrayOutput {
	return o
}

func (o SystemctlArrayOutput) ToSystemctlArrayOutputWithContext(ctx context.Context) SystemctlArrayOutput {
	return o
}

func (o SystemctlArrayOutput) Index(i pulumi.IntInput) SystemctlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Systemctl {
		return vs[0].([]*Systemctl)[vs[1].(int)]
	}).(SystemctlOutput)
}

type SystemctlMapOutput struct{ *pulumi.OutputState }

func (SystemctlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Systemctl)(nil)).Elem()
}

func (o SystemctlMapOutput) ToSystemctlMapOutput() SystemctlMapOutput {
	return o
}

func (o SystemctlMapOutput) ToSystemctlMapOutputWithContext(ctx context.Context) SystemctlMapOutput {
	return o
}

func (o SystemctlMapOutput) MapIndex(k pulumi.StringInput) SystemctlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Systemctl {
		return vs[0].(map[string]*Systemctl)[vs[1].(string)]
	}).(SystemctlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemctlInput)(nil)).Elem(), &Systemctl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemctlArrayInput)(nil)).Elem(), SystemctlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemctlMapInput)(nil)).Elem(), SystemctlMap{})
	pulumi.RegisterOutputType(SystemctlOutput{})
	pulumi.RegisterOutputType(SystemctlArrayOutput{})
	pulumi.RegisterOutputType(SystemctlMapOutput{})
}
