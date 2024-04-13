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
)

// Etcd systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.
type EtcdService struct {
	pulumi.ResourceState

	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Optional systemd unit description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The location to create the service file.
	Directory pulumi.StringPtrOutput `pulumi:"directory"`
	// Optional systemd unit documentation
	Documentation pulumi.StringPtrOutput `pulumi:"documentation"`
	// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
	Restart SystemdServiceRestartPtrOutput `pulumi:"restart"`
	// Optionally override the systemd service RestartSec. Defaults to `5`.
	RestartSec pulumi.StringPtrOutput `pulumi:"restartSec"`
	// The remote systemd service.
	Service SystemdServiceOutput `pulumi:"service"`
	// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
	WantedBy pulumi.StringPtrOutput `pulumi:"wantedBy"`
}

// NewEtcdService registers a new resource with the given unique name, arguments, and options.
func NewEtcdService(ctx *pulumi.Context,
	name string, args *EtcdServiceArgs, opts ...pulumi.ResourceOption) (*EtcdService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EtcdService
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:EtcdService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type etcdServiceArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Optional systemd unit description.
	Description *string `pulumi:"description"`
	// The location to create the service file.
	Directory *string `pulumi:"directory"`
	// Optional systemd unit documentation
	Documentation *string `pulumi:"documentation"`
	// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
	Restart *SystemdServiceRestart `pulumi:"restart"`
	// Optionally override the systemd service RestartSec. Defaults to `5`.
	RestartSec *string `pulumi:"restartSec"`
	// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
	WantedBy *string `pulumi:"wantedBy"`
}

// The set of arguments for constructing a EtcdService resource.
type EtcdServiceArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// Optional systemd unit description.
	Description pulumi.StringPtrInput
	// The location to create the service file.
	Directory pulumi.StringPtrInput
	// Optional systemd unit documentation
	Documentation pulumi.StringPtrInput
	// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
	Restart SystemdServiceRestartPtrInput
	// Optionally override the systemd service RestartSec. Defaults to `5`.
	RestartSec pulumi.StringPtrInput
	// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
	WantedBy pulumi.StringPtrInput
}

func (EtcdServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*etcdServiceArgs)(nil)).Elem()
}

type EtcdServiceInput interface {
	pulumi.Input

	ToEtcdServiceOutput() EtcdServiceOutput
	ToEtcdServiceOutputWithContext(ctx context.Context) EtcdServiceOutput
}

func (*EtcdService) ElementType() reflect.Type {
	return reflect.TypeOf((**EtcdService)(nil)).Elem()
}

func (i *EtcdService) ToEtcdServiceOutput() EtcdServiceOutput {
	return i.ToEtcdServiceOutputWithContext(context.Background())
}

func (i *EtcdService) ToEtcdServiceOutputWithContext(ctx context.Context) EtcdServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdServiceOutput)
}

// EtcdServiceArrayInput is an input type that accepts EtcdServiceArray and EtcdServiceArrayOutput values.
// You can construct a concrete instance of `EtcdServiceArrayInput` via:
//
//	EtcdServiceArray{ EtcdServiceArgs{...} }
type EtcdServiceArrayInput interface {
	pulumi.Input

	ToEtcdServiceArrayOutput() EtcdServiceArrayOutput
	ToEtcdServiceArrayOutputWithContext(context.Context) EtcdServiceArrayOutput
}

type EtcdServiceArray []EtcdServiceInput

func (EtcdServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EtcdService)(nil)).Elem()
}

func (i EtcdServiceArray) ToEtcdServiceArrayOutput() EtcdServiceArrayOutput {
	return i.ToEtcdServiceArrayOutputWithContext(context.Background())
}

func (i EtcdServiceArray) ToEtcdServiceArrayOutputWithContext(ctx context.Context) EtcdServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdServiceArrayOutput)
}

// EtcdServiceMapInput is an input type that accepts EtcdServiceMap and EtcdServiceMapOutput values.
// You can construct a concrete instance of `EtcdServiceMapInput` via:
//
//	EtcdServiceMap{ "key": EtcdServiceArgs{...} }
type EtcdServiceMapInput interface {
	pulumi.Input

	ToEtcdServiceMapOutput() EtcdServiceMapOutput
	ToEtcdServiceMapOutputWithContext(context.Context) EtcdServiceMapOutput
}

