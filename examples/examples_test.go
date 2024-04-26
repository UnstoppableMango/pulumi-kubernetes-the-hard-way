package examples

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/network"
	"golang.org/x/exp/maps"
)

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Bin: path.Join(getCwd(t), "..", ".pulumi", "bin", "pulumi"),
		LocalProviders: []integration.LocalDependency{{
			Package: "kubernetes-the-hard-way",
			Path:    path.Join(getCwd(t), "..", "bin"),
		}},
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}

type node struct {
	Server SshServer
	Port   string
	Ip     string
}

func (n node) Exec(t *testing.T, command []string) string {
	result, err := n.Server.Exec(context.Background(), command)
	require.NoError(t, err, "failed to execute command")
	return result
}

func (n node) ReadFile(t *testing.T, file string) string {
	result, err := n.Server.ReadFile(context.Background(), file)
	require.NoError(t, err, "failed to read file")
	return result
}

func newNode(t *testing.T, opts ...SshServerOption) node {
	ctx := context.Background()
	server, err := StartSshServer(ctx, opts...)
	require.NoError(t, err, "failed to start ssh server")

	t.Cleanup(func() {
		err := StopSshServer(ctx, server)
		require.NoError(t, err)
	})

	port, err := server.Port(ctx)
	require.NoError(t, err, "failed to get node port")

	ip, err := server.Ip(ctx)
	require.NoError(t, err, "failed to retrieve node IP")

	return node{Server: server, Port: port, Ip: ip}
}

func newCluster(t *testing.T, config map[string][]SshServerOption) map[string]node {
	ctx := context.Background()
	network, err := network.New(ctx, network.WithCheckDuplicate())
	require.NoError(t, err, "failed to create network")

	nodes := map[string]node{}
	for _, key := range maps.Keys(config) {
		opts := append(config[key], WithNetwork(network.Name))
		nodes[key] = newNode(t, opts...)
	}

	return nodes
}
