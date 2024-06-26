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

// A systemd service on a remote system.
type SystemdService struct {
	pulumi.ResourceState

	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The location to create the service file.
	Directory pulumi.StringOutput `pulumi:"directory"`
	// The service file on the remote machine.
	File FileOutput `pulumi:"file"`
	// Describes the [Install] section of a systemd service file.
	Install SystemdInstallSectionPtrOutput `pulumi:"install"`
	// Describes the [Service] section of a systemd service file.
	Service SystemdServiceSectionOutput `pulumi:"service"`
	// Describes the [Unit] section of a systemd service file.
	Unit SystemdUnitSectionPtrOutput `pulumi:"unit"`
	// Name of the systemd unit.
	UnitName pulumi.StringPtrOutput `pulumi:"unitName"`
}

// NewSystemdService registers a new resource with the given unique name, arguments, and options.
func NewSystemdService(ctx *pulumi.Context,
	name string, args *SystemdServiceArgs, opts ...pulumi.ResourceOption) (*SystemdService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	if args.Directory == nil {
		args.Directory = pulumi.StringPtr("/etc/systemd/system")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemdService
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:SystemdService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type systemdServiceArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The location to create the service file.
	Directory *string `pulumi:"directory"`
	// Describes the [Install] section of a systemd service file.
	Install *SystemdInstallSection `pulumi:"install"`
	// Describes the [Service] section of a systemd service file.
	Service SystemdServiceSection `pulumi:"service"`
	// Describes the [Unit] section of a systemd service file.
	Unit *SystemdUnitSection `pulumi:"unit"`
	// Name of the systemd unit.
	UnitName *string `pulumi:"unitName"`
}

// The set of arguments for constructing a SystemdService resource.
type SystemdServiceArgs struct {
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The location to create the service file.
	Directory pulumi.StringPtrInput
	// Describes the [Install] section of a systemd service file.
	Install SystemdInstallSectionPtrInput
	// Describes the [Service] section of a systemd service file.
	Service SystemdServiceSectionInput
	// Describes the [Unit] section of a systemd service file.
	Unit SystemdUnitSectionPtrInput
	// Name of the systemd unit.
	UnitName pulumi.StringPtrInput
}

func (SystemdServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemdServiceArgs)(nil)).Elem()
}

type SystemdServiceInput interface {
	pulumi.Input

	ToSystemdServiceOutput() SystemdServiceOutput
	ToSystemdServiceOutputWithContext(ctx context.Context) SystemdServiceOutput
}

func (*SystemdService) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdService)(nil)).Elem()
}

func (i *SystemdService) ToSystemdServiceOutput() SystemdServiceOutput {
	return i.ToSystemdServiceOutputWithContext(context.Background())
}

func (i *SystemdService) ToSystemdServiceOutputWithContext(ctx context.Context) SystemdServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdServiceOutput)
}

// SystemdServiceArrayInput is an input type that accepts SystemdServiceArray and SystemdServiceArrayOutput values.
// You can construct a concrete instance of `SystemdServiceArrayInput` via:
//
//	SystemdServiceArray{ SystemdServiceArgs{...} }
type SystemdServiceArrayInput interface {
	pulumi.Input

	ToSystemdServiceArrayOutput() SystemdServiceArrayOutput
	ToSystemdServiceArrayOutputWithContext(context.Context) SystemdServiceArrayOutput
}

type SystemdServiceArray []SystemdServiceInput

func (SystemdServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemdService)(nil)).Elem()
}

func (i SystemdServiceArray) ToSystemdServiceArrayOutput() SystemdServiceArrayOutput {
	return i.ToSystemdServiceArrayOutputWithContext(context.Background())
}

func (i SystemdServiceArray) ToSystemdServiceArrayOutputWithContext(ctx context.Context) SystemdServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdServiceArrayOutput)
}

// SystemdServiceMapInput is an input type that accepts SystemdServiceMap and SystemdServiceMapOutput values.
// You can construct a concrete instance of `SystemdServiceMapInput` via:
//
//	SystemdServiceMap{ "key": SystemdServiceArgs{...} }
type SystemdServiceMapInput interface {
	pulumi.Input

	ToSystemdServiceMapOutput() SystemdServiceMapOutput
	ToSystemdServiceMapOutputWithContext(context.Context) SystemdServiceMapOutput
}

