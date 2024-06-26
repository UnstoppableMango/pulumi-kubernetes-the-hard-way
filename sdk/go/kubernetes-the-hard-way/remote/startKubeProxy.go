// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/tools"
)

// Starts `kube-proxy` on a remote system
type StartKubeProxy struct {
	pulumi.ResourceState

	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The daemon-reload command.
	DaemonReload tools.SystemctlOutput `pulumi:"daemonReload"`
	// The enable command.
	Enable tools.SystemctlOutput `pulumi:"enable"`
	// The start command.
	Start tools.SystemctlOutput `pulumi:"start"`
}

// NewStartKubeProxy registers a new resource with the given unique name, arguments, and options.
func NewStartKubeProxy(ctx *pulumi.Context,
	name string, args *StartKubeProxyArgs, opts ...pulumi.ResourceOption) (*StartKubeProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StartKubeProxy
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:StartKubeProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type startKubeProxyArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
}

// The set of arguments for constructing a StartKubeProxy resource.
type StartKubeProxyArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
}

func (StartKubeProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*startKubeProxyArgs)(nil)).Elem()
}

type StartKubeProxyInput interface {
	pulumi.Input

	ToStartKubeProxyOutput() StartKubeProxyOutput
	ToStartKubeProxyOutputWithContext(ctx context.Context) StartKubeProxyOutput
}

func (*StartKubeProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**StartKubeProxy)(nil)).Elem()
}

func (i *StartKubeProxy) ToStartKubeProxyOutput() StartKubeProxyOutput {
	return i.ToStartKubeProxyOutputWithContext(context.Background())
}

func (i *StartKubeProxy) ToStartKubeProxyOutputWithContext(ctx context.Context) StartKubeProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartKubeProxyOutput)
}

// StartKubeProxyArrayInput is an input type that accepts StartKubeProxyArray and StartKubeProxyArrayOutput values.
// You can construct a concrete instance of `StartKubeProxyArrayInput` via:
//
//	StartKubeProxyArray{ StartKubeProxyArgs{...} }
type StartKubeProxyArrayInput interface {
	pulumi.Input

	ToStartKubeProxyArrayOutput() StartKubeProxyArrayOutput
	ToStartKubeProxyArrayOutputWithContext(context.Context) StartKubeProxyArrayOutput
}

type StartKubeProxyArray []StartKubeProxyInput

func (StartKubeProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartKubeProxy)(nil)).Elem()
}

func (i StartKubeProxyArray) ToStartKubeProxyArrayOutput() StartKubeProxyArrayOutput {
	return i.ToStartKubeProxyArrayOutputWithContext(context.Background())
}

func (i StartKubeProxyArray) ToStartKubeProxyArrayOutputWithContext(ctx context.Context) StartKubeProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartKubeProxyArrayOutput)
}

// StartKubeProxyMapInput is an input type that accepts StartKubeProxyMap and StartKubeProxyMapOutput values.
// You can construct a concrete instance of `StartKubeProxyMapInput` via:
//
//	StartKubeProxyMap{ "key": StartKubeProxyArgs{...} }
type StartKubeProxyMapInput interface {
	pulumi.Input

	ToStartKubeProxyMapOutput() StartKubeProxyMapOutput
	ToStartKubeProxyMapOutputWithContext(context.Context) StartKubeProxyMapOutput
}

type StartKubeProxyMap map[string]StartKubeProxyInput

func (StartKubeProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartKubeProxy)(nil)).Elem()
}

func (i StartKubeProxyMap) ToStartKubeProxyMapOutput() StartKubeProxyMapOutput {
	return i.ToStartKubeProxyMapOutputWithContext(context.Background())
}

func (i StartKubeProxyMap) ToStartKubeProxyMapOutputWithContext(ctx context.Context) StartKubeProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartKubeProxyMapOutput)
}

type StartKubeProxyOutput struct{ *pulumi.OutputState }

func (StartKubeProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartKubeProxy)(nil)).Elem()
}

func (o StartKubeProxyOutput) ToStartKubeProxyOutput() StartKubeProxyOutput {
	return o
}

func (o StartKubeProxyOutput) ToStartKubeProxyOutputWithContext(ctx context.Context) StartKubeProxyOutput {
	return o
}

// The parameters with which to connect to the remote host.
func (o StartKubeProxyOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *StartKubeProxy) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The daemon-reload command.
func (o StartKubeProxyOutput) DaemonReload() tools.SystemctlOutput {
	return o.ApplyT(func(v *StartKubeProxy) tools.SystemctlOutput { return v.DaemonReload }).(tools.SystemctlOutput)
}

// The enable command.
func (o StartKubeProxyOutput) Enable() tools.SystemctlOutput {
	return o.ApplyT(func(v *StartKubeProxy) tools.SystemctlOutput { return v.Enable }).(tools.SystemctlOutput)
}

// The start command.
func (o StartKubeProxyOutput) Start() tools.SystemctlOutput {
	return o.ApplyT(func(v *StartKubeProxy) tools.SystemctlOutput { return v.Start }).(tools.SystemctlOutput)
}

type StartKubeProxyArrayOutput struct{ *pulumi.OutputState }

func (StartKubeProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartKubeProxy)(nil)).Elem()
}

func (o StartKubeProxyArrayOutput) ToStartKubeProxyArrayOutput() StartKubeProxyArrayOutput {
	return o
}

func (o StartKubeProxyArrayOutput) ToStartKubeProxyArrayOutputWithContext(ctx context.Context) StartKubeProxyArrayOutput {
	return o
}

func (o StartKubeProxyArrayOutput) Index(i pulumi.IntInput) StartKubeProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StartKubeProxy {
		return vs[0].([]*StartKubeProxy)[vs[1].(int)]
	}).(StartKubeProxyOutput)
}

type StartKubeProxyMapOutput struct{ *pulumi.OutputState }

func (StartKubeProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartKubeProxy)(nil)).Elem()
}

func (o StartKubeProxyMapOutput) ToStartKubeProxyMapOutput() StartKubeProxyMapOutput {
	return o
}

func (o StartKubeProxyMapOutput) ToStartKubeProxyMapOutputWithContext(ctx context.Context) StartKubeProxyMapOutput {
	return o
}

func (o StartKubeProxyMapOutput) MapIndex(k pulumi.StringInput) StartKubeProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StartKubeProxy {
		return vs[0].(map[string]*StartKubeProxy)[vs[1].(string)]
	}).(StartKubeProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StartKubeProxyInput)(nil)).Elem(), &StartKubeProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartKubeProxyArrayInput)(nil)).Elem(), StartKubeProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartKubeProxyMapInput)(nil)).Elem(), StartKubeProxyMap{})
	pulumi.RegisterOutputType(StartKubeProxyOutput{})
	pulumi.RegisterOutputType(StartKubeProxyArrayOutput{})
	pulumi.RegisterOutputType(StartKubeProxyMapOutput{})
}
