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

type EtcdConfiguration struct {
	pulumi.ResourceState

	// The remote certificate authority file.
	CaFile FileOutput `pulumi:"caFile"`
	// The remote certificate file.
	CertFile FileOutput `pulumi:"certFile"`
	// The directory to store etcd configuration.
	ConfigurationDirectory pulumi.StringOutput `pulumi:"configurationDirectory"`
	// The command used to create the configuration directory.
	ConfigurationMkdir tools.MkdirOutput `pulumi:"configurationMkdir"`
	// The directory etcd will use.
	DataDirectory pulumi.StringOutput `pulumi:"dataDirectory"`
	// The command used to create the data directory.
	DataMkdir tools.MkdirOutput `pulumi:"dataMkdir"`
	// IP used to serve client requests and communicate with etcd peers.
	InternalIp pulumi.StringOutput `pulumi:"internalIp"`
	// The remote key file.
	KeyFile FileOutput `pulumi:"keyFile"`
	// The remote systemd service.
	SystemdService SystemdServiceOutput `pulumi:"systemdService"`
}

// NewEtcdConfiguration registers a new resource with the given unique name, arguments, and options.
func NewEtcdConfiguration(ctx *pulumi.Context,
	name string, args *EtcdConfigurationArgs, opts ...pulumi.ResourceOption) (*EtcdConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPem == nil {
		return nil, errors.New("invalid value for required argument 'CaPem'")
	}
	if args.CertPem == nil {
		return nil, errors.New("invalid value for required argument 'CertPem'")
	}
	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.EtcdPath == nil {
		return nil, errors.New("invalid value for required argument 'EtcdPath'")
	}
	if args.InternalIp == nil {
		return nil, errors.New("invalid value for required argument 'InternalIp'")
	}
	if args.KeyPem == nil {
		return nil, errors.New("invalid value for required argument 'KeyPem'")
	}
	if args.ConfigurationDirectory == nil {
		args.ConfigurationDirectory = pulumi.StringPtr("/etc/etcd")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	if args.DataDirectory == nil {
		args.DataDirectory = pulumi.StringPtr("/var/lib/etcd")
	}
	if args.SystemdDirectory == nil {
		args.SystemdDirectory = pulumi.StringPtr("/etc/system/systemd")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EtcdConfiguration
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:EtcdConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type etcdConfigurationArgs struct {
	// The PEM encoded CA data.
	CaPem string `pulumi:"caPem"`
	// The PEM encoded certificate data.
	CertPem string `pulumi:"certPem"`
	// The directory to store etcd configuration.
	ConfigurationDirectory *string `pulumi:"configurationDirectory"`
	// The connection details.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The directory etcd will use.
	DataDirectory *string `pulumi:"dataDirectory"`
	EtcdPath      string  `pulumi:"etcdPath"`
	// IP used to serve client requests and communicate with etcd peers.
	InternalIp string `pulumi:"internalIp"`
	// The PEM encoded key data.
	KeyPem string `pulumi:"keyPem"`
	// The systemd service file dirctory.
	SystemdDirectory *string `pulumi:"systemdDirectory"`
}

// The set of arguments for constructing a EtcdConfiguration resource.
type EtcdConfigurationArgs struct {
	// The PEM encoded CA data.
	CaPem pulumi.StringInput
	// The PEM encoded certificate data.
	CertPem pulumi.StringInput
	// The directory to store etcd configuration.
	ConfigurationDirectory pulumi.StringPtrInput
	// The connection details.
	Connection pulumiCommand.ConnectionInput
	// The directory etcd will use.
	DataDirectory pulumi.StringPtrInput
	EtcdPath      pulumi.StringInput
	// IP used to serve client requests and communicate with etcd peers.
	InternalIp pulumi.StringInput
	// The PEM encoded key data.
	KeyPem pulumi.StringInput
	// The systemd service file dirctory.
	SystemdDirectory pulumi.StringPtrInput
}

func (EtcdConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*etcdConfigurationArgs)(nil)).Elem()
}

type EtcdConfigurationInput interface {
	pulumi.Input

	ToEtcdConfigurationOutput() EtcdConfigurationOutput
	ToEtcdConfigurationOutputWithContext(ctx context.Context) EtcdConfigurationOutput
}

func (*EtcdConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**EtcdConfiguration)(nil)).Elem()
}

func (i *EtcdConfiguration) ToEtcdConfigurationOutput() EtcdConfigurationOutput {
	return i.ToEtcdConfigurationOutputWithContext(context.Background())
}

func (i *EtcdConfiguration) ToEtcdConfigurationOutputWithContext(ctx context.Context) EtcdConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdConfigurationOutput)
}

// EtcdConfigurationArrayInput is an input type that accepts EtcdConfigurationArray and EtcdConfigurationArrayOutput values.
// You can construct a concrete instance of `EtcdConfigurationArrayInput` via:
//
//	EtcdConfigurationArray{ EtcdConfigurationArgs{...} }
type EtcdConfigurationArrayInput interface {
	pulumi.Input

	ToEtcdConfigurationArrayOutput() EtcdConfigurationArrayOutput
	ToEtcdConfigurationArrayOutputWithContext(context.Context) EtcdConfigurationArrayOutput
}

type EtcdConfigurationArray []EtcdConfigurationInput

func (EtcdConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EtcdConfiguration)(nil)).Elem()
}

func (i EtcdConfigurationArray) ToEtcdConfigurationArrayOutput() EtcdConfigurationArrayOutput {
	return i.ToEtcdConfigurationArrayOutputWithContext(context.Background())
}

