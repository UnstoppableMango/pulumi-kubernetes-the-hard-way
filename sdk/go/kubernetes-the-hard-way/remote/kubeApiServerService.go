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

// Kube API Server systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.
type KubeApiServerService struct {
	pulumi.ResourceState

	// If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
	ClientCaFile pulumi.StringPtrOutput `pulumi:"clientCaFile"`
	// KubeApiServer configuration.
	Configuration KubeApiServerPropsOutput `pulumi:"configuration"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// Optional systemd unit description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The location to create the service file.
	Directory pulumi.StringPtrOutput `pulumi:"directory"`
	// Optional systemd unit documentation
	Documentation pulumi.StringPtrOutput `pulumi:"documentation"`
	// List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
	EtcdServers pulumi.StringPtrOutput `pulumi:"etcdServers"`
	// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
	Restart SystemdServiceRestartPtrOutput `pulumi:"restart"`
	// Optionally override the systemd service RestartSec. Defaults to `5`.
	RestartSec pulumi.StringPtrOutput `pulumi:"restartSec"`
	// The remote systemd service.
	Service SystemdServiceOutput `pulumi:"service"`
	// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
	WantedBy pulumi.StringPtrOutput `pulumi:"wantedBy"`
}

// NewKubeApiServerService registers a new resource with the given unique name, arguments, and options.
func NewKubeApiServerService(ctx *pulumi.Context,
	name string, args *KubeApiServerServiceArgs, opts ...pulumi.ResourceOption) (*KubeApiServerService, error) {
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
	var resource KubeApiServerService
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:KubeApiServerService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type kubeApiServerServiceArgs struct {
	// If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
	ClientCaFile *string `pulumi:"clientCaFile"`
	// KubeApiServer configuration.
	Configuration KubeApiServerProps `pulumi:"configuration"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Optional systemd unit description.
	Description *string `pulumi:"description"`
	// The location to create the service file.
	Directory *string `pulumi:"directory"`
	// Optional systemd unit documentation
	Documentation *string `pulumi:"documentation"`
	// List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
	EtcdServers *string `pulumi:"etcdServers"`
	// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
	Restart *SystemdServiceRestart `pulumi:"restart"`
	// Optionally override the systemd service RestartSec. Defaults to `5`.
	RestartSec *string `pulumi:"restartSec"`
	// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
	WantedBy *string `pulumi:"wantedBy"`
}

// The set of arguments for constructing a KubeApiServerService resource.
type KubeApiServerServiceArgs struct {
	// If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
	ClientCaFile pulumi.StringPtrInput
	// KubeApiServer configuration.
	Configuration KubeApiServerPropsInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// Optional systemd unit description.
	Description pulumi.StringPtrInput
	// The location to create the service file.
	Directory pulumi.StringPtrInput
	// Optional systemd unit documentation
	Documentation pulumi.StringPtrInput
	// List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
	EtcdServers pulumi.StringPtrInput
	// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
	Restart SystemdServiceRestartPtrInput
	// Optionally override the systemd service RestartSec. Defaults to `5`.
	RestartSec pulumi.StringPtrInput
	// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
	WantedBy pulumi.StringPtrInput
}

func (KubeApiServerServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeApiServerServiceArgs)(nil)).Elem()
}

type KubeApiServerServiceInput interface {
	pulumi.Input

	ToKubeApiServerServiceOutput() KubeApiServerServiceOutput
	ToKubeApiServerServiceOutputWithContext(ctx context.Context) KubeApiServerServiceOutput
}

func (*KubeApiServerService) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeApiServerService)(nil)).Elem()
}

func (i *KubeApiServerService) ToKubeApiServerServiceOutput() KubeApiServerServiceOutput {
	return i.ToKubeApiServerServiceOutputWithContext(context.Background())
}

func (i *KubeApiServerService) ToKubeApiServerServiceOutputWithContext(ctx context.Context) KubeApiServerServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerServiceOutput)
}

// KubeApiServerServiceArrayInput is an input type that accepts KubeApiServerServiceArray and KubeApiServerServiceArrayOutput values.
// You can construct a concrete instance of `KubeApiServerServiceArrayInput` via:
//
//	KubeApiServerServiceArray{ KubeApiServerServiceArgs{...} }
type KubeApiServerServiceArrayInput interface {
	pulumi.Input

	ToKubeApiServerServiceArrayOutput() KubeApiServerServiceArrayOutput
	ToKubeApiServerServiceArrayOutputWithContext(context.Context) KubeApiServerServiceArrayOutput
}

type KubeApiServerServiceArray []KubeApiServerServiceInput

func (KubeApiServerServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeApiServerService)(nil)).Elem()
}

func (i KubeApiServerServiceArray) ToKubeApiServerServiceArrayOutput() KubeApiServerServiceArrayOutput {
	return i.ToKubeApiServerServiceArrayOutputWithContext(context.Background())
}

func (i KubeApiServerServiceArray) ToKubeApiServerServiceArrayOutputWithContext(ctx context.Context) KubeApiServerServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerServiceArrayOutput)
}

// KubeApiServerServiceMapInput is an input type that accepts KubeApiServerServiceMap and KubeApiServerServiceMapOutput values.
// You can construct a concrete instance of `KubeApiServerServiceMapInput` via:
//
//	KubeApiServerServiceMap{ "key": KubeApiServerServiceArgs{...} }
type KubeApiServerServiceMapInput interface {
	pulumi.Input

	ToKubeApiServerServiceMapOutput() KubeApiServerServiceMapOutput
	ToKubeApiServerServiceMapOutputWithContext(context.Context) KubeApiServerServiceMapOutput
}

