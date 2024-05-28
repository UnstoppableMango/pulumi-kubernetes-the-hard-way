package rt

import (
	"fmt"
	"testing"
	"time"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/node"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdContext struct {
	ResourceContext
	EtcdClient clientv3.Client
}

func EtcdNodeTest(t *testing.T, project string, options integration.ProgramTestOptions, validation func(*EtcdContext)) {
	const (
		username = "root"
		password = "root"
	)

	node := newNode(t,
		node.WithSshUsername(username),
		node.WithSshPassword(password),
	)

	options = options.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"test:container": "single",
			"host":           "localhost",
			"port":           node.SshPort,
			"user":           username,
			"password":       password,
		},
	})

	ResourceTest(t, project, options, func(ctx *ResourceContext) {
		etcdClient, err := clientv3.New(clientv3.Config{
			Endpoints: []string{
				fmt.Sprintf("localhost:%s", node.EtcdPort),
			},
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			t.Fatal("failed to create etcd client")
		}

		validation(&EtcdContext{
			ResourceContext: *ctx,
			EtcdClient:      *etcdClient,
		})
	})
}

func EtcdClusterTest(t *testing.T, project string, options integration.ProgramTestOptions, validation func(*EtcdContext)) {
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

	node1 := nodes["node1"]
	node2 := nodes["node2"]
	node3 := nodes["node3"]

	options = options.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"test:container": "multi",
			"node1-host":     "localhost",
			"node1-ip":       node1.Ip,
			"node1-port":     node1.SshPort,
			"node1-user":     username,
			"node1-password": password,
			"node2-host":     "localhost",
			"node2-ip":       node2.Ip,
			"node2-port":     node2.SshPort,
			"node2-user":     username,
			"node2-password": password,
			"node3-host":     "localhost",
			"node3-ip":       node3.Ip,
			"node3-port":     node3.SshPort,
			"node3-user":     username,
			"node3-password": password,
		},
	})

	ResourceTest(t, project, options, func(ctx *ResourceContext) {
		etcdClient, err := clientv3.New(clientv3.Config{
			Endpoints: []string{
				fmt.Sprintf("localhost:%s", node1.EtcdPort),
				fmt.Sprintf("localhost:%s", node2.EtcdPort),
				fmt.Sprintf("localhost:%s", node3.EtcdPort),
			},
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			t.Fatal("failed to create etcd client")
		}

		validation(&EtcdContext{
			ResourceContext: *ctx,
			EtcdClient:      *etcdClient,
		})
	})
}
