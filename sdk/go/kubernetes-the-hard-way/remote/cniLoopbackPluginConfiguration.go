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

// The CNI loopback plugin configuration.
type CniLoopbackPluginConfiguration struct {
	pulumi.CustomResourceState

	// CNI version.
	CniVersion pulumi.StringOutput `pulumi:"cniVersion"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The file on the remote system.
	File FileOutput `pulumi:"file"`
	// CNI plugin name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to put the configuration file on the remote system
	Path pulumi.StringOutput `pulumi:"path"`
	// CNI plugin type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCniLoopbackPluginConfiguration registers a new resource with the given unique name, arguments, and options.
func NewCniLoopbackPluginConfiguration(ctx *pulumi.Context,
	name string, args *CniLoopbackPluginConfigurationArgs, opts ...pulumi.ResourceOption) (*CniLoopbackPluginConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CniLoopbackPluginConfiguration
	err := ctx.RegisterResource("kubernetes-the-hard-way:remote:CniLoopbackPluginConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCniLoopbackPluginConfiguration gets an existing CniLoopbackPluginConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCniLoopbackPluginConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CniLoopbackPluginConfigurationState, opts ...pulumi.ResourceOption) (*CniLoopbackPluginConfiguration, error) {
	var resource CniLoopbackPluginConfiguration
	err := ctx.ReadResource("kubernetes-the-hard-way:remote:CniLoopbackPluginConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CniLoopbackPluginConfiguration resources.
type cniLoopbackPluginConfigurationState struct {
}

type CniLoopbackPluginConfigurationState struct {
}

func (CniLoopbackPluginConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*cniLoopbackPluginConfigurationState)(nil)).Elem()
}

type cniLoopbackPluginConfigurationArgs struct {
	// CNI version.
	CniVersion *string `pulumi:"cniVersion"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// CNI plugin name.
	Name *string `pulumi:"name"`
	// Path to put the configuration file on the remote system
	Path *string `pulumi:"path"`
	// CNI plugin type.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a CniLoopbackPluginConfiguration resource.
type CniLoopbackPluginConfigurationArgs struct {
	// CNI version.
	CniVersion pulumi.StringPtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// CNI plugin name.
	Name pulumi.StringPtrInput
	// Path to put the configuration file on the remote system
	Path pulumi.StringPtrInput
	// CNI plugin type.
	Type pulumi.StringPtrInput
}

func (CniLoopbackPluginConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cniLoopbackPluginConfigurationArgs)(nil)).Elem()
}

type CniLoopbackPluginConfigurationInput interface {
	pulumi.Input

	ToCniLoopbackPluginConfigurationOutput() CniLoopbackPluginConfigurationOutput
	ToCniLoopbackPluginConfigurationOutputWithContext(ctx context.Context) CniLoopbackPluginConfigurationOutput
}

func (*CniLoopbackPluginConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**CniLoopbackPluginConfiguration)(nil)).Elem()
}

func (i *CniLoopbackPluginConfiguration) ToCniLoopbackPluginConfigurationOutput() CniLoopbackPluginConfigurationOutput {
	return i.ToCniLoopbackPluginConfigurationOutputWithContext(context.Background())
}

func (i *CniLoopbackPluginConfiguration) ToCniLoopbackPluginConfigurationOutputWithContext(ctx context.Context) CniLoopbackPluginConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniLoopbackPluginConfigurationOutput)
}

// CniLoopbackPluginConfigurationArrayInput is an input type that accepts CniLoopbackPluginConfigurationArray and CniLoopbackPluginConfigurationArrayOutput values.
// You can construct a concrete instance of `CniLoopbackPluginConfigurationArrayInput` via:
//
//	CniLoopbackPluginConfigurationArray{ CniLoopbackPluginConfigurationArgs{...} }
type CniLoopbackPluginConfigurationArrayInput interface {
	pulumi.Input

	ToCniLoopbackPluginConfigurationArrayOutput() CniLoopbackPluginConfigurationArrayOutput
	ToCniLoopbackPluginConfigurationArrayOutputWithContext(context.Context) CniLoopbackPluginConfigurationArrayOutput
}

type CniLoopbackPluginConfigurationArray []CniLoopbackPluginConfigurationInput

func (CniLoopbackPluginConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CniLoopbackPluginConfiguration)(nil)).Elem()
}

func (i CniLoopbackPluginConfigurationArray) ToCniLoopbackPluginConfigurationArrayOutput() CniLoopbackPluginConfigurationArrayOutput {
	return i.ToCniLoopbackPluginConfigurationArrayOutputWithContext(context.Background())
}

func (i CniLoopbackPluginConfigurationArray) ToCniLoopbackPluginConfigurationArrayOutputWithContext(ctx context.Context) CniLoopbackPluginConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniLoopbackPluginConfigurationArrayOutput)
}