type SystemdServiceMap map[string]SystemdServiceInput

func (SystemdServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemdService)(nil)).Elem()
}

func (i SystemdServiceMap) ToSystemdServiceMapOutput() SystemdServiceMapOutput {
	return i.ToSystemdServiceMapOutputWithContext(context.Background())
}

func (i SystemdServiceMap) ToSystemdServiceMapOutputWithContext(ctx context.Context) SystemdServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdServiceMapOutput)
}

type SystemdServiceOutput struct{ *pulumi.OutputState }

func (SystemdServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdService)(nil)).Elem()
}

func (o SystemdServiceOutput) ToSystemdServiceOutput() SystemdServiceOutput {
	return o
}

func (o SystemdServiceOutput) ToSystemdServiceOutputWithContext(ctx context.Context) SystemdServiceOutput {
	return o
}

// The parameters with which to connect to the remote host.
func (o SystemdServiceOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *SystemdService) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The location to create the service file.
func (o SystemdServiceOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemdService) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

// The service file on the remote machine.
func (o SystemdServiceOutput) File() FileOutput {
	return o.ApplyT(func(v *SystemdService) FileOutput { return v.File }).(FileOutput)
}

// Describes the [Install] section of a systemd service file.
func (o SystemdServiceOutput) Install() SystemdInstallSectionPtrOutput {
	return o.ApplyT(func(v *SystemdService) SystemdInstallSectionPtrOutput { return v.Install }).(SystemdInstallSectionPtrOutput)
}

// Describes the [Service] section of a systemd service file.
func (o SystemdServiceOutput) Service() SystemdServiceSectionOutput {
	return o.ApplyT(func(v *SystemdService) SystemdServiceSectionOutput { return v.Service }).(SystemdServiceSectionOutput)
}

// Describes the [Unit] section of a systemd service file.
func (o SystemdServiceOutput) Unit() SystemdUnitSectionPtrOutput {
	return o.ApplyT(func(v *SystemdService) SystemdUnitSectionPtrOutput { return v.Unit }).(SystemdUnitSectionPtrOutput)
}

// Name of the systemd unit.
func (o SystemdServiceOutput) UnitName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemdService) pulumi.StringPtrOutput { return v.UnitName }).(pulumi.StringPtrOutput)
}

type SystemdServiceArrayOutput struct{ *pulumi.OutputState }

func (SystemdServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemdService)(nil)).Elem()
}

func (o SystemdServiceArrayOutput) ToSystemdServiceArrayOutput() SystemdServiceArrayOutput {
	return o
}

func (o SystemdServiceArrayOutput) ToSystemdServiceArrayOutputWithContext(ctx context.Context) SystemdServiceArrayOutput {
	return o
}

func (o SystemdServiceArrayOutput) Index(i pulumi.IntInput) SystemdServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemdService {
		return vs[0].([]*SystemdService)[vs[1].(int)]
	}).(SystemdServiceOutput)
}

type SystemdServiceMapOutput struct{ *pulumi.OutputState }

func (SystemdServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemdService)(nil)).Elem()
}

func (o SystemdServiceMapOutput) ToSystemdServiceMapOutput() SystemdServiceMapOutput {
	return o
}

func (o SystemdServiceMapOutput) ToSystemdServiceMapOutputWithContext(ctx context.Context) SystemdServiceMapOutput {
	return o
}

func (o SystemdServiceMapOutput) MapIndex(k pulumi.StringInput) SystemdServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemdService {
		return vs[0].(map[string]*SystemdService)[vs[1].(string)]
	}).(SystemdServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceInput)(nil)).Elem(), &SystemdService{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceArrayInput)(nil)).Elem(), SystemdServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceMapInput)(nil)).Elem(), SystemdServiceMap{})
	pulumi.RegisterOutputType(SystemdServiceOutput{})
	pulumi.RegisterOutputType(SystemdServiceArrayOutput{})
	pulumi.RegisterOutputType(SystemdServiceMapOutput{})
}
