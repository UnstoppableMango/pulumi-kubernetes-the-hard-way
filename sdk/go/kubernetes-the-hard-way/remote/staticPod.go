// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/config"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/tools"
)

// Create a static pod manifest on a remote system.
type StaticPod struct {
	pulumi.ResourceState

	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The remote manifest file.
	File FileOutput `pulumi:"file"`
	// The name of the file on the remote system.
	FileName pulumi.StringOutput `pulumi:"fileName"`
	// The mkdir operation to ensure /etc/kubernetes/manifests exists.
	Mkdir tools.MkdirOutput `pulumi:"mkdir"`
	// The path to the manifest on the remote system.
	Path pulumi.StringOutput `pulumi:"path"`
	// The pod manifest.
	Pod config.PodManifestOutput `pulumi:"pod"`
}

// NewStaticPod registers a new resource with the given unique name, arguments, and options.
func NewStaticPod(ctx *pulumi.Context,
	name string, args *StaticPodArgs, opts ...pulumi.ResourceOption) (*StaticPod, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Pod == nil {
		return nil, errors.New("invalid value for required argument 'Pod'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StaticPod
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:StaticPod", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type staticPodArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The name of the file on the remote system.
	FileName *string `pulumi:"fileName"`
	// The pod manifest.
	Pod config.PodManifest `pulumi:"pod"`
}

// The set of arguments for constructing a StaticPod resource.
type StaticPodArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The name of the file on the remote system.
	FileName pulumi.StringPtrInput
	// The pod manifest.
	Pod config.PodManifestInput
}

func (StaticPodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticPodArgs)(nil)).Elem()
}

type StaticPodInput interface {
	pulumi.Input

	ToStaticPodOutput() StaticPodOutput
	ToStaticPodOutputWithContext(ctx context.Context) StaticPodOutput
}

func (*StaticPod) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticPod)(nil)).Elem()
}

func (i *StaticPod) ToStaticPodOutput() StaticPodOutput {
	return i.ToStaticPodOutputWithContext(context.Background())
}

func (i *StaticPod) ToStaticPodOutputWithContext(ctx context.Context) StaticPodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticPodOutput)
}

// StaticPodArrayInput is an input type that accepts StaticPodArray and StaticPodArrayOutput values.
// You can construct a concrete instance of `StaticPodArrayInput` via:
//
//	StaticPodArray{ StaticPodArgs{...} }
type StaticPodArrayInput interface {
	pulumi.Input

	ToStaticPodArrayOutput() StaticPodArrayOutput
	ToStaticPodArrayOutputWithContext(context.Context) StaticPodArrayOutput
}

type StaticPodArray []StaticPodInput

func (StaticPodArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticPod)(nil)).Elem()
}

func (i StaticPodArray) ToStaticPodArrayOutput() StaticPodArrayOutput {
	return i.ToStaticPodArrayOutputWithContext(context.Background())
}

func (i StaticPodArray) ToStaticPodArrayOutputWithContext(ctx context.Context) StaticPodArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticPodArrayOutput)
}

// StaticPodMapInput is an input type that accepts StaticPodMap and StaticPodMapOutput values.
// You can construct a concrete instance of `StaticPodMapInput` via:
//
//	StaticPodMap{ "key": StaticPodArgs{...} }
type StaticPodMapInput interface {
	pulumi.Input

	ToStaticPodMapOutput() StaticPodMapOutput
	ToStaticPodMapOutputWithContext(context.Context) StaticPodMapOutput
}

type StaticPodMap map[string]StaticPodInput

func (StaticPodMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticPod)(nil)).Elem()
}

func (i StaticPodMap) ToStaticPodMapOutput() StaticPodMapOutput {
	return i.ToStaticPodMapOutputWithContext(context.Background())
}

func (i StaticPodMap) ToStaticPodMapOutputWithContext(ctx context.Context) StaticPodMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticPodMapOutput)
}

type StaticPodOutput struct{ *pulumi.OutputState }

func (StaticPodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticPod)(nil)).Elem()
}

func (o StaticPodOutput) ToStaticPodOutput() StaticPodOutput {
	return o
}

func (o StaticPodOutput) ToStaticPodOutputWithContext(ctx context.Context) StaticPodOutput {
	return o
}

// The parameters with which to connect to the remote host.
func (o StaticPodOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *StaticPod) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The remote manifest file.
func (o StaticPodOutput) File() FileOutput {
	return o.ApplyT(func(v *StaticPod) FileOutput { return v.File }).(FileOutput)
}

// The name of the file on the remote system.
func (o StaticPodOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticPod) pulumi.StringOutput { return v.FileName }).(pulumi.StringOutput)
}

// The mkdir operation to ensure /etc/kubernetes/manifests exists.
func (o StaticPodOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *StaticPod) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

// The path to the manifest on the remote system.
func (o StaticPodOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticPod) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The pod manifest.
func (o StaticPodOutput) Pod() config.PodManifestOutput {
	return o.ApplyT(func(v *StaticPod) config.PodManifestOutput { return v.Pod }).(config.PodManifestOutput)
}

type StaticPodArrayOutput struct{ *pulumi.OutputState }

func (StaticPodArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticPod)(nil)).Elem()
}

func (o StaticPodArrayOutput) ToStaticPodArrayOutput() StaticPodArrayOutput {
	return o
}

func (o StaticPodArrayOutput) ToStaticPodArrayOutputWithContext(ctx context.Context) StaticPodArrayOutput {
	return o
}

func (o StaticPodArrayOutput) Index(i pulumi.IntInput) StaticPodOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StaticPod {
		return vs[0].([]*StaticPod)[vs[1].(int)]
	}).(StaticPodOutput)
}

type StaticPodMapOutput struct{ *pulumi.OutputState }

func (StaticPodMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticPod)(nil)).Elem()
}

func (o StaticPodMapOutput) ToStaticPodMapOutput() StaticPodMapOutput {
	return o
}

func (o StaticPodMapOutput) ToStaticPodMapOutputWithContext(ctx context.Context) StaticPodMapOutput {
	return o
}

func (o StaticPodMapOutput) MapIndex(k pulumi.StringInput) StaticPodOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StaticPod {
		return vs[0].(map[string]*StaticPod)[vs[1].(string)]
	}).(StaticPodOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticPodInput)(nil)).Elem(), &StaticPod{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticPodArrayInput)(nil)).Elem(), StaticPodArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticPodMapInput)(nil)).Elem(), StaticPodMap{})
	pulumi.RegisterOutputType(StaticPodOutput{})
	pulumi.RegisterOutputType(StaticPodArrayOutput{})
	pulumi.RegisterOutputType(StaticPodMapOutput{})
}
