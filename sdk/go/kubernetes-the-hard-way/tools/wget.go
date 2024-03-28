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

// Abstraction over the `wget` utility on a remote system.
type Wget struct {
	pulumi.ResourceState

	// Represents the remote `tar` operation.
	Command remote.CommandOutput `pulumi:"command"`
	// Corresponds to the --directory-prefix option.
	DirectoryPrefix pulumi.StringPtrOutput `pulumi:"directoryPrefix"`
	// Corresponds to the --https-only option.
	HttpsOnly pulumi.BoolOutput `pulumi:"httpsOnly"`
	// Corresponds to the --output-document option.
	OutputDocument pulumi.StringPtrOutput `pulumi:"outputDocument"`
	// Corresponds to the --quiet option.
	Quiet pulumi.BoolOutput `pulumi:"quiet"`
	// The process' stderr.
	Stderr pulumi.StringOutput `pulumi:"stderr"`
	// The process' stdin.
	Stdin pulumi.StringPtrOutput `pulumi:"stdin"`
	// The process' stdout.
	Stdout pulumi.StringOutput `pulumi:"stdout"`
	// Corresponds to the --timestamping option.
	Timestamping pulumi.BoolOutput `pulumi:"timestamping"`
	// Corresponse to the [URL] argument.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewWget registers a new resource with the given unique name, arguments, and options.
func NewWget(ctx *pulumi.Context,
	name string, args *WgetArgs, opts ...pulumi.ResourceOption) (*Wget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v remote.Connection) remote.Connection { return *v.Defaults() }).(remote.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Wget
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:tools:Wget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type wgetArgs struct {
	// Connection details for the remote system.
	Connection remote.Connection `pulumi:"connection"`
	// Corresponds to the --directory-prefix option.
	DirectoryPrefix *string `pulumi:"directoryPrefix"`
	// Corresponds to the --https-only option.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// Corresponds to the --output-document option.
	OutputDocument *string `pulumi:"outputDocument"`
	// Corresponds to the --quiet option.
	Quiet *bool `pulumi:"quiet"`
	// Corresponds to the --timestamping option.
	Timestamping *bool `pulumi:"timestamping"`
	// Corresponse to the [URL] argument.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Wget resource.
type WgetArgs struct {
	// Connection details for the remote system.
	Connection remote.ConnectionInput
	// Corresponds to the --directory-prefix option.
	DirectoryPrefix pulumi.StringPtrInput
	// Corresponds to the --https-only option.
	HttpsOnly pulumi.BoolPtrInput
	// Corresponds to the --output-document option.
	OutputDocument pulumi.StringPtrInput
	// Corresponds to the --quiet option.
	Quiet pulumi.BoolPtrInput
	// Corresponds to the --timestamping option.
	Timestamping pulumi.BoolPtrInput
	// Corresponse to the [URL] argument.
	Url pulumi.StringInput
}

func (WgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wgetArgs)(nil)).Elem()
}

type WgetInput interface {
	pulumi.Input

	ToWgetOutput() WgetOutput
	ToWgetOutputWithContext(ctx context.Context) WgetOutput
}

func (*Wget) ElementType() reflect.Type {
	return reflect.TypeOf((**Wget)(nil)).Elem()
}

func (i *Wget) ToWgetOutput() WgetOutput {
	return i.ToWgetOutputWithContext(context.Background())
}

func (i *Wget) ToWgetOutputWithContext(ctx context.Context) WgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WgetOutput)
}

// WgetArrayInput is an input type that accepts WgetArray and WgetArrayOutput values.
// You can construct a concrete instance of `WgetArrayInput` via:
//
//	WgetArray{ WgetArgs{...} }
type WgetArrayInput interface {
	pulumi.Input

	ToWgetArrayOutput() WgetArrayOutput
	ToWgetArrayOutputWithContext(context.Context) WgetArrayOutput
}

type WgetArray []WgetInput

func (WgetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wget)(nil)).Elem()
}

func (i WgetArray) ToWgetArrayOutput() WgetArrayOutput {
	return i.ToWgetArrayOutputWithContext(context.Background())
}

func (i WgetArray) ToWgetArrayOutputWithContext(ctx context.Context) WgetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WgetArrayOutput)
}

