//go:build worker || all

package examples

import (
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/rt"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

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
