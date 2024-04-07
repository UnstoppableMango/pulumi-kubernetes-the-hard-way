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

// Installs cni-plugins on a remote system.
type CniPluginsInstall struct {
	pulumi.ResourceState

	// The CPU architecture.
	Architecture  ArchitectureOutput  `pulumi:"architecture"`
	ArchiveName   pulumi.StringOutput `pulumi:"archiveName"`
	BandwidthPath pulumi.StringOutput `pulumi:"bandwidthPath"`
	BridgePath    pulumi.StringOutput `pulumi:"bridgePath"`
	// The connection details.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	DhcpPath   pulumi.StringOutput            `pulumi:"dhcpPath"`
	// Directory to install the binary.
	Directory      pulumi.StringOutput `pulumi:"directory"`
	Download       DownloadOutput      `pulumi:"download"`
	DummyPath      pulumi.StringOutput `pulumi:"dummyPath"`
	FirewallPath   pulumi.StringOutput `pulumi:"firewallPath"`
	HostDevicePath pulumi.StringOutput `pulumi:"hostDevicePath"`
	HostLocalPath  pulumi.StringOutput `pulumi:"hostLocalPath"`
	IpvlanPath     pulumi.StringOutput `pulumi:"ipvlanPath"`
	LoopbackPath   pulumi.StringOutput `pulumi:"loopbackPath"`
	MacvlanPath    pulumi.StringOutput `pulumi:"macvlanPath"`
	Mkdir          tools.MkdirOutput   `pulumi:"mkdir"`
	Mktemp         tools.MktempOutput  `pulumi:"mktemp"`
	Mv             tools.MvOutput      `pulumi:"mv"`
	PortmapPath    pulumi.StringOutput `pulumi:"portmapPath"`
	PtpPath        pulumi.StringOutput `pulumi:"ptpPath"`
	Rm             tools.RmOutput      `pulumi:"rm"`
	SbrPath        pulumi.StringOutput `pulumi:"sbrPath"`
	StaticPath     pulumi.StringOutput `pulumi:"staticPath"`
	TapPath        pulumi.StringOutput `pulumi:"tapPath"`
	Tar            tools.TarOutput     `pulumi:"tar"`
	TuningPath     pulumi.StringOutput `pulumi:"tuningPath"`
	Url            pulumi.StringOutput `pulumi:"url"`
	// The version to install.
	Version  pulumi.StringOutput `pulumi:"version"`
	VlanPath pulumi.StringOutput `pulumi:"vlanPath"`
	VrfPath  pulumi.StringOutput `pulumi:"vrfPath"`
}

// NewCniPluginsInstall registers a new resource with the given unique name, arguments, and options.
func NewCniPluginsInstall(ctx *pulumi.Context,
	name string, args *CniPluginsInstallArgs, opts ...pulumi.ResourceOption) (*CniPluginsInstall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	if args.Directory == nil {
		args.Directory = pulumi.StringPtr("/usr/local/bin")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CniPluginsInstall
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:CniPluginsInstall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type cniPluginsInstallArgs struct {
	// The CPU architecture.
	Architecture *Architecture `pulumi:"architecture"`
	// The connection details.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// Directory to install the binary.
	Directory *string `pulumi:"directory"`
	// The version of to install.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a CniPluginsInstall resource.
type CniPluginsInstallArgs struct {
	// The CPU architecture.
	Architecture ArchitecturePtrInput
	// The connection details.
	Connection pulumiCommand.ConnectionInput
	// Directory to install the binary.
	Directory pulumi.StringPtrInput
	// The version of to install.
	Version pulumi.StringPtrInput
}

func (CniPluginsInstallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cniPluginsInstallArgs)(nil)).Elem()
}

type CniPluginsInstallInput interface {
	pulumi.Input

	ToCniPluginsInstallOutput() CniPluginsInstallOutput
	ToCniPluginsInstallOutputWithContext(ctx context.Context) CniPluginsInstallOutput
}

func (*CniPluginsInstall) ElementType() reflect.Type {
	return reflect.TypeOf((**CniPluginsInstall)(nil)).Elem()
}

func (i *CniPluginsInstall) ToCniPluginsInstallOutput() CniPluginsInstallOutput {
	return i.ToCniPluginsInstallOutputWithContext(context.Background())
}

func (i *CniPluginsInstall) ToCniPluginsInstallOutputWithContext(ctx context.Context) CniPluginsInstallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniPluginsInstallOutput)
}

