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

// Kubelet systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.
type KubeProxyService struct {
	pulumi.ResourceState

	// Kubelet configuration.
	Configuration KubeProxyConfigurationPropsOutput `pulumi:"configuration"`
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

// NewKubeProxyService registers a new resource with the given unique name, arguments, and options.
func NewKubeProxyService(ctx *pulumi.Context,
	name string, args *KubeProxyServiceArgs, opts ...pulumi.ResourceOption) (*KubeProxyService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KubeProxyService
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeProxyService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeProxyServiceArgs struct {
	// Kubelet configuration.
	Configuration KubeProxyConfigurationProps `pulumi:"configuration"`
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

// The set of arguments for constructing a KubeProxyService resource.
type KubeProxyServiceArgs struct {
	// Kubelet configuration.
	Configuration KubeProxyConfigurationPropsInput
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

func (KubeProxyServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeProxyServiceArgs)(nil)).Elem()
}

type KubeProxyServiceInput interface {
	pulumi.Input

	ToKubeProxyServiceOutput() KubeProxyServiceOutput
	ToKubeProxyServiceOutputWithContext(ctx context.Context) KubeProxyServiceOutput
}

func (*KubeProxyService) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeProxyService)(nil)).Elem()
}

func (i *KubeProxyService) ToKubeProxyServiceOutput() KubeProxyServiceOutput {
	return i.ToKubeProxyServiceOutputWithContext(context.Background())
}

func (i *KubeProxyService) ToKubeProxyServiceOutputWithContext(ctx context.Context) KubeProxyServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeProxyServiceOutput)
}

// KubeProxyServiceArrayInput is an input type that accepts KubeProxyServiceArray and KubeProxyServiceArrayOutput values.
// You can construct a concrete instance of `KubeProxyServiceArrayInput` via:
//
//	KubeProxyServiceArray{ KubeProxyServiceArgs{...} }
type KubeProxyServiceArrayInput interface {
	pulumi.Input

	ToKubeProxyServiceArrayOutput() KubeProxyServiceArrayOutput
	ToKubeProxyServiceArrayOutputWithContext(context.Context) KubeProxyServiceArrayOutput
}

type KubeProxyServiceArray []KubeProxyServiceInput

func (KubeProxyServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeProxyService)(nil)).Elem()
}

func (i KubeProxyServiceArray) ToKubeProxyServiceArrayOutput() KubeProxyServiceArrayOutput {
	return i.ToKubeProxyServiceArrayOutputWithContext(context.Background())
}

func (i KubeProxyServiceArray) ToKubeProxyServiceArrayOutputWithContext(ctx context.Context) KubeProxyServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeProxyServiceArrayOutput)
}

// KubeProxyServiceMapInput is an input type that accepts KubeProxyServiceMap and KubeProxyServiceMapOutput values.
// You can construct a concrete instance of `KubeProxyServiceMapInput` via:
//
//	KubeProxyServiceMap{ "key": KubeProxyServiceArgs{...} }
type KubeProxyServiceMapInput interface {
	pulumi.Input

	ToKubeProxyServiceMapOutput() KubeProxyServiceMapOutput
	ToKubeProxyServiceMapOutputWithContext(context.Context) KubeProxyServiceMapOutput
}

type KubeProxyServiceMap map[string]KubeProxyServiceInput

func (KubeProxyServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeProxyService)(nil)).Elem()
}

func (i KubeProxyServiceMap) ToKubeProxyServiceMapOutput() KubeProxyServiceMapOutput {
	return i.ToKubeProxyServiceMapOutputWithContext(context.Background())
}

func (i KubeProxyServiceMap) ToKubeProxyServiceMapOutputWithContext(ctx context.Context) KubeProxyServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeProxyServiceMapOutput)
}

type KubeProxyServiceOutput struct{ *pulumi.OutputState }

func (KubeProxyServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeProxyService)(nil)).Elem()
}

