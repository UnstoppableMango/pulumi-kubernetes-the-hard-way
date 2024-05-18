//go:build nodejs || all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestCniPluginsTs(t *testing.T) {
	ResourceTest(t, "remote/cni-plugins-ts", getJSBaseOptions(t), func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:remote:CniBridgePluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "bridge", "cni0")
			expectOutput(t, res, "name", "bridge")
			expectOutput(t, res, "type", "bridge")
			expectOutput(t, res, "cniVersion", "1.0.0")
			expectOutput(t, res, "path", "/etc/cni/net.d/10-bridge.conf")
			expectOutput(t, res, "subnet", "10.0.69.0/24")
			expectOutput(t, res, "isGateway", true)
			expectOutput(t, res, "ipMasq", true)

			assert.Contains(t, res.Outputs, "file")
			assert.Contains(t, res.Outputs, "ipam")
		})
		Validate(ctx, "kubernetes-the-hard-way:remote:CniLoopbackPluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "name", "lo")
			expectOutput(t, res, "type", "loopback")
			expectOutput(t, res, "cniVersion", "1.1.0")
			expectOutput(t, res, "path", "/etc/cni/net.d/99-loopback.conf")

			assert.Contains(t, res.Outputs, "file")
		})
		Validate(ctx, "kubernetes-the-hard-way:remote:CniPluginConfiguration", "all", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			expectOutput(t, res, "subnet", "10.0.69.0/24")
			expectOutput(t, res, "directory", "/etc/cni2/net.d")

			assert.Contains(t, res.Outputs, "bridge")
			assert.Contains(t, res.Outputs, "connection")
			assert.Contains(t, res.Outputs, "loopback")
			assert.Contains(t, res.Outputs, "mkdir")
		})
	})
}

func TestRemoteEtcdClusterMultiTs(t *testing.T) {
	skipIfShort(t)

	const (
		username = "root"
		password = "root"
	)

	opts := []SshServerOption{
		WithSshUsername(username),
		WithSshPassword(password),
	}
	nodes := newCluster(t, map[string][]SshServerOption{
		"node1": opts,
		"node2": opts,
		"node3": opts,
	})

	validateNode := func(t *testing.T, name string, outputs map[string]interface{}) {
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

	validateSimple := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)
		validateNode(t, "node1", res.Outputs)
		validateNode(t, "node2", res.Outputs)
		validateNode(t, "node3", res.Outputs)
		assert.Contains(t, res.Outputs, "bundle")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "remote", "etcd-cluster-multi-ts"),
			Config: map[string]string{
				"node1-host":     "localhost",
				"node1-ip":       nodes["node1"].Ip,
				"node1-port":     nodes["node1"].Port,
				"node1-user":     username,
				"node1-password": password,
				"node2-host":     "localhost",
				"node2-ip":       nodes["node2"].Ip,
				"node2-port":     nodes["node2"].Port,
				"node2-user":     username,
				"node2-password": password,
				"node3-host":     "localhost",
				"node3-ip":       nodes["node3"].Ip,
				"node3-port":     nodes["node3"].Port,
				"node3-user":     username,
				"node3-password": password,
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					if res.Type != "kubernetes-the-hard-way:remote:EtcdCluster" {
						continue
					}
					switch res.URN.Name() {
					case "simple":
						validateSimple(t, res)
						validatedResources = append(validatedResources, "simple")
					}
				}

				assert.Equal(t, []string{"simple"}, validatedResources, "Not all resources were validated")
			},
		})

	integration.ProgramTest(t, &test)
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

func TestRemoteWorkerTs(t *testing.T) {
	ResourceTest(t, "remote/worker-ts", getJSBaseOptions(t), func(ctx *ResourceContext) {
		Validate(ctx, "kubernetes-the-hard-way:remote:ContainerdConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			assert.Contains(t, res.Outputs, "connection")
			assert.Contains(t, res.Outputs, "cri")
			assert.Contains(t, res.Outputs, "file")

			expectOutput(t, res, "path", "/etc/containerd/containerd-config.toml")
		})
	})
}