// CniPluginsInstallArrayInput is an input type that accepts CniPluginsInstallArray and CniPluginsInstallArrayOutput values.
// You can construct a concrete instance of `CniPluginsInstallArrayInput` via:
//
//	CniPluginsInstallArray{ CniPluginsInstallArgs{...} }
type CniPluginsInstallArrayInput interface {
	pulumi.Input

	ToCniPluginsInstallArrayOutput() CniPluginsInstallArrayOutput
	ToCniPluginsInstallArrayOutputWithContext(context.Context) CniPluginsInstallArrayOutput
}

type CniPluginsInstallArray []CniPluginsInstallInput

func (CniPluginsInstallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CniPluginsInstall)(nil)).Elem()
}

func (i CniPluginsInstallArray) ToCniPluginsInstallArrayOutput() CniPluginsInstallArrayOutput {
	return i.ToCniPluginsInstallArrayOutputWithContext(context.Background())
}

func (i CniPluginsInstallArray) ToCniPluginsInstallArrayOutputWithContext(ctx context.Context) CniPluginsInstallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniPluginsInstallArrayOutput)
}

// CniPluginsInstallMapInput is an input type that accepts CniPluginsInstallMap and CniPluginsInstallMapOutput values.
// You can construct a concrete instance of `CniPluginsInstallMapInput` via:
//
//	CniPluginsInstallMap{ "key": CniPluginsInstallArgs{...} }
type CniPluginsInstallMapInput interface {
	pulumi.Input

	ToCniPluginsInstallMapOutput() CniPluginsInstallMapOutput
	ToCniPluginsInstallMapOutputWithContext(context.Context) CniPluginsInstallMapOutput
}

type CniPluginsInstallMap map[string]CniPluginsInstallInput

func (CniPluginsInstallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CniPluginsInstall)(nil)).Elem()
}

func (i CniPluginsInstallMap) ToCniPluginsInstallMapOutput() CniPluginsInstallMapOutput {
	return i.ToCniPluginsInstallMapOutputWithContext(context.Background())
}

func (i CniPluginsInstallMap) ToCniPluginsInstallMapOutputWithContext(ctx context.Context) CniPluginsInstallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniPluginsInstallMapOutput)
}

type CniPluginsInstallOutput struct{ *pulumi.OutputState }

func (CniPluginsInstallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CniPluginsInstall)(nil)).Elem()
}

func (o CniPluginsInstallOutput) ToCniPluginsInstallOutput() CniPluginsInstallOutput {
	return o
}

func (o CniPluginsInstallOutput) ToCniPluginsInstallOutputWithContext(ctx context.Context) CniPluginsInstallOutput {
	return o
}

// The CPU architecture.
func (o CniPluginsInstallOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *CniPluginsInstall) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

func (o CniPluginsInstallOutput) ArchiveName() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.ArchiveName }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) BandwidthPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.BandwidthPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) BridgePath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.BridgePath }).(pulumi.StringOutput)
}

// The connection details.
func (o CniPluginsInstallOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

func (o CniPluginsInstallOutput) DhcpPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.DhcpPath }).(pulumi.StringOutput)
}

// Directory to install the binary.
func (o CniPluginsInstallOutput) Directory() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.Directory }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) Download() DownloadOutput {
	return o.ApplyT(func(v *CniPluginsInstall) DownloadOutput { return v.Download }).(DownloadOutput)
}

