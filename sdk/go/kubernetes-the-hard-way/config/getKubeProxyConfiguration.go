// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

// kube-proxy configuration.
func GetKubeProxyConfiguration(ctx *pulumi.Context, args *GetKubeProxyConfigurationArgs, opts ...pulumi.InvokeOption) (*GetKubeProxyConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubeProxyConfigurationResult
	err := ctx.Invoke("kubernetes-the-hard-way:config:getKubeProxyConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetKubeProxyConfigurationArgs struct {
	// Cluster CIDR.
	ClusterCIDR string `pulumi:"clusterCIDR"`
	// Path to the kubeconfig.
	Kubeconfig string `pulumi:"kubeconfig"`
	// TODO
	Mode *string `pulumi:"mode"`
}

// kube-proxy configuration.
type GetKubeProxyConfigurationResult struct {
	Result KubeProxyConfigurationType `pulumi:"result"`
}

func GetKubeProxyConfigurationOutput(ctx *pulumi.Context, args GetKubeProxyConfigurationOutputArgs, opts ...pulumi.InvokeOption) GetKubeProxyConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubeProxyConfigurationResult, error) {
			args := v.(GetKubeProxyConfigurationArgs)
			r, err := GetKubeProxyConfiguration(ctx, &args, opts...)
			var s GetKubeProxyConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubeProxyConfigurationResultOutput)
}

type GetKubeProxyConfigurationOutputArgs struct {
	// Cluster CIDR.
	ClusterCIDR pulumi.StringInput `pulumi:"clusterCIDR"`
	// Path to the kubeconfig.
	Kubeconfig pulumi.StringInput `pulumi:"kubeconfig"`
	// TODO
	Mode pulumi.StringPtrInput `pulumi:"mode"`
}

func (GetKubeProxyConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeProxyConfigurationArgs)(nil)).Elem()
}

// kube-proxy configuration.
type GetKubeProxyConfigurationResultOutput struct{ *pulumi.OutputState }

func (GetKubeProxyConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeProxyConfigurationResult)(nil)).Elem()
}

func (o GetKubeProxyConfigurationResultOutput) ToGetKubeProxyConfigurationResultOutput() GetKubeProxyConfigurationResultOutput {
	return o
}

func (o GetKubeProxyConfigurationResultOutput) ToGetKubeProxyConfigurationResultOutputWithContext(ctx context.Context) GetKubeProxyConfigurationResultOutput {
	return o
}

func (o GetKubeProxyConfigurationResultOutput) Result() KubeProxyConfigurationTypeOutput {
	return o.ApplyT(func(v GetKubeProxyConfigurationResult) KubeProxyConfigurationType { return v.Result }).(KubeProxyConfigurationTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubeProxyConfigurationResultOutput{})
}
