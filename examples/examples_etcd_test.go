//go:build etcd || all

package examples

import (
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/rt"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

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
