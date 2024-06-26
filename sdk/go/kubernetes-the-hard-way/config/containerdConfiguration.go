// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Get the containerd configuration.
type ContainerdConfiguration struct {
	pulumi.ResourceState

	Result ContainerdConfigurationTypeOutput `pulumi:"result"`
	// The toml representation of the configuration.
	Toml pulumi.StringOutput `pulumi:"toml"`
}

// NewContainerdConfiguration registers a new resource with the given unique name, arguments, and options.
func NewContainerdConfiguration(ctx *pulumi.Context,
	name string, args *ContainerdConfigurationArgs, opts ...pulumi.ResourceOption) (*ContainerdConfiguration, error) {
	if args == nil {
		args = &ContainerdConfigurationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerdConfiguration
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:config:ContainerdConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type containerdConfigurationArgs struct {
	// The cri configuration.
	Cri *ContainerdCriPluginConfiguration `pulumi:"cri"`
}

// The set of arguments for constructing a ContainerdConfiguration resource.
type ContainerdConfigurationArgs struct {
	// The cri configuration.
	Cri *ContainerdCriPluginConfigurationArgs
}

func (ContainerdConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerdConfigurationArgs)(nil)).Elem()
}

type ContainerdConfigurationInput interface {
	pulumi.Input

	ToContainerdConfigurationOutput() ContainerdConfigurationOutput
	ToContainerdConfigurationOutputWithContext(ctx context.Context) ContainerdConfigurationOutput
}

func (*ContainerdConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerdConfiguration)(nil)).Elem()
}

func (i *ContainerdConfiguration) ToContainerdConfigurationOutput() ContainerdConfigurationOutput {
	return i.ToContainerdConfigurationOutputWithContext(context.Background())
}

func (i *ContainerdConfiguration) ToContainerdConfigurationOutputWithContext(ctx context.Context) ContainerdConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerdConfigurationOutput)
}

// ContainerdConfigurationArrayInput is an input type that accepts ContainerdConfigurationArray and ContainerdConfigurationArrayOutput values.
// You can construct a concrete instance of `ContainerdConfigurationArrayInput` via:
//
//	ContainerdConfigurationArray{ ContainerdConfigurationArgs{...} }
type ContainerdConfigurationArrayInput interface {
	pulumi.Input

	ToContainerdConfigurationArrayOutput() ContainerdConfigurationArrayOutput
	ToContainerdConfigurationArrayOutputWithContext(context.Context) ContainerdConfigurationArrayOutput
}

type ContainerdConfigurationArray []ContainerdConfigurationInput

func (ContainerdConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerdConfiguration)(nil)).Elem()
}

func (i ContainerdConfigurationArray) ToContainerdConfigurationArrayOutput() ContainerdConfigurationArrayOutput {
	return i.ToContainerdConfigurationArrayOutputWithContext(context.Background())
}

func (i ContainerdConfigurationArray) ToContainerdConfigurationArrayOutputWithContext(ctx context.Context) ContainerdConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerdConfigurationArrayOutput)
}

// ContainerdConfigurationMapInput is an input type that accepts ContainerdConfigurationMap and ContainerdConfigurationMapOutput values.
// You can construct a concrete instance of `ContainerdConfigurationMapInput` via:
//
//	ContainerdConfigurationMap{ "key": ContainerdConfigurationArgs{...} }
type ContainerdConfigurationMapInput interface {
	pulumi.Input

	ToContainerdConfigurationMapOutput() ContainerdConfigurationMapOutput
	ToContainerdConfigurationMapOutputWithContext(context.Context) ContainerdConfigurationMapOutput
}

type ContainerdConfigurationMap map[string]ContainerdConfigurationInput

func (ContainerdConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerdConfiguration)(nil)).Elem()
}

func (i ContainerdConfigurationMap) ToContainerdConfigurationMapOutput() ContainerdConfigurationMapOutput {
	return i.ToContainerdConfigurationMapOutputWithContext(context.Background())
}

func (i ContainerdConfigurationMap) ToContainerdConfigurationMapOutputWithContext(ctx context.Context) ContainerdConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerdConfigurationMapOutput)
}

type ContainerdConfigurationOutput struct{ *pulumi.OutputState }

func (ContainerdConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerdConfiguration)(nil)).Elem()
}

func (o ContainerdConfigurationOutput) ToContainerdConfigurationOutput() ContainerdConfigurationOutput {
	return o
}

func (o ContainerdConfigurationOutput) ToContainerdConfigurationOutputWithContext(ctx context.Context) ContainerdConfigurationOutput {
	return o
}

func (o ContainerdConfigurationOutput) Result() ContainerdConfigurationTypeOutput {
	return o.ApplyT(func(v *ContainerdConfiguration) ContainerdConfigurationTypeOutput { return v.Result }).(ContainerdConfigurationTypeOutput)
}

// The toml representation of the configuration.
func (o ContainerdConfigurationOutput) Toml() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerdConfiguration) pulumi.StringOutput { return v.Toml }).(pulumi.StringOutput)
}

type ContainerdConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ContainerdConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerdConfiguration)(nil)).Elem()
}

func (o ContainerdConfigurationArrayOutput) ToContainerdConfigurationArrayOutput() ContainerdConfigurationArrayOutput {
	return o
}

func (o ContainerdConfigurationArrayOutput) ToContainerdConfigurationArrayOutputWithContext(ctx context.Context) ContainerdConfigurationArrayOutput {
	return o
}

func (o ContainerdConfigurationArrayOutput) Index(i pulumi.IntInput) ContainerdConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerdConfiguration {
		return vs[0].([]*ContainerdConfiguration)[vs[1].(int)]
	}).(ContainerdConfigurationOutput)
}

type ContainerdConfigurationMapOutput struct{ *pulumi.OutputState }

func (ContainerdConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerdConfiguration)(nil)).Elem()
}

func (o ContainerdConfigurationMapOutput) ToContainerdConfigurationMapOutput() ContainerdConfigurationMapOutput {
	return o
}

func (o ContainerdConfigurationMapOutput) ToContainerdConfigurationMapOutputWithContext(ctx context.Context) ContainerdConfigurationMapOutput {
	return o
}

func (o ContainerdConfigurationMapOutput) MapIndex(k pulumi.StringInput) ContainerdConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerdConfiguration {
		return vs[0].(map[string]*ContainerdConfiguration)[vs[1].(string)]
	}).(ContainerdConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerdConfigurationInput)(nil)).Elem(), &ContainerdConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerdConfigurationArrayInput)(nil)).Elem(), ContainerdConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerdConfigurationMapInput)(nil)).Elem(), ContainerdConfigurationMap{})
	pulumi.RegisterOutputType(ContainerdConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerdConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ContainerdConfigurationMapOutput{})
}