// WgetMapInput is an input type that accepts WgetMap and WgetMapOutput values.
// You can construct a concrete instance of `WgetMapInput` via:
//
//	WgetMap{ "key": WgetArgs{...} }
type WgetMapInput interface {
	pulumi.Input

	ToWgetMapOutput() WgetMapOutput
	ToWgetMapOutputWithContext(context.Context) WgetMapOutput
}

type WgetMap map[string]WgetInput

func (WgetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wget)(nil)).Elem()
}

func (i WgetMap) ToWgetMapOutput() WgetMapOutput {
	return i.ToWgetMapOutputWithContext(context.Background())
}

func (i WgetMap) ToWgetMapOutputWithContext(ctx context.Context) WgetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WgetMapOutput)
}

type WgetOutput struct{ *pulumi.OutputState }

func (WgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wget)(nil)).Elem()
}

func (o WgetOutput) ToWgetOutput() WgetOutput {
	return o
}

func (o WgetOutput) ToWgetOutputWithContext(ctx context.Context) WgetOutput {
	return o
}

// Represents the remote `tar` operation.
func (o WgetOutput) Command() remote.CommandOutput {
	return o.ApplyT(func(v *Wget) remote.CommandOutput { return v.Command }).(remote.CommandOutput)
}

// Corresponds to the --directory-prefix option.
func (o WgetOutput) DirectoryPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringPtrOutput { return v.DirectoryPrefix }).(pulumi.StringPtrOutput)
}

// Corresponds to the --https-only option.
func (o WgetOutput) HttpsOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *Wget) pulumi.BoolOutput { return v.HttpsOnly }).(pulumi.BoolOutput)
}

// Corresponds to the --output-document option.
func (o WgetOutput) OutputDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringPtrOutput { return v.OutputDocument }).(pulumi.StringPtrOutput)
}

// Corresponds to the --quiet option.
func (o WgetOutput) Quiet() pulumi.BoolOutput {
	return o.ApplyT(func(v *Wget) pulumi.BoolOutput { return v.Quiet }).(pulumi.BoolOutput)
}

// The process' stderr.
func (o WgetOutput) Stderr() pulumi.StringOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringOutput { return v.Stderr }).(pulumi.StringOutput)
}

// The process' stdin.
func (o WgetOutput) Stdin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringPtrOutput { return v.Stdin }).(pulumi.StringPtrOutput)
}

// The process' stdout.
func (o WgetOutput) Stdout() pulumi.StringOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringOutput { return v.Stdout }).(pulumi.StringOutput)
}

// Corresponds to the --timestamping option.
func (o WgetOutput) Timestamping() pulumi.BoolOutput {
	return o.ApplyT(func(v *Wget) pulumi.BoolOutput { return v.Timestamping }).(pulumi.BoolOutput)
}

// Corresponse to the [URL] argument.
func (o WgetOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Wget) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type WgetArrayOutput struct{ *pulumi.OutputState }

func (WgetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wget)(nil)).Elem()
}

func (o WgetArrayOutput) ToWgetArrayOutput() WgetArrayOutput {
	return o
}

func (o WgetArrayOutput) ToWgetArrayOutputWithContext(ctx context.Context) WgetArrayOutput {
	return o
}

func (o WgetArrayOutput) Index(i pulumi.IntInput) WgetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wget {
		return vs[0].([]*Wget)[vs[1].(int)]
	}).(WgetOutput)
}

type WgetMapOutput struct{ *pulumi.OutputState }

func (WgetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wget)(nil)).Elem()
}

func (o WgetMapOutput) ToWgetMapOutput() WgetMapOutput {
	return o
}

func (o WgetMapOutput) ToWgetMapOutputWithContext(ctx context.Context) WgetMapOutput {
	return o
}

func (o WgetMapOutput) MapIndex(k pulumi.StringInput) WgetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wget {
		return vs[0].(map[string]*Wget)[vs[1].(string)]
	}).(WgetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WgetInput)(nil)).Elem(), &Wget{})
	pulumi.RegisterInputType(reflect.TypeOf((*WgetArrayInput)(nil)).Elem(), WgetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WgetMapInput)(nil)).Elem(), WgetMap{})
	pulumi.RegisterOutputType(WgetOutput{})
	pulumi.RegisterOutputType(WgetArrayOutput{})
	pulumi.RegisterOutputType(WgetMapOutput{})
}