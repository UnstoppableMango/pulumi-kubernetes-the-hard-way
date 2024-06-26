// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"errors"
	pulumiCommand "github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/config"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/tools"
)

// A Kubernetes worker node.
type WorkerNode struct {
	pulumi.ResourceState

	// The CPU architecture of the node.
	Architecture ArchitectureOutput `pulumi:"architecture"`
	// The path to the cluster certificate authority file.
	CaPath pulumi.StringOutput `pulumi:"caPath"`
	// The CIDR to use for the cluster.
	ClusterCIDR pulumi.StringPtrOutput `pulumi:"clusterCIDR"`
	// The domain for the cluster to use. Defaults to cluster.local.
	ClusterDomain pulumi.StringPtrOutput `pulumi:"clusterDomain"`
	// The CNI bridge plugin configuration.
	CniBridgeConfiguration config.CniBridgePluginConfigurationOutput `pulumi:"cniBridgeConfiguration"`
	// The CNI bridge plugin configuration file.
	CniBridgeConfigurationFile FileOutput `pulumi:"cniBridgeConfigurationFile"`
	// The directory to store CNI plugin configuration files. Defaults to /etc/cni/net.d.
	CniConfigurationDirectory pulumi.StringPtrOutput `pulumi:"cniConfigurationDirectory"`
	// The directory to store CNI plugin binaries. Defaults to /opt/cni/bin.
	CniInstallDirectory pulumi.StringPtrOutput `pulumi:"cniInstallDirectory"`
	// The CNI loopback plugin configuration.
	CniLoopbackConfiguration config.CniLoopbackPluginConfigurationOutput `pulumi:"cniLoopbackConfiguration"`
	// The CNI loopback plugin configuration file.
	CniLoopbackConfigurationFile FileOutput `pulumi:"cniLoopbackConfigurationFile"`
	// The CNI configuration mkdir operation.
	CniMkdir tools.MkdirOutput `pulumi:"cniMkdir"`
	// The CNI plugin install.
	CniPluginsInstall CniPluginsInstallOutput `pulumi:"cniPluginsInstall"`
	// The CNI version to use.
	CniVersion pulumi.StringPtrOutput `pulumi:"cniVersion"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionOutput `pulumi:"connection"`
	// The containerd configuration.
	ContainerdConfiguration config.ContainerdConfigurationOutput `pulumi:"containerdConfiguration"`
	// The directory to store containerd configuration files. Defaults to /etc/containerd.
	ContainerdConfigurationDirectory pulumi.StringPtrOutput `pulumi:"containerdConfigurationDirectory"`
	// The containerd configuration file.
	ContainerdConfigurationFile FileOutput `pulumi:"containerdConfigurationFile"`
	// The containerd install.
	ContainerdInstall ContainerdInstallOutput `pulumi:"containerdInstall"`
	// The directory to store the containerd binary. Defaults to /bin.
	ContainerdInstallDirectory pulumi.StringPtrOutput `pulumi:"containerdInstallDirectory"`
	// The containerd configuration mkdir operation.
	ContainerdMkdir tools.MkdirOutput `pulumi:"containerdMkdir"`
	// The containerd systemd service.
	ContainerdService ContainerdServiceOutput `pulumi:"containerdService"`
	// The containerd version to use.
	ContainerdVersion pulumi.StringPtrOutput `pulumi:"containerdVersion"`
	// The crictl install.
	CrictlInstall CrictlInstallOutput `pulumi:"crictlInstall"`
	// The directory to store the crictl binary. Defaults to /usr/local/bin.
	CrictlInstallDirectory pulumi.StringPtrOutput `pulumi:"crictlInstallDirectory"`
	// The kube-proxy configuration
	KubeProxyConfiguration config.KubeProxyConfigurationOutput `pulumi:"kubeProxyConfiguration"`
	// The directory to store kube-proxy configuration files. Defaults to /var/lib/kube-proxy.
	KubeProxyConfigurationDirectory pulumi.StringPtrOutput `pulumi:"kubeProxyConfigurationDirectory"`
	// The kube-proxy configuration file.
	KubeProxyConfigurationFile FileOutput `pulumi:"kubeProxyConfigurationFile"`
	// The kube-proxy install.
	KubeProxyInstall KubeProxyInstallOutput `pulumi:"kubeProxyInstall"`
	// The directory to store the kube-proxy binary. Defaults to /usr/local/bin.
	KubeProxyInstallDirectory pulumi.StringPtrOutput `pulumi:"kubeProxyInstallDirectory"`
	// The path to the kube-proxy's kubeconfig file.
	KubeProxyKubeconfigPath pulumi.StringPtrOutput `pulumi:"kubeProxyKubeconfigPath"`
	// The kube-proxy configuration mkdir operation.
	KubeProxyMkdir tools.MkdirOutput `pulumi:"kubeProxyMkdir"`
	// The kubelet systemd service.
	KubeProxyService KubeProxyServiceOutput `pulumi:"kubeProxyService"`
	// The kubectl install.
	KubectlInstall KubectlInstallOutput `pulumi:"kubectlInstall"`
	// The directory to store the kubectl binary. Defaults to /usr/local/bin.
	KubectlInstallDirectory pulumi.StringPtrOutput `pulumi:"kubectlInstallDirectory"`
	// The path to the kubelet certificate.
	KubeletCertificatePath pulumi.StringOutput `pulumi:"kubeletCertificatePath"`
	// The kubelet configuration
	KubeletConfiguration config.KubeletConfigurationOutput `pulumi:"kubeletConfiguration"`
	// The directory to store kubelet configuration files. Defaults to /var/lib/kubelet.
	KubeletConfigurationDirectory pulumi.StringPtrOutput `pulumi:"kubeletConfigurationDirectory"`
	// The kubelet configuration file.
	KubeletConfigurationFile FileOutput `pulumi:"kubeletConfigurationFile"`
	// The kubelet install.
	KubeletInstall KubeletInstallOutput `pulumi:"kubeletInstall"`
	// The directory to store the kubelet binary. Defaults to /usr/local/bin.
	KubeletInstallDirectory pulumi.StringPtrOutput `pulumi:"kubeletInstallDirectory"`
	// The path to the kubelet's kubeconfig file.
	KubeletKubeconfigPath pulumi.StringPtrOutput `pulumi:"kubeletKubeconfigPath"`
	// The kubelet configuration mkdir operation.
	KubeletMkdir tools.MkdirOutput `pulumi:"kubeletMkdir"`
	// The path to the kubelet private key file.
	KubeletPrivateKeyPath pulumi.StringOutput `pulumi:"kubeletPrivateKeyPath"`
	// The kubelet systemd service.
	KubeletService KubeletServiceOutput `pulumi:"kubeletService"`
	// The kubernetes version to use.
	KubernetesVersion pulumi.StringPtrOutput `pulumi:"kubernetesVersion"`
	// The runc install.
	RuncInstall RuncInstallOutput `pulumi:"runcInstall"`
	// The subnet for the cluster.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// The /var/lib/kubernetes mkdir operation.
	VarLibKubernetesMkdir tools.MkdirOutput `pulumi:"varLibKubernetesMkdir"`
	// The /var/run/kubernetes mkdir operation.
	VarRunKubernetesMkdir tools.MkdirOutput `pulumi:"varRunKubernetesMkdir"`
}

// NewWorkerNode registers a new resource with the given unique name, arguments, and options.
func NewWorkerNode(ctx *pulumi.Context,
	name string, args *WorkerNodeArgs, opts ...pulumi.ResourceOption) (*WorkerNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Architecture == nil {
		return nil, errors.New("invalid value for required argument 'Architecture'")
	}
	if args.CaPath == nil {
		return nil, errors.New("invalid value for required argument 'CaPath'")
	}
	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	if args.KubeletCertificatePath == nil {
		return nil, errors.New("invalid value for required argument 'KubeletCertificatePath'")
	}
	if args.KubeletPrivateKeyPath == nil {
		return nil, errors.New("invalid value for required argument 'KubeletPrivateKeyPath'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	args.Connection = args.Connection.ToConnectionOutput().ApplyT(func(v pulumiCommand.Connection) pulumiCommand.Connection { return *v.Defaults() }).(pulumiCommand.ConnectionOutput)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkerNode
	err := ctx.RegisterRemoteComponentResource("kubernetes-the-hard-way:remote:WorkerNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type workerNodeArgs struct {
	// The CPU architecture of the node.
	Architecture Architecture `pulumi:"architecture"`
	// The path to the cluster certificate authority file.
	CaPath string `pulumi:"caPath"`
	// The CIDR to use for the cluster.
	ClusterCIDR *string `pulumi:"clusterCIDR"`
	// The domain for the cluster to use. Defaults to cluster.local.
	ClusterDomain *string `pulumi:"clusterDomain"`
	// The directory to store CNI plugin configuration files. Defaults to /etc/cni/net.d.
	CniConfigurationDirectory *string `pulumi:"cniConfigurationDirectory"`
	// The directory to store CNI plugin binaries. Defaults to /opt/cni/bin.
	CniInstallDirectory *string `pulumi:"cniInstallDirectory"`
	// The CNI version to use.
	CniVersion *string `pulumi:"cniVersion"`
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.Connection `pulumi:"connection"`
	// The directory to store containerd configuration files. Defaults to /etc/containerd.
	ContainerdConfigurationDirectory *string `pulumi:"containerdConfigurationDirectory"`
	// The directory to store the containerd binary. Defaults to /bin.
	ContainerdInstallDirectory *string `pulumi:"containerdInstallDirectory"`
	// The containerd version to use.
	ContainerdVersion *string `pulumi:"containerdVersion"`
	// The directory to store the crictl binary. Defaults to /usr/local/bin.
	CrictlInstallDirectory *string `pulumi:"crictlInstallDirectory"`
	// The directory to store kube-proxy configuration files. Defaults to /var/lib/kube-proxy.
	KubeProxyConfigurationDirectory *string `pulumi:"kubeProxyConfigurationDirectory"`
	// The directory to store the kube-proxy binary. Defaults to /usr/local/bin.
	KubeProxyInstallDirectory *string `pulumi:"kubeProxyInstallDirectory"`
	// The path to the kube-proxy's kubeconfig file.
	KubeProxyKubeconfigPath *string `pulumi:"kubeProxyKubeconfigPath"`
	// The directory to store the kubectl binary. Defaults to /usr/local/bin.
	KubectlInstallDirectory *string `pulumi:"kubectlInstallDirectory"`
	// The path to the kubelet certificate.
	KubeletCertificatePath string `pulumi:"kubeletCertificatePath"`
	// The directory to store kubelet configuration files. Defaults to /var/lib/kubelet.
	KubeletConfigurationDirectory *string `pulumi:"kubeletConfigurationDirectory"`
	// The directory to store the kubelet binary. Defaults to /usr/local/bin.
	KubeletInstallDirectory *string `pulumi:"kubeletInstallDirectory"`
	// The path to the kubelet's kubeconfig file.
	KubeletKubeconfigPath *string `pulumi:"kubeletKubeconfigPath"`
	// The path to the kubelet private key file.
	KubeletPrivateKeyPath string `pulumi:"kubeletPrivateKeyPath"`
	// The kubernetes version to use.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// The subnet for the cluster.
	Subnet string `pulumi:"subnet"`
}

// The set of arguments for constructing a WorkerNode resource.
type WorkerNodeArgs struct {
	// The CPU architecture of the node.
	Architecture ArchitectureInput
	// The path to the cluster certificate authority file.
	CaPath pulumi.StringInput
	// The CIDR to use for the cluster.
	ClusterCIDR pulumi.StringPtrInput
	// The domain for the cluster to use. Defaults to cluster.local.
	ClusterDomain pulumi.StringPtrInput
	// The directory to store CNI plugin configuration files. Defaults to /etc/cni/net.d.
	CniConfigurationDirectory pulumi.StringPtrInput
	// The directory to store CNI plugin binaries. Defaults to /opt/cni/bin.
	CniInstallDirectory pulumi.StringPtrInput
	// The CNI version to use.
	CniVersion pulumi.StringPtrInput
	// The parameters with which to connect to the remote host.
	Connection pulumiCommand.ConnectionInput
	// The directory to store containerd configuration files. Defaults to /etc/containerd.
	ContainerdConfigurationDirectory pulumi.StringPtrInput
	// The directory to store the containerd binary. Defaults to /bin.
	ContainerdInstallDirectory pulumi.StringPtrInput
	// The containerd version to use.
	ContainerdVersion pulumi.StringPtrInput
	// The directory to store the crictl binary. Defaults to /usr/local/bin.
	CrictlInstallDirectory pulumi.StringPtrInput
	// The directory to store kube-proxy configuration files. Defaults to /var/lib/kube-proxy.
	KubeProxyConfigurationDirectory pulumi.StringPtrInput
	// The directory to store the kube-proxy binary. Defaults to /usr/local/bin.
	KubeProxyInstallDirectory pulumi.StringPtrInput
	// The path to the kube-proxy's kubeconfig file.
	KubeProxyKubeconfigPath pulumi.StringPtrInput
	// The directory to store the kubectl binary. Defaults to /usr/local/bin.
	KubectlInstallDirectory pulumi.StringPtrInput
	// The path to the kubelet certificate.
	KubeletCertificatePath pulumi.StringInput
	// The directory to store kubelet configuration files. Defaults to /var/lib/kubelet.
	KubeletConfigurationDirectory pulumi.StringPtrInput
	// The directory to store the kubelet binary. Defaults to /usr/local/bin.
	KubeletInstallDirectory pulumi.StringPtrInput
	// The path to the kubelet's kubeconfig file.
	KubeletKubeconfigPath pulumi.StringPtrInput
	// The path to the kubelet private key file.
	KubeletPrivateKeyPath pulumi.StringInput
	// The kubernetes version to use.
	KubernetesVersion pulumi.StringPtrInput
	// The subnet for the cluster.
	Subnet pulumi.StringInput
}

func (WorkerNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workerNodeArgs)(nil)).Elem()
}

type WorkerNodeInput interface {
	pulumi.Input

	ToWorkerNodeOutput() WorkerNodeOutput
	ToWorkerNodeOutputWithContext(ctx context.Context) WorkerNodeOutput
}

func (*WorkerNode) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkerNode)(nil)).Elem()
}

func (i *WorkerNode) ToWorkerNodeOutput() WorkerNodeOutput {
	return i.ToWorkerNodeOutputWithContext(context.Background())
}

func (i *WorkerNode) ToWorkerNodeOutputWithContext(ctx context.Context) WorkerNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerNodeOutput)
}

// WorkerNodeArrayInput is an input type that accepts WorkerNodeArray and WorkerNodeArrayOutput values.
// You can construct a concrete instance of `WorkerNodeArrayInput` via:
//
//	WorkerNodeArray{ WorkerNodeArgs{...} }
type WorkerNodeArrayInput interface {
	pulumi.Input

	ToWorkerNodeArrayOutput() WorkerNodeArrayOutput
	ToWorkerNodeArrayOutputWithContext(context.Context) WorkerNodeArrayOutput
}

type WorkerNodeArray []WorkerNodeInput

func (WorkerNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkerNode)(nil)).Elem()
}

func (i WorkerNodeArray) ToWorkerNodeArrayOutput() WorkerNodeArrayOutput {
	return i.ToWorkerNodeArrayOutputWithContext(context.Background())
}

func (i WorkerNodeArray) ToWorkerNodeArrayOutputWithContext(ctx context.Context) WorkerNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerNodeArrayOutput)
}

// WorkerNodeMapInput is an input type that accepts WorkerNodeMap and WorkerNodeMapOutput values.
// You can construct a concrete instance of `WorkerNodeMapInput` via:
//
//	WorkerNodeMap{ "key": WorkerNodeArgs{...} }
type WorkerNodeMapInput interface {
	pulumi.Input

	ToWorkerNodeMapOutput() WorkerNodeMapOutput
	ToWorkerNodeMapOutputWithContext(context.Context) WorkerNodeMapOutput
}

type WorkerNodeMap map[string]WorkerNodeInput

func (WorkerNodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkerNode)(nil)).Elem()
}

func (i WorkerNodeMap) ToWorkerNodeMapOutput() WorkerNodeMapOutput {
	return i.ToWorkerNodeMapOutputWithContext(context.Background())
}

func (i WorkerNodeMap) ToWorkerNodeMapOutputWithContext(ctx context.Context) WorkerNodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerNodeMapOutput)
}

type WorkerNodeOutput struct{ *pulumi.OutputState }

func (WorkerNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkerNode)(nil)).Elem()
}

func (o WorkerNodeOutput) ToWorkerNodeOutput() WorkerNodeOutput {
	return o
}

func (o WorkerNodeOutput) ToWorkerNodeOutputWithContext(ctx context.Context) WorkerNodeOutput {
	return o
}

// The CPU architecture of the node.
func (o WorkerNodeOutput) Architecture() ArchitectureOutput {
	return o.ApplyT(func(v *WorkerNode) ArchitectureOutput { return v.Architecture }).(ArchitectureOutput)
}

// The path to the cluster certificate authority file.
func (o WorkerNodeOutput) CaPath() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringOutput { return v.CaPath }).(pulumi.StringOutput)
}

// The CIDR to use for the cluster.
func (o WorkerNodeOutput) ClusterCIDR() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.ClusterCIDR }).(pulumi.StringPtrOutput)
}

// The domain for the cluster to use. Defaults to cluster.local.
func (o WorkerNodeOutput) ClusterDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.ClusterDomain }).(pulumi.StringPtrOutput)
}

// The CNI bridge plugin configuration.
func (o WorkerNodeOutput) CniBridgeConfiguration() config.CniBridgePluginConfigurationOutput {
	return o.ApplyT(func(v *WorkerNode) config.CniBridgePluginConfigurationOutput { return v.CniBridgeConfiguration }).(config.CniBridgePluginConfigurationOutput)
}

// The CNI bridge plugin configuration file.
func (o WorkerNodeOutput) CniBridgeConfigurationFile() FileOutput {
	return o.ApplyT(func(v *WorkerNode) FileOutput { return v.CniBridgeConfigurationFile }).(FileOutput)
}

// The directory to store CNI plugin configuration files. Defaults to /etc/cni/net.d.
func (o WorkerNodeOutput) CniConfigurationDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.CniConfigurationDirectory }).(pulumi.StringPtrOutput)
}

// The directory to store CNI plugin binaries. Defaults to /opt/cni/bin.
func (o WorkerNodeOutput) CniInstallDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.CniInstallDirectory }).(pulumi.StringPtrOutput)
}

// The CNI loopback plugin configuration.
func (o WorkerNodeOutput) CniLoopbackConfiguration() config.CniLoopbackPluginConfigurationOutput {
	return o.ApplyT(func(v *WorkerNode) config.CniLoopbackPluginConfigurationOutput { return v.CniLoopbackConfiguration }).(config.CniLoopbackPluginConfigurationOutput)
}

// The CNI loopback plugin configuration file.
func (o WorkerNodeOutput) CniLoopbackConfigurationFile() FileOutput {
	return o.ApplyT(func(v *WorkerNode) FileOutput { return v.CniLoopbackConfigurationFile }).(FileOutput)
}

// The CNI configuration mkdir operation.
func (o WorkerNodeOutput) CniMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *WorkerNode) tools.MkdirOutput { return v.CniMkdir }).(tools.MkdirOutput)
}

// The CNI plugin install.
func (o WorkerNodeOutput) CniPluginsInstall() CniPluginsInstallOutput {
	return o.ApplyT(func(v *WorkerNode) CniPluginsInstallOutput { return v.CniPluginsInstall }).(CniPluginsInstallOutput)
}

// The CNI version to use.
func (o WorkerNodeOutput) CniVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.CniVersion }).(pulumi.StringPtrOutput)
}

// The parameters with which to connect to the remote host.
func (o WorkerNodeOutput) Connection() pulumiCommand.ConnectionOutput {
	return o.ApplyT(func(v *WorkerNode) pulumiCommand.ConnectionOutput { return v.Connection }).(pulumiCommand.ConnectionOutput)
}

// The containerd configuration.
func (o WorkerNodeOutput) ContainerdConfiguration() config.ContainerdConfigurationOutput {
	return o.ApplyT(func(v *WorkerNode) config.ContainerdConfigurationOutput { return v.ContainerdConfiguration }).(config.ContainerdConfigurationOutput)
}

// The directory to store containerd configuration files. Defaults to /etc/containerd.
func (o WorkerNodeOutput) ContainerdConfigurationDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.ContainerdConfigurationDirectory }).(pulumi.StringPtrOutput)
}

// The containerd configuration file.
func (o WorkerNodeOutput) ContainerdConfigurationFile() FileOutput {
	return o.ApplyT(func(v *WorkerNode) FileOutput { return v.ContainerdConfigurationFile }).(FileOutput)
}

// The containerd install.
func (o WorkerNodeOutput) ContainerdInstall() ContainerdInstallOutput {
	return o.ApplyT(func(v *WorkerNode) ContainerdInstallOutput { return v.ContainerdInstall }).(ContainerdInstallOutput)
}

// The directory to store the containerd binary. Defaults to /bin.
func (o WorkerNodeOutput) ContainerdInstallDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.ContainerdInstallDirectory }).(pulumi.StringPtrOutput)
}

// The containerd configuration mkdir operation.
func (o WorkerNodeOutput) ContainerdMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *WorkerNode) tools.MkdirOutput { return v.ContainerdMkdir }).(tools.MkdirOutput)
}

// The containerd systemd service.
func (o WorkerNodeOutput) ContainerdService() ContainerdServiceOutput {
	return o.ApplyT(func(v *WorkerNode) ContainerdServiceOutput { return v.ContainerdService }).(ContainerdServiceOutput)
}

// The containerd version to use.
func (o WorkerNodeOutput) ContainerdVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.ContainerdVersion }).(pulumi.StringPtrOutput)
}

// The crictl install.
func (o WorkerNodeOutput) CrictlInstall() CrictlInstallOutput {
	return o.ApplyT(func(v *WorkerNode) CrictlInstallOutput { return v.CrictlInstall }).(CrictlInstallOutput)
}

// The directory to store the crictl binary. Defaults to /usr/local/bin.
func (o WorkerNodeOutput) CrictlInstallDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.CrictlInstallDirectory }).(pulumi.StringPtrOutput)
}

// The kube-proxy configuration
func (o WorkerNodeOutput) KubeProxyConfiguration() config.KubeProxyConfigurationOutput {
	return o.ApplyT(func(v *WorkerNode) config.KubeProxyConfigurationOutput { return v.KubeProxyConfiguration }).(config.KubeProxyConfigurationOutput)
}

// The directory to store kube-proxy configuration files. Defaults to /var/lib/kube-proxy.
func (o WorkerNodeOutput) KubeProxyConfigurationDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubeProxyConfigurationDirectory }).(pulumi.StringPtrOutput)
}

// The kube-proxy configuration file.
func (o WorkerNodeOutput) KubeProxyConfigurationFile() FileOutput {
	return o.ApplyT(func(v *WorkerNode) FileOutput { return v.KubeProxyConfigurationFile }).(FileOutput)
}

// The kube-proxy install.
func (o WorkerNodeOutput) KubeProxyInstall() KubeProxyInstallOutput {
	return o.ApplyT(func(v *WorkerNode) KubeProxyInstallOutput { return v.KubeProxyInstall }).(KubeProxyInstallOutput)
}

// The directory to store the kube-proxy binary. Defaults to /usr/local/bin.
func (o WorkerNodeOutput) KubeProxyInstallDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubeProxyInstallDirectory }).(pulumi.StringPtrOutput)
}

// The path to the kube-proxy's kubeconfig file.
func (o WorkerNodeOutput) KubeProxyKubeconfigPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubeProxyKubeconfigPath }).(pulumi.StringPtrOutput)
}

// The kube-proxy configuration mkdir operation.
func (o WorkerNodeOutput) KubeProxyMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *WorkerNode) tools.MkdirOutput { return v.KubeProxyMkdir }).(tools.MkdirOutput)
}

// The kubelet systemd service.
func (o WorkerNodeOutput) KubeProxyService() KubeProxyServiceOutput {
	return o.ApplyT(func(v *WorkerNode) KubeProxyServiceOutput { return v.KubeProxyService }).(KubeProxyServiceOutput)
}

// The kubectl install.
func (o WorkerNodeOutput) KubectlInstall() KubectlInstallOutput {
	return o.ApplyT(func(v *WorkerNode) KubectlInstallOutput { return v.KubectlInstall }).(KubectlInstallOutput)
}

// The directory to store the kubectl binary. Defaults to /usr/local/bin.
func (o WorkerNodeOutput) KubectlInstallDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubectlInstallDirectory }).(pulumi.StringPtrOutput)
}

// The path to the kubelet certificate.
func (o WorkerNodeOutput) KubeletCertificatePath() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringOutput { return v.KubeletCertificatePath }).(pulumi.StringOutput)
}

// The kubelet configuration
func (o WorkerNodeOutput) KubeletConfiguration() config.KubeletConfigurationOutput {
	return o.ApplyT(func(v *WorkerNode) config.KubeletConfigurationOutput { return v.KubeletConfiguration }).(config.KubeletConfigurationOutput)
}

// The directory to store kubelet configuration files. Defaults to /var/lib/kubelet.
func (o WorkerNodeOutput) KubeletConfigurationDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubeletConfigurationDirectory }).(pulumi.StringPtrOutput)
}

// The kubelet configuration file.
func (o WorkerNodeOutput) KubeletConfigurationFile() FileOutput {
	return o.ApplyT(func(v *WorkerNode) FileOutput { return v.KubeletConfigurationFile }).(FileOutput)
}

// The kubelet install.
func (o WorkerNodeOutput) KubeletInstall() KubeletInstallOutput {
	return o.ApplyT(func(v *WorkerNode) KubeletInstallOutput { return v.KubeletInstall }).(KubeletInstallOutput)
}

// The directory to store the kubelet binary. Defaults to /usr/local/bin.
func (o WorkerNodeOutput) KubeletInstallDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubeletInstallDirectory }).(pulumi.StringPtrOutput)
}

// The path to the kubelet's kubeconfig file.
func (o WorkerNodeOutput) KubeletKubeconfigPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubeletKubeconfigPath }).(pulumi.StringPtrOutput)
}

// The kubelet configuration mkdir operation.
func (o WorkerNodeOutput) KubeletMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *WorkerNode) tools.MkdirOutput { return v.KubeletMkdir }).(tools.MkdirOutput)
}

// The path to the kubelet private key file.
func (o WorkerNodeOutput) KubeletPrivateKeyPath() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringOutput { return v.KubeletPrivateKeyPath }).(pulumi.StringOutput)
}

// The kubelet systemd service.
func (o WorkerNodeOutput) KubeletService() KubeletServiceOutput {
	return o.ApplyT(func(v *WorkerNode) KubeletServiceOutput { return v.KubeletService }).(KubeletServiceOutput)
}

// The kubernetes version to use.
func (o WorkerNodeOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringPtrOutput { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

// The runc install.
func (o WorkerNodeOutput) RuncInstall() RuncInstallOutput {
	return o.ApplyT(func(v *WorkerNode) RuncInstallOutput { return v.RuncInstall }).(RuncInstallOutput)
}

// The subnet for the cluster.
func (o WorkerNodeOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkerNode) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// The /var/lib/kubernetes mkdir operation.
func (o WorkerNodeOutput) VarLibKubernetesMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *WorkerNode) tools.MkdirOutput { return v.VarLibKubernetesMkdir }).(tools.MkdirOutput)
}

// The /var/run/kubernetes mkdir operation.
func (o WorkerNodeOutput) VarRunKubernetesMkdir() tools.MkdirOutput {
	return o.ApplyT(func(v *WorkerNode) tools.MkdirOutput { return v.VarRunKubernetesMkdir }).(tools.MkdirOutput)
}

type WorkerNodeArrayOutput struct{ *pulumi.OutputState }

func (WorkerNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkerNode)(nil)).Elem()
}

func (o WorkerNodeArrayOutput) ToWorkerNodeArrayOutput() WorkerNodeArrayOutput {
	return o
}

func (o WorkerNodeArrayOutput) ToWorkerNodeArrayOutputWithContext(ctx context.Context) WorkerNodeArrayOutput {
	return o
}

func (o WorkerNodeArrayOutput) Index(i pulumi.IntInput) WorkerNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkerNode {
		return vs[0].([]*WorkerNode)[vs[1].(int)]
	}).(WorkerNodeOutput)
}

type WorkerNodeMapOutput struct{ *pulumi.OutputState }

func (WorkerNodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkerNode)(nil)).Elem()
}

func (o WorkerNodeMapOutput) ToWorkerNodeMapOutput() WorkerNodeMapOutput {
	return o
}

func (o WorkerNodeMapOutput) ToWorkerNodeMapOutputWithContext(ctx context.Context) WorkerNodeMapOutput {
	return o
}

func (o WorkerNodeMapOutput) MapIndex(k pulumi.StringInput) WorkerNodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkerNode {
		return vs[0].(map[string]*WorkerNode)[vs[1].(string)]
	}).(WorkerNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkerNodeInput)(nil)).Elem(), &WorkerNode{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkerNodeArrayInput)(nil)).Elem(), WorkerNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkerNodeMapInput)(nil)).Elem(), WorkerNodeMap{})
	pulumi.RegisterOutputType(WorkerNodeOutput{})
	pulumi.RegisterOutputType(WorkerNodeArrayOutput{})
	pulumi.RegisterOutputType(WorkerNodeMapOutput{})
}
