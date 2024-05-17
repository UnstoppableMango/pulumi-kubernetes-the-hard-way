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
	skipIfShort(t)

	const (
		username = "root"
		password = "root"
	)

	node := newNode(t,
		WithSshUsername(username),
		WithSshPassword(password),
	)

	validateBridge := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)

		bridge, ok := res.Outputs["bridge"]
		assert.True(t, ok, "Output `bridge` was not set")
		assert.Equal(t, "cni0", bridge)

		name, ok := res.Outputs["name"]
		assert.True(t, ok, "Output `name` was not set")
		assert.Equal(t, "bridge", name)

		typ, ok := res.Outputs["type"]
		assert.True(t, ok, "Output `type` was not set")
		assert.Equal(t, "bridge", typ)

		cniVersion, ok := res.Outputs["cniVersion"]
		assert.True(t, ok, "Output `cniVersion` was not set")
		assert.Equal(t, "1.0.0", cniVersion)

		path, ok := res.Outputs["path"]
		assert.True(t, ok, "Output `path` was not set")
		assert.Equal(t, "/var/lib/kubernetes", path)

		subnet, ok := res.Outputs["subnet"]
		assert.True(t, ok, "Output `subnet` was not set")
		assert.Equal(t, "10.0.69.0/24", subnet)

		isGateway, ok := res.Outputs["isGateway"]
		assert.True(t, ok, "Output `isGateway` was not set")
		assert.Equal(t, true, isGateway)

		ipMasq, ok := res.Outputs["ipMasq"]
		assert.True(t, ok, "Output `ipMasq` was not set")
		assert.Equal(t, true, ipMasq)

		assert.Contains(t, res.Outputs, "file")
		assert.Contains(t, res.Outputs, "ipam")
	}

	validateLoopback := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)

		cniVersion, ok := res.Outputs["cniVersion"]
		assert.True(t, ok, "Output `cniVersion` was not set")
		assert.Equal(t, "1.1.0", cniVersion)

		name, ok := res.Outputs["name"]
		assert.True(t, ok, "Output `name` was not set")
		assert.Equal(t, "lo", name)

		path, ok := res.Outputs["path"]
		assert.True(t, ok, "Output `path` was not set")
		assert.Equal(t, "/var/lib/kubernetes", path)

		typ, ok := res.Outputs["type"]
		assert.True(t, ok, "Output `type` was not set")
		assert.Equal(t, "loopback", typ)

		assert.Contains(t, res.Outputs, "file")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "remote", "cni-plugins-ts"),
			Config: map[string]string{
				"host":     "localhost",
				"port":     node.Port,
				"user":     username,
				"password": password,
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					switch res.Type {
					case "kubernetes-the-hard-way:remote:CniBridgePluginConfiguration":
						switch res.URN.Name() {
						case "simple":
							validateBridge(t, res)
							validatedResources = append(validatedResources, "simple")
						}
					case "kubernetes-the-hard-way:remote:CniLoopbackPluginConfiguration":
						switch res.URN.Name() {
						case "simple":
							validateLoopback(t, res)
							validatedResources = append(validatedResources, "simple")
						}
					}
				}

				assert.Equal(t, []string{
					"simple",
					"simple",
				}, validatedResources, "Not all resources were validated")
			},
		})

	integration.ProgramTest(t, &test)
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
	skipIfShort(t)

	const (
		username = "root"
		password = "root"
	)

	node := newNode(t, WithSshUsername(username), WithSshPassword(password))

	validateSimple := func(t *testing.T, res apitype.ResourceV3) {
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
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "remote", "etcd-cluster-single-ts"),
			Config: map[string]string{
				"host":     "localhost",
				"port":     node.Port,
				"user":     username,
				"password": password,
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

func TestRemoteEtcdInstallTs(t *testing.T) {
	skipIfShort(t)

	const (
		username = "root"
		password = "root"
	)

	node := newNode(t,
		WithSshUsername(username),
		WithSshPassword(password),
	)

	validateSimple := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)

		arch, ok := res.Outputs["architecture"]
		assert.True(t, ok, "Output `architecture` was not set")
		assert.Equal(t, "amd64", arch)

		archiveName, ok := res.Outputs["archiveName"]
		assert.True(t, ok, "Output `archiveName` was not set")
		assert.Equal(t, "etcd-v3.4.15-linux-amd64.tar.gz", archiveName)

		directory, ok := res.Outputs["directory"]
		assert.True(t, ok, "Output `directory` was not set")
		assert.Equal(t, "/usr/local/bin", directory)

		etcdPath, ok := res.Outputs["etcdPath"]
		assert.True(t, ok, "Output `etcdPath` was not set")
		assert.Equal(t, "/usr/local/bin/etcd", etcdPath)

		etcdctlPath, ok := res.Outputs["etcdctlPath"]
		assert.True(t, ok, "Output `etcdctlPath` was not set")
		assert.Equal(t, "/usr/local/bin/etcdctl", etcdctlPath)

		name, ok := res.Outputs["name"]
		assert.True(t, ok, "Output `name` was not set")
		assert.Equal(t, "simple", name)

		url, ok := res.Outputs["url"]
		assert.True(t, ok, "Output `url` was not set")
		assert.Equal(t, "https://github.com/etcd-io/etcd/releases/download/v3.4.15/etcd-v3.4.15-linux-amd64.tar.gz", url)

		version, ok := res.Outputs["version"]
		assert.True(t, ok, "Output `version` was not set")
		assert.Equal(t, "3.4.15", version)

		assert.Contains(t, res.Outputs, "download")
		assert.Contains(t, res.Outputs, "mkdir")
		assert.Contains(t, res.Outputs, "mvEtcd")
		assert.Contains(t, res.Outputs, "mvEtcdctl")
		assert.Contains(t, res.Outputs, "tar")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "remote", "etcd-install-ts"),
			SkipPreview:   true,  // TODO: There is some bug surrounding this
			RunUpdateTest: false, // TODO: Enable
			Config: map[string]string{
				"host":     "localhost",
				"port":     node.Port,
				"user":     username,
				"password": password,
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				etcdctlVersion, ok := stack.Outputs["etcdctlVersion"]
				assert.True(t, ok, "Output `etcdctlVersion` was not set")
				assert.Equal(t, "etcdctl version: 3.4.15\nAPI version: 3.4", etcdctlVersion) // TODO: No hardcoded versions

				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					if res.Type != "kubernetes-the-hard-way:remote:EtcdInstall" {
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
