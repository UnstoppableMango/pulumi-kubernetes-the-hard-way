// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// TODO
type Kubeconfig struct {
	pulumi.ResourceState

	Result KubeconfigTypeOutput `pulumi:"result"`
	// The yaml representation of the manifest.
	Yaml pulumi.StringOutput `pulumi:"yaml"`
}

// NewKubeconfig registers a new resource with the given unique name, arguments, and options.
func NewKubeconfig(ctx *pulumi.Context,
	name string, args *KubeconfigArgs, opts ...pulumi.ResourceOption) (*Kubeconfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPem == nil {
		return nil, errors.New("invalid value for required argument 'CaPem'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Kubeconfig
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:config:Kubeconfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeconfigArgs struct {
	// Certificate authority data.
	CaPem string `pulumi:"caPem"`
	// Options for creating the kubeconfig.
	Options interface{} `pulumi:"options"`
}

// The set of arguments for constructing a Kubeconfig resource.
type KubeconfigArgs struct {
	// Certificate authority data.
	CaPem pulumi.StringInput
	// Options for creating the kubeconfig.
	Options interface{}
}

func (KubeconfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeconfigArgs)(nil)).Elem()
}

type KubeconfigInput interface {
	pulumi.Input

	ToKubeconfigOutput() KubeconfigOutput
	ToKubeconfigOutputWithContext(ctx context.Context) KubeconfigOutput
}

func (*Kubeconfig) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubeconfig)(nil)).Elem()
}

func (i *Kubeconfig) ToKubeconfigOutput() KubeconfigOutput {
	return i.ToKubeconfigOutputWithContext(context.Background())
}

func (i *Kubeconfig) ToKubeconfigOutputWithContext(ctx context.Context) KubeconfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeconfigOutput)
}

// KubeconfigArrayInput is an input type that accepts KubeconfigArray and KubeconfigArrayOutput values.
// You can construct a concrete instance of `KubeconfigArrayInput` via:
//
//	KubeconfigArray{ KubeconfigArgs{...} }
type KubeconfigArrayInput interface {
	pulumi.Input

	ToKubeconfigArrayOutput() KubeconfigArrayOutput
	ToKubeconfigArrayOutputWithContext(context.Context) KubeconfigArrayOutput
}

type KubeconfigArray []KubeconfigInput

func (KubeconfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubeconfig)(nil)).Elem()
}

func (i KubeconfigArray) ToKubeconfigArrayOutput() KubeconfigArrayOutput {
	return i.ToKubeconfigArrayOutputWithContext(context.Background())
}

func (i KubeconfigArray) ToKubeconfigArrayOutputWithContext(ctx context.Context) KubeconfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeconfigArrayOutput)
}

// KubeconfigMapInput is an input type that accepts KubeconfigMap and KubeconfigMapOutput values.
// You can construct a concrete instance of `KubeconfigMapInput` via:
//
//	KubeconfigMap{ "key": KubeconfigArgs{...} }
type KubeconfigMapInput interface {
	pulumi.Input

	ToKubeconfigMapOutput() KubeconfigMapOutput
	ToKubeconfigMapOutputWithContext(context.Context) KubeconfigMapOutput
}

type KubeconfigMap map[string]KubeconfigInput

func (KubeconfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubeconfig)(nil)).Elem()
}

func (i KubeconfigMap) ToKubeconfigMapOutput() KubeconfigMapOutput {
	return i.ToKubeconfigMapOutputWithContext(context.Background())
}

func (i KubeconfigMap) ToKubeconfigMapOutputWithContext(ctx context.Context) KubeconfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeconfigMapOutput)
}

type KubeconfigOutput struct{ *pulumi.OutputState }

func (KubeconfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubeconfig)(nil)).Elem()
}

func (o KubeconfigOutput) ToKubeconfigOutput() KubeconfigOutput {
	return o
}

func (o KubeconfigOutput) ToKubeconfigOutputWithContext(ctx context.Context) KubeconfigOutput {
	return o
}

func (o KubeconfigOutput) Result() KubeconfigTypeOutput {
	return o.ApplyT(func(v *Kubeconfig) KubeconfigTypeOutput { return v.Result }).(KubeconfigTypeOutput)
}

// The yaml representation of the manifest.
func (o KubeconfigOutput) Yaml() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubeconfig) pulumi.StringOutput { return v.Yaml }).(pulumi.StringOutput)
}

type KubeconfigArrayOutput struct{ *pulumi.OutputState }

func (KubeconfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubeconfig)(nil)).Elem()
}

func (o KubeconfigArrayOutput) ToKubeconfigArrayOutput() KubeconfigArrayOutput {
	return o
}

func (o KubeconfigArrayOutput) ToKubeconfigArrayOutputWithContext(ctx context.Context) KubeconfigArrayOutput {
	return o
}

func (o KubeconfigArrayOutput) Index(i pulumi.IntInput) KubeconfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Kubeconfig {
		return vs[0].([]*Kubeconfig)[vs[1].(int)]
	}).(KubeconfigOutput)
}

type KubeconfigMapOutput struct{ *pulumi.OutputState }

func (KubeconfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubeconfig)(nil)).Elem()
}

func (o KubeconfigMapOutput) ToKubeconfigMapOutput() KubeconfigMapOutput {
	return o
}

func (o KubeconfigMapOutput) ToKubeconfigMapOutputWithContext(ctx context.Context) KubeconfigMapOutput {
	return o
}

func (o KubeconfigMapOutput) MapIndex(k pulumi.StringInput) KubeconfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Kubeconfig {
		return vs[0].(map[string]*Kubeconfig)[vs[1].(string)]
	}).(KubeconfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeconfigInput)(nil)).Elem(), &Kubeconfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeconfigArrayInput)(nil)).Elem(), KubeconfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeconfigMapInput)(nil)).Elem(), KubeconfigMap{})
	pulumi.RegisterOutputType(KubeconfigOutput{})
	pulumi.RegisterOutputType(KubeconfigArrayOutput{})
	pulumi.RegisterOutputType(KubeconfigMapOutput{})
}
