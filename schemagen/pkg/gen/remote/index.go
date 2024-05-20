package remote

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/internal"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func Generate(commandSpec schema.PackageSpec) schema.PackageSpec {
	functions := map[string]schema.FunctionSpec{}

	resources := map[string]schema.ResourceSpec{
		name("CniPluginsInstall"): generateArchiveInstall(commandSpec, "Installs cni-plugins on a remote system",
			"bandwidth", "bridge", "dhcp", "dummy", "firewall", "hostDevice",
			"hostLocal", "ipvlan", "loopback", "macvlan", "portmap", "ptp", "sbr",
			"static", "tap", "tuning", "vlan", "vrf"),
		name("ContainerdConfiguration"):        generateContainerdConfiguration(commandSpec),
		name("ContainerdInstall"):              generateArchiveInstall(commandSpec, "Installs containerd on a remote system", "containerd"),
		name("CniBridgePluginConfiguration"):   generateCniBridgePluginConfiguration(commandSpec),
		name("CniLoopbackPluginConfiguration"): generateCniLoopbackPluginConfiguration(commandSpec),
		name("CniPluginConfiguration"):         generateCniPluginConfiguration(commandSpec),
		name("CrictlInstall"):                  generateArchiveInstall(commandSpec, "Installs crictl on a remote system", "crictl"),
		name("EtcdInstall"):                    generateArchiveInstall(commandSpec, "Installs etcd on a remote system", "etcd", "etcdctl"),
		name("KubeApiServerInstall"):           generateBinaryInstall(commandSpec, "Installs kube-apiserver on a remote system."),
		name("KubeControllerManagerInstall"):   generateBinaryInstall(commandSpec, "Installs kube-controller-manager on a remote system."),
		name("KubectlInstall"):                 generateBinaryInstall(commandSpec, "Installs kubectl on a remote system."),
		name("KubeletInstall"):                 generateBinaryInstall(commandSpec, "Installs kubelet on a remote system."),
		name("KubeProxyInstall"):               generateBinaryInstall(commandSpec, "Installs kube-proxy on a remote system."),
		name("KubeSchedulerInstall"):           generateBinaryInstall(commandSpec, "Installs kube-scheduler on a remote system."),
		name("RuncInstall"):                    generateBinaryInstall(commandSpec, "Installs runc on a remote system."),
		name("StaticPod"):                      generateStaticPod(commandSpec),
		name("SystemdService"):                 generateSystemdService(commandSpec),
	}

	base := schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     generateTypes(),
	}

	return internal.ExtendSchemas(base,
		generateDownload(commandSpec),
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
	)
}
