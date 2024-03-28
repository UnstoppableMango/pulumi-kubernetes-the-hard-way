//go:build nodejs || all

package examples

import (
	"context"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

// func TestSimpleTs(t *testing.T) {
// 	test := getJSBaseOptions(t).
// 		With(integration.ProgramTestOptions{
// 			Dir:           path.Join(getCwd(t), "simple-ts"),
// 			Quick:         true,
// 			SkipRefresh:   true,
// 			RunUpdateTest: false,
// 		})

// 	integration.ProgramTest(t, &test)
// }

func TestRemoteTs(t *testing.T) {
	const (
		username = "test-user"
		password = "test-password"
		content  = "Some content idk"
	)

	ctx := context.Background()
	server, err := StartSshServer(ctx,
		WithSshUsername(username),
		WithSshPassword(password))
	assert.NoError(t, err)

	port, err := server.Port(ctx)
	assert.NoError(t, err)

	defer StopSshServer(ctx, server) // TODO: Error handling?

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "remote-ts"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
			Config: map[string]string{
				"host":     "localhost",
				"port":     port,
				"user":     username,
				"password": password,
				"content":  content,
				"path":     "/config/remote-file-test",
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				assert.Equal(t, content, stack.Outputs["stdout"])
				assert.Empty(t, stack.Outputs["stderr"])
			},
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@unmango/pulumi-kubernetes-the-hard-way",
		},
	})

	return baseJS
}
