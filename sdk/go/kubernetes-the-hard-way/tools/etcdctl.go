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

// Abstraction over the `etcdctl` utility on a remote system.
type Etcdctl struct {
	pulumi.ResourceState

	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringOutput `pulumi:"binaryPath"`
	// The underlying command
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The command to run on create.
	Create EtcdctlOptsPtrOutput `pulumi:"create"`
	// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
	// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
	// Command resource from previous create or update steps.
	Delete EtcdctlOptsPtrOutput `pulumi:"delete"`
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
	Update EtcdctlOptsPtrOutput `pulumi:"update"`
}

// NewEtcdctl registers a new resource with the given unique name, arguments, and options.
func NewEtcdctl(ctx *pulumi.Context,
	name string, args *EtcdctlArgs, opts ...pulumi.ResourceOption) (*Etcdctl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Etcdctl
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Etcdctl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type etcdctlArgs struct {
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath *string `pulumi:"binaryPath"`
	// Connection details for the remote system
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The command to run on create.
	Create *EtcdctlOpts `pulumi:"create"`
	// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
	// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
	// Command resource from previous create or update steps.
	Delete *EtcdctlOpts `pulumi:"delete"`
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
	Update *EtcdctlOpts `pulumi:"update"`
}

// The set of arguments for constructing a Etcdctl resource.
type EtcdctlArgs struct {
	// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
	BinaryPath pulumi.StringPtrInput
	// Connection details for the remote system
	Connection pulumiCommand.ConnectionInput
	// The command to run on create.
	Create *EtcdctlOptsArgs
	// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
	// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
	// Command resource from previous create or update steps.
	Delete *EtcdctlOptsArgs
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
	Update *EtcdctlOptsArgs
}

func (EtcdctlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*etcdctlArgs)(nil)).Elem()
}

type EtcdctlInput interface {
	pulumi.Input

	ToEtcdctlOutput() EtcdctlOutput
	ToEtcdctlOutputWithContext(ctx context.Context) EtcdctlOutput
}

func (*Etcdctl) ElementType() reflect.Type {
	return reflect.TypeOf((**Etcdctl)(nil)).Elem()
}

func (i *Etcdctl) ToEtcdctlOutput() EtcdctlOutput {
	return i.ToEtcdctlOutputWithContext(context.Background())
}

func (i *Etcdctl) ToEtcdctlOutputWithContext(ctx context.Context) EtcdctlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdctlOutput)
}

// EtcdctlArrayInput is an input type that accepts EtcdctlArray and EtcdctlArrayOutput values.
// You can construct a concrete instance of `EtcdctlArrayInput` via:
//
//	EtcdctlArray{ EtcdctlArgs{...} }
type EtcdctlArrayInput interface {
	pulumi.Input

	ToEtcdctlArrayOutput() EtcdctlArrayOutput
	ToEtcdctlArrayOutputWithContext(context.Context) EtcdctlArrayOutput
}

type EtcdctlArray []EtcdctlInput

func (EtcdctlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Etcdctl)(nil)).Elem()
}

func (i EtcdctlArray) ToEtcdctlArrayOutput() EtcdctlArrayOutput {
	return i.ToEtcdctlArrayOutputWithContext(context.Background())
}

func (i EtcdctlArray) ToEtcdctlArrayOutputWithContext(ctx context.Context) EtcdctlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdctlArrayOutput)
}

// EtcdctlMapInput is an input type that accepts EtcdctlMap and EtcdctlMapOutput values.
// You can construct a concrete instance of `EtcdctlMapInput` via:
//
//	EtcdctlMap{ "key": EtcdctlArgs{...} }
type EtcdctlMapInput interface {
	pulumi.Input

	ToEtcdctlMapOutput() EtcdctlMapOutput
	ToEtcdctlMapOutputWithContext(context.Context) EtcdctlMapOutput
}

type EtcdctlMap map[string]EtcdctlInput

func (EtcdctlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Etcdctl)(nil)).Elem()
}

func (i EtcdctlMap) ToEtcdctlMapOutput() EtcdctlMapOutput {
	return i.ToEtcdctlMapOutputWithContext(context.Background())
}

func (i EtcdctlMap) ToEtcdctlMapOutputWithContext(ctx context.Context) EtcdctlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdctlMapOutput)
}

