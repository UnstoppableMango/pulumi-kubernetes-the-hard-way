package examples

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func skipIfCi(t *testing.T) {
	if x, ok := os.LookupEnv("CI"); ok && x == "true" {
		t.Skip("skipping resource-intensive test in CI")
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

func expectKey[T comparable, V any](t *testing.T, m map[T]V, key T, value V) {
	actual, ok := m[key]
	assert.Truef(t, ok, "Key `%s` was not set", key)
	assert.Equal(t, value, actual)
}