// CniLoopbackPluginConfigurationMapInput is an input type that accepts CniLoopbackPluginConfigurationMap and CniLoopbackPluginConfigurationMapOutput values.
// You can construct a concrete instance of `CniLoopbackPluginConfigurationMapInput` via:
//
//	CniLoopbackPluginConfigurationMap{ "key": CniLoopbackPluginConfigurationArgs{...} }
type CniLoopbackPluginConfigurationMapInput interface {
	pulumi.Input

	ToCniLoopbackPluginConfigurationMapOutput() CniLoopbackPluginConfigurationMapOutput
	ToCniLoopbackPluginConfigurationMapOutputWithContext(context.Context) CniLoopbackPluginConfigurationMapOutput
}

type CniLoopbackPluginConfigurationMap map[string]CniLoopbackPluginConfigurationInput

func (CniLoopbackPluginConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CniLoopbackPluginConfiguration)(nil)).Elem()
}

func (i CniLoopbackPluginConfigurationMap) ToCniLoopbackPluginConfigurationMapOutput() CniLoopbackPluginConfigurationMapOutput {
	return i.ToCniLoopbackPluginConfigurationMapOutputWithContext(context.Background())
}

func (i CniLoopbackPluginConfigurationMap) ToCniLoopbackPluginConfigurationMapOutputWithContext(ctx context.Context) CniLoopbackPluginConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CniLoopbackPluginConfigurationMapOutput)
}

type CniLoopbackPluginConfigurationOutput struct{ *pulumi.OutputState }

func (CniLoopbackPluginConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CniLoopbackPluginConfiguration)(nil)).Elem()
}

func (o CniLoopbackPluginConfigurationOutput) ToCniLoopbackPluginConfigurationOutput() CniLoopbackPluginConfigurationOutput {
	return o
}

func (o CniLoopbackPluginConfigurationOutput) ToCniLoopbackPluginConfigurationOutputWithContext(ctx context.Context) CniLoopbackPluginConfigurationOutput {
	return o
}

// CNI version.
func (o CniLoopbackPluginConfigurationOutput) CniVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CniLoopbackPluginConfiguration) pulumi.StringOutput { return v.CniVersion }).(pulumi.StringOutput)
}

// The parameters with which to connect to the remote host.
func (o CniLoopbackPluginConfigurationOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *CniLoopbackPluginConfiguration) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The file on the remote system.
func (o CniLoopbackPluginConfigurationOutput) File() FileOutput {
	return o.ApplyT(func(v *CniLoopbackPluginConfiguration) FileOutput { return v.File }).(FileOutput)
}

// CNI plugin name.
func (o CniLoopbackPluginConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CniLoopbackPluginConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path to put the configuration file on the remote system
func (o CniLoopbackPluginConfigurationOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *CniLoopbackPluginConfiguration) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// CNI plugin type.
func (o CniLoopbackPluginConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CniLoopbackPluginConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type CniLoopbackPluginConfigurationArrayOutput struct{ *pulumi.OutputState }

func (CniLoopbackPluginConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CniLoopbackPluginConfiguration)(nil)).Elem()
}

func (o CniLoopbackPluginConfigurationArrayOutput) ToCniLoopbackPluginConfigurationArrayOutput() CniLoopbackPluginConfigurationArrayOutput {
	return o
}

func (o CniLoopbackPluginConfigurationArrayOutput) ToCniLoopbackPluginConfigurationArrayOutputWithContext(ctx context.Context) CniLoopbackPluginConfigurationArrayOutput {
	return o
}

func (o CniLoopbackPluginConfigurationArrayOutput) Index(i pulumi.IntInput) CniLoopbackPluginConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CniLoopbackPluginConfiguration {
		return vs[0].([]*CniLoopbackPluginConfiguration)[vs[1].(int)]
	}).(CniLoopbackPluginConfigurationOutput)
}

type CniLoopbackPluginConfigurationMapOutput struct{ *pulumi.OutputState }

func (CniLoopbackPluginConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CniLoopbackPluginConfiguration)(nil)).Elem()
}

func (o CniLoopbackPluginConfigurationMapOutput) ToCniLoopbackPluginConfigurationMapOutput() CniLoopbackPluginConfigurationMapOutput {
	return o
}

func (o CniLoopbackPluginConfigurationMapOutput) ToCniLoopbackPluginConfigurationMapOutputWithContext(ctx context.Context) CniLoopbackPluginConfigurationMapOutput {
	return o
}

func (o CniLoopbackPluginConfigurationMapOutput) MapIndex(k pulumi.StringInput) CniLoopbackPluginConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CniLoopbackPluginConfiguration {
		return vs[0].(map[string]*CniLoopbackPluginConfiguration)[vs[1].(string)]
	}).(CniLoopbackPluginConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CniLoopbackPluginConfigurationInput)(nil)).Elem(), &CniLoopbackPluginConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*CniLoopbackPluginConfigurationArrayInput)(nil)).Elem(), CniLoopbackPluginConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CniLoopbackPluginConfigurationMapInput)(nil)).Elem(), CniLoopbackPluginConfigurationMap{})
	pulumi.RegisterOutputType(CniLoopbackPluginConfigurationOutput{})
	pulumi.RegisterOutputType(CniLoopbackPluginConfigurationArrayOutput{})
	pulumi.RegisterOutputType(CniLoopbackPluginConfigurationMapOutput{})
}