func (o CniPluginsInstallOutput) DummyPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.DummyPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) FirewallPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.FirewallPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) HostDevicePath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.HostDevicePath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) HostLocalPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.HostLocalPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) IpvlanPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.IpvlanPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) LoopbackPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.LoopbackPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) MacvlanPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.MacvlanPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) Mkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *CniPluginsInstall) tools.MkdirOutput { return v.Mkdir }).(tools.MkdirOutput)
}

func (o CniPluginsInstallOutput) Mktemp() tools.MktempOutput {
	return o.ApplyT(func(v *CniPluginsInstall) tools.MktempOutput { return v.Mktemp }).(tools.MktempOutput)
}

func (o CniPluginsInstallOutput) Mv() tools.MvOutput {
	return o.ApplyT(func(v *CniPluginsInstall) tools.MvOutput { return v.Mv }).(tools.MvOutput)
}

func (o CniPluginsInstallOutput) PortmapPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.PortmapPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) PtpPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.PtpPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) Rm() tools.RmOutput {
	return o.ApplyT(func(v *CniPluginsInstall) tools.RmOutput { return v.Rm }).(tools.RmOutput)
}

func (o CniPluginsInstallOutput) SbrPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.SbrPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) StaticPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.StaticPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) TapPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.TapPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) Tar() tools.TarOutput {
	return o.ApplyT(func(v *CniPluginsInstall) tools.TarOutput { return v.Tar }).(tools.TarOutput)
}

func (o CniPluginsInstallOutput) TuningPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.TuningPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The version to install.
func (o CniPluginsInstallOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) VlanPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.VlanPath }).(pulumi.StringOutput)
}

func (o CniPluginsInstallOutput) VrfPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CniPluginsInstall) pulumi.StringOutput { return v.VrfPath }).(pulumi.StringOutput)
}

type CniPluginsInstallArrayOutput struct{ *pulumi.OutputState }

func (CniPluginsInstallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CniPluginsInstall)(nil)).Elem()
}

func (o CniPluginsInstallArrayOutput) ToCniPluginsInstallArrayOutput() CniPluginsInstallArrayOutput {
	return o
}

func (o CniPluginsInstallArrayOutput) ToCniPluginsInstallArrayOutputWithContext(ctx context.Context) CniPluginsInstallArrayOutput {
	return o
}

func (o CniPluginsInstallArrayOutput) Index(i pulumi.IntInput) CniPluginsInstallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CniPluginsInstall {
		return vs[0].([]*CniPluginsInstall)[vs[1].(int)]
	}).(CniPluginsInstallOutput)
}

type CniPluginsInstallMapOutput struct{ *pulumi.OutputState }

func (CniPluginsInstallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CniPluginsInstall)(nil)).Elem()
}

func (o CniPluginsInstallMapOutput) ToCniPluginsInstallMapOutput() CniPluginsInstallMapOutput {
	return o
}

func (o CniPluginsInstallMapOutput) ToCniPluginsInstallMapOutputWithContext(ctx context.Context) CniPluginsInstallMapOutput {
	return o
}

func (o CniPluginsInstallMapOutput) MapIndex(k pulumi.StringInput) CniPluginsInstallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CniPluginsInstall {
		return vs[0].(map[string]*CniPluginsInstall)[vs[1].(string)]
	}).(CniPluginsInstallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CniPluginsInstallInput)(nil)).Elem(), &CniPluginsInstall{})
	pulumi.RegisterInputType(reflect.TypeOf((*CniPluginsInstallArrayInput)(nil)).Elem(), CniPluginsInstallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CniPluginsInstallMapInput)(nil)).Elem(), CniPluginsInstallMap{})
	pulumi.RegisterOutputType(CniPluginsInstallOutput{})
	pulumi.RegisterOutputType(CniPluginsInstallArrayOutput{})
	pulumi.RegisterOutputType(CniPluginsInstallMapOutput{})
}