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

// Starts etcd on a remote system.
type StartEtcd struct {
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

// NewStartEtcd registers a new resource with the given unique name, arguments, and options.
func NewStartEtcd(ctx *pulumi.Context,
	name string, args *StartEtcdArgs, opts ...pulumi.ResourceOption) (*StartEtcd, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StartEtcd
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:StartEtcd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type startEtcdArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
}

// The set of arguments for constructing a StartEtcd resource.
type StartEtcdArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
}

func (StartEtcdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*startEtcdArgs)(nil)).Elem()
}

type StartEtcdInput interface {
	pulumi.Input

	ToStartEtcdOutput() StartEtcdOutput
	ToStartEtcdOutputWithContext(ctx context.Context) StartEtcdOutput
}

func (*StartEtcd) ElementType() reflect.Type {
	return reflect.TypeOf((**StartEtcd)(nil)).Elem()
}

func (i *StartEtcd) ToStartEtcdOutput() StartEtcdOutput {
	return i.ToStartEtcdOutputWithContext(context.Background())
}

func (i *StartEtcd) ToStartEtcdOutputWithContext(ctx context.Context) StartEtcdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartEtcdOutput)
}

// StartEtcdArrayInput is an input type that accepts StartEtcdArray and StartEtcdArrayOutput values.
// You can construct a concrete instance of `StartEtcdArrayInput` via:
//
//	StartEtcdArray{ StartEtcdArgs{...} }
type StartEtcdArrayInput interface {
	pulumi.Input

	ToStartEtcdArrayOutput() StartEtcdArrayOutput
	ToStartEtcdArrayOutputWithContext(context.Context) StartEtcdArrayOutput
}

type StartEtcdArray []StartEtcdInput

func (StartEtcdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartEtcd)(nil)).Elem()
}

func (i StartEtcdArray) ToStartEtcdArrayOutput() StartEtcdArrayOutput {
	return i.ToStartEtcdArrayOutputWithContext(context.Background())
}

func (i StartEtcdArray) ToStartEtcdArrayOutputWithContext(ctx context.Context) StartEtcdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartEtcdArrayOutput)
}

// StartEtcdMapInput is an input type that accepts StartEtcdMap and StartEtcdMapOutput values.
// You can construct a concrete instance of `StartEtcdMapInput` via:
//
//	StartEtcdMap{ "key": StartEtcdArgs{...} }
type StartEtcdMapInput interface {
	pulumi.Input

	ToStartEtcdMapOutput() StartEtcdMapOutput
	ToStartEtcdMapOutputWithContext(context.Context) StartEtcdMapOutput
}

type StartEtcdMap map[string]StartEtcdInput

func (StartEtcdMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartEtcd)(nil)).Elem()
}

func (i StartEtcdMap) ToStartEtcdMapOutput() StartEtcdMapOutput {
	return i.ToStartEtcdMapOutputWithContext(context.Background())
}

func (i StartEtcdMap) ToStartEtcdMapOutputWithContext(ctx context.Context) StartEtcdMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StartEtcdMapOutput)
}

type StartEtcdOutput struct{ *pulumi.OutputState }

func (StartEtcdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StartEtcd)(nil)).Elem()
}

func (o StartEtcdOutput) ToStartEtcdOutput() StartEtcdOutput {
	return o
}

func (o StartEtcdOutput) ToStartEtcdOutputWithContext(ctx context.Context) StartEtcdOutput {
	return o
}

// The parameters with which to connect to the remote host.
func (o StartEtcdOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *StartEtcd) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The daemon-reload command.
func (o StartEtcdOutput) DaemonReload() tools.SystemctlOutput {
	return o.ApplyT(func(v *StartEtcd) tools.SystemctlOutput { return v.DaemonReload }).(tools.SystemctlOutput)
}

// The enable command.
func (o StartEtcdOutput) Enable() tools.SystemctlOutput {
	return o.ApplyT(func(v *StartEtcd) tools.SystemctlOutput { return v.Enable }).(tools.SystemctlOutput)
}

// The start command.
func (o StartEtcdOutput) Start() tools.SystemctlOutput {
	return o.ApplyT(func(v *StartEtcd) tools.SystemctlOutput { return v.Start }).(tools.SystemctlOutput)
}

type StartEtcdArrayOutput struct{ *pulumi.OutputState }

func (StartEtcdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StartEtcd)(nil)).Elem()
}

func (o StartEtcdArrayOutput) ToStartEtcdArrayOutput() StartEtcdArrayOutput {
	return o
}

func (o StartEtcdArrayOutput) ToStartEtcdArrayOutputWithContext(ctx context.Context) StartEtcdArrayOutput {
	return o
}

func (o StartEtcdArrayOutput) Index(i pulumi.IntInput) StartEtcdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StartEtcd {
		return vs[0].([]*StartEtcd)[vs[1].(int)]
	}).(StartEtcdOutput)
}

type StartEtcdMapOutput struct{ *pulumi.OutputState }

func (StartEtcdMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StartEtcd)(nil)).Elem()
}

func (o StartEtcdMapOutput) ToStartEtcdMapOutput() StartEtcdMapOutput {
	return o
}

func (o StartEtcdMapOutput) ToStartEtcdMapOutputWithContext(ctx context.Context) StartEtcdMapOutput {
	return o
}

func (o StartEtcdMapOutput) MapIndex(k pulumi.StringInput) StartEtcdOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StartEtcd {
		return vs[0].(map[string]*StartEtcd)[vs[1].(string)]
	}).(StartEtcdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StartEtcdInput)(nil)).Elem(), &StartEtcd{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartEtcdArrayInput)(nil)).Elem(), StartEtcdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StartEtcdMapInput)(nil)).Elem(), StartEtcdMap{})
	pulumi.RegisterOutputType(StartEtcdOutput{})
	pulumi.RegisterOutputType(StartEtcdArrayOutput{})
	pulumi.RegisterOutputType(StartEtcdMapOutput{})
}
