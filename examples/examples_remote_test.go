//go:build nodejs || all

package examples

import (
	"context"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestRemoteEtcdClusterTs(t *testing.T) {
	skipIfShort(t)

	const (
		username = "root"
		password = "root"
	)

	ctx := context.Background()
	server, err := StartSshServer(ctx,
		WithSshUsername(username),
		WithSshPassword(password))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	port, err := server.Port(ctx)
	assert.NoError(t, err)

	defer StopSshServer(ctx, server) // TODO: Error handling?

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
			Dir: path.Join(getCwd(t), "remote", "etcd-cluster-ts"),
			Config: map[string]string{
				"host":     "localhost",
				"port":     port,
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

	ctx := context.Background()
	server, err := StartSshServer(ctx,
		WithSshUsername(username),
		WithSshPassword(password))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	port, err := server.Port(ctx)
	assert.NoError(t, err)

	defer StopSshServer(ctx, server) // TODO: Error handling?

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
				"port":     port,
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
