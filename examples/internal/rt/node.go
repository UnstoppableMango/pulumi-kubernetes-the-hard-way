package rt

import (
	"context"
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/node"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/network"
	"golang.org/x/exp/maps"
)

type containerNode struct {
	Server   node.SshServer
	SshPort  string
	EtcdPort string
	Ip       string
}

func newNode(t *testing.T, opts ...node.SshServerOption) containerNode {
	ctx := context.Background()
	server, err := node.StartSshServer(ctx, opts...)
	require.NoError(t, err, "failed to start ssh server")

	t.Cleanup(func() {
		err := node.StopSshServer(ctx, server)
		require.NoError(t, err)
	})

	sshPort, err := server.SshPort(ctx)
	require.NoError(t, err, "failed to get ssh port")

	etcdPort, err := server.EtcdPort(ctx)
	require.NoError(t, err, "failed to get etcd port")

	ip, err := server.Ip(ctx)
	require.NoError(t, err, "failed to retrieve node IP")

	return containerNode{
		Server:   server,
		SshPort:  sshPort,
		EtcdPort: etcdPort,
		Ip:       ip,
	}
}

func newCluster(t *testing.T, config map[string][]node.SshServerOption) map[string]containerNode {
	ctx := context.Background()
	network, err := network.New(ctx, network.WithCheckDuplicate())
	require.NoError(t, err, "failed to create network")

	t.Cleanup(func() {
		require.NoError(t, network.Remove(ctx))
	})

	nodes := map[string]containerNode{}
	for _, key := range maps.Keys(config) {
		opts := append(config[key], node.WithNetwork(network.Name))
		nodes[key] = newNode(t, opts...)
	}

	return nodes
}

func (n containerNode) Exec(t *testing.T, command []string) string {
	result, err := n.Server.Exec(context.Background(), command)
	require.NoError(t, err, "failed to execute command")
	return result
}

func (n containerNode) ReadFile(t *testing.T, file string) string {
	result, err := n.Server.ReadFile(context.Background(), file)
	require.NoError(t, err, "failed to read file")
	return result
}
