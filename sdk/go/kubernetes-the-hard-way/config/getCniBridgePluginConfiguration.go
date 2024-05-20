// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// Get the `bridge` configuration.
func GetCniBridgePluginConfiguration(ctx *pulumi.Context, args *GetCniBridgePluginConfigurationArgs, opts ...pulumi.InvokeOption) (*GetCniBridgePluginConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCniBridgePluginConfigurationResult
	err := ctx.Invoke("kubernetes-the-hard-way:config:getCniBridgePluginConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetCniBridgePluginConfigurationArgs struct {
	// Bridge name.
	Bridge *string `pulumi:"bridge"`
	// CNI version.
	CniVersion *string `pulumi:"cniVersion"`
	// IP masq.
	IpMasq *bool `pulumi:"ipMasq"`
	// IPAM
	Ipam *CniBridgeIpam `pulumi:"ipam"`
	// Is gateway.
	IsGateway *bool `pulumi:"isGateway"`
	// CNI plugin name.
	Name *string `pulumi:"name"`
	// Path to put the configuration file on the remote system
	Path *string `pulumi:"path"`
	// The subnet to use.
	Subnet string `pulumi:"subnet"`
	// CNI plugin type.
	Type *string `pulumi:"type"`
}

// Get the `bridge` configuration.
type GetCniBridgePluginConfigurationResult struct {
	Result CniBridgePluginConfigurationType `pulumi:"result"`
}

func GetCniBridgePluginConfigurationOutput(ctx *pulumi.Context, args GetCniBridgePluginConfigurationOutputArgs, opts ...pulumi.InvokeOption) GetCniBridgePluginConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCniBridgePluginConfigurationResult, error) {
			args := v.(GetCniBridgePluginConfigurationArgs)
			r, err := GetCniBridgePluginConfiguration(ctx, &args, opts...)
			var s GetCniBridgePluginConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCniBridgePluginConfigurationResultOutput)
}

type GetCniBridgePluginConfigurationOutputArgs struct {
	// Bridge name.
	Bridge pulumi.StringPtrInput `pulumi:"bridge"`
	// CNI version.
	CniVersion pulumi.StringPtrInput `pulumi:"cniVersion"`
	// IP masq.
	IpMasq pulumi.BoolPtrInput `pulumi:"ipMasq"`
	// IPAM
	Ipam CniBridgeIpamPtrInput `pulumi:"ipam"`
	// Is gateway.
	IsGateway pulumi.BoolPtrInput `pulumi:"isGateway"`
	// CNI plugin name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Path to put the configuration file on the remote system
	Path pulumi.StringPtrInput `pulumi:"path"`
	// The subnet to use.
	Subnet pulumi.StringInput `pulumi:"subnet"`
	// CNI plugin type.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetCniBridgePluginConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCniBridgePluginConfigurationArgs)(nil)).Elem()
}

// Get the `bridge` configuration.
type GetCniBridgePluginConfigurationResultOutput struct{ *pulumi.OutputState }

func (GetCniBridgePluginConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCniBridgePluginConfigurationResult)(nil)).Elem()
}

func (o GetCniBridgePluginConfigurationResultOutput) ToGetCniBridgePluginConfigurationResultOutput() GetCniBridgePluginConfigurationResultOutput {
	return o
}

func (o GetCniBridgePluginConfigurationResultOutput) ToGetCniBridgePluginConfigurationResultOutputWithContext(ctx context.Context) GetCniBridgePluginConfigurationResultOutput {
	return o
}

func (o GetCniBridgePluginConfigurationResultOutput) Result() CniBridgePluginConfigurationTypeOutput {
	return o.ApplyT(func(v GetCniBridgePluginConfigurationResult) CniBridgePluginConfigurationType { return v.Result }).(CniBridgePluginConfigurationTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCniBridgePluginConfigurationResultOutput{})
}