func (i EtcdConfigurationArray) ToEtcdConfigurationArrayOutputWithContext(ctx context.Context) EtcdConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdConfigurationArrayOutput)
}

// EtcdConfigurationMapInput is an input type that accepts EtcdConfigurationMap and EtcdConfigurationMapOutput values.
// You can construct a concrete instance of `EtcdConfigurationMapInput` via:
//
//	EtcdConfigurationMap{ "key": EtcdConfigurationArgs{...} }
type EtcdConfigurationMapInput interface {
	pulumi.Input

	ToEtcdConfigurationMapOutput() EtcdConfigurationMapOutput
	ToEtcdConfigurationMapOutputWithContext(context.Context) EtcdConfigurationMapOutput
}

type EtcdConfigurationMap map[string]EtcdConfigurationInput

func (EtcdConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EtcdConfiguration)(nil)).Elem()
}

func (i EtcdConfigurationMap) ToEtcdConfigurationMapOutput() EtcdConfigurationMapOutput {
	return i.ToEtcdConfigurationMapOutputWithContext(context.Background())
}

func (i EtcdConfigurationMap) ToEtcdConfigurationMapOutputWithContext(ctx context.Context) EtcdConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdConfigurationMapOutput)
}

type EtcdConfigurationOutput struct{ *pulumi.OutputState }

func (EtcdConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EtcdConfiguration)(nil)).Elem()
}

func (o EtcdConfigurationOutput) ToEtcdConfigurationOutput() EtcdConfigurationOutput {
	return o
}

func (o EtcdConfigurationOutput) ToEtcdConfigurationOutputWithContext(ctx context.Context) EtcdConfigurationOutput {
	return o
}

// The remote certificate authority file.
func (o EtcdConfigurationOutput) CaFile() FileOutput {
	return o.ApplyT(func(v *EtcdConfiguration) FileOutput { return v.CaFile }).(FileOutput)
}

// The remote certificate file.
func (o EtcdConfigurationOutput) CertFile() FileOutput {
	return o.ApplyT(func(v *EtcdConfiguration) FileOutput { return v.CertFile }).(FileOutput)
}

// The directory to store etcd configuration.
func (o EtcdConfigurationOutput) ConfigurationDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *EtcdConfiguration) pulumi.StringOutput { return v.ConfigurationDirectory }).(pulumi.StringOutput)
}

// The command used to create the configuration directory.
func (o EtcdConfigurationOutput) ConfigurationMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *EtcdConfiguration) tools.MkdirOutput { return v.ConfigurationMkdir }).(tools.MkdirOutput)
}

// The directory etcd will use.
func (o EtcdConfigurationOutput) DataDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *EtcdConfiguration) pulumi.StringOutput { return v.DataDirectory }).(pulumi.StringOutput)
}

// The command used to create the data directory.
func (o EtcdConfigurationOutput) DataMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *EtcdConfiguration) tools.MkdirOutput { return v.DataMkdir }).(tools.MkdirOutput)
}

// IP used to serve client requests and communicate with etcd peers.
func (o EtcdConfigurationOutput) InternalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *EtcdConfiguration) pulumi.StringOutput { return v.InternalIp }).(pulumi.StringOutput)
}

// The remote key file.
func (o EtcdConfigurationOutput) KeyFile() FileOutput {
	return o.ApplyT(func(v *EtcdConfiguration) FileOutput { return v.KeyFile }).(FileOutput)
}

// The remote systemd service.
func (o EtcdConfigurationOutput) SystemdService() SystemdServiceOutput {
	return o.ApplyT(func(v *EtcdConfiguration) SystemdServiceOutput { return v.SystemdService }).(SystemdServiceOutput)
}

type EtcdConfigurationArrayOutput struct{ *pulumi.OutputState }

func (EtcdConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EtcdConfiguration)(nil)).Elem()
}

func (o EtcdConfigurationArrayOutput) ToEtcdConfigurationArrayOutput() EtcdConfigurationArrayOutput {
	return o
}

func (o EtcdConfigurationArrayOutput) ToEtcdConfigurationArrayOutputWithContext(ctx context.Context) EtcdConfigurationArrayOutput {
	return o
}

func (o EtcdConfigurationArrayOutput) Index(i pulumi.IntInput) EtcdConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EtcdConfiguration {
		return vs[0].([]*EtcdConfiguration)[vs[1].(int)]
	}).(EtcdConfigurationOutput)
}

type EtcdConfigurationMapOutput struct{ *pulumi.OutputState }

func (EtcdConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EtcdConfiguration)(nil)).Elem()
}

func (o EtcdConfigurationMapOutput) ToEtcdConfigurationMapOutput() EtcdConfigurationMapOutput {
	return o
}

func (o EtcdConfigurationMapOutput) ToEtcdConfigurationMapOutputWithContext(ctx context.Context) EtcdConfigurationMapOutput {
	return o
}

func (o EtcdConfigurationMapOutput) MapIndex(k pulumi.StringInput) EtcdConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EtcdConfiguration {
		return vs[0].(map[string]*EtcdConfiguration)[vs[1].(string)]
	}).(EtcdConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdConfigurationInput)(nil)).Elem(), &EtcdConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdConfigurationArrayInput)(nil)).Elem(), EtcdConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdConfigurationMapInput)(nil)).Elem(), EtcdConfigurationMap{})
	pulumi.RegisterOutputType(EtcdConfigurationOutput{})
	pulumi.RegisterOutputType(EtcdConfigurationArrayOutput{})
	pulumi.RegisterOutputType(EtcdConfigurationMapOutput{})
}