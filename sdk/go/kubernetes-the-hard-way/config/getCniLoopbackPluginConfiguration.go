// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Get the `loopback` configuration.
func GetCniLoopbackPluginConfiguration(ctx *pulumi.Context, args *GetCniLoopbackPluginConfigurationArgs, opts ...pulumi.InvokeOption) (*GetCniLoopbackPluginConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCniLoopbackPluginConfigurationResult
	err := ctx.Invoke("kubernetes-the-hard-way:config:getCniLoopbackPluginConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetCniLoopbackPluginConfigurationArgs struct {
	// CNI version.
	CniVersion *string `pulumi:"cniVersion"`
	// CNI plugin name.
	Name *string `pulumi:"name"`
	// Path to put the configuration file on the remote system
	Path *string `pulumi:"path"`
	// CNI plugin type.
	Type *string `pulumi:"type"`
}

// Get the `loopback` configuration.
type GetCniLoopbackPluginConfigurationResult struct {
	Result CniLoopbackPluginConfigurationType `pulumi:"result"`
}

func GetCniLoopbackPluginConfigurationOutput(ctx *pulumi.Context, args GetCniLoopbackPluginConfigurationOutputArgs, opts ...pulumi.InvokeOption) GetCniLoopbackPluginConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCniLoopbackPluginConfigurationResult, error) {
			args := v.(GetCniLoopbackPluginConfigurationArgs)
			r, err := GetCniLoopbackPluginConfiguration(ctx, &args, opts...)
			var s GetCniLoopbackPluginConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCniLoopbackPluginConfigurationResultOutput)
}

type GetCniLoopbackPluginConfigurationOutputArgs struct {
	// CNI version.
	CniVersion pulumi.StringPtrInput `pulumi:"cniVersion"`
	// CNI plugin name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Path to put the configuration file on the remote system
	Path pulumi.StringPtrInput `pulumi:"path"`
	// CNI plugin type.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetCniLoopbackPluginConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCniLoopbackPluginConfigurationArgs)(nil)).Elem()
}

// Get the `loopback` configuration.
type GetCniLoopbackPluginConfigurationResultOutput struct{ *pulumi.OutputState }

func (GetCniLoopbackPluginConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCniLoopbackPluginConfigurationResult)(nil)).Elem()
}

func (o GetCniLoopbackPluginConfigurationResultOutput) ToGetCniLoopbackPluginConfigurationResultOutput() GetCniLoopbackPluginConfigurationResultOutput {
	return o
}

func (o GetCniLoopbackPluginConfigurationResultOutput) ToGetCniLoopbackPluginConfigurationResultOutputWithContext(ctx context.Context) GetCniLoopbackPluginConfigurationResultOutput {
	return o
}

func (o GetCniLoopbackPluginConfigurationResultOutput) Result() CniLoopbackPluginConfigurationTypeOutput {
	return o.ApplyT(func(v GetCniLoopbackPluginConfigurationResult) CniLoopbackPluginConfigurationType { return v.Result }).(CniLoopbackPluginConfigurationTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCniLoopbackPluginConfigurationResultOutput{})
}