package remote

import (
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
		name("Download"):                       generateDownload(commandSpec),
		name("EtcdCluster"):                    generateEtcdCluster(),
		name("EtcdConfiguration"):              generateEtcdConfiguration(commandSpec),
		name("EtcdInstall"):                    generateArchiveInstall(commandSpec, "Installs etcd on a remote system", "etcd", "etcdctl"),
		name("EtcdService"):                    generateEtcdService(commandSpec),
		name("File"):                           generateFile(commandSpec),
		name("KubeApiServerInstall"):           generateBinaryInstall(commandSpec, "Installs kube-apiserver on a remote system."),
		name("KubeControllerManagerInstall"):   generateBinaryInstall(commandSpec, "Installs kube-controller-manager on a remote system."),
		name("KubectlInstall"):                 generateBinaryInstall(commandSpec, "Installs kubectl on a remote system."),
		name("KubeletInstall"):                 generateBinaryInstall(commandSpec, "Installs kubelet on a remote system."),
		name("KubeletService"):                 generateKubeletService(commandSpec),
		name("KubeProxyInstall"):               generateBinaryInstall(commandSpec, "Installs kube-proxy on a remote system."),
		name("KubeProxyService"):               generateKubeProxyService(commandSpec),
		name("KubeSchedulerInstall"):           generateBinaryInstall(commandSpec, "Installs kube-scheduler on a remote system."),
		name("RuncInstall"):                    generateBinaryInstall(commandSpec, "Installs runc on a remote system."),
		name("ProvisionEtcd"):                  generateProvisionEtcd(commandSpec),
		name("StartContainerd"):                generateStartSystemdService(commandSpec, "containerd"),
		name("StartEtcd"):                      generateStartSystemdService(commandSpec, "etcd"),
		name("StartKubelet"):                   generateStartSystemdService(commandSpec, "kubelet"),
		name("StartKubeProxy"):                 generateStartSystemdService(commandSpec, "kube-proxy"),
		name("StaticPod"):                      generateStaticPod(commandSpec),
		name("SystemdService"):                 generateSystemdService(commandSpec),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     generateTypes(commandSpec),
	}
}