func (o KubeProxyServiceOutput) ToKubeProxyServiceOutput() KubeProxyServiceOutput {
	return o
}

func (o KubeProxyServiceOutput) ToKubeProxyServiceOutputWithContext(ctx context.Context) KubeProxyServiceOutput {
	return o
}

// Kubelet configuration.
func (o KubeProxyServiceOutput) Configuration() KubeProxyConfigurationPropsOutput {
	return o.ApplyT(func(v *KubeProxyService) KubeProxyConfigurationPropsOutput { return v.Configuration }).(KubeProxyConfigurationPropsOutput)
}

// The parameters with which to connect to the remote host.
func (o KubeProxyServiceOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeProxyService) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Optional systemd unit description.
func (o KubeProxyServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeProxyService) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The location to create the service file.
func (o KubeProxyServiceOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeProxyService) pulumi.StringPtrOutput { return v.Directory }).(pulumi.StringPtrOutput)
}

// Optional systemd unit documentation
func (o KubeProxyServiceOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeProxyService) pulumi.StringPtrOutput { return v.Documentation }).(pulumi.StringPtrOutput)
}

// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
func (o KubeProxyServiceOutput) Restart() SystemdServiceRestartPtrOutput {
	return o.ApplyT(func(v *KubeProxyService) SystemdServiceRestartPtrOutput { return v.Restart }).(SystemdServiceRestartPtrOutput)
}

// Optionally override the systemd service RestartSec. Defaults to `5`.
func (o KubeProxyServiceOutput) RestartSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeProxyService) pulumi.StringPtrOutput { return v.RestartSec }).(pulumi.StringPtrOutput)
}

// The remote systemd service.
func (o KubeProxyServiceOutput) Service() SystemdServiceOutput {
	return o.ApplyT(func(v *KubeProxyService) SystemdServiceOutput { return v.Service }).(SystemdServiceOutput)
}

// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
func (o KubeProxyServiceOutput) WantedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeProxyService) pulumi.StringPtrOutput { return v.WantedBy }).(pulumi.StringPtrOutput)
}

type KubeProxyServiceArrayOutput struct{ *pulumi.OutputState }

func (KubeProxyServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeProxyService)(nil)).Elem()
}

func (o KubeProxyServiceArrayOutput) ToKubeProxyServiceArrayOutput() KubeProxyServiceArrayOutput {
	return o
}

func (o KubeProxyServiceArrayOutput) ToKubeProxyServiceArrayOutputWithContext(ctx context.Context) KubeProxyServiceArrayOutput {
	return o
}

func (o KubeProxyServiceArrayOutput) Index(i pulumi.IntInput) KubeProxyServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeProxyService {
		return vs[0].([]*KubeProxyService)[vs[1].(int)]
	}).(KubeProxyServiceOutput)
}

type KubeProxyServiceMapOutput struct{ *pulumi.OutputState }

func (KubeProxyServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeProxyService)(nil)).Elem()
}

func (o KubeProxyServiceMapOutput) ToKubeProxyServiceMapOutput() KubeProxyServiceMapOutput {
	return o
}

func (o KubeProxyServiceMapOutput) ToKubeProxyServiceMapOutputWithContext(ctx context.Context) KubeProxyServiceMapOutput {
	return o
}

func (o KubeProxyServiceMapOutput) MapIndex(k pulumi.StringInput) KubeProxyServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeProxyService {
		return vs[0].(map[string]*KubeProxyService)[vs[1].(string)]
	}).(KubeProxyServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeProxyServiceInput)(nil)).Elem(), &KubeProxyService{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeProxyServiceArrayInput)(nil)).Elem(), KubeProxyServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeProxyServiceMapInput)(nil)).Elem(), KubeProxyServiceMap{})
	pulumi.RegisterOutputType(KubeProxyServiceOutput{})
	pulumi.RegisterOutputType(KubeProxyServiceArrayOutput{})
	pulumi.RegisterOutputType(KubeProxyServiceMapOutput{})
}
