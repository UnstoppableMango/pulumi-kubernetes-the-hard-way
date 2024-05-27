//go:build nodejs || all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestCniPluginsTs(t *testing.T) {
	ResourceTest(t, "remote/cni-plugins-ts", getJSBaseOptions(t), func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:config:CniBridgePluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "result", map[string]interface{}{
				"name":       "bridge",
				"type":       "bridge",
				"bridge":     "cni0",
				"cniVersion": "1.0.0",
				"isGateway":  true,
				"ipMasq":     true,
				"subnet":     "10.0.69.0/24",
				"ipam": map[string]interface{}{
					"type": "host-local",
					"ranges": []interface{}{
						map[string]interface{}{"subnet": "10.0.69.0/24"},
					},
					"routes": []interface{}{
						map[string]interface{}{"dst": "0.0.0.0/0"},
					},
				},
			})

			assert.Contains(t, res.Outputs, "yaml")
		})
		Validate(ctx, "kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "result", map[string]interface{}{
				"name":       "lo",
				"type":       "loopback",
				"cniVersion": "1.1.0",
			})

			assert.Contains(t, res.Outputs, "yaml")
		})
	})
}

func TestRemoteEtcdClusterMultiTs(t *testing.T) {
	opts := getJSBaseOptions(t).With(MultiContainerSetup(t))
	ResourceTest(t, "remote/etcd-cluster-multi-ts", opts, func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:remote:EtcdCluster", "simple", func(t *testing.T, res apitype.ResourceV3) {
			validateNode := func(name string, outputs map[string]interface{}) {
				install, ok := outputs["install"]
				assert.True(t, ok, "Output `install` was not set")
				assert.Contains(t, install, name)

				configuration, ok := outputs["configuration"]
				assert.True(t, ok, "Output `configuration` was not set")
				assert.Contains(t, configuration, name)

				nodes, ok := outputs["nodes"]
				assert.True(t, ok, "Output `nodes` was not set")
				assert.Contains(t, nodes, name)

				service, ok := outputs["service"]
				assert.True(t, ok, "Output `service` was not set")
				assert.Contains(t, service, name)

				start, ok := outputs["start"]
				assert.True(t, ok, "Output `start` was not set")
				assert.Contains(t, start, name)
			}

			assert.NotEmpty(t, res.Outputs)

			validateNode("node1", res.Outputs)
			validateNode("node2", res.Outputs)
			validateNode("node3", res.Outputs)

			assert.Contains(t, res.Outputs, "bundle")
		})
	})
}

func TestRemoteEtcdClusterSingleTs(t *testing.T) {
	ResourceTest(t, "remote/etcd-cluster-single-ts", getJSBaseOptions(t), func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:remote:EtcdCluster", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			install, ok := res.Outputs["install"]
			assert.True(t, ok, "Output `install` was not set")
			assert.Contains(t, install, "node0")

			configuration, ok := res.Outputs["configuration"]
			assert.True(t, ok, "Output `configuration` was not set")
			assert.Contains(t, configuration, "node0")

			nodes, ok := res.Outputs["nodes"]
			assert.True(t, ok, "Output `nodes` was not set")
			assert.Contains(t, nodes, "node0")

			service, ok := res.Outputs["service"]
			assert.True(t, ok, "Output `service` was not set")
			assert.Contains(t, service, "node0")

			start, ok := res.Outputs["start"]
			assert.True(t, ok, "Output `start` was not set")
			assert.Contains(t, start, "node0")

			assert.Contains(t, res.Outputs, "bundle")
		})
	})
}

func TestRemoteEtcdInstallTs(t *testing.T) {
	options := getJSBaseOptions(t).With(integration.ProgramTestOptions{
		SkipPreview:   true,
		RunUpdateTest: false,
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			etcdctlVersion, ok := stack.Outputs["etcdctlVersion"]
			assert.True(t, ok, "Output `etcdctlVersion` was not set")
			assert.Equal(t, "etcdctl version: 3.4.15\nAPI version: 3.4", etcdctlVersion) // TODO: No hardcoded versions
		},
	})

	ResourceTest(t, "remote/etcd-install-ts", options, func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:remote:EtcdInstall", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "architecture", "amd64")
			expectOutput(t, res, "archiveName", "etcd-v3.4.15-linux-amd64.tar.gz")
			expectOutput(t, res, "directory", "/usr/local/bin")
			expectOutput(t, res, "etcdPath", "/usr/local/bin/etcd")
			expectOutput(t, res, "etcdctlPath", "/usr/local/bin/etcdctl")
			expectOutput(t, res, "name", "simple")
			expectOutput(t, res, "url", "https://github.com/etcd-io/etcd/releases/download/v3.4.15/etcd-v3.4.15-linux-amd64.tar.gz")
			expectOutput(t, res, "version", "3.4.15")

			assert.Contains(t, res.Outputs, "download")
			assert.Contains(t, res.Outputs, "mkdir")
			assert.Contains(t, res.Outputs, "mvEtcd")
			assert.Contains(t, res.Outputs, "mvEtcdctl")
			assert.Contains(t, res.Outputs, "tar")
		})
	})
}