type EtcdServiceMap map[string]EtcdServiceInput

func (EtcdServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EtcdService)(nil)).Elem()
}

func (i EtcdServiceMap) ToEtcdServiceMapOutput() EtcdServiceMapOutput {
	return i.ToEtcdServiceMapOutputWithContext(context.Background())
}

func (i EtcdServiceMap) ToEtcdServiceMapOutputWithContext(ctx context.Context) EtcdServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdServiceMapOutput)
}

type EtcdServiceOutput struct{ *pulumi.OutputState }

func (EtcdServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EtcdService)(nil)).Elem()
}

func (o EtcdServiceOutput) ToEtcdServiceOutput() EtcdServiceOutput {
	return o
}

func (o EtcdServiceOutput) ToEtcdServiceOutputWithContext(ctx context.Context) EtcdServiceOutput {
	return o
}

// The parameters with which to connect to the remote host.
func (o EtcdServiceOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *EtcdService) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Optional systemd unit description.
func (o EtcdServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EtcdService) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The location to create the service file.
func (o EtcdServiceOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EtcdService) pulumi.StringPtrOutput { return v.Directory }).(pulumi.StringPtrOutput)
}

// Optional systemd unit documentation
func (o EtcdServiceOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EtcdService) pulumi.StringPtrOutput { return v.Documentation }).(pulumi.StringPtrOutput)
}

// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
func (o EtcdServiceOutput) Restart() SystemdServiceRestartPtrOutput {
	return o.ApplyT(func(v *EtcdService) SystemdServiceRestartPtrOutput { return v.Restart }).(SystemdServiceRestartPtrOutput)
}

// Optionally override the systemd service RestartSec. Defaults to `5`.
func (o EtcdServiceOutput) RestartSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EtcdService) pulumi.StringPtrOutput { return v.RestartSec }).(pulumi.StringPtrOutput)
}

// The remote systemd service.
func (o EtcdServiceOutput) Service() SystemdServiceOutput {
	return o.ApplyT(func(v *EtcdService) SystemdServiceOutput { return v.Service }).(SystemdServiceOutput)
}

// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
func (o EtcdServiceOutput) WantedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EtcdService) pulumi.StringPtrOutput { return v.WantedBy }).(pulumi.StringPtrOutput)
}

type EtcdServiceArrayOutput struct{ *pulumi.OutputState }

func (EtcdServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EtcdService)(nil)).Elem()
}

func (o EtcdServiceArrayOutput) ToEtcdServiceArrayOutput() EtcdServiceArrayOutput {
	return o
}

func (o EtcdServiceArrayOutput) ToEtcdServiceArrayOutputWithContext(ctx context.Context) EtcdServiceArrayOutput {
	return o
}

func (o EtcdServiceArrayOutput) Index(i pulumi.IntInput) EtcdServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EtcdService {
		return vs[0].([]*EtcdService)[vs[1].(int)]
	}).(EtcdServiceOutput)
}

type EtcdServiceMapOutput struct{ *pulumi.OutputState }

func (EtcdServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EtcdService)(nil)).Elem()
}

func (o EtcdServiceMapOutput) ToEtcdServiceMapOutput() EtcdServiceMapOutput {
	return o
}

func (o EtcdServiceMapOutput) ToEtcdServiceMapOutputWithContext(ctx context.Context) EtcdServiceMapOutput {
	return o
}

func (o EtcdServiceMapOutput) MapIndex(k pulumi.StringInput) EtcdServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EtcdService {
		return vs[0].(map[string]*EtcdService)[vs[1].(string)]
	}).(EtcdServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdServiceInput)(nil)).Elem(), &EtcdService{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdServiceArrayInput)(nil)).Elem(), EtcdServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdServiceMapInput)(nil)).Elem(), EtcdServiceMap{})
	pulumi.RegisterOutputType(EtcdServiceOutput{})
	pulumi.RegisterOutputType(EtcdServiceArrayOutput{})
	pulumi.RegisterOutputType(EtcdServiceMapOutput{})
}
