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

// Abstracion over the `mktemp` utility on a remote system.
type Mktemp struct {
	pulumi.ResourceState

	// Represents the remote `tar` operation.
	Command pulumiCommand.CommandOutput `pulumi:"command"`
	// Corresponds to the --directory option.
	Directory pulumi.BoolOutput `pulumi:"directory"`
	// Corresponds to the --dry-run option.
	DryRun pulumi.BoolOutput `pulumi:"dryRun"`
	// Corresponds to the --quiet option.
	Quiet pulumi.BoolOutput `pulumi:"quiet"`
	// Corresponds to the --suffix option.
	Suffix pulumi.StringPtrOutput `pulumi:"suffix"`
	// Corresponds to the [TEMPLATE] arg.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// Corresponds to the --tmpdir option.
	Tmpdir pulumi.StringPtrOutput `pulumi:"tmpdir"`
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
	// Connection details for the remote system.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Corresponds to the --directory option.
	Directory *bool `pulumi:"directory"`
	// Corresponds to the --dry-run option.
	DryRun *bool `pulumi:"dryRun"`
	// Corresponds to the --quiet option.
	Quiet *bool `pulumi:"quiet"`
	// Corresponds to the --suffix option.
	Suffix *string `pulumi:"suffix"`
	// Corresponds to the [TEMPLATE] arg.
	Template *string `pulumi:"template"`
	// Corresponds to the --tmpdir option.
	Tmpdir *string `pulumi:"tmpdir"`
}

// The set of arguments for constructing a Mktemp resource.
type MktempArgs struct {
	// Connection details for the remote system.
	Connection pulumiCommand.ConnectionInput
	// Corresponds to the --directory option.
	Directory pulumi.BoolPtrInput
	// Corresponds to the --dry-run option.
	DryRun pulumi.BoolPtrInput
	// Corresponds to the --quiet option.
	Quiet pulumi.BoolPtrInput
	// Corresponds to the --suffix option.
	Suffix pulumi.StringPtrInput
	// Corresponds to the [TEMPLATE] arg.
	Template pulumi.StringPtrInput
	// Corresponds to the --tmpdir option.
	Tmpdir pulumi.StringPtrInput
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

// Represents the remote `tar` operation.
func (o MktempOutput) Command() pulumiCommand.CommandOutput {
	return o.ApplyT(func(v *Mktemp) pulumiCommand.CommandOutput { return v.Command }).(pulumiCommand.CommandOutput)
}

// Corresponds to the --directory option.
func (o MktempOutput) Directory() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.BoolOutput { return v.Directory }).(pulumi.BoolOutput)
}

// Corresponds to the --dry-run option.
func (o MktempOutput) DryRun() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.BoolOutput { return v.DryRun }).(pulumi.BoolOutput)
}

// Corresponds to the --quiet option.
func (o MktempOutput) Quiet() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.BoolOutput { return v.Quiet }).(pulumi.BoolOutput)
}

// Corresponds to the --suffix option.
func (o MktempOutput) Suffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringPtrOutput { return v.Suffix }).(pulumi.StringPtrOutput)
}

// Corresponds to the [TEMPLATE] arg.
func (o MktempOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

// Corresponds to the --tmpdir option.
func (o MktempOutput) Tmpdir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Mktemp) pulumi.StringPtrOutput { return v.Tmpdir }).(pulumi.StringPtrOutput)
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