func TestRemoteControlPlaneTs(t *testing.T) {
	ResourceTest(t, "remote/control-plane-ts", getJSBaseOptions(t), func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:remote:ControlPlaneNode", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "apiServerCount", 1.)
			expectOutput(t, res, "architecture", "amd64")
			expectOutput(t, res, "auditLogPath", "/var/log/audit.log")
			expectOutput(t, res, "caCertificatePath", "TODO")
			expectOutput(t, res, "caPrivateKeyPath", "TODO")
			expectOutput(t, res, "clusterCIDR", "10.200.0.0/16")
			expectOutput(t, res, "clusterName", "kubernetes")
			expectOutput(t, res, "encryptionConfigYaml", "TODO")
			expectOutput(t, res, "kubeApiServerCertificatePath", "TODO")
			expectOutput(t, res, "kubeApiServerInstallDirectory", "/usr/local/bin")
			expectOutput(t, res, "kubeApiServerPrivateKeyPath", "TODO")
			expectOutput(t, res, "kubeControllerManagerInstallDirectory", "/usr/local/bin")
			expectOutput(t, res, "kubeControllerManagerKubeconfigPath", "TODO")
			expectOutput(t, res, "kubectlInstallDirectory", "/usr/local/bin")
			expectOutput(t, res, "kubernetesVersion", "1.30.0")
			expectOutput(t, res, "kubeSchedulerConfigYaml", "TODO")
			expectOutput(t, res, "kubeSchedulerInstallDirectory", "/usr/local/bin")
			expectOutput(t, res, "kubeSchedulerKubeconfigPath", "TODO")
			expectOutput(t, res, "nodeName", "server")
			expectOutput(t, res, "serviceAccountsCertificatePath", "TODO")
			expectOutput(t, res, "serviceAccountsPrivateKeyPath", "TODO")
			expectOutput(t, res, "serviceClusterIpRange", "10.32.0.0/24")

			assert.Contains(t, res.Outputs, "kubeApiServerInstall")
			assert.Contains(t, res.Outputs, "kubeApiServerService")
			assert.Contains(t, res.Outputs, "kubeControllerManagerInstall")
			assert.Contains(t, res.Outputs, "kubeControllerManagerService")
			assert.Contains(t, res.Outputs, "kubectlInstall")
			assert.Contains(t, res.Outputs, "kubernetesConfigurationMkdir")
			assert.Contains(t, res.Outputs, "kubeSchedulerInstall")
			assert.Contains(t, res.Outputs, "varLibKubernetesMkdir")
		})
	})
}

func TestRemoteWorkerTs(t *testing.T) {
	ResourceTest(t, "remote/worker-ts", getJSBaseOptions(t), func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:config:ContainerdConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			assert.Contains(t, res.Outputs, "result")
			assert.Contains(t, res.Outputs, "toml")
		})
		Validate(ctx, "kubernetes-the-hard-way:remote:WorkerNode", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "subnet", "0.0.0.0/24")
			expectOutput(t, res, "clusterCIDR", "10.200.0.0/16")
			expectOutput(t, res, "architecture", "amd64")
			expectOutput(t, res, "caPath", "TODO")
			expectOutput(t, res, "kubeletCertificatePath", "TODO")
			expectOutput(t, res, "kubeletPrivateKeyPath", "TODO")
			expectOutput(t, res, "clusterCIDR", "10.200.0.0/16")
			expectOutput(t, res, "cniConfigurationDirectory", "/etc/cni/net.d")
			expectOutput(t, res, "containerdConfigurationDirectory", "/etc/containerd")
			expectOutput(t, res, "kubeletConfigurationDirectory", "/var/lib/kubelet")
			expectOutput(t, res, "kubeletKubeconfigPath", "/var/lib/kubelet/kubeconfig")
			expectOutput(t, res, "kubeProxyConfigurationDirectory", "/var/lib/kube-proxy")
			expectOutput(t, res, "kubeProxyKubeconfigPath", "/var/lib/kube-proxy/kubeconfig")
			expectOutput(t, res, "kubernetesVersion", "1.30.0")

			assert.Contains(t, res.Outputs, "cniMkdir")
			assert.Contains(t, res.Outputs, "containerdMkdir")
			assert.Contains(t, res.Outputs, "kubeletMkdir")
			assert.Contains(t, res.Outputs, "kubeProxyMkdir")
			assert.Contains(t, res.Outputs, "varLibKubernetesMkdir")
			assert.Contains(t, res.Outputs, "varRunKubernetesMkdir")
			assert.Contains(t, res.Outputs, "cniPluginsInstall")
			assert.Contains(t, res.Outputs, "cniBridgeConfiguration")
			assert.Contains(t, res.Outputs, "cniBridgeConfigurationFile")
			assert.Contains(t, res.Outputs, "cniLoopbackConfiguration")
			assert.Contains(t, res.Outputs, "cniLoopbackConfigurationFile")
			assert.Contains(t, res.Outputs, "containerdConfiguration")
			assert.Contains(t, res.Outputs, "containerdConfigurationFile")
			assert.Contains(t, res.Outputs, "containerdInstall")
			assert.Contains(t, res.Outputs, "containerdService")
			assert.Contains(t, res.Outputs, "crictlInstall")
			assert.Contains(t, res.Outputs, "kubectlInstall")
			assert.Contains(t, res.Outputs, "kubeletInstall")
			assert.Contains(t, res.Outputs, "kubeletConfiguration")
			assert.Contains(t, res.Outputs, "kubeletConfigurationFile")
			assert.Contains(t, res.Outputs, "kubeletService")
			assert.Contains(t, res.Outputs, "kubeProxyConfiguration")
			assert.Contains(t, res.Outputs, "kubeProxyConfigurationFile")
			assert.Contains(t, res.Outputs, "kubeProxyInstall")
			assert.Contains(t, res.Outputs, "kubeProxyService")
		})
	})
}
