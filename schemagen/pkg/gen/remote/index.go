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
		name("ContainerdInstall"):            generateArchiveInstall(commandSpec, "Installs containerd on a remote system", "containerd"),
		name("CrictlInstall"):                generateArchiveInstall(commandSpec, "Installs crictl on a remote system", "crictl"),
		name("Download"):                     generateDownload(commandSpec),
		name("EtcdConfiguration"):            generateEtcdConfiguration(commandSpec),
		name("EtcdInstall"):                  generateArchiveInstall(commandSpec, "Installs etcd on a remote system", "etcd", "etcdctl"),
		name("File"):                         generateFile(commandSpec),
		name("KubeApiServerInstall"):         generateBinaryInstall(commandSpec, "Installs kube-apiserver on a remote system."),
		name("KubeControllerManagerInstall"): generateBinaryInstall(commandSpec, "Installs kube-controller-manager on a remote system."),
		name("KubectlInstall"):               generateBinaryInstall(commandSpec, "Installs kubectl on a remote system."),
		name("KubeletInstall"):               generateBinaryInstall(commandSpec, "Installs kubelet on a remote system."),
		name("KubeProxyInstall"):             generateBinaryInstall(commandSpec, "Installs kube-proxy on a remote system."),
		name("KubeSchedulerInstall"):         generateBinaryInstall(commandSpec, "Installs kube-scheduler on a remote system."),
		name("RuncInstall"):                  generateBinaryInstall(commandSpec, "Installs runc on a remote system."),
		name("StaticPod"):                    generateStaticPod(commandSpec),
		name("SystemdService"):               generateSystemdService(commandSpec),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     generateTypes(),
	}
}
