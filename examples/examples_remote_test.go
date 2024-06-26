//go:build nodejs || all

package examples

import (
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/rt"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestCniPluginsTs(t *testing.T) {
	rt.ResourceTest(t, "remote/cni-plugins-ts", getJSBaseOptions(t), func(ctx *rt.ResourceContext) {
		rt.Validate(ctx, "kubernetes-the-hard-way:config:CniBridgePluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			rt.ExpectOutput(t, res, "result", map[string]interface{}{
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
		rt.Validate(ctx, "kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			rt.ExpectOutput(t, res, "result", map[string]interface{}{
				"name":       "lo",
				"type":       "loopback",
				"cniVersion": "1.1.0",
			})

			assert.Contains(t, res.Outputs, "yaml")
		})
	})
}

func TestRemoteEtcdClusterMultiTs(t *testing.T) {
	skipIfCi(t) // This test breaks all the time and we know it works to skip it for now
	opts := getJSBaseOptions(t).With(rt.MultiContainerSetup(t))
	rt.ResourceTest(t, "remote/etcd-cluster-multi-ts", opts, func(ctx *rt.ResourceContext) {
		rt.Validate(ctx, "kubernetes-the-hard-way:remote:EtcdCluster", "simple", func(t *testing.T, res apitype.ResourceV3) {
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
	rt.ResourceTest(t, "remote/etcd-cluster-single-ts", getJSBaseOptions(t), func(ctx *rt.ResourceContext) {
		rt.Validate(ctx, "kubernetes-the-hard-way:remote:EtcdCluster", "simple", func(t *testing.T, res apitype.ResourceV3) {
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

	rt.ResourceTest(t, "remote/etcd-install-ts", options, func(ctx *rt.ResourceContext) {
		rt.Validate(ctx, "kubernetes-the-hard-way:remote:EtcdInstall", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			rt.ExpectOutput(t, res, "architecture", "amd64")
			rt.ExpectOutput(t, res, "archiveName", "etcd-v3.4.15-linux-amd64.tar.gz")
			rt.ExpectOutput(t, res, "directory", "/usr/local/bin")
			rt.ExpectOutput(t, res, "etcdPath", "/usr/local/bin/etcd")
			rt.ExpectOutput(t, res, "etcdctlPath", "/usr/local/bin/etcdctl")
			rt.ExpectOutput(t, res, "name", "simple")
			rt.ExpectOutput(t, res, "url", "https://github.com/etcd-io/etcd/releases/download/v3.4.15/etcd-v3.4.15-linux-amd64.tar.gz")
			rt.ExpectOutput(t, res, "version", "3.4.15")

			assert.Contains(t, res.Outputs, "download")
			assert.Contains(t, res.Outputs, "mkdir")
			assert.Contains(t, res.Outputs, "mvEtcd")
			assert.Contains(t, res.Outputs, "mvEtcdctl")
			assert.Contains(t, res.Outputs, "tar")
		})
	})
}

func TestRemoteControlPlaneTs(t *testing.T) {
	rt.ResourceTest(t, "remote/control-plane-ts", getJSBaseOptions(t), func(ctx *rt.ResourceContext) {
		rt.Validate(ctx, "kubernetes-the-hard-way:remote:ControlPlaneNode", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			rt.ExpectOutput(t, res, "apiServerCount", 1.)
			rt.ExpectOutput(t, res, "architecture", "amd64")
			rt.ExpectOutput(t, res, "auditLogPath", "/var/log/audit.log")
			rt.ExpectOutput(t, res, "caCertificatePath", "TODO")
			rt.ExpectOutput(t, res, "caPrivateKeyPath", "TODO")
			rt.ExpectOutput(t, res, "clusterCIDR", "10.200.0.0/16")
			rt.ExpectOutput(t, res, "clusterName", "kubernetes")
			rt.ExpectOutput(t, res, "encryptionConfigYaml", "TODO")
			rt.ExpectOutput(t, res, "kubeApiServerCertificatePath", "TODO")
			rt.ExpectOutput(t, res, "kubeApiServerInstallDirectory", "/usr/local/bin")
			rt.ExpectOutput(t, res, "kubeApiServerPrivateKeyPath", "TODO")
			rt.ExpectOutput(t, res, "kubeControllerManagerInstallDirectory", "/usr/local/bin")
			rt.ExpectOutput(t, res, "kubeControllerManagerKubeconfigPath", "TODO")
			rt.ExpectOutput(t, res, "kubectlInstallDirectory", "/usr/local/bin")
			rt.ExpectOutput(t, res, "kubernetesVersion", "1.30.0")
			rt.ExpectOutput(t, res, "kubeSchedulerConfigYaml", "TODO")
			rt.ExpectOutput(t, res, "kubeSchedulerInstallDirectory", "/usr/local/bin")
			rt.ExpectOutput(t, res, "kubeSchedulerKubeconfigPath", "TODO")
			rt.ExpectOutput(t, res, "nodeName", "server")
			rt.ExpectOutput(t, res, "serviceAccountsCertificatePath", "TODO")
			rt.ExpectOutput(t, res, "serviceAccountsPrivateKeyPath", "TODO")
			rt.ExpectOutput(t, res, "serviceClusterIpRange", "10.32.0.0/24")

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
	rt.ResourceTest(t, "remote/worker-ts", getJSBaseOptions(t), func(ctx *rt.ResourceContext) {
		rt.Validate(ctx, "kubernetes-the-hard-way:config:ContainerdConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			assert.Contains(t, res.Outputs, "result")
			assert.Contains(t, res.Outputs, "toml")
		})
		rt.Validate(ctx, "kubernetes-the-hard-way:remote:WorkerNode", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			rt.ExpectOutput(t, res, "subnet", "0.0.0.0/24")
			rt.ExpectOutput(t, res, "clusterCIDR", "10.200.0.0/16")
			rt.ExpectOutput(t, res, "architecture", "amd64")
			rt.ExpectOutput(t, res, "caPath", "TODO")
			rt.ExpectOutput(t, res, "kubeletCertificatePath", "TODO")
			rt.ExpectOutput(t, res, "kubeletPrivateKeyPath", "TODO")
			rt.ExpectOutput(t, res, "clusterCIDR", "10.200.0.0/16")
			rt.ExpectOutput(t, res, "cniConfigurationDirectory", "/etc/cni/net.d")
			rt.ExpectOutput(t, res, "containerdConfigurationDirectory", "/etc/containerd")
			rt.ExpectOutput(t, res, "kubeletConfigurationDirectory", "/var/lib/kubelet")
			rt.ExpectOutput(t, res, "kubeletKubeconfigPath", "/var/lib/kubelet/kubeconfig")
			rt.ExpectOutput(t, res, "kubeProxyConfigurationDirectory", "/var/lib/kube-proxy")
			rt.ExpectOutput(t, res, "kubeProxyKubeconfigPath", "/var/lib/kube-proxy/kubeconfig")
			rt.ExpectOutput(t, res, "kubernetesVersion", "1.30.0")

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
