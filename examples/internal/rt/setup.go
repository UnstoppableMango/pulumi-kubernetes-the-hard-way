package rt

import (
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/node"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

type ResourceTestSetup struct {
	Config map[string]string
	Setup  func(t *testing.T)
}

func SingleContainerSetup(t *testing.T) integration.ProgramTestOptions {
	const (
		username = "root"
		password = "root"
	)

	node := newNode(t,
		node.WithSshUsername(username),
		node.WithSshPassword(password),
	)

	return integration.ProgramTestOptions{
		Config: map[string]string{
			"test:container": "single",
			"host":           "localhost",
			"port":           node.SshPort,
			"user":           username,
			"password":       password,
		},
	}
}

func MultiContainerSetup(t *testing.T) integration.ProgramTestOptions {
	const (
		username = "root"
		password = "root"
	)

	opts := []node.SshServerOption{
		node.WithSshUsername(username),
		node.WithSshPassword(password),
	}
	nodes := newCluster(t, map[string][]node.SshServerOption{
		"node1": opts,
		"node2": opts,
		"node3": opts,
	})

	return integration.ProgramTestOptions{
		Config: map[string]string{
			"test:container": "multi",
			"node1-host":     "localhost",
			"node1-ip":       nodes["node1"].Ip,
			"node1-port":     nodes["node1"].SshPort,
			"node1-user":     username,
			"node1-password": password,
			"node2-host":     "localhost",
			"node2-ip":       nodes["node2"].Ip,
			"node2-port":     nodes["node2"].SshPort,
			"node2-user":     username,
			"node2-password": password,
			"node3-host":     "localhost",
			"node3-ip":       nodes["node3"].Ip,
			"node3-port":     nodes["node3"].SshPort,
			"node3-user":     username,
			"node3-password": password,
		},
	}
}
