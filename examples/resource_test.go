package examples

import (
	"context"
	"fmt"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/network"
	clientV3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/exp/maps"
)

type ResourceValidator func(t *testing.T, res apitype.ResourceV3)

type ResourceTestSetup struct {
	Config map[string]string
	Setup  func(t *testing.T)
}

type ResourceContext struct {
	EtcdClient clientV3.Client
	tokens     map[validatorKey]ResourceValidator
}

type validatorKey struct {
	Type tokens.Type
	Name string
}

type node struct {
	Server   SshServer
	SshPort  string
	EtcdPort string
	Ip       string
}

func Validate(ctx *ResourceContext, typ tokens.Type, name string, validator ResourceValidator) {
	ctx.tokens[validatorKey{Type: typ, Name: name}] = validator
}

func ResourceTest(t *testing.T, project string, baseOptions integration.ProgramTestOptions, validation func(ctx *ResourceContext)) {
	skipIfShort(t)

	if containerType, ok := baseOptions.Config["test:container"]; !ok {
		baseOptions = baseOptions.With(SingleContainerSetup(t))
	}

	test := baseOptions.With(integration.ProgramTestOptions{
		Dir: path.Join(getCwd(t), project),
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			var etcdClient clientV3.Client

			ctx := ResourceContext{tokens: map[validatorKey]ResourceValidator{}}
			validation(&ctx)
			lookup := ctx.tokens

			validated := []string{}
			for _, res := range stack.Deployment.Resources {
				key := validatorKey{Type: res.Type, Name: res.URN.Name()}
				testName := fmt.Sprintf("%s::%s", res.Type, res.URN.Name())
				if validator, ok := lookup[key]; ok && t.Run(testName, func(t *testing.T) {
					validator(t, res)
				}) {
					validated = append(validated, fmt.Sprintf("%s:%s", key.Name, key.Type))
				}
			}

			expected := []string{}
			for _, key := range maps.Keys(lookup) {
				expected = append(expected, fmt.Sprintf("%s:%s", key.Name, key.Type))
			}

			assert.ElementsMatch(t, expected, validated, "Not all resources were validated")
		},
	})

	integration.ProgramTest(t, &test)
}

func SingleContainerSetup(t *testing.T) integration.ProgramTestOptions {
	const (
		username = "root"
		password = "root"
	)

	node := newNode(t,
		WithSshUsername(username),
		WithSshPassword(password),
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

	opts := []SshServerOption{
		WithSshUsername(username),
		WithSshPassword(password),
	}
	nodes := newCluster(t, map[string][]SshServerOption{
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

func expectOutput(t *testing.T, res apitype.ResourceV3, key string, value interface{}) {
	expectKey(t, res.Outputs, key, value)
}

func newNode(t *testing.T, opts ...SshServerOption) node {
	ctx := context.Background()
	server, err := StartSshServer(ctx, opts...)
	require.NoError(t, err, "failed to start ssh server")

	t.Cleanup(func() {
		err := StopSshServer(ctx, server)
		require.NoError(t, err)
	})

	sshPort, err := server.SshPort(ctx)
	require.NoError(t, err, "failed to get ssh port")

	etcdPort, err := server.EtcdPort(ctx)
	require.NoError(t, err, "failed to get etcd port")

	ip, err := server.Ip(ctx)
	require.NoError(t, err, "failed to retrieve node IP")

	return node{
		Server:   server,
		SshPort:  sshPort,
		EtcdPort: etcdPort,
		Ip:       ip,
	}
}

func newCluster(t *testing.T, config map[string][]SshServerOption) map[string]node {
	ctx := context.Background()
	network, err := network.New(ctx, network.WithCheckDuplicate())
	require.NoError(t, err, "failed to create network")

	t.Cleanup(func() {
		require.NoError(t, network.Remove(ctx))
	})

	nodes := map[string]node{}
	for _, key := range maps.Keys(config) {
		opts := append(config[key], WithNetwork(network.Name))
		nodes[key] = newNode(t, opts...)
	}

	return nodes
}
