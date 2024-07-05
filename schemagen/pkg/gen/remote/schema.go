package remote

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/internal"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func Generate(commandSpec, commandxSpec schema.PackageSpec) schema.PackageSpec {
	base := schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{},
		Types: map[string]schema.ComplexTypeSpec{
			name("Architecture"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "CPU architecture",
					Type:        "string",
				},
				Enum: []schema.EnumValueSpec{
					{Value: "amd64"},
					{Value: "arm64"},
				},
			},
		},
	}

	return internal.ExtendSchemas(base,
		generateDownload(commandSpec, commandxSpec),
		generateContainerdService(commandSpec),
		generateControlPlaneNode(commandSpec, commandxSpec),
		generateEtcdCluster(commandSpec),
		generateEtcdConfiguration(commandSpec, commandxSpec),
		generateEtcdService(commandSpec),
		generateFile(commandSpec),
		generateKubeletService(commandSpec),
		generateKubeProxyService(commandSpec),
		generateProvisionEtcd(commandSpec),
		generateStartSystemdService(commandSpec, commandxSpec, "StartContainerd", "containerd"),
		generateStartSystemdService(commandSpec, commandxSpec, "StartEtcd", "etcd"),
		generateStartSystemdService(commandSpec, commandxSpec, "StartKubelet", "kubelet"),
		generateStartSystemdService(commandSpec, commandxSpec, "StartKubeProxy", "kube-proxy"),
		generateStaticPod(commandSpec, commandxSpec),
		generateSystemdService(commandSpec),
		generateWorkerNode(commandSpec, commandxSpec),
		generateWorkerPreRequisites(commandSpec),
		generateArchiveInstall(commandSpec, commandxSpec, "CniPluginsInstall", "Installs cni-plugins on a remote system",
			"bandwidth", "bridge", "dhcp", "dummy", "firewall", "hostDevice",
			"hostLocal", "ipvlan", "loopback", "macvlan", "portmap", "ptp", "sbr",
			"static", "tap", "tuning", "vlan", "vrf"),
		generateArchiveInstall(commandSpec, commandxSpec, "ContainerdInstall", "Installs containerd on a remote system", "containerd"),
		generateArchiveInstall(commandSpec, commandxSpec, "CrictlInstall", "Installs crictl on a remote system", "crictl"),
		generateArchiveInstall(commandSpec, commandxSpec, "EtcdInstall", "Installs etcd on a remote system", "etcd", "etcdctl"),
		generateBinaryInstall(commandSpec, "KubeadmInstall", "Installs kubeadm on a remote system."),
		generateBinaryInstall(commandSpec, commandxSpec, "KubeControllerManagerInstall", "Installs kube-controller-manager on a remote system."),
		generateBinaryInstall(commandSpec, commandxSpec, "KubectlInstall", "Installs kubectl on a remote system."),
		generateBinaryInstall(commandSpec, commandxSpec, "KubeletInstall", "Installs kubelet on a remote system."),
		generateBinaryInstall(commandSpec, commandxSpec, "KubeProxyInstall", "Installs kube-proxy on a remote system."),
		generateBinaryInstall(commandSpec, commandxSpec, "KubeSchedulerInstall", "Installs kube-scheduler on a remote system."),
		generateBinaryInstall(commandSpec, commandxSpec, "RuncInstall", "Installs runc on a remote system."),
	)
}
