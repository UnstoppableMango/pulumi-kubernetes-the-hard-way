package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateWorkerNode(commandSpec, commandxSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"architecture": {
			Description: "The CPU architecture of the node.",
			TypeSpec:    types.LocalType("Architecture", "remote"),
		},
		"caPath":                           props.String("The path to the cluster certificate authority file."),
		"clusterCIDR":                      props.String("The CIDR to use for the cluster."),
		"clusterDomain":                    props.String("The domain for the cluster to use. Defaults to cluster.local."),
		"cniConfigurationDirectory":        props.String("The directory to store CNI plugin configuration files. Defaults to /etc/cni/net.d."),
		"cniInstallDirectory":              props.String("The directory to store CNI plugin binaries. Defaults to /opt/cni/bin."),
		"cniVersion":                       props.String("The CNI version to use."),
		"connection":                       props.Connection(commandSpec),
		"containerdConfigurationDirectory": props.String("The directory to store containerd configuration files. Defaults to /etc/containerd."),
		"containerdInstallDirectory":       props.String("The directory to store the containerd binary. Defaults to /bin."),
		"containerdVersion":                props.String("The containerd version to use."),
		"crictlInstallDirectory":           props.String("The directory to store the crictl binary. Defaults to /usr/local/bin."),
		"kubectlInstallDirectory":          props.String("The directory to store the kubectl binary. Defaults to /usr/local/bin."),
		"kubeletCertificatePath":           props.String("The path to the kubelet certificate."),
		"kubeletConfigurationDirectory":    props.String("The directory to store kubelet configuration files. Defaults to /var/lib/kubelet."),
		"kubeletInstallDirectory":          props.String("The directory to store the kubelet binary. Defaults to /usr/local/bin."),
		"kubeletKubeconfigPath":            props.String("The path to the kubelet's kubeconfig file."),
		"kubeletPrivateKeyPath":            props.String("The path to the kubelet private key file."),
		"kubeProxyConfigurationDirectory":  props.String("The directory to store kube-proxy configuration files. Defaults to /var/lib/kube-proxy."),
		"kubeProxyInstallDirectory":        props.String("The directory to store the kube-proxy binary. Defaults to /usr/local/bin."),
		"kubeProxyKubeconfigPath":          props.String("The path to the kube-proxy's kubeconfig file."),
		"kubernetesVersion":                props.String("The kubernetes version to use."),
		"subnet":                           props.String("The subnet for the cluster."),
	}

	requiredInputs := []string{
		"architecture",
		"caPath",
		"connection",
		"kubeletCertificatePath",
		"kubeletPrivateKeyPath",
		"subnet",
	}

	outputs := map[string]schema.PropertySpec{
		"cniPluginsInstall": {
			Description: "The CNI plugin install.",
			TypeSpec:    types.LocalResource("CniPluginsInstall", "remote"),
		},
		"cniBridgeConfiguration": {
			Description: "The CNI bridge plugin configuration.",
			TypeSpec:    types.LocalResource("CniBridgePluginConfiguration", "config"),
		},
		"cniBridgeConfigurationFile": {
			Description: "The CNI bridge plugin configuration file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"cniLoopbackConfiguration": {
			Description: "The CNI loopback plugin configuration.",
			TypeSpec:    types.LocalResource("CniLoopbackPluginConfiguration", "config"),
		},
		"cniLoopbackConfigurationFile": {
			Description: "The CNI loopback plugin configuration file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"cniMkdir": {
			Description: "The CNI configuration mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"containerdConfiguration": {
			Description: "The containerd configuration.",
			TypeSpec:    types.LocalResource("ContainerdConfiguration", "config"),
		},
		"containerdConfigurationFile": {
			Description: "The containerd configuration file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"containerdMkdir": {
			Description: "The containerd configuration mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"containerdInstall": {
			Description: "The containerd install.",
			TypeSpec:    types.LocalResource("ContainerdInstall", "remote"),
		},
		"containerdService": {
			Description: "The containerd systemd service.",
			TypeSpec:    types.LocalResource("ContainerdService", "remote"),
		},
		"crictlInstall": {
			Description: "The crictl install.",
			TypeSpec:    types.LocalResource("CrictlInstall", "remote"),
		},
		"kubectlInstall": {
			Description: "The kubectl install.",
			TypeSpec:    types.LocalResource("KubectlInstall", "remote"),
		},
		"kubeletConfiguration": {
			Description: "The kubelet configuration",
			TypeSpec:    types.LocalResource("KubeletConfiguration", "config"),
		},
		"kubeletConfigurationFile": {
			Description: "The kubelet configuration file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"kubeletInstall": {
			Description: "The kubelet install.",
			TypeSpec:    types.LocalResource("KubeletInstall", "remote"),
		},
		"kubeletMkdir": {
			Description: "The kubelet configuration mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"kubeletService": {
			Description: "The kubelet systemd service.",
			TypeSpec:    types.LocalResource("KubeletService", "remote"),
		},
		"kubeProxyConfiguration": {
			Description: "The kube-proxy configuration",
			TypeSpec:    types.LocalResource("KubeProxyConfiguration", "config"),
		},
		"kubeProxyConfigurationFile": {
			Description: "The kube-proxy configuration file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"kubeProxyInstall": {
			Description: "The kube-proxy install.",
			TypeSpec:    types.LocalResource("KubeProxyInstall", "remote"),
		},
		"kubeProxyMkdir": {
			Description: "The kube-proxy configuration mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"kubeProxyService": {
			Description: "The kubelet systemd service.",
			TypeSpec:    types.LocalResource("KubeProxyService", "remote"),
		},
		"runcInstall": {
			Description: "The runc install.",
			TypeSpec:    types.LocalResource("RuncInstall", "remote"),
		},
		"varLibKubernetesMkdir": {
			Description: "The /var/lib/kubernetes mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"varRunKubernetesMkdir": {
			Description: "The /var/run/kubernetes mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"cniMkdir",
		"cniBridgeConfiguration",
		"cniBridgeConfigurationFile",
		"cniLoopbackConfiguration",
		"cniLoopbackConfigurationFile",
		"containerdConfiguration",
		"containerdConfigurationFile",
		"containerdMkdir",
		"containerdInstall",
		"containerdService",
		"crictlInstall",
		"kubectlInstall",
		"kubeletConfiguration",
		"kubeletConfigurationFile",
		"kubeletInstall",
		"kubeletMkdir",
		"kubeletService",
		"kubeProxyConfiguration",
		"kubeProxyConfigurationFile",
		"kubeProxyMkdir",
		"kubeProxyService",
		"varLibKubernetesMkdir",
		"varRunKubernetesMkdir",
	})

	return schema.PackageSpec{
		Types:     map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("WorkerNode"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "A Kubernetes worker node.",
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}
