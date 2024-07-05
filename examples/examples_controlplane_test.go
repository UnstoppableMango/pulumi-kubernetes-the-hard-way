//go:build controlplane || all

package examples

import (
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/rt"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

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
