//go:build nodejs || all

package examples

import (
	"context"
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

func TestRemoteFileTs(t *testing.T) {
	ctx := context.Background()
	server, err := startSshServer(ctx)
	assert.NoError(t, err)

	defer stopSshServer(ctx, server) // TODO: Error handling?

	// test := getJSBaseOptions(t).
	// 	With(integration.ProgramTestOptions{
	// 		Dir:           path.Join(getCwd(t), "remote-file"),
	// 		Quick:         true,
	// 		SkipRefresh:   true,
	// 		RunUpdateTest: false,
	// 		Config: map[string]string{
	// 			"host":    "localhost",
	// 			"content": "Some test content idk",
	// 			"path":    "/home/testuser/remote-file-test",
	// 		},
	// 	})

	// integration.ProgramTest(t, &test)
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