type EtcdctlOutput struct{ *pulumi.OutputState }

func (EtcdctlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Etcdctl)(nil)).Elem()
}

func (o EtcdctlOutput) ToEtcdctlOutput() EtcdctlOutput {
	return o
}

func (o EtcdctlOutput) ToEtcdctlOutputWithContext(ctx context.Context) EtcdctlOutput {
	return o
}

// Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH
func (o EtcdctlOutput) BinaryPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Etcdctl) pulumi.StringOutput { return v.BinaryPath }).(pulumi.StringOutput)
}

// The underlying command
func (o EtcdctlOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Etcdctl) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Connection details for the remote system
func (o EtcdctlOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *Etcdctl) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The command to run on create.
func (o EtcdctlOutput) Create() EtcdctlOptsPtrOutput {
	return o.ApplyT(func(v *Etcdctl) EtcdctlOptsPtrOutput { return v.Create }).(EtcdctlOptsPtrOutput)
}

// The command to run on delete. The environment variables PULUMI_COMMAND_STDOUT
// and PULUMI_COMMAND_STDERR are set to the stdout and stderr properties of the
// Command resource from previous create or update steps.
func (o EtcdctlOutput) Delete() EtcdctlOptsPtrOutput {
	return o.ApplyT(func(v *Etcdctl) EtcdctlOptsPtrOutput { return v.Delete }).(EtcdctlOptsPtrOutput)
}

// Environment variables
func (o EtcdctlOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Etcdctl) pulumi.StringMapOutput { return v.Environment }).(pulumi.StringMapOutput)
}

// TODO
func (o EtcdctlOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Etcdctl) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// TODO
func (o EtcdctlOutput) Stdin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Etcdctl) pulumi.StringPtrOutput { return v.Stdin }).(pulumi.StringPtrOutput)
}

// TODO
func (o EtcdctlOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Etcdctl) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// TODO
func (o EtcdctlOutput) Triggers() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Etcdctl) pulumi.ArrayOutput { return v.Triggers }).(pulumi.ArrayOutput)
}

// The command to run on update, if empty, create will
// run again. The environment variables PULUMI_COMMAND_STDOUT and PULUMI_COMMAND_STDERR
// are set to the stdout and stderr properties of the Command resource from previous
// create or update steps.
func (o EtcdctlOutput) Update() EtcdctlOptsPtrOutput {
	return o.ApplyT(func(v *Etcdctl) EtcdctlOptsPtrOutput { return v.Update }).(EtcdctlOptsPtrOutput)
}

type EtcdctlArrayOutput struct{ *pulumi.OutputState }

func (EtcdctlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Etcdctl)(nil)).Elem()
}

func (o EtcdctlArrayOutput) ToEtcdctlArrayOutput() EtcdctlArrayOutput {
	return o
}

func (o EtcdctlArrayOutput) ToEtcdctlArrayOutputWithContext(ctx context.Context) EtcdctlArrayOutput {
	return o
}

func (o EtcdctlArrayOutput) Index(i pulumi.IntInput) EtcdctlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Etcdctl {
		return vs[0].([]*Etcdctl)[vs[1].(int)]
	}).(EtcdctlOutput)
}

type EtcdctlMapOutput struct{ *pulumi.OutputState }

func (EtcdctlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Etcdctl)(nil)).Elem()
}

func (o EtcdctlMapOutput) ToEtcdctlMapOutput() EtcdctlMapOutput {
	return o
}

func (o EtcdctlMapOutput) ToEtcdctlMapOutputWithContext(ctx context.Context) EtcdctlMapOutput {
	return o
}

func (o EtcdctlMapOutput) MapIndex(k pulumi.StringInput) EtcdctlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Etcdctl {
		return vs[0].(map[string]*Etcdctl)[vs[1].(string)]
	}).(EtcdctlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdctlInput)(nil)).Elem(), &Etcdctl{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdctlArrayInput)(nil)).Elem(), EtcdctlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdctlMapInput)(nil)).Elem(), EtcdctlMap{})
	pulumi.RegisterOutputType(EtcdctlOutput{})
	pulumi.RegisterOutputType(EtcdctlArrayOutput{})
	pulumi.RegisterOutputType(EtcdctlMapOutput{})
}
