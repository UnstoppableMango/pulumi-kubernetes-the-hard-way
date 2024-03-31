// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tools

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Abstraction over the `rm` utility on a remote system.
type Rm struct {
	pulumi.ResourceState

	// Corresponds to the --dir option.
	Dir pulumi.StringOutput `pulumi:"dir"`
	// Corresponds to the [FILE] argument.
	Files pulumi.StringArrayOutput `pulumi:"files"`
	// Corresponds to the --force option.
	Force pulumi.BoolOutput `pulumi:"force"`
	// Whether rm should be run when the resource is created or deleted.
	OnDelete pulumi.BoolOutput `pulumi:"onDelete"`
	// Corresponds to the --recursive option.
	Recursive pulumi.BoolOutput `pulumi:"recursive"`
	// Corresponds to the --verbose option.
	Verbose pulumi.BoolOutput `pulumi:"verbose"`
}

// NewRm registers a new resource with the given unique name, arguments, and options.
func NewRm(ctx *pulumi.Context,
	name string, args *RmArgs, opts ...pulumi.ResourceOption) (*Rm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Files == nil {
		return nil, errors.New("invalid value for required argument 'Files'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v remote.Connection) remote.Connection { return *v.Defaults() }).(remote.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rm
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Rm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type rmArgs struct {
	// Connection details for the remote system.
	Connection remote.Connection `pulumi:"connection"`
	// Corresponds to the --dir option.
	Dir *bool `pulumi:"dir"`
	// Corresponds to the [FILE] argument.
	Files interface{} `pulumi:"files"`
	// Corresponds to the --force option.
	Force *bool `pulumi:"force"`
	// Whether rm should be run when the resource is created or deleted.
	OnDelete *bool `pulumi:"onDelete"`
	// Corresponds to the --recursive option.
	Recursive *bool `pulumi:"recursive"`
	// Corresponds to the --verbose option.
	Verbose *bool `pulumi:"verbose"`
}

// The set of arguments for constructing a Rm resource.
type RmArgs struct {
	// Connection details for the remote system.
	Connection remote.ConnectionInput
	// Corresponds to the --dir option.
	Dir pulumi.BoolPtrInput
	// Corresponds to the [FILE] argument.
	Files pulumi.Input
	// Corresponds to the --force option.
	Force pulumi.BoolPtrInput
	// Whether rm should be run when the resource is created or deleted.
	OnDelete *bool
	// Corresponds to the --recursive option.
	Recursive pulumi.BoolPtrInput
	// Corresponds to the --verbose option.
	Verbose pulumi.BoolPtrInput
}

func (RmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rmArgs)(nil)).Elem()
}

type RmInput interface {
	pulumi.Input

	ToRmOutput() RmOutput
	ToRmOutputWithContext(ctx context.Context) RmOutput
}

func (*Rm) ElementType() reflect.Type {
	return reflect.TypeOf((**Rm)(nil)).Elem()
}

func (i *Rm) ToRmOutput() RmOutput {
	return i.ToRmOutputWithContext(context.Background())
}

func (i *Rm) ToRmOutputWithContext(ctx context.Context) RmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RmOutput)
}

// RmArrayInput is an input type that accepts RmArray and RmArrayOutput values.
// You can construct a concrete instance of `RmArrayInput` via:
//
//	RmArray{ RmArgs{...} }
type RmArrayInput interface {
	pulumi.Input

	ToRmArrayOutput() RmArrayOutput
	ToRmArrayOutputWithContext(context.Context) RmArrayOutput
}

type RmArray []RmInput

func (RmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rm)(nil)).Elem()
}

func (i RmArray) ToRmArrayOutput() RmArrayOutput {
	return i.ToRmArrayOutputWithContext(context.Background())
}

func (i RmArray) ToRmArrayOutputWithContext(ctx context.Context) RmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RmArrayOutput)
}

// RmMapInput is an input type that accepts RmMap and RmMapOutput values.
// You can construct a concrete instance of `RmMapInput` via:
//
//	RmMap{ "key": RmArgs{...} }
type RmMapInput interface {
	pulumi.Input

	ToRmMapOutput() RmMapOutput
	ToRmMapOutputWithContext(context.Context) RmMapOutput
}

type RmMap map[string]RmInput

func (RmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rm)(nil)).Elem()
}

func (i RmMap) ToRmMapOutput() RmMapOutput {
	return i.ToRmMapOutputWithContext(context.Background())
}

func (i RmMap) ToRmMapOutputWithContext(ctx context.Context) RmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RmMapOutput)
}

type RmOutput struct{ *pulumi.OutputState }

func (RmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rm)(nil)).Elem()
}

func (o RmOutput) ToRmOutput() RmOutput {
	return o
}

func (o RmOutput) ToRmOutputWithContext(ctx context.Context) RmOutput {
	return o
}

// Corresponds to the --dir option.
func (o RmOutput) Dir() pulumi.StringOutput {
	return o.ApplyT(func(v *Rm) pulumi.StringOutput { return v.Dir }).(pulumi.StringOutput)
}

// Corresponds to the [FILE] argument.
func (o RmOutput) Files() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rm) pulumi.StringArrayOutput { return v.Files }).(pulumi.StringArrayOutput)
}

// Corresponds to the --force option.
func (o RmOutput) Force() pulumi.BoolOutput {
	return o.ApplyT(func(v *Rm) pulumi.BoolOutput { return v.Force }).(pulumi.BoolOutput)
}

// Whether rm should be run when the resource is created or deleted.
func (o RmOutput) OnDelete() pulumi.BoolOutput {
	return o.ApplyT(func(v *Rm) pulumi.BoolOutput { return v.OnDelete }).(pulumi.BoolOutput)
}

// Corresponds to the --recursive option.
func (o RmOutput) Recursive() pulumi.BoolOutput {
	return o.ApplyT(func(v *Rm) pulumi.BoolOutput { return v.Recursive }).(pulumi.BoolOutput)
}

// Corresponds to the --verbose option.
func (o RmOutput) Verbose() pulumi.BoolOutput {
	return o.ApplyT(func(v *Rm) pulumi.BoolOutput { return v.Verbose }).(pulumi.BoolOutput)
}

type RmArrayOutput struct{ *pulumi.OutputState }

func (RmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rm)(nil)).Elem()
}

func (o RmArrayOutput) ToRmArrayOutput() RmArrayOutput {
	return o
}

func (o RmArrayOutput) ToRmArrayOutputWithContext(ctx context.Context) RmArrayOutput {
	return o
}

func (o RmArrayOutput) Index(i pulumi.IntInput) RmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rm {
		return vs[0].([]*Rm)[vs[1].(int)]
	}).(RmOutput)
}

type RmMapOutput struct{ *pulumi.OutputState }

func (RmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rm)(nil)).Elem()
}

func (o RmMapOutput) ToRmMapOutput() RmMapOutput {
	return o
}

func (o RmMapOutput) ToRmMapOutputWithContext(ctx context.Context) RmMapOutput {
	return o
}

func (o RmMapOutput) MapIndex(k pulumi.StringInput) RmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rm {
		return vs[0].(map[string]*Rm)[vs[1].(string)]
	}).(RmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RmInput)(nil)).Elem(), &Rm{})
	pulumi.RegisterInputType(reflect.TypeOf((*RmArrayInput)(nil)).Elem(), RmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RmMapInput)(nil)).Elem(), RmMap{})
	pulumi.RegisterOutputType(RmOutput{})
	pulumi.RegisterOutputType(RmArrayOutput{})
	pulumi.RegisterOutputType(RmMapOutput{})
}