type KubeApiServerServiceMap map[string]KubeApiServerServiceInput

func (KubeApiServerServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeApiServerService)(nil)).Elem()
}

func (i KubeApiServerServiceMap) ToKubeApiServerServiceMapOutput() KubeApiServerServiceMapOutput {
	return i.ToKubeApiServerServiceMapOutputWithContext(context.Background())
}

func (i KubeApiServerServiceMap) ToKubeApiServerServiceMapOutputWithContext(ctx context.Context) KubeApiServerServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeApiServerServiceMapOutput)
}

type KubeApiServerServiceOutput struct{ *pulumi.OutputState }

func (KubeApiServerServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeApiServerService)(nil)).Elem()
}

func (o KubeApiServerServiceOutput) ToKubeApiServerServiceOutput() KubeApiServerServiceOutput {
	return o
}

func (o KubeApiServerServiceOutput) ToKubeApiServerServiceOutputWithContext(ctx context.Context) KubeApiServerServiceOutput {
	return o
}

// If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate
func (o KubeApiServerServiceOutput) ClientCaFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumi.StringPtrOutput { return v.ClientCaFile }).(pulumi.StringPtrOutput)
}

// KubeApiServer configuration.
func (o KubeApiServerServiceOutput) Configuration() KubeApiServerPropsOutput {
	return o.ApplyT(func(v *KubeApiServerService) KubeApiServerPropsOutput { return v.Configuration }).(KubeApiServerPropsOutput)
}

// The parameters with which to connect to the remote host.
func (o KubeApiServerServiceOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// Optional systemd unit description.
func (o KubeApiServerServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The location to create the service file.
func (o KubeApiServerServiceOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumi.StringPtrOutput { return v.Directory }).(pulumi.StringPtrOutput)
}

// Optional systemd unit documentation
func (o KubeApiServerServiceOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumi.StringPtrOutput { return v.Documentation }).(pulumi.StringPtrOutput)
}

// List of etcd servers to connect with (scheme://ip:port), comma separatedList of etcd servers to connect with (scheme://ip:port), comma separated
func (o KubeApiServerServiceOutput) EtcdServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumi.StringPtrOutput { return v.EtcdServers }).(pulumi.StringPtrOutput)
}

// Optionally override the systemd service restart behaviour. Defaults to `on-failure`.
func (o KubeApiServerServiceOutput) Restart() SystemdServiceRestartPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) SystemdServiceRestartPtrOutput { return v.Restart }).(SystemdServiceRestartPtrOutput)
}

// Optionally override the systemd service RestartSec. Defaults to `5`.
func (o KubeApiServerServiceOutput) RestartSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumi.StringPtrOutput { return v.RestartSec }).(pulumi.StringPtrOutput)
}

// The remote systemd service.
func (o KubeApiServerServiceOutput) Service() SystemdServiceOutput {
	return o.ApplyT(func(v *KubeApiServerService) SystemdServiceOutput { return v.Service }).(SystemdServiceOutput)
}

// Optionally override the systemd service wanted-by. Defaults to `multi-user.target`.
func (o KubeApiServerServiceOutput) WantedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeApiServerService) pulumi.StringPtrOutput { return v.WantedBy }).(pulumi.StringPtrOutput)
}

type KubeApiServerServiceArrayOutput struct{ *pulumi.OutputState }

func (KubeApiServerServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeApiServerService)(nil)).Elem()
}

func (o KubeApiServerServiceArrayOutput) ToKubeApiServerServiceArrayOutput() KubeApiServerServiceArrayOutput {
	return o
}

func (o KubeApiServerServiceArrayOutput) ToKubeApiServerServiceArrayOutputWithContext(ctx context.Context) KubeApiServerServiceArrayOutput {
	return o
}

func (o KubeApiServerServiceArrayOutput) Index(i pulumi.IntInput) KubeApiServerServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeApiServerService {
		return vs[0].([]*KubeApiServerService)[vs[1].(int)]
	}).(KubeApiServerServiceOutput)
}

type KubeApiServerServiceMapOutput struct{ *pulumi.OutputState }

func (KubeApiServerServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeApiServerService)(nil)).Elem()
}

func (o KubeApiServerServiceMapOutput) ToKubeApiServerServiceMapOutput() KubeApiServerServiceMapOutput {
	return o
}

func (o KubeApiServerServiceMapOutput) ToKubeApiServerServiceMapOutputWithContext(ctx context.Context) KubeApiServerServiceMapOutput {
	return o
}

func (o KubeApiServerServiceMapOutput) MapIndex(k pulumi.StringInput) KubeApiServerServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeApiServerService {
		return vs[0].(map[string]*KubeApiServerService)[vs[1].(string)]
	}).(KubeApiServerServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerServiceInput)(nil)).Elem(), &KubeApiServerService{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerServiceArrayInput)(nil)).Elem(), KubeApiServerServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeApiServerServiceMapInput)(nil)).Elem(), KubeApiServerServiceMap{})
	pulumi.RegisterOutputType(KubeApiServerServiceOutput{})
	pulumi.RegisterOutputType(KubeApiServerServiceArrayOutput{})
	pulumi.RegisterOutputType(KubeApiServerServiceMapOutput{})
}