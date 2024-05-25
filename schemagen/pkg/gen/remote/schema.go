package remote

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/internal"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func Generate(commandSpec schema.PackageSpec) schema.PackageSpec {
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
		generateDownload(commandSpec),
		generateContainerdService(commandSpec),
		generateControlPlaneNode(commandSpec),
		generateEtcdCluster(commandSpec),
		generateEtcdConfiguration(commandSpec),
		generateEtcdService(commandSpec),
		generateFile(commandSpec),
		generateKubeletService(commandSpec),
		generateKubeProxyService(commandSpec),
		generateProvisionEtcd(commandSpec),
		generateStartSystemdService(commandSpec, "StartContainerd", "containerd"),
		generateStartSystemdService(commandSpec, "StartEtcd", "etcd"),
		generateStartSystemdService(commandSpec, "StartKubelet", "kubelet"),
		generateStartSystemdService(commandSpec, "StartKubeProxy", "kube-proxy"),
		generateStaticPod(commandSpec),
		generateSystemdService(commandSpec),
		generateWorkerNode(commandSpec),
		generateArchiveInstall(commandSpec, "CniPluginsInstall", "Installs cni-plugins on a remote system",
			"bandwidth", "bridge", "dhcp", "dummy", "firewall", "hostDevice",
			"hostLocal", "ipvlan", "loopback", "macvlan", "portmap", "ptp", "sbr",
			"static", "tap", "tuning", "vlan", "vrf"),
		generateArchiveInstall(commandSpec, "ContainerdInstall", "Installs containerd on a remote system", "containerd"),
		generateArchiveInstall(commandSpec, "CrictlInstall", "Installs crictl on a remote system", "crictl"),
		generateArchiveInstall(commandSpec, "EtcdInstall", "Installs etcd on a remote system", "etcd", "etcdctl"),
		generateBinaryInstall(commandSpec, "KubeApiServerInstall", "Installs kube-apiserver on a remote system."),
		generateBinaryInstall(commandSpec, "KubeControllerManagerInstall", "Installs kube-controller-manager on a remote system."),
		generateBinaryInstall(commandSpec, "KubectlInstall", "Installs kubectl on a remote system."),
		generateBinaryInstall(commandSpec, "KubeletInstall", "Installs kubelet on a remote system."),
		generateBinaryInstall(commandSpec, "KubeProxyInstall", "Installs kube-proxy on a remote system."),
		generateBinaryInstall(commandSpec, "KubeSchedulerInstall", "Installs kube-scheduler on a remote system."),
		generateBinaryInstall(commandSpec, "RuncInstall", "Installs runc on a remote system."),
	)
}